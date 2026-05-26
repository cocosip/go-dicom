// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
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
	pendingRequests   map[uint16]*pendingRequest
	pendingRequestsMu sync.RWMutex

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
	message  dimse.Message
	resultCh chan error // Channel to receive send result
}

// pendingRequest tracks a request waiting for a response.
type pendingRequest struct {
	request    dimse.Request
	responseCh chan dimse.Response
	cancelCh   chan struct{}
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
	// Returns multiple responses (Pending + final Success/Failed).
	CFindHandler func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error)

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

	s := &Service{
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
		ctx:                        ctx,
		cancel:                     cancel,
		handlers:                   config.handlers,
		associationNegotiator:      config.associationNegotiator,
		associationReleaseHandler:  config.associationReleaseHandler,
		connectionLifecycleHandler: config.connectionLifecycleHandler,
	}

	return s
}

// SetAssociation sets the association for this service.
// This should be called after association negotiation is complete.
func (s *Service) SetAssociation(assoc *association.Association) {
	s.assocMu.Lock()
	defer s.assocMu.Unlock()
	s.assoc = assoc
}

// GetAssociation returns the current association.
// Returns nil if no association is established.
func (s *Service) GetAssociation() *association.Association {
	s.assocMu.RLock()
	defer s.assocMu.RUnlock()
	return s.assoc
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

// cancelPendingRequests cancels all pending requests.
func (s *Service) cancelPendingRequests() {
	s.pendingRequestsMu.Lock()
	defer s.pendingRequestsMu.Unlock()

	for _, pending := range s.pendingRequests {
		close(pending.cancelCh)
	}
	s.pendingRequests = make(map[uint16]*pendingRequest)
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

		s.cancelPendingRequests()

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
			done := make(chan struct{})
			go func() {
				s.requestWg.Wait()
				close(done)
			}()

			timeout := s.config.dimseTimeout
			if timeout <= 0 {
				timeout = 60 * time.Second
			}
			select {
			case <-done:
				// All handlers completed normally.
			case <-time.After(timeout):
				// Handlers timed out; proceed with shutdown to avoid hanging.
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

// deadlineFromContext calculates a deadline from context and timeout duration.
// If context has a deadline, returns the earlier of context deadline or timeout.
// If context has no deadline, returns time.Now() + timeout.
func deadlineFromContext(ctx context.Context, timeout time.Duration) time.Time {
	if timeout <= 0 {
		return time.Time{} // No deadline
	}

	deadline := time.Now().Add(timeout)
	if ctxDeadline, ok := ctx.Deadline(); ok {
		if ctxDeadline.Before(deadline) {
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
