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
	// Patient Root vs Study Root UIDs for GET
	sopClassUID := sopClassUIDForLevel(level, [2]string{
		sopClassUIDPatientRootGet, // Patient Root GET
		sopClassUIDStudyRootGet,   // Study Root GET
	})
	return newCGetRequestWithSOPClassUID(level, identifier, sopClassUID)
}

func newCGetRequestWithSOPClassUID(level QueryRetrieveLevel, identifier *dataset.Dataset, sopClassUID string) *CGetRequest {
	command := createQueryRetrieveRequest(uint16(CommandCGetRQ), sopClassUID, level, identifier)
	return &CGetRequest{
		BaseRequest:         NewBaseRequest(command, identifier),
		affectedSOPClassUID: sopClassUID,
		priority:            uint16(PriorityMedium),
		queryLevel:          level,
	}
}

// NewCGetRequestPatientRoot creates a C-GET-RQ for Patient Root Query/Retrieve.
func NewCGetRequestPatientRoot(level QueryRetrieveLevel, identifier *dataset.Dataset) *CGetRequest {
	return newCGetRequestWithSOPClassUID(level, identifier, sopClassUIDPatientRootGet)
}

// NewCGetRequestStudyRoot creates a C-GET-RQ for Study Root Query/Retrieve.
func NewCGetRequestStudyRoot(level QueryRetrieveLevel, identifier *dataset.Dataset) *CGetRequest {
	return newCGetRequestWithSOPClassUID(level, identifier, sopClassUIDStudyRootGet)
}

// AffectedSOPClassUID returns the affected SOP class UID.
func (r *CGetRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// Priority returns the priority.
func (r *CGetRequest) Priority() uint16 {
	return r.priority
}

// SetAffectedSOPClassUID overwrites the affected SOP Class UID in the cached
// field and in the command set so a rebuilt request echoes the UID the peer
// actually sent.
func (r *CGetRequest) SetAffectedSOPClassUID(uid string) {
	r.affectedSOPClassUID = uid
	_ = r.command.AddOrUpdate(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{uid}))
}

// SetPriority sets the priority of the request.
func (r *CGetRequest) SetPriority(priority uint16) {
	r.priority = priority
	_ = r.command.AddOrUpdate(element.NewUnsignedShort(tag.Priority, []uint16{priority}))
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

	// Sub-operation counters (always present; zero when not reported)
	numberOfRemainingSubOperations uint16
	numberOfCompletedSubOperations uint16
	numberOfFailedSubOperations    uint16
	numberOfWarningSubOperations   uint16
}

// subOperationCountTags are the four C-GET/C-MOVE-RSP sub-operation count
// command elements (PS3.4 C.4.2.2.1). Every response constructor adds them so
// that the counts are always present on the wire.
var subOperationCountTags = [...]*tag.Tag{
	tag.NumberOfRemainingSuboperations,
	tag.NumberOfCompletedSuboperations,
	tag.NumberOfFailedSuboperations,
	tag.NumberOfWarningSuboperations,
}

func addSubOperationCounts(command *dataset.Dataset, remaining, completed, failed, warning uint16) {
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfRemainingSuboperations, []uint16{remaining}))
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfCompletedSuboperations, []uint16{completed}))
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfFailedSuboperations, []uint16{failed}))
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.NumberOfWarningSuboperations, []uint16{warning}))
}

func hasSubOperationCounts(command *dataset.Dataset) bool {
	if command == nil {
		return false
	}
	for _, countTag := range subOperationCountTags {
		if _, ok := command.Get(countTag); !ok {
			return false
		}
	}
	return true
}

// NewCGetResponse creates a new C-GET-RSP message.
//
// Status codes:
//   - status.CGetPending (0xFF00): Pending (sub-operations are continuing)
//   - status.Success (0x0000): Success (all sub-operations completed successfully)
//   - status.CGetWarningSubOperationsComplete (0xB000): Warning (completed with some sub-operation warnings)
//   - status.CGetRefusedOutOfResources (0xA701): Out of resources - unable to calculate number of matches
//   - status.CGetRefusedOutOfResourcesSubOps (0xA702): Out of resources - unable to perform sub-operations
//   - status.CGetFailedIdentifierDoesNotMatchSOPClass (0xA900): Identifier does not match SOP Class
//   - status.CGetFailedUnableToProcess (0xC000): Unable to process
//
// The four sub-operation count elements are always added (zero by default);
// use SetSubOperationCounts to report progress.
func NewCGetResponse(messageIDBeingRespondedTo uint16, s *status.Status, sopClassUID string) *CGetResponse {
	s = resolveStatus(s)
	// Create command dataset
	command := CreateCommandDataset(uint16(CommandCGetRSP), 0)

	// Set message ID being responded to
	_ = command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{messageIDBeingRespondedTo}))

	// Affected SOP Class UID is Type 1 in C-GET-RSP (PS3.7 Table 9.3-7).
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{sopClassUID}))

	// Set status
	_ = command.Add(element.NewUnsignedShort(tag.Status, []uint16{s.Code}))

	// CommandDataSetType - no dataset
	_ = command.AddOrUpdate(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))

	// Sub-operation counts are always present (PS3.4 C.4.2.2.1).
	addSubOperationCounts(command, 0, 0, 0, 0)

	return &CGetResponse{
		BaseResponse:              NewBaseResponse(command, nil),
		statusCode:                s.Code,
		affectedSOPClassUID:       sopClassUID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// SetSubOperationCounts sets the four sub-operation count elements and the
// cached values. It can be used on any C-GET-RSP, pending or final.
func (r *CGetResponse) SetSubOperationCounts(remaining, completed, failed, warning uint16) {
	addSubOperationCounts(r.command, remaining, completed, failed, warning)
	r.numberOfRemainingSubOperations = remaining
	r.numberOfCompletedSubOperations = completed
	r.numberOfFailedSubOperations = failed
	r.numberOfWarningSubOperations = warning
}

// NewCGetResponsePending creates a pending C-GET-RSP message with sub-operation counts.
func NewCGetResponsePending(messageIDBeingRespondedTo uint16, sopClassUID string,
	remaining, completed, failed, warning uint16) *CGetResponse {

	resp := NewCGetResponse(messageIDBeingRespondedTo, status.CGetPending, sopClassUID)
	resp.SetSubOperationCounts(remaining, completed, failed, warning)
	return resp
}

// NewCGetResponseSuccess creates a successful C-GET-RSP message.
func NewCGetResponseSuccess(messageIDBeingRespondedTo uint16, sopClassUID string) *CGetResponse {
	return NewCGetResponse(messageIDBeingRespondedTo, status.Success, sopClassUID)
}

// NewCGetResponseFromRequest creates a C-GET-RSP message from the corresponding request.
func NewCGetResponseFromRequest(req *CGetRequest, s *status.Status) *CGetResponse {
	resp := NewCGetResponse(req.MessageID(), s, req.AffectedSOPClassUID())
	resp.SetPresentationContextID(req.PresentationContextID())
	return resp
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

// HasSubOperationCounts reports whether the four sub-operation count elements
// are present in the command set. Constructors always add them.
func (r *CGetResponse) HasSubOperationCounts() bool {
	return hasSubOperationCounts(r.command)
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
