// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package observability

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

func TestIDsAreUniqueUnderConcurrency(t *testing.T) {
	const count = 256
	connectionIDs := make(chan ConnectionID, count)
	associationIDs := make(chan AssociationID, count)

	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			connectionIDs <- NewConnectionID()
			associationIDs <- NewAssociationID()
		}()
	}
	wg.Wait()
	close(connectionIDs)
	close(associationIDs)

	assertUniqueConnections(t, connectionIDs, count)
	assertUniqueAssociations(t, associationIDs, count)
}

func assertUniqueConnections(t *testing.T, ids <-chan ConnectionID, want int) {
	t.Helper()
	seen := make(map[ConnectionID]struct{}, want)
	for id := range ids {
		if id == 0 {
			t.Fatal("NewConnectionID() returned zero")
		}
		if _, exists := seen[id]; exists {
			t.Fatalf("duplicate connection ID %d", id)
		}
		seen[id] = struct{}{}
	}
	if len(seen) != want {
		t.Fatalf("connection ID count = %d, want %d", len(seen), want)
	}
}

func assertUniqueAssociations(t *testing.T, ids <-chan AssociationID, want int) {
	t.Helper()
	seen := make(map[AssociationID]struct{}, want)
	for id := range ids {
		if id == 0 {
			t.Fatal("NewAssociationID() returned zero")
		}
		if _, exists := seen[id]; exists {
			t.Fatalf("duplicate association ID %d", id)
		}
		seen[id] = struct{}{}
	}
	if len(seen) != want {
		t.Fatalf("association ID count = %d, want %d", len(seen), want)
	}
}

func TestFunctionAdaptersDeliverRecords(t *testing.T) {
	ctx := context.Background()
	event := Event{Kind: EventRequestSent, MessageID: 17}
	metric := Metric{Kind: MetricDIMSE, Value: 1}

	var logCalls atomic.Int32
	logger := LoggerFunc(func(_ context.Context, record LogRecord) {
		logCalls.Add(1)
		if record.Event.MessageID != 17 {
			t.Errorf("log MessageID = %d, want 17", record.Event.MessageID)
		}
	})
	if !logger.Enabled(LevelInfo) {
		t.Fatal("non-nil LoggerFunc is disabled")
	}
	if !EmitLog(ctx, logger, LogRecord{Level: LevelInfo, Event: event}) {
		t.Fatal("EmitLog() = false, want true")
	}

	var eventCalls atomic.Int32
	observer := EventObserverFunc(func(_ context.Context, got Event) {
		eventCalls.Add(1)
		if got.Kind != EventRequestSent {
			t.Errorf("event kind = %q, want %q", got.Kind, EventRequestSent)
		}
	})
	if !EmitEvent(ctx, observer, event) {
		t.Fatal("EmitEvent() = false, want true")
	}

	var metricCalls atomic.Int32
	metrics := MetricsObserverFunc(func(_ context.Context, got Metric) {
		metricCalls.Add(1)
		if got != metric {
			t.Errorf("metric = %#v, want %#v", got, metric)
		}
	})
	if !EmitMetric(ctx, metrics, metric) {
		t.Fatal("EmitMetric() = false, want true")
	}

	if logCalls.Load() != 1 || eventCalls.Load() != 1 || metricCalls.Load() != 1 {
		t.Fatalf("calls = log:%d event:%d metric:%d, want 1 each",
			logCalls.Load(), eventCalls.Load(), metricCalls.Load())
	}
}

func TestNoOpAndNilHooksRemainDisabled(t *testing.T) {
	ctx := context.Background()
	record := LogRecord{Level: LevelError}

	if (NopLogger{}).Enabled(LevelError) {
		t.Fatal("NopLogger.Enabled() = true")
	}
	if EmitLog(ctx, NopLogger{}, record) {
		t.Fatal("EmitLog(NopLogger) = true")
	}
	if EmitLog(ctx, nil, record) {
		t.Fatal("EmitLog(nil) = true")
	}
	if EmitEvent(ctx, nil, Event{}) {
		t.Fatal("EmitEvent(nil) = true")
	}
	if EmitMetric(ctx, nil, Metric{}) {
		t.Fatal("EmitMetric(nil) = true")
	}
}

func TestEmitHelpersRecoverHookPanics(t *testing.T) {
	ctx := context.Background()

	logger := LoggerFunc(func(context.Context, LogRecord) { panic("logger") })
	if EmitLog(ctx, logger, LogRecord{Level: LevelInfo}) {
		t.Fatal("EmitLog() = true after panic")
	}
	observer := EventObserverFunc(func(context.Context, Event) { panic("event") })
	if EmitEvent(ctx, observer, Event{}) {
		t.Fatal("EmitEvent() = true after panic")
	}
	metrics := MetricsObserverFunc(func(context.Context, Metric) { panic("metric") })
	if EmitMetric(ctx, metrics, Metric{}) {
		t.Fatal("EmitMetric() = true after panic")
	}
}

type enabledPanicLogger struct{}

func (enabledPanicLogger) Enabled(Level) bool { panic("enabled") }
func (enabledPanicLogger) Log(context.Context, LogRecord) {
	panic("Log must not be called after Enabled panic")
}

func TestEmitLogRecoversEnabledPanic(t *testing.T) {
	if EmitLog(context.Background(), enabledPanicLogger{}, LogRecord{Level: LevelInfo}) {
		t.Fatal("EmitLog() = true after Enabled panic")
	}
}
