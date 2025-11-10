// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"testing"
)

func TestNewBitDepth(t *testing.T) {
	bd := NewBitDepth(16, 12, 11, true)

	if bd.BitsAllocated != 16 {
		t.Errorf("Expected BitsAllocated=16, got %d", bd.BitsAllocated)
	}
	if bd.BitsStored != 12 {
		t.Errorf("Expected BitsStored=12, got %d", bd.BitsStored)
	}
	if bd.HighBit != 11 {
		t.Errorf("Expected HighBit=11, got %d", bd.HighBit)
	}
	if !bd.IsSigned {
		t.Error("Expected IsSigned=true")
	}
}

func TestBitDepth_MinMaxValue(t *testing.T) {
	testCases := []struct {
		name      string
		bd        *BitDepth
		expectMin float64
		expectMax float64
	}{
		{
			name:      "8-bit unsigned",
			bd:        NewBitDepth(8, 8, 7, false),
			expectMin: 0,
			expectMax: 255,
		},
		{
			name:      "8-bit signed",
			bd:        NewBitDepth(8, 8, 7, true),
			expectMin: -128,
			expectMax: 127,
		},
		{
			name:      "12-bit unsigned",
			bd:        NewBitDepth(16, 12, 11, false),
			expectMin: 0,
			expectMax: 4095,
		},
		{
			name:      "12-bit signed",
			bd:        NewBitDepth(16, 12, 11, true),
			expectMin: -2048,
			expectMax: 2047,
		},
		{
			name:      "16-bit unsigned",
			bd:        NewBitDepth(16, 16, 15, false),
			expectMin: 0,
			expectMax: 65535,
		},
		{
			name:      "16-bit signed",
			bd:        NewBitDepth(16, 16, 15, true),
			expectMin: -32768,
			expectMax: 32767,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			minimumValue := tc.bd.MinimumValue()
			maximumValue := tc.bd.MaximumValue()

			if minimumValue != tc.expectMin {
				t.Errorf("MinimumValue: expected %f, got %f", tc.expectMin, minimumValue)
			}
			if maximumValue != tc.expectMax {
				t.Errorf("MaximumValue: expected %f, got %f", tc.expectMax, maximumValue)
			}
		})
	}
}

func TestBitDepth_BytesAllocated(t *testing.T) {
	testCases := []struct {
		bitsAllocated uint16
		expected      int
	}{
		{1, 1},
		{8, 1},
		{16, 2},
		{24, 3},
		{32, 4},
	}

	for _, tc := range testCases {
		bd := NewBitDepth(tc.bitsAllocated, tc.bitsAllocated, tc.bitsAllocated-1, false)
		result := bd.BytesAllocated()
		if result != tc.expected {
			t.Errorf("BytesAllocated for %d bits: expected %d, got %d",
				tc.bitsAllocated, tc.expected, result)
		}
	}
}

func TestBitDepth_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		bd       *BitDepth
		expected bool
	}{
		{
			name:     "valid 8-bit",
			bd:       NewBitDepth(8, 8, 7, false),
			expected: true,
		},
		{
			name:     "valid 16-bit",
			bd:       NewBitDepth(16, 12, 11, true),
			expected: true,
		},
		{
			name:     "invalid: BitsStored > BitsAllocated",
			bd:       NewBitDepth(8, 12, 11, false),
			expected: false,
		},
		{
			name:     "invalid: HighBit >= BitsAllocated",
			bd:       NewBitDepth(8, 8, 8, false),
			expected: false,
		},
		{
			name:     "invalid: HighBit < BitsStored-1",
			bd:       NewBitDepth(16, 12, 10, false),
			expected: false,
		},
		{
			name:     "invalid: uncommon BitsAllocated",
			bd:       NewBitDepth(7, 7, 6, false),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bd.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestBitDepth_Mask(t *testing.T) {
	testCases := []struct {
		bitsStored uint16
		expected   uint32
	}{
		{8, 0xFF},
		{12, 0xFFF},
		{16, 0xFFFF},
		{32, 0xFFFFFFFF},
	}

	for _, tc := range testCases {
		bd := NewBitDepth(16, tc.bitsStored, tc.bitsStored-1, false)
		result := bd.Mask()
		if result != tc.expected {
			t.Errorf("Mask for %d bits: expected 0x%X, got 0x%X",
				tc.bitsStored, tc.expected, result)
		}
	}
}

func TestBitDepth_SignMask(t *testing.T) {
	testCases := []struct {
		bitsStored uint16
		isSigned   bool
		expected   uint32
	}{
		{8, true, 0x80},
		{12, true, 0x800},
		{16, true, 0x8000},
		{8, false, 0}, // Unsigned should return 0
	}

	for _, tc := range testCases {
		bd := NewBitDepth(16, tc.bitsStored, tc.bitsStored-1, tc.isSigned)
		result := bd.SignMask()
		if result != tc.expected {
			t.Errorf("SignMask for %d bits (signed=%v): expected 0x%X, got 0x%X",
				tc.bitsStored, tc.isSigned, tc.expected, result)
		}
	}
}

func TestBitDepth_ExtendSign(t *testing.T) {
	testCases := []struct {
		name     string
		bd       *BitDepth
		value    uint32
		expected int32
	}{
		{
			name:     "positive 12-bit signed",
			bd:       NewBitDepth(16, 12, 11, true),
			value:    0x7FF, // 2047
			expected: 2047,
		},
		{
			name:     "negative 12-bit signed",
			bd:       NewBitDepth(16, 12, 11, true),
			value:    0x800, // -2048 in 12-bit
			expected: -2048,
		},
		{
			name:     "unsigned value",
			bd:       NewBitDepth(16, 12, 11, false),
			value:    0xFFF, // 4095
			expected: 4095,
		},
		{
			name:     "8-bit signed positive",
			bd:       NewBitDepth(8, 8, 7, true),
			value:    0x7F, // 127
			expected: 127,
		},
		{
			name:     "8-bit signed negative",
			bd:       NewBitDepth(8, 8, 7, true),
			value:    0x80, // -128 in 8-bit
			expected: -128,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bd.ExtendSign(tc.value)
			if result != tc.expected {
				t.Errorf("ExtendSign(0x%X): expected %d, got %d",
					tc.value, tc.expected, result)
			}
		})
	}
}

func TestBitDepth_Range(t *testing.T) {
	bd := NewBitDepth(16, 12, 11, true)
	minimumValue, maximumValue := bd.Range()

	if minimumValue != -2048 {
		t.Errorf("Range min: expected -2048, got %f", minimumValue)
	}
	if maximumValue != 2047 {
		t.Errorf("Range max: expected 2047, got %f", maximumValue)
	}
}
