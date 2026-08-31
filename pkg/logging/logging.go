// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package logging provides the logger used by go-dicom packages.
package logging

import (
	"context"
	"log/slog"
	"sync/atomic"
)

var activeLogger atomic.Pointer[slog.Logger]

// SetLogger configures the logger used by go-dicom packages. Passing nil
// disables go-dicom logging. This does not modify slog.Default().
func SetLogger(logger *slog.Logger) {
	activeLogger.Store(logger)
}

// Logger returns the logger configured for go-dicom packages.
func Logger() *slog.Logger {
	return activeLogger.Load()
}

// Enabled reports whether go-dicom logging is enabled for level.
func Enabled(ctx context.Context, level slog.Level) bool {
	logger := Logger()
	return logger != nil && logger.Enabled(ctx, level)
}

// LogAttrs records a message through the logger configured for go-dicom.
func LogAttrs(ctx context.Context, level slog.Level, message string, attrs ...slog.Attr) {
	if logger := Logger(); logger != nil {
		logger.LogAttrs(ctx, level, message, attrs...)
	}
}
