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
)

func TestValidateRequestForSend(t *testing.T) {
	const (
		sopClassUID = "1.2.840.10008.5.1.4.1.1.1"
		instanceUID = "1.2.3.4.5.6"
		opEvent     = "N-EVENT-REPORT"
	)

	newCStore := func() *CStoreRequest {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{sopClassUID}))
		_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{instanceUID}))
		req, err := NewCStoreRequest(ds)
		if err != nil {
			t.Fatalf("NewCStoreRequest: %v", err)
		}
		return req
	}

	storeMissingClass := newCStore()
	storeMissingClass.SetAffectedSOPClassUID("")
	storeMissingInstance := newCStore()
	storeMissingInstance.SetAffectedSOPInstanceUID("")

	find := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	findMissingClass := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	findMissingClass.SetAffectedSOPClassUID("")

	get := NewCGetRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	getMissingClass := NewCGetRequestStudyRoot(QueryRetrieveLevelStudy, dataset.New())
	getMissingClass.SetAffectedSOPClassUID("")

	move := NewCMoveRequestStudyRoot(QueryRetrieveLevelStudy, "DEST_AE", dataset.New())
	moveMissingClass := NewCMoveRequestStudyRoot(QueryRetrieveLevelStudy, "DEST_AE", dataset.New())
	moveMissingClass.SetAffectedSOPClassUID("")

	event := NewNEventReportRequest(sopClassUID, instanceUID, 7, nil)
	eventZeroType := NewNEventReportRequest(sopClassUID, instanceUID, 0, nil)
	eventMissingType := NewNEventReportRequest(sopClassUID, instanceUID, 7, nil)
	eventMissingType.CommandDataset().Remove(tag.EventTypeID)

	tests := []struct {
		name    string
		req     Request
		wantErr bool
		wantOp  string
	}{
		{"C-ECHO passes", NewCEchoRequest(), false, ""},
		{"C-CANCEL passes", NewCCancelRequest(1), false, ""},
		{"nil request fails", nil, true, "dimse"},
		{"C-STORE valid", newCStore(), false, ""},
		{"C-STORE missing class", storeMissingClass, true, "C-STORE"},
		{"C-STORE missing instance", storeMissingInstance, true, "C-STORE"},
		{"C-FIND valid", find, false, ""},
		{"C-FIND missing class", findMissingClass, true, "C-FIND"},
		{"C-GET valid", get, false, ""},
		{"C-GET missing class", getMissingClass, true, "C-GET"},
		{"C-MOVE valid", move, false, ""},
		{"C-MOVE missing class", moveMissingClass, true, "C-MOVE"},
		{"C-MOVE missing destination", NewCMoveRequestStudyRoot(QueryRetrieveLevelStudy, "", dataset.New()), true, "C-MOVE"},
		{"N-GET valid", NewNGetRequest(sopClassUID, instanceUID, nil), false, ""},
		{"N-GET missing class", NewNGetRequest("", instanceUID, nil), true, "N-GET"},
		{"N-GET missing instance", NewNGetRequest(sopClassUID, "", nil), true, "N-GET"},
		{"N-SET valid", NewNSetRequest(sopClassUID, instanceUID, nil), false, ""},
		{"N-SET missing class", NewNSetRequest("", instanceUID, nil), true, "N-SET"},
		{"N-SET missing instance", NewNSetRequest(sopClassUID, "", nil), true, "N-SET"},
		{"N-DELETE valid", NewNDeleteRequest(sopClassUID, instanceUID), false, ""},
		{"N-DELETE missing class", NewNDeleteRequest("", instanceUID), true, "N-DELETE"},
		{"N-DELETE missing instance", NewNDeleteRequest(sopClassUID, ""), true, "N-DELETE"},
		{"N-ACTION valid", NewNActionRequest(sopClassUID, instanceUID, 1, nil), false, ""},
		{"N-ACTION missing class", NewNActionRequest("", instanceUID, 1, nil), true, "N-ACTION"},
		{"N-ACTION missing instance", NewNActionRequest(sopClassUID, "", 1, nil), true, "N-ACTION"},
		{"N-CREATE valid", NewNCreateRequest(sopClassUID, instanceUID, nil), false, ""},
		{"N-CREATE server-assigned instance", NewNCreateRequest(sopClassUID, "", nil), false, ""},
		{"N-CREATE missing class", NewNCreateRequest("", instanceUID, nil), true, "N-CREATE"},
		{"N-EVENT-REPORT valid", event, false, ""},
		{"N-EVENT-REPORT zero type ID", eventZeroType, false, ""},
		{"N-EVENT-REPORT missing class", NewNEventReportRequest("", instanceUID, 1, nil), true, opEvent},
		{"N-EVENT-REPORT missing instance", NewNEventReportRequest(sopClassUID, "", 1, nil), true, opEvent},
		{"N-EVENT-REPORT missing type ID", eventMissingType, true, opEvent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequestForSend(tt.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateRequestForSend() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantOp != "" && !strings.Contains(err.Error(), tt.wantOp) {
				t.Errorf("error %q does not name operation %q", err, tt.wantOp)
			}
		})
	}
}
