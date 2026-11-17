// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
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
