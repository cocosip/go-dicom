// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// countingListener records how many times Accept was called.
type countingListener struct {
	mu      sync.Mutex
	accepts int
	conns   chan net.Conn
	closed  chan struct{}
	once    sync.Once
}

func newCountingListener() *countingListener {
	return &countingListener{conns: make(chan net.Conn, 4), closed: make(chan struct{})}
}

func (l *countingListener) Accept(ctx context.Context) (net.Conn, error) {
	l.mu.Lock()
	l.accepts++
	l.mu.Unlock()
	select {
	case conn := <-l.conns:
		return conn, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-l.closed:
		return nil, net.ErrClosed
	}
}

func (l *countingListener) Close() error {
	l.once.Do(func() { close(l.closed) })
	return nil
}

func (l *countingListener) acceptCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.accepts
}

// addressedPipeConn gives a net.Pipe endpoint a distinct remote address so the
// server's connID (remote address) does not collide, unlike real TCP peers.
type addressedPipeConn struct {
	net.Conn
	remote net.Addr
}

func (c addressedPipeConn) RemoteAddr() net.Addr { return c.remote }

func pipeWithRemoteAddress(conn net.Conn, port int) net.Conn {
	return addressedPipeConn{
		Conn:   conn,
		remote: &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: port},
	}
}

func waitForServerCondition(t *testing.T, condition func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("server did not reach the expected state")
		}
		time.Sleep(time.Millisecond)
	}
}

// TestAcceptLoopRejectsOverLimitConnectionWithRJ verifies that a connection
// arriving while MaxConnections is exhausted is still accepted: the server
// reads its A-ASSOCIATE-RQ and answers with A-ASSOCIATE-RJ (rejected-transient,
// service-provider-presentation, local-limit-exceeded) instead of leaving the
// peer hanging, and the rejected connection never becomes active.
func TestAcceptLoopRejectsOverLimitConnectionWithRJ(t *testing.T) {
	srv := New(WithMaxConnections(1), WithAssociationTimeout(10*time.Second))
	listener := newCountingListener()
	srv.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- srv.ListenAndServe(context.Background()) }()
	t.Cleanup(func() {
		_ = srv.Close()
		<-serveDone
	})

	// The first connection never sends A-ASSOCIATE-RQ, so it keeps the only
	// slot in the negotiating state for the whole test.
	firstServerConn, firstClientConn := net.Pipe()
	t.Cleanup(func() { _ = firstClientConn.Close() })
	listener.conns <- firstServerConn
	waitForServerCondition(t, func() bool {
		return listener.acceptCount() >= 1 && srv.connectionLimitReached()
	})

	// The second connection must be accepted and rejected with an RJ.
	secondServerConn, secondClientConn := net.Pipe()
	t.Cleanup(func() { _ = secondClientConn.Close() })
	listener.conns <- secondServerConn
	waitForServerCondition(t, func() bool { return listener.acceptCount() >= 2 })

	rawRQ, err := verificationAssociationRequest().Encode()
	if err != nil {
		t.Fatalf("encode A-ASSOCIATE-RQ: %v", err)
	}
	_ = secondClientConn.SetDeadline(time.Now().Add(2 * time.Second))
	if err := pdu.WritePDU(secondClientConn, rawRQ); err != nil {
		t.Fatalf("write A-ASSOCIATE-RQ: %v", err)
	}
	rawRJ, err := transport.ReadPDU(secondClientConn, 2*time.Second)
	if err != nil {
		t.Fatalf("read A-ASSOCIATE-RJ: %v", err)
	}
	rj := &pdu.AAssociateRJ{}
	if err := rj.Decode(rawRJ); err != nil {
		t.Fatalf("decode A-ASSOCIATE-RJ: %v", err)
	}
	if rj.Result != pdu.ResultRejectedTransient ||
		rj.Source != pdu.SourceServiceProviderPresentation ||
		rj.Reason != pdu.ReasonServiceProviderPresentationLocalLimitExceeded {
		t.Fatalf("RJ = result %d, source %d, reason %d; want 2/3/2 (local-limit-exceeded)",
			rj.Result, rj.Source, rj.Reason)
	}
	if got := srv.ActiveConnections(); got != 0 {
		t.Fatalf("ActiveConnections() = %d, want 0 for a rejected connection", got)
	}

	srv.connectionsMu.RLock()
	inflight := srv.inflightConnections
	srv.connectionsMu.RUnlock()
	if inflight != 1 {
		t.Fatalf("inflightConnections = %d, want only the negotiating connection", inflight)
	}
}

