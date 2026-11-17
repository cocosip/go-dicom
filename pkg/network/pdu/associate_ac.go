// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
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
	UserInformation         *UserInformation
	applicationContextCount int
	userInformationCount    int

	// CalledAETitleRaw and CallingAETitleRaw, when 16 bytes long, are written
	// verbatim into the Reserved AE Title fields. PS3.8 Table 9-17 requires the
	// value received in the A-ASSOCIATE-RQ, so an acceptor copies
	// AAssociateRQ.CalledAETitleRaw/CallingAETitleRaw here; otherwise the
	// trimmed titles are space-padded as usual.
	CalledAETitleRaw  []byte
	CallingAETitleRaw []byte
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

// String returns a compact summary of the association acceptance.
func (a *AAssociateAC) String() string {
	if a == nil {
		return "<nil A-ASSOCIATE-AC>"
	}
	accepted := 0
	for _, context := range a.PresentationContexts {
		if context.Result == ResultAcceptance {
			accepted++
		}
	}
	maximumLength := uint32(0)
	if a.UserInformation != nil {
		maximumLength = a.UserInformation.MaximumLength
	}
	contextLabel := "contexts"
	if len(a.PresentationContexts) == 1 {
		contextLabel = "context"
	}
	return fmt.Sprintf("A-ASSOCIATE-AC [%s -> %s, %d/%d %s accepted, MaxPDU=%d]",
		a.CallingAETitle, a.CalledAETitle, accepted, len(a.PresentationContexts), contextLabel, maximumLength)
}

// String returns a compact summary of the accepted presentation context.
func (pc PresentationContextAC) String() string {
	transferSyntax := pc.TransferSyntax
	if transferSyntax == "" {
		transferSyntax = "<none>"
	}
	return fmt.Sprintf("PC[%d] -> %s [%s]", pc.ID, transferSyntax, ResultString(pc.Result))
}

// NewAAssociateAC creates a new A-ASSOCIATE-AC PDU with default values.
func NewAAssociateAC() *AAssociateAC {
	return &AAssociateAC{
		ProtocolVersion:      0x0001,
		ApplicationContext:   applicationContextUID, // DICOM Application Context
		PresentationContexts: make([]PresentationContextAC, 0),
		UserInformation: &UserInformation{
			MaximumLength:          16384, // Default 16KB
			ImplementationClassUID: DefaultImplementationClassUID,
			hasMaximumLength:       true,
		},
	}
}

// validateForEncode checks the mandatory A-ASSOCIATE-AC structure before
// encoding (PS3.8 Table 9-17, PS3.7 Annex D.3.3.2), so the encoder never emits
// a PDU that a conformant peer is entitled to reject.
func (a *AAssociateAC) validateForEncode() error {
	if a == nil {
		return fmt.Errorf("A-ASSOCIATE-AC is nil")
	}
	if a.ProtocolVersion&0x0001 == 0 {
		return fmt.Errorf("A-ASSOCIATE-AC protocol version 0x%04X does not support version 1", a.ProtocolVersion)
	}
	// PS3.8 Table 9-17: the Called/Calling AE Title fields of an A-ASSOCIATE-AC are
	// Reserved and "shall not be tested when received", so they are not AE-validated
	// here; writeString pads/truncates them to the fixed 16-byte width safely.
	if err := validateApplicationContext(a.ApplicationContext); err != nil {
		return fmt.Errorf("A-ASSOCIATE-AC: %w", err)
	}
	if a.UserInformation == nil {
		return fmt.Errorf("A-ASSOCIATE-AC is missing user information")
	}
	if err := validateImplementationIdentification(a.UserInformation, "A-ASSOCIATE-AC"); err != nil {
		return err
	}
	if len(a.PresentationContexts) == 0 {
		return fmt.Errorf("A-ASSOCIATE-AC must contain at least one presentation context")
	}
	return validatePresentationContextsAC(a.PresentationContexts)
}

