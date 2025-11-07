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
		{AbortSourceServiceUser, "DICOM UL service-user"},
		{AbortSourceServiceProvider, "DICOM UL service-provider"},
		{0xFF, "unknown(0xFF)"},
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
	tests := []struct {
		name     string
		source   byte
		reason   byte
		expected string
	}{
		{
			name:     "Service user - not specified",
			source:   AbortSourceServiceUser,
			reason:   AbortReasonServiceUserNotSpecified,
			expected: "reason-not-specified",
		},
		{
			name:     "Service provider - not specified",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderNotSpecified,
			expected: "reason-not-specified",
		},
		{
			name:     "Service provider - unrecognized PDU",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderUnrecognizedPDU,
			expected: "unrecognized-PDU",
		},
		{
			name:     "Service provider - unexpected PDU",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderUnexpectedPDU,
			expected: "unexpected-PDU",
		},
		{
			name:     "Service provider - unrecognized PDU parameter",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderUnrecognizedPDUParam,
			expected: "unrecognized-PDU-parameter",
		},
		{
			name:     "Service provider - unexpected PDU parameter",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderUnexpectedPDUParam,
			expected: "unexpected-PDU-parameter",
		},
		{
			name:     "Service provider - invalid PDU parameter",
			source:   AbortSourceServiceProvider,
			reason:   AbortReasonServiceProviderInvalidPDUParam,
			expected: "invalid-PDU-parameter",
		},
		{
			name:     "Unknown source and reason",
			source:   0xFF,
			reason:   0xFF,
			expected: "unknown(0xFF)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			abort := &AAbort{Source: tt.source, Reason: tt.reason}
			result := abort.ReasonString()
			if result != tt.expected {
				t.Errorf("ReasonString: expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestAAbort_String(t *testing.T) {
	abort := NewAAbortServiceProvider(AbortReasonServiceProviderUnexpectedPDU)

	str := abort.String()
	expected := "A-ABORT: Source=DICOM UL service-provider, Reason=unexpected-PDU"

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
		{AbortReasonServiceProviderNotSpecified, "reason-not-specified"},
		{AbortReasonServiceProviderUnrecognizedPDU, "unrecognized-PDU"},
		{AbortReasonServiceProviderUnexpectedPDU, "unexpected-PDU"},
		{AbortReasonServiceProviderUnrecognizedPDUParam, "unrecognized-PDU-parameter"},
		{AbortReasonServiceProviderUnexpectedPDUParam, "unexpected-PDU-parameter"},
		{AbortReasonServiceProviderInvalidPDUParam, "invalid-PDU-parameter"},
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
