// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package observability

import "context"

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
