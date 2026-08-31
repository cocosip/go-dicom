// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package server provides DICOM SCP (Service Class Provider) server functionality.
// It handles incoming DICOM connections, negotiates associations, and processes DIMSE requests.
package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

//var _ service.AssociationResponder = (*associationReleaseAdapter)(nil)

// Server represents a DICOM SCP (Service Class Provider) server.
// It listens for incoming DICOM connections and handles DIMSE requests.
//
// Example usage:
//
//	server := server.New(
//	    server.WithPort(104),
//	)
//
//	// Set handlers for DIMSE operations
//	server.SetCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
//	    return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
//	})
//
//	// Set custom association negotiator (optional, for AE validation, etc.)
//	server.SetAssociationNegotiatorFunc(func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
//	    // Validate AE titles, presentation contexts, etc.
//	    // For example, check CallingAE from database
//	    for _, pc := range assoc.PresentationContexts {
//	        if len(pc.ProposedTransferSyntaxes) > 0 {
//	            pc.Accept(pc.ProposedTransferSyntaxes[0])
//	        }
//	    }
//	    return responder.SendAccept(ctx, assoc)
//	})
//
//	// Start the server
//	if err := server.ListenAndServe(ctx); err != nil {
//	    log.Fatal(err)
//	}
type Server struct {
	// Configuration
	config *Config

	// Network listener
	listener serverListener
	listenFn listenerFactory

	// Active connections
	connections   map[string]*serverConnection
	connectionsMu sync.RWMutex

	// Service options to apply to each connection
	serviceOptions []service.Option
	optionsMu      sync.RWMutex

	// Server state
	running    bool
	starting   bool
	runningMu  sync.RWMutex
	startDone  chan struct{}
	acceptDone chan struct{}

	// Context for server lifecycle
	ctx    context.Context
	cancel context.CancelFunc

	// Wait group for active connections
	wg sync.WaitGroup
}

type serverListener interface {
	Accept(context.Context) (net.Conn, error)
	Close() error
}

type listenerFactory func(string, string, ...transport.ListenOption) (serverListener, error)

func defaultListenerFactory(network, address string, opts ...transport.ListenOption) (serverListener, error) {
	return transport.Listen(network, address, opts...)
}

// Config contains configuration options for the DICOM server.
type Config struct {
	// EventObserver receives connection, association, and DIMSE lifecycle events.
	EventObserver observability.EventObserver

	// MetricsObserver receives vendor-neutral network metrics.
	MetricsObserver observability.MetricsObserver

	// Port is the TCP port to listen on
	// Default: 104 (standard DICOM port)
	Port int

	// MaxPDULength is the maximum PDU length to accept
	// Default: 16384 bytes
	MaxPDULength uint32

	// AcceptTimeout is the timeout for accepting new connections
	// Default: no timeout
	AcceptTimeout time.Duration

	// AssociationTimeout is the timeout for association negotiation.
	// It does not apply after the association is established.
	// Default: 10 seconds
	AssociationTimeout time.Duration

	// RequestTimeout is the default response idle timeout for outgoing
	// DIMSE requests initiated by the server-side service.
	// Default: 30 seconds
	RequestTimeout time.Duration

	// TransportReadTimeout limits a single PDU read. A zero value permits an
	// established association to remain idle until it is released or closed.
	// Default: 0 (disabled)
	TransportReadTimeout time.Duration

	// TransportWriteTimeout limits a single PDU write.
	// Default: 30 seconds
	TransportWriteTimeout time.Duration

	// ImplementationClassUID identifies the implementation
	// Default: "1.2.826.0.1.3680043.10.854"
	ImplementationClassUID string

	// ImplementationVersionName identifies the implementation version
	// Default: "GO-DICOM-1.0"
	ImplementationVersionName string

	// MaxConnections limits the number of concurrent connections
	// 0 means no limit
	// Default: 0 (no limit)
	MaxConnections int

	// TLSConfig for secure connections (optional)
	TLSConfig *tls.Config
}

