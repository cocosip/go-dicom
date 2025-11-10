// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
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
	config *ServerConfig

	// Network listener
	listener *transport.Listener

	// Active connections
	connections   map[string]*serverConnection
	connectionsMu sync.RWMutex

	// Service options to apply to each connection
	serviceOptions []service.ServiceOption
	optionsMu      sync.RWMutex

	// Server state
	running   bool
	runningMu sync.RWMutex

	// Context for server lifecycle
	ctx    context.Context
	cancel context.CancelFunc

	// Wait group for active connections
	wg sync.WaitGroup
}

// ServerConfig contains configuration options for the DICOM server.
type ServerConfig struct {
	// Port is the TCP port to listen on
	// Default: 104 (standard DICOM port)
	Port int

	// MaxPDULength is the maximum PDU length to accept
	// Default: 16384 bytes
	MaxPDULength uint32

	// AcceptTimeout is the timeout for accepting new connections
	// Default: no timeout
	AcceptTimeout time.Duration

	// AssociationTimeout is the timeout for association negotiation
	// Default: 10 seconds
	AssociationTimeout time.Duration

	// RequestTimeout is the default timeout for processing DIMSE requests
	// Default: 30 seconds
	RequestTimeout time.Duration

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

// ServerOption is a function that modifies server configuration.
type ServerOption func(*ServerConfig)

// WithPort sets the listening port.
func WithPort(port int) ServerOption {
	return func(o *ServerConfig) {
		o.Port = port
	}
}

// WithMaxPDULength sets the maximum PDU length.
func WithMaxPDULength(length uint32) ServerOption {
	return func(o *ServerConfig) {
		o.MaxPDULength = length
	}
}

// WithAcceptTimeout sets the accept timeout.
func WithAcceptTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerConfig) {
		o.AcceptTimeout = timeout
	}
}

// WithAssociationTimeout sets the association timeout.
func WithAssociationTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerConfig) {
		o.AssociationTimeout = timeout
	}
}

// WithRequestTimeout sets the request timeout.
func WithRequestTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerConfig) {
		o.RequestTimeout = timeout
	}
}

// WithImplementationClassUID sets the implementation class UID.
func WithImplementationClassUID(uid string) ServerOption {
	return func(o *ServerConfig) {
		o.ImplementationClassUID = uid
	}
}

// WithImplementationVersionName sets the implementation version name.
func WithImplementationVersionName(name string) ServerOption {
	return func(o *ServerConfig) {
		o.ImplementationVersionName = name
	}
}

// WithMaxConnections sets the maximum number of concurrent connections.
func WithMaxConnections(max int) ServerOption {
	return func(o *ServerConfig) {
		o.MaxConnections = max
	}
}

// WithTLSConfig sets the TLS configuration for secure connections.
func WithTLSConfig(tlsConfig *tls.Config) ServerOption {
	return func(o *ServerConfig) {
		o.TLSConfig = tlsConfig
	}
}

// defaultServerConfig returns the default server configuration.
func defaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:                      104,
		MaxPDULength:              16384,
		AcceptTimeout:             0, // No timeout
		AssociationTimeout:        10 * time.Second,
		RequestTimeout:            30 * time.Second,
		ImplementationClassUID:    "1.2.826.0.1.3680043.10.854",
		ImplementationVersionName: "GO-DICOM-1.0",
		MaxConnections:            0, // No limit
		TLSConfig:                 nil,
	}
}

// New creates a new DICOM server with the specified options.
func New(opts ...ServerOption) *Server {
	config := defaultServerConfig()
	for _, opt := range opts {
		opt(config)
	}

	return &Server{
		config:         config,
		connections:    make(map[string]*serverConnection),
		serviceOptions: make([]service.ServiceOption, 0),
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
func (s *Server) GetConfig() *ServerConfig {
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

// SetCFindHandler sets the C-FIND request handler.
func (s *Server) SetCFindHandler(handler func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error)) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCFindHandler(handler))
}

// SetCMoveHandler sets the C-MOVE request handler.
func (s *Server) SetCMoveHandler(handler func(context.Context, *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error)) {
	s.optionsMu.Lock()
	defer s.optionsMu.Unlock()
	s.serviceOptions = append(s.serviceOptions, service.WithCMoveHandler(handler))
}

// SetCGetHandler sets the C-GET request handler.
func (s *Server) SetCGetHandler(handler func(context.Context, *dimse.CGetRequest) ([]*dimse.CGetResponse, error)) {
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
	if s.running {
		s.runningMu.Unlock()
		return fmt.Errorf("server is already running")
	}
	s.running = true
	s.runningMu.Unlock()

	// Create server context
	s.ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	// Create listener
	var err error
	address := fmt.Sprintf(":%d", s.config.Port)

	if s.config.TLSConfig != nil {
		s.listener, err = transport.Listen("tcp", address, transport.WithListenTLSConfig(s.config.TLSConfig))
	} else {
		s.listener, err = transport.Listen("tcp", address)
	}

	if err != nil {
		s.runningMu.Lock()
		s.running = false
		s.runningMu.Unlock()
		return fmt.Errorf("failed to start listener: %w", err)
	}

	defer s.listener.Close()

	// Accept connections loop
	for {
		// Check if server was stopped
		select {
		case <-s.ctx.Done():
			s.runningMu.Lock()
			s.running = false
			s.runningMu.Unlock()
			// Wait for all connections to finish
			s.wg.Wait()
			return s.ctx.Err()
		default:
		}

		// Check connection limit
		if s.config.MaxConnections > 0 && s.ActiveConnections() >= s.config.MaxConnections {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Accept connection
		conn, err := s.listener.Accept(s.ctx)
		if err != nil {
			select {
			case <-s.ctx.Done():
				s.runningMu.Lock()
				s.running = false
				s.runningMu.Unlock()
				s.wg.Wait()
				return s.ctx.Err()
			default:
				// Log error and continue
				continue
			}
		}

		// Handle connection in goroutine
		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	if !s.IsRunning() {
		return nil
	}

	// Cancel server context
	if s.cancel != nil {
		s.cancel()
	}

	// Close listener
	if s.listener != nil {
		_ = s.listener.Close()
	}

	// Wait for connections to finish with timeout
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

// handleConnection handles a single client connection.
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	// Create connection ID
	connID := conn.RemoteAddr().String()

	// Build service options from server configuration
	svcOpts := []service.ServiceOption{
		service.WithMaxPDULength(s.config.MaxPDULength),
		service.WithReadTimeout(s.config.AssociationTimeout),
		service.WithWriteTimeout(s.config.AssociationTimeout),
	}

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
