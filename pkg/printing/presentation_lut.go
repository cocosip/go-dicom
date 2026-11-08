// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

// PresentationLUTShape represents the shape of the Presentation LUT
type PresentationLUTShape string

const (
	// PresentationLUTShapeIdentity - IDENTITY shape
	PresentationLUTShapeIdentity PresentationLUTShape = "IDENTITY"
	// PresentationLUTShapeLinOD - LIN OD (Linear Optical Density) shape
	PresentationLUTShapeLinOD PresentationLUTShape = "LIN OD"
	// PresentationLUTShapeInverse - INVERSE shape
	PresentationLUTShapeInverse PresentationLUTShape = "INVERSE"
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

	// BitsStored is the image Bits Stored this LUT applies to when known. A
	// print Presentation LUT maps 8-bit images with 256 entries and 12-bit
	// images with 4096 entries (PS3.4 H.4.9.2.1.1.1); zero leaves it unspecified.
	BitsStored uint16

	// PresentationLUTShape specifies the shape of the Presentation LUT.
	// Print presentation LUTs accept only IDENTITY or LIN OD (PS3.3 C.11.4,
	// Table C.11-4); INVERSE is a softcopy (C.11.6) value.
	PresentationLUTShape PresentationLUTShape
}

// NewPresentationLUT creates a new Presentation LUT instance
func NewPresentationLUT(sopInstanceUID string) *PresentationLUT {
	if sopInstanceUID == "" {
		sopInstanceUID = newSOPInstanceUID()
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
	p.PresentationLUTShape = ""
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
	hasShape := p.PresentationLUTShape != ""
	hasLUT := len(p.LUTDescriptor) != 0 || len(p.LUTData) != 0
	if hasShape {
		// PS3.3 C.11.4 Table C.11-4: print Presentation LUT Shape is
		// IDENTITY or LIN OD only. INVERSE is not permitted here.
		return !hasLUT && (p.PresentationLUTShape == PresentationLUTShapeIdentity ||
			p.PresentationLUTShape == PresentationLUTShapeLinOD)
	}
	if !hasLUT {
		return false
	}

	if len(p.LUTDescriptor) != 3 {
		return false
	}

	numberOfEntries := p.GetNumberOfEntries()
	bitsPerEntry := p.GetBitsPerEntry()

	// Bits per entry must be between 10 and 16 inclusive
	if bitsPerEntry < 10 || bitsPerEntry > 16 {
		return false
	}

	// PS3.4 H.4.9.2.1.1.1: the first LUT Descriptor value is 256 for 8-bit
	// images and 4096 for 12-bit images; no other count is permitted for print.
	switch numberOfEntries {
	case 256:
		if p.BitsStored != 0 && p.BitsStored != 8 {
			return false
		}
	case 4096:
		if p.BitsStored != 0 && p.BitsStored != 12 {
			return false
		}
	default:
		return false
	}
	if len(p.LUTData) != int(numberOfEntries) {
		return false
	}
	maxValue := uint16((uint32(1) << bitsPerEntry) - 1)
	for _, value := range p.LUTData {
		if value > maxValue {
			return false
		}
	}

	// First value mapped should be 0
	if p.GetFirstValueMapped() != 0 {
		return false
	}

	return true
}

// TransformValue applies the LUT transformation to an input value
func (p *PresentationLUT) TransformValue(inputValue uint16) uint16 {
	if len(p.LUTData) == 0 || len(p.LUTDescriptor) != 3 {
		if p.PresentationLUTShape == PresentationLUTShapeInverse {
			return ^inputValue
		}
		return inputValue
	}

	firstValue := p.GetFirstValueMapped()

	// Calculate index
	index := int(inputValue) - int(firstValue)

	// Clamp to LUT range
	if index < 0 {
		return p.LUTData[0]
	}
	if index >= len(p.LUTData) {
		return p.LUTData[len(p.LUTData)-1]
	}

	return p.LUTData[index]
}
