// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"fmt"
)

// AAssociateRJ represents an A-ASSOCIATE-RJ (Association Reject) PDU.
// This PDU is sent to reject an association request.
//
// PDU Structure (4 bytes + 6-byte header = 10 bytes total):
//   - Bytes 0-5: PDU header (Type=0x03, Reserved, Length=4)
//   - Byte 6: Reserved (0x00)
//   - Byte 7: Result (1=rejected-permanent, 2=rejected-transient)
//   - Byte 8: Source (1, 2, or 3)
//   - Byte 9: Reason/Diag (depends on Source)
type AAssociateRJ struct {
	// Result indicates whether the rejection is permanent or transient:
	//   1 = rejected-permanent
	//   2 = rejected-transient
	Result byte

	// Source indicates the source of the rejection:
	//   1 = DICOM UL service-user
	//   2 = DICOM UL service-provider (ACSE related function)
	//   3 = DICOM UL service-provider (Presentation related function)
	Source byte

	// Reason indicates the reason for rejection (meaning depends on Source):
	//
	// If Source = 1 (service-user):
	//   1 = no-reason-given
	//   2 = application-context-name-not-supported
	//   3 = calling-AE-title-not-recognized
	//   7 = called-AE-title-not-recognized
	//
	// If Source = 2 (service-provider, ACSE):
	//   1 = no-reason-given
	//   2 = protocol-version-not-supported
	//
	// If Source = 3 (service-provider, Presentation):
	//   1 = temporary-congestion
	//   2 = local-limit-exceeded
	Reason byte
}

// Result codes for AAssociateRJ
const (
	ResultRejectedPermanent  byte = 1
	ResultRejectedTransient  byte = 2
)

// Source codes for AAssociateRJ
const (
	SourceServiceUser                     byte = 1
	SourceServiceProviderACSE             byte = 2
	SourceServiceProviderPresentation     byte = 3
)

// Reason codes when Source = 1 (service-user)
const (
	ReasonServiceUserNoReasonGiven                     byte = 1
	ReasonServiceUserApplicationContextNotSupported    byte = 2
	ReasonServiceUserCallingAETitleNotRecognized       byte = 3
	ReasonServiceUserCalledAETitleNotRecognized        byte = 7
)

// Reason codes when Source = 2 (service-provider, ACSE)
const (
	ReasonServiceProviderACSENoReasonGiven            byte = 1
	ReasonServiceProviderACSEProtocolVersionNotSupported byte = 2
)

// Reason codes when Source = 3 (service-provider, Presentation)
const (
	ReasonServiceProviderPresentationTemporaryCongestion  byte = 1
	ReasonServiceProviderPresentationLocalLimitExceeded   byte = 2
)

// NewAAssociateRJ creates a new A-ASSOCIATE-RJ PDU.
func NewAAssociateRJ(result, source, reason byte) *AAssociateRJ {
	return &AAssociateRJ{
		Result: result,
		Source: source,
		Reason: reason,
	}
}

// Encode encodes the A-ASSOCIATE-RJ to a RawPDU.
func (a *AAssociateRJ) Encode() (*RawPDU, error) {
	data := make([]byte, 4)
	data[0] = 0x00        // Reserved
	data[1] = a.Result
	data[2] = a.Source
	data[3] = a.Reason

	return NewRawPDU(TypeAAssociateRJ, data), nil
}

// Decode decodes an A-ASSOCIATE-RJ from a RawPDU.
func (a *AAssociateRJ) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAAssociateRJ {
		return fmt.Errorf("invalid PDU type: expected A-ASSOCIATE-RJ (0x03), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	if len(data) != 4 {
		return fmt.Errorf("invalid A-ASSOCIATE-RJ data length: %d (expected 4 bytes)", len(data))
	}

	// data[0] is reserved
	a.Result = data[1]
	a.Source = data[2]
	a.Reason = data[3]

	return nil
}

// String returns a human-readable representation of the rejection.
func (a *AAssociateRJ) String() string {
	return fmt.Sprintf("A-ASSOCIATE-RJ: Result=%s, Source=%s, Reason=%s",
		a.ResultString(), a.SourceString(), a.ReasonString())
}

// ResultString returns the string representation of the Result field.
func (a *AAssociateRJ) ResultString() string {
	switch a.Result {
	case ResultRejectedPermanent:
		return "rejected-permanent"
	case ResultRejectedTransient:
		return "rejected-transient"
	default:
		return fmt.Sprintf("unknown(0x%02X)", a.Result)
	}
}

// SourceString returns the string representation of the Source field.
func (a *AAssociateRJ) SourceString() string {
	switch a.Source {
	case SourceServiceUser:
		return "DICOM UL service-user"
	case SourceServiceProviderACSE:
		return "DICOM UL service-provider (ACSE)"
	case SourceServiceProviderPresentation:
		return "DICOM UL service-provider (Presentation)"
	default:
		return fmt.Sprintf("unknown(0x%02X)", a.Source)
	}
}

// ReasonString returns the string representation of the Reason field.
// The meaning of Reason depends on the Source field.
func (a *AAssociateRJ) ReasonString() string {
	switch a.Source {
	case SourceServiceUser:
		switch a.Reason {
		case ReasonServiceUserNoReasonGiven:
			return "no-reason-given"
		case ReasonServiceUserApplicationContextNotSupported:
			return "application-context-name-not-supported"
		case ReasonServiceUserCallingAETitleNotRecognized:
			return "calling-AE-title-not-recognized"
		case ReasonServiceUserCalledAETitleNotRecognized:
			return "called-AE-title-not-recognized"
		default:
			return fmt.Sprintf("unknown(0x%02X)", a.Reason)
		}

	case SourceServiceProviderACSE:
		switch a.Reason {
		case ReasonServiceProviderACSENoReasonGiven:
			return "no-reason-given"
		case ReasonServiceProviderACSEProtocolVersionNotSupported:
			return "protocol-version-not-supported"
		default:
			return fmt.Sprintf("unknown(0x%02X)", a.Reason)
		}

	case SourceServiceProviderPresentation:
		switch a.Reason {
		case ReasonServiceProviderPresentationTemporaryCongestion:
			return "temporary-congestion"
		case ReasonServiceProviderPresentationLocalLimitExceeded:
			return "local-limit-exceeded"
		default:
			return fmt.Sprintf("unknown(0x%02X)", a.Reason)
		}

	default:
		return fmt.Sprintf("unknown(0x%02X)", a.Reason)
	}
}
