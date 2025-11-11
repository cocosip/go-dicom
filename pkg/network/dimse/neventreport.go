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

// NEventReportRequest represents an N-EVENT-REPORT-RQ message.
// N-EVENT-REPORT is used to report an event about a managed SOP Instance.
type NEventReportRequest struct {
	*BaseRequest
	affectedSOPClassUID    string
	affectedSOPInstanceUID string
	eventTypeID            uint16
}

// NewNEventReportRequest creates a new N-EVENT-REPORT-RQ message.
// The MessageID will be automatically assigned by the Association/Client when sending.
//
// Parameters:
//   - affectedSOPClassUID: The SOP Class UID of the managed object
//   - affectedSOPInstanceUID: The SOP Instance UID of the managed object
//   - eventTypeID: Application-specific event type identifier
//   - eventInformation: Optional dataset containing event-specific information
func NewNEventReportRequest(
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	eventTypeID uint16,
	eventInformation *dataset.Dataset,
) *NEventReportRequest {
	// Create command dataset with MessageID=0 (unassigned)
	command := CreateCommandDataset(uint16(CommandNEventReportRQ), 0)

	// Set affected SOP Class UID and Instance UID
	_ = command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{affectedSOPClassUID}))
	_ = command.Add(element.NewString(tag.AffectedSOPInstanceUID, vr.UI, []string{affectedSOPInstanceUID}))

	// Set event type ID
	_ = command.Add(element.NewUnsignedShort(tag.EventTypeID, []uint16{eventTypeID}))

	// CommandDataSetType
	if eventInformation != nil {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0001}))
	} else {
		_ = command.Add(element.NewUnsignedShort(tag.CommandDataSetType, []uint16{0x0101}))
	}

	return &NEventReportRequest{
		BaseRequest:            NewBaseRequest(command, eventInformation),
		affectedSOPClassUID:    affectedSOPClassUID,
		affectedSOPInstanceUID: affectedSOPInstanceUID,
		eventTypeID:            eventTypeID,
	}
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NEventReportRequest) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NEventReportRequest) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// EventTypeID returns the event type identifier.
func (r *NEventReportRequest) EventTypeID() uint16 {
	return r.eventTypeID
}

// HasEventInformation returns true if this request contains event information dataset.
func (r *NEventReportRequest) HasEventInformation() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NEventReportRequest) String() string {
	hasInfo := "no event info"
	if r.HasEventInformation() {
		hasInfo = "with event info"
	}
	return fmt.Sprintf("N-EVENT-REPORT-RQ [MessageID=%d, SOP Class=%s, SOP Instance=%s, EventTypeID=%d, %s]",
		r.MessageID(), r.affectedSOPClassUID, r.affectedSOPInstanceUID, r.eventTypeID, hasInfo)
}

// NEventReportResponse represents an N-EVENT-REPORT-RSP message.
type NEventReportResponse struct {
	*BaseResponse
	statusCode                uint16
	affectedSOPClassUID       string
	affectedSOPInstanceUID    string
	eventTypeID               uint16
	messageIDBeingRespondedTo uint16
}

// NewNEventReportResponse creates a new N-EVENT-REPORT-RSP message.
//
// Parameters:
//   - messageIDBeingRespondedTo: MessageID of the request being responded to
//   - statusCode: Status code (0x0000 = Success, etc.)
//   - affectedSOPClassUID: The SOP Class UID
//   - affectedSOPInstanceUID: The SOP Instance UID
//   - eventTypeID: Event type identifier from the request
//   - eventReply: Optional dataset containing event-specific reply information
func NewNEventReportResponse(
	messageIDBeingRespondedTo uint16,
	statusCode uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	eventTypeID uint16,
	eventReply *dataset.Dataset,
) *NEventReportResponse {
	command := createNOperationResponseCommand(
		uint16(CommandNEventReportRSP),
		messageIDBeingRespondedTo,
		statusCode,
		affectedSOPClassUID,
		affectedSOPInstanceUID,
		tag.EventTypeID,
		eventTypeID,
		eventReply != nil,
	)

	return &NEventReportResponse{
		BaseResponse:              NewBaseResponse(command, eventReply),
		statusCode:                statusCode,
		affectedSOPClassUID:       affectedSOPClassUID,
		affectedSOPInstanceUID:    affectedSOPInstanceUID,
		eventTypeID:               eventTypeID,
		messageIDBeingRespondedTo: messageIDBeingRespondedTo,
	}
}

// NewNEventReportResponseSuccess creates a successful N-EVENT-REPORT-RSP message.
func NewNEventReportResponseSuccess(
	messageIDBeingRespondedTo uint16,
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	eventTypeID uint16,
	eventReply *dataset.Dataset,
) *NEventReportResponse {
	return NewNEventReportResponse(messageIDBeingRespondedTo, 0x0000, affectedSOPClassUID,
		affectedSOPInstanceUID, eventTypeID, eventReply)
}

// StatusCode returns the status code.
func (r *NEventReportResponse) StatusCode() uint16 {
	return r.statusCode
}

// AffectedSOPClassUID returns the affected SOP Class UID.
func (r *NEventReportResponse) AffectedSOPClassUID() string {
	return r.affectedSOPClassUID
}

// AffectedSOPInstanceUID returns the affected SOP Instance UID.
func (r *NEventReportResponse) AffectedSOPInstanceUID() string {
	return r.affectedSOPInstanceUID
}

// EventTypeID returns the event type identifier.
func (r *NEventReportResponse) EventTypeID() uint16 {
	return r.eventTypeID
}

// HasEventReply returns true if this response contains event reply dataset.
func (r *NEventReportResponse) HasEventReply() bool {
	return r.data != nil
}

// String returns a human-readable representation.
func (r *NEventReportResponse) String() string {
	st := r.Status()
	hasReply := "no event reply"
	if r.HasEventReply() {
		hasReply = "with event reply"
	}
	return fmt.Sprintf("N-EVENT-REPORT-RSP [MessageID=%d, Status=%s (0x%04X), EventTypeID=%d, %s]",
		r.MessageIDBeingRespondedTo(), st.State, r.statusCode, r.eventTypeID, hasReply)
}
