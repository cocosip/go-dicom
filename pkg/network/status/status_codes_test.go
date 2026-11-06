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
}
