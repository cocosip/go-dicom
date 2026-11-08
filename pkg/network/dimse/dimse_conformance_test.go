// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package dimse

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func commandUInt16(t *testing.T, ds *dataset.Dataset, tg *tag.Tag) uint16 {
	t.Helper()
	v, err := ds.GetUInt16(tg, 0)
	if err != nil {
		t.Fatalf("command element %s missing: %v", tg, err)
	}
	return v
}

// TestCEchoRequestHasNoPriorityElement verifies the C-ECHO-RQ command set does
// not carry a Priority element (PS3.7 defines none for C-ECHO).
func TestCEchoRequestHasNoPriorityElement(t *testing.T) {
	req := NewCEchoRequest()
	if _, ok := req.CommandDataset().Get(tag.Priority); ok {
		t.Error("C-ECHO-RQ must not contain a Priority element")
	}
	if got := commandUInt16(t, req.CommandDataset(), tag.CommandDataSetType); got != 0x0101 {
		t.Errorf("C-ECHO-RQ CommandDataSetType = 0x%04X, want 0x0101", got)
	}
}

// TestNOperationResponsesKeepZeroTypeID verifies presence of the required Type
// ID is not inferred from its value: 0 is a valid Type ID.
func TestNOperationResponsesKeepZeroTypeID(t *testing.T) {
	const uid = "1.2.840.10008.5.1.4.1.1.1"

	action := NewNActionResponse(1, status.NActionSuccess, uid, uid, 0, nil)
	if _, ok := action.CommandDataset().Get(tag.ActionTypeID); !ok {
		t.Error("N-ACTION-RSP must contain Action Type ID even when it is 0")
	}
	if got := action.ActionTypeID(); got != 0 {
		t.Errorf("ActionTypeID = %d, want 0", got)
	}

	event := NewNEventReportResponse(2, status.NEventReportSuccess, uid, uid, 0, nil)
	if _, ok := event.CommandDataset().Get(tag.EventTypeID); !ok {
		t.Error("N-EVENT-REPORT-RSP must contain Event Type ID even when it is 0")
	}
	if got := event.EventTypeID(); got != 0 {
		t.Errorf("EventTypeID = %d, want 0", got)
	}
}

// TestCFindResponseIdentifierOnlyWhenPending verifies the Identifier is dropped
// for non-Pending statuses so the command flag and data set stay consistent.
func TestCFindResponseIdentifierOnlyWhenPending(t *testing.T) {
	const uid = "1.2.840.10008.5.1.4.1.2.1.1"
	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))

	final := NewCFindResponse(1, status.Success, uid, identifier)
	if final.DataDataset() != nil || final.HasIdentifier() {
		t.Error("final C-FIND-RSP must not attach the identifier")
	}
	if got := commandUInt16(t, final.CommandDataset(), tag.CommandDataSetType); got != 0x0101 {
		t.Errorf("final C-FIND-RSP CommandDataSetType = 0x%04X, want 0x0101", got)
	}

	pending := NewCFindResponse(1, status.CFindPending, uid, identifier)
	if pending.DataDataset() == nil || !pending.HasIdentifier() {
		t.Error("pending C-FIND-RSP must attach the identifier")
	}
	if got := commandUInt16(t, pending.CommandDataset(), tag.CommandDataSetType); got != 0x0001 {
		t.Errorf("pending C-FIND-RSP CommandDataSetType = 0x%04X, want 0x0001", got)
	}

	// PendingWarning may also carry an identifier.
	pendingWarning := NewCFindResponse(1, status.CFindPendingWarningOptionalKeysNotSupported, uid, identifier)
	if pendingWarning.DataDataset() == nil {
		t.Error("pending warning C-FIND-RSP must attach the identifier")
	}
}

// TestNResponsesAttachAttributeListOnlyOnSuccess verifies N-GET/N-SET/N-CREATE
// responses never attach data while advertising CommandDataSetType = 0x0101.
func TestNResponsesAttachAttributeListOnlyOnSuccess(t *testing.T) {
	const uid = "1.2.840.10008.5.1.4.1.1.1"
	attributeList := dataset.New()
	_ = attributeList.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))

	check := func(name string, cmd *dataset.Dataset, data *dataset.Dataset, wantData bool) {
		t.Helper()
		present := commandUInt16(t, cmd, tag.CommandDataSetType) == 0x0001
		if present != wantData {
			t.Errorf("%s: CommandDataSetType present=%v, want %v", name, present, wantData)
		}
		if (data != nil) != wantData {
			t.Errorf("%s: attached data=%v, want %v", name, data != nil, wantData)
		}
	}

	getOK := NewNGetResponse(1, status.NGetSuccess, uid, uid, attributeList)
	check("N-GET success", getOK.CommandDataset(), getOK.DataDataset(), true)
	getWarn := NewNGetResponse(2, status.NGetWarningAttributeListError, uid, uid, attributeList)
	// PS3.7 10.1.2 + Annex C: a 0x0107 Attribute List Error still returns the
	// attribute list that could be read.
	check("N-GET warning", getWarn.CommandDataset(), getWarn.DataDataset(), true)

	setOK := NewNSetResponse(3, status.NSetSuccess, uid, uid, attributeList)
	check("N-SET success", setOK.CommandDataset(), setOK.DataDataset(), true)
	setFail := NewNSetResponse(4, status.NSetFailureProcessingFailure, uid, uid, attributeList)
	check("N-SET failure", setFail.CommandDataset(), setFail.DataDataset(), false)

	createOK := NewNCreateResponse(5, status.NCreateSuccess, uid, uid, attributeList)
	check("N-CREATE success", createOK.CommandDataset(), createOK.DataDataset(), true)
	createWarn := NewNCreateResponse(6, status.NCreateWarningAttributeListError, uid, uid, attributeList)
	check("N-CREATE warning", createWarn.CommandDataset(), createWarn.DataDataset(), false)
}

