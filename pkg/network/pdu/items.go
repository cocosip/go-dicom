// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Item type constants for A-ASSOCIATE PDUs
const (
	// ItemTypeApplicationContext represents Application Context Item (0x10)
	ItemTypeApplicationContext byte = 0x10

	// ItemTypePresentationContextRQ represents Presentation Context Item (RQ) (0x20)
	ItemTypePresentationContextRQ byte = 0x20

	// ItemTypePresentationContextAC represents Presentation Context Item (AC) (0x21)
	ItemTypePresentationContextAC byte = 0x21

	// ItemTypeAbstractSyntax represents Abstract Syntax Sub-item (0x30)
	ItemTypeAbstractSyntax byte = 0x30

	// ItemTypeTransferSyntax represents Transfer Syntax Sub-item (0x40)
	ItemTypeTransferSyntax byte = 0x40

	// ItemTypeUserInformation represents User Information Item (0x50)
	ItemTypeUserInformation byte = 0x50

	// ItemTypeMaximumLength represents Maximum Length Sub-item (0x51)
	ItemTypeMaximumLength byte = 0x51

	// ItemTypeImplementationClassUID represents Implementation Class UID Sub-item (0x52)
	ItemTypeImplementationClassUID byte = 0x52

	// ItemTypeAsynchronousOperationsWindow represents Asynchronous Operations Window Sub-item (0x53)
	ItemTypeAsynchronousOperationsWindow byte = 0x53

	// ItemTypeSCPSCURoleSelection represents SCP/SCU Role Selection Sub-item (0x54)
	ItemTypeSCPSCURoleSelection byte = 0x54

	// ItemTypeImplementationVersionName represents Implementation Version Name Sub-item (0x55)
	ItemTypeImplementationVersionName byte = 0x55

	// ItemTypeExtendedNegotiation represents Extended Negotiation Sub-item (0x56)
	ItemTypeExtendedNegotiation byte = 0x56

	// ItemTypeCommonExtendedNegotiation represents Common Extended Negotiation Sub-item (0x57)
	ItemTypeCommonExtendedNegotiation byte = 0x57

	// ItemTypeUserIdentityNegotiation represents User Identity Negotiation Sub-item (0x58)
	ItemTypeUserIdentityNegotiation byte = 0x58

	// ItemTypeUserIdentityNegotiationResponse represents User Identity Negotiation Response Sub-item (0x59)
	ItemTypeUserIdentityNegotiationResponse byte = 0x59
)

// ItemTypeString returns the string representation of an item type.
func ItemTypeString(itemType byte) string {
	switch itemType {
	case ItemTypeApplicationContext:
		return "Application Context"
	case ItemTypePresentationContextRQ:
		return "Presentation Context (RQ)"
	case ItemTypePresentationContextAC:
		return "Presentation Context (AC)"
	case ItemTypeAbstractSyntax:
		return "Abstract Syntax"
	case ItemTypeTransferSyntax:
		return "Transfer Syntax"
	case ItemTypeUserInformation:
		return "User Information"
	case ItemTypeMaximumLength:
		return "Maximum Length"
	case ItemTypeImplementationClassUID:
		return "Implementation Class UID"
	case ItemTypeAsynchronousOperationsWindow:
		return "Asynchronous Operations Window"
	case ItemTypeSCPSCURoleSelection:
		return "SCP/SCU Role Selection"
	case ItemTypeImplementationVersionName:
		return "Implementation Version Name"
	case ItemTypeExtendedNegotiation:
		return "Extended Negotiation"
	case ItemTypeCommonExtendedNegotiation:
		return "Common Extended Negotiation"
	case ItemTypeUserIdentityNegotiation:
		return "User Identity Negotiation"
	case ItemTypeUserIdentityNegotiationResponse:
		return "User Identity Negotiation Response"
	default:
		return fmt.Sprintf("Unknown(0x%02X)", itemType)
	}
}

