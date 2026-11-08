// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

import (
	"testing"
)

func TestNewPresentationLUT(t *testing.T) {
	lut := NewPresentationLUT(testSOPInstanceUID)

	if lut.SOPInstanceUID != testSOPInstanceUID {
		t.Errorf("Expected SOPInstanceUID='1.2.3.4.5', got '%s'", lut.SOPInstanceUID)
	}

	if lut.PresentationLUTShape != PresentationLUTShapeIdentity {
		t.Errorf("Expected PresentationLUTShape=IDENTITY, got '%s'", lut.PresentationLUTShape)
	}
	if !lut.IsValid() {
		t.Fatal("default IDENTITY Presentation LUT is not valid")
	}
}

func TestPresentationLUT_SetLUT(t *testing.T) {
	lut := NewPresentationLUT(testSOPInstanceUID)

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
	if lut.PresentationLUTShape != "" {
		t.Fatalf("SetLUT() left PresentationLUTShape=%q, want empty sequence mode", lut.PresentationLUTShape)
	}
}

// makeTestLUTData returns a conformant identity-ramp LUT Data slice.
func makeTestLUTData(entries uint16) []uint16 {
	data := make([]uint16, entries)
	for index := range data {
		data[index] = uint16(index)
	}
	return data
}

// setTestLUT configures a conformant print LUT: 256 or 4096 entries with a
// 10-16 bit entry width (PS3.4 H.4.9.2.1.1.1).
func setTestLUT(t *testing.T, lut *PresentationLUT, entries, bitsPerEntry uint16) {
	t.Helper()
	if err := lut.SetLUT(entries, 0, bitsPerEntry, makeTestLUTData(entries)); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}
}

func TestPresentationLUT_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(*testing.T, *PresentationLUT)
		expected bool
	}{
		{
			name:     "valid IDENTITY shape",
			setup:    func(_ *testing.T, _ *PresentationLUT) {},
			expected: true,
		},
		{
			name: "invalid: shape and LUT both present",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				lut.LUTDescriptor = []uint16{256, 0, 12}
				lut.LUTData = makeTestLUTData(256)
			},
			expected: false,
		},
		{
			name: "valid 256-entry LUT",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 256, 12)
			},
			expected: true,
		},
		{
			name: "valid 4096-entry LUT",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 4096, 12)
				lut.BitsStored = 12
			},
			expected: true,
		},
		{
			name: "valid 256-entry LUT for 8-bit images",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 256, 10)
				lut.BitsStored = 8
			},
			expected: true,
		},
		{
			name: "invalid: 256 entries for 12-bit images",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 256, 12)
				lut.BitsStored = 12
			},
			expected: false,
		},
		{
			name: "invalid: 4096 entries for 8-bit images",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 4096, 12)
				lut.BitsStored = 8
			},
			expected: false,
		},
		{
			name: "invalid: 1024 entries",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 1024, 12)
			},
			expected: false,
		},
		{
			name: "invalid: zero entries",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				lut.LUTDescriptor = []uint16{0, 0, 12}
				lut.LUTData = makeTestLUTData(256)
			},
			expected: false,
		},
		{
			name: "invalid: bits per entry too low",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 256, 8) // 8 < 10
			},
			expected: false,
		},
		{
			name: "invalid: bits per entry too high",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				setTestLUT(t, lut, 256, 20) // 20 > 16
			},
			expected: false,
		},
		{
			name: "invalid: LUT data size mismatch",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				lut.LUTDescriptor = []uint16{256, 0, 12}
				lut.LUTData = makeTestLUTData(2) // Intentionally short
			},
			expected: false,
		},
		{
			name: "invalid: first value mapped not 0",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				_ = lut.SetLUT(256, 10, 12, makeTestLUTData(256))
			},
			expected: false,
		},
		{
			name: "invalid: LUT data exceeds declared output range",
			setup: func(_ *testing.T, lut *PresentationLUT) {
				data := makeTestLUTData(256)
				data[255] = 1024
				_ = lut.SetLUT(256, 0, 10, data)
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lut := NewPresentationLUT(testSOPInstanceUID)
			tc.setup(t, lut)

			result := lut.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestPresentationLUT_TransformValue_Identity(t *testing.T) {
	lut := NewPresentationLUT(testSOPInstanceUID)
	lut.PresentationLUTShape = PresentationLUTShapeIdentity

	testValues := []uint16{0, 100, 500, 1000, 4095}

	for _, val := range testValues {
		result := lut.TransformValue(val)
		if result != val {
			t.Errorf("IDENTITY transform of %d: expected %d, got %d", val, val, result)
		}
	}
}

func TestPresentationLUT_InverseShapeIsNotValidForPrint(t *testing.T) {
	lut := NewPresentationLUT(testSOPInstanceUID)
	lut.PresentationLUTShape = PresentationLUTShapeInverse
	// PS3.3 C.11.4 Table C.11-4: INVERSE is not a valid print Presentation LUT
	// shape (only IDENTITY and LIN OD).
	if lut.IsValid() {
		t.Fatal("INVERSE Presentation LUT shape must not be valid for print")
	}
	if got := lut.TransformValue(0); got != ^uint16(0) {
		t.Fatalf("INVERSE transform of zero = %d, want %d", got, ^uint16(0))
	}
	if got := lut.TransformValue(^uint16(0)); got != 0 {
		t.Fatalf("INVERSE transform of max = %d, want 0", got)
	}
}

func TestPresentationLUT_TransformValue_LUT(t *testing.T) {
	lut := NewPresentationLUT(testSOPInstanceUID)

	// Create a simple LUT: input 0-4 maps to output 0, 100, 200, 300, 400
	lutData := []uint16{0, 100, 200, 300, 400}
	if err := lut.SetLUT(5, 0, 12, lutData); err != nil {
		t.Fatalf("SetLUT failed: %v", err)
	}

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
	lut := NewPresentationLUT(testSOPInstanceUID)

	lutData := []uint16{100, 200, 300}
	if err := lut.SetLUT(3, 0, 12, lutData); err != nil {
		t.Fatalf("SetLUT failed: %v", err)
	}

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

