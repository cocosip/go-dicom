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

// NSetRequest represents an N-SET-RQ message.
// N-SET is used to modify attribute values of a managed SOP Instance.
type NSetRequest struct {
	*BaseRequest
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
}

// NewNSetRequest creates a new N-SET-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - requestedSOPClassUID: The SOP Class UID of the managed object
//   - requestedSOPInstanceUID: The SOP Instance UID of the managed object
//   - modificationList: Dataset containing the attributes to be modified with new values
func NewNSetRequest(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	modificationList *dataset.Dataset,
) *NSetRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNSetRQ), 0)

	// Set requested SOP Class UID and Instance UID
	command.Add(element.NewString(tag.RequestedSOPClassUID, vr.UI, []string{requestedSOPClassUID}))
	command.Add(element.NewString(tag.RequestedSOPInstanceUID, vr.UI, []string{requestedSOPInstanceUID}))

	// CommandDataSetType - dataset present (modification list)
	if modificationList != nil {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NSetRequest{
		BaseRequest:             NewBaseRequest(command, modificationList),
		requestedSOPClassUID:    requestedSOPClassUID,
		requestedSOPInstanceUID: requestedSOPInstanceUID,
	}
}

// RequestedSOPClassUID returns the requested SOP Class UID.
func (r *NSetRequest) RequestedSOPClassUID() string {
	return r.requestedSOPClassUID
}

// RequestedSOPInstanceUID returns the requested SOP Instance UID.
func (r *NSetRequest) RequestedSOPInstanceUID() string {
	return r.requestedSOPInstanceUID
}

// HasModificationList returns true if this request contains a modification list dataset.
func (r *NSetRequest) HasModificationList() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NSetRequest) String() string {
	hasMods := "no modifications"
	if r.HasModificationList() {
		hasMods = "with modifications"
	}
	return fmt.Sprintf("N-SET-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s, %s]",
		r.MessageID(), r.requestedSOPClassUID, r.requestedSOPInstanceUID, hasMods)
}

// NSetResponse represents an N-SET-RSP message.
type NSetResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	messageIDBeingRespondedTo uint16
}

// NewNSetResponse creates a new N-SET-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID
//   - attributeList: Optional dataset containing attribute values after modification
func NewNSetResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NSetResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandNSetRSP), 0)

	// Set affected SOP Class UID and Instance UID
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))
	command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))

	// MessageIDBeingRespondedTo
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType
	if attributeList != nil && statusCode == 0x0000 {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NSetResponse{
		BaseResponse:              NewBaseResponse(command, attributeList),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNSetResponseSuccess creates a successful N-SET-RSP message.
func NewNSetResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NSetResponse {
	return NewNSetResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID,
		affectedSOPInstanceUID, attributeList)
}

// StatusCode returns the status code.
func (r *NSetResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NSetResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NSetResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// HasAttributeList returns true if this response contains an attribute list dataset.
func (r *NSetResponse) HasAttributeList() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NSetResponse) String() string {
	st := r.Status()
	hasData := "no attributes"
	if r.HasAttributeList() {
		hasData = "with attributes"
	}
	return fmt.Sprintf("N-SET-RSP [MessageID=%d, Status=%s (0x%04X), %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, hasData)
}
