// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package association provides DICOM association negotiation and management.
// It handles the initial association request, response, and context negotiation.
package association

import (
	"fmt"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// Association represents a DICOM association between two Application Entities.
// This is a pure data structure that describes the negotiated parameters of an association.
// It does not contain any network I/O logic.
type Association struct {
	// Basic information
	CallingAE string // Calling Application Entity Title
	CalledAE  string // Called Application Entity Title

	// Remote endpoint (for informational purposes)
	RemoteHost string
	RemotePort int

	// Protocol negotiation
	ProtocolVersion uint16 // Usually 0x0001
	MaxPDULength    uint32 // Maximum PDU length (default 16384)

	// Implementation identification
	ImplementationClassUID    string // Implementation Class UID
	ImplementationVersionName string // Implementation Version Name (optional)

	// Presentation Contexts
	PresentationContexts []*PresentationContext

	// Extended Negotiation (optional)
	ExtendedNegotiations []*ExtendedNegotiation

	// SCP/SCU Role Selection (optional)
	RoleSelections []*RoleSelection

	// User Identity (optional)
	UserIdentity *UserIdentity

	// Asynchronous Operations Window (optional)
	AsynchronousOperations *AsynchronousOperationsWindow

	// Association state
	IsEstablished bool

	// MessageID generator for this association
	// Each association maintains its own MessageID counter that is
	// unique within the scope of this association/connection
	messageIDGen *dimse.MessageIDGenerator

	// Thread safety
	mu sync.RWMutex
}

// NewAssociation creates a new Association with default values.
func NewAssociation(callingAE, calledAE string) *Association {
	return &Association{
		CallingAE:                 callingAE,
		CalledAE:                  calledAE,
		ProtocolVersion:           0x0001,
		MaxPDULength:              16384,                         // 16KB default
		ImplementationClassUID:    "1.2.826.0.1.3680043.8.498.1", // go-dicom implementation UID
		ImplementationVersionName: "go-dicom-1.0",
		PresentationContexts:      make([]*PresentationContext, 0),
		ExtendedNegotiations:      make([]*ExtendedNegotiation, 0),
		RoleSelections:            make([]*RoleSelection, 0),
		IsEstablished:             false,
		messageIDGen:              dimse.NewMessageIDGenerator(), // One generator per association
	}
}

// AddPresentationContext adds a presentation context to the association.
// For A-ASSOCIATE-RQ, this adds a proposed context.
// For A-ASSOCIATE-AC, this adds an accepted/rejected context.
func (a *Association) AddPresentationContext(pc *PresentationContext) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Verify context ID is odd
	if pc.ID%2 == 0 {
		return fmt.Errorf("presentation context ID must be odd, got %d", pc.ID)
	}

	// Check for duplicate IDs
	for _, existing := range a.PresentationContexts {
		if existing.ID == pc.ID {
			return fmt.Errorf("presentation context ID %d already exists", pc.ID)
		}
	}

	a.PresentationContexts = append(a.PresentationContexts, pc)
	return nil
}

// FindPresentationContextByID finds a presentation context by its ID.
func (a *Association) FindPresentationContextByID(id byte) *PresentationContext {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, pc := range a.PresentationContexts {
		if pc.ID == id {
			return pc
		}
	}
	return nil
}

// FindPresentationContextByAbstractSyntax finds the first accepted presentation context
// for the given abstract syntax (SOP Class UID).
func (a *Association) FindPresentationContextByAbstractSyntax(abstractSyntax string) *PresentationContext {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, pc := range a.PresentationContexts {
		if pc.AbstractSyntax == abstractSyntax && pc.IsAccepted() {
			return pc
		}
	}
	return nil
}

// AddExtendedNegotiation adds an extended negotiation item.
func (a *Association) AddExtendedNegotiation(en *ExtendedNegotiation) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.ExtendedNegotiations = append(a.ExtendedNegotiations, en)
}

