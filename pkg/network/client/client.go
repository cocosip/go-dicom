// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package client provides a DICOM SCU (Service Class User) client implementation.
// It supports association negotiation, DIMSE operations, and graceful release.
package client

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
	SendCCancel(ctx context.Context, messageID uint16, presentationContextID byte) error
	SendNCreate(ctx context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error)
	SendNGet(ctx context.Context, req *dimse.NGetRequest) (*dimse.NGetResponse, error)
	SendNSet(ctx context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error)
	SendNDelete(ctx context.Context, req *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error)
	SendNAction(ctx context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error)
	SendNEventReport(ctx context.Context, req *dimse.NEventReportRequest) (*dimse.NEventReportResponse, error)
}

// cFindErrorService is implemented by service.Service. It is intentionally
// separate from serviceInterface so existing test and third-party service
// adapters remain source compatible while Client can surface asynchronous
// C-FIND terminal errors when the underlying implementation supports them.
type cFindErrorService interface {
	SendCFindWithError(context.Context, *dimse.CFindRequest) (<-chan *dimse.CFindResponse, <-chan error, error)
}

type cMoveErrorService interface {
	SendCMoveWithError(context.Context, *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, <-chan error, error)
}

type cGetErrorService interface {
	SendCGetWithError(context.Context, *dimse.CGetRequest) (<-chan *dimse.CGetResponse, <-chan error, error)
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
	connectionID observability.ConnectionID

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

	// Connection state. All fields below are protected by mu.
	mu             sync.Mutex
	connected      bool
	state          clientState
	connectCancel  context.CancelFunc
	transitionDone chan struct{}
}

type clientState uint8

const (
	clientDisconnected clientState = iota
	clientConnecting
	clientConnected
	clientClosing
)

// Config contains configuration options for the DICOM client.
type Config struct {
	// EventObserver receives connection, association, and DIMSE lifecycle events.
	EventObserver observability.EventObserver

	// MetricsObserver receives vendor-neutral network metrics.
	MetricsObserver observability.MetricsObserver

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

	// TransportReadTimeout limits a single PDU read. A zero value permits an
	// established association to remain idle until it is released or closed.
	// Default: 0 (disabled).
	TransportReadTimeout time.Duration

	// TransportWriteTimeout limits a single PDU write.
	// Default: 30 seconds.
	TransportWriteTimeout time.Duration

	// AssociationTimeout is the timeout for association negotiation
	// Default: 10 seconds
	AssociationTimeout time.Duration

	// TLSConfig enables TLS when non-nil. The configuration is caller-owned
	// and is cloned by the transport layer before use.
	// Default: nil (plain TCP)
	TLSConfig *tls.Config

	// ImplementationClassUID identifies the implementation
	// Default: "1.2.826.0.1.3680043.10.854"
	ImplementationClassUID string

	// ImplementationVersionName identifies the implementation version
	// Default: "GO-DICOM-1.0"
	ImplementationVersionName string

	// CStoreHandler handles incoming C-STORE requests on this connection.
	// Required for C-GET: the SCP sends C-STORE sub-operations back over the
	// same association, so this handler must be set to receive the images.
	CStoreHandler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)

	// KeepConnectionOnPeerRelease keeps the TCP connection open after replying
	// to a peer A-RELEASE-RQ for compatibility with non-conformant PACS.
	// Default: false.
	KeepConnectionOnPeerRelease bool

	// AsynchronousOperations requests the DIMSE asynchronous operations window.
	AsynchronousOperations *association.AsynchronousOperationsWindow

	// RoleSelections requests SCU/SCP roles by SOP Class UID.
	RoleSelections []*association.RoleSelection

	// ExtendedNegotiations contains requested SOP Class application information
	// and optional SOP Class Common Extended Negotiation values.
	ExtendedNegotiations []*association.ExtendedNegotiation

	// UserIdentity contains the optional association user identity request.
	UserIdentity *association.UserIdentity

	// RequireSuccessfulUserIdentityNegotiation rejects an association when a
	// requested positive user identity response is omitted.
	RequireSuccessfulUserIdentityNegotiation bool
}

