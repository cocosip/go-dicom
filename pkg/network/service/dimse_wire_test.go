// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"fmt"
	"net"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// writeDIMSEPDU encodes msg and writes it to conn as P-DATA-TF PDUs on
// contextID. maxPDULength 0 keeps the message in one PDV, which is enough for
// the small messages used by the wire-level tests.
func writeDIMSEPDU(conn net.Conn, msg dimse.Message, contextID byte, ts *transfer.Syntax) error {
	command, data, err := EncodeDIMSEMessage(msg, ts)
	if err != nil {
		return fmt.Errorf("encode %T: %w", msg, err)
	}
	pdvs := FragmentData(command, contextID, true, 0)
	if len(data) > 0 {
		pdvs = append(pdvs, FragmentData(data, contextID, false, 0)...)
	}
	raw, err := CreatePDataTFPDU(pdvs)
	if err != nil {
		return fmt.Errorf("create P-DATA-TF PDU: %w", err)
	}
	if err := transport.WritePDU(conn, 2*time.Second, raw); err != nil {
		return fmt.Errorf("write P-DATA-TF PDU: %w", err)
	}
	return nil
}

// readNextDIMSECommand reads the next PDU, requires it to be a P-DATA-TF with a
// command PDV and returns the decoded command set (Implicit VR Little Endian,
// PS3.7 6.3.1).
func readNextDIMSECommand(conn net.Conn, timeout time.Duration) (*dataset.Dataset, error) {
	raw, err := transport.ReadPDU(conn, timeout)
	if err != nil {
		return nil, err
	}
	if raw.Type != pdu.TypePDataTF {
		return nil, fmt.Errorf("PDU type = 0x%02X, want P-DATA-TF", raw.Type)
	}
	pdvs, err := ParsePDataTFPDU(raw)
	if err != nil {
		return nil, fmt.Errorf("parse P-DATA-TF PDU: %w", err)
	}
	var commandData []byte
	for _, pdv := range pdvs {
		if pdv.IsCommand {
			commandData = append(commandData, pdv.Data...)
		}
	}
	if len(commandData) == 0 {
		return nil, fmt.Errorf("P-DATA-TF PDU carries no command PDV")
	}
	command, _, err := DecodeDIMSEMessage(commandData, nil, transfer.ImplicitVRLittleEndian)
	if err != nil {
		return nil, fmt.Errorf("decode command set: %w", err)
	}
	return command, nil
}
