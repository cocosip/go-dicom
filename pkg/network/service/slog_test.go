// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"log/slog"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/logging"
	"github.com/cocosip/go-dicom/pkg/network/observability"
)

func TestServiceWritesLifecycleToGoDicomLogger(t *testing.T) {
	handler := &serviceSlogHandler{}
	previous := logging.Logger()
	logging.SetLogger(slog.New(handler))
	t.Cleanup(func() { logging.SetLogger(previous) })

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

func serviceSlogAttrs(record slog.Record) map[string]any {
	attrs := make(map[string]any)
	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Any()
		return true
	})
	return attrs
}