// Encode encodes the A-ASSOCIATE-AC to a RawPDU.
func (a *AAssociateAC) Encode() (*RawPDU, error) {
	if err := a.validateForEncode(); err != nil {
		return nil, err
	}
	// The encoder always writes a Maximum Length sub-item, so mark it present
	// to keep ValidateStructure() consistent for programmatically built PDUs.
	if a.UserInformation != nil {
		a.UserInformation.hasMaximumLength = true
	}
	var buf bytes.Buffer

	// Protocol Version (2 bytes)
	if err := binary.Write(&buf, binary.BigEndian, a.ProtocolVersion); err != nil {
		return nil, fmt.Errorf("encoding protocol version: %w", err)
	}

	// Reserved (2 bytes)
	buf.Write([]byte{0x00, 0x00})

	// Called AE Title (16 bytes, space-padded)
	writeAETitleField(&buf, a.CalledAETitle, a.CalledAETitleRaw)

	// Calling AE Title (16 bytes, space-padded)
	writeAETitleField(&buf, a.CallingAETitle, a.CallingAETitleRaw)

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

		// Transfer Syntax Sub-item. PS3.8 Table 9-18 requires exactly one to be
		// present even for a rejected context; a zero-length value is invalid,
		// so fall back to Implicit VR Little Endian when none was set.
		transferSyntax := pc.TransferSyntax
		if transferSyntax == "" {
			transferSyntax = uid.ImplicitVRLittleEndian.UID()
		}
		tsData := []byte(transferSyntax)
		if err := writeItem(&pcBuf, ItemTypeTransferSyntax, tsData); err != nil {
			return nil, fmt.Errorf("encoding transfer syntax for context %d: %w", i, err)
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

// encodeUserInformation encodes the User Information item; the sub-items are
// shared with A-ASSOCIATE-RQ (see encodeUserInformationItems).
func (a *AAssociateAC) encodeUserInformation() ([]byte, error) {
	return encodeUserInformationItems(a.UserInformation, false)
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

	// A decoded PDU carries parsed state only, so reset it before parsing to
	// keep two decodes on the same receiver independent.
	*a = AAssociateAC{}

	// Protocol Version (2 bytes)
	a.ProtocolVersion = binary.BigEndian.Uint16(data[0:2])

	// Reserved (2 bytes) - skip

	// Called AE Title (16 bytes)
	a.CalledAETitle = trimAETitleSpaces(string(data[4:20]))

	// Calling AE Title (16 bytes)
	a.CallingAETitle = trimAETitleSpaces(string(data[20:36]))

	// Reserved (32 bytes) - skip

	applicationContextCount := 0
	userInformationCount := 0
	seenPresentationContexts := make(map[byte]struct{})
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
			applicationContextCount++
			// UIDs are padded with a trailing NUL and some peers use spaces;
			// trim both before comparing against the DICOM Application Context.
			a.ApplicationContext = trimUIDPadding(string(itemData))

		case ItemTypePresentationContextAC:
			// Bound the count while decoding: the association PDU is read with a
			// hard size cap, so without this a peer could make the decoder
			// allocate millions of context structs.
			if len(a.PresentationContexts) >= maximumPresentationContexts {
				return fmt.Errorf("%w: got more than %d", ErrTooManyPresentationContexts, maximumPresentationContexts)
			}
			pc, err := a.decodePresentationContextAC(itemData)
			if err != nil {
				return fmt.Errorf("decoding presentation context: %w", err)
			}
			if _, exists := seenPresentationContexts[pc.ID]; exists {
				return fmt.Errorf("duplicate presentation context ID %d", pc.ID)
			}
			seenPresentationContexts[pc.ID] = struct{}{}
			a.PresentationContexts = append(a.PresentationContexts, pc)

		case ItemTypeUserInformation:
			userInformationCount++
			ui, err := a.decodeUserInformation(itemData)
			if err != nil {
				return fmt.Errorf("decoding user information: %w", err)
			}
			a.UserInformation = ui

		default:
			// Ignore unknown item types
		}
	}

	a.applicationContextCount = applicationContextCount
	a.userInformationCount = userInformationCount
	return nil
}

// ValidateStructure validates the mandatory A-ASSOCIATE-AC structure.
func (a *AAssociateAC) ValidateStructure() error {
	if a == nil {
		return fmt.Errorf("A-ASSOCIATE-AC is nil")
	}
	// PS3.8 9.3.3: a receiver shall only test that protocol-version bit 0 is set.
	if a.ProtocolVersion&0x0001 == 0 {
		return fmt.Errorf("A-ASSOCIATE-AC protocol version 0x%04X does not support version 1", a.ProtocolVersion)
	}
	if a.applicationContextCount != 0 && a.applicationContextCount != 1 {
		return fmt.Errorf("A-ASSOCIATE-AC must contain exactly one application context item, got %d", a.applicationContextCount)
	}
	if err := validateApplicationContext(a.ApplicationContext); err != nil {
		return fmt.Errorf("A-ASSOCIATE-AC: %w", err)
	}
	if len(a.PresentationContexts) == 0 {
		return fmt.Errorf("A-ASSOCIATE-AC must contain at least one presentation context")
	}
	if a.userInformationCount != 0 && a.userInformationCount != 1 {
		return fmt.Errorf("A-ASSOCIATE-AC must contain exactly one user information item, got %d", a.userInformationCount)
	}
	if a.UserInformation == nil || !a.UserInformation.hasMaximumLength {
		return fmt.Errorf("A-ASSOCIATE-AC is missing maximum length user information")
	}
	if err := validateImplementationIdentification(a.UserInformation, "A-ASSOCIATE-AC"); err != nil {
		return err
	}
	return nil
}

