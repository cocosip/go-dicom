// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
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

func setupCapturingTestService(t *testing.T) (*Service, context.Context, context.CancelFunc, <-chan dimse.Message) {
	t.Helper()

	service := NewService(nil, nil)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	sent := make(chan dimse.Message, 1)
	go func() {
		select {
		case req := <-service.sendQueue:
			sent <- req.message
			req.resultCh <- nil
		case <-service.closeCh:
		case <-ctx.Done():
		}
	}()

	return service, ctx, cancel, sent
}

func TestSetGetHandlers(t *testing.T) {
	service := NewService(nil, nil)
	defer func() { _ = service.Close() }()

	// Initially should be nil
	if h := service.GetHandlers(); h != nil {
		t.Error("Expected nil handlers initially")
	}

	// Set handlers
	handlers := &Handlers{
		CEchoHandler: func(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
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
	defer func() { _ = service.Close() }()
	defer cancel()

	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// No custom handlers - should use default
	err := service.handleCEchoRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCEchoRequest failed: %v", err)
	}
}

func TestHandleRequest_DispatchErrorDoesNotDeadlockShutdown(t *testing.T) {
	service := NewService(nil, nil)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	command := dimse.CreateCommandDataset(uint16(dimse.CommandNGetRQ), 1)
	req := &unsupportedServiceRequest{BaseRequest: dimse.NewBaseRequest(command, nil)}

	if err := service.handleRequest(context.Background(), req); err != nil {
		t.Fatalf("handleRequest failed: %v", err)
	}

	select {
	case <-service.shutdownCh:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("dispatch error shutdown deadlocked waiting for the failing request handler")
	}
}

type unsupportedServiceRequest struct {
	*dimse.BaseRequest
}

func TestHandleCEchoRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Custom handler that returns success
	handlerCalled := false
	handlers := &Handlers{
		CEchoHandler: func(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			handlerCalled = true
			return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
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
	defer func() { _ = service.Close() }()
	defer cancel()

	// Create dataset with required tags
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})) // CT Image Storage
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE request: %v", err)
	}
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// No custom handlers - should use default
	err = service.handleCStoreRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCStoreRequest failed: %v", err)
	}
}

func TestUnhandledDIMSERequestsReturnFailure(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))
	storeRequest, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest failed: %v", err)
	}

	const managedSOPClassUID = "1.2.840.10008.5.1.1.1"
	const managedSOPInstanceUID = "2.25.400"
	tests := []struct {
		name       string
		wantStatus uint16
		handle     func(*Service, context.Context) error
	}{
		{"C-STORE", status.CStoreErrorCannotUnderstand.Code, func(s *Service, ctx context.Context) error {
			return s.handleCStoreRequest(ctx, storeRequest, nil)
		}},
		{"C-FIND", status.CFindFailedUnableToProcess.Code, func(s *Service, ctx context.Context) error {
			return s.handleCFindRequest(ctx, dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()), nil)
		}},
		{"C-MOVE", status.CMoveFailedUnableToProcess.Code, func(s *Service, ctx context.Context) error {
			return s.handleCMoveRequest(ctx, dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST", dataset.New()), nil)
		}},
		{"C-GET", status.CGetFailedUnableToProcess.Code, func(s *Service, ctx context.Context) error {
			return s.handleCGetRequest(ctx, dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, dataset.New()), nil)
		}},
		{"N-EVENT-REPORT", status.NEventReportFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNEventReportRequest(ctx, dimse.NewNEventReportRequest(managedSOPClassUID, managedSOPInstanceUID, 1, nil), nil)
		}},
		{"N-GET", status.NGetFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNGetRequest(ctx, dimse.NewNGetRequest(managedSOPClassUID, managedSOPInstanceUID, nil), nil)
		}},
		{testNSetOperation, status.NSetFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNSetRequest(ctx, dimse.NewNSetRequest(managedSOPClassUID, managedSOPInstanceUID, nil), nil)
		}},
		{testNActionOperation, status.NActionFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNActionRequest(ctx, dimse.NewNActionRequest(managedSOPClassUID, managedSOPInstanceUID, 1, nil), nil)
		}},
		{"N-CREATE", status.NCreateFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNCreateRequest(ctx, dimse.NewNCreateRequest(managedSOPClassUID, managedSOPInstanceUID, nil), nil)
		}},
		{"N-DELETE", status.NDeleteFailureProcessingFailure.Code, func(s *Service, ctx context.Context) error {
			return s.handleNDeleteRequest(ctx, dimse.NewNDeleteRequest(managedSOPClassUID, managedSOPInstanceUID), nil)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, ctx, cancel, sent := setupCapturingTestService(t)
			defer func() { _ = service.Close() }()
			defer cancel()

			if err := tt.handle(service, ctx); err != nil {
				t.Fatalf("handler failed: %v", err)
			}
			response, ok := (<-sent).(dimse.Response)
			if !ok {
				t.Fatalf("sent message is not a DIMSE response")
			}
			if response.Status().Code != tt.wantStatus || !response.Status().IsFailure() {
				t.Fatalf("status = %v, want failure 0x%04X", response.Status(), tt.wantStatus)
			}
		})
	}
}

func TestHandleCStoreRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	// Create dataset with required tags
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})) // CT Image Storage
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testSOPInstanceUID}))

	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE request: %v", err)
	}
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Custom handler
	handlerCalled := false
	handlers := &Handlers{
		CStoreHandler: func(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
			handlerCalled = true
			return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
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
	defer func() { _ = service.Close() }()
	defer cancel()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// No custom handlers - should use default (returns success with no results)
	err := service.handleCFindRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCFindRequest failed: %v", err)
	}
}

func TestHandleCFindRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Custom handler that returns multiple responses
	handlerCalled := false
	handlers := &Handlers{
		CFindHandler: func(_ context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			handlerCalled = true
			// Return 2 pending + 1 success
			return []*dimse.CFindResponse{
				dimse.NewCFindResponseFromRequest(req, status.CFindPending, dataset.New()), // Pending
				dimse.NewCFindResponseFromRequest(req, status.CFindPending, dataset.New()), // Pending
				dimse.NewCFindResponseFromRequest(req, status.Success, nil),                // Success
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

func TestHandleCFindRequest_StreamHandlerSendsPendingBeforeHandlerReturns(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(91); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}

	firstSent := make(chan struct{})
	releaseHandler := make(chan struct{})
	handlerDone := make(chan struct{})
	handlers := &Handlers{
		CFindStreamHandler: func(_ context.Context, operation CFindOperation) error {
			if err := operation.SendPending(dataset.New()); err != nil {
				return err
			}
			close(firstSent)
			<-releaseHandler
			close(handlerDone)
			return operation.SendFinal(status.Success)
		},
	}

	done := make(chan error, 1)
	go func() { done <- service.handleCFindRequest(ctx, request, handlers) }()

	select {
	case <-firstSent:
	case <-time.After(time.Second):
		t.Fatal("stream handler did not send its first pending response")
	}

	select {
	case <-handlerDone:
		t.Fatal("handler returned before the test released it")
	default:
	}

	close(releaseHandler)
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("handleCFindRequest() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for stream handler")
	}
}

func TestCFindOperationRejectsResponsesAfterFinal(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(92); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}

	var responses []*dimse.CFindResponse
	operation := newCFindOperation(request, func(response *dimse.CFindResponse) error {
		responses = append(responses, response)
		return nil
	})
	if err := operation.SendFinal(status.Success); err != nil {
		t.Fatalf("SendFinal() error = %v", err)
	}
	if err := operation.SendPending(dataset.New()); err == nil {
		t.Fatal("SendPending() after final error = nil")
	}
	if err := operation.SendFinal(status.Success); err == nil {
		t.Fatal("second SendFinal() error = nil")
	}
	if len(responses) != 1 || responses[0].IsPending() {
		t.Fatalf("responses = %#v, want one final response", responses)
	}
}

func TestCFindOperationRejectsPendingFinalStatus(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	operation := newCFindOperation(request, func(*dimse.CFindResponse) error { return nil })

	if err := operation.SendFinal(status.CFindPending); err == nil {
		t.Fatal("SendFinal(Pending) error = nil")
	}
	if err := operation.SendFinal(status.Success); err != nil {
		t.Fatalf("SendFinal(Success) error = %v", err)
	}
}

func TestServiceCloseReportsCFindHandlerShutdownTimeout(t *testing.T) {
	handlerStarted := make(chan struct{})
	releaseHandler := make(chan struct{})
	service := NewService(nil, nil,
		WithHandlerShutdownTimeout(20*time.Millisecond),
		WithCFindStreamHandler(func(context.Context, CFindOperation) error {
			close(handlerStarted)
			<-releaseHandler // Intentionally ignores the cancellation context.
			return nil
		}),
	)
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(93); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	if err := service.handleRequest(context.Background(), request); err != nil {
		t.Fatalf("handleRequest() error = %v", err)
	}
	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("stream handler did not start")
	}

	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !errors.Is(service.CloseError(), ErrHandlerShutdownTimeout) {
		t.Fatalf("CloseError() = %v, want ErrHandlerShutdownTimeout", service.CloseError())
	}
	close(releaseHandler)
}

