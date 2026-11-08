// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package association

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

const (
	testStorageSOPClassUID         = "1.2.840.10008.5.1.4.1.1.2"
	testExplicitVRLittleEndianUID  = "1.2.840.10008.1.2.1"
	testExtendedSOPClassUID        = "1.2.3"
	testExtendedServiceClassUID    = "4.5"
	testRelatedGeneralSOPClassUID1 = "6.7"
	testRelatedGeneralSOPClassUID2 = "8.9"
)

func TestApplyAAssociateACPreservesRequestedAndAcceptedNegotiation(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.CallingAETitle = "SCU"
	rq.CalledAETitle = "SCP"
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   testStorageSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}
	rq.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   4,
		MaximumNumberOperationsPerformed: 3,
	}
	rq.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: testStorageSOPClassUID,
		SCURole:     1,
		SCPRole:     1,
	}}
	rq.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         testStorageSOPClassUID,
		ServiceClassAppInfo: []byte{1, 1, 0},
	}}
	rq.UserInformation.UserIdentity = &pdu.UserIdentityNegotiation{
		UserIdentityType:          5,
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("request-token"),
	}

	ac := pdu.NewAAssociateAC()
	ac.CallingAETitle = "SCU"
	ac.CalledAETitle = "SCP"
	ac.PresentationContexts = []pdu.PresentationContextAC{{
		ID:             1,
		Result:         pdu.ResultAcceptance,
		TransferSyntax: testExplicitVRLittleEndianUID,
	}}
	ac.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   2,
		MaximumNumberOperationsPerformed: 1,
	}
	ac.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: testStorageSOPClassUID,
		SCURole:     1,
		SCPRole:     0,
	}}
	ac.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         testStorageSOPClassUID,
		ServiceClassAppInfo: []byte{1, 0, 1, 1},
	}}
	ac.UserInformation.UserIdentityResponse = &pdu.UserIdentityNegotiationResponse{
		ServerResponse: []byte("server-token"),
	}

	registry := transfer.NewRegistry()
	assoc := FromAAssociateRQ(rq, registry)
	if err := ApplyAAssociateAC(assoc, ac, registry); err != nil {
		t.Fatalf("ApplyAAssociateAC() error = %v", err)
	}

	if assoc.RequestedAsynchronousOperations == nil ||
		assoc.RequestedAsynchronousOperations.MaxInvokedOperations != 4 ||
		assoc.RequestedAsynchronousOperations.MaxPerformedOperations != 3 {
		t.Fatalf("requested async window = %#v", assoc.RequestedAsynchronousOperations)
	}
	if assoc.AsynchronousOperations == nil || assoc.AsynchronousOperations.MaxInvokedOperations != 2 ||
		assoc.AsynchronousOperations.MaxPerformedOperations != 1 {
		t.Fatalf("accepted async window = %#v", assoc.AsynchronousOperations)
	}

	pc := assoc.FindPresentationContextByID(1)
	if pc == nil || pc.RequestedRole == nil || pc.AcceptedRole == nil {
		t.Fatalf("presentation context roles = requested %#v, accepted %#v", pc.RequestedRole, pc.AcceptedRole)
	}
	if pc.RequestedRole.SCPRole != 1 || pc.AcceptedRole.SCPRole != 0 {
		t.Fatalf("presentation context roles = requested %#v, accepted %#v", pc.RequestedRole, pc.AcceptedRole)
	}
	if pc.AbstractSyntax != testStorageSOPClassUID || len(pc.ProposedTransferSyntaxes) != 1 || pc.AcceptedTransferSyntax == nil {
		t.Fatalf("presentation context proposal/acceptance was not preserved: %#v", pc)
	}

	extended := assoc.FindExtendedNegotiation(testStorageSOPClassUID)
	if extended == nil || !bytes.Equal(extended.RequestedApplicationInfo, []byte{1, 1, 0}) {
		t.Fatalf("requested extended negotiation = %#v", extended)
	}
	if !bytes.Equal(extended.AcceptedApplicationInfo, []byte{1, 0, 1, 1}) {
		t.Fatalf("accepted extended negotiation = %v, want [1 0 1 1]", extended.AcceptedApplicationInfo)
	}

	if assoc.UserIdentity == nil || assoc.UserIdentity.Type != UserIdentityTypeJWT ||
		!bytes.Equal(assoc.UserIdentity.PrimaryField, []byte("request-token")) ||
		!bytes.Equal(assoc.UserIdentity.ServerResponse, []byte("server-token")) {
		t.Fatalf("merged user identity = %#v", assoc.UserIdentity)
	}
}

