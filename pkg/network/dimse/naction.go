// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// NActionRequest represents an N-ACTION-RQ message.
// N-ACTION is used to request a managed SOP Instance to perform an action.
type NActionRequest struct {
	*BaseRequest
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	actionTypeID            uint16
}

// NewNActionRequest creates a new N-ACTION-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - requestedSOPClassUID: The SOP Class UID of the managed object
//   - requestedSOPInstanceUID: The SOP Instance UID of the managed object
//   - actionTypeID: Application-specific action type identifier
//   - actionInformation: Optional dataset containing action-specific information
func NewNActionRequest(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	actionTypeID uint16,
	actionInformation *dataset.Dataset,
) *NActionRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNActionRQ), 0)

	// Set requested SOP Class UID and Instance UID
	command.Add(element.NewString(tag.RequestedSOPClassUID, vr.UI, []string{requestedSOPClassUID}))
	command.Add(element.NewString(tag.RequestedSOPInstanceUID, vr.UI, []string{requestedSOPInstanceUID}))

	// Set action type ID
	command.Add(element.NewUnsignedShort(tag.ActionTypeID, []uint16{actionTypeID}))

	// CommandDataSetType
	if actionInformation != nil {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NActionRequest{
		BaseRequest:             NewBaseRequest(command, actionInformation),
		requestedSOPClassUID:    requestedSOPClassUID,
		requestedSOPInstanceUID: requestedSOPInstanceUID,
		actionTypeID:            actionTypeID,
	}
}

// RequestedSOPClassUID returns the requested SOP Class UID.
func (r *NActionRequest) RequestedSOPClassUID() string {
	return r.requestedSOPClassUID
}

// RequestedSOPInstanceUID returns the requested SOP Instance UID.
func (r *NActionRequest) RequestedSOPInstanceUID() string {
	return r.requestedSOPInstanceUID
}

// ActionTypeID returns the action type identifier.
func (r *NActionRequest) ActionTypeID() uint16 {
	return r.actionTypeID
}

// HasActionInformation returns true if this request contains action information dataset.
func (r *NActionRequest) HasActionInformation() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NActionRequest) String() string {
	hasInfo := "no action info"
	if r.HasActionInformation() {
		hasInfo = "with action info"
	}
	return fmt.Sprintf("N-ACTION-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s, ActionTypeID=%d, %s]",
		r.MessageID(), r.requestedSOPClassUID, r.requestedSOPInstanceUID, r.actionTypeID, hasInfo)
}

// NActionResponse represents an N-ACTION-RSP message.
type NActionResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	actionTypeID              uint16
	messageIDBeingRespondedTo uint16
}

// NewNActionResponse creates a new N-ACTION-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID
//   - actionTypeID: Action type identifier from the request
//   - actionReply: Optional dataset containing action-specific reply information
func NewNActionResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	actionTypeID uint16,
	actionReply *dataset.Dataset,
) *NActionResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandNActionRSP), 0)

	// Set affected SOP Class UID and Instance UID
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))
	command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))

	// MessageIDBeingRespondedTo
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// Action type ID
	if actionTypeID != 0 {
		command.Add(element.NewUnsignedShort(tag.ActionTypeID, []uint16{actionTypeID}))
	}

	// CommandDataSetType
	if actionReply != nil {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NActionResponse{
		BaseResponse:              NewBaseResponse(command, actionReply),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		actionTypeID:              actionTypeID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNActionResponseSuccess creates a successful N-ACTION-RSP message.
func NewNActionResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	actionTypeID uint16,
	actionReply *dataset.Dataset,
) *NActionResponse {
	return NewNActionResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID,
		affectedSOPInstanceUID, actionTypeID, actionReply)
}

// StatusCode returns the status code.
func (r *NActionResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NActionResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NActionResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// ActionTypeID returns the action type identifier.
func (r *NActionResponse) ActionTypeID() uint16 {
	return r.actionTypeID
}

// HasActionReply returns true if this response contains action reply dataset.
func (r *NActionResponse) HasActionReply() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NActionResponse) String() string {
	st := r.Status()
	hasReply := "no action reply"
	if r.HasActionReply() {
		hasReply = "with action reply"
	}
	return fmt.Sprintf("N-ACTION-RSP [MessageID=%d, Status=%s (0x%04X), ActionTypeID=%d, %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, r.actionTypeID, hasReply)
}
