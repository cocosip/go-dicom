// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import "math"

// VOILUTFunction represents the VOI LUT function type
type VOILUTFunction string

const (
	// VOILinear represents the LINEAR VOI LUT function
	VOILinear VOILUTFunction = "LINEAR"
	// VOILinearExact represents the LINEAR_EXACT VOI LUT function (DICOM C.11.2.1.3.2)
	VOILinearExact VOILUTFunction = "LINEAR_EXACT"
	// VOISigmoid represents the SIGMOID VOI LUT function
	VOISigmoid VOILUTFunction = "SIGMOID"
)

// VOILUT defines the interface for VOI (Value of Interest) LUT
type VOILUT interface {
	LUT
	// WindowCenter returns the window center
	WindowCenter() float64
	// WindowWidth returns the window width
	WindowWidth() float64
	// SetWindow updates the window center and width
	SetWindow(center, width float64)
}

// baseVOILUT contains common fields and methods for all VOI LUT implementations
type baseVOILUT struct {
	windowCenter      float64
	windowWidth       float64
	windowCenterMin05 float64
	windowWidthMin1   float64
	windowWidthDiv2   float64
	windowStart       int
	windowEnd         int
	outputRange       float64
}

// newBaseVOILUT creates a new base VOI LUT
func newBaseVOILUT(center, width float64) *baseVOILUT {
	v := &baseVOILUT{
		outputRange: 255.0,
	}
	v.SetWindow(center, width)
	return v
}

// WindowCenter returns the window center
func (v *baseVOILUT) WindowCenter() float64 {
	return v.windowCenter
}

// WindowWidth returns the window width
func (v *baseVOILUT) WindowWidth() float64 {
	return v.windowWidth
}

// SetWindow updates the window center and width
func (v *baseVOILUT) SetWindow(center, width float64) {
	v.windowCenter = center
	v.windowWidth = width
	v.recalculate()
}

// IsValid returns false (always recalculate)
func (v *baseVOILUT) IsValid() bool {
	return false
}

// MinimumOutputValue returns the minimum output value
func (v *baseVOILUT) MinimumOutputValue() float64 {
	return 0
}

// MaximumOutputValue returns the maximum output value
func (v *baseVOILUT) MaximumOutputValue() float64 {
	return v.outputRange
}

// Recalculate forces the recalculation of LUT
func (v *baseVOILUT) Recalculate() {
	v.recalculate()
}

func (v *baseVOILUT) recalculate() {
	v.windowCenterMin05 = v.windowCenter - 0.5
	v.windowWidthMin1 = v.windowWidth - 1
	v.windowWidthDiv2 = v.windowWidthMin1 / 2
	v.windowStart = int(v.windowCenterMin05 - v.windowWidthDiv2)
	v.windowEnd = int(v.windowCenterMin05 + v.windowWidthDiv2)
}

// LinearVOILUT implements the LINEAR VOI LUT function
type LinearVOILUT struct {
	*baseVOILUT
}

// NewLinearVOILUT creates a new LINEAR VOI LUT
func NewLinearVOILUT(center, width float64) *LinearVOILUT {
	return &LinearVOILUT{
		baseVOILUT: newBaseVOILUT(center, width),
	}
}

// Transform applies the LINEAR transformation
func (v *LinearVOILUT) Transform(value float64) float64 {
	if v.windowWidth == 1 {
		if value < v.windowCenterMin05 {
			return v.MinimumOutputValue()
		}
		return v.MaximumOutputValue()
	}

	result := (((value-v.windowCenterMin05)/v.windowWidthMin1)+0.5)*v.outputRange + v.MinimumOutputValue()
	return math.Max(v.MinimumOutputValue(), math.Min(v.MaximumOutputValue(), result))
}

// LinearExactVOILUT implements the LINEAR_EXACT VOI LUT function (DICOM C.11.2.1.3.2)
type LinearExactVOILUT struct {
	*baseVOILUT
}

// NewLinearExactVOILUT creates a new LINEAR_EXACT VOI LUT
func NewLinearExactVOILUT(center, width float64) *LinearExactVOILUT {
	return &LinearExactVOILUT{
		baseVOILUT: newBaseVOILUT(center, width),
	}
}

// Transform applies the LINEAR_EXACT transformation
func (v *LinearExactVOILUT) Transform(value float64) float64 {
	if v.windowWidth == 1 {
		if value < v.windowCenterMin05 {
			return v.MinimumOutputValue()
		}
		return v.MaximumOutputValue()
	}

	result := ((value-v.windowCenter)/v.windowWidth+0.5)*v.outputRange + v.MinimumOutputValue()
	return math.Max(v.MinimumOutputValue(), math.Min(v.MaximumOutputValue(), result))
}

// SigmoidVOILUT implements the SIGMOID VOI LUT function
type SigmoidVOILUT struct {
	*baseVOILUT
}

// NewSigmoidVOILUT creates a new SIGMOID VOI LUT
func NewSigmoidVOILUT(center, width float64) *SigmoidVOILUT {
	return &SigmoidVOILUT{
		baseVOILUT: newBaseVOILUT(center, width),
	}
}

// Transform applies the SIGMOID transformation
func (v *SigmoidVOILUT) Transform(value float64) float64 {
	return 255.0 / (1.0 + math.Exp(-4.0*((value-v.windowCenter)/v.windowWidth)))
}

// CreateVOILUT creates a new VOILUT according to the specified function type
func CreateVOILUT(function VOILUTFunction, center, width float64) VOILUT {
	// LINEAR_EXACT is required when windowWidth < 1.0 with LINEAR function
	if width < 1.0 && function == VOILinear {
		function = VOILinearExact
	}

	switch function {
	case VOISigmoid:
		return NewSigmoidVOILUT(center, width)
	case VOILinearExact:
		return NewLinearExactVOILUT(center, width)
	case VOILinear:
		fallthrough
	default:
		return NewLinearVOILUT(center, width)
	}
}
