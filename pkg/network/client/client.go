// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package client provides a DICOM SCU (Service Class User) client implementation.
// It supports association negotiation, DIMSE operations, and graceful release.
package client

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
)

// serviceInterface defines the interface for DIMSE service operations.
// This allows for mocking in tests.
type serviceInterface interface {
	// Association management
	GetAssociation() *association.Association
	SetAssociation(assoc *association.Association)
	SendAssociationRequest(ctx context.Context, rq *pdu.AAssociateRQ) error
	ReceiveAssociationResponse(ctx context.Context) (*pdu.AAssociateAC, error)
	Start() error
	GracefulRelease(ctx context.Context) error
	Abort(ctx context.Context, source byte, reason byte) error

	// DIMSE operations
	SendCEcho(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error)
	SendCStore(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error)
	SendCFind(ctx context.Context, req *dimse.CFindRequest) (<-chan *dimse.CFindResponse, error)
	SendCMove(ctx context.Context, req *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, error)
	SendCGet(ctx context.Context, req *dimse.CGetRequest) (<-chan *dimse.CGetResponse, error)
}

// Client represents a DICOM SCU (Service Class User) client.
// It provides a high-level API for connecting to DICOM servers
// and performing DIMSE operations.
//
// Example usage:
//
//	client := client.New(client.WithCallingAE("MY_SCU"))
//	client.AddPresentationContext("1.2.840.10008.1.1", // Verification SOP Class
//	    transfer.ImplicitVRLittleEndian.UID)
//
//	if err := client.Connect(ctx, "192.168.1.100", 104); err != nil {
//	    return err
//	}
//	defer client.Close()
//
//	if err := client.CEcho(ctx); err != nil {
//	    return err
//	}
type Client struct {
	// Network connection
	conn net.Conn

	// Service layer
	service serviceInterface

	// Association information
	assoc *association.Association

	// Configuration options
	config *Config

	// Presentation contexts to negotiate
	presentationContexts []*pdu.PresentationContextRQ

	// Connection state
	connected bool
}

// Config contains configuration options for the DICOM client.
type Config struct {
	// CallingAE is the AE Title of this client (SCU)
	CallingAE string

	// CalledAE is the AE Title of the remote server (SCP)
	CalledAE string

	// MaxPDULength is the maximum PDU length to negotiate
	// Default: 16384 bytes
	MaxPDULength uint32

	// ConnectTimeout is the timeout for TCP connection
	// Default: 10 seconds
	ConnectTimeout time.Duration

	// RequestTimeout is the default timeout for DIMSE requests
	// Default: 30 seconds
	RequestTimeout time.Duration

	// AssociationTimeout is the timeout for association negotiation
	// Default: 10 seconds
	AssociationTimeout time.Duration

	// ImplementationClassUID identifies the implementation
	// Default: "1.2.826.0.1.3680043.10.854"
	ImplementationClassUID string

	// ImplementationVersionName identifies the implementation version
	// Default: "GO-DICOM-1.0"
	ImplementationVersionName string
}

// Option is a function that modifies client configuration.
type Option func(*Config)

// WithCallingAE sets the calling AE title.
func WithCallingAE(ae string) Option {
	return func(o *Config) {
		o.CallingAE = ae
	}
}

// WithCalledAE sets the called AE title.
func WithCalledAE(ae string) Option {
	return func(o *Config) {
		o.CalledAE = ae
	}
}

// WithMaxPDULength sets the maximum PDU length.
func WithMaxPDULength(length uint32) Option {
	return func(o *Config) {
		o.MaxPDULength = length
	}
}

// WithConnectTimeout sets the connection timeout.
func WithConnectTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.ConnectTimeout = timeout
	}
}

// WithRequestTimeout sets the request timeout.
func WithRequestTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.RequestTimeout = timeout
	}
}

// WithAssociationTimeout sets the association timeout.
func WithAssociationTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.AssociationTimeout = timeout
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

// defaultClientConfig returns the default client configuration.
func defaultClientConfig() *Config {
	return &Config{
		CallingAE:                 "GO_DICOM_SCU",
		CalledAE:                  "ANY_SCP",
		MaxPDULength:              16384,
		ConnectTimeout:            10 * time.Second,
		RequestTimeout:            30 * time.Second,
		AssociationTimeout:        10 * time.Second,
		ImplementationClassUID:    "1.2.826.0.1.3680043.10.854",
		ImplementationVersionName: "GO-DICOM-1.0",
	}
}

