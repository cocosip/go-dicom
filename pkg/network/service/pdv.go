// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// PDV (Presentation Data Value) represents a DICOM data fragment.
// PDVs are used to fragment DIMSE messages into smaller chunks that fit within
// the maximum PDU length negotiated during association.
//
// PDV structure:
//   - Length (4 bytes): Length of PDV data (including presentation context ID and header)
//   - Presentation Context ID (1 byte)
//   - Message Control Header (1 byte), per PS3.8 E.2:
//   - Bit 0: 1=Command, 0=Data set
//   - Bit 1: 1=Last fragment, 0=More fragments follow
//   - PDV Data (variable length)
type PDV struct {
	PresentationContextID byte
	IsCommand             bool   // true if this is a command PDV, false if dataset
	IsLastFragment        bool   // true if this is the last fragment
	Data                  []byte // PDV data
}

// encodedSize returns the number of bytes the PDV occupies on the wire.
func (p *PDV) encodedSize() int { return 6 + len(p.Data) }

// encodeInto writes the PDV into dst and returns the number of bytes written.
// dst must be at least encodedSize() bytes long.
func (p *PDV) encodeInto(dst []byte) int {
	// PDV length covers context ID + control header + data. The fragment size is
	// bounded by the negotiated maximum PDU length, so it always fits in uint32.
	total := 1 + 1 + len(p.Data)
	binary.BigEndian.PutUint32(dst[0:4], uint32(total)) // #nosec G115 -- bounded by the negotiated PDU length

	dst[4] = p.PresentationContextID

	var header byte
	if p.IsLastFragment {
		header |= 0x02 // Set bit 1
	}
	if p.IsCommand {
		header |= 0x01 // Set bit 0
	}
	dst[5] = header

	copy(dst[6:], p.Data)
	return 4 + total
}

// Encode encodes the PDV into bytes for transmission.
func (p *PDV) Encode() []byte {
	buf := make([]byte, p.encodedSize())
	p.encodeInto(buf)
	return buf
}

// DecodePDV decodes a PDV from bytes. The returned Data aliases data; callers
// must not retain it after data is released.
func DecodePDV(data []byte) (*PDV, error) {
	if len(data) < 6 {
		return nil, fmt.Errorf("PDV data too short: %d bytes", len(data))
	}

	// Read length
	length := binary.BigEndian.Uint32(data[0:4])
	if length < 2 {
		return nil, fmt.Errorf("invalid PDV length: %d (minimum 2 bytes)", length)
	}
	expectedLength := uint64(4) + uint64(length)
	if uint64(len(data)) < expectedLength {
		return nil, fmt.Errorf("PDV data incomplete: expected %d bytes, got %d", expectedLength, len(data))
	}

	// Read presentation context ID
	contextID := data[4]

	// Read message control header
	header := data[5]
	isLastFragment := (header & 0x02) != 0
	isCommand := (header & 0x01) != 0

	// Read PDV data (no copy: the caller consumes it before the PDU buffer is
	// released).
	pdvData := data[6 : 6+length-2]

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

	// Calculate total size once; then encode each PDV exactly once into the
	// final buffer (previously each PDV was encoded twice).
	totalSize := 0
	for _, pdv := range pdvs {
		totalSize += pdv.encodedSize()
	}

	// Create buffer
	buf := make([]byte, totalSize)

	// Encode all PDVs
	offset := 0
	for _, pdv := range pdvs {
		offset += pdv.encodeInto(buf[offset:])
	}

	return &pdu.RawPDU{
		Type: pdu.TypePDataTF,
		Data: buf,
	}, nil
}

// ParsePDataTFPDU parses PDVs from a P-DATA-TF PDU.
func ParsePDataTFPDU(p *pdu.RawPDU) ([]*PDV, error) {
	if p == nil {
		return nil, fmt.Errorf("P-DATA-TF PDU is nil")
	}
	if p.Type != pdu.TypePDataTF {
		return nil, fmt.Errorf("not a P-DATA-TF PDU: type=%d", p.Type)
	}
	if len(p.Data) == 0 {
		return nil, fmt.Errorf("P-DATA-TF PDU contains no PDV")
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
		pdvEnd := uint64(offset) + 4 + uint64(pdvLength)
		if pdvEnd > uint64(len(data)) {
			return nil, fmt.Errorf("PDV extends beyond buffer: offset=%d, length=%d, total=%d",
				offset, pdvLength, len(data))
		}

		// Decode PDV
		pdvData := data[offset:pdvEnd]
		pdv, err := DecodePDV(pdvData)
		if err != nil {
			return nil, fmt.Errorf("failed to decode PDV at offset %d: %w", offset, err)
		}

		pdvs = append(pdvs, pdv)
		offset = int(pdvEnd)
	}

	return pdvs, nil
}

// FragmentData fragments data into multiple PDVs based on max PDU length.
// Each PDV will be at most maxPDVSize bytes (accounting for PDV header overhead).
func FragmentData(data []byte, contextID byte, isCommand bool, maxPDULength uint32) []*PDV {
	if len(data) == 0 {
		// PS3.8 E.1 forbids empty PDVs; the caller must not send a dataset PDV
		// at all in this case.
		return nil
	}

	// Calculate max PDV data size
	// PDV header per fragment: 6 bytes (length + context ID + control header)
	const pdvHeaderSize = 6

	// MaximumLengthReceived is the maximum variable-field length and excludes
	// the six-byte PDU header (PS3.8). Only the PDV header consumes this budget.
	maxPDVDataSize := int(maxPDULength) - pdvHeaderSize

	// PS3.8 E.1: every fragment shall consist of an even number of bytes.
	maxPDVDataSize &^= 1

	if maxPDVDataSize <= 0 {
		if maxPDULength == 0 {
			// No max PDU length configured — return all data in one PDV.
			return []*PDV{
				{
					PresentationContextID: contextID,
					IsCommand:             isCommand,
					IsLastFragment:        true,
					Data:                  data,
				},
			}
		}
		// maxPDULength is positive but too small to hold even a one-byte PDV.
		return nil
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
