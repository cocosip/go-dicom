// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestPDataTFDecodeRejectsOversizedPDVLength ensures a crafted PDV length that
// exceeds the remaining PDU payload is rejected instead of driving a huge
// allocation.
func TestPDataTFDecodeRejectsOversizedPDVLength(t *testing.T) {
	var data []byte
	var lenBuf [4]byte
	binary.BigEndian.PutUint32(lenBuf[:], 0xFFFFFFF0)
	data = append(data, lenBuf[:]...)
	data = append(data, 0x01, 0x03) // presentation context ID + control header

	pdt := &PDataTF{}
	if err := pdt.Decode(NewRawPDU(TypePDataTF, data)); err == nil {
		t.Fatal("expected error for oversized PDV length, got nil")
	}
}

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
