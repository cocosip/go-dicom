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
	sendQueue chan *sendRequest
	recvQueue chan dimse.Message
	closeOnce sync.Once
	closeCh   chan struct{}
	errCh     chan error

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

	// CMoveHandler handles C-MOVE requests.
	// Returns multiple responses (Pending with progress + final Success/Failed).
	CMoveHandler func(context.Context, *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error)

	// CGetHandler handles C-GET requests.
	// Returns multiple responses (Pending with progress + final Success/Failed).
	CGetHandler func(context.Context, *dimse.CGetRequest) ([]*dimse.CGetResponse, error)

	// Additional handlers for N-* operations can be added here
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
func NewService(conn net.Conn, assoc *association.Association, opts ...ServiceOption) *Service {
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
		recvQueue:                  make(chan dimse.Message, config.recvQueueSize),
		closeCh:                    make(chan struct{}),
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
	var err error
	s.closeOnce.Do(func() {
		// Cancel context to stop goroutines
		s.cancel()

		// Close channels
		close(s.closeCh)

		// Set state to closed
		s.stateMu.Lock()
		s.state = StateClosed
		s.stateMu.Unlock()

		// Cancel all pending requests
		s.cancelPendingRequests()

		// Close connection
		if s.conn != nil {
			err = s.conn.Close()
		}

		// Call OnConnectionClosed callback if set
		s.callbacksMu.RLock()
		lifecycleHandler := s.connectionLifecycleHandler
		s.callbacksMu.RUnlock()

		if lifecycleHandler != nil {
			lifecycleHandler.OnConnectionClosed(s.ctx, err)
		}
	})
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
