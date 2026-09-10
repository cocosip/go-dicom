// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package colorconv

import (
	"encoding/binary"
	"fmt"
)

// PlanarToInterleaved converts planar data to interleaved layout.
// Supports arbitrary samplesPerPixel and bytesPerSample (1 or 2).
func PlanarToInterleaved(data []byte, samplesPerPixel int, bytesPerSample int) ([]byte, error) {
	if samplesPerPixel <= 1 {
		return data, nil
	}
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for planar conversion", bytesPerSample)
	}
	pixelSize := samplesPerPixel * bytesPerSample
	if len(data)%pixelSize != 0 {
		return nil, fmt.Errorf("planar data length %d is not divisible by pixel size %d", len(data), pixelSize)
	}
	pixelCount := len(data) / pixelSize
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

// InterleavedToPlanar converts interleaved data to planar layout.
// Supports arbitrary samplesPerPixel and bytesPerSample (1 or 2).
func InterleavedToPlanar(data []byte, samplesPerPixel int, bytesPerSample int) ([]byte, error) {
	if samplesPerPixel <= 1 {
		return data, nil
	}
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for planar conversion", bytesPerSample)
	}
	pixelSize := samplesPerPixel * bytesPerSample
	if len(data)%pixelSize != 0 {
		return nil, fmt.Errorf("interleaved data length %d is not divisible by pixel size %d", len(data), pixelSize)
	}
	pixelCount := len(data) / pixelSize
	if pixelCount == 0 {
		return data, nil
	}
	dst := make([]byte, len(data))
	for i := 0; i < pixelCount; i++ {
		for s := 0; s < samplesPerPixel; s++ {
			srcOffset := (i*samplesPerPixel + s) * bytesPerSample
			dstOffset := (s*pixelCount + i) * bytesPerSample
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
	if bitsStored == 0 || int(bitsStored) > bytesPerSample*8 {
		return nil, fmt.Errorf("invalid BitsStored=%d for BytesAllocated=%d", bitsStored, bytesPerSample)
	}
	if len(data)%bytesPerSample != 0 {
		return nil, fmt.Errorf("mono data length %d is not divisible by BytesAllocated=%d", len(data), bytesPerSample)
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

// ConvertYBRFullToRGB converts interleaved YBR_FULL samples to RGB.
func ConvertYBRFullToRGB(data []byte) ([]byte, error) {
	if len(data)%3 != 0 {
		return nil, fmt.Errorf("invalid YBR_FULL length %d", len(data))
	}
	result := make([]byte, len(data))
	for i := 0; i < len(data); i += 3 {
		result[i], result[i+1], result[i+2] = YBRFullToRGB(data[i], data[i+1], data[i+2])
	}
	return result, nil
}

// ConvertYBRFull422ToRGB wraps the converter logic with error reporting.
func ConvertYBRFull422ToRGB(data []byte, width int) ([]byte, error) {
	if width <= 0 {
		return nil, fmt.Errorf("width must be positive")
	}
	if len(data)%4 != 0 {
		return nil, fmt.Errorf("invalid YBR_FULL_422 length %d", len(data))
	}
	result := make([]byte, len(data)/4*6)
	var j, column int
	for i := 0; i < len(data); i += 4 {
		r1, g1, b1, r2, g2, b2 := YBRFull422ToRGB(data[i], data[i+1], data[i+2], data[i+3])
		result[j], result[j+1], result[j+2] = r1, g1, b1
		j += 3
		column++
		if column == width {
			column = 0
			continue
		}
		result[j], result[j+1], result[j+2] = r2, g2, b2
		j += 3
		column++
		if column == width {
			column = 0
		}
	}
	return result[:j], nil
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

// ConvertYBRPartialToRGB converts interleaved YBR_PARTIAL samples to RGB.
func ConvertYBRPartialToRGB(data []byte) ([]byte, error) {
	if len(data)%3 != 0 {
		return nil, fmt.Errorf("invalid YBR_PARTIAL length %d", len(data))
	}
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

	return newPixels, nil
}

// ConvertRGBToYBRFull converts interleaved RGB samples to YBR_FULL.
func ConvertRGBToYBRFull(data []byte) ([]byte, error) {
	if len(data)%3 != 0 {
		return nil, fmt.Errorf("invalid RGB length %d", len(data))
	}
	newPixels := make([]byte, len(data))

	for n := 0; n < len(data); n += 3 {
		newPixels[n], newPixels[n+1], newPixels[n+2] = RGBToYBRFull(data[n], data[n+1], data[n+2])
	}

	return newPixels, nil
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
