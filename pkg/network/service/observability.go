// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/logging"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

type queuedRequestEvent struct {
	ctx   context.Context
	event observability.Event
}

type requestLifecycle struct {
	service     *Service
	association observability.Association
	direction   observability.Direction
	messageID   uint16
	command     uint16
	started     time.Time

	mu       sync.Mutex
	initial  bool
	terminal bool
	emitting bool
	queue    []queuedRequestEvent
}

func newRequestLifecycle(s *Service, direction observability.Direction, req dimse.Request) *requestLifecycle {
	if s == nil || (!logging.Configured() && s.eventObserver == nil && s.metricsObserver == nil) {
		return nil
	}
	messageID := req.MessageID()
	if cancel, ok := req.(*dimse.CCancelRequest); ok {
		messageID = cancel.MessageIDBeingRespondedTo()
	}
	return &requestLifecycle{
		service:     s,
		association: s.observationAssociation(),
		direction:   direction,
		messageID:   messageID,
		command:     req.CommandField(),
		started:     time.Now(),
	}
}

func (l *requestLifecycle) received(ctx context.Context) {
	if l == nil {
		return
	}
	l.mu.Lock()
	if !l.initial && !l.terminal {
		l.initial = true
		l.enqueueLocked(ctx, observability.EventRequestReceived, 0, "", observability.OutcomeSuccess, nil)
	}
	l.startDrainerLocked()
}

func (l *requestLifecycle) sent(ctx context.Context) {
	if l == nil {
		return
	}
	l.mu.Lock()
	if !l.initial && !l.terminal {
		l.initial = true
		l.enqueueLocked(ctx, observability.EventRequestSent, 0, "", observability.OutcomeSuccess, nil)
	}
	l.startDrainerLocked()
}

func (l *requestLifecycle) response(ctx context.Context, resp dimse.Response) {
	if l == nil || resp == nil {
		return
	}
	statusValue := resp.Status()
	statusCode := uint16(0)
	statusState := ""
	outcome := observability.OutcomeFailure
	if statusValue != nil {
		statusCode = statusValue.Code
		statusState = statusValue.State
		outcome = outcomeForResponse(resp)
	}

	l.mu.Lock()
	if l.terminal {
		l.mu.Unlock()
		return
	}
	if !l.initial {
		l.initial = true
		kind := observability.EventRequestReceived
		if l.direction == observability.DirectionOutbound {
			kind = observability.EventRequestSent
		}
		l.enqueueLocked(ctx, kind, 0, "", observability.OutcomeSuccess, nil)
	}
	if resp.IsPending() {
		l.enqueueLocked(ctx, observability.EventRequestPending, statusCode, statusState, observability.OutcomePending, nil)
	} else {
		l.terminal = true
		l.enqueueLocked(ctx, observability.EventRequestCompleted, statusCode, statusState, outcome, nil)
	}
	l.startDrainerLocked()
}

func outcomeForResponse(resp dimse.Response) observability.Outcome {
	if resp.IsPending() {
		return observability.OutcomePending
	}
	statusValue := resp.Status()
	if statusValue == nil {
		return observability.OutcomeFailure
	}
	switch {
	case statusValue.IsSuccess():
		return observability.OutcomeSuccess
	case statusValue.IsWarning():
		return observability.OutcomeWarning
	case statusValue.IsCancel():
		return observability.OutcomeCancelled
	default:
		return observability.OutcomeFailure
	}
}

func (l *requestLifecycle) finishContext(ctx context.Context, err error) {
	if errors.Is(err, context.DeadlineExceeded) {
		l.finish(ctx, observability.EventRequestTimedOut, observability.OutcomeTimeout, err)
		return
	}
	if errors.Is(err, context.Canceled) {
		l.finish(ctx, observability.EventRequestCancelled, observability.OutcomeCancelled, err)
		return
	}
	l.finishError(ctx, err)
}

func (l *requestLifecycle) finishError(ctx context.Context, err error) {
	l.finish(ctx, observability.EventRequestFailed, observability.OutcomeFailure, err)
}

func (l *requestLifecycle) finish(ctx context.Context, kind observability.EventKind, outcome observability.Outcome, err error) {
	if l == nil {
		return
	}
	l.mu.Lock()
	if l.terminal {
		l.mu.Unlock()
		return
	}
	l.terminal = true
	l.enqueueLocked(ctx, kind, 0, "", outcome, err)
	l.startDrainerLocked()
}

func (l *requestLifecycle) enqueueLocked(
	ctx context.Context,
	kind observability.EventKind,
	statusCode uint16,
	statusState string,
	outcome observability.Outcome,
	err error,
) {
	if ctx == nil {
		ctx = context.Background()
	}
	l.queue = append(l.queue, queuedRequestEvent{ctx: ctx, event: observability.Event{
		Kind:        kind,
		Association: l.association,
		Direction:   l.direction,
		MessageID:   l.messageID,
		Command:     l.command,
		StatusCode:  statusCode,
		StatusState: statusState,
		Duration:    time.Since(l.started),
		Outcome:     outcome,
		Error:       err,
	}})
}

