// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// ReadPDU reads a complete PDU from the connection with an optional timeout.
//
// This function:
//  1. Sets a read deadline if timeout > 0
//  2. Reads the PDU header (6 bytes: type + reserved + length)
//  3. Reads the PDU data based on the length field
//  4. Returns a RawPDU containing the complete PDU
//
// Example:
//
//	pduData, err := transport.ReadPDU(conn, 30*time.Second)
//	if err != nil {
//	    return err
//	}
//
//	switch pduData.Type {
//	case pdu.TypeAssociateRQ:
//	    // Handle A-ASSOCIATE-RQ
//	case pdu.TypePDataTF:
//	    // Handle P-DATA-TF
//	}
func ReadPDU(conn net.Conn, timeout time.Duration) (*pdu.RawPDU, error) {
	// Set read deadline if timeout is specified
	if timeout > 0 {
		deadline := time.Now().Add(timeout)
		if err := conn.SetReadDeadline(deadline); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
		// Clear deadline after reading
		defer func() { _ = conn.SetReadDeadline(time.Time{}) }()
	}

	// Read PDU header (6 bytes)
	// Format: [type:1][reserved:1][length:4]
	header := make([]byte, 6)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, fmt.Errorf("failed to read PDU header: %w", err)
	}

	// Parse header
	pduType := header[0]
	// reserved := header[1]  // Reserved byte, not used
	length := binary.BigEndian.Uint32(header[2:6])

	// Validate length (reasonable upper bound to prevent memory exhaustion)
	// DICOM standard doesn't specify a hard limit, but 100MB seems reasonable
	const maxPDULength = 100 * 1024 * 1024 // 100 MB
	if length > maxPDULength {
		return nil, fmt.Errorf("PDU length %d exceeds maximum allowed %d", length, maxPDULength)
	}

	// Read PDU data
	data := make([]byte, length)
	if length > 0 {
		if _, err := io.ReadFull(conn, data); err != nil {
			return nil, fmt.Errorf("failed to read PDU data: %w", err)
		}
	}

	return &pdu.RawPDU{
		Type: pduType,
		Data: data,
	}, nil
}

// WritePDU writes a complete PDU to the connection with an optional timeout.
//
// This function:
//  1. Sets a write deadline if timeout > 0
//  2. Encodes the PDU (type + reserved + length + data)
//  3. Writes the complete PDU to the connection
//
// Example:
//
//	pduData := &pdu.RawPDU{
//	    Type: pdu.TypeAssociateRQ,
//	    Data: encodedData,
//	}
//	err := transport.WritePDU(conn, 30*time.Second, pduData)
func WritePDU(conn net.Conn, timeout time.Duration, p *pdu.RawPDU) error {
	// Set write deadline if timeout is specified
	if timeout > 0 {
		deadline := time.Now().Add(timeout)
		if err := conn.SetWriteDeadline(deadline); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
		// Clear deadline after writing
		defer func() { _ = conn.SetWriteDeadline(time.Time{}) }()
	}

	// Encode PDU header
	// Format: [type:1][reserved:1][length:4]
	length := uint32(len(p.Data))
	header := make([]byte, 6)
	header[0] = p.Type
	header[1] = 0 // Reserved
	binary.BigEndian.PutUint32(header[2:6], length)

	// Write header
	if _, err := conn.Write(header); err != nil {
		return fmt.Errorf("failed to write PDU header: %w", err)
	}

	// Write data
	if length > 0 {
		if _, err := conn.Write(p.Data); err != nil {
			return fmt.Errorf("failed to write PDU data: %w", err)
		}
	}

	return nil
}
