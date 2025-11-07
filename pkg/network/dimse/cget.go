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

// CGetRequest represents a C-GET-RQ message.
// C-GET is used to request that an SCP retrieve DICOM instances and send them back
// to the requestor (SCU) over the same association.
type CGetRequest struct {
	*BaseRequest
	affectedSOPClassUID string
	priority            uint16
	queryLevel          QueryRetrieveLevel
}

// NewCGetRequest creates a new C-GET-RQ message.
// The SOP Class UID is automatically determined based on the query level:
//   - Patient level: Uses Patient Root Query/Retrieve Information Model - GET
//   - Study/Series/Image levels: Use Study Root Query/Retrieve Information Model - GET
//
// The MessageID will be automatically assigned by the Association/Client when sending.
func NewCGetRequest(level QueryRetrieveLevel, identifier *dataset.Dataset) *CGetRequest {
	// Determine SOP Class UID based on query level
	var sopClassUID string
	switch level {
	case QueryRetrieveLevelPatient:
		// Patient Root Query/Retrieve Information Model - GET
		sopClassUID = "1.2.840.10008.5.1.4.1.2.1.3"
	case QueryRetrieveLevelStudy, QueryRetrieveLevelSeries, QueryRetrieveLevelImage:
		// Study Root Query/Retrieve Information Model - GET
		sopClassUID = "1.2.840.10008.5.1.4.1.2.2.3"
	default:
		// Default to Study Root
		sopClassUID = "1.2.840.10008.5.1.4.1.2.2.3"
	}

	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandCGetRQ), 0)

	// Set affected SOP Class UID
	command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))

	// Priority (optional, default to medium)
	command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))

	// CommandDataSetType - dataset is present (identifier)
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))

	// Add QueryRetrieveLevel to the identifier dataset if not already present
	if identifier != nil {
		_, exists := identifier.Get(tag.QueryRetrieveLevel)
		if !exists {
			identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(level)}))
		}
	}

	return &CGetRequest{
		BaseRequest:         NewBaseRequest(command, identifier),
		affectedSOPClassUID: sopClassUID,
		priority:            uint16(PriorityMedium),
		queryLevel:          level,
	}
}

// NewCGetRequestPatientRoot creates a C-GET-RQ for Patient Root Query/Retrieve.
func NewCGetRequestPatientRoot(level QueryRetrieveLevel, identifier *dataset.Dataset) *CGetRequest {
	return NewCGetRequest(level, identifier)
}

// NewCGetRequestStudyRoot creates a C-GET-RQ for Study Root Query/Retrieve.
func NewCGetRequestStudyRoot(level QueryRetrieveLevel, identifier *dataset.Dataset) *CGetRequest {
	return NewCGetRequest(level, identifier)
}

// AffectedSOPClassUID returns the affected SOP class UID.
func (r *CGetRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// Priority returns the priority.
func (r *CGetRequest) Priority() uint16 {
	return r.priority
}

// SetPriority sets the priority of the request.
func (r *CGetRequest) SetPriority(priority uint16) {
	r.priority = priority
	r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
}

// QueryLevel returns the query/retrieve level.
func (r *CGetRequest) QueryLevel() QueryRetrieveLevel {
	return r.queryLevel
}

// String returns a human-readable representation.
func (r *CGetRequest) String() string {
	return fmt.Sprintf("C-GET-RQ [MessageID=%d, SOPClass=%s, Level=%s]",
		r.MessageID(), r.affectedSOPClassUID, r.queryLevel)
}

// CGetResponse represents a C-GET-RSP message.
// C-GET responses can be:
//   - Pending (0xFF00) with sub-operations progress
//   - Success (0x0000) when all sub-operations are complete
//   - Warning (0xB000) when completed with warnings
//   - Failure (0xXXXX) when operation fails
type CGetResponse struct {
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

// NewCGetResponse creates a new C-GET-RSP message.
//
// Status codes:
//   - 0xFF00: Pending (sub-operations are continuing)
//   - 0x0000: Success (all sub-operations completed successfully)
//   - 0xB000: Warning (completed with some sub-operation warnings)
//   - 0xA701: Out of resources - unable to calculate number of matches
//   - 0xA702: Out of resources - unable to perform sub-operations
//   - 0xA900: Identifier does not match SOP Class
//   - 0xC000: Unable to process
func NewCGetResponse(messageIDBeingRespondedTo uint16, statusCode uint16, sopClassUID string) *CGetResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCGetRSP), 0)

	// Set message ID being responded to
	command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Set affected SOP Class UID (optional but recommended)
	if sopClassUID != "" {
		command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	}

	// Set status
	command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType - no dataset
	command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	return &CGetResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                statusCode,
		affectedSOPClassUID:       sopClassUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCGetResponsePending creates a pending C-GET-RSP message with sub-operation counts.