func TestFromAAssociateRQCombinesCommonAndApplicationNegotiation(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         testExtendedSOPClassUID,
		ServiceClassAppInfo: []byte{1, 0, 1},
	}}
	rq.UserInformation.CommonExtendedNegotiations = []pdu.CommonExtendedNegotiation{{
		SOPClassUID:                testExtendedSOPClassUID,
		ServiceClassUID:            testExtendedServiceClassUID,
		RelatedGeneralSOPClassUIDs: []string{testRelatedGeneralSOPClassUID1, testRelatedGeneralSOPClassUID2},
	}}

	assoc := FromAAssociateRQ(rq, transfer.NewRegistry())
	if len(assoc.ExtendedNegotiations) != 1 {
		t.Fatalf("extended negotiations = %#v, want one combined entry", assoc.ExtendedNegotiations)
	}
	negotiation := assoc.FindExtendedNegotiation(testExtendedSOPClassUID)
	if negotiation == nil || !bytes.Equal(negotiation.RequestedApplicationInfo, []byte{1, 0, 1}) ||
		negotiation.ServiceClassUID != testExtendedServiceClassUID || len(negotiation.RelatedGeneralSOPClassUIDs) != 2 ||
		negotiation.RelatedGeneralSOPClassUIDs[0] != testRelatedGeneralSOPClassUID1 ||
		negotiation.RelatedGeneralSOPClassUIDs[1] != testRelatedGeneralSOPClassUID2 {
		t.Fatalf("combined negotiation = %#v", negotiation)
	}

	rq.UserInformation.ExtendedNegotiations[0].ServiceClassAppInfo[0] = 9
	rq.UserInformation.CommonExtendedNegotiations[0].RelatedGeneralSOPClassUIDs[0] = "changed"
	if negotiation.RequestedApplicationInfo[0] != 1 ||
		negotiation.RelatedGeneralSOPClassUIDs[0] != testRelatedGeneralSOPClassUID1 {
		t.Fatalf("association retained caller-owned slices: %#v", negotiation)
	}
}

func TestFromAAssociateRQUsesTransferSyntaxRegistry(t *testing.T) {
	registry := transfer.NewRegistry()
	replacement := transfer.NewBuilder(transfer.ExplicitVRLittleEndian.UID()).
		SetExplicitVR(true).
		Build()
	if _, err := registry.Replace(replacement); err != nil {
		t.Fatalf("Registry.Replace() error = %v", err)
	}
	rq := pdu.NewAAssociateRQ()
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   testStorageSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}

	assoc := FromAAssociateRQ(rq, registry)
	if got := assoc.PresentationContexts[0].ProposedTransferSyntaxes[0]; got != replacement {
		t.Fatal("FromAAssociateRQ() did not use the configured Transfer Syntax Registry")
	}
}

func TestFromAAssociateACStoresAcceptedApplicationInfo(t *testing.T) {
	ac := pdu.NewAAssociateAC()
	ac.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         testExtendedSOPClassUID,
		ServiceClassAppInfo: []byte{1, 0, 1},
	}}

	assoc := FromAAssociateAC(ac, transfer.NewRegistry())
	negotiation := assoc.FindExtendedNegotiation(testExtendedSOPClassUID)
	if negotiation == nil {
		t.Fatal("accepted extended negotiation was not retained")
	}
	if negotiation.RequestedApplicationInfo != nil {
		t.Fatalf("requested application info = %v, want nil for an AC-only association", negotiation.RequestedApplicationInfo)
	}
	if !bytes.Equal(negotiation.AcceptedApplicationInfo, []byte{1, 0, 1}) {
		t.Fatalf("accepted application info = %v, want [1 0 1]", negotiation.AcceptedApplicationInfo)
	}
	ac.UserInformation.ExtendedNegotiations[0].ServiceClassAppInfo[0] = 9
	if negotiation.AcceptedApplicationInfo[0] != 1 {
		t.Fatal("association retained the AC PDU application-info slice")
	}
}

