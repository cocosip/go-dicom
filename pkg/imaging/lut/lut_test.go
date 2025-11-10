// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging"
)

func TestInvertLUT(t *testing.T) {
	lut := NewInvertLUT(0.0, 255.0)

	if !lut.IsValid() {
		t.Error("InvertLUT should be valid")
	}

	testCases := []struct {
		input    float64
		expected float64
	}{
		{0.0, 255.0},
		{128.0, 127.0},
		{255.0, 0.0},
		{100.0, 155.0},
	}

	for _, tc := range testCases {
		result := lut.Transform(tc.input)
		if result != tc.expected {
			t.Errorf("Transform(%f): expected %f, got %f", tc.input, tc.expected, result)
		}
	}
}

func TestInvertLUT_OutputRange(t *testing.T) {
	lut := NewInvertLUT(0.0, 255.0)

	if lut.MinimumOutputValue() != 0.0 {
		t.Errorf("Expected MinimumOutputValue=0.0, got %f", lut.MinimumOutputValue())
	}

	if lut.MaximumOutputValue() != 255.0 {
		t.Errorf("Expected MaximumOutputValue=255.0, got %f", lut.MaximumOutputValue())
	}
}

func TestOutputLUT_Grayscale(t *testing.T) {
	lut := NewGrayscaleOutputLUT()

	if !lut.IsValid() {
		t.Error("OutputLUT should be valid")
	}

	// Test grayscale mapping (R=G=B=input)
	testValues := []int{0, 50, 128, 200, 255}

	for _, val := range testValues {
		color := lut.GetColor(val)

		if color.A != 255 {
			t.Errorf("Color at %d: expected A=255, got %d", val, color.A)
		}

		if color.R != uint8(val) || color.G != uint8(val) || color.B != uint8(val) {
			t.Errorf("Color at %d: expected RGB=(%d,%d,%d), got RGB=(%d,%d,%d)",
				val, val, val, val, color.R, color.G, color.B)
		}
	}
}

func TestOutputLUT_CustomColorMap(t *testing.T) {
	// Create a simple hot metal color map (simplified)
	var colorMap [256]imaging.Color32

	for i := 0; i < 256; i++ {
		// Simple gradient: Red -> Yellow -> White
		r := uint8(i)
		g := uint8(0)
		b := uint8(0)

		if i > 128 {
			g = uint8((i - 128) * 2)
		}

		colorMap[i] = imaging.Color32{A: 255, R: r, G: g, B: b}
	}

	lut := NewOutputLUT(colorMap)

	// Test that colors match
	color0 := lut.GetColor(0)
	if color0.R != 0 || color0.G != 0 || color0.B != 0 {
		t.Error("Color at 0 should be black")
	}

	color255 := lut.GetColor(255)
	if color255.R != 255 {
		t.Error("Color at 255 should have R=255")
	}
}

func TestOutputLUT_Clamping(t *testing.T) {
	lut := NewGrayscaleOutputLUT()

	// Test clamping below range
	colorBelow := lut.GetColor(-10)
	if colorBelow.R != 0 {
		t.Error("Values below 0 should clamp to 0")
	}

	// Test clamping above range
	colorAbove := lut.GetColor(300)
	if colorAbove.R != 255 {
		t.Error("Values above 255 should clamp to 255")
	}
}

func TestCompositeLUT(t *testing.T) {
	composite := NewCompositeLUT()

	if composite.Count() != 0 {
		t.Error("New CompositeLUT should have count 0")
	}

	// Add a modality rescale LUT
	bitDepth := &imaging.BitDepth{
		BitsAllocated: 16,
		BitsStored:    12,
		HighBit:       11,
		IsSigned:      false,
	}
	modalityLUT := NewModalityRescaleLUT(1.0, -1024.0, bitDepth)
	composite.Add(modalityLUT)

	if composite.Count() != 1 {
		t.Error("After adding 1 LUT, count should be 1")
	}

	// Add a VOI LUT
	voiLUT := NewVOILinearLUT(0.0, 2000.0)
	composite.Add(voiLUT)

	if composite.Count() != 2 {
		t.Error("After adding 2 LUTs, count should be 2")
	}

	// Test transformation through pipeline
	input := 1024.0
	result := composite.Transform(input)

	// Should apply: 1024 * 1 + (-1024) = 0 (modality)
	// Then VOI windowing on 0
	// Result should be in [0, 255]
	if result < 0.0 || result > 255.0 {
		t.Errorf("Composite transform result %f out of range [0, 255]", result)
	}
}

