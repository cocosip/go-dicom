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

// NGetRequest represents an N-GET-RQ message.
// N-GET is used to retrieve attribute values from a managed SOP Instance.
type NGetRequest struct {
	*BaseRequest
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	attributeIdentifierList []*tag.Tag // Optional list of specific attributes to retrieve
}

// NewNGetRequest creates a new N-GET-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - requestedSOPClassUID: The SOP Class UID of the managed object
//   - requestedSOPInstanceUID: The SOP Instance UID of the managed object
//   - attributeIdentifierList: Optional list of attribute tags to retrieve (nil = all attributes)
func NewNGetRequest(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	attributeIdentifierList []*tag.Tag,
) *NGetRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNGetRQ), 0)

	// Set requested SOP Class UID and Instance UID
	command.Add(element.NewString(tag.RequestedSOPClassUID, vr.UI, []string{requestedSOPClassUID}))
	command.Add(element.NewString(tag.RequestedSOPInstanceUID, vr.UI, []string{requestedSOPInstanceUID}))

	// Set attribute identifier list if provided
	if len(attributeIdentifierList) > 0 {
		command.Add(element.NewAttributeTag(tag.AttributeIdentifierList, attributeIdentifierList))
	}

	// CommandDataSetType - no dataset in request
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &NGetRequest{
		BaseRequest:             NewBaseRequest(command, nil),
		requestedSOPClassUID:    requestedSOPClassUID,
		requestedSOPInstanceUID: requestedSOPInstanceUID,
		attributeIdentifierList: attributeIdentifierList,
	}
}

// RequestedSOPClassUID returns the requested SOP Class UID.
func (r *NGetRequest) RequestedSOPClassUID() string {
	return r.requestedSOPClassUID
}

// RequestedSOPInstanceUID returns the requested SOP Instance UID.
func (r *NGetRequest) RequestedSOPInstanceUID() string {
	return r.requestedSOPInstanceUID
}

// AttributeIdentifierList returns the list of requested attributes.
// Returns nil if all attributes are requested.
func (r *NGetRequest) AttributeIdentifierList() []*tag.Tag {
	return r.attributeIdentifierList
}

// String returns a human-readable representation.
func (r *NGetRequest) String() string {
	attrInfo := "all attributes"
	if len(r.attributeIdentifierList) > 0 {
		attrInfo = fmt.Sprintf("%d specific attributes", len(r.attributeIdentifierList))
	}
	return fmt.Sprintf("N-GET-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s, %s]",
		r.MessageID(), r.requestedSOPClassUID, r.requestedSOPInstanceUID, attrInfo)
}

// NGetResponse represents an N-GET-RSP message.
type NGetResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	messageIDBeingRespondedTo uint16
}

// NewNGetResponse creates a new N-GET-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID
//   - attributeList: Dataset containing the requested attribute values
func NewNGetResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NGetResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandNGetRSP), 0)

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

	return &NGetResponse{
		BaseResponse:              NewBaseResponse(command, attributeList),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNGetResponseSuccess creates a successful N-GET-RSP message.
func NewNGetResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NGetResponse {
	return NewNGetResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID,
		affectedSOPInstanceUID, attributeList)
}

// StatusCode returns the status code.
func (r *NGetResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NGetResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NGetResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// HasAttributeList returns true if this response contains an attribute list dataset.
func (r *NGetResponse) HasAttributeList() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NGetResponse) String() string {
	st := r.Status()
	hasData := "no attributes"
	if r.HasAttributeList() {
		hasData = "with attributes"
	}
	return fmt.Sprintf("N-GET-RSP [MessageID=%d, Status=%s (0x%04X), %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, hasData)
}
