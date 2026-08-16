// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

type observationRecorder struct {
	mu      sync.Mutex
	events  []observability.Event
	metrics []observability.Metric
	logs    []observability.LogRecord
}

func (r *observationRecorder) observeEvent(_ context.Context, event observability.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, event)
}

func (r *observationRecorder) observeMetric(_ context.Context, metric observability.Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.metrics = append(r.metrics, metric)
}

func (r *observationRecorder) log(_ context.Context, record observability.LogRecord) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logs = append(r.logs, record)
}

func (r *observationRecorder) snapshot() ([]observability.Event, []observability.Metric, []observability.LogRecord) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]observability.Event(nil), r.events...),
		append([]observability.Metric(nil), r.metrics...),
		append([]observability.LogRecord(nil), r.logs...)
}

func (r *observationRecorder) waitForSnapshot(
	t *testing.T,
	description string,
	predicate func([]observability.Event, []observability.Metric, []observability.LogRecord) bool,
) ([]observability.Event, []observability.Metric, []observability.LogRecord) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for {
		events, metrics, logs := r.snapshot()
		if predicate(events, metrics, logs) {
			return events, metrics, logs
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s; events=%#v metrics=%#v logs=%#v", description, events, metrics, logs)
		}
		time.Sleep(time.Millisecond)
	}
}

func TestObservabilityConnectionLifecycle(t *testing.T) {
	recorder := &observationRecorder{}
	conn := &mockConn{}
	service := NewService(conn, nil,
		WithConnectionID(42),
		WithLogger(observability.LoggerFunc(recorder.log)),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.observeMetric)),
	)

	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	events, metrics, logs := recorder.snapshot()
	if len(events) != 2 {
		t.Fatalf("event count = %d, want 2: %#v", len(events), events)
	}
	if events[0].Kind != observability.EventConnectionOpened || events[1].Kind != observability.EventConnectionClosed {
		t.Fatalf("event kinds = [%q, %q], want [connection_opened, connection_closed]",
			events[0].Kind, events[1].Kind)
	}
	for _, event := range events {
		if event.Association.ConnectionID != 42 {
			t.Errorf("connection ID = %d, want 42", event.Association.ConnectionID)
		}
	}
	if len(metrics) != 2 || metrics[0].Kind != observability.MetricConnection || metrics[1].Kind != observability.MetricConnection {
		t.Fatalf("connection metrics = %#v, want open and close", metrics)
	}
	if len(logs) != 2 {
		t.Fatalf("log count = %d, want 2", len(logs))
	}
}

func TestObservabilityAssociationRequestAndBytes(t *testing.T) {
	recorder := &observationRecorder{}
	conn := &mockConn{}
	service := NewService(conn, nil,
		WithConnectionID(77),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.observeMetric)),
	)
	defer func() { _ = service.Close() }()

	rq := &pdu.AAssociateRQ{
		ProtocolVersion:    1,
		CalledAETitle:      testCalledAE,
		CallingAETitle:     testCallingAE,
		ApplicationContext: testApplicationContext,
	}
	if err := service.SendAssociationRequest(context.Background(), rq); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}

	events, metrics, _ := recorder.snapshot()
	var requestEvent *observability.Event
	for i := range events {
		if events[i].Kind == observability.EventAssociationRequested {
			requestEvent = &events[i]
			break
		}
	}
	if requestEvent == nil {
		t.Fatalf("association request event not found: %#v", events)
	}
	if requestEvent.Direction != observability.DirectionOutbound ||
		requestEvent.Association.ConnectionID != 77 ||
		requestEvent.Association.AssociationID == 0 ||
		requestEvent.Association.CallingAE != testCallingAE ||
		requestEvent.Association.CalledAE != testCalledAE {
		t.Fatalf("association request event = %#v", *requestEvent)
	}

	var bytesMetric *observability.Metric
	for i := range metrics {
		if metrics[i].Kind == observability.MetricBytes {
			bytesMetric = &metrics[i]
			break
		}
	}
	if bytesMetric == nil {
		t.Fatalf("bytes metric not found: %#v", metrics)
	}
	if bytesMetric.Direction != observability.DirectionOutbound || bytesMetric.Value != int64(len(conn.writeData)) {
		t.Fatalf("bytes metric = %#v, written bytes = %d", *bytesMetric, len(conn.writeData))
	}
}

func TestObservabilityHookPanicDoesNotBreakServiceCreation(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil,
		WithEventObserver(observability.EventObserverFunc(func(context.Context, observability.Event) {
			panic("observer failure")
		})),
	)
	if err := service.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