// FindExtendedNegotiation finds extended negotiation by SOP Class UID.
func (a *Association) FindExtendedNegotiation(sopClassUID string) *ExtendedNegotiation {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, en := range a.ExtendedNegotiations {
		if en.SOPClassUID == sopClassUID {
			return en
		}
	}
	return nil
}

// AddRoleSelection adds a role selection item.
func (a *Association) AddRoleSelection(rs *RoleSelection) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.RoleSelections = append(a.RoleSelections, rs)
}

// FindRoleSelection finds role selection by SOP Class UID.
func (a *Association) FindRoleSelection(sopClassUID string) *RoleSelection {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, rs := range a.RoleSelections {
		if rs.SOPClassUID == sopClassUID {
			return rs
		}
	}
	return nil
}

// GetAcceptedPresentationContexts returns all accepted presentation contexts.
func (a *Association) GetAcceptedPresentationContexts() []*PresentationContext {
	a.mu.RLock()
	defer a.mu.RUnlock()

	var accepted []*PresentationContext
	for _, pc := range a.PresentationContexts {
		if pc.IsAccepted() {
			accepted = append(accepted, pc)
		}
	}
	return accepted
}

// GetTransferSyntaxForAbstractSyntax returns the accepted transfer syntax for the given abstract syntax.
func (a *Association) GetTransferSyntaxForAbstractSyntax(abstractSyntax string) *transfer.Syntax {
	pc := a.FindPresentationContextByAbstractSyntax(abstractSyntax)
	if pc != nil && pc.AcceptedTransferSyntax != nil {
		return pc.AcceptedTransferSyntax
	}
	return nil
}

// SetEstablished marks the association as established.
func (a *Association) SetEstablished(isEstablished bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.IsEstablished = isEstablished
}

// NextMessageID returns the next available MessageID for this association.
// String returns a human-readable representation of the association.
// This is thread-safe and can be called concurrently.
func (a *Association) NextMessageID() uint16 {
	return a.messageIDGen.Next()
}

// AssignMessageID assigns a MessageID to a DIMSE message if it doesn't have one.
// If the message already has a MessageID, it will be preserved.
// This is the recommended way to assign MessageIDs to ensure uniqueness within the association.
func (a *Association) AssignMessageID(msg dimse.Message) (uint16, error) {
	return a.messageIDGen.AssignMessageID(msg)
}

func (a *Association) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()

	acceptedCount := 0
	for _, pc := range a.PresentationContexts {
		if pc.IsAccepted() {
			acceptedCount++
		}
	}

	return fmt.Sprintf("Association[%s -> %s, %d/%d contexts accepted, MaxPDU=%d]",
		a.CallingAE, a.CalledAE, acceptedCount, len(a.PresentationContexts), a.MaxPDULength)
}

// PresentationContext represents a presentation context in an association.
type PresentationContext struct {
	// ID is the presentation context ID (odd numbers: 1, 3, 5, ...)
	ID byte

	// AbstractSyntax is the SOP Class UID
	AbstractSyntax string

	// ProposedTransferSyntaxes is the list of proposed transfer syntaxes (for A-ASSOCIATE-RQ)
	ProposedTransferSyntaxes []*transfer.Syntax

	// AcceptedTransferSyntax is the accepted transfer syntax (for A-ASSOCIATE-AC)
	AcceptedTransferSyntax *transfer.Syntax

	// Result indicates the acceptance status (for A-ASSOCIATE-AC)
	// 0 = acceptance
	// 1 = user-rejection
	// 2 = no-reason (provider rejection)
	// 3 = abstract-syntax-not-supported (provider rejection)
	// 4 = transfer-syntaxes-not-supported (provider rejection)
	Result byte
}

// Acceptance result codes
const (
	ResultAcceptance                   byte = 0
	ResultUserRejection                byte = 1
	ResultNoReason                     byte = 2
	ResultAbstractSyntaxNotSupported   byte = 3
	ResultTransferSyntaxesNotSupported byte = 4
)

