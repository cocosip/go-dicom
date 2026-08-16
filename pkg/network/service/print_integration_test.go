// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/printing"
)

const (
	printFilmSessionClassUID = "1.2.840.10008.5.1.1.1"
	printFilmBoxClassUID     = "1.2.840.10008.5.1.1.2"
	printImageBoxClassUID    = "1.2.840.10008.5.1.1.4"
	printLUTClassUID         = "1.2.840.10008.5.1.1.23"
	printNCreateOperation    = "N-CREATE"
	printRemoteImageBoxUID   = "2.25.699"
)

type receivedPrintOperation struct {
	operation string
	classUID  string
	instance  string
	label     string
	format    string
	position  uint16
	actionID  uint16
}

func printTestAssociation(t *testing.T) *association.Association {
	t.Helper()
	assoc := association.NewAssociation("PRINT-SCU", "PRINT-SCP")
	for index, classUID := range []string{
		printFilmSessionClassUID,
		printLUTClassUID,
		printFilmBoxClassUID,
		printImageBoxClassUID,
	} {
		pc := association.NewPresentationContext(byte(index*2+1), classUID, transfer.ExplicitVRLittleEndian)
		pc.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
		pc.Result = association.ResultAcceptance
		if err := assoc.AddPresentationContext(pc); err != nil {
			t.Fatalf("AddPresentationContext() error = %v", err)
		}
	}
	assoc.SetEstablished(true)
	return assoc
}

func printImageBoxReferenceDataset(t *testing.T) *dataset.Dataset {
	t.Helper()
	item := dataset.New()
	if err := item.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{printImageBoxClassUID})); err != nil {
		t.Fatalf("add Referenced SOP Class UID: %v", err)
	}
	if err := item.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{printRemoteImageBoxUID})); err != nil {
		t.Fatalf("add Referenced SOP Instance UID: %v", err)
	}
	result := dataset.New()
	if err := result.Add(dataset.NewSequenceWithItems(tag.ReferencedImageBoxSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add Referenced Image Box Sequence: %v", err)
	}
	return result
}

func TestPrintClientDIMSEWorkflowEndToEnd(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	clientService := NewService(clientConn, printTestAssociation(t), WithAssociationRequestor(true))
	serverService := NewService(serverConn, printTestAssociation(t), WithAssociationRequestor(false))
	t.Cleanup(func() {
		_ = clientService.Close()
		_ = serverService.Close()
	})

	var mu sync.Mutex
	operations := make([]receivedPrintOperation, 0, 5)
	record := func(operation receivedPrintOperation) {
		mu.Lock()
		operations = append(operations, operation)
		mu.Unlock()
	}
	filmBoxResponseDataset := printImageBoxReferenceDataset(t)
	serverService.SetHandlers(&Handlers{
		NCreateHandler: func(_ context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error) {
			op := receivedPrintOperation{
				operation: printNCreateOperation,
				classUID:  req.AffectedSOPClassUID(),
				instance:  req.AffectedSOPInstanceUID(),
			}
			if req.DataDataset() != nil {
				op.label, _ = req.DataDataset().GetString(tag.FilmSessionLabel)
				op.format, _ = req.DataDataset().GetString(tag.ImageDisplayFormat)
			}
			record(op)
			var responseDataset *dataset.Dataset
			if req.AffectedSOPClassUID() == printFilmBoxClassUID {
				responseDataset = filmBoxResponseDataset
			}
			return dimse.NewNCreateResponseSuccess(req.MessageID(), req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), responseDataset), nil
		},
		NSetHandler: func(_ context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error) {
			position, _ := req.DataDataset().GetUInt16(tag.ImageBoxPosition, 0)
			record(receivedPrintOperation{
				operation: "N-SET",
				classUID:  req.RequestedSOPClassUID(),
				instance:  req.RequestedSOPInstanceUID(),
				position:  position,
			})
			return dimse.NewNSetResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil), nil
		},
		NActionHandler: func(_ context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error) {
			record(receivedPrintOperation{
				operation: "N-ACTION",
				classUID:  req.RequestedSOPClassUID(),
				instance:  req.RequestedSOPInstanceUID(),
				actionID:  req.ActionTypeID(),
			})
			return dimse.NewNActionResponseSuccess(req.MessageID(), req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID(), nil), nil
		},
	})

	if err := serverService.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("server setState() error = %v", err)
	}
	if err := clientService.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("client setState() error = %v", err)
	}
	if err := serverService.Start(); err != nil {
		t.Fatalf("server Start() error = %v", err)
	}
	if err := clientService.Start(); err != nil {
		t.Fatalf("client Start() error = %v", err)
	}

	session := printing.NewFilmSession(printFilmSessionClassUID, "2.25.601", false)
	session.FilmSessionLabel = "wire workflow"
	lut := printing.NewPresentationLUT("2.25.602")
	if err := lut.SetLUT(2, 0, 12, []uint16{10, 20}); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}
	session.AddPresentationLUT(lut)
	filmBox := printing.NewFilmBox("2.25.603", `STANDARD\1,1`)
	imageBox := printing.NewImageBox("2.25.604", false)
	imageBox.ImageBoxPosition = 1
	imageBox.SetImageData([]byte{1, 2, 3, 4})
	filmBox.AddImageBox(imageBox)
	session.AddFilmBox(filmBox)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := printing.NewClient(clientService).Print(ctx, session); err != nil {
		t.Fatalf("Print() error = %v", err)
	}

	mu.Lock()
	got := append([]receivedPrintOperation(nil), operations...)
	mu.Unlock()
	if len(got) != 5 {
		t.Fatalf("received %d operations, want 5: %#v", len(got), got)
	}
	if got[0].operation != printNCreateOperation || got[0].classUID != printFilmSessionClassUID || got[0].instance != "2.25.601" || got[0].label != "wire workflow" {
		t.Errorf("Film Session operation = %#v", got[0])
	}
	if got[1].operation != printNCreateOperation || got[1].classUID != printLUTClassUID || got[1].instance != "2.25.602" {
		t.Errorf("Presentation LUT operation = %#v", got[1])
	}
	if got[2].operation != printNCreateOperation || got[2].classUID != printFilmBoxClassUID || got[2].instance != "2.25.603" || got[2].format != `STANDARD\1,1` {
		t.Errorf("Film Box operation = %#v", got[2])
	}
	if got[3].operation != "N-SET" || got[3].classUID != printImageBoxClassUID || got[3].instance != printRemoteImageBoxUID || got[3].position != 1 {
		t.Errorf("Image Box operation = %#v", got[3])
	}
	if got[4].operation != "N-ACTION" || got[4].classUID != printFilmSessionClassUID || got[4].instance != "2.25.601" || got[4].actionID != 1 {
		t.Errorf("Film Session print operation = %#v", got[4])
	}
}
