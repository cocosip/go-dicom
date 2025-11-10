// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// AAssociateRQ represents an A-ASSOCIATE-RQ PDU.
// This PDU is sent by the SCU to request an association with an SCP.
type AAssociateRQ struct {
	ProtocolVersion      uint16                  // Protocol version (always 0x0001)
	CalledAETitle        string                  // Called AE Title (max 16 bytes)
	CallingAETitle       string                  // Calling AE Title (max 16 bytes)
	ApplicationContext   string                  // Application Context Name (UID)
	PresentationContexts []PresentationContextRQ // Presentation contexts
	UserInformation      *UserInformation        // User information
}

// PresentationContextRQ represents a Presentation Context item in A-ASSOCIATE-RQ.
type PresentationContextRQ struct {
	ID               byte     // Presentation Context ID (odd numbers: 1, 3, 5, ...)
	AbstractSyntax   string   // Abstract Syntax UID (SOP Class UID)
	TransferSyntaxes []string // List of proposed Transfer Syntax UIDs
}

// UserInformation represents the User Information item.
type UserInformation struct {
	MaximumLength             uint32                        // Maximum PDU length (0 = unlimited)
	ImplementationClassUID    string                        // Implementation Class UID
	ImplementationVersionName string                        // Implementation Version Name
	AsynchronousOperations    *AsynchronousOperationsWindow // Optional: Async operations
	SCPSCURoleSelections      []SCPSCURoleSelection         // Optional: Role selections
	ExtendedNegotiations      []ExtendedNegotiation         // Optional: Extended negotiations
	UserIdentity              *UserIdentityNegotiation      // Optional: User identity
}

// AsynchronousOperationsWindow represents async operations sub-item.
type AsynchronousOperationsWindow struct {
	MaximumNumberOperationsInvoked   uint16
	MaximumNumberOperationsPerformed uint16
}

// SCPSCURoleSelection represents SCP/SCU role selection sub-item.
type SCPSCURoleSelection struct {
	SOPClassUID string
	SCURole     byte // 0 = non-support, 1 = support
	SCPRole     byte // 0 = non-support, 1 = support
}

// ExtendedNegotiation represents extended negotiation sub-item.
type ExtendedNegotiation struct {
	SOPClassUID         string
	ServiceClassAppInfo []byte // Application-specific information
}

// UserIdentityNegotiation represents user identity negotiation sub-item.
type UserIdentityNegotiation struct {
	UserIdentityType          byte   // 1=Username, 2=Username+Password, 3=Kerberos, 4=SAML, 5=JWT
	PositiveResponseRequested byte   // 0=no response, 1=response requested
	PrimaryField              []byte // Username or ticket
	SecondaryField            []byte // Password (only if Type=2)
}

// NewAAssociateRQ creates a new A-ASSOCIATE-RQ PDU with default values.
func NewAAssociateRQ() *AAssociateRQ {
	return &AAssociateRQ{
		ProtocolVersion:      0x0001,
		ApplicationContext:   "1.2.840.10008.3.1.1.1", // Default Application Context
		PresentationContexts: []PresentationContextRQ{},
		UserInformation: &UserInformation{
			MaximumLength:          16384,                                // Default 16KB
			ImplementationClassUID: "1.2.826.0.1.3680043.9.7278.1.0.0.0", // Placeholder
		},
	}
}

// Encode encodes the A-ASSOCIATE-RQ into a RawPDU.
func (a *AAssociateRQ) Encode() (*RawPDU, error) {
	var buf bytes.Buffer

	// Write fixed fields
	// Protocol version (2 bytes)
	if err := binary.Write(&buf, binary.BigEndian, a.ProtocolVersion); err != nil {
		return nil, fmt.Errorf("writing protocol version: %w", err)
	}

	// Reserved (2 bytes)
	buf.Write([]byte{0x00, 0x00})

	// Called AE Title (16 bytes, space-padded)
	buf.Write(writeString(a.CalledAETitle))

	// Calling AE Title (16 bytes, space-padded)
	buf.Write(writeString(a.CallingAETitle))

	// Reserved (32 bytes)
	buf.Write(make([]byte, 32))

	// Write variable items
	// Application Context Item
	if err := writeItem(&buf, ItemTypeApplicationContext, []byte(a.ApplicationContext)); err != nil {
		return nil, fmt.Errorf("writing application context: %w", err)
	}

	// Presentation Context Items
	for _, pc := range a.PresentationContexts {
		if err := a.encodePresentationContext(&buf, &pc); err != nil {
			return nil, fmt.Errorf("writing presentation context: %w", err)
		}
	}

	// User Information Item
	if a.UserInformation != nil {
		if err := a.encodeUserInformation(&buf, a.UserInformation); err != nil {
			return nil, fmt.Errorf("writing user information: %w", err)
		}
	}

	return NewRawPDU(TypeAAssociateRQ, buf.Bytes()), nil
}

