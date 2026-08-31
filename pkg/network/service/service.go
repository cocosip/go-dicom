// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
)

// Service represents a DICOM network service.
// It manages the lifecycle of a DICOM association, including:
// - State management
// - PDU send/receive loops
// - DIMSE message encoding/decoding
// - Message routing and handling
//
// Service uses goroutines for concurrent send and receive operations,
// and channels for message queuing.
type Service struct {
	eventObserver   observability.EventObserver
	metricsObserver observability.MetricsObserver
	connectionID    observability.ConnectionID
	associationID   observability.AssociationID
	callingAE       string
	calledAE        string

	// Connection
	conn net.Conn

	// Association information
	assoc   *association.Association
	assocMu sync.RWMutex

	// State management
	state   State
	stateMu sync.RWMutex

	// Goroutine communication
	sendQueue    chan *sendRequest
	closeOnce    sync.Once
	closeCh      chan struct{}
	releaseCh    chan struct{}
	shutdownCh   chan struct{}
	errCh        chan error
	closeErr     error
	closeConnErr error
	closeErrMu   sync.RWMutex

	// Configuration
	config *serviceConfig

	// Message tracking
	pendingRequests       map[uint16]*pendingRequest
	pendingRequestsMu     sync.RWMutex
	activeOperations      map[uint16]*activeOperation
	activeOperationsMu    sync.RWMutex
	inboundRequests       map[uint16]*requestLifecycle
	inboundRequestsMu     sync.RWMutex
	asyncOperationSlots   chan struct{}
	asyncOperationSlotsMu sync.RWMutex

	// DIMSE message handlers (optional, for server mode)
	handlers   *Handlers
	handlersMu sync.RWMutex

	// Lifecycle callbacks (optional)
	associationNegotiator      AssociationNegotiator
	associationReleaseHandler  AssociationReleaseHandler
	connectionLifecycleHandler ConnectionLifecycleHandler
	callbacksMu                sync.RWMutex

	// Context for goroutine lifecycle
	ctx    context.Context
	cancel context.CancelFunc

	// Tracks in-flight request handler goroutines so Close() can wait for them.
	requestWg sync.WaitGroup

	connectionClosedOnce sync.Once
	shutdownOnce         sync.Once
	startMu              sync.Mutex
	started              bool
}

// sendRequest represents a request to send a DIMSE message.
type sendRequest struct {
	message   dimse.Message
	resultCh  chan error // Channel to receive send result
	ctx       context.Context
	lifecycle *requestLifecycle
}

// pendingRequest tracks a request waiting for a response.
type pendingRequest struct {
	request        dimse.Request
	responseCh     chan dimse.Response
	doneCh         chan error
	cancelCh       chan struct{}
	lifecycle      *requestLifecycle
	timeoutTimer   *time.Timer
	timeoutVersion uint64
}

// activeOperation tracks an incoming operation that can be cancelled by C-CANCEL-RQ.
type activeOperation struct {
	token     *struct{}
	cancel    context.CancelFunc
	lifecycle *requestLifecycle
}

