// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

type lifecycleTestListener struct {
	acceptStarted chan struct{}
	releaseAccept chan struct{}
	conn          net.Conn
	closeOnce     sync.Once
	closed        chan struct{}
}

func (l *lifecycleTestListener) Accept(ctx context.Context) (net.Conn, error) {
	if l.acceptStarted != nil {
		l.closeOnce.Do(func() { close(l.acceptStarted) })
	}
	if l.releaseAccept != nil {
		<-l.releaseAccept
		return l.conn, nil
	}
	<-ctx.Done()
	return nil, ctx.Err()
}

func (l *lifecycleTestListener) Close() error {
	if l.closed != nil {
		select {
		case <-l.closed:
		default:
			close(l.closed)
		}
	}
	return nil
}

func TestServerShutdownWaitsForStartupPublication(t *testing.T) {
	server := New(WithPort(0))
	listenStarted := make(chan struct{})
	releaseListen := make(chan struct{})
	listener := &lifecycleTestListener{}
	server.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		close(listenStarted)
		<-releaseListen
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- server.ListenAndServe(context.Background()) }()
	<-listenStarted
	shutdownDone := make(chan error, 1)
	go func() { shutdownDone <- server.Shutdown(context.Background()) }()

	returnedEarly := false
	select {
	case <-shutdownDone:
		returnedEarly = true
	case <-time.After(25 * time.Millisecond):
	}
	close(releaseListen)
	if !returnedEarly {
		if err := <-shutdownDone; err != nil {
			t.Fatalf("Shutdown() error = %v", err)
		}
	}
	if err := <-serveDone; !errors.Is(err, context.Canceled) {
		t.Fatalf("ListenAndServe() error = %v, want context.Canceled", err)
	}
	if returnedEarly {
		t.Fatal("Shutdown() returned before startup published its listener and lifecycle context")
	}
}

func TestServerShutdownWaitsForAcceptLoopBeforeConnectionWait(t *testing.T) {
	serverConn, peerConn := net.Pipe()
	_ = peerConn.Close()
	listener := &lifecycleTestListener{
		acceptStarted: make(chan struct{}),
		releaseAccept: make(chan struct{}),
		conn:          serverConn,
	}
	server := New(WithPort(0))
	server.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- server.ListenAndServe(context.Background()) }()
	<-listener.acceptStarted
	shutdownDone := make(chan error, 1)
	go func() { shutdownDone <- server.Shutdown(context.Background()) }()

	returnedBeforeAcceptExited := false
	select {
	case <-shutdownDone:
		returnedBeforeAcceptExited = true
	case <-time.After(25 * time.Millisecond):
	}
	close(listener.releaseAccept)
	if !returnedBeforeAcceptExited {
		if err := <-shutdownDone; err != nil {
			t.Fatalf("Shutdown() error = %v", err)
		}
	}
	if err := <-serveDone; !errors.Is(err, context.Canceled) {
		t.Fatalf("ListenAndServe() error = %v, want context.Canceled", err)
	}
	if returnedBeforeAcceptExited {
		t.Fatal("Shutdown() called connection Wait before the accept loop could finish adding work")
	}
}

func TestServerAcceptTimeoutRearmsListener(t *testing.T) {
	timedOut := make(chan struct{}, 1)
	server := New(
		WithPort(0),
		WithAcceptTimeout(25*time.Millisecond),
		WithLogger(observability.LoggerFunc(func(_ context.Context, record observability.LogRecord) {
			if record.Message == "accept_failed" {
				select {
				case timedOut <- struct{}{}:
				default:
				}
			}
		})),
	)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- server.ListenAndServe(ctx) }()

	select {
	case <-timedOut:
	case <-time.After(time.Second):
		cancel()
		<-done
		t.Fatal("AcceptTimeout did not expire an idle accept")
	}

	if !server.IsRunning() {
		cancel()
		<-done
		t.Fatal("server stopped after an accept timeout")
	}
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("ListenAndServe() error = %v, want context.Canceled", err)
	}
}

func TestServerAssociationTimeoutDoesNotCloseEstablishedIdleAssociation(t *testing.T) {
	const associationTimeout = 30 * time.Millisecond
	server, clientService, _, done := establishPipeAssociation(t,
		WithAssociationTimeout(associationTimeout),
		WithTransportReadTimeout(0),
	)
	defer stopPipeAssociation(t, server, clientService, done)

	select {
	case <-done:
		t.Fatal("established association closed at the association negotiation timeout")
	case <-time.After(4 * associationTimeout):
	}
}

func TestServerTransportReadTimeoutClosesEstablishedIdleAssociation(t *testing.T) {
	const transportReadTimeout = 30 * time.Millisecond
	server, clientService, _, done := establishPipeAssociation(t,
		WithAssociationTimeout(time.Second),
		WithTransportReadTimeout(transportReadTimeout),
	)
	defer stopPipeAssociation(t, server, clientService, done)

	select {
	case <-done:
	case <-time.After(10 * transportReadTimeout):
		t.Fatal("established association remained open beyond TransportReadTimeout")
	}
}

