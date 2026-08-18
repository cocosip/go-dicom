// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/testutil"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	bufferio "github.com/cocosip/go-dicom/pkg/io/buffer"
)

const dicmPrefix = "DICM"

func addTestSOPUIDs(t *testing.T, ds *dataset.Dataset) {
	t.Helper()
	if err := ds.AddOrUpdate(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("add SOPClassUID: %v", err)
	}
	if err := ds.AddOrUpdate(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.1142.1"})); err != nil {
		t.Fatalf("add SOPInstanceUID: %v", err)
	}
}

// TestWritePreamble tests preamble writing
func TestWritePreamble(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	if err := w.writePreamble(); err != nil {
		t.Fatalf("writePreamble() error = %v", err)
	}

	// Check preamble length (128 + 4 = 132 bytes)
	if buf.Len() != 132 {
		t.Errorf("Preamble length = %d, want 132", buf.Len())
	}

	// Check DICM prefix
	data := buf.Bytes()
	actualPrefix := string(data[128:132])
	if actualPrefix != dicmPrefix {
		t.Errorf("DICM prefix = %q, want %q", actualPrefix, dicmPrefix)
	}
}

func TestWriteUsesDatasetTransferSyntaxAndSyncsFileMeta(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Meta^Sync"}))

	fmi := dataset.New()
	_ = fmi.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
		[]string{transfer.ExplicitVRLittleEndian.UID().String()}))

	var buf bytes.Buffer
	if err := Write(&buf, ds, WithFileMetaInfo(fmi)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	originalTS, ok := fmi.GetString(tag.TransferSyntaxUID)
	if !ok || originalTS != transfer.ExplicitVRLittleEndian.UID().String() {
		t.Fatalf("caller file meta was mutated to %q", originalTS)
	}

	result, err := parser.Parse(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	gotTS, ok := result.FileMetaInformationDataset().GetString(tag.TransferSyntaxUID)
	if !ok {
		t.Fatal("written file meta missing TransferSyntaxUID")
	}
	if gotTS != transfer.ExplicitVRBigEndian.UID().String() {
		t.Fatalf("TransferSyntaxUID = %q, want %q", gotTS, transfer.ExplicitVRBigEndian.UID().String())
	}
	if result.TransferSyntax != transfer.ExplicitVRBigEndian {
		t.Fatalf("parsed transfer syntax = %v, want ExplicitVRBigEndian", result.TransferSyntax)
	}
}

func TestWriteFragmentSequenceStreamsFragments(t *testing.T) {
	fs := element.NewOtherByteFragment(tag.PixelData).FragmentSequence
	frag := &trackingBuffer{data: []byte{0x01, 0x02, 0x03}}
	fs.AddFragment(frag)

	var buf bytes.Buffer
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = &buf
	if err := w.writeFragmentSequence(fs); err != nil {
		t.Fatalf("writeFragmentSequence() error = %v", err)
	}

	if frag.dataCalled {
		t.Fatal("writeFragmentSequence() called Data(); want streaming through WriteTo")
	}
	if !frag.writeToCalled {
		t.Fatal("writeFragmentSequence() did not call WriteTo")
	}
}

func TestWriteAndCloseReturnsCloseError(t *testing.T) {
	closeErr := errors.New("close failed")
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Close^Error"}))

	err := writeAndClose(&closeErrorWriter{closeErr: closeErr}, ds)
	if !errors.Is(err, closeErr) {
		t.Fatalf("writeAndClose() error = %v, want close error %v", err, closeErr)
	}
}

type closeErrorWriter struct {
	bytes.Buffer
	closeErr error
}

func (w *closeErrorWriter) Close() error {
	return w.closeErr
}

type trackingBuffer struct {
	data          []byte
	dataCalled    bool
	writeToCalled bool
}

func (b *trackingBuffer) IsMemory() bool {
	return false
}

func (b *trackingBuffer) Size() uint32 {
	return uint32(len(b.data)) //nolint:gosec // test data is small
}

func (b *trackingBuffer) Data() []byte {
	b.dataCalled = true
	return b.data
}

func (b *trackingBuffer) GetByteRange(offset, count uint32, output []byte) error {
	copy(output[:count], b.data[offset:offset+count])
	return nil
}

func (b *trackingBuffer) WriteTo(w io.Writer) (int64, error) {
	b.writeToCalled = true
	n, err := w.Write(b.data)
	return int64(n), err
}

// TestWriteTag tests tag writing
func TestWriteTag(t *testing.T) {
	t.Run("LittleEndian", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		testTag := tag.New(0x0010, 0x0010)
		if err := w.writeTag(testTag); err != nil {
			t.Fatalf("writeTag() error = %v", err)
		}

		if buf.Len() != 4 {
			t.Errorf("Tag length = %d, want 4", buf.Len())
		}

		// Verify bytes
		data := buf.Bytes()
		group := binary.LittleEndian.Uint16(data[0:2])
		elem := binary.LittleEndian.Uint16(data[2:4])

		if group != 0x0010 {
			t.Errorf("Group = %04X, want 0010", group)
		}
		if elem != 0x0010 {
			t.Errorf("Element = %04X, want 0010", elem)
		}
	})
}

