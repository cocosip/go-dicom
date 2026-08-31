// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package observability defines vendor-neutral hooks for DICOM network logs,
// lifecycle events, and metrics. Records contain protocol metadata only; they
// never contain DIMSE datasets or raw PDU payloads.
package observability

import (
	"context"
	"sync/atomic"
	"time"
)

// ConnectionID uniquely identifies a connection within the current process.
type ConnectionID uint64

// AssociationID uniquely identifies an association within the current process.
type AssociationID uint64

var (
	nextConnectionID  atomic.Uint64
	nextAssociationID atomic.Uint64
)

// NewConnectionID returns a new non-zero process-local connection identifier.
func NewConnectionID() ConnectionID {
	return ConnectionID(nextConnectionID.Add(1))
}

// NewAssociationID returns a new non-zero process-local association identifier.
func NewAssociationID() AssociationID {
	return AssociationID(nextAssociationID.Add(1))
}

// Level classifies the severity of a structured log record.
type Level string

const (
	// LevelDebug enables detailed diagnostic records.
	LevelDebug Level = "debug"
	// LevelInfo reports ordinary lifecycle records.
	LevelInfo Level = "info"
	// LevelWarn reports recoverable or forward-compatible conditions.
	LevelWarn Level = "warn"
	// LevelError reports failed operations.
	LevelError Level = "error"
)

// Direction identifies whether an observation entered or left the local AE.
type Direction string

const (
	// DirectionNone indicates that direction does not apply.
	DirectionNone Direction = ""
	// DirectionInbound identifies data received by the local AE.
	DirectionInbound Direction = "inbound"
	// DirectionOutbound identifies data sent by the local AE.
	DirectionOutbound Direction = "outbound"
)

// Outcome classifies the result of an observed operation.
type Outcome string

const (
	// OutcomeNone indicates that an outcome does not apply.
	OutcomeNone Outcome = ""
	// OutcomeSuccess indicates successful completion.
	OutcomeSuccess Outcome = "success"
	// OutcomePending indicates that more responses may follow.
	OutcomePending Outcome = "pending"
	// OutcomeWarning indicates completion with a warning.
	OutcomeWarning Outcome = "warning"
	// OutcomeFailure indicates unsuccessful completion.
	OutcomeFailure Outcome = "failure"
	// OutcomeTimeout indicates deadline expiry.
	OutcomeTimeout Outcome = "timeout"
	// OutcomeCancelled indicates cancellation.
	OutcomeCancelled Outcome = "cancelled"
	// OutcomeRejected indicates association rejection.
	OutcomeRejected Outcome = "rejected"
	// OutcomeAborted indicates association abort.
	OutcomeAborted Outcome = "aborted"
)

// EventKind identifies a connection, association, or DIMSE lifecycle event.
type EventKind string

const (
	// EventConnectionAttempted reports an outbound connection attempt.
	EventConnectionAttempted EventKind = "connection_attempted"
	// EventConnectionOpened reports an established connection.
	EventConnectionOpened EventKind = "connection_opened"
	// EventConnectionClosed reports a closed connection.
	EventConnectionClosed EventKind = "connection_closed"
	// EventAssociationRequested reports an association request.
	EventAssociationRequested EventKind = "association_requested"
	// EventAssociationAccepted reports association acceptance.
	EventAssociationAccepted EventKind = "association_accepted"
	// EventAssociationRejected reports association rejection.
	EventAssociationRejected EventKind = "association_rejected"
	// EventAssociationReleased reports graceful association release.
	EventAssociationReleased EventKind = "association_released"
	// EventAssociationAborted reports association abort.
	EventAssociationAborted EventKind = "association_aborted"
	// EventRequestSent reports a fully written outbound DIMSE request.
	EventRequestSent EventKind = "request_sent"
	// EventRequestReceived reports an inbound DIMSE request.
	EventRequestReceived EventKind = "request_received"
	// EventRequestPending reports a pending DIMSE response.
	EventRequestPending EventKind = "request_pending"
	// EventRequestCompleted reports a terminal DIMSE response.
	EventRequestCompleted EventKind = "request_completed"
	// EventRequestTimedOut reports a request deadline expiry.
	EventRequestTimedOut EventKind = "request_timed_out"
	// EventRequestCancelled reports a cancelled request.
	EventRequestCancelled EventKind = "request_cancelled"
	// EventRequestFailed reports a failed request.
	EventRequestFailed EventKind = "request_failed"
	// EventDecodeWarning reports a forward-compatible decode warning.
	EventDecodeWarning EventKind = "decode_warning"
)

