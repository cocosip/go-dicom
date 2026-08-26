// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"bytes"
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

const verificationSOPClassUID = "1.2.840.10008.1.1"

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
	if opts.TransportReadTimeout != 0 {
		t.Errorf("Expected TransportReadTimeout 0, got %v", opts.TransportReadTimeout)
	}
	if opts.TransportWriteTimeout != 30*time.Second {
		t.Errorf("Expected TransportWriteTimeout 30s, got %v", opts.TransportWriteTimeout)
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
		WithTransportReadTimeout(20*time.Second),
		WithTransportWriteTimeout(25*time.Second),
		WithAssociationTimeout(15*time.Second),
		WithImplementationClassUID(testImplementationClassUID),
		WithImplementationVersionName("TEST-1.0"),
		WithKeepConnectionOnPeerRelease(true),
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
	if opts.TransportReadTimeout != 20*time.Second {
		t.Errorf("Expected TransportReadTimeout 20s, got %v", opts.TransportReadTimeout)
	}
	if opts.TransportWriteTimeout != 25*time.Second {
		t.Errorf("Expected TransportWriteTimeout 25s, got %v", opts.TransportWriteTimeout)
	}
	if opts.AssociationTimeout != 15*time.Second {
		t.Errorf("Expected AssociationTimeout 15s, got %v", opts.AssociationTimeout)
	}
	if opts.ImplementationClassUID != testImplementationClassUID {
		t.Errorf("Expected ImplementationClassUID '1.2.3.4.5', got '%s'", opts.ImplementationClassUID)
	}
	if opts.ImplementationVersionName != "TEST-1.0" {
		t.Errorf("Expected ImplementationVersionName 'TEST-1.0', got '%s'", opts.ImplementationVersionName)
	}
	if !opts.KeepConnectionOnPeerRelease {
		t.Error("Expected KeepConnectionOnPeerRelease to be true")
	}
}

func TestAddPresentationContext(t *testing.T) {
	client := New()

	// Add first context
	if err := client.AddPresentationContext(verificationSOPClassUID, // Verification
		"1.2.840.10008.1.2",   // Implicit VR Little Endian
		"1.2.840.10008.1.2.1", // Explicit VR Little Endian
	); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}

	if len(client.presentationContexts) != 1 {
		t.Fatalf("Expected 1 presentation context, got %d", len(client.presentationContexts))
	}

	pc := client.presentationContexts[0]
	if pc.ID != 1 {
		t.Errorf("Expected ID 1, got %d", pc.ID)
	}
	if pc.AbstractSyntax != verificationSOPClassUID {
		t.Errorf("Expected AbstractSyntax %q, got %q", verificationSOPClassUID, pc.AbstractSyntax)
	}
	if len(pc.TransferSyntaxes) != 2 {
		t.Errorf("Expected 2 transfer syntaxes, got %d", len(pc.TransferSyntaxes))
	}

	// Add second context
	if err := client.AddPresentationContext(testCTImageStorageUID, // CT Image Storage
		"1.2.840.10008.1.2", // Implicit VR Little Endian
	); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}

	if len(client.presentationContexts) != 2 {
		t.Fatalf("Expected 2 presentation contexts, got %d", len(client.presentationContexts))
	}

	pc2 := client.presentationContexts[1]
	if pc2.ID != 3 {
		t.Errorf("Expected ID 3, got %d", pc2.ID)
	}

	// Add third context
	if err := client.AddPresentationContext("1.2.840.10008.5.1.4.1.1.4", // MR Image Storage
		"1.2.840.10008.1.2.1", // Explicit VR Little Endian
	); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}

	if len(client.presentationContexts) != 3 {
		t.Fatalf("Expected 3 presentation contexts, got %d", len(client.presentationContexts))
	}

	pc3 := client.presentationContexts[2]
	if pc3.ID != 5 {
		t.Errorf("Expected ID 5, got %d", pc3.ID)
	}
}

