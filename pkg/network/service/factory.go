// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
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
		return dimse.NewCCancelRequest(commandDS), nil

	default:
		return nil, fmt.Errorf("unsupported command type: 0x%04X", commandField)
	}
}

func requiredString(commandDS *dataset.Dataset, fieldName string, fieldTag *tag.Tag) (string, error) {
	value, ok := commandDS.GetString(fieldTag)
	if !ok {
		return "", fmt.Errorf("%s not found", fieldName)
	}
	return value, nil
}

func requiredUInt16(commandDS *dataset.Dataset, fieldName string, fieldTag *tag.Tag) (uint16, error) {
	value, err := commandDS.GetUInt16(fieldTag, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to get %s: %w", fieldName, err)
	}
	return value, nil
}

func optionalUInt16(commandDS *dataset.Dataset, fieldTag *tag.Tag) (uint16, error) {
	if !commandDS.Contains(fieldTag) {
		return 0, nil
	}
	return commandDS.GetUInt16(fieldTag, 0)
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
	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	return dimse.NewCEchoResponse(messageID, status.LookupStatus(statusCode)), nil
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
	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
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

	return dimse.NewCStoreResponse(messageID, status.LookupStatus(statusCode), sopClassUID, sopInstanceUID), nil
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
	sopClassUID, err := requiredString(commandDS, "AffectedSOPClassUID", tag.AffectedSOPClassUID)
	if err != nil {
		return nil, err
	}

	// Create request
	var req *dimse.CFindRequest
	switch sopClassUID {
	case "1.2.840.10008.5.1.4.1.2.1.1":
		req = dimse.NewCFindRequestPatientRoot(level, dataDS)
	case "1.2.840.10008.5.1.4.1.2.2.1":
		req = dimse.NewCFindRequestStudyRoot(level, dataDS)
	default:
		req = dimse.NewCFindRequest(level, dataDS)
	}
	if err := req.SetMessageID(messageID); err != nil {
		return nil, fmt.Errorf("failed to set MessageID: %w", err)
	}
	return req, nil
}

func createCGetRequest(commandDS, dataDS *dataset.Dataset) (*dimse.CGetRequest, error) {
	if dataDS == nil {
		return nil, fmt.Errorf("C-GET-RQ requires data dataset (identifier)")
	}

	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	levelStr, ok := dataDS.GetString(tag.QueryRetrieveLevel)
	if !ok {
		return nil, fmt.Errorf("QueryRetrieveLevel not found in identifier")
	}
	level := dimse.QueryRetrieveLevel(levelStr)
	sopClassUID, err := requiredString(commandDS, "AffectedSOPClassUID", tag.AffectedSOPClassUID)
	if err != nil {
		return nil, err
	}

	var req *dimse.CGetRequest
	switch sopClassUID {
	case "1.2.840.10008.5.1.4.1.2.1.3":
		req = dimse.NewCGetRequestPatientRoot(level, dataDS)
	case "1.2.840.10008.5.1.4.1.2.2.3":
		req = dimse.NewCGetRequestStudyRoot(level, dataDS)
	default:
		req = dimse.NewCGetRequest(level, dataDS)
	}
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

	if statusCode == status.CGetPending.Code {
		remaining, err := requiredUInt16(commandDS, "NumberOfRemainingSuboperations", tag.NumberOfRemainingSuboperations)
		if err != nil {
			return nil, err
		}
		completed, err := requiredUInt16(commandDS, "NumberOfCompletedSuboperations", tag.NumberOfCompletedSuboperations)
		if err != nil {
			return nil, err
		}
		failed, err := requiredUInt16(commandDS, "NumberOfFailedSuboperations", tag.NumberOfFailedSuboperations)
		if err != nil {
			return nil, err
		}
		warning, err := requiredUInt16(commandDS, "NumberOfWarningSuboperations", tag.NumberOfWarningSuboperations)
		if err != nil {
			return nil, err
		}
		return dimse.NewCGetResponsePending(messageID, sopClassUID, remaining, completed, failed, warning), nil
	}
	return dimse.NewCGetResponse(messageID, status.LookupStatus(statusCode), sopClassUID), nil
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

	moveDestination, err := requiredString(commandDS, "MoveDestination", tag.MoveDestination)
	if err != nil {
		return nil, err
	}
	levelStr, ok := dataDS.GetString(tag.QueryRetrieveLevel)
	if !ok {
		return nil, fmt.Errorf("QueryRetrieveLevel not found in identifier")
	}
	level := dimse.QueryRetrieveLevel(levelStr)
	sopClassUID, err := requiredString(commandDS, "AffectedSOPClassUID", tag.AffectedSOPClassUID)
	if err != nil {
		return nil, err
	}

	var req *dimse.CMoveRequest
	switch sopClassUID {
	case "1.2.840.10008.5.1.4.1.2.1.2":
		req = dimse.NewCMoveRequestPatientRoot(level, moveDestination, dataDS)
	case "1.2.840.10008.5.1.4.1.2.2.2":
		req = dimse.NewCMoveRequestStudyRoot(level, moveDestination, dataDS)
	default:
		req = dimse.NewCMoveRequest(level, moveDestination, dataDS)
	}
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

	if statusCode == status.CMovePending.Code {
		remaining, err := requiredUInt16(commandDS, "NumberOfRemainingSuboperations", tag.NumberOfRemainingSuboperations)
		if err != nil {
			return nil, err
		}
		completed, err := requiredUInt16(commandDS, "NumberOfCompletedSuboperations", tag.NumberOfCompletedSuboperations)
		if err != nil {
			return nil, err
		}
		failed, err := requiredUInt16(commandDS, "NumberOfFailedSuboperations", tag.NumberOfFailedSuboperations)
		if err != nil {
			return nil, err
		}
		warning, err := requiredUInt16(commandDS, "NumberOfWarningSuboperations", tag.NumberOfWarningSuboperations)
		if err != nil {
			return nil, err
		}
		return dimse.NewCMoveResponsePending(messageID, sopClassUID, remaining, completed, failed, warning), nil
	}
	return dimse.NewCMoveResponse(messageID, status.LookupStatus(statusCode), sopClassUID), nil
}

