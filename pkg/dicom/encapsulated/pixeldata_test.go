// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package encapsulated

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestPixelDataRejectsAmbiguousEmptyOffsetTable(t *testing.T) {
	pixelData := openTestPixelData(t, nil, nil, nil)

	_, err := pixelData.Frames(2)
	if err == nil || !strings.Contains(err.Error(), "cannot determine frame boundaries") {
		t.Fatalf("Frames() error = %v, want indeterminate-boundaries error", err)
	}
}

func TestPixelDataUsesExtendedOffsetTableBeforeBasicOffsetTable(t *testing.T) {
	pixelData := openTestPixelData(t, []uint32{0, 10}, []uint64{0, 20}, []uint64{4, 4})

	frames, err := pixelData.Frames(2)
	if err != nil {
		t.Fatalf("Frames() error = %v", err)
	}
	want := [][]byte{[]byte("AABB"), []byte("CCDD")}
	if len(frames) != len(want) {
		t.Fatalf("Frames() returned %d frames, want %d", len(frames), len(want))
	}
	for i := range want {
		if !bytes.Equal(frames[i], want[i]) {
			t.Fatalf("Frames()[%d] = %q, want %q", i, frames[i], want[i])
		}
	}
}

func TestPixelDataReadsSelectedExtendedOffsetRange(t *testing.T) {
	pixelData := openTestPixelData(t, nil, []uint64{0, 20}, []uint64{4, 4})

	got, err := pixelData.Frame(2, 1)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}
	if !bytes.Equal(got, []byte("CCDD")) {
		t.Fatalf("Frame() = %q, want %q", got, []byte("CCDD"))
	}
}

func TestPixelDataUsesEncodedFragmentItemOffsets(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable([]uint32{0, 24})
	fragments.AddFragment(buffer.NewMemory([]byte{0xAA, 0xBB, 0xCC}))
	fragments.AddFragment(buffer.NewMemory([]byte{0x11, 0x22, 0x33, 0x44}))
	fragments.AddFragment(buffer.NewMemory([]byte{0x55, 0x66, 0x77}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	frames, err := pixelData.Frames(2)
	if err != nil {
		t.Fatalf("Frames() error = %v", err)
	}
	want := [][]byte{
		{0xAA, 0xBB, 0xCC, 0x11, 0x22, 0x33, 0x44},
		{0x55, 0x66, 0x77},
	}
	if !reflect.DeepEqual(frames, want) {
		t.Fatalf("Frames() = %v, want %v", frames, want)
	}
}

func TestPixelDataConcatenatesSingleFrameWithoutOffsetTable(t *testing.T) {
	pixelData := openTestPixelData(t, nil, nil, nil)

	frames, err := pixelData.Frames(1)
	if err != nil {
		t.Fatalf("Frames() error = %v", err)
	}
	if want := [][]byte{[]byte("AABBCCDD")}; !reflect.DeepEqual(frames, want) {
		t.Fatalf("Frames() = %q, want %q", frames, want)
	}
}

func TestOpenDatasetUsesElementByteOrderForExtendedTable(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	offsets := element.NewOtherVeryLong(tag.ExtendedOffsetTable, encodeBigEndianUint64s(0))
	lengths := element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encodeBigEndianUint64s(2))
	element.SetByteOrder(offsets, binary.BigEndian)
	element.SetByteOrder(lengths, binary.BigEndian)
	if err := ds.Add(offsets); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(lengths); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte("AA")))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	got, err := pixelData.Frame(1, 0)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}
	if !reflect.DeepEqual(got, []byte("AA")) {
		t.Fatalf("Frame() = %q, want AA", got)
	}
}

func TestPixelDataExtendedOffsetLengthExcludesPadding(t *testing.T) {
	// One frame held in a single odd-length fragment. The Extended Offset Table
	// Length is the unpadded frame length (3), not the padded item length (4).
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, encodeUint64s(0))); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encodeUint64s(3))); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0xAA, 0xBB, 0xCC}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	frames, err := pixelData.Frames(1)
	if err != nil {
		t.Fatalf("Frames() error = %v (odd-length frame length must not be padded)", err)
	}
	if want := [][]byte{{0xAA, 0xBB, 0xCC}}; !reflect.DeepEqual(frames, want) {
		t.Fatalf("Frames() = %v, want %v", frames, want)
	}
}

func TestPixelDataExtendedOffsetTrimsStoredFragmentPad(t *testing.T) {
	// A padded fragment item stores 4 bytes for a 3-byte frame; the Extended
	// Offset Table Length records the unpadded 3 bytes (PS3.3 Table C.7-11a).
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, encodeUint64s(0))); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encodeUint64s(3))); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0xAA, 0xBB, 0xCC, 0x00}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	frames, err := pixelData.Frames(1)
	if err != nil {
		t.Fatalf("Frames() error = %v (stored pad byte must be trimmed)", err)
	}
	if want := [][]byte{{0xAA, 0xBB, 0xCC}}; !reflect.DeepEqual(frames, want) {
		t.Fatalf("Frames() = %v, want %v", frames, want)
	}
}

func openTestPixelData(t *testing.T, basic []uint32, extended, lengths []uint64) *PixelData {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if extended != nil {
		if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, encodeUint64s(extended...))); err != nil {
			t.Fatal(err)
		}
	}
	if lengths != nil {
		if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encodeUint64s(lengths...))); err != nil {
			t.Fatal(err)
		}
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable(basic)
	for _, data := range [][]byte{[]byte("AA"), []byte("BB"), []byte("CC"), []byte("DD")} {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}
	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	return pixelData
}