// Option is a function that modifies server configuration.
type Option func(*Config)

// WithEventObserver sets the network lifecycle event observer.
func WithEventObserver(observer observability.EventObserver) Option {
	return func(o *Config) { o.EventObserver = observer }
}

// WithMetricsObserver sets the vendor-neutral network metrics observer.
func WithMetricsObserver(observer observability.MetricsObserver) Option {
	return func(o *Config) { o.MetricsObserver = observer }
}

// WithPort sets the listening port.
func WithPort(port int) Option {
	return func(o *Config) {
		o.Port = port
	}
}

// WithMaxPDULength sets the maximum PDU length.
func WithMaxPDULength(length uint32) Option {
	return func(o *Config) {
		o.MaxPDULength = length
	}
}

// WithAcceptTimeout sets the accept timeout.
func WithAcceptTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.AcceptTimeout = timeout
	}
}

// WithAssociationTimeout sets the association timeout.
func WithAssociationTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.AssociationTimeout = timeout
	}
}

// WithRequestTimeout sets the default response idle timeout for outgoing
// DIMSE requests initiated by the server-side service.
func WithRequestTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.RequestTimeout = timeout
	}
}

// WithTransportReadTimeout sets the timeout for each transport PDU read.
// It does not affect association negotiation or an individual DIMSE request.
func WithTransportReadTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.TransportReadTimeout = timeout
	}
}

// WithTransportWriteTimeout sets the timeout for each transport PDU write.
// It does not affect association negotiation or an individual DIMSE request.
func WithTransportWriteTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.TransportWriteTimeout = timeout
	}
}

// WithImplementationClassUID sets the implementation class UID.
func WithImplementationClassUID(uid string) Option {
	return func(o *Config) {
		o.ImplementationClassUID = uid
	}
}

// WithImplementationVersionName sets the implementation version name.
func WithImplementationVersionName(name string) Option {
	return func(o *Config) {
		o.ImplementationVersionName = name
	}
}

// WithMaxConnections sets the maximum number of concurrent connections.
func WithMaxConnections(maxConnections int) Option {
	return func(o *Config) {
		o.MaxConnections = maxConnections
	}
}

// WithTLSConfig sets the TLS configuration for secure connections.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(o *Config) {
		o.TLSConfig = tlsConfig
	}
}

// defaultServerConfig returns the default server configuration.
func defaultServerConfig() *Config {
	return &Config{
		Port:                      104,
		MaxPDULength:              16384,
		AcceptTimeout:             0, // No timeout
		AssociationTimeout:        10 * time.Second,
		RequestTimeout:            30 * time.Second,
		TransportReadTimeout:      0,
		TransportWriteTimeout:     30 * time.Second,
		ImplementationClassUID:    "1.2.826.0.1.3680043.10.854",
		ImplementationVersionName: "GO-DICOM-1.0",
		MaxConnections:            0, // No limit
		TLSConfig:                 nil,
	}
}

// New creates a new DICOM server with the specified options.
func New(opts ...Option) *Server {
	config := defaultServerConfig()
	for _, opt := range opts {
		opt(config)
	}

	return &Server{
		config:         config,
		connections:    make(map[string]*serverConnection),
		serviceOptions: make([]service.Option, 0),
		listenFn:       defaultListenerFactory,
	}
}

// serverConnection represents an active client connection.
type serverConnection struct {
	conn    net.Conn
	service *service.Service
	assoc   *association.Association
	id      string
}

// GetConfig returns the server configuration.
func (s *Server) GetConfig() *Config {
	return s.config
}

// SetCEchoHandler sets the C-ECHO request handler.
func (s *Server) SetCEchoHandler(handler func(context.Context, *dimse.CEchoRequest) (*dimse.CEchoResponse, error)) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCEchoHandler(handler))
}

// SetCStoreHandler sets the C-STORE request handler.
func (s *Server) SetCStoreHandler(handler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCStoreHandler(handler))
}