func TestAddPresentationContext_Rejects129thContextWithoutMutatingClient(t *testing.T) {
	client := New()
	for range 128 {
		if err := client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1"); err != nil {
			t.Fatalf("AddPresentationContext() before limit error = %v", err)
		}
	}

	if got := client.presentationContexts[127].ID; got != 255 {
		t.Fatalf("128th presentation context ID = %d, want 255", got)
	}

	err := client.AddPresentationContext("1.2.840.10008.5.1.4.1.1.2", "1.2.840.10008.1.2.1")
	if !errors.Is(err, ErrTooManyPresentationContexts) {
		t.Fatalf("129th AddPresentationContext() error = %v, want ErrTooManyPresentationContexts", err)
	}
	if got := len(client.presentationContexts); got != 128 {
		t.Fatalf("presentation context count after rejected add = %d, want 128", got)
	}
}

func TestAddPresentationContextWithRoles_DoesNotAddRoleWhenContextIsRejected(t *testing.T) {
	client := New()
	for range 128 {
		if err := client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1"); err != nil {
			t.Fatalf("AddPresentationContext() before limit error = %v", err)
		}
	}

	err := client.AddPresentationContextWithRoles("1.2.840.10008.5.1.4.1.1.2", true, false, "1.2.840.10008.1.2.1")
	if !errors.Is(err, ErrTooManyPresentationContexts) {
		t.Fatalf("AddPresentationContextWithRoles() error = %v, want ErrTooManyPresentationContexts", err)
	}
	if got := len(client.config.RoleSelections); got != 0 {
		t.Fatalf("role selection count after rejected add = %d, want 0", got)
	}
}

func TestAddPresentationContext_RejectsInvalidSyntaxesWithoutMutatingClient(t *testing.T) {
	tests := []struct {
		name             string
		abstractSyntax   string
		transferSyntaxes []string
	}{
		{name: "empty abstract syntax", transferSyntaxes: []string{testExplicitVRLittleEndianUID}},
		{name: "no transfer syntax", abstractSyntax: verificationSOPClassUID},
		{name: "empty transfer syntax", abstractSyntax: verificationSOPClassUID, transferSyntaxes: []string{""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New()
			err := client.AddPresentationContext(tt.abstractSyntax, tt.transferSyntaxes...)
			if !errors.Is(err, ErrInvalidPresentationContext) {
				t.Fatalf("AddPresentationContext() error = %v, want ErrInvalidPresentationContext", err)
			}
			if got := len(client.presentationContexts); got != 0 {
				t.Fatalf("presentation context count after rejected add = %d, want 0", got)
			}
		})
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
		WithImplementationClassUID(testImplementationClassUID),
		WithImplementationVersionName("TEST-1.0"),
	)

	userInfo := client.buildUserInformation()

	if userInfo.MaximumLength != 32768 {
		t.Errorf("Expected MaximumLength 32768, got %d", userInfo.MaximumLength)
	}
	if userInfo.ImplementationClassUID != testImplementationClassUID {
		t.Errorf("Expected ImplementationClassUID '1.2.3.4.5', got '%s'", userInfo.ImplementationClassUID)
	}
	if userInfo.ImplementationVersionName != "TEST-1.0" {
		t.Errorf("Expected ImplementationVersionName 'TEST-1.0', got '%s'", userInfo.ImplementationVersionName)
	}
}

func TestBuildAssociateRQIncludesAdvancedNegotiation(t *testing.T) {
	identity := association.NewUserIdentityJWT([]byte("request-token"), true)
	client := New(
		WithAsynchronousOperations(4, 3),
		WithExtendedNegotiation(association.NewExtendedNegotiation(
			"1.2.840.10008.5.1.4.1.1.2", []byte{1, 1, 0},
		)),
		WithUserIdentity(identity),
	)
	if err := client.AddPresentationContextWithRoles(
		"1.2.840.10008.5.1.4.1.1.2", true, true, "1.2.840.10008.1.2.1",
	); err != nil {
		t.Fatalf("AddPresentationContextWithRoles() error = %v", err)
	}

	rq := client.buildAssociateRQ()
	async := rq.UserInformation.AsynchronousOperations
	if async == nil || async.MaximumNumberOperationsInvoked != 4 || async.MaximumNumberOperationsPerformed != 3 {
		t.Fatalf("async operations = %#v", async)
	}
	roles := rq.UserInformation.SCPSCURoleSelections
	if len(roles) != 1 || roles[0].SOPClassUID != testCTImageStorageUID || roles[0].SCURole != 1 || roles[0].SCPRole != 1 {
		t.Fatalf("role selections = %#v", roles)
	}
	extended := rq.UserInformation.ExtendedNegotiations
	if len(extended) != 1 || !bytes.Equal(extended[0].ServiceClassAppInfo, []byte{1, 1, 0}) {
		t.Fatalf("extended negotiations = %#v", extended)
	}
	if rq.UserInformation.UserIdentity == nil || rq.UserInformation.UserIdentity.UserIdentityType != association.UserIdentityTypeJWT ||
		!bytes.Equal(rq.UserInformation.UserIdentity.PrimaryField, []byte("request-token")) {
		t.Fatalf("user identity = %#v", rq.UserInformation.UserIdentity)
	}
	if !client.GetConfig().RequireSuccessfulUserIdentityNegotiation {
		t.Fatal("successful user identity negotiation must be required by default")
	}
}

