// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"testing"
)

func TestNewAAbort(t *testing.T) {
	abort := NewAAbort(AbortSourceServiceUser, AbortReasonServiceUserNotSpecified)

	if abort.Source != AbortSourceServiceUser {
		t.Errorf("Expected Source %d, got %d", AbortSourceServiceUser, abort.Source)
	}

	if abort.Reason != AbortReasonServiceUserNotSpecified {
		t.Errorf("Expected Reason %d, got %d", AbortReasonServiceUserNotSpecified, abort.Reason)
	}
}

func TestNewAAbortServiceUser(t *testing.T) {
	abort := NewAAbortServiceUser()

	if abort.Source != AbortSourceServiceUser {
		t.Errorf("Expected Source %d, got %d", AbortSourceServiceUser, abort.Source)
	}

	if abort.Reason != AbortReasonServiceUserNotSpecified {
		t.Errorf("Expected Reason %d, got %d", AbortReasonServiceUserNotSpecified, abort.Reason)
	}
}

func TestNewAAbortServiceProvider(t *testing.T) {
	abort := NewAAbortServiceProvider(AbortReasonServiceProviderUnexpectedPDU)

	if abort.Source != AbortSourceServiceProvider {
		t.Errorf("Expected Source %d, got %d", AbortSourceServiceProvider, abort.Source)
	}

	if abort.Reason != AbortReasonServiceProviderUnexpectedPDU {
		t.Errorf("Expected Reason %d, got %d", AbortReasonServiceProviderUnexpectedPDU, abort.Reason)
	}
}

func TestAAbort_EncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		source byte
		reason byte
	}{
		{
			name:   "Service user abort",
			source: AbortSourceServiceUser,
			reason: AbortReasonServiceUserNotSpecified,
		},
		{
			name:   "Service provider - unrecognized PDU",
			source: AbortSourceServiceProvider,
			reason: AbortReasonServiceProviderUnrecognizedPDU,
		},
		{
			name:   "Service provider - unexpected PDU",
			source: AbortSourceServiceProvider,
			reason: AbortReasonServiceProviderUnexpectedPDU,
		},
		{
			name:   "Service provider - unrecognized PDU parameter",
			source: AbortSourceServiceProvider,
			reason: AbortReasonServiceProviderUnrecognizedPDUParam,
		},
		{
			name:   "Service provider - unexpected PDU parameter",
			source: AbortSourceServiceProvider,
			reason: AbortReasonServiceProviderUnexpectedPDUParam,
		},
		{
			name:   "Service provider - invalid PDU parameter",
			source: AbortSourceServiceProvider,
			reason: AbortReasonServiceProviderInvalidPDUParam,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create
			original := NewAAbort(tt.source, tt.reason)

			// Encode
			pdu, err := original.Encode()
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			if pdu.Type != TypeAAbort {
				t.Errorf("Expected PDU type 0x07, got 0x%02X", pdu.Type)
			}

			if pdu.Length != 4 {
				t.Errorf("Expected PDU length 4, got %d", pdu.Length)
			}

			// Decode
			decoded := &AAbort{}
			if err := decoded.Decode(pdu); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			// Compare
			if decoded.Source != original.Source {
				t.Errorf("Source mismatch: expected %d, got %d", original.Source, decoded.Source)
			}

			if decoded.Reason != original.Reason {
				t.Errorf("Reason mismatch: expected %d, got %d", original.Reason, decoded.Reason)
			}
		})
	}
}

func TestAAbort_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAReleaseRQ, []byte{0x00, 0x00, 0x00, 0x00})

	abort := &AAbort{}
	err := abort.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAAbort_DecodeInvalidLength(t *testing.T) {
	pdu := NewRawPDU(TypeAAbort, []byte{0x00, 0x00})

	abort := &AAbort{}
	err := abort.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid data length, got nil")
	}
}

func TestAAbort_SourceString(t *testing.T) {
	tests := []struct {
		source   byte
		expected string
	}{
		{AbortSourceServiceUser, dicomULServiceUserName},
		{AbortSourceServiceProvider, "DICOM UL service-provider"},
		{0xFF, unknownFFName},
	}

	for _, tt := range tests {
		abort := &AAbort{Source: tt.source}
		result := abort.SourceString()
		if result != tt.expected {
			t.Errorf("SourceString(%d): expected %s, got %s", tt.source, tt.expected, result)
		}
	}
}

