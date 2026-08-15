// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
)

// EncodeDIMSEMessage encodes a DIMSE message into command and data datasets.
// The command dataset is always encoded in Implicit VR Little Endian. The data
// dataset is encoded using the negotiated transfer syntax.
func EncodeDIMSEMessage(msg dimse.Message, transferSyntax *transfer.Syntax) (commandData, datasetData []byte, err error) {
	commandDS := msg.CommandDataset()
	if commandDS == nil {
		return nil, nil, fmt.Errorf("message has no command dataset")
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
	commandData = commandBuf.Bytes()

	dataDS := msg.DataDataset()
	if dataDS == nil {
		return commandData, nil, nil
	}
	if transferSyntax == nil {
		transferSyntax = transfer.ExplicitVRLittleEndian
	}
	dataBuf := &bytes.Buffer{}
	if err := writer.Write(dataBuf, dataDS,
		writer.WithTransferSyntax(transferSyntax),
		writer.WithoutPreamble()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode data dataset: %w", err)
	}
	return commandData, dataBuf.Bytes(), nil
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
