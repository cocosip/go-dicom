// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Test N-EVENT-REPORT
func TestNewNEventReportRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"  // Procedure Step SOP Class
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"
	eventTypeID := uint16(1)

	eventInfo := dataset.New()
	eventInfo.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))

	req := NewNEventReportRequest(sopClassUID, sopInstanceUID, eventTypeID, eventInfo)

	// MessageID should be 0 (unassigned) initially
	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNEventReportRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNEventReportRQ, req.CommandField())
	}

	if req.AffectedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.AffectedSOPClassUID())
	}

	if req.AffectedSOPInstanceUID() != sopInstanceUID {
		t.Errorf("Wrong SOP Instance UID: %s", req.AffectedSOPInstanceUID())
	}

	if req.EventTypeID() != eventTypeID {
		t.Errorf("Expected EventTypeID %d, got %d", eventTypeID, req.EventTypeID())
	}

	if !req.HasEventInformation() {
		t.Error("Expected HasEventInformation() to be true")
	}

	str := req.String()
	if str == "" {
		t.Error("String() should not be empty")
	}
}

func TestNewNEventReportResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"
	eventTypeID := uint16(1)

	eventReply := dataset.New()
	eventReply.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	resp := NewNEventReportResponseSuccess(100, sopClassUID, sopInstanceUID, eventTypeID, eventReply)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNEventReportRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNEventReportRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	if resp.EventTypeID() != eventTypeID {
		t.Errorf("Expected EventTypeID %d, got %d", eventTypeID, resp.EventTypeID())
	}

	if !resp.HasEventReply() {
		t.Error("Expected HasEventReply() to be true")
	}
}

// Test N-GET
func TestNewNGetRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	// Request specific attributes
	attrList := []*tag.Tag{
		tag.PatientName,
		tag.PatientID,
		tag.StudyDate,
	}

	req := NewNGetRequest(sopClassUID, sopInstanceUID, attrList)

	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNGetRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNGetRQ, req.CommandField())
	}

	if req.RequestedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.RequestedSOPClassUID())
	}

	if req.RequestedSOPInstanceUID() != sopInstanceUID {
		t.Errorf("Wrong SOP Instance UID: %s", req.RequestedSOPInstanceUID())
	}

	if len(req.AttributeIdentifierList()) != 3 {
		t.Errorf("Expected 3 attributes, got %d", len(req.AttributeIdentifierList()))
	}
}

func TestNewNGetRequestAllAttributes(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	// Request all attributes (nil list)
	req := NewNGetRequest(sopClassUID, sopInstanceUID, nil)

	if req.AttributeIdentifierList() != nil {
		t.Error("Expected nil attribute list for all attributes request")
	}
}

func TestNewNGetResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	attrList := dataset.New()
	attrList.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
	attrList.Add(element.NewString(tag.PatientID, vr.LO, []string{"123456"}))

	resp := NewNGetResponseSuccess(100, sopClassUID, sopInstanceUID, attrList)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNGetRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNGetRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	if !resp.HasAttributeList() {
		t.Error("Expected HasAttributeList() to be true")
	}
}

// Test N-SET
func TestNewNSetRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	modList := dataset.New()
	modList.Add(element.NewString(tag.PatientName, vr.PN, []string{"Updated^Patient"}))

	req := NewNSetRequest(sopClassUID, sopInstanceUID, modList)

	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNSetRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNSetRQ, req.CommandField())
	}

	if req.RequestedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.RequestedSOPClassUID())
	}

	if !req.HasModificationList() {
		t.Error("Expected HasModificationList() to be true")
	}
}

func TestNewNSetResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	attrList := dataset.New()
	attrList.Add(element.NewString(tag.PatientName, vr.PN, []string{"Updated^Patient"}))

	resp := NewNSetResponseSuccess(100, sopClassUID, sopInstanceUID, attrList)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNSetRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNSetRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	if !resp.HasAttributeList() {
		t.Error("Expected HasAttributeList() to be true")
	}
}

// Test N-ACTION
func TestNewNActionRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"
	actionTypeID := uint16(1)

	actionInfo := dataset.New()
	actionInfo.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	req := NewNActionRequest(sopClassUID, sopInstanceUID, actionTypeID, actionInfo)

	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNActionRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNActionRQ, req.CommandField())
	}

	if req.RequestedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.RequestedSOPClassUID())
	}

	if req.ActionTypeID() != actionTypeID {
		t.Errorf("Expected ActionTypeID %d, got %d", actionTypeID, req.ActionTypeID())
	}

	if !req.HasActionInformation() {
		t.Error("Expected HasActionInformation() to be true")
	}
}

func TestNewNActionResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"
	actionTypeID := uint16(1)

	actionReply := dataset.New()
	actionReply.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	resp := NewNActionResponseSuccess(100, sopClassUID, sopInstanceUID, actionTypeID, actionReply)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNActionRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNActionRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if resp.ActionTypeID() != actionTypeID {
		t.Errorf("Expected ActionTypeID %d, got %d", actionTypeID, resp.ActionTypeID())
	}

	if !resp.HasActionReply() {
		t.Error("Expected HasActionReply() to be true")
	}
}

