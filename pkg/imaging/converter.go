// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"
)

// PixelDataConverter provides utilities to convert pixel data between different formats
type PixelDataConverter struct{}

// NewPixelDataConverter creates a new PixelDataConverter
func NewPixelDataConverter() *PixelDataConverter {
	return &PixelDataConverter{}
}

// ConvertPlanarToInterleavedGeneric converts planar configuration 1 data to interleaved layout.
// Supports arbitrary samplesPerPixel and bytesPerSample (1 or 2).
func ConvertPlanarToInterleavedGeneric(data []byte, samplesPerPixel int, bytesPerSample int) ([]byte, error) {
	if samplesPerPixel <= 1 {
		return data, nil
	}
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for planar conversion", bytesPerSample)
	}
	pixelCount := len(data) / (samplesPerPixel * bytesPerSample)
	if pixelCount == 0 {
		return data, nil
	}
	dst := make([]byte, len(data))
	for i := 0; i < pixelCount; i++ {
		for s := 0; s < samplesPerPixel; s++ {
			srcOffset := (s*pixelCount + i) * bytesPerSample
			dstOffset := (i*samplesPerPixel + s) * bytesPerSample
			copy(dst[dstOffset:dstOffset+bytesPerSample], data[srcOffset:srcOffset+bytesPerSample])
		}
	}
	return dst, nil
}

// ConvertMono1ToMono2 inverts grayscale samples for MONOCHROME1 to MONOCHROME2.
// Supports 8/16-bit, signed/unsigned.
func ConvertMono1ToMono2(data []byte, bitsStored uint16, bytesPerSample int, signed bool) ([]byte, error) {
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for mono conversion", bytesPerSample)
	}
	maxVal := uint32((1 << bitsStored) - 1)
	out := make([]byte, len(data))
	for off := 0; off+bytesPerSample <= len(data); off += bytesPerSample {
		var val uint32
		if bytesPerSample == 1 {
			if signed {
				val = uint32(int8(data[off]))
			} else {
				val = uint32(data[off])
			}
		} else {
			if signed {
				val = uint32(int16(binary.LittleEndian.Uint16(data[off:])))
			} else {
				val = uint32(binary.LittleEndian.Uint16(data[off:]))
			}
		}
		inv := (maxVal - val) & maxVal
		if bytesPerSample == 1 {
			out[off] = byte(inv)
		} else {
			binary.LittleEndian.PutUint16(out[off:], uint16(inv))
		}
	}
	return out, nil
}

// InterleavedToPlanar24 converts 24-bit pixels from interleaved (RGBRGBRGB...) to planar (RRR...GGG...BBB...)
func (c *PixelDataConverter) InterleavedToPlanar24(data []byte) []byte {
	newPixels := make([]byte, len(data))
	pixelCount := len(data) / 3

	for n := 0; n < pixelCount; n++ {
		newPixels[n+pixelCount*0] = data[n*3+0] // R
		newPixels[n+pixelCount*1] = data[n*3+1] // G
		newPixels[n+pixelCount*2] = data[n*3+2] // B
	}

	return newPixels
}

// PlanarToInterleaved24 converts 24-bit pixels from planar (RRR...GGG...BBB...) to interleaved (RGBRGBRGB...)
func (c *PixelDataConverter) PlanarToInterleaved24(data []byte) []byte {
	newPixels := make([]byte, len(data))
	pixelCount := len(data) / 3

	for n := 0; n < pixelCount; n++ {
		newPixels[n*3+0] = data[n+pixelCount*0] // R
		newPixels[n*3+1] = data[n+pixelCount*1] // G
		newPixels[n*3+2] = data[n+pixelCount*2] // B
	}

	return newPixels
}

