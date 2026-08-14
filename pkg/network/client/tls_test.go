// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

func TestNewWithTLSConfig(t *testing.T) {
	tlsConfig := &tls.Config{ServerName: "pacs.example.com", MinVersion: tls.VersionTLS12}

	client := New(WithTLSConfig(tlsConfig))

	if client.GetConfig().TLSConfig != tlsConfig {
		t.Fatal("client did not retain the caller-provided TLS configuration")
	}
}

func TestDialUsesTLSWhenConfigured(t *testing.T) {
	tlsServer := newQuietTLSServer()
	defer tlsServer.Close()

	roots := x509.NewCertPool()
	roots.AddCert(tlsServer.Certificate())
	host, portText, err := net.SplitHostPort(tlsServer.Listener.Addr().String())
	if err != nil {
		t.Fatalf("split TLS server address: %v", err)
	}
	port, err := net.LookupPort("tcp", portText)
	if err != nil {
		t.Fatalf("parse TLS server port: %v", err)
	}

	client := New(
		WithTLSConfig(&tls.Config{
			RootCAs:    roots,
			ServerName: host,
			MinVersion: tls.VersionTLS12,
		}),
		WithConnectTimeout(2*time.Second),
	)
	defer func() {
		if client.conn != nil {
			_ = client.conn.Close()
		}
	}()

	if err := client.dial(context.Background(), host, port); err != nil {
		t.Fatalf("dial TLS server: %v", err)
	}
	tlsConn, ok := client.conn.(*tls.Conn)
	if !ok {
		t.Fatalf("client connection type = %T, want *tls.Conn", client.conn)
	}
	if !tlsConn.ConnectionState().HandshakeComplete {
		t.Fatal("TLS handshake was not completed before dial returned")
	}
}

func TestDialTLSConnectTimeoutIncludesHandshake(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer func() { _ = listener.Close() }()

	serverDone := make(chan struct{})
	defer close(serverDone)
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			return
		}
		defer func() { _ = conn.Close() }()
		<-serverDone
	}()

	address := listener.Addr().(*net.TCPAddr)
	client := New(
		WithTLSConfig(&tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS12}), //nolint:gosec // stalled local test server has no certificate
		WithConnectTimeout(50*time.Millisecond),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	started := time.Now()
	err = client.dial(ctx, address.IP.String(), address.Port)
	elapsed := time.Since(started)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("dial error = %v, want context deadline exceeded", err)
	}
	if elapsed >= 500*time.Millisecond {
		t.Fatalf("TLS handshake returned after %v, want client connect timeout", elapsed)
	}
}

func TestConnectNegotiatesAssociationOverTLS(t *testing.T) {
	seedServer := newQuietTLSServer()
	serverTLSConfig := &tls.Config{
		Certificates: seedServer.TLS.Certificates,
		MinVersion:   tls.VersionTLS12,
	}
	roots := x509.NewCertPool()
	roots.AddCert(seedServer.Certificate())
	seedServer.Close()

	listener, err := transport.Listen("tcp", "127.0.0.1:0", transport.WithListenTLSConfig(serverTLSConfig))
	if err != nil {
		t.Fatalf("listen with TLS: %v", err)
	}
	defer func() { _ = listener.Close() }()

	serverCtx, cancelServer := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelServer()
	serverStarted := make(chan error, 1)
	go func() {
		conn, acceptErr := listener.Accept(serverCtx)
		if acceptErr != nil {
			serverStarted <- acceptErr
			return
		}
		serverService := service.NewService(conn, nil,
			service.WithReadTimeout(time.Second),
			service.WithWriteTimeout(time.Second),
		)
		defer func() { _ = serverService.Close() }()
		if _, receiveErr := serverService.ReceiveAssociationRequest(serverCtx); receiveErr != nil {
			serverStarted <- receiveErr
			return
		}
		if startErr := serverService.Start(); startErr != nil {
			serverStarted <- startErr
			return
		}
		serverStarted <- nil
		<-serverService.Context().Done()
	}()

	address := listener.Addr().(*net.TCPAddr)
	client := New(
		WithTLSConfig(&tls.Config{
			RootCAs:    roots,
			ServerName: address.IP.String(),
			MinVersion: tls.VersionTLS12,
		}),
		WithConnectTimeout(time.Second),
		WithAssociationTimeout(time.Second),
	)
	client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1")

	if err := client.Connect(context.Background(), address.IP.String(), address.Port); err != nil {
		t.Fatalf("connect over TLS: %v", err)
	}
	defer func() { _ = client.Close() }()
	if err := <-serverStarted; err != nil {
		t.Fatalf("start association service: %v", err)
	}
	if !client.IsConnected() {
		t.Fatal("client is not connected after TLS association negotiation")
	}
	if client.GetAssociation() == nil {
		t.Fatal("client has no negotiated association")
	}
	if _, ok := client.conn.(*tls.Conn); !ok {
		t.Fatalf("client connection type = %T, want *tls.Conn", client.conn)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("close TLS association: %v", err)
	}
	if client.IsConnected() {
		t.Fatal("client remains connected after closing TLS association")
	}
}