// readItem reads an item header (type + reserved + length) and returns the item data.
// Item structure:
//   - Byte 0: Item-type (1 byte)
//   - Byte 1: Reserved (1 byte, always 0x00)
//   - Bytes 2-3: Item-length (2 bytes, big-endian unsigned integer)
//   - Bytes 4+: Item data (variable length)
//
// Note: Item-length does NOT include the 4-byte header.
// Returns io.EOF when there is no more data to read.
func readItem(r io.Reader) (itemType byte, data []byte, err error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(r, header); err != nil {
		// Return EOF only when no header bytes remain. A partial header is malformed.
		if err == io.EOF {
			return 0, nil, io.EOF
		}
		return 0, nil, fmt.Errorf("reading item header: %w", err)
	}

	itemType = header[0]
	// reserved := header[1] (ignored)
	length := binary.BigEndian.Uint16(header[2:4])

	if length > 0 {
		data = make([]byte, length)
		if _, err := io.ReadFull(r, data); err != nil {
			return 0, nil, fmt.Errorf("reading item data (type=0x%02X, length=%d): %w", itemType, length, err)
		}
	}

	return itemType, data, nil
}

// writeItem writes an item header and data to the writer.
func writeItem(w io.Writer, itemType byte, data []byte) error {
	var header [4]byte
	header[0] = itemType
	header[1] = 0x00 // reserved
	if len(data) > int(math.MaxUint16) {
		return fmt.Errorf("item data too large: %d > %d", len(data), math.MaxUint16)
	}
	binary.BigEndian.PutUint16(header[2:4], uint16(len(data))) // #nosec G115 -- bounded by check above

	if err := writeFull(w, header[:]); err != nil {
		return fmt.Errorf("writing item header (type=0x%02X): %w", itemType, err)
	}

	if len(data) > 0 {
		if err := writeFull(w, data); err != nil {
			return fmt.Errorf("writing item data (type=0x%02X): %w", itemType, err)
		}
	}

	return nil
}

// trimUIDPadding removes trailing NUL/space padding that a peer may append to an
// odd-length UID value (PS3.5 UID padding is a single NUL). Decoding must trim it
// so exact string comparisons against the local UID registry keep working.
func trimUIDPadding(s string) string {
	return strings.TrimRight(s, "\x00 ")
}

// writeString writes a string, padding with spaces to 16 bytes (DICOM AE Title length).
func writeString(s string) []byte {
	const length = 16
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		if i < len(s) {
			result[i] = s[i]
		} else {
			result[i] = ' '
		}
	}
	return result
}

// writeAETitleField writes the 16-byte AE Title field of an A-ASSOCIATE-AC.
// PS3.8 Table 9-17 makes this field Reserved and requires it to carry the value
// received in the A-ASSOCIATE-RQ, so a captured 16-byte field is echoed
// verbatim; an absent raw field falls back to the space-padded title.
func writeAETitleField(w io.Writer, title string, raw []byte) {
	if len(raw) == 16 {
		_, _ = w.Write(raw)
		return
	}
	_, _ = w.Write(writeString(title))
}

// validateAETitles validates the Called and Calling AE Titles before encoding.
// PS3.8 Table 9-11 requires 16 characters with leading/trailing spaces
// non-significant and explicitly forbids the all-space value ("no Application
// Name specified").
func validateAETitles(called, calling string) error {
	if err := validateAETitle("called", called); err != nil {
		return err
	}
	return validateAETitle("calling", calling)
}

func validateAETitle(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("invalid %s AE Title: value must not be empty or consist only of spaces", field)
	}
	if err := vr.ValidateAE(value); err != nil {
		return fmt.Errorf("invalid %s AE Title: %w", field, err)
	}
	return nil
}

