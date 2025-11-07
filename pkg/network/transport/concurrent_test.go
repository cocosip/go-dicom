// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"testing"
	"time"
)

func TestDial_Concurrent(t *testing.T) {
	// Start a test server
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Accept connections in background
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conn.Close()
		}
	}()

	// Test concurrent dials
	numDials := 50
	var wg sync.WaitGroup
	wg.Add(numDials)

	errChan := make(chan error, numDials)

	for i := 0; i < numDials; i++ {
		go func(id int) {
			defer wg.Done()

			conn, err := Dial(context.Background(), "tcp", addr)
			if err != nil {
				errChan <- err
				return
			}
			defer conn.Close()

			// Verify connection works
			if conn == nil {
				errChan <- err
			}
		}(i)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		t.Errorf("Concurrent dial error: %v", err)
	}
}

func TestListener_ConcurrentAccept(t *testing.T) {
	// Create a listener
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Number of concurrent connections
	numConns := 20

	// Channel to track accepted connections
	acceptedConns := make(chan int, numConns)

	// Accept connections in multiple goroutines
	var acceptWg sync.WaitGroup
	numAcceptors := 5

	for i := 0; i < numAcceptors; i++ {
		acceptWg.Add(1)
		go func(acceptorID int) {
			defer acceptWg.Done()

			for j := 0; j < numConns/numAcceptors+1; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
				conn, err := listener.Accept(ctx)
				cancel()

				if err == nil {
					conn.Close()
					acceptedConns <- acceptorID
				}
			}
		}(i)
	}

	// Give acceptors time to start
	time.Sleep(50 * time.Millisecond)

	// Connect multiple clients concurrently
	var clientWg sync.WaitGroup
	for i := 0; i < numConns; i++ {
		clientWg.Add(1)
		go func(id int) {
			defer clientWg.Done()

			conn, err := net.Dial("tcp", addr)
			if err != nil {
				t.Errorf("Client %d dial failed: %v", id, err)
				return
			}
			defer conn.Close()

			// Hold connection briefly
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	// Wait for all clients to finish
	clientWg.Wait()

	// Give time for all accepts to complete
	time.Sleep(100 * time.Millisecond)

	// Close listener to stop acceptors
	listener.Close()
	acceptWg.Wait()

	close(acceptedConns)

	// Count accepted connections
	acceptCount := 0
	for range acceptedConns {
		acceptCount++
	}

	if acceptCount == 0 {
		t.Error("No connections were accepted")
	} else {
		t.Logf("Accepted %d connections from %d acceptor goroutines", acceptCount, numAcceptors)
	}
}

func TestDialTLS_Concurrent(t *testing.T) {
	// Create a self-signed certificate for testing
	cert, err := generateTestCertificate()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Start a TLS listener
	listener, err := Listen("tcp", "127.0.0.1:0", WithListenTLSConfig(serverConfig))
	if err != nil {
		t.Fatalf("Failed to create TLS listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Accept TLS connections in background
	go func() {
		for {
			conn, err := listener.Accept(context.Background())
			if err != nil {
				return
			}
			conn.Close()
		}
	}()

	// Test concurrent TLS dials
	numDials := 20
	var wg sync.WaitGroup
	wg.Add(numDials)

	errChan := make(chan error, numDials)
	clientConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	for i := 0; i < numDials; i++ {
		go func(id int) {
			defer wg.Done()

			conn, err := DialTLS(context.Background(), "tcp", addr, WithTLSConfig(clientConfig))
			if err != nil {
				errChan <- err
				return
			}
			defer conn.Close()
		}(i)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		t.Errorf("Concurrent TLS dial error: %v", err)
	}
}

func TestListener_ConcurrentCloseAndAccept(t *testing.T) {
	// Test that Close() is safe to call concurrently with Accept()
	listener, err := Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	var wg sync.WaitGroup

	// Start multiple acceptors
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			listener.Accept(ctx)
			// We expect this to fail after Close() is called
		}()
	}

	// Close the listener after a short delay
	time.Sleep(50 * time.Millisecond)
	listener.Close()

	// Wait for all acceptors to finish
	wg.Wait()

	// This test passes if no panic occurs
}
