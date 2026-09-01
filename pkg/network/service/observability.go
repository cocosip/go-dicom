// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/logging"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

type queuedRequestEvent struct {
	ctx   context.Context
	event observability.Event
}

type requestLifecycle struct {
	service                *Service
	association            observability.Association
	direction              observability.Direction
	messageID              uint16
	command                uint16
	started                time.Time
	request                dimse.Request
	commandName            string
	operation              string
	operationID            uint64
	parentOperationID      uint64
	sopClassUID            string
	queryRetrieveLevel     string
	moveDestinationAE      string
	presentationContextID  byte
	subOperationCounts     bool
	remainingSubOperations uint16
	completedSubOperations uint16
	failedSubOperations    uint16
	warningSubOperations   uint16
	resultCount            uint64
	transferSyntax         string

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
		service:               s,
		association:           s.observationAssociation(),
		direction:             direction,
		messageID:             messageID,
		command:               req.CommandField(),
		started:               time.Now(),
		request:               req,
		commandName:           dimse.CommandField(req.CommandField()).String(),
		operation:             dimseOperationName(req.CommandField()),
		operationID:           observability.NewOperationID(),
		sopClassUID:           requestSOPClassUID(req),
		queryRetrieveLevel:    requestQueryRetrieveLevel(req),
		moveDestinationAE:     requestMoveDestination(req),
		presentationContextID: req.PresentationContextID(),
		transferSyntax:        requestTransferSyntax(s, req),
	}
}

type operationContextKey struct{}

func withOperationID(ctx context.Context, operationID uint64) context.Context {
	if operationID == 0 {
		return ctx
	}
	return context.WithValue(ctx, operationContextKey{}, operationID)
}

func operationIDFromContext(ctx context.Context) uint64 {
	if ctx == nil {
		return 0
	}
	operationID, _ := ctx.Value(operationContextKey{}).(uint64)
	return operationID
}

