// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"fmt"
)

// AAbort represents an A-ABORT PDU.
// This PDU is sent to abort an association (non-orderly release).
//
// PDU Structure (4 bytes + 6-byte header = 10 bytes total):
//   - Bytes 0-5: PDU header (Type=0x07, Reserved, Length=4)
//   - Bytes 6-7: Reserved (0x0000)
//   - Byte 8: Source (0=DICOM UL service-user, 2=DICOM UL service-provider)
//   - Byte 9: Reason/Diag (depends on Source)
type AAbort struct {
	// Source indicates the source of the abort:
	//   0 = DICOM UL service-user (initiated by the application)
	//   2 = DICOM UL service-provider (initiated by the DICOM UL service-provider)
	Source byte

	// Reason indicates the reason for abort (meaning depends on Source):
	//
	// If Source = 0 (service-user):
	//   0 = reason-not-specified
	//
	// If Source = 2 (service-provider):
	//   0 = reason-not-specified
	//   1 = unrecognized-PDU
	//   2 = unexpected-PDU
	//   4 = unrecognized-PDU-parameter
	//   5 = unexpected-PDU-parameter
	//   6 = invalid-PDU-parameter value
	Reason byte
}

// Source codes for AAbort
const (
	AbortSourceServiceUser     byte = 0
	AbortSourceServiceProvider byte = 2
)

// Reason codes when Source = 0 (service-user)
const (
	AbortReasonServiceUserNotSpecified byte = 0
)

// Reason codes when Source = 2 (service-provider)
const (
	AbortReasonServiceProviderNotSpecified         byte = 0
	AbortReasonServiceProviderUnrecognizedPDU      byte = 1
	AbortReasonServiceProviderUnexpectedPDU        byte = 2
	AbortReasonServiceProviderUnrecognizedPDUParam byte = 4
	AbortReasonServiceProviderUnexpectedPDUParam   byte = 5
	AbortReasonServiceProviderInvalidPDUParam      byte = 6
)

// NewAAbort creates a new A-ABORT PDU.
func NewAAbort(source, reason byte) *AAbort {
	return &AAbort{
		Source: source,
		Reason: reason,
	}
}

// NewAAbortServiceUser creates an A-ABORT initiated by the service user.
func NewAAbortServiceUser() *AAbort {
	return &AAbort{
		Source: AbortSourceServiceUser,
		Reason: AbortReasonServiceUserNotSpecified,
	}
}

// NewAAbortServiceProvider creates an A-ABORT initiated by the service provider with a specific reason.
func NewAAbortServiceProvider(reason byte) *AAbort {
	return &AAbort{
		Source: AbortSourceServiceProvider,
		Reason: reason,
	}
}

// Encode encodes the A-ABORT to a RawPDU.
func (a *AAbort) Encode() (*RawPDU, error) {
	data := make([]byte, 4)
	data[0] = 0x00 // Reserved
	data[1] = 0x00 // Reserved
	data[2] = a.Source
	data[3] = a.Reason

	return NewRawPDU(TypeAAbort, data), nil
}

// Decode decodes an A-ABORT from a RawPDU.
func (a *AAbort) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAAbort {
		return fmt.Errorf("invalid PDU type: expected A-ABORT (0x07), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	if len(data) != 4 {
		return fmt.Errorf("invalid A-ABORT data length: %d (expected 4 bytes)", len(data))
	}

	// data[0] and data[1] are reserved
	a.Source = data[2]
	a.Reason = data[3]

	return nil
}

// String returns a human-readable representation of the abort.
func (a *AAbort) String() string {
	return fmt.Sprintf("A-ABORT: Source=%s, Reason=%s",
		a.SourceString(), a.ReasonString())
}

// SourceString returns the string representation of the Source field.
func (a *AAbort) SourceString() string {
	switch a.Source {
	case AbortSourceServiceUser:
		return "DICOM UL service-user"
	case AbortSourceServiceProvider:
		return "DICOM UL service-provider"
	default:
		return fmt.Sprintf("unknown(0x%02X)", a.Source)
	}
}

// ReasonString returns the string representation of the Reason field.
// The meaning of Reason depends on the Source field.
func (a *AAbort) ReasonString() string {
	switch a.Source {
	case AbortSourceServiceUser:
		switch a.Reason {
		case AbortReasonServiceUserNotSpecified:
			return "reason-not-specified"
		default:
			return fmt.Sprintf("unknown(0x%02X)", a.Reason)
		}

	case AbortSourceServiceProvider:
		switch a.Reason {
		case AbortReasonServiceProviderNotSpecified:
			return "reason-not-specified"
		case AbortReasonServiceProviderUnrecognizedPDU:
			return "unrecognized-PDU"
		case AbortReasonServiceProviderUnexpectedPDU:
			return "unexpected-PDU"
		case AbortReasonServiceProviderUnrecognizedPDUParam:
			return "unrecognized-PDU-parameter"
		case AbortReasonServiceProviderUnexpectedPDUParam:
			return "unexpected-PDU-parameter"
		case AbortReasonServiceProviderInvalidPDUParam:
			return "invalid-PDU-parameter"
		default:
			return fmt.Sprintf("unknown(0x%02X)", a.Reason)
		}

	default:
		return fmt.Sprintf("unknown(0x%02X)", a.Reason)
	}
}
