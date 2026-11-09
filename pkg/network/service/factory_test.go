// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestCreateCGetRequestRequiresQueryRetrieveLevel(t *testing.T) {
	command := dataset.New()
	if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{1})); err != nil {
		t.Fatalf("failed to add MessageID: %v", err)
	}

	identifier := dataset.New()

	_, err := createCGetRequest(command, identifier)
	if err == nil {
		t.Fatal("expected error when QueryRetrieveLevel is missing")
	}
	if !strings.Contains(err.Error(), "QueryRetrieveLevel") {
		t.Fatalf("expected QueryRetrieveLevel error, got %v", err)
	}
}

func TestCreateMessageFromDatasetsRejectsZeroRequestMessageID(t *testing.T) {
	command := dimse.CreateCommandDataset(uint16(dimse.CommandCEchoRQ), 0)
	if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{0})); err != nil {
		t.Fatalf("failed to add zero MessageID: %v", err)
	}

	_, err := createMessageFromDatasets(command, nil)
	if err == nil {
		t.Fatal("expected zero request MessageID to be rejected")
	}
	if !strings.Contains(err.Error(), "MessageID") {
		t.Fatalf("expected MessageID error, got %v", err)
	}
}

func TestCreateMessageFromDatasetsRejectsInvalidPriority(t *testing.T) {
	command := dimse.CreateCommandDataset(uint16(dimse.CommandCFindRQ), 1)
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.2.1.1"})); err != nil {
		t.Fatalf("failed to add AffectedSOPClassUID: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.Priority, []uint16{0xFFFF})); err != nil {
		t.Fatalf("failed to add Priority: %v", err)
	}
	identifier := dataset.New()
	if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(dimse.QueryRetrieveLevelStudy)})); err != nil {
		t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
	}

	_, err := createMessageFromDatasets(command, identifier)
	if err == nil {
		t.Fatal("expected invalid Priority to be rejected")
	}
	if !strings.Contains(err.Error(), "Priority") {
		t.Fatalf("expected Priority error, got %v", err)
	}
}

func TestCreateCMoveRequestRequiresMoveDestinationAndQueryRetrieveLevel(t *testing.T) {
	t.Run("missing MoveDestination", func(t *testing.T) {
		command := dataset.New()
		if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{1})); err != nil {
			t.Fatalf("failed to add MessageID: %v", err)
		}

		identifier := dataset.New()
		if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{"STUDY"})); err != nil {
			t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
		}

		_, err := createCMoveRequest(command, identifier)
		if err == nil {
			t.Fatal("expected error when MoveDestination is missing")
		}
		if !strings.Contains(err.Error(), "MoveDestination") {
			t.Fatalf("expected MoveDestination error, got %v", err)
		}
	})

	t.Run("missing QueryRetrieveLevel", func(t *testing.T) {
		command := dataset.New()
		if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{1})); err != nil {
			t.Fatalf("failed to add MessageID: %v", err)
		}
		if err := command.Add(element.NewString(tag.MoveDestination, vr.AE, []string{testMoveDestinationAE})); err != nil {
			t.Fatalf("failed to add MoveDestination: %v", err)
		}

		identifier := dataset.New()

		_, err := createCMoveRequest(command, identifier)
		if err == nil {
			t.Fatal("expected error when QueryRetrieveLevel is missing")
		}
		if !strings.Contains(err.Error(), "QueryRetrieveLevel") {
			t.Fatalf("expected QueryRetrieveLevel error, got %v", err)
		}
	})
}

func TestCreateNGetRequestRequiresRequestedSOPUIDs(t *testing.T) {
	command := dataset.New()
	if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{1})); err != nil {
		t.Fatalf("failed to add MessageID: %v", err)
	}

	_, err := createNGetRequest(command)
	if err == nil {
		t.Fatal("expected error when requested SOP UIDs are missing")
	}
	if !strings.Contains(err.Error(), "RequestedSOPClassUID") {
		t.Fatalf("expected RequestedSOPClassUID error, got %v", err)
	}
}