func TestAAbort_ReasonString(t *testing.T) {
	// Test cases for A-ABORT PDU reason strings - A-ABORT is used for abnormal termination of associations
	testReasonStrings(t, []struct {
		name     string
		source   byte
		reason   byte
		expected string
	}{
		{"Service user - not specified", AbortSourceServiceUser, AbortReasonServiceUserNotSpecified, abortReasonNotSpecifiedName},
		{"Service provider - not specified", AbortSourceServiceProvider, AbortReasonServiceProviderNotSpecified, abortReasonNotSpecifiedName},
		{"Service provider - unrecognized PDU", AbortSourceServiceProvider, AbortReasonServiceProviderUnrecognizedPDU, abortReasonUnrecognizedPDUName},
		{"Service provider - unexpected PDU", AbortSourceServiceProvider, AbortReasonServiceProviderUnexpectedPDU, abortReasonUnexpectedPDUName},
		{"Service provider - unrecognized PDU parameter", AbortSourceServiceProvider, AbortReasonServiceProviderUnrecognizedPDUParam, abortReasonUnrecognizedPDUParameterName},
		{"Service provider - unexpected PDU parameter", AbortSourceServiceProvider, AbortReasonServiceProviderUnexpectedPDUParam, abortReasonUnexpectedPDUParameterName},
		{"Service provider - invalid PDU parameter", AbortSourceServiceProvider, AbortReasonServiceProviderInvalidPDUParam, abortReasonInvalidPDUParameterName},
		{"Unknown source and reason", 0xFF, 0xFF, unknownFFName},
	}, func(source, reason byte) reasonStringer {
		return &AAbort{Source: source, Reason: reason}
	})
}

func TestAAbort_String(t *testing.T) {
	abort := NewAAbortServiceProvider(AbortReasonServiceProviderUnexpectedPDU)

	str := abort.String()
	expected := pduTypeAAbortName + ": Source=DICOM UL service-provider, Reason=" + abortReasonUnexpectedPDUName

	if str != expected {
		t.Errorf("String() mismatch:\nexpected: %s\ngot:      %s", expected, str)
	}
}

func TestAAbort_AllServiceProviderReasons(t *testing.T) {
	// Test all service-provider abort reasons
	reasons := []struct {
		code byte
		name string
	}{
		{AbortReasonServiceProviderNotSpecified, abortReasonNotSpecifiedName},
		{AbortReasonServiceProviderUnrecognizedPDU, abortReasonUnrecognizedPDUName},
		{AbortReasonServiceProviderUnexpectedPDU, abortReasonUnexpectedPDUName},
		{AbortReasonServiceProviderUnrecognizedPDUParam, abortReasonUnrecognizedPDUParameterName},
		{AbortReasonServiceProviderUnexpectedPDUParam, abortReasonUnexpectedPDUParameterName},
		{AbortReasonServiceProviderInvalidPDUParam, abortReasonInvalidPDUParameterName},
	}

	for _, r := range reasons {
		abort := NewAAbortServiceProvider(r.code)

		// Encode and decode
		pdu, err := abort.Encode()
		if err != nil {
			t.Fatalf("Encode failed for reason %s: %v", r.name, err)
		}

		decoded := &AAbort{}
		if err := decoded.Decode(pdu); err != nil {
			t.Fatalf("Decode failed for reason %s: %v", r.name, err)
		}

		// Verify
		if decoded.Source != AbortSourceServiceProvider {
			t.Errorf("Source mismatch for %s: expected %d, got %d", r.name, AbortSourceServiceProvider, decoded.Source)
		}

		if decoded.Reason != r.code {
			t.Errorf("Reason code mismatch for %s: expected %d, got %d", r.name, r.code, decoded.Reason)
		}

		reasonStr := decoded.ReasonString()
		if reasonStr != r.name {
			t.Errorf("ReasonString mismatch: expected %s, got %s", r.name, reasonStr)
		}
	}
}
