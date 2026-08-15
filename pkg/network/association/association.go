// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package association provides DICOM association negotiation and management.
// It handles the initial association request, response, and context negotiation.
package association

import (
	"fmt"
	"sort"
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
	// RequestedRoleSelections retains the roles proposed by the requestor.
	RequestedRoleSelections []*RoleSelection

	// User Identity (optional)
	UserIdentity *UserIdentity

	// Asynchronous Operations Window (optional)
	AsynchronousOperations *AsynchronousOperationsWindow
	// RequestedAsynchronousOperations retains the requestor's proposed values.
	RequestedAsynchronousOperations *AsynchronousOperationsWindow

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
		RequestedRoleSelections:   make([]*RoleSelection, 0),
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
	if en == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, existing := range a.ExtendedNegotiations {
		if existing != nil && existing.SOPClassUID == en.SOPClassUID {
			mergeExtendedNegotiation(existing, en)
			return
		}
	}
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

	// RequestedRole is the role selection proposed for this abstract syntax.
	RequestedRole *RoleSelection

	// AcceptedRole is the role selection returned by the association acceptor.
	AcceptedRole *RoleSelection
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
	// Deprecated: use RequestedApplicationInfo and AcceptedApplicationInfo.
	ServiceClassAppInfo []byte

	// RequestedApplicationInfo is the application information proposed by the requestor.
	RequestedApplicationInfo []byte

	// AcceptedApplicationInfo is the application information returned by the acceptor.
	AcceptedApplicationInfo []byte

	// ServiceClassUID is the optional Common Service Class UID requested through
	// SOP Class Common Extended Negotiation.
	ServiceClassUID string

	// RelatedGeneralSOPClassUIDs contains the ordered Related General SOP Class UIDs.
	RelatedGeneralSOPClassUIDs []string
}

// NewExtendedNegotiation creates a new extended negotiation item.
func NewExtendedNegotiation(sopClassUID string, appInfo []byte) *ExtendedNegotiation {
	requested := append([]byte(nil), appInfo...)
	return &ExtendedNegotiation{
		SOPClassUID:              sopClassUID,
		ServiceClassAppInfo:      requested,
		RequestedApplicationInfo: requested,
	}
}

// NewCommonExtendedNegotiation creates a SOP Class Common Extended Negotiation item.
func NewCommonExtendedNegotiation(
	sopClassUID, serviceClassUID string,
	relatedGeneralSOPClassUIDs ...string,
) *ExtendedNegotiation {
	return &ExtendedNegotiation{
		SOPClassUID:                sopClassUID,
		ServiceClassUID:            serviceClassUID,
		RelatedGeneralSOPClassUIDs: append([]string{}, relatedGeneralSOPClassUIDs...),
	}
}

// HasCommonExtendedNegotiation reports whether the common negotiation values
// were explicitly configured, including an invalid empty Service Class UID.
func (e *ExtendedNegotiation) HasCommonExtendedNegotiation() bool {
	return e != nil && (e.ServiceClassUID != "" || e.RelatedGeneralSOPClassUIDs != nil)
}

// Clone returns an ownership-independent copy of the negotiation.
func (e *ExtendedNegotiation) Clone() *ExtendedNegotiation {
	if e == nil {
		return nil
	}
	requested := append([]byte(nil), e.RequestedApplicationInfo...)
	legacy := append([]byte(nil), e.ServiceClassAppInfo...)
	if e.RequestedApplicationInfo != nil {
		legacy = requested
	}
	var relatedGeneralSOPClassUIDs []string
	if e.RelatedGeneralSOPClassUIDs != nil {
		relatedGeneralSOPClassUIDs = append([]string{}, e.RelatedGeneralSOPClassUIDs...)
	}
	return &ExtendedNegotiation{
		SOPClassUID:                e.SOPClassUID,
		ServiceClassAppInfo:        legacy,
		RequestedApplicationInfo:   requested,
		AcceptedApplicationInfo:    append([]byte(nil), e.AcceptedApplicationInfo...),
		ServiceClassUID:            e.ServiceClassUID,
		RelatedGeneralSOPClassUIDs: relatedGeneralSOPClassUIDs,
	}
}

