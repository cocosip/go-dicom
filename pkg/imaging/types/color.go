// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package types provides common types used across the imaging package.
package types

// Color32 represents a 32-bit RGBA color
type Color32 struct {
	A uint8 // Alpha
	R uint8 // Red
	G uint8 // Green
	B uint8 // Blue
}

// NewColor32 creates a new Color32
func NewColor32(a, r, g, b uint8) Color32 {
	return Color32{A: a, R: r, G: g, B: b}
}

// ToInt32 converts the Color32 to a packed int32 value (ARGB format)
func (c Color32) ToInt32() int32 {
	return int32(c.A)<<24 | int32(c.R)<<16 | int32(c.G)<<8 | int32(c.B)
}
