// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package status

import (
	"testing"
)

func TestStatus_String(t *testing.T) {
	s := NewStatus(0x0000, StateSuccess, "Success")
	expected := "0x0000 (Success): Success"
	if s.String() != expected {
		t.Errorf("Expected %s, got %s", expected, s.String())
	}
}

func TestStatus_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		expected bool
	}{
		{"Success", 0x0000, true},
		{"Pending", 0xFF00, false},
		{"Failure", 0x0110, false},
		{"Warning", 0x0107, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, "", "")
			if s.IsSuccess() != tt.expected {
				t.Errorf("Expected IsSuccess() = %v for code 0x%04X", tt.expected, tt.code)
			}
		})
	}
}

func TestStatus_IsPending(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		expected bool
	}{
		{"Pending", 0xFF00, true},
		{"PendingWarning", 0xFF01, true},
		{"Success", 0x0000, false},
		{"Failure", 0x0110, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, "", "")
			if s.IsPending() != tt.expected {
				t.Errorf("Expected IsPending() = %v for code 0x%04X", tt.expected, tt.code)
			}
		})
	}
}

func TestStatus_IsWarning(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		expected bool
	}{
		{"Warning 0x0001", 0x0001, true},
		{"Warning 0x00FF", 0x00FF, true},
		{"Warning 0xB000", 0xB000, true},
		{"Warning 0xB006", 0xB006, true},
		{"Success", 0x0000, false},
		{"Failure 0x0107", 0x0107, false},
		{"Failure 0x0110", 0x0110, false},
		{"Pending", 0xFF00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, "", "")
			if s.IsWarning() != tt.expected {
				t.Errorf("Expected IsWarning() = %v for code 0x%04X", tt.expected, tt.code)
			}
		})
	}
}

func TestStatus_IsFailure(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		expected bool
	}{
		{"Failure 0x0107", 0x0107, true},
		{"Failure 0x0110", 0x0110, true},
		{"Failure 0xA700", 0xA700, true},
		{"Failure 0xC000", 0xC000, true},
		{"Success", 0x0000, false},
		{"Warning 0x0001", 0x0001, false},
		{"Warning 0xB000", 0xB000, false},
		{"Pending", 0xFF00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, "", "")
			if s.IsFailure() != tt.expected {
				t.Errorf("Expected IsFailure() = %v for code 0x%04X", tt.expected, tt.code)
			}
		})
	}
}

func TestStatus_IsCancel(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		expected bool
	}{
		{"Cancel", 0xFE00, true},
		{"Success", 0x0000, false},
		{"Pending", 0xFF00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, "", "")
			if s.IsCancel() != tt.expected {
				t.Errorf("Expected IsCancel() = %v for code 0x%04X", tt.expected, tt.code)
			}
		})
	}
}

func TestCommonStatuses(t *testing.T) {
	tests := []struct {
		name   string
		status *Status
		code   uint16
		state  string
	}{
		{"Success", Success, 0x0000, StateSuccess},
		{"Pending", Pending, 0xFF00, StatePending},
		{"Cancel", Cancel, 0xFE00, StateCancel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.status.Code != tt.code {
				t.Errorf("Expected code 0x%04X, got 0x%04X", tt.code, tt.status.Code)
			}
			if tt.status.State != tt.state {
				t.Errorf("Expected state %s, got %s", tt.state, tt.status.State)
			}
		})
	}
}

func TestCStoreStatuses(t *testing.T) {
	// Test C-STORE specific statuses
	if CStoreRefusedOutOfResources.Code != 0xA700 {
		t.Errorf("Expected CStoreRefusedOutOfResources code 0xA700, got 0x%04X", CStoreRefusedOutOfResources.Code)
	}
	if CStoreErrorCannotUnderstand.Code != 0xC000 {
		t.Errorf("Expected CStoreErrorCannotUnderstand code 0xC000, got 0x%04X", CStoreErrorCannotUnderstand.Code)
	}
	if CStoreWarningCoercionOfDataElements.Code != 0xB000 {
		t.Errorf("Expected CStoreWarningCoercionOfDataElements code 0xB000, got 0x%04X", CStoreWarningCoercionOfDataElements.Code)
	}
}

func TestCFindStatuses(t *testing.T) {
	// Test C-FIND specific statuses
	if CFindPending.Code != 0xFF00 {
		t.Errorf("Expected CFindPending code 0xFF00, got 0x%04X", CFindPending.Code)
	}
	if CFindCancelMatchingTerminated.Code != 0xFE00 {
		t.Errorf("Expected CFindCancelMatchingTerminated code 0xFE00, got 0x%04X", CFindCancelMatchingTerminated.Code)
	}
}

func TestLookupStatus(t *testing.T) {
	tests := []struct {
		name          string
		code          uint16
		expectedState string
	}{
		{"Success", 0x0000, StateSuccess},
		{"Pending", 0xFF00, StatePending},
		{"PendingWarning", 0xFF01, StatePending},
		{"Cancel", 0xFE00, StateCancel},
		{"Warning", 0x0050, StateWarning},
		{"Failure", 0x0110, StateFailure},
		{"Unknown Failure", 0x1234, StateFailure},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LookupStatus(tt.code)
			if s.State != tt.expectedState {
				t.Errorf("Expected state %s for code 0x%04X, got %s", tt.expectedState, tt.code, s.State)
			}
			if s.Code != tt.code {
				t.Errorf("Expected code 0x%04X, got 0x%04X", tt.code, s.Code)
			}
		})
	}
}

func TestNOperationStatuses(t *testing.T) {
	// Test N-EVENT-REPORT
	if NEventReportSuccess.Code != 0x0000 {
		t.Errorf("Expected NEventReportSuccess code 0x0000, got 0x%04X", NEventReportSuccess.Code)
	}

	// Test N-GET
	if NGetFailureNoSuchObjectInstance.Code != 0x0112 {
		t.Errorf("Expected NGetFailureNoSuchObjectInstance code 0x0112, got 0x%04X", NGetFailureNoSuchObjectInstance.Code)
	}

	// Test N-CREATE
	if NCreateFailureDuplicateSOPInstance.Code != 0x0111 {
		t.Errorf("Expected NCreateFailureDuplicateSOPInstance code 0x0111, got 0x%04X", NCreateFailureDuplicateSOPInstance.Code)
	}
}
