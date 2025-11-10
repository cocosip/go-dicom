// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"math"
	"testing"
)

func TestModalityRescaleLUT(t *testing.T) {
	// Test rescale with slope=2, intercept=10
	lut := NewModalityRescaleLUT(2.0, 10.0, 0, 100)

	if lut.RescaleSlope() != 2.0 {
		t.Errorf("Expected slope 2.0, got %f", lut.RescaleSlope())
	}

	if lut.RescaleIntercept() != 10.0 {
		t.Errorf("Expected intercept 10.0, got %f", lut.RescaleIntercept())
	}

	// Test transformation: output = input * 2 + 10
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 10},
		{5, 20},
		{10, 30},
		{50, 110},
	}

	for _, tc := range testCases {
		result := lut.Transform(tc.input)
		if result != tc.expected {
			t.Errorf("Transform(%f): expected %f, got %f", tc.input, tc.expected, result)
		}
	}

	if !lut.IsValid() {
		t.Error("Expected LUT to be valid")
	}
}

func TestModalitySequenceLUT(t *testing.T) {
	// Create a simple LUT: [10, 20, 30, 40, 50]
	lutData := []float64{10, 20, 30, 40, 50}
	lut := NewModalitySequenceLUT(lutData, 0, false)

	if lut == nil {
		t.Fatal("Failed to create ModalitySequenceLUT")
	}

	if !lut.IsValid() {
		t.Error("Expected LUT to be valid")
	}

	// Test transformations
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 10},  // First value
		{2, 30},  // Middle value
		{4, 50},  // Last value
		{-1, 10}, // Below range, should return first
		{10, 50}, // Above range, should return last
	}

	for _, tc := range testCases {
		result := lut.Transform(tc.input)
		if result != tc.expected {
			t.Errorf("Transform(%f): expected %f, got %f", tc.input, tc.expected, result)
		}
	}

	// Check min/max output values
	if lut.MinimumOutputValue() != 10 {
		t.Errorf("Expected min output 10, got %f", lut.MinimumOutputValue())
	}
	if lut.MaximumOutputValue() != 50 {
		t.Errorf("Expected max output 50, got %f", lut.MaximumOutputValue())
	}
}

func TestLinearVOILUT(t *testing.T) {
	// Window center = 128, window width = 256
	lut := NewLinearVOILUT(128, 256)

	if lut.WindowCenter() != 128 {
		t.Errorf("Expected center 128, got %f", lut.WindowCenter())
	}

	if lut.WindowWidth() != 256 {
		t.Errorf("Expected width 256, got %f", lut.WindowWidth())
	}

	// Test some transformations
	// For center=128, width=256: values from 0 to 256 should map to 0-255
	testCases := []struct {
		input       float64
		shouldClamp bool
	}{
		{0, true},    // Should be clamped to 0
		{128, false}, // Middle should be around 127
		{255, true},  // Should be clamped to 255
	}

	for _, tc := range testCases {
		result := lut.Transform(tc.input)
		if tc.shouldClamp {
			if result < 0 || result > 255 {
				t.Errorf("Transform(%f): expected clamped value, got %f", tc.input, result)
			}
		}
	}

	// Test window width = 1
	lut.SetWindow(128, 1)
	below := lut.Transform(127)
	above := lut.Transform(128)

	if below != 0 {
		t.Errorf("For width=1, value below center should be 0, got %f", below)
	}
	if above != 255 {
		t.Errorf("For width=1, value at/above center should be 255, got %f", above)
	}
}

