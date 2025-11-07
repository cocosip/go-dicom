// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

func TestSend(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create association
	assoc := createTestAssociation()

	// Create service
	service := NewService(client, assoc)
	defer service.Close()

	// Set state to allow DIMSE sending
	service.setState(StateAssociationAccepted)

	// Start service
	if err := service.Start(); err != nil {
		t.Fatalf("Failed to start service: %v", err)
	}

	// Start goroutine to read from server side
	go func() {
		buf := make([]byte, 4096)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	// Create C-ECHO request
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// Send message
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := service.Send(ctx, req)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
}

func TestSend_ServiceClosed(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	service := NewService(client, nil)

	// Close service before sending
	service.Close()

	// Try to send message
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := service.Send(ctx, req)
	if err != ErrServiceClosed {
		t.Errorf("Expected ErrServiceClosed, got %v", err)
	}
}

func TestSend_ContextCancellation(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service (don't start it so send queue blocks)
	assoc := createTestAssociation()
	service := NewService(client, assoc)
	defer service.Close()

	// Set state to allow DIMSE sending
	service.setState(StateAssociationAccepted)

	// Create C-ECHO request
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// Create context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Try to send message
	err := service.Send(ctx, req)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func TestSendWithTimeout(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create association
	assoc := createTestAssociation()

	// Create service
	service := NewService(client, assoc)
	defer service.Close()

	// Set state to allow DIMSE sending
	service.setState(StateAssociationAccepted)

	// Start service
	if err := service.Start(); err != nil {
		t.Fatalf("Failed to start service: %v", err)
	}

	// Start goroutine to read from server side
	go func() {
		buf := make([]byte, 4096)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	// Create C-ECHO request
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// Send message with timeout
	err := service.SendWithTimeout(req, 1*time.Second)
	if err != nil {
		t.Fatalf("SendWithTimeout failed: %v", err)
	}
}

func TestRegisterUnregisterPendingRequest(t *testing.T) {
	service := NewService(nil, nil)
	defer service.Close()

	// Register a pending request
	respCh := make(chan dimse.Response, 1)
	messageID := uint16(123)
	req := dimse.NewCEchoRequest()

	service.registerPendingRequest(messageID, req, respCh)

	// Check it was registered
	service.pendingRequestsMu.Lock()
	_, exists := service.pendingRequests[messageID]
	service.pendingRequestsMu.Unlock()

	if !exists {
		t.Error("Pending request was not registered")
	}

	// Unregister it
	service.unregisterPendingRequest(messageID)

	// Check it was unregistered
	service.pendingRequestsMu.Lock()
	_, exists = service.pendingRequests[messageID]
	service.pendingRequestsMu.Unlock()

	if exists {
		t.Error("Pending request was not unregistered")
	}
}

func TestSend_WrongState(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	assoc := createTestAssociation()
	service := NewService(client, assoc)
	defer service.Close()

	// Set state to one that doesn't allow sending
	service.setState(StateClosed)

	// Try to send message
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := service.Send(ctx, req)
	if err == nil {
		t.Error("Expected error when sending in closed state, got nil")
	}
}

func TestSendCEcho_NoAssociation(t *testing.T) {
	// Create service without association
	service := NewService(nil, nil)
	defer service.Close()

	req := dimse.NewCEchoRequest()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := service.SendCEcho(ctx, req)
	if err == nil {
		t.Error("Expected error when no association available, got nil")
	}
}

func TestSendCStore_NoAssociation(t *testing.T) {
	// Create service without association
	service := NewService(nil, nil)
	defer service.Close()

	// For this test, we just want to verify the "no association" error,
	// so we can skip creating a full valid C-STORE request
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Pass nil to trigger the association check first
	_, err := service.SendCStore(ctx, nil)
	if err == nil {
		t.Error("Expected error when no association available, got nil")
	}
}

func TestSendCFind_NoAssociation(t *testing.T) {
	// Create service without association
	service := NewService(nil, nil)
	defer service.Close()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := service.SendCFind(ctx, req)
	if err == nil {
		t.Error("Expected error when no association available, got nil")
	}
}
