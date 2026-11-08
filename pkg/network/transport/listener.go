// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"
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

	// Clone the TLS config so every connection shares a private copy (no data
	// race on a caller-owned pointer) and enforce at least TLS 1.2, mirroring
	// DialTLS.
	tlsConfig := config.tlsConfig
	if tlsConfig != nil {
		tlsConfig = tlsConfig.Clone()
		if tlsConfig.MinVersion == 0 {
			tlsConfig.MinVersion = tls.VersionTLS12
		}
	}

	return &Listener{
		listener:  listener,
		tlsConfig: tlsConfig,
	}, nil
}

// Accept waits for and returns the next connection to the listener.
// If the listener was created with TLS support, the returned connection is a
// *tls.Conn. The TLS handshake is performed lazily by crypto/tls on the first
// Read or Write, i.e. in the per-connection goroutine under that goroutine's
// read/write deadline, so a stalled handshake can never block the accept loop.
//
// The provided Context can be used to cancel the accept operation. It does not
// cover the (deferred) TLS handshake: crypto/tls performs it lazily on the first
// Read/Write under the caller's per-connection deadline. A server bounds it via
// its association timeout, and Close can interrupt a stalled handshake because
// not-yet-negotiated connections are tracked.
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

	return tls.Server(conn, l.tlsConfig), nil
}

func (l *Listener) acceptWithContext(ctx context.Context) (net.Conn, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("accept cancelled: %w", err)
	}

	dler, deadlineCapable := l.listener.(deadlineListener)
	if !deadlineCapable || ctx.Done() == nil {
		conn, err := l.listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
			}
			return nil, fmt.Errorf("accept failed: %w", err)
		}
		return conn, nil
	}

	// Interrupt a blocking Accept as soon as the context is done, instead of
	// polling on a fixed interval (which woke the accept loop ~10 times/second
	// even when idle). The watcher forces an immediate deadline so Accept
	// returns at once.
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			_ = dler.SetDeadline(time.Now())
		case <-stop:
		}
	}()

	conn, err := l.listener.Accept()
	close(stop)
	wg.Wait()
	_ = dler.SetDeadline(time.Time{})

	if err != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
		}
		return nil, fmt.Errorf("accept failed: %w", err)
	}
	if ctx.Err() != nil {
		// The context raced with a successful accept; drop the connection.
		_ = conn.Close()
		return nil, fmt.Errorf("accept cancelled: %w", ctx.Err())
	}
	return conn, nil
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