func encodeUint64s(values ...uint64) []byte {
	data := make([]byte, len(values)*8)
	for i, value := range values {
		binary.LittleEndian.PutUint64(data[i*8:], value)
	}
	return data
}

func encodeBigEndianUint64s(values ...uint64) []byte {
	data := make([]byte, len(values)*8)
	for i, value := range values {
		binary.BigEndian.PutUint64(data[i*8:], value)
	}
	return data
}

// TestPixelDataSplitsFramesAtEndOfStreamMarker verifies PS3.5 A.4: an empty
// Basic Offset Table is allowed, so frames that span several fragments must be
// recovered from the codestream's end-of-stream marker.
func TestPixelDataSplitsFramesAtEndOfStreamMarker(t *testing.T) {
	first := []byte{0xff, 0xd8, 0x01, 0x02, 0xff, 0xd9}
	second := []byte{0xff, 0xd8, 0x03, 0x04, 0xff, 0xd9}

	pixelData := openTestPixelDataWithFragments(t, transfer.JPEGBaseline8Bit, nil,
		first[:2], first[2:], second[:5], second[5:])

	frames, err := pixelData.Frames(2)
	if err != nil {
		t.Fatalf("Frames() error = %v", err)
	}
	if len(frames) != 2 {
		t.Fatalf("Frames() returned %d frames, want 2", len(frames))
	}
	for index, want := range [][]byte{first, second} {
		if !bytes.Equal(frames[index], want) {
			t.Fatalf("Frames()[%d] = %x, want %x", index, frames[index], want)
		}
	}
}

// TestPixelDataWithoutOffsetsStillRejectsNonCodestreamSyntax keeps the
// indeterminate-boundaries error for syntaxes whose frames have no common end
// marker.
func TestPixelDataWithoutOffsetsStillRejectsNonCodestreamSyntax(t *testing.T) {
	pixelData := openTestPixelDataWithFragments(t, transfer.RLELossless, nil,
		[]byte{1, 2}, []byte{3, 4}, []byte{5, 6}, []byte{7, 8})

	if _, err := pixelData.Frames(2); err == nil ||
		!strings.Contains(err.Error(), "cannot determine frame boundaries") {
		t.Fatalf("Frames() error = %v, want indeterminate-boundaries error", err)
	}
}

// TestPixelDataToleratesZeroLengthFragments verifies that an empty fragment
// item inside a frame contributes no bytes instead of failing the whole read.
// The parser keeps such items and the writer skips them, so the frame reader
// must tolerate them too (PS3.5 A.4 permits an empty Fragment Item Value).
func TestPixelDataToleratesZeroLengthFragments(t *testing.T) {
	pixelData := openTestPixelDataWithFragments(t, transfer.JPEGBaseline8Bit, []uint32{0},
		[]byte{0xff, 0xd8, 0x01, 0x02}, []byte{}, []byte{0xff, 0xd9})

	frames, err := pixelData.Frames(1)
	if err != nil {
		t.Fatalf("Frames() error = %v", err)
	}
	if len(frames) != 1 {
		t.Fatalf("Frames() returned %d frames, want 1", len(frames))
	}
	want := []byte{0xff, 0xd8, 0x01, 0x02, 0xff, 0xd9}
	if !bytes.Equal(frames[0], want) {
		t.Fatalf("Frames()[0] = %x, want %x", frames[0], want)
	}
}

func openTestPixelDataWithFragments(
	t *testing.T, syntax *transfer.Syntax, basic []uint32, fragmentData ...[]byte,
) *PixelData {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(syntax)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable(basic)
	for _, data := range fragmentData {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}
	pixelData, err := OpenDataset(ds)
	if err != nil {
		t.Fatalf("OpenDataset() error = %v", err)
	}
	return pixelData
}

func TestUsesFrameEndMarkerIsLimitedToCodestreams(t *testing.T) {
	tests := map[string]bool{
		"1.2.840.10008.1.2.4.50":  true,  // JPEG Baseline
		"1.2.840.10008.1.2.4.70":  true,  // JPEG Lossless SV1
		"1.2.840.10008.1.2.4.80":  true,  // JPEG-LS Lossless
		"1.2.840.10008.1.2.4.90":  true,  // JPEG 2000 Lossless
		"1.2.840.10008.1.2.4.203": true,  // HTJ2K
		"1.2.840.10008.1.2.4.100": false, // MPEG2 MP@ML
		"1.2.840.10008.1.2.4.110": false, // JPEG XL Lossless
		"1.2.840.10008.1.2.4.94":  false, // JPIP Referenced
		"1.2.840.10008.1.2.4.99":  false, // Deflated Explicit VR Little Endian
	}
	for uidString, want := range tests {
		syntax, err := transfer.Parse(uidString)
		if err != nil {
			t.Fatalf("transfer.Parse(%q) error = %v", uidString, err)
		}
		if got := usesFrameEndMarker(syntax); got != want {
			t.Errorf("usesFrameEndMarker(%s) = %v, want %v", uidString, got, want)
		}
	}
	if usesFrameEndMarker(nil) {
		t.Error("usesFrameEndMarker(nil) = true, want false")
	}
}
