// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// createMessageFromDatasets creates a DIMSE message from command and data datasets.
// This factory function examines the CommandField to determine the message type
// and constructs the appropriate request or response object.
func createMessageFromDatasets(commandDS, dataDS *dataset.Dataset) (dimse.Message, error) {
	// Get CommandField to determine message type
	commandField, err := commandDS.GetUInt16(tag.CommandField, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get CommandField: %w", err)
	}

	// Dispatch based on command type
	switch dimse.CommandField(commandField) {
	// C-ECHO
	case dimse.CommandCEchoRQ:
		return createCEchoRequest(commandDS)
	case dimse.CommandCEchoRSP:
		return createCEchoResponse(commandDS)

	// C-STORE
	case dimse.CommandCStoreRQ:
		return createCStoreRequest(commandDS, dataDS)
	case dimse.CommandCStoreRSP:
		return createCStoreResponse(commandDS)

	// C-FIND
	case dimse.CommandCFindRQ:
		return createCFindRequest(commandDS, dataDS)
	case dimse.CommandCFindRSP:
		return createCFindResponse(commandDS, dataDS)

	default:
		return nil, fmt.Errorf("unsupported command type: 0x%04X", commandField)
	}
}

// createCEchoRequest creates a C-ECHO-RQ from datasets.
func createCEchoRequest(commandDS *dataset.Dataset) (*dimse.CEchoRequest, error) {
	// Get MessageID
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	req := dimse.NewCEchoRequest()
	_ = req.SetMessageID(messageID)
	return req, nil
}

// createCEchoResponse creates a C-ECHO-RSP from datasets.
func createCEchoResponse(commandDS *dataset.Dataset) (*dimse.CEchoResponse, error) {
	// Get MessageIDBeingRespondedTo
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	// Get Status
	status, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	return dimse.NewCEchoResponse(messageID, status), nil
}

// createCStoreRequest creates a C-STORE-RQ from datasets.
func createCStoreRequest(commandDS, dataDS *dataset.Dataset) (*dimse.CStoreRequest, error) {
	if dataDS == nil {
		return nil, fmt.Errorf("C-STORE-RQ requires data dataset")
	}

	// Get MessageID
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	// Create request from data dataset
	req, err := dimse.NewCStoreRequest(dataDS)
	if err != nil {
		return nil, fmt.Errorf("failed to create C-STORE request: %w", err)
	}

	_ = req.SetMessageID(messageID)
	return req, nil
}

// createCStoreResponse creates a C-STORE-RSP from datasets.
func createCStoreResponse(commandDS *dataset.Dataset) (*dimse.CStoreResponse, error) {
	// Get MessageIDBeingRespondedTo
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	// Get Status
	status, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	// Get AffectedSOPClassUID
	sopClassUID, ok := commandDS.GetString(tag.AffectedSOPClassUID)
	if !ok {
		return nil, fmt.Errorf("AffectedSOPClassUID not found")
	}

	// Get AffectedSOPInstanceUID
	sopInstanceUID, ok := commandDS.GetString(tag.AffectedSOPInstanceUID)
	if !ok {
		return nil, fmt.Errorf("AffectedSOPInstanceUID not found")
	}

	return dimse.NewCStoreResponse(messageID, status, sopClassUID, sopInstanceUID), nil
}

// createCFindRequest creates a C-FIND-RQ from datasets.
func createCFindRequest(commandDS, dataDS *dataset.Dataset) (*dimse.CFindRequest, error) {
	if dataDS == nil {
		return nil, fmt.Errorf("C-FIND-RQ requires data dataset (query identifier)")
	}

	// Get MessageID
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	// Get QueryRetrieveLevel from data dataset
	levelStr, ok := dataDS.GetString(tag.QueryRetrieveLevel)
	if !ok {
		return nil, fmt.Errorf("QueryRetrieveLevel not found in identifier")
	}
	level := dimse.QueryRetrieveLevel(levelStr)

	// Create request
	req := dimse.NewCFindRequest(level, dataDS)
	_ = req.SetMessageID(messageID)
	return req, nil
}

// createCFindResponse creates a C-FIND-RSP from datasets.
func createCFindResponse(commandDS, dataDS *dataset.Dataset) (*dimse.CFindResponse, error) {
	// Get MessageIDBeingRespondedTo
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	// Get Status
	status, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	// Get AffectedSOPClassUID
	sopClassUID, ok := commandDS.GetString(tag.AffectedSOPClassUID)
	if !ok {
		return nil, fmt.Errorf("AffectedSOPClassUID not found")
	}

	// dataDS is the identifier (may be nil for final response)
	return dimse.NewCFindResponse(messageID, status, sopClassUID, dataDS), nil
}
