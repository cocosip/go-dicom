// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

// InvertLUT inverts pixel values by subtracting from the maximum value.
//
// This is used to invert grayscale images (make dark pixels light and vice versa).
//
// Transformation: output = maxValue - input
//
// Based on fo-dicom InvertLUT
type InvertLUT struct {
	minValue float64
	maxValue float64
}

// NewInvertLUT creates a new Invert LUT
//
// Parameters:
//   - minValue: Minimum output value (becomes maximum after inversion)
//   - maxValue: Maximum output value (becomes minimum after inversion)
func NewInvertLUT(minValue, maxValue float64) *InvertLUT {
	return &InvertLUT{
		minValue: minValue,
		maxValue: maxValue,
	}
}

// IsValid returns true (invert LUT is always valid)
func (i *InvertLUT) IsValid() bool {
	return true
}

// MinimumOutputValue returns the minimum output value
func (i *InvertLUT) MinimumOutputValue() float64 {
	return i.minValue
}

// MaximumOutputValue returns the maximum output value
func (i *InvertLUT) MaximumOutputValue() float64 {
	return i.maxValue
}

// Transform inverts the input value: output = maxValue - input
func (i *InvertLUT) Transform(input float64) float64 {
	return i.maxValue - input
}

// Recalculate is a no-op for InvertLUT (values don't change)
func (i *InvertLUT) Recalculate() {
	// No recalculation needed
}