func mergeExtendedNegotiation(existing, incoming *ExtendedNegotiation) {
	if incoming.RequestedApplicationInfo != nil || incoming.ServiceClassAppInfo != nil {
		requested := incoming.RequestedApplicationInfo
		if requested == nil {
			requested = incoming.ServiceClassAppInfo
		}
		existing.RequestedApplicationInfo = append([]byte(nil), requested...)
		existing.ServiceClassAppInfo = existing.RequestedApplicationInfo
	}
	if incoming.AcceptedApplicationInfo != nil {
		existing.AcceptedApplicationInfo = append([]byte(nil), incoming.AcceptedApplicationInfo...)
	}
	if incoming.HasCommonExtendedNegotiation() {
		existing.ServiceClassUID = incoming.ServiceClassUID
		existing.RelatedGeneralSOPClassUIDs = append([]string{}, incoming.RelatedGeneralSOPClassUIDs...)
	}
}

// AcceptApplicationInfo records the accepted application information while
// discarding fields that were not present in the request.
func (e *ExtendedNegotiation) AcceptApplicationInfo(appInfo []byte) {
	if e == nil || len(e.RequestedApplicationInfo) == 0 {
		return
	}
	length := min(len(appInfo), len(e.RequestedApplicationInfo))
	e.AcceptedApplicationInfo = append([]byte(nil), appInfo[:length]...)
}

// ServiceApplicationInfo encapsulates the Service Class Application Information field
// for the SOP Class Extended Negotiation Sub-item.
// See: http://dicom.nema.org/medical/dicom/current/output/chtml/part07/sect_D.3.3.5.html
type ServiceApplicationInfo struct {
	fields map[byte]byte
}

// NewServiceApplicationInfo creates a new ServiceApplicationInfo from raw bytes.
func NewServiceApplicationInfo(data []byte) *ServiceApplicationInfo {
	info := &ServiceApplicationInfo{
		fields: make(map[byte]byte),
	}
	for i, b := range data {
		info.fields[byte(i+1)] = b
	}
	return info
}

// NewServiceApplicationInfoEmpty creates an empty ServiceApplicationInfo.
func NewServiceApplicationInfoEmpty() *ServiceApplicationInfo {
	return &ServiceApplicationInfo{
		fields: make(map[byte]byte),
	}
}

// Count returns the number of fields.
func (s *ServiceApplicationInfo) Count() int {
	return len(s.fields)
}

// Get returns the value at the given index (1-based).
func (s *ServiceApplicationInfo) Get(index byte) (byte, bool) {
	v, ok := s.fields[index]
	return v, ok
}

// Set sets the value at the given index.
func (s *ServiceApplicationInfo) Set(index byte, value byte) error {
	if index == 0 {
		return fmt.Errorf("index 0 is not valid")
	}
	s.fields[index] = value
	s.fillGaps()
	return nil
}

// SetBool sets a boolean value at the given index.
func (s *ServiceApplicationInfo) SetBool(index byte, value bool) error {
	if value {
		return s.Set(index, 1)
	}
	return s.Set(index, 0)
}

// Contains returns true if the index exists.
func (s *ServiceApplicationInfo) Contains(index byte) bool {
	_, ok := s.fields[index]
	return ok
}

// GetBool returns the value as a boolean, or the default if not present.
func (s *ServiceApplicationInfo) GetBool(index byte, defaultValue bool) bool {
	if v, ok := s.fields[index]; ok {
		return v == 1
	}
	return defaultValue
}

// Remove removes the field at the given index.
func (s *ServiceApplicationInfo) Remove(index byte) bool {
	if _, ok := s.fields[index]; ok {
		delete(s.fields, index)
		return true
	}
	return false
}

// Values returns the raw bytes representation.
func (s *ServiceApplicationInfo) Values() []byte {
	if len(s.fields) == 0 {
		return nil
	}

	maxKey := byte(0)
	for k := range s.fields {
		if k > maxKey {
			maxKey = k
		}
	}

	result := make([]byte, maxKey)
	for k, v := range s.fields {
		result[k-1] = v
	}
	return result
}

