// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

func TestSend(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create association
	assoc := createTestAssociation()

	// Create service
	service := NewService(client, assoc, WithAssociationRequestor(true))
	defer func() { _ = service.Close() }()

	// Set state to allow DIMSE sending
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState failed: %v", err)
	}

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
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

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
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)

	// Close service before sending
	_ = service.Close()

	// Try to send message
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

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
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service (don't start it so send queue blocks)
	assoc := createTestAssociation()
	service := NewService(client, assoc, WithAssociationRequestor(true))
	defer func() { _ = service.Close() }()

	// Set state to allow DIMSE sending
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState failed: %v", err)
	}

	// Create C-ECHO request
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

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
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create association
	assoc := createTestAssociation()

	// Create service
	service := NewService(client, assoc, WithAssociationRequestor(true))
	defer func() { _ = service.Close() }()

	// Set state to allow DIMSE sending
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState failed: %v", err)
	}

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
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Send message with timeout
	err := service.SendWithTimeout(req, 1*time.Second)
	if err != nil {
		t.Fatalf("SendWithTimeout failed: %v", err)
	}
}

func TestRegisterUnregisterPendingRequest(t *testing.T) {
	service := NewService(nil, nil)
	defer func() { _ = service.Close() }()

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
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	assoc := createTestAssociation()
	service := NewService(client, assoc, WithAssociationRequestor(true))
	defer func() { _ = service.Close() }()

	// Set state to one that doesn't allow sending
	if err := service.setState(StateClosed); err != nil {
		t.Fatalf("setState failed: %v", err)
	}

	// Try to send message
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

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
	defer func() { _ = service.Close() }()

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
	defer func() { _ = service.Close() }()

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
	defer func() { _ = service.Close() }()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := service.SendCFind(ctx, req)
	if err == nil {
		t.Error("Expected error when no association available, got nil")
	}
}

func TestSendCMove_SendFailureCleansPendingRequest(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	defer func() { _ = service.Close() }()

	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch, err := service.SendCMove(ctx, req)
	if err == nil {
		t.Fatal("SendCMove() error = nil, want send failure in Idle state")
	}
	if ch != nil {
		t.Fatalf("SendCMove() channel = %v, want nil on send failure", ch)
	}

	service.pendingRequestsMu.RLock()
	pendingCount := len(service.pendingRequests)
	service.pendingRequestsMu.RUnlock()
	if pendingCount != 0 {
		t.Fatalf("pending request count = %d, want 0", pendingCount)
	}
}

