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

// CMoveRequest represents a C-MOVE-RQ message.
// C-MOVE is used to request that an SCP move DICOM instances to a specified destination.
type CMoveRequest struct {
	*BaseRequest
	affectedSOPClassUID string
	priority            uint16
	moveDestination     string
	queryLevel          QueryRetrieveLevel
}

// NewCMoveRequest creates a new C-MOVE-RQ message.
// The SOP Class UID is automatically determined based on the query level:
//   - Patient level: Uses Patient Root Query/Retrieve Information Model - MOVE
//   - Study/Series/Image levels: Use Study Root Query/Retrieve Information Model - MOVE
//
// The moveDestination specifies the AE title where the instances should be moved to.
// The MessageID will be automatically assigned by the Association/Client when sending.
func NewCMoveRequest(level QueryRetrieveLevel, moveDestination string, identifier *dataset.Dataset) *CMoveRequest {
	// Determine SOP Class UID based on query level
	var sopClassUID string
	switch level {
	case QueryRetrieveLevelPatient:
		// Patient Root Query/Retrieve Information Model - MOVE
		sopClassUID = "1.2.840.10008.5.1.4.1.2.1.2"
	case QueryRetrieveLevelStudy, QueryRetrieveLevelSeries, QueryRetrieveLevelImage:
		// Study Root Query/Retrieve Information Model - MOVE
		sopClassUID = "1.2.840.10008.5.1.4.1.2.2.2"
	default:
		// Default to Study Root
		sopClassUID = "1.2.840.10008.5.1.4.1.2.2.2"
	}

	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandCMoveRQ), 0)

	// Set affected SOP Class UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))

	// Priority (optional, default to medium)
	_ = command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))

	// Move Destination AE Title
	_ = command.Add(element.NewString(tag.MoveDestination, vr.AE, []string{moveDestination}))

	// CommandDataSetType - dataset is present (identifier)
	_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))

	// Add QueryRetrieveLevel to the identifier dataset if not already present
	if identifier != nil {
		_, exists := identifier.Get(tag.QueryRetrieveLevel)
		if !exists {
			_ = identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(level)}))
		}
	}

	return &CMoveRequest{
		BaseRequest:         NewBaseRequest(command, identifier),
		affectedSOPClassUID: sopClassUID,
		priority:            uint16(PriorityMedium),
		moveDestination:     moveDestination,
		queryLevel:          level,
	}
}

// NewCMoveRequestPatientRoot creates a C-MOVE-RQ for Patient Root Query/Retrieve.
func NewCMoveRequestPatientRoot(level QueryRetrieveLevel, moveDestination string, identifier *dataset.Dataset) *CMoveRequest {
	return NewCMoveRequest(level, moveDestination, identifier)
}

// NewCMoveRequestStudyRoot creates a C-MOVE-RQ for Study Root Query/Retrieve.
func NewCMoveRequestStudyRoot(level QueryRetrieveLevel, moveDestination string, identifier *dataset.Dataset) *CMoveRequest {
	return NewCMoveRequest(level, moveDestination, identifier)
}

// AffectedSOPClassUID returns the affected SOP class UID.
func (r *CMoveRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// Priority returns the priority.
func (r *CMoveRequest) Priority() uint16 {
	return r.priority
}

// SetPriority sets the priority of the request.
func (r *CMoveRequest) SetPriority(priority uint16) {
	r.priority = priority
	_ = r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
}

// MoveDestination returns the move destination AE title.
func (r *CMoveRequest) MoveDestination() string {
	return r.moveDestination
}

// QueryLevel returns the query/retrieve level.
func (r *CMoveRequest) QueryLevel() QueryRetrieveLevel {
	return r.queryLevel
}

// String returns a human-readable representation.
func (r *CMoveRequest) String() string {
	return fmt.Sprintf("C-MOVE-RQ [MessageID=%d, SOPClass=%s, Destination=%s, Level=%s]",
		r.MessageID(), r.affectedSOPClassUID, r.moveDestination, r.queryLevel)
}

// CMoveResponse represents a C-MOVE-RSP message.
// C-MOVE responses can be:
//   - Pending (0xFF00) with sub-operations progress
//   - Success (0x0000) when all sub-operations are complete
//   - Warning (0xB000) when completed with warnings
//   - Failure (0xXXXX) when operation fails
type CMoveResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	messageIDBeingRespondedTo uint16

	// Sub-operation counters (optional, present in pending responses)
	numberOfRemainingSubOperations uint16
	numberOfCompletedSubOperations uint16
	numberOfFailedSubOperations    uint16
	numberOfWarningSubOperations   uint16
}

