// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

// PrecalculatedLUT is a performance optimization that pre-computes
// all LUT values for a given input range.
//
// Instead of applying the LUT transformation for each pixel,
// this calculates all possible outputs once and stores them in a table.
// Subsequent lookups are O(1) array accesses.
//
// This is particularly useful when:
// - Applying complex LUT pipelines to large images
// - The input value range is known and limited (e.g., 12-bit or 16-bit)
// - The same LUT will be applied to many pixels
//
// Based on fo-dicom PrecalculatedLUT
type PrecalculatedLUT struct {
	sourceLUT LUT
	minValue  int
	maxValue  int
	offset    int
	table     []int
}

// NewPrecalculatedLUT creates a new precalculated LUT wrapping the source LUT.
//
// Parameters:
//   - sourceLUT: The LUT to precalculate
//   - minValue: Minimum input value to precalculate
//   - maxValue: Maximum input value to precalculate
//
// The table will have (maxValue - minValue + 1) entries.
// For example, for 12-bit data: minValue=0, maxValue=4095 → 4096 entries
func NewPrecalculatedLUT(sourceLUT LUT, minValue, maxValue int) *PrecalculatedLUT {
	p := &PrecalculatedLUT{
		sourceLUT: sourceLUT,
		minValue:  minValue,
		maxValue:  maxValue,
		offset:    -minValue,
		table:     make([]int, maxValue-minValue+1),
	}

	// Pre-calculate all values immediately
	p.Recalculate()

	return p
}

// IsValid returns the validity of the source LUT
func (p *PrecalculatedLUT) IsValid() bool {
	if p.sourceLUT == nil {
		return false
	}
	return p.sourceLUT.IsValid()
}

// MinimumOutputValue returns the minimum output value from the source LUT
func (p *PrecalculatedLUT) MinimumOutputValue() float64 {
	if p.sourceLUT == nil {
		return 0
	}
	return p.sourceLUT.MinimumOutputValue()
}

// MaximumOutputValue returns the maximum output value from the source LUT
func (p *PrecalculatedLUT) MaximumOutputValue() float64 {
	if p.sourceLUT == nil {
		return 0
	}
	return p.sourceLUT.MaximumOutputValue()
}

// Transform performs a fast table lookup.
//
// Input values outside [minValue, maxValue] are clamped to the
// nearest table entry.
func (p *PrecalculatedLUT) Transform(input float64) float64 {
	index := int(input) + p.offset

	// Clamp to table bounds
	if index < 0 {
		return float64(p.table[0])
	}
	if index >= len(p.table) {
		return float64(p.table[len(p.table)-1])
	}

	return float64(p.table[index])
}

// Recalculate rebuilds the precalculated table.
//
// This should be called if the source LUT parameters change
// (e.g., window/level adjusted for a VOILUT).
func (p *PrecalculatedLUT) Recalculate() {
	if p.sourceLUT == nil {
		return
	}

	// Only recalculate if source LUT is invalid (needs recalc)
	if p.IsValid() {
		return
	}

	// Force source LUT to recalculate
	p.sourceLUT.Recalculate()

	// Pre-calculate all values in range
	for i := p.minValue; i <= p.maxValue; i++ {
		tableIndex := i + p.offset
		p.table[tableIndex] = int(p.sourceLUT.Transform(float64(i)))
	}
}

// TableSize returns the size of the precalculated table
func (p *PrecalculatedLUT) TableSize() int {
	return len(p.table)
}

// MinValue returns the minimum input value in the table
func (p *PrecalculatedLUT) MinValue() int {
	return p.minValue
}

// MaxValue returns the maximum input value in the table
func (p *PrecalculatedLUT) MaxValue() int {
	return p.maxValue
}
