// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

// TestWriteThenRead tests writing a DICOM file and then reading it back
func TestWriteThenRead(t *testing.T) {
	// Create dataset with various element types
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
	if err := ds.Add(element.NewFloat(tag.RecommendedDisplayFrameRateInFloat, []float32{1.5})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

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
	tsUID, exists := result.FileMetaInformation.TransferSyntaxUID()
	if !exists {
		t.Fatal("TransferSyntaxUID not found in FileMetaInformation")
	}
	if tsUID != testExplicitVRLittleLE {
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, testExplicitVRLittleLE)
	}

	// Verify main dataset
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if !exists {
		t.Fatal("PatientName not found")
	}
	if patientName != testPatientNameJohn {
		t.Errorf("PatientName = %q, want %q", patientName, testPatientNameJohn)
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
	position, err := result.Dataset.GetFloat32(tag.RecommendedDisplayFrameRateInFloat, 0)
	if err != nil {
		t.Fatalf("GetFloat32(RecommendedDisplayFrameRateInFloat, 0) error = %v", err)
	}
	if position != 1.5 {
		t.Errorf("RecommendedDisplayFrameRateInFloat = %f, want 1.5", position)
	}
}

func TestWriteThenReadChineseTextWithDeclaredCharacterSet(t *testing.T) {
	for _, characterSet := range []string{"ISO_IR 192", "GB18030"} {
		t.Run(characterSet, func(t *testing.T) {
			ds := dataset.New()
			addTestSOPUIDs(t, ds)
			if err := ds.SetSpecificCharacterSet(characterSet); err != nil {
				t.Fatalf("SetSpecificCharacterSet() error = %v", err)
			}
			values := []struct {
				tag   *tag.Tag
				value string
			}{
				{tag.PatientName, "张三"},
				{tag.StudyDescription, "头部增强检查"},
				{tag.AdditionalPatientHistory, "患者头痛三天"},
				{tag.UnformattedTextValue, "检查过程中患者配合良好"},
			}
			for _, value := range values {
				if err := ds.AddValue(value.tag, value.value); err != nil {
					t.Fatalf("AddValue(%s) error = %v", value.tag, err)
				}
			}

			var encoded bytes.Buffer
			if err := Write(&encoded, ds); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			result, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			if got := result.Dataset.SpecificCharacterSet(); len(got) != 1 || got[0] != characterSet {
				t.Fatalf("SpecificCharacterSet() = %v, want [%s]", got, characterSet)
			}
			for _, value := range values {
				got, ok := result.Dataset.GetString(value.tag)
				if !ok || got != value.value {
					t.Fatalf("GetString(%s) = %q, %v; want %q, true", value.tag, got, ok, value.value)
				}
			}
		})
	}
}

func TestParsedSequenceItemInheritsSpecificCharacterSetForNewValues(t *testing.T) {
	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.SetSpecificCharacterSet("ISO_IR 192"); err != nil {
		t.Fatalf("SetSpecificCharacterSet() error = %v", err)
	}
	item := dataset.New()
	if err := item.Add(element.NewStringWithEncoding(tag.PatientName, vr.PN, []string{"张三"}, unicode.UTF8)); err != nil {
		t.Fatalf("Add(PatientName) error = %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.RequestAttributesSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("Add(RequestAttributesSequence) error = %v", err)
	}

	var encoded bytes.Buffer
	if err := Write(&encoded, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	result, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	sequence, err := result.Dataset.GetSequence(tag.RequestAttributesSequence)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}
	parsedItem := sequence.GetItem(0)
	if err := parsedItem.AddValue(tag.StudyDescription, "中文描述"); err != nil {
		t.Fatalf("sequence item AddValue() error = %v", err)
	}
	if got, ok := parsedItem.GetString(tag.StudyDescription); !ok || got != "中文描述" {
		t.Fatalf("GetString(StudyDescription) = %q, %v; want 中文描述, true", got, ok)
	}
}

func TestWriteThenReadPreservesTextBytesWhenDeclaredCharacterSetIsWrong(t *testing.T) {
	raw, err := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("张三"))
	if err != nil {
		t.Fatalf("encode GB18030 fixture: %v", err)
	}
	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.SetSpecificCharacterSet("ISO_IR 192"); err != nil {
		t.Fatalf("SetSpecificCharacterSet() error = %v", err)
	}
	if err := ds.Add(element.NewStringFromBufferWithEncodings(
		tag.PatientName,
		vr.PN,
		buffer.NewMemory(raw),
		[]encoding.Encoding{unicode.UTF8},
	)); err != nil {
		t.Fatalf("Add(PatientName) error = %v", err)
	}

	var encoded bytes.Buffer
	if err := Write(&encoded, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	result, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	gotRaw, err := result.Dataset.GetBytes(tag.PatientName)
	if err != nil {
		t.Fatalf("GetBytes(PatientName) error = %v", err)
	}
	if !bytes.Equal(gotRaw, raw) {
		t.Fatalf("round-tripped PatientName bytes = %x, want %x", gotRaw, raw)
	}
	decoded, err := charset.DecodeString(gotRaw, charset.GetEncodings([]string{"GB18030"}))
	if err != nil {
		t.Fatalf("DecodeString(GB18030) error = %v", err)
	}
	if decoded != "张三" {
		t.Fatalf("DecodeString(GB18030) = %q, want 张三", decoded)
	}
}

func TestWriteThenReadDeflatedExplicitVRLittleEndian(t *testing.T) {
	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Deflated^Patient"})); err != nil {
		t.Fatalf("Add(PatientName) error = %v", err)
	}
	if err := ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"DEF123"})); err != nil {
		t.Fatalf("Add(PatientID) error = %v", err)
	}

	buf := &bytes.Buffer{}
	if err := Write(buf, ds, WithTransferSyntax(transfer.DeflatedExplicitVRLittleEndian)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	result, err := parser.Parse(buf)
	if err != nil {
		t.Fatalf("Parse(deflated output) error = %v", err)
	}

	tsUID, exists := result.FileMetaInformation.TransferSyntaxUID()
	if !exists {
		t.Fatal("TransferSyntaxUID not found")
	}
	if tsUID != transfer.DeflatedExplicitVRLittleEndian.UID().String() {
		t.Fatalf("TransferSyntaxUID = %q, want %q", tsUID, transfer.DeflatedExplicitVRLittleEndian.UID().String())
	}

	if got, exists := result.Dataset.GetString(tag.PatientName); !exists || got != "Deflated^Patient" {
		t.Fatalf("PatientName = %q, exists=%v; want Deflated^Patient", got, exists)
	}
}

// TestWriteThenReadWithSequence tests writing and reading a dataset with a sequence
func TestWriteThenReadWithSequence(t *testing.T) {
	// Create dataset with a sequence
	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName})); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

	// Create a Referenced Image Sequence
	seq := dataset.NewSequence(tag.New(0x0008, 0x1140))
	item1 := dataset.New()
	if err := item1.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{testReferencedUID})); err != nil {
		t.Fatalf("Item.Add() error: %v", err)
	}
	if err := item1.Add(element.NewString(tag.New(0x0008, 0x1150), vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("Item.Add() error: %v", err)
	}
	seq.AddItem(item1)

	item2 := dataset.New()
	if err := item2.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{"1.2.3.5"})); err != nil {
		t.Fatalf("Item.Add() error: %v", err)
	}
	seq.AddItem(item2)

	if err := ds.Add(seq); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

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
	if patientName != testPatientName {
		t.Errorf("PatientName = %q, want %q", patientName, testPatientName)
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
	if sopInstanceUID != testReferencedUID {
		t.Errorf("SOP Instance UID = %q, want %q", sopInstanceUID, testReferencedUID)
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