func TestSendCCancel_NoAssociation(t *testing.T) {
	service := NewService(nil, nil)
	defer func() { _ = service.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := service.SendCCancel(ctx, 123, 1)
	if err == nil {
		t.Fatal("SendCCancel() error = nil, want no association error")
	}
}

func TestSendCCancel_SendsRequestWithoutRegisteringPendingResponse(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	defer func() { _ = service.Close() }()

	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState failed: %v", err)
	}

	sent := make(chan dimse.Message, 1)
	go func() {
		req := <-service.sendQueue
		sent <- req.message
		req.resultCh <- nil
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	const messageIDBeingRespondedTo uint16 = 123
	const presentationContextID byte = 5
	if err := service.SendCCancel(ctx, messageIDBeingRespondedTo, presentationContextID); err != nil {
		t.Fatalf("SendCCancel() error = %v", err)
	}

	select {
	case msg := <-sent:
		req, ok := msg.(*dimse.CCancelRequest)
		if !ok {
			t.Fatalf("sent message type = %T, want *dimse.CCancelRequest", msg)
		}
		if req.MessageIDBeingRespondedTo() != messageIDBeingRespondedTo {
			t.Fatalf("MessageIDBeingRespondedTo = %d, want %d", req.MessageIDBeingRespondedTo(), messageIDBeingRespondedTo)
		}
		if req.PresentationContextID() != presentationContextID {
			t.Fatalf("PresentationContextID = %d, want %d", req.PresentationContextID(), presentationContextID)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-CANCEL send")
	}

	service.pendingRequestsMu.RLock()
	pendingCount := len(service.pendingRequests)
	service.pendingRequestsMu.RUnlock()
	if pendingCount != 0 {
		t.Fatalf("pending request count = %d, want 0", pendingCount)
	}
}

func newAcceptedTestService(t *testing.T) *Service {
	t.Helper()
	service := NewService(nil, createTestAssociation())
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	t.Cleanup(func() { _ = service.Close() })
	return service
}

func TestServiceSendNCreate(t *testing.T) {
	service := newAcceptedTestService(t)
	attributes := dataset.New()
	if err := attributes.Add(element.NewString(tag.FilmSessionLabel, vr.LO, []string{"session"})); err != nil {
		t.Fatalf("Dataset.Add() error = %v", err)
	}
	req := dimse.NewNCreateRequest("1.2.840.10008.5.1.1.1", "2.25.401", attributes)
	go func() {
		send := <-service.sendQueue
		got, ok := send.message.(*dimse.NCreateRequest)
		if !ok {
			t.Errorf("sent message type = %T, want *dimse.NCreateRequest", send.message)
			send.resultCh <- nil
			return
		}
		if got.AffectedSOPInstanceUID() != "2.25.401" || got.DataDataset() != attributes {
			t.Errorf("sent N-CREATE did not preserve UID and attributes")
		}
		send.resultCh <- nil
		_ = service.handleResponse(dimse.NewNCreateResponseSuccess(got.MessageID(), got.AffectedSOPClassUID(), got.AffectedSOPInstanceUID(), nil))
	}()

	resp, err := service.SendNCreate(context.Background(), req)
	if err != nil {
		t.Fatalf("SendNCreate() error = %v", err)
	}
	if resp.AffectedSOPInstanceUID() != "2.25.401" {
		t.Errorf("response SOP Instance UID = %q", resp.AffectedSOPInstanceUID())
	}
}

func TestServiceSendNSet(t *testing.T) {
	service := newAcceptedTestService(t)
	modifications := dataset.New()
	if err := modifications.Add(element.NewUnsignedShort(tag.ImageBoxPosition, []uint16{3})); err != nil {
		t.Fatalf("Dataset.Add() error = %v", err)
	}
	req := dimse.NewNSetRequest("1.2.840.10008.5.1.1.4", "2.25.402", modifications)
	go func() {
		send := <-service.sendQueue
		got, ok := send.message.(*dimse.NSetRequest)
		if !ok {
			t.Errorf("sent message type = %T, want *dimse.NSetRequest", send.message)
			send.resultCh <- nil
			return
		}
		if got.RequestedSOPInstanceUID() != "2.25.402" || got.DataDataset() != modifications {
			t.Errorf("sent N-SET did not preserve UID and modifications")
		}
		send.resultCh <- nil
		_ = service.handleResponse(dimse.NewNSetResponseSuccess(got.MessageID(), got.RequestedSOPClassUID(), got.RequestedSOPInstanceUID(), nil))
	}()

	resp, err := service.SendNSet(context.Background(), req)
	if err != nil {
		t.Fatalf("SendNSet() error = %v", err)
	}
	if resp.AffectedSOPInstanceUID() != "2.25.402" {
		t.Errorf("response SOP Instance UID = %q", resp.AffectedSOPInstanceUID())
	}
}

func TestServiceSendNAction(t *testing.T) {
	service := newAcceptedTestService(t)
	req := dimse.NewNActionRequest("1.2.840.10008.5.1.1.1", "2.25.403", 1, nil)
	go func() {
		send := <-service.sendQueue
		got, ok := send.message.(*dimse.NActionRequest)
		if !ok {
			t.Errorf("sent message type = %T, want *dimse.NActionRequest", send.message)
			send.resultCh <- nil
			return
		}
		if got.RequestedSOPInstanceUID() != "2.25.403" || got.ActionTypeID() != 1 {
			t.Errorf("sent N-ACTION did not preserve UID and action type")
		}
		send.resultCh <- nil
		_ = service.handleResponse(dimse.NewNActionResponseSuccess(got.MessageID(), got.RequestedSOPClassUID(), got.RequestedSOPInstanceUID(), got.ActionTypeID(), nil))
	}()

	resp, err := service.SendNAction(context.Background(), req)
	if err != nil {
		t.Fatalf("SendNAction() error = %v", err)
	}
	if resp.ActionTypeID() != 1 {
		t.Errorf("response Action Type ID = %d, want 1", resp.ActionTypeID())
	}
}

func TestServiceSendNDelete(t *testing.T) {
	service := newAcceptedTestService(t)
	req := dimse.NewNDeleteRequest("1.2.840.10008.5.1.1.2", "2.25.404")
	go func() {
		send := <-service.sendQueue
		got, ok := send.message.(*dimse.NDeleteRequest)
		if !ok {
			t.Errorf("sent message type = %T, want *dimse.NDeleteRequest", send.message)
			send.resultCh <- nil
			return
		}
		if got.RequestedSOPInstanceUID() != "2.25.404" {
			t.Errorf("sent N-DELETE SOP Instance UID = %q", got.RequestedSOPInstanceUID())
		}
		send.resultCh <- nil
		_ = service.handleResponse(dimse.NewNDeleteResponseSuccess(got.MessageID(), got.RequestedSOPClassUID(), got.RequestedSOPInstanceUID()))
	}()

	resp, err := service.SendNDelete(context.Background(), req)
	if err != nil {
		t.Fatalf("SendNDelete() error = %v", err)
	}
	if resp.AffectedSOPInstanceUID() != "2.25.404" {
		t.Errorf("response SOP Instance UID = %q", resp.AffectedSOPInstanceUID())
	}
}
