// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import "math"

// ModalitySequenceLUT maps stored values through a DICOM Modality LUT Sequence.
type ModalitySequenceLUT struct {
	values           []float64
	firstValueMapped int
	isSigned         bool
	minOutputValue   float64
	maxOutputValue   float64
}

// NewModalitySequenceLUT creates a Modality Sequence LUT with an owned copy of values.
func NewModalitySequenceLUT(values []float64, firstValueMapped int, isSigned bool) *ModalitySequenceLUT {
	if len(values) == 0 {
		return nil
	}

	owned := append([]float64(nil), values...)
	result := &ModalitySequenceLUT{
		values:           owned,
		firstValueMapped: firstValueMapped,
		isSigned:         isSigned,
	}
	result.Recalculate()
	return result
}

// IsValid reports whether the LUT contains at least one mapped value.
func (m *ModalitySequenceLUT) IsValid() bool {
	return m != nil && len(m.values) > 0
}

// MinimumOutputValue returns the smallest mapped output value.
func (m *ModalitySequenceLUT) MinimumOutputValue() float64 {
	return m.minOutputValue
}

// MaximumOutputValue returns the largest mapped output value.
func (m *ModalitySequenceLUT) MaximumOutputValue() float64 {
	return m.maxOutputValue
}

// Transform maps a stored value and clamps values outside the LUT range.
func (m *ModalitySequenceLUT) Transform(value float64) float64 {
	index := int(value) - m.firstValueMapped
	if index < 0 {
		return m.values[0]
	}
	if index >= len(m.values) {
		return m.values[len(m.values)-1]
	}
	return m.values[index]
}

// Recalculate refreshes the cached output range from the mapped values.
func (m *ModalitySequenceLUT) Recalculate() {
	if len(m.values) == 0 {
		return
	}

	minimum, maximum := m.values[0], m.values[0]
	for _, value := range m.values[1:] {
		minimum = math.Min(minimum, value)
		maximum = math.Max(maximum, value)
	}
	m.minOutputValue = minimum
	m.maxOutputValue = maximum
}

// IsSigned reports how the LUT descriptor's first mapped value was interpreted.
func (m *ModalitySequenceLUT) IsSigned() bool {
	return m.isSigned
}
