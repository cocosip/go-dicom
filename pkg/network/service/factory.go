// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
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

	// C-GET
	case dimse.CommandCGetRQ:
		return createCGetRequest(commandDS, dataDS)
	case dimse.CommandCGetRSP:
		return createCGetResponse(commandDS)

	// C-MOVE
	case dimse.CommandCMoveRQ:
		return createCMoveRequest(commandDS, dataDS)
	case dimse.CommandCMoveRSP:
		return createCMoveResponse(commandDS)

	// N-EVENT-REPORT
	case dimse.CommandNEventReportRQ:
		return createNEventReportRequest(commandDS, dataDS)
	case dimse.CommandNEventReportRSP:
		return createNEventReportResponse(commandDS, dataDS)

	// N-GET
	case dimse.CommandNGetRQ:
		return createNGetRequest(commandDS)
	case dimse.CommandNGetRSP:
		return createNGetResponse(commandDS, dataDS)

	// N-SET
	case dimse.CommandNSetRQ:
		return createNSetRequest(commandDS, dataDS)
	case dimse.CommandNSetRSP:
		return createNSetResponse(commandDS, dataDS)

	// N-ACTION
	case dimse.CommandNActionRQ:
		return createNActionRequest(commandDS, dataDS)
	case dimse.CommandNActionRSP:
		return createNActionResponse(commandDS, dataDS)

	// N-CREATE
	case dimse.CommandNCreateRQ:
		return createNCreateRequest(commandDS, dataDS)
	case dimse.CommandNCreateRSP:
		return createNCreateResponse(commandDS, dataDS)

	// N-DELETE
	case dimse.CommandNDeleteRQ:
		return createNDeleteRequest(commandDS)
	case dimse.CommandNDeleteRSP:
		return createNDeleteResponse(commandDS)

	// C-CANCEL
	case dimse.CommandCCancelRQ:
		return createCCancelRequest(commandDS)

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
    if err := req.SetMessageID(messageID); err != nil {
        return nil, fmt.Errorf("failed to set MessageID: %w", err)
    }
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

    if err := req.SetMessageID(messageID); err != nil {
        return nil, fmt.Errorf("failed to set MessageID: %w", err)
    }
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
    if err := req.SetMessageID(messageID); err != nil {
        return nil, fmt.Errorf("failed to set MessageID: %w", err)
    }
    return req, nil
}

