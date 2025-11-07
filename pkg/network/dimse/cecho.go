// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// CEchoRequest represents a C-ECHO-RQ message.
// C-ECHO is used to verify DICOM connectivity (like a ping).
type CEchoRequest struct {
	*BaseRequest
	affectedSOPClassUID string
}

// NewCEchoRequest creates a new C-ECHO-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
func NewCEchoRequest() *CEchoRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandCEchoRQ), 0)

	// Set affected SOP Class UID to Verification SOP Class
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{"1.2.840.10008.1.1"}))

	// Priority (optional, default to medium)
	command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))

	// CommandDataSetType is already set to 0x0101 (no dataset)

	return &CEchoRequest{
		BaseRequest:         NewBaseRequest(command, nil),
		affectedSOPClassUID: "1.2.840.10008.1.1",
	}
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *CEchoRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// String returns a human-readable representation.
func (r *CEchoRequest) String() string {
	return fmt.Sprintf("C-ECHO-RQ [MessageID=%d, SOP Class=%s]",
		r.MessageID(), r.affectedSOPClassUID)
}

// CEchoResponse represents a C-ECHO-RSP message.
type CEchoResponse struct {
	*BaseResponse
	statusCode              uint16
	affectedSOPClassUID     string
	messageIDBeingRespondedTo uint16
}

// NewCEchoResponse creates a new C-ECHO-RSP message.
func NewCEchoResponse(messageIDBeingRespondedTo uint16, statusCode uint16) *CEchoResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCEchoRSP), 0) // Response doesn't have its own MessageID

	// Set affected SOP Class UID to Verification SOP Class
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{"1.2.840.10008.1.1"}))

	// MessageIDBeingRespondedTo
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType is already set to 0x0101 (no dataset)

	return &CEchoResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                statusCode,
		affectedSOPClassUID:       "1.2.840.10008.1.1",
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCEchoResponseSuccess creates a successful C-ECHO-RSP message.
func NewCEchoResponseSuccess(messageIDBeingRespondedTo uint16) *CEchoResponse {
	return NewCEchoResponse(messageIDBeingRespondedTo, 0x0000)
}

// NewCEchoResponseFromRequest creates a C-ECHO-RSP message from the corresponding request.
// This is a convenience function that automatically uses the request's MessageID.
//
// Example:
//
//	// When receiving a C-ECHO request
//	resp := dimse.NewCEchoResponseFromRequest(req, 0x0000) // Success
func NewCEchoResponseFromRequest(req *CEchoRequest, statusCode uint16) *CEchoResponse {
	return NewCEchoResponse(req.MessageID(), statusCode)
}

// StatusCode returns the status code.
func (r *CEchoResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *CEchoResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// String returns a human-readable representation.
func (r *CEchoResponse) String() string {
	st := r.Status()
	return fmt.Sprintf("C-ECHO-RSP [MessageID=%d, Status=%s (0x%04X)]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode)
}
