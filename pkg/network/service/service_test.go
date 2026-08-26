// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

func TestNewService(t *testing.T) {
	// Create a mock connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service with default options
	service := NewService(client, nil)
	if service == nil {
		t.Fatal("NewService returned nil")
	}

	// Verify default config was applied
	if service.config == nil {
		t.Fatal("Config should not be nil")
	}

	if service.config.maxPDULength != 16384 {
		t.Errorf("Expected maxPDULength 16384, got %d", service.config.maxPDULength)
	}

	if service.config.readTimeout != 0 {
		t.Errorf("Expected readTimeout 0, got %v", service.config.readTimeout)
	}

	// Verify initial state
	if service.GetState() != StateIdle {
		t.Errorf("Expected initial state Idle, got %s", service.GetState())
	}

	// Verify channels are created
	if service.sendQueue == nil {
		t.Error("sendQueue should not be nil")
	}

	if service.closeCh == nil {
		t.Error("closeCh should not be nil")
	}

	// Clean up
	_ = service.Close()
}

func TestNewServiceWithCustomOptions(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil,
		WithMaxPDULength(32768),
		WithReadTimeout(10*time.Second),
		WithWriteTimeout(10*time.Second),
		WithDIMSETimeout(30*time.Second),
		WithSendQueueSize(50))
	defer func() { _ = service.Close() }()

	if service.config.maxPDULength != 32768 {
		t.Errorf("Expected maxPDULength 32768, got %d", service.config.maxPDULength)
	}

	if service.config.readTimeout != 10*time.Second {
		t.Errorf("Expected readTimeout 10s, got %v", service.config.readTimeout)
	}

	if cap(service.sendQueue) != 50 {
		t.Errorf("Expected sendQueue capacity 50, got %d", cap(service.sendQueue))
	}

}

func TestServiceSetGetAssociation(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Initially no association
	if service.GetAssociation() != nil {
		t.Error("Expected nil association initially")
	}

	// Set association
	assoc := association.NewAssociation("CALLING_AE", "CALLED_AE")
	service.SetAssociation(assoc)

	// Verify association is set
	retrieved := service.GetAssociation()
	if retrieved != assoc {
		t.Error("GetAssociation returned different association")
	}
}

func TestServiceSetHandlers(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Set handlers
	handlers := &Handlers{
		CEchoHandler: func(_ context.Context, _ *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			return nil, nil
		},
	}

	service.SetHandlers(handlers)

	if service.handlers == nil {
		t.Error("Handlers should not be nil after SetHandlers")
	}
}

func TestServiceGetState(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Initial state
	if service.GetState() != StateIdle {
		t.Errorf("Expected initial state Idle, got %s", service.GetState())
	}
}

func TestServiceSetState(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Valid transition: Idle -> AssociationRequested
	err := service.setState(StateAssociationRequested)
	if err != nil {
		t.Errorf("Valid transition failed: %v", err)
	}

	if service.GetState() != StateAssociationRequested {
		t.Errorf("State not updated correctly, got %s", service.GetState())
	}

	// Invalid transition: AssociationRequested -> Transferring
	err = service.setState(StateTransferring)
	if err == nil {
		t.Error("Invalid transition should return error")
	}

	// State should not change on invalid transition
	if service.GetState() != StateAssociationRequested {
		t.Errorf("State should not change on invalid transition, got %s", service.GetState())
	}
}

func TestServiceClose(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()

	service := NewService(client, nil)

	// Verify service is not closed initially
	if service.IsClosed() {
		t.Error("Service should not be closed initially")
	}

	// Close service
	err := service.Close()
	if err != nil {
		t.Errorf("Close returned error: %v", err)
	}

	// Verify service is closed
	if !service.IsClosed() {
		t.Error("Service should be closed after Close()")
	}

	// Verify state is closed
	if service.GetState() != StateClosed {
		t.Errorf("Expected state Closed, got %s", service.GetState())
	}

	// Verify context is cancelled
	select {
	case <-service.Context().Done():
		// Expected
	default:
		t.Error("Service context should be cancelled after Close()")
	}

	// Multiple Close() calls should be safe
	err = service.Close()
	if err != nil {
		t.Errorf("Second Close() returned error: %v", err)
	}
}

func TestServiceIsClosed(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()

	service := NewService(client, nil)

	if service.IsClosed() {
		t.Error("Service should not be closed initially")
	}

	_ = service.Close()

	if !service.IsClosed() {
		t.Error("Service should be closed after Close()")
	}
}

func TestServiceContext(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	ctx := service.Context()
	if ctx == nil {
		t.Fatal("Context should not be nil")
	}

	// Context should not be cancelled initially
	select {
	case <-ctx.Done():
		t.Error("Context should not be cancelled initially")
	default:
		// Expected
	}

	// Close service
	_ = service.Close()

	// Context should be cancelled after close
	select {
	case <-ctx.Done():
		// Expected
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should be cancelled after Close()")
	}
}

