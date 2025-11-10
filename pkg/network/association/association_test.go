// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package association

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

func TestNewAssociation(t *testing.T) {
	assoc := NewAssociation("MY_SCU", "PACS_SERVER")

	if assoc.CallingAE != "MY_SCU" {
		t.Errorf("Expected CallingAE MY_SCU, got %s", assoc.CallingAE)
	}

	if assoc.CalledAE != "PACS_SERVER" {
		t.Errorf("Expected CalledAE PACS_SERVER, got %s", assoc.CalledAE)
	}

	if assoc.ProtocolVersion != 0x0001 {
		t.Errorf("Expected protocol version 0x0001, got 0x%04X", assoc.ProtocolVersion)
	}

	if assoc.MaxPDULength != 16384 {
		t.Errorf("Expected MaxPDULength 16384, got %d", assoc.MaxPDULength)
	}

	if assoc.IsEstablished {
		t.Error("Expected IsEstablished to be false")
	}

	if len(assoc.PresentationContexts) != 0 {
		t.Errorf("Expected 0 presentation contexts, got %d", len(assoc.PresentationContexts))
	}
}

func TestAssociation_AddPresentationContext(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	// Add valid context (odd ID)
	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	if err := assoc.AddPresentationContext(pc1); err != nil {
		t.Fatalf("Failed to add presentation context: %v", err)
	}

	if len(assoc.PresentationContexts) != 1 {
		t.Errorf("Expected 1 presentation context, got %d", len(assoc.PresentationContexts))
	}

	// Try to add context with even ID (should fail)
	pc2 := NewPresentationContext(2, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian)
	if err := assoc.AddPresentationContext(pc2); err == nil {
		t.Error("Expected error for even presentation context ID")
	}

	// Try to add duplicate ID (should fail)
	pc3 := NewPresentationContext(1, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian)
	if err := assoc.AddPresentationContext(pc3); err == nil {
		t.Error("Expected error for duplicate presentation context ID")
	}
}

func TestAssociation_FindPresentationContextByID(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc3 := NewPresentationContext(3, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian)

	assoc.AddPresentationContext(pc1)
	assoc.AddPresentationContext(pc3)

	// Find existing context
	found := assoc.FindPresentationContextByID(1)
	if found == nil {
		t.Fatal("Expected to find context ID 1")
	}
	if found.AbstractSyntax != "1.2.840.10008.5.1.4.1.1.2" {
		t.Errorf("Wrong abstract syntax: %s", found.AbstractSyntax)
	}

	// Find non-existing context
	notFound := assoc.FindPresentationContextByID(5)
	if notFound != nil {
		t.Error("Expected nil for non-existing context ID")
	}
}

func TestAssociation_FindPresentationContextByAbstractSyntax(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc1.Accept(transfer.ExplicitVRLittleEndian)

	pc3 := NewPresentationContext(3, "1.2.840.10008.1.1", transfer.ImplicitVRLittleEndian)
	pc3.Reject(ResultAbstractSyntaxNotSupported)

	assoc.AddPresentationContext(pc1)
	assoc.AddPresentationContext(pc3)

	// Find accepted context
	found := assoc.FindPresentationContextByAbstractSyntax("1.2.840.10008.5.1.4.1.1.2")
	if found == nil {
		t.Error("Expected to find accepted context")
	}

	// Rejected context should not be found
	notFound := assoc.FindPresentationContextByAbstractSyntax("1.2.840.10008.1.1")
	if notFound != nil {
		t.Error("Expected nil for rejected context")
	}
}

func TestAssociation_GetAcceptedPresentationContexts(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc1.Accept(transfer.ExplicitVRLittleEndian)

	pc3 := NewPresentationContext(3, "1.2.840.10008.1.1", transfer.ImplicitVRLittleEndian)
	pc3.Reject(ResultUserRejection)

	pc5 := NewPresentationContext(5, "1.2.840.10008.5.1.4.1.1.4", transfer.ExplicitVRLittleEndian)
	pc5.Accept(transfer.ExplicitVRLittleEndian)

	assoc.AddPresentationContext(pc1)
	assoc.AddPresentationContext(pc3)
	assoc.AddPresentationContext(pc5)

	accepted := assoc.GetAcceptedPresentationContexts()
	if len(accepted) != 2 {
		t.Errorf("Expected 2 accepted contexts, got %d", len(accepted))
	}
}

func TestAssociation_ExtendedNegotiation(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	en1 := NewExtendedNegotiation("1.2.840.10008.5.1.4.1.1.2", []byte{0x01, 0x02, 0x03})
	en2 := NewExtendedNegotiation("1.2.840.10008.1.1", []byte{0x04, 0x05})

	assoc.AddExtendedNegotiation(en1)
	assoc.AddExtendedNegotiation(en2)

	// Find existing
	found := assoc.FindExtendedNegotiation("1.2.840.10008.5.1.4.1.1.2")
	if found == nil {
		t.Fatal("Expected to find extended negotiation")
	}
	if len(found.ServiceClassAppInfo) != 3 {
		t.Errorf("Expected 3 bytes of app info, got %d", len(found.ServiceClassAppInfo))
	}

	// Find non-existing
	notFound := assoc.FindExtendedNegotiation("1.2.3.4.5")
	if notFound != nil {
		t.Error("Expected nil for non-existing extended negotiation")
	}
}

