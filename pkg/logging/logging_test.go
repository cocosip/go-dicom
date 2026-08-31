// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package logging

import (
	"context"
	"log/slog"
	"sync"
	"testing"
)

func TestSetLoggerScopesRecordsToGoDicom(t *testing.T) {
	standardDefault := slog.Default()
	handler := &recordingHandler{}
	configured := slog.New(handler)
	previous := Logger()
	SetLogger(configured)
	t.Cleanup(func() { SetLogger(previous) })

	LogAttrs(context.Background(), slog.LevelInfo, "DICOM operation completed", slog.String("event", "completed"))

	if slog.Default() != standardDefault {
		t.Fatal("SetLogger() changed slog.Default()")
	}
	if records := handler.snapshot(); len(records) != 1 || records[0].Message != "DICOM operation completed" {
		t.Fatalf("records = %#v, want one go-dicom record", records)
	}
}

func TestSetLoggerNilDisablesGoDicomRecords(t *testing.T) {
	handler := &recordingHandler{}
	previous := Logger()
	SetLogger(slog.New(handler))
	t.Cleanup(func() { SetLogger(previous) })

	SetLogger(nil)
	if Logger() != nil {
		t.Fatal("Logger() = non-nil after SetLogger(nil), want nil")
	}
	LogAttrs(context.Background(), slog.LevelError, "DICOM operation failed", slog.String("event", "failed"))

	if records := handler.snapshot(); len(records) != 0 {
		t.Fatalf("records = %#v, want no records", records)
	}
}

type recordingHandler struct {
	mu      sync.Mutex
	records []slog.Record
}

func (h *recordingHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *recordingHandler) Handle(_ context.Context, record slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, record.Clone())
	return nil
}

func (h *recordingHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *recordingHandler) WithGroup(string) slog.Handler { return h }

func (h *recordingHandler) snapshot() []slog.Record {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]slog.Record(nil), h.records...)
}
