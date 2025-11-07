// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
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
	defer listener.Close()

	if listener == nil {
		t.Fatal("Expected non-nil listener")
	}

	addr := listener.Addr()
	if addr == nil {
		t.Fatal("Expected non-nil address")
	}
}

func TestListener_Accept(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Connect from a client in a goroutine
	go func() {
		time.Sleep(50 * time.Millisecond)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return
		}
		conn.Close()
	}()

	// Accept the connection
	conn, err := listener.Accept(context.Background())
	if err != nil {
		t.Fatalf("Accept failed: %v", err)
	}
	defer conn.Close()

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
	defer listener.Close()

	// Create a context that we'll cancel
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Try to accept (should fail due to cancelled context)
	_, err = listener.Accept(ctx)
	if err == nil {
		t.Fatal("Expected error due to cancelled context, got nil")
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
	}

	// Create a TLS listener
	listener, err := Listen("tcp", "127.0.0.1:0", WithListenTLSConfig(serverConfig))
	if err != nil {
		t.Fatalf("Failed to create TLS listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Connect from a TLS client in a goroutine
	go func() {
		time.Sleep(50 * time.Millisecond)
		clientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		conn, err := tls.Dial("tcp", addr, clientConfig)
		if err != nil {
			return
		}
		conn.Close()
	}()

	// Accept the TLS connection
	conn, err := listener.Accept(context.Background())
	if err != nil {
		t.Fatalf("Accept TLS failed: %v", err)
	}
	defer conn.Close()

	// Verify it's a TLS connection
	if _, ok := conn.(*tls.Conn); !ok {
		t.Fatal("Expected *tls.Conn type")
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
	}

	// Create a TLS listener
	listener, err := Listen("tcp", "127.0.0.1:0", WithListenTLSConfig(serverConfig))
	if err != nil {
		t.Fatalf("Failed to create TLS listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Connect with a plain TCP client (no TLS)
	go func() {
		time.Sleep(50 * time.Millisecond)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return
		}
		defer conn.Close()
		// Just hold the connection for a bit
		time.Sleep(200 * time.Millisecond)
	}()

	// Try to accept (should fail TLS handshake)
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	_, err = listener.Accept(ctx)
	if err == nil {
		t.Fatal("Expected TLS handshake failure, got nil")
	}
}

func TestListener_MultipleConnections(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

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
			defer conn.Close()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// Accept multiple connections
	for i := 0; i < numConns; i++ {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			t.Fatalf("Accept %d failed: %v", i, err)
		}
		conn.Close()
	}
}
