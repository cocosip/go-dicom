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

	responses, err := service.SendCFind(context.Background(), dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()))
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	request := receiveCFindRequest(t, sent)
	time.Sleep(30 * time.Millisecond)
	if err := service.handleResponse(dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New())); err != nil {
		t.Fatalf("pending handleResponse() error = %v", err)
	}
	if response := <-responses; !response.IsPending() {
		t.Fatalf("first response = %#v, want pending C-FIND response", response)
	}

	time.Sleep(30 * time.Millisecond)
	if err := service.handleResponse(dimse.NewCFindResponseFromRequest(request, status.Success, nil)); err != nil {
		t.Fatalf("final handleResponse() error = %v", err)
	}
	response, ok := <-responses
	if !ok {
		t.Fatal("C-FIND response channel closed before final response")
	}
	if response.IsPending() || response.StatusCode() != status.Success.Code {
		t.Fatalf("final response = %#v, want C-FIND success", response)
	}
	if _, ok := <-responses; ok {
		t.Fatal("C-FIND response channel remained open after final response")
	}
}

func TestSendCFindWithErrorReportsRequestTimeout(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	responses, terminalErrors, err := service.SendCFindWithError(
		context.Background(), dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCFindWithError() error = %v", err)
	}

	select {
	case _, ok := <-responses:
		if ok {
			t.Fatal("responses yielded a response without a peer reply")
		}
	case <-time.After(time.Second):
		t.Fatal("responses did not close after request timeout")
	}
	select {
	case terminalErr, ok := <-terminalErrors:
		if !ok {
			t.Fatal("terminal errors closed without the request timeout")
		}
		if !errors.Is(terminalErr, ErrRequestTimeout) {
			t.Fatalf("terminal error = %v, want ErrRequestTimeout", terminalErr)
		}
	case <-time.After(time.Second):
		t.Fatal("request timeout was not reported")
	}
}

func TestSendCMoveWithErrorReportsRequestTimeout(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	responses, terminalErrors, err := service.SendCMoveWithError(
		context.Background(),
		dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST_AE", dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCMoveWithError() error = %v", err)
	}

	assertProgressRequestTimeout(t, responses, terminalErrors)
}

func TestSendCGetWithErrorReportsRequestTimeout(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	go drainSuccessfulSends(service)
	responses, terminalErrors, err := service.SendCGetWithError(
		context.Background(),
		dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, dataset.New()),
	)
	if err != nil {
		t.Fatalf("SendCGetWithError() error = %v", err)
	}

	assertProgressRequestTimeout(t, responses, terminalErrors)
}

func assertProgressRequestTimeout[T any](t *testing.T, responses <-chan T, terminalErrors <-chan error) {
	t.Helper()
	select {
	case _, ok := <-responses:
		if ok {
			t.Fatal("responses yielded a response without a peer reply")
		}
	case <-time.After(time.Second):
		t.Fatal("responses did not close after request timeout")
	}
	select {
	case terminalErr, ok := <-terminalErrors:
		if !ok {
			t.Fatal("terminal errors closed without the request timeout")
		}
		if !errors.Is(terminalErr, ErrRequestTimeout) {
			t.Fatalf("terminal error = %v, want ErrRequestTimeout", terminalErr)
		}
	case <-time.After(time.Second):
		t.Fatal("request timeout was not reported")
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
