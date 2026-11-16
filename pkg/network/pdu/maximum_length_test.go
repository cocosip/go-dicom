// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import "testing"

// TestMaximumLengthZeroEncodesEnforcedReceiveCap verifies that an unlimited
// configured maximum (0, PS3.8 Table D.1-1) is never advertised as 0: the encoded
// Maximum Length sub-item carries the cap the receive path actually enforces.
func TestMaximumLengthZeroEncodesEnforcedReceiveCap(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = testCalledAETitle
	rq.CallingAETitle = testCallingAETitle
	rq.UserInformation.MaximumLength = 0
	rq.PresentationContexts = []PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   testCTImageStorageUID,
		TransferSyntaxes: []string{testExplicitVRLittleLE},
	}}
	rawRQ, err := rq.Encode()
	if err != nil {
		t.Fatalf("AAssociateRQ.Encode() error = %v", err)
	}
	decodedRQ := &AAssociateRQ{}
	if err := decodedRQ.Decode(rawRQ); err != nil {
		t.Fatalf("AAssociateRQ.Decode() error = %v", err)
	}
	if got := decodedRQ.UserInformation.MaximumLength; got != MaxPDUDataLength || got == 0 {
		t.Fatalf("A-ASSOCIATE-RQ MaximumLength = %d, want enforced cap %d", got, MaxPDUDataLength)
	}

	ac := NewAAssociateAC()
	ac.CalledAETitle = testCalledAETitle
	ac.CallingAETitle = testCallingAETitle
	ac.UserInformation.MaximumLength = 0
	ac.PresentationContexts = []PresentationContextAC{{
		ID:             1,
		Result:         ResultAcceptance,
		TransferSyntax: testExplicitVRLittleLE,
	}}
	rawAC, err := ac.Encode()
	if err != nil {
		t.Fatalf("AAssociateAC.Encode() error = %v", err)
	}
	decodedAC := &AAssociateAC{}
	if err := decodedAC.Decode(rawAC); err != nil {
		t.Fatalf("AAssociateAC.Decode() error = %v", err)
	}
	if got := decodedAC.UserInformation.MaximumLength; got != MaxPDUDataLength || got == 0 {
		t.Fatalf("A-ASSOCIATE-AC MaximumLength = %d, want enforced cap %d", got, MaxPDUDataLength)
	}
}
