// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func drainSuccessfulSends(service *Service) {
	go func() {
		for {
			select {
			case request := <-service.sendQueue:
				request.resultCh <- nil
			case <-service.closeCh:
				return
			}
		}
	}()
}

func TestSendCEchoLimitsOutstandingRequestsToNegotiatedMaximum(t *testing.T) {
	service := newAsyncOperationsTestService(t, 2)
	sent := consumeCEchoSendQueue(service)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	results := startCEchoRequests(ctx, service, 3)

	first := receiveCEchoRequest(t, sent)
	second := receiveCEchoRequest(t, sent)
	select {
	case third := <-sent:
		t.Fatalf("third request %d was sent before an outstanding request completed", third.MessageID())
	case <-time.After(100 * time.Millisecond):
	}

	completeCEchoRequest(t, service, first)
	third := receiveCEchoRequest(t, sent)
	completeCEchoRequest(t, service, second)
	completeCEchoRequest(t, service, third)
	assertCEchoResults(t, results, 3)
}

func TestSendCEchoTreatsZeroMaximumAsUnlimited(t *testing.T) {
	service := newAsyncOperationsTestService(t, 0)
	sent := consumeCEchoSendQueue(service)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	results := startCEchoRequests(ctx, service, 3)

	requests := make([]*dimse.CEchoRequest, 0, 3)
	for range 3 {
		requests = append(requests, receiveCEchoRequest(t, sent))
	}
	for _, request := range requests {
		completeCEchoRequest(t, service, request)
	}
	assertCEchoResults(t, results, 3)
}

func TestSendCEchoStopsWaitingForSlotWhenContextIsCancelled(t *testing.T) {
	service := newAsyncOperationsTestService(t, 1)
	sent := consumeCEchoSendQueue(service)

	firstCtx, cancelFirst := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFirst()
	firstResult := startCEchoRequests(firstCtx, service, 1)
	first := receiveCEchoRequest(t, sent)

	waitingCtx, cancelWaiting := context.WithCancel(context.Background())
	waitingResult := startCEchoRequests(waitingCtx, service, 1)
	cancelWaiting()

	select {
	case err := <-waitingResult:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("waiting request error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("waiting request did not stop after context cancellation")
	}
	select {
	case request := <-sent:
		t.Fatalf("cancelled request %d was sent", request.MessageID())
	case <-time.After(100 * time.Millisecond):
	}

	completeCEchoRequest(t, service, first)
	assertCEchoResults(t, firstResult, 1)
}

func TestSendCEchoRequestTimeoutDoesNotCloseAssociationAndIgnoresLateResponse(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(25*time.Millisecond))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	sent := consumeCEchoSendQueue(service)

	firstResult := make(chan error, 1)
	go func() {
		_, err := service.SendCEcho(context.Background(), dimse.NewCEchoRequest())
		firstResult <- err
	}()
	first := receiveCEchoRequest(t, sent)

	select {
	case err := <-firstResult:
		if !errors.Is(err, ErrRequestTimeout) {
			t.Fatalf("SendCEcho() error = %v, want ErrRequestTimeout", err)
		}
		var timeoutErr *RequestTimeoutError
		if !errors.As(err, &timeoutErr) || timeoutErr.MessageID != first.MessageID() {
			t.Fatalf("SendCEcho() timeout error = %#v, want message ID %d", err, first.MessageID())
		}
	case <-time.After(time.Second):
		t.Fatal("SendCEcho() did not return its request timeout")
	}
	if service.IsClosed() {
		t.Fatal("service closed after one request timed out")
	}
	if err := service.handleResponse(dimse.NewCEchoResponseFromRequest(first, status.Success)); err != nil {
		t.Fatalf("late handleResponse() error = %v", err)
	}

	secondResult := make(chan error, 1)
	go func() {
		_, err := service.SendCEcho(context.Background(), dimse.NewCEchoRequest())
		secondResult <- err
	}()
	second := receiveCEchoRequest(t, sent)
	completeCEchoRequest(t, service, second)
	if err := <-secondResult; err != nil {
		t.Fatalf("second SendCEcho() error = %v", err)
	}
}

func TestSendCFindPendingResponseResetsRequestTimeout(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(50*time.Millisecond))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	sent := make(chan *dimse.CFindRequest, 1)
	go func() {
		request := <-service.sendQueue
		find, ok := request.message.(*dimse.CFindRequest)
		if ok {
			sent <- find
		}
		request.resultCh <- nil
	}()

	events, err := service.SendCFind(context.Background(), dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()))
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	request := receiveCFindRequest(t, sent)
	time.Sleep(30 * time.Millisecond)
	if err := service.handleResponse(dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New())); err != nil {
		t.Fatalf("pending handleResponse() error = %v", err)
	}
	if event := receiveResponseEvent(t, events); event.Response == nil || !event.Response.IsPending() {
		t.Fatalf("first event = %#v, want pending C-FIND response", event)
	}

	time.Sleep(30 * time.Millisecond)
	if err := service.handleResponse(dimse.NewCFindResponseFromRequest(request, status.Success, nil)); err != nil {
		t.Fatalf("final handleResponse() error = %v", err)
	}
	event, ok := <-events
	if !ok {
		t.Fatal("C-FIND event channel closed before final response")
	}
	assertResponseEvent(t, event)
	if event.Response.IsPending() || event.Response.StatusCode() != status.Success.Code {
		t.Fatalf("final event = %#v, want C-FIND success", event)
	}
	if _, ok := <-events; ok {
		t.Fatal("C-FIND event channel remained open after final response")
	}
}