func TestCompositeLUT_FinalLUT(t *testing.T) {
	composite := NewCompositeLUT()

	// Empty composite
	if composite.FinalLUT() != nil {
		t.Error("Empty composite should have nil FinalLUT")
	}

	// Add LUTs
	lut1 := NewInvertLUT(0.0, 255.0)
	lut2 := NewInvertLUT(0.0, 255.0)

	composite.Add(lut1)
	composite.Add(lut2)

	finalLUT := composite.FinalLUT()
	if finalLUT != lut2 {
		t.Error("FinalLUT should return the last added LUT")
	}
}

func TestCompositeLUT_GetLUT(t *testing.T) {
	composite := NewCompositeLUT()

	lut1 := NewInvertLUT(0.0, 255.0)
	lut2 := NewInvertLUT(0.0, 255.0)

	composite.Add(lut1)
	composite.Add(lut2)

	// Test valid indices
	if composite.GetLUT(0) != lut1 {
		t.Error("GetLUT(0) should return first LUT")
	}

	if composite.GetLUT(1) != lut2 {
		t.Error("GetLUT(1) should return second LUT")
	}

	// Test invalid indices
	if composite.GetLUT(-1) != nil {
		t.Error("GetLUT(-1) should return nil")
	}

	if composite.GetLUT(2) != nil {
		t.Error("GetLUT(2) should return nil")
	}
}

func TestCompositeLUT_IsValid(t *testing.T) {
	composite := NewCompositeLUT()

	// Empty composite is valid
	if !composite.IsValid() {
		t.Error("Empty composite should be valid")
	}

	// Add valid LUT
	lut := NewInvertLUT(0.0, 255.0)
	composite.Add(lut)

	if !composite.IsValid() {
		t.Error("Composite with valid LUTs should be valid")
	}
}

func TestCompositeLUT_OutputRange(t *testing.T) {
	composite := NewCompositeLUT()

	// Add a VOI LUT (output range 0-255)
	voiLUT := NewVOILinearLUT(128.0, 256.0)
	composite.Add(voiLUT)

	// Output range should match final LUT
	if composite.MinimumOutputValue() != 0.0 {
		t.Errorf("Expected MinimumOutputValue=0.0, got %f", composite.MinimumOutputValue())
	}

	if composite.MaximumOutputValue() != 255.0 {
		t.Errorf("Expected MaximumOutputValue=255.0, got %f", composite.MaximumOutputValue())
	}
}

func TestPrecalculatedLUT(t *testing.T) {
	// Create a VOI LUT
	voiLUT := NewVOILinearLUT(2048.0, 4096.0)

	// Precalculate for 12-bit range
	precalc := NewPrecalculatedLUT(voiLUT, 0, 4095)

	if precalc.TableSize() != 4096 {
		t.Errorf("Expected table size=4096, got %d", precalc.TableSize())
	}

	if precalc.MinValue() != 0 {
		t.Errorf("Expected MinValue=0, got %d", precalc.MinValue())
	}

	if precalc.MaxValue() != 4095 {
		t.Errorf("Expected MaxValue=4095, got %d", precalc.MaxValue())
	}

	// Test that precalculated values match source LUT (within tolerance)
	// Note: Precalculated LUT stores int values, so there's rounding
	testValues := []float64{0, 1000, 2048, 3000, 4095}

	for _, val := range testValues {
		precalcResult := precalc.Transform(val)
		directResult := voiLUT.Transform(val)

		// Allow for int conversion difference (should be within 1.0)
		if math.Abs(precalcResult-directResult) > 1.0 {
			t.Errorf("Value %f: precalc=%f, direct=%f (diff too large)", val, precalcResult, directResult)
		}
	}
}

func TestPrecalculatedLUT_Clamping(t *testing.T) {
	voiLUT := NewVOILinearLUT(128.0, 256.0)
	precalc := NewPrecalculatedLUT(voiLUT, 0, 255)

	// Test clamping below range
	result := precalc.Transform(-10.0)
	expected := precalc.Transform(0.0)
	if result != expected {
		t.Error("Values below min should clamp to first table entry")
	}

	// Test clamping above range
	result = precalc.Transform(300.0)
	expected = precalc.Transform(255.0)
	if result != expected {
		t.Error("Values above max should clamp to last table entry")
	}
}

func TestCompositeLUT_AddNil(t *testing.T) {
	composite := NewCompositeLUT()
	composite.Add(nil)

	if composite.Count() != 0 {
		t.Error("Adding nil should not change count")
	}
}