// startDrainerLocked transfers ownership of queued event delivery to one
// goroutine. User hooks run after the lifecycle lock has been released.
func (l *requestLifecycle) startDrainerLocked() {
	if l.emitting {
		l.mu.Unlock()
		return
	}
	l.emitting = true
	l.mu.Unlock()
	l.drain()
}

func (l *requestLifecycle) drain() {
	for {
		l.mu.Lock()
		if len(l.queue) == 0 {
			l.emitting = false
			l.mu.Unlock()
			return
		}
		queued := l.queue[0]
		l.queue = l.queue[1:]
		l.mu.Unlock()
		l.service.emitRequestObservation(queued.ctx, queued.event)
	}
}

func (s *Service) emitRequestObservation(ctx context.Context, event observability.Event) {
	s.emitEvent(ctx, event)
	s.emitMetric(ctx, observability.Metric{
		Kind:        observability.MetricDIMSE,
		Association: event.Association,
		Direction:   event.Direction,
		Command:     event.Command,
		StatusCode:  event.StatusCode,
		Outcome:     event.Outcome,
		Value:       1,
	})
	switch event.Kind {
	case observability.EventRequestCompleted,
		observability.EventRequestTimedOut,
		observability.EventRequestCancelled,
		observability.EventRequestFailed:
		s.emitMetric(ctx, observability.Metric{
			Kind:        observability.MetricLatency,
			Association: event.Association,
			Direction:   event.Direction,
			Command:     event.Command,
			StatusCode:  event.StatusCode,
			Outcome:     event.Outcome,
			Duration:    event.Duration,
			Value:       1,
		})
	}
	if event.Kind == observability.EventRequestTimedOut || event.Kind == observability.EventRequestFailed {
		s.emitMetric(ctx, observability.Metric{
			Kind:        observability.MetricError,
			Association: event.Association,
			Direction:   event.Direction,
			Command:     event.Command,
			StatusCode:  event.StatusCode,
			Outcome:     event.Outcome,
			ErrorKind:   observability.ErrorDIMSE,
			Value:       1,
		})
	}
}

func (s *Service) lifecycleForSend(message dimse.Message) *requestLifecycle {
	if request, ok := message.(dimse.Request); ok {
		s.pendingRequestsMu.RLock()
		pending := s.pendingRequests[request.MessageID()]
		s.pendingRequestsMu.RUnlock()
		if pending != nil {
			return pending.lifecycle
		}
		if _, ok := request.(*dimse.CCancelRequest); ok {
			return newRequestLifecycle(s, observability.DirectionOutbound, request)
		}
		return nil
	}
	if response, ok := message.(dimse.Response); ok {
		return s.inboundRequest(response.MessageIDBeingRespondedTo())
	}
	return nil
}

func (s *Service) observeSentMessage(message dimse.Message, lifecycle *requestLifecycle, err error) {
	if lifecycle == nil {
		return
	}
	if err != nil {
		lifecycle.finishContext(context.Background(), err)
		return
	}
	if request, ok := message.(dimse.Request); ok {
		lifecycle.sent(s.ctx)
		if _, ok := request.(*dimse.CCancelRequest); ok {
			lifecycle.finish(s.ctx, observability.EventRequestCompleted, observability.OutcomeSuccess, nil)
		}
		return
	}
	response, ok := message.(dimse.Response)
	if !ok {
		return
	}
	lifecycle.response(s.ctx, response)
	if !response.IsPending() {
		s.unregisterInboundRequest(response.MessageIDBeingRespondedTo(), lifecycle)
	}
}

func (s *Service) registerInboundRequest(req dimse.Request) *requestLifecycle {
	lifecycle := newRequestLifecycle(s, observability.DirectionInbound, req)
	if lifecycle == nil {
		return nil
	}
	s.inboundRequestsMu.Lock()
	s.inboundRequests[req.MessageID()] = lifecycle
	s.inboundRequestsMu.Unlock()
	return lifecycle
}

func (s *Service) inboundRequest(messageID uint16) *requestLifecycle {
	s.inboundRequestsMu.RLock()
	defer s.inboundRequestsMu.RUnlock()
	return s.inboundRequests[messageID]
}

func (s *Service) unregisterInboundRequest(messageID uint16, lifecycle *requestLifecycle) {
	if lifecycle == nil {
		return
	}
	s.inboundRequestsMu.Lock()
	if current := s.inboundRequests[messageID]; current == lifecycle {
		delete(s.inboundRequests, messageID)
	}
	s.inboundRequestsMu.Unlock()
}

func (s *Service) observationAssociation() observability.Association {
	s.assocMu.RLock()
	defer s.assocMu.RUnlock()
	return observability.Association{
		ConnectionID:  s.connectionID,
		AssociationID: s.associationID,
		CallingAE:     s.callingAE,
		CalledAE:      s.calledAE,
	}
}