// NewCMoveResponse creates a new C-MOVE-RSP message.
//
// Status codes:
//   - 0xFF00: Pending (sub-operations are continuing)
//   - 0x0000: Success (all sub-operations completed successfully)
//   - 0xB000: Warning (completed with some sub-operation warnings)
//   - 0xA701: Out of resources - unable to calculate number of matches
//   - 0xA702: Out of resources - unable to perform sub-operations
//   - 0xA801: Move destination unknown
//   - 0xA900: Identifier does not match SOP Class
//   - 0xC000: Unable to process
func NewCMoveResponse(messageIDBeingRespondedTo uint16, statusCode uint16, sopClassUID string) *CMoveResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCMoveRSP), 0)

	// Set message ID being responded to
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Set affected SOP Class UID (optional but recommended)
	if sopClassUID != "" {
		_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	}

	// Set status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType - no dataset
	_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &CMoveResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                statusCode,
		affectedSOPClassUID:       sopClassUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCMoveResponsePending creates a pending C-MOVE-RSP message with sub-operation counts.
func NewCMoveResponsePending(messageIDBeingRespondedTo uint16, sopClassUID string,
	remaining, completed, failed, warning uint16) *CMoveResponse {

	resp := NewCMoveResponse(messageIDBeingRespondedTo, 0xFF00, sopClassUID)

	// Add sub-operation counters
	_ = resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfRemainingSuboperations, []uint16{remaining}))
	_ = resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfCompletedSuboperations, []uint16{completed}))
	_ = resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfFailedSuboperations, []uint16{failed}))
	_ = resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfWarningSuboperations, []uint16{warning}))

	resp.numberOfRemainingSubOperations = remaining
	resp.numberOfCompletedSubOperations = completed
	resp.numberOfFailedSubOperations = failed
	resp.numberOfWarningSubOperations = warning

	return resp
}

// NewCMoveResponseSuccess creates a successful C-MOVE-RSP message.
func NewCMoveResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID string) *CMoveResponse {
	return NewCMoveResponse(messageIDBeingRespondedTo, 0x0000, sopClassUID)
}

// NewCMoveResponseFromRequest creates a C-MOVE-RSP message from the corresponding request.
func NewCMoveResponseFromRequest(req *CMoveRequest, statusCode uint16) *CMoveResponse {
	return NewCMoveResponse(req.MessageID(), statusCode, req.AffectedSOPClassUID())
}

// StatusCode returns the status code.
func (r *CMoveResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP class UID.
func (r *CMoveResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// MessageIDBeingRespondedTo returns the message ID being responded to.
func (r *CMoveResponse) MessageIDBeingRespondedTo() uint16 {
	return r.messageIDBeingRespondedTo
}

// NumberOfRemainingSubOperations returns the number of remaining sub-operations.
func (r *CMoveResponse) NumberOfRemainingSubOperations() uint16 {
	return r.numberOfRemainingSubOperations
}

// NumberOfCompletedSubOperations returns the number of completed sub-operations.
func (r *CMoveResponse) NumberOfCompletedSubOperations() uint16 {
	return r.numberOfCompletedSubOperations
}

// NumberOfFailedSubOperations returns the number of failed sub-operations.
func (r *CMoveResponse) NumberOfFailedSubOperations() uint16 {
	return r.numberOfFailedSubOperations
}

// NumberOfWarningSubOperations returns the number of warning sub-operations.
func (r *CMoveResponse) NumberOfWarningSubOperations() uint16 {
	return r.numberOfWarningSubOperations
}

// HasSubOperationCounts returns true if this response includes sub-operation counts.
func (r *CMoveResponse) HasSubOperationCounts() bool {
	// Sub-operation counts are typically present in pending responses
	return r.statusCode == 0xFF00
}

// String returns a human-readable representation.
func (r *CMoveResponse) String() string {
	st := r.Status()
	if r.HasSubOperationCounts() {
		return fmt.Sprintf("C-MOVE-RSP [MessageID=%d, Status=%s (0x%04X), Remaining=%d, Completed=%d, Failed=%d, Warning=%d]",
			r.MessageIDBeingRespondedTo(), st.State, r.statusCode,
			r.numberOfRemainingSubOperations, r.numberOfCompletedSubOperations,
			r.numberOfFailedSubOperations, r.numberOfWarningSubOperations)
	}
	return fmt.Sprintf("C-MOVE-RSP [MessageID=%d, Status=%s (0x%04X)]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode)
}
