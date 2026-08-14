// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package association

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

func TestApplyAAssociateACPreservesRequestedAndAcceptedNegotiation(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.CallingAETitle = "SCU"
	rq.CalledAETitle = "SCP"
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   "1.2.840.10008.5.1.4.1.1.2",
		TransferSyntaxes: []string{"1.2.840.10008.1.2.1"},
	}}
	rq.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   4,
		MaximumNumberOperationsPerformed: 3,
	}
	rq.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: "1.2.840.10008.5.1.4.1.1.2",
		SCURole:     1,
		SCPRole:     1,
	}}
	rq.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         "1.2.840.10008.5.1.4.1.1.2",
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
		TransferSyntax: "1.2.840.10008.1.2.1",
	}}
	ac.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   2,
		MaximumNumberOperationsPerformed: 1,
	}
	ac.UserInformation.SCPSCURoleSelections = []pdu.SCPSCURoleSelection{{
		SOPClassUID: "1.2.840.10008.5.1.4.1.1.2",
		SCURole:     1,
		SCPRole:     0,
	}}
	ac.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         "1.2.840.10008.5.1.4.1.1.2",
		ServiceClassAppInfo: []byte{1, 0, 1, 1},
	}}
	ac.UserInformation.UserIdentityResponse = &pdu.UserIdentityNegotiationResponse{
		ServerResponse: []byte("server-token"),
	}

	assoc := FromAAssociateRQ(rq)
	if err := ApplyAAssociateAC(assoc, ac); err != nil {
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
	if pc.AbstractSyntax != "1.2.840.10008.5.1.4.1.1.2" || len(pc.ProposedTransferSyntaxes) != 1 || pc.AcceptedTransferSyntax == nil {
		t.Fatalf("presentation context proposal/acceptance was not preserved: %#v", pc)
	}

	extended := assoc.FindExtendedNegotiation("1.2.840.10008.5.1.4.1.1.2")
	if extended == nil || !bytes.Equal(extended.RequestedApplicationInfo, []byte{1, 1, 0}) {
		t.Fatalf("requested extended negotiation = %#v", extended)
	}
	if !bytes.Equal(extended.AcceptedApplicationInfo, []byte{1, 0, 1}) {
		t.Fatalf("accepted extended negotiation = %v, want [1 0 1]", extended.AcceptedApplicationInfo)
	}

	if assoc.UserIdentity == nil || assoc.UserIdentity.Type != UserIdentityTypeJWT ||
		!bytes.Equal(assoc.UserIdentity.PrimaryField, []byte("request-token")) ||
		!bytes.Equal(assoc.UserIdentity.ServerResponse, []byte("server-token")) {
		t.Fatalf("merged user identity = %#v", assoc.UserIdentity)
	}
}

func TestToAAssociateACIncludesOnlyExplicitlyAcceptedNegotiation(t *testing.T) {
	rq := pdu.NewAAssociateRQ()
	rq.PresentationContexts = []pdu.PresentationContextRQ{{
		ID:               1,
		AbstractSyntax:   "1.2.840.10008.5.1.4.1.1.2",
		TransferSyntaxes: []string{"1.2.840.10008.1.2.1"},
	}}
	rq.UserInformation.ExtendedNegotiations = []pdu.ExtendedNegotiation{{
		SOPClassUID:         "1.2.840.10008.5.1.4.1.1.2",
		ServiceClassAppInfo: []byte{1, 1},
	}}
	rq.UserInformation.UserIdentity = &pdu.UserIdentityNegotiation{
		UserIdentityType:          5,
		PositiveResponseRequested: 1,
		PrimaryField:              []byte("request-token"),
	}

	assoc := FromAAssociateRQ(rq)
	assoc.PresentationContexts[0].Accept(assoc.PresentationContexts[0].ProposedTransferSyntaxes[0])
	withoutAcceptance := ToAAssociateAC(assoc)
	if got := withoutAcceptance.UserInformation.ExtendedNegotiations; len(got) != 0 {
		t.Fatalf("unaccepted extended negotiation was echoed: %#v", got)
	}
	if withoutAcceptance.UserInformation.UserIdentityResponse != nil {
		t.Fatal("unset user identity response was encoded")
	}

	extended := assoc.FindExtendedNegotiation("1.2.840.10008.5.1.4.1.1.2")
	extended.AcceptApplicationInfo([]byte{1, 0, 1})
	assoc.UserIdentity.ServerResponse = []byte{}
	accepted := ToAAssociateAC(assoc)
	if got := accepted.UserInformation.ExtendedNegotiations; len(got) != 1 ||
		!bytes.Equal(got[0].ServiceClassAppInfo, []byte{1, 0}) {
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

	ac := ToAAssociateAC(assoc)
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

	err := ApplyAAssociateAC(FromAAssociateRQ(rq), ac)
	if err == nil || !strings.Contains(err.Error(), "unrequested SCP role") {
		t.Fatalf("ApplyAAssociateAC error = %v, want unrequested SCP role error", err)
	}
}