// Option is a function that modifies client configuration.
type Option func(*Config)

var (
	// ErrTooManyPresentationContexts reports an attempt to configure more than
	// the 128 presentation contexts permitted in one DICOM association.
	ErrTooManyPresentationContexts = pdu.ErrTooManyPresentationContexts
	// ErrInvalidPresentationContext reports a context with no abstract syntax
	// or no usable transfer syntax.
	ErrInvalidPresentationContext = pdu.ErrInvalidPresentationContext
	// ErrClientConnecting indicates that an association attempt is already in progress.
	ErrClientConnecting = errors.New("client is connecting")
	// ErrClientConnected indicates that the client already owns an established association.
	ErrClientConnected = errors.New("client is already connected")
	// ErrClientClosing indicates that the client is releasing or cancelling an association.
	ErrClientClosing = errors.New("client is closing")
	// ErrClientNotConnected indicates that no usable association is established.
	ErrClientNotConnected = errors.New("client not connected")
)

// WithEventObserver sets the network lifecycle event observer.
func WithEventObserver(observer observability.EventObserver) Option {
	return func(o *Config) { o.EventObserver = observer }
}

// WithMetricsObserver sets the vendor-neutral network metrics observer.
func WithMetricsObserver(observer observability.MetricsObserver) Option {
	return func(o *Config) { o.MetricsObserver = observer }
}

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

// WithTransportReadTimeout sets the timeout for each transport PDU read.
// It does not affect the lifetime of an individual DIMSE request.
func WithTransportReadTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.TransportReadTimeout = timeout
	}
}

// WithTransportWriteTimeout sets the timeout for each transport PDU write.
// It does not affect the lifetime of an individual DIMSE request.
func WithTransportWriteTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.TransportWriteTimeout = timeout
	}
}

// WithAssociationTimeout sets the association timeout.
func WithAssociationTimeout(timeout time.Duration) Option {
	return func(o *Config) {
		o.AssociationTimeout = timeout
	}
}

// WithTLSConfig sets the TLS configuration for secure connections.
// A nil configuration keeps plain TCP as the default.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(o *Config) {
		o.TLSConfig = tlsConfig
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

// WithCStoreHandler sets a handler for incoming C-STORE requests on the client connection.
// This is required for C-GET operations: the SCP sends C-STORE sub-operations back
// over the same association, and without this handler the images are silently accepted
// but discarded.
func WithCStoreHandler(handler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)) Option {
	return func(o *Config) {
		o.CStoreHandler = handler
	}
}

// WithKeepConnectionOnPeerRelease controls whether the client keeps the TCP
// connection open after responding to an A-RELEASE-RQ from the peer. The
// default is false, which closes the connection according to the normal
// release lifecycle.
func WithKeepConnectionOnPeerRelease(keep bool) Option {
	return func(o *Config) {
		o.KeepConnectionOnPeerRelease = keep
	}
}

// WithAsynchronousOperations requests an asynchronous operations window.
// A value of zero means unlimited operations in that direction.
func WithAsynchronousOperations(maxInvoked, maxPerformed uint16) Option {
	return func(o *Config) {
		o.AsynchronousOperations = association.NewAsynchronousOperationsWindow(maxInvoked, maxPerformed)
	}
}

// WithRoleSelection adds or replaces the requested roles for a SOP Class UID.
func WithRoleSelection(selection *association.RoleSelection) Option {
	return func(o *Config) {
		o.RoleSelections = upsertRoleSelection(o.RoleSelections, selection)
	}
}

// WithExtendedNegotiation adds or merges negotiation values for a SOP Class UID.
func WithExtendedNegotiation(negotiation *association.ExtendedNegotiation) Option {
	return func(o *Config) {
		o.ExtendedNegotiations = upsertExtendedNegotiation(o.ExtendedNegotiations, negotiation)
	}
}

