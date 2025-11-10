// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"
)

func TestNewPresentationLUT(t *testing.T) {
	lut := NewPresentationLUT("1.2.3.4.5")

	if lut.SOPInstanceUID != "1.2.3.4.5" {
		t.Errorf("Expected SOPInstanceUID='1.2.3.4.5', got '%s'", lut.SOPInstanceUID)
	}

	if lut.PresentationLUTShape != PresentationLUTShapeIdentity {
		t.Errorf("Expected PresentationLUTShape=IDENTITY, got '%s'", lut.PresentationLUTShape)
	}
}

func TestPresentationLUT_SetLUT(t *testing.T) {
	lut := NewPresentationLUT("1.2.3.4.5")

	lutData := []uint16{0, 100, 200, 300, 400}
	err := lut.SetLUT(5, 0, 10, lutData)
	if err != nil {
		t.Fatalf("SetLUT failed: %v", err)
	}

	if lut.GetNumberOfEntries() != 5 {
		t.Errorf("Expected 5 entries, got %d", lut.GetNumberOfEntries())
	}

	if lut.GetFirstValueMapped() != 0 {
		t.Errorf("Expected first value mapped=0, got %d", lut.GetFirstValueMapped())
	}

	if lut.GetBitsPerEntry() != 10 {
		t.Errorf("Expected bits per entry=10, got %d", lut.GetBitsPerEntry())
	}

	if len(lut.LUTData) != 5 {
		t.Errorf("Expected LUTData length=5, got %d", len(lut.LUTData))
	}
}

func TestPresentationLUT_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(*PresentationLUT)
		expected bool
	}{
		{
			name: "valid LUT",
			setup: func(lut *PresentationLUT) {
				lutData := []uint16{0, 100, 200, 300, 400}
				lut.SetLUT(5, 0, 12, lutData)
			},
			expected: true,
		},
		{
			name: "invalid: bits per entry too low",
			setup: func(lut *PresentationLUT) {
				lutData := []uint16{0, 100, 200}
				lut.SetLUT(3, 0, 8, lutData) // 8 < 10
			},
			expected: false,
		},
		{
			name: "invalid: bits per entry too high",
			setup: func(lut *PresentationLUT) {
				lutData := []uint16{0, 100, 200}
				lut.SetLUT(3, 0, 20, lutData) // 20 > 16
			},
			expected: false,
		},
		{
			name: "invalid: LUT data size mismatch",
			setup: func(lut *PresentationLUT) {
				lutData := []uint16{0, 100}   // Only 2 entries
				lut.SetLUT(5, 0, 12, lutData) // Claims 5 entries
			},
			expected: false,
		},
		{
			name: "invalid: first value mapped not 0",
			setup: func(lut *PresentationLUT) {
				lutData := []uint16{0, 100, 200}
				lut.SetLUT(3, 10, 12, lutData) // First value = 10
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lut := NewPresentationLUT("1.2.3.4.5")
			tc.setup(lut)

			result := lut.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestPresentationLUT_TransformValue_Identity(t *testing.T) {
	lut := NewPresentationLUT("1.2.3.4.5")
	lut.PresentationLUTShape = PresentationLUTShapeIdentity

	testValues := []uint16{0, 100, 500, 1000, 4095}

	for _, val := range testValues {
		result := lut.TransformValue(val)
		if result != val {
			t.Errorf("IDENTITY transform of %d: expected %d, got %d", val, val, result)
		}
	}
}

func TestPresentationLUT_TransformValue_LUT(t *testing.T) {
	lut := NewPresentationLUT("1.2.3.4.5")
	lut.PresentationLUTShape = PresentationLUTShapeLinOD

	// Create a simple LUT: input 0-4 maps to output 0, 100, 200, 300, 400
	lutData := []uint16{0, 100, 200, 300, 400}
	lut.SetLUT(5, 0, 12, lutData)

	testCases := []struct {
		input    uint16
		expected uint16
	}{
		{0, 0},
		{1, 100},
		{2, 200},
		{3, 300},
		{4, 400},
		{5, 400},   // Clamped to last entry
		{100, 400}, // Clamped to last entry
	}

	for _, tc := range testCases {
		result := lut.TransformValue(tc.input)
		if result != tc.expected {
			t.Errorf("TransformValue(%d): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func TestPresentationLUT_TransformValue_Clamp(t *testing.T) {
	lut := NewPresentationLUT("1.2.3.4.5")
	lut.PresentationLUTShape = PresentationLUTShapeLinOD

	lutData := []uint16{100, 200, 300}
	lut.SetLUT(3, 0, 12, lutData)

	// Test values below range
	result := lut.TransformValue(0)
	if result != 100 {
		t.Errorf("Value below range should clamp to first entry (100), got %d", result)
	}

	// Test values above range
	result = lut.TransformValue(10)
	if result != 300 {
		t.Errorf("Value above range should clamp to last entry (300), got %d", result)
	}
}