// NewPresentationContext creates a new presentation context for A-ASSOCIATE-RQ.
func NewPresentationContext(id byte, abstractSyntax string, transferSyntaxes ...*transfer.Syntax) *PresentationContext {
	return &PresentationContext{
		ID:                       id,
		AbstractSyntax:           abstractSyntax,
		ProposedTransferSyntaxes: transferSyntaxes,
		Result:                   ResultAcceptance, // Default to acceptance
	}
}

// NewPresentationContextFromUID creates a presentation context using UID objects.
func NewPresentationContextFromUID(id byte, abstractSyntaxUID *uid.UID, transferSyntaxUIDs ...*uid.UID) (*PresentationContext, error) {
	// Convert abstract syntax UID
	abstractSyntax := abstractSyntaxUID.UID()

	// Convert transfer syntax UIDs to TransferSyntax objects
	transferSyntaxes := make([]*transfer.Syntax, len(transferSyntaxUIDs))
	for i, tsUID := range transferSyntaxUIDs {
		ts, err := transfer.Lookup(tsUID)
		if err != nil {
			return nil, fmt.Errorf("unknown transfer syntax: %s: %w", tsUID.UID(), err)
		}
		transferSyntaxes[i] = ts
	}

	return NewPresentationContext(id, abstractSyntax, transferSyntaxes...), nil
}

// Accept accepts this presentation context with the given transfer syntax.
func (pc *PresentationContext) Accept(transferSyntax *transfer.Syntax) {
	pc.Result = ResultAcceptance
	pc.AcceptedTransferSyntax = transferSyntax
}

// Reject rejects this presentation context with the given reason.
func (pc *PresentationContext) Reject(result byte) {
	pc.Result = result
	pc.AcceptedTransferSyntax = nil
}

// SetResult sets the result and optionally the accepted transfer syntax.
// This is a convenience method similar to fo-dicom's SetResult.
//
// Example:
//
//	pc.SetResult(ResultAcceptance, transferSyntax)
//	pc.SetResult(ResultAbstractSyntaxNotSupported, nil)
func (pc *PresentationContext) SetResult(result byte, transferSyntax *transfer.Syntax) {
	pc.Result = result
	if result == ResultAcceptance {
		pc.AcceptedTransferSyntax = transferSyntax
	} else {
		pc.AcceptedTransferSyntax = nil
	}
}

// AcceptTransferSyntaxes compares accepted transfer syntaxes against the proposed transfer syntaxes.
// Returns true if a matching transfer syntax was found and accepted, false otherwise.
//
// This method is similar to fo-dicom's AcceptTransferSyntaxes method.
// If scpPriority is false, transfer syntaxes are accepted in the order proposed by the SCU.
// If scpPriority is true, transfer syntaxes are accepted in the order specified by acceptedTransferSyntaxes.
//
// Example:
//
//	// Accept in SCU-proposed order (prefer SCU's preference)
//	if pc.AcceptTransferSyntaxes(false, ts1, ts2, ts3) {
//	    // Accepted
//	}
//
//	// Accept in SCP-specified order (prefer SCP's preference)
//	if pc.AcceptTransferSyntaxes(true, ts1, ts2, ts3) {
//	    // Accepted
//	}
func (pc *PresentationContext) AcceptTransferSyntaxes(scpPriority bool, acceptedTransferSyntaxes ...*transfer.Syntax) bool {
	// If already accepted, return true
	if pc.Result == ResultAcceptance {
		return true
	}

	if scpPriority {
		// SCP decides priority - iterate through SCP's preferred order
		for _, ts := range acceptedTransferSyntaxes {
			if ts != nil && pc.HasTransferSyntax(ts) {
				pc.SetResult(ResultAcceptance, ts)
				return true
			}
		}
	} else {
		// SCU priority - iterate through proposed transfer syntaxes in order
		for _, proposedTS := range pc.ProposedTransferSyntaxes {
			for _, acceptedTS := range acceptedTransferSyntaxes {
				if proposedTS.UID().UID() == acceptedTS.UID().UID() {
					pc.SetResult(ResultAcceptance, acceptedTS)
					return true
				}
			}
		}
	}

	// No matching transfer syntax found
	pc.SetResult(ResultTransferSyntaxesNotSupported, nil)
	return false
}

