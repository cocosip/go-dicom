// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/logging"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestServiceWritesLifecycleToGoDicomLogger(t *testing.T) {
	handler := &serviceSlogHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	svc := NewService(&mockConn{}, nil, WithConnectionID(42))
	if err := svc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	records := handler.snapshot()
	if len(records) != 2 {
		t.Fatalf("slog records = %d, want 2", len(records))
	}
	for index, wantEvent := range []observability.EventKind{
		observability.EventConnectionOpened,
		observability.EventConnectionClosed,
	} {
		record := records[index]
		if record.Message != string(wantEvent) {
			t.Errorf("record %d message = %q, want %q", index, record.Message, wantEvent)
		}
		attrs := serviceSlogAttrs(record)
		if attrs["component"] != "network.service" || attrs["connection_id"] != uint64(42) {
			t.Errorf("record %d attrs = %#v", index, attrs)
		}
	}
}

func TestGlobalLoggerCreatesDIMSERequestLifecycle(t *testing.T) {
	handler := &serviceSlogHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	svc := NewService(&mockConn{}, createTestAssociation(), WithConnectionID(43))
	t.Cleanup(func() { _ = svc.Close() })
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(7); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	lifecycle := newRequestLifecycle(svc, observability.DirectionOutbound, req)
	if lifecycle == nil {
		t.Fatal("newRequestLifecycle() = nil with global logger configured")
	}
	lifecycle.sent(context.Background())
	lifecycle.finishError(context.Background(), errors.New("test failure"))

	records := handler.snapshot()
	if len(records) < 2 || records[len(records)-2].Message != string(observability.EventRequestSent) || records[len(records)-1].Message != string(observability.EventRequestFailed) {
		t.Fatalf("request records = %#v, want request_sent and request_failed", records)
	}
}

func TestErrorOnlyGlobalLoggerReceivesTerminalDIMSEFailure(t *testing.T) {
	handler := &serviceErrorHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	svc := NewService(&mockConn{}, createTestAssociation(), WithConnectionID(44))
	t.Cleanup(func() { _ = svc.Close() })
	lifecycle := newRequestLifecycle(svc, observability.DirectionOutbound, dimse.NewCEchoRequest())
	if lifecycle == nil {
		t.Fatal("newRequestLifecycle() = nil with Error-only global logger configured")
	}
	lifecycle.sent(context.Background())
	lifecycle.finishError(context.Background(), errors.New("test failure"))

	records := handler.snapshot()
	if len(records) != 1 || records[0].Message != string(observability.EventRequestFailed) || records[0].Level != slog.LevelError {
		t.Fatalf("Error-only records = %#v, want one request_failed Error record", records)
	}
}

func TestPendingDIMSELogUsesInfoLevel(t *testing.T) {
	handler := &serviceSlogHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	svc := NewService(&mockConn{}, createTestAssociation(), WithConnectionID(45))
	t.Cleanup(func() { _ = svc.Close() })
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, nil)
	if err := req.SetMessageID(8); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	lifecycle := newRequestLifecycle(svc, observability.DirectionOutbound, req)
	lifecycle.sent(context.Background())
	lifecycle.response(context.Background(), dimse.NewCFindResponseFromRequest(req, status.Pending, nil))
	lifecycle.response(context.Background(), dimse.NewCFindResponseFromRequest(req, status.Success, nil))

	records := handler.snapshot()
	var pending, completed *slog.Record
	for i := range records {
		switch records[i].Message {
		case string(observability.EventRequestPending):
			pending = &records[i]
		case string(observability.EventRequestCompleted):
			completed = &records[i]
		}
	}
	if pending == nil || pending.Level != slog.LevelInfo {
		t.Fatalf("pending record = %#v, want INFO", pending)
	}
	if completed == nil || completed.Level != slog.LevelInfo {
		t.Fatalf("completed record = %#v, want INFO", completed)
	}
}

func TestRequestLifecycleCapturesQueryRetrieveMetadata(t *testing.T) {
	recorder := &observationRecorder{}
	svc := NewService(&mockConn{}, createTestAssociation(),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
	)
	t.Cleanup(func() { _ = svc.Close() })

	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, testMoveDestinationAE, nil)
	if err := req.SetMessageID(9); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	lifecycle := newRequestLifecycle(svc, observability.DirectionOutbound, req)
	if lifecycle == nil {
		t.Fatal("newRequestLifecycle() = nil")
	}
	lifecycle.sent(context.Background())

	events, _, _ := recorder.snapshot()
	var event observability.Event
	for _, candidate := range events {
		if candidate.Kind == observability.EventRequestSent {
			event = candidate
			break
		}
	}
	if event.Kind != observability.EventRequestSent {
		t.Fatalf("request event not found: %#v", events)
	}
	if event.CommandName != "C-MOVE-RQ" || event.Operation != "C-MOVE" ||
		event.QueryRetrieveLevel != "STUDY" || event.MoveDestinationAE != testMoveDestinationAE ||
		event.SOPClassUID == "" || event.OperationID == 0 {
		t.Fatalf("request metadata = %#v", event)
	}
}

type serviceSlogHandler struct {
	mu      sync.Mutex
	records []slog.Record
}

func (h *serviceSlogHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *serviceSlogHandler) Handle(_ context.Context, record slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, record.Clone())
	return nil
}

func (h *serviceSlogHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *serviceSlogHandler) WithGroup(string) slog.Handler { return h }

func (h *serviceSlogHandler) snapshot() []slog.Record {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]slog.Record(nil), h.records...)
}

type serviceErrorHandler struct {
	serviceSlogHandler
}

func (h *serviceErrorHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= slog.LevelError
}

func serviceSlogAttrs(record slog.Record) map[string]any {
	attrs := make(map[string]any)
	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Any()
		return true
	})
	return attrs
}