// WithUserIdentity configures association user identity negotiation.
func WithUserIdentity(identity *association.UserIdentity) Option {
	return func(o *Config) {
		o.UserIdentity = identity.Clone()
	}
}

// WithRequireSuccessfulUserIdentityNegotiation controls whether an omitted
// positive identity response rejects association establishment.
func WithRequireSuccessfulUserIdentityNegotiation(require bool) Option {
	return func(o *Config) {
		o.RequireSuccessfulUserIdentityNegotiation = require
	}
}

// defaultClientConfig returns the default client configuration.
func defaultClientConfig() *Config {
	return &Config{
		CallingAE:                                "GO_DICOM_SCU",
		CalledAE:                                 "ANY_SCP",
		MaxPDULength:                             16384,
		ConnectTimeout:                           10 * time.Second,
		RequestTimeout:                           30 * time.Second,
		TransportWriteTimeout:                    30 * time.Second,
		AssociationTimeout:                       10 * time.Second,
		ImplementationClassUID:                   "1.2.826.0.1.3680043.10.854",
		ImplementationVersionName:                "GO-DICOM-1.0",
		AsynchronousOperations:                   association.NewAsynchronousOperationsWindow(1, 1),
		RoleSelections:                           make([]*association.RoleSelection, 0),
		ExtendedNegotiations:                     make([]*association.ExtendedNegotiation, 0),
		RequireSuccessfulUserIdentityNegotiation: true,
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
func (c *Client) AddPresentationContext(abstractSyntax string, transferSyntaxes ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.addPresentationContextLocked(abstractSyntax, transferSyntaxes...)
}

func (c *Client) addPresentationContextLocked(abstractSyntax string, transferSyntaxes ...string) error {
	if err := c.configurationStateErrorLocked(); err != nil {
		return err
	}

	if len(c.presentationContexts) >= 128 {
		return fmt.Errorf("%w: got %d, maximum is 128", ErrTooManyPresentationContexts, len(c.presentationContexts)+1)
	}
	if abstractSyntax == "" {
		return fmt.Errorf("%w: abstract syntax is empty", ErrInvalidPresentationContext)
	}
	if len(transferSyntaxes) == 0 {
		return fmt.Errorf("%w: transfer syntaxes are empty", ErrInvalidPresentationContext)
	}
	for _, transferSyntax := range transferSyntaxes {
		if transferSyntax == "" {
			return fmt.Errorf("%w: transfer syntax is empty", ErrInvalidPresentationContext)
		}
	}
	// The count is bounded above, so this conversion cannot wrap.
	contextID := byte(len(c.presentationContexts)*2 + 1) // #nosec G115 -- bounded by the 128-context limit

	pc := &pdu.PresentationContextRQ{
		ID:               contextID,
		AbstractSyntax:   abstractSyntax,
		TransferSyntaxes: transferSyntaxes,
	}

	c.presentationContexts = append(c.presentationContexts, pc)
	return nil
}

// AddPresentationContextWithRoles adds a presentation context and requests the
// calling AE's SCU and SCP roles for its abstract syntax.
func (c *Client) AddPresentationContextWithRoles(
	abstractSyntax string,
	scuRole, scpRole bool,
	transferSyntaxes ...string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.addPresentationContextLocked(abstractSyntax, transferSyntaxes...); err != nil {
		return err
	}
	selection := association.NewRoleSelection(abstractSyntax, boolByte(scuRole), boolByte(scpRole))
	c.config.RoleSelections = upsertRoleSelection(c.config.RoleSelections, selection)
	return nil
}

// GetConfig returns the client configuration.
func (c *Client) GetConfig() *Config {
	return c.config
}

// GetAssociation returns the current association.
// Returns nil if not connected.
func (c *Client) GetAssociation() *association.Association {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.assoc
}

// IsConnected returns true if the client is connected.
func (c *Client) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected && c.state != clientClosing
}

func (c *Client) activeService() (serviceInterface, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.connected || c.service == nil {
		return nil, ErrClientNotConnected
	}
	return c.service, nil
}

// Close closes the client connection.
// If connected, it attempts a graceful release first using a short timeout.
// If the peer does not acknowledge the release in time, Close returns that
// error after cleaning up the local connection state.
func (c *Client) Close() error {
	c.mu.Lock()
	if c.state == clientConnecting {
		cancel := c.connectCancel
		conn := c.conn
		done := c.transitionDone
		c.state = clientClosing
		c.connected = false
		c.mu.Unlock()
		if cancel != nil {
			cancel()
		}
		if conn != nil {
			_ = conn.Close()
		}
		if done != nil {
			<-done
		}
		return nil
	}
	if c.state == clientClosing {
		done := c.transitionDone
		c.mu.Unlock()
		if done != nil {
			<-done
		}
		return nil
	}
	if !c.connected {
		c.mu.Unlock()
		return nil
	}
	svc := c.service
	conn := c.conn
	c.connected = false
	c.service = nil
	c.assoc = nil
	c.conn = nil
	c.state = clientClosing
	c.transitionDone = make(chan struct{})
	done := c.transitionDone
	c.mu.Unlock()

	// Try graceful release with a short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if svc != nil {
		err = svc.GracefulRelease(ctx)
		// Service layer closes the connection via GracefulRelease; avoid double close.
		conn = nil
	}

	if conn != nil {
		if closeErr := conn.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}

	c.finishTransition(done)
	return err
}