// Test N-CREATE
func TestNewNCreateRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	attrList := dataset.New()
	attrList.Add(element.NewString(tag.PatientName, vr.PN, []string{"New^Patient"}))

	req := NewNCreateRequest(sopClassUID, sopInstanceUID, attrList)

	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNCreateRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNCreateRQ, req.CommandField())
	}

	if req.AffectedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.AffectedSOPClassUID())
	}

	if req.AffectedSOPInstanceUID() != sopInstanceUID {
		t.Errorf("Wrong SOP Instance UID: %s", req.AffectedSOPInstanceUID())
	}

	if !req.HasAttributeList() {
		t.Error("Expected HasAttributeList() to be true")
	}
}

func TestNewNCreateRequestServerAssignedUID(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"

	attrList := dataset.New()
	attrList.Add(element.NewString(tag.PatientName, vr.PN, []string{"New^Patient"}))

	// Empty instance UID - server should assign
	req := NewNCreateRequest(sopClassUID, "", attrList)

	if req.AffectedSOPInstanceUID() != "" {
		t.Errorf("Expected empty SOP Instance UID, got %s", req.AffectedSOPInstanceUID())
	}
}

func TestNewNCreateResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	attrList := dataset.New()
	attrList.Add(element.NewString(tag.PatientName, vr.PN, []string{"New^Patient"}))

	resp := NewNCreateResponseSuccess(100, sopClassUID, sopInstanceUID, attrList)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNCreateRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNCreateRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if resp.AffectedSOPInstanceUID() != sopInstanceUID {
		t.Errorf("Wrong SOP Instance UID: %s", resp.AffectedSOPInstanceUID())
	}

	if !resp.HasAttributeList() {
		t.Error("Expected HasAttributeList() to be true")
	}
}

// Test N-DELETE
func TestNewNDeleteRequest(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	req := NewNDeleteRequest(sopClassUID, sopInstanceUID)

	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandNDeleteRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNDeleteRQ, req.CommandField())
	}

	if req.RequestedSOPClassUID() != sopClassUID {
		t.Errorf("Wrong SOP Class UID: %s", req.RequestedSOPClassUID())
	}

	if req.RequestedSOPInstanceUID() != sopInstanceUID {
		t.Errorf("Wrong SOP Instance UID: %s", req.RequestedSOPInstanceUID())
	}

	str := req.String()
	if str == "" {
		t.Error("String() should not be empty")
	}
}

func TestNewNDeleteResponse(t *testing.T) {
	sopClassUID := "1.2.840.10008.5.1.1.40"
	sopInstanceUID := "1.2.3.4.5.6.7.8.9.10"

	resp := NewNDeleteResponseSuccess(100, sopClassUID, sopInstanceUID)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandNDeleteRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandNDeleteRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	str := resp.String()
	if str == "" {
		t.Error("String() should not be empty")
	}
}

// Test MessageID assignment for N-operations
func TestNOperations_MessageIDAssignment(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Test N-EVENT-REPORT
	eventReq := NewNEventReportRequest("1.2.3", "1.2.3.4", 1, nil)
	msgID, err := gen.AssignMessageID(eventReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 1 {
		t.Errorf("Expected MessageID 1, got %d", msgID)
	}
	if eventReq.MessageID() != 1 {
		t.Errorf("Request MessageID not updated, got %d", eventReq.MessageID())
	}

	// Test N-GET
	getReq := NewNGetRequest("1.2.3", "1.2.3.4", nil)
	msgID, err = gen.AssignMessageID(getReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 2 {
		t.Errorf("Expected MessageID 2, got %d", msgID)
	}
	if getReq.MessageID() != 2 {
		t.Errorf("Request MessageID not updated, got %d", getReq.MessageID())
	}

	// Test N-SET
	setReq := NewNSetRequest("1.2.3", "1.2.3.4", nil)
	msgID, err = gen.AssignMessageID(setReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 3 {
		t.Errorf("Expected MessageID 3, got %d", msgID)
	}
	if setReq.MessageID() != 3 {
		t.Errorf("Request MessageID not updated, got %d", setReq.MessageID())
	}

	// Test N-ACTION
	actionReq := NewNActionRequest("1.2.3", "1.2.3.4", 1, nil)
	msgID, err = gen.AssignMessageID(actionReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 4 {
		t.Errorf("Expected MessageID 4, got %d", msgID)
	}
	if actionReq.MessageID() != 4 {
		t.Errorf("Request MessageID not updated, got %d", actionReq.MessageID())
	}

	// Test N-CREATE
	createReq := NewNCreateRequest("1.2.3", "1.2.3.4", nil)
	msgID, err = gen.AssignMessageID(createReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 5 {
		t.Errorf("Expected MessageID 5, got %d", msgID)
	}
	if createReq.MessageID() != 5 {
		t.Errorf("Request MessageID not updated, got %d", createReq.MessageID())
	}

	// Test N-DELETE
	deleteReq := NewNDeleteRequest("1.2.3", "1.2.3.4")
	msgID, err = gen.AssignMessageID(deleteReq)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}
	if msgID != 6 {
		t.Errorf("Expected MessageID 6, got %d", msgID)
	}
	if deleteReq.MessageID() != 6 {
		t.Errorf("Request MessageID not updated, got %d", deleteReq.MessageID())
	}
}
