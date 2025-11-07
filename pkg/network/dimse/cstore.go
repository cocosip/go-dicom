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

// CStoreRequest represents a C-STORE-RQ message.
// C-STORE is used to store (send) DICOM instances.
type CStoreRequest struct {
	*BaseRequest
	affectedSOPClassUID    string
	affectedSOPInstanceUID string
	priority               uint16
	moveOriginatorAET      string
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
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	// Priority (optional, default to medium)
	command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))

	// CommandDataSetType - dataset is present
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))

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
	r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
}

// SetMoveOriginator sets the move originator AET and message ID (for C-MOVE sub-operations).
func (r *CStoreRequest) SetMoveOriginator(aet string, messageID uint16) {
	r.moveOriginatorAET = aet
	r.moveOriginatorMessageID = messageID

	r.command.Add(element.NewString(tag.MoveOriginatorApplicationEntityTitle, vr.AE, []string{aet}))
	r.command.Add(element.NewUnsignedShort(tag.MoveOriginatorMessageID, []uint16{messageID}))
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
func NewCStoreResponse(messageIDBeingRespondedTo uint16, statusCode uint16, sopClassUID, sopInstanceUID string) *CStoreResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCStoreRSP), 0)

	// Set affected SOP Class UID and Instance UID
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	// MessageIDBeingRespondedTo
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType - no dataset in response
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &CStoreResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                statusCode,
		affectedSOPClassUID:       sopClassUID,
		affectedSOPInstanceUID:    sopInstanceUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCStoreResponseSuccess creates a successful C-STORE-RSP message.
func NewCStoreResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID, sopInstanceUID string) *CStoreResponse {
	return NewCStoreResponse(messageIDBeingRespondedTo, 0x0000, sopClassUID, sopInstanceUID)
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
