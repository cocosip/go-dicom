// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package status

import (
	"testing"
)

func TestStatus_String(t *testing.T) {
	s := NewStatus(0x0000, StateSuccess, StateSuccess)
	expected := "0x0000 (Success): Success"
	if s.String() != expected {
		t.Errorf("Expected %s, got %s", expected, s.String())
	}
}

func TestStatus_IsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		code     uint16
		state    string
		expected bool
	}{
		{StateSuccess, 0x0000, StateSuccess, true},
		{StatePending, 0xFF00, StatePending, false},
		{StateFailure, 0x0110, StateFailure, false},
		{"Warning", 0x0107, StateWarning, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, tt.state, "")
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
		state    string
		expected bool
	}{
		{StatePending, 0xFF00, StatePending, true},
		{"PendingWarning", 0xFF01, StatePending, true},
		{StateSuccess, 0x0000, StateSuccess, false},
		{StateFailure, 0x0110, StateFailure, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, tt.state, "")
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
		state    string
		expected bool
	}{
		{"Warning 0x0001", 0x0001, StateWarning, true},
		{"Warning 0x0107", 0x0107, StateWarning, true},
		{"Warning 0x0116", 0x0116, StateWarning, true},
		{"Warning 0xB006", 0xB006, StateWarning, true},
		{StateSuccess, 0x0000, StateSuccess, false},
		{"Failure 0x0110", 0x0110, StateFailure, false},
		{StatePending, 0xFF00, StatePending, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, tt.state, "")
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
		state    string
		expected bool
	}{
		{"Warning 0x0107", 0x0107, StateWarning, false},
		{"Failure 0x0110", 0x0110, StateFailure, true},
		{"Failure 0xA700", 0xA700, StateFailure, true},
		{"Failure 0xC000", 0xC000, StateFailure, true},
		{StateSuccess, 0x0000, StateSuccess, false},
		{"Failure 0x0122", 0x0122, StateFailure, true},
		{"Warning 0xB000", 0xB000, StateWarning, false},
		{StatePending, 0xFF00, StatePending, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, tt.state, "")
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
		state    string
		expected bool
	}{
		{StateCancel, 0xFE00, StateCancel, true},
		{StateSuccess, 0x0000, StateSuccess, false},
		{StatePending, 0xFF00, StatePending, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStatus(tt.code, tt.state, "")
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
		{StateSuccess, Success, 0x0000, StateSuccess},
		{StatePending, Pending, 0xFF00, StatePending},
		{StateCancel, Cancel, 0xFE00, StateCancel},
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
		{StateSuccess, 0x0000, StateSuccess},
		{StatePending, 0xFF00, StatePending},
		{"PendingWarning", 0xFF01, StatePending},
		{StateCancel, 0xFE00, StateCancel},
		{"Reserved 00xx classifies as failure", 0x0050, StateFailure},
		{"Requested optional Attributes warning", 0x0001, StateWarning},
		{"Attribute list warning", 0x0107, StateWarning},
		{"Attribute value warning", 0x0116, StateWarning},
		{"Refused SOP class not supported", 0x0122, StateFailure},
		{"Refused not authorized", 0x0124, StateFailure},
		{"Coercion warning", 0xB000, StateWarning},
		{"Unknown Bxxx warning", 0xB001, StateWarning},
		{StateFailure, 0x0110, StateFailure},
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

// TestLookupStatus_ReturnsNamedStatus verifies that named codes resolve to their
// package-level Status values instead of a generic re-created one.
func TestLookupStatus_ReturnsNamedStatus(t *testing.T) {
	tests := []struct {
		code uint16
		want *Status
	}{
		{0x0000, Success},
		{0x0001, WarningRequestedOptionalAttributesNotSupported},
		{0x0107, AttributeListError},
		{0x0110, ErrorProcessingFailure},
		{0x0122, RefusedSOPClassNotSupported},
		{0x0124, RefusedNotAuthorized},
		{0xFE00, Cancel},
		{0xFF00, Pending},
	}

	for _, tt := range tests {
		if got := LookupStatus(tt.code); got != tt.want {
			t.Errorf("LookupStatus(0x%04X) = 0x%04X %q, want package status 0x%04X %q",
				tt.code, got.Code, got.Description, tt.want.Code, tt.want.Description)
		}
	}
}

// TestLookupStatus_NoDuplicateCode guards against an alias reusing a code and
// silently changing LookupStatus output.
func TestLookupStatus_NoDuplicateCode(t *testing.T) {
	for code, s := range statusByCode {
		if s == nil {
			t.Fatalf("statusByCode[0x%04X] is nil", code)
		}
		if s.Code != code {
			t.Errorf("statusByCode[0x%04X] has Code 0x%04X", code, s.Code)
		}
	}
	if got := statusByCode[0x0110]; got != ErrorProcessingFailure {
		t.Errorf("statusByCode[0x0110] = %q, want ErrorProcessingFailure", got.Description)
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
