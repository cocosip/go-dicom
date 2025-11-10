// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// Option configures a Service.
type Option func(*serviceConfig)

// serviceConfig contains configuration options for a DICOM service.
type serviceConfig struct {
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

	// recvQueueSize is the size of the receive queue channel.
	// Default: 100
	recvQueueSize int

	// Lifecycle callbacks (optional)
	associationNegotiator      AssociationNegotiator
	associationReleaseHandler  AssociationReleaseHandler
	connectionLifecycleHandler ConnectionLifecycleHandler

	// DIMSE message handlers (optional)
	handlers *Handlers
}

// defaultServiceConfig returns the default service configuration.
func defaultServiceConfig() *serviceConfig {
	return &serviceConfig{
		maxPDULength:  16384, // 16 KB
		readTimeout:   30 * time.Second,
		writeTimeout:  30 * time.Second,
		dimseTimeout:  60 * time.Second,
		sendQueueSize: 100,
		recvQueueSize: 100,
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

// WithRecvQueueSize sets the size of the receive queue channel.
// Default: 100.
func WithRecvQueueSize(size int) Option {
	return func(c *serviceConfig) {
		c.recvQueueSize = size
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

// WithCMoveHandler sets the C-MOVE request handler.
func WithCMoveHandler(handler func(context.Context, *dimse.CMoveRequest) ([]*dimse.CMoveResponse, error)) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CMoveHandler = handler
	}
}

// WithCGetHandler sets the C-GET request handler.
func WithCGetHandler(handler func(context.Context, *dimse.CGetRequest) ([]*dimse.CGetResponse, error)) Option {
	return func(c *serviceConfig) {
		if c.handlers == nil {
			c.handlers = &Handlers{}
		}
		c.handlers.CGetHandler = handler
	}
}