func TestExtendedNegotiationCloneOwnsRequestedAndAcceptedApplicationInfo(t *testing.T) {
	original := &ExtendedNegotiation{
		SOPClassUID:              testExtendedSOPClassUID,
		RequestedApplicationInfo: []byte{1, 1, 0},
		AcceptedApplicationInfo:  []byte{1, 0, 1},
	}

	clone := original.Clone()
	clone.RequestedApplicationInfo[0] = 9
	clone.AcceptedApplicationInfo[0] = 8

	if !bytes.Equal(original.RequestedApplicationInfo, []byte{1, 1, 0}) {
		t.Fatalf("clone retained requested application-info storage: %v", original.RequestedApplicationInfo)
	}
	if !bytes.Equal(original.AcceptedApplicationInfo, []byte{1, 0, 1}) {
		t.Fatalf("clone retained accepted application-info storage: %v", original.AcceptedApplicationInfo)
	}
}

func TestToAAssociateRQEmitsCombinedExtendedNegotiation(t *testing.T) {
	assoc := NewAssociation("SCU", "SCP")
	assoc.AddExtendedNegotiation(NewExtendedNegotiation(testExtendedSOPClassUID, []byte{1, 0, 1}))
	assoc.AddExtendedNegotiation(NewCommonExtendedNegotiation(
		testExtendedSOPClassUID, testExtendedServiceClassUID, testRelatedGeneralSOPClassUID1, testRelatedGeneralSOPClassUID2,
	))

	rq := ToAAssociateRQ(assoc)
	if got := rq.UserInformation.ExtendedNegotiations; len(got) != 1 ||
		got[0].SOPClassUID != testExtendedSOPClassUID || !bytes.Equal(got[0].ServiceClassAppInfo, []byte{1, 0, 1}) {
		t.Fatalf("extended negotiations = %#v", got)
	}
	if got := rq.UserInformation.CommonExtendedNegotiations; len(got) != 1 ||
		got[0].SOPClassUID != testExtendedSOPClassUID || got[0].ServiceClassUID != testExtendedServiceClassUID ||
		len(got[0].RelatedGeneralSOPClassUIDs) != 2 ||
		got[0].RelatedGeneralSOPClassUIDs[0] != testRelatedGeneralSOPClassUID1 {
		t.Fatalf("common extended negotiations = %#v", got)
	}

	rq.UserInformation.ExtendedNegotiations[0].ServiceClassAppInfo[0] = 9
	rq.UserInformation.CommonExtendedNegotiations[0].RelatedGeneralSOPClassUIDs[0] = "changed"
	negotiation := assoc.FindExtendedNegotiation(testExtendedSOPClassUID)
	if negotiation.RequestedApplicationInfo[0] != 1 ||
		negotiation.RelatedGeneralSOPClassUIDs[0] != testRelatedGeneralSOPClassUID1 {
		t.Fatalf("RQ retained association-owned slices: %#v", negotiation)
	}
}

