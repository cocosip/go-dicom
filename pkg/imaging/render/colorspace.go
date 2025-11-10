// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import "math"

// ColorSpaceConverter provides color space conversion functions
type ColorSpaceConverter struct{}

// RGBToYBRFull converts RGB to YBR_FULL color space
// Y = +0.299*R + 0.587*G + 0.114*B
// Cb = -0.168736*R - 0.331264*G + 0.5*B + 128
// Cr = +0.5*R - 0.418688*G - 0.081312*B + 128
func (c *ColorSpaceConverter) RGBToYBRFull(r, g, b uint8) (y, cb, cr uint8) {
	rf, gf, bf := float64(r), float64(g), float64(b)

	yf := 0.299*rf + 0.587*gf + 0.114*bf
	cbf := -0.168736*rf - 0.331264*gf + 0.5*bf + 128
	crf := 0.5*rf - 0.418688*gf - 0.081312*bf + 128

	y = clampUint8(yf)
	cb = clampUint8(cbf)
	cr = clampUint8(crf)
	return
}

// YBRFullToRGB converts YBR_FULL to RGB color space
// R = Y + 1.402 * (Cr - 128)
// G = Y - 0.344136 * (Cb - 128) - 0.714136 * (Cr - 128)
// B = Y + 1.772 * (Cb - 128)
func (c *ColorSpaceConverter) YBRFullToRGB(y, cb, cr uint8) (r, g, b uint8) {
	yf, cbf, crf := float64(y), float64(cb), float64(cr)

	cbShift := cbf - 128
	crShift := crf - 128

	rf := yf + 1.402*crShift
	gf := yf - 0.344136*cbShift - 0.714136*crShift
	bf := yf + 1.772*cbShift

	r = clampUint8(rf)
	g = clampUint8(gf)
	b = clampUint8(bf)
	return
}

// RGBToYBRFull422 converts RGB to YBR_FULL_422 color space (4:2:2 subsampling)
// Same conversion as YBR_FULL but with horizontal chroma subsampling
func (c *ColorSpaceConverter) RGBToYBRFull422(r1, g1, b1, r2, g2, b2 uint8) (y1, y2, cb, cr uint8) {
	// Convert both pixels to YCbCr
	y1Val, cb1, cr1 := c.RGBToYBRFull(r1, g1, b1)
	y2Val, cb2, cr2 := c.RGBToYBRFull(r2, g2, b2)

	// Average the chroma values
	y1 = y1Val
	y2 = y2Val
	cb = uint8((int(cb1) + int(cb2)) / 2)
	cr = uint8((int(cr1) + int(cr2)) / 2)
	return
}

// YBRFull422ToRGB converts YBR_FULL_422 to RGB (4:2:2 subsampling)
// Both pixels share the same Cb and Cr values
func (c *ColorSpaceConverter) YBRFull422ToRGB(y1, y2, cb, cr uint8) (r1, g1, b1, r2, g2, b2 uint8) {
	r1, g1, b1 = c.YBRFullToRGB(y1, cb, cr)
	r2, g2, b2 = c.YBRFullToRGB(y2, cb, cr)
	return
}

// RGBToYBRPartial422 converts RGB to YBR_PARTIAL_422
// This uses the BT.601 limited range encoding
// Y' = 0.2568*R + 0.5041*G + 0.0979*B + 16
// Cb = -0.1482*R - 0.2910*G + 0.4392*B + 128
// Cr = 0.4392*R - 0.3678*G - 0.0714*B + 128
func (c *ColorSpaceConverter) RGBToYBRPartial422(r, g, b uint8) (y, cb, cr uint8) {
	rf, gf, bf := float64(r), float64(g), float64(b)

	yf := 0.2568*rf + 0.5041*gf + 0.0979*bf + 16
	cbf := -0.1482*rf - 0.2910*gf + 0.4392*bf + 128
	crf := 0.4392*rf - 0.3678*gf - 0.0714*bf + 128

	y = clampUint8(yf)
	cb = clampUint8(cbf)
	cr = clampUint8(crf)
	return
}

// YBRPartial422ToRGB converts YBR_PARTIAL_422 to RGB
// This uses the BT.601 limited range decoding
// R = 1.1644*(Y' - 16) + 1.5960*(Cr - 128)
// G = 1.1644*(Y' - 16) - 0.3918*(Cb - 128) - 0.8130*(Cr - 128)
// B = 1.1644*(Y' - 16) + 2.0172*(Cb - 128)
func (c *ColorSpaceConverter) YBRPartial422ToRGB(y, cb, cr uint8) (r, g, b uint8) {
	yf, cbf, crf := float64(y), float64(cb), float64(cr)

	yShift := yf - 16
	cbShift := cbf - 128
	crShift := crf - 128

	rf := 1.1644*yShift + 1.5960*crShift
	gf := 1.1644*yShift - 0.3918*cbShift - 0.8130*crShift
	bf := 1.1644*yShift + 2.0172*cbShift

	r = clampUint8(rf)
	g = clampUint8(gf)
	b = clampUint8(bf)
	return
}

// RGBToYBRICT converts RGB to YBR_ICT (Integer Color Transform, used in JPEG 2000)
// Y = (R + 2*G + B) / 4
// Cb = B - G
// Cr = R - G
func (c *ColorSpaceConverter) RGBToYBRICT(r, g, b uint8) (y, cb, cr int16) {
	ri, gi, bi := int16(r), int16(g), int16(b)

	y = (ri + 2*gi + bi) / 4
	cb = bi - gi
	cr = ri - gi
	return
}

