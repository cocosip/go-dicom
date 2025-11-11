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

// PDataTF represents a P-DATA-TF (Presentation Data Transfer) PDU.
// This PDU carries DIMSE message data.
//
// PDU Structure (variable length):
//   - Bytes 0-5: PDU header (Type=0x04, Reserved, Length)
//   - Bytes 6+: One or more Presentation-data-value items (PDVs)
//
// Each PDV has the structure:
//   - Bytes 0-3: PDV length (4 bytes, big-endian, includes header byte)
//   - Byte 4: Presentation Context ID
//   - Byte 5: Message Control Header (bit flags)
//   - Bytes 6+: PDV data
type PDataTF struct {
	// PDVs contains one or more Presentation Data Values
	PDVs []PDV
}

// PDV represents a Presentation Data Value.
// A PDV is a fragment of a DIMSE message.
type PDV struct {
	// PresentationContextID identifies which presentation context this data belongs to
	PresentationContextID byte

	// IsCommand indicates whether this PDV contains command data (true) or dataset data (false)
	// Bit 0 of Message Control Header
	IsCommand bool

	// IsLastFragment indicates whether this is the last fragment of the message
	// Bit 1 of Message Control Header
	IsLastFragment bool

	// Data contains the actual DIMSE message fragment
	Data []byte
}

// Message Control Header bit masks
const (
	MessageControlHeaderCommand      byte = 0x01 // Bit 0: 1=command, 0=data
	MessageControlHeaderLastFragment byte = 0x02 // Bit 1: 1=last fragment, 0=not last
)

// NewPDataTF creates a new P-DATA-TF PDU.
func NewPDataTF() *PDataTF {
	return &PDataTF{
		PDVs: make([]PDV, 0),
	}
}

// AddPDV adds a PDV to the P-DATA-TF PDU.
func (p *PDataTF) AddPDV(pdv PDV) {
	p.PDVs = append(p.PDVs, pdv)
}

// NewPDV creates a new Presentation Data Value.
func NewPDV(presentationContextID byte, isCommand, isLastFragment bool, data []byte) PDV {
	return PDV{
		PresentationContextID: presentationContextID,
		IsCommand:             isCommand,
		IsLastFragment:        isLastFragment,
		Data:                  data,
	}
}

// Encode encodes the P-DATA-TF to a RawPDU.
func (p *PDataTF) Encode() (*RawPDU, error) {
	if len(p.PDVs) == 0 {
		return nil, fmt.Errorf("P-DATA-TF must contain at least one PDV")
	}

	var buf bytes.Buffer

	// Encode each PDV
	for i, pdv := range p.PDVs {
		// PDV Length (4 bytes) - includes Presentation Context ID + Message Control Header + Data
        pdvLen := 1 + 1 + len(pdv.Data)
        if pdvLen > int(math.MaxUint32) {
            return nil, fmt.Errorf("PDV length too large: %d", pdvLen)
        }
        // #nosec G115 -- bounded by check above
        pdvLength := uint32(pdvLen) // Context ID (1) + Header (1) + Data
		if err := binary.Write(&buf, binary.BigEndian, pdvLength); err != nil {
			return nil, fmt.Errorf("encoding PDV %d length: %w", i, err)
		}

		// Presentation Context ID (1 byte)
		buf.WriteByte(pdv.PresentationContextID)

		// Message Control Header (1 byte)
		header := byte(0x00)
		if pdv.IsCommand {
			header |= MessageControlHeaderCommand
		}
		if pdv.IsLastFragment {
			header |= MessageControlHeaderLastFragment
		}
		buf.WriteByte(header)

		// PDV Data
		if len(pdv.Data) > 0 {
			buf.Write(pdv.Data)
		}
	}

	return NewRawPDU(TypePDataTF, buf.Bytes()), nil
}