// encodePresentationContext encodes a presentation context item.
func (a *AAssociateRQ) encodePresentationContext(w io.Writer, pc *PresentationContextRQ) error {
	var buf bytes.Buffer

	// Presentation Context ID (1 byte)
	buf.WriteByte(pc.ID)

	// Reserved (1 byte)
	buf.WriteByte(0x00)

	// Reserved (1 byte)
	buf.WriteByte(0x00)

	// Reserved (1 byte)
	buf.WriteByte(0x00)

	// Abstract Syntax Sub-item
	if err := writeItem(&buf, ItemTypeAbstractSyntax, []byte(pc.AbstractSyntax)); err != nil {
		return fmt.Errorf("writing abstract syntax: %w", err)
	}

	// Transfer Syntax Sub-items
	for _, ts := range pc.TransferSyntaxes {
		if err := writeItem(&buf, ItemTypeTransferSyntax, []byte(ts)); err != nil {
			return fmt.Errorf("writing transfer syntax: %w", err)
		}
	}

	// Write the presentation context item
	return writeItem(w, ItemTypePresentationContextRQ, buf.Bytes())
}

// encodeUserInformation encodes the user information item.
func (a *AAssociateRQ) encodeUserInformation(w io.Writer, ui *UserInformation) error {
	var buf bytes.Buffer

	// Maximum Length Sub-item
	maxLenData := make([]byte, 4)
	binary.BigEndian.PutUint32(maxLenData, ui.MaximumLength)
	if err := writeItem(&buf, ItemTypeMaximumLength, maxLenData); err != nil {
		return fmt.Errorf("writing maximum length: %w", err)
	}

	// Implementation Class UID Sub-item
	if ui.ImplementationClassUID != "" {
		if err := writeItem(&buf, ItemTypeImplementationClassUID, []byte(ui.ImplementationClassUID)); err != nil {
			return fmt.Errorf("writing implementation class UID: %w", err)
		}
	}

	// Implementation Version Name Sub-item (optional)
	if ui.ImplementationVersionName != "" {
		if err := writeItem(&buf, ItemTypeImplementationVersionName, []byte(ui.ImplementationVersionName)); err != nil {
			return fmt.Errorf("writing implementation version name: %w", err)
		}
	}

	// Asynchronous Operations Window Sub-item (optional)
	if ui.AsynchronousOperations != nil {
		asyncData := make([]byte, 4)
		binary.BigEndian.PutUint16(asyncData[0:2], ui.AsynchronousOperations.MaximumNumberOperationsInvoked)
		binary.BigEndian.PutUint16(asyncData[2:4], ui.AsynchronousOperations.MaximumNumberOperationsPerformed)
		if err := writeItem(&buf, ItemTypeAsynchronousOperationsWindow, asyncData); err != nil {
			return fmt.Errorf("writing async operations: %w", err)
		}
	}

	// SCP/SCU Role Selection Sub-items (optional)
	for _, role := range ui.SCPSCURoleSelections {
		var roleData bytes.Buffer
		// SOP Class UID length (2 bytes)
		binary.Write(&roleData, binary.BigEndian, uint16(len(role.SOPClassUID)))
		// SOP Class UID
		roleData.WriteString(role.SOPClassUID)
		// SCU Role (1 byte)
		roleData.WriteByte(role.SCURole)
		// SCP Role (1 byte)
		roleData.WriteByte(role.SCPRole)

		if err := writeItem(&buf, ItemTypeSCPSCURoleSelection, roleData.Bytes()); err != nil {
			return fmt.Errorf("writing role selection: %w", err)
		}
	}

	// Extended Negotiation Sub-items (optional)
	for _, ext := range ui.ExtendedNegotiations {
		var extData bytes.Buffer
		// SOP Class UID length (2 bytes)
		binary.Write(&extData, binary.BigEndian, uint16(len(ext.SOPClassUID)))
		// SOP Class UID
		extData.WriteString(ext.SOPClassUID)
		// Service Class Application Information
		extData.Write(ext.ServiceClassAppInfo)

		if err := writeItem(&buf, ItemTypeExtendedNegotiation, extData.Bytes()); err != nil {
			return fmt.Errorf("writing extended negotiation: %w", err)
		}
	}

	// User Identity Negotiation Sub-item (optional)
	if ui.UserIdentity != nil {
		var idData bytes.Buffer
		// User Identity Type (1 byte)
		idData.WriteByte(ui.UserIdentity.UserIdentityType)
		// Positive Response Requested (1 byte)
		idData.WriteByte(ui.UserIdentity.PositiveResponseRequested)
		// Primary Field Length (2 bytes)
		binary.Write(&idData, binary.BigEndian, uint16(len(ui.UserIdentity.PrimaryField)))
		// Primary Field
		idData.Write(ui.UserIdentity.PrimaryField)
		// Secondary Field Length (2 bytes)
		binary.Write(&idData, binary.BigEndian, uint16(len(ui.UserIdentity.SecondaryField)))
		// Secondary Field (only if present)
		if len(ui.UserIdentity.SecondaryField) > 0 {
			idData.Write(ui.UserIdentity.SecondaryField)
		}

		if err := writeItem(&buf, ItemTypeUserIdentityNegotiation, idData.Bytes()); err != nil {
			return fmt.Errorf("writing user identity: %w", err)
		}
	}

	// Write the user information item
	return writeItem(w, ItemTypeUserInformation, buf.Bytes())
}