// YBRICTToRGB converts YBR_ICT to RGB
// G = Y - (Cb + Cr) / 4
// R = Cr + G
// B = Cb + G
func (c *ColorSpaceConverter) YBRICTToRGB(y, cb, cr int16) (r, g, b uint8) {
	g16 := y - (cb+cr)/4
	r16 := cr + g16
	b16 := cb + g16

	r = clampUint8(float64(r16))
	g = clampUint8(float64(g16))
	b = clampUint8(float64(b16))
	return
}

// RGBToYBRRCT converts RGB to YBR_RCT (Reversible Color Transform, used in JPEG 2000 lossless)
// Y = ⌊(R + 2*G + B) / 4⌋
// Cb = B - G
// Cr = R - G
func (c *ColorSpaceConverter) RGBToYBRRCT(r, g, b uint8) (y, cb, cr int16) {
	ri, gi, bi := int16(r), int16(g), int16(b)

	y = (ri + 2*gi + bi) >> 2 // Floor division by 4
	cb = bi - gi
	cr = ri - gi
	return
}

// YBRRCTToRGB converts YBR_RCT to RGB
// G = Y - ⌊(Cb + Cr) / 4⌋
// R = Cr + G
// B = Cb + G
func (c *ColorSpaceConverter) YBRRCTToRGB(y, cb, cr int16) (r, g, b uint8) {
	g16 := y - ((cb + cr) >> 2)
	r16 := cr + g16
	b16 := cb + g16

	r = clampUint8(float64(r16))
	g = clampUint8(float64(g16))
	b = clampUint8(float64(b16))
	return
}

// ConvertToRGB converts pixel data from various color spaces to RGB
// photometric specifies the source color space
// planarConfig indicates if color channels are interleaved (0) or planar (1)
func (c *ColorSpaceConverter) ConvertToRGB(data []byte, width, height int, photometric string, planarConfig int) ([]byte, error) {
	switch photometric {
	case "RGB":
		// Already RGB, no conversion needed
		return data, nil

	case "YBR_FULL":
		return c.convertYBRFullToRGB(data, width, height, planarConfig)

	case "YBR_FULL_422":
		return c.convertYBRFull422ToRGB(data, width, height, planarConfig)

	case "YBR_PARTIAL_422":
		return c.convertYBRPartial422ToRGB(data, width, height, planarConfig)

	default:
		// Unsupported or grayscale, return as-is
		return data, nil
	}
}

func (c *ColorSpaceConverter) convertYBRFullToRGB(data []byte, width, height, planarConfig int) ([]byte, error) {
	result := make([]byte, len(data))
	pixelCount := width * height

	if planarConfig == 0 {
		// Interleaved: YCBCR YCBCR YCBCR...
		for i := 0; i < pixelCount; i++ {
			idx := i * 3
			y, cb, cr := data[idx], data[idx+1], data[idx+2]
			r, g, b := c.YBRFullToRGB(y, cb, cr)
			result[idx] = r
			result[idx+1] = g
			result[idx+2] = b
		}
	} else {
		// Planar: YYY... CbCbCb... CrCrCr...
		for i := 0; i < pixelCount; i++ {
			y := data[i]
			cb := data[pixelCount+i]
			cr := data[2*pixelCount+i]
			r, g, b := c.YBRFullToRGB(y, cb, cr)
			result[i*3] = r
			result[i*3+1] = g
			result[i*3+2] = b
		}
	}

	return result, nil
}

func (c *ColorSpaceConverter) convertYBRFull422ToRGB(data []byte, width, height, planarConfig int) ([]byte, error) {
	result := make([]byte, width*height*3)

	if planarConfig == 0 {
		// Interleaved 4:2:2: Y1 Cb Y2 Cr Y3 Cb Y4 Cr...
		outIdx := 0
		for i := 0; i < len(data); i += 4 {
			y1, cb, y2, cr := data[i], data[i+1], data[i+2], data[i+3]
			r1, g1, b1, r2, g2, b2 := c.YBRFull422ToRGB(y1, y2, cb, cr)

			result[outIdx] = r1
			result[outIdx+1] = g1
			result[outIdx+2] = b1
			result[outIdx+3] = r2
			result[outIdx+4] = g2
			result[outIdx+5] = b2
			outIdx += 6
		}
	}

	return result, nil
}

func (c *ColorSpaceConverter) convertYBRPartial422ToRGB(data []byte, width, height, planarConfig int) ([]byte, error) {
	result := make([]byte, width*height*3)

	if planarConfig == 0 {
		// Interleaved 4:2:2: Y1 Cb Y2 Cr...
		outIdx := 0
		for i := 0; i < len(data); i += 4 {
			y1, cb, y2, cr := data[i], data[i+1], data[i+2], data[i+3]

			r1, g1, b1 := c.YBRPartial422ToRGB(y1, cb, cr)
			r2, g2, b2 := c.YBRPartial422ToRGB(y2, cb, cr)

			result[outIdx] = r1
			result[outIdx+1] = g1
			result[outIdx+2] = b1
			result[outIdx+3] = r2
			result[outIdx+4] = g2
			result[outIdx+5] = b2
			outIdx += 6
		}
	}

	return result, nil
}

// clampUint8 clamps a float64 value to the uint8 range [0, 255]
func clampUint8(value float64) uint8 {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return uint8(math.Round(value))
}

// NewColorSpaceConverter creates a new ColorSpaceConverter
func NewColorSpaceConverter() *ColorSpaceConverter {
	return &ColorSpaceConverter{}
}
