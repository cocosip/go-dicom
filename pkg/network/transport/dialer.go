// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

// DialOption configures how we dial a DICOM connection.
type DialOption func(*dialConfig)

// dialConfig holds the configuration for dialing.
type dialConfig struct {
	timeout   time.Duration
	keepAlive time.Duration
	tlsConfig *tls.Config
}

// WithTimeout sets the maximum amount of time a dial will wait for
// a connect to complete. If zero, no timeout is applied.
// Default: 30 seconds.
func WithTimeout(timeout time.Duration) DialOption {
	return func(c *dialConfig) {
		c.timeout = timeout
	}
}

// WithKeepAlive sets the interval between keep-alive probes
// for an active network connection. If zero, keep-alive probes
// are sent with a default interval.
// Default: 15 seconds.
func WithKeepAlive(keepAlive time.Duration) DialOption {
	return func(c *dialConfig) {
		c.keepAlive = keepAlive
	}
}

// WithTLSConfig specifies the TLS configuration to use.
// If nil, the default configuration is used for TLS connections.
func WithTLSConfig(config *tls.Config) DialOption {
	return func(c *dialConfig) {
		c.tlsConfig = config
	}
}

// defaultDialConfig returns the default dial configuration.
func defaultDialConfig() *dialConfig {
	return &dialConfig{
		timeout:   30 * time.Second,
		keepAlive: 15 * time.Second,
	}
}

// Dial connects to the address on the named network (tcp, tcp4, tcp6).
// The address format is "host:port".
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned.
//
// Example:
//
//	conn, err := transport.Dial(context.Background(), "tcp", "192.168.1.100:104")
//
//	// With custom timeout
//	conn, err := transport.Dial(ctx, "tcp", "192.168.1.100:104",
//	    transport.WithTimeout(10*time.Second))
func Dial(ctx context.Context, network, address string, opts ...DialOption) (net.Conn, error) {
	// Apply options
	config := defaultDialConfig()
	for _, opt := range opts {
		opt(config)
	}

	// Create a net.Dialer with our settings
	dialer := &net.Dialer{
		Timeout:   config.timeout,
		KeepAlive: config.keepAlive,
	}

	// Dial the connection
	conn, err := dialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s:%s: %w", network, address, err)
	}

	// Enable TCP keep-alive if it's a TCP connection
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if err := tcpConn.SetKeepAlive(true); err != nil {
			_ = conn.Close()
			return nil, fmt.Errorf("failed to set keep-alive: %w", err)
		}
		if config.keepAlive > 0 {
			if err := tcpConn.SetKeepAlivePeriod(config.keepAlive); err != nil {
				_ = conn.Close()
				return nil, fmt.Errorf("failed to set keep-alive period: %w", err)
			}
		}
	}

	return conn, nil
}

// DialTLS connects to the given address and initiates a TLS handshake,
// returning the resulting TLS connection.
//
// The provided Context must be non-nil. If the context expires before
// the connection is complete, an error is returned.
//
// Example:
//
//	conn, err := transport.DialTLS(ctx, "tcp", "pacs.example.com:11112",
//	    transport.WithTLSConfig(&tls.Config{ServerName: "pacs.example.com"}))
func DialTLS(ctx context.Context, network, address string, opts ...DialOption) (net.Conn, error) {
	// Apply options
	config := defaultDialConfig()
	for _, opt := range opts {
		opt(config)
	}

	// First establish a plain TCP connection
	rawConn, err := Dial(ctx, network, address,
		WithTimeout(config.timeout),
		WithKeepAlive(config.keepAlive))
	if err != nil {
		return nil, err
	}

	// Get TLS config (use default if not specified)
    tlsConfig := config.tlsConfig
    if tlsConfig == nil {
        tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
    } else if tlsConfig.MinVersion == 0 {
        tlsConfig.MinVersion = tls.VersionTLS12
    }

	// Wrap the connection with TLS
	tlsConn := tls.Client(rawConn, tlsConfig)

	// Perform the TLS handshake with context
	errChan := make(chan error, 1)
	go func() {
		errChan <- tlsConn.Handshake()
	}()

	select {
	case <-ctx.Done():
		_ = tlsConn.Close()
		return nil, fmt.Errorf("TLS handshake cancelled: %w", ctx.Err())
	case err := <-errChan:
		if err != nil {
			_ = tlsConn.Close()
			return nil, fmt.Errorf("TLS handshake failed: %w", err)
		}
	}

	return tlsConn, nil
}
