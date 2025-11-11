// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"testing"
)

const (
	testServerAETitle      = "SERVER"
	testClientAETitle      = "CLIENT"
	testImplVersionName    = "GO_DICOM_1_0"
	testShortAETitle       = "SHORT"
	testAnotherAETitle     = "ANOTHER"
	testExplicitVRLittleLE = "1.2.840.10008.1.2.1"
)

func TestNewAAssociateAC(t *testing.T) {
	ac := NewAAssociateAC()

	if ac.ProtocolVersion != 0x0001 {
		t.Errorf("Expected protocol version 0x0001, got 0x%04X", ac.ProtocolVersion)
	}

	if ac.ApplicationContext != "1.2.840.10008.3.1.1.1" {
		t.Errorf("Expected default application context, got %s", ac.ApplicationContext)
	}

	if ac.UserInformation == nil {
		t.Error("Expected UserInformation to be initialized")
	}

	if ac.UserInformation.MaximumLength != 16384 {
		t.Errorf("Expected default MaximumLength 16384, got %d", ac.UserInformation.MaximumLength)
	}
}

func TestAAssociateAC_EncodeDecodeBasic(t *testing.T) {
	// Create a basic A-ASSOCIATE-AC
	ac := NewAAssociateAC()
	ac.CalledAETitle = "CALLED_AE"
	ac.CallingAETitle = "CALLING_AE"

	// Add accepted presentation context
	ac.PresentationContexts = []PresentationContextAC{
		{
			ID:             1,
			Result:         ResultAcceptance,
			TransferSyntax: testExplicitVRLittleLE, // Explicit VR Little Endian
		},
	}

	// Encode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypeAAssociateAC {
		t.Errorf("Expected PDU type 0x02, got 0x%02X", pdu.Type)
	}

	// Decode
	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Compare
	if decoded.ProtocolVersion != ac.ProtocolVersion {
		t.Errorf("ProtocolVersion mismatch: expected 0x%04X, got 0x%04X", ac.ProtocolVersion, decoded.ProtocolVersion)
	}

	if decoded.CalledAETitle != ac.CalledAETitle {
		t.Errorf("CalledAETitle mismatch: expected %s, got %s", ac.CalledAETitle, decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != ac.CallingAETitle {
		t.Errorf("CallingAETitle mismatch: expected %s, got %s", ac.CallingAETitle, decoded.CallingAETitle)
	}

	if decoded.ApplicationContext != ac.ApplicationContext {
		t.Errorf("ApplicationContext mismatch: expected %s, got %s", ac.ApplicationContext, decoded.ApplicationContext)
	}

	if len(decoded.PresentationContexts) != 1 {
		t.Fatalf("Expected 1 presentation context, got %d", len(decoded.PresentationContexts))
	}

	pc := decoded.PresentationContexts[0]
	if pc.ID != 1 {
		t.Errorf("PresentationContext ID mismatch: expected 1, got %d", pc.ID)
	}

	if pc.Result != ResultAcceptance {
		t.Errorf("PresentationContext Result mismatch: expected %d, got %d", ResultAcceptance, pc.Result)
	}

	if pc.TransferSyntax != testExplicitVRLittleLE {
		t.Errorf("TransferSyntax mismatch: expected 1.2.840.10008.1.2.1, got %s", pc.TransferSyntax)
	}

	if decoded.UserInformation.MaximumLength != ac.UserInformation.MaximumLength {
		t.Errorf("MaximumLength mismatch: expected %d, got %d", ac.UserInformation.MaximumLength, decoded.UserInformation.MaximumLength)
	}
}

func TestAAssociateAC_EncodeDecodeMultiplePresentationContexts(t *testing.T) {
	ac := NewAAssociateAC()
	ac.CalledAETitle = "SCP_AE"
	ac.CallingAETitle = "SCU_AE"

	// Add multiple presentation contexts with different results
	ac.PresentationContexts = []PresentationContextAC{
		{
			ID:             1,
			Result:         ResultAcceptance,
			TransferSyntax: testExplicitVRLittleLE,
		},
		{
			ID:     3,
			Result: ResultAbstractSyntaxNotSupported,
		},
		{
			ID:             5,
			Result:         ResultAcceptance,
			TransferSyntax: "1.2.840.10008.1.2",
		},
	}

	// Encode and decode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.PresentationContexts) != 3 {
		t.Fatalf("Expected 3 presentation contexts, got %d", len(decoded.PresentationContexts))
	}

	// Check accepted context
	if decoded.PresentationContexts[0].Result != ResultAcceptance {
		t.Errorf("Context 0: expected acceptance")
	}
	if decoded.PresentationContexts[0].TransferSyntax != testExplicitVRLittleLE {
		t.Errorf("Context 0: wrong transfer syntax")
	}

	// Check rejected context
	if decoded.PresentationContexts[1].Result != ResultAbstractSyntaxNotSupported {
		t.Errorf("Context 1: expected abstract syntax not supported")
	}
	if decoded.PresentationContexts[1].TransferSyntax != "" {
		t.Errorf("Context 1: rejected context should have no transfer syntax")
	}
}

