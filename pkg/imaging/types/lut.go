// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

//nolint:revive // "types" is a descriptive and appropriate package name for common imaging types
package types

// LUT is the interface for all Lookup Table implementations.
// It transforms input pixel values to output values through various transformations.
//
// Reference: Based on fo-dicom ILUT interface
type LUT interface {
	// IsValid returns true if the lookup table is valid and ready to use
	IsValid() bool

	// MinimumOutputValue returns the minimum possible output value
	MinimumOutputValue() float64

	// MaximumOutputValue returns the maximum possible output value
	MaximumOutputValue() float64

	// Transform transforms an input value to an output value
	Transform(input float64) float64

	// Recalculate forces recalculation of the LUT if needed
	Recalculate()
}

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
