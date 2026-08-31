// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package logging

import (
	"context"
	"log/slog"
	"sync"
	"testing"
)

const (
	testComponent         = "test"
	testMessageNotEmitted = "not emitted"
)

func TestUnconfiguredLoggingIsSilent(t *testing.T) {
	Disable()
	t.Cleanup(Disable)

	if Configured() {
		t.Fatal("Configured() = true without configuration")
	}
	if Enabled(context.Background(), slog.LevelError) {
		t.Fatal("Enabled() = true without configuration")
	}

	Emit(context.Background(), Record{
		Level: slog.LevelError, Component: testComponent, Event: "not_configured", Message: testMessageNotEmitted,
	})
}

func TestConfigureDoesNotUseSlogDefault(t *testing.T) {
	defaultHandler := &recordingHandler{}
	configuredHandler := &recordingHandler{}
	previousDefault := slog.Default()
	slog.SetDefault(slog.New(defaultHandler))
	t.Cleanup(func() {
		Disable()
		slog.SetDefault(previousDefault)
	})

	if err := Configure(Config{Handler: configuredHandler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	Emit(context.Background(), Record{
		Level: slog.LevelInfo, Component: "dicom.parser", Event: "parse_completed", Message: "completed",
	})

	if got := len(configuredHandler.snapshot()); got != 1 {
		t.Fatalf("configured handler records = %d, want 1", got)
	}
	if got := len(defaultHandler.snapshot()); got != 0 {
		t.Fatalf("default handler records = %d, want 0", got)
	}
}

func TestDisableStopsEmission(t *testing.T) {
	handler := &recordingHandler{}
	if err := Configure(Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(Disable)

	Disable()
	Emit(context.Background(), Record{
		Level: slog.LevelError, Component: "dicom.writer", Event: "write_failed", Message: "failed",
	})

	if Configured() {
		t.Fatal("Configured() = true after Disable()")
	}
	if got := len(handler.snapshot()); got != 0 {
		t.Fatalf("records after Disable() = %d, want 0", got)
	}
}

func TestConfigureRejectsNilHandler(t *testing.T) {
	Disable()
	t.Cleanup(Disable)

	if err := Configure(Config{}); err == nil {
		t.Fatal("Configure() error = nil, want nil-handler error")
	}
	if Configured() {
		t.Fatal("Configured() = true after rejected configuration")
	}
}

func TestConcurrentConfigureEmitAndDisable(t *testing.T) {
	Disable()
	t.Cleanup(Disable)

	const goroutines = 8
	const iterations = 100
	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range iterations {
				_ = Configure(Config{Handler: &recordingHandler{}})
				Emit(context.Background(), Record{
					Level: slog.LevelDebug, Component: testComponent, Event: "concurrent", Message: "concurrent",
				})
				Disable()
			}
		}()
	}
	wg.Wait()
}

func TestHandlerPanicDoesNotEscapeLoggingBoundary(t *testing.T) {
	t.Run("Enabled", func(t *testing.T) {
		if err := Configure(Config{Handler: panicHandler{panicEnabled: true}}); err != nil {
			t.Fatalf("Configure() error = %v", err)
		}
		t.Cleanup(Disable)

		if Enabled(context.Background(), slog.LevelInfo) {
			t.Fatal("Enabled() = true after Handler.Enabled panic")
		}
		Emit(context.Background(), Record{
			Level: slog.LevelInfo, Component: testComponent, Event: "enabled_panic", Message: testMessageNotEmitted,
		})
	})

	t.Run("Handle", func(t *testing.T) {
		if err := Configure(Config{Handler: panicHandler{panicHandle: true}}); err != nil {
			t.Fatalf("Configure() error = %v", err)
		}
		t.Cleanup(Disable)

		Emit(context.Background(), Record{
			Level: slog.LevelInfo, Component: testComponent, Event: "handle_panic", Message: testMessageNotEmitted,
		})
	})
}

type panicHandler struct {
	panicEnabled bool
	panicHandle  bool
}

func (h panicHandler) Enabled(context.Context, slog.Level) bool {
	if h.panicEnabled {
		panic("Enabled")
	}
	return true
}

func (h panicHandler) Handle(context.Context, slog.Record) error {
	if h.panicHandle {
		panic("Handle")
	}
	return nil
}

func (h panicHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h panicHandler) WithGroup(string) slog.Handler { return h }

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
