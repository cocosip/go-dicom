// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"testing"
)

func TestNewAAssociateRQ(t *testing.T) {
	rq := NewAAssociateRQ()

	if rq.ProtocolVersion != 0x0001 {
		t.Errorf("Expected protocol version 0x0001, got 0x%04X", rq.ProtocolVersion)
	}

	if rq.ApplicationContext != "1.2.840.10008.3.1.1.1" {
		t.Errorf("Expected default application context, got %s", rq.ApplicationContext)
	}

	if rq.UserInformation == nil {
		t.Error("Expected UserInformation to be initialized")
	}

	if rq.UserInformation.MaximumLength != 16384 {
		t.Errorf("Expected default MaximumLength 16384, got %d", rq.UserInformation.MaximumLength)
	}
}

func TestAAssociateRQ_EncodeDecodeBasic(t *testing.T) {
	// Create a basic A-ASSOCIATE-RQ
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "CALLED_AE"
	rq.CallingAETitle = "CALLING_AE"

	// Add a presentation context
	rq.PresentationContexts = []PresentationContextRQ{
		{
			ID:             1,
			AbstractSyntax: "1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
			TransferSyntaxes: []string{
				"1.2.840.10008.1.2.1", // Explicit VR Little Endian
				"1.2.840.10008.1.2",   // Implicit VR Little Endian
			},
		},
	}

	// Encode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypeAAssociateRQ {
		t.Errorf("Expected PDU type 0x01, got 0x%02X", pdu.Type)
	}

	// Decode
	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Compare
	if decoded.ProtocolVersion != rq.ProtocolVersion {
		t.Errorf("ProtocolVersion mismatch: expected 0x%04X, got 0x%04X", rq.ProtocolVersion, decoded.ProtocolVersion)
	}

	if decoded.CalledAETitle != rq.CalledAETitle {
		t.Errorf("CalledAETitle mismatch: expected %s, got %s", rq.CalledAETitle, decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != rq.CallingAETitle {
		t.Errorf("CallingAETitle mismatch: expected %s, got %s", rq.CallingAETitle, decoded.CallingAETitle)
	}

	if decoded.ApplicationContext != rq.ApplicationContext {
		t.Errorf("ApplicationContext mismatch: expected %s, got %s", rq.ApplicationContext, decoded.ApplicationContext)
	}

	if len(decoded.PresentationContexts) != len(rq.PresentationContexts) {
		t.Fatalf("PresentationContexts count mismatch: expected %d, got %d", len(rq.PresentationContexts), len(decoded.PresentationContexts))
	}

	pc := decoded.PresentationContexts[0]
	if pc.ID != 1 {
		t.Errorf("PresentationContext ID mismatch: expected 1, got %d", pc.ID)
	}

	if pc.AbstractSyntax != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("AbstractSyntax mismatch: expected 1.2.840.10008.5.1.4.1.1.2, got %s", pc.AbstractSyntax)
	}

	if len(pc.TransferSyntaxes) != 2 {
		t.Fatalf("TransferSyntaxes count mismatch: expected 2, got %d", len(pc.TransferSyntaxes))
	}

	if decoded.UserInformation == nil {
		t.Fatal("UserInformation is nil")
	}

	if decoded.UserInformation.MaximumLength != rq.UserInformation.MaximumLength {
		t.Errorf("MaximumLength mismatch: expected %d, got %d", rq.UserInformation.MaximumLength, decoded.UserInformation.MaximumLength)
	}
}

