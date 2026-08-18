// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// AAssociateAC represents an A-ASSOCIATE-AC (Association Accept) PDU.
// This is the response to an A-ASSOCIATE-RQ, indicating the remote AE has accepted the association.
//
// PDU Structure (variable length):
//   - Bytes 0-5: PDU header (Type=0x02, Reserved, Length)
//   - Bytes 6-7: Protocol Version (0x0001)
//   - Bytes 8-9: Reserved (0x0000)
//   - Bytes 10-25: Called AE Title (16 bytes, space-padded)
//   - Bytes 26-41: Calling AE Title (16 bytes, space-padded)
//   - Bytes 42-73: Reserved (32 bytes, all 0x00)
//   - Bytes 74+: Variable items (Application Context, Presentation Contexts, User Information)
type AAssociateAC struct {
	// ProtocolVersion is the DICOM protocol version (always 0x0001)
	ProtocolVersion uint16

	// CalledAETitle is the AE title of the called application (max 16 bytes)
	CalledAETitle string

	// CallingAETitle is the AE title of the calling application (max 16 bytes)
	CallingAETitle string

	// ApplicationContext is the application context UID
	// Usually "1.2.840.10008.3.1.1.1" (DICOM Application Context)
	ApplicationContext string

	// PresentationContexts contains the negotiated presentation contexts
	PresentationContexts []PresentationContextAC

	// UserInformation contains optional user information items
	UserInformation *UserInformation
}

// PresentationContextAC represents an accepted presentation context in A-ASSOCIATE-AC.
type PresentationContextAC struct {
	// ID is the presentation context ID (odd numbers: 1, 3, 5, ...)
	// Must match the ID from the corresponding A-ASSOCIATE-RQ
	ID byte

	// Result indicates the acceptance status:
	//   0 = acceptance
	//   1 = user-rejection
	//   2 = no-reason (provider rejection)
	//   3 = abstract-syntax-not-supported (provider rejection)
	//   4 = transfer-syntaxes-not-supported (provider rejection)
	Result byte

	// TransferSyntax is the selected transfer syntax UID
	// Only present if Result = 0 (acceptance)
	TransferSyntax string
}

// Result codes for PresentationContextAC
const (
	ResultAcceptance                   byte = 0
	ResultUserRejection                byte = 1
	ResultNoReason                     byte = 2
	ResultAbstractSyntaxNotSupported   byte = 3
	ResultTransferSyntaxesNotSupported byte = 4
)

// ResultString returns the string representation of a result code.
func ResultString(result byte) string {
	switch result {
	case ResultAcceptance:
		return "Acceptance"
	case ResultUserRejection:
		return "User Rejection"
	case ResultNoReason:
		return "No Reason (Provider Rejection)"
	case ResultAbstractSyntaxNotSupported:
		return "Abstract Syntax Not Supported (Provider Rejection)"
	case ResultTransferSyntaxesNotSupported:
		return "Transfer Syntaxes Not Supported (Provider Rejection)"
	default:
		return fmt.Sprintf("Unknown(0x%02X)", result)
	}
}

// NewAAssociateAC creates a new A-ASSOCIATE-AC PDU with default values.
func NewAAssociateAC() *AAssociateAC {
	return &AAssociateAC{
		ProtocolVersion:      0x0001,
		ApplicationContext:   applicationContextUID, // DICOM Application Context
		PresentationContexts: make([]PresentationContextAC, 0),
		UserInformation: &UserInformation{
			MaximumLength:          16384,                         // Default 16KB
			ImplementationClassUID: "1.2.826.0.1.3680043.8.498.1", // go-dicom implementation UID
		},
	}
}

