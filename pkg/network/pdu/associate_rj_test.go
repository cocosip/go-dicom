// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"testing"
)

func TestNewAAssociateRJ(t *testing.T) {
	rj := NewAAssociateRJ(ResultRejectedPermanent, SourceServiceUser, ReasonServiceUserCalledAETitleNotRecognized)

	if rj.Result != ResultRejectedPermanent {
		t.Errorf("Expected Result %d, got %d", ResultRejectedPermanent, rj.Result)
	}

	if rj.Source != SourceServiceUser {
		t.Errorf("Expected Source %d, got %d", SourceServiceUser, rj.Source)
	}

	if rj.Reason != ReasonServiceUserCalledAETitleNotRecognized {
		t.Errorf("Expected Reason %d, got %d", ReasonServiceUserCalledAETitleNotRecognized, rj.Reason)
	}
}

func TestAAssociateRJ_EncodeDecode(t *testing.T) {
	// Test all source and reason combinations
	tests := []struct {
		name   string
		result byte
		source byte
		reason byte
	}{
		{
			name:   "Permanent rejection - called AE not recognized",
			result: ResultRejectedPermanent,
			source: SourceServiceUser,
			reason: ReasonServiceUserCalledAETitleNotRecognized,
		},
		{
			name:   "Transient rejection - protocol version not supported",
			result: ResultRejectedTransient,
			source: SourceServiceProviderACSE,
			reason: ReasonServiceProviderACSEProtocolVersionNotSupported,
		},
		{
			name:   "Temporary congestion",
			result: ResultRejectedTransient,
			source: SourceServiceProviderPresentation,
			reason: ReasonServiceProviderPresentationTemporaryCongestion,
		},
		{
			name:   "Local limit exceeded",
			result: ResultRejectedPermanent,
			source: SourceServiceProviderPresentation,
			reason: ReasonServiceProviderPresentationLocalLimitExceeded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create
			original := NewAAssociateRJ(tt.result, tt.source, tt.reason)

			// Encode
			pdu, err := original.Encode()
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if pdu.Type != TypeAAssociateRJ {
				t.Errorf("Expected PDU type 0x03, got 0x%02X", pdu.Type)
			}

			if pdu.Length != 4 {
				t.Errorf("Expected PDU length 4, got %d", pdu.Length)
			}

			// Decode
			decoded := &AAssociateRJ{}
			if err := decoded.Decode(pdu); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if decoded.Result != original.Result {
				t.Errorf("Result mismatch: expected %d, got %d", original.Result, decoded.Result)
			}

			if decoded.Source != original.Source {
				t.Errorf("Source mismatch: expected %d, got %d", original.Source, decoded.Source)
			}

			if decoded.Reason != original.Reason {
				t.Errorf("Reason mismatch: expected %d, got %d", original.Reason, decoded.Reason)
			}
		})
	}
}

func TestAAssociateRJ_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateRQ, []byte{0x00, 0x01, 0x01, 0x01})

	rj := &AAssociateRJ{}
	err := rj.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAAssociateRJ_DecodeInvalidLength(t *testing.T) {
	// Too short
	pdu := NewRawPDU(TypeAAssociateRJ, []byte{0x00, 0x01})

	rj := &AAssociateRJ{}
	err := rj.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid data length, got nil")
	}
}

func TestAAssociateRJ_ResultString(t *testing.T) {
	tests := []struct {
		result   byte
		expected string
	}{
		{ResultRejectedPermanent, "rejected-permanent"},
		{ResultRejectedTransient, "rejected-transient"},
		{0xFF, "unknown(0xFF)"},
	}

	for _, tt := range tests {
		rj := &AAssociateRJ{Result: tt.result}
		result := rj.ResultString()
		if result != tt.expected {
			t.Errorf("ResultString(%d): expected %s, got %s", tt.result, tt.expected, result)
		}
	}
}