// HasTransferSyntax checks if this presentation context has the specified transfer syntax
// in its proposed transfer syntaxes.
func (pc *PresentationContext) HasTransferSyntax(ts *transfer.Syntax) bool {
	for _, proposedTS := range pc.ProposedTransferSyntaxes {
		if proposedTS.UID().UID() == ts.UID().UID() {
			return true
		}
	}
	return false
}

// IsAccepted returns true if this presentation context is accepted.
func (pc *PresentationContext) IsAccepted() bool {
	return pc.Result == ResultAcceptance
}

// ExtendedNegotiation represents SOP Class Extended Negotiation.
type ExtendedNegotiation struct {
	// SOPClassUID is the SOP Class UID
	SOPClassUID string

	// ServiceClassAppInfo is the application information (SOP Class specific)
	ServiceClassAppInfo []byte
}

// NewExtendedNegotiation creates a new extended negotiation item.
func NewExtendedNegotiation(sopClassUID string, appInfo []byte) *ExtendedNegotiation {
	return &ExtendedNegotiation{
		SOPClassUID:         sopClassUID,
		ServiceClassAppInfo: appInfo,
	}
}

// RoleSelection represents SCP/SCU Role Selection negotiation.
type RoleSelection struct {
	// SOPClassUID is the SOP Class UID
	SOPClassUID string

	// SCURole indicates if the requesting AE can act as SCU (1=yes, 0=no)
	SCURole byte

	// SCPRole indicates if the requesting AE can act as SCP (1=yes, 0=no)
	SCPRole byte
}

// NewRoleSelection creates a new role selection item.
func NewRoleSelection(sopClassUID string, scuRole, scpRole byte) *RoleSelection {
	return &RoleSelection{
		SOPClassUID: sopClassUID,
		SCURole:     scuRole,
		SCPRole:     scpRole,
	}
}

// UserIdentity represents User Identity Negotiation.
type UserIdentity struct {
	// Type indicates the type of user identity:
	//   1 = Username
	//   2 = Username + Password
	//   3 = Kerberos Service Ticket
	//   4 = SAML Assertion
	//   5 = JSON Web Token (JWT)
	Type byte

	// PositiveResponseRequested indicates if a positive response is requested
	PositiveResponseRequested bool

	// PrimaryField contains:
	//   - Username (Type=1,2)
	//   - Kerberos ticket (Type=3)
	//   - SAML assertion (Type=4)
	//   - JWT (Type=5)
	PrimaryField []byte

	// SecondaryField contains the password (only for Type=2)
	SecondaryField []byte

	// ServerResponse contains the server's response (if PositiveResponseRequested=true)
	ServerResponse []byte
}

// User identity types
const (
	UserIdentityTypeUsername         byte = 1
	UserIdentityTypeUsernamePassword byte = 2
	UserIdentityTypeKerberos         byte = 3
	UserIdentityTypeSAML             byte = 4
	UserIdentityTypeJWT              byte = 5
)

// NewUserIdentityUsername creates a username-only user identity.
func NewUserIdentityUsername(username string, responseRequested bool) *UserIdentity {
	return &UserIdentity{
		Type:                      UserIdentityTypeUsername,
		PositiveResponseRequested: responseRequested,
		PrimaryField:              []byte(username),
	}
}

// NewUserIdentityUsernamePassword creates a username+password user identity.
func NewUserIdentityUsernamePassword(username, password string) *UserIdentity {
	return &UserIdentity{
		Type:           UserIdentityTypeUsernamePassword,
		PrimaryField:   []byte(username),
		SecondaryField: []byte(password),
	}
}

// AsynchronousOperationsWindow represents the Asynchronous Operations Window negotiation.
type AsynchronousOperationsWindow struct {
	// MaxInvokedOperations is the maximum number of outstanding operations invoked
	MaxInvokedOperations uint16

	// MaxPerformedOperations is the maximum number of outstanding operations performed
	MaxPerformedOperations uint16
}

