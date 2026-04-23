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
		if err := command.Add(element.NewString(tag.MoveDestination, vr.AE, []string{"DEST_AE"})); err != nil {
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