func TestLinearExactVOILUT(t *testing.T) {
	// Window center = 128, window width = 0.5 (test LINEAR_EXACT)
	lut := NewLinearExactVOILUT(128, 0.5)

	if lut.WindowCenter() != 128 {
		t.Errorf("Expected center 128, got %f", lut.WindowCenter())
	}

	if lut.WindowWidth() != 0.5 {
		t.Errorf("Expected width 0.5, got %f", lut.WindowWidth())
	}

	// LINEAR_EXACT formula: ((value - center) / width + 0.5) * 255
	// For center=128, width=0.5:
	// value=128 -> ((128-128)/0.5 + 0.5) * 255 = 0.5 * 255 = 127.5
	result := lut.Transform(128)
	expected := 127.5
	if math.Abs(result-expected) > 0.1 {
		t.Errorf("Transform(128): expected ~%f, got %f", expected, result)
	}
}

func TestSigmoidVOILUT(t *testing.T) {
	// Window center = 128, window width = 64
	lut := NewSigmoidVOILUT(128, 64)

	// Sigmoid should produce values in range [0, 255]
	// At center, output should be around 127.5
	resultCenter := lut.Transform(128)
	if math.Abs(resultCenter-127.5) > 1.0 {
		t.Errorf("Sigmoid at center: expected ~127.5, got %f", resultCenter)
	}

	// Test extreme values
	resultLow := lut.Transform(0)
	resultHigh := lut.Transform(255)

	if resultLow >= resultCenter {
		t.Errorf("Sigmoid should be monotonic: low=%f should be < center=%f", resultLow, resultCenter)
	}
	if resultHigh <= resultCenter {
		t.Errorf("Sigmoid should be monotonic: high=%f should be > center=%f", resultHigh, resultCenter)
	}
}

func TestCreateVOILUT(t *testing.T) {
	// Test LINEAR
	lut1 := CreateVOILUT(VOILinear, 128, 256)
	if _, ok := lut1.(*LinearVOILUT); !ok {
		t.Error("Expected LinearVOILUT for LINEAR function")
	}

	// Test LINEAR_EXACT
	lut2 := CreateVOILUT(VOILinearExact, 128, 256)
	if _, ok := lut2.(*LinearExactVOILUT); !ok {
		t.Error("Expected LinearExactVOILUT for LINEAR_EXACT function")
	}

	// Test SIGMOID
	lut3 := CreateVOILUT(VOISigmoid, 128, 64)
	if _, ok := lut3.(*SigmoidVOILUT); !ok {
		t.Error("Expected SigmoidVOILUT for SIGMOID function")
	}

	// Test automatic LINEAR_EXACT for width < 1.0
	lut4 := CreateVOILUT(VOILinear, 128, 0.5)
	if _, ok := lut4.(*LinearExactVOILUT); !ok {
		t.Error("Expected automatic LINEAR_EXACT when width < 1.0")
	}
}

func TestCompositeLUT(t *testing.T) {
	// Create a chain: Modality (slope=2, intercept=0) -> VOI (center=100, width=200)
	modalityLUT := NewModalityRescaleLUT(2.0, 0, 0, 100)
	voiLUT := NewLinearVOILUT(100, 200)

	compositeLUT := NewCompositeLUT(modalityLUT, voiLUT)

	// Test: input=50 -> modality: 50*2+0=100 -> VOI: should map to middle range
	result := compositeLUT.Transform(50)
	if result < 0 || result > 255 {
		t.Errorf("Composite transform: expected value in [0,255], got %f", result)
	}

	// Test min/max
	if compositeLUT.MinimumOutputValue() != 0 {
		t.Errorf("Expected min output 0, got %f", compositeLUT.MinimumOutputValue())
	}
	if compositeLUT.MaximumOutputValue() != 255 {
		t.Errorf("Expected max output 255, got %f", compositeLUT.MaximumOutputValue())
	}
}

func TestInvertLUT(t *testing.T) {
	lut := NewInvertLUT(0, 255)

	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 255},
		{255, 0},
		{128, 127},
	}

	for _, tc := range testCases {
		result := lut.Transform(tc.input)
		if result != tc.expected {
			t.Errorf("Invert(%f): expected %f, got %f", tc.input, tc.expected, result)
		}
	}

	if !lut.IsValid() {
		t.Error("Expected InvertLUT to be valid")
	}
}
