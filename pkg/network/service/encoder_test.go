// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
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
