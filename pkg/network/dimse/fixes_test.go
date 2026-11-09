// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// TestRequestSetAffectedSOPClassUIDUpdatesCacheAndString verifies the setter
// updates both the command element (used by AffectedSOPClassUID) and the cached
// field (used by String) for the query/retrieve requests.
func TestRequestSetAffectedSOPClassUIDUpdatesCacheAndString(t *testing.T) {
	const newUID = "1.2.840.10008.9.9.9"

	check := func(t *testing.T, name string, req Request, set func(string), str func() string) {
		t.Helper()
		set(newUID)
		if got := req.AffectedSOPClassUID(); got != newUID {
			t.Errorf("%s: AffectedSOPClassUID() = %q, want %q", name, got, newUID)
		}
		if got, ok := req.CommandDataset().GetString(tag.AffectedSOPClassUID); !ok || got != newUID {
			t.Errorf("%s: command AffectedSOPClassUID = %q (ok=%v), want %q", name, got, ok, newUID)
		}
		if got := str(); !strings.Contains(got, newUID) {
			t.Errorf("%s: String() = %q, want it to contain %q", name, got, newUID)
		}
	}

	find := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	check(t, "C-FIND", find, find.SetAffectedSOPClassUID, find.String)

	get := NewCGetRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	check(t, "C-GET", get, get.SetAffectedSOPClassUID, get.String)

	move := NewCMoveRequestStudyRoot(QueryRetrieveLevelStudy, "DEST_AE", dataset.New())
	check(t, "C-MOVE", move, move.SetAffectedSOPClassUID, move.String)
}

// TestCStoreResponseOmitsEmptyUIDs verifies (0000,0002)/(0000,1000) are only
// written when non-empty (Type U in PS3.7 Table 9.3-3).
func TestCStoreResponseOmitsEmptyUIDs(t *testing.T) {
	minimal := NewCStoreResponse(7, status.Success, "", "")
	if minimal.CommandDataset().Contains(tag.AffectedSOPClassUID) {
		t.Error("empty C-STORE-RSP must omit Affected SOP Class UID")
	}
	if minimal.CommandDataset().Contains(tag.AffectedSOPInstanceUID) {
		t.Error("empty C-STORE-RSP must omit Affected SOP Instance UID")
	}
	if got := minimal.AffectedSOPClassUID(); got != "" {
		t.Errorf("AffectedSOPClassUID() = %q, want empty", got)
	}
	if got := minimal.AffectedSOPInstanceUID(); got != "" {
		t.Errorf("AffectedSOPInstanceUID() = %q, want empty", got)
	}

	full := NewCStoreResponse(7, status.Success, testCTImageStorageUID, testSOPInstanceUID)
	if got, ok := full.CommandDataset().GetString(tag.AffectedSOPClassUID); !ok || got != testCTImageStorageUID {
		t.Errorf("Affected SOP Class UID = %q (ok=%v), want %q", got, ok, testCTImageStorageUID)
	}
	if got, ok := full.CommandDataset().GetString(tag.AffectedSOPInstanceUID); !ok || got != testSOPInstanceUID {
		t.Errorf("Affected SOP Instance UID = %q (ok=%v), want %q", got, ok, testSOPInstanceUID)
	}

	// FromRequest propagates the request UIDs, including their absence.
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))
	req, err := NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest: %v", err)
	}
	req.SetAffectedSOPClassUID("")
	req.SetAffectedSOPInstanceUID("")
	fromReq := NewCStoreResponseFromRequest(req, status.Success)
	if fromReq.CommandDataset().Contains(tag.AffectedSOPClassUID) ||
		fromReq.CommandDataset().Contains(tag.AffectedSOPInstanceUID) {
		t.Error("C-STORE-RSP from a request without UIDs must omit both UID elements")
	}
}

// TestCFindResponseOmitsEmptyAffectedSOPClassUID verifies (0000,0002) is only
// written when non-empty (Type U in PS3.7 Table 9.3-5).
func TestCFindResponseOmitsEmptyAffectedSOPClassUID(t *testing.T) {
	minimal := NewCFindResponse(9, status.Success, "", nil)
	if minimal.CommandDataset().Contains(tag.AffectedSOPClassUID) {
		t.Error("empty C-FIND-RSP must omit Affected SOP Class UID")
	}
	if got := minimal.AffectedSOPClassUID(); got != "" {
		t.Errorf("AffectedSOPClassUID() = %q, want empty", got)
	}

	const uid = "1.2.840.10008.5.1.4.1.2.1.1"
	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	full := NewCFindResponse(9, status.CFindPending, uid, identifier)
	if got, ok := full.CommandDataset().GetString(tag.AffectedSOPClassUID); !ok || got != uid {
		t.Errorf("Affected SOP Class UID = %q (ok=%v), want %q", got, ok, uid)
	}

	req := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	req.SetAffectedSOPClassUID("")
	fromReq := NewCFindResponseFromRequest(req, status.Success, nil)
	if fromReq.CommandDataset().Contains(tag.AffectedSOPClassUID) {
		t.Error("C-FIND-RSP from a request without a UID must omit Affected SOP Class UID")
	}
}

// TestNewCStoreRequestRejectsPresentButEmptyUIDs verifies a zero-length
// SOPClassUID/SOPInstanceUID element is rejected, not just a missing one.
func TestNewCStoreRequestRejectsPresentButEmptyUIDs(t *testing.T) {
	tests := []struct {
		name       string
		classUID   string
		instanceID string
		wantErr    bool
	}{
		{"both set", testCTImageStorageUID, testSOPInstanceUID, false},
		{"empty class", "", testSOPInstanceUID, true},
		{"empty instance", testCTImageStorageUID, "", true},
		{"both empty", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{tt.classUID}))
			_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{tt.instanceID}))
			_, err := NewCStoreRequest(ds)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewCStoreRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestBaseResponseStatusMissingIsFailure verifies a response without a Status
// element is classified as Failure (PS3.7 Annex C) instead of Unknown.
func TestBaseResponseStatusMissingIsFailure(t *testing.T) {
	command := CreateCommandDataset(uint16(CommandCStoreRSP), 0)
	resp := NewBaseResponse(command, nil)

	st := resp.Status()
	if st.Code != 0xFFFF {
		t.Errorf("missing Status code = 0x%04X, want 0xFFFF", st.Code)
	}
	if !st.IsFailure() || st.State != status.StateFailure {
		t.Errorf("missing Status state = %q, want %q", st.State, status.StateFailure)
	}
	if !resp.IsFailure() {
		t.Error("IsFailure() = false, want true")
	}
	if resp.IsSuccess() || resp.IsPending() {
		t.Error("missing Status must not be reported as success or pending")
	}
}