// TestConnectionLimitAllowsAssociationWithinLimit verifies that the
// connection-limit rejection path does not disturb connections accepted while a
// slot is free.
func TestConnectionLimitAllowsAssociationWithinLimit(t *testing.T) {
	server, clientService, _, done := establishPipeAssociation(t, WithMaxConnections(1))
	defer stopPipeAssociation(t, server, clientService, done)

	waitForActiveService(t, server)
	if got := server.ActiveConnections(); got != 1 {
		t.Fatalf("ActiveConnections() = %d, want 1", got)
	}
}

// TestOverLimitSilentConnectionClosedByNegotiationTimeout verifies that a
// connection over MaxConnections that never sends an A-ASSOCIATE-RQ is closed
// within the bounded reject-negotiation wait instead of pinning a goroutine
// forever.
func TestOverLimitSilentConnectionClosedByNegotiationTimeout(t *testing.T) {
	previous := rejectNegotiationTimeout
	rejectNegotiationTimeout = 50 * time.Millisecond
	t.Cleanup(func() { rejectNegotiationTimeout = previous })

	srv := New(WithMaxConnections(1), WithAssociationTimeout(5*time.Second))
	listener := newCountingListener()
	srv.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- srv.ListenAndServe(context.Background()) }()
	t.Cleanup(func() {
		_ = srv.Close()
		<-serveDone
	})

	// Occupy the only slot with a fully negotiated association.
	firstServerConn, firstClientConn := net.Pipe()
	listener.conns <- firstServerConn
	clientService := service.NewService(firstClientConn, nil, service.WithAssociationRequestor(true))
	t.Cleanup(func() { _ = clientService.Close() })
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := clientService.SendAssociationRequest(ctx, verificationAssociationRequest()); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}
	if _, err := clientService.ReceiveAssociationResponse(ctx); err != nil {
		t.Fatalf("ReceiveAssociationResponse() error = %v", err)
	}
	waitForActiveService(t, srv)

	// The second peer stays silent: the rejection path must time it out.
	secondServerConn, secondClientConn := net.Pipe()
	t.Cleanup(func() { _ = secondClientConn.Close() })
	listener.conns <- secondServerConn
	waitForServerCondition(t, func() bool { return listener.acceptCount() >= 2 })

	_ = secondClientConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	var buf [1]byte
	if _, err := secondClientConn.Read(buf[:]); err == nil {
		t.Fatal("over-limit silent connection was not closed")
	} else {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			t.Fatal("over-limit silent connection was not closed by the reject negotiation timeout")
		}
	}
	if got := srv.ActiveConnections(); got != 1 {
		t.Fatalf("ActiveConnections() = %d, want only the established association", got)
	}
}

// TestOverLimitRejectNegotiationsAreBounded verifies that only MaxConnections
// over-limit reject negotiations run concurrently. A further over-limit
// connection is dropped immediately instead of spawning another goroutine.
func TestOverLimitRejectNegotiationsAreBounded(t *testing.T) {
	srv := New(WithMaxConnections(1), WithAssociationTimeout(10*time.Second))
	listener := newCountingListener()
	srv.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- srv.ListenAndServe(context.Background()) }()
	t.Cleanup(func() {
		_ = srv.Close()
		<-serveDone
	})

	// Keep the only connection slot busy with a silent negotiating peer.
	firstServerConn, firstClientConn := net.Pipe()
	t.Cleanup(func() { _ = firstClientConn.Close() })
	listener.conns <- firstServerConn
	waitForServerCondition(t, func() bool {
		return listener.acceptCount() >= 1 && srv.connectionLimitReached()
	})

	// The first over-limit connection may hold the single reject slot.
	secondServerConn, secondClientConn := net.Pipe()
	t.Cleanup(func() { _ = secondClientConn.Close() })
	listener.conns <- secondServerConn
	waitForServerCondition(t, func() bool { return listener.acceptCount() >= 2 })
	waitForServerCondition(t, func() bool { return len(srv.rejectSlots) == 1 })

	// The next over-limit connection must be dropped without negotiation.
	thirdServerConn, thirdClientConn := net.Pipe()
	t.Cleanup(func() { _ = thirdClientConn.Close() })
	listener.conns <- thirdServerConn
	waitForServerCondition(t, func() bool { return listener.acceptCount() >= 3 })

	_ = thirdClientConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	var buf [1]byte
	_, err := thirdClientConn.Read(buf[:])
	if err == nil {
		t.Fatal("connection beyond the reject-slot bound was not closed")
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		t.Fatal("connection beyond the reject-slot bound was not closed promptly")
	}
}