// SetCFindHandler sets the legacy C-FIND request handler.
//
// Deprecated: use SetCFindStreamHandler for large result sets.
func (s *Server) SetCFindHandler(handler func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error)) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCFindHandler(handler))
}

// SetCFindStreamHandler sets a C-FIND handler that sends results as they
// become available without first constructing a complete response slice.
func (s *Server) SetCFindStreamHandler(handler func(context.Context, service.CFindOperation) error) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCFindStreamHandler(handler))
}

// SetCMoveHandler sets the C-MOVE handler via a CMoveOperation interface.
// The handler calls op.SendPending after each sub-operation and finishes with
// op.SendSuccess(), op.SendWarning(), or op.SendFailure(code).
func (s *Server) SetCMoveHandler(handler func(ctx context.Context, op service.CMoveOperation) error) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCMoveHandler(handler))
}

// SetCGetHandler sets the C-GET handler via a CGetOperation interface.
// The handler calls op.SendCStore for each file, op.SendPending after each result,
// and finishes with op.SendSuccess(), op.SendWarning(), or op.SendFailure(code).
func (s *Server) SetCGetHandler(handler func(ctx context.Context, op service.CGetOperation) error) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCGetHandler(handler))
}

// SetAssociationNegotiator sets the association negotiator for this server.
// The negotiator controls which associations are accepted.
func (s *Server) SetAssociationNegotiator(negotiator service.AssociationNegotiator) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	// Append new option (later options override earlier ones in Service.NewService)
	s.serviceOptions = append(s.serviceOptions, service.WithAssociationNegotiator(negotiator))
}

// SetAssociationReleaseHandler sets the association release handler for this server.
func (s *Server) SetAssociationReleaseHandler(handler service.AssociationReleaseHandler) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithAssociationReleaseHandler(handler))
}

// SetConnectionLifecycleHandler sets the connection lifecycle handler for this server.
func (s *Server) SetConnectionLifecycleHandler(handler service.ConnectionLifecycleHandler) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithConnectionLifecycleHandler(handler))
}

// SetAssociationNegotiatorFunc is a convenience method that accepts a function and wraps it
// as an AssociationNegotiator. This allows using a simple function instead of implementing the interface.
//
// Example:
//
//	server.SetAssociationNegotiatorFunc(func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
//	    // Validate and negotiate
//	    for _, pc := range assoc.PresentationContexts {
//	        if len(pc.ProposedTransferSyntaxes) > 0 {
//	            pc.Accept(pc.ProposedTransferSyntaxes[0])
//	        }
//	    }
//	    return responder.SendAccept(ctx, assoc)
//	})
func (s *Server) SetAssociationNegotiatorFunc(fn func(context.Context, *association.Association, service.AssociationResponder) error) {
	s.SetAssociationNegotiator(service.FuncAssociationNegotiator(fn))
}

// SetAssociationReleaseHandlerFunc is a convenience method that accepts a function and wraps it
// as an AssociationReleaseHandler.
//
// Example:
//
//	server.SetAssociationReleaseHandlerFunc(func(ctx context.Context) error {
//	    // Perform cleanup
//	    return nil // Accept release
//	})
func (s *Server) SetAssociationReleaseHandlerFunc(fn func(context.Context) error) {
	s.SetAssociationReleaseHandler(service.FuncAssociationReleaseHandler(fn))
}

// SetConnectionLifecycleHandlerFuncs is a convenience method that accepts individual functions
// for connection lifecycle events.
//
// Example:
//
//	server.SetConnectionLifecycleHandlerFuncs(
//	    func(ctx context.Context, source, reason byte) {
//	        log.Printf("Aborted: source=%d, reason=%d", source, reason)
//	    },
//	    func(ctx context.Context, err error) {
//	        log.Printf("Closed: %v", err)
//	    },
//	)
func (s *Server) SetConnectionLifecycleHandlerFuncs(onAbort func(context.Context, byte, byte), onConnectionClosed func(context.Context, error)) {
	handler := &service.ConnectionLifecycleHandlerFuncs{
		OnAbortFunc:            onAbort,
		OnConnectionClosedFunc: onConnectionClosed,
	}
	s.SetConnectionLifecycleHandler(handler)
}