// String returns a string representation.
func (s *ServiceApplicationInfo) String() string {
	if len(s.fields) == 0 {
		return ""
	}

	keys := make([]byte, 0, len(s.fields))
	for k := range s.fields {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	result := ""
	for i, k := range keys {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", s.fields[k])
	}
	return result
}

// fillGaps ensures sequential fields from 1 to max index.
func (s *ServiceApplicationInfo) fillGaps() {
	if len(s.fields) == 0 {
		return
	}

	maxKey := byte(0)
	for k := range s.fields {
		if k > maxKey {
			maxKey = k
		}
	}

	for i := byte(1); i < maxKey; i++ {
		if _, ok := s.fields[i]; !ok {
			s.fields[i] = 0
		}
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
	return NewUserIdentity(UserIdentityTypeUsername, []byte(username), nil, responseRequested)
}

// NewUserIdentityUsernamePassword creates a username+password user identity.
func NewUserIdentityUsernamePassword(username, password string) *UserIdentity {
	return NewUserIdentityUsernamePasswordWithResponse(username, password, false)
}

// NewUserIdentityUsernamePasswordWithResponse creates a username and password identity.
func NewUserIdentityUsernamePasswordWithResponse(username, password string, responseRequested bool) *UserIdentity {
	return NewUserIdentity(UserIdentityTypeUsernamePassword, []byte(username), []byte(password), responseRequested)
}

// NewUserIdentityKerberos creates a Kerberos service ticket identity.
func NewUserIdentityKerberos(ticket []byte, responseRequested bool) *UserIdentity {
	return NewUserIdentity(UserIdentityTypeKerberos, ticket, nil, responseRequested)
}

// NewUserIdentitySAML creates a SAML assertion identity.
func NewUserIdentitySAML(assertion []byte, responseRequested bool) *UserIdentity {
	return NewUserIdentity(UserIdentityTypeSAML, assertion, nil, responseRequested)
}

// NewUserIdentityJWT creates a JSON Web Token identity.
func NewUserIdentityJWT(token []byte, responseRequested bool) *UserIdentity {
	return NewUserIdentity(UserIdentityTypeJWT, token, nil, responseRequested)
}

// NewUserIdentity creates a user identity from raw fields.
func NewUserIdentity(identityType byte, primary, secondary []byte, responseRequested bool) *UserIdentity {
	return &UserIdentity{
		Type:                      identityType,
		PositiveResponseRequested: responseRequested,
		PrimaryField:              append([]byte(nil), primary...),
		SecondaryField:            append([]byte(nil), secondary...),
	}
}

// Validate checks the user identity type and field combination.
func (u *UserIdentity) Validate() error {
	if u == nil {
		return fmt.Errorf("user identity cannot be nil")
	}
	if u.Type < UserIdentityTypeUsername || u.Type > UserIdentityTypeJWT {
		return fmt.Errorf("unsupported user identity type %d", u.Type)
	}
	if u.Type != UserIdentityTypeUsernamePassword && len(u.SecondaryField) != 0 {
		return fmt.Errorf("secondary user identity field is only valid for username and password")
	}
	return nil

}

// Clone returns a defensive copy of the identity fields and server response.
func (u *UserIdentity) Clone() *UserIdentity {
	if u == nil {
		return nil
	}
	return &UserIdentity{
		Type:                      u.Type,
		PositiveResponseRequested: u.PositiveResponseRequested,
		PrimaryField:              append([]byte(nil), u.PrimaryField...),
		SecondaryField:            append([]byte(nil), u.SecondaryField...),
		ServerResponse:            append([]byte(nil), u.ServerResponse...),
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
			requested := &AsynchronousOperationsWindow{
				MaxInvokedOperations:   rq.UserInformation.AsynchronousOperations.MaximumNumberOperationsInvoked,
				MaxPerformedOperations: rq.UserInformation.AsynchronousOperations.MaximumNumberOperationsPerformed,
			}
			assoc.RequestedAsynchronousOperations = requested
			assoc.AsynchronousOperations = cloneAsynchronousOperationsWindow(requested)
		}

		// Role selections
		for _, rs := range rq.UserInformation.SCPSCURoleSelections {
			requested := &RoleSelection{
				SOPClassUID: rs.SOPClassUID,
				SCURole:     rs.SCURole,
				SCPRole:     rs.SCPRole,
			}
			assoc.RequestedRoleSelections = append(assoc.RequestedRoleSelections, requested)
			assoc.AddRoleSelection(cloneRoleSelection(requested))
		}

		// Extended negotiations
		for _, en := range rq.UserInformation.ExtendedNegotiations {
			assoc.AddExtendedNegotiation(NewExtendedNegotiation(en.SOPClassUID, en.ServiceClassAppInfo))
		}
		for _, en := range rq.UserInformation.CommonExtendedNegotiations {
			assoc.AddExtendedNegotiation(NewCommonExtendedNegotiation(
				en.SOPClassUID,
				en.ServiceClassUID,
				en.RelatedGeneralSOPClassUIDs...,
			))
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
	for _, role := range assoc.RequestedRoleSelections {
		for _, pc := range assoc.PresentationContexts {
			if pc.AbstractSyntax == role.SOPClassUID {
				pc.RequestedRole = cloneRoleSelection(role)
			}
		}
	}

	return assoc
}

// ApplyAAssociateAC merges an association acceptance into the original request
// association so callers can inspect both proposed and accepted values.
func ApplyAAssociateAC(assoc *Association, ac *pdu.AAssociateAC) error {
	if assoc == nil || ac == nil {
		return fmt.Errorf("association request and acceptance cannot be nil")
	}
	if ac.UserInformation == nil {
		return fmt.Errorf("A-ASSOCIATE-AC User Information is missing")
	}

	assoc.ProtocolVersion = ac.ProtocolVersion
	assoc.MaxPDULength = ac.UserInformation.MaximumLength
	assoc.ImplementationClassUID = ac.UserInformation.ImplementationClassUID
	assoc.ImplementationVersionName = ac.UserInformation.ImplementationVersionName

	seenContexts := make(map[byte]struct{}, len(ac.PresentationContexts))
	for i := range ac.PresentationContexts {
		accepted := &ac.PresentationContexts[i]
		if _, duplicate := seenContexts[accepted.ID]; duplicate {
			return fmt.Errorf("A-ASSOCIATE-AC contains duplicate presentation context ID %d", accepted.ID)
		}
		seenContexts[accepted.ID] = struct{}{}
		pc := assoc.FindPresentationContextByID(accepted.ID)
		if pc == nil {
			return fmt.Errorf("A-ASSOCIATE-AC contains unknown presentation context ID %d", accepted.ID)
		}
		pc.Result = accepted.Result
		pc.AcceptedTransferSyntax = nil
		if accepted.Result == pdu.ResultAcceptance {
			if accepted.TransferSyntax == "" {
				return fmt.Errorf("accepted presentation context %d has no transfer syntax", accepted.ID)
			}
			syntax, err := transfer.Parse(accepted.TransferSyntax)
			if err != nil {
				return fmt.Errorf("accepted presentation context %d transfer syntax: %w", accepted.ID, err)
			}
			if !pc.HasTransferSyntax(syntax) {
				return fmt.Errorf("presentation context %d accepted an unproposed transfer syntax", accepted.ID)
			}
			pc.AcceptedTransferSyntax = syntax
		}
	}
	if len(seenContexts) != len(assoc.PresentationContexts) {
		return fmt.Errorf("A-ASSOCIATE-AC returned %d presentation contexts, want %d", len(seenContexts), len(assoc.PresentationContexts))
	}

	assoc.AsynchronousOperations = &AsynchronousOperationsWindow{MaxInvokedOperations: 1, MaxPerformedOperations: 1}
	if async := ac.UserInformation.AsynchronousOperations; async != nil {
		assoc.AsynchronousOperations = &AsynchronousOperationsWindow{
			MaxInvokedOperations:   async.MaximumNumberOperationsInvoked,
			MaxPerformedOperations: async.MaximumNumberOperationsPerformed,
		}
	}

	assoc.RoleSelections = assoc.RoleSelections[:0]
	for _, pc := range assoc.PresentationContexts {
		pc.AcceptedRole = nil
	}
	for _, accepted := range ac.UserInformation.SCPSCURoleSelections {
		if accepted.SCURole > 1 || accepted.SCPRole > 1 {
			return fmt.Errorf("A-ASSOCIATE-AC returned invalid role values for SOP Class %s", accepted.SOPClassUID)
		}
		requested := findRoleSelection(assoc.RequestedRoleSelections, accepted.SOPClassUID)
		if requested == nil {
			return fmt.Errorf("A-ASSOCIATE-AC returned role selection for unrequested SOP Class %s", accepted.SOPClassUID)
		}
		if accepted.SCURole > requested.SCURole {
			return fmt.Errorf("A-ASSOCIATE-AC accepted an unrequested SCU role for SOP Class %s", accepted.SOPClassUID)
		}
		if accepted.SCPRole > requested.SCPRole {
			return fmt.Errorf("A-ASSOCIATE-AC accepted an unrequested SCP role for SOP Class %s", accepted.SOPClassUID)
		}
		role := &RoleSelection{SOPClassUID: accepted.SOPClassUID, SCURole: accepted.SCURole, SCPRole: accepted.SCPRole}
		assoc.RoleSelections = append(assoc.RoleSelections, role)
		for _, pc := range assoc.PresentationContexts {
			if pc.AbstractSyntax == accepted.SOPClassUID {
				pc.AcceptedRole = cloneRoleSelection(role)
			}
		}
	}

	for _, accepted := range ac.UserInformation.ExtendedNegotiations {
		if requested := assoc.FindExtendedNegotiation(accepted.SOPClassUID); requested != nil {
			requested.AcceptApplicationInfo(accepted.ServiceClassAppInfo)
		}
	}
	if assoc.UserIdentity != nil {
		response := ac.UserInformation.UserIdentityResponse
		if response == nil && ac.UserInformation.UserIdentity != nil && ac.UserInformation.UserIdentity.PositiveResponseRequested == 1 {
			response = &pdu.UserIdentityNegotiationResponse{ServerResponse: ac.UserInformation.UserIdentity.PrimaryField}
		}
		if response != nil {
			assoc.UserIdentity.ServerResponse = append([]byte{}, response.ServerResponse...)
		}
	}

	assoc.IsEstablished = true
	return nil
}

func cloneAsynchronousOperationsWindow(value *AsynchronousOperationsWindow) *AsynchronousOperationsWindow {
	if value == nil {
		return nil
	}
	return &AsynchronousOperationsWindow{
		MaxInvokedOperations:   value.MaxInvokedOperations,
		MaxPerformedOperations: value.MaxPerformedOperations,
	}
}

func cloneRoleSelection(value *RoleSelection) *RoleSelection {
	if value == nil {
		return nil
	}
	return &RoleSelection{SOPClassUID: value.SOPClassUID, SCURole: value.SCURole, SCPRole: value.SCPRole}
}

func findRoleSelection(values []*RoleSelection, sopClassUID string) *RoleSelection {
	for _, value := range values {
		if value != nil && value.SOPClassUID == sopClassUID {
			return value
		}
	}
	return nil
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

	// Convert presentation contexts — use index so mutations on the PDU
	// slice are visible to any later code that re-reads the slice.
	for i := range ac.PresentationContexts {
		pcAC := &ac.PresentationContexts[i]
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
	for _, pc := range assoc.PresentationContexts {
		pc.AcceptedRole = nil
	}
	for _, rs := range assoc.RoleSelections {
		if rs == nil {
			continue
		}
		requested := findRoleSelection(assoc.RequestedRoleSelections, rs.SOPClassUID)
		if requested == nil {
			continue
		}
		accepted := &RoleSelection{
			SOPClassUID: rs.SOPClassUID,
			SCURole:     rs.SCURole & requested.SCURole,
			SCPRole:     rs.SCPRole & requested.SCPRole,
		}
		for _, pc := range assoc.PresentationContexts {
			if pc.AbstractSyntax == accepted.SOPClassUID {
				pc.AcceptedRole = cloneRoleSelection(accepted)
			}
		}
		ac.UserInformation.SCPSCURoleSelections = append(ac.UserInformation.SCPSCURoleSelections, pdu.SCPSCURoleSelection{
			SOPClassUID: accepted.SOPClassUID,
			SCURole:     accepted.SCURole,
			SCPRole:     accepted.SCPRole,
		})
	}

	// Extended negotiations
	for _, en := range assoc.ExtendedNegotiations {
		if en == nil || en.AcceptedApplicationInfo == nil {
			continue
		}
		ac.UserInformation.ExtendedNegotiations = append(ac.UserInformation.ExtendedNegotiations, pdu.ExtendedNegotiation{
			SOPClassUID:         en.SOPClassUID,
			ServiceClassAppInfo: append([]byte(nil), en.AcceptedApplicationInfo...),
		})
	}

	// User identity response (if explicitly supplied by the acceptor).
	if assoc.UserIdentity != nil && assoc.UserIdentity.ServerResponse != nil {
		ac.UserInformation.UserIdentityResponse = &pdu.UserIdentityNegotiationResponse{
			ServerResponse: append([]byte{}, assoc.UserIdentity.ServerResponse...),
		}
	}

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
		if en == nil {
			continue
		}
		if len(en.RequestedApplicationInfo) > 0 {
			rq.UserInformation.ExtendedNegotiations = append(rq.UserInformation.ExtendedNegotiations, pdu.ExtendedNegotiation{
				SOPClassUID:         en.SOPClassUID,
				ServiceClassAppInfo: append([]byte(nil), en.RequestedApplicationInfo...),
			})
		}
		if en.HasCommonExtendedNegotiation() {
			rq.UserInformation.CommonExtendedNegotiations = append(
				rq.UserInformation.CommonExtendedNegotiations,
				pdu.CommonExtendedNegotiation{
					SOPClassUID:                en.SOPClassUID,
					ServiceClassUID:            en.ServiceClassUID,
					RelatedGeneralSOPClassUIDs: append([]string(nil), en.RelatedGeneralSOPClassUIDs...),
				},
			)
		}
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