// Abort aborts the association and closes the connection.
func (c *Client) Abort(ctx context.Context) error {
	c.mu.Lock()
	if c.state == clientConnecting {
		cancel := c.connectCancel
		conn := c.conn
		done := c.transitionDone
		c.state = clientClosing
		c.connected = false
		c.mu.Unlock()
		if cancel != nil {
			cancel()
		}
		if conn != nil {
			_ = conn.Close()
		}
		if done != nil {
			<-done
		}
		return nil
	}
	if c.state == clientClosing {
		done := c.transitionDone
		c.mu.Unlock()
		if done != nil {
			<-done
		}
		return nil
	}
	if !c.connected {
		c.mu.Unlock()
		return nil
	}
	svc := c.service
	conn := c.conn
	c.connected = false
	c.service = nil
	c.assoc = nil
	c.conn = nil
	c.state = clientClosing
	c.transitionDone = make(chan struct{})
	done := c.transitionDone
	c.mu.Unlock()

	var err error
	if svc != nil {
		// Source: 0 = service-user (SCU initiated)
		// Reason: 0 = not-specified
		err = svc.Abort(ctx, 0, 0)
		// Service layer closes the connection via Abort; avoid double close.
		conn = nil
	}

	if conn != nil {
		if closeErr := conn.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}

	c.finishTransition(done)
	return err
}

func (c *Client) finishTransition(done chan struct{}) {
	c.mu.Lock()
	c.state = clientDisconnected
	c.connected = false
	c.connectCancel = nil
	if c.transitionDone == done {
		c.transitionDone = nil
	}
	c.mu.Unlock()
	if done != nil {
		close(done)
	}
}

// handleServiceClosed receives the lifecycle notification emitted by the
// current service. The identity check prevents an old association's receive
// loop from clearing a newer client session.
func (c *Client) handleServiceClosed(closed serviceInterface) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.service != closed {
		return
	}
	c.service = nil
	c.assoc = nil
	c.conn = nil
	c.connected = false
	if c.state == clientConnected {
		c.state = clientDisconnected
	}
}

func (c *Client) configurationStateErrorLocked() error {
	switch c.state {
	case clientConnecting:
		return ErrClientConnecting
	case clientConnected:
		return ErrClientConnected
	case clientClosing:
		return ErrClientClosing
	}
	if c.connected {
		return ErrClientConnected
	}
	return nil
}