func NewCGetResponsePending(messageIDBeingRespondedTo uint16, sopClassUID string,
	remaining, completed, failed, warning uint16) *CGetResponse {

	resp := NewCGetResponse(messageIDBeingRespondedTo, 0xFF00, sopClassUID)

	// Add sub-operation counters
	resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfRemainingSuboperations, []uint16{remaining}))
	resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfCompletedSuboperations, []uint16{completed}))
	resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfFailedSuboperations, []uint16{failed}))
	resp.command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfWarningSuboperations, []uint16{warning}))

	resp.numberOfRemainingSubOperations = remaining
	resp.numberOfCompletedSubOperations = completed
	resp.numberOfFailedSubOperations = failed
	resp.numberOfWarningSubOperations = warning

	return resp
}

// NewCGetResponseSuccess creates a successful C-GET-RSP message.
func NewCGetResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID string) *CGetResponse {
	return NewCGetResponse(messageIDBeingRespondedTo, 0x0000, sopClassUID)
}

// NewCGetResponseFromRequest creates a C-GET-RSP message from the corresponding request.
func NewCGetResponseFromRequest(req *CGetRequest, statusCode uint16) *CGetResponse {
	return NewCGetResponse(req.MessageID(), statusCode, req.AffectedSOPClassUID())
}

// StatusCode returns the status code.
func (r *CGetResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP class UID.
func (r *CGetResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// MessageIDBeingRespondedTo returns the message ID being responded to.
func (r *CGetResponse) MessageIDBeingRespondedTo() uint16 {
	return r.messageIDBeingRespondedTo
}

// NumberOfRemainingSubOperations returns the number of remaining sub-operations.
func (r *CGetResponse) NumberOfRemainingSubOperations() uint16 {
	return r.numberOfRemainingSubOperations
}

// NumberOfCompletedSubOperations returns the number of completed sub-operations.
func (r *CGetResponse) NumberOfCompletedSubOperations() uint16 {
	return r.numberOfCompletedSubOperations
}

// NumberOfFailedSubOperations returns the number of failed sub-operations.
func (r *CGetResponse) NumberOfFailedSubOperations() uint16 {
	return r.numberOfFailedSubOperations
}

// NumberOfWarningSubOperations returns the number of warning sub-operations.
func (r *CGetResponse) NumberOfWarningSubOperations() uint16 {
	return r.numberOfWarningSubOperations
}

// HasSubOperationCounts returns true if this response includes sub-operation counts.
func (r *CGetResponse) HasSubOperationCounts() bool {
	// Sub-operation counts are typically present in pending responses
	return r.statusCode == 0xFF00
}

// String returns a human-readable representation.
func (r *CGetResponse) String() string {
	st := r.Status()
	if r.HasSubOperationCounts() {
		return fmt.Sprintf("C-GET-RSP [MessageID=%d, Status=%s (0x%04X), Remaining=%d, Completed=%d, Failed=%d, Warning=%d]",
			r.MessageIDBeingRespondedTo(), st.State, r.statusCode,
			r.numberOfRemainingSubOperations, r.numberOfCompletedSubOperations,
			r.numberOfFailedSubOperations, r.numberOfWarningSubOperations)
	}
	return fmt.Sprintf("C-GET-RSP [MessageID=%d, Status=%s (0x%04X)]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode)
}