func TestAAssociateAC_ResultString(t *testing.T) {
	tests := []struct {
		result   byte
		expected string
	}{
		{ResultAcceptance, "Acceptance"},
		{ResultUserRejection, "User Rejection"},
		{ResultNoReason, "No Reason (Provider Rejection)"},
		{ResultAbstractSyntaxNotSupported, "Abstract Syntax Not Supported (Provider Rejection)"},
		{ResultTransferSyntaxesNotSupported, "Transfer Syntaxes Not Supported (Provider Rejection)"},
		{0xFF, "Unknown(0xFF)"},
	}

	for _, tt := range tests {
		result := ResultString(tt.result)
		if result != tt.expected {
			t.Errorf("ResultString(0x%02X): expected %s, got %s", tt.result, tt.expected, result)
		}
	}
}

func TestAAssociateAC_EncodeDecodeWithImplementationVersion(t *testing.T) {
	ac := NewAAssociateAC()
	ac.CalledAETitle = testServerAETitle
	ac.CallingAETitle = testClientAETitle
	ac.UserInformation.ImplementationVersionName = testImplVersionName

	// Encode and decode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if decoded.UserInformation.ImplementationVersionName != testImplVersionName {
		t.Errorf("ImplementationVersionName mismatch: expected GO_DICOM_1_0, got %s", decoded.UserInformation.ImplementationVersionName)
	}
}

func TestAAssociateAC_EncodeDecodeWithUserIdentityResponse(t *testing.T) {
	ac := NewAAssociateAC()
	ac.CalledAETitle = testServerAETitle
	ac.CallingAETitle = testClientAETitle

	// Add user identity response (e.g., Kerberos token)
	ac.UserInformation.UserIdentity = &UserIdentityNegotiation{
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("server-response-token"),
	}

	// Encode and decode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify user identity response
	if decoded.UserInformation.UserIdentity == nil {
		t.Fatal("UserIdentity is nil after decode")
	}

	ui := decoded.UserInformation.UserIdentity
	if ui.PositiveResponseRequested != 1 {
		t.Errorf("PositiveResponseRequested mismatch: expected 1, got %d", ui.PositiveResponseRequested)
	}

	if !bytes.Equal(ui.PrimaryField, []byte("server-response-token")) {
		t.Errorf("PrimaryField mismatch: expected 'server-response-token', got %s", string(ui.PrimaryField))
	}
}