// buildUserInformation builds the UserInformation structure for A-ASSOCIATE-RQ.
func (c *Client) buildUserInformation() *pdu.UserInformation {
	userInfo := &pdu.UserInformation{
		MaximumLength:             c.config.MaxPDULength,
		ImplementationClassUID:    c.config.ImplementationClassUID,
		ImplementationVersionName: c.config.ImplementationVersionName,
	}
	if async := c.config.AsynchronousOperations; async != nil &&
		(async.MaxInvokedOperations != 1 || async.MaxPerformedOperations != 1) {
		userInfo.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
			MaximumNumberOperationsInvoked:   async.MaxInvokedOperations,
			MaximumNumberOperationsPerformed: async.MaxPerformedOperations,
		}
	}
	for _, role := range c.config.RoleSelections {
		if role == nil {
			continue
		}
		userInfo.SCPSCURoleSelections = append(userInfo.SCPSCURoleSelections, pdu.SCPSCURoleSelection{
			SOPClassUID: role.SOPClassUID,
			SCURole:     role.SCURole,
			SCPRole:     role.SCPRole,
		})
	}
	for _, negotiation := range c.config.ExtendedNegotiations {
		if negotiation == nil {
			continue
		}
		if len(negotiation.RequestedApplicationInfo) > 0 {
			userInfo.ExtendedNegotiations = append(userInfo.ExtendedNegotiations, pdu.ExtendedNegotiation{
				SOPClassUID:         negotiation.SOPClassUID,
				ServiceClassAppInfo: append([]byte(nil), negotiation.RequestedApplicationInfo...),
			})
		}
		if negotiation.HasCommonExtendedNegotiation() {
			userInfo.CommonExtendedNegotiations = append(
				userInfo.CommonExtendedNegotiations,
				pdu.CommonExtendedNegotiation{
					SOPClassUID:                negotiation.SOPClassUID,
					ServiceClassUID:            negotiation.ServiceClassUID,
					RelatedGeneralSOPClassUIDs: append([]string(nil), negotiation.RelatedGeneralSOPClassUIDs...),
				},
			)
		}
	}
	if identity := c.config.UserIdentity; identity != nil {
		userInfo.UserIdentity = &pdu.UserIdentityNegotiation{
			UserIdentityType:          identity.Type,
			PositiveResponseRequested: boolByte(identity.PositiveResponseRequested),
			PrimaryField:              append([]byte(nil), identity.PrimaryField...),
			SecondaryField:            append([]byte(nil), identity.SecondaryField...),
		}
	}

	return userInfo
}

func boolByte(value bool) byte {
	if value {
		return 1
	}
	return 0
}

func upsertRoleSelection(values []*association.RoleSelection, selection *association.RoleSelection) []*association.RoleSelection {
	if selection == nil {
		return values
	}
	copySelection := association.NewRoleSelection(selection.SOPClassUID, selection.SCURole, selection.SCPRole)
	for i, value := range values {
		if value != nil && value.SOPClassUID == selection.SOPClassUID {
			values[i] = copySelection
			return values
		}
	}
	return append(values, copySelection)
}