// NewAsynchronousOperationsWindow creates a new asynchronous operations window.
func NewAsynchronousOperationsWindow(maxInvoked, maxPerformed uint16) *AsynchronousOperationsWindow {
	return &AsynchronousOperationsWindow{
		MaxInvokedOperations:   maxInvoked,
		MaxPerformedOperations: maxPerformed,
	}
}

// FromAAssociateRQ creates an Association from an A-ASSOCIATE-RQ PDU.
// This is typically used by an SCP (server) after receiving a connection request from an SCU (client).
// The returned Association will have all presentation contexts in "Proposed" state (Result=255).
// The server should then negotiate these contexts (accept/reject) before sending A-ASSOCIATE-AC.
func FromAAssociateRQ(rq *pdu.AAssociateRQ) *Association {
	assoc := NewAssociation(rq.CallingAETitle, rq.CalledAETitle)
	assoc.ProtocolVersion = rq.ProtocolVersion

	// Set user information
	if rq.UserInformation != nil {
		assoc.MaxPDULength = rq.UserInformation.MaximumLength
		assoc.ImplementationClassUID = rq.UserInformation.ImplementationClassUID
		assoc.ImplementationVersionName = rq.UserInformation.ImplementationVersionName

		// Asynchronous operations
		if rq.UserInformation.AsynchronousOperations != nil {
			assoc.AsynchronousOperations = &AsynchronousOperationsWindow{
				MaxInvokedOperations:   rq.UserInformation.AsynchronousOperations.MaximumNumberOperationsInvoked,
				MaxPerformedOperations: rq.UserInformation.AsynchronousOperations.MaximumNumberOperationsPerformed,
			}
		}

		// Role selections
		for _, rs := range rq.UserInformation.SCPSCURoleSelections {
			assoc.AddRoleSelection(&RoleSelection{
				SOPClassUID: rs.SOPClassUID,
				SCURole:     rs.SCURole,
				SCPRole:     rs.SCPRole,
			})
		}

		// Extended negotiations
		for _, en := range rq.UserInformation.ExtendedNegotiations {
			assoc.AddExtendedNegotiation(&ExtendedNegotiation{
				SOPClassUID:         en.SOPClassUID,
				ServiceClassAppInfo: en.ServiceClassAppInfo,
			})
		}

		// User identity
		if rq.UserInformation.UserIdentity != nil {
			assoc.UserIdentity = &UserIdentity{
				Type:                      rq.UserInformation.UserIdentity.UserIdentityType,
				PositiveResponseRequested: rq.UserInformation.UserIdentity.PositiveResponseRequested != 0,
				PrimaryField:              rq.UserInformation.UserIdentity.PrimaryField,
				SecondaryField:            rq.UserInformation.UserIdentity.SecondaryField,
			}
		}
	}

	// Convert presentation contexts (from RQ)
	for _, pcRQ := range rq.PresentationContexts {
		// Parse transfer syntaxes
		transferSyntaxes := make([]*transfer.Syntax, 0, len(pcRQ.TransferSyntaxes))
		for _, tsUID := range pcRQ.TransferSyntaxes {
			ts, err := transfer.Parse(tsUID)
			if err != nil {
				// Skip invalid transfer syntaxes
				continue
			}
			transferSyntaxes = append(transferSyntaxes, ts)
		}

		pc := &PresentationContext{
			ID:                       pcRQ.ID,
			AbstractSyntax:           pcRQ.AbstractSyntax,
			ProposedTransferSyntaxes: transferSyntaxes,
			Result:                   255, // Proposed (not yet negotiated)
		}

		_ = assoc.AddPresentationContext(pc)
	}

	return assoc
}

