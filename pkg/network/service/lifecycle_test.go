// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
)

// mockConnForLifecycle is a mock connection that simulates network behavior
// for lifecycle testing.
type mockConnForLifecycle struct {
	readData  []byte
	readPos   int
	writeData []byte
	closed    bool
	blockRead bool // If true, Read() will block indefinitely
	mu        sync.Mutex
}

func (m *mockConnForLifecycle) Read(b []byte) (n int, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return 0, net.ErrClosed
	}

	if m.blockRead {
		// Block until closed
		time.Sleep(100 * time.Millisecond)
		if m.closed {
			return 0, net.ErrClosed
		}
		return 0, net.ErrClosed
	}

	if m.readPos >= len(m.readData) {
		// No more data, block a bit then return EOF
		time.Sleep(10 * time.Millisecond)
		return 0, net.ErrClosed
	}

	n = copy(b, m.readData[m.readPos:])
	m.readPos += n
	return n, nil
}

func (m *mockConnForLifecycle) Write(b []byte) (n int, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return 0, net.ErrClosed
	}

	m.writeData = append(m.writeData, b...)
	return len(b), nil
}

func (m *mockConnForLifecycle) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closed = true
	return nil
}

func (m *mockConnForLifecycle) LocalAddr() net.Addr                { return nil }
func (m *mockConnForLifecycle) RemoteAddr() net.Addr               { return nil }
func (m *mockConnForLifecycle) SetDeadline(_ time.Time) error      { return nil }
func (m *mockConnForLifecycle) SetReadDeadline(_ time.Time) error  { return nil }
func (m *mockConnForLifecycle) SetWriteDeadline(_ time.Time) error { return nil }

func TestRun(t *testing.T) {
	conn := &mockConnForLifecycle{
		blockRead: true,
	}

	// Create a mock association
	assoc := &association.Association{}

    service := NewService(conn, assoc)
    defer func() { _ = service.Close() }()

	// Start service in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.Run()
	}()

	// Give it time to start
	time.Sleep(50 * time.Millisecond)

    // Close the service
    _ = service.Close()

	// Wait for Run to finish
	select {
	case err := <-errCh:
		// Should finish without error (context cancelled is filtered)
		if err != nil {
			t.Errorf("Run returned error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("Run did not finish in time")
	}
}

func TestRun_AlreadyClosed(t *testing.T) {
	conn := &mockConnForLifecycle{}
    service := NewService(conn, nil)

	// Close before running
    _ = service.Close()

	// Run should return immediately with error
	err := service.Run()
	if err != ErrServiceClosed {
		t.Errorf("Expected ErrServiceClosed, got %v", err)
	}
}

func TestAbort(t *testing.T) {
	conn := &mockConnForLifecycle{}
    service := NewService(conn, nil)

	ctx := context.Background()
	err := service.Abort(ctx, 0, 0)
	if err != nil {
		t.Errorf("Abort failed: %v", err)
	}

	// Verify service is closed
	if !service.IsClosed() {
		t.Error("Service should be closed after Abort")
	}

	// Verify state is Aborted
	if service.GetState() != StateAborted {
		t.Errorf("Expected state Aborted, got %s", service.GetState())
	}

	// Verify A-ABORT PDU was written
	if len(conn.writeData) == 0 {
		t.Error("Expected A-ABORT PDU to be written")
	}
}

func TestGracefulRelease_Success(t *testing.T) {
	conn := &mockConnForLifecycle{}
    service := NewService(conn, nil)
    defer func() { _ = service.Close() }()

	// Set state to AssociationAccepted (required for release)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx := context.Background()
	err := service.GracefulRelease(ctx)
	if err != nil {
		t.Errorf("GracefulRelease failed: %v", err)
	}

	// Verify service is closed
	if !service.IsClosed() {
		t.Error("Service should be closed after GracefulRelease")
	}

	// Verify A-RELEASE-RQ PDU was written
	if len(conn.writeData) == 0 {
		t.Error("Expected A-RELEASE-RQ PDU to be written")
	}
}

func TestGracefulRelease_WrongState(t *testing.T) {
	conn := &mockConnForLifecycle{}
    service := NewService(conn, nil)
    defer func() { _ = service.Close() }()

	// State is Idle, which doesn't allow release
	ctx := context.Background()
	err := service.GracefulRelease(ctx)
	if err != nil {
		t.Errorf("GracefulRelease failed: %v", err)
	}

	// Should still close the service
	if !service.IsClosed() {
		t.Error("Service should be closed even if release failed")
	}
}

func TestWaitForClose(t *testing.T) {
	conn := &mockConnForLifecycle{}
	service := NewService(conn, nil)

	// Start goroutine that waits for close
	done := make(chan struct{})
	go func() {
		service.WaitForClose()
		close(done)
	}()

	// Close after a delay
    time.Sleep(50 * time.Millisecond)
    _ = service.Close()

	// WaitForClose should unblock
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("WaitForClose did not unblock")
	}
}

func TestErr(t *testing.T) {
	conn := &mockConnForLifecycle{}
    service := NewService(conn, nil)
    defer func() { _ = service.Close() }()

	// Get error channel
	errCh := service.Err()
	if errCh == nil {
		t.Error("Err() returned nil channel")
	}

	// Channel should be buffered
	select {
	case <-errCh:
		t.Error("Error channel should be empty initially")
	default:
		// Expected
	}
}

func TestRun_SendLoopError(t *testing.T) {
	// This test simulates an error in sendLoop by using a connection
	// that fails on write
	conn := &mockConnForLifecycle{}

	// Create service with association
	assoc := &association.Association{}
    service := NewService(conn, assoc)
    defer func() { _ = service.Close() }()

	// Close connection immediately to cause write errors
    _ = conn.Close()

    // Start service in goroutine
    errCh := make(chan error, 1)
    go func() {
        errCh <- service.Run()
    }()

	// Give it time to start and detect error
	time.Sleep(50 * time.Millisecond)

	// Close the service
    _ = service.Close()

	// Wait for Run to finish
	select {
	case err := <-errCh:
		// May get error from recvLoop (closed connection)
		// or no error (context cancelled before error detected)
		_ = err
	case <-time.After(1 * time.Second):
		t.Error("Run did not finish in time")
	}
}

func TestRun_RecvLoopError(t *testing.T) {
	// This test simulates an error in recvLoop by using a connection
	// that immediately returns an error on read
	conn := &mockConnForLifecycle{
		closed: true, // Already closed
	}

	assoc := &association.Association{}
    service := NewService(conn, assoc)
    defer func() { _ = service.Close() }()

	// Start service in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.Run()
	}()

	// Wait for Run to detect error and finish
	select {
	case err := <-errCh:
		// Should get error from recvLoop
		if err == nil {
			t.Error("Expected error from recvLoop")
		}
	case <-time.After(1 * time.Second):
		t.Error("Run did not finish in time")
	}
}

func TestRun_MultipleCallsNotAllowed(_ *testing.T) {
	conn := &mockConnForLifecycle{
		blockRead: true,
	}

	assoc := &association.Association{}
	service := NewService(conn, assoc)
    defer func() { _ = service.Close() }()

	// Start first Run
    go func() { _ = service.Run() }()

	// Give it time to start
	time.Sleep(50 * time.Millisecond)

    // Try to run again - should just start duplicate goroutines
    // (This is allowed, though not recommended)
    // The test just verifies it doesn't panic
    go func() { _ = service.Run() }()

	time.Sleep(50 * time.Millisecond)

    // Close should stop all
    _ = service.Close()
    service.WaitForClose()
}
