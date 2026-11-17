// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package writer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

// The interop fixture is a small part-10 file that fo-dicom v6.0.0-alpha1 was
// verified to decode and re-encode to the same dataset bytes (see
// test-data/interop/README.md). The writer must keep producing those exact
// bytes, so a change that alters padding, element order, sequence lengths or the
// file meta format fails here instead of only in a future interop run.
const (
	interopGoldenPath      = "../../../test-data/interop/ct_small_explicit_le.dcm"
	interopFoDicomPath     = "../../../test-data/interop/ct_small_explicit_le_fo-dicom.dcm"
	interopImplementation  = "1.2.826.0.1.3680043.10.999.1"
	interopVersionName     = "GODICOM-INTEROP"
	interopUpdateGoldenVar = "GO_DICOM_UPDATE_INTEROP_GOLDEN"
	interopPatientName     = "Doe^John"
	interopCTImageStorage  = "1.2.840.10008.5.1.4.1.1.2"
)

// interopWriteOptions returns the write options the golden fixture uses.
func interopWriteOptions() []writer.WriteOption {
	return []writer.WriteOption{
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
		writer.WithImplementationClassUID(interopImplementation),
		writer.WithImplementationVersionName(interopVersionName),
		// Defined lengths are the canonical form fo-dicom produces with
		// DicomWriteOptions.Default.ExplicitLengthSequences(true).
		writer.WithExplicitLengthSequences(true),
		writer.WithExplicitLengthSequenceItems(),
	}
}

func TestWriteMatchesInteropGolden(t *testing.T) {
	ds := buildInteropDataset(t)

	var buf bytes.Buffer
	if err := writer.Write(&buf,
		ds,
		interopWriteOptions()...,
	); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	if os.Getenv(interopUpdateGoldenVar) != "" {
		if err := os.WriteFile(interopGoldenPath, buf.Bytes(), 0o644); err != nil {
			t.Fatalf("write golden fixture: %v", err)
		}
		t.Logf("regenerated %s (%d bytes)", interopGoldenPath, buf.Len())
	}

	golden, err := os.ReadFile(interopGoldenPath)
	if err != nil {
		t.Fatalf("read golden fixture: %v", err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		t.Fatalf("encoded file differs from the interop golden fixture (%d vs %d bytes); "+
			"run the test with %s=1 only after re-verifying the new encoding with fo-dicom",
			buf.Len(), len(golden), interopUpdateGoldenVar)
	}
}

func TestInteropGoldenRoundTrip(t *testing.T) {
	file, err := parser.ParseFile(interopGoldenPath)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}

	if got := file.TransferSyntax.UID().UID(); got != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Fatalf("TransferSyntaxUID = %q, want Explicit VR Little Endian", got)
	}

	ds := file.Dataset
	if got := ds.TryGetString(tag.PatientName); got != interopPatientName {
		t.Fatalf("PatientName = %q, want Doe^John", got)
	}
	if got := ds.TryGetString(tag.PatientID); got != "P12345" {
		t.Fatalf("PatientID = %q, want P12345", got)
	}
	if got := ds.TryGetString(tag.SOPInstanceUID); got != "1.2.826.0.1.3680043.10.999.1.2.3.4.5" {
		t.Fatalf("SOPInstanceUID = %q", got)
	}
	rows, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil || rows != 2 {
		t.Fatalf("Rows = %d, %v; want 2", rows, err)
	}
	pixels, err := ds.GetBytes(tag.PixelData)
	if err != nil {
		t.Fatalf("GetBytes(PixelData) error = %v", err)
	}
	if len(pixels) != 8 {
		t.Fatalf("PixelData length = %d, want 8", len(pixels))
	}
	sequence, err := ds.GetSequence(tag.ReferencedImageSequence)
	if err != nil {
		t.Fatalf("GetSequence(ReferencedImageSequence) error = %v", err)
	}
	if sequence.Count() != 1 {
		t.Fatalf("ReferencedImageSequence items = %d, want 1", sequence.Count())
	}
	item := sequence.GetItem(0)
	if got := item.TryGetString(tag.ReferencedSOPClassUID); got != interopCTImageStorage {
		t.Fatalf("ReferencedSOPClassUID = %q", got)
	}

	// The fixture must also re-encode to itself, so parsing and writing agree.
	var buf bytes.Buffer
	if err := writer.Write(&buf,
		ds,
		interopWriteOptions()...,
	); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	golden, err := os.ReadFile(interopGoldenPath)
	if err != nil {
		t.Fatalf("read golden fixture: %v", err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		t.Fatal("parsed fixture re-encodes to different bytes")
	}
}