// FromAAssociateAC creates an Association from an A-ASSOCIATE-AC PDU.
// This is typically used by an SCU (client) after receiving acceptance from an SCP (server).
func FromAAssociateAC(ac *pdu.AAssociateAC) *Association {
	assoc := NewAssociation(ac.CallingAETitle, ac.CalledAETitle)
	assoc.ProtocolVersion = ac.ProtocolVersion

	// Set user information
	if ac.UserInformation != nil {
		assoc.MaxPDULength = ac.UserInformation.MaximumLength
		assoc.ImplementationClassUID = ac.UserInformation.ImplementationClassUID
		assoc.ImplementationVersionName = ac.UserInformation.ImplementationVersionName

		// Asynchronous operations
		if ac.UserInformation.AsynchronousOperations != nil {
			assoc.AsynchronousOperations = &AsynchronousOperationsWindow{
				MaxInvokedOperations:   ac.UserInformation.AsynchronousOperations.MaximumNumberOperationsInvoked,
				MaxPerformedOperations: ac.UserInformation.AsynchronousOperations.MaximumNumberOperationsPerformed,
			}
		}

		// Role selections
		for _, rs := range ac.UserInformation.SCPSCURoleSelections {
			assoc.AddRoleSelection(&RoleSelection{
				SOPClassUID: rs.SOPClassUID,
				SCURole:     rs.SCURole,
				SCPRole:     rs.SCPRole,
			})
		}

		// Extended negotiations
		for _, en := range ac.UserInformation.ExtendedNegotiations {
			assoc.AddExtendedNegotiation(&ExtendedNegotiation{
				SOPClassUID:         en.SOPClassUID,
				ServiceClassAppInfo: en.ServiceClassAppInfo,
			})
		}
	}

	// Convert presentation contexts
	for _, pcAC := range ac.PresentationContexts {
		var acceptedTS *transfer.Syntax

		// Only set transfer syntax if accepted
		if pcAC.Result == pdu.ResultAcceptance && pcAC.TransferSyntax != "" {
			ts, err := transfer.Parse(pcAC.TransferSyntax)
			if err != nil {
				// If we can't parse it, treat as rejected
				pcAC.Result = pdu.ResultTransferSyntaxesNotSupported
			} else {
				acceptedTS = ts
			}
		}

		pc := &PresentationContext{
			ID:                     pcAC.ID,
			AcceptedTransferSyntax: acceptedTS,
			Result:                 pcAC.Result,
		}

		// Note: AbstractSyntax is not in A-ASSOCIATE-AC
		// The SCU must remember the original abstract syntax from the RQ
		// For now, we leave it empty as the client will need to map it

		_ = assoc.AddPresentationContext(pc)
	}

	assoc.IsEstablished = true
	return assoc
}

// ToAAssociateAC converts an Association to an A-ASSOCIATE-AC PDU.
// This is typically used by an SCP (server) to send the association acceptance response.
// The Association should have been negotiated (presentation contexts accepted/rejected).
func ToAAssociateAC(assoc *Association) *pdu.AAssociateAC {
	ac := pdu.NewAAssociateAC()
	ac.CallingAETitle = assoc.CallingAE
	ac.CalledAETitle = assoc.CalledAE
	ac.ProtocolVersion = assoc.ProtocolVersion

	// Set user information
	ac.UserInformation = &pdu.UserInformation{
		MaximumLength:             assoc.MaxPDULength,
		ImplementationClassUID:    assoc.ImplementationClassUID,
		ImplementationVersionName: assoc.ImplementationVersionName,
	}

	// Asynchronous operations
	if assoc.AsynchronousOperations != nil {
		ac.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
			MaximumNumberOperationsInvoked:   assoc.AsynchronousOperations.MaxInvokedOperations,
			MaximumNumberOperationsPerformed: assoc.AsynchronousOperations.MaxPerformedOperations,
		}
	}

	// Role selections
	for _, rs := range assoc.RoleSelections {
		ac.UserInformation.SCPSCURoleSelections = append(ac.UserInformation.SCPSCURoleSelections, pdu.SCPSCURoleSelection{
			SOPClassUID: rs.SOPClassUID,
			SCURole:     rs.SCURole,
			SCPRole:     rs.SCPRole,
		})
	}

	// Extended negotiations
	for _, en := range assoc.ExtendedNegotiations {
		ac.UserInformation.ExtendedNegotiations = append(ac.UserInformation.ExtendedNegotiations, pdu.ExtendedNegotiation{
			SOPClassUID:         en.SOPClassUID,
			ServiceClassAppInfo: en.ServiceClassAppInfo,
		})
	}

	// User identity response (if applicable)
	// Note: Server response is typically handled separately in user identity negotiation
	// and is not included in A-ASSOCIATE-AC in the current implementation

	// Convert presentation contexts
	for _, pc := range assoc.PresentationContexts {
		var transferSyntaxUID string
		if pc.AcceptedTransferSyntax != nil {
			transferSyntaxUID = pc.AcceptedTransferSyntax.UID().UID()
		}

		ac.PresentationContexts = append(ac.PresentationContexts, pdu.PresentationContextAC{
			ID:             pc.ID,
			Result:         pc.Result,
			TransferSyntax: transferSyntaxUID,
		})
	}

	return ac
}