func TestAssociation_RoleSelection(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	rs1 := NewRoleSelection("1.2.840.10008.5.1.4.1.1.2", 1, 0) // SCU only
	rs2 := NewRoleSelection("1.2.840.10008.1.1", 1, 1)         // Both

	assoc.AddRoleSelection(rs1)
	assoc.AddRoleSelection(rs2)

	// Find existing
	found := assoc.FindRoleSelection("1.2.840.10008.5.1.4.1.1.2")
	if found == nil {
		t.Fatal("Expected to find role selection")
	}
	if found.SCURole != 1 || found.SCPRole != 0 {
		t.Errorf("Wrong role selection: SCU=%d, SCP=%d", found.SCURole, found.SCPRole)
	}

	// Find non-existing
	notFound := assoc.FindRoleSelection("1.2.3.4.5")
	if notFound != nil {
		t.Error("Expected nil for non-existing role selection")
	}
}

func TestAssociation_SetEstablished(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	if assoc.IsEstablished {
		t.Error("Expected IsEstablished to be false initially")
	}

	assoc.SetEstablished(true)
	if !assoc.IsEstablished {
		t.Error("Expected IsEstablished to be true")
	}

	assoc.SetEstablished(false)
	if assoc.IsEstablished {
		t.Error("Expected IsEstablished to be false")
	}
}

func TestAssociation_String(t *testing.T) {
	assoc := NewAssociation("MY_SCU", "PACS_SERVER")

	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc1.Accept(transfer.ExplicitVRLittleEndian)

	pc3 := NewPresentationContext(3, "1.2.840.10008.1.1", transfer.ImplicitVRLittleEndian)
	pc3.Reject(ResultUserRejection)

	assoc.AddPresentationContext(pc1)
	assoc.AddPresentationContext(pc3)

	str := assoc.String()
	expected := "Association[MY_SCU -> PACS_SERVER, 1/2 contexts accepted, MaxPDU=16384]"

	if str != expected {
		t.Errorf("String() mismatch:\nexpected: %s\ngot:      %s", expected, str)
	}
}

func TestPresentationContext_AcceptReject(t *testing.T) {
	pc := NewPresentationContext(1, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian)

	// Initially should be accepted (default)
	if !pc.IsAccepted() {
		t.Error("Expected initial state to be accepted")
	}

	// Reject
	pc.Reject(ResultUserRejection)
	if pc.IsAccepted() {
		t.Error("Expected context to be rejected")
	}
	if pc.Result != ResultUserRejection {
		t.Errorf("Expected result %d, got %d", ResultUserRejection, pc.Result)
	}
	if pc.AcceptedTransferSyntax != nil {
		t.Error("Expected AcceptedTransferSyntax to be nil after rejection")
	}

	// Accept
	pc.Accept(transfer.ExplicitVRLittleEndian)
	if !pc.IsAccepted() {
		t.Error("Expected context to be accepted")
	}
	if pc.AcceptedTransferSyntax != transfer.ExplicitVRLittleEndian {
		t.Error("Wrong accepted transfer syntax")
	}
}

func TestNewUserIdentityUsername(t *testing.T) {
	ui := NewUserIdentityUsername("testuser", true)

	if ui.Type != UserIdentityTypeUsername {
		t.Errorf("Expected type %d, got %d", UserIdentityTypeUsername, ui.Type)
	}

	if !ui.PositiveResponseRequested {
		t.Error("Expected PositiveResponseRequested to be true")
	}

	if string(ui.PrimaryField) != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", string(ui.PrimaryField))
	}
}

func TestNewUserIdentityUsernamePassword(t *testing.T) {
	ui := NewUserIdentityUsernamePassword("admin", "password123")

	if ui.Type != UserIdentityTypeUsernamePassword {
		t.Errorf("Expected type %d, got %d", UserIdentityTypeUsernamePassword, ui.Type)
	}

	if string(ui.PrimaryField) != "admin" {
		t.Errorf("Expected username 'admin', got %s", string(ui.PrimaryField))
	}

	if string(ui.SecondaryField) != "password123" {
		t.Errorf("Expected password 'password123', got %s", string(ui.SecondaryField))
	}
}

func TestNewAsynchronousOperationsWindow(t *testing.T) {
	async := NewAsynchronousOperationsWindow(5, 10)

	if async.MaxInvokedOperations != 5 {
		t.Errorf("Expected MaxInvokedOperations 5, got %d", async.MaxInvokedOperations)
	}

	if async.MaxPerformedOperations != 10 {
		t.Errorf("Expected MaxPerformedOperations 10, got %d", async.MaxPerformedOperations)
	}
}

func TestAssociation_GetTransferSyntaxForAbstractSyntax(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")

	pc1 := NewPresentationContext(1, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc1.Accept(transfer.ExplicitVRLittleEndian)

	assoc.AddPresentationContext(pc1)

	// Find transfer syntax for accepted context
	ts := assoc.GetTransferSyntaxForAbstractSyntax("1.2.840.10008.5.1.4.1.1.2")
	if ts == nil {
		t.Error("Expected to find transfer syntax")
	}
	if ts != transfer.ExplicitVRLittleEndian {
		t.Error("Wrong transfer syntax")
	}

	// Find transfer syntax for non-existing context
	tsNotFound := assoc.GetTransferSyntaxForAbstractSyntax("1.2.3.4.5")
	if tsNotFound != nil {
		t.Error("Expected nil for non-existing context")
	}
}