func (s *Service) ensureAssociationIdentity(callingAE, calledAE string) observability.Association {
	s.assocMu.Lock()
	defer s.assocMu.Unlock()
	if s.associationID == 0 {
		s.associationID = observability.NewAssociationID()
	}
	if callingAE != "" {
		s.callingAE = callingAE
	}
	if calledAE != "" {
		s.calledAE = calledAE
	}
	return observability.Association{
		ConnectionID:  s.connectionID,
		AssociationID: s.associationID,
		CallingAE:     s.callingAE,
		CalledAE:      s.calledAE,
	}
}

func (s *Service) emitEvent(ctx context.Context, event observability.Event) {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	if event.Association.ConnectionID == 0 {
		event.Association = s.observationAssociation()
	}
	observability.EmitEvent(ctx, s.eventObserver, event)
	record := observability.LogRecord{
		Level:   eventLogLevel(event),
		Message: string(event.Kind),
		Event:   event,
	}
	observability.EmitSlog(ctx, "network.service", record)
}

func (s *Service) emitDecodeWarning(ctx context.Context, association observability.Association, warning pdu.DecodeWarning) {
	event := observability.Event{
		Timestamp:   time.Now(),
		Kind:        observability.EventDecodeWarning,
		Association: association,
		Direction:   observability.DirectionInbound,
		Outcome:     observability.OutcomeWarning,
	}
	observability.EmitEvent(ctx, s.eventObserver, event)
	record := observability.LogRecord{
		Level:    observability.LevelWarn,
		Message:  string(event.Kind),
		Code:     string(warning.Code),
		ItemType: warning.ItemType,
		Event:    event,
	}
	observability.EmitSlog(ctx, "network.service", record)
}

func eventLogLevel(event observability.Event) observability.Level {
	if event.Error != nil || event.Outcome == observability.OutcomeFailure {
		return observability.LevelError
	}
	switch event.Kind {
	case observability.EventAssociationRejected,
		observability.EventAssociationAborted,
		observability.EventRequestTimedOut,
		observability.EventRequestCancelled,
		observability.EventDecodeWarning:
		return observability.LevelWarn
	default:
		return observability.LevelInfo
	}
}

func (s *Service) emitMetric(ctx context.Context, metric observability.Metric) {
	if metric.Timestamp.IsZero() {
		metric.Timestamp = time.Now()
	}
	if metric.Association.ConnectionID == 0 {
		metric.Association = s.observationAssociation()
	}
	observability.EmitMetric(ctx, s.metricsObserver, metric)
}

func (s *Service) emitConnectionOpened() {
	event := observability.Event{
		Kind:    observability.EventConnectionOpened,
		Outcome: observability.OutcomeSuccess,
	}
	s.emitEvent(s.ctx, event)
	s.emitMetric(s.ctx, observability.Metric{
		Kind:    observability.MetricConnection,
		Outcome: observability.OutcomeSuccess,
		Value:   1,
	})
}

func (s *Service) emitConnectionClosed(err error) {
	outcome := observability.OutcomeSuccess
	if err != nil {
		outcome = observability.OutcomeFailure
	}
	event := observability.Event{
		Kind:    observability.EventConnectionClosed,
		Outcome: outcome,
		Error:   err,
	}
	s.emitEvent(context.Background(), event)
	s.emitMetric(context.Background(), observability.Metric{
		Kind:    observability.MetricConnection,
		Outcome: outcome,
		Value:   1,
	})
	if err != nil {
		s.emitMetric(context.Background(), observability.Metric{
			Kind:      observability.MetricError,
			Outcome:   observability.OutcomeFailure,
			ErrorKind: observability.ErrorTransport,
			Value:     1,
		})
	}
}

func (s *Service) emitPDUBytes(ctx context.Context, direction observability.Direction, raw *pdu.RawPDU) {
	if raw == nil {
		return
	}
	s.emitMetric(ctx, observability.Metric{
		Kind:      observability.MetricBytes,
		Direction: direction,
		Value:     int64(6 + len(raw.Data)),
	})
}

func (s *Service) emitAssociationObservation(
	ctx context.Context,
	kind observability.EventKind,
	direction observability.Direction,
	outcome observability.Outcome,
	err error,
) {
	s.emitEvent(ctx, observability.Event{
		Kind:      kind,
		Direction: direction,
		Outcome:   outcome,
		Error:     err,
	})
	s.emitMetric(ctx, observability.Metric{
		Kind:      observability.MetricAssociation,
		Direction: direction,
		Outcome:   outcome,
		Value:     1,
	})
	if err != nil {
		s.emitMetric(ctx, observability.Metric{
			Kind:      observability.MetricError,
			Direction: direction,
			Outcome:   observability.OutcomeFailure,
			ErrorKind: observability.ErrorAssociation,
			Value:     1,
		})
	}
}
