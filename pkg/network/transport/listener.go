// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

// Listener wraps a net.Listener with optional TLS support for DICOM connections.
type Listener struct {
	listener  net.Listener
	tlsConfig *tls.Config
}

type deadlineListener interface {
	SetDeadline(time.Time) error
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
	if ctx == nil {
		return nil, fmt.Errorf("accept failed: context cannot be nil")
	}

	conn, err := l.acceptWithContext(ctx)
	if err != nil {
		return nil, err
	}

	if l.tlsConfig == nil {
		return conn, nil
	}

	tlsConn := tls.Server(conn, l.tlsConfig)
	if err := tlsConn.HandshakeContext(ctx); err != nil {
		_ = tlsConn.Close()
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("TLS handshake cancelled: %w", err)
		}
		return nil, fmt.Errorf("TLS handshake failed: %w", err)
	}

	return tlsConn, nil
}

func (l *Listener) acceptWithContext(ctx context.Context) (net.Conn, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("accept cancelled: %w", err)
	}

	deadlineCapable, ok := l.listener.(deadlineListener)
	if !ok {
		conn, err := l.listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
			}
			return nil, fmt.Errorf("accept failed: %w", err)
		}
		return conn, nil
	}

	for {
		deadline := time.Now().Add(100 * time.Millisecond)
		if deadlineErr := deadlineCapable.SetDeadline(deadline); deadlineErr != nil {
			return nil, fmt.Errorf("accept failed: %w", deadlineErr)
		}

		conn, err := l.listener.Accept()
		if err == nil {
			if clearErr := deadlineCapable.SetDeadline(time.Time{}); clearErr != nil {
				_ = conn.Close()
				return nil, fmt.Errorf("accept failed: %w", clearErr)
			}
			return conn, nil
		}

		if ctx.Err() != nil {
			_ = deadlineCapable.SetDeadline(time.Time{})
			return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
		}

		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			continue
		}

		if errors.Is(err, net.ErrClosed) || errors.Is(err, os.ErrClosed) {
			_ = deadlineCapable.SetDeadline(time.Time{})
			return nil, fmt.Errorf("accept failed: %w", err)
		}

		_ = deadlineCapable.SetDeadline(time.Time{})
		return nil, fmt.Errorf("accept failed: %w", err)
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