// YBRFullToRGB converts YBR_FULL photometric interpretation pixels to RGB
func (c *PixelDataConverter) YBRFullToRGB(data []byte) []byte {
	newPixels := make([]byte, len(data))

	for n := 0; n < len(data); n += 3 {
		y := int(data[n+0])
		cb := int(data[n+1])
		cr := int(data[n+2])

		// YBR_FULL to RGB conversion (BT.601 full range)
		// Note: fo-dicom uses slightly different coefficients
		r := y + int(1.4020*(float64(cr)-128)+0.5)
		g := y - int(0.3441*(float64(cb)-128)+0.7141*(float64(cr)-128)-0.5)
		b := y + int(1.7720*(float64(cb)-128)+0.5)

		newPixels[n+0] = clampByte(r)
		newPixels[n+1] = clampByte(g)
		newPixels[n+2] = clampByte(b)
	}

	return newPixels
}

// YBRFull422ToRGB converts YBR_FULL_422 photometric interpretation pixels to RGB (4:2:2 subsampling)
func (c *PixelDataConverter) YBRFull422ToRGB(data []byte, _ int) []byte {
	// YBR_FULL_422 format (fo-dicom): Y1 Y2 Cb Cr for every 2 pixels
	pixelCount := (len(data) * 2) / 4
	newPixels := make([]byte, pixelCount*3)

	for i, j := 0, 0; i < len(data); i += 4 {
		y1 := int(data[i])
		y2 := int(data[i+1])
		cb := int(data[i+2])
		cr := int(data[i+3])

		// Convert first pixel
		r1 := y1 + int(1.4020*(float64(cr)-128)+0.5)
		g1 := y1 - int(0.3441*(float64(cb)-128)+0.7141*(float64(cr)-128)-0.5)
		b1 := y1 + int(1.7720*(float64(cb)-128)+0.5)

		newPixels[j+0] = clampByte(r1)
		newPixels[j+1] = clampByte(g1)
		newPixels[j+2] = clampByte(b1)

		// Convert second pixel
		r2 := y2 + int(1.4020*(float64(cr)-128)+0.5)
		g2 := y2 - int(0.3441*(float64(cb)-128)+0.7141*(float64(cr)-128)-0.5)
		b2 := y2 + int(1.7720*(float64(cb)-128)+0.5)

		newPixels[j+3] = clampByte(r2)
		newPixels[j+4] = clampByte(g2)
		newPixels[j+5] = clampByte(b2)

		j += 6
	}

	return newPixels
}

// ConvertYBRFullToRGB wraps the converter logic with error reporting.
func ConvertYBRFullToRGB(data []byte) ([]byte, error) {
	if len(data)%3 != 0 {
		return nil, fmt.Errorf("invalid YBR_FULL length %d", len(data))
	}
	c := NewPixelDataConverter()
	return c.YBRFullToRGB(data), nil
}

// ConvertYBRFull422ToRGB wraps the converter logic with error reporting.
func ConvertYBRFull422ToRGB(data []byte, width int) ([]byte, error) {
	if width <= 0 {
		return nil, fmt.Errorf("width must be positive")
	}
	if len(data)%4 != 0 {
		return nil, fmt.Errorf("invalid YBR_FULL_422 length %d", len(data))
	}
	c := NewPixelDataConverter()
	return c.YBRFull422ToRGB(data, width), nil
}

// ConvertYBRPartial422ToRGB converts YBR_PARTIAL_422 (8-bit) to RGB interleaved.
func ConvertYBRPartial422ToRGB(data []byte, width int) ([]byte, error) {
	if width <= 0 {
		return nil, fmt.Errorf("width must be positive")
	}
	if len(data)%4 != 0 {
		return nil, fmt.Errorf("invalid YBR_PARTIAL_422 length %d", len(data))
	}
	pixelCount := (len(data) * 2) / 4
	newPixels := make([]byte, pixelCount*3)
	var j, col int
	for i := 0; i < len(data); i += 4 {
		y1 := float64(data[i])
		y2 := float64(data[i+1])
		cb := float64(data[i+2])
		cr := float64(data[i+3])

		r1 := 1.1644*(y1-16) + 1.5960*(cr-128)
		g1 := 1.1644*(y1-16) - 0.3917*(cb-128) - 0.8130*(cr-128)
		b1 := 1.1644*(y1-16) + 2.0173*(cb-128)

		newPixels[j+0] = clampByte(int(r1 + 0.5))
		newPixels[j+1] = clampByte(int(g1 + 0.5))
		newPixels[j+2] = clampByte(int(b1 + 0.5))
		j += 3
		col++
		if col == width {
			col = 0
			continue
		}

		r2 := 1.1644*(y2-16) + 1.5960*(cr-128)
		g2 := 1.1644*(y2-16) - 0.3917*(cb-128) - 0.8130*(cr-128)
		b2 := 1.1644*(y2-16) + 2.0173*(cb-128)

		newPixels[j+0] = clampByte(int(r2 + 0.5))
		newPixels[j+1] = clampByte(int(g2 + 0.5))
		newPixels[j+2] = clampByte(int(b2 + 0.5))
		j += 3
		col++
		if col == width {
			col = 0
		}
	}
	return newPixels[:j], nil
}

