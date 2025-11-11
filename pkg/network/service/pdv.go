// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
    "encoding/binary"
    "fmt"
    "math"

    "github.com/cocosip/go-dicom/pkg/network/pdu"
)

// PDV (Presentation Data Value) represents a DICOM data fragment.
// PDVs are used to fragment DIMSE messages into smaller chunks that fit within
// the maximum PDU length negotiated during association.
//
// PDV structure:
//   - Length (4 bytes): Length of PDV data (including presentation context ID and header)
//   - Presentation Context ID (1 byte)
//   - Message Control Header (1 byte):
//   - Bit 0: 0=More fragments follow, 1=Last fragment
//   - Bit 1: 0=Data set, 1=Command
//   - PDV Data (variable length)
type PDV struct {
	PresentationContextID byte
	IsCommand             bool   // true if this is a command PDV, false if dataset
	IsLastFragment        bool   // true if this is the last fragment
	Data                  []byte // PDV data
}

// Encode encodes the PDV into bytes for transmission.
func (p *PDV) Encode() []byte {
    // Calculate total length (1 byte context ID + 1 byte header + data)
    total := 1 + 1 + len(p.Data)
    if total > int(math.MaxUint32) {
        // defensive: extremely unlikely in tests/typical usage
        total = int(math.MaxUint32)
    }
    length := uint32(total) // #nosec G115 -- bounded above

	// Create buffer
	buf := make([]byte, 4+length)

	// Write length
	binary.BigEndian.PutUint32(buf[0:4], length)

	// Write presentation context ID
	buf[4] = p.PresentationContextID

	// Write message control header
	var header byte
	if p.IsLastFragment {
		header |= 0x02 // Set bit 1
	}
	if p.IsCommand {
		header |= 0x01 // Set bit 0
	}
	buf[5] = header

	// Write data
	copy(buf[6:], p.Data)

	return buf
}

// DecodePDV decodes a PDV from bytes.
func DecodePDV(data []byte) (*PDV, error) {
    if len(data) < 6 {
        return nil, fmt.Errorf("PDV data too short: %d bytes", len(data))
    }

    // Read length
    length := binary.BigEndian.Uint32(data[0:4])
    // Safe conversion: compare using int after bounds check
    if len(data) < int(4+length) {
        return nil, fmt.Errorf("PDV data incomplete: expected %d bytes, got %d", 4+length, len(data))
    }

	// Read presentation context ID
	contextID := data[4]

	// Read message control header
	header := data[5]
	isLastFragment := (header & 0x02) != 0
	isCommand := (header & 0x01) != 0

	// Read PDV data
	pdvData := make([]byte, length-2)
	copy(pdvData, data[6:6+length-2])

	return &PDV{
		PresentationContextID: contextID,
		IsCommand:             isCommand,
		IsLastFragment:        isLastFragment,
		Data:                  pdvData,
	}, nil
}

// CreatePDataTFPDU creates a P-DATA-TF PDU from a list of PDVs.
func CreatePDataTFPDU(pdvs []*PDV) (*pdu.RawPDU, error) {
	if len(pdvs) == 0 {
		return nil, fmt.Errorf("no PDVs to encode")
	}

	// Calculate total size
	totalSize := 0
	for _, pdv := range pdvs {
		totalSize += len(pdv.Encode())
	}

	// Create buffer
	buf := make([]byte, totalSize)

	// Encode all PDVs
	offset := 0
	for _, pdv := range pdvs {
		encoded := pdv.Encode()
		copy(buf[offset:], encoded)
		offset += len(encoded)
	}

	return &pdu.RawPDU{
		Type: pdu.TypePDataTF,
		Data: buf,
	}, nil
}

// ParsePDataTFPDU parses PDVs from a P-DATA-TF PDU.
func ParsePDataTFPDU(p *pdu.RawPDU) ([]*PDV, error) {
	if p.Type != pdu.TypePDataTF {
		return nil, fmt.Errorf("not a P-DATA-TF PDU: type=%d", p.Type)
	}

	var pdvs []*PDV
	data := p.Data
	offset := 0

	for offset < len(data) {
		if offset+4 > len(data) {
			return nil, fmt.Errorf("incomplete PDV at offset %d", offset)
		}

		// Read PDV length
		pdvLength := binary.BigEndian.Uint32(data[offset : offset+4])
		if offset+4+int(pdvLength) > len(data) {
			return nil, fmt.Errorf("PDV extends beyond buffer: offset=%d, length=%d, total=%d",
				offset, pdvLength, len(data))
		}

		// Decode PDV
		pdvData := data[offset : offset+4+int(pdvLength)]
		pdv, err := DecodePDV(pdvData)
		if err != nil {
			return nil, fmt.Errorf("failed to decode PDV at offset %d: %w", offset, err)
		}

		pdvs = append(pdvs, pdv)
		offset += 4 + int(pdvLength)
	}

	return pdvs, nil
}

// FragmentData fragments data into multiple PDVs based on max PDU length.
// Each PDV will be at most maxPDVSize bytes (accounting for PDV header overhead).
func FragmentData(data []byte, contextID byte, isCommand bool, maxPDULength uint32) []*PDV {
	if len(data) == 0 {
		// Return single empty PDV
		return []*PDV{
			{
				PresentationContextID: contextID,
				IsCommand:             isCommand,
				IsLastFragment:        true,
				Data:                  []byte{},
			},
		}
	}

	// Calculate max PDV data size
	// PDU header: 6 bytes (type + reserved + length)
	// PDV header per fragment: 6 bytes (length + context ID + control header)
	// We need to leave room for at least one PDV header
	const pduHeaderSize = 6
	const pdvHeaderSize = 6

	maxPDVDataSize := int(maxPDULength) - pduHeaderSize - pdvHeaderSize

	if maxPDVDataSize <= 0 {
		maxPDVDataSize = 1024 // Fallback to 1KB minimum
	}

	var pdvs []*PDV
	offset := 0

	for offset < len(data) {
		// Calculate chunk size
		chunkSize := maxPDVDataSize
		if offset+chunkSize > len(data) {
			chunkSize = len(data) - offset
		}

		// Create PDV
		pdv := &PDV{
			PresentationContextID: contextID,
			IsCommand:             isCommand,
			IsLastFragment:        offset+chunkSize >= len(data),
			Data:                  data[offset : offset+chunkSize],
		}

		pdvs = append(pdvs, pdv)
		offset += chunkSize
	}

	return pdvs
}