// ValidateAETitle validates a single AE Title (PS3.8 Table 9-11). The service
// uses it to turn an invalid peer AE Title into an A-ASSOCIATE-RJ instead of
// failing while encoding the A-ASSOCIATE-AC.
func ValidateAETitle(field, value string) error {
	return validateAETitle(field, value)
}

// trimAETitleSpaces removes leading and trailing spaces from an AE Title.
// PS3.8 Table 9-11 makes both leading and trailing spaces non-significant.
func trimAETitleSpaces(s string) string {
	return strings.Trim(s, " ")
}

// encodeUserInformationItems encodes the sub-items of the User Information item
// (0x50). A-ASSOCIATE-RQ and A-ASSOCIATE-AC share sub-items 0x51-0x56; only the
// RQ carries SOP Class Common Extended Negotiation (0x57) and User Identity
// Negotiation (0x58), and only the AC carries User Identity Negotiation
// Response (0x59) (PS3.8 9.3.2.3, 9.3.3.3). Sub-items are written in ascending
// item type order because older applications may expect it (PS3.8 9.3.2.3).
func encodeUserInformationItems(ui *UserInformation, request bool) ([]byte, error) {
	var buf bytes.Buffer
	writers := []func(*bytes.Buffer) error{
		func(w *bytes.Buffer) error { return writeMaximumLengthItem(w, ui.MaximumLength) },
		func(w *bytes.Buffer) error { return writeImplementationClassUIDItem(w, ui.ImplementationClassUID) },
		func(w *bytes.Buffer) error { return writeAsynchronousOperationsItem(w, ui.AsynchronousOperations) },
		func(w *bytes.Buffer) error { return writeRoleSelectionItems(w, ui.SCPSCURoleSelections) },
		func(w *bytes.Buffer) error {
			return writeImplementationVersionNameItem(w, ui.ImplementationVersionName)
		},
		func(w *bytes.Buffer) error { return writeExtendedNegotiationItems(w, ui.ExtendedNegotiations) },
	}
	if request {
		writers = append(writers,
			func(w *bytes.Buffer) error {
				return writeCommonExtendedNegotiationItems(w, ui.CommonExtendedNegotiations)
			},
			func(w *bytes.Buffer) error { return writeUserIdentityItem(w, ui.UserIdentity) },
		)
	} else {
		writers = append(writers,
			func(w *bytes.Buffer) error { return writeUserIdentityResponseItem(w, userIdentityResponse(ui)) },
		)
	}
	for _, writeItems := range writers {
		if err := writeItems(&buf); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// userIdentityResponse returns the AC user identity response, falling back to
// the request structure for callers that represented an AC response with
// UserIdentity before UserIdentityResponse existed.
func userIdentityResponse(ui *UserInformation) *UserIdentityNegotiationResponse {
	if ui.UserIdentityResponse != nil {
		return ui.UserIdentityResponse
	}
	if ui.UserIdentity != nil && ui.UserIdentity.PositiveResponseRequested == 1 {
		return &UserIdentityNegotiationResponse{ServerResponse: ui.UserIdentity.PrimaryField}
	}
	return nil
}

// writeMaximumLengthItem writes the Maximum Length Sub-item (PS3.8 D.1.1).
func writeMaximumLengthItem(w *bytes.Buffer, maximumLength uint32) error {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, advertisedMaximumLength(maximumLength))
	if err := writeItem(w, ItemTypeMaximumLength, data); err != nil {
		return fmt.Errorf("encoding maximum length: %w", err)
	}
	return nil
}

// writeImplementationClassUIDItem writes the Implementation Class UID Sub-item
// (PS3.8 D.1.2).
func writeImplementationClassUIDItem(w *bytes.Buffer, implementationClassUID string) error {
	if implementationClassUID == "" {
		return nil
	}
	if err := writeItem(w, ItemTypeImplementationClassUID, []byte(implementationClassUID)); err != nil {
		return fmt.Errorf("encoding implementation class UID: %w", err)
	}
	return nil
}

// writeAsynchronousOperationsItem writes the Asynchronous Operations Window
// Sub-item (PS3.8 D.3.3.3).
func writeAsynchronousOperationsItem(w *bytes.Buffer, operations *AsynchronousOperationsWindow) error {
	if operations == nil {
		return nil
	}
	data := make([]byte, 4)
	binary.BigEndian.PutUint16(data[0:2], operations.MaximumNumberOperationsInvoked)
	binary.BigEndian.PutUint16(data[2:4], operations.MaximumNumberOperationsPerformed)
	if err := writeItem(w, ItemTypeAsynchronousOperationsWindow, data); err != nil {
		return fmt.Errorf("encoding asynchronous operations window: %w", err)
	}
	return nil
}

// writeRoleSelectionItems writes the SCP/SCU Role Selection Sub-items
// (PS3.7 D.3.3.4).
func writeRoleSelectionItems(w *bytes.Buffer, roles []SCPSCURoleSelection) error {
	for _, role := range roles {
		if len(role.SOPClassUID) > math.MaxUint16 {
			return fmt.Errorf("SOP Class UID length too large: %d", len(role.SOPClassUID))
		}
		var data bytes.Buffer
		_ = binary.Write(&data, binary.BigEndian, uint16(len(role.SOPClassUID))) // #nosec G115 -- bounded above
		data.WriteString(role.SOPClassUID)
		data.WriteByte(role.SCURole)
		data.WriteByte(role.SCPRole)
		if err := writeItem(w, ItemTypeSCPSCURoleSelection, data.Bytes()); err != nil {
			return fmt.Errorf("encoding role selection: %w", err)
		}
	}
	return nil
}

// writeImplementationVersionNameItem writes the Implementation Version Name
// Sub-item (PS3.8 D.1.3).
func writeImplementationVersionNameItem(w *bytes.Buffer, implementationVersionName string) error {
	if implementationVersionName == "" {
		return nil
	}
	if err := writeItem(w, ItemTypeImplementationVersionName, []byte(implementationVersionName)); err != nil {
		return fmt.Errorf("encoding implementation version name: %w", err)
	}
	return nil
}

// writeExtendedNegotiationItems writes the SOP Class Extended Negotiation
// Sub-items (PS3.7 D.3.3.5).
func writeExtendedNegotiationItems(w *bytes.Buffer, negotiations []ExtendedNegotiation) error {
	for _, negotiation := range negotiations {
		if len(negotiation.SOPClassUID) > math.MaxUint16 {
			return fmt.Errorf("extended negotiation SOP Class UID length too large: %d", len(negotiation.SOPClassUID))
		}
		var data bytes.Buffer
		_ = binary.Write(&data, binary.BigEndian, uint16(len(negotiation.SOPClassUID))) // #nosec G115 -- bounded above
		data.WriteString(negotiation.SOPClassUID)
		data.Write(negotiation.ServiceClassAppInfo)
		if err := writeItem(w, ItemTypeExtendedNegotiation, data.Bytes()); err != nil {
			return fmt.Errorf("encoding extended negotiation: %w", err)
		}
	}
	return nil
}

// writeCommonExtendedNegotiationItems writes the SOP Class Common Extended
// Negotiation Sub-items, which are valid only in A-ASSOCIATE-RQ (PS3.8 9.3.2.3.5).
func writeCommonExtendedNegotiationItems(w *bytes.Buffer, negotiations []CommonExtendedNegotiation) error {
	for _, negotiation := range negotiations {
		if err := encodeCommonExtendedNegotiation(w, negotiation); err != nil {
			return fmt.Errorf("encoding common extended negotiation: %w", err)
		}
	}
	return nil
}

// writeUserIdentityItem writes the User Identity Negotiation Sub-item of an
// A-ASSOCIATE-RQ (PS3.7 D.3.3.4).
func writeUserIdentityItem(w *bytes.Buffer, identity *UserIdentityNegotiation) error {
	if identity == nil {
		return nil
	}
	data, err := encodeUserIdentityNegotiation(identity)
	if err != nil {
		return err
	}
	if err := writeItem(w, ItemTypeUserIdentityNegotiation, data); err != nil {
		return fmt.Errorf("encoding user identity: %w", err)
	}
	return nil
}

// writeUserIdentityResponseItem writes the User Identity Negotiation Response
// Sub-item of an A-ASSOCIATE-AC (PS3.8 D.1.7).
func writeUserIdentityResponseItem(w *bytes.Buffer, response *UserIdentityNegotiationResponse) error {
	if response == nil {
		return nil
	}
	data, err := encodeUserIdentityResponse(response)
	if err != nil {
		return err
	}
	if err := writeItem(w, ItemTypeUserIdentityNegotiationResponse, data); err != nil {
		return fmt.Errorf("encoding user identity response: %w", err)
	}
	return nil
}

// encodeUserIdentityNegotiation encodes the User Identity Negotiation Sub-item
// value (PS3.7 D.3.3.4.1).
func encodeUserIdentityNegotiation(identity *UserIdentityNegotiation) ([]byte, error) {
	var data bytes.Buffer
	data.WriteByte(identity.UserIdentityType)
	data.WriteByte(identity.PositiveResponseRequested)
	if len(identity.PrimaryField) > math.MaxUint16 {
		return nil, fmt.Errorf("user identity primary field length too large: %d", len(identity.PrimaryField))
	}
	_ = binary.Write(&data, binary.BigEndian, uint16(len(identity.PrimaryField))) // #nosec G115 -- bounded above
	data.Write(identity.PrimaryField)
	if len(identity.SecondaryField) > math.MaxUint16 {
		return nil, fmt.Errorf("user identity secondary field length too large: %d", len(identity.SecondaryField))
	}
	_ = binary.Write(&data, binary.BigEndian, uint16(len(identity.SecondaryField))) // #nosec G115 -- bounded above
	data.Write(identity.SecondaryField)
	return data.Bytes(), nil
}

// encodeUserIdentityResponse encodes the User Identity Negotiation Response
// Sub-item value: a two-byte length followed by the server response.
func encodeUserIdentityResponse(response *UserIdentityNegotiationResponse) ([]byte, error) {
	if len(response.ServerResponse) > math.MaxUint16 {
		return nil, fmt.Errorf("user identity response length too large: %d", len(response.ServerResponse))
	}
	var data bytes.Buffer
	_ = binary.Write(&data, binary.BigEndian, uint16(len(response.ServerResponse))) // #nosec G115 -- bounded above
	data.Write(response.ServerResponse)
	return data.Bytes(), nil
}

// userInformationDecoder walks the sub-items of a User Information item and
// tracks the duplicates each Sub-Item forbids. request selects the direction
// described in encodeUserInformationItems.
type userInformationDecoder struct {
	ui                            *UserInformation
	request                       bool
	handler                       DecodeWarningHandler
	maximumLengthSeen             bool
	implementationClassUIDSeen    bool
	implementationVersionNameSeen bool
	asynchronousOperationsSeen    bool
	userIdentitySeen              bool
	seenRoleSelections            map[string]struct{}
}

// decodeUserInformationItems decodes the sub-items of the User Information item.
// The RQ decodes into empty (non-nil) slices, matching what callers have always
// seen from a decoded RQ; the AC keeps nil slices.
func decodeUserInformationItems(data []byte, request bool, handler DecodeWarningHandler) (*UserInformation, error) {
	decoder := &userInformationDecoder{
		ui:                 &UserInformation{},
		request:            request,
		handler:            handler,
		seenRoleSelections: make(map[string]struct{}),
	}
	if request {
		decoder.ui.SCPSCURoleSelections = []SCPSCURoleSelection{}
		decoder.ui.ExtendedNegotiations = []ExtendedNegotiation{}
		decoder.ui.CommonExtendedNegotiations = []CommonExtendedNegotiation{}
	}

	reader := bytes.NewReader(data)
	for {
		itemType, itemData, err := readItem(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading user information sub-item: %w", err)
		}
		if err := decoder.decodeSubItem(itemType, itemData); err != nil {
			return nil, err
		}
	}

	decoder.ui.hasMaximumLength = decoder.maximumLengthSeen
	return decoder.ui, nil
}

// decodeSubItem decodes one sub-item. Unknown sub-items only produce a decode
// warning: the AC ignores them silently because its caller passes no handler.
func (d *userInformationDecoder) decodeSubItem(itemType byte, itemData []byte) error {
	switch itemType {
	case ItemTypeMaximumLength:
		return d.decodeMaximumLength(itemData)
	case ItemTypeImplementationClassUID:
		return d.decodeImplementationClassUID(itemData)
	case ItemTypeImplementationVersionName:
		return d.decodeImplementationVersionName(itemData)
	case ItemTypeAsynchronousOperationsWindow:
		return d.decodeAsynchronousOperations(itemData)
	case ItemTypeSCPSCURoleSelection:
		return d.decodeRoleSelectionItem(itemData)
	case ItemTypeExtendedNegotiation:
		return d.decodeExtendedNegotiationItem(itemData)
	case ItemTypeCommonExtendedNegotiation:
		return d.decodeCommonExtendedNegotiationItem(itemData)
	case ItemTypeUserIdentityNegotiation:
		return d.decodeUserIdentityItem(itemData)
	case ItemTypeUserIdentityNegotiationResponse:
		return d.decodeUserIdentityResponse(itemData)
	default:
		reportDecodeWarning(d.handler, DecodeWarningUnknownUserInformationSubItem, itemType)
		return nil
	}
}

func (d *userInformationDecoder) decodeMaximumLength(data []byte) error {
	// PS3.8 D.1.1: only one Maximum Length Sub-Item shall be present.
	if d.maximumLengthSeen {
		return fmt.Errorf("duplicate maximum length sub-item in user information")
	}
	if len(data) != 4 {
		return fmt.Errorf("invalid maximum length data: %d bytes", len(data))
	}
	d.maximumLengthSeen = true
	d.ui.MaximumLength = binary.BigEndian.Uint32(data)
	return nil
}

func (d *userInformationDecoder) decodeImplementationClassUID(data []byte) error {
	// PS3.8 D.1.2: only one Implementation Class UID Sub-Item shall be present.
	if d.implementationClassUIDSeen {
		return fmt.Errorf("duplicate implementation class UID sub-item in user information")
	}
	d.implementationClassUIDSeen = true
	d.ui.ImplementationClassUID = trimUIDPadding(string(data))
	return nil
}

func (d *userInformationDecoder) decodeImplementationVersionName(data []byte) error {
	// PS3.8 D.1.3: only one Implementation Version Name Sub-Item shall be present.
	if d.implementationVersionNameSeen {
		return fmt.Errorf("duplicate implementation version name sub-item in user information")
	}
	d.implementationVersionNameSeen = true
	d.ui.ImplementationVersionName = string(data)
	return nil
}

func (d *userInformationDecoder) decodeAsynchronousOperations(data []byte) error {
	// PS3.8 D.1.4: only one Asynchronous Operations Window Sub-Item shall be present.
	if d.asynchronousOperationsSeen {
		return fmt.Errorf("duplicate asynchronous operations window sub-item in user information")
	}
	d.asynchronousOperationsSeen = true
	if len(data) != 4 {
		return fmt.Errorf("invalid asynchronous operations window: %d bytes", len(data))
	}
	d.ui.AsynchronousOperations = &AsynchronousOperationsWindow{
		MaximumNumberOperationsInvoked:   binary.BigEndian.Uint16(data[0:2]),
		MaximumNumberOperationsPerformed: binary.BigEndian.Uint16(data[2:4]),
	}
	return nil
}

func (d *userInformationDecoder) decodeRoleSelectionItem(data []byte) error {
	role, err := decodeRoleSelection(data)
	if err != nil {
		return fmt.Errorf("decoding role selection: %w", err)
	}
	// PS3.7 D.3.3.4: at most one SCP/SCU Role Selection Sub-Item per SOP Class UID.
	if _, duplicate := d.seenRoleSelections[role.SOPClassUID]; duplicate {
		return fmt.Errorf("duplicate SCP/SCU role selection sub-item for SOP Class %s", role.SOPClassUID)
	}
	d.seenRoleSelections[role.SOPClassUID] = struct{}{}
	d.ui.SCPSCURoleSelections = append(d.ui.SCPSCURoleSelections, *role)
	return nil
}

func (d *userInformationDecoder) decodeExtendedNegotiationItem(data []byte) error {
	negotiation, err := decodeExtendedNegotiation(data)
	if err != nil {
		return fmt.Errorf("decoding extended negotiation: %w", err)
	}
	d.ui.ExtendedNegotiations = append(d.ui.ExtendedNegotiations, *negotiation)
	return nil
}

func (d *userInformationDecoder) decodeCommonExtendedNegotiationItem(data []byte) error {
	// The AC has no SOP Class Common Extended Negotiation Sub-item.
	if !d.request {
		reportDecodeWarning(d.handler, DecodeWarningUnknownUserInformationSubItem, ItemTypeCommonExtendedNegotiation)
		return nil
	}
	negotiation, err := decodeCommonExtendedNegotiation(data)
	if err != nil {
		return fmt.Errorf("decoding common extended negotiation: %w", err)
	}
	d.ui.CommonExtendedNegotiations = append(d.ui.CommonExtendedNegotiations, *negotiation)
	return nil
}

func (d *userInformationDecoder) decodeUserIdentityItem(data []byte) error {
	// A-ASSOCIATE-AC answers user identity with the response sub-item only.
	if !d.request {
		reportDecodeWarning(d.handler, DecodeWarningUnknownUserInformationSubItem, ItemTypeUserIdentityNegotiation)
		return nil
	}
	// PS3.7 D.3.3.4: at most one User Identity Sub-Item.
	if d.userIdentitySeen {
		return fmt.Errorf("duplicate user identity sub-item in user information")
	}
	d.userIdentitySeen = true
	identity, err := decodeUserIdentity(data)
	if err != nil {
		return fmt.Errorf("decoding user identity: %w", err)
	}
	d.ui.UserIdentity = identity
	return nil
}

func (d *userInformationDecoder) decodeUserIdentityResponse(data []byte) error {
	// The RQ has no User Identity Negotiation Response Sub-item.
	if d.request {
		reportDecodeWarning(d.handler, DecodeWarningUnknownUserInformationSubItem, ItemTypeUserIdentityNegotiationResponse)
		return nil
	}
	// PS3.8 D.1.7: only one User Identity Sub-Item.
	if d.userIdentitySeen {
		return fmt.Errorf("duplicate user identity sub-item in user information")
	}
	d.userIdentitySeen = true
	if len(data) < 2 {
		return fmt.Errorf("invalid user identity response: %d bytes", len(data))
	}
	serverResponseLength := binary.BigEndian.Uint16(data[0:2])
	if len(data) != int(2+serverResponseLength) {
		return fmt.Errorf("invalid user identity response length")
	}
	serverResponse := data[2:]
	d.ui.UserIdentityResponse = &UserIdentityNegotiationResponse{ServerResponse: serverResponse}
	// Preserve the legacy decoded representation.
	d.ui.UserIdentity = &UserIdentityNegotiation{
		PositiveResponseRequested: 1, // A response was provided
		PrimaryField:              serverResponse,
	}
	return nil
}