func TestDialTLSRejectsUntrustedCertificate(t *testing.T) {
	tlsServer := newQuietTLSServer()
	defer tlsServer.Close()
	host, portText, err := net.SplitHostPort(tlsServer.Listener.Addr().String())
	if err != nil {
		t.Fatalf("split TLS server address: %v", err)
	}
	port, err := net.LookupPort("tcp", portText)
	if err != nil {
		t.Fatalf("parse TLS server port: %v", err)
	}

	client := New(
		WithTLSConfig(&tls.Config{
			RootCAs:    x509.NewCertPool(),
			ServerName: host,
			MinVersion: tls.VersionTLS12,
		}),
		WithConnectTimeout(time.Second),
	)

	err = client.dial(context.Background(), host, port)
	var unknownAuthorityError x509.UnknownAuthorityError
	if !errors.As(err, &unknownAuthorityError) {
		t.Fatalf("dial error = %v, want x509.UnknownAuthorityError", err)
	}
}

func TestDialTLSRejectsHostnameMismatch(t *testing.T) {
	tlsServer := newQuietTLSServer()
	defer tlsServer.Close()
	roots := x509.NewCertPool()
	roots.AddCert(tlsServer.Certificate())
	host, portText, err := net.SplitHostPort(tlsServer.Listener.Addr().String())
	if err != nil {
		t.Fatalf("split TLS server address: %v", err)
	}
	port, err := net.LookupPort("tcp", portText)
	if err != nil {
		t.Fatalf("parse TLS server port: %v", err)
	}

	client := New(
		WithTLSConfig(&tls.Config{
			RootCAs:    roots,
			ServerName: "not-example.invalid",
			MinVersion: tls.VersionTLS12,
		}),
		WithConnectTimeout(time.Second),
	)

	err = client.dial(context.Background(), host, port)
	var hostnameError x509.HostnameError
	if !errors.As(err, &hostnameError) {
		t.Fatalf("dial error = %v, want x509.HostnameError", err)
	}
}

func TestDialUsesPlainTCPByDefault(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer func() { _ = listener.Close() }()
	serverDone := make(chan struct{})
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			_ = conn.Close()
		}
		close(serverDone)
	}()

	address := listener.Addr().(*net.TCPAddr)
	client := New(WithConnectTimeout(time.Second))
	if err := client.dial(context.Background(), address.IP.String(), address.Port); err != nil {
		t.Fatalf("dial TCP server: %v", err)
	}
	defer func() { _ = client.conn.Close() }()
	if _, ok := client.conn.(*net.TCPConn); !ok {
		t.Fatalf("client connection type = %T, want *net.TCPConn", client.conn)
	}
	<-serverDone
}

func TestDialTLSPreservesCallerCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	client := New(
		WithTLSConfig(&tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS12}), //nolint:gosec // no connection is established
		WithConnectTimeout(time.Second),
	)

	err := client.dial(ctx, "127.0.0.1", 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("dial error = %v, want context canceled", err)
	}
}

func TestDialTLSConcurrentClientsShareConfig(t *testing.T) {
	tlsServer := newQuietTLSServer()
	defer tlsServer.Close()
	roots := x509.NewCertPool()
	roots.AddCert(tlsServer.Certificate())
	host, portText, err := net.SplitHostPort(tlsServer.Listener.Addr().String())
	if err != nil {
		t.Fatalf("split TLS server address: %v", err)
	}
	port, err := net.LookupPort("tcp", portText)
	if err != nil {
		t.Fatalf("parse TLS server port: %v", err)
	}
	sharedTLSConfig := &tls.Config{RootCAs: roots, ServerName: host}

	const clientCount = 8
	errorsByClient := make(chan error, clientCount)
	var wg sync.WaitGroup
	wg.Add(clientCount)
	for range clientCount {
		go func() {
			defer wg.Done()
			client := New(
				WithTLSConfig(sharedTLSConfig),
				WithConnectTimeout(2*time.Second),
			)
			if dialErr := client.dial(context.Background(), host, port); dialErr != nil {
				errorsByClient <- dialErr
				return
			}
			_ = client.conn.Close()
		}()
	}
	wg.Wait()
	close(errorsByClient)
	for clientErr := range errorsByClient {
		t.Errorf("concurrent TLS dial: %v", clientErr)
	}
	if sharedTLSConfig.MinVersion != 0 {
		t.Fatalf("shared TLS config MinVersion = %d, want caller value 0", sharedTLSConfig.MinVersion)
	}
}

func newQuietTLSServer() *httptest.Server {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	server.StartTLS()
	return server
}
