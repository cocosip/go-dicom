// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package lut provides Lookup Table implementations for DICOM imaging.
// Lookup Tables (LUTs) are used to map input pixel values to output values
// in DICOM imaging. They are commonly used for windowing, normalization,
// and color mapping.
package lut

// LUT transforms input pixel values to output values.
type LUT interface {
	IsValid() bool
	MinimumOutputValue() float64
	MaximumOutputValue() float64
	Transform(input float64) float64
	Recalculate()
}

// VOILUTFunction identifies a DICOM VOI LUT function.
type VOILUTFunction string

const (
	// VOILUTFunctionLinear applies the DICOM LINEAR windowing function.
	VOILUTFunctionLinear VOILUTFunction = "LINEAR"
	// VOILUTFunctionLinearExact applies the DICOM LINEAR_EXACT windowing function.
	VOILUTFunctionLinearExact VOILUTFunction = "LINEAR_EXACT"
	// VOILUTFunctionSigmoid applies the DICOM SIGMOID windowing function.
	VOILUTFunctionSigmoid VOILUTFunction = "SIGMOID"
)