// ConvertYBRICTToRGB converts YBR_ICT (JPEG2000 irreversible transform) to RGB.
// Supports BitsAllocated 8 or 16; samples are assumed interleaved Y,Cb,Cr.
func ConvertYBRICTToRGB(data []byte, bitsAllocated int) ([]byte, error) {
	if bitsAllocated != 8 && bitsAllocated != 16 {
		return nil, fmt.Errorf("YBR_ICT conversion supports 8 or 16 bits, got %d", bitsAllocated)
	}
	bytesPerSample := bitsAllocated / 8
	if len(data)%(3*bytesPerSample) != 0 {
		return nil, fmt.Errorf("invalid YBR_ICT length %d", len(data))
	}
	dst := make([]byte, len(data))
	for i, j := 0, 0; i < len(data); i += 3 * bytesPerSample {
		var y, cb, cr float64
		if bytesPerSample == 1 {
			y = float64(int8(data[i]))
			cb = float64(int8(data[i+1]))
			cr = float64(int8(data[i+2]))
		} else {
			y = float64(int16(binary.LittleEndian.Uint16(data[i:])))
			cb = float64(int16(binary.LittleEndian.Uint16(data[i+2:])))
			cr = float64(int16(binary.LittleEndian.Uint16(data[i+4:])))
		}
		r := y + 1.40200*cr
		g := y - 0.344136*cb - 0.714136*cr
		b := y + 1.77200*cb
		if bytesPerSample == 1 {
			dst[j] = clampByte(int(r + 0.5))
			dst[j+1] = clampByte(int(g + 0.5))
			dst[j+2] = clampByte(int(b + 0.5))
			j += 3
		} else {
			if r < 0 {
				r = 0
			} else if r > 65535 {
				r = 65535
			}
			if g < 0 {
				g = 0
			} else if g > 65535 {
				g = 65535
			}
			if b < 0 {
				b = 0
			} else if b > 65535 {
				b = 65535
			}
			binary.LittleEndian.PutUint16(dst[j:], uint16(r+0.5))
			binary.LittleEndian.PutUint16(dst[j+2:], uint16(g+0.5))
			binary.LittleEndian.PutUint16(dst[j+4:], uint16(b+0.5))
			j += 6
		}
	}
	return dst, nil
}

// ConvertYBRRCTToRGB converts YBR_RCT (JPEG2000 reversible transform) to RGB.
// Supports BitsAllocated 8 or 16; samples assumed interleaved Y,Cb,Cr.
func ConvertYBRRCTToRGB(data []byte, bitsAllocated int) ([]byte, error) {
	if bitsAllocated != 8 && bitsAllocated != 16 {
		return nil, fmt.Errorf("YBR_RCT conversion supports 8 or 16 bits, got %d", bitsAllocated)
	}
	bytesPerSample := bitsAllocated / 8
	if len(data)%(3*bytesPerSample) != 0 {
		return nil, fmt.Errorf("invalid YBR_RCT length %d", len(data))
	}
	dst := make([]byte, len(data))
	for i, j := 0, 0; i < len(data); i += 3 * bytesPerSample {
		var y, cb, cr int32
		if bytesPerSample == 1 {
			y = int32(int8(data[i]))
			cb = int32(int8(data[i+1]))
			cr = int32(int8(data[i+2]))
		} else {
			y = int32(int16(binary.LittleEndian.Uint16(data[i:])))
			cb = int32(int16(binary.LittleEndian.Uint16(data[i+2:])))
			cr = int32(int16(binary.LittleEndian.Uint16(data[i+4:])))
		}
		g := y - ((cb + cr) >> 2)
		r := cr + g
		b := cb + g
		if bytesPerSample == 1 {
			dst[j] = clampByte(int(r))
			dst[j+1] = clampByte(int(g))
			dst[j+2] = clampByte(int(b))
			j += 3
		} else {
			if r < 0 {
				r = 0
			} else if r > 65535 {
				r = 65535
			}
			if g < 0 {
				g = 0
			} else if g > 65535 {
				g = 65535
			}
			if b < 0 {
				b = 0
			} else if b > 65535 {
				b = 65535
			}
			binary.LittleEndian.PutUint16(dst[j:], uint16(r))
			binary.LittleEndian.PutUint16(dst[j+2:], uint16(g))
			binary.LittleEndian.PutUint16(dst[j+4:], uint16(b))
			j += 6
		}
	}
	return dst, nil
}

