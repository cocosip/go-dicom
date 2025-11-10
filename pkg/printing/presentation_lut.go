// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

// PresentationLUTShape represents the shape of the Presentation LUT
type PresentationLUTShape string

const (
	// PresentationLUTShapeIdentity - IDENTITY shape
	PresentationLUTShapeIdentity PresentationLUTShape = "IDENTITY"
	// PresentationLUTShapeLinOD - LIN OD (Linear Optical Density) shape
	PresentationLUTShapeLinOD PresentationLUTShape = "LIN OD"
)

// PresentationLUT represents a Presentation LUT Information Object
// Used for DICOM Print Management to define how pixel values are
// transformed for display on film
//
// Reference: DICOM Part 3, Section B.18
// http://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_B.18.html
type PresentationLUT struct {
	// SOPInstanceUID is the SOP Instance UID of the Presentation LUT object
	SOPInstanceUID string

	// LUTDescriptor contains three values:
	// - Number of entries in the lookup table
	// - First input value mapped (always 0)
	// - Number of bits for each entry in LUT Data (10-16 inclusive)
	LUTDescriptor []uint16

	// LUTExplanation is a free-form text explanation of the meaning of the LUT
	LUTExplanation string

	// LUTData contains the LUT entry values (P-Values)
	LUTData []uint16

	// PresentationLUTShape specifies the shape of the Presentation LUT
	// Enumerated values: 'IDENTITY' or 'LIN OD'
	PresentationLUTShape PresentationLUTShape
}

// NewPresentationLUT creates a new Presentation LUT instance
func NewPresentationLUT(sopInstanceUID string) *PresentationLUT {
	if sopInstanceUID == "" {
		// Generate a UID if not provided
		// In a real implementation, use proper UID generation
		sopInstanceUID = "1.2.840.10008.5.1.4.1.1.11.1" // Placeholder
	}

	return &PresentationLUT{
		SOPInstanceUID:       sopInstanceUID,
		LUTDescriptor:        []uint16{},
		LUTExplanation:       "",
		LUTData:              []uint16{},
		PresentationLUTShape: PresentationLUTShapeIdentity,
	}
}

// SetLUT sets the LUT descriptor and data
func (p *PresentationLUT) SetLUT(numberOfEntries, firstValueMapped, bitsPerEntry uint16, lutData []uint16) error {
	p.LUTDescriptor = []uint16{numberOfEntries, firstValueMapped, bitsPerEntry}
	p.LUTData = make([]uint16, len(lutData))
	copy(p.LUTData, lutData)
	return nil
}

// GetNumberOfEntries returns the number of entries in the LUT
func (p *PresentationLUT) GetNumberOfEntries() uint16 {
	if len(p.LUTDescriptor) >= 1 {
		return p.LUTDescriptor[0]
	}
	return 0
}

// GetFirstValueMapped returns the first input value mapped
func (p *PresentationLUT) GetFirstValueMapped() uint16 {
	if len(p.LUTDescriptor) >= 2 {
		return p.LUTDescriptor[1]
	}
	return 0
}

// GetBitsPerEntry returns the number of bits for each LUT entry
func (p *PresentationLUT) GetBitsPerEntry() uint16 {
	if len(p.LUTDescriptor) >= 3 {
		return p.LUTDescriptor[2]
	}
	return 0
}

// IsValid checks if the Presentation LUT configuration is valid
func (p *PresentationLUT) IsValid() bool {
	// Check descriptor
	if len(p.LUTDescriptor) != 3 {
		return false
	}

	numberOfEntries := p.GetNumberOfEntries()
	bitsPerEntry := p.GetBitsPerEntry()

	// Bits per entry must be between 10 and 16 inclusive
	if bitsPerEntry < 10 || bitsPerEntry > 16 {
		return false
	}

	// LUT data size should match number of entries
	if len(p.LUTData) != int(numberOfEntries) {
		return false
	}

	// First value mapped should be 0
	if p.GetFirstValueMapped() != 0 {
		return false
	}

	return true
}

// TransformValue applies the LUT transformation to an input value
func (p *PresentationLUT) TransformValue(inputValue uint16) uint16 {
	if p.PresentationLUTShape == PresentationLUTShapeIdentity {
		// IDENTITY: output = input
		return inputValue
	}

	// Apply LUT transformation
	if len(p.LUTData) == 0 {
		return inputValue
	}

	numberOfEntries := p.GetNumberOfEntries()
	firstValue := p.GetFirstValueMapped()

	// Calculate index
	index := int(inputValue) - int(firstValue)

	// Clamp to LUT range
	if index < 0 {
		return p.LUTData[0]
	}
	if index >= int(numberOfEntries) {
		return p.LUTData[numberOfEntries-1]
	}

	return p.LUTData[index]
}