// New creates a new DICOM client with the specified options.
func New(opts ...Option) *Client {
	config := defaultClientConfig()
	for _, opt := range opts {
		opt(config)
	}

	return &Client{
		config:               config,
		presentationContexts: make([]*pdu.PresentationContextRQ, 0),
	}
}

// AddPresentationContext adds a presentation context to be negotiated
// during association.
//
// Parameters:
//   - abstractSyntax: The SOP Class UID (e.g., "1.2.840.10008.1.1" for Verification)
//   - transferSyntaxes: List of transfer syntax UIDs to propose
//
// The presentation context ID will be automatically assigned (odd numbers: 1, 3, 5, ...).
func (c *Client) AddPresentationContext(abstractSyntax string, transferSyntaxes ...string) {
	// Calculate next context ID (odd numbers: 1, 3, 5, ...)
	contextID := byte(len(c.presentationContexts)*2 + 1)

	pc := &pdu.PresentationContextRQ{
		ID:               contextID,
		AbstractSyntax:   abstractSyntax,
		TransferSyntaxes: transferSyntaxes,
	}

	c.presentationContexts = append(c.presentationContexts, pc)
}

// GetConfig returns the client configuration.
func (c *Client) GetConfig() *Config {
	return c.config
}

// GetAssociation returns the current association.
// Returns nil if not connected.
func (c *Client) GetAssociation() *association.Association {
	return c.assoc
}

// IsConnected returns true if the client is connected.
func (c *Client) IsConnected() bool {
	return c.connected
}

// Close closes the client connection.
// If connected, it will attempt a graceful release first.
func (c *Client) Close() error {
	if !c.connected {
		return nil
	}

	// Try graceful release with a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if c.service != nil {
		err = c.service.GracefulRelease(ctx)
	}

	// Clean up
	c.connected = false
	c.service = nil
	c.assoc = nil

	if c.conn != nil {
		if closeErr := c.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		c.conn = nil
	}

	return err
}

// Abort aborts the association and closes the connection.
func (c *Client) Abort(ctx context.Context) error {
	if !c.connected {
		return nil
	}

	var err error
	if c.service != nil {
		// Source: 0 = service-user (SCU initiated)
		// Reason: 0 = not-specified
		err = c.service.Abort(ctx, 0, 0)
	}

	// Clean up
	c.connected = false
	c.service = nil
	c.assoc = nil

	if c.conn != nil {
		if closeErr := c.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		c.conn = nil
	}

	return err
}

// buildUserInformation builds the UserInformation structure for A-ASSOCIATE-RQ.
func (c *Client) buildUserInformation() *pdu.UserInformation {
	userInfo := &pdu.UserInformation{
		MaximumLength:             c.config.MaxPDULength,
		ImplementationClassUID:    c.config.ImplementationClassUID,
		ImplementationVersionName: c.config.ImplementationVersionName,
	}

	return userInfo
}

// buildAssociateRQ builds an A-ASSOCIATE-RQ PDU with the configured
// presentation contexts and user information.
func (c *Client) buildAssociateRQ() *pdu.AAssociateRQ {
	rq := pdu.NewAAssociateRQ()
	rq.CalledAETitle = c.config.CalledAE
	rq.CallingAETitle = c.config.CallingAE
	rq.ApplicationContext = "1.2.840.10008.3.1.1.1" // DICOM Application Context

	// Copy presentation contexts (convert from slice of pointers to slice of values)
	rq.PresentationContexts = make([]pdu.PresentationContextRQ, len(c.presentationContexts))
	for i, pc := range c.presentationContexts {
		rq.PresentationContexts[i] = *pc
	}

	rq.UserInformation = c.buildUserInformation()

	return rq
}

// validateAssociateAC validates the A-ASSOCIATE-AC response.
// Returns an error if no presentation contexts were accepted.
func (c *Client) validateAssociateAC(ac *pdu.AAssociateAC) error {
	if len(ac.PresentationContexts) == 0 {
		return fmt.Errorf("no presentation contexts in A-ASSOCIATE-AC")
	}

	// Check if at least one context was accepted
	acceptedCount := 0
	for _, pc := range ac.PresentationContexts {
		if pc.Result == pdu.ResultAcceptance {
			acceptedCount++
		}
	}

	if acceptedCount == 0 {
		return fmt.Errorf("all presentation contexts were rejected")
	}

	return nil
}