// YBRPartialToRGB converts YBR_PARTIAL_422 photometric interpretation pixels to RGB
func (c *PixelDataConverter) YBRPartialToRGB(data []byte) []byte {
	newPixels := make([]byte, len(data))

	for n := 0; n < len(data); n += 3 {
		y := int(data[n+0])
		cb := int(data[n+1])
		cr := int(data[n+2])

		// YBR_PARTIAL (BT.601 limited range) to RGB conversion
		// R = 1.1644*(Y - 16) + 1.5960*(Cr - 128)
		// G = 1.1644*(Y - 16) - 0.3918*(Cb - 128) - 0.8130*(Cr - 128)
		// B = 1.1644*(Y - 16) + 2.0172*(Cb - 128)
		r := int(1.1644*float64(y-16) + 1.5960*float64(cr-128) + 0.5)
		g := int(1.1644*float64(y-16) - 0.3918*float64(cb-128) - 0.8130*float64(cr-128) + 0.5)
		b := int(1.1644*float64(y-16) + 2.0172*float64(cb-128) + 0.5)

		newPixels[n+0] = clampByte(r)
		newPixels[n+1] = clampByte(g)
		newPixels[n+2] = clampByte(b)
	}

	return newPixels
}

// RGBToYBRFull converts RGB photometric interpretation pixels to YBR_FULL
func (c *PixelDataConverter) RGBToYBRFull(data []byte) []byte {
	newPixels := make([]byte, len(data))

	for n := 0; n < len(data); n += 3 {
		r := float64(data[n+0])
		g := float64(data[n+1])
		b := float64(data[n+2])

		// RGB to YBR_FULL conversion (BT.601 full range)
		y := 0.299*r + 0.587*g + 0.114*b
		cb := -0.168736*r - 0.331264*g + 0.5*b + 128
		cr := 0.5*r - 0.418688*g - 0.081312*b + 128

		newPixels[n+0] = clampByte(int(y + 0.5))
		newPixels[n+1] = clampByte(int(cb + 0.5))
		newPixels[n+2] = clampByte(int(cr + 0.5))
	}

	return newPixels
}

// SwapBytes16 swaps bytes for 16-bit pixel data (Little Endian <-> Big Endian)
func (c *PixelDataConverter) SwapBytes16(data []byte) []byte {
	newPixels := make([]byte, len(data))

	for i := 0; i < len(data); i += 2 {
		newPixels[i] = data[i+1]
		newPixels[i+1] = data[i]
	}

	return newPixels
}

// SwapBytes32 swaps bytes for 32-bit pixel data (Little Endian <-> Big Endian)
func (c *PixelDataConverter) SwapBytes32(data []byte) []byte {
	newPixels := make([]byte, len(data))

	for i := 0; i < len(data); i += 4 {
		newPixels[i] = data[i+3]
		newPixels[i+1] = data[i+2]
		newPixels[i+2] = data[i+1]
		newPixels[i+3] = data[i]
	}

	return newPixels
}

// clampByte clamps an integer value to the byte range [0, 255]
func clampByte(value int) byte {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return byte(value)
}
