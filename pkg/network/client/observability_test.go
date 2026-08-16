// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/observability"
)

type clientObservationRecorder struct {
	mu      sync.Mutex
	events  []observability.Event
	metrics []observability.Metric
	logs    []observability.LogRecord
}

func (r *clientObservationRecorder) event(_ context.Context, event observability.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, event)
}

func (r *clientObservationRecorder) metric(_ context.Context, metric observability.Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.metrics = append(r.metrics, metric)
}

func (r *clientObservationRecorder) log(_ context.Context, record observability.LogRecord) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logs = append(r.logs, record)
}

func TestObservabilityDialFailureBeforeServiceCreation(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	portText := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	if err := listener.Close(); err != nil {
		t.Fatalf("listener Close() error = %v", err)
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		t.Fatalf("Atoi(%q) error = %v", portText, err)
	}

	recorder := &clientObservationRecorder{}
	client := New(
		WithCallingAE("OBS-SCU"),
		WithCalledAE("OBS-SCP"),
		WithConnectTimeout(200*time.Millisecond),
		WithLogger(observability.LoggerFunc(recorder.log)),
		WithEventObserver(observability.EventObserverFunc(recorder.event)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.metric)),
	)
	client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := client.Connect(ctx, "127.0.0.1", port); err == nil {
		t.Fatal("Connect() error = nil, want refused connection")
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	if len(recorder.events) != 2 {
		t.Fatalf("event count = %d, want 2: %#v", len(recorder.events), recorder.events)
	}
	if recorder.events[0].Kind != observability.EventConnectionAttempted ||
		recorder.events[1].Kind != observability.EventConnectionClosed ||
		recorder.events[1].Outcome != observability.OutcomeFailure {
		t.Fatalf("connection events = %#v", recorder.events)
	}
	if recorder.events[0].Association.ConnectionID == 0 ||
		recorder.events[0].Association.ConnectionID != recorder.events[1].Association.ConnectionID {
		t.Fatalf("connection IDs = [%d, %d]", recorder.events[0].Association.ConnectionID, recorder.events[1].Association.ConnectionID)
	}
	if len(recorder.metrics) < 2 {
		t.Fatalf("metric count = %d, want connection and error metrics", len(recorder.metrics))
	}
	if len(recorder.logs) != 2 {
		t.Fatalf("log count = %d, want 2", len(recorder.logs))
	}
}