// dial establishes a TCP connection to the remote host.
func (c *Client) dial(ctx context.Context, host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)

	// Create dialer with timeout
	dialer := &net.Dialer{
		Timeout: c.config.ConnectTimeout,
	}

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", address, err)
	}

	c.conn = conn
	return nil
}

// negotiateAssociation performs DICOM association negotiation.
func (c *Client) negotiateAssociation(ctx context.Context) error {
	// Create timeout context for association
	assocCtx, cancel := context.WithTimeout(ctx, c.config.AssociationTimeout)
	defer cancel()

	// Create service
	c.service = service.NewService(c.conn, nil,
		service.WithMaxPDULength(c.config.MaxPDULength),
		service.WithReadTimeout(c.config.AssociationTimeout),
		service.WithWriteTimeout(c.config.AssociationTimeout),
	)

	// Build and send A-ASSOCIATE-RQ
	rq := c.buildAssociateRQ()
	if err := c.service.SendAssociationRequest(assocCtx, rq); err != nil {
		return fmt.Errorf("failed to send A-ASSOCIATE-RQ: %w", err)
	}

	// Wait for A-ASSOCIATE-AC or A-ASSOCIATE-RJ
	ac, err := c.service.ReceiveAssociationResponse(assocCtx)
	if err != nil {
		return fmt.Errorf("association rejected: %w", err)
	}

	// Validate the response
	if err := c.validateAssociateAC(ac); err != nil {
		return fmt.Errorf("invalid A-ASSOCIATE-AC: %w", err)
	}

	// Build association object from the accepted response
	c.assoc = association.FromAAssociateAC(ac)

	// Map abstract syntaxes from RQ to the accepted contexts in AC
	// (A-ASSOCIATE-AC doesn't include AbstractSyntax, so we need to restore it from RQ)
	for _, pcRQ := range rq.PresentationContexts {
		pcAssoc := c.assoc.FindPresentationContextByID(pcRQ.ID)
		if pcAssoc != nil {
			pcAssoc.AbstractSyntax = pcRQ.AbstractSyntax
		}
	}

	c.service.SetAssociation(c.assoc)

	return nil
}

// Connect establishes a connection to a DICOM server and negotiates association.
//
// Steps:
//  1. Establish TCP connection
//  2. Send A-ASSOCIATE-RQ with presentation contexts
//  3. Receive and validate A-ASSOCIATE-AC
//  4. Start the service send/receive loops
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - host: Server hostname or IP address
//   - port: Server port (typically 104 for DICOM)
//
// Returns an error if connection or association fails.
func (c *Client) Connect(ctx context.Context, host string, port int) error {
	if c.connected {
		return fmt.Errorf("client is already connected")
	}

	if len(c.presentationContexts) == 0 {
		return fmt.Errorf("no presentation contexts configured (use AddPresentationContext)")
	}

	// Step 1: Establish TCP connection
	if err := c.dial(ctx, host, port); err != nil {
		return err
	}

	// Step 2-3: Negotiate association
	if err := c.negotiateAssociation(ctx); err != nil {
		// Clean up connection on failure
		if c.conn != nil {
			_ = c.conn.Close()
			c.conn = nil
		}
		return err
	}

	// Step 4: Start service loops
	if err := c.service.Start(); err != nil {
		// Clean up on failure
		if c.conn != nil {
			_ = c.conn.Close()
			c.conn = nil
		}
		return fmt.Errorf("failed to start service: %w", err)
	}

	c.connected = true
	return nil
}

// Dial is a convenience function that creates a new client, adds presentation contexts,
// connects to the server, and returns the connected client.
//
// This is useful for simple use cases where you just want to connect quickly.
//
// Example:
//
//	client, err := client.Dial(ctx, "192.168.1.100", 104,
//	    client.WithCallingAE("MY_SCU"),
//	    client.WithCalledAE("REMOTE_SCP"))
func Dial(ctx context.Context, host string, port int, opts ...Option) (*Client, error) {
	c := New(opts...)

	// Add default verification context if no contexts specified
	if len(c.presentationContexts) == 0 {
		c.AddPresentationContext(
			"1.2.840.10008.1.1",   // Verification SOP Class
			"1.2.840.10008.1.2",   // Implicit VR Little Endian
			"1.2.840.10008.1.2.1", // Explicit VR Little Endian
		)
	}

	if err := c.Connect(ctx, host, port); err != nil {
		return nil, err
	}

	return c, nil
}
