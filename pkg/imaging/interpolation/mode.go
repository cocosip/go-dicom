// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package interpolation defines the interpolation modes shared by rendering
// transforms and image sampling.
package interpolation

// Mode selects an interpolation algorithm.
type Mode int

const (
	// ModeNearestNeighbor samples the nearest pixel center.
	ModeNearestNeighbor Mode = iota
	// ModeBilinear blends the four neighboring pixel centers.
	ModeBilinear
)
