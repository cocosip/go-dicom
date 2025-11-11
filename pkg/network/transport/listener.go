// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
)

// Listener wraps a net.Listener with optional TLS support for DICOM connections.
type Listener struct {
	listener  net.Listener
	tlsConfig *tls.Config
}

// ListenOption configures how we create a listener.
type ListenOption func(*listenConfig)

// listenConfig holds the configuration for listening.
type listenConfig struct {
	tlsConfig *tls.Config
}

// WithListenTLSConfig specifies the TLS configuration for the listener.
// If provided, all accepted connections will be upgraded to TLS.
func WithListenTLSConfig(config *tls.Config) ListenOption {
	return func(c *listenConfig) {
		c.tlsConfig = config
	}
}

// Listen creates a new Listener on the specified network and address.
// The network must be "tcp", "tcp4", or "tcp6".
// The address format is "host:port", e.g., ":104" or "192.168.1.100:104".
//
// Example:
//
//	listener, err := transport.Listen("tcp", ":104")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer listener.Close()
//
//	for {
//	    conn, err := listener.Accept(context.Background())
//	    if err != nil {
//	        continue
//	    }
//	    go handleConnection(conn)
//	}
func Listen(network, address string, opts ...ListenOption) (*Listener, error) {
	// Apply options
	config := &listenConfig{}
	for _, opt := range opts {
		opt(config)
	}

	// Create the listener
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s:%s: %w", network, address, err)
	}

	return &Listener{
		listener:  listener,
		tlsConfig: config.tlsConfig,
	}, nil
}

// Accept waits for and returns the next connection to the listener.
// If the listener was created with TLS support, the connection will be
// upgraded to TLS before being returned.
//
// The provided Context can be used to cancel the accept operation.
//
// Example:
//
//	conn, err := listener.Accept(ctx)
//	if err != nil {
//	    return err
//	}
//	defer conn.Close()
func (l *Listener) Accept(ctx context.Context) (net.Conn, error) {
	// Channel to receive the connection or error
	type result struct {
		conn net.Conn
		err  error
	}
	resultChan := make(chan result, 1)

	// Accept in a goroutine
	go func() {
		conn, err := l.listener.Accept()
		resultChan <- result{conn: conn, err: err}
	}()

	// Wait for either accept to complete or context to be cancelled
	select {
	case <-ctx.Done():
		// Context was cancelled, but we can't cancel Accept()
		// The connection will be accepted anyway, so we close it
		go func() {
			res := <-resultChan
			if res.conn != nil {
				_ = res.conn.Close()
			}
		}()
		return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
	case res := <-resultChan:
		if res.err != nil {
			return nil, fmt.Errorf("accept failed: %w", res.err)
		}

		// If TLS is enabled, upgrade the connection
		if l.tlsConfig != nil {
			tlsConn := tls.Server(res.conn, l.tlsConfig)

			// Perform TLS handshake with context
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

		return res.conn, nil
	}
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *Listener) Close() error {
	return l.listener.Close()
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return l.listener.Addr()
}
