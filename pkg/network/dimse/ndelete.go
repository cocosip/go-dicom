// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// NDeleteRequest represents an N-DELETE-RQ message.
// N-DELETE is used to request the deletion of a managed SOP Instance.
type NDeleteRequest struct {
	*BaseRequest
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
}

// NewNDeleteRequest creates a new N-DELETE-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - requestedSOPClassUID: The SOP Class UID of the managed object to delete
//   - requestedSOPInstanceUID: The SOP Instance UID of the managed object to delete
func NewNDeleteRequest(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
) *NDeleteRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNDeleteRQ), 0)

	// Set requested SOP Class UID and Instance UID
	command.Add(element.NewString(tag.RequestedSOPClassUID, vr.UI, []string{requestedSOPClassUID}))
	command.Add(element.NewString(tag.RequestedSOPInstanceUID, vr.UI, []string{requestedSOPInstanceUID}))

	// CommandDataSetType - no dataset in request
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &NDeleteRequest{
		BaseRequest:             NewBaseRequest(command, nil),
		requestedSOPClassUID:    requestedSOPClassUID,
		requestedSOPInstanceUID: requestedSOPInstanceUID,
	}
}

// RequestedSOPClassUID returns the requested SOP Class UID.
func (r *NDeleteRequest) RequestedSOPClassUID() string {
	return r.requestedSOPClassUID
}

// RequestedSOPInstanceUID returns the requested SOP Instance UID.
func (r *NDeleteRequest) RequestedSOPInstanceUID() string {
	return r.requestedSOPInstanceUID
}

// String returns a human-readable representation.
func (r *NDeleteRequest) String() string {
	return fmt.Sprintf("N-DELETE-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s]",
		r.MessageID(), r.requestedSOPClassUID, r.requestedSOPInstanceUID)
}

// NDeleteResponse represents an N-DELETE-RSP message.
type NDeleteResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	messageIDBeingRespondedTo uint16
}

// NewNDeleteResponse creates a new N-DELETE-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID of the deleted object
func NewNDeleteResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
) *NDeleteResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandNDeleteRSP), 0)

	// Set affected SOP Class UID and Instance UID
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))
	command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))

	// MessageIDBeingRespondedTo
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType - no dataset in response
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &NDeleteResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNDeleteResponseSuccess creates a successful N-DELETE-RSP message.
func NewNDeleteResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
) *NDeleteResponse {
	return NewNDeleteResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID, affectedSOPInstanceUID)
}

// StatusCode returns the status code.
func (r *NDeleteResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NDeleteResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NDeleteResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// String returns a human-readable representation.
func (r *NDeleteResponse) String() string {
	st := r.Status()
	return fmt.Sprintf("N-DELETE-RSP [MessageID=%d, Status=%s (0x%04X), SOP Instance=%s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, r.affectedSOPInstanceUID)
}