// validatePresentationContextsAC validates presentation contexts before encoding
// an A-ASSOCIATE-AC, mirroring the checks performed for A-ASSOCIATE-RQ so this
// library never emits an AC that its own decoder would reject.
func validatePresentationContextsAC(contexts []PresentationContextAC) error {
	if len(contexts) > maximumPresentationContexts {
		return fmt.Errorf("%w: got %d, maximum is %d", ErrTooManyPresentationContexts, len(contexts), maximumPresentationContexts)
	}
	seen := make(map[byte]struct{}, len(contexts))
	for _, pc := range contexts {
		if pc.ID == 0 || pc.ID%2 == 0 {
			return fmt.Errorf("%w: ID must be a non-zero odd value, got %d", ErrInvalidPresentationContext, pc.ID)
		}
		if _, exists := seen[pc.ID]; exists {
			return fmt.Errorf("%w: duplicate ID %d", ErrInvalidPresentationContext, pc.ID)
		}
		seen[pc.ID] = struct{}{}
		if pc.Result > ResultTransferSyntaxesNotSupported {
			return fmt.Errorf("%w: invalid result %d for ID %d", ErrInvalidPresentationContext, pc.Result, pc.ID)
		}
		if pc.Result == ResultAcceptance && pc.TransferSyntax == "" {
			return fmt.Errorf("%w: accepted context %d has no transfer syntax", ErrInvalidPresentationContext, pc.ID)
		}
		if pc.TransferSyntax != "" && !uid.IsValid(pc.TransferSyntax) {
			return fmt.Errorf("%w: invalid transfer syntax %q for ID %d", ErrInvalidPresentationContext, pc.TransferSyntax, pc.ID)
		}
		// A Transfer Syntax value is always written by the encoder; an empty
		// value on a rejected context is replaced by the default syntax there.
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
	if pc.ID == 0 || pc.ID%2 == 0 {
		return PresentationContextAC{}, fmt.Errorf("invalid presentation context ID %d", pc.ID)
	}
	if pc.Result > ResultTransferSyntaxesNotSupported {
		return PresentationContextAC{}, fmt.Errorf("invalid presentation context result %d", pc.Result)
	}

	// Parse sub-items. PS3.8 Table 9-18: each presentation context shall contain
	// exactly one Transfer Syntax Sub-Item, even when the context is rejected.
	transferSyntaxCount := 0
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
				transferSyntaxCount++
				if transferSyntaxCount > 1 {
					return PresentationContextAC{}, fmt.Errorf("presentation context %d has more than one Transfer Syntax Sub-Item", pc.ID)
				}
				pc.TransferSyntax = trimUIDPadding(string(itemData))
			}
		}
	}
	if transferSyntaxCount != 1 {
		return PresentationContextAC{}, fmt.Errorf("presentation context %d must contain exactly one Transfer Syntax Sub-Item", pc.ID)
	}
	if pc.Result == ResultAcceptance && pc.TransferSyntax == "" {
		return PresentationContextAC{}, fmt.Errorf("accepted presentation context %d has no transfer syntax", pc.ID)
	}
	// PS3.8 Table 9-18 requires a Transfer Syntax Sub-Item even for a rejected
	// context; its value is not significant in that case, so drop it.
	if pc.Result != ResultAcceptance {
		pc.TransferSyntax = ""
	}

	return pc, nil
}

// decodeUserInformation decodes the User Information item; the sub-items are
// shared with A-ASSOCIATE-RQ (see decodeUserInformationItems).
func (a *AAssociateAC) decodeUserInformation(data []byte) (*UserInformation, error) {
	return decodeUserInformationItems(data, false, nil)
}
