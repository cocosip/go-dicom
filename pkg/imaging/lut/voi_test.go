// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"
	"testing"
)

func TestNewVOILinearLUT(t *testing.T) {
	lut := NewVOILinearLUT(128.0, 256.0)

	if lut == nil {
		t.Fatal("NewVOILinearLUT returned nil")
	}

	if lut.WindowCenter() != 128.0 {
		t.Errorf("Expected WindowCenter=128.0, got %f", lut.WindowCenter())
	}

	if lut.WindowWidth() != 256.0 {
		t.Errorf("Expected WindowWidth=256.0, got %f", lut.WindowWidth())
	}
}

func TestVOILinearLUT_Transform(t *testing.T) {
	testCases := []struct {
		name         string
		windowCenter float64
		windowWidth  float64
		input        float64
		expectedMin  float64
		expectedMax  float64
	}{
		{
			name:         "center of window",
			windowCenter: 128.0,
			windowWidth:  256.0,
			input:        128.0,
			expectedMin:  120.0,
			expectedMax:  135.0,
		},
		{
			name:         "below window - should clamp to 0",
			windowCenter: 128.0,
			windowWidth:  256.0,
			input:        -100.0,
			expectedMin:  0.0,
			expectedMax:  0.0,
		},
		{
			name:         "above window - should clamp to 255",
			windowCenter: 128.0,
			windowWidth:  256.0,
			input:        500.0,
			expectedMin:  255.0,
			expectedMax:  255.0,
		},
		{
			name:         "narrow window",
			windowCenter: 100.0,
			windowWidth:  50.0,
			input:        100.0,
			expectedMin:  120.0,
			expectedMax:  135.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lut := NewVOILinearLUT(tc.windowCenter, tc.windowWidth)
			result := lut.Transform(tc.input)

			if result < tc.expectedMin || result > tc.expectedMax {
				t.Errorf("Transform(%f): expected range [%f, %f], got %f",
					tc.input, tc.expectedMin, tc.expectedMax, result)
			}

			// Should always be in [0, 255]
			if result < 0.0 || result > 255.0 {
				t.Errorf("Result %f out of valid range [0, 255]", result)
			}
		})
	}
}

func TestVOILinearLUT_WindowWidth1(t *testing.T) {
	// Special case: window width of 1 creates a step function
	lut := NewVOILinearLUT(128.0, 1.0)

	// Below center: should be 0
	result := lut.Transform(127.0)
	if result != 0.0 {
		t.Errorf("Expected 0.0 for value below center, got %f", result)
	}

	// At or above center: should be 255
	result = lut.Transform(128.0)
	if result != 255.0 {
		t.Errorf("Expected 255.0 for value at center, got %f", result)
	}

	result = lut.Transform(129.0)
	if result != 255.0 {
		t.Errorf("Expected 255.0 for value above center, got %f", result)
	}
}

func TestVOILinearLUT_SetWindow(t *testing.T) {
	lut := NewVOILinearLUT(100.0, 200.0)

	// Change window parameters
	lut.SetWindow(150.0, 100.0)

	if lut.WindowCenter() != 150.0 {
		t.Errorf("SetWindow failed: expected center=150.0, got %f", lut.WindowCenter())
	}

	if lut.WindowWidth() != 100.0 {
		t.Errorf("SetWindow failed: expected width=100.0, got %f", lut.WindowWidth())
	}
}

func TestNewVOILinearExactLUT(t *testing.T) {
	lut := NewVOILinearExactLUT(128.0, 256.0)

	if lut == nil {
		t.Fatal("NewVOILinearExactLUT returned nil")
	}

	if lut.WindowCenter() != 128.0 {
		t.Errorf("Expected WindowCenter=128.0, got %f", lut.WindowCenter())
	}
}

func TestVOILinearExactLUT_Transform(t *testing.T) {
	lut := NewVOILinearExactLUT(128.0, 256.0)

	// Test that it produces values in valid range
	result := lut.Transform(128.0)
	if result < 0.0 || result > 255.0 {
		t.Errorf("Transform result %f out of valid range [0, 255]", result)
	}

	// Test clamping at extremes
	result = lut.Transform(-1000.0)
	if result != 0.0 {
		t.Errorf("Expected clamping to 0.0, got %f", result)
	}

	result = lut.Transform(10000.0)
	if result != 255.0 {
		t.Errorf("Expected clamping to 255.0, got %f", result)
	}
}

func TestNewVOISigmoidLUT(t *testing.T) {
	lut := NewVOISigmoidLUT(128.0, 256.0)

	if lut == nil {
		t.Fatal("NewVOISigmoidLUT returned nil")
	}
}

func TestVOISigmoidLUT_Transform(t *testing.T) {
	lut := NewVOISigmoidLUT(128.0, 256.0)

	testCases := []struct {
		name  string
		input float64
	}{
		{"below center", 50.0},
		{"at center", 128.0},
		{"above center", 200.0},
		{"far below", -100.0},
		{"far above", 500.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lut.Transform(tc.input)

			// Sigmoid should produce smooth values in [0, 255]
			if result < 0.0 || result > 255.0 {
				t.Errorf("Transform(%f): result %f out of range [0, 255]", tc.input, result)
			}
		})
	}

	// Test sigmoid property: center point should be around 127.5
	center := lut.Transform(128.0)
	if math.Abs(center-127.5) > 10.0 {
		t.Errorf("Sigmoid at center should be near 127.5, got %f", center)
	}
}

func TestCreateVOILUT(t *testing.T) {
	testCases := []struct {
		name         string
		function     VOILUTFunction
		expectedType string
	}{
		{
			name:         "LINEAR",
			function:     VOILUTFunctionLinear,
			expectedType: "*lut.VOILinearLUT",
		},
		{
			name:         "LINEAR_EXACT",
			function:     VOILUTFunctionLinearExact,
			expectedType: "*lut.VOILinearExactLUT",
		},
		{
			name:         "SIGMOID",
			function:     VOILUTFunctionSigmoid,
			expectedType: "*lut.VOISigmoidLUT",
		},
		{
			name:         "default to LINEAR",
			function:     "UNKNOWN",
			expectedType: "*lut.VOILinearLUT",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lut := CreateVOILUT(tc.function, 128.0, 256.0)

			if lut == nil {
				t.Fatal("CreateVOILUT returned nil")
			}

			// Test that it works
			result := lut.Transform(128.0)
			if result < 0.0 || result > 255.0 {
				t.Errorf("Transform result %f out of valid range", result)
			}
		})
	}
}

func TestVOILUT_OutputRange(t *testing.T) {
	lut := NewVOILinearLUT(128.0, 256.0)

	if lut.MinimumOutputValue() != 0.0 {
		t.Errorf("Expected MinimumOutputValue=0.0, got %f", lut.MinimumOutputValue())
	}

	if lut.MaximumOutputValue() != 255.0 {
		t.Errorf("Expected MaximumOutputValue=255.0, got %f", lut.MaximumOutputValue())
	}
}

func TestVOILUT_IsValid(t *testing.T) {
	lut := NewVOILinearLUT(128.0, 256.0)

	// VOI LUTs always return false for IsValid (always recalculate)
	if lut.IsValid() {
		t.Error("VOI LUT IsValid should return false (always recalculate)")
	}
}
