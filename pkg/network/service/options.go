// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"strings"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/transcode"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
)

// Option configures a Service.
type Option func(*serviceConfig)

// serviceConfig contains configuration options for a DICOM service.
type serviceConfig struct {
	// Observability hooks are optional and silent by default.
	eventObserver          observability.EventObserver
	metricsObserver        observability.MetricsObserver
	connectionID           observability.ConnectionID
	transcodeManager       *transcode.Manager
	transferSyntaxRegistry *transfer.Registry

	// supportedSOPClasses is the set of SOP Class UIDs this SCP implements.
	// A non-empty set drives A-ASSOCIATE negotiation (unsupported abstract
	// syntaxes are rejected with result 3) and the request-time 0x0122
	// fallback. An empty set means "no restriction" (accept all).
	supportedSOPClasses map[string]struct{}

	// associationRequestor indicates that the local AE initiated the association.
	// The default is false for services accepting inbound associations.
	associationRequestor bool

	// maxPDULength is the maximum PDU length this AE advertises it can receive.
	// It is a per-direction limit (PS3.7 D.3.3.1); the peer advertises its own.
	// Default: 16384 (16 KB). DICOM allows up to 2^32-1, but practical limits are lower.
	maxPDULength uint32

	// maxMessageSize bounds the total reassembled size of a single inbound DIMSE
	// message spread across many PDVs. A peer that keeps sending fragments for
	// one message cannot grow memory beyond this bound.
	// Default: 512 MiB (0 means unlimited).
	maxMessageSize uint32

	// implementationClassUID / implementationVersionName identify the local AE
	// in the A-ASSOCIATE-RQ and A-ASSOCIATE-AC (PS3.7 D.3.3.2). They never take
	// the remote AE's values.
	implementationClassUID    string
	implementationVersionName string

	// maxConcurrentRequests bounds the number of inbound request handler
	// goroutines. Exceeding it aborts the association instead of letting a peer
	// spawn unbounded goroutines. Zero means unlimited.
	maxConcurrentRequests int

	// asynchronousOperations is the local AE's negotiated operation window.
	asynchronousOperations *association.AsynchronousOperationsWindow

	// readTimeout is the timeout for reading PDUs from the connection.
	// Default: disabled
	// Set to 0 to disable timeout.
	readTimeout time.Duration

	// writeTimeout is the timeout for writing PDUs to the connection.
	// Default: 30 seconds
	// Set to 0 to disable timeout.
	writeTimeout time.Duration

	// requestTimeout is the default response idle timeout for an outgoing
	// DIMSE request. It does not set a transport socket deadline.
	requestTimeout time.Duration

	// handlerShutdownTimeout bounds Close while inbound request handlers unwind.
	// It does not control outgoing request/response timeouts.
	handlerShutdownTimeout time.Duration

	// sendQueueSize is the size of the send queue channel.
	// Default: 100
	sendQueueSize int

	// keepConnectionOnPeerRelease is retained for source compatibility and has
	// no effect. PS3.8 7.2.2 forbids any further P-DATA after the A-RELEASE-RP,
	// so replying to a peer release always ends the association.
	keepConnectionOnPeerRelease bool

	// Lifecycle callbacks (optional)
	associationNegotiator      AssociationNegotiator
	associationReleaseHandler  AssociationReleaseHandler
	connectionLifecycleHandler ConnectionLifecycleHandler

	// DIMSE message handlers (optional)
	handlers *Handlers
}

// WithTranscodeManager sets the Dataset transcode manager used when a C-STORE
// dataset does not match any accepted transfer syntax directly.
func WithTranscodeManager(manager *transcode.Manager) Option {
	return func(c *serviceConfig) {
		c.transcodeManager = manager
	}
}

// WithTransferSyntaxRegistry sets the isolated registry used to resolve
// presentation-context Transfer Syntax UIDs.
func WithTransferSyntaxRegistry(registry *transfer.Registry) Option {
	return func(c *serviceConfig) {
		c.transferSyntaxRegistry = registry
	}
}

// WithSupportedSOPClasses declares the SOP Class UIDs this SCP implements. When
// at least one UID is given, A-ASSOCIATE negotiation rejects any other abstract
// syntax with abstract-syntax-not-supported (result 3, PS3.8 9.3.3.2) and inbound
// requests for other SOP Classes are refused with 0x0122. A zero-length list
// keeps the default accept-all behaviour.
func WithSupportedSOPClasses(uids ...string) Option {
	return func(c *serviceConfig) {
		c.supportedSOPClasses = normalizeSOPClassSet(uids)
	}
}

