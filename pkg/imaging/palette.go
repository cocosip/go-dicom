// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
)

// Color32 represents a 32-bit RGBA color
type Color32 struct {
	A uint8 // Alpha
	R uint8 // Red
	G uint8 // Green
	B uint8 // Blue
}

// NewColor32 creates a new Color32
func NewColor32(a, r, g, b uint8) Color32 {
	return Color32{A: a, R: r, G: g, B: b}
}

// ToInt32 converts the Color32 to a packed int32 value (ARGB format)
func (c Color32) ToInt32() int32 {
	return int32(c.A)<<24 | int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}

// PaletteColorLUT represents a palette color lookup table
type PaletteColorLUT struct {
	// Size is the number of entries in the LUT (0 means 65536)
	Size int
	// Bits is the number of bits per entry (8 or 16)
	Bits int
	// Red channel LUT data
	Red []byte
	// Green channel LUT data
	Green []byte
	// Blue channel LUT data
	Blue []byte
	// LUT is the parsed color lookup table
	LUT []Color32
}

// NewPaletteColorLUT creates a new palette color LUT from descriptor and data
func NewPaletteColorLUT(descriptorRed []uint16, red, green, blue []byte) (*PaletteColorLUT, error) {
	if len(descriptorRed) < 3 {
		return nil, fmt.Errorf("palette descriptor must have at least 3 values")
	}

	size := int(descriptorRed[0])
	bits := int(descriptorRed[2])

	// If the LUT size is 0, that means it's 65536 in size
	if size == 0 {
		size = 65536
	}

	p := &PaletteColorLUT{
		Size:  size,
		Bits:  bits,
		Red:   red,
		Green: green,
		Blue:  blue,
		LUT:   make([]Color32, size),
	}

	// Parse the LUT
	if err := p.parseLUT(); err != nil {
		return nil, err
	}

	return p, nil
}

// parseLUT parses the raw palette data into Color32 entries
func (p *PaletteColorLUT) parseLUT() error {
	if len(p.Red) == p.Size && len(p.Green) == p.Size && len(p.Blue) == p.Size {
		// 8-bit LUT entries
		for i := 0; i < p.Size; i++ {
			p.LUT[i] = NewColor32(0xFF, p.Red[i], p.Green[i], p.Blue[i])
		}
	} else if len(p.Red) >= p.Size*2 && len(p.Green) >= p.Size*2 && len(p.Blue) >= p.Size*2 {
		// 16-bit LUT entries... we only support 8-bit until someone can find a sample image with a 16-bit palette

		// 8-bit entries with 16-bits allocated
		offset := 0

		// 16-bit entries with 8-bits stored
		if p.Bits == 16 {
			offset = 1
		}

		for i := 0; i < p.Size; i++ {
			idx := i*2 + offset
			if idx >= len(p.Red) || idx >= len(p.Green) || idx >= len(p.Blue) {
				return fmt.Errorf("palette LUT index out of range")
			}
			p.LUT[i] = NewColor32(0xFF, p.Red[idx], p.Green[idx], p.Blue[idx])
		}
	} else {
		return fmt.Errorf("invalid palette color LUT data size: red=%d, green=%d, blue=%d, expected size=%d",
			len(p.Red), len(p.Green), len(p.Blue), p.Size)
	}

	return nil
}

// GetColor returns the color for the specified pixel value
func (p *PaletteColorLUT) GetColor(pixelValue uint16) Color32 {
	if int(pixelValue) >= len(p.LUT) {
		// Return black for out of range values
		return NewColor32(0xFF, 0, 0, 0)
	}
	return p.LUT[pixelValue]
}

// ApplyToPixelData applies the palette LUT to pixel data
// Input: grayscale pixel values
// Output: RGB pixel data (interleaved)
func (p *PaletteColorLUT) ApplyToPixelData(pixelData []byte, bitsAllocated uint16) ([]byte, error) {
	var pixelCount int
	var rgbData []byte

	switch bitsAllocated {
	case 8:
		// 8-bit pixels
		pixelCount = len(pixelData)
		rgbData = make([]byte, pixelCount*3)

		for i := 0; i < pixelCount; i++ {
			pixelValue := uint16(pixelData[i])
			color := p.GetColor(pixelValue)
			rgbData[i*3] = color.R
			rgbData[i*3+1] = color.G
			rgbData[i*3+2] = color.B
		}
	case 16:
		// 16-bit pixels
		pixelCount = len(pixelData) / 2
		rgbData = make([]byte, pixelCount*3)

		for i := 0; i < pixelCount; i++ {
			// Little endian
			pixelValue := uint16(pixelData[i*2]) | uint16(pixelData[i*2+1])<<8
			color := p.GetColor(pixelValue)
			rgbData[i*3] = color.R
			rgbData[i*3+1] = color.G
			rgbData[i*3+2] = color.B
		}
	default:
		return nil, fmt.Errorf("unsupported bits allocated for palette color: %d", bitsAllocated)
	}

	return rgbData, nil
}