// TestParseFoDicomEncoding covers the other direction: a file written by
// fo-dicom (with its own Implementation Class UID/Version Name and file-meta
// group length) parses here, and re-encoding it reproduces the go-dicom golden.
func TestParseFoDicomEncoding(t *testing.T) {
	file, err := parser.ParseFile(interopFoDicomPath)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if got := file.TransferSyntax.UID().UID(); got != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Fatalf("TransferSyntaxUID = %q, want Explicit VR Little Endian", got)
	}
	if got := file.Dataset.TryGetString(tag.PatientName); got != interopPatientName {
		t.Fatalf("PatientName = %q, want Doe^John", got)
	}
	pixels, err := file.Dataset.GetBytes(tag.PixelData)
	if err != nil {
		t.Fatalf("GetBytes(PixelData) error = %v", err)
	}
	want := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	if !bytes.Equal(pixels[:len(want)], want) {
		t.Fatalf("PixelData = %x, want %x", pixels, want)
	}

	var buf bytes.Buffer
	if err := writer.Write(&buf, file.Dataset, interopWriteOptions()...); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	golden, err := os.ReadFile(interopGoldenPath)
	if err != nil {
		t.Fatalf("read golden fixture: %v", err)
	}
	if !bytes.Equal(buf.Bytes(), golden) {
		t.Fatal("re-encoding the fo-dicom file does not reproduce the golden bytes")
	}
}

// buildInteropDataset builds the fixture deterministically: strings that need// padding, a multi-valued UI/IS element, 16-bit pixel data, and a defined-length
// sequence item. Keeping it small makes byte-level differences readable.
func buildInteropDataset(t *testing.T) *dataset.Dataset {
	t.Helper()

	ds := dataset.New()
	add := func(elem element.Element) {
		t.Helper()
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	add(element.NewString(tag.SOPClassUID, vr.UI, []string{interopCTImageStorage}))
	add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.999.1.2.3.4.5"}))
	add(element.NewString(tag.PatientName, vr.PN, []string{interopPatientName}))
	add(element.NewString(tag.PatientID, vr.LO, []string{"P12345"}))
	add(element.NewString(tag.PatientBirthDate, vr.DA, []string{"19700101"}))
	add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.999.1.2.3"}))
	add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.999.1.2.3.4"}))
	add(element.NewString(tag.StudyDate, vr.DA, []string{"20260101"}))
	add(element.NewString(tag.StudyTime, vr.TM, []string{"120000"}))
	add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))
	add(element.NewString(tag.SeriesNumber, vr.IS, []string{"1"}))
	add(element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"}))
	add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))
	add(element.NewUnsignedShort(tag.Rows, []uint16{2}))
	add(element.NewUnsignedShort(tag.Columns, []uint16{2}))
	add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
	add(element.NewUnsignedShort(tag.BitsStored, []uint16{12}))
	add(element.NewUnsignedShort(tag.HighBit, []uint16{11}))
	add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))

	reference := dataset.New()
	if err := reference.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{interopCTImageStorage})); err != nil {
		t.Fatalf("add referenced SOP Class UID: %v", err)
	}
	if err := reference.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.999.1.2.3.4.9"})); err != nil {
		t.Fatalf("add referenced SOP Instance UID: %v", err)
	}
	add(dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{reference}))

	add(element.NewOtherWord(tag.PixelData, []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
	}))

	return ds
}