func TestWithExtendedNegotiationMergesCommonFieldsInEitherOrder(t *testing.T) {
	orders := []struct {
		name             string
		applicationFirst bool
	}{
		{name: "application then common", applicationFirst: true},
		{name: "common then application", applicationFirst: false},
	}

	for _, order := range orders {
		t.Run(order.name, func(t *testing.T) {
			application := association.NewExtendedNegotiation("1.2.3", []byte{1, 0, 1})
			common := association.NewCommonExtendedNegotiation("1.2.3", "4.5", "6.7", "8.9")
			options := []Option{WithExtendedNegotiation(application), WithExtendedNegotiation(common)}
			if !order.applicationFirst {
				options[0], options[1] = options[1], options[0]
			}
			client := New(options...)

			application.RequestedApplicationInfo[0] = 9
			common.RelatedGeneralSOPClassUIDs[0] = "changed"
			rq := client.buildAssociateRQ()

			if got := rq.UserInformation.ExtendedNegotiations; len(got) != 1 ||
				!bytes.Equal(got[0].ServiceClassAppInfo, []byte{1, 0, 1}) {
				t.Fatalf("extended negotiations = %#v", got)
			}
			if got := rq.UserInformation.CommonExtendedNegotiations; len(got) != 1 ||
				got[0].ServiceClassUID != "4.5" || len(got[0].RelatedGeneralSOPClassUIDs) != 2 ||
				got[0].RelatedGeneralSOPClassUIDs[0] != "6.7" || got[0].RelatedGeneralSOPClassUIDs[1] != "8.9" {
				t.Fatalf("common extended negotiations = %#v", got)
			}
			if len(client.GetConfig().ExtendedNegotiations) != 1 {
				t.Fatalf("configured negotiations = %#v", client.GetConfig().ExtendedNegotiations)
			}
		})
	}
}

func TestWithExtendedNegotiationDoesNotDiscardEmptyCommonServiceClassUID(t *testing.T) {
	client := New(WithExtendedNegotiation(
		association.NewCommonExtendedNegotiation("1.2.3", ""),
	))

	rq := client.buildAssociateRQ()
	if got := rq.UserInformation.CommonExtendedNegotiations; len(got) != 1 {
		t.Fatalf("common extended negotiations = %#v, want explicit invalid request preserved", got)
	}
	if _, err := rq.Encode(); err == nil {
		t.Fatal("Encode() error = nil, want empty Service Class UID rejection")
	}
}