// WithAsynchronousOperations sets the local AE's asynchronous operation window.
// A value of zero means unlimited operations in that direction.
func WithAsynchronousOperations(maxInvoked, maxPerformed uint16) Option {
	return func(c *serviceConfig) {
		c.asynchronousOperations = association.NewAsynchronousOperationsWindow(maxInvoked, maxPerformed)
	}
}

// WithEventObserver sets the network lifecycle event observer.
func WithEventObserver(observer observability.EventObserver) Option {
	return func(c *serviceConfig) {
		c.eventObserver = observer
	}
}

// WithMetricsObserver sets the vendor-neutral network metrics observer.
func WithMetricsObserver(observer observability.MetricsObserver) Option {
	return func(c *serviceConfig) {
		c.metricsObserver = observer
	}
}

// WithConnectionID assigns the process-local identifier used to correlate
// observations for this service. A zero value causes NewService to allocate one.
func WithConnectionID(id observability.ConnectionID) Option {
	return func(c *serviceConfig) {
		c.connectionID = id
	}
}

// defaultServiceConfig returns the default service configuration.
func defaultServiceConfig() *serviceConfig {
	return &serviceConfig{
		maxPDULength:              16384, // 16 KB
		maxMessageSize:            512 << 20,
		maxConcurrentRequests:     1024,
		implementationClassUID:    association.DefaultImplementationClassUID,
		implementationVersionName: association.DefaultImplementationVersionName,
		asynchronousOperations:    association.NewAsynchronousOperationsWindow(1, 1),
		transferSyntaxRegistry:    transfer.NewRegistry(),
		readTimeout:               0,
		writeTimeout:              30 * time.Second,
		handlerShutdownTimeout:    60 * time.Second,
		sendQueueSize:             100,
	}
}

// WithImplementationClassUID sets the Implementation Class UID the local AE
// advertises in the User Information item (PS3.7 D.3.3.2).
func WithImplementationClassUID(uid string) Option {
	return func(c *serviceConfig) {
		c.implementationClassUID = uid
	}
}

// WithImplementationVersionName sets the Implementation Version Name the local
// AE advertises in the User Information item.
func WithImplementationVersionName(name string) Option {
	return func(c *serviceConfig) {
		c.implementationVersionName = name
	}
}

// WithMaxConcurrentRequests bounds the number of concurrently running inbound
// request handlers. When the limit is exceeded the association is aborted with
// A-ABORT instead of spawning more goroutines. Zero disables the limit.
// Default: 1024.
func WithMaxConcurrentRequests(limit int) Option {
	return func(c *serviceConfig) {
		c.maxConcurrentRequests = limit
	}
}

// WithMaxPDULength sets the maximum PDU length this AE advertises it can
// receive. Outgoing PDUs are fragmented to the smaller of this value and the
// peer's advertised limit. Default: 16384 bytes (16 KB); 0 means unlimited.
func WithMaxPDULength(length uint32) Option {
	return func(c *serviceConfig) {
		c.maxPDULength = length
	}
}

// WithMaxMessageSize bounds the total reassembled size of a single inbound DIMSE
// message spread across many PDVs. The default is 512 MiB. Zero disables the
// bound, which is only safe when every peer is trusted.
func WithMaxMessageSize(size uint32) Option {
	return func(c *serviceConfig) {
		c.maxMessageSize = size
	}
}

// WithReadTimeout sets the read timeout for PDU operations.
// Default: disabled (no socket deadline). Set to 0 to disable the timeout.
func WithReadTimeout(timeout time.Duration) Option {
	return func(c *serviceConfig) {
		c.readTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout for PDU operations.
// Default: 30 seconds. Set to 0 to disable timeout.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(c *serviceConfig) {
		c.writeTimeout = timeout
	}
}

// WithRequestTimeout sets the default response idle timeout for outgoing
// DIMSE requests. Set it to zero to rely only on the caller's context.
func WithRequestTimeout(timeout time.Duration) Option {
	return func(c *serviceConfig) {
		c.requestTimeout = timeout
	}
}

// WithHandlerShutdownTimeout sets the time Close waits for inbound request
// handlers to return before completing shutdown. Set it to zero to use the
// default bounded wait.
func WithHandlerShutdownTimeout(timeout time.Duration) Option {
	return func(c *serviceConfig) {
		c.handlerShutdownTimeout = timeout
	}
}