// TestWriteVR tests VR writing
func TestWriteVR(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	if err := w.writeVR(vr.PN); err != nil {
		t.Fatalf("writeVR() error = %v", err)
	}

	if buf.Len() != 2 {
		t.Errorf("VR length = %d, want 2", buf.Len())
	}

	vrCode := buf.String()
	if vrCode != "PN" {
		t.Errorf("VR = %q, want %q", vrCode, "PN")
	}
}

// TestWriteLength tests length writing
func TestWriteLength(t *testing.T) {
	t.Run("16BitLength", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = true

		if err := w.writeLength(vr.PN, 100); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		if buf.Len() != 2 {
			t.Errorf("Length field size = %d, want 2", buf.Len())
		}

		length := binary.LittleEndian.Uint16(buf.Bytes())
		if length != 100 {
			t.Errorf("Length = %d, want 100", length)
		}
	})

	t.Run("32BitLength", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = true

		if err := w.writeLength(vr.SQ, 1000); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		// 2 bytes reserved + 4 bytes length = 6 bytes
		if buf.Len() != 6 {
			t.Errorf("Length field size = %d, want 6", buf.Len())
		}

		data := buf.Bytes()
		reserved := binary.LittleEndian.Uint16(data[0:2])
		length := binary.LittleEndian.Uint32(data[2:6])

		if reserved != 0 {
			t.Errorf("Reserved = %d, want 0", reserved)
		}
		if length != 1000 {
			t.Errorf("Length = %d, want 1000", length)
		}
	})

	t.Run("ImplicitVR", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = false

		if err := w.writeLength(vr.PN, 2000); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		// Implicit VR always uses 32-bit length
		if buf.Len() != 4 {
			t.Errorf("Length field size = %d, want 4", buf.Len())
		}

		length := binary.LittleEndian.Uint32(buf.Bytes())
		if length != 2000 {
			t.Errorf("Length = %d, want 2000", length)
		}
	})
}

// TestWriteElement tests element writing
func TestWriteElement(t *testing.T) {
	t.Run("StringElement", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		elem := element.NewString(tag.PatientName, vr.PN, []string{testPatientNameJohn})
		if err := w.writeElement(elem); err != nil {
			t.Fatalf("writeElement() error = %v", err)
		}

		// Tag (4) + VR (2) + Length (2) + Value (8) = 16 bytes
		if buf.Len() != 16 {
			t.Errorf("Element size = %d, want 16", buf.Len())
		}
	})

	t.Run("UnsignedShortElement", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		elem := element.NewUnsignedShort(tag.Rows, []uint16{512})
		if err := w.writeElement(elem); err != nil {
			t.Fatalf("writeElement() error = %v", err)
		}

		// Tag (4) + VR (2) + Length (2) + Value (2) = 10 bytes
		if buf.Len() != 10 {
			t.Errorf("Element size = %d, want 10", buf.Len())
		}
	})

	t.Run("OddLengthStringElementIsPadded", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		elem := element.NewString(tag.PatientID, vr.LO, []string{"A"})
		if err := w.writeElement(elem); err != nil {
			t.Fatalf("writeElement() error = %v", err)
		}

		data := buf.Bytes()
		if got := binary.LittleEndian.Uint16(data[6:8]); got != 2 {
			t.Fatalf("VL = %d, want 2", got)
		}
		if got := data[8:10]; !bytes.Equal(got, []byte{'A', ' '}) {
			t.Fatalf("value bytes = %v, want %v", got, []byte{'A', ' '})
		}
	})
}