func TestAAssociateRQ_EncodeDecodeMultiplePresentationContexts(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SCP_AE"
	rq.CallingAETitle = "SCU_AE"

	// Add multiple presentation contexts
	rq.PresentationContexts = []PresentationContextRQ{
		{
			ID:             1,
			AbstractSyntax: "1.2.840.10008.5.1.4.1.1.2",
			TransferSyntaxes: []string{
				"1.2.840.10008.1.2.1",
			},
		},
		{
			ID:             3,
			AbstractSyntax: "1.2.840.10008.5.1.4.1.1.4",
			TransferSyntaxes: []string{
				"1.2.840.10008.1.2",
				"1.2.840.10008.1.2.1",
			},
		},
		{
			ID:             5,
			AbstractSyntax: "1.2.840.10008.1.1", // Verification SOP Class
			TransferSyntaxes: []string{
				"1.2.840.10008.1.2",
			},
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.PresentationContexts) != 3 {
		t.Fatalf("Expected 3 presentation contexts, got %d", len(decoded.PresentationContexts))
	}

	// Check each context
	for i, expected := range rq.PresentationContexts {
		got := decoded.PresentationContexts[i]
		if got.ID != expected.ID {
			t.Errorf("Context %d: ID mismatch: expected %d, got %d", i, expected.ID, got.ID)
		}
		if got.AbstractSyntax != expected.AbstractSyntax {
			t.Errorf("Context %d: AbstractSyntax mismatch", i)
		}
		if len(got.TransferSyntaxes) != len(expected.TransferSyntaxes) {
			t.Errorf("Context %d: TransferSyntaxes count mismatch", i)
		}
	}
}

func TestAAssociateRQ_EncodeDecodeWithUserIdentity(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SERVER"
	rq.CallingAETitle = "CLIENT"

	// Add user identity (username only)
	rq.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType:          1, // Username
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("testuser"),
		SecondaryField:            nil,
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify user identity
	if decoded.UserInformation.UserIdentity == nil {
		t.Fatal("UserIdentity is nil after decode")
	}

	ui := decoded.UserInformation.UserIdentity
	if ui.UserIdentityType != 1 {
		t.Errorf("UserIdentityType mismatch: expected 1, got %d", ui.UserIdentityType)
	}

	if ui.PositiveResponseRequested != 1 {
		t.Errorf("PositiveResponseRequested mismatch: expected 1, got %d", ui.PositiveResponseRequested)
	}

	if !bytes.Equal(ui.PrimaryField, []byte("testuser")) {
		t.Errorf("PrimaryField mismatch: expected 'testuser', got %s", string(ui.PrimaryField))
	}
}

func TestAAssociateRQ_EncodeDecodeWithUsernamePassword(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SERVER"
	rq.CallingAETitle = "CLIENT"

	// Add user identity (username + password)
	rq.UserInformation.UserIdentity = &UserIdentityNegotiation{
		UserIdentityType:          2, // Username + Password
		PositiveResponseRequested: 0,
		PrimaryField:              []byte("admin"),
		SecondaryField:            []byte("password123"),
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	ui := decoded.UserInformation.UserIdentity
	if ui == nil {
		t.Fatal("UserIdentity is nil")
	}

	if ui.UserIdentityType != 2 {
		t.Errorf("UserIdentityType mismatch")
	}

	if !bytes.Equal(ui.PrimaryField, []byte("admin")) {
		t.Errorf("PrimaryField mismatch")
	}

	if !bytes.Equal(ui.SecondaryField, []byte("password123")) {
		t.Errorf("SecondaryField mismatch")
	}
}

func TestAAssociateRQ_EncodeDecodeWithExtendedNegotiation(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SERVER"
	rq.CallingAETitle = "CLIENT"

	// Add extended negotiation
	rq.UserInformation.ExtendedNegotiations = []ExtendedNegotiation{
		{
			SOPClassUID:         "1.2.840.10008.5.1.4.1.1.2",
			ServiceClassAppInfo: []byte{0x01, 0x02, 0x03},
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.UserInformation.ExtendedNegotiations) != 1 {
		t.Fatalf("Expected 1 extended negotiation, got %d", len(decoded.UserInformation.ExtendedNegotiations))
	}

	ext := decoded.UserInformation.ExtendedNegotiations[0]
	if ext.SOPClassUID != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("SOPClassUID mismatch")
	}

	if !bytes.Equal(ext.ServiceClassAppInfo, []byte{0x01, 0x02, 0x03}) {
		t.Errorf("ServiceClassAppInfo mismatch")
	}
}

func TestAAssociateRQ_EncodeDecodeWithRoleSelection(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SERVER"
	rq.CallingAETitle = "CLIENT"

	// Add role selection
	rq.UserInformation.SCPSCURoleSelections = []SCPSCURoleSelection{
		{
			SOPClassUID: "1.2.840.10008.5.1.4.1.1.2",
			SCURole:     1, // Support SCU
			SCPRole:     0, // No SCP support
		},
	}

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.UserInformation.SCPSCURoleSelections) != 1 {
		t.Fatalf("Expected 1 role selection, got %d", len(decoded.UserInformation.SCPSCURoleSelections))
	}

	role := decoded.UserInformation.SCPSCURoleSelections[0]
	if role.SOPClassUID != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("SOPClassUID mismatch")
	}

	if role.SCURole != 1 {
		t.Errorf("SCURole mismatch: expected 1, got %d", role.SCURole)
	}

	if role.SCPRole != 0 {
		t.Errorf("SCPRole mismatch: expected 0, got %d", role.SCPRole)
	}
}

func TestAAssociateRQ_EncodeDecodeWithImplementationVersion(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SERVER"
	rq.CallingAETitle = "CLIENT"
	rq.UserInformation.ImplementationVersionName = "GO_DICOM_1_0"

	// Encode and decode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if decoded.UserInformation.ImplementationVersionName != "GO_DICOM_1_0" {
		t.Errorf("ImplementationVersionName mismatch: expected GO_DICOM_1_0, got %s", decoded.UserInformation.ImplementationVersionName)
	}
}

func TestAAssociateRQ_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateAC, []byte{0x00})

	rq := &AAssociateRQ{}
	err := rq.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestAETitleSpacePadding(t *testing.T) {
	rq := NewAAssociateRQ()
	rq.CalledAETitle = "SHORT"    // Only 5 chars
	rq.CallingAETitle = "ANOTHER" // 7 chars

	// Encode
	pdu, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Decode
	decoded := &AAssociateRQ{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify no extra spaces
	if decoded.CalledAETitle != "SHORT" {
		t.Errorf("CalledAETitle should trim spaces: expected 'SHORT', got '%s'", decoded.CalledAETitle)
	}

	if decoded.CallingAETitle != "ANOTHER" {
		t.Errorf("CallingAETitle should trim spaces: expected 'ANOTHER', got '%s'", decoded.CallingAETitle)
	}
}
