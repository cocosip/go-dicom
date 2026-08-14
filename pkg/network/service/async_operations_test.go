// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

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
