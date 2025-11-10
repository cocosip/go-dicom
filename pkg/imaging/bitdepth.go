// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "math"

// BitDepth represents the bit depth information of pixel data
type BitDepth struct {
	// BitsAllocated is the number of bits allocated for each pixel sample
	BitsAllocated uint16
	// BitsStored is the number of bits actually stored for each pixel sample
	BitsStored uint16
	// HighBit is the index of the most significant bit (MSB)
	HighBit uint16
	// IsSigned indicates whether the pixel representation is signed
	IsSigned bool
}

// NewBitDepth creates a new BitDepth instance
func NewBitDepth(bitsAllocated, bitsStored, highBit uint16, isSigned bool) *BitDepth {
	return &BitDepth{
		BitsAllocated: bitsAllocated,
		BitsStored:    bitsStored,
		HighBit:       highBit,
		IsSigned:      isSigned,
	}
}

// MinimumValue returns the minimum pixel value based on bit depth
func (b *BitDepth) MinimumValue() float64 {
	if b.IsSigned {
		// Signed: -2^(n-1)
		return -math.Pow(2, float64(b.BitsStored-1))
	}
	// Unsigned: 0
	return 0
}

// MaximumValue returns the maximum pixel value based on bit depth
func (b *BitDepth) MaximumValue() float64 {
	if b.IsSigned {
		// Signed: 2^(n-1) - 1
		return math.Pow(2, float64(b.BitsStored-1)) - 1
	}
	// Unsigned: 2^n - 1
	return math.Pow(2, float64(b.BitsStored)) - 1
}

// BytesAllocated returns the number of bytes allocated per pixel sample
func (b *BitDepth) BytesAllocated() int {
	return int((b.BitsAllocated-1)/8 + 1)
}

// IsValid checks if the bit depth configuration is valid
func (b *BitDepth) IsValid() bool {
	// BitsStored must be <= BitsAllocated
	if b.BitsStored > b.BitsAllocated {
		return false
	}

	// HighBit must be < BitsAllocated
	if b.HighBit >= b.BitsAllocated {
		return false
	}

	// HighBit should typically be BitsStored - 1
	// But we allow some flexibility
	if b.HighBit < b.BitsStored-1 {
		return false
	}

	// Common bit allocations: 1, 8, 16, 32
	if b.BitsAllocated != 1 && b.BitsAllocated != 8 && b.BitsAllocated != 16 && b.BitsAllocated != 32 {
		return false
	}

	return true
}

// Mask returns the bitmask for the stored bits
func (b *BitDepth) Mask() uint32 {
	if b.BitsStored == 32 {
		return 0xFFFFFFFF
	}
	return (1 << b.BitsStored) - 1
}

// SignMask returns the sign bit mask for signed pixels
func (b *BitDepth) SignMask() uint32 {
	if !b.IsSigned || b.BitsStored == 0 {
		return 0
	}
	return 1 << (b.BitsStored - 1)
}

// ExtendSign extends the sign bit for signed pixels
// This converts a value stored in BitsStored to a full signed integer
func (b *BitDepth) ExtendSign(value uint32) int32 {
	if !b.IsSigned {
		return int32(value)
	}

	mask := b.Mask()
	signMask := b.SignMask()

	// Mask to stored bits
	value &= mask

	// Check if sign bit is set
	if (value & signMask) != 0 {
		// Extend sign bit
		return int32(value | ^mask)
	}

	return int32(value)
}

// Range returns the range of valid pixel values
func (b *BitDepth) Range() (minimumValue, maximumValue float64) {
	return b.MinimumValue(), b.MaximumValue()
}
