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

// Test BaseMessage
func TestBaseMessage_CommandField(t *testing.T) {
	command := CreateCommandDataset(uint16(CommandCEchoRQ), 123)
	msg := NewBaseMessage(command, nil)

	if msg.CommandField() != uint16(CommandCEchoRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCEchoRQ, msg.CommandField())
	}

	if msg.MessageID() != 123 {
		t.Errorf("Expected MessageID 123, got %d", msg.MessageID())
	}
}

func TestBaseMessage_PresentationContextID(t *testing.T) {
	command := CreateCommandDataset(uint16(CommandCEchoRQ), 1)
	msg := NewBaseMessage(command, nil)

	msg.SetPresentationContextID(5)
	if msg.PresentationContextID() != 5 {
		t.Errorf("Expected PresentationContextID 5, got %d", msg.PresentationContextID())
	}
}

// Test C-ECHO
func TestNewCEchoRequest(t *testing.T) {
	req := NewCEchoRequest()

	// MessageID should be 0 (unassigned) initially
	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	// Assign MessageID via generator
	gen := NewMessageIDGenerator()
	msgID, err := gen.AssignMessageID(req)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}

	if msgID != 1 {
		t.Errorf("Expected assigned MessageID to be 1, got %d", msgID)
	}

	if req.MessageID() != 1 {
		t.Errorf("Expected MessageID 1 after assignment, got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandCEchoRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCEchoRQ, req.CommandField())
	}

	if req.AffectedSOPClassUID() != "1.2.840.10008.1.1" {
		t.Errorf("Expected Verification SOP Class, got %s", req.AffectedSOPClassUID())
	}

	if req.Priority() != uint16(PriorityMedium) {
		t.Errorf("Expected Priority %d, got %d", PriorityMedium, req.Priority())
	}

	str := req.String()
	if str == "" {
		t.Error("String() should not be empty")
	}
}

func TestNewCEchoResponse(t *testing.T) {
	resp := NewCEchoResponseSuccess(100)

	if resp.MessageIDBeingRespondedTo() != 100 {
		t.Errorf("Expected MessageIDBeingRespondedTo 100, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandCEchoRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCEchoRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	if resp.IsPending() {
		t.Error("Expected IsPending() to be false")
	}

	if resp.IsFailure() {
		t.Error("Expected IsFailure() to be false")
	}
}

// Test C-STORE
func TestNewCStoreRequest(t *testing.T) {
	// Create a DICOM dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"})) // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))

	req, err := NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE-RQ: %v", err)
	}

	req.SetMessageID(101)

	if req.MessageID() != 101 {
		t.Errorf("Expected MessageID 101, got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandCStoreRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCStoreRQ, req.CommandField())
	}

	if req.AffectedSOPClassUID() != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("Wrong SOP Class UID: %s", req.AffectedSOPClassUID())
	}

	if req.AffectedSOPInstanceUID() != "1.2.3.4.5.6.7.8.9" {
		t.Errorf("Wrong SOP Instance UID: %s", req.AffectedSOPInstanceUID())
	}

	if req.DataDataset() == nil {
		t.Error("Expected data dataset to be present")
	}
}

func TestCStoreRequest_SetPriority(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

	req, _ := NewCStoreRequest(ds)
	req.SetPriority(uint16(PriorityHigh))

	if req.Priority() != uint16(PriorityHigh) {
		t.Errorf("Expected Priority %d, got %d", PriorityHigh, req.Priority())
	}
}

func TestCStoreRequest_SetMoveOriginator(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

	req, _ := NewCStoreRequest(ds)
	req.SetMoveOriginator("MOVE_SCU", 999)

	// Verify move originator fields are in command dataset
	moveAET, ok := req.CommandDataset().GetString(tag.MoveOriginatorApplicationEntityTitle)
	if !ok {
		t.Errorf("MoveOriginatorAET not found")
	}
	if moveAET != "MOVE_SCU" {
		t.Errorf("Expected MoveOriginatorAET 'MOVE_SCU', got %s", moveAET)
	}

	moveMsgID, err := req.CommandDataset().GetUInt16(tag.MoveOriginatorMessageID, 0)
	if err != nil {
		t.Errorf("MoveOriginatorMessageID not found: %v", err)
	}
	if moveMsgID != 999 {
		t.Errorf("Expected MoveOriginatorMessageID 999, got %d", moveMsgID)
	}
}

func TestNewCStoreResponse(t *testing.T) {
	resp := NewCStoreResponseSuccess(200, "1.2.840.10008.5.1.4.1.1.2", "1.2.3.4.5.6.7.8.9")

	if resp.MessageIDBeingRespondedTo() != 200 {
		t.Errorf("Expected MessageIDBeingRespondedTo 200, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandCStoreRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCStoreRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}
}

func TestNewCStoreRequest_MissingSOPClass(t *testing.T) {
	// Dataset without SOP Class UID
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

	_, err := NewCStoreRequest(ds)
	if err == nil {
		t.Error("Expected error for missing SOP Class UID")
	}
}

func TestNewCStoreRequest_MissingSOPInstance(t *testing.T) {
	// Dataset without SOP Instance UID
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))

	_, err := NewCStoreRequest(ds)
	if err == nil {
		t.Error("Expected error for missing SOP Instance UID")
	}
}

