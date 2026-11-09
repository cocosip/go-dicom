// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// countingAcceptListener records accept attempts and delegates to a real
// listener so AcceptTimeout produces the same wrapped deadline error as in
// production.
type countingAcceptListener struct {
	serverListener
	mu      sync.Mutex
	accepts int
}

func (l *countingAcceptListener) Accept(ctx context.Context) (net.Conn, error) {
	l.mu.Lock()
	l.accepts++
	l.mu.Unlock()
	return l.serverListener.Accept(ctx)
}

func (l *countingAcceptListener) acceptCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.accepts
}

// TestAcceptLoopIgnoresAcceptTimeoutErrors verifies that an AcceptTimeout
// deadline (a polling timeout, not a listener failure) neither emits an
// accept_failed error metric nor stops the loop.
func TestAcceptLoopIgnoresAcceptTimeoutErrors(t *testing.T) {
	base, err := transport.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	t.Cleanup(func() { _ = base.Close() })
	listener := &countingAcceptListener{serverListener: base}

	recorder := &serverObservationRecorder{}
	srv := New(
		WithAcceptTimeout(5*time.Millisecond),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.metric)),
	)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- srv.acceptLoop(ctx, listener) }()

	deadline := time.Now().Add(time.Second)
	for listener.acceptCount() < 3 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("acceptLoop() error = %v, want context.Canceled", err)
	}
	if got := listener.acceptCount(); got < 3 {
		t.Fatalf("accept count = %d, want the loop to keep accepting after accept timeouts", got)
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	for _, metric := range recorder.metrics {
		if metric.Kind == observability.MetricError {
			t.Fatalf("accept timeout emitted error metric %#v", metric)
		}
	}
}

// failingAcceptListener fails every Accept immediately.
type failingAcceptListener struct {
	mu      sync.Mutex
	accepts int
}

func (l *failingAcceptListener) Accept(context.Context) (net.Conn, error) {
	l.mu.Lock()
	l.accepts++
	l.mu.Unlock()
	return nil, errors.New("listener failed")
}

func (l *failingAcceptListener) Close() error { return nil }

func (l *failingAcceptListener) acceptCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.accepts
}

// TestAcceptLoopBacksOffOnPersistentAcceptErrors verifies that repeated accept
// failures are rate-limited by exponential backoff instead of spinning.
func TestAcceptLoopBacksOffOnPersistentAcceptErrors(t *testing.T) {
	listener := &failingAcceptListener{}
	recorder := &serverObservationRecorder{}
	srv := New(WithMetricsObserver(observability.MetricsObserverFunc(recorder.metric)))
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- srv.acceptLoop(ctx, listener) }()

	time.Sleep(120 * time.Millisecond)
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("acceptLoop() error = %v, want context.Canceled", err)
	}
	// 5ms, 10ms, 20ms, 40ms, 80ms: at most a handful of attempts fit in the
	// window. A busy loop would report thousands.
	if got := listener.acceptCount(); got < 2 || got > 12 {
		t.Fatalf("accept count = %d in 120ms, want exponential backoff", got)
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	errorMetrics := 0
	for _, metric := range recorder.metrics {
		if metric.Kind == observability.MetricError {
			errorMetrics++
		}
	}
	if errorMetrics == 0 {
		t.Fatal("persistent accept errors did not emit any error metric")
	}
	if errorMetrics > listener.acceptCount() {
		t.Fatalf("error metrics = %d exceeds accept attempts = %d", errorMetrics, listener.acceptCount())
	}
}
