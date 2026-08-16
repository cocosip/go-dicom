// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"encoding/binary"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
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
	clientConn, serverConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()

	service := NewService(clientConn, nil)
	defer func() { _ = service.Close() }()

	// Set state to AssociationAccepted (required for release)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	if err := service.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	releaseSeen := make(chan struct{})
	go func() {
		raw, err := transport.ReadPDU(serverConn, time.Second)
		if err != nil {
			t.Errorf("ReadPDU() error = %v", err)
			return
		}
		if raw.Type != pdu.TypeAReleaseRQ {
			t.Errorf("PDU type = 0x%02x, want 0x%02x", raw.Type, pdu.TypeAReleaseRQ)
			return
		}
		close(releaseSeen)
		time.Sleep(50 * time.Millisecond)
		rp, err := pdu.NewAReleaseRP().Encode()
		if err != nil {
			t.Errorf("Encode() error = %v", err)
			return
		}
		if err := transport.WritePDU(serverConn, time.Second, rp); err != nil {
			t.Errorf("WritePDU() error = %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- service.GracefulRelease(ctx)
	}()

	select {
	case <-releaseSeen:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("timed out waiting for A-RELEASE-RQ")
	}

	select {
	case err := <-done:
		t.Fatalf("GracefulRelease() returned before A-RELEASE-RP: %v", err)
	case <-time.After(20 * time.Millisecond):
	}

	err := <-done
	if err != nil {
		t.Errorf("GracefulRelease failed: %v", err)
	}

	// Verify service is closed
	if !service.IsClosed() {
		t.Error("Service should be closed after GracefulRelease")
	}
}

func TestGracefulReleaseHandlesResponseBeforeRequestWriteReturns(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()

	responseHandled := make(chan struct{})
	service := NewService(&releaseResponseGateConn{
		Conn:            clientConn,
		responseHandled: responseHandled,
	}, nil)
	defer func() { _ = service.Close() }()

	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("set association state: %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	go func() {
		raw, err := transport.ReadPDU(serverConn, time.Second)
		if err != nil {
			t.Errorf("ReadPDU() error = %v", err)
			return
		}
		if raw.Type != pdu.TypeAReleaseRQ {
			t.Errorf("PDU type = 0x%02x, want 0x%02x", raw.Type, pdu.TypeAReleaseRQ)
			return
		}

		rp, err := pdu.NewAReleaseRP().Encode()
		if err != nil {
			t.Errorf("Encode() error = %v", err)
			return
		}
		if err := transport.WritePDU(serverConn, time.Second, rp); err != nil {
			t.Errorf("WritePDU() error = %v", err)
		}
	}()

	go func() {
		deadline := time.Now().Add(time.Second)
		for service.GetState() != StateClosed && time.Now().Before(deadline) {
			time.Sleep(time.Millisecond)
		}
		close(responseHandled)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := service.GracefulRelease(ctx); err != nil {
		t.Fatalf("GracefulRelease() error = %v", err)
	}
}

type releaseResponseGateConn struct {
	net.Conn
	responseHandled         <-chan struct{}
	releaseRequestRemaining int
}

func (c *releaseResponseGateConn) Write(data []byte) (int, error) {
	n, err := c.Conn.Write(data)
	if err != nil || n == 0 {
		return n, err
	}

	written := data[:n]
	if c.releaseRequestRemaining > 0 {
		c.releaseRequestRemaining -= len(written)
		if c.releaseRequestRemaining <= 0 {
			<-c.responseHandled
		}
		return n, nil
	}
	if len(written) < 6 || written[0] != pdu.TypeAReleaseRQ {
		return n, nil
	}

	c.releaseRequestRemaining = int(binary.BigEndian.Uint32(written[2:6]))
	if c.releaseRequestRemaining == 0 {
		<-c.responseHandled
	}
	return n, nil
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

func TestGracefulRelease_TimesOutWithoutReleaseResponse(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()

	service := NewService(clientConn, nil)
	defer func() { _ = service.Close() }()

	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	requestSeen := make(chan struct{})
	go func() {
		raw, err := transport.ReadPDU(serverConn, time.Second)
		if err != nil {
			t.Errorf("ReadPDU() error = %v", err)
			return
		}
		if raw.Type != pdu.TypeAReleaseRQ {
			t.Errorf("PDU type = 0x%02x, want 0x%02x", raw.Type, pdu.TypeAReleaseRQ)
			return
		}
		close(requestSeen)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- service.GracefulRelease(ctx)
	}()

	select {
	case <-requestSeen:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("timed out waiting for A-RELEASE-RQ")
	}

	select {
	case err := <-done:
		t.Fatalf("GracefulRelease() returned before context timeout: %v", err)
	case <-time.After(20 * time.Millisecond):
	}

	err := <-done
	if err == nil {
		t.Fatal("GracefulRelease() error = nil, want timeout")
	}
	if !service.IsClosed() {
		t.Error("Service should be closed after timed out graceful release")
	}
}

func TestSendReleaseRequest_ReleaseResponseClosesService(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()

	service := NewService(clientConn, nil)
	defer func() { _ = service.Close() }()

	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	peerDone := make(chan struct{})
	go func() {
		defer close(peerDone)

		raw, err := transport.ReadPDU(serverConn, time.Second)
		if err != nil {
			t.Errorf("ReadPDU() error = %v", err)
			return
		}
		if raw.Type != pdu.TypeAReleaseRQ {
			t.Errorf("PDU type = 0x%02x, want 0x%02x", raw.Type, pdu.TypeAReleaseRQ)
			return
		}

		rp, err := pdu.NewAReleaseRP().Encode()
		if err != nil {
			t.Errorf("Encode() error = %v", err)
			return
		}
		if err := transport.WritePDU(serverConn, time.Second, rp); err != nil {
			t.Errorf("WritePDU() error = %v", err)
		}
	}()

	if err := service.SendReleaseRequest(ctx); err != nil {
		t.Fatalf("SendReleaseRequest() error = %v", err)
	}

	done := make(chan struct{})
	go func() {
		service.WaitForClose()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("service did not close after receiving A-RELEASE-RP")
	}

	<-peerDone
}

func TestLoopError_OnConnectionClosedAfterRequestHandlersExit(t *testing.T) {
	requestRelease := make(chan struct{})
	requestDone := make(chan struct{})
	callbackDone := make(chan struct{})

	conn := &mockConnForLifecycle{closed: true}
	service := NewService(conn, nil, WithConnectionLifecycleHandler(&ConnectionLifecycleHandlerFuncs{
		OnConnectionClosedFunc: func(_ context.Context, _ error) {
			select {
			case <-requestDone:
			default:
				t.Error("OnConnectionClosed fired before request handlers exited")
			}
			close(callbackDone)
		},
	}))

	service.requestWg.Add(1)
	go func() {
		defer service.requestWg.Done()
		defer close(requestDone)
		<-requestRelease
	}()
	time.AfterFunc(50*time.Millisecond, func() {
		close(requestRelease)
	})

	if err := service.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	select {
	case <-callbackDone:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for OnConnectionClosed callback")
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

func TestRunReturnsLoopErrorAfterServiceClose(t *testing.T) {
	conn := &mockConnForLifecycle{closed: true}
	service := NewService(conn, &association.Association{})
	defer func() { _ = service.Close() }()

	err := service.Run()
	if err == nil {
		t.Fatal("Run() error = nil, want loop error")
	}
}

func TestRun_MultipleCallsNotAllowed(t *testing.T) {
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

	// Try to run again - should be rejected to avoid duplicate goroutines
	err := service.Run()
	if err != ErrServiceAlreadyStarted {
		t.Fatalf("Run() error = %v, want %v", err, ErrServiceAlreadyStarted)
	}

	// Close should stop all
	_ = service.Close()
	service.WaitForClose()
}
