// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
	ImplementationClassUID   string // Implementation Class UID
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
		CallingAE:                callingAE,
		CalledAE:                 calledAE,
		ProtocolVersion:          0x0001,
		MaxPDULength:             16384, // 16KB default
		ImplementationClassUID:   "1.2.826.0.1.3680043.8.498.1", // go-dicom implementation UID
		ImplementationVersionName: "go-dicom-1.0",
		PresentationContexts:     make([]*PresentationContext, 0),
		ExtendedNegotiations:     make([]*ExtendedNegotiation, 0),
		RoleSelections:           make([]*RoleSelection, 0),
		IsEstablished:            false,
		messageIDGen:             dimse.NewMessageIDGenerator(), // One generator per association
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
func (a *Association) GetTransferSyntaxForAbstractSyntax(abstractSyntax string) *transfer.TransferSyntax {
	pc := a.FindPresentationContextByAbstractSyntax(abstractSyntax)
	if pc != nil && pc.AcceptedTransferSyntax != nil {
		return pc.AcceptedTransferSyntax
	}
	return nil
}

// SetEstablished marks the association as established.
func (a *Association) SetEstablished(established bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.IsEstablished = established
}

// String returns a human-readable representation of the association.
// NextMessageID returns the next available MessageID for this association.
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
	ProposedTransferSyntaxes []*transfer.TransferSyntax

	// AcceptedTransferSyntax is the accepted transfer syntax (for A-ASSOCIATE-AC)
	AcceptedTransferSyntax *transfer.TransferSyntax

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
func NewPresentationContext(id byte, abstractSyntax string, transferSyntaxes ...*transfer.TransferSyntax) *PresentationContext {
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
	transferSyntaxes := make([]*transfer.TransferSyntax, len(transferSyntaxUIDs))
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
func (pc *PresentationContext) Accept(transferSyntax *transfer.TransferSyntax) {
	pc.Result = ResultAcceptance
	pc.AcceptedTransferSyntax = transferSyntax
}

// Reject rejects this presentation context with the given reason.
func (pc *PresentationContext) Reject(result byte) {
	pc.Result = result
	pc.AcceptedTransferSyntax = nil
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
		var acceptedTS *transfer.TransferSyntax

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
