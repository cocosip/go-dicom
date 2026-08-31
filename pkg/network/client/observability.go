// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/observability"
)

func (c *Client) observationAssociation() observability.Association {
	return observability.Association{
		ConnectionID: c.connectionID,
		CallingAE:    c.config.CallingAE,
		CalledAE:     c.config.CalledAE,
	}
}

func (c *Client) emitClientEvent(ctx context.Context, event observability.Event, level observability.Level) {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	if event.Association.ConnectionID == 0 {
		event.Association = c.observationAssociation()
	}
	observability.EmitEvent(ctx, c.config.EventObserver, event)
	record := observability.LogRecord{
		Level: level, Message: string(event.Kind), Event: event,
	}
	observability.EmitSlog(ctx, "network.client", record)
}

func (c *Client) emitClientMetric(ctx context.Context, metric observability.Metric) {
	if metric.Timestamp.IsZero() {
		metric.Timestamp = time.Now()
	}
	if metric.Association.ConnectionID == 0 {
		metric.Association = c.observationAssociation()
	}
	observability.EmitMetric(ctx, c.config.MetricsObserver, metric)
}

func (c *Client) emitConnectionAttempted(ctx context.Context) {
	c.emitClientEvent(ctx, observability.Event{
		Kind: observability.EventConnectionAttempted,
	}, observability.LevelInfo)
	c.emitClientMetric(ctx, observability.Metric{
		Kind: observability.MetricConnection, Value: 1,
	})
}

func (c *Client) emitConnectionFailure(ctx context.Context, err error) {
	c.emitClientEvent(ctx, observability.Event{
		Kind: observability.EventConnectionClosed, Outcome: observability.OutcomeFailure, Error: err,
	}, observability.LevelError)
	c.emitClientMetric(ctx, observability.Metric{
		Kind: observability.MetricConnection, Outcome: observability.OutcomeFailure, Value: 1,
	})
	c.emitClientMetric(ctx, observability.Metric{
		Kind: observability.MetricError, Outcome: observability.OutcomeFailure,
		ErrorKind: observability.ErrorConnection, Value: 1,
	})
}