func TestAAssociateRJ_SourceString(t *testing.T) {
	tests := []struct {
		source   byte
		expected string
	}{
		{SourceServiceUser, "DICOM UL service-user"},
		{SourceServiceProviderACSE, "DICOM UL service-provider (ACSE)"},
		{SourceServiceProviderPresentation, "DICOM UL service-provider (Presentation)"},
		{0xFF, "unknown(0xFF)"},
	}

	for _, tt := range tests {
		rj := &AAssociateRJ{Source: tt.source}
		result := rj.SourceString()
		if result != tt.expected {
			t.Errorf("SourceString(%d): expected %s, got %s", tt.source, tt.expected, result)
		}
	}
}

func TestAAssociateRJ_ReasonString(t *testing.T) {
	// Test cases for A-ASSOCIATE-RJ PDU reason strings - A-ASSOCIATE-RJ is used to reject association requests
	testReasonStrings(t, []struct {
		name     string
		source   byte
		reason   byte
		expected string
	}{
		{"Service user - no reason", SourceServiceUser, ReasonServiceUserNoReasonGiven, "no-reason-given"},
		{"Service user - application context not supported", SourceServiceUser, ReasonServiceUserApplicationContextNotSupported, "application-context-name-not-supported"},
		{"Service user - calling AE not recognized", SourceServiceUser, ReasonServiceUserCallingAETitleNotRecognized, "calling-AE-title-not-recognized"},
		{"Service user - called AE not recognized", SourceServiceUser, ReasonServiceUserCalledAETitleNotRecognized, "called-AE-title-not-recognized"},
		{"Service provider ACSE - protocol version not supported", SourceServiceProviderACSE, ReasonServiceProviderACSEProtocolVersionNotSupported, "protocol-version-not-supported"},
		{"Service provider Presentation - temporary congestion", SourceServiceProviderPresentation, ReasonServiceProviderPresentationTemporaryCongestion, "temporary-congestion"},
		{"Service provider Presentation - local limit exceeded", SourceServiceProviderPresentation, ReasonServiceProviderPresentationLocalLimitExceeded, "local-limit-exceeded"},
		{"Unknown source and reason", 0xFF, 0xFF, "unknown(0xFF)"},
	}, func(source, reason byte) reasonStringer {
		return &AAssociateRJ{Source: source, Reason: reason}
	})
}

func TestAAssociateRJ_String(t *testing.T) {
	rj := NewAAssociateRJ(
		ResultRejectedPermanent,
		SourceServiceUser,
		ReasonServiceUserCalledAETitleNotRecognized,
	)

	str := rj.String()
	expected := "A-ASSOCIATE-RJ: Result=rejected-permanent, Source=DICOM UL service-user, Reason=called-AE-title-not-recognized"

	if str != expected {
		t.Errorf("String() mismatch:\nexpected: %s\ngot:      %s", expected, str)
	}
}

func TestAAssociateRJ_AllServiceUserReasons(t *testing.T) {
	// Test all service-user rejection reasons
	reasons := []struct {
		code byte
		name string
	}{
		{ReasonServiceUserNoReasonGiven, "no-reason-given"},
		{ReasonServiceUserApplicationContextNotSupported, "application-context-name-not-supported"},
		{ReasonServiceUserCallingAETitleNotRecognized, "calling-AE-title-not-recognized"},
		{ReasonServiceUserCalledAETitleNotRecognized, "called-AE-title-not-recognized"},
	}

	for _, r := range reasons {
		rj := NewAAssociateRJ(ResultRejectedPermanent, SourceServiceUser, r.code)

		// Encode and decode
		pdu, err := rj.Encode()
		if err != nil {
			t.Fatalf("Encode failed for reason %s: %v", r.name, err)
		}

		decoded := &AAssociateRJ{}
		if err := decoded.Decode(pdu); err != nil {
			t.Fatalf("Decode failed for reason %s: %v", r.name, err)
		}

		// Verify
		if decoded.Reason != r.code {
			t.Errorf("Reason code mismatch for %s: expected %d, got %d", r.name, r.code, decoded.Reason)
		}

		reasonStr := decoded.ReasonString()
		if reasonStr != r.name {
			t.Errorf("ReasonString mismatch: expected %s, got %s", r.name, reasonStr)
		}
	}
}
