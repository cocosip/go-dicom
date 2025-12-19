// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"
)

// ModalityRescaleLUT implements a modality LUT using Rescale Slope and Intercept.
//
// Transforms pixel values using the formula:
//
//	output = input * RescaleSlope + RescaleIntercept
//
// This is used to convert stored pixel values to modality-specific units
// (e.g., Hounsfield Units for CT).
//
// Reference: DICOM Part 3, Section C.11.1 (Modality LUT Module)
// Based on fo-dicom ModalityRescaleLUT
type ModalityRescaleLUT struct {
	rescaleSlope     float64
	rescaleIntercept float64
	minOutputValue   float64
	maxOutputValue   float64
}

// NewModalityRescaleLUT creates a new Modality Rescale LUT.
//
// Parameters:
//   - rescaleSlope: Modality Rescale Slope (0028,1053)
//   - rescaleIntercept: Modality Rescale Intercept (0028,1052)
//   - minInput, maxInput: optional input range used to pre-compute output min/max (pass NaN to skip)
func NewModalityRescaleLUT(rescaleSlope, rescaleIntercept float64, minInput, maxInput float64) *ModalityRescaleLUT {
	lut := &ModalityRescaleLUT{
		rescaleSlope:     rescaleSlope,
		rescaleIntercept: rescaleIntercept,
	}

	// Calculate output range based on bit depth
	if !math.IsNaN(minInput) && !math.IsNaN(maxInput) {
		lut.minOutputValue = lut.Transform(minInput)
		lut.maxOutputValue = lut.Transform(maxInput)
	}

	return lut
}

// IsValid returns true (modality rescale LUT is always valid)
func (m *ModalityRescaleLUT) IsValid() bool {
	return true
}

// MinimumOutputValue returns the minimum output value
func (m *ModalityRescaleLUT) MinimumOutputValue() float64 {
	return m.minOutputValue
}

// MaximumOutputValue returns the maximum output value
func (m *ModalityRescaleLUT) MaximumOutputValue() float64 {
	return m.maxOutputValue
}

// Transform applies the rescale transformation: output = input * slope + intercept
func (m *ModalityRescaleLUT) Transform(input float64) float64 {
	return (input * m.rescaleSlope) + m.rescaleIntercept
}

// Recalculate is a no-op for ModalityRescaleLUT (values don't change)
func (m *ModalityRescaleLUT) Recalculate() {
	// No recalculation needed
}

// RescaleSlope returns the rescale slope
func (m *ModalityRescaleLUT) RescaleSlope() float64 {
	return m.rescaleSlope
}

// RescaleIntercept returns the rescale intercept
func (m *ModalityRescaleLUT) RescaleIntercept() float64 {
	return m.rescaleIntercept
}
