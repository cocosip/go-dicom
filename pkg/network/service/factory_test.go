// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
