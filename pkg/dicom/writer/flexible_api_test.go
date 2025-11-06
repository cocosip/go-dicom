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
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestWriteWithDefaults tests writing with default options
func TestWriteWithDefaults(t *testing.T) {
	buf := &bytes.Buffer{}

	// Create a simple dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))

	// Write with defaults (no options)
	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Parse it back
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

	// Verify Transfer Syntax UID was auto-generated
	tsUID, exists := result.FileMetaInformation.GetString(tag.TransferSyntaxUID)
	if !exists {
		t.Fatal("TransferSyntaxUID not found in auto-generated FileMetaInformation")
	}
	if tsUID != "1.2.840.10008.1.2.1" { // Explicit VR Little Endian
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, "1.2.840.10008.1.2.1")
	}
}

// TestWriteWithFileMetaInfo tests writing with custom file meta info
func TestWriteWithFileMetaInfo(t *testing.T) {
	buf := &bytes.Buffer{}

	// Create custom file meta info (without TransferSyntaxUID)
	fileMetaInfo := dataset.New()
	// Note: TransferSyntaxUID will be added automatically

	// Create a simple dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))

	// Write with custom file meta info
	if err := Write(buf, ds, WithFileMetaInfo(fileMetaInfo)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Parse it back
	result, err := parser.Parse(buf)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify TransferSyntaxUID was added automatically
	tsUID, exists := result.FileMetaInformation.GetString(tag.TransferSyntaxUID)
	if !exists {
		t.Fatal("TransferSyntaxUID not found - should have been added automatically")
	}
	if tsUID != "1.2.840.10008.1.2.1" {
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, "1.2.840.10008.1.2.1")
	}
}

// TestWriteWithTransferSyntax tests writing with a specific transfer syntax
func TestWriteWithTransferSyntax(t *testing.T) {
	buf := &bytes.Buffer{}

	// Create a dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))

	// Write with explicit transfer syntax option
	ts := transfer.ExplicitVRLittleEndian
	if err := Write(buf, ds, WithTransferSyntax(ts)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	t.Logf("Wrote %d bytes", buf.Len())

	// Parse it back
	result, err := parser.Parse(buf)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify all elements
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if !exists || patientName != "Doe^John" {
		t.Errorf("PatientName = %q, exists = %v, want %q", patientName, exists, "Doe^John")
	}

	patientID, exists := result.Dataset.GetString(tag.PatientID)
	if !exists || patientID != "12345" {
		t.Errorf("PatientID = %q, exists = %v, want %q", patientID, exists, "12345")
	}

	rows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil || rows != 512 {
		t.Errorf("Rows = %d, err = %v, want 512", rows, err)
	}
}

// TestWriteWithoutPreamble tests writing without preamble
func TestWriteWithoutPreambleOption(t *testing.T) {
	buf := &bytes.Buffer{}

	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test"}))

	// Write without preamble
	if err := Write(buf, ds, WithoutPreamble()); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Should not start with preamble
	data := buf.Bytes()
	if len(data) >= 132 {
		dicmCheck := string(data[128:132])
		if dicmCheck == "DICM" {
			t.Error("Should not contain DICM prefix when WithoutPreamble is used")
		}
	}
}
