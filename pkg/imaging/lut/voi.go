// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import (
	"math"
	"strings"
)

// VOILUTFunction represents the type of VOI LUT function
type VOILUTFunction string

const (
	// VOILUTFunctionLinear is the standard linear windowing function
	VOILUTFunctionLinear VOILUTFunction = "LINEAR"
	// VOILUTFunctionLinearExact is the DICOM C.11.2.1.3.2 exact linear function
	VOILUTFunctionLinearExact VOILUTFunction = "LINEAR_EXACT"
	// VOILUTFunctionSigmoid is a sigmoid windowing function
	VOILUTFunctionSigmoid VOILUTFunction = "SIGMOID"
)

// VOILUT is the base interface for Value of Interest LUT implementations.
//
// VOI LUT applies window/level transformations to convert modality values
// to presentation values suitable for display.
//
// Reference: DICOM Part 3, Section C.11.2 (VOI LUT Module)
// Based on fo-dicom VOILUT
type VOILUT interface {
	LUT
	WindowCenter() float64
	WindowWidth() float64
}

// baseVOILUT contains common VOI LUT fields and logic
type baseVOILUT struct {
	windowCenter    float64
	windowWidth     float64
	windowCenterMin float64 // windowCenter - 0.5
	windowWidthMin  float64 // windowWidth - 1
	windowWidthDiv2 float64 // (windowWidth - 1) / 2
	windowStart     int
	windowEnd       int
	outputRange     float64
	needsRecalc     bool
}

func (v *baseVOILUT) WindowCenter() float64 {
	return v.windowCenter
}

func (v *baseVOILUT) WindowWidth() float64 {
	return v.windowWidth
}

func (v *baseVOILUT) MinimumOutputValue() float64 {
	return 0
}

func (v *baseVOILUT) MaximumOutputValue() float64 {
	return 255
}

func (v *baseVOILUT) IsValid() bool {
	// Always recalculate for VOI LUTs
	return false
}

func (v *baseVOILUT) Recalculate() {
	if v.needsRecalc {
		v.windowCenterMin = v.windowCenter - 0.5
		v.windowWidthMin = v.windowWidth - 1
		v.windowWidthDiv2 = v.windowWidthMin / 2
		v.windowStart = int(v.windowCenterMin - v.windowWidthDiv2)
		v.windowEnd = int(v.windowCenterMin + v.windowWidthDiv2)
		v.outputRange = 255
		v.needsRecalc = false
	}
}

func (v *baseVOILUT) updateWindow(center, width float64) {
	if v.windowCenter != center || v.windowWidth != width {
		v.windowCenter = center
		v.windowWidth = width
		v.needsRecalc = true
	}
}

// VOILinearLUT implements linear VOI LUT transformation.
//
// This is the standard DICOM windowing function.
//
// Reference: DICOM Part 3, Section C.11.2.1.2
type VOILinearLUT struct {
	baseVOILUT
}

// NewVOILinearLUT creates a new linear VOI LUT
func NewVOILinearLUT(windowCenter, windowWidth float64) *VOILinearLUT {
	lut := &VOILinearLUT{}
	lut.updateWindow(windowCenter, windowWidth)
	lut.Recalculate()
	return lut
}

// Transform applies linear windowing transformation
func (v *VOILinearLUT) Transform(input float64) float64 {
	v.Recalculate()

	if v.windowWidth == 1 {
		if input < v.windowCenterMin {
			return v.MinimumOutputValue()
		}
		return v.MaximumOutputValue()
	}

	// Linear transformation with clamping
	output := (((input - v.windowCenterMin) / v.windowWidthMin) + 0.5) * v.outputRange
	return math.Min(v.MaximumOutputValue(), math.Max(v.MinimumOutputValue(), output))
}

// SetWindow updates the window center and width
func (v *VOILinearLUT) SetWindow(center, width float64) {
	v.updateWindow(center, width)
}

// VOILinearExactLUT implements the DICOM C.11.2.1.3.2 exact linear transformation.
//
// This differs slightly from the standard linear LUT in how it handles
// the window center calculation.
//
// Reference: DICOM Part 3, Section C.11.2.1.3.2
type VOILinearExactLUT struct {
	baseVOILUT
}

// NewVOILinearExactLUT creates a new exact linear VOI LUT
func NewVOILinearExactLUT(windowCenter, windowWidth float64) *VOILinearExactLUT {
	lut := &VOILinearExactLUT{}
	lut.updateWindow(windowCenter, windowWidth)
	lut.Recalculate()
	return lut
}

// Transform applies exact linear windowing transformation
func (v *VOILinearExactLUT) Transform(input float64) float64 {
	v.Recalculate()

	if v.windowWidth == 1 {
		if input < v.windowCenterMin {
			return v.MinimumOutputValue()
		}
		return v.MaximumOutputValue()
	}

	// Exact linear transformation per DICOM spec
	output := ((input - v.windowCenter) / v.windowWidth + 0.5) * v.outputRange
	return math.Min(v.MaximumOutputValue(), math.Max(v.MinimumOutputValue(), output))
}

// SetWindow updates the window center and width
func (v *VOILinearExactLUT) SetWindow(center, width float64) {
	v.updateWindow(center, width)
}

// VOISigmoidLUT implements sigmoid VOI LUT transformation.
//
// This provides a smooth, non-linear windowing function that can be
// useful for certain types of images.
type VOISigmoidLUT struct {
	baseVOILUT
}

// NewVOISigmoidLUT creates a new sigmoid VOI LUT
func NewVOISigmoidLUT(windowCenter, windowWidth float64) *VOISigmoidLUT {
	lut := &VOISigmoidLUT{}
	lut.updateWindow(windowCenter, windowWidth)
	lut.Recalculate()
	return lut
}

// Transform applies sigmoid windowing transformation
func (v *VOISigmoidLUT) Transform(input float64) float64 {
	v.Recalculate()

	// Sigmoid transformation: 255 / (1 + exp(-4 * (input - center) / width))
	return 255.0 / (1.0 + math.Exp(-4.0*((input-v.windowCenter)/v.windowWidth)))
}

// SetWindow updates the window center and width
func (v *VOISigmoidLUT) SetWindow(center, width float64) {
	v.updateWindow(center, width)
}

// CreateVOILUT creates a VOI LUT based on the specified function type
func CreateVOILUT(function VOILUTFunction, windowCenter, windowWidth float64) VOILUT {
	switch strings.ToUpper(string(function)) {
	case "SIGMOID":
		return NewVOISigmoidLUT(windowCenter, windowWidth)
	case "LINEAR_EXACT":
		return NewVOILinearExactLUT(windowCenter, windowWidth)
	default:
		return NewVOILinearLUT(windowCenter, windowWidth)
	}
}
