// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

// CompositeLUT chains multiple LUTs together into a pipeline.
//
// When transforming a value, it passes through each LUT in sequence:
//
//	output = LUT_n(...LUT_2(LUT_1(input)))
//
// This is commonly used to chain:
// 1. ModalityRescaleLUT (stored values → modality values)
// 2. VOILUT (modality values → presentation values with windowing)
// 3. InvertLUT (optional inversion)
// 4. OutputLUT (grayscale → RGB/color)
//
// Based on fo-dicom CompositeLUT
type CompositeLUT struct {
	luts []LUT
}

// NewCompositeLUT creates a new empty composite LUT
func NewCompositeLUT() *CompositeLUT {
	return &CompositeLUT{
		luts: make([]LUT, 0),
	}
}

// Add appends a LUT to the pipeline
func (c *CompositeLUT) Add(lut LUT) {
	if lut != nil {
		c.luts = append(c.luts, lut)
	}
}

// IsValid returns true if all constituent LUTs are valid
func (c *CompositeLUT) IsValid() bool {
	for _, lut := range c.luts {
		if !lut.IsValid() {
			return false
		}
	}
	return true
}

// MinimumOutputValue returns the minimum output value from the final LUT
func (c *CompositeLUT) MinimumOutputValue() float64 {
	if len(c.luts) == 0 {
		return 0
	}
	return c.luts[len(c.luts)-1].MinimumOutputValue()
}

// MaximumOutputValue returns the maximum output value from the final LUT
func (c *CompositeLUT) MaximumOutputValue() float64 {
	if len(c.luts) == 0 {
		return 0
	}
	return c.luts[len(c.luts)-1].MaximumOutputValue()
}

// Transform applies all LUTs in sequence
func (c *CompositeLUT) Transform(input float64) float64 {
	value := input
	for _, lut := range c.luts {
		value = lut.Transform(value)
	}
	return value
}

// Recalculate forces recalculation of all constituent LUTs
func (c *CompositeLUT) Recalculate() {
	for _, lut := range c.luts {
		lut.Recalculate()
	}
}

// FinalLUT returns the last LUT in the pipeline, or nil if empty
func (c *CompositeLUT) FinalLUT() LUT {
	if len(c.luts) == 0 {
		return nil
	}
	return c.luts[len(c.luts)-1]
}

// Count returns the number of LUTs in the pipeline
func (c *CompositeLUT) Count() int {
	return len(c.luts)
}

// GetLUT returns the LUT at the specified index, or nil if out of range
func (c *CompositeLUT) GetLUT(index int) LUT {
	if index < 0 || index >= len(c.luts) {
		return nil
	}
	return c.luts[index]
}