func (l *requestLifecycle) setParentOperationID(operationID uint64) {
	if l == nil || operationID == 0 {
		return
	}
	l.mu.Lock()
	if l.parentOperationID == 0 {
		l.parentOperationID = operationID
	}
	l.mu.Unlock()
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
	if l.request != nil {
		l.presentationContextID = l.request.PresentationContextID()
		l.transferSyntax = requestTransferSyntax(l.service, l.request)
	}
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
	if resp.IsPending() {
		if progress, ok := resp.(interface {
			HasSubOperationCounts() bool
			NumberOfRemainingSubOperations() uint16
			NumberOfCompletedSubOperations() uint16
			NumberOfFailedSubOperations() uint16
			NumberOfWarningSubOperations() uint16
		}); ok && progress.HasSubOperationCounts() {
			l.subOperationCounts = true
			l.remainingSubOperations = progress.NumberOfRemainingSubOperations()
			l.completedSubOperations = progress.NumberOfCompletedSubOperations()
			l.failedSubOperations = progress.NumberOfFailedSubOperations()
			l.warningSubOperations = progress.NumberOfWarningSubOperations()
		}
		if _, ok := resp.(*dimse.CFindResponse); ok && resp.DataDataset() != nil {
			l.resultCount++
		}
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
		Kind:                   kind,
		Association:            l.association,
		Direction:              l.direction,
		MessageID:              l.messageID,
		Command:                l.command,
		CommandName:            l.commandName,
		Operation:              l.operation,
		OperationID:            l.operationID,
		ParentOperationID:      l.parentOperationID,
		SOPClassUID:            l.sopClassUID,
		TransferSyntax:         l.transferSyntax,
		QueryRetrieveLevel:     l.queryRetrieveLevel,
		MoveDestinationAE:      l.moveDestinationAE,
		PresentationContextID:  l.presentationContextID,
		SubOperationCounts:     l.subOperationCounts,
		RemainingSubOperations: l.remainingSubOperations,
		CompletedSubOperations: l.completedSubOperations,
		FailedSubOperations:    l.failedSubOperations,
		WarningSubOperations:   l.warningSubOperations,
		ResultCount:            l.resultCount,
		StatusCode:             statusCode,
		StatusState:            statusState,
		Duration:               time.Since(l.started),
		Outcome:                outcome,
		Error:                  err,
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
			lifecycle := newRequestLifecycle(s, observability.DirectionOutbound, request)
			if cancel, ok := request.(*dimse.CCancelRequest); ok {
				if target := s.pendingRequestOperation(cancel.MessageIDBeingRespondedTo()); target != 0 {
					lifecycle.setParentOperationID(target)
				}
			}
			return lifecycle
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
	if parent := s.parentOperationForInbound(req); parent != 0 {
		lifecycle.setParentOperationID(parent)
	}
	s.inboundRequestsMu.Lock()
	s.inboundRequests[req.MessageID()] = lifecycle
	s.inboundRequestsMu.Unlock()
	return lifecycle
}

func (s *Service) pendingRequestOperation(messageID uint16) uint64 {
	s.pendingRequestsMu.RLock()
	defer s.pendingRequestsMu.RUnlock()
	if pending := s.pendingRequests[messageID]; pending != nil && pending.lifecycle != nil {
		return pending.lifecycle.operationID
	}
	return 0
}

func (s *Service) parentOperationForInbound(req dimse.Request) uint64 {
	if req.CommandField() != uint16(dimse.CommandCStoreRQ) {
		return 0
	}
	s.pendingRequestsMu.RLock()
	defer s.pendingRequestsMu.RUnlock()
	var parent uint64
	for _, pending := range s.pendingRequests {
		if pending == nil || pending.lifecycle == nil || pending.request.CommandField() != uint16(dimse.CommandCGetRQ) {
			continue
		}
		if parent != 0 {
			return 0
		}
		parent = pending.lifecycle.operationID
	}
	return parent
}

func dimseOperationName(command uint16) string {
	name := dimse.CommandField(command).String()
	name = strings.TrimSuffix(name, "-RQ")
	name = strings.TrimSuffix(name, "-RSP")
	return name
}

func requestSOPClassUID(req dimse.Request) string {
	if req == nil {
		return ""
	}
	if uid := req.AffectedSOPClassUID(); uid != "" {
		return uid
	}
	if command := req.CommandDataset(); command != nil {
		uid, _ := command.GetString(tag.RequestedSOPClassUID)
		return uid
	}
	return ""
}

func requestQueryRetrieveLevel(req dimse.Request) string {
	switch request := req.(type) {
	case *dimse.CFindRequest:
		return string(request.QueryLevel())
	case *dimse.CMoveRequest:
		return string(request.QueryLevel())
	case *dimse.CGetRequest:
		return string(request.QueryLevel())
	default:
		return ""
	}
}

func requestMoveDestination(req dimse.Request) string {
	if request, ok := req.(*dimse.CMoveRequest); ok {
		return request.MoveDestination()
	}
	return ""
}

func requestTransferSyntax(s *Service, req dimse.Request) string {
	if s == nil || req == nil {
		return ""
	}
	assoc := s.GetAssociation()
	if assoc == nil {
		return ""
	}
	pc := assoc.FindPresentationContextByID(req.PresentationContextID())
	if pc == nil {
		pc = assoc.FindPresentationContextByAbstractSyntax(requestSOPClassUID(req))
	}
	if pc == nil || pc.AcceptedTransferSyntax == nil || pc.AcceptedTransferSyntax.UID() == nil {
		return ""
	}
	return pc.AcceptedTransferSyntax.UID().UID()
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
	info := observability.Association{
		ConnectionID:  s.connectionID,
		AssociationID: s.associationID,
		CallingAE:     s.callingAE,
		CalledAE:      s.calledAE,
	}
	if s.conn != nil {
		info.LocalAddr, _, _ = endpointMetadata(s.conn.LocalAddr())
		var remoteHost string
		var remotePort int
		info.RemoteAddr, remoteHost, remotePort = endpointMetadata(s.conn.RemoteAddr())
		info.RemoteHost = remoteHost
		info.RemotePort = remotePort
	}
	if assoc := s.assoc; assoc != nil {
		if assoc.RemoteHost != "" {
			info.RemoteHost = assoc.RemoteHost
		}
		if assoc.RemotePort != 0 {
			info.RemotePort = assoc.RemotePort
		}
		info.ImplementationClassUID = assoc.ImplementationClassUID
		info.ImplementationVersionName = assoc.ImplementationVersionName
		info.MaxPDULength = assoc.MaxPDULength
		if assoc.AsynchronousOperations != nil {
			info.AsyncOpsInvoked = assoc.AsynchronousOperations.MaxInvokedOperations
			info.AsyncOpsPerformed = assoc.AsynchronousOperations.MaxPerformedOperations
		}
		info.PresentationContextCount = uint16(len(assoc.PresentationContexts))
		info.PresentationContextSummary = snapshotPresentationContexts(assoc)
		for _, pc := range assoc.PresentationContexts {
			if pc != nil && pc.IsAccepted() {
				info.AcceptedPresentationContextCount++
			}
		}
	}
	return info
}

func endpointMetadata(addr net.Addr) (address, host string, port int) {
	if addr == nil {
		return "", "", 0
	}
	address = addr.String()
	if tcp, ok := addr.(*net.TCPAddr); ok {
		return address, tcp.IP.String(), tcp.Port
	}
	host, portText, err := net.SplitHostPort(address)
	if err != nil {
		return address, "", 0
	}
	port, err = strconv.Atoi(portText)
	if err != nil {
		return address, host, 0
	}
	return address, host, port
}

func (s *Service) associationObservationForRQ(rq *pdu.AAssociateRQ) observability.Association {
	info := s.observationAssociation()
	if rq == nil {
		return info
	}
	info.CallingAE = rq.CallingAETitle
	info.CalledAE = rq.CalledAETitle
	info.PresentationContextCount = uint16(len(rq.PresentationContexts))
	contexts := make([]observability.PresentationContext, 0, len(rq.PresentationContexts))
	for _, pc := range rq.PresentationContexts {
		contexts = append(contexts, observability.PresentationContext{
			ID:               pc.ID,
			AbstractSyntax:   pc.AbstractSyntax,
			TransferSyntaxes: append([]string(nil), pc.TransferSyntaxes...),
		})
	}
	info.PresentationContextSummary = marshalPresentationContexts(contexts)
	applyUserInformation(&info, rq.UserInformation)
	return info
}

func (s *Service) associationObservationForAC(ac *pdu.AAssociateAC) observability.Association {
	info := s.observationAssociation()
	if ac == nil {
		return info
	}
	if ac.CallingAETitle != "" {
		info.CallingAE = ac.CallingAETitle
	}
	if ac.CalledAETitle != "" {
		info.CalledAE = ac.CalledAETitle
	}
	contexts := make([]observability.PresentationContext, 0, len(ac.PresentationContexts))
	for _, pc := range ac.PresentationContexts {
		contexts = append(contexts, observability.PresentationContext{
			ID:               pc.ID,
			TransferSyntaxes: nonEmptyStrings(pc.TransferSyntax),
			Result:           pc.Result,
		})
		if pc.Result == pdu.ResultAcceptance {
			info.AcceptedPresentationContextCount++
		}
	}
	info.PresentationContextCount = uint16(len(ac.PresentationContexts))
	info.PresentationContextSummary = mergeAcceptedPresentationContexts(info.PresentationContextSummary, contexts)
	applyUserInformation(&info, ac.UserInformation)
	return info
}

func mergeAcceptedPresentationContexts(base string, accepted []observability.PresentationContext) string {
	if base == "" {
		return marshalPresentationContexts(accepted)
	}
	var contexts []observability.PresentationContext
	if err := json.Unmarshal([]byte(base), &contexts); err != nil {
		return marshalPresentationContexts(accepted)
	}
	byID := make(map[byte]*observability.PresentationContext, len(contexts))
	for i := range contexts {
		byID[contexts[i].ID] = &contexts[i]
	}
	for _, item := range accepted {
		current := byID[item.ID]
		if current == nil {
			contexts = append(contexts, item)
			continue
		}
		current.Result = item.Result
		current.TransferSyntaxes = append([]string(nil), item.TransferSyntaxes...)
	}
	return marshalPresentationContexts(contexts)
}

func snapshotPresentationContexts(assoc *association.Association) string {
	if assoc == nil || len(assoc.PresentationContexts) == 0 {
		return ""
	}
	contexts := make([]observability.PresentationContext, 0, len(assoc.PresentationContexts))
	for _, pc := range assoc.PresentationContexts {
		if pc == nil {
			continue
		}
		item := observability.PresentationContext{
			ID:             pc.ID,
			AbstractSyntax: pc.AbstractSyntax,
			Result:         pc.Result,
		}
		for _, ts := range pc.ProposedTransferSyntaxes {
			if ts != nil && ts.UID() != nil {
				item.TransferSyntaxes = append(item.TransferSyntaxes, ts.UID().UID())
			}
		}
		if pc.AcceptedTransferSyntax != nil && pc.AcceptedTransferSyntax.UID() != nil {
			item.TransferSyntaxes = []string{pc.AcceptedTransferSyntax.UID().UID()}
		}
		contexts = append(contexts, item)
	}
	return marshalPresentationContexts(contexts)
}

func marshalPresentationContexts(contexts []observability.PresentationContext) string {
	if len(contexts) == 0 {
		return ""
	}
	data, err := json.Marshal(contexts)
	if err != nil {
		return ""
	}
	return string(data)
}

func nonEmptyStrings(value string) []string {
	if value == "" {
		return nil
	}
	return []string{value}
}

func applyUserInformation(info *observability.Association, userInfo *pdu.UserInformation) {
	if info == nil || userInfo == nil {
		return
	}
	info.ImplementationClassUID = userInfo.ImplementationClassUID
	info.ImplementationVersionName = userInfo.ImplementationVersionName
	info.MaxPDULength = userInfo.MaximumLength
	if userInfo.AsynchronousOperations != nil {
		info.AsyncOpsInvoked = userInfo.AsynchronousOperations.MaximumNumberOperationsInvoked
		info.AsyncOpsPerformed = userInfo.AsynchronousOperations.MaximumNumberOperationsPerformed
	}
}

func (s *Service) ensureAssociationIdentity(callingAE, calledAE string) {
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
	s.emitAssociationObservationWithInfo(ctx, kind, direction, outcome, err, s.observationAssociation())
}

func (s *Service) emitAssociationObservationWithInfo(
	ctx context.Context,
	kind observability.EventKind,
	direction observability.Direction,
	outcome observability.Outcome,
	err error,
	associationInfo observability.Association,
) {
	s.emitEvent(ctx, observability.Event{
		Association: associationInfo,
		Kind:        kind,
		Direction:   direction,
		Outcome:     outcome,
		Error:       err,
	})
	s.emitMetric(ctx, observability.Metric{
		Kind:        observability.MetricAssociation,
		Association: associationInfo,
		Direction:   direction,
		Outcome:     outcome,
		Value:       1,
	})
	if err != nil {
		s.emitMetric(ctx, observability.Metric{
			Kind:        observability.MetricError,
			Association: associationInfo,
			Direction:   direction,
			Outcome:     observability.OutcomeFailure,
			ErrorKind:   observability.ErrorAssociation,
			Value:       1,
		})
	}
}
