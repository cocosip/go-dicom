// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// setupTestService creates a service ready for testing handlers.
// It sets the service to AssociationAccepted state and starts a mock sendLoop
// that consumes from sendQueue and responds with success.
func setupTestService(t *testing.T) (*Service, context.Context, context.CancelFunc) {
	service := NewService(nil, nil)

	// Set service to AssociationAccepted state (required to send responses)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	// Start a mock sendLoop that drains sendQueue and responds with success
	// This simulates the real sendLoop without needing a network connection
	go func() {
		for {
			select {
			case req := <-service.sendQueue:
				// Simulate successful send
				req.resultCh <- nil
			case <-service.closeCh:
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return service, ctx, cancel
}

func TestSetGetHandlers(t *testing.T) {
	service := NewService(nil, nil)
	defer service.Close()

	// Initially should be nil
	if h := service.GetHandlers(); h != nil {
		t.Error("Expected nil handlers initially")
	}

	// Set handlers
	handlers := &Handlers{
		CEchoHandler: func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
		},
	}
	service.SetHandlers(handlers)

	// Get handlers
	retrieved := service.GetHandlers()
	if retrieved == nil {
		t.Fatal("Expected handlers, got nil")
	}
	if retrieved.CEchoHandler == nil {
		t.Error("Expected CEchoHandler to be set")
	}
}

func TestHandleCEchoRequest_DefaultHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// No custom handlers - should use default
	err := service.handleCEchoRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCEchoRequest failed: %v", err)
	}
}

func TestHandleCEchoRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	req := dimse.NewCEchoRequest()
	req.SetMessageID(1)

	// Custom handler that returns success
	handlerCalled := false
	handlers := &Handlers{
		CEchoHandler: func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			handlerCalled = true
			return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
		},
	}

	err := service.handleCEchoRequest(ctx, req, handlers)
	if err != nil {
		t.Errorf("handleCEchoRequest failed: %v", err)
	}

	if !handlerCalled {
		t.Error("Custom handler was not called")
	}
}

func TestHandleCStoreRequest_DefaultHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	// Create dataset with required tags
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))       // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE request: %v", err)
	}
	req.SetMessageID(1)

	// No custom handlers - should use default
	err = service.handleCStoreRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCStoreRequest failed: %v", err)
	}
}

func TestHandleCStoreRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	// Create dataset with required tags
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))       // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))

	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE request: %v", err)
	}
	req.SetMessageID(1)

	// Custom handler
	handlerCalled := false
	handlers := &Handlers{
		CStoreHandler: func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
			handlerCalled = true
			return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
		},
	}

	err = service.handleCStoreRequest(ctx, req, handlers)
	if err != nil {
		t.Errorf("handleCStoreRequest failed: %v", err)
	}

	if !handlerCalled {
		t.Error("Custom handler was not called")
	}
}

func TestHandleCFindRequest_DefaultHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
	req.SetMessageID(1)

	// No custom handlers - should use default (returns success with no results)
	err := service.handleCFindRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCFindRequest failed: %v", err)
	}
}

func TestHandleCFindRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer service.Close()
	defer cancel()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
	req.SetMessageID(1)

	// Custom handler that returns multiple responses
	handlerCalled := false
	handlers := &Handlers{
		CFindHandler: func(ctx context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			handlerCalled = true
			// Return 2 pending + 1 success
			return []*dimse.CFindResponse{
				dimse.NewCFindResponseFromRequest(req, 0xFF00, dataset.New()), // Pending
				dimse.NewCFindResponseFromRequest(req, 0xFF00, dataset.New()), // Pending
				dimse.NewCFindResponseFromRequest(req, 0x0000, nil),           // Success
			}, nil
		},
	}

	err := service.handleCFindRequest(ctx, req, handlers)
	if err != nil {
		t.Errorf("handleCFindRequest failed: %v", err)
	}

	if !handlerCalled {
		t.Error("Custom handler was not called")
	}
}

func TestHandleResponse(t *testing.T) {
	service := NewService(nil, nil)
	defer service.Close()

	// Register a pending request
	msgID := uint16(123)
	req := dimse.NewCEchoRequest()
	respCh := make(chan dimse.Response, 1)
	service.registerPendingRequest(msgID, req, respCh)

	// Create a response
	resp := dimse.NewCEchoResponse(msgID, 0x0000)

	// Handle the response
	err := service.handleResponse(resp)
	if err != nil {
		t.Errorf("handleResponse failed: %v", err)
	}

	// Verify response was delivered
	select {
	case receivedResp := <-respCh:
		if echoResp, ok := receivedResp.(*dimse.CEchoResponse); ok {
			if echoResp.MessageIDBeingRespondedTo() != msgID {
				t.Errorf("Expected MessageID %d, got %d", msgID, echoResp.MessageIDBeingRespondedTo())
			}
		} else {
			t.Errorf("Expected *CEchoResponse, got %T", receivedResp)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Response was not delivered")
	}

	// Clean up
	service.unregisterPendingRequest(msgID)
}

func TestHandleResponse_UnknownMessageID(t *testing.T) {
	service := NewService(nil, nil)
	defer service.Close()

	// Create a response for a MessageID that has no pending request
	resp := dimse.NewCEchoResponse(999, 0x0000)

	// Should return error
	err := service.handleResponse(resp)
	if err == nil {
		t.Error("Expected error for unknown MessageID, got nil")
	}
}