func TestToAAssociateACIncludesOnlyExplicitlyAcceptedNegotiation(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   testStorageSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}
	rq.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         testStorageSOPClassUID,
		ServiceClassAppInfo: []byte{1, 1},
	}}
	rq.UserInformation.UserIdentity = &pdu.UserIdentityNegotiation{
		UserIdentityType:          5,
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("request-token"),
	}

	assoc := FromAAssociateRQ(rq, transfer.NewRegistry())
	assoc.PresentationContexts[0].Accept(assoc.PresentationContexts[0].ProposedTransferSyntaxes[0])
	withoutAcceptance := ToAAssociateACWithMaxPDULength(assoc, assoc.MaxPDULength)
	if got := withoutAcceptance.UserInformation.ExtendedNegotiations; len(got) != 0 {
		t.Fatalf("unaccepted extended negotiation was echoed: %#v", got)
	}
	if withoutAcceptance.UserInformation.UserIdentityResponse != nil {
		t.Fatal("unset user identity response was encoded")
	}

	extended := assoc.FindExtendedNegotiation(testStorageSOPClassUID)
	extended.AcceptApplicationInfo([]byte{1, 0, 1})
	assoc.UserIdentity.ServerResponse = []byte{}
	accepted := ToAAssociateACWithMaxPDULength(assoc, assoc.MaxPDULength)
	if got := accepted.UserInformation.ExtendedNegotiations; len(got) != 1 ||
		!bytes.Equal(got[0].ServiceClassAppInfo, []byte{1, 0, 1}) {
		t.Fatalf("accepted extended negotiation = %#v", got)
	}
	if accepted.UserInformation.UserIdentityResponse == nil ||
		accepted.UserInformation.UserIdentityResponse.ServerResponse == nil {
		t.Fatal("explicit empty user identity response was not encoded")
	}
}

func TestToAAssociateACBindsAcceptedRoleToPresentationContext(t *testing.T) {
	const storageSOPClassUID = "1.2.840.10008.5.1.4.1.1.2"
	assoc := NewAssociation("REQUESTOR", "ACCEPTOR")
	pc := NewPresentationContext(1, storageSOPClassUID, transfer.ExplicitVRLittleEndian)
	pc.Accept(transfer.ExplicitVRLittleEndian)
	pc.RequestedRole = NewRoleSelection(storageSOPClassUID, 1, 1)
	if err := assoc.AddPresentationContext(pc); err != nil {
		t.Fatalf("AddPresentationContext failed: %v", err)
	}
	assoc.RequestedRoleSelections = []*RoleSelection{NewRoleSelection(storageSOPClassUID, 1, 1)}
	assoc.RoleSelections = []*RoleSelection{NewRoleSelection(storageSOPClassUID, 1, 1)}

	ac := ToAAssociateACWithMaxPDULength(assoc, assoc.MaxPDULength)
	if len(ac.UserInformation.SCPSCURoleSelections) != 1 {
		t.Fatalf("role selections = %#v, want one accepted role", ac.UserInformation.SCPSCURoleSelections)
	}
	if pc.AcceptedRole == nil || pc.AcceptedRole.SCURole != 1 || pc.AcceptedRole.SCPRole != 1 {
		t.Fatalf("presentation context accepted role = %#v, want SCU/SCP", pc.AcceptedRole)
	}
}

func TestApplyAAssociateACRejectsRoleNotRequested(t *testing.T) {
	const storageSOPClassUID = "1.2.840.10008.5.1.4.1.1.2"
	rq := pdu.NewAAssociateRQ()
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   storageSOPClassUID,
		TransferSyntaxes: []string{transfer.ExplicitVRLittleEndian.UID().UID()},
	}}
	rq.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: storageSOPClassUID,
		SCURole:     1,
		SCPRole:     0,
	}}
	ac := pdu.NewAAssociateAC()
	ac.PresentationContexts = []pdu.PresentationContextAC{{
		ID:             1,
		Result:         pdu.ResultAcceptance,
		TransferSyntax: transfer.ExplicitVRLittleEndian.UID().UID(),
	}}
	ac.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: storageSOPClassUID,
		SCURole:     1,
		SCPRole:     1,
	}}

	registry := transfer.NewRegistry()
	err := ApplyAAssociateAC(FromAAssociateRQ(rq, registry), ac, registry)
	if err == nil || !strings.Contains(err.Error(), "unrequested SCP role") {
		t.Fatalf("ApplyAAssociateAC error = %v, want unrequested SCP role error", err)
	}
}

// newAsyncNegotiationPDUs returns a request/acceptance pair sharing one
// presentation context; the caller sets the Asynchronous Operations Window.
func newAsyncNegotiationPDUs() (*pdu.AAssociateRQ, *pdu.AAssociateAC) {
	rq := pdu.NewAAssociateRQ()
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   testStorageSOPClassUID,
		TransferSyntaxes: []string{testExplicitVRLittleEndianUID},
	}}
	ac := pdu.NewAAssociateAC()
	ac.PresentationContexts = []pdu.PresentationContextAC{{
		ID:             1,
		Result:         pdu.ResultAcceptance,
		TransferSyntax: testExplicitVRLittleEndianUID,
	}}
	return rq, ac
}