func upsertExtendedNegotiation(values []*association.ExtendedNegotiation, negotiation *association.ExtendedNegotiation) []*association.ExtendedNegotiation {
	if negotiation == nil {
		return values
	}
	copyNegotiation := negotiation.Clone()
	for i, value := range values {
		if value != nil && value.SOPClassUID == negotiation.SOPClassUID {
			merged := value.Clone()
			if copyNegotiation.RequestedApplicationInfo != nil {
				merged.RequestedApplicationInfo = append([]byte(nil), copyNegotiation.RequestedApplicationInfo...)
				merged.ServiceClassAppInfo = merged.RequestedApplicationInfo
			}
			if copyNegotiation.HasCommonExtendedNegotiation() {
				merged.ServiceClassUID = copyNegotiation.ServiceClassUID
				merged.RelatedGeneralSOPClassUIDs = append([]string{}, copyNegotiation.RelatedGeneralSOPClassUIDs...)
			}
			values[i] = merged
			return values
		}
	}
	return append(values, copyNegotiation)
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

func (c *Client) buildAcceptedAssociation(rq *pdu.AAssociateRQ, ac *pdu.AAssociateAC) (*association.Association, error) {
	if err := c.validateAssociateAC(ac); err != nil {
		return nil, err
	}
	assoc := association.FromAAssociateRQ(rq)
	if err := association.ApplyAAssociateAC(assoc, ac); err != nil {
		return nil, err
	}
	identity := assoc.UserIdentity
	if c.config.RequireSuccessfulUserIdentityNegotiation && identity != nil &&
		identity.PositiveResponseRequested && identity.ServerResponse == nil {
		return nil, fmt.Errorf("positive user identity response was requested but not returned")
	}
	return assoc, nil
}

// dial establishes a TCP connection to the remote host.
func (c *Client) dial(ctx context.Context, host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)

	if c.config.TLSConfig != nil {
		dialCtx := ctx
		if c.config.ConnectTimeout > 0 {
			var cancel context.CancelFunc
			dialCtx, cancel = context.WithTimeout(ctx, c.config.ConnectTimeout)
			defer cancel()
		}

		conn, err := transport.DialTLS(dialCtx, "tcp", address,
			transport.WithTimeout(c.config.ConnectTimeout),
			transport.WithTLSConfig(c.config.TLSConfig),
		)
		if err != nil {
			return fmt.Errorf("failed to connect to %s: %w", address, err)
		}

		c.mu.Lock()
		if c.state == clientClosing {
			c.mu.Unlock()
			_ = conn.Close()
			return ErrClientClosing
		}
		c.conn = conn
		c.mu.Unlock()
		return nil
	}

	// Create dialer with timeout
	dialer := &net.Dialer{
		Timeout: c.config.ConnectTimeout,
	}

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", address, err)
	}

	c.mu.Lock()
	if c.state == clientClosing {
		c.mu.Unlock()
		_ = conn.Close()
		return ErrClientClosing
	}
	c.conn = conn
	c.mu.Unlock()
	return nil
}

