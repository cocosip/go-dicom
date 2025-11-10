// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

// PixelRepresentation represents signed/unsigned data of the pixel samples.
// This corresponds to DICOM tag (0028,0103).
// Each sample shall have the same pixel representation.
type PixelRepresentation uint16

const (
	// UnsignedPixels represents unsigned integer pixel data.
	UnsignedPixels PixelRepresentation = 0

	// SignedPixels represents 2's complement (signed) integer pixel data.
	SignedPixels PixelRepresentation = 1
)

// String returns a string representation of the pixel representation.
func (pr PixelRepresentation) String() string {
	switch pr {
	case UnsignedPixels:
		return "Unsigned"
	case SignedPixels:
		return "Signed"
	default:
		return "Unknown"
	}
}

// IsSigned returns true if the pixel representation is signed.
func (pr PixelRepresentation) IsSigned() bool {
	return pr == SignedPixels
}

// PlanarConfiguration indicates whether the color pixel data are sent
// color-by-plane or color-by-pixel. This corresponds to DICOM tag (0028,0006).
//
// This attribute shall be present if Samples per Pixel (0028,0002) has a value
// greater than 1. It shall not be present otherwise.
type PlanarConfiguration uint16

const (
	// InterleavedPlanar means the sample values for the first pixel are followed
	// by the sample values for the second pixel, etc.
	// For RGB images, this means the order is: R1, G1, B1, R2, G2, B2, ..., etc.
	InterleavedPlanar PlanarConfiguration = 0

	// PlanarPlanar means each color plane shall be sent contiguously.
	// For RGB images, this means the order is: R1, R2, R3, ..., G1, G2, G3, ..., B1, B2, B3, etc.
	PlanarPlanar PlanarConfiguration = 1
)

// String returns a string representation of the planar configuration.
func (pc PlanarConfiguration) String() string {
	switch pc {
	case InterleavedPlanar:
		return "Interleaved"
	case PlanarPlanar:
		return "Planar"
	default:
		return "Unknown"
	}
}

// IsInterleaved returns true if the planar configuration is interleaved.
func (pc PlanarConfiguration) IsInterleaved() bool {
	return pc == InterleavedPlanar
}

// IsPlanar returns true if the planar configuration is planar.
func (pc PlanarConfiguration) IsPlanar() bool {
	return pc == PlanarPlanar
}
