// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import "testing"

// TestNewCEchoResponseNilStatusMeansSuccess verifies that a nil status produces
// the only status a C-ECHO-RSP may carry (PS3.7 9.3.5.2) instead of the generic
// 0x0110 failure that the service encoder rejects as unsendable.
func TestNewCEchoResponseNilStatusMeansSuccess(t *testing.T) {
	resp := NewCEchoResponse(7, nil)
	if !resp.Status().IsSuccess() {
		t.Fatalf("Status() = %s, want Success", resp.Status())
	}
	if got := resp.StatusCode(); got != 0x0000 {
		t.Fatalf("StatusCode() = 0x%04X, want 0x0000", got)
	}
	if got := resp.MessageIDBeingRespondedTo(); got != 7 {
		t.Fatalf("MessageIDBeingRespondedTo() = %d, want 7", got)
	}
}
