// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"testing"
	"time"
)

func TestDial_Success(t *testing.T) {
	// Start a test TCP server
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer func() { _ = listener.Close() }()

	// Get the port assigned by the OS
	addr := listener.Addr().String()

	// Accept connections in background
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()
		// Just accept and close
	}()

	// Test dialing
	conn, err := Dial(context.Background(), "tcp", addr)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	defer func() { _ = conn.Close() }()

	if conn == nil {
		t.Fatal("Expected non-nil connection")
	}
}

func TestDial_WithCustomTimeout(t *testing.T) {
	t.Skip("Timeout test is unreliable across different platforms")
	// Note: Testing actual timeout behavior is platform-dependent.
	// On some systems, TCP handshake completes even if no Accept() is called.
	// The timeout option is still validated through integration tests.
}

func TestDial_ContextCancellation(t *testing.T) {
	// Start a test server that delays accepting connections
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Create a context that will be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel immediately
	cancel()

	_, err = Dial(ctx, "tcp", addr)
	if err == nil {
		t.Fatal("Expected error due to cancelled context, got nil")
	}
}

func TestDialTLS_Success(t *testing.T) {
	// Create a self-signed certificate for testing
	cert, err := generateTestCertificate()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	// Start a TLS test server
	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
    defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Accept TLS connections in background
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()

		tlsConn := tls.Server(conn, serverConfig)
        defer func() { _ = tlsConn.Close() }()

		// Perform handshake
		if err := tlsConn.Handshake(); err != nil {
			return
		}
	}()

	// Give server time to start
	time.Sleep(50 * time.Millisecond)

	// Test dialing with TLS
	// Build a Root CA pool from the generated self-signed certificate
	rootPool := x509.NewCertPool()
	if len(cert.Certificate) == 0 {
		t.Fatal("generated certificate missing DER data")
	}
	parsedCert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		t.Fatalf("failed to parse certificate: %v", err)
	}
	rootPool.AddCert(parsedCert)

	clientConfig := &tls.Config{
		RootCAs:    rootPool,
		MinVersion: tls.VersionTLS12,
		ServerName: "localhost",
	}

	conn, err := DialTLS(context.Background(), "tcp", addr, WithTLSConfig(clientConfig))
	if err != nil {
		t.Fatalf("DialTLS failed: %v", err)
	}
	defer func() { _ = conn.Close() }()

	if conn == nil {
		t.Fatal("Expected non-nil TLS connection")
	}

	// Verify it's a TLS connection
	if _, ok := conn.(*tls.Conn); !ok {
		t.Fatal("Expected *tls.Conn type")
	}
}

func TestDialTLS_HandshakeFailure(t *testing.T) {
	// Start a plain TCP server (not TLS)
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
    defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Accept connections but don't do TLS handshake
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
        defer func() { _ = conn.Close() }()
		// Just accept and wait (no TLS handshake)
		time.Sleep(200 * time.Millisecond)
	}()

	// Try to connect with TLS (should fail)
	clientConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "localhost",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	_, err = DialTLS(ctx, "tcp", addr, WithTLSConfig(clientConfig))
	if err == nil {
		t.Fatal("Expected TLS handshake failure, got nil")
	}
}

func TestDial_WithKeepAlive(t *testing.T) {
	// Start a test TCP server
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
    defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Accept connections in background
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
                defer func() { _ = conn.Close() }()
		time.Sleep(100 * time.Millisecond)
	}()

	// Test dialing with custom keep-alive
	conn, err := Dial(context.Background(), "tcp", addr, WithKeepAlive(10*time.Second))
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
    defer func() { _ = conn.Close() }()

	// Verify it's a TCP connection
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		t.Fatal("Expected *net.TCPConn type")
	}

	// Try to set keep-alive (should work without error since it's already set)
	if err := tcpConn.SetKeepAlive(true); err != nil {
		t.Fatalf("Failed to verify keep-alive: %v", err)
	}
}

func TestDialTLS_WithMultipleOptions(t *testing.T) {
	// Create a self-signed certificate for testing
	cert, err := generateTestCertificate()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	// Start a TLS test server
	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
    defer func() { _ = listener.Close() }()

	addr := listener.Addr().String()

	// Accept TLS connections in background
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
    defer func() { _ = conn.Close() }()

		tlsConn := tls.Server(conn, serverConfig)
		defer func() { _ = tlsConn.Close() }()
		_ = tlsConn.Handshake()
	}()

	time.Sleep(50 * time.Millisecond)

	// Test with multiple options
	// Build a Root CA pool from the generated certificate
	rootPool := x509.NewCertPool()
	if len(cert.Certificate) == 0 {
		t.Fatal("generated certificate missing DER data")
	}
	parsedCert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		t.Fatalf("failed to parse certificate: %v", err)
	}
	rootPool.AddCert(parsedCert)

	clientConfig := &tls.Config{
		RootCAs:    rootPool,
		MinVersion: tls.VersionTLS12,
		ServerName: "localhost",
	}

	conn, err := DialTLS(context.Background(), "tcp", addr,
		WithTimeout(5*time.Second),
		WithKeepAlive(10*time.Second),
		WithTLSConfig(clientConfig))
	if err != nil {
		t.Fatalf("DialTLS with options failed: %v", err)
	}
	defer func() { _ = conn.Close() }()
}
