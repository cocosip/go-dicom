// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

// LUT defines the interface for lookup table operations
type LUT interface {
	// IsValid returns true if the lookup table is valid
	IsValid() bool

	// MinimumOutputValue returns the minimum output value
	MinimumOutputValue() float64

	// MaximumOutputValue returns the maximum output value
	MaximumOutputValue() float64

	// Transform transforms input value into output value
	Transform(input float64) float64

	// Recalculate forces the recalculation of LUT
	Recalculate()
}

// ModalityLUT defines the interface for Modality lookup tables
// Either Modality Rescale LUT or Modality Sequence LUT
type ModalityLUT interface {
	LUT
}
