// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

// PixelDataConverter provides utilities to convert pixel data between different formats
type PixelDataConverter struct{}

// NewPixelDataConverter creates a new PixelDataConverter
func NewPixelDataConverter() *PixelDataConverter {
	return &PixelDataConverter{}
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
func (c *PixelDataConverter) YBRFull422ToRGB(data []byte, width int) []byte {
	// YBR_FULL_422 format: Y0 Cb Y1 Cr for every 2 pixels
	pixelCount := (len(data) * 2) / 4
	newPixels := make([]byte, pixelCount*3)

	for i, j := 0, 0; i < len(data); i += 4 {
		y1 := int(data[i])
		cb := int(data[i+1])
		y2 := int(data[i+2])
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
