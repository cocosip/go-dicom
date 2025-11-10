// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

// Component represents a single component of a color space with subsampling information.
type Component struct {
	Name       string
	SubSampleX int
	SubSampleY int
}

// NewComponent creates a new color space component.
func NewComponent(name string, subSampleX, subSampleY int) Component {
	return Component{
		Name:       name,
		SubSampleX: subSampleX,
		SubSampleY: subSampleY,
	}
}

// ColorSpace represents a color space with its components.
type ColorSpace struct {
	Name       string
	Components []Component
}

// NewColorSpace creates a new color space with the given name and components.
func NewColorSpace(name string, components ...Component) *ColorSpace {
	return &ColorSpace{
		Name:       name,
		Components: components,
	}
}

// String returns the name of the color space.
func (cs *ColorSpace) String() string {
	return cs.Name
}

// Equals compares two color spaces for equality.
func (cs *ColorSpace) Equals(other *ColorSpace) bool {
	if cs == nil || other == nil {
		return cs == other
	}
	return cs.Name == other.Name
}

// Standard color space definitions

// OneBit represents a 1-bit color space.
var OneBit = NewColorSpace("1-bit", NewComponent("Value", 1, 1))

// Grayscale represents a grayscale color space.
var Grayscale = NewColorSpace("Grayscale", NewComponent("Value", 1, 1))

// Indexed represents an indexed (palette) color space.
var Indexed = NewColorSpace("Indexed", NewComponent("Value", 1, 1))

// RGB represents an RGB color space with full resolution for all components.
var RGB = NewColorSpace(
	"RGB",
	NewComponent("Red", 1, 1),
	NewComponent("Green", 1, 1),
	NewComponent("Blue", 1, 1),
)

// BGR represents a BGR color space (reverse of RGB).
var BGR = NewColorSpace(
	"BGR",
	NewComponent("Blue", 1, 1),
	NewComponent("Green", 1, 1),
	NewComponent("Red", 1, 1),
)

// RGBA represents an RGBA color space with alpha channel.
var RGBA = NewColorSpace(
	"RGBA",
	NewComponent("Red", 1, 1),
	NewComponent("Green", 1, 1),
	NewComponent("Blue", 1, 1),
	NewComponent("Alpha", 1, 1),
)

// YCbCrJPEG represents a YCbCr 4:4:4 color space used in JPEG.
var YCbCrJPEG = NewColorSpace(
	"Y'CbCr 4:4:4 (JPEG)",
	NewComponent("Luminance", 1, 1),
	NewComponent("Blue Chroma", 1, 1),
	NewComponent("Red Chroma", 1, 1),
)