func TestBuildAcceptedAssociationRequiresRequestedIdentityResponse(t *testing.T) {
	newPair := func(require bool, response *pdu.UserIdentityNegotiationResponse) (*Client, *pdu.AAssociateRQ, *pdu.AAssociateAC) {
		client := New(
			WithUserIdentity(association.NewUserIdentityJWT([]byte("sensitive-request-token"), true)),
			WithRequireSuccessfulUserIdentityNegotiation(require),
		)
		if err := client.AddPresentationContext(verificationSOPClassUID, "1.2.840.10008.1.2.1"); err != nil {
			t.Fatalf("AddPresentationContext() error = %v", err)
		}
		rq := client.buildAssociateRQ()
		ac := pdu.NewAAssociateAC()
		ac.PresentationContexts = []pdu.PresentationContextAC{{ID: 1, Result: pdu.ResultAcceptance, TransferSyntax: testExplicitVRLittleEndianUID}}
		ac.UserInformation.UserIdentityResponse = response
		return client, rq, ac
	}

	client, rq, ac := newPair(true, nil)
	_, err := client.buildAcceptedAssociation(rq, ac)
	if err == nil || !strings.Contains(err.Error(), "positive user identity response") {
		t.Fatalf("buildAcceptedAssociation() error = %v", err)
	}
	if strings.Contains(err.Error(), "sensitive-request-token") {
		t.Fatalf("identity negotiation error exposed request payload: %v", err)
	}

	client, rq, ac = newPair(false, nil)
	if _, err := client.buildAcceptedAssociation(rq, ac); err != nil {
		t.Fatalf("optional identity response was rejected: %v", err)
	}

	client, rq, ac = newPair(true, &pdu.UserIdentityNegotiationResponse{ServerResponse: []byte{}})
	assoc, err := client.buildAcceptedAssociation(rq, ac)
	if err != nil {
		t.Fatalf("present empty identity response was rejected: %v", err)
	}
	if assoc.UserIdentity == nil || assoc.UserIdentity.ServerResponse == nil {
		t.Fatal("present empty identity response was not retained")
	}
}

func TestBuildAssociateRQ(t *testing.T) {
	client := New(
		WithCallingAE("MY_SCU"),
		WithCalledAE("REMOTE_SCP"),
		WithMaxPDULength(16384),
	)

	// Add presentation context
	if err := client.AddPresentationContext(verificationSOPClassUID, "1.2.840.10008.1.2"); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}

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
	err := client.Connect(context.TODO(), "localhost", 104)
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

	err := client.Connect(context.TODO(), "localhost", 104)
	if err == nil {
		t.Error("Expected error when already connected")
	}

	expectedMsg := "client is already connected"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestConnectRejectsConcurrentAttemptAndCloseCancelsConnectingAttempt(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	t.Cleanup(func() { _ = listener.Close() })
	accepted := make(chan net.Conn, 1)
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			accepted <- conn
		}
	}()

	client := New(WithAssociationTimeout(5 * time.Second))
	if err := client.AddPresentationContext("1.2.840.10008.1.1", "1.2.840.10008.1.2.1"); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}
	connectResult := make(chan error, 1)
	go func() {
		port := listener.Addr().(*net.TCPAddr).Port
		connectResult <- client.Connect(context.Background(), "127.0.0.1", port)
	}()

	select {
	case conn := <-accepted:
		t.Cleanup(func() { _ = conn.Close() })
	case <-time.After(time.Second):
		t.Fatal("client did not begin the first connection attempt")
	}
	concurrentResults := make(chan error, 50)
	for range 50 {
		go func() { concurrentResults <- client.Connect(context.Background(), "127.0.0.1", 1) }()
	}
	for range 50 {
		if err := <-concurrentResults; !errors.Is(err, ErrClientConnecting) {
			t.Fatalf("concurrent Connect() error = %v, want ErrClientConnecting", err)
		}
	}
	if err := client.AddPresentationContext("1.2.840.10008.5.1.4.1.1.2", "1.2.840.10008.1.2.1"); !errors.Is(err, ErrClientConnecting) {
		t.Fatalf("AddPresentationContext() during Connect error = %v, want ErrClientConnecting", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("Close() during Connect error = %v", err)
	}
	select {
	case err := <-connectResult:
		if err == nil {
			t.Fatal("connecting attempt succeeded after Close()")
		}
	case <-time.After(time.Second):
		t.Fatal("Close() did not cancel the connecting attempt")
	}
}

func TestHandleServiceClosedOnlyClearsCurrentSession(t *testing.T) {
	client := New()
	current := &mockServiceForDIMSE{}
	old := &mockServiceForDIMSE{}
	client.mu.Lock()
	client.service = current
	client.assoc = &association.Association{}
	client.connected = true
	client.state = clientConnected
	client.mu.Unlock()

	client.handleServiceClosed(old)
	if !client.IsConnected() {
		t.Fatal("old service closure cleared the current connection")
	}

	client.handleServiceClosed(current)
	if client.IsConnected() {
		t.Fatal("current service closure left client connected")
	}
	if association := client.GetAssociation(); association != nil {
		t.Fatalf("association = %#v after service closure, want nil", association)
	}
}