// Handlers contains optional DIMSE message handlers for server mode.
// If a handler is nil, the service will send a default response.
//
// Note: Lifecycle callbacks (association negotiation, release, abort, connection close)
// are now provided via separate interfaces:
//   - AssociationNegotiator: for association request/accept callbacks
//   - AssociationReleaseHandler: for release request callbacks
//   - ConnectionLifecycleHandler: for abort and connection close callbacks
type Handlers struct {
	// CEchoHandler handles C-ECHO requests.
	CEchoHandler func(context.Context, *dimse.CEchoRequest) (*dimse.CEchoResponse, error)

	// CStoreHandler handles C-STORE requests.
	CStoreHandler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)

	// CFindHandler handles C-FIND requests.
	//
	// Deprecated: use CFindStreamHandler so results can be sent without first
	// accumulating the complete response set in memory.
	// Returns multiple responses (Pending + final Success/Failed).
	CFindHandler func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error)

	// CFindStreamHandler handles C-FIND requests as a streaming operation.
	// SendPending blocks until its response has entered the service send queue.
	CFindStreamHandler func(context.Context, CFindOperation) error

	// CMoveHandler handles C-MOVE requests via a CMoveOperation interface.
	// The handler calls op.SendPending after each sub-operation completes, enabling
	// per-instance progress streaming analogous to fo-dicom's IAsyncEnumerable pattern.
	// Finish with op.SendSuccess(), op.SendWarning(), or op.SendFailure(code).
	CMoveHandler func(ctx context.Context, op CMoveOperation) error

	// CGetHandler handles C-GET requests via a CGetOperation interface.
	// The handler calls op.SendCStore to push each file back to the SCU over the same
	// association, then op.SendPending to report progress. Finish with op.SendSuccess(),
	// op.SendWarning(), or op.SendFailure(code).
	CGetHandler func(ctx context.Context, op CGetOperation) error

	// NEventReportHandler handles N-EVENT-REPORT requests.
	NEventReportHandler func(context.Context, *dimse.NEventReportRequest) (*dimse.NEventReportResponse, error)

	// NGetHandler handles N-GET requests.
	NGetHandler func(context.Context, *dimse.NGetRequest) (*dimse.NGetResponse, error)

	// NSetHandler handles N-SET requests.
	NSetHandler func(context.Context, *dimse.NSetRequest) (*dimse.NSetResponse, error)

	// NActionHandler handles N-ACTION requests.
	NActionHandler func(context.Context, *dimse.NActionRequest) (*dimse.NActionResponse, error)

	// NCreateHandler handles N-CREATE requests.
	NCreateHandler func(context.Context, *dimse.NCreateRequest) (*dimse.NCreateResponse, error)

	// NDeleteHandler handles N-DELETE requests.
	NDeleteHandler func(context.Context, *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error)
}

// NewService creates a new DICOM service with the given connection and options.
// The connection should already be established (e.g., via transport.Dial).
// The association parameter can be nil initially and set later via SetAssociation.
//
// Example:
//
//	conn, err := transport.Dial(ctx, "tcp", "192.168.1.100:104")
//	if err != nil {
//	    return err
//	}
//
//	service := service.NewService(conn, nil,
//	    service.WithMaxPDULength(32768),
//	    service.WithReadTimeout(60*time.Second))
//	defer service.Close()
func NewService(conn net.Conn, assoc *association.Association, opts ...Option) *Service {
	// Apply default config
	config := defaultServiceConfig()

	// Apply custom options
	for _, opt := range opts {
		opt(config)
	}

	ctx, cancel := context.WithCancel(context.Background())
	connectionID := config.connectionID
	if connectionID == 0 {
		connectionID = observability.NewConnectionID()
	}
	var associationID observability.AssociationID
	var callingAE, calledAE string
	if assoc != nil {
		associationID = observability.NewAssociationID()
		callingAE = assoc.CallingAE
		calledAE = assoc.CalledAE
	}

	s := &Service{
		eventObserver:              config.eventObserver,
		metricsObserver:            config.metricsObserver,
		connectionID:               connectionID,
		associationID:              associationID,
		callingAE:                  callingAE,
		calledAE:                   calledAE,
		conn:                       conn,
		assoc:                      assoc,
		state:                      StateIdle,
		sendQueue:                  make(chan *sendRequest, config.sendQueueSize),
		closeCh:                    make(chan struct{}),
		releaseCh:                  make(chan struct{}, 1),
		shutdownCh:                 make(chan struct{}),
		errCh:                      make(chan error, 1),
		config:                     config,
		pendingRequests:            make(map[uint16]*pendingRequest),
		activeOperations:           make(map[uint16]*activeOperation),
		inboundRequests:            make(map[uint16]*requestLifecycle),
		ctx:                        ctx,
		cancel:                     cancel,
		handlers:                   config.handlers,
		associationNegotiator:      config.associationNegotiator,
		associationReleaseHandler:  config.associationReleaseHandler,
		connectionLifecycleHandler: config.connectionLifecycleHandler,
	}
	s.configureAsyncOperationSlots(assoc)
	s.emitConnectionOpened()

	return s
}