func TestCFindStreamHandlerCancelBeforeFirstPendingSendsSingleCancelFinal(t *testing.T) {
	service, ctx, cancel, sent := setupCapturingTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(94); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	handlerStarted := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- service.handleCFindRequest(ctx, request, &Handlers{
			CFindStreamHandler: func(ctx context.Context, _ CFindOperation) error {
				close(handlerStarted)
				<-ctx.Done()
				return ctx.Err()
			},
		})
	}()
	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("stream handler did not start")
	}
	if err := service.handleCCancelRequest(dimse.NewCCancelRequest(request.MessageID())); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("handleCFindRequest() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("stream handler did not finish after C-CANCEL")
	}
	select {
	case message := <-sent:
		response, ok := message.(*dimse.CFindResponse)
		if !ok || response.StatusCode() != status.Cancel.Code {
			t.Fatalf("response = %T %#v, want C-FIND Cancel final", message, message)
		}
	case <-time.After(time.Second):
		t.Fatal("C-CANCEL did not produce a final response")
	}
	select {
	case duplicate := <-sent:
		t.Fatalf("C-CANCEL produced duplicate final response: %T", duplicate)
	case <-time.After(50 * time.Millisecond):
	}
}

func TestCFindStreamHandlerCancelUnblocksPendingSendWhenQueueIsFull(t *testing.T) {
	service := NewService(nil, nil, WithSendQueueSize(0))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := request.SetMessageID(95); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	handlerStarted := make(chan struct{})
	pendingReturned := make(chan error, 1)
	handlerDone := make(chan error, 1)
	go func() {
		handlerDone <- service.handleCFindRequest(context.Background(), request, &Handlers{
			CFindStreamHandler: func(_ context.Context, operation CFindOperation) error {
				close(handlerStarted)
				err := operation.SendPending(dataset.New())
				pendingReturned <- err
				return err
			},
		})
	}()
	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("stream handler did not start")
	}

	if err := service.handleCCancelRequest(dimse.NewCCancelRequest(request.MessageID())); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}
	select {
	case err := <-pendingReturned:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("SendPending() error = %v, want context.Canceled", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("C-CANCEL did not unblock SendPending while the send queue was full")
	}
	select {
	case send := <-service.sendQueue:
		response, ok := send.message.(*dimse.CFindResponse)
		if !ok || response.StatusCode() != status.Cancel.Code {
			t.Fatalf("final message = %T %#v, want C-FIND Cancel final", send.message, send.message)
		}
		send.resultCh <- nil
	case <-time.After(time.Second):
		t.Fatal("C-CANCEL did not enqueue a final response")
	}
	select {
	case <-handlerDone:
	case <-time.After(time.Second):
		t.Fatal("C-FIND handler did not return after C-CANCEL")
	}
}

func TestServiceStartRejectsConflictingCFindHandlers(t *testing.T) {
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	service := NewService(client, nil, WithHandlers(&Handlers{
		CFindHandler: func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			return nil, nil
		},
		CFindStreamHandler: func(context.Context, CFindOperation) error { return nil },
	}))
	defer func() { _ = service.Close() }()

	if err := service.Start(); !errors.Is(err, ErrCFindHandlerConflict) {
		t.Fatalf("Start() error = %v, want ErrCFindHandlerConflict", err)
	}
}

func TestHandleCCancelRequest_CancelsActiveCFindHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer func() { _ = service.Close() }()
	defer cancel()

	query := dataset.New()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
	if err := req.SetMessageID(77); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	handlerStarted := make(chan struct{})
	handlerCanceled := make(chan struct{})
	handlers := &Handlers{
		CFindHandler: func(ctx context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			close(handlerStarted)
			<-ctx.Done()
			close(handlerCanceled)
			return []*dimse.CFindResponse{
				dimse.NewCFindResponseFromRequest(req, status.Cancel, nil),
			}, nil
		},
	}

	done := make(chan error, 1)
	go func() {
		done <- service.handleCFindRequest(ctx, req, handlers)
	}()

	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-FIND handler to start")
	}

	cancelReq := dimse.NewCCancelRequest(req.MessageID())
	if err := service.handleCCancelRequest(cancelReq); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}

	select {
	case <-handlerCanceled:
	case <-time.After(time.Second):
		t.Fatal("C-FIND handler context was not canceled")
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("handleCFindRequest() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-FIND handler to finish")
	}
}

