// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"bytes"
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestAdvancedNegotiationRoundTripEnforcesAsyncWindow(t *testing.T) {
	const verificationSOPClassUID = "1.2.840.10008.1.1"
	serverConn, clientConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()
	defer func() { _ = clientConn.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	received := make(chan *dimse.CEchoRequest, 3)
	releaseHandler := make(chan struct{}, 3)
	serverReady := make(chan *association.Association, 1)
	serverErr := make(chan error, 1)

	negotiator := service.FuncAssociationNegotiator(func(
		ctx context.Context,
		assoc *association.Association,
		responder service.AssociationResponder,
	) error {
		for _, pc := range assoc.PresentationContexts {
			pc.Accept(pc.ProposedTransferSyntaxes[0])
		}
		assoc.AsynchronousOperations = association.NewAsynchronousOperationsWindow(2, 1)
		for _, role := range assoc.RoleSelections {
			role.SCURole = 1
			role.SCPRole = 1
		}
		for _, extended := range assoc.ExtendedNegotiations {
			extended.AcceptApplicationInfo([]byte{1, 0, 1})
		}
		assoc.UserIdentity.ServerResponse = []byte("server-response")
		return responder.SendAccept(ctx, assoc)
	})
	serverService := service.NewService(serverConn, nil,
		service.WithReadTimeout(time.Second),
		service.WithWriteTimeout(time.Second),
		service.WithAssociationNegotiator(negotiator),
		service.WithCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			select {
			case received <- req:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			select {
			case <-releaseHandler:
				return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}),
	)
	defer func() { _ = serverService.Close() }()
	go func() {
		if _, err := serverService.ReceiveAssociationRequest(ctx); err != nil {
			serverErr <- err
			return
		}
		if err := serverService.Start(); err != nil {
			serverErr <- err
			return
		}
		serverReady <- serverService.GetAssociation()
	}()

	client := New(
		WithAsynchronousOperations(3, 2),
		WithExtendedNegotiation(association.NewExtendedNegotiation(
			verificationSOPClassUID, []byte{1, 1, 0},
		)),
		WithExtendedNegotiation(association.NewCommonExtendedNegotiation(
			verificationSOPClassUID,
			"1.2.840.10008.4.2",
			"1.2.840.10008.5.1.4.1.1.2",
			"1.2.840.10008.5.1.4.1.1.4",
		)),
		WithUserIdentity(association.NewUserIdentityJWT([]byte("request-token"), true)),
	)
	client.AddPresentationContextWithRoles(
		verificationSOPClassUID, true, true, "1.2.840.10008.1.2.1",
	)
	client.conn = clientConn
	if err := client.negotiateAssociation(ctx); err != nil {
		t.Fatalf("negotiateAssociation failed: %v", err)
	}
	if err := client.service.Start(); err != nil {
		t.Fatalf("client service Start failed: %v", err)
	}
	client.connected = true
	defer func() { _ = client.Close() }()

	var serverAssociation *association.Association
	select {
	case serverAssociation = <-serverReady:
	case err := <-serverErr:
		t.Fatalf("server association failed: %v", err)
	case <-ctx.Done():
		t.Fatal("timed out waiting for server association")
	}
	assertAdvancedNegotiation(t, serverAssociation, true)
	assertAdvancedNegotiation(t, client.GetAssociation(), false)

	results := make(chan error, 3)
	for range 3 {
		go func() { results <- client.CEcho(ctx) }()
	}
	first := receiveIntegratedCEcho(ctx, t, received)
	second := receiveIntegratedCEcho(ctx, t, received)
	select {
	case third := <-received:
		t.Fatalf("third request %d arrived while requests %d and %d were outstanding",
			third.MessageID(), first.MessageID(), second.MessageID())
	case <-time.After(100 * time.Millisecond):
	}

	releaseHandler <- struct{}{}
	receiveIntegratedCEcho(ctx, t, received)
	releaseHandler <- struct{}{}
	releaseHandler <- struct{}{}
	for range 3 {
		select {
		case err := <-results:
			if err != nil {
				t.Fatalf("CEcho failed: %v", err)
			}
		case <-ctx.Done():
			t.Fatal("timed out waiting for C-ECHO completion")
		}
	}
}

func assertAdvancedNegotiation(t *testing.T, assoc *association.Association, serverSide bool) {
	t.Helper()
	if assoc == nil {
		t.Fatal("association is nil")
	}
	if assoc.RequestedAsynchronousOperations == nil ||
		assoc.RequestedAsynchronousOperations.MaxInvokedOperations != 3 ||
		assoc.RequestedAsynchronousOperations.MaxPerformedOperations != 2 {
		t.Fatalf("requested async window = %#v", assoc.RequestedAsynchronousOperations)
	}
	if assoc.AsynchronousOperations == nil || assoc.AsynchronousOperations.MaxInvokedOperations != 2 ||
		assoc.AsynchronousOperations.MaxPerformedOperations != 1 {
		t.Fatalf("accepted async window = %#v", assoc.AsynchronousOperations)
	}
	pc := assoc.FindPresentationContextByID(1)
	if pc == nil || pc.RequestedRole == nil || pc.AcceptedRole == nil ||
		pc.RequestedRole.SCURole != 1 || pc.RequestedRole.SCPRole != 1 ||
		pc.AcceptedRole.SCURole != 1 || pc.AcceptedRole.SCPRole != 1 {
		t.Fatalf("presentation context roles = requested %#v, accepted %#v", pc.RequestedRole, pc.AcceptedRole)
	}
	extended := assoc.FindExtendedNegotiation("1.2.840.10008.1.1")
	if extended == nil || !bytes.Equal(extended.RequestedApplicationInfo, []byte{1, 1, 0}) ||
		!bytes.Equal(extended.AcceptedApplicationInfo, []byte{1, 0, 1}) ||
		extended.ServiceClassUID != "1.2.840.10008.4.2" ||
		len(extended.RelatedGeneralSOPClassUIDs) != 2 ||
		extended.RelatedGeneralSOPClassUIDs[0] != "1.2.840.10008.5.1.4.1.1.2" ||
		extended.RelatedGeneralSOPClassUIDs[1] != "1.2.840.10008.5.1.4.1.1.4" {
		t.Fatalf("extended negotiation = %#v", extended)
	}
	if assoc.UserIdentity == nil || assoc.UserIdentity.Type != association.UserIdentityTypeJWT ||
		!bytes.Equal(assoc.UserIdentity.ServerResponse, []byte("server-response")) {
		t.Fatalf("user identity negotiation = %#v", assoc.UserIdentity)
	}
	if serverSide && !bytes.Equal(assoc.UserIdentity.PrimaryField, []byte("request-token")) {
		t.Fatal("server did not retain the requested user identity")
	}
}

func receiveIntegratedCEcho(ctx context.Context, t *testing.T, received <-chan *dimse.CEchoRequest) *dimse.CEchoRequest {
	t.Helper()
	select {
	case req := <-received:
		return req
	case <-ctx.Done():
		t.Fatal("timed out waiting for C-ECHO request")
		return nil
	}
}
