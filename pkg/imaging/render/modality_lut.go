// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import "math"

// ModalityRescaleLUT implements ModalityLUT using rescale slope and intercept
type ModalityRescaleLUT struct {
	rescaleSlope     float64
	rescaleIntercept float64
	minOutputValue   float64
	maxOutputValue   float64
	minInputValue    float64
	maxInputValue    float64
}

// NewModalityRescaleLUT creates a new ModalityRescaleLUT with slope, intercept, and input range
func NewModalityRescaleLUT(slope, intercept, minInput, maxInput float64) *ModalityRescaleLUT {
	lut := &ModalityRescaleLUT{
		rescaleSlope:     slope,
		rescaleIntercept: intercept,
		minInputValue:    minInput,
		maxInputValue:    maxInput,
	}
	lut.minOutputValue = lut.Transform(minInput)
	lut.maxOutputValue = lut.Transform(maxInput)
	return lut
}

// RescaleSlope returns the modality rescale slope
func (m *ModalityRescaleLUT) RescaleSlope() float64 {
	return m.rescaleSlope
}

// RescaleIntercept returns the modality rescale intercept
func (m *ModalityRescaleLUT) RescaleIntercept() float64 {
	return m.rescaleIntercept
}

// IsValid returns true (always valid)
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

// Transform transforms input value using the formula: output = input * slope + intercept
func (m *ModalityRescaleLUT) Transform(value float64) float64 {
	return (value * m.rescaleSlope) + m.rescaleIntercept
}

// Recalculate forces the recalculation of LUT (no-op for rescale LUT)
func (m *ModalityRescaleLUT) Recalculate() {
	// No recalculation needed for rescale LUT
}

// ModalitySequenceLUT implements ModalityLUT using a LUT sequence
type ModalitySequenceLUT struct {
	lut              []float64
	firstValueMapped int
	isSigned         bool
	minOutputValue   float64
	maxOutputValue   float64
}

// NewModalitySequenceLUT creates a new ModalitySequenceLUT
func NewModalitySequenceLUT(lut []float64, firstValueMapped int, isSigned bool) *ModalitySequenceLUT {
	if len(lut) == 0 {
		return nil
	}

	minVal := lut[0]
	maxVal := lut[0]
	for _, v := range lut {
		minVal = math.Min(minVal, v)
		maxVal = math.Max(maxVal, v)
	}

	return &ModalitySequenceLUT{
		lut:              lut,
		firstValueMapped: firstValueMapped,
		isSigned:         isSigned,
		minOutputValue:   minVal,
		maxOutputValue:   maxVal,
	}
}

// IsValid returns true if the LUT is valid
func (m *ModalitySequenceLUT) IsValid() bool {
	return len(m.lut) > 0
}

// MinimumOutputValue returns the minimum output value
func (m *ModalitySequenceLUT) MinimumOutputValue() float64 {
	return m.minOutputValue
}

// MaximumOutputValue returns the maximum output value
func (m *ModalitySequenceLUT) MaximumOutputValue() float64 {
	return m.maxOutputValue
}

// Transform transforms input value using the LUT sequence
func (m *ModalitySequenceLUT) Transform(value float64) float64 {
	index := int(value) - m.firstValueMapped
	if index < 0 {
		return m.lut[0]
	}
	if index >= len(m.lut) {
		return m.lut[len(m.lut)-1]
	}
	return m.lut[index]
}

// Recalculate forces the recalculation of LUT
func (m *ModalitySequenceLUT) Recalculate() {
	if len(m.lut) == 0 {
		return
	}

	minVal := m.lut[0]
	maxVal := m.lut[0]
	for _, v := range m.lut {
		minVal = math.Min(minVal, v)
		maxVal = math.Max(maxVal, v)
	}
	m.minOutputValue = minVal
	m.maxOutputValue = maxVal
}