// ToAAssociateRQ converts an Association back to an A-ASSOCIATE-RQ PDU.
// This is useful for backward compatibility with deprecated APIs that expect the PDU format.
// Note: This creates a new RQ from the Association's current state, which may differ from
// the original request if the Association was modified after creation.
func ToAAssociateRQ(assoc *Association) *pdu.AAssociateRQ {
	rq := pdu.NewAAssociateRQ()
	rq.CallingAETitle = assoc.CallingAE
	rq.CalledAETitle = assoc.CalledAE
	rq.ProtocolVersion = assoc.ProtocolVersion
	rq.ApplicationContext = "1.2.840.10008.3.1.1.1" // DICOM Application Context

	// Convert presentation contexts
	for _, pc := range assoc.PresentationContexts {
		transferSyntaxes := make([]string, len(pc.ProposedTransferSyntaxes))
		for i, ts := range pc.ProposedTransferSyntaxes {
			if ts != nil {
				transferSyntaxes[i] = ts.UID().UID()
			}
		}
		rq.PresentationContexts = append(rq.PresentationContexts, pdu.PresentationContextRQ{
			ID:               pc.ID,
			AbstractSyntax:   pc.AbstractSyntax,
			TransferSyntaxes: transferSyntaxes,
		})
	}

	// Set user information
	rq.UserInformation = &pdu.UserInformation{
		MaximumLength:             assoc.MaxPDULength,
		ImplementationClassUID:    assoc.ImplementationClassUID,
		ImplementationVersionName: assoc.ImplementationVersionName,
	}

	// Asynchronous operations
	if assoc.AsynchronousOperations != nil {
		rq.UserInformation.AsynchronousOperations = &pdu.AsynchronousOperationsWindow{
			MaximumNumberOperationsInvoked:   assoc.AsynchronousOperations.MaxInvokedOperations,
			MaximumNumberOperationsPerformed: assoc.AsynchronousOperations.MaxPerformedOperations,
		}
	}

	// Role selections
	for _, rs := range assoc.RoleSelections {
		rq.UserInformation.SCPSCURoleSelections = append(rq.UserInformation.SCPSCURoleSelections, pdu.SCPSCURoleSelection{
			SOPClassUID: rs.SOPClassUID,
			SCURole:     rs.SCURole,
			SCPRole:     rs.SCPRole,
		})
	}

	// Extended negotiations
	for _, en := range assoc.ExtendedNegotiations {
		rq.UserInformation.ExtendedNegotiations = append(rq.UserInformation.ExtendedNegotiations, pdu.ExtendedNegotiation{
			SOPClassUID:         en.SOPClassUID,
			ServiceClassAppInfo: en.ServiceClassAppInfo,
		})
	}

	// User identity
	if assoc.UserIdentity != nil {
		rq.UserInformation.UserIdentity = &pdu.UserIdentityNegotiation{
			UserIdentityType: assoc.UserIdentity.Type,
			PositiveResponseRequested: func() byte {
				if assoc.UserIdentity.PositiveResponseRequested {
					return 1
				}
				return 0
			}(),
			PrimaryField:   assoc.UserIdentity.PrimaryField,
			SecondaryField: assoc.UserIdentity.SecondaryField,
		}
	}

	return rq
}
