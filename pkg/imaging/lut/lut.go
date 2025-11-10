// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package lut provides Lookup Table implementations for DICOM imaging.
// Lookup Tables (LUTs) are used to map input pixel values to output values
// in DICOM imaging. They are commonly used for windowing, normalization,
// and color mapping.
package lut

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
