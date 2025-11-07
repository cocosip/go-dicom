// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"fmt"
)

// AReleaseRQ represents an A-RELEASE-RQ (Association Release Request) PDU.
// This PDU is sent to request the orderly release of an association.
//
// PDU Structure (4 bytes + 6-byte header = 10 bytes total):
//   - Bytes 0-5: PDU header (Type=0x05, Reserved, Length=4)
//   - Bytes 6-9: Reserved (all 0x00)
type AReleaseRQ struct {
	// No fields - this is a simple signaling PDU
}

// NewAReleaseRQ creates a new A-RELEASE-RQ PDU.
func NewAReleaseRQ() *AReleaseRQ {
	return &AReleaseRQ{}
}

// Encode encodes the A-RELEASE-RQ to a RawPDU.
func (a *AReleaseRQ) Encode() (*RawPDU, error) {
	// Data is 4 bytes of reserved (all 0x00)
	data := make([]byte, 4)
	return NewRawPDU(TypeAReleaseRQ, data), nil
}

// Decode decodes an A-RELEASE-RQ from a RawPDU.
func (a *AReleaseRQ) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAReleaseRQ {
		return fmt.Errorf("invalid PDU type: expected A-RELEASE-RQ (0x05), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	if len(data) != 4 {
		return fmt.Errorf("invalid A-RELEASE-RQ data length: %d (expected 4 bytes)", len(data))
	}

	// No fields to decode
	return nil
}

// String returns a human-readable representation.
func (a *AReleaseRQ) String() string {
	return "A-RELEASE-RQ"
}

// AReleaseRP represents an A-RELEASE-RP (Association Release Response) PDU.
// This PDU is sent in response to an A-RELEASE-RQ to confirm the release of an association.
//
// PDU Structure (4 bytes + 6-byte header = 10 bytes total):
//   - Bytes 0-5: PDU header (Type=0x06, Reserved, Length=4)
//   - Bytes 6-9: Reserved (all 0x00)
type AReleaseRP struct {
	// No fields - this is a simple signaling PDU
}

// NewAReleaseRP creates a new A-RELEASE-RP PDU.
func NewAReleaseRP() *AReleaseRP {
	return &AReleaseRP{}
}

// Encode encodes the A-RELEASE-RP to a RawPDU.
func (a *AReleaseRP) Encode() (*RawPDU, error) {
	// Data is 4 bytes of reserved (all 0x00)
	data := make([]byte, 4)
	return NewRawPDU(TypeAReleaseRP, data), nil
}

// Decode decodes an A-RELEASE-RP from a RawPDU.
func (a *AReleaseRP) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAReleaseRP {
		return fmt.Errorf("invalid PDU type: expected A-RELEASE-RP (0x06), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	if len(data) != 4 {
		return fmt.Errorf("invalid A-RELEASE-RP data length: %d (expected 4 bytes)", len(data))
	}

	// No fields to decode
	return nil
}

// String returns a human-readable representation.
func (a *AReleaseRP) String() string {
	return "A-RELEASE-RP"
}
