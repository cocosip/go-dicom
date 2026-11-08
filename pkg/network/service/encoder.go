// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"bytes"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// EncodeDIMSEMessage encodes a DIMSE message into command and data datasets.
// The command dataset is always encoded in Implicit VR Little Endian. The data
// dataset is encoded using the negotiated transfer syntax.
//
// CommandDataSetType (0000,0800) is derived from the bytes actually produced:
// it is 0x0101 when no data dataset is sent and 0x0001 otherwise. This keeps the
// command set and the P-DATA stream consistent (PS3.7 6.3, PS3.8 Annex E) even
// when a constructor set a flag that does not match its dataset.
func EncodeDIMSEMessage(msg dimse.Message, transferSyntax *transfer.Syntax) (commandData, datasetData []byte, err error) {
	commandDS := msg.CommandDataset()
	if commandDS == nil {
		return nil, nil, fmt.Errorf("message has no command dataset")
	}
	if err := validateRequiredDataset(msg); err != nil {
		return nil, nil, err
	}
	if err := validateCEchoResponseStatus(msg); err != nil {
		return nil, nil, err
	}
	if _, hasPriority := commandDS.Get(tag.Priority); hasPriority {
		priority, priorityErr := commandDS.GetUInt16(tag.Priority, 0)
		if priorityErr != nil {
			return nil, nil, fmt.Errorf("invalid DICOM priority: %w", priorityErr)
		}
		if !dimse.IsValidPriority(priority) {
			return nil, nil, fmt.Errorf("invalid DICOM priority 0x%04X", priority)
		}
	}

	dataDS := msg.DataDataset()
	if dataDS != nil {
		if transferSyntax == nil {
			transferSyntax = transfer.ExplicitVRLittleEndian
		}
		dataBuf := &bytes.Buffer{}
		if err := writer.Write(dataBuf, dataDS,
			writer.WithTransferSyntax(transferSyntax),
			writer.WithoutPreamble()); err != nil {
			return nil, nil, fmt.Errorf("failed to encode data dataset: %w", err)
		}
		datasetData = dataBuf.Bytes()
	}

	dataSetType := uint16(0x0101)
	if len(datasetData) > 0 {
		dataSetType = 0x0001
	}
	if err := commandDS.AddOrUpdate(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{dataSetType})); err != nil {
		return nil, nil, fmt.Errorf("failed to update CommandDataSetType: %w", err)
	}

	// Encode once to calculate the bytes following CommandGroupLength, then
	// encode again with the correct value.
	tempBuf := &bytes.Buffer{}
	if err := writer.Write(tempBuf, commandDS,
		writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian),
		writer.WithoutPreamble(),
		writer.WithKeepGroupLengths()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode command dataset for length calculation: %w", err)
	}
	commandGroupLength := uint32(len(tempBuf.Bytes()) - 12) // #nosec G115 -- command datasets include this 12-byte element
	if err := commandDS.AddOrUpdate(element.NewUnsignedLong(tag.CommandGroupLength, []uint32{commandGroupLength})); err != nil {
		return nil, nil, fmt.Errorf("failed to update CommandGroupLength: %w", err)
	}

	commandBuf := &bytes.Buffer{}
	if err := writer.Write(commandBuf, commandDS,
		writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian),
		writer.WithoutPreamble(),
		writer.WithKeepGroupLengths()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode command dataset: %w", err)
	}
	return commandBuf.Bytes(), datasetData, nil
}

// DecodeDIMSEMessage decodes raw command and data Dataset bytes. Command bytes
// always use Implicit VR Little Endian; data bytes use the negotiated syntax.
func DecodeDIMSEMessage(commandData, datasetData []byte, transferSyntax *transfer.Syntax) (*dataset.Dataset, *dataset.Dataset, error) {
	commandDS, err := decodeRawDataset(commandData, transfer.ImplicitVRLittleEndian)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode command dataset: %w", err)
	}

	var dataDS *dataset.Dataset
	if len(datasetData) > 0 {
		if transferSyntax == nil {
			transferSyntax = transfer.ExplicitVRLittleEndian
		}
		dataDS, err = decodeRawDataset(datasetData, transferSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode data dataset: %w", err)
		}
		dataDS.SetInternalTransferSyntax(transferSyntax)
	}
	return commandDS, dataDS, nil
}

// validateRequiredDataset enforces that DIMSE requests whose data set is
// mandatory (PS3.7 §9.3) actually carry one, so the encoder never advertises
// "no data set" for a command that requires it.
func validateRequiredDataset(msg dimse.Message) error {
	switch dimse.CommandField(msg.CommandField()) {
	case dimse.CommandCStoreRQ, dimse.CommandCFindRQ, dimse.CommandCGetRQ,
		dimse.CommandCMoveRQ, dimse.CommandNSetRQ:
		if msg.DataDataset() == nil {
			return fmt.Errorf("%s requires a data set", dimse.CommandField(msg.CommandField()))
		}
	case dimse.CommandCFindRSP:
		// A Pending C-FIND-RSP must carry an Identifier (PS3.7 9.1.2.1.5).
		code, err := msg.CommandDataset().GetUInt16(tag.Status, 0)
		if err == nil && code == status.CFindPending.Code && msg.DataDataset() == nil {
			return fmt.Errorf("C-FIND-RSP Pending requires a data set")
		}
	}
	return nil
}

// validateCEchoResponseStatus rejects an outgoing C-ECHO-RSP whose Status is not
// Success. PS3.7 9.3.5.2 (Table 9.3-13): the C-ECHO-RSP Status "shall have a
// value of Success". NewCEchoResponse keeps a non-Success status only so a
// non-conformant peer response can be represented for diagnostics; the wire
// encoder is the rejection point.
func validateCEchoResponseStatus(msg dimse.Message) error {
	if dimse.CommandField(msg.CommandField()) != dimse.CommandCEchoRSP {
		return nil
	}
	code, err := msg.CommandDataset().GetUInt16(tag.Status, 0)
	if err != nil {
		return fmt.Errorf("C-ECHO-RSP has no Status element: %w", err)
	}
	if code != status.Success.Code {
		return fmt.Errorf("C-ECHO-RSP status 0x%04X is not allowed: PS3.7 9.3.5.2 requires Success", code)
	}
	return nil
}

func decodeRawDataset(data []byte, ts *transfer.Syntax) (*dataset.Dataset, error) {
	if len(data) == 0 {
		return dataset.New(), nil
	}
	if ts == nil {
		return nil, fmt.Errorf("transfer syntax cannot be nil")
	}
	result, err := parser.Parse(bytes.NewReader(data),
		parser.WithAssumedTransferSyntax(ts),
		parser.WithReadOption(parser.ReadAll),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse raw Dataset: %w", err)
	}
	return result.Dataset, nil
}
