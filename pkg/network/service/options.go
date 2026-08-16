// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
)

// Option configures a Service.
type Option func(*serviceConfig)

// serviceConfig contains configuration options for a DICOM service.
type serviceConfig struct {
	// Observability hooks are optional and silent by default.
	logger          observability.Logger
	eventObserver   observability.EventObserver
	metricsObserver observability.MetricsObserver
	connectionID    observability.ConnectionID

	// associationRequestor indicates that the local AE initiated the association.
	// The default is false for services accepting inbound associations.
	associationRequestor bool

	// maxPDULength is the maximum PDU length in bytes.
	// Default: 16384 (16 KB)
	// DICOM standard allows up to 2^32-1, but practical limits are lower.
	maxPDULength uint32

	// readTimeout is the timeout for reading PDUs from the connection.
	// Default: 30 seconds
	// Set to 0 to disable timeout.
	readTimeout time.Duration

	// writeTimeout is the timeout for writing PDUs to the connection.
	// Default: 30 seconds
	// Set to 0 to disable timeout.
	writeTimeout time.Duration

	// dimseTimeout is the timeout for DIMSE operations (request/response).
	// Default: 60 seconds
	// Set to 0 to disable timeout.
	dimseTimeout time.Duration

	// sendQueueSize is the size of the send queue channel.
	// Default: 100
	sendQueueSize int

	// keepConnectionOnPeerRelease keeps the TCP connection open after replying
	// to a peer A-RELEASE-RQ for compatibility with non-conformant PACS.
	keepConnectionOnPeerRelease bool

	// Lifecycle callbacks (optional)
	associationNegotiator      AssociationNegotiator
	associationReleaseHandler  AssociationReleaseHandler
	connectionLifecycleHandler ConnectionLifecycleHandler

	// DIMSE message handlers (optional)
	handlers *Handlers
}

// WithLogger sets the structured network logger.
func WithLogger(logger observability.Logger) Option {
	return func(c *serviceConfig) {
		c.logger = logger
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
		maxPDULength:  16384, // 16 KB
		readTimeout:   30 * time.Second,
		writeTimeout:  30 * time.Second,
		dimseTimeout:  60 * time.Second,
		sendQueueSize: 100,
	}
}

// WithMaxPDULength sets the maximum PDU length.
// Default: 16384 bytes (16 KB).
func WithMaxPDULength(length uint32) Option {
	return func(c *serviceConfig) {
		c.maxPDULength = length
	}
}

// WithReadTimeout sets the read timeout for PDU operations.
// Default: 30 seconds. Set to 0 to disable timeout.
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

// WithDIMSETimeout sets the timeout for DIMSE request/response operations.
// Default: 60 seconds. Set to 0 to disable timeout.
func WithDIMSETimeout(timeout time.Duration) Option {
	return func(c *serviceConfig) {
		c.dimseTimeout = timeout
	}
}

// WithSendQueueSize sets the size of the send queue channel.
// Default: 100.
func WithSendQueueSize(size int) Option {
	return func(c *serviceConfig) {
		c.sendQueueSize = size
	}
}

// WithKeepConnectionOnPeerRelease controls whether the service closes the TCP
// connection after responding to an A-RELEASE-RQ from the peer. The default is
// false, which closes the connection according to the normal release lifecycle.
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

// WithCFindHandler sets the C-FIND request handler.
func WithCFindHandler(handler func(context.Context, *dimse.CFindRequest) ([]*dimse.CFindResponse, error)) Option {
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