func TestHandleResponse(t *testing.T) {
	service := NewService(nil, nil)
	defer func() { _ = service.Close() }()

	// Register a pending request
	msgID := uint16(123)
	req := dimse.NewCEchoRequest()
	respCh := make(chan dimse.Response, 1)
	service.registerPendingRequest(msgID, req, respCh)

	// Create a response
	resp := dimse.NewCEchoResponse(msgID, status.Success)

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
	defer func() { _ = service.Close() }()

	// Create a response for a MessageID that has no pending request
	resp := dimse.NewCEchoResponse(999, status.Success)

	err := service.handleResponse(resp)
	if err != nil {
		t.Errorf("handleResponse() error = %v, want nil for late/unknown response", err)
	}
}

func TestHandleCMoveRequest_DefaultHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, testMoveDestinationAE, identifier)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Call with nil handlers - should use default
	err := service.handleCMoveRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCMoveRequest with default handler failed: %v", err)
	}
}

func TestHandleCMoveRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, testMoveDestinationAE, identifier)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Custom handler using CMoveOperation
	handlerCalled := false
	handlers := &Handlers{
		CMoveHandler: func(_ context.Context, op CMoveOperation) error {
			handlerCalled = true
			// Send 2 pending updates then success
			_ = op.SendPending(10, 0, 0, 0)
			_ = op.SendPending(0, 10, 0, 0)
			return op.SendSuccess()
		},
	}

	err := service.handleCMoveRequest(ctx, req, handlers)
	if err != nil {
		t.Errorf("handleCMoveRequest failed: %v", err)
	}

	if !handlerCalled {
		t.Error("Custom handler was not called")
	}
}

func TestHandleCCancelRequest_CancelsActiveCMoveHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, testMoveDestinationAE, identifier)
	if err := req.SetMessageID(88); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	handlerStarted := make(chan struct{})
	handlerCanceled := make(chan struct{})
	handlers := &Handlers{
		CMoveHandler: func(ctx context.Context, op CMoveOperation) error {
			close(handlerStarted)
			<-ctx.Done()
			close(handlerCanceled)
			return op.SendFailure(status.Cancel)
		},
	}

	done := make(chan error, 1)
	go func() {
		done <- service.handleCMoveRequest(ctx, req, handlers)
	}()

	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-MOVE handler to start")
	}

	cancelReq := dimse.NewCCancelRequest(req.MessageID())
	if err := service.handleCCancelRequest(cancelReq); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}

	select {
	case <-handlerCanceled:
	case <-time.After(time.Second):
		t.Fatal("C-MOVE handler context was not canceled")
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("handleCMoveRequest() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-MOVE handler to finish")
	}
}

func TestHandleCGetRequest_DefaultHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, identifier)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Call with nil handlers - should use default
	err := service.handleCGetRequest(ctx, req, nil)
	if err != nil {
		t.Errorf("handleCGetRequest with default handler failed: %v", err)
	}
}

func TestHandleCGetRequest_CustomHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, identifier)
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Custom handler using CGetOperation
	handlerCalled := false
	handlers := &Handlers{
		CGetHandler: func(_ context.Context, op CGetOperation) error {
			handlerCalled = true
			// Send 2 pending updates then success (no real C-STORE in unit test)
			_ = op.SendPending(5, 0, 0, 0)
			_ = op.SendPending(0, 5, 0, 0)
			return op.SendSuccess()
		},
	}

	err := service.handleCGetRequest(ctx, req, handlers)
	if err != nil {
		t.Errorf("handleCGetRequest failed: %v", err)
	}

	if !handlerCalled {
		t.Error("Custom handler was not called")
	}
}

func TestHandleCCancelRequest_CancelsActiveCGetHandler(t *testing.T) {
	service, ctx, cancel := setupTestService(t)
	defer cancel()
	defer func() { _ = service.Close() }()

	identifier := dataset.New()
	_ = identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	req := dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, identifier)
	if err := req.SetMessageID(99); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	handlerStarted := make(chan struct{})
	handlerCanceled := make(chan struct{})
	handlers := &Handlers{
		CGetHandler: func(ctx context.Context, op CGetOperation) error {
			close(handlerStarted)
			<-ctx.Done()
			close(handlerCanceled)
			return op.SendFailure(status.Cancel)
		},
	}

	done := make(chan error, 1)
	go func() {
		done <- service.handleCGetRequest(ctx, req, handlers)
	}()

	select {
	case <-handlerStarted:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-GET handler to start")
	}

	cancelReq := dimse.NewCCancelRequest(req.MessageID())
	if err := service.handleCCancelRequest(cancelReq); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}

	select {
	case <-handlerCanceled:
	case <-time.After(time.Second):
		t.Fatal("C-GET handler context was not canceled")
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("handleCGetRequest() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-GET handler to finish")
	}
}
