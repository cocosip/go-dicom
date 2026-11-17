// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package logging provides the logger used by go-dicom packages.
package logging

import (
	"context"
	"errors"
	"log/slog"
	"sync/atomic"
)

// Config defines the process-wide logger used by go-dicom.
type Config struct {
	Handler slog.Handler
}

// Record is one structured go-dicom log event.
type Record struct {
	Level     slog.Level
	Component string
	Event     string
	Message   string
	Attrs     []slog.Attr
}

type runtime struct {
	logger *slog.Logger
}

var (
	activeRuntime atomic.Pointer[runtime]
	handlerPanics atomic.Uint64
)

// HandlerPanics reports how many times the configured handler panicked while
// logging was being written. Such records are dropped instead of failing the
// caller, so a non-zero count means the handler should be replaced.
func HandlerPanics() uint64 {
	return handlerPanics.Load()
}

// Configure atomically replaces the process-wide go-dicom logger. The logger
// is private to go-dicom and never reads or modifies slog.Default().
func Configure(config Config) error {
	if config.Handler == nil {
		return errors.New("logging handler cannot be nil")
	}
	activeRuntime.Store(&runtime{logger: slog.New(config.Handler)})
	return nil
}

// Disable atomically disables go-dicom logging. It does not close the
// configured handler or its underlying writer.
func Disable() {
	activeRuntime.Store(nil)
}

// Configured reports whether go-dicom logging is currently configured.
func Configured() bool {
	return activeRuntime.Load() != nil
}

// Enabled reports whether go-dicom logging is enabled for level.
//
// A panic from the configured handler is recovered and reported as disabled, so
// a faulty handler can never crash the library caller; the panic is counted by
// HandlerPanics.
func Enabled(ctx context.Context, level slog.Level) (enabled bool) {
	defer func() {
		if recover() != nil {
			handlerPanics.Add(1)
			enabled = false
		}
	}()
	current := activeRuntime.Load()
	return current != nil && current.logger.Enabled(normalizeContext(ctx), level)
}

// Emit writes a structured record through the private go-dicom logger.
//
// A panic from the configured handler is recovered, the record is dropped, and
// the panic is counted by HandlerPanics, so logging can never crash the
// operation being logged.
func Emit(ctx context.Context, record Record) {
	defer func() {
		if recover() != nil {
			handlerPanics.Add(1)
		}
	}()
	current := activeRuntime.Load()
	ctx = normalizeContext(ctx)
	if current == nil || !current.logger.Enabled(ctx, record.Level) {
		return
	}

	attrs := []slog.Attr{
		slog.String("component", record.Component),
		slog.String("event", record.Event),
	}
	attrs = append(attrs, record.Attrs...)
	current.logger.LogAttrs(ctx, record.Level, record.Message, attrs...)
}

func normalizeContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
