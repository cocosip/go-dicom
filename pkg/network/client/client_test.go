// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	client := New()

	if client == nil {
		t.Fatal("New() returned nil")
	}

	// Verify default options
	opts := client.GetConfig()
	if opts.CallingAE != "GO_DICOM_SCU" {
		t.Errorf("Expected CallingAE 'GO_DICOM_SCU', got '%s'", opts.CallingAE)
	}
	if opts.CalledAE != "ANY_SCP" {
		t.Errorf("Expected CalledAE 'ANY_SCP', got '%s'", opts.CalledAE)
	}
	if opts.MaxPDULength != 16384 {
		t.Errorf("Expected MaxPDULength 16384, got %d", opts.MaxPDULength)
	}
	if opts.ConnectTimeout != 10*time.Second {
		t.Errorf("Expected ConnectTimeout 10s, got %v", opts.ConnectTimeout)
	}
	if opts.RequestTimeout != 30*time.Second {
		t.Errorf("Expected RequestTimeout 30s, got %v", opts.RequestTimeout)
	}
	if opts.AssociationTimeout != 10*time.Second {
		t.Errorf("Expected AssociationTimeout 10s, got %v", opts.AssociationTimeout)
	}
}

func TestNewWithOptions(t *testing.T) {
	client := New(
		WithCallingAE("MY_SCU"),
		WithCalledAE("REMOTE_SCP"),
		WithMaxPDULength(32768),
		WithConnectTimeout(5*time.Second),
		WithRequestTimeout(60*time.Second),
		WithAssociationTimeout(15*time.Second),
		WithImplementationClassUID("1.2.3.4.5"),
		WithImplementationVersionName("TEST-1.0"),
	)

	opts := client.GetConfig()
	if opts.CallingAE != "MY_SCU" {
		t.Errorf("Expected CallingAE 'MY_SCU', got '%s'", opts.CallingAE)
	}
	if opts.CalledAE != "REMOTE_SCP" {
		t.Errorf("Expected CalledAE 'REMOTE_SCP', got '%s'", opts.CalledAE)
	}
	if opts.MaxPDULength != 32768 {
		t.Errorf("Expected MaxPDULength 32768, got %d", opts.MaxPDULength)
	}
	if opts.ConnectTimeout != 5*time.Second {
		t.Errorf("Expected ConnectTimeout 5s, got %v", opts.ConnectTimeout)
	}
	if opts.RequestTimeout != 60*time.Second {
		t.Errorf("Expected RequestTimeout 60s, got %v", opts.RequestTimeout)
	}
	if opts.AssociationTimeout != 15*time.Second {
		t.Errorf("Expected AssociationTimeout 15s, got %v", opts.AssociationTimeout)
	}
	if opts.ImplementationClassUID != "1.2.3.4.5" {
		t.Errorf("Expected ImplementationClassUID '1.2.3.4.5', got '%s'", opts.ImplementationClassUID)
	}
	if opts.ImplementationVersionName != "TEST-1.0" {
		t.Errorf("Expected ImplementationVersionName 'TEST-1.0', got '%s'", opts.ImplementationVersionName)
	}
}

