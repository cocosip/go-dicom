// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package imaging

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
)

// PaletteColorLUT represents a palette color lookup table
type PaletteColorLUT struct {
	// Size is the number of entries in the LUT (0 means 65536)
	Size int
	// Bits is the number of bits per entry (8 or 16)
	Bits int
	// FirstMappedValue is the input value mapped to LUT entry zero.
	FirstMappedValue int
	// ByteOrder is the byte order of 16-bit LUT Data. Defaults to
	// binary.LittleEndian.
	ByteOrder binary.ByteOrder
	// Red channel LUT data
	Red []byte
	// Green channel LUT data
	Green []byte
	// Blue channel LUT data
	Blue []byte
	// Alpha channel LUT data. When absent, output is opaque RGB.
	Alpha []byte
	// HasAlpha reports whether an Alpha Palette Color LUT was supplied.
	HasAlpha bool
	// LUT is the parsed color lookup table
	LUT []colorconv.Color32
}

// PaletteColorLUTOptions controls descriptor interpretation for palette LUTs
// whose LUT Descriptor uses Signed Short for the first mapped value and/or
// whose LUT Data uses Big Endian byte order.
type PaletteColorLUTOptions struct {
	ByteOrder              binary.ByteOrder
	SignedFirstMappedValue bool
}

// NewPaletteColorLUT creates a new palette color LUT from descriptor and data
func NewPaletteColorLUT(descriptorRed []uint16, red, green, blue []byte) (*PaletteColorLUT, error) {
	return newPaletteColorLUT(descriptorRed, nil, red, green, blue, nil, PaletteColorLUTOptions{})
}

// NewPaletteColorLUTWithOptions creates a palette color LUT with an explicit
// byte order and first-mapped-value interpretation.
func NewPaletteColorLUTWithOptions(descriptorRed []uint16, red, green, blue []byte, options PaletteColorLUTOptions) (*PaletteColorLUT, error) {
	return newPaletteColorLUT(descriptorRed, nil, red, green, blue, nil, options)
}

// NewPaletteColorLUTWithAlpha creates a palette color LUT with the optional
// DICOM Alpha Palette Color LUT. Alpha descriptors must share the RGB entry
// count and first mapped value, and must declare 8 bits per entry.
func NewPaletteColorLUTWithAlpha(descriptorRed, descriptorAlpha []uint16, red, green, blue, alpha []byte) (*PaletteColorLUT, error) {
	return newPaletteColorLUT(descriptorRed, descriptorAlpha, red, green, blue, alpha, PaletteColorLUTOptions{})
}

// NewPaletteColorLUTWithAlphaAndOptions creates an alpha-capable palette color
// LUT with an explicit byte order and first-mapped-value interpretation.
func NewPaletteColorLUTWithAlphaAndOptions(descriptorRed, descriptorAlpha []uint16, red, green, blue, alpha []byte, options PaletteColorLUTOptions) (*PaletteColorLUT, error) {
	return newPaletteColorLUT(descriptorRed, descriptorAlpha, red, green, blue, alpha, options)
}

func newPaletteColorLUT(descriptorRed, descriptorAlpha []uint16, red, green, blue, alpha []byte, options PaletteColorLUTOptions) (*PaletteColorLUT, error) {
	byteOrder := options.ByteOrder
	if byteOrder == nil {
		byteOrder = binary.LittleEndian
	}
	if len(descriptorRed) != 3 {
		return nil, fmt.Errorf("palette descriptor must have exactly 3 values")
	}

	size := int(descriptorRed[0])
	bits := int(descriptorRed[2])
	if bits != 8 && bits != 16 {
		return nil, fmt.Errorf("palette LUT bits per entry must be 8 or 16, got %d", bits)
	}

	// If the LUT size is 0, that means it's 65536 in size
	if size == 0 {
		size = 65536
	}
	if descriptorAlpha != nil {
		if len(descriptorAlpha) != 3 {
			return nil, fmt.Errorf("alpha palette descriptor must have exactly 3 values")
		}
		alphaSize := int(descriptorAlpha[0])
		if alphaSize == 0 {
			alphaSize = 65536
		}
		if alphaSize != size || descriptorAlpha[1] != descriptorRed[1] {
			return nil, fmt.Errorf("alpha palette descriptor must match RGB entry count and first mapped value")
		}
		if descriptorAlpha[2] != 8 {
			return nil, fmt.Errorf("alpha palette LUT bits per entry must be 8, got %d", descriptorAlpha[2])
		}
	}

	firstMapped := int(descriptorRed[1])
	if options.SignedFirstMappedValue {
		firstMapped = int(int16(descriptorRed[1]))
	}

	p := &PaletteColorLUT{
		Size:             size,
		Bits:             bits,
		FirstMappedValue: firstMapped,
		ByteOrder:        byteOrder,
		Red:              red,
		Green:            green,
		Blue:             blue,
		Alpha:            alpha,
		HasAlpha:         descriptorAlpha != nil,
		LUT:              make([]colorconv.Color32, size),
	}

	// Parse the LUT
	if err := p.parseLUT(); err != nil {
		return nil, err
	}

	return p, nil
}

