// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package pdu implements DICOM Protocol Data Unit (PDU) encoding and decoding.
package pdu

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
)

// PDU type constants as defined in DICOM standard PS3.8
const (
	pduTypeAAssociateRQName = "A-ASSOCIATE-RQ"
	pduTypeAAssociateACName = "A-ASSOCIATE-AC"
	pduTypeAAssociateRJName = "A-ASSOCIATE-RJ"
	pduTypePDataTFName      = "P-DATA-TF"
	pduTypeAReleaseRQName   = "A-RELEASE-RQ"
	pduTypeAReleaseRPName   = "A-RELEASE-RP"
	pduTypeAAbortName       = "A-ABORT"

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

	// MaxPDUDataLength is the hard safety cap for inbound PDU payloads.
	MaxPDUDataLength uint32 = 100 * 1024 * 1024
)

// pduDataChunkSize is the allocation/read granularity used for large PDU
// payloads. Payloads up to this size are read in a single allocation; larger
// payloads grow in chunks so a peer that declares a huge PDU length but sends
// little or no data cannot force a large up-front allocation.
const pduDataChunkSize = 1 << 20 // 1 MiB

// TypeString returns the string representation of a PDU type.
func TypeString(pduType byte) string {
	switch pduType {
	case TypeAAssociateRQ:
		return pduTypeAAssociateRQName
	case TypeAAssociateAC:
		return pduTypeAAssociateACName
	case TypeAAssociateRJ:
		return pduTypeAAssociateRJName
	case TypePDataTF:
		return pduTypePDataTFName
	case TypeAReleaseRQ:
		return pduTypeAReleaseRQName
	case TypeAReleaseRP:
		return pduTypeAReleaseRPName
	case TypeAAbort:
		return pduTypeAAbortName
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
	if uint64(len(data)) > math.MaxUint32 {
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
// The payload buffer grows incrementally, so a peer that declares a large PDU
// but sends little or no data cannot force a large up-front allocation.
func (p *RawPDU) Read(r io.Reader) error {
	// Read 6-byte header
	var header [6]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return fmt.Errorf("reading PDU header: %w", err)
	}

	p.Type = header[0]
	p.Reserved = header[1]
	p.Length = binary.BigEndian.Uint32(header[2:6])
	if p.Length > MaxPDUDataLength {
		return fmt.Errorf("PDU length %d exceeds maximum allowed %d", p.Length, MaxPDUDataLength)
	}

	// Validate PDU type
	if !isValidPDUType(p.Type) {
		return fmt.Errorf("invalid PDU type: 0x%02X", p.Type)
	}

	// Read PDU data
	data, err := ReadPDUData(r, p.Length)
	if err != nil {
		return fmt.Errorf("reading PDU data (length=%d): %w", p.Length, err)
	}
	p.Data = data

	return nil
}

// ReadPDUData reads exactly length payload bytes from r.
//
// Payloads up to 1 MiB are read into a single allocation. Larger payloads are
// read in 1 MiB chunks and only appended once a chunk has been received, so an
// announced length that is never delivered holds at most about one chunk of
// memory instead of the full announced size.
func ReadPDUData(r io.Reader, length uint32) ([]byte, error) {
	if length == 0 {
		return nil, nil
	}
	if length <= pduDataChunkSize {
		data := make([]byte, length)
		if _, err := io.ReadFull(r, data); err != nil {
			return nil, err
		}
		return data, nil
	}

	chunk := make([]byte, pduDataChunkSize)
	data := make([]byte, 0, pduDataChunkSize)
	for remaining := length; remaining > 0; {
		size := remaining
		if size > pduDataChunkSize {
			size = pduDataChunkSize
		}
		if _, err := io.ReadFull(r, chunk[:size]); err != nil {
			return nil, err
		}
		data = append(data, chunk[:size]...)
		remaining -= size
	}
	return data, nil
}

// Write writes the RawPDU to an io.Writer.
// It writes the 6-byte header followed by the PDU data.
func (p *RawPDU) Write(w io.Writer) error {
	// Update length from data
	if uint64(len(p.Data)) > math.MaxUint32 {
		return fmt.Errorf("raw PDU data too large: %d", len(p.Data))
	}
	p.Length = uint32(len(p.Data)) // #nosec G115 -- bounded by check above

	// Build 6-byte header using fixed-size array to avoid index warnings
	var header [6]byte
	header[0] = p.Type     // #nosec G602 -- fixed-size header array [6]byte
	header[1] = p.Reserved // #nosec G602 -- fixed-size header array [6]byte
	binary.BigEndian.PutUint32(header[2:6], p.Length)

	// Fast path: a single writev call on a TCP connection avoids two syscalls
	// per PDU. Only *net.TCPConn is used: it never short-writes without an
	// error, unlike an arbitrary io.Writer.
	if conn, ok := w.(*net.TCPConn); ok && len(p.Data) > 0 {
		buffers := net.Buffers{header[:], p.Data}
		n, err := buffers.WriteTo(conn)
		if err != nil {
			return fmt.Errorf("writing PDU: %w", err)
		}
		if n != int64(len(header)+len(p.Data)) {
			return fmt.Errorf("writing PDU: wrote %d of %d bytes", n, len(header)+len(p.Data))
		}
		return nil
	}

	// Write header
	if err := writeFull(w, header[:]); err != nil {
		return fmt.Errorf("writing PDU header: %w", err)
	}

	// Write data
	if len(p.Data) > 0 {
		if err := writeFull(w, p.Data); err != nil {
			return fmt.Errorf("writing PDU data: %w", err)
		}
	}

	return nil
}

func writeFull(w io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		if n <= 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
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