func TestAddPresentationContext(t *testing.T) {
	client := New()

	// Add first context
	client.AddPresentationContext("1.2.840.10008.1.1", // Verification
		"1.2.840.10008.1.2",   // Implicit VR Little Endian
		"1.2.840.10008.1.2.1", // Explicit VR Little Endian
	)

	if len(client.presentationContexts) != 1 {
		t.Fatalf("Expected 1 presentation context, got %d", len(client.presentationContexts))
	}

	pc := client.presentationContexts[0]
	if pc.ID != 1 {
		t.Errorf("Expected ID 1, got %d", pc.ID)
	}
	if pc.AbstractSyntax != "1.2.840.10008.1.1" {
		t.Errorf("Expected AbstractSyntax '1.2.840.10008.1.1', got '%s'", pc.AbstractSyntax)
	}
	if len(pc.TransferSyntaxes) != 2 {
		t.Errorf("Expected 2 transfer syntaxes, got %d", len(pc.TransferSyntaxes))
	}

	// Add second context
	client.AddPresentationContext("1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
		"1.2.840.10008.1.2", // Implicit VR Little Endian
	)

	if len(client.presentationContexts) != 2 {
		t.Fatalf("Expected 2 presentation contexts, got %d", len(client.presentationContexts))
	}

	pc2 := client.presentationContexts[1]
	if pc2.ID != 3 {
		t.Errorf("Expected ID 3, got %d", pc2.ID)
	}

	// Add third context
	client.AddPresentationContext("1.2.840.10008.5.1.4.1.1.4", // MR Image Storage
		"1.2.840.10008.1.2.1", // Explicit VR Little Endian
	)

	if len(client.presentationContexts) != 3 {
		t.Fatalf("Expected 3 presentation contexts, got %d", len(client.presentationContexts))
	}

	pc3 := client.presentationContexts[2]
	if pc3.ID != 5 {
		t.Errorf("Expected ID 5, got %d", pc3.ID)
	}
}

func TestIsConnected(t *testing.T) {
	client := New()

	if client.IsConnected() {
		t.Error("Expected IsConnected() to return false initially")
	}

	// Simulate connected state
	client.connected = true

	if !client.IsConnected() {
		t.Error("Expected IsConnected() to return true after setting connected")
	}
}

func TestGetAssociation(t *testing.T) {
	client := New()

	if assoc := client.GetAssociation(); assoc != nil {
		t.Error("Expected GetAssociation() to return nil initially")
	}
}

func TestCloseWhenNotConnected(t *testing.T) {
	client := New()

	err := client.Close()
	if err != nil {
		t.Errorf("Close() should not return error when not connected, got: %v", err)
	}
}

func TestBuildUserInformation(t *testing.T) {
	client := New(
		WithMaxPDULength(32768),
		WithImplementationClassUID("1.2.3.4.5"),
		WithImplementationVersionName("TEST-1.0"),
	)

	userInfo := client.buildUserInformation()

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

func TestBuildAssociateRQ(t *testing.T) {
	client := New(
		WithCallingAE("MY_SCU"),
		WithCalledAE("REMOTE_SCP"),
		WithMaxPDULength(16384),
	)

	// Add presentation context
	client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2")

	rq := client.buildAssociateRQ()

	if rq.CallingAETitle != "MY_SCU" {
		t.Errorf("Expected CallingAETitle 'MY_SCU', got '%s'", rq.CallingAETitle)
	}
	if rq.CalledAETitle != "REMOTE_SCP" {
		t.Errorf("Expected CalledAETitle 'REMOTE_SCP', got '%s'", rq.CalledAETitle)
	}
	if rq.ApplicationContext != "1.2.840.10008.3.1.1.1" {
		t.Errorf("Expected ApplicationContext '1.2.840.10008.3.1.1.1', got '%s'", rq.ApplicationContext)
	}
	if len(rq.PresentationContexts) != 1 {
		t.Errorf("Expected 1 presentation context, got %d", len(rq.PresentationContexts))
	}
	if rq.UserInformation == nil {
		t.Error("Expected UserInformation to be set")
	}
	if rq.UserInformation.MaximumLength != 16384 {
		t.Errorf("Expected MaximumLength 16384, got %d", rq.UserInformation.MaximumLength)
	}
}

func TestConnectWithoutPresentationContexts(t *testing.T) {
	client := New()

	// Try to connect without adding presentation contexts
	// Should fail immediately without attempting network connection
	err := client.Connect(nil, "localhost", 104)
	if err == nil {
		t.Error("Expected error when connecting without presentation contexts")
	}

	expectedMsg := "no presentation contexts configured (use AddPresentationContext)"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestConnectAlreadyConnected(t *testing.T) {
	client := New()
	client.connected = true

	err := client.Connect(nil, "localhost", 104)
	if err == nil {
		t.Error("Expected error when already connected")
	}

	expectedMsg := "client is already connected"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}
