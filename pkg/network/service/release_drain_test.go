// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/status"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// TestReleaseRequestWaitsForRunningHandler verifies that an A-RELEASE-RQ that
// arrives while a request handler is still running lets the handler write its
// final response first: PS3.8 9.3.5 makes the A-RELEASE-RP the last PDU, so the
// final C-FIND-RSP must reach the peer before it.
func TestReleaseRequestWaitsForRunningHandler(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	t.Cleanup(func() {
		_ = clientConn.Close()
		_ = serverConn.Close()
	})

	const cFindSOPClassUID = "1.2.840.10008.5.1.4.1.2.2.1"
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	pc := association.NewPresentationContext(1, cFindSOPClassUID, transfer.ExplicitVRLittleEndian)
	pc.Accept(transfer.ExplicitVRLittleEndian)
	if err := assoc.AddPresentationContext(pc); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}
	assoc.SetEstablished(true)

	handlerRunning := make(chan struct{})
	releaseHandler := make(chan struct{})
	var releaseOnce sync.Once
	release := func() { releaseOnce.Do(func() { close(releaseHandler) }) }
	// Unblock the handler on every exit path so Close cannot wait for it.
	defer release()

	service := acceptedService(t, serverConn, assoc,
		WithAssociationRequestor(false),
		WithHandlerShutdownTimeout(2*time.Second),
		WithHandlers(&Handlers{
			CFindHandler: func(_ context.Context, op CFindOperation) error {
				close(handlerRunning)
				<-releaseHandler
				return op.SendFinal(status.Success)
			},
		}))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	request.SetPresentationContextID(1)
	if err := writeDIMSEPDU(clientConn, request, 1, transfer.ExplicitVRLittleEndian); err != nil {
		t.Fatalf("write C-FIND-RQ: %v", err)
	}

	// Wait until the handler is running before releasing, so the release is
	// processed while an operation is still outstanding.
	select {
	case <-handlerRunning:
	case <-time.After(2 * time.Second):
		t.Fatal("C-FIND handler did not start")
	}

	releaseRQ, err := pdu.NewAReleaseRQ().Encode()
	if err != nil {
		t.Fatalf("encode A-RELEASE-RQ: %v", err)
	}
	if err := transport.WritePDU(clientConn, 2*time.Second, releaseRQ); err != nil {
		t.Fatalf("write A-RELEASE-RQ: %v", err)
	}

	// With the handler blocked and the release processed, nothing may reach the
	// wire until the handler is allowed to finish. Without the wait in
	// handleReleaseRequest, the A-RELEASE-RP (and no response) arrives here.
	if raw, err := transport.ReadPDU(clientConn, 200*time.Millisecond); err == nil {
		t.Fatalf("received PDU type 0x%02X before the running handler finished", raw.Type)
	}

	release()

	command, err := readNextDIMSECommand(clientConn, 2*time.Second)
	if err != nil {
		t.Fatalf("read final C-FIND-RSP: %v", err)
	}
	commandField, err := command.GetUInt16(tag.CommandField, 0)
	if err != nil {
		t.Fatalf("read Command Field: %v", err)
	}
	if dimse.CommandField(commandField) != dimse.CommandCFindRSP {
		t.Fatalf("first PDU after the release was %s, want C-FIND-RSP", dimse.CommandField(commandField))
	}
	code, err := command.GetUInt16(tag.Status, 0)
	if err != nil {
		t.Fatalf("read Status: %v", err)
	}
	if code != status.Success.Code {
		t.Fatalf("final C-FIND-RSP status = 0x%04X, want Success", code)
	}

	rawRP, err := transport.ReadPDU(clientConn, 2*time.Second)
	if err != nil {
		t.Fatalf("read A-RELEASE-RP: %v", err)
	}
	if rawRP.Type != pdu.TypeAReleaseRP {
		t.Fatalf("PDU after the final response type = 0x%02X, want A-RELEASE-RP", rawRP.Type)
	}
}
