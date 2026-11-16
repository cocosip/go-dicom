// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package server

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
)

// Shared literals for association-request tests in this package.
const (
	testApplicationContext     = "1.2.840.10008.3.1.1.1"
	testExplicitVRLittleEndian = "1.2.840.10008.1.2.1"
	testImplementationClassUID = "1.2.826.0.1.3680043.10.854"
)

// rejectingNegotiator refuses every association with an A-ASSOCIATE-RJ.
type rejectingNegotiator struct{}

func (rejectingNegotiator) OnAssociationRequest(
	ctx context.Context,
	_ *association.Association,
	responder service.AssociationResponder,
) error {
	return responder.SendReject(
		ctx,
		pdu.ResultRejectedPermanent,
		pdu.SourceServiceUser,
		pdu.ReasonServiceUserCalledAETitleNotRecognized,
	)
}

// TestRejectedAssociationEmitsConnectionClosed verifies that the server closes
// the per-connection service when the application negotiator rejects the
// association, so the connection_opened event emitted by NewService is always
// paired with connection_closed. Without the close the pair is unbalanced and a
// ConnectionLifecycleHandler never sees the rejection.
func TestRejectedAssociationEmitsConnectionClosed(t *testing.T) {
	recorder := &serverObservationRecorder{}
	server := New(
		WithAssociationTimeout(time.Second),
		WithEventObserver(observability.EventObserverFunc(recorder.event)),
	)
	server.SetAssociationNegotiator(rejectingNegotiator{})
	server.ctx, server.cancel = context.WithCancel(context.Background())
	defer server.cancel()

	clientConn, serverConn := net.Pipe()
	done := make(chan struct{})
	server.wg.Add(1)
	go func() {
		server.handleConnection(serverConn)
		close(done)
	}()

	clientService := service.NewService(clientConn, nil, service.WithAssociationRequestor(true))
	rq := &pdu.AAssociateRQ{
		ProtocolVersion:    1,
		CalledAETitle:      "REJECT-SCP",
		CallingAETitle:     "REJECT-SCU",
		ApplicationContext: testApplicationContext,
		PresentationContexts: []pdu.PresentationContextRQ{{
			ID:               1,
			AbstractSyntax:   testVerificationSOPClassUID,
			TransferSyntaxes: []string{testExplicitVRLittleEndian},
		}},
		UserInformation: &pdu.UserInformation{
			MaximumLength:          16384,
			ImplementationClassUID: testImplementationClassUID,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := clientService.SendAssociationRequest(ctx, rq); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}
	if _, err := clientService.ReceiveAssociationResponse(ctx); err == nil {
		t.Fatal("ReceiveAssociationResponse() error = nil, want rejection")
	}
	_ = clientService.Close()
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("server connection did not stop")
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	var opened, closed bool
	for _, event := range recorder.events {
		switch event.Kind {
		case observability.EventConnectionOpened:
			opened = true
		case observability.EventConnectionClosed:
			closed = true
		}
	}
	if !opened || !closed {
		t.Fatalf("events = %#v, want a balanced connection opened/closed pair", recorder.events)
	}
}
