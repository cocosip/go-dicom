// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// TestHandleResponseDoesNotBlockOnSlowConsumer verifies that delivering more
// responses than both the per-request channel and the application event channel
// can hold does not block the receive loop. A stalled consumer must only grow
// the per-request queue (as fo-dicom's unbounded request channel does), and the
// queued responses must still arrive in order once the application reads them.
func TestHandleResponseDoesNotBlockOnSlowConsumer(t *testing.T) {
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

	// 30 responses exceed the response channel (10) and the event channel (10),
	// and the application below has not read anything yet.
	const responseCount = 30
	responses := make([]*dimse.CFindResponse, 0, responseCount)
	for range responseCount {
		responses = append(responses, dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New()))
	}

	pushDone := make(chan error, 1)
	go func() {
		for _, resp := range responses {
			if err := service.handleResponse(resp); err != nil {
				pushDone <- err
				return
			}
		}
		pushDone <- nil
	}()

	select {
	case err := <-pushDone:
		if err != nil {
			t.Fatalf("handleResponse() error = %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("handleResponse() blocked the receive loop on a slow consumer")
	}

	for i, want := range responses {
		event := receiveResponseEvent(t, events)
		if event.Response != want {
			t.Fatalf("event %d = %p, want %p (response order changed)", i, event.Response, want)
		}
	}

	if err := service.handleResponse(dimse.NewCFindResponseFromRequest(request, status.Success, nil)); err != nil {
		t.Fatalf("final handleResponse() error = %v", err)
	}
	final := receiveResponseEvent(t, events)
	if final.Response == nil || final.Response.IsPending() {
		t.Fatalf("final event = %#v, want a final response", final)
	}
	if _, ok := <-events; ok {
		t.Fatal("event channel remained open after the final response")
	}
}

// TestResponseQueueOverflowFailsRequestOnly verifies that a stalled consumer and
// a peer streaming more pending responses than maxQueuedResponses fail only that
// request (ErrResponseQueueOverflow) while the association stays usable.
func TestResponseQueueOverflowFailsRequestOnly(t *testing.T) {
	service := NewService(nil, createTestAssociation())
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	// Answer C-ECHOs so the post-overflow health check can complete; other sends
	// are just acknowledged.
	go func() {
		for {
			select {
			case send := <-service.sendQueue:
				echo, isEcho := send.message.(*dimse.CEchoRequest)
				send.resultCh <- nil
				if isEcho {
					_ = service.handleResponse(dimse.NewCEchoResponse(echo.MessageID(), status.Success))
				}
			case <-service.closeCh:
				return
			}
		}
	}()

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	events, err := service.SendCFind(context.Background(), request)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}

	// Do not read events: the per-request queue fills up and must fail the
	// request instead of growing without bound. The slack on top of the bound
	// covers the responses already buffered in responseCh/eventCh.
	const overflowSlack = 128
	for i := 0; i < maxQueuedResponses+overflowSlack; i++ {
		resp := dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New())
		if err := service.handleResponse(resp); err != nil {
			t.Fatalf("handleResponse(%d) error = %v, want nil (overflow fails only the request)", i, err)
		}
	}
	service.pendingRequestsMu.Lock()
	_, stillPending := service.pendingRequests[request.MessageID()]
	service.pendingRequestsMu.Unlock()
	if stillPending {
		t.Fatal("request is still pending after exceeding the response queue bound")
	}

	var terminal error
	for terminal == nil {
		terminal = receiveResponseEvent(t, events).Err
	}
	if !errors.Is(terminal, ErrResponseQueueOverflow) {
		t.Fatalf("terminal event error = %v, want ErrResponseQueueOverflow", terminal)
	}
	if _, ok := <-events; ok {
		t.Fatal("event channel remained open after the overflow error")
	}

	// The association itself must remain usable.
	resp, err := service.SendCEcho(context.Background(), dimse.NewCEchoRequest())
	if err != nil {
		t.Fatalf("SendCEcho() after overflow error = %v", err)
	}
	if !resp.Status().IsSuccess() {
		t.Fatalf("SendCEcho() status = %s, want Success", resp.Status())
	}
	if state := service.GetState(); state != StateAssociationAccepted {
		t.Fatalf("state = %s, want %s (overflow must not abort the association)", state, StateAssociationAccepted)
	}
}