// TestWriteSimpleDataset tests writing a simple dataset
func TestWriteSimpleDataset(t *testing.T) {
	buf := &bytes.Buffer{}

	// Create dataset
	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientNameJohn})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}
	if err := ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

	// Write (using defaults)
	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Check that something was written
	if buf.Len() == 0 {
		t.Error("Nothing was written")
	}

	// Check preamble
	data := buf.Bytes()
	if len(data) < 132 {
		t.Error("Output too short to contain preamble")
	}

	actualPrefix := string(data[128:132])
	if actualPrefix != dicmPrefix {
		t.Errorf("DICM prefix = %q, want %q", actualPrefix, dicmPrefix)
	}
}

func TestWriteNilDataset(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Write(buf, nil); err == nil {
		t.Fatal("Write() should fail for nil dataset")
	}
}

func TestWritePart10RequiresDatasetSOPUIDs(t *testing.T) {
	tests := []struct {
		name string
		ds   *dataset.Dataset
	}{
		{"missing SOP Class UID", func() *dataset.Dataset {
			ds := dataset.New()
			_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))
			return ds
		}()},
		{"missing SOP Instance UID", func() *dataset.Dataset {
			ds := dataset.New()
			_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
			return ds
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(&bytes.Buffer{}, tt.ds); err == nil {
				t.Fatal("Write() succeeded without required Dataset SOP UID")
			}
		})
	}
}

func TestWritePart10RejectsMismatchedMediaStorageSOPUIDs(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	tests := []struct {
		name string
		tag  *tag.Tag
	}{
		{"SOP Class UID", tag.MediaStorageSOPClassUID},
		{"SOP Instance UID", tag.MediaStorageSOPInstanceUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmi := dataset.New()
			_ = fmi.Add(element.NewString(tt.tag, vr.UI, []string{"9.9.9"}))
			if err := Write(&bytes.Buffer{}, ds, WithFileMetaInfo(fmi)); err == nil {
				t.Fatal("Write() succeeded with mismatched File Meta SOP UID")
			}
		})
	}
}

func TestWritePart10AllowsDICOMDIRIdentityInFileMetaOnly(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.FileSetID, vr.CS, []string{"TEST"}))
	fmi := dataset.New()
	_ = fmi.Add(element.NewString(tag.MediaStorageSOPClassUID, vr.UI, []string{uid.MediaStorageDirectoryStorage.UID()}))
	_ = fmi.Add(element.NewString(tag.MediaStorageSOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	if err := Write(&bytes.Buffer{}, ds, WithFileMetaInfo(fmi)); err != nil {
		t.Fatalf("Write() rejected standard DICOMDIR identity placement: %v", err)
	}
}

func TestWriteRejectsPixelDataRepresentationTransferSyntaxMismatch(t *testing.T) {
	newDataset := func(pixelData element.Element) *dataset.Dataset {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
		_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))
		_ = ds.Add(pixelData)
		return ds
	}

	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(bufferio.NewMemory([]byte{0x01, 0x02}))
	tests := []struct {
		name string
		ds   *dataset.Dataset
		ts   *transfer.Syntax
	}{
		{"native pixels with encapsulated syntax", newDataset(element.NewOtherByte(tag.PixelData, []byte{0x01, 0x02})), transfer.JPEG2000Lossless},
		{"fragments with native syntax", newDataset(fragments), transfer.ExplicitVRLittleEndian},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(&bytes.Buffer{}, tt.ds, WithTransferSyntax(tt.ts)); err == nil {
				t.Fatal("Write() succeeded with incompatible Pixel Data representation")
			}
		})
	}
}