func TestCreateNActionResponseAllowsZeroActionTypeID(t *testing.T) {
	resp := dimse.NewNActionResponse(100, status.NActionSuccess, "1.2.840.10008.5.1.1.40", "1.2.3", 0, nil)

	decoded, err := createNActionResponse(resp.CommandDataset(), resp.DataDataset())
	if err != nil {
		t.Fatalf("expected zero ActionTypeID to round-trip, got error: %v", err)
	}
	if decoded.ActionTypeID() != 0 {
		t.Fatalf("expected ActionTypeID 0, got %d", decoded.ActionTypeID())
	}
}

func TestCreateNEventReportResponseAllowsZeroEventTypeID(t *testing.T) {
	resp := dimse.NewNEventReportResponse(100, status.NEventReportSuccess, "1.2.840.10008.5.1.1.40", "1.2.3", 0, nil)

	decoded, err := createNEventReportResponse(resp.CommandDataset(), resp.DataDataset())
	if err != nil {
		t.Fatalf("expected zero EventTypeID to round-trip, got error: %v", err)
	}
	if decoded.EventTypeID() != 0 {
		t.Fatalf("expected EventTypeID 0, got %d", decoded.EventTypeID())
	}
}

func TestCreateCFindRequestPreservesAffectedSOPClassUID(t *testing.T) {
	const patientRootFind = "1.2.840.10008.5.1.4.1.2.1.1"

	command := dimse.CreateCommandDataset(uint16(dimse.CommandCFindRQ), 1)
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{patientRootFind})); err != nil {
		t.Fatalf("failed to add AffectedSOPClassUID: %v", err)
	}

	identifier := dataset.New()
	if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(dimse.QueryRetrieveLevelStudy)})); err != nil {
		t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
	}

	req, err := createCFindRequest(command, identifier)
	if err != nil {
		t.Fatalf("createCFindRequest() error = %v", err)
	}
	if req.AffectedSOPClassUID() != patientRootFind {
		t.Fatalf("AffectedSOPClassUID = %q, want %q", req.AffectedSOPClassUID(), patientRootFind)
	}
}

func TestCreateCGetRequestPreservesAffectedSOPClassUID(t *testing.T) {
	const patientRootGet = "1.2.840.10008.5.1.4.1.2.1.3"

	command := dimse.CreateCommandDataset(uint16(dimse.CommandCGetRQ), 1)
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{patientRootGet})); err != nil {
		t.Fatalf("failed to add AffectedSOPClassUID: %v", err)
	}

	identifier := dataset.New()
	if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(dimse.QueryRetrieveLevelStudy)})); err != nil {
		t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
	}

	req, err := createCGetRequest(command, identifier)
	if err != nil {
		t.Fatalf("createCGetRequest() error = %v", err)
	}
	if req.AffectedSOPClassUID() != patientRootGet {
		t.Fatalf("AffectedSOPClassUID = %q, want %q", req.AffectedSOPClassUID(), patientRootGet)
	}
}

func TestCreateCMoveRequestPreservesAffectedSOPClassUID(t *testing.T) {
	const patientRootMove = "1.2.840.10008.5.1.4.1.2.1.2"

	command := dimse.CreateCommandDataset(uint16(dimse.CommandCMoveRQ), 1)
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{patientRootMove})); err != nil {
		t.Fatalf("failed to add AffectedSOPClassUID: %v", err)
	}
	if err := command.Add(element.NewString(tag.MoveDestination, vr.AE, []string{testMoveDestinationAE})); err != nil {
		t.Fatalf("failed to add MoveDestination: %v", err)
	}

	identifier := dataset.New()
	if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(dimse.QueryRetrieveLevelStudy)})); err != nil {
		t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
	}

	req, err := createCMoveRequest(command, identifier)
	if err != nil {
		t.Fatalf("createCMoveRequest() error = %v", err)
	}
	if req.AffectedSOPClassUID() != patientRootMove {
		t.Fatalf("AffectedSOPClassUID = %q, want %q", req.AffectedSOPClassUID(), patientRootMove)
	}
}

