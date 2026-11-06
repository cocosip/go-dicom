// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package transport

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"testing"
	"time"
)

func TestListen_Success(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	if listener == nil {
		t.Fatal("Expected non-nil listener")
	}

	addr := listener.Addr()
	if addr == nil {
		t.Fatal("Expected non-nil address")
	}
}

func TestListener_Accept(t *testing.T) {
	var err error
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Connect from a client in a goroutine
	go func() {
		time.Sleep(50 * time.Millisecond)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Logf("Failed to connect from client: %v", err)
			return
		}
		_ = conn.Close()
	}()

	// Accept the connection
	conn, err := listener.Accept(context.Background())
	if err != nil {
		t.Fatalf("Accept failed: %v", err)
	}
	defer func() { _ = conn.Close() }()

	if conn == nil {
		t.Fatal("Expected non-nil connection")
	}
}

func TestListener_AcceptContextCancellation(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	// Create a context that we'll cancel
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Try to accept (should fail due to cancelled context)
	_, err = listener.Accept(ctx)
	if err == nil {
		t.Fatal("Expected error due to cancelled context, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

func TestListener_Close(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	// Close the listener
	err = listener.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Try to accept after close (should fail)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err = listener.Accept(ctx)
	if err == nil {
		t.Fatal("Expected error after close, got nil")
	}
}

func TestListenTLS_Success(t *testing.T) {
	// Create a self-signed certificate for testing
	cert, err := generateTestCertificate()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	// Create a TLS listener
	listener, err := Listen("tcp", "127.0.0.1:0", WithListenTLSConfig(serverConfig))
	if err != nil {
		t.Fatalf("Failed to create TLS listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Connect from a TLS client in a goroutine
	go func() {
		time.Sleep(50 * time.Millisecond)
		// Build a Root CA pool from the generated cert
		rootPool := x509.NewCertPool()
		if len(cert.Certificate) == 0 {
			return
		}
		parsedCert, err := x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			return
		}
		rootPool.AddCert(parsedCert)
		clientConfig := &tls.Config{
			RootCAs:    rootPool,
			MinVersion: tls.VersionTLS12,
			ServerName: localhostName,
		}
		conn, err := tls.Dial("tcp", addr, clientConfig)
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()
		_, _ = conn.Write([]byte("ping"))
	}()

	// Accept the TLS connection. The handshake is lazy, so Accept must return
	// without waiting for the client.
	conn, err := listener.Accept(context.Background())
	if err != nil {
		t.Fatalf("Accept TLS failed: %v", err)
	}
	defer func() { _ = conn.Close() }()

	// Verify it's a TLS connection
	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		t.Fatalf("Expected *tls.Conn type, got %T", conn)
	}

	// Trigger the deferred handshake with the caller's read deadline.
	if err := conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("SetReadDeadline failed: %v", err)
	}
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		t.Fatalf("TLS read failed: %v", err)
	}
	if string(buf) != "ping" {
		t.Fatalf("read %q, want %q", buf, "ping")
	}
	if !tlsConn.ConnectionState().HandshakeComplete {
		t.Fatal("TLS handshake was not completed by the first read")
	}
}

func TestListenTLS_HandshakeFailure(t *testing.T) {
	// Create a self-signed certificate for testing
	cert, err := generateTestCertificate()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	// Create a TLS listener
	listener, err := Listen("tcp", "127.0.0.1:0", WithListenTLSConfig(serverConfig))
	if err != nil {
		t.Fatalf("Failed to create TLS listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Connect with a plain TCP client that sends non-TLS bytes.
	go func() {
		time.Sleep(50 * time.Millisecond)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()
		_, _ = conn.Write([]byte("this is not a TLS record\r\n"))
		// Hold the connection open long enough for the server to read.
		time.Sleep(200 * time.Millisecond)
	}()

	// Accept succeeds immediately: the handshake is deferred.
	conn, err := listener.Accept(context.Background())
	if err != nil {
		t.Fatalf("Accept failed before handshake: %v", err)
	}
	defer func() { _ = conn.Close() }()

	// The handshake failure is reported by the first read, not by Accept.
	if err := conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("SetReadDeadline failed: %v", err)
	}
	if _, err := conn.Read(make([]byte, 1)); err == nil {
		t.Fatal("Expected TLS handshake failure on first read, got nil")
	}
}

func TestListener_AcceptContextTimeoutDoesNotConsumeNextConnection(t *testing.T) {
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err = listener.Accept(ctx)
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}

	client, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to dial listener: %v", err)
	}
	defer func() { _ = client.Close() }()

	acceptCtx, acceptCancel := context.WithTimeout(context.Background(), time.Second)
	defer acceptCancel()

	conn, err := listener.Accept(acceptCtx)
	if err != nil {
		t.Fatalf("Accept after timeout failed: %v", err)
	}
	defer func() { _ = conn.Close() }()
}

func TestListener_MultipleConnections(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Number of connections to test
	numConns := 5

	// Connect multiple clients
	for i := 0; i < numConns; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id*20) * time.Millisecond)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			defer func() { _ = conn.Close() }()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// Accept multiple connections
	for i := 0; i < numConns; i++ {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			t.Fatalf("Accept %d failed: %v", i, err)
		}
		_ = conn.Close()
	}
}
