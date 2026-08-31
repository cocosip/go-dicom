// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package server

import (
	"context"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/observability"
)

func (s *Server) emitServerError(ctx context.Context, message string, err error) {
	event := observability.Event{
		Timestamp: time.Now(),
		Outcome:   observability.OutcomeFailure,
		Error:     err,
	}
	record := observability.LogRecord{
		Level:   observability.LevelError,
		Message: message,
		Event:   event,
	}
	observability.EmitSlog(ctx, "network.server", record)
	observability.EmitLog(ctx, s.config.Logger, record)
	observability.EmitMetric(ctx, s.config.MetricsObserver, observability.Metric{
		Timestamp: time.Now(),
		Kind:      observability.MetricError,
		Outcome:   observability.OutcomeFailure,
		ErrorKind: observability.ErrorConnection,
		Value:     1,
	})
}
