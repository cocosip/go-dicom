// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package pdu implements DICOM Protocol Data Unit (PDU) encoding and decoding.
package pdu

import (
    "encoding/binary"
    "fmt"
    "io"
    "math"
)

// PDU type constants as defined in DICOM standard PS3.8
const (
	// TypeAAssociateRQ represents A-ASSOCIATE-RQ PDU (0x01)
	TypeAAssociateRQ byte = 0x01

	// TypeAAssociateAC represents A-ASSOCIATE-AC PDU (0x02)
	TypeAAssociateAC byte = 0x02

	// TypeAAssociateRJ represents A-ASSOCIATE-RJ PDU (0x03)
	TypeAAssociateRJ byte = 0x03

	// TypePDataTF represents P-DATA-TF PDU (0x04)
	TypePDataTF byte = 0x04

	// TypeAReleaseRQ represents A-RELEASE-RQ PDU (0x05)
	TypeAReleaseRQ byte = 0x05

	// TypeAReleaseRP represents A-RELEASE-RP PDU (0x06)
	TypeAReleaseRP byte = 0x06

	// TypeAAbort represents A-ABORT PDU (0x07)
	TypeAAbort byte = 0x07
)

// TypeString returns the string representation of a PDU type.
func TypeString(pduType byte) string {
	switch pduType {
	case TypeAAssociateRQ:
		return "A-ASSOCIATE-RQ"
	case TypeAAssociateAC:
		return "A-ASSOCIATE-AC"
	case TypeAAssociateRJ:
		return "A-ASSOCIATE-RJ"
	case TypePDataTF:
		return "P-DATA-TF"
	case TypeAReleaseRQ:
		return "A-RELEASE-RQ"
	case TypeAReleaseRP:
		return "A-RELEASE-RP"
	case TypeAAbort:
		return "A-ABORT"
	default:
		return fmt.Sprintf("Unknown(0x%02X)", pduType)
	}
}

// RawPDU represents a raw DICOM PDU with type, length, and data.
// The PDU structure is:
//   - Byte 0: PDU-type (1 byte)
//   - Byte 1: Reserved (1 byte, always 0x00)
//   - Bytes 2-5: PDU-length (4 bytes, big-endian unsigned integer)
//   - Bytes 6+: PDU data (variable length)
//
// Note: PDU-length does NOT include the 6-byte header (type + reserved + length).
type RawPDU struct {
	Type     byte   // PDU type (0x01-0x07)
	Reserved byte   // Reserved field (should be 0x00)
	Length   uint32 // Length of PDU data (excludes 6-byte header)
	Data     []byte // PDU data
}

// NewRawPDU creates a new RawPDU with the given type and data.
func NewRawPDU(pduType byte, data []byte) *RawPDU {
    if len(data) > int(math.MaxUint32) {
        // Infeasible in practice, but guard to satisfy static analysis
        panic(fmt.Sprintf("raw PDU data too large: %d", len(data)))
    }
    return &RawPDU{
        Type:     pduType,
        Reserved: 0x00,
        Length:   uint32(len(data)), // #nosec G115 -- guarded above
        Data:     data,
    }
}

// String returns a string representation of the PDU.
func (p *RawPDU) String() string {
	return fmt.Sprintf("%s (length=%d)", TypeString(p.Type), p.Length)
}

// Read reads a RawPDU from an io.Reader.
// It reads the 6-byte header first, then reads the PDU data based on the length field.
func (p *RawPDU) Read(r io.Reader) error {
	// Read 6-byte header
	header := make([]byte, 6)
	if _, err := io.ReadFull(r, header); err != nil {
		return fmt.Errorf("reading PDU header: %w", err)
	}

	p.Type = header[0]
	p.Reserved = header[1]
	p.Length = binary.BigEndian.Uint32(header[2:6])

	// Validate PDU type
	if !isValidPDUType(p.Type) {
		return fmt.Errorf("invalid PDU type: 0x%02X", p.Type)
	}

	// Read PDU data
	if p.Length > 0 {
		p.Data = make([]byte, p.Length)
		if _, err := io.ReadFull(r, p.Data); err != nil {
			return fmt.Errorf("reading PDU data (length=%d): %w", p.Length, err)
		}
	} else {
		p.Data = nil
	}

	return nil
}

// Write writes the RawPDU to an io.Writer.
// It writes the 6-byte header followed by the PDU data.
func (p *RawPDU) Write(w io.Writer) error {
    // Update length from data
    if len(p.Data) > int(math.MaxUint32) {
        return fmt.Errorf("raw PDU data too large: %d", len(p.Data))
    }
    p.Length = uint32(len(p.Data)) // #nosec G115 -- bounded by check above

    // Build 6-byte header using fixed-size array to avoid index warnings
    var header [6]byte
    header[0] = p.Type        // #nosec G602 -- fixed-size header array [6]byte
    header[1] = p.Reserved    // #nosec G602 -- fixed-size header array [6]byte
    binary.BigEndian.PutUint32(header[2:6], p.Length)

	// Write header
    if _, err := w.Write(header[:]); err != nil {
        return fmt.Errorf("writing PDU header: %w", err)
    }

	// Write data
	if len(p.Data) > 0 {
		if _, err := w.Write(p.Data); err != nil {
			return fmt.Errorf("writing PDU data: %w", err)
		}
	}

	return nil
}

// TotalLength returns the total length of the PDU including the 6-byte header.
func (p *RawPDU) TotalLength() uint32 {
	return 6 + p.Length
}

// isValidPDUType checks if the given PDU type is valid.
func isValidPDUType(pduType byte) bool {
	switch pduType {
	case TypeAAssociateRQ, TypeAAssociateAC, TypeAAssociateRJ,
		TypePDataTF, TypeAReleaseRQ, TypeAReleaseRP, TypeAAbort:
		return true
	default:
		return false
	}
}

// ReadPDU is a convenience function to read a RawPDU from an io.Reader.
func ReadPDU(r io.Reader) (*RawPDU, error) {
	pdu := &RawPDU{}
	if err := pdu.Read(r); err != nil {
		return nil, err
	}
	return pdu, nil
}

// WritePDU is a convenience function to write a RawPDU to an io.Writer.
func WritePDU(w io.Writer, pdu *RawPDU) error {
	return pdu.Write(w)
}
