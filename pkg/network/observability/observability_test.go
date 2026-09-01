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
	operationIDs := make(chan uint64, count)

	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			connectionIDs <- NewConnectionID()
			associationIDs <- NewAssociationID()
			operationIDs <- NewOperationID()
		}()
	}
	wg.Wait()
	close(connectionIDs)
	close(associationIDs)
	close(operationIDs)

	assertUniqueConnections(t, connectionIDs, count)
	assertUniqueAssociations(t, associationIDs, count)
	assertUniqueOperations(t, operationIDs, count)
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

func assertUniqueOperations(t *testing.T, ids <-chan uint64, want int) {
	t.Helper()
	seen := make(map[uint64]struct{}, want)
	for id := range ids {
		if id == 0 {
			t.Fatal("NewOperationID() returned zero")
		}
		if _, exists := seen[id]; exists {
			t.Fatalf("duplicate operation ID %d", id)
		}
		seen[id] = struct{}{}
	}
	if len(seen) != want {
		t.Fatalf("operation ID count = %d, want %d", len(seen), want)
	}
}

func TestFunctionAdaptersDeliverRecords(t *testing.T) {
	ctx := context.Background()
	event := Event{Kind: EventRequestSent, MessageID: 17}
	metric := Metric{Kind: MetricDIMSE, Value: 1}

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

	if eventCalls.Load() != 1 || metricCalls.Load() != 1 {
		t.Fatalf("calls = event:%d metric:%d, want 1 each",
			eventCalls.Load(), metricCalls.Load())
	}
}

func TestNoOpAndNilHooksRemainDisabled(t *testing.T) {
	ctx := context.Background()
	if EmitEvent(ctx, nil, Event{}) {
		t.Fatal("EmitEvent(nil) = true")
	}
	if EmitMetric(ctx, nil, Metric{}) {
		t.Fatal("EmitMetric(nil) = true")
	}
}

func TestEmitHelpersRecoverHookPanics(t *testing.T) {
	ctx := context.Background()
	observer := EventObserverFunc(func(context.Context, Event) { panic("event") })
	if EmitEvent(ctx, observer, Event{}) {
		t.Fatal("EmitEvent() = true after panic")
	}
	metrics := MetricsObserverFunc(func(context.Context, Metric) { panic("metric") })
	if EmitMetric(ctx, metrics, Metric{}) {
		t.Fatal("EmitMetric() = true after panic")
	}
}