// createCGetRequest creates a C-GET-RQ from datasets.
func createCGetRequest(commandDS, dataDS *dataset.Dataset) (*dimse.CGetRequest, error) {
	if dataDS == nil {
		return nil, fmt.Errorf("C-GET-RQ requires data dataset (identifier)")
	}

	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	levelStr, _ := dataDS.GetString(tag.QueryRetrieveLevel)
	level := dimse.QueryRetrieveLevel(levelStr)

	req := dimse.NewCGetRequest(level, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createCGetResponse creates a C-GET-RSP from datasets.
func createCGetResponse(commandDS *dataset.Dataset) (*dimse.CGetResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	sopClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)

	if statusCode == 0xFF00 {
		remaining, _ := commandDS.GetUInt16(tag.NumberOfRemainingSuboperations, 0)
		completed, _ := commandDS.GetUInt16(tag.NumberOfCompletedSuboperations, 0)
		failed, _ := commandDS.GetUInt16(tag.NumberOfFailedSuboperations, 0)
		warning, _ := commandDS.GetUInt16(tag.NumberOfWarningSuboperations, 0)
		return dimse.NewCGetResponsePending(messageID, sopClassUID, remaining, completed, failed, warning), nil
	}
	return dimse.NewCGetResponse(messageID, statusCode, sopClassUID), nil
}

// createCMoveRequest creates a C-MOVE-RQ from datasets.
func createCMoveRequest(commandDS, dataDS *dataset.Dataset) (*dimse.CMoveRequest, error) {
	if dataDS == nil {
		return nil, fmt.Errorf("C-MOVE-RQ requires data dataset (identifier)")
	}

	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	moveDestination, _ := commandDS.GetString(tag.MoveDestination)
	levelStr, _ := dataDS.GetString(tag.QueryRetrieveLevel)
	level := dimse.QueryRetrieveLevel(levelStr)

	req := dimse.NewCMoveRequest(level, moveDestination, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createCMoveResponse creates a C-MOVE-RSP from datasets.
func createCMoveResponse(commandDS *dataset.Dataset) (*dimse.CMoveResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	sopClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)

	if statusCode == 0xFF00 {
		remaining, _ := commandDS.GetUInt16(tag.NumberOfRemainingSuboperations, 0)
		completed, _ := commandDS.GetUInt16(tag.NumberOfCompletedSuboperations, 0)
		failed, _ := commandDS.GetUInt16(tag.NumberOfFailedSuboperations, 0)
		warning, _ := commandDS.GetUInt16(tag.NumberOfWarningSuboperations, 0)
		return dimse.NewCMoveResponsePending(messageID, sopClassUID, remaining, completed, failed, warning), nil
	}
	return dimse.NewCMoveResponse(messageID, statusCode, sopClassUID), nil
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

// createNEventReportRequest creates an N-EVENT-REPORT-RQ from datasets.
func createNEventReportRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NEventReportRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)
	eventTypeID, _ := commandDS.GetUInt16(tag.EventTypeID, 0)

	req := dimse.NewNEventReportRequest(affectedSOPClassUID, affectedSOPInstanceUID, eventTypeID, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNEventReportResponse creates an N-EVENT-REPORT-RSP from datasets.
func createNEventReportResponse(commandDS, dataDS *dataset.Dataset) (*dimse.NEventReportResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)
	eventTypeID, _ := commandDS.GetUInt16(tag.EventTypeID, 0)

	return dimse.NewNEventReportResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID, eventTypeID, dataDS), nil
}

// createNGetRequest creates an N-GET-RQ from datasets.
func createNGetRequest(commandDS *dataset.Dataset) (*dimse.NGetRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, _ := commandDS.GetString(tag.RequestedSOPClassUID)
	requestedSOPInstanceUID, _ := commandDS.GetString(tag.RequestedSOPInstanceUID)

	// AttributeIdentifierList is optional
	var attrList []*tag.Tag
	if elem, ok := commandDS.Get(tag.AttributeIdentifierList); ok {
		if atElem, ok := elem.(*element.AttributeTag); ok {
			attrList, _ = atElem.GetValues()
		}
	}

	req := dimse.NewNGetRequest(requestedSOPClassUID, requestedSOPInstanceUID, attrList)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNGetResponse creates an N-GET-RSP from datasets.
func createNGetResponse(commandDS, dataDS *dataset.Dataset) (*dimse.NGetResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)

	return dimse.NewNGetResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNSetRequest creates an N-SET-RQ from datasets.
func createNSetRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NSetRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, _ := commandDS.GetString(tag.RequestedSOPClassUID)
	requestedSOPInstanceUID, _ := commandDS.GetString(tag.RequestedSOPInstanceUID)

	req := dimse.NewNSetRequest(requestedSOPClassUID, requestedSOPInstanceUID, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNSetResponse creates an N-SET-RSP from datasets.
func createNSetResponse(commandDS, dataDS *dataset.Dataset) (*dimse.NSetResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)

	return dimse.NewNSetResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNActionRequest creates an N-ACTION-RQ from datasets.
func createNActionRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NActionRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, _ := commandDS.GetString(tag.RequestedSOPClassUID)
	requestedSOPInstanceUID, _ := commandDS.GetString(tag.RequestedSOPInstanceUID)
	actionTypeID, _ := commandDS.GetUInt16(tag.ActionTypeID, 0)

	req := dimse.NewNActionRequest(requestedSOPClassUID, requestedSOPInstanceUID, actionTypeID, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNActionResponse creates an N-ACTION-RSP from datasets.
func createNActionResponse(commandDS, dataDS *dataset.Dataset) (*dimse.NActionResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)
	actionTypeID, _ := commandDS.GetUInt16(tag.ActionTypeID, 0)

	return dimse.NewNActionResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID, actionTypeID, dataDS), nil
}

// createNCreateRequest creates an N-CREATE-RQ from datasets.
func createNCreateRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NCreateRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)

	req := dimse.NewNCreateRequest(affectedSOPClassUID, affectedSOPInstanceUID, dataDS)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNCreateResponse creates an N-CREATE-RSP from datasets.
func createNCreateResponse(commandDS, dataDS *dataset.Dataset) (*dimse.NCreateResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)

	return dimse.NewNCreateResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNDeleteRequest creates an N-DELETE-RQ from datasets.
func createNDeleteRequest(commandDS *dataset.Dataset) (*dimse.NDeleteRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, _ := commandDS.GetString(tag.RequestedSOPClassUID)
	requestedSOPInstanceUID, _ := commandDS.GetString(tag.RequestedSOPInstanceUID)

	req := dimse.NewNDeleteRequest(requestedSOPClassUID, requestedSOPInstanceUID)
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

// createNDeleteResponse creates an N-DELETE-RSP from datasets.
func createNDeleteResponse(commandDS *dataset.Dataset) (*dimse.NDeleteResponse, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	affectedSOPClassUID, _ := commandDS.GetString(tag.AffectedSOPClassUID)
	affectedSOPInstanceUID, _ := commandDS.GetString(tag.AffectedSOPInstanceUID)

	return dimse.NewNDeleteResponse(messageID, statusCode, affectedSOPClassUID, affectedSOPInstanceUID), nil
}

// createCCancelRequest creates a C-CANCEL-RQ from datasets.
// C-CANCEL uses MessageIDBeingRespondedTo to identify which pending operation to cancel.
func createCCancelRequest(commandDS *dataset.Dataset) (*dimse.BaseRequest, error) {
	return dimse.NewBaseRequest(commandDS, nil), nil
}