// TestQueryRetrieveRequest_LevelWinsOverStaleValue verifies the constructor
// level replaces a stale QueryRetrieveLevel already present in the identifier.
func TestQueryRetrieveRequest_LevelWinsOverStaleValue(t *testing.T) {
	newStale := func() *dataset.Dataset {
		ds := dataset.New()
		_ = ds.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{"STUDY"}))
		return ds
	}

	findDS := newStale()
	NewCFindRequestStudyRoot(QueryRetrieveLevelSeries, findDS)
	if got, _ := findDS.GetString(tag.QueryRetrieveLevel); got != "SERIES" {
		t.Errorf("C-FIND QueryRetrieveLevel = %q, want SERIES", got)
	}

	getDS := newStale()
	NewCGetRequest(QueryRetrieveLevelImage, getDS)
	if got, _ := getDS.GetString(tag.QueryRetrieveLevel); got != "IMAGE" {
		t.Errorf("C-GET QueryRetrieveLevel = %q, want IMAGE", got)
	}

	moveDS := newStale()
	NewCMoveRequest(QueryRetrieveLevelSeries, "DEST_AE", moveDS)
	if got, _ := moveDS.GetString(tag.QueryRetrieveLevel); got != "SERIES" {
		t.Errorf("C-MOVE QueryRetrieveLevel = %q, want SERIES", got)
	}
}

// TestQueryRetrieveRequest_NilIdentifier verifies the command flag is 0x0101
// when there is no identifier to attach.
func TestQueryRetrieveRequest_NilIdentifier(t *testing.T) {
	find := NewCFindRequestStudyRoot(QueryRetrieveLevelStudy, nil)
	if got := commandUInt16(t, find.CommandDataset(), tag.CommandDataSetType); got != 0x0101 {
		t.Errorf("C-FIND with nil identifier CommandDataSetType = 0x%04X, want 0x0101", got)
	}

	move := NewCMoveRequest(QueryRetrieveLevelStudy, "DEST_AE", nil)
	if got := commandUInt16(t, move.CommandDataset(), tag.CommandDataSetType); got != 0x0101 {
		t.Errorf("C-MOVE with nil identifier CommandDataSetType = 0x%04X, want 0x0101", got)
	}
	if move.DataDataset() != nil {
		t.Error("C-MOVE with nil identifier must not attach a data set")
	}
}

// TestCStoreRequest_MoveOriginatorRoundTrip verifies SetMoveOriginator updates
// existing elements (AddOrUpdate) and MoveOriginator reads them back.
func TestCStoreRequest_MoveOriginatorRoundTrip(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	req, err := NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest: %v", err)
	}

	if err := req.SetMoveOriginator("MOVE_SCU", 999); err != nil {
		t.Fatalf("SetMoveOriginator: %v", err)
	}
	aet, messageID, ok := req.MoveOriginator()
	if !ok || aet != "MOVE_SCU" || messageID != 999 {
		t.Fatalf("MoveOriginator = (%q, %d, %v), want (MOVE_SCU, 999, true)", aet, messageID, ok)
	}

	// A second call must replace, not silently keep, the old values.
	if err := req.SetMoveOriginator("OTHER_SCU", 7); err != nil {
		t.Fatalf("SetMoveOriginator (second): %v", err)
	}
	aet, messageID, ok = req.MoveOriginator()
	if !ok || aet != "OTHER_SCU" || messageID != 7 {
		t.Fatalf("MoveOriginator after update = (%q, %d, %v), want (OTHER_SCU, 7, true)", aet, messageID, ok)
	}
	if got, _ := req.CommandDataset().GetString(tag.MoveOriginatorApplicationEntityTitle); got != "OTHER_SCU" {
		t.Errorf("command AE title = %q, want OTHER_SCU", got)
	}
	if got := commandUInt16(t, req.CommandDataset(), tag.MoveOriginatorMessageID); got != 7 {
		t.Errorf("command message ID = %d, want 7", got)
	}
}

// TestCStoreRequest_SetAffectedUIDs verifies the setters update both the cached
// fields and the command elements.
func TestCStoreRequest_SetAffectedUIDs(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	req, err := NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest: %v", err)
	}

	const newClass = "1.2.840.10008.5.1.4.1.1.2"
	const newInstance = "1.2.3.4.5.6.7"
	req.SetAffectedSOPClassUID(newClass)
	req.SetAffectedSOPInstanceUID(newInstance)

	if req.AffectedSOPClassUID() != newClass || req.AffectedSOPInstanceUID() != newInstance {
		t.Error("setters did not update the cached UIDs")
	}
	if got, _ := req.CommandDataset().GetString(tag.AffectedSOPClassUID); got != newClass {
		t.Errorf("command Affected SOP Class UID = %q, want %q", got, newClass)
	}
	if got, _ := req.CommandDataset().GetString(tag.AffectedSOPInstanceUID); got != newInstance {
		t.Errorf("command Affected SOP Instance UID = %q, want %q", got, newInstance)
	}
}
