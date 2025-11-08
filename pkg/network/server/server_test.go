// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package server

import (
	"context"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

func TestNew(t *testing.T) {
	server := New()

	if server == nil {
		t.Fatal("New() returned nil")
	}

	// Verify default configuration
	config := server.GetConfig()
	if config.Port != 104 {
		t.Errorf("Expected Port 104, got %d", config.Port)
	}
	if config.MaxPDULength != 16384 {
		t.Errorf("Expected MaxPDULength 16384, got %d", config.MaxPDULength)
	}
	if config.AssociationTimeout != 10*time.Second {
		t.Errorf("Expected AssociationTimeout 10s, got %v", config.AssociationTimeout)
	}
	if config.RequestTimeout != 30*time.Second {
		t.Errorf("Expected RequestTimeout 30s, got %v", config.RequestTimeout)
	}
}

func TestNewWithOptions(t *testing.T) {
	server := New(
		WithPort(11112),
		WithMaxPDULength(32768),
		WithAssociationTimeout(15*time.Second),
		WithRequestTimeout(60*time.Second),
		WithImplementationClassUID("1.2.3.4.5"),
		WithImplementationVersionName("TEST-1.0"),
		WithMaxConnections(10),
	)

	config := server.GetConfig()
	if config.Port != 11112 {
		t.Errorf("Expected Port 11112, got %d", config.Port)
	}
	if config.MaxPDULength != 32768 {
		t.Errorf("Expected MaxPDULength 32768, got %d", config.MaxPDULength)
	}
	if config.AssociationTimeout != 15*time.Second {
		t.Errorf("Expected AssociationTimeout 15s, got %v", config.AssociationTimeout)
	}
	if config.RequestTimeout != 60*time.Second {
		t.Errorf("Expected RequestTimeout 60s, got %v", config.RequestTimeout)
	}
	if config.ImplementationClassUID != "1.2.3.4.5" {
		t.Errorf("Expected ImplementationClassUID '1.2.3.4.5', got '%s'", config.ImplementationClassUID)
	}
	if config.ImplementationVersionName != "TEST-1.0" {
		t.Errorf("Expected ImplementationVersionName 'TEST-1.0', got '%s'", config.ImplementationVersionName)
	}
	if config.MaxConnections != 10 {
		t.Errorf("Expected MaxConnections 10, got %d", config.MaxConnections)
	}
}

func TestSetHandlers(t *testing.T) {
	server := New()

	// Set C-ECHO handler
	server.SetCEchoHandler(func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
		return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
	})

	// Set C-STORE handler
	server.SetCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
		return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
	})

	// Set C-FIND handler
	server.SetCFindHandler(func(ctx context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
		return []*dimse.CFindResponse{
			dimse.NewCFindResponseFromRequest(req, 0x0000, nil),
		}, nil
	})

	// Note: Handlers are now stored as service options and cannot be directly verified.
	// They will be tested through integration tests that actually invoke the handlers.
}

func TestIsRunning(t *testing.T) {
	server := New()

	if server.IsRunning() {
		t.Error("Expected IsRunning() to return false initially")
	}
}

func TestActiveConnections(t *testing.T) {
	server := New()

	if server.ActiveConnections() != 0 {
		t.Errorf("Expected 0 active connections initially, got %d", server.ActiveConnections())
	}
}

func TestBuildUserInformation(t *testing.T) {
	server := New(
		WithMaxPDULength(32768),
		WithImplementationClassUID("1.2.3.4.5"),
		WithImplementationVersionName("TEST-1.0"),
	)

	userInfo := server.buildUserInformation()

	if userInfo.MaximumLength != 32768 {
		t.Errorf("Expected MaximumLength 32768, got %d", userInfo.MaximumLength)
	}
	if userInfo.ImplementationClassUID != "1.2.3.4.5" {
		t.Errorf("Expected ImplementationClassUID '1.2.3.4.5', got '%s'", userInfo.ImplementationClassUID)
	}
	if userInfo.ImplementationVersionName != "TEST-1.0" {
		t.Errorf("Expected ImplementationVersionName 'TEST-1.0', got '%s'", userInfo.ImplementationVersionName)
	}
}

func TestShutdownWhenNotRunning(t *testing.T) {
	server := New()

	ctx := context.Background()
	err := server.Shutdown(ctx)
	if err != nil {
		t.Errorf("Shutdown() should not return error when not running, got: %v", err)
	}
}