func TestSendCFindReportsRequestTimeoutAsSingleTerminalEvent(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	events, err := service.SendCFind(
		context.Background(), dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}

	assertProgressRequestTimeout(t, events)
}

func TestSendCMoveReportsRequestTimeoutAsSingleTerminalEvent(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	events, err := service.SendCMove(
		context.Background(),
		dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, testMoveDestinationAE, dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCMove() error = %v", err)
	}

	assertProgressRequestTimeout(t, events)
}

func TestSendCGetReportsRequestTimeoutAsSingleTerminalEvent(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	events, err := service.SendCGet(
		context.Background(),
		dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCGet() error = %v", err)
	}

	assertProgressRequestTimeout(t, events)
}

func TestSendCFindContextCancellationEmitsOneTerminalEvent(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	drainSuccessfulSends(service)

	ctx, cancel := context.WithCancel(context.Background())
	events, err := service.SendCFind(ctx, dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()))
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	cancel()
	assertSingleTerminalEvent(t, events, context.Canceled)
}

func TestSendCFindServiceCloseEmitsOneTerminalEvent(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	drainSuccessfulSends(service)

	events, err := service.SendCFind(context.Background(), dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()))
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	assertSingleTerminalEvent(t, events, ErrServiceClosed)
}

func TestSendCFindDoesNotSilentlyCloseWithoutFinalResponseOrError(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	drainSuccessfulSends(service)

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	events, err := service.SendCFind(context.Background(), request)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	service.pendingRequestsMu.RLock()
	pending := service.pendingRequests[request.MessageID()]
	service.pendingRequestsMu.RUnlock()
	if pending == nil {
		t.Fatal("pending request was not registered")
	}
	close(pending.responseCh)

	event := receiveResponseEvent(t, events)
	if event.Response != nil || event.Err == nil {
		t.Fatalf("terminal event = %#v, want error-only event", event)
	}
	if _, ok := <-events; ok {
		t.Fatal("event channel remained open after terminal error")
	}
}

type testProgressResponse interface {
	dimse.Response
	*dimse.CFindResponse | *dimse.CMoveResponse | *dimse.CGetResponse
}

func assertProgressRequestTimeout[T testProgressResponse](t *testing.T, events <-chan ResponseEvent[T]) {
	t.Helper()
	assertSingleTerminalEvent(t, events, ErrRequestTimeout)
}

func assertSingleTerminalEvent[T testProgressResponse](t *testing.T, events <-chan ResponseEvent[T], want error) {
	t.Helper()
	event := receiveResponseEvent(t, events)
	if event.Response != nil || !errors.Is(event.Err, want) {
		t.Fatalf("terminal event = %#v, want error-only event matching %v", event, want)
	}
	if _, ok := <-events; ok {
		t.Fatal("event channel yielded more than one terminal event")
	}
}

func receiveResponseEvent[T testProgressResponse](t *testing.T, events <-chan ResponseEvent[T]) ResponseEvent[T] {
	t.Helper()
	select {
	case event, ok := <-events:
		if !ok {
			t.Fatal("event channel closed without a final response or terminal error")
		}
		if (event.Response == nil) == (event.Err == nil) {
			t.Fatalf("event = %#v, want exactly one of Response or Err", event)
		}
		return event
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for response event")
	}
	var zero ResponseEvent[T]
	return zero
}

func assertResponseEvent[T testProgressResponse](t *testing.T, event ResponseEvent[T]) {
	t.Helper()
	if event.Response == nil || event.Err != nil {
		t.Fatalf("event = %#v, want response-only event", event)
	}
}

func newAsyncOperationsTestService(t *testing.T, maxInvoked uint16) *Service {
	t.Helper()
	assoc := createTestAssociation()
	assoc.AsynchronousOperations = &association.AsynchronousOperationsWindow{
		MaxInvokedOperations:   maxInvoked,
		MaxPerformedOperations: 1,
	}
	service := NewService(nil, assoc)
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState failed: %v", err)
	}
	return service
}

func consumeCEchoSendQueue(service *Service) <-chan *dimse.CEchoRequest {
	sent := make(chan *dimse.CEchoRequest, 10)
	go func() {
		for {
			select {
			case request := <-service.sendQueue:
				echo, ok := request.message.(*dimse.CEchoRequest)
				if ok {
					sent <- echo
				}
				request.resultCh <- nil
			case <-service.closeCh:
				return
			}
		}
	}()
	return sent
}

func startCEchoRequests(ctx context.Context, service *Service, count int) <-chan error {
	results := make(chan error, count)
	for range count {
		go func() {
			_, err := service.SendCEcho(ctx, dimse.NewCEchoRequest())
			results <- err
		}()
	}
	return results
}

func receiveCEchoRequest(t *testing.T, sent <-chan *dimse.CEchoRequest) *dimse.CEchoRequest {
	t.Helper()
	select {
	case request := <-sent:
		return request
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-ECHO request")
		return nil
	}
}

func receiveCFindRequest(t *testing.T, sent <-chan *dimse.CFindRequest) *dimse.CFindRequest {
	t.Helper()
	select {
	case request := <-sent:
		return request
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for C-FIND request")
		return nil
	}
}

func completeCEchoRequest(t *testing.T, service *Service, request *dimse.CEchoRequest) {
	t.Helper()
	response := dimse.NewCEchoResponseFromRequest(request, status.Success)
	if err := service.handleResponse(response); err != nil {
		t.Fatalf("handleResponse failed: %v", err)
	}
}

func assertCEchoResults(t *testing.T, results <-chan error, count int) {
	t.Helper()
	for range count {
		select {
		case err := <-results:
			if err != nil {
				t.Fatalf("SendCEcho failed: %v", err)
			}
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for SendCEcho result")
		}
	}
}
