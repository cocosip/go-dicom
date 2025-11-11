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

const testStudyRootMoveUID = "1.2.840.10008.5.1.4.1.2.2.2"

func TestNewCMoveRequest(t *testing.T) {
    identifier := dataset.New()
    _ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	req := NewCMoveRequest(QueryRetrieveLevelStudy, "DEST_AE", identifier)

	if req == nil {
		t.Fatal("NewCMoveRequest returned nil")
	}

	if req.AffectedSOPClassUID() != testStudyRootMoveUID {
		t.Errorf("Expected SOP Class UID '1.2.840.10008.5.1.4.1.2.2.2', got '%s'", req.AffectedSOPClassUID())
	}

	if req.MoveDestination() != "DEST_AE" {
		t.Errorf("Expected move destination 'DEST_AE', got '%s'", req.MoveDestination())
	}

	if req.QueryLevel() != QueryRetrieveLevelStudy {
		t.Errorf("Expected query level 'STUDY', got '%s'", req.QueryLevel())
	}

	if req.Priority() != uint16(PriorityMedium) {
		t.Errorf("Expected priority %d, got %d", PriorityMedium, req.Priority())
	}
}

func TestNewCMoveRequestPatientRoot(t *testing.T) {
    identifier := dataset.New()
    _ = identifier.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	req := NewCMoveRequestPatientRoot(QueryRetrieveLevelPatient, "DEST_AE", identifier)

	if req.AffectedSOPClassUID() != "1.2.840.10008.5.1.4.1.2.1.2" {
		t.Errorf("Expected Patient Root SOP Class UID, got '%s'", req.AffectedSOPClassUID())
	}
}

func TestCMoveRequest_SetPriority(t *testing.T) {
	identifier := dataset.New()
	req := NewCMoveRequest(QueryRetrieveLevelStudy, "DEST_AE", identifier)

	req.SetPriority(uint16(PriorityHigh))

	if req.Priority() != uint16(PriorityHigh) {
		t.Errorf("Expected priority %d, got %d", PriorityHigh, req.Priority())
	}
}

func TestNewCMoveResponse(t *testing.T) {
	resp := NewCMoveResponse(123, 0x0000, testStudyRootMoveUID)

	if resp == nil {
		t.Fatal("NewCMoveResponse returned nil")
	}

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected status code 0x0000, got 0x%04X", resp.StatusCode())
	}

	if resp.MessageIDBeingRespondedTo() != 123 {
		t.Errorf("Expected message ID 123, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.AffectedSOPClassUID() != testStudyRootMoveUID {
		t.Errorf("Expected SOP Class UID '1.2.840.10008.5.1.4.1.2.2.2', got '%s'", resp.AffectedSOPClassUID())
	}
}

func TestNewCMoveResponsePending(t *testing.T) {
	resp := NewCMoveResponsePending(123, testStudyRootMoveUID, 10, 5, 1, 2)

	if resp.StatusCode() != 0xFF00 {
		t.Errorf("Expected status code 0xFF00, got 0x%04X", resp.StatusCode())
	}

	if !resp.HasSubOperationCounts() {
		t.Error("Expected response to have sub-operation counts")
	}

	if resp.NumberOfRemainingSubOperations() != 10 {
		t.Errorf("Expected 10 remaining, got %d", resp.NumberOfRemainingSubOperations())
	}

	if resp.NumberOfCompletedSubOperations() != 5 {
		t.Errorf("Expected 5 completed, got %d", resp.NumberOfCompletedSubOperations())
	}

	if resp.NumberOfFailedSubOperations() != 1 {
		t.Errorf("Expected 1 failed, got %d", resp.NumberOfFailedSubOperations())
	}

	if resp.NumberOfWarningSubOperations() != 2 {
		t.Errorf("Expected 2 warning, got %d", resp.NumberOfWarningSubOperations())
	}
}

func TestNewCMoveResponseSuccess(t *testing.T) {
	resp := NewCMoveResponseSuccess(123, testStudyRootMoveUID)

	if resp.StatusCode() != 0x0000 {
		t.Errorf("Expected status code 0x0000, got 0x%04X", resp.StatusCode())
	}

	if resp.HasSubOperationCounts() {
		t.Error("Success response should not have sub-operation counts")
	}
}

func TestNewCMoveResponseFromRequest(t *testing.T) {
	identifier := dataset.New()
    req := NewCMoveRequest(QueryRetrieveLevelStudy, "DEST_AE", identifier)
    _ = req.SetMessageID(456)

	resp := NewCMoveResponseFromRequest(req, 0x0000)

	if resp.MessageIDBeingRespondedTo() != 456 {
		t.Errorf("Expected message ID 456, got %d", resp.MessageIDBeingRespondedTo())
	}

	if resp.AffectedSOPClassUID() != req.AffectedSOPClassUID() {
		t.Errorf("Expected SOP Class UID '%s', got '%s'", req.AffectedSOPClassUID(), resp.AffectedSOPClassUID())
	}
}

func TestCMoveRequest_String(t *testing.T) {
	identifier := dataset.New()
    req := NewCMoveRequest(QueryRetrieveLevelStudy, "DEST_AE", identifier)
    _ = req.SetMessageID(789)

	str := req.String()
	if str == "" {
		t.Error("String() returned empty string")
	}

	// Should contain key information
	if !contains(str, "C-MOVE-RQ") {
		t.Error("String should contain 'C-MOVE-RQ'")
	}
	if !contains(str, "DEST_AE") {
		t.Error("String should contain destination AE title")
	}
}

func TestCMoveResponse_String(t *testing.T) {
	resp := NewCMoveResponsePending(123, testStudyRootMoveUID, 10, 5, 1, 2)

	str := resp.String()
	if str == "" {
		t.Error("String() returned empty string")
	}

	// Should contain key information
	if !contains(str, "C-MOVE-RSP") {
		t.Error("String should contain 'C-MOVE-RSP'")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
