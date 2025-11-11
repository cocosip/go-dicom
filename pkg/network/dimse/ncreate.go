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

const (
	noAttributesMsg   = "no attributes"
	withAttributesMsg = "with attributes"
)

// NCreateRequest represents an N-CREATE-RQ message.
// N-CREATE is used to request the creation of a new managed SOP Instance.
type NCreateRequest struct {
	*BaseRequest
	affectedSOPClassUID    string
	affectedSOPInstanceUID string // Optional - may be empty if server assigns UID
}

// NewNCreateRequest creates a new N-CREATE-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - affectedSOPClassUID: The SOP Class UID of the object to be created
//   - affectedSOPInstanceUID: Optional SOP Instance UID (empty string if server should assign)
//   - attributeList: Dataset containing the initial attribute values for the new instance
func NewNCreateRequest(
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NCreateRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNCreateRQ), 0)

	// Set affected SOP Class UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))

	// Set affected SOP Instance UID if provided
	if affectedSOPInstanceUID != "" {
		_ = command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))
	}

	// CommandDataSetType
	if attributeList != nil {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NCreateRequest{
		BaseRequest:            NewBaseRequest(command, attributeList),
		affectedSOPClassUID:    affectedSOPClassUID,
		affectedSOPInstanceUID: affectedSOPInstanceUID,
	}
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NCreateRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
// Returns empty string if the server should assign the UID.
func (r *NCreateRequest) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// HasAttributeList returns true if this request contains an attribute list dataset.
func (r *NCreateRequest) HasAttributeList() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NCreateRequest) String() string {
	instanceInfo := "server-assigned UID"
	if r.affectedSOPInstanceUID != "" {
		instanceInfo = fmt.Sprintf("UID=%s", r.affectedSOPInstanceUID)
	}
	hasAttrs := noAttributesMsg
	if r.HasAttributeList() {
		hasAttrs = withAttributesMsg
	}
	return fmt.Sprintf("N-CREATE-RQ [MessageID=%d, SOP Class=%s, %s, %s]",
		r.MessageID(), r.affectedSOPClassUID, instanceInfo, hasAttrs)
}

// NCreateResponse represents an N-CREATE-RSP message.
type NCreateResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	messageIDBeingRespondedTo uint16
}

// NewNCreateResponse creates a new N-CREATE-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID of the created object
//   - attributeList: Optional dataset containing attribute values of the created instance
func NewNCreateResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NCreateResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandNCreateRSP), 0)

	// Set affected SOP Class UID and Instance UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))
	if affectedSOPInstanceUID != "" {
		_ = command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))
	}

	// MessageIDBeingRespondedTo
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType
	if attributeList != nil && statusCode == 0x0000 {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NCreateResponse{
		BaseResponse:              NewBaseResponse(command, attributeList),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNCreateResponseSuccess creates a successful N-CREATE-RSP message.
func NewNCreateResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
) *NCreateResponse {
	return NewNCreateResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID,
		affectedSOPInstanceUID, attributeList)
}

// StatusCode returns the status code.
func (r *NCreateResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NCreateResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NCreateResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// HasAttributeList returns true if this response contains an attribute list dataset.
func (r *NCreateResponse) HasAttributeList() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NCreateResponse) String() string {
	st := r.Status()
	hasData := noAttributesMsg
	if r.HasAttributeList() {
		hasData = withAttributesMsg
	}
	return fmt.Sprintf("N-CREATE-RSP [MessageID=%d, Status=%s (0x%04X), SOP Instance=%s, %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, r.affectedSOPInstanceUID, hasData)
}