func TestAAssociateAC_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateRQ, []byte{0x00})

	ac := &AAssociateAC{}
	err := ac.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAAssociateAC_RoundtripWithRQ(t *testing.T) {
	// Simulate a complete association negotiation

	// 1. Create A-ASSOCIATE-RQ
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "PACS_SERVER"
	rq.CallingAETitle = "MY_SCU"
	rq.PresentationContexts = []PresentationContextRQ{
		{
			ID:             1,
			AbstractSyntax: "1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
			TransferSyntaxes: []string{
				testExplicitVRLittleLE, // Explicit VR Little Endian
				"1.2.840.10008.1.2",   // Implicit VR Little Endian
			},
		},
		{
			ID:             3,
			AbstractSyntax: "1.2.840.10008.1.1", // Verification SOP Class
			TransferSyntaxes: []string{
				"1.2.840.10008.1.2",
			},
		},
	}

	// 2. Create A-ASSOCIATE-AC response
	ac := NewAAssociateAC()
	ac.CalledAETitle = rq.CalledAETitle
	ac.CallingAETitle = rq.CallingAETitle
	ac.ApplicationContext = rq.ApplicationContext

	// Accept context 1, reject context 3
	ac.PresentationContexts = []PresentationContextAC{
		{
			ID:             1,
			Result:         ResultAcceptance,
			TransferSyntax: testExplicitVRLittleLE, // Choose first transfer syntax
		},
		{
			ID:     3,
			Result: ResultAbstractSyntaxNotSupported,
		},
	}

	// 3. Encode AC
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// 4. Decode AC
	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// 5. Verify negotiation results
	if len(decoded.PresentationContexts) != 2 {
		t.Fatalf("Expected 2 presentation contexts, got %d", len(decoded.PresentationContexts))
	}

	// Check accepted context
	accepted := decoded.PresentationContexts[0]
	if accepted.ID != 1 {
		t.Errorf("Expected context ID 1, got %d", accepted.ID)
	}
	if accepted.Result != ResultAcceptance {
		t.Errorf("Expected acceptance, got result %d", accepted.Result)
	}
	if accepted.TransferSyntax != testExplicitVRLittleLE {
		t.Errorf("Wrong transfer syntax: %s", accepted.TransferSyntax)
	}

	// Check rejected context
	rejected := decoded.PresentationContexts[1]
	if rejected.ID != 3 {
		t.Errorf("Expected context ID 3, got %d", rejected.ID)
	}
	if rejected.Result != ResultAbstractSyntaxNotSupported {
		t.Errorf("Expected abstract syntax not supported, got result %d", rejected.Result)
	}
	if rejected.TransferSyntax != "" {
		t.Errorf("Rejected context should have empty transfer syntax, got %s", rejected.TransferSyntax)
	}
}

func TestAAssociateAC_AETitleSpacePadding(t *testing.T) {
	ac := NewAAssociateAC()
	ac.CalledAETitle = testShortAETitle    // Only 5 chars
	ac.CallingAETitle = testAnotherAETitle // 7 chars

	// Encode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Decode
	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify no extra spaces
	if decoded.CalledAETitle != testShortAETitle {
		t.Errorf("CalledAETitle should trim spaces: expected 'SHORT', got '%s'", decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != testAnotherAETitle {
		t.Errorf("CallingAETitle should trim spaces: expected 'ANOTHER', got '%s'", decoded.CallingAETitle)
	}
}

func TestAAssociateAC_AllRejectionReasons(t *testing.T) {
	// Test all possible rejection reasons
	ac := NewAAssociateAC()
	ac.CalledAETitle = testServerAETitle
	ac.CallingAETitle = testClientAETitle

	ac.PresentationContexts = []PresentationContextAC{
		{ID: 1, Result: ResultUserRejection},
		{ID: 3, Result: ResultNoReason},
		{ID: 5, Result: ResultAbstractSyntaxNotSupported},
		{ID: 7, Result: ResultTransferSyntaxesNotSupported},
	}

	// Encode and decode
	pdu, err := ac.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateAC{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify all rejection reasons are preserved
	if len(decoded.PresentationContexts) != 4 {
		t.Fatalf("Expected 4 contexts, got %d", len(decoded.PresentationContexts))
	}

	expected := []byte{
		ResultUserRejection,
		ResultNoReason,
		ResultAbstractSyntaxNotSupported,
		ResultTransferSyntaxesNotSupported,
	}

	for i, pc := range decoded.PresentationContexts {
		if pc.Result != expected[i] {
			t.Errorf("Context %d: expected result %d (%s), got %d (%s)",
				i, expected[i], ResultString(expected[i]), pc.Result, ResultString(pc.Result))
		}
		// Rejected contexts should have no transfer syntax
		if pc.TransferSyntax != "" {
			t.Errorf("Context %d: rejected context should have empty transfer syntax", i)
		}
	}
}