// IsRunning returns true if the server is running.
func (s *Server) IsRunning() bool {
	s.runningMu.RLock()
	defer s.runningMu.RUnlock()
	return s.running
}

// ActiveConnections returns the number of active connections.
func (s *Server) ActiveConnections() int {
	s.connectionsMu.RLock()
	defer s.connectionsMu.RUnlock()
	return len(s.connections)
}

// buildUserInformation builds the UserInformation structure for A-ASSOCIATE-AC.
func (s *Server) buildUserInformation() *pdu.UserInformation {
	return &pdu.UserInformation{
		MaximumLength:             s.config.MaxPDULength,
		ImplementationClassUID:    s.config.ImplementationClassUID,
		ImplementationVersionName: s.config.ImplementationVersionName,
	}
}

// ListenAndServe starts the server and blocks until it's stopped or an error occurs.
func (s *Server) ListenAndServe(ctx context.Context) error {
	s.runningMu.Lock()
	if s.running || s.starting {
		s.runningMu.Unlock()
		return fmt.Errorf("server is already running")
	}
	startDone := make(chan struct{})
	s.starting = true
	s.startDone = startDone
	s.runningMu.Unlock()

	runCtx, cancel := context.WithCancel(ctx)
	address := fmt.Sprintf(":%d", s.config.Port)
	listenOptions := make([]transport.ListenOption, 0, 1)
	if s.config.TLSConfig != nil {
		listenOptions = append(listenOptions, transport.WithListenTLSConfig(s.config.TLSConfig))
	}
	listenFn := s.listenFn
	if listenFn == nil {
		listenFn = defaultListenerFactory
	}
	listener, err := listenFn("tcp", address, listenOptions...)
	if err != nil {
		cancel()
		s.runningMu.Lock()
		if s.startDone == startDone {
			s.starting = false
			s.startDone = nil
			close(startDone)
		}
		s.runningMu.Unlock()
		s.emitServerError(runCtx, "listener_failed", err)
		return fmt.Errorf("failed to start listener: %w", err)
	}
	acceptDone := make(chan struct{})
	s.runningMu.Lock()
	s.ctx = runCtx
	s.cancel = cancel
	s.listener = listener
	s.acceptDone = acceptDone
	s.running = true
	s.starting = false
	s.startDone = nil
	close(startDone)
	s.runningMu.Unlock()

	err = s.acceptLoop(runCtx, listener)
	close(acceptDone)
	s.wg.Wait()
	_ = listener.Close()
	cancel()

	s.runningMu.Lock()
	if s.acceptDone == acceptDone {
		s.running = false
		s.ctx = nil
		s.cancel = nil
		s.listener = nil
		s.acceptDone = nil
	}
	s.runningMu.Unlock()
	return err
}

