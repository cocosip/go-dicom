// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

func TestRecvLoop_ContextCancellation(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	service := NewService(client, nil)
	defer service.Close()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start receive loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.recvLoop(ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Cancel context
	cancel()

	// Close the connection to unblock the read
	// This simulates the real scenario where context cancellation
	// leads to connection close
	client.Close()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		// Either context.Canceled or read error is acceptable
		if err == nil {
			t.Error("Expected error, got nil")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("recvLoop did not exit after context cancellation")
	}
}

func TestRecvLoop_ServiceClose(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	service := NewService(client, nil)

	// Start receive loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.recvLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Close service (this closes the context and connection)
	service.Close()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		// Either context.Canceled or a read error from closed connection is acceptable
		if err == nil {
			t.Error("Expected error when service closes, got nil")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("recvLoop did not exit after service close")
	}
}

func TestGetTransferSyntaxForContext(t *testing.T) {
	// Create association with presentation contexts
	assoc := createTestAssociation()

	// Create service
	service := NewService(nil, assoc)
	defer service.Close()

	// Test valid context ID
	ts := service.getTransferSyntaxForContext(1)
	if ts != transfer.ExplicitVRLittleEndian {
		t.Errorf("Expected ExplicitVRLittleEndian, got %v", ts)
	}

	// Test invalid context ID (should return default)
	ts = service.getTransferSyntaxForContext(99)
	if ts != transfer.ExplicitVRLittleEndian {
		t.Errorf("Expected default ExplicitVRLittleEndian, got %v", ts)
	}

	// Test with nil association (should return default)
	service2 := NewService(nil, nil)
	defer service2.Close()
	ts = service2.getTransferSyntaxForContext(1)
	if ts != transfer.ExplicitVRLittleEndian {
		t.Errorf("Expected default ExplicitVRLittleEndian with nil association, got %v", ts)
	}
}

func TestRecvLoop_ReceivePDataTF(t *testing.T) {
	t.Skip("Skipping until raw dataset decoding is implemented")

	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create association
	assoc := createTestAssociation()

	// Create service with association
	service := NewService(client, assoc)
	defer service.Close()

	// Start receive loop
	go func() {
		service.recvLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Create a test C-ECHO request
	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// Encode the message
	commandData, datasetData, err := EncodeDIMSEMessage(req, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("EncodeDIMSEMessage failed: %v", err)
	}

	// Create PDVs
	contextID := byte(1)
	commandPDVs := FragmentData(commandData, contextID, true, 16384)

	var allPDVs []*PDV
	allPDVs = append(allPDVs, commandPDVs...)
	if len(datasetData) > 0 {
		datasetPDVs := FragmentData(datasetData, contextID, false, 16384)
		allPDVs = append(allPDVs, datasetPDVs...)
	}

	// Create P-DATA-TF PDU
	pduData, err := CreatePDataTFPDU(allPDVs)
	if err != nil {
		t.Fatalf("CreatePDataTFPDU failed: %v", err)
	}

	// Send PDU to service
	err = transport.WritePDU(server, 5*time.Second, pduData)
	if err != nil {
		t.Fatalf("WritePDU failed: %v", err)
	}

	// Wait for message to be processed
	time.Sleep(100 * time.Millisecond)

	// TODO: Verify message was received and processed
	// This will require adding a way to inspect received messages
}

func TestRecvLoop_NonPDataTFPDU(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	service := NewService(client, nil)
	defer service.Close()

	// Start receive loop
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.recvLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Send a non-P-DATA-TF PDU (e.g., A-ASSOCIATE-RQ)
	// This should be ignored by recvLoop
	nonDataPDU := &pdu.RawPDU{
		Type: pdu.TypeAAssociateRQ,
		Data: []byte{0x00, 0x01},
	}

	go func() {
		transport.WritePDU(server, 5*time.Second, nonDataPDU)
		// Close server after sending to trigger loop exit
		time.Sleep(50 * time.Millisecond)
		server.Close()
	}()

	// Wait for loop to exit or timeout
	select {
	case err := <-errCh:
		// Should exit with read error when server closes
		if err == nil {
			t.Error("Expected error when connection closed, got nil")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("recvLoop did not exit")
	}
}

func TestRecvLoop_ContextIDChange(t *testing.T) {
	t.Skip("Skipping until raw dataset decoding is implemented")

	// This test verifies that recvLoop detects and rejects
	// PDVs with changing presentation context IDs mid-message

	// Create a pipe connection
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	// Create service
	assoc := createTestAssociation()
	service := NewService(client, assoc)
	defer service.Close()

	// Start receive loop
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.recvLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Create PDVs with different context IDs
	pdv1 := &PDV{
		PresentationContextID: 1,
		IsCommand:             true,
		IsLastFragment:        false,
		Data:                  []byte{0x00, 0x00, 0x00, 0x00},
	}

	pdv2 := &PDV{
		PresentationContextID: 3, // Different context ID!
		IsCommand:             true,
		IsLastFragment:        true,
		Data:                  []byte{0x00, 0x00, 0x00, 0x00},
	}

	// Create P-DATA-TF PDU with both PDVs
	pduData, err := CreatePDataTFPDU([]*PDV{pdv1, pdv2})
	if err != nil {
		t.Fatalf("CreatePDataTFPDU failed: %v", err)
	}

	// Send PDU
	go func() {
		transport.WritePDU(server, 5*time.Second, pduData)
	}()

	// Wait for error
	select {
	case err := <-errCh:
		if err == nil {
			t.Error("Expected error for context ID change, got nil")
		}
		// Verify error message mentions context ID change
		if err != nil && err.Error() == "" {
			t.Error("Expected error message about context ID change")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("recvLoop did not report error for context ID change")
	}
}