// MetricKind identifies a vendor-neutral network measurement.
type MetricKind string

const (
	// MetricConnection counts connection lifecycle observations.
	MetricConnection MetricKind = "connection"
	// MetricAssociation counts association lifecycle observations.
	MetricAssociation MetricKind = "association"
	// MetricDIMSE counts DIMSE lifecycle observations.
	MetricDIMSE MetricKind = "dimse"
	// MetricBytes records complete PDU byte counts.
	MetricBytes MetricKind = "bytes"
	// MetricError counts classified errors.
	MetricError MetricKind = "error"
	// MetricLatency records terminal request duration.
	MetricLatency MetricKind = "latency"
)

// ErrorKind classifies an error without tying it to a telemetry vendor.
type ErrorKind string

const (
	// ErrorNone indicates that error classification does not apply.
	ErrorNone ErrorKind = ""
	// ErrorConnection classifies connection setup or teardown errors.
	ErrorConnection ErrorKind = "connection"
	// ErrorAssociation classifies association negotiation errors.
	ErrorAssociation ErrorKind = "association"
	// ErrorTransport classifies network transport errors.
	ErrorTransport ErrorKind = "transport"
	// ErrorProtocol classifies malformed or unexpected protocol data.
	ErrorProtocol ErrorKind = "protocol"
	// ErrorDIMSE classifies DIMSE operation errors.
	ErrorDIMSE ErrorKind = "dimse"
)

// Association is a safe immutable snapshot used to correlate observations.
type Association struct {
	ConnectionID  ConnectionID
	AssociationID AssociationID
	CallingAE     string
	CalledAE      string
}

// Event describes one network lifecycle transition.
type Event struct {
	Timestamp   time.Time
	Kind        EventKind
	Association Association
	Direction   Direction
	MessageID   uint16
	Command     uint16
	StatusCode  uint16
	StatusState string
	Duration    time.Duration
	Outcome     Outcome
	Error       error
}

// LogRecord is a structured logger input.
type LogRecord struct {
	Level    Level
	Message  string
	Code     string
	ItemType byte
	Event    Event
}

// Metric describes one count, byte, error, or latency observation.
type Metric struct {
	Timestamp   time.Time
	Kind        MetricKind
	Association Association
	Direction   Direction
	Command     uint16
	StatusCode  uint16
	Outcome     Outcome
	ErrorKind   ErrorKind
	Value       int64
	Duration    time.Duration
}

// EventObserver accepts structured lifecycle events.
type EventObserver interface {
	ObserveEvent(context.Context, Event)
}

// EventObserverFunc adapts a function to EventObserver.
type EventObserverFunc func(context.Context, Event)

// ObserveEvent delivers an event to the adapted function.
func (f EventObserverFunc) ObserveEvent(ctx context.Context, event Event) { f(ctx, event) }

// NopEventObserver discards all lifecycle events.
type NopEventObserver struct{}

// ObserveEvent discards the event.
func (NopEventObserver) ObserveEvent(context.Context, Event) {}

// MetricsObserver accepts vendor-neutral network metrics.
type MetricsObserver interface {
	ObserveMetric(context.Context, Metric)
}

// MetricsObserverFunc adapts a function to MetricsObserver.
type MetricsObserverFunc func(context.Context, Metric)

// ObserveMetric delivers a metric to the adapted function.
func (f MetricsObserverFunc) ObserveMetric(ctx context.Context, metric Metric) { f(ctx, metric) }

// NopMetricsObserver discards all metrics.
type NopMetricsObserver struct{}

// ObserveMetric discards the metric.
func (NopMetricsObserver) ObserveMetric(context.Context, Metric) {}
