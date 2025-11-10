// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// EncodeDIMSEMessage encodes a DIMSE message into command and data datasets.
// Returns:
//   - commandData: Encoded command dataset (always present)
//   - datasetData: Encoded data dataset (may be nil if no dataset)
//   - error: Any encoding error
//
// The command dataset is always encoded in Implicit VR Little Endian.
// The data dataset is encoded according to the transfer syntax.
func EncodeDIMSEMessage(msg dimse.Message, transferSyntax *transfer.Syntax) (commandData, datasetData []byte, err error) {
	// Get command dataset
	commandDS := msg.CommandDataset()
	if commandDS == nil {
		return nil, nil, fmt.Errorf("message has no command dataset")
	}

	// Encode command dataset (always Implicit VR Little Endian)
	commandBuf := &bytes.Buffer{}
	if err := writer.Write(commandBuf, commandDS,
		writer.WithTransferSyntax(transfer.ImplicitVRLittleEndian),
		writer.WithoutPreamble()); err != nil {
		return nil, nil, fmt.Errorf("failed to encode command dataset: %w", err)
	}
	commandData = commandBuf.Bytes()

	// Get data dataset (may be nil)
	dataDS := msg.DataDataset()
	if dataDS != nil {
		// Use the negotiated transfer syntax for data dataset
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

	return commandData, datasetData, nil
}

// DecodeDIMSEMessage decodes command and data bytes into datasets.
// This function parses raw DICOM datasets without file preamble or file meta information.
func DecodeDIMSEMessage(commandData, datasetData []byte, transferSyntax *transfer.Syntax) (*dataset.Dataset, *dataset.Dataset, error) {
	// Decode command dataset (always Implicit VR Little Endian per DICOM standard)
	commandDS, err := decodeRawDataset(commandData, transfer.ImplicitVRLittleEndian)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode command dataset: %w", err)
	}

	// Decode data dataset if present
	var dataDS *dataset.Dataset
	if len(datasetData) > 0 {
		if transferSyntax == nil {
			transferSyntax = transfer.ExplicitVRLittleEndian
		}

		dataDS, err = decodeRawDataset(datasetData, transferSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode data dataset: %w", err)
		}
	}

	return commandDS, dataDS, nil
}

// decodeRawDataset decodes a raw DICOM dataset from bytes.
// This is a simplified decoder that reads DICOM elements directly without
// expecting file preamble or file meta information.
//
// TODO: Implement full raw dataset parsing with proper element-by-element reading.
// This will be completed in a future phase when we need to actually receive and parse messages.
// For now, the recvLoop is functional but cannot fully decode received datasets.
func decodeRawDataset(data []byte, ts *transfer.Syntax) (*dataset.Dataset, error) {
	if len(data) == 0 {
		return dataset.New(), nil
	}

	// Placeholder implementation - return empty dataset
	// TODO: Implement proper DICOM element parsing:
	// 1. Read tag (group, element)
	// 2. Read VR (explicit) or lookup VR (implicit)
	// 3. Read length
	// 4. Read value
	// 5. Create appropriate element type
	// 6. Add to dataset
	// 7. Repeat until EOF

	_ = ts // Will be used for determining VR mode and byte order

	return dataset.New(), fmt.Errorf("raw dataset decoding not yet fully implemented - will be completed in future phase")
}
