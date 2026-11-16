// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package transport

import (
	"fmt"
	"net"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// ErrPDULengthExceeded indicates that a peer sent a PDU larger than the active limit.
var ErrPDULengthExceeded = pdu.ErrPDULengthExceeded

// ErrUnrecognizedPDUType indicates that a peer sent a PDU whose type byte is
// not one of the seven DICOM Upper Layer PDU types. Callers can use errors.Is
// to react by sending an A-ABORT with reason unrecognized-PDU (PS3.8 Table 9-26).
var ErrUnrecognizedPDUType = pdu.ErrUnrecognizedPDUType

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
	return ReadPDUMaxLength(conn, timeout, pdu.MaxPDUDataLength)
}

// ReadPDUMaxLength reads a PDU while applying the caller's negotiated payload limit.
// A zero maxLength keeps the transport safety cap but does not add a negotiated limit.
func ReadPDUMaxLength(conn net.Conn, timeout time.Duration, maxLength uint32) (*pdu.RawPDU, error) {
	// Set read deadline if timeout is specified
	if timeout > 0 {
		deadline := time.Now().Add(timeout)
		if err := conn.SetReadDeadline(deadline); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
		// Clear deadline after reading
		defer func() { _ = conn.SetReadDeadline(time.Time{}) }()
	}

	rawPDU := &pdu.RawPDU{}
	if err := rawPDU.ReadMax(conn, maxLength); err != nil {
		return nil, err
	}
	return rawPDU, nil
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

	return p.Write(conn)
}
