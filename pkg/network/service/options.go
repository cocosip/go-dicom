// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import "time"

// ServiceOption configures a Service.
type ServiceOption func(*serviceConfig)

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
func WithMaxPDULength(length uint32) ServiceOption {
	return func(c *serviceConfig) {
		c.maxPDULength = length
	}
}

// WithReadTimeout sets the read timeout for PDU operations.
// Default: 30 seconds. Set to 0 to disable timeout.
func WithReadTimeout(timeout time.Duration) ServiceOption {
	return func(c *serviceConfig) {
		c.readTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout for PDU operations.
// Default: 30 seconds. Set to 0 to disable timeout.
func WithWriteTimeout(timeout time.Duration) ServiceOption {
	return func(c *serviceConfig) {
		c.writeTimeout = timeout
	}
}

// WithDIMSETimeout sets the timeout for DIMSE request/response operations.
// Default: 60 seconds. Set to 0 to disable timeout.
func WithDIMSETimeout(timeout time.Duration) ServiceOption {
	return func(c *serviceConfig) {
		c.dimseTimeout = timeout
	}
}

// WithSendQueueSize sets the size of the send queue channel.
// Default: 100.
func WithSendQueueSize(size int) ServiceOption {
	return func(c *serviceConfig) {
		c.sendQueueSize = size
	}
}

// WithRecvQueueSize sets the size of the receive queue channel.
// Default: 100.
func WithRecvQueueSize(size int) ServiceOption {
	return func(c *serviceConfig) {
		c.recvQueueSize = size
	}
}
