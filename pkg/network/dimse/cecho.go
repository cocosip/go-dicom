// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/status"
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
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUIDVerification}))

	// The C-ECHO-RQ command set does not define a Priority element.

	// CommandDataSetType is already set to 0x0101 (no dataset)

	return &CEchoRequest{
		BaseRequest:         NewBaseRequest(command, nil),
		affectedSOPClassUID: sopClassUIDVerification,
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
	statusCode                uint16
	affectedSOPClassUID       string
	messageIDBeingRespondedTo uint16
}

// NewCEchoResponse creates a new C-ECHO-RSP message.
//
// PS3.7 9.3.5.2 (Table 9.3-13) requires the C-ECHO-RSP Status to be Success. A
// non-Success status is accepted here only so that a non-conformant peer's
// response can be represented for diagnostics; EncodeDIMSEMessage rejects such
// a message, so it can never be sent over the wire.
func NewCEchoResponse(messageIDBeingRespondedTo uint16, s *status.Status) *CEchoResponse {
	s = resolveStatus(s)
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCEchoRSP), 0) // Response doesn't have its own MessageID

	// Set affected SOP Class UID to Verification SOP Class
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUIDVerification}))

	// MessageIDBeingRespondedTo
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{s.Code}))

	// CommandDataSetType is already set to 0x0101 (no dataset)

	return &CEchoResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                s.Code,
		affectedSOPClassUID:       sopClassUIDVerification,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCEchoResponseSuccess creates a successful C-ECHO-RSP message.
func NewCEchoResponseSuccess(messageIDBeingRespondedTo uint16) *CEchoResponse {
	return NewCEchoResponse(messageIDBeingRespondedTo, status.Success)
}

// NewCEchoResponseFromRequest creates a C-ECHO-RSP message from the corresponding request.
// This is a convenience function that automatically uses the request's MessageID.
//
// Example:
//
//	// When receiving a C-ECHO request
//	resp := dimse.NewCEchoResponseFromRequest(req, status.Success) // Success
func NewCEchoResponseFromRequest(req *CEchoRequest, s *status.Status) *CEchoResponse {
	resp := NewCEchoResponse(req.MessageID(), s)
	resp.SetPresentationContextID(req.PresentationContextID())
	return resp
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