// Encode encodes the A-ASSOCIATE-AC to a RawPDU.
func (a *AAssociateAC) Encode() (*RawPDU, error) {
	if err := validateAETitles(a.CalledAETitle, a.CallingAETitle); err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	// Protocol Version (2 bytes)
	if err := binary.Write(&buf, binary.BigEndian, a.ProtocolVersion); err != nil {
		return nil, fmt.Errorf("encoding protocol version: %w", err)
	}

	// Reserved (2 bytes)
	buf.Write([]byte{0x00, 0x00})

	// Called AE Title (16 bytes, space-padded)
	buf.Write(writeString(a.CalledAETitle))

	// Calling AE Title (16 bytes, space-padded)
	buf.Write(writeString(a.CallingAETitle))

	// Reserved (32 bytes)
	buf.Write(make([]byte, 32))

	// Application Context Item
	appContextData := []byte(a.ApplicationContext)
	if err := writeItem(&buf, ItemTypeApplicationContext, appContextData); err != nil {
		return nil, fmt.Errorf("encoding application context: %w", err)
	}

	// Presentation Context Items (AC format)
	for i, pc := range a.PresentationContexts {
		// Presentation Context Item
		var pcBuf bytes.Buffer

		// Context ID (1 byte)
		pcBuf.WriteByte(pc.ID)

		// Reserved (1 byte)
		pcBuf.WriteByte(0x00)

		// Result/Reason (1 byte)
		pcBuf.WriteByte(pc.Result)

		// Reserved (1 byte)
		pcBuf.WriteByte(0x00)

		// Transfer Syntax Sub-item (only if accepted)
		if pc.Result == ResultAcceptance && pc.TransferSyntax != "" {
			tsData := []byte(pc.TransferSyntax)
			if err := writeItem(&pcBuf, ItemTypeTransferSyntax, tsData); err != nil {
				return nil, fmt.Errorf("encoding transfer syntax for context %d: %w", i, err)
			}
		}

		if err := writeItem(&buf, ItemTypePresentationContextAC, pcBuf.Bytes()); err != nil {
			return nil, fmt.Errorf("encoding presentation context %d: %w", i, err)
		}
	}

	// User Information Item
	if a.UserInformation != nil {
		userInfoData, err := a.encodeUserInformation()
		if err != nil {
			return nil, fmt.Errorf("encoding user information: %w", err)
		}
		if err := writeItem(&buf, ItemTypeUserInformation, userInfoData); err != nil {
			return nil, fmt.Errorf("writing user information item: %w", err)
		}
	}

	return NewRawPDU(TypeAAssociateAC, buf.Bytes()), nil
}

