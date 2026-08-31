// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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

var activeRuntime atomic.Pointer[runtime]

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
func Enabled(ctx context.Context, level slog.Level) (enabled bool) {
	defer func() {
		if recover() != nil {
			enabled = false
		}
	}()
	current := activeRuntime.Load()
	return current != nil && current.logger.Enabled(normalizeContext(ctx), level)
}

// Emit writes a structured record through the private go-dicom logger.
func Emit(ctx context.Context, record Record) {
	defer func() {
		_ = recover()
	}()
	current := activeRuntime.Load()
	ctx = normalizeContext(ctx)
	if current == nil || !current.logger.Enabled(ctx, record.Level) {
		return
	}

	attrs := make([]slog.Attr, 0, len(record.Attrs)+2)
	attrs = append(attrs,
		slog.String("component", record.Component),
		slog.String("event", record.Event),
	)
	attrs = append(attrs, record.Attrs...)
	current.logger.LogAttrs(ctx, record.Level, record.Message, attrs...)
}

func normalizeContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