func TestObservabilityRoutesAssociationDecodeWarningToLogger(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.CallingAETitle = testCallingAE
	rq.CalledAETitle = testCalledAE
	raw, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	raw.Data = append(raw.Data, 0x99, 0x00, 0x00, 0x00)
	raw.Length = uint32(len(raw.Data))

	recorder := &observationRecorder{}
	conn := &mockConn{readData: serializeRawPDUForTest(raw)}
	service := NewService(conn, nil,
		WithLogger(observability.LoggerFunc(recorder.log)),
	)
	defer func() { _ = service.Close() }()

	if _, err := service.ReceiveAssociationRequest(context.Background()); err != nil {
		t.Fatalf("ReceiveAssociationRequest() error = %v", err)
	}
	_, _, logs := recorder.snapshot()
	var warning *observability.LogRecord
	for i := range logs {
		if logs[i].Event.Kind == observability.EventDecodeWarning {
			warning = &logs[i]
			break
		}
	}
	if warning == nil {
		t.Fatalf("decode warning log not found: %#v", logs)
	}
	if warning.Level != observability.LevelWarn ||
		warning.Code != string(pdu.DecodeWarningUnknownItem) ||
		warning.ItemType != 0x99 ||
		warning.Event.Association.AssociationID == 0 {
		t.Fatalf("decode warning = %#v", *warning)
	}
}

func TestObservabilityCountsDIMSEPDUBytesBothDirections(t *testing.T) {
	clientRecorder := &observationRecorder{}
	serverRecorder := &observationRecorder{}
	clientService := startObservedServicePair(t, clientRecorder, serverRecorder, &Handlers{
		CEchoHandler: func(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err := clientService.SendCEcho(ctx, dimse.NewCEchoRequest()); err != nil {
		t.Fatalf("SendCEcho() error = %v", err)
	}

	assertByteDirections := func(name string, recorder *observationRecorder) {
		t.Helper()
		recorder.waitForSnapshot(t, name+" inbound and outbound byte metrics", func(
			_ []observability.Event,
			metrics []observability.Metric,
			_ []observability.LogRecord,
		) bool {
			directions := map[observability.Direction]bool{}
			for _, metric := range metrics {
				if metric.Kind == observability.MetricBytes && metric.Value >= 6 {
					directions[metric.Direction] = true
				}
			}
			return directions[observability.DirectionInbound] && directions[observability.DirectionOutbound]
		})
	}
	assertByteDirections("client", clientRecorder)
	assertByteDirections("server", serverRecorder)
}

func TestObservabilityAssociationReleaseLifecycle(t *testing.T) {
	clientRecorder := &observationRecorder{}
	serverRecorder := &observationRecorder{}
	clientService := startObservedServicePair(t, clientRecorder, serverRecorder, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := clientService.GracefulRelease(ctx); err != nil {
		t.Fatalf("GracefulRelease() error = %v", err)
	}

	clientEvents, _, _ := clientRecorder.waitForSnapshot(t, "client inbound association release", func(
		events []observability.Event,
		_ []observability.Metric,
		_ []observability.LogRecord,
	) bool {
		return containsAssociationEvent(events, observability.EventAssociationReleased, observability.DirectionInbound, observability.OutcomeSuccess)
	})
	serverEvents, _, _ := serverRecorder.waitForSnapshot(t, "server outbound association release", func(
		events []observability.Event,
		_ []observability.Metric,
		_ []observability.LogRecord,
	) bool {
		return containsAssociationEvent(events, observability.EventAssociationReleased, observability.DirectionOutbound, observability.OutcomeSuccess)
	})
	assertAssociationEvent(t, clientEvents, observability.EventAssociationReleased, observability.DirectionInbound, observability.OutcomeSuccess)
	assertAssociationEvent(t, serverEvents, observability.EventAssociationReleased, observability.DirectionOutbound, observability.OutcomeSuccess)
}

func TestObservabilityAssociationAbortLifecycle(t *testing.T) {
	recorder := &observationRecorder{}
	service := NewService(&mockConn{}, createTestAssociation(),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.observeMetric)),
	)
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	if err := service.SendAbort(context.Background(), pdu.AbortSourceServiceUser, pdu.AbortReasonServiceUserNotSpecified); err != nil {
		t.Fatalf("SendAbort() error = %v", err)
	}

	events, metrics, _ := recorder.snapshot()
	assertAssociationEvent(t, events, observability.EventAssociationAborted, observability.DirectionOutbound, observability.OutcomeAborted)
	foundBytes := false
	for _, metric := range metrics {
		if metric.Kind == observability.MetricBytes && metric.Direction == observability.DirectionOutbound && metric.Value == 10 {
			foundBytes = true
		}
	}
	if !foundBytes {
		t.Fatalf("outbound A-ABORT byte metric not found: %#v", metrics)
	}
}

func assertAssociationEvent(
	t *testing.T,
	events []observability.Event,
	kind observability.EventKind,
	direction observability.Direction,
	outcome observability.Outcome,
) {
	t.Helper()
	if containsAssociationEvent(events, kind, direction, outcome) {
		return
	}
	t.Fatalf("association event %q/%q/%q not found: %#v", kind, direction, outcome, events)
}

func containsAssociationEvent(
	events []observability.Event,
	kind observability.EventKind,
	direction observability.Direction,
	outcome observability.Outcome,
) bool {
	for _, event := range events {
		if event.Kind == kind && event.Direction == direction && event.Outcome == outcome {
			return true
		}
	}
	return false
}