// encodeUserInformation encodes the UserInformation structure.
// This is shared between A-ASSOCIATE-RQ and A-ASSOCIATE-AC.
func (a *AAssociateAC) encodeUserInformation() ([]byte, error) {
	var buf bytes.Buffer

	ui := a.UserInformation

	// Maximum Length Sub-item (always present)
	maxLenData := make([]byte, 4)
	binary.BigEndian.PutUint32(maxLenData, ui.MaximumLength)
	if err := writeItem(&buf, ItemTypeMaximumLength, maxLenData); err != nil {
		return nil, fmt.Errorf("encoding maximum length: %w", err)
	}

	// Implementation Class UID Sub-item (always present)
	if ui.ImplementationClassUID != "" {
		implClassData := []byte(ui.ImplementationClassUID)
		if err := writeItem(&buf, ItemTypeImplementationClassUID, implClassData); err != nil {
			return nil, fmt.Errorf("encoding implementation class UID: %w", err)
		}
	}

	// Implementation Version Name Sub-item (optional)
	if ui.ImplementationVersionName != "" {
		implVersionData := []byte(ui.ImplementationVersionName)
		if err := writeItem(&buf, ItemTypeImplementationVersionName, implVersionData); err != nil {
			return nil, fmt.Errorf("encoding implementation version name: %w", err)
		}
	}

	// Asynchronous Operations Window Sub-item (optional)
	if ui.AsynchronousOperations != nil {
		asyncData := make([]byte, 4)
		binary.BigEndian.PutUint16(asyncData[0:2], ui.AsynchronousOperations.MaximumNumberOperationsInvoked)
		binary.BigEndian.PutUint16(asyncData[2:4], ui.AsynchronousOperations.MaximumNumberOperationsPerformed)
		if err := writeItem(&buf, ItemTypeAsynchronousOperationsWindow, asyncData); err != nil {
			return nil, fmt.Errorf("encoding async operations: %w", err)
		}
	}

	// SCP/SCU Role Selection Sub-items (optional)
	for _, role := range ui.SCPSCURoleSelections {
		if len(role.SOPClassUID) > int(math.MaxUint16) {
			return nil, fmt.Errorf("SOP Class UID length too large: %d", len(role.SOPClassUID))
		}
		var roleData bytes.Buffer
		_ = binary.Write(&roleData, binary.BigEndian, uint16(len(role.SOPClassUID))) // #nosec G115 -- bounded above
		roleData.WriteString(role.SOPClassUID)
		roleData.WriteByte(role.SCURole)
		roleData.WriteByte(role.SCPRole)
		if err := writeItem(&buf, ItemTypeSCPSCURoleSelection, roleData.Bytes()); err != nil {
			return nil, fmt.Errorf("encoding role selection: %w", err)
		}
	}

	// SOP Class Extended Negotiation Sub-items (optional)
	for _, negotiation := range ui.ExtendedNegotiations {
		if len(negotiation.SOPClassUID) > int(math.MaxUint16) {
			return nil, fmt.Errorf("extended negotiation SOP Class UID length too large: %d", len(negotiation.SOPClassUID))
		}
		var negotiationData bytes.Buffer
		_ = binary.Write(&negotiationData, binary.BigEndian, uint16(len(negotiation.SOPClassUID))) // #nosec G115 -- bounded above
		negotiationData.WriteString(negotiation.SOPClassUID)
		negotiationData.Write(negotiation.ServiceClassAppInfo)
		if err := writeItem(&buf, ItemTypeExtendedNegotiation, negotiationData.Bytes()); err != nil {
			return nil, fmt.Errorf("encoding extended negotiation: %w", err)
		}
	}

	// User Identity Negotiation Response (optional).
	response := ui.UserIdentityResponse
	if response == nil && ui.UserIdentity != nil && ui.UserIdentity.PositiveResponseRequested == 1 {
		// Preserve compatibility with callers that used the request structure to
		// represent an AC response before UserIdentityResponse was available.
		response = &UserIdentityNegotiationResponse{ServerResponse: ui.UserIdentity.PrimaryField}
	}
	if response != nil {
		var uiBuf bytes.Buffer

		// Server response length (2 bytes)
		if len(response.ServerResponse) > int(math.MaxUint16) {
			return nil, fmt.Errorf("user identity response length too large: %d", len(response.ServerResponse))
		}
		serverResponseLen := uint16(len(response.ServerResponse)) // #nosec G115 -- bounded by check above
		if err := binary.Write(&uiBuf, binary.BigEndian, serverResponseLen); err != nil {
			return nil, fmt.Errorf("encoding user identity response length: %w", err)
		}

		// Server response data
		if serverResponseLen > 0 {
			uiBuf.Write(response.ServerResponse)
		}

		if err := writeItem(&buf, ItemTypeUserIdentityNegotiationResponse, uiBuf.Bytes()); err != nil {
			return nil, fmt.Errorf("encoding user identity response: %w", err)
		}
	}

	return buf.Bytes(), nil
}