func TestServerAssociationTimeoutBoundsNegotiationWriteWhenTransportWriteDisabled(t *testing.T) {
	const associationTimeout = 30 * time.Millisecond
	server := New(
		WithAssociationTimeout(associationTimeout),
		WithTransportWriteTimeout(0),
	)
	server.ctx, server.cancel = context.WithCancel(context.Background())

	clientConn, serverConn := net.Pipe()
	done := make(chan struct{})
	server.wg.Add(1)
	go func() {
		server.handleConnection(serverConn)
		close(done)
	}()
	clientService := service.NewService(clientConn, nil, service.WithAssociationRequestor(true))
	defer func() {
		_ = clientService.Close()
		server.cancel()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("server connection did not stop during cleanup")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := clientService.SendAssociationRequest(ctx, verificationAssociationRequest()); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}

	select {
	case <-done:
	case <-time.After(10 * associationTimeout):
		t.Fatal("blocked association response write outlived AssociationTimeout")
	}
}

func TestServerRequestTimeoutAppliesToOutgoingDIMSERequest(t *testing.T) {
	const associationTimeout = 30 * time.Millisecond
	server, clientService, clientConn, done := establishPipeAssociation(t,
		WithAssociationTimeout(associationTimeout),
		WithRequestTimeout(30*time.Millisecond),
		WithTransportReadTimeout(0),
		WithTransportWriteTimeout(0),
	)
	defer stopPipeAssociation(t, server, clientService, done)
	svc := waitForActiveService(t, server)
	time.Sleep(2 * associationTimeout)

	requestRead := make(chan error, 1)
	go func() {
		_, err := transport.ReadPDU(clientConn, time.Second)
		requestRead <- err
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := svc.SendCEcho(ctx, dimse.NewCEchoRequest())
	if !errors.Is(err, service.ErrRequestTimeout) {
		t.Fatalf("SendCEcho() error = %v, want ErrRequestTimeout", err)
	}
	if err := <-requestRead; err != nil {
		t.Fatalf("peer ReadPDU() error = %v", err)
	}
	if svc.IsClosed() {
		t.Fatal("request timeout closed the association")
	}
}

func TestServerCloseClosesAssociationAfterShutdownTimeout(t *testing.T) {
	server, clientService, _, done := establishPipeAssociation(t,
		WithAssociationTimeout(time.Second),
		WithTransportReadTimeout(0),
	)
	defer func() { _ = clientService.Close() }()

	server.runningMu.Lock()
	server.running = true
	server.runningMu.Unlock()
	waitForActiveService(t, server)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Shutdown() error = %v, want context deadline exceeded", err)
	}
	if server.ActiveConnections() != 1 {
		t.Fatalf("ActiveConnections() = %d, want 1 after graceful shutdown timeout", server.ActiveConnections())
	}

	if err := server.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Close() did not close the active association")
	}
	if server.ActiveConnections() != 0 {
		t.Fatalf("ActiveConnections() = %d, want 0 after Close()", server.ActiveConnections())
	}
}

func TestServerCloseIsIdempotentWhenNotRunning(t *testing.T) {
	server := New()

	if err := server.Close(); err != nil {
		t.Fatalf("first Close() error = %v", err)
	}
	if err := server.Close(); err != nil {
		t.Fatalf("second Close() error = %v", err)
	}
}

func establishPipeAssociation(t *testing.T, opts ...Option) (*Server, *service.Service, net.Conn, <-chan struct{}) {
	t.Helper()
	server := New(opts...)
	server.ctx, server.cancel = context.WithCancel(context.Background())

	clientConn, serverConn := net.Pipe()
	done := make(chan struct{})
	server.wg.Add(1)
	go func() {
		server.handleConnection(serverConn)
		close(done)
	}()

	clientService := service.NewService(clientConn, nil, service.WithAssociationRequestor(true))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := clientService.SendAssociationRequest(ctx, verificationAssociationRequest()); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}
	if _, err := clientService.ReceiveAssociationResponse(ctx); err != nil {
		t.Fatalf("ReceiveAssociationResponse() error = %v", err)
	}
	return server, clientService, clientConn, done
}

func stopPipeAssociation(t *testing.T, server *Server, clientService *service.Service, done <-chan struct{}) {
	t.Helper()
	_ = clientService.Close()
	server.cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("server connection did not stop")
	}
}

func verificationAssociationRequest() *pdu.AAssociateRQ {
	return &pdu.AAssociateRQ{
		ProtocolVersion:    1,
		CalledAETitle:      "TIMEOUT-SCP",
		CallingAETitle:     "TIMEOUT-SCU",
		ApplicationContext: "1.2.840.10008.3.1.1.1",
		PresentationContexts: []pdu.PresentationContextRQ{{
			ID:               1,
			AbstractSyntax:   testVerificationSOPClassUID,
			TransferSyntaxes: []string{"1.2.840.10008.1.2.1"},
		}},
		UserInformation: &pdu.UserInformation{
			MaximumLength: 16384,
			SCPSCURoleSelections: []pdu.SCPSCURoleSelection{{
				SOPClassUID: testVerificationSOPClassUID,
				SCURole:     1,
				SCPRole:     1,
			}},
		},
	}
}

func waitForActiveService(t *testing.T, server *Server) *service.Service {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		server.connectionsMu.RLock()
		for _, connection := range server.connections {
			svc := connection.service
			server.connectionsMu.RUnlock()
			return svc
		}
		server.connectionsMu.RUnlock()
		time.Sleep(time.Millisecond)
	}
	t.Fatal("server did not register the negotiated association")
	return nil
}