func TestDefaultConfig(t *testing.T) {
	config := defaultServiceConfig()

	if config.maxPDULength != 16384 {
		t.Errorf("Expected maxPDULength 16384, got %d", config.maxPDULength)
	}

	if config.readTimeout != 0 {
		t.Errorf("Expected readTimeout 0, got %v", config.readTimeout)
	}

	if config.writeTimeout != 30*time.Second {
		t.Errorf("Expected writeTimeout 30s, got %v", config.writeTimeout)
	}

	if config.handlerShutdownTimeout != 60*time.Second {
		t.Errorf("Expected handlerShutdownTimeout 60s, got %v", config.handlerShutdownTimeout)
	}

	if config.requestTimeout != 0 {
		t.Errorf("Expected requestTimeout 0, got %v", config.requestTimeout)
	}

	if config.sendQueueSize != 100 {
		t.Errorf("Expected sendQueueSize 100, got %d", config.sendQueueSize)
	}

}

func TestServiceOptions(t *testing.T) {
	config := defaultServiceConfig()

	// Test WithMaxPDULength
	WithMaxPDULength(32768)(config)
	if config.maxPDULength != 32768 {
		t.Errorf("WithMaxPDULength failed")
	}

	// Test WithReadTimeout
	WithReadTimeout(10 * time.Second)(config)
	if config.readTimeout != 10*time.Second {
		t.Errorf("WithReadTimeout failed")
	}

	// Test WithWriteTimeout
	WithWriteTimeout(15 * time.Second)(config)
	if config.writeTimeout != 15*time.Second {
		t.Errorf("WithWriteTimeout failed")
	}

	// Test WithDIMSETimeout
	WithDIMSETimeout(45 * time.Second)(config)
	if config.handlerShutdownTimeout != 45*time.Second {
		t.Errorf("WithDIMSETimeout compatibility alias failed")
	}

	WithHandlerShutdownTimeout(50 * time.Second)(config)
	if config.handlerShutdownTimeout != 50*time.Second {
		t.Errorf("WithHandlerShutdownTimeout failed")
	}

	WithRequestTimeout(45 * time.Second)(config)
	if config.requestTimeout != 45*time.Second {
		t.Errorf("WithRequestTimeout failed")
	}

	// Test WithSendQueueSize
	WithSendQueueSize(200)(config)
	if config.sendQueueSize != 200 {
		t.Errorf("WithSendQueueSize failed")
	}

}

// TestReadTimeoutFromContext verifies the edge-case logic in readTimeoutFromContext.
func TestReadTimeoutFromContext(t *testing.T) {
	const tolerance = 10 * time.Millisecond

	inDelta := func(t *testing.T, got, want time.Duration, tol time.Duration, label string) {
		t.Helper()
		diff := got - want
		if diff < 0 {
			diff = -diff
		}
		if diff > tol {
			t.Errorf("%s: got %v, want ~%v (tolerance ±%v)", label, got, want, tol)
		}
	}

	tests := []struct {
		name       string
		configured time.Duration
		// makeCtx returns the context to use; deadline expressed as offset from now
		makeCtx func() context.Context
		// check is called with the returned duration; use wantExact for exact match
		check func(t *testing.T, got time.Duration)
	}{
		{
			name:       "no ctx deadline, configured > 0",
			configured: 5 * time.Second,
			makeCtx:    context.Background,
			check: func(t *testing.T, got time.Duration) {
				inDelta(t, got, 5*time.Second, tolerance, "no-ctx-deadline configured>0")
			},
		},
		{
			name:       "no ctx deadline, configured == 0",
			configured: 0,
			makeCtx:    context.Background,
			check: func(t *testing.T, got time.Duration) {
				if got != 0 {
					t.Errorf("no-ctx-deadline configured==0: got %v, want 0", got)
				}
			},
		},
		{
			name:       "ctx deadline > configured",
			configured: 5 * time.Second,
			makeCtx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
				t.Cleanup(cancel)
				return ctx
			},
			check: func(t *testing.T, got time.Duration) {
				inDelta(t, got, 5*time.Second, tolerance, "ctx-deadline>configured")
			},
		},
		{
			name:       "ctx deadline < configured",
			configured: 10 * time.Second,
			makeCtx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
				t.Cleanup(cancel)
				return ctx
			},
			check: func(t *testing.T, got time.Duration) {
				inDelta(t, got, 3*time.Second, tolerance, "ctx-deadline<configured")
			},
		},
		{
			name:       "ctx deadline only (configured == 0)",
			configured: 0,
			makeCtx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
				t.Cleanup(cancel)
				return ctx
			},
			check: func(t *testing.T, got time.Duration) {
				inDelta(t, got, 5*time.Second, tolerance, "ctx-only configured==0")
			},
		},
		{
			name:       "ctx deadline already expired",
			configured: 5 * time.Second,
			makeCtx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
				t.Cleanup(cancel)
				return ctx
			},
			check: func(t *testing.T, got time.Duration) {
				if got != time.Nanosecond {
					t.Errorf("expired ctx: got %v, want %v", got, time.Nanosecond)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.makeCtx()
			got := readTimeoutFromContext(ctx, tt.configured)
			tt.check(t, got)
		})
	}
}
