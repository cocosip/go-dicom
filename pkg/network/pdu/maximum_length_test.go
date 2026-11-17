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

// TestMaximumLengthOverCapEncodesEnforcedReceiveCap verifies that a configured
// Maximum Length above the enforced receive cap is advertised as the cap, not
// as the configured value: a peer that honoured the advertisement would send
// PDUs that ReadMax rejects as oversized. Values within the cap are unchanged.
func TestMaximumLengthOverCapEncodesEnforcedReceiveCap(t *testing.T) {
	const oversized = 200 * 1024 * 1024
	const withinCap = 32 * 1024

	for _, tc := range []struct {
		name       string
		configured uint32
		want       uint32
	}{
		{name: "over cap is clamped", configured: oversized, want: MaxPDUDataLength},
		{name: "within cap is kept", configured: withinCap, want: withinCap},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rq := NewAAssociateRQ()
			rq.CalledAETitle = testCalledAETitle
			rq.CallingAETitle = testCallingAETitle
			rq.UserInformation.MaximumLength = tc.configured
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
			if got := decodedRQ.UserInformation.MaximumLength; got != tc.want {
				t.Fatalf("A-ASSOCIATE-RQ MaximumLength = %d, want %d", got, tc.want)
			}

			ac := NewAAssociateAC()
			ac.CalledAETitle = testCalledAETitle
			ac.CallingAETitle = testCallingAETitle
			ac.UserInformation.MaximumLength = tc.configured
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
			if got := decodedAC.UserInformation.MaximumLength; got != tc.want {
				t.Fatalf("A-ASSOCIATE-AC MaximumLength = %d, want %d", got, tc.want)
			}
		})
	}
}
