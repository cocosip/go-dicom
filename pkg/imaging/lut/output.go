// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging"
)

// OutputLUT maps grayscale values to RGB or pseudo-color values.
//
// This is used for:
// - Mapping grayscale to RGB grays (default grayscale display)
// - Applying color maps for pseudo-color visualization (hot metal, rainbow, etc.)
//
// Based on fo-dicom OutputLUT
type OutputLUT struct {
	colorMap [256]imaging.Color32
}

// NewOutputLUT creates a new Output LUT with the specified color map.
//
// The color map must be a 256-entry array mapping grayscale values (0-255)
// to RGB colors.
//
// For standard grayscale display, use a linear gray map where RGB channels
// all equal the input value.
func NewOutputLUT(colorMap [256]imaging.Color32) *OutputLUT {
	return &OutputLUT{
		colorMap: colorMap,
	}
}

// NewGrayscaleOutputLUT creates a standard grayscale Output LUT.
//
// Maps each value to an RGB gray where R=G=B=value.
func NewGrayscaleOutputLUT() *OutputLUT {
	var colorMap [256]imaging.Color32
	for i := 0; i < 256; i++ {
		gray := uint8(i)
		colorMap[i] = imaging.Color32{
			A: 255,
			R: gray,
			G: gray,
			B: gray,
		}
	}
	return &OutputLUT{colorMap: colorMap}
}

// IsValid returns true if the color map is initialized
func (o *OutputLUT) IsValid() bool {
	return true
}

// MinimumOutputValue returns the minimum output value (int32 min as a Color32 value)
func (o *OutputLUT) MinimumOutputValue() float64 {
	return float64(math.MinInt32)
}

// MaximumOutputValue returns the maximum output value (int32 max as a Color32 value)
func (o *OutputLUT) MaximumOutputValue() float64 {
	return float64(math.MaxInt32)
}

// Transform maps the input grayscale value to a Color32 value.
//
// The input value is clamped to [0, 255] range before lookup.
// The returned float64 represents the Color32.Value (packed ARGB int32).
func (o *OutputLUT) Transform(input float64) float64 {
	// Clamp input to [0, 255]
	var index int
	if input < 0 {
		index = 0
	} else if input > 255 {
		index = 255
	} else {
		index = int(input)
	}

	color := o.colorMap[index]
	return float64(color.ToInt32())
}

// Recalculate is a no-op for OutputLUT (color map is static)
func (o *OutputLUT) Recalculate() {
	// No recalculation needed
}

// GetColor returns the Color32 for the given grayscale value (clamped to [0, 255])
func (o *OutputLUT) GetColor(value int) imaging.Color32 {
	if value < 0 {
		return o.colorMap[0]
	}
	if value > 255 {
		return o.colorMap[255]
	}
	return o.colorMap[value]
}

// ColorMap returns a copy of the color map
func (o *OutputLUT) ColorMap() [256]imaging.Color32 {
	return o.colorMap
}
