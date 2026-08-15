// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding/unicode"
)

func TestDecodeDIMSEMessage_PreservesSequenceItems(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"})); err != nil {
		t.Fatalf("failed to build item dataset: %v", err)
	}

	seq := dataset.NewSequenceWithItems(tag.ReferencedStudySequence, []*dataset.Dataset{item})
	ds := dataset.New()
	if err := ds.Add(seq); err != nil {
		t.Fatalf("failed to build test dataset: %v", err)
	}

	var raw bytes.Buffer
	if err := writer.Write(&raw, ds,
		writer.WithoutPreamble(),
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		t.Fatalf("failed to encode test dataset: %v", err)
	}

	_, decoded, err := DecodeDIMSEMessage(nil, raw.Bytes(), transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}

	el, exists := decoded.Get(tag.ReferencedStudySequence)
	if !exists {
		t.Fatal("decoded dataset missing ReferencedStudySequence")
	}

	decodedSeq, ok := el.(*dataset.Sequence)
	if !ok {
		t.Fatalf("decoded element type = %T, want *dataset.Sequence", el)
	}
	if decodedSeq.Count() != 1 {
		t.Fatalf("sequence item count = %d, want 1", decodedSeq.Count())
	}

	gotPatientID, exists := decodedSeq.GetItem(0).GetString(tag.PatientID)
	if !exists {
		t.Fatal("decoded item missing PatientID")
	}
	if gotPatientID != "12345" {
		t.Fatalf("PatientID = %q, want %q", gotPatientID, "12345")
	}
}

func TestDecodeDIMSEMessageUsesParserCompatibilityFeatures(t *testing.T) {
	ds := dataset.New()
	for _, elem := range []element.Element{
		element.NewString(tag.SpecificCharacterSet, vr.CS, []string{"ISO_IR 192"}),
		element.NewStringWithEncoding(tag.PatientName, vr.PN, []string{"张三"}, unicode.UTF8),
		element.NewString(tag.New(0x0011, 0x0010), vr.LO, []string{"ACME"}),
		element.NewString(tag.New(0x0011, 0x1010), vr.LO, []string{"private value"}),
		element.NewUnsignedShortWithEndian(tag.Rows, []uint16{0x1234}, endian.Big),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatal(err)
		}
	}

	var raw bytes.Buffer
	if err := writer.Write(&raw, ds,
		writer.WithoutPreamble(),
		writer.WithTransferSyntax(transfer.ExplicitVRBigEndian)); err != nil {
		t.Fatal(err)
	}
	_, decoded, err := DecodeDIMSEMessage(nil, raw.Bytes(), transfer.ExplicitVRBigEndian)
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := decoded.GetString(tag.PatientName); !ok || got != "张三" {
		t.Fatalf("PatientName = %q, exists=%v", got, ok)
	}
	rows, err := decoded.GetUInt16(tag.Rows, 0)
	if err != nil || rows != 0x1234 {
		t.Fatalf("Rows = %#x, err=%v", rows, err)
	}
	privateElement, ok := decoded.Get(tag.New(0x0011, 0x1010))
	if !ok || privateElement.Tag().PrivateCreator() == nil ||
		privateElement.Tag().PrivateCreator().Creator() != "ACME" {
		t.Fatalf("private creator was not attached: %#v", privateElement)
	}
}

func TestDecodeDIMSEMessagePreservesEmptyPayloadAndErrorContext(t *testing.T) {
	command, data, err := DecodeDIMSEMessage(nil, nil, nil)
	if err != nil || command == nil || command.Count() != 0 || data != nil {
		t.Fatalf("empty decode = command %#v, data %#v, err %v", command, data, err)
	}

	_, _, err = DecodeDIMSEMessage([]byte{1, 2, 3, 4, 5}, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "failed to decode command dataset") {
		t.Fatalf("malformed command error = %v", err)
	}
}

func TestDecodeDIMSEMessagePreservesFragmentSequence(t *testing.T) {
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable([]uint32{0})
	fragments.AddFragment(buffer.NewMemory([]byte{1, 2, 3, 4}))
	ds := dataset.New()
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	var raw bytes.Buffer
	if err := writer.Write(&raw, ds,
		writer.WithoutPreamble(),
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		t.Fatal(err)
	}

	_, decoded, err := DecodeDIMSEMessage(nil, raw.Bytes(), transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	got, ok := decoded.Get(tag.PixelData)
	if !ok {
		t.Fatal("decoded dataset missing PixelData")
	}
	parsed, ok := got.(*element.OtherByteFragment)
	if !ok || parsed.FragmentCount() != 1 {
		t.Fatalf("decoded PixelData = %T with %#v, want one fragment", got, got)
	}
}

func TestDecodeDIMSEMessage_PreservesBinaryVRType(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewOtherWord(tag.PixelData, []byte{1, 2, 3, 4})); err != nil {
		t.Fatalf("failed to build test dataset: %v", err)
	}

	var raw bytes.Buffer
	if err := writer.Write(&raw, ds,
		writer.WithoutPreamble(),
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		t.Fatalf("failed to encode test dataset: %v", err)
	}

	_, decoded, err := DecodeDIMSEMessage(nil, raw.Bytes(), transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}

	el, exists := decoded.Get(tag.PixelData)
	if !exists {
		t.Fatal("decoded dataset missing PixelData")
	}

	if _, ok := el.(*element.OtherWord); !ok {
		t.Fatalf("decoded element type = %T, want *element.OtherWord", el)
	}
}
