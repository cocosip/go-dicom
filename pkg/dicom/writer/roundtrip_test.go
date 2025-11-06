// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestWriteThenRead tests writing a DICOM file and then reading it back
func TestWriteThenRead(t *testing.T) {
	// Create dataset with various element types
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))
	ds.Add(element.NewFloat(tag.ImagePositionPatient, []float32{1.5, 2.5, 3.5}))

	// Write to buffer (using defaults)
	buf := &bytes.Buffer{}
	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	t.Logf("Wrote %d bytes", buf.Len())

	// Read it back
	result, err := parser.Parse(buf)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify file meta information
	tsUID, exists := result.FileMetaInformation.GetString(tag.TransferSyntaxUID)
	if !exists {
		t.Fatal("TransferSyntaxUID not found in FileMetaInformation")
	}
	if tsUID != "1.2.840.10008.1.2.1" {
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, "1.2.840.10008.1.2.1")
	}

	// Verify main dataset
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if !exists {
		t.Fatal("PatientName not found")
	}
	if patientName != "Doe^John" {
		t.Errorf("PatientName = %q, want %q", patientName, "Doe^John")
	}

	patientID, exists := result.Dataset.GetString(tag.PatientID)
	if !exists {
		t.Fatal("PatientID not found")
	}
	if patientID != "12345" {
		t.Errorf("PatientID = %q, want %q", patientID, "12345")
	}

	rows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil {
		t.Fatalf("GetUInt16(Rows) error = %v", err)
	}
	if rows != 512 {
		t.Errorf("Rows = %d, want 512", rows)
	}

	columns, err := result.Dataset.GetUInt16(tag.Columns, 0)
	if err != nil {
		t.Fatalf("GetUInt16(Columns) error = %v", err)
	}
	if columns != 512 {
		t.Errorf("Columns = %d, want 512", columns)
	}

	// Verify float array
	position, err := result.Dataset.GetFloat32(tag.ImagePositionPatient, 0)
	if err != nil {
		t.Fatalf("GetFloat32(ImagePositionPatient, 0) error = %v", err)
	}
	if position != 1.5 {
		t.Errorf("ImagePositionPatient[0] = %f, want 1.5", position)
	}
}

// TestWriteThenReadWithSequence tests writing and reading a dataset with a sequence
func TestWriteThenReadWithSequence(t *testing.T) {
	// Create dataset with a sequence
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))

	// Create a Referenced Image Sequence
	seq := dataset.NewSequence(tag.New(0x0008, 0x1140))
	item1 := dataset.New()
	item1.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{"1.2.3.4"}))
	item1.Add(element.NewString(tag.New(0x0008, 0x1150), vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	seq.AddItem(item1)

	item2 := dataset.New()
	item2.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{"1.2.3.5"}))
	seq.AddItem(item2)

	ds.Add(seq)

	// Write to buffer (using defaults)
	buf := &bytes.Buffer{}
	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	t.Logf("Wrote %d bytes", buf.Len())

	// Read it back
	result, err := parser.Parse(buf)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify patient name
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if !exists {
		t.Fatal("PatientName not found")
	}
	if patientName != "Test^Patient" {
		t.Errorf("PatientName = %q, want %q", patientName, "Test^Patient")
	}

	// Verify sequence
	seqElem, exists := result.Dataset.Get(tag.New(0x0008, 0x1140))
	if !exists {
		t.Fatal("Referenced Image Sequence not found")
	}

	readSeq, ok := seqElem.(*dataset.Sequence)
	if !ok {
		t.Fatal("Element is not a Sequence")
	}

	if readSeq.Count() != 2 {
		t.Errorf("Sequence count = %d, want 2", readSeq.Count())
	}

	// Verify first item
	readItem1 := readSeq.GetItem(0)
	if readItem1 == nil {
		t.Fatal("First item is nil")
	}

	sopInstanceUID, exists := readItem1.GetString(tag.New(0x0008, 0x1155))
	if !exists {
		t.Fatal("SOP Instance UID not found in first item")
	}
	if sopInstanceUID != "1.2.3.4" {
		t.Errorf("SOP Instance UID = %q, want %q", sopInstanceUID, "1.2.3.4")
	}

	// Verify second item
	readItem2 := readSeq.GetItem(1)
	if readItem2 == nil {
		t.Fatal("Second item is nil")
	}

	sopInstanceUID2, exists := readItem2.GetString(tag.New(0x0008, 0x1155))
	if !exists {
		t.Fatal("SOP Instance UID not found in second item")
	}
	if sopInstanceUID2 != "1.2.3.5" {
		t.Errorf("SOP Instance UID = %q, want %q", sopInstanceUID2, "1.2.3.5")
	}
}