// Test C-FIND
func TestNewCFindRequest(t *testing.T) {
	query := dataset.New()
	query.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^*"}))
	query.Add(element.NewString(tag.PatientID, vr.LO, []string{""}))

	req := NewCFindRequestPatientRoot(QueryRetrieveLevelPatient, query)

	// MessageID should be 0 (unassigned) initially
	if req.MessageID() != 0 {
		t.Errorf("Expected MessageID 0 (unassigned), got %d", req.MessageID())
	}

	if req.CommandField() != uint16(CommandCFindRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCFindRQ, req.CommandField())
	}

	if req.AffectedSOPClassUID() != "1.2.840.10008.5.1.4.1.2.1.1" {
		t.Errorf("Wrong SOP Class UID: %s", req.AffectedSOPClassUID())
	}

	if req.QueryLevel() != QueryRetrieveLevelPatient {
		t.Errorf("Expected PATIENT level, got %s", req.QueryLevel())
	}

	if req.DataDataset() == nil {
		t.Error("Expected query dataset to be present")
	}

	// Verify QueryRetrieveLevel was added to dataset
	level, ok := query.GetString(tag.QueryRetrieveLevel)
	if !ok {
		t.Errorf("QueryRetrieveLevel not found")
	}
	if level != "PATIENT" {
		t.Errorf("Expected QueryRetrieveLevel 'PATIENT', got %s", level)
	}
}

func TestNewCFindRequestStudyRoot(t *testing.T) {
	query := dataset.New()
	req := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, query)

	if req.AffectedSOPClassUID() != "1.2.840.10008.5.1.4.1.2.2.1" {
		t.Errorf("Wrong Study Root SOP Class UID: %s", req.AffectedSOPClassUID())
	}
}

func TestCFindRequest_SetPriority(t *testing.T) {
	query := dataset.New()
	req := NewCFindRequestPatientRoot(QueryRetrieveLevelPatient, query)

	req.SetPriority(uint16(PriorityLow))
	if req.Priority() != uint16(PriorityLow) {
		t.Errorf("Expected Priority %d, got %d", PriorityLow, req.Priority())
	}
}

func TestNewCFindResponsePending(t *testing.T) {
	identifier := dataset.New()
	identifier.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
	identifier.Add(element.NewString(tag.PatientID, vr.LO, []string{"123456"}))

	resp := NewCFindResponsePending(300, "1.2.840.10008.5.1.4.1.2.1.1", identifier)

	if resp.MessageIDBeingRespondedTo() != 300 {
		t.Errorf("Expected MessageIDBeingRespondedTo 300, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.CommandField() != uint16(CommandCFindRSP) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCFindRSP, resp.CommandField())
	}

	if resp.StatusCode() != 0xFF00 {
		t.Errorf("Expected Pending status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsPending() {
		t.Error("Expected IsPending() to be true")
	}

	if resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be false for pending")
	}

	if !resp.HasIdentifier() {
		t.Error("Expected HasIdentifier() to be true")
	}

	if resp.DataDataset() == nil {
		t.Error("Expected identifier dataset to be present")
	}
}

func TestNewCFindResponseSuccess(t *testing.T) {
	resp := NewCFindResponseSuccess(300, "1.2.840.10008.5.1.4.1.2.1.1")

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected Success status, got 0x%04X", resp.StatusCode())
	}

	if !resp.IsSuccess() {
		t.Error("Expected IsSuccess() to be true")
	}

	if resp.IsPending() {
		t.Error("Expected IsPending() to be false")
	}

	if resp.HasIdentifier() {
		t.Error("Expected HasIdentifier() to be false for final response")
	}

	if resp.DataDataset() != nil {
		t.Error("Expected no identifier dataset in final response")
	}
}

// Test CreateCommandDataset
func TestCreateCommandDataset(t *testing.T) {
	ds := CreateCommandDataset(uint16(CommandCEchoRQ), 123)

	// Verify CommandField
	cmdField, err := ds.GetUInt16(tag.CommandField, 0)
	if err != nil {
		t.Errorf("CommandField not found: %v", err)
	}
	if cmdField != uint16(CommandCEchoRQ) {
		t.Errorf("Expected CommandField 0x%04X, got 0x%04X", CommandCEchoRQ, cmdField)
	}

	// Verify MessageID
	msgID, err := ds.GetUInt16(tag.MessageID, 0)
	if err != nil {
		t.Errorf("MessageID not found: %v", err)
	}
	if msgID != 123 {
		t.Errorf("Expected MessageID 123, got %d", msgID)
	}

	// Verify CommandDataSetType
	dataSetType, err := ds.GetUInt16(tag.CommandDataSetType, 0)
	if err != nil {
		t.Errorf("CommandDataSetType not found: %v", err)
	}
	if dataSetType != 0x0101 {
		t.Errorf("Expected CommandDataSetType 0x0101, got 0x%04X", dataSetType)
	}
}
