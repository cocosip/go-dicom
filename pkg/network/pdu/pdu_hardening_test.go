// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import (
	"bytes"
	"errors"
	"testing"
)

// TestAAssociateACDecodeToleratesTransferSyntaxOnRejectedContext verifies that a
// rejected presentation context carrying a transfer syntax sub-item (emitted by
// some real PACS) is tolerated rather than aborting the whole association.
func TestAAssociateACDecodeToleratesTransferSyntaxOnRejectedContext(t *testing.T) {
	var pcBuf bytes.Buffer
	pcBuf.WriteByte(0x01) // presentation context ID (odd)
	pcBuf.WriteByte(0x00) // reserved
	pcBuf.WriteByte(ResultUserRejection)
	pcBuf.WriteByte(0x00) // reserved
	if err := writeItem(&pcBuf, ItemTypeTransferSyntax, []byte("1.2.840.10008.1.2")); err != nil {
		t.Fatalf("writeItem: %v", err)
	}

	pc, err := (&AAssociateAC{}).decodePresentationContextAC(pcBuf.Bytes())
	if err != nil {
		t.Fatalf("expected tolerant decode, got error: %v", err)
	}
	if pc.TransferSyntax != "" {
		t.Errorf("expected transfer syntax dropped on rejected context, got %q", pc.TransferSyntax)
	}
}

// TestAAssociateACEncodeRejectsInvalidPresentationContext verifies the AC
// encoder rejects an even presentation context ID, matching the RQ encoder.
func TestAAssociateACEncodeRejectsInvalidPresentationContext(t *testing.T) {
	ac := NewAAssociateAC()
	ac.CalledAETitle = "SERVER"
	ac.CallingAETitle = "CLIENT"
	ac.PresentationContexts = []PresentationContextAC{
		{ID: 2, Result: ResultAcceptance, TransferSyntax: "1.2.840.10008.1.2"},
	}
	if _, err := ac.Encode(); err == nil {
		t.Fatal("expected error for even presentation context ID, got nil")
	}
}

// appendPresentationContextItems appends count presentation context items. Each
// body carries a minimal 4-byte header and, when abstractSyntax is not empty, an
// abstract syntax sub-item (required by the A-ASSOCIATE-RQ decoder).
func appendPresentationContextItems(data []byte, itemType byte, count int, abstractSyntax string) []byte {
	var items bytes.Buffer
	for i := 0; i < count; i++ {
		body := []byte{byte(2*i + 1), 0x00, ResultUserRejection, 0x00} // odd ID, 3 reserved bytes
		if abstractSyntax != "" {
			var sub bytes.Buffer
			if err := writeItem(&sub, ItemTypeAbstractSyntax, []byte(abstractSyntax)); err != nil {
				panic(err)
			}
			if err := writeItem(&sub, ItemTypeTransferSyntax, []byte("1.2.840.10008.1.2")); err != nil {
				panic(err)
			}
			body = append(body, sub.Bytes()...)
		}
		if err := writeItem(&items, itemType, body); err != nil {
			panic(err)
		}
	}
	return append(data, items.Bytes()...)
}

// TestAAssociateRQDecodeBoundsPresentationContexts verifies the decoder enforces
// the presentation context limit itself, so a large association PDU cannot make
// it allocate an unbounded number of contexts before negotiation policy runs.
func TestAAssociateRQDecodeBoundsPresentationContexts(t *testing.T) {
	const abstractSyntax = "1.2.840.10008.5.1.4.1.1.2"
	build := func(count int) *RawPDU {
		data := make([]byte, 68)
		return NewRawPDU(TypeAAssociateRQ,
			appendPresentationContextItems(data, ItemTypePresentationContextRQ, count, abstractSyntax))
	}

	atLimit := &AAssociateRQ{}
	if err := atLimit.Decode(build(maximumPresentationContexts)); err != nil {
		t.Fatalf("Decode() at %d contexts error = %v, want nil", maximumPresentationContexts, err)
	}
	if got := len(atLimit.PresentationContexts); got != maximumPresentationContexts {
		t.Fatalf("decoded %d contexts, want %d", got, maximumPresentationContexts)
	}

	overLimit := &AAssociateRQ{}
	err := overLimit.Decode(build(maximumPresentationContexts + 1))
	if !errors.Is(err, ErrTooManyPresentationContexts) {
		t.Fatalf("Decode() over limit error = %v, want %v", err, ErrTooManyPresentationContexts)
	}
}

// TestAAssociateACDecodeBoundsPresentationContexts is the A-ASSOCIATE-AC
// counterpart of the request-side bound test.
func TestAAssociateACDecodeBoundsPresentationContexts(t *testing.T) {
	build := func(count int) *RawPDU {
		data := make([]byte, 68)
		return NewRawPDU(TypeAAssociateAC, appendPresentationContextItems(data, ItemTypePresentationContextAC, count, ""))
	}

	if err := (&AAssociateAC{}).Decode(build(maximumPresentationContexts)); err != nil {
		t.Fatalf("Decode() at %d contexts error = %v, want nil", maximumPresentationContexts, err)
	}

	err := (&AAssociateAC{}).Decode(build(maximumPresentationContexts + 1))
	if !errors.Is(err, ErrTooManyPresentationContexts) {
		t.Fatalf("Decode() over limit error = %v, want %v", err, ErrTooManyPresentationContexts)
	}
}
