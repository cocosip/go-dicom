// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
	"strings"
)

// PhotometricInterpretation represents the photometric interpretation of pixel data.
// This corresponds to DICOM tag (0028,0004).
type PhotometricInterpretation struct {
	Value       string      // The DICOM defined term value
	Description string      // Human-readable description
	IsColor     bool        // Whether this represents color (true) or grayscale (false)
	IsPalette   bool        // Whether this uses a color palette
	IsYBR       bool        // Whether this uses YBR color scheme
	ColorSpace  *ColorSpace // Associated color space (may be nil)
}

// String returns the description of the photometric interpretation.
func (pi *PhotometricInterpretation) String() string {
	return pi.Description
}

// Equals compares two photometric interpretations for equality.
func (pi *PhotometricInterpretation) Equals(other *PhotometricInterpretation) bool {
	if pi == nil || other == nil {
		return pi == other
	}
	return pi.Value == other.Value
}

// ParsePhotometricInterpretation parses a photometric interpretation string.
// The string is typically obtained from DICOM tag (0028,0004).
func ParsePhotometricInterpretation(value string) (*PhotometricInterpretation, error) {
	// Trim spaces and null characters
	value = strings.Trim(value, " \x00")

	switch value {
	case "MONOCHROME1":
		return Monochrome1, nil
	case "MONOCHROME", "MONOCHROME2":
		return Monochrome2, nil
	case "PALETTE COLOR":
		return PaletteColor, nil
	case "RGB":
		return RGBPhotometric, nil
	case "YBR_FULL":
		return YbrFull, nil
	case "YBR_FULL_422":
		return YbrFull422, nil
	case "YBR_PARTIAL_422":
		return YbrPartial422, nil
	case "YBR_PARTIAL_420":
		return YbrPartial420, nil
	case "YBR_ICT":
		return YbrIct, nil
	case "YBR_RCT":
		return YbrRct, nil
	default:
		return nil, fmt.Errorf("unknown photometric interpretation: %s", value)
	}
}

// MustParsePhotometricInterpretation parses a photometric interpretation string
// and panics if the value is invalid.
func MustParsePhotometricInterpretation(value string) *PhotometricInterpretation {
	pi, err := ParsePhotometricInterpretation(value)
	if err != nil {
		panic(err)
	}
	return pi
}

// Standard photometric interpretation definitions

// Monochrome1 represents a single monochrome image plane where the minimum sample
// value is intended to be displayed as white after any VOI gray scale transformations.
// This value may be used only when Samples per Pixel (0028,0002) has a value of 1.
var Monochrome1 = &PhotometricInterpretation{
	Value:       "MONOCHROME1",
	Description: "Monochrome 1",
	IsColor:     false,
	IsPalette:   false,
	IsYBR:       false,
	ColorSpace:  Grayscale,
}

// Monochrome2 represents a single monochrome image plane where the minimum sample
// value is intended to be displayed as black after any VOI gray scale transformations.
// This value may be used only when Samples per Pixel (0028,0002) has a value of 1.
var Monochrome2 = &PhotometricInterpretation{
	Value:       "MONOCHROME2",
	Description: "Monochrome 2",
	IsColor:     false,
	IsPalette:   false,
	IsYBR:       false,
	ColorSpace:  Grayscale,
}

// PaletteColor represents a color image with a single sample per pixel.
// The pixel value is used as an index into Red, Blue, and Green Palette Color Lookup Tables.
// This value may be used only when Samples per Pixel (0028,0002) has a value of 1.
var PaletteColor = &PhotometricInterpretation{
	Value:       "PALETTE COLOR",
	Description: "Palette Color",
	IsColor:     true,
	IsPalette:   true,
	IsYBR:       false,
	ColorSpace:  Indexed,
}

// RGBPhotometric represents a color image described by red, green, and blue image planes.
// The minimum sample value for each color plane represents minimum intensity of the color.
// This value may be used only when Samples per Pixel (0028,0002) has a value of 3.
var RGBPhotometric = &PhotometricInterpretation{
	Value:       "RGB",
	Description: "RGB",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       false,
	ColorSpace:  RGB,
}

// YbrFull represents a color image described by one luminance (Y) and two chrominance
// planes (Cb and Cr). Black is represented by Y equal to zero. The absence of color
// is represented by both Cb and Cr values equal to half full scale.
//
// Conversion equations (for Bits Allocated = 8):
//
//	Y  = +0.2990*R + 0.5870*G + 0.1140*B
//	Cb = -0.1687*R - 0.3313*G + 0.5000*B + 128
//	Cr = +0.5000*R - 0.4187*G - 0.0813*B + 128
var YbrFull = &PhotometricInterpretation{
	Value:       "YBR_FULL",
	Description: "YBR Full",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  YCbCrJPEG,
}

// YbrFull422 is the same as YBR_FULL except that the Cb and Cr values are sampled
// horizontally at half the Y rate (4:2:2 subsampling).
// This Photometric Interpretation is only allowed with Planar Configuration equal to 0.
var YbrFull422 = &PhotometricInterpretation{
	Value:       "YBR_FULL_422",
	Description: "YBR Full 4:2:2",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  nil,
}

// YbrPartial422 is the same as YBR_FULL_422 except that:
//   - Black corresponds to Y = 16
//   - Y is restricted to 220 levels (maximum value is 235)
//   - Cb and Cr each has a minimum value of 16
//   - Cb and Cr are restricted to 225 levels (maximum value is 240)
//   - Lack of color is represented by Cb and Cr equal to 128
//
// Conversion equations (for Bits Allocated = 8):
//
//	Y  = +0.2568*R + 0.5041*G + 0.0979*B + 16
//	Cb = -0.1482*R - 0.2910*G + 0.4392*B + 128
//	Cr = +0.4392*R - 0.3678*G - 0.0714*B + 128
var YbrPartial422 = &PhotometricInterpretation{
	Value:       "YBR_PARTIAL_422",
	Description: "YBR Partial 4:2:2",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  nil,
}

// YbrPartial420 is the same as YBR_PARTIAL_422 except that the Cb and Cr values
// are sampled horizontally and vertically at half the Y rate (4:2:0 subsampling).
// This results in four times less Cb and Cr values than Y values.
var YbrPartial420 = &PhotometricInterpretation{
	Value:       "YBR_PARTIAL_420",
	Description: "YBR Partial 4:2:0",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  nil,
}

// YbrIct represents YBR with Irreversible Color Transformation (used in JPEG 2000).
//
// Conversion equations:
//
//	Y  = +0.29900*R + 0.58700*G + 0.11400*B
//	Cb = -0.16875*R - 0.33126*G + 0.50000*B
//	Cr = +0.50000*R - 0.41869*G - 0.08131*B
var YbrIct = &PhotometricInterpretation{
	Value:       "YBR_ICT",
	Description: "YBR Irreversible Color Transformation (JPEG 2000)",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  nil,
}

// YbrRct represents YBR with Reversible Color Transformation (used in JPEG 2000).
//
// Conversion equations:
//
//	Y  = floor((R + 2*G + B) / 4)
//	Cb = B - G
//	Cr = R - G
//
// Reverse conversion:
//
//	R = Cr + G
//	G = Y - floor((Cb + Cr) / 4)
//	B = Cb + G
var YbrRct = &PhotometricInterpretation{
	Value:       "YBR_RCT",
	Description: "YBR Reversible Color Transformation (JPEG 2000)",
	IsColor:     true,
	IsPalette:   false,
	IsYBR:       true,
	ColorSpace:  nil,
}