// Decode decodes a P-DATA-TF from a RawPDU.
func (p *PDataTF) Decode(pdu *RawPDU) error {
	if pdu.Type != TypePDataTF {
		return fmt.Errorf("invalid PDU type: expected P-DATA-TF (0x04), got 0x%02X", pdu.Type)
	}

	data := pdu.Data
	reader := bytes.NewReader(data)

	p.PDVs = make([]PDV, 0)

	// Read PDVs until end of data
	for {
		// Read PDV Length (4 bytes)
		var pdvLength uint32
		if err := binary.Read(reader, binary.BigEndian, &pdvLength); err != nil {
			if err == io.EOF {
				break // End of PDVs
			}
			return fmt.Errorf("reading PDV length: %w", err)
		}

		if pdvLength < 2 {
			return fmt.Errorf("invalid PDV length: %d (minimum 2 bytes)", pdvLength)
		}

		// Read Presentation Context ID (1 byte)
		contextID, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading presentation context ID: %w", err)
		}

		// Read Message Control Header (1 byte)
		header, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading message control header: %w", err)
		}

		isCommand := (header & MessageControlHeaderCommand) != 0
		isLastFragment := (header & MessageControlHeaderLastFragment) != 0

		// Read PDV Data (remaining bytes)
		dataLength := pdvLength - 2 // Subtract Context ID (1) + Header (1)
		pdvData := make([]byte, dataLength)
		if dataLength > 0 {
			if _, err := io.ReadFull(reader, pdvData); err != nil {
				return fmt.Errorf("reading PDV data: %w", err)
			}
		}

		p.PDVs = append(p.PDVs, PDV{
			PresentationContextID: contextID,
			IsCommand:             isCommand,
			IsLastFragment:        isLastFragment,
			Data:                  pdvData,
		})
	}

	if len(p.PDVs) == 0 {
		return fmt.Errorf("P-DATA-TF contains no PDVs")
	}

	return nil
}

// String returns a human-readable representation.
func (p *PDataTF) String() string {
	return fmt.Sprintf("P-DATA-TF (%d PDVs)", len(p.PDVs))
}

// TotalDataSize returns the total size of all PDV data.
func (p *PDataTF) TotalDataSize() int {
	total := 0
	for _, pdv := range p.PDVs {
		total += len(pdv.Data)
	}
	return total
}

// GetCommandPDVs returns all PDVs that contain command data.
func (p *PDataTF) GetCommandPDVs() []PDV {
	var result []PDV
	for _, pdv := range p.PDVs {
		if pdv.IsCommand {
			result = append(result, pdv)
		}
	}
	return result
}

// GetDataPDVs returns all PDVs that contain dataset data.
func (p *PDataTF) GetDataPDVs() []PDV {
	var result []PDV
	for _, pdv := range p.PDVs {
		if !pdv.IsCommand {
			result = append(result, pdv)
		}
	}
	return result
}

// FragmentData splits data into multiple PDVs based on maxPDVSize.
// This is used when a DIMSE message is larger than the maximum PDU length.
//
// maxPDVSize is the maximum size of data in a single PDV (excluding PDV header).
// The actual PDU size will be: 6 (PDU header) + 4 (PDV length) + 2 (Context ID + Header) + maxPDVSize
//
// For a maximum PDU length of 16384 bytes, maxPDVSize should be 16384 - 6 - 6 = 16372 bytes.
func FragmentData(presentationContextID byte, isCommand bool, data []byte, maxPDVSize uint32) []PDV {
	if maxPDVSize == 0 {
		// No fragmentation, return single PDV
		return []PDV{
			NewPDV(presentationContextID, isCommand, true, data),
		}
	}

	var pdvs []PDV
	remaining := data
	offset := 0

	for len(remaining) > 0 {
		// Determine fragment size
		fragmentSize := int(maxPDVSize)
		if fragmentSize > len(remaining) {
			fragmentSize = len(remaining)
		}

		// Extract fragment
		fragment := remaining[:fragmentSize]
		remaining = remaining[fragmentSize:]

		// Create PDV
		isLast := len(remaining) == 0
		pdv := NewPDV(presentationContextID, isCommand, isLast, fragment)
		pdvs = append(pdvs, pdv)

		offset += fragmentSize
	}

	return pdvs
}

// ReassembleData combines multiple PDVs into a single data buffer.
// Returns an error if PDVs are incomplete (no last fragment) or from different contexts.
func ReassembleData(pdvs []PDV) ([]byte, error) {
	if len(pdvs) == 0 {
		return nil, fmt.Errorf("no PDVs to reassemble")
	}

	// Verify all PDVs are from the same context and type
	firstContext := pdvs[0].PresentationContextID
	firstType := pdvs[0].IsCommand

	var buf bytes.Buffer
	hasLastFragment := false

	for i, pdv := range pdvs {
		if pdv.PresentationContextID != firstContext {
			return nil, fmt.Errorf("PDV %d has different context ID: expected %d, got %d",
				i, firstContext, pdv.PresentationContextID)
		}
		if pdv.IsCommand != firstType {
			return nil, fmt.Errorf("PDV %d has different type: expected command=%v, got command=%v",
				i, firstType, pdv.IsCommand)
		}

		buf.Write(pdv.Data)

		if pdv.IsLastFragment {
			hasLastFragment = true
			// Verify this is the last PDV
			if i != len(pdvs)-1 {
				return nil, fmt.Errorf("last fragment flag set on PDV %d, but not the last PDV", i)
			}
		}
	}

	if !hasLastFragment {
		return nil, fmt.Errorf("incomplete message: no last fragment")
	}

	return buf.Bytes(), nil
}