// Decode decodes an A-ASSOCIATE-RQ from a RawPDU.
func (a *AAssociateRQ) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAAssociateRQ {
		return fmt.Errorf("invalid PDU type: expected 0x01, got 0x%02X", pdu.Type)
	}

	r := bytes.NewReader(pdu.Data)

	// Read fixed fields (68 bytes total)
	fixedFields := make([]byte, 68)
	if _, err := io.ReadFull(r, fixedFields); err != nil {
		return fmt.Errorf("reading fixed fields: %w", err)
	}

	// Protocol version (bytes 0-1)
	a.ProtocolVersion = binary.BigEndian.Uint16(fixedFields[0:2])

	// Reserved (bytes 2-3) - ignored

	// Called AE Title (bytes 4-19)
	a.CalledAETitle = trimTrailingSpaces(string(fixedFields[4:20]))

	// Calling AE Title (bytes 20-35)
	a.CallingAETitle = trimTrailingSpaces(string(fixedFields[20:36]))

	// Reserved (bytes 36-67) - ignored

	// Read variable items
	for {
		itemType, data, err := readItem(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("reading item: %w", err)
		}

		switch itemType {
		case ItemTypeApplicationContext:
			a.ApplicationContext = string(data)

		case ItemTypePresentationContextRQ:
			pc, err := a.decodePresentationContext(data)
			if err != nil {
				return fmt.Errorf("decoding presentation context: %w", err)
			}
			a.PresentationContexts = append(a.PresentationContexts, *pc)

		case ItemTypeUserInformation:
			ui, err := a.decodeUserInformation(data)
			if err != nil {
				return fmt.Errorf("decoding user information: %w", err)
			}
			a.UserInformation = ui

		default:
			// Unknown item type - skip
			fmt.Printf("Warning: Unknown item type 0x%02X in A-ASSOCIATE-RQ\n", itemType)
		}
	}

	return nil
}

// decodePresentationContext decodes a presentation context item.
func (a *AAssociateRQ) decodePresentationContext(data []byte) (*PresentationContextRQ, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("presentation context data too short: %d bytes", len(data))
	}

	pc := &PresentationContextRQ{
		ID:               data[0],
		TransferSyntaxes: []string{},
	}

	// Skip reserved bytes (1-3)
	r := bytes.NewReader(data[4:])

	// Read sub-items
	for {
		itemType, itemData, err := readItem(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading presentation context sub-item: %w", err)
		}

		switch itemType {
		case ItemTypeAbstractSyntax:
			pc.AbstractSyntax = string(itemData)
		case ItemTypeTransferSyntax:
			pc.TransferSyntaxes = append(pc.TransferSyntaxes, string(itemData))
		default:
			fmt.Printf("Warning: Unknown sub-item type 0x%02X in presentation context\n", itemType)
		}
	}

	return pc, nil
}