func TestCreateCStoreRequestRejectsEmptyCommandUIDs(t *testing.T) {
	data := dataset.New()
	if err := data.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("failed to add SOPClassUID: %v", err)
	}
	if err := data.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4"})); err != nil {
		t.Fatalf("failed to add SOPInstanceUID: %v", err)
	}

	command := dataset.New()
	if err := command.Add(element.NewUnsignedShort(tag.CommandField, []uint16{uint16(dimse.CommandCStoreRQ)})); err != nil {
		t.Fatalf("failed to add CommandField: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.MessageID, []uint16{1})); err != nil {
		t.Fatalf("failed to add MessageID: %v", err)
	}
	// Present-but-empty Type 1 identifier: the dataset fallback must not hide it.
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{""})); err != nil {
		t.Fatalf("failed to add AffectedSOPClassUID: %v", err)
	}

	if _, err := createCStoreRequest(command, data); err == nil {
		t.Fatal("createCStoreRequest() accepted a present-but-empty AffectedSOPClassUID")
	}
}

func TestCreateCStoreResponseToleratesMissingOptionalUIDs(t *testing.T) {
	command := dataset.New()
	if err := command.Add(element.NewUnsignedShort(tag.CommandField, []uint16{uint16(dimse.CommandCStoreRSP)})); err != nil {
		t.Fatalf("failed to add CommandField: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{7})); err != nil {
		t.Fatalf("failed to add MessageIDBeingRespondedTo: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.Status, []uint16{status.Success.Code})); err != nil {
		t.Fatalf("failed to add Status: %v", err)
	}

	resp, err := createCStoreResponse(command)
	if err != nil {
		t.Fatalf("createCStoreResponse() error = %v, want nil for a minimal C-STORE-RSP", err)
	}
	if resp.MessageIDBeingRespondedTo() != 7 {
		t.Fatalf("MessageIDBeingRespondedTo = %d, want 7", resp.MessageIDBeingRespondedTo())
	}
	if resp.AffectedSOPClassUID() != "" || resp.AffectedSOPInstanceUID() != "" {
		t.Fatalf("optional UIDs = %q/%q, want empty", resp.AffectedSOPClassUID(), resp.AffectedSOPInstanceUID())
	}
}

func TestCreateCFindResponseToleratesMissingAffectedSOPClassUID(t *testing.T) {
	command := dataset.New()
	if err := command.Add(element.NewUnsignedShort(tag.CommandField, []uint16{uint16(dimse.CommandCFindRSP)})); err != nil {
		t.Fatalf("failed to add CommandField: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.MessageIDBeingRespondedTo, []uint16{9})); err != nil {
		t.Fatalf("failed to add MessageIDBeingRespondedTo: %v", err)
	}
	if err := command.Add(element.NewUnsignedShort(tag.Status, []uint16{status.Success.Code})); err != nil {
		t.Fatalf("failed to add Status: %v", err)
	}

	resp, err := createCFindResponse(command, nil)
	if err != nil {
		t.Fatalf("createCFindResponse() error = %v, want nil for a minimal C-FIND-RSP", err)
	}
	if resp.AffectedSOPClassUID() != "" {
		t.Fatalf("AffectedSOPClassUID = %q, want empty", resp.AffectedSOPClassUID())
	}
}

func TestCreateMessageFromDatasetsRejectsEmptyRequiredString(t *testing.T) {
	command := dimse.CreateCommandDataset(uint16(dimse.CommandCFindRQ), 3)
	if err := command.Add(element.NewString(tag.AffectedSOPClassUID, vr.UI, []string{""})); err != nil {
		t.Fatalf("failed to add empty AffectedSOPClassUID: %v", err)
	}
	identifier := dataset.New()
	if err := identifier.Add(element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(dimse.QueryRetrieveLevelStudy)})); err != nil {
		t.Fatalf("failed to add QueryRetrieveLevel: %v", err)
	}

	_, err := createMessageFromDatasets(command, identifier)
	if err == nil {
		t.Fatal("expected a present-but-empty AffectedSOPClassUID to be rejected")
	}
	if !strings.Contains(err.Error(), "AffectedSOPClassUID is empty") {
		t.Fatalf("error = %v, want empty AffectedSOPClassUID error", err)
	}
}