// WithSendQueueSize sets the size of the send queue channel.
// Default: 100.
func WithSendQueueSize(size int) Option {
	return func(c *serviceConfig) {
		c.sendQueueSize = size
	}
}

// WithKeepConnectionOnPeerRelease is deprecated and has no effect.
//
// After the local AE answers a peer A-RELEASE-RQ with A-RELEASE-RP it must not
// send any further P-DATA and the transport is closed (PS3.8 7.2.2), so the
// association can never stay usable on the same connection.
//
// Deprecated: keeping the connection open after a release is not conformant.
func WithKeepConnectionOnPeerRelease(keep bool) Option {
	return func(c *serviceConfig) {
		c.keepConnectionOnPeerRelease = keep
	}
}

// WithAssociationRequestor identifies whether the local AE initiated the
// association. This controls the direction of negotiated SCU/SCP roles.
func WithAssociationRequestor(requestor bool) Option {
	return func(c *serviceConfig) {
		c.associationRequestor = requestor
	}
}

// WithAssociationNegotiator sets the association negotiator callback.
// The negotiator controls which associations are accepted and how presentation contexts are negotiated.
func WithAssociationNegotiator(negotiator AssociationNegotiator) Option {
	return func(c *serviceConfig) {
		c.associationNegotiator = negotiator
	}
}

// WithAssociationReleaseHandler sets the association release handler callback.
// The handler is called when an A-RELEASE-RQ is received.
func WithAssociationReleaseHandler(handler AssociationReleaseHandler) Option {
	return func(c *serviceConfig) {
		c.associationReleaseHandler = handler
	}
}

// WithConnectionLifecycleHandler sets the connection lifecycle handler callback.
// The handler is called for connection lifecycle events (abort, close).
func WithConnectionLifecycleHandler(handler ConnectionLifecycleHandler) Option {
	return func(c *serviceConfig) {
		c.connectionLifecycleHandler = handler
	}
}

// WithHandlers sets the DIMSE message handlers.
// The handlers process incoming DIMSE requests (C-ECHO, C-STORE, C-FIND, C-MOVE, C-GET).
func WithHandlers(handlers *Handlers) Option {
	return func(c *serviceConfig) {
		c.handlers = handlers
	}
}

// WithCEchoHandler sets the C-ECHO request handler.
func WithCEchoHandler(handler func(context.Context, *dimse.CEchoRequest) (*dimse.CEchoResponse, error)) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CEchoHandler = handler
	}
}

// WithCStoreHandler sets the C-STORE request handler.
func WithCStoreHandler(handler func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CStoreHandler = handler
	}
}

// WithCFindHandler sets the streaming C-FIND request handler.
func WithCFindHandler(handler func(context.Context, CFindOperation) error) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CFindHandler = handler
	}
}

// WithCMoveHandler sets the C-MOVE request handler via a CMoveOperation interface.
// The handler calls op.SendPending after each sub-operation and finishes with
// op.SendSuccess(), op.SendWarning(), or op.SendFailure(code).
func WithCMoveHandler(handler func(ctx context.Context, op CMoveOperation) error) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CMoveHandler = handler
	}
}

// WithCGetHandler sets the C-GET handler via a CGetOperation interface.
// The handler calls op.SendCStore for each file, op.SendPending after each result,
// and finishes with op.SendSuccess(), op.SendWarning(), or op.SendFailure(code).
func WithCGetHandler(handler func(ctx context.Context, op CGetOperation) error) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CGetHandler = handler
	}
}

// normalizeSOPClassSet trims and indexes UIDs. A nil or empty result means "no
// restriction".
func normalizeSOPClassSet(uids []string) map[string]struct{} {
	if len(uids) == 0 {
		return nil
	}
	set := make(map[string]struct{}, len(uids))
	for _, sopClassUID := range uids {
		if sopClassUID = trimSOPClassUID(sopClassUID); sopClassUID != "" {
			set[sopClassUID] = struct{}{}
		}
	}
	if len(set) == 0 {
		return nil
	}
	return set
}

// trimSOPClassUID removes the padding and surrounding spaces from a UID read
// from a command/data set.
func trimSOPClassUID(sopClassUID string) string {
	return strings.TrimSpace(strings.TrimRight(sopClassUID, "\x00"))
}
