// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
    "encoding/binary"
    "fmt"
    "io"
    "math"
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
		// Return EOF directly without wrapping if we're at the end of the stream
		if err == io.EOF || err == io.ErrUnexpectedEOF {
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
    header := make([]byte, 4)
    header[0] = itemType
    header[1] = 0x00 // reserved
    if len(data) > int(math.MaxUint16) {
        return fmt.Errorf("item data too large: %d > %d", len(data), math.MaxUint16)
    }
    binary.BigEndian.PutUint16(header[2:4], uint16(len(data))) // #nosec G115 -- bounded by check above

	if _, err := w.Write(header); err != nil {
		return fmt.Errorf("writing item header (type=0x%02X): %w", itemType, err)
	}

	if len(data) > 0 {
		if _, err := w.Write(data); err != nil {
			return fmt.Errorf("writing item data (type=0x%02X): %w", itemType, err)
		}
	}

	return nil
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

// trimTrailingSpaces removes trailing spaces from a string.
func trimTrailingSpaces(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}
