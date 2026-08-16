// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package observability

import "context"

// EmitLog delivers a log record when its level is enabled. It returns false
// when the logger is absent, disabled, or panics.
func EmitLog(ctx context.Context, logger Logger, record LogRecord) (delivered bool) {
	defer func() {
		if recover() != nil {
			delivered = false
		}
	}()
	if logger == nil || !logger.Enabled(record.Level) {
		return false
	}
	logger.Log(ctx, record)
	return true
}

// EmitEvent delivers an event and returns false when the observer is absent or panics.
func EmitEvent(ctx context.Context, observer EventObserver, event Event) (delivered bool) {
	if observer == nil {
		return false
	}
	defer func() {
		if recover() != nil {
			delivered = false
		}
	}()
	observer.ObserveEvent(ctx, event)
	return true
}

// EmitMetric delivers a metric and returns false when the observer is absent or panics.
func EmitMetric(ctx context.Context, observer MetricsObserver, metric Metric) (delivered bool) {
	if observer == nil {
		return false
	}
	defer func() {
		if recover() != nil {
			delivered = false
		}
	}()
	observer.ObserveMetric(ctx, metric)
	return true
}