// TestApplyAAssociateACRejectsUnproposedAsyncWindow verifies PS3.7 D.3.3.3: the
// A-ASSOCIATE-AC may not carry an Asynchronous Operations Window when the
// request omitted it.
func TestApplyAAssociateACRejectsUnproposedAsyncWindow(t *testing.T) {
	rq, ac := newAsyncNegotiationPDUs()
	ac.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   5,
		MaximumNumberOperationsPerformed: 5,
	}

	registry := transfer.NewRegistry()
	err := ApplyAAssociateAC(FromAAssociateRQ(rq, registry), ac, registry)
	if err == nil || !strings.Contains(err.Error(), "not proposed") {
		t.Fatalf("ApplyAAssociateAC error = %v, want unproposed async window error", err)
	}
}

// TestApplyAAssociateACRejectsAsyncWindowAboveProposal verifies PS3.7 D.3.3.3:
// accepted values must be equal to or less than the proposed values, with zero
// meaning unlimited in both directions.
func TestApplyAAssociateACRejectsAsyncWindowAboveProposal(t *testing.T) {
	tests := []struct {
		name               string
		requestedInvoked   uint16
		requestedPerformed uint16
		acceptedInvoked    uint16
		acceptedPerformed  uint16
		wantErr            bool
	}{
		{"invoked above proposal", 4, 3, 5, 3, true},
		{"performed above proposal", 4, 3, 4, 4, true},
		{"unlimited acceptance above finite proposal", 4, 3, 0, 3, true},
		{"within proposal", 4, 3, 2, 1, false},
		{"equal to proposal", 4, 0, 4, 0, false},
		{"finite acceptance within unlimited proposal", 0, 0, 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq, ac := newAsyncNegotiationPDUs()
			rq.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
				MaximumNumberOperationsInvoked:   tt.requestedInvoked,
				MaximumNumberOperationsPerformed: tt.requestedPerformed,
			}
			ac.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
				MaximumNumberOperationsInvoked:   tt.acceptedInvoked,
				MaximumNumberOperationsPerformed: tt.acceptedPerformed,
			}

			registry := transfer.NewRegistry()
			assoc := FromAAssociateRQ(rq, registry)
			err := ApplyAAssociateAC(assoc, ac, registry)
			if tt.wantErr {
				if err == nil || !strings.Contains(err.Error(), "above the proposal") {
					t.Fatalf("ApplyAAssociateAC error = %v, want above-proposal error", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ApplyAAssociateAC error = %v", err)
			}
			if assoc.AsynchronousOperations == nil ||
				assoc.AsynchronousOperations.MaxInvokedOperations != tt.acceptedInvoked ||
				assoc.AsynchronousOperations.MaxPerformedOperations != tt.acceptedPerformed {
				t.Fatalf("accepted async window = %#v", assoc.AsynchronousOperations)
			}
		})
	}
}

// TestApplyAAssociateACDefaultsAsyncWindowWhenOmitted verifies the PS3.7
// D.3.3.3 default of one operation in each direction.
func TestApplyAAssociateACDefaultsAsyncWindowWhenOmitted(t *testing.T) {
	rq, ac := newAsyncNegotiationPDUs()
	rq.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   4,
		MaximumNumberOperationsPerformed: 3,
	}

	registry := transfer.NewRegistry()
	assoc := FromAAssociateRQ(rq, registry)
	if err := ApplyAAssociateAC(assoc, ac, registry); err != nil {
		t.Fatalf("ApplyAAssociateAC error = %v", err)
	}
	if assoc.AsynchronousOperations == nil ||
		assoc.AsynchronousOperations.MaxInvokedOperations != 1 ||
		assoc.AsynchronousOperations.MaxPerformedOperations != 1 {
		t.Fatalf("default async window = %#v, want 1,1", assoc.AsynchronousOperations)
	}
}