// negotiateAssociation performs DICOM association negotiation.
func (c *Client) negotiateAssociation(ctx context.Context) error {
	// Create timeout context for association
	assocCtx, cancel := context.WithTimeout(ctx, c.config.AssociationTimeout)
	defer cancel()

	// Create service
	svcOpts := []service.Option{
		service.WithAssociationRequestor(true),
		service.WithConnectionID(c.connectionID),
		service.WithEventObserver(c.config.EventObserver),
		service.WithMetricsObserver(c.config.MetricsObserver),
		service.WithMaxPDULength(c.config.MaxPDULength),
		service.WithRequestTimeout(c.config.RequestTimeout),
		service.WithReadTimeout(c.config.TransportReadTimeout),
		service.WithWriteTimeout(c.config.TransportWriteTimeout),
		service.WithKeepConnectionOnPeerRelease(c.config.KeepConnectionOnPeerRelease),
	}
	if c.config.CStoreHandler != nil {
		svcOpts = append(svcOpts, service.WithCStoreHandler(c.config.CStoreHandler))
	}
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return ErrClientNotConnected
	}
	var svc *service.Service
	svcOpts = append(svcOpts, service.WithConnectionLifecycleHandler(
		&service.ConnectionLifecycleHandlerFuncs{
			OnConnectionClosedFunc: func(_ context.Context, _ error) {
				c.handleServiceClosed(svc)
			},
		},
	))
	svc = service.NewService(conn, nil, svcOpts...)
	c.mu.Lock()
	c.service = svc
	c.mu.Unlock()

	// Build and send A-ASSOCIATE-RQ
	rq := c.buildAssociateRQ()
	if err := svc.SendAssociationRequest(assocCtx, rq); err != nil {
		return fmt.Errorf("failed to send A-ASSOCIATE-RQ: %w", err)
	}

	// Wait for A-ASSOCIATE-AC or A-ASSOCIATE-RJ
	ac, err := svc.ReceiveAssociationResponse(assocCtx)
	if err != nil {
		return fmt.Errorf("association rejected: %w", err)
	}

	// Validate and merge the response with the original request.
	acceptedAssociation, err := c.buildAcceptedAssociation(rq, ac)
	if err != nil {
		return fmt.Errorf("invalid A-ASSOCIATE-AC: %w", err)
	}

	// Map abstract syntaxes from RQ to the accepted contexts in AC
	// (A-ASSOCIATE-AC doesn't include AbstractSyntax, so we need to restore it from RQ)
	for _, pcRQ := range rq.PresentationContexts {
		pcAssoc := acceptedAssociation.FindPresentationContextByID(pcRQ.ID)
		if pcAssoc != nil {
			pcAssoc.AbstractSyntax = pcRQ.AbstractSyntax
		}
	}

	svc.SetAssociation(acceptedAssociation)
	c.mu.Lock()
	if c.state == clientClosing {
		c.mu.Unlock()
		return ErrClientClosing
	}
	c.assoc = acceptedAssociation
	c.mu.Unlock()

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
	c.mu.Lock()
	if err := c.configurationStateErrorLocked(); err != nil {
		c.mu.Unlock()
		return err
	}
	if len(c.presentationContexts) == 0 {
		c.mu.Unlock()
		return fmt.Errorf("no presentation contexts configured (use AddPresentationContext)")
	}
	attemptCtx, cancel := context.WithCancel(ctx)
	c.connectCancel = cancel
	c.transitionDone = make(chan struct{})
	done := c.transitionDone
	c.state = clientConnecting
	c.connected = false
	c.connectionID = observability.NewConnectionID()
	c.mu.Unlock()
	defer cancel()
	c.emitConnectionAttempted(attemptCtx)

	fail := func(err error) error {
		c.finishConnectingAttempt(done)
		c.emitConnectionFailure(attemptCtx, err)
		return err
	}

	// Step 1: Establish TCP connection
	if err := c.dial(attemptCtx, host, port); err != nil {
		return fail(err)
	}

	// Step 2-3: Negotiate association
	if err := c.negotiateAssociation(attemptCtx); err != nil {
		return fail(err)
	}

	// Step 4: Start service loops
	c.mu.Lock()
	svc := c.service
	c.mu.Unlock()
	if err := svc.Start(); err != nil {
		return fail(fmt.Errorf("failed to start service: %w", err))
	}

	c.mu.Lock()
	if c.state != clientConnecting || c.transitionDone != done {
		c.mu.Unlock()
		return fail(ErrClientClosing)
	}
	c.connected = true
	c.state = clientConnected
	c.connectCancel = nil
	c.transitionDone = nil
	c.mu.Unlock()
	close(done)
	return nil
}

func (c *Client) finishConnectingAttempt(done chan struct{}) {
	c.mu.Lock()
	conn := c.conn
	c.conn = nil
	c.service = nil
	c.assoc = nil
	c.connected = false
	c.state = clientDisconnected
	c.connectCancel = nil
	if c.transitionDone == done {
		c.transitionDone = nil
	}
	c.mu.Unlock()
	if conn != nil {
		_ = conn.Close()
	}
	if done != nil {
		close(done)
	}
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
		if err := c.AddPresentationContext(
			"1.2.840.10008.1.1",   // Verification SOP Class
			"1.2.840.10008.1.2",   // Implicit VR Little Endian
			"1.2.840.10008.1.2.1", // Explicit VR Little Endian
		); err != nil {
			return nil, err
		}
	}

	if err := c.Connect(ctx, host, port); err != nil {
		return nil, err
	}

	return c, nil
}
