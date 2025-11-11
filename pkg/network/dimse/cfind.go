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

// QueryRetrieveLevel represents the Query/Retrieve Information Model level.
type QueryRetrieveLevel string

const (
	// QueryRetrieveLevelPatient represents the PATIENT level
	QueryRetrieveLevelPatient QueryRetrieveLevel = "PATIENT"

	// QueryRetrieveLevelStudy represents the STUDY level
	QueryRetrieveLevelStudy QueryRetrieveLevel = "STUDY"

	// QueryRetrieveLevelSeries represents the SERIES level
	QueryRetrieveLevelSeries QueryRetrieveLevel = "SERIES"

	// QueryRetrieveLevelImage represents the IMAGE level
	QueryRetrieveLevelImage QueryRetrieveLevel = "IMAGE"
)

// sopClassUIDForLevel returns the appropriate SOP Class UID for a given query/retrieve level.
// The uids parameter should contain [patientRootUID, studyRootUID].
func sopClassUIDForLevel(level QueryRetrieveLevel, uids [2]string) string {
	switch level {
	case QueryRetrieveLevelPatient:
		return uids[0] // Patient Root
	case QueryRetrieveLevelStudy, QueryRetrieveLevelSeries, QueryRetrieveLevelImage:
		return uids[1] // Study Root
	default:
		return uids[1] // Default to Study Root
	}
}

// createQueryRetrieveRequest is a helper to create C-FIND/C-GET/C-MOVE request command datasets.
// This eliminates code duplication across NewCFindRequest, NewCGetRequest, and NewCMoveRequest.
func createQueryRetrieveRequest(
	commandType uint16,
	sopClassUID string,
	level QueryRetrieveLevel,
	data *dataset.Dataset,
) *dataset.Dataset {
	command := CreateCommandDataset(commandType, 0)
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))
	_ = command.Add(element.NewUnsignedShort(tag.Priority, []uint16{uint16(PriorityMedium)}))
	_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	if data != nil {
		if _, exists := data.Get(tag.QueryRetrieveLevel); !exists {
			_ = data.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(level)}))
		}
	}
	return command
}

// CFindRequest represents a C-FIND-RQ message.
// C-FIND is used to query DICOM archives.
type CFindRequest struct {
	*BaseRequest
	affectedSOPClassUID string
	priority            uint16
	queryLevel          QueryRetrieveLevel
}

// NewCFindRequest creates a new C-FIND-RQ message.
// The SOP Class UID is automatically determined based on the query level:
//   - Patient level: Uses Patient Root Query/Retrieve Information Model - FIND
//   - Study/Series/Image levels: Use Study Root Query/Retrieve Information Model - FIND
//
// The MessageID will be automatically assigned by the Association/Client when sending.
func NewCFindRequest(level QueryRetrieveLevel, query *dataset.Dataset) *CFindRequest {
	// Patient Root vs Study Root UIDs for FIND
	sopClassUID := sopClassUIDForLevel(level, [2]string{
		"1.2.840.10008.5.1.4.1.2.1.1", // Patient Root FIND
		"1.2.840.10008.5.1.4.1.2.2.1", // Study Root FIND
	})
	command := createQueryRetrieveRequest(uint16(CommandCFindRQ), sopClassUID, level, query)
	return &CFindRequest{
		BaseRequest:         NewBaseRequest(command, query),
		affectedSOPClassUID: sopClassUID,
		priority:            uint16(PriorityMedium),
		queryLevel:          level,
	}
}

// NewCFindRequestPatientRoot creates a C-FIND-RQ for Patient Root Query/Retrieve.
// This is a convenience function that's equivalent to NewCFindRequest(QueryRetrieveLevelPatient, query).
func NewCFindRequestPatientRoot(level QueryRetrieveLevel, query *dataset.Dataset) *CFindRequest {
	return NewCFindRequest(level, query)
}

// NewCFindRequestStudyRoot creates a C-FIND-RQ for Study Root Query/Retrieve.
// This is a convenience function that's equivalent to NewCFindRequest(level, query) for Study/Series/Image levels.
func NewCFindRequestStudyRoot(level QueryRetrieveLevel, query *dataset.Dataset) *CFindRequest {
	return NewCFindRequest(level, query)
}

// SetPriority sets the priority of the request.
func (r *CFindRequest) SetPriority(priority uint16) {
	r.priority = priority
	_ = r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
}

// QueryLevel returns the query/retrieve level.
func (r *CFindRequest) QueryLevel() QueryRetrieveLevel {
	return r.queryLevel
}

// String returns a human-readable representation.
func (r *CFindRequest) String() string {
	return fmt.Sprintf("C-FIND-RQ [MessageID=%d, SOP Class=%s, Level=%s]",
		r.MessageID(), r.affectedSOPClassUID, r.queryLevel)
}

// CFindResponse represents a C-FIND-RSP message.
type CFindResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	messageIDBeingRespondedTo uint16
}

// NewCFindResponse creates a new C-FIND-RSP message.
func NewCFindResponse(messageIDBeingRespondedTo uint16, statusCode uint16, sopClassUID string, identifier *dataset.Dataset) *CFindResponse {
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCFindRSP), 0)

	// Set affected SOP Class UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))

	// MessageIDBeingRespondedTo
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{statusCode}))

	// CommandDataSetType
	if identifier != nil && statusCode == 0xFF00 {
		// Pending response with identifier
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		// Final response or no identifier
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &CFindResponse{
		BaseResponse:              NewBaseResponse(command, identifier),
		statusCode:                statusCode,
		affectedSOPClassUID:       sopClassUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewCFindResponsePending creates a pending C-FIND-RSP message with a result.
func NewCFindResponsePending(messageIDBeingRespondedTo uint16, sopClassUID string, identifier *dataset.Dataset) *CFindResponse {
	return NewCFindResponse(messageIDBeingRespondedTo, 0xFF00, sopClassUID, identifier)
}

// NewCFindResponseSuccess creates a successful C-FIND-RSP message (no more results).
func NewCFindResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID string) *CFindResponse {
	return NewCFindResponse(messageIDBeingRespondedTo, 0x0000, sopClassUID, nil)
}

// NewCFindResponseFromRequest creates a C-FIND-RSP message from the corresponding request.
// This is a convenience function that automatically extracts the SOP Class UID from the request.
//
// Example:
//
//	// When receiving a C-FIND request and sending a pending response with results
//	resp := dimse.NewCFindResponseFromRequest(req, 0xFF00, resultDataset) // Pending
//
//	// When sending the final success response
//	resp := dimse.NewCFindResponseFromRequest(req, 0x0000, nil) // Success
func NewCFindResponseFromRequest(req *CFindRequest, statusCode uint16, identifier *dataset.Dataset) *CFindResponse {
	return NewCFindResponse(
		req.MessageID(),
		statusCode,
		req.AffectedSOPClassUID(),
		identifier,
	)
}

// StatusCode returns the status code.
func (r *CFindResponse) StatusCode() uint16 {
	return r.statusCode
}

// HasIdentifier returns true if this response contains an identifier dataset.
func (r *CFindResponse) HasIdentifier() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *CFindResponse) String() string {
	st := r.Status()
	hasData := "no data"
	if r.HasIdentifier() {
		hasData = "with data"
	}
	return fmt.Sprintf("C-FIND-RSP [MessageID=%d, Status=%s (0x%04X), %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, hasData)
}
