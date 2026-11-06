// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package server

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

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

// TestAcceptLoopStopsAtMaxConnectionsWhileNegotiating verifies that
// MaxConnections is enforced while a connection is still negotiating, so a peer
// cannot exceed the limit by opening connections that never associate.
func TestAcceptLoopStopsAtMaxConnectionsWhileNegotiating(t *testing.T) {
	srv := New(WithMaxConnections(1))
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

	// Hand the server one connection that never sends A-ASSOCIATE-RQ, so it
	// stays in the negotiating state for the whole test.
	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() { _ = clientConn.Close() })
	listener.conns <- serverConn

	deadline := time.Now().Add(2 * time.Second)
	for listener.acceptCount() == 0 {
		if time.Now().After(deadline) {
			t.Fatal("server never accepted the first connection")
		}
		time.Sleep(5 * time.Millisecond)
	}

	// Give the accept loop several chances to over-accept.
	time.Sleep(300 * time.Millisecond)

	if got := listener.acceptCount(); got != 1 {
		t.Fatalf("Accept called %d times, want 1 while one connection is negotiating", got)
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
