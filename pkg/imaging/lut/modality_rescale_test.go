// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging"
)

func TestNewModalityRescaleLUT(t *testing.T) {
	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    12,
		HighBit:       11,
		IsSigned:      false,
	}

	lut := NewModalityRescaleLUT(1.0, 0.0, bitDepth)

	if lut == nil {
		t.Fatal("NewModalityRescaleLUT returned nil")
	}

	if !lut.IsValid() {
		t.Error("ModalityRescaleLUT should be valid")
	}

	if lut.RescaleSlope() != 1.0 {
		t.Errorf("Expected slope=1.0, got %f", lut.RescaleSlope())
	}

	if lut.RescaleIntercept() != 0.0 {
		t.Errorf("Expected intercept=0.0, got %f", lut.RescaleIntercept())
	}
}

func TestModalityRescaleLUT_Transform(t *testing.T) {
	testCases := []struct {
		name       string
		slope      float64
		intercept  float64
		input      float64
		expected   float64
	}{
		{
			name:       "identity transform (slope=1, intercept=0)",
			slope:      1.0,
			intercept:  0.0,
			input:      100.0,
			expected:   100.0,
		},
		{
			name:       "CT Hounsfield Units (typical)",
			slope:      1.0,
			intercept:  -1024.0,
			input:      1024.0,
			expected:   0.0, // Water
		},
		{
			name:       "positive slope and intercept",
			slope:      2.0,
			intercept:  10.0,
			input:      50.0,
			expected:   110.0, // 50 * 2 + 10
		},
		{
			name:       "negative intercept",
			slope:      1.5,
			intercept:  -100.0,
			input:      200.0,
			expected:   200.0, // 200 * 1.5 - 100
		},
		{
			name:       "fractional slope",
			slope:      0.5,
			intercept:  0.0,
			input:      100.0,
			expected:   50.0,
		},
	}

	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    12,
		HighBit:       11,
		IsSigned:      false,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lut := NewModalityRescaleLUT(tc.slope, tc.intercept, bitDepth)
			result := lut.Transform(tc.input)

			if math.Abs(result-tc.expected) > 0.0001 {
				t.Errorf("Transform(%f): expected %f, got %f", tc.input, tc.expected, result)
			}
		})
	}
}

func TestModalityRescaleLUT_OutputRange(t *testing.T) {
	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    12,
		HighBit:       11,
		IsSigned:      false,
	}

	// CT with Hounsfield Units
	lut := NewModalityRescaleLUT(1.0, -1024.0, bitDepth)

	minOut := lut.MinimumOutputValue()
	maxOut := lut.MaximumOutputValue()

	// For unsigned 12-bit data (0-4095) with HU transform:
	// Min: 0 * 1 + (-1024) = -1024
	// Max: 4095 * 1 + (-1024) = 3071
	expectedMin := -1024.0
	expectedMax := 3071.0

	if math.Abs(minOut-expectedMin) > 0.0001 {
		t.Errorf("Expected MinimumOutputValue=%f, got %f", expectedMin, minOut)
	}

	if math.Abs(maxOut-expectedMax) > 0.0001 {
		t.Errorf("Expected MaximumOutputValue=%f, got %f", expectedMax, maxOut)
	}
}

func TestModalityRescaleLUT_Recalculate(t *testing.T) {
	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    12,
		HighBit:       11,
		IsSigned:      false,
	}

	lut := NewModalityRescaleLUT(1.0, 0.0, bitDepth)

	// Recalculate should be a no-op
	beforeMin := lut.MinimumOutputValue()
	beforeMax := lut.MaximumOutputValue()

	lut.Recalculate()

	if lut.MinimumOutputValue() != beforeMin {
		t.Error("Recalculate changed MinimumOutputValue")
	}

	if lut.MaximumOutputValue() != beforeMax {
		t.Error("Recalculate changed MaximumOutputValue")
	}
}

func TestModalityRescaleLUT_SignedData(t *testing.T) {
	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    16,
		HighBit:       15,
		IsSigned:      true,
	}

	lut := NewModalityRescaleLUT(1.0, 0.0, bitDepth)

	// Test negative values
	result := lut.Transform(-1000.0)
	if result != -1000.0 {
		t.Errorf("Expected -1000.0, got %f", result)
	}

	// Check output range includes negative values
	minOut := lut.MinimumOutputValue()
	if minOut >= 0 {
		t.Error("Signed data should have negative minimum output value")
	}
}