// createCFindResponse creates a C-FIND-RSP from datasets.
func createCFindResponse(commandDS, dataDS *dataset.Dataset) (*dimse.CFindResponse, error) {
	// Get MessageIDBeingRespondedTo
	messageID, err := commandDS.GetUInt16(tag.MessageIDBeingRespondedTo, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageIDBeingRespondedTo: %w", err)
	}

	// Get Status
	statusCode, err := commandDS.GetUInt16(tag.Status, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get Status: %w", err)
	}

	// Get AffectedSOPClassUID
	sopClassUID, ok := commandDS.GetString(tag.AffectedSOPClassUID)
	if !ok {
		return nil, fmt.Errorf("AffectedSOPClassUID not found")
	}

	// dataDS is the identifier (may be nil for final response)
	return dimse.NewCFindResponse(messageID, status.LookupStatus(statusCode), sopClassUID, dataDS), nil
}

// createNEventReportRequest creates an N-EVENT-REPORT-RQ from datasets.
func createNEventReportRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NEventReportRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	affectedSOPClassUID, err := requiredString(commandDS, "AffectedSOPClassUID", tag.AffectedSOPClassUID)
	if err != nil {
		return nil, err
	}
	affectedSOPInstanceUID, err := requiredString(commandDS, "AffectedSOPInstanceUID", tag.AffectedSOPInstanceUID)
	if err != nil {
		return nil, err
	}
	eventTypeID, err := requiredUInt16(commandDS, "EventTypeID", tag.EventTypeID)
	if err != nil {
		return nil, err
	}

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
	eventTypeID, err := optionalUInt16(commandDS, tag.EventTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get EventTypeID: %w", err)
	}

	return dimse.NewNEventReportResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID, eventTypeID, dataDS), nil
}

// createNGetRequest creates an N-GET-RQ from datasets.
func createNGetRequest(commandDS *dataset.Dataset) (*dimse.NGetRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, err := requiredString(commandDS, "RequestedSOPClassUID", tag.RequestedSOPClassUID)
	if err != nil {
		return nil, err
	}
	requestedSOPInstanceUID, err := requiredString(commandDS, "RequestedSOPInstanceUID", tag.RequestedSOPInstanceUID)
	if err != nil {
		return nil, err
	}

	// AttributeIdentifierList is optional
	var attrList []*tag.Tag
	if elem, ok := commandDS.Get(tag.AttributeIdentifierList); ok {
		if atElem, ok := elem.(*element.AttributeTag); ok {
			attrList, err = atElem.GetValues()
			if err != nil {
				return nil, fmt.Errorf("failed to get AttributeIdentifierList: %w", err)
			}
		} else {
			return nil, fmt.Errorf("AttributeIdentifierList is not AttributeTag")
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

	return dimse.NewNGetResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNSetRequest creates an N-SET-RQ from datasets.
func createNSetRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NSetRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, err := requiredString(commandDS, "RequestedSOPClassUID", tag.RequestedSOPClassUID)
	if err != nil {
		return nil, err
	}
	requestedSOPInstanceUID, err := requiredString(commandDS, "RequestedSOPInstanceUID", tag.RequestedSOPInstanceUID)
	if err != nil {
		return nil, err
	}

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

	return dimse.NewNSetResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNActionRequest creates an N-ACTION-RQ from datasets.
func createNActionRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NActionRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, err := requiredString(commandDS, "RequestedSOPClassUID", tag.RequestedSOPClassUID)
	if err != nil {
		return nil, err
	}
	requestedSOPInstanceUID, err := requiredString(commandDS, "RequestedSOPInstanceUID", tag.RequestedSOPInstanceUID)
	if err != nil {
		return nil, err
	}
	actionTypeID, err := requiredUInt16(commandDS, "ActionTypeID", tag.ActionTypeID)
	if err != nil {
		return nil, err
	}

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
	actionTypeID, err := optionalUInt16(commandDS, tag.ActionTypeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ActionTypeID: %w", err)
	}

	return dimse.NewNActionResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID, actionTypeID, dataDS), nil
}

// createNCreateRequest creates an N-CREATE-RQ from datasets.
func createNCreateRequest(commandDS, dataDS *dataset.Dataset) (*dimse.NCreateRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	affectedSOPClassUID, err := requiredString(commandDS, "AffectedSOPClassUID", tag.AffectedSOPClassUID)
	if err != nil {
		return nil, err
	}
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

	return dimse.NewNCreateResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID, dataDS), nil
}

// createNDeleteRequest creates an N-DELETE-RQ from datasets.
func createNDeleteRequest(commandDS *dataset.Dataset) (*dimse.NDeleteRequest, error) {
	messageID, err := commandDS.GetUInt16(tag.MessageID, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get MessageID: %w", err)
	}

	requestedSOPClassUID, err := requiredString(commandDS, "RequestedSOPClassUID", tag.RequestedSOPClassUID)
	if err != nil {
		return nil, err
	}
	requestedSOPInstanceUID, err := requiredString(commandDS, "RequestedSOPInstanceUID", tag.RequestedSOPInstanceUID)
	if err != nil {
		return nil, err
	}

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

	return dimse.NewNDeleteResponse(messageID, status.LookupStatus(statusCode), affectedSOPClassUID, affectedSOPInstanceUID), nil
}