// TestShutdownClosesNegotiatingConnections verifies that graceful shutdown does
// not wait for the negotiation bound of silent peers: raw negotiating
// connections (including over-limit rejects) are closed before wg.Wait.
func TestShutdownClosesNegotiatingConnections(t *testing.T) {
	srv := New(WithMaxConnections(1), WithAssociationTimeout(10*time.Second))
	listener := newCountingListener()
	srv.listenFn = func(string, string, ...transport.ListenOption) (serverListener, error) {
		return listener, nil
	}

	serveDone := make(chan error, 1)
	go func() { serveDone <- srv.ListenAndServe(context.Background()) }()
	t.Cleanup(func() {
		_ = srv.Close()
		<-serveDone
	})

	firstServerConn, firstClientConn := net.Pipe()
	t.Cleanup(func() { _ = firstClientConn.Close() })
	listener.conns <- pipeWithRemoteAddress(firstServerConn, 40001)
	waitForServerCondition(t, func() bool {
		return listener.acceptCount() >= 1 && srv.connectionLimitReached()
	})

	secondServerConn, secondClientConn := net.Pipe()
	t.Cleanup(func() { _ = secondClientConn.Close() })
	listener.conns <- pipeWithRemoteAddress(secondServerConn, 40002)
	waitForServerCondition(t, func() bool { return listener.acceptCount() >= 2 })

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	shutdownDone := make(chan error, 1)
	go func() { shutdownDone <- srv.Shutdown(ctx) }()
	select {
	case err := <-shutdownDone:
		if err != nil {
			t.Fatalf("Shutdown() error = %v, want nil after closing negotiating connections", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Shutdown waited for the negotiation bound of a silent peer")
	}
}

// TestRejectNegotiationContextBounds verifies the min(AssociationTimeout,
// rejectNegotiationTimeout) rule that bounds silent over-limit connections.
func TestRejectNegotiationContextBounds(t *testing.T) {
	tests := []struct {
		name               string
		associationTimeout time.Duration
		min                time.Duration
		max                time.Duration
	}{
		{"disabled association timeout uses default bound", 0, 29 * time.Second, 31 * time.Second},
		{"shorter association timeout wins", 50 * time.Millisecond, 0, time.Second},
		{"longer association timeout is clamped", time.Hour, 29 * time.Second, 31 * time.Second},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := New(WithAssociationTimeout(test.associationTimeout))
			ctx, cancel := srv.rejectNegotiationContext(context.Background())
			deadline, ok := ctx.Deadline()
			cancel()
			if !ok {
				t.Fatal("reject negotiation context has no deadline")
			}
			remaining := time.Until(deadline)
			if remaining < test.min || remaining > test.max {
				t.Fatalf("remaining = %s, want within [%s, %s]", remaining, test.min, test.max)
			}
		})
	}
}

// TestConnectionSlotAccountingEnforcesMaxConnections verifies the slot counter
// that backs the limit.
func TestConnectionSlotAccountingEnforcesMaxConnections(t *testing.T) {
	s := New(WithMaxConnections(1))

	if !s.acquireConnectionSlot() {
		t.Fatal("first connection slot must be granted")
	}
	if s.acquireConnectionSlot() {
		t.Fatal("second connection slot must be refused while the first is in flight")
	}
	if !s.connectionLimitReached() {
		t.Fatal("connectionLimitReached() = false, want true")
	}

	s.releaseConnectionSlot()
	if s.connectionLimitReached() {
		t.Fatal("connectionLimitReached() = true after release, want false")
	}
	if !s.acquireConnectionSlot() {
		t.Fatal("slot must be reusable after release")
	}
}

// TestConnectionSlotUnlimitedWhenMaxIsZero verifies that the default
// configuration never refuses a connection.
func TestConnectionSlotUnlimitedWhenMaxIsZero(t *testing.T) {
	s := New()
	for i := 0; i < 1000; i++ {
		if !s.acquireConnectionSlot() {
			t.Fatalf("connection %d refused with MaxConnections=0", i)
		}
	}
}
