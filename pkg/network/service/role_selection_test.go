// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestSendMessageRejectsRequestWhenRequestorSCURoleWasNotAccepted(t *testing.T) {
	assoc := createTestAssociation()
	pc := assoc.FindPresentationContextByAbstractSyntax("1.2.840.10008.1.1")
	pc.AcceptedRole = association.NewRoleSelection(pc.AbstractSyntax, 0, 1)
	service := NewService(&recordingConn{}, assoc, WithAssociationRequestor(true))

	err := service.sendMessage(&sendRequest{
		message:  dimse.NewCEchoRequest(),
		resultCh: make(chan error, 1),
	})
	if err == nil || !strings.Contains(err.Error(), "SCU role") {
		t.Fatalf("sendMessage error = %v, want rejected SCU role error", err)
	}
}

func TestHandleRequestRejectsReverseCStoreWithoutRequestorSCPRole(t *testing.T) {
	assoc := createTestAssociation()
	pc := assoc.FindPresentationContextByAbstractSyntax(testCTImageStorageUID)
	pc.AcceptedRole = association.NewRoleSelection(pc.AbstractSyntax, 1, 0)
	called := make(chan struct{}, 1)
	service := NewService(nil, assoc,
		WithAssociationRequestor(true),
		WithCStoreHandler(func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
			called <- struct{}{}
			return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
		}))

	req := newRoleTestCStoreRequest(t)
	err := service.handleRequest(context.Background(), req)
	if err == nil || !strings.Contains(err.Error(), "SCP role") {
		t.Fatalf("handleRequest error = %v, want rejected SCP role error", err)
	}
	select {
	case <-called:
		t.Fatal("C-STORE handler was called without an accepted SCP role")
	case <-time.After(100 * time.Millisecond):
	}
}

func TestSendMessageRejectsReverseCStoreWithoutPeerSCPRole(t *testing.T) {
	assoc := createTestAssociation()
	pc := assoc.FindPresentationContextByAbstractSyntax(testCTImageStorageUID)
	pc.AcceptedRole = association.NewRoleSelection(pc.AbstractSyntax, 1, 0)
	service := NewService(&recordingConn{}, assoc)

	err := service.sendMessage(&sendRequest{
		message:  newRoleTestCStoreRequest(t),
		resultCh: make(chan error, 1),
	})
	if err == nil || !strings.Contains(err.Error(), "SCU role") {
		t.Fatalf("sendMessage error = %v, want rejected SCU role error", err)
	}
}

func TestDefaultRolesAllowOnlyRequestorToInvokeAndAcceptorToPerform(t *testing.T) {
	request := dimse.NewCEchoRequest()
	request.SetPresentationContextID(1)

	requestor := NewService(&recordingConn{}, createTestAssociation(), WithAssociationRequestor(true))
	if err := requestor.sendMessage(&sendRequest{message: request, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("requestor sendMessage failed with default roles: %v", err)
	}
	if err := requestor.handleRequest(context.Background(), request); err == nil || !strings.Contains(err.Error(), "SCP role") {
		t.Fatalf("requestor handleRequest error = %v, want default SCP role rejection", err)
	}

	acceptor := NewService(&recordingConn{}, createTestAssociation())
	if err := acceptor.sendMessage(&sendRequest{message: request, resultCh: make(chan error, 1)}); err == nil || !strings.Contains(err.Error(), "SCU role") {
		t.Fatalf("acceptor sendMessage error = %v, want default SCU role rejection", err)
	}
}

func newRoleTestCStoreRequest(t *testing.T) *dimse.CStoreRequest {
	t.Helper()
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("add SOP Class UID: %v", err)
	}
	if err := ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID})); err != nil {
		t.Fatalf("add SOP Instance UID: %v", err)
	}
	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest failed: %v", err)
	}
	req.SetPresentationContextID(3)
	return req
}
