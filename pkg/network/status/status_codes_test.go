// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package status

import "testing"

// TestStatus_CorrectedCodes locks in the DICOM PS3.7 values for a couple of
// generic status variables that were previously mislabelled.
func TestStatus_CorrectedCodes(t *testing.T) {
	if got := RefusedMoveDestinationUnknown.Code; got != 0xA801 {
		t.Errorf("RefusedMoveDestinationUnknown.Code = 0x%04X, want 0xA801", got)
	}
	if got := ErrorDataSetDoesNotMatchSOPClass.Code; got != 0xA900 {
		t.Errorf("ErrorDataSetDoesNotMatchSOPClass.Code = 0x%04X, want 0xA900", got)
	}
	if got := WarningRequestedOptionalAttributesNotSupported; got.Code != 0x0001 || got.State != StateWarning {
		t.Errorf("WarningRequestedOptionalAttributesNotSupported = 0x%04X/%s, want 0x0001/%s",
			got.Code, got.State, StateWarning)
	}
	if got := RefusedSOPClassNotSupported; got.Code != 0x0122 || got.Description != "Refused: SOP class not supported" {
		t.Errorf("RefusedSOPClassNotSupported = 0x%04X %q, want 0x0122 \"Refused: SOP class not supported\"",
			got.Code, got.Description)
	}
	if got := RefusedNotAuthorized; got.Code != 0x0124 || got.Description != "Refused: Not authorized" {
		t.Errorf("RefusedNotAuthorized = 0x%04X %q, want 0x0124 \"Refused: Not authorized\"",
			got.Code, got.Description)
	}
	if got := ErrorProcessingFailure.Code; got != 0x0110 {
		t.Errorf("ErrorProcessingFailure.Code = 0x%04X, want 0x0110", got)
	}
	if got := NSetFailureMissingAttribute; got.Code != 0x0120 || got.Description != "Missing attribute" {
		t.Errorf("NSetFailureMissingAttribute = 0x%04X %q, want 0x0120 \"Missing attribute\"",
			got.Code, got.Description)
	}
	if NSetFailureMissingAttribute == ErrorMissingAttribute {
		t.Error("NSetFailureMissingAttribute must not alias ErrorMissingAttribute with a different description")
	}
}
