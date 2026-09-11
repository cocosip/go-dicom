// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
	pixelData := openTestPixelData(t, []uint32{0, 10}, []uint64{0, 20}, []uint64{20, 20})

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
	pixelData := openTestPixelData(t, nil, []uint64{0, 20}, []uint64{20, 20})

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
	lengths := element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encodeBigEndianUint64s(10))
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