// parseLUT parses the raw palette data into Color32 entries
func (p *PaletteColorLUT) parseLUT() error {
	compactLength := p.Size
	if compactLength%2 != 0 {
		compactLength++
	}
	wordLength := p.Size * 2

	if p.Bits == 8 && paletteDataLengthsEqual(p, p.Size, compactLength) {
		// 8-bit LUT entries
		for i := 0; i < p.Size; i++ {
			p.LUT[i] = colorconv.NewColor32(0xFF, p.Red[i], p.Green[i], p.Blue[i])
		}
	} else if paletteDataLengthsEqual(p, wordLength) {
		// 8-bit entries packed in 16-bit words occupy the word's low byte; a
		// 16-bit entry's approximation is its high byte. The byte positions
		// depend on the LUT Data byte order.
		wordOffset := 1
		if p.Bits == 8 {
			wordOffset = 0
		}
		if p.ByteOrder == binary.BigEndian {
			wordOffset = 1 - wordOffset
		}

		for i := 0; i < p.Size; i++ {
			idx := i*2 + wordOffset
			p.LUT[i] = colorconv.NewColor32(0xFF, p.Red[idx], p.Green[idx], p.Blue[idx])
		}
	} else {
		return fmt.Errorf("invalid palette color LUT data size: red=%d, green=%d, blue=%d, expected compact length %d or word length %d",
			len(p.Red), len(p.Green), len(p.Blue), compactLength, wordLength)
	}
	if p.HasAlpha {
		if len(p.Alpha) == compactLength || len(p.Alpha) == p.Size {
			for i := 0; i < p.Size; i++ {
				p.LUT[i].A = p.Alpha[i]
			}
		} else if len(p.Alpha) == wordLength {
			alphaOffset := 0
			if p.ByteOrder == binary.BigEndian {
				alphaOffset = 1
			}
			for i := 0; i < p.Size; i++ {
				p.LUT[i].A = p.Alpha[i*2+alphaOffset]
			}
		} else {
			return fmt.Errorf("invalid alpha palette LUT data size: got %d, expected compact length %d or word length %d",
				len(p.Alpha), compactLength, wordLength)
		}
	}

	return nil
}

func paletteDataLengthsEqual(p *PaletteColorLUT, lengths ...int) bool {
	for _, length := range lengths {
		if len(p.Red) == length && len(p.Green) == length && len(p.Blue) == length {
			return true
		}
	}
	return false
}

// GetColor returns the color for the specified pixel value
func (p *PaletteColorLUT) GetColor(pixelValue uint16) colorconv.Color32 {
	return p.GetColorValue(int32(pixelValue))
}

// GetColorValue returns the color for a signed-or-unsigned pixel value. Unlike
// GetColor it can express the negative samples that signed pixel data and a
// Signed Short LUT Descriptor can produce.
func (p *PaletteColorLUT) GetColorValue(pixelValue int32) colorconv.Color32 {
	if len(p.LUT) == 0 {
		return colorconv.NewColor32(0xFF, 0, 0, 0)
	}
	index := int(pixelValue) - p.FirstMappedValue
	if index < 0 {
		index = 0
	}
	if index >= len(p.LUT) {
		index = len(p.LUT) - 1
	}
	return p.LUT[index]
}

// ApplyToPixelData applies the palette LUT to pixel data
// Input: grayscale pixel values
// Output: RGB pixel data (interleaved)
func (p *PaletteColorLUT) ApplyToPixelData(pixelData []byte, bitsAllocated uint16) ([]byte, error) {
	return p.ApplyToPixelDataWithByteOrder(pixelData, bitsAllocated, binary.LittleEndian)
}

// ApplyToPixelDataWithByteOrder applies the palette using the supplied byte
// order for 16-bit input samples. Output is always RGB (3 samples per pixel);
// when an Alpha palette is present its opacity is composited over black.
func (p *PaletteColorLUT) ApplyToPixelDataWithByteOrder(pixelData []byte, bitsAllocated uint16, order binary.ByteOrder) ([]byte, error) {
	if order == nil {
		order = binary.LittleEndian
	}
	var pixelCount int
	const channels = 3
	var colorData []byte

	switch bitsAllocated {
	case 8:
		// 8-bit pixels
		pixelCount = len(pixelData)
		colorData = make([]byte, pixelCount*channels)

		for i := 0; i < pixelCount; i++ {
			pixelValue := uint16(pixelData[i])
			writePaletteColor(colorData, i*channels, p.GetColor(pixelValue), p.HasAlpha)
		}
	case 16:
		if len(pixelData)%2 != 0 {
			return nil, fmt.Errorf("16-bit palette pixel data has odd byte length %d", len(pixelData))
		}
		// 16-bit pixels
		pixelCount = len(pixelData) / 2
		colorData = make([]byte, pixelCount*channels)

		for i := 0; i < pixelCount; i++ {
			pixelValue := order.Uint16(pixelData[i*2:])
			writePaletteColor(colorData, i*channels, p.GetColor(pixelValue), p.HasAlpha)
		}
	default:
		return nil, fmt.Errorf("unsupported bits allocated for palette color: %d", bitsAllocated)
	}

	return colorData, nil
}

func writePaletteColor(output []byte, offset int, value colorconv.Color32, alpha bool) {
	r, g, b := value.R, value.G, value.B
	if alpha {
		// Composite over black: c * a / 255 (rounded).
		r = uint8((uint32(value.R)*uint32(value.A) + 127) / 255)
		g = uint8((uint32(value.G)*uint32(value.A) + 127) / 255)
		b = uint8((uint32(value.B)*uint32(value.A) + 127) / 255)
	}
	output[offset] = r
	output[offset+1] = g
	output[offset+2] = b
}