// Decode decodes an A-ASSOCIATE-AC from a RawPDU.
func (a *AAssociateAC) Decode(pdu *RawPDU) error {
	if pdu.Type != TypeAAssociateAC {
		return fmt.Errorf("invalid PDU type: expected A-ASSOCIATE-AC (0x02), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	if len(data) < 68 {
		return fmt.Errorf("invalid A-ASSOCIATE-AC data length: %d (minimum 68 bytes)", len(data))
	}

	// Protocol Version (2 bytes)
	a.ProtocolVersion = binary.BigEndian.Uint16(data[0:2])

	// Reserved (2 bytes) - skip

	// Called AE Title (16 bytes)
	a.CalledAETitle = trimTrailingSpaces(string(data[4:20]))

	// Calling AE Title (16 bytes)
	a.CallingAETitle = trimTrailingSpaces(string(data[20:36]))

	// Reserved (32 bytes) - skip

	// Parse variable items
	itemsData := data[68:]
	reader := bytes.NewReader(itemsData)

	for {
		itemType, itemData, err := readItem(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("reading item: %w", err)
		}

		switch itemType {
		case ItemTypeApplicationContext:
			a.ApplicationContext = string(itemData)

		case ItemTypePresentationContextAC:
			pc, err := a.decodePresentationContextAC(itemData)
			if err != nil {
				return fmt.Errorf("decoding presentation context: %w", err)
			}
			a.PresentationContexts = append(a.PresentationContexts, pc)

		case ItemTypeUserInformation:
			ui, err := a.decodeUserInformation(itemData)
			if err != nil {
				return fmt.Errorf("decoding user information: %w", err)
			}
			a.UserInformation = ui

		default:
			// Ignore unknown item types
		}
	}

	return nil
}

// decodePresentationContextAC decodes a Presentation Context (AC) item.
func (a *AAssociateAC) decodePresentationContextAC(data []byte) (PresentationContextAC, error) {
	if len(data) < 4 {
		return PresentationContextAC{}, fmt.Errorf("invalid presentation context data length: %d", len(data))
	}

	pc := PresentationContextAC{
		ID:     data[0],
		Result: data[2],
	}

	// Parse sub-items (only Transfer Syntax for accepted contexts)
	if len(data) > 4 {
		reader := bytes.NewReader(data[4:])
		for {
			itemType, itemData, err := readItem(reader)
			if err != nil {
				if err == io.EOF {
					break
				}
				return PresentationContextAC{}, fmt.Errorf("reading sub-item: %w", err)
			}

			switch itemType {
			case ItemTypeTransferSyntax:
				pc.TransferSyntax = string(itemData)
			}
		}
	}

	return pc, nil
}

// decodeUserInformation decodes the User Information item.
func (a *AAssociateAC) decodeUserInformation(data []byte) (*UserInformation, error) {
	ui := &UserInformation{}
	reader := bytes.NewReader(data)

	for {
		itemType, itemData, err := readItem(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading user info sub-item: %w", err)
		}

		switch itemType {
		case ItemTypeMaximumLength:
			if len(itemData) != 4 {
				return nil, fmt.Errorf("invalid maximum length data: %d bytes", len(itemData))
			}
			ui.MaximumLength = binary.BigEndian.Uint32(itemData)

		case ItemTypeImplementationClassUID:
			ui.ImplementationClassUID = string(itemData)

		case ItemTypeImplementationVersionName:
			ui.ImplementationVersionName = string(itemData)

		case ItemTypeAsynchronousOperationsWindow:
			if len(itemData) != 4 {
				return nil, fmt.Errorf("invalid asynchronous operations window: %d bytes", len(itemData))
			}
			ui.AsynchronousOperations = &AsynchronousOperationsWindow{
				MaximumNumberOperationsInvoked:   binary.BigEndian.Uint16(itemData[0:2]),
				MaximumNumberOperationsPerformed: binary.BigEndian.Uint16(itemData[2:4]),
			}

		case ItemTypeSCPSCURoleSelection:
			role, err := decodeRoleSelection(itemData)
			if err != nil {
				return nil, fmt.Errorf("decoding role selection: %w", err)
			}
			ui.SCPSCURoleSelections = append(ui.SCPSCURoleSelections, *role)

		case ItemTypeExtendedNegotiation:
			negotiation, err := decodeExtendedNegotiation(itemData)
			if err != nil {
				return nil, fmt.Errorf("decoding extended negotiation: %w", err)
			}
			ui.ExtendedNegotiations = append(ui.ExtendedNegotiations, *negotiation)

		case ItemTypeUserIdentityNegotiationResponse:
			// Server response to user identity
			if len(itemData) < 2 {
				return nil, fmt.Errorf("invalid user identity response: %d bytes", len(itemData))
			}
			serverResponseLen := binary.BigEndian.Uint16(itemData[0:2])
			if len(itemData) != int(2+serverResponseLen) {
				return nil, fmt.Errorf("invalid user identity response length")
			}
			serverResponse := make([]byte, int(serverResponseLen))
			copy(serverResponse, itemData[2:])
			ui.UserIdentityResponse = &UserIdentityNegotiationResponse{ServerResponse: serverResponse}
			// Preserve the legacy decoded representation.
			ui.UserIdentity = &UserIdentityNegotiation{
				PositiveResponseRequested: 1, // Response was provided
				PrimaryField:              serverResponse,
			}
		}
	}

	return ui, nil
}