// SetAssociation sets the association for this service.
// This should be called after association negotiation is complete.
func (s *Service) SetAssociation(assoc *association.Association) {
	s.assocMu.Lock()
	s.assoc = assoc
	if assoc != nil {
		if s.associationID == 0 {
			s.associationID = observability.NewAssociationID()
		}
		s.callingAE = assoc.CallingAE
		s.calledAE = assoc.CalledAE
	}
	s.assocMu.Unlock()
	s.configureAsyncOperationSlots(assoc)
}

// GetAssociation returns the current association.
// Returns nil if no association is established.
func (s *Service) GetAssociation() *association.Association {
	s.assocMu.RLock()
	defer s.assocMu.RUnlock()
	return s.assoc
}

func (s *Service) configureAsyncOperationSlots(assoc *association.Association) {
	maxInvoked := uint16(1)
	if assoc != nil && assoc.AsynchronousOperations != nil {
		maxInvoked = assoc.AsynchronousOperations.MaxInvokedOperations
	}

	var slots chan struct{}
	if maxInvoked > 0 {
		slots = make(chan struct{}, maxInvoked)
	}
	s.asyncOperationSlotsMu.Lock()
	s.asyncOperationSlots = slots
	s.asyncOperationSlotsMu.Unlock()
}

func (s *Service) acquireAsyncOperation(ctx context.Context) (func(), error) {
	s.asyncOperationSlotsMu.RLock()
	slots := s.asyncOperationSlots
	s.asyncOperationSlotsMu.RUnlock()
	if slots == nil {
		return func() {}, nil
	}

	select {
	case slots <- struct{}{}:
		var once sync.Once
		return func() {
			once.Do(func() { <-slots })
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closeCh:
		return nil, s.CloseError()
	}
}

func (s *Service) requireLocalRole(pc *association.PresentationContext, scu bool) error {
	localSCU, localSCP := localRoles(pc, s.config.associationRequestor)
	if (scu && localSCU) || (!scu && localSCP) {
		return nil
	}
	role := "SCP"
	if scu {
		role = "SCU"
	}
	return fmt.Errorf("local AE did not negotiate the %s role for presentation context %d (SOP Class %s)",
		role, pc.ID, pc.AbstractSyntax)
}

func localRoles(pc *association.PresentationContext, associationRequestor bool) (scu, scp bool) {
	if pc.AcceptedRole == nil {
		return associationRequestor, !associationRequestor
	}
	if associationRequestor {
		return pc.AcceptedRole.SCURole == 1, pc.AcceptedRole.SCPRole == 1
	}
	return pc.AcceptedRole.SCPRole == 1, pc.AcceptedRole.SCURole == 1
}

func presentationContextForRequest(assoc *association.Association, req dimse.Request) (*association.PresentationContext, error) {
	if contextID := req.PresentationContextID(); contextID != 0 {
		pc := assoc.FindPresentationContextByID(contextID)
		if pc == nil || !pc.IsAccepted() {
			return nil, fmt.Errorf("presentation context ID %d is not accepted", contextID)
		}
		return pc, nil
	}

	sopClassUID := req.AffectedSOPClassUID()
	if sopClassUID == "" && req.CommandDataset() != nil {
		sopClassUID, _ = req.CommandDataset().GetString(tag.RequestedSOPClassUID)
	}
	if sopClassUID == "" {
		return nil, fmt.Errorf("request has no presentation context ID or SOP Class UID")
	}
	pc := assoc.FindPresentationContextByAbstractSyntax(sopClassUID)
	if pc == nil {
		return nil, fmt.Errorf("no accepted presentation context found for SOP Class UID: %s", sopClassUID)
	}
	return pc, nil
}

// GetState returns the current state of the service.
func (s *Service) GetState() State {
	s.stateMu.RLock()
	defer s.stateMu.RUnlock()
	return s.state
}

// setState sets the state of the service.
// Returns an error if the transition is invalid.
func (s *Service) setState(newState State) error {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()

	if !IsValidTransition(s.state, newState) {
		return fmt.Errorf("%w: cannot transition from %s to %s",
			ErrInvalidState, s.state, newState)
	}

	s.state = newState
	return nil
}

// Close closes the service and releases all resources.
// This will:
// - Cancel all goroutines
// - Close the network connection
// - Cancel all pending requests
// - Drain all queues
func (s *Service) Close() error {
	err := s.initiateClose(StateClosed, nil)
	s.waitForShutdown()
	return err
}

// cancelPendingRequests cancels all pending requests and returns their
// lifecycles so observers can be notified after the map lock is released.
func (s *Service) cancelPendingRequests() []*requestLifecycle {
	s.pendingRequestsMu.Lock()
	lifecycles := make([]*requestLifecycle, 0, len(s.pendingRequests))
	for _, pending := range s.pendingRequests {
		if pending.timeoutTimer != nil {
			pending.timeoutTimer.Stop()
		}
		close(pending.cancelCh)
		lifecycles = append(lifecycles, pending.lifecycle)
	}
	s.pendingRequests = make(map[uint16]*pendingRequest)
	s.pendingRequestsMu.Unlock()
	return lifecycles
}

func (s *Service) takeInboundRequestLifecycles() []*requestLifecycle {
	s.inboundRequestsMu.Lock()
	lifecycles := make([]*requestLifecycle, 0, len(s.inboundRequests))
	for _, lifecycle := range s.inboundRequests {
		lifecycles = append(lifecycles, lifecycle)
	}
	s.inboundRequests = make(map[uint16]*requestLifecycle)
	s.inboundRequestsMu.Unlock()
	return lifecycles
}

// cancelActiveOperations cancels all incoming operations that can receive C-CANCEL-RQ.
func (s *Service) cancelActiveOperations() {
	s.activeOperationsMu.Lock()
	defer s.activeOperationsMu.Unlock()

	for _, op := range s.activeOperations {
		op.cancel()
	}
	s.activeOperations = make(map[uint16]*activeOperation)
}

func (s *Service) initiateClose(targetState State, recordErr error) error {
	if recordErr != nil {
		s.setCloseError(recordErr)
	}

	var err error
	s.closeOnce.Do(func() {
		s.cancel()
		close(s.closeCh)

		s.stateMu.Lock()
		s.state = targetState
		s.stateMu.Unlock()

		terminalErr := recordErr
		if terminalErr == nil {
			terminalErr = ErrServiceClosed
		}
		requestLifecycles := s.cancelPendingRequests()
		s.cancelActiveOperations()
		requestLifecycles = append(requestLifecycles, s.takeInboundRequestLifecycles()...)
		for _, lifecycle := range requestLifecycles {
			lifecycle.finishError(context.Background(), terminalErr)
		}

		if s.conn != nil {
			err = s.conn.Close()
			s.setCloseConnError(err)
		}

		s.startShutdownFinalizer()
	})

	return err
}

func (s *Service) startShutdownFinalizer() {
	s.shutdownOnce.Do(func() {
		go func() {
			// Wait for in-flight request handlers with a timeout to prevent
			// permanent shutdown blocking when a handler never returns.
			//
			// requestWg.Wait() is called in a separate goroutine so that a
			// handler calling Close() does not self-deadlock: the handler
			// goroutine is counted in requestWg, so if Wait() blocked the
			// finalizer goroutine directly, the handler would be stuck in
			// waitForShutdown() and never reach requestWg.Done().
			done := make(chan struct{})
			go func() {
				s.requestWg.Wait()
				close(done)
			}()

			timeout := s.config.handlerShutdownTimeout
			if timeout <= 0 {
				timeout = 60 * time.Second
			}
			select {
			case <-done:
				// All handlers completed normally.
			case <-time.After(timeout):
				// Handlers timed out; proceed with shutdown to avoid hanging.
				// This also breaks the self-deadlock when a handler calls
				// Close() — the handler goroutine can't reach Done() while
				// blocked in waitForShutdown, so the timeout unblocks it.
				s.setCloseError(ErrHandlerShutdownTimeout)
			}

			s.notifyConnectionClosed(s.shutdownError())
			close(s.shutdownCh)
		}()
	})
}

func (s *Service) waitForShutdown() {
	<-s.shutdownCh
}

func (s *Service) notifyConnectionClosed(err error) {
	s.connectionClosedOnce.Do(func() {
		s.emitConnectionClosed(err)

		s.callbacksMu.RLock()
		lifecycleHandler := s.connectionLifecycleHandler
		s.callbacksMu.RUnlock()

		if lifecycleHandler != nil {
			lifecycleHandler.OnConnectionClosed(s.ctx, err)
		}
	})
}

// setCloseError stores the error that caused the service to close.
// Must be called before closing closeCh so readers always see it.
func (s *Service) setCloseError(err error) {
	s.closeErrMu.Lock()
	defer s.closeErrMu.Unlock()
	if s.closeErr == nil {
		s.closeErr = err
	}
}

func (s *Service) setCloseConnError(err error) {
	s.closeErrMu.Lock()
	defer s.closeErrMu.Unlock()
	if s.closeConnErr == nil {
		s.closeConnErr = err
	}
}

func (s *Service) shutdownError() error {
	s.closeErrMu.RLock()
	defer s.closeErrMu.RUnlock()
	if s.closeErr != nil {
		return s.closeErr
	}
	return s.closeConnErr
}

// CloseError returns the error that caused the service to close, or nil if it
// closed normally. Returns ErrServiceClosed when the service is closed but no
// specific error was recorded.
func (s *Service) CloseError() error {
	s.closeErrMu.RLock()
	defer s.closeErrMu.RUnlock()
	if s.closeErr != nil {
		return s.closeErr
	}
	if s.IsClosed() {
		return ErrServiceClosed
	}
	return nil
}

// IsClosed returns true if the service is closed.
func (s *Service) IsClosed() bool {
	select {
	case <-s.closeCh:
		return true
	default:
		return false
	}
}

// Context returns the service's context.
// This context is cancelled when the service is closed.
func (s *Service) Context() context.Context {
	return s.ctx
}

// deadlineFromContext calculates the earlier deadline supplied by the context
// or configured timeout. A zero configured timeout disables only the configured
// deadline; it does not discard a deadline carried by the context.
func deadlineFromContext(ctx context.Context, timeout time.Duration) time.Time {
	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(timeout)
	}
	if ctx != nil {
		if ctxDeadline, ok := ctx.Deadline(); ok && (deadline.IsZero() || ctxDeadline.Before(deadline)) {
			return ctxDeadline
		}
	}
	return deadline
}

// readTimeoutFromContext returns the effective read timeout for a context and configured duration.
// If the context carries a deadline that is sooner than configured, the remaining time until that
// deadline is returned. Otherwise the configured duration is returned.
//
// Edge cases:
//   - configured == 0 but context has deadline → use context deadline
//   - Context already expired → return 1ns so ReadPDU arms the deadline and the read fails fast
//   - Both configured and context deadline → use the earlier one
func readTimeoutFromContext(ctx context.Context, configured time.Duration) time.Duration {
	if ctxDeadline, ok := ctx.Deadline(); ok {
		remaining := time.Until(ctxDeadline)
		if remaining <= 0 {
			// Context already expired; return a minimal positive duration so ReadPDU
			// arms the deadline and the subsequent read fails immediately.
			return time.Nanosecond
		}
		if configured <= 0 || remaining < configured {
			return remaining
		}
	}
	return configured
}