func (s *Server) acceptLoop(ctx context.Context, listener serverListener) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if s.config.MaxConnections > 0 && s.ActiveConnections() >= s.config.MaxConnections {
			select {
			case <-time.After(100 * time.Millisecond):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}

		acceptCtx, cancelAccept := s.acceptContext(ctx)
		conn, err := listener.Accept(acceptCtx)
		cancelAccept()
		if err != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				s.emitServerError(ctx, "accept_failed", err)
				continue
			}
		}
		if err := ctx.Err(); err != nil {
			_ = conn.Close()
			return err
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *Server) acceptContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if s.config.AcceptTimeout <= 0 {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, s.config.AcceptTimeout)
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.waitForStartup(ctx); err != nil {
		return err
	}
	acceptDone, stopErr := s.stopAccepting()
	if stopErr != nil {
		return stopErr
	}
	if acceptDone != nil {
		select {
		case <-acceptDone:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) waitForStartup(ctx context.Context) error {
	for {
		s.runningMu.RLock()
		starting := s.starting
		startDone := s.startDone
		s.runningMu.RUnlock()
		if !starting || startDone == nil {
			return nil
		}
		select {
		case <-startDone:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Close immediately closes the listener and all active inbound connections.
// It does not send an A-RELEASE-RQ or wait for active handlers to return.
// Close may be called after Shutdown returns because its context expired.
func (s *Server) Close() error {
	var errs []error
	if _, err := s.stopAccepting(); err != nil {
		errs = append(errs, err)
	}

	s.connectionsMu.RLock()
	connections := make([]*serverConnection, 0, len(s.connections))
	for _, connection := range s.connections {
		connections = append(connections, connection)
	}
	s.connectionsMu.RUnlock()

	for _, connection := range connections {
		if err := connection.conn.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (s *Server) stopAccepting() (<-chan struct{}, error) {
	s.runningMu.RLock()
	if !s.running {
		s.runningMu.RUnlock()
		return nil, nil
	}
	cancel := s.cancel
	listener := s.listener
	acceptDone := s.acceptDone
	s.runningMu.RUnlock()

	if cancel != nil {
		cancel()
	}
	if listener == nil {
		return acceptDone, nil
	}
	if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
		return acceptDone, err
	}
	return acceptDone, nil
}

// handleConnection handles a single client connection.
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer func() { _ = conn.Close() }()

	// Create connection ID
	connID := conn.RemoteAddr().String()
	observationConnectionID := observability.NewConnectionID()

	// Build service options from server configuration.
	svcOpts := s.serviceOptionsForConnection(observationConnectionID)

	// Add all configured service options (handlers, callbacks, etc.)
	s.optionsMu.RLock()
	svcOpts = append(svcOpts, s.serviceOptions...)
	s.optionsMu.RUnlock()

	// Create service with all options
	svc := service.NewService(conn, nil, svcOpts...)

	// Receive A-ASSOCIATE-RQ and handle negotiation
	// The service will call the AssociationNegotiator callback (if set by user), which will:
	// - Validate the association (AE titles, presentation contexts, etc.)
	// - Send A-ASSOCIATE-AC or A-ASSOCIATE-RJ
	// - Set the association on the service if accepted
	// If no negotiator is set, the service uses DefaultAssociationNegotiator which accepts
	// all presentation contexts with their first proposed transfer syntax.
	assocCtx, cancel := context.WithTimeout(s.ctx, s.config.AssociationTimeout)
	defer cancel()

	_, err := svc.ReceiveAssociationRequest(assocCtx)
	if err != nil {
		// Association was rejected or failed
		return
	}

	// Get the negotiated association from the service
	// (it was set by the responder.SendAccept call in the negotiator)
	assoc := svc.GetAssociation()
	if assoc == nil {
		// This shouldn't happen if ReceiveAssociationRequest succeeded
		return
	}

	// Register connection
	serverConn := &serverConnection{
		conn:    conn,
		service: svc,
		assoc:   assoc,
		id:      connID,
	}

	s.connectionsMu.Lock()
	if s.ctx != nil && s.ctx.Err() != nil {
		s.connectionsMu.Unlock()
		return
	}
	s.connections[connID] = serverConn
	s.connectionsMu.Unlock()

	defer func() {
		s.connectionsMu.Lock()
		delete(s.connections, connID)
		s.connectionsMu.Unlock()
	}()

	// Start service loops
	if err := svc.Start(); err != nil {
		return
	}

	// Wait for service to finish
	// The service will handle incoming DIMSE requests automatically
	// and call the registered handlers
	<-svc.Context().Done()
}

func (s *Server) serviceOptionsForConnection(connectionID observability.ConnectionID) []service.Option {
	return []service.Option{
		service.WithConnectionID(connectionID),
		service.WithEventObserver(s.config.EventObserver),
		service.WithMetricsObserver(s.config.MetricsObserver),
		service.WithMaxPDULength(s.config.MaxPDULength),
		service.WithRequestTimeout(s.config.RequestTimeout),
		service.WithReadTimeout(s.config.TransportReadTimeout),
		service.WithWriteTimeout(s.config.TransportWriteTimeout),
	}
}