// TestWriteWithoutPreamble tests writing without preamble
func TestWriteWithoutPreamble(t *testing.T) {
	buf := &bytes.Buffer{}

	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test"})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

	if err := Write(buf, ds, WithoutPreamble()); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Should not start with preamble
	data := buf.Bytes()
	if len(data) >= 132 {
		dicmCheck := string(data[128:132])
		if dicmCheck == dicmPrefix {
			t.Error("Should not contain DICM prefix when WithoutPreamble is used")
		}
	}
}

// TestWriteSequence tests writing a sequence
func TestWriteSequence(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	// Create a sequence with one item
	seq := dataset.NewSequence(tag.New(0x0008, 0x1140))
	item := dataset.New()
	if err := item.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{testReferencedUID})); err != nil {
		t.Fatalf("Item.Add() error: %v", err)
	}
	seq.AddItem(item)

	if err := w.writeSequence(seq); err != nil {
		t.Fatalf("writeSequence() error = %v", err)
	}

	// Sequence should have been written
	if buf.Len() == 0 {
		t.Error("Nothing was written")
	}

	// Check for sequence delimitation tag (FFFE,E0DD) at the end
	data := buf.Bytes()
	if len(data) < 8 {
		t.Error("Sequence output too short")
	}

	// The last 8 bytes should be the delimitation tag and length
	lastBytes := data[len(data)-8:]
	delimGroup := binary.LittleEndian.Uint16(lastBytes[0:2])
	delimElem := binary.LittleEndian.Uint16(lastBytes[2:4])
	delimLength := binary.LittleEndian.Uint32(lastBytes[4:8])

	if delimGroup != 0xFFFE || delimElem != 0xE0DD {
		t.Errorf("Expected sequence delimitation tag (FFFE,E0DD), got (%04X,%04X)",
			delimGroup, delimElem)
	}
	if delimLength != 0 {
		t.Errorf("Delimitation length = %d, want 0", delimLength)
	}
}

// TestRoundTrip tests write then read
func TestRoundTrip(t *testing.T) {
	buf := &bytes.Buffer{}

	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientNameJohn})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}
	if err := ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}
	if err := ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}
	if err := ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	t.Logf("Wrote %d bytes", buf.Len())

	// Verify basic structure
	data := buf.Bytes()

	// Should have preamble
	if len(data) < 132 {
		t.Fatal("Output too short")
	}

	// Should have DICM
	if string(data[128:132]) != dicmPrefix {
		t.Error("Missing DICM prefix")
	}
}

// Benchmark tests for Writer

func BenchmarkWriteSmallDataset(b *testing.B) {
	// Create a small dataset
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientNameJohn}))
	_ = ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))
	_ = ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteMediumDataset(b *testing.B) {
	// Create a medium dataset with 50 elements
	ds := dataset.New()
	for i := 0; i < 50; i++ {
		// Use safe conversion helper to satisfy gosec
		elem := testutil.SafeUint16FromInt(i)
		t := tag.New(0x0010, elem)
		_ = ds.Add(element.NewString(t, vr.LO, []string{"TestValue"}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteLargeDataset(b *testing.B) {
	// Create a large dataset with 200 elements
	ds := dataset.New()
	for i := 0; i < 200; i++ {
		// Use safe conversion helper
		elem := testutil.SafeUint16FromInt(i % 256)
		t := tag.New(0x0010, elem)
		_ = ds.Add(element.NewString(t, vr.LO, []string{"TestValue"}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteTag(b *testing.B) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf
	testTag := tag.New(0x0010, 0x0010)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = w.writeTag(testTag)
	}
}

func BenchmarkWriteElement(b *testing.B) {
	elem := element.NewString(tag.PatientName, vr.PN, []string{testPatientNameJohn})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		_ = w.writeElement(elem)
	}
}

func BenchmarkWritePreamble(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		_ = w.writePreamble()
	}
}
