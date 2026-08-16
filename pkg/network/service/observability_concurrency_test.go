// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestRequestLifecycleConcurrentTerminalExactlyOnce(t *testing.T) {
	for iteration := 0; iteration < 50; iteration++ {
		recorder := &observationRecorder{}
		service := NewService(&mockConn{}, createTestAssociation(),
			WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		)
		req := dimse.NewCEchoRequest()
		if err := req.SetMessageID(uint16(iteration + 1)); err != nil {
			t.Fatalf("SetMessageID() error = %v", err)
		}
		lifecycle := newRequestLifecycle(service, observability.DirectionOutbound, req)
		lifecycle.sent(context.Background())
		response := dimse.NewCEchoResponseFromRequest(req, status.Success)

		start := make(chan struct{})
		var wg sync.WaitGroup
		for _, finish := range []func(){
			func() { lifecycle.response(context.Background(), response) },
			func() { lifecycle.finishContext(context.Background(), context.DeadlineExceeded) },
			func() { lifecycle.finishContext(context.Background(), context.Canceled) },
			func() { lifecycle.finishError(context.Background(), errors.New("connection closed")) },
		} {
			wg.Add(1)
			go func(f func()) {
				defer wg.Done()
				<-start
				f()
			}(finish)
		}
		close(start)
		wg.Wait()

		events := requestEvents(recorder, dimse.CommandCEchoRQ)
		if len(events) != 2 || events[0].Kind != observability.EventRequestSent || !isTerminalRequestEvent(events[1].Kind) {
			t.Fatalf("iteration %d events = %#v, want sent plus one terminal", iteration, events)
		}
		_ = service.Close()
	}
}

func TestObservabilityCloseFinalizesPendingAndInboundOnce(t *testing.T) {
	recorder := &observationRecorder{}
	service := NewService(&mockConn{}, createTestAssociation(),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
	)

	outbound := dimse.NewCEchoRequest()
	if err := outbound.SetMessageID(11); err != nil {
		t.Fatalf("outbound SetMessageID() error = %v", err)
	}
	pending := service.registerPendingRequest(11, outbound, make(chan dimse.Response, 1))
	pending.lifecycle.sent(context.Background())

	inbound := dimse.NewCEchoRequest()
	if err := inbound.SetMessageID(12); err != nil {
		t.Fatalf("inbound SetMessageID() error = %v", err)
	}
	inboundLifecycle := service.registerInboundRequest(inbound)
	inboundLifecycle.received(context.Background())

	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	for _, messageID := range []uint16{11, 12} {
		events, _, _ := recorder.snapshot()
		var requestEventsForID []observability.Event
		for _, event := range events {
			if event.MessageID == messageID {
				requestEventsForID = append(requestEventsForID, event)
			}
		}
		if len(requestEventsForID) != 2 || !isTerminalRequestEvent(requestEventsForID[1].Kind) {
			t.Errorf("message %d events = %#v, want initial plus one terminal", messageID, requestEventsForID)
		}
	}
}

func TestObservabilityHooksAreReentrantSlowAndPanicIsolated(t *testing.T) {
	var service *Service
	var ready atomic.Bool
	var calls atomic.Int32
	observer := observability.EventObserverFunc(func(context.Context, observability.Event) {
		calls.Add(1)
		if ready.Load() {
			_ = service.GetState()
			time.Sleep(time.Millisecond)
		}
	})
	service = NewService(&mockConn{}, nil,
		WithLogger(observability.LoggerFunc(func(context.Context, observability.LogRecord) {
			panic("logger failure")
		})),
		WithEventObserver(observer),
		WithMetricsObserver(observability.MetricsObserverFunc(func(context.Context, observability.Metric) {
			if ready.Load() {
				_ = service.GetAssociation()
			}
		})),
	)
	ready.Store(true)

	done := make(chan struct{})
	go func() {
		service.emitAssociationObservation(
			context.Background(),
			observability.EventAssociationAccepted,
			observability.DirectionInbound,
			observability.OutcomeSuccess,
			nil,
		)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("reentrant or slow hook deadlocked observation delivery")
	}
	if calls.Load() < 2 {
		t.Fatalf("event calls = %d, want connection and association events", calls.Load())
	}
	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

func isTerminalRequestEvent(kind observability.EventKind) bool {
	switch kind {
	case observability.EventRequestCompleted,
		observability.EventRequestTimedOut,
		observability.EventRequestCancelled,
		observability.EventRequestFailed:
		return true
	default:
		return false
	}
}