// decodeUserInformation decodes the user information item.
func (a *AAssociateRQ) decodeUserInformation(data []byte) (*UserInformation, error) {
	ui := &UserInformation{
		SCPSCURoleSelections: []SCPSCURoleSelection{},
		ExtendedNegotiations: []ExtendedNegotiation{},
	}

	r := bytes.NewReader(data)

	// Read sub-items
	for {
		itemType, itemData, err := readItem(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading user information sub-item: %w", err)
		}

		switch itemType {
		case ItemTypeMaximumLength:
			if len(itemData) >= 4 {
				ui.MaximumLength = binary.BigEndian.Uint32(itemData)
			}

		case ItemTypeImplementationClassUID:
			ui.ImplementationClassUID = string(itemData)

		case ItemTypeImplementationVersionName:
			ui.ImplementationVersionName = string(itemData)

		case ItemTypeAsynchronousOperationsWindow:
			if len(itemData) >= 4 {
				ui.AsynchronousOperations = &AsynchronousOperationsWindow{
					MaximumNumberOperationsInvoked:   binary.BigEndian.Uint16(itemData[0:2]),
					MaximumNumberOperationsPerformed: binary.BigEndian.Uint16(itemData[2:4]),
				}
			}

		case ItemTypeSCPSCURoleSelection:
			role, err := decodeRoleSelection(itemData)
			if err != nil {
				return nil, fmt.Errorf("decoding role selection: %w", err)
			}
			ui.SCPSCURoleSelections = append(ui.SCPSCURoleSelections, *role)

		case ItemTypeExtendedNegotiation:
			ext, err := decodeExtendedNegotiation(itemData)
			if err != nil {
				return nil, fmt.Errorf("decoding extended negotiation: %w", err)
			}
			ui.ExtendedNegotiations = append(ui.ExtendedNegotiations, *ext)

		case ItemTypeUserIdentityNegotiation:
			identity, err := decodeUserIdentity(itemData)
			if err != nil {
				return nil, fmt.Errorf("decoding user identity: %w", err)
			}
			ui.UserIdentity = identity

		default:
			fmt.Printf("Warning: Unknown sub-item type 0x%02X in user information\n", itemType)
		}
	}

	return ui, nil
}

// decodeRoleSelection decodes a role selection sub-item.
func decodeRoleSelection(data []byte) (*SCPSCURoleSelection, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("role selection data too short")
	}

	r := bytes.NewReader(data)

	// SOP Class UID Length
	var uidLen uint16
	if err := binary.Read(r, binary.BigEndian, &uidLen); err != nil {
		return nil, err
	}

	// SOP Class UID
	uidBytes := make([]byte, uidLen)
	if _, err := io.ReadFull(r, uidBytes); err != nil {
		return nil, err
	}

	// SCU Role and SCP Role
	var scuRole, scpRole byte
	if err := binary.Read(r, binary.BigEndian, &scuRole); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &scpRole); err != nil {
		return nil, err
	}

	return &SCPSCURoleSelection{
		SOPClassUID: string(uidBytes),
		SCURole:     scuRole,
		SCPRole:     scpRole,
	}, nil
}

// decodeExtendedNegotiation decodes an extended negotiation sub-item.
func decodeExtendedNegotiation(data []byte) (*ExtendedNegotiation, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("extended negotiation data too short")
	}

	r := bytes.NewReader(data)

	// SOP Class UID Length
	var uidLen uint16
	if err := binary.Read(r, binary.BigEndian, &uidLen); err != nil {
		return nil, err
	}

	// SOP Class UID
	uidBytes := make([]byte, uidLen)
	if _, err := io.ReadFull(r, uidBytes); err != nil {
		return nil, err
	}

	// Service Class Application Information (rest of data)
	appInfo := make([]byte, r.Len())
	if _, err := io.ReadFull(r, appInfo); err != nil && err != io.EOF {
		return nil, err
	}

	return &ExtendedNegotiation{
		SOPClassUID:         string(uidBytes),
		ServiceClassAppInfo: appInfo,
	}, nil
}

// decodeUserIdentity decodes a user identity negotiation sub-item.
func decodeUserIdentity(data []byte) (*UserIdentityNegotiation, error) {
	if len(data) < 6 {
		return nil, fmt.Errorf("user identity data too short")
	}

	r := bytes.NewReader(data)

	identity := &UserIdentityNegotiation{}

	// User Identity Type
	if err := binary.Read(r, binary.BigEndian, &identity.UserIdentityType); err != nil {
		return nil, err
	}

	// Positive Response Requested
	if err := binary.Read(r, binary.BigEndian, &identity.PositiveResponseRequested); err != nil {
		return nil, err
	}

	// Primary Field Length
	var primaryLen uint16
	if err := binary.Read(r, binary.BigEndian, &primaryLen); err != nil {
		return nil, err
	}

	// Primary Field
	identity.PrimaryField = make([]byte, primaryLen)
	if _, err := io.ReadFull(r, identity.PrimaryField); err != nil {
		return nil, err
	}

	// Secondary Field Length
	var secondaryLen uint16
	if err := binary.Read(r, binary.BigEndian, &secondaryLen); err != nil {
		return nil, err
	}

	// Secondary Field (if present)
	if secondaryLen > 0 {
		identity.SecondaryField = make([]byte, secondaryLen)
		if _, err := io.ReadFull(r, identity.SecondaryField); err != nil {
			return nil, err
		}
	}

	return identity, nil
}
