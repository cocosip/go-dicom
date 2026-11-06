// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// CStoreRequest represents a C-STORE-RQ message.
// C-STORE is used to store (send) DICOM instances.
type CStoreRequest struct {
	*BaseRequest
	affectedSOPClassUID     string
	affectedSOPInstanceUID  string
	priority                uint16
	moveOriginatorAET       string
	moveOriginatorMessageID uint16
}

// NewCStoreRequest creates a new C-STORE-RQ message from a DICOM dataset.
// The MessageID will be automatically assigned by the Association/Client when sending.
func NewCStoreRequest(dicomData *dataset.Dataset) (*CStoreRequest, error) {
	// Extract required fields from the dataset
	sopClassUID, ok := dicomData.GetString(tag.SOPClassUID)
	if !ok {
		return nil, fmt.Errorf("SOPClassUID not found in dataset")
	}

	sopInstanceUID, ok := dicomData.GetString(tag.SOPInstanceUID)
	if !ok {
		return nil, fmt.Errorf("SOPInstanceUID not found in dataset")
	}

	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandCStoreRQ), 0)

	// Set affected SOP Class UID and Instance UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	_ = command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	// Priority (optional, default to medium)
	_ = command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))

	// CommandDataSetType - dataset is present (use AddOrUpdate since CreateCommandDataset sets it to 0x0101)
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))

	return &CStoreRequest{
		BaseRequest:            NewBaseRequest(command, dicomData),
		affectedSOPClassUID:    sopClassUID,
		affectedSOPInstanceUID: sopInstanceUID,
		priority:               uint16(PriorityMedium),
	}, nil
}

// SetPriority sets the priority of the request.
func (r *CStoreRequest) SetPriority(priority uint16) {
	r.priority = priority
	_ = r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
}

// SetMoveOriginator sets the move originator AET and message ID (for C-MOVE sub-operations).
// It updates the command elements in place and returns an error if they cannot be written.
func (r *CStoreRequest) SetMoveOriginator(aet string, messageID uint16) error {
	if err := r.command.AddOrUpdate(element.NewString(tag.MoveOriginatorApplicationEntityTitle, vr.AE, []string{aet})); err != nil {
		return fmt.Errorf("failed to set move originator AE title: %w", err)
	}
	if err := r.command.AddOrUpdate(element.NewUnsignedShort(tag.MoveOriginatorMessageID, []uint16{messageID})); err != nil {
		return fmt.Errorf("failed to set move originator message ID: %w", err)
	}
	r.moveOriginatorAET = aet
	r.moveOriginatorMessageID = messageID
	return nil
}

// MoveOriginator returns the Move Originator AE Title and Message ID from the
// command set. ok is false when either element is absent.
func (r *CStoreRequest) MoveOriginator() (aet string, messageID uint16, ok bool) {
	aet, hasAET := r.command.GetString(tag.MoveOriginatorApplicationEntityTitle)
	id, err := r.command.GetUInt16(tag.MoveOriginatorMessageID, 0)
	if !hasAET || err != nil {
		return "", 0, false
	}
	return aet, id, true
}

// SetAffectedSOPClassUID updates the affected SOP Class UID in the cached
// field and in the command set.
func (r *CStoreRequest) SetAffectedSOPClassUID(uid string) {
	r.affectedSOPClassUID = uid
	_ = r.command.AddOrUpdate(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{uid}))
}

// SetAffectedSOPInstanceUID updates the affected SOP Instance UID in the cached
// field and in the command set.
func (r *CStoreRequest) SetAffectedSOPInstanceUID(uid string) {
	r.affectedSOPInstanceUID = uid
	_ = r.command.AddOrUpdate(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{uid}))
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *CStoreRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *CStoreRequest) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// String returns a human-readable representation.
func (r *CStoreRequest) String() string {
	return fmt.Sprintf("C-STORE-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s]",
		r.MessageID(), r.affectedSOPClassUID, r.affectedSOPInstanceUID)
}

// CStoreResponse represents a C-STORE-RSP message.
type CStoreResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	messageIDBeingRespondedTo uint16
}

// NewCStoreResponse creates a new C-STORE-RSP message.
func NewCStoreResponse(messageIDBeingRespondedTo uint16, s *status.Status, sopClassUID, sopInstanceUID string) *CStoreResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCStoreRSP), 0)

	// Set affected SOP Class UID and Instance UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	_ = command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	// MessageIDBeingRespondedTo
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{s.Code}))

	// CommandDataSetType - no dataset in response
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &CStoreResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                s.Code,
		affectedSOPClassUID:       sopClassUID,
		affectedSOPInstanceUID:    sopInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCStoreResponseSuccess creates a successful C-STORE-RSP message.
func NewCStoreResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID, sopInstanceUID string) *CStoreResponse {
	return NewCStoreResponse(messageIDBeingRespondedTo, status.Success, sopClassUID, sopInstanceUID)
}

// NewCStoreResponseFromRequest creates a C-STORE-RSP message from the corresponding request.
// This is a convenience function that automatically extracts the SOP Class UID and SOP Instance UID
// from the request.
//
// Example:
//
//	// When receiving a C-STORE request
//	resp := dimse.NewCStoreResponseFromRequest(req, status.Success) // Success
func NewCStoreResponseFromRequest(req *CStoreRequest, s *status.Status) *CStoreResponse {
	resp := NewCStoreResponse(
		req.MessageID(),
		s,
		req.AffectedSOPClassUID(),
		req.AffectedSOPInstanceUID(),
	)
	// Copy the presentation context ID so the response is sent on the same
	// context as the request (required by DICOM PS3.7).
	resp.SetPresentationContextID(req.PresentationContextID())
	return resp
}

// StatusCode returns the status code.
func (r *CStoreResponse) StatusCode() uint16 {
	return r.statusCode
}

// String returns a human-readable representation.
func (r *CStoreResponse) String() string {
	st := r.Status()
	return fmt.Sprintf("C-STORE-RSP [MessageID=%d, Status=%s (0x%04X), SOP Instance=%s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, r.affectedSOPInstanceUID)
}
