// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"
	"strconv"
	"strings"
)

// FilmOrientation represents the orientation of film
type FilmOrientation string

const (
	// FilmOrientationPortrait - vertical film position
	FilmOrientationPortrait FilmOrientation = "PORTRAIT"
	// FilmOrientationLandscape - horizontal film position
	FilmOrientationLandscape FilmOrientation = "LANDSCAPE"
)

// FilmSize represents standard film sizes
type FilmSize string

const (
	FilmSize8INX10IN   FilmSize = "8INX10IN"
	FilmSize8_5INX11IN FilmSize = "8_5INX11IN"
	FilmSize10INX12IN  FilmSize = "10INX12IN"
	FilmSize10INX14IN  FilmSize = "10INX14IN"
	FilmSize11INX14IN  FilmSize = "11INX14IN"
	FilmSize11INX17IN  FilmSize = "11INX17IN"
	FilmSize14INX14IN  FilmSize = "14INX14IN"
	FilmSize14INX17IN  FilmSize = "14INX17IN"
	FilmSize24CMX24CM  FilmSize = "24CMX24CM"
	FilmSize24CMX30CM  FilmSize = "24CMX30CM"
	FilmSize24CMX36CM  FilmSize = "24CMX36CM"
	FilmSize26CMX30CM  FilmSize = "26CMX30CM"
	FilmSizeA4         FilmSize = "A4"
	FilmSizeA3         FilmSize = "A3"
)

// MagnificationType represents the type of image magnification
type MagnificationType string

const (
	// MagnificationReplicate - pixel replication
	MagnificationReplicate MagnificationType = "REPLICATE"
	// MagnificationBilinear - bilinear interpolation
	MagnificationBilinear MagnificationType = "BILINEAR"
	// MagnificationCubic - cubic interpolation
	MagnificationCubic MagnificationType = "CUBIC"
	// MagnificationNone - no magnification
	MagnificationNone MagnificationType = "NONE"
)

// TrimMode represents the trimming mode for the film
type TrimMode string

const (
	// TrimYes - trim the film
	TrimYes TrimMode = "YES"
	// TrimNo - do not trim the film
	TrimNo TrimMode = "NO"
)

// BorderDensity represents the density of the film border
type BorderDensity string

const (
	BorderDensityBlack BorderDensity = "BLACK"
	BorderDensityWhite BorderDensity = "WHITE"
)

// EmptyImageDensity represents the density of empty image boxes
type EmptyImageDensity string

const (
	EmptyImageDensityBlack EmptyImageDensity = "BLACK"
	EmptyImageDensityWhite EmptyImageDensity = "WHITE"
)

// FilmBox represents a Basic Film Box in DICOM Print Management
//
// A Film Box defines the layout and characteristics of a single sheet of film
// and contains one or more Image Boxes arranged according to the Image Display Format.
//
// Reference: DICOM Part 3, Section C.13.3
type FilmBox struct {
	// filmSession is the parent Film Session
	filmSession *FilmSession

	// SOPInstanceUID is the Film Box SOP Instance UID
	SOPInstanceUID string

	// ImageDisplayFormat specifies the type of image display format
	// Examples: "STANDARD\2,2", "ROW\1,2,1", "COL\2,2", "SLIDE", "CUSTOM\1"
	ImageDisplayFormat string

	// FilmOrientation specifies the film orientation
	// Enumerated Values: PORTRAIT, LANDSCAPE
	FilmOrientation FilmOrientation

	// FilmSizeID specifies the film size
	// Defined Terms: 8INX10IN, A4, 24CMX30CM, etc.
	FilmSizeID FilmSize

	// MagnificationType specifies the type of magnification
	// Enumerated Values: REPLICATE, BILINEAR, CUBIC, NONE
	MagnificationType MagnificationType

	// MaxDensity specifies the maximum density (0-255, or 0 for printer default)
	MaxDensity uint16

	// MinDensity specifies the minimum density (0-255, or 0 for printer default)
	MinDensity uint16

	// Trim specifies whether to trim the film
	// Enumerated Values: YES, NO
	Trim TrimMode

	// BorderDensity specifies the density of the film border
	BorderDensity BorderDensity

	// EmptyImageDensity specifies the density of empty image boxes
	EmptyImageDensity EmptyImageDensity

	// ConfigurationInformation contains implementation-specific configuration
	ConfigurationInformation string

	// BasicImageBoxes is the list of Image Boxes in this Film Box
	BasicImageBoxes []*ImageBox
}

// NewFilmBox creates a new Film Box
func NewFilmBox(sopInstanceUID string, imageDisplayFormat string) *FilmBox {
	if sopInstanceUID == "" {
		// Generate a UID if not provided
		sopInstanceUID = "1.2.840.10008.5.1.1.2.1" // Placeholder
	}

	return &FilmBox{
		SOPInstanceUID:           sopInstanceUID,
		ImageDisplayFormat:       imageDisplayFormat,
		FilmOrientation:          FilmOrientationPortrait,
		FilmSizeID:               FilmSize14INX17IN,
		MagnificationType:        MagnificationReplicate,
		MaxDensity:               0, // Printer default
		MinDensity:               0, // Printer default
		Trim:                     TrimNo,
		BorderDensity:            BorderDensityBlack,
		EmptyImageDensity:        EmptyImageDensityBlack,
		ConfigurationInformation: "",
		BasicImageBoxes:          make([]*ImageBox, 0),
	}
}

// ParseImageDisplayFormat parses the image display format string and returns
// the total number of image boxes required
func ParseImageDisplayFormat(format string) (int, error) {
	parts := strings.Split(format, "\\")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid image display format: %s", format)
	}

	formatType := parts[0]
	params := parts[1]

	switch formatType {
	case "STANDARD":
		// STANDARD\C,R - C columns, R rows
		dimensions := strings.Split(params, ",")
		if len(dimensions) != 2 {
			return 0, fmt.Errorf("invalid STANDARD format: %s", format)
		}
		cols, err1 := strconv.Atoi(dimensions[0])
		rows, err2 := strconv.Atoi(dimensions[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid STANDARD dimensions: %s", format)
		}
		return cols * rows, nil

	case "ROW":
		// ROW\R1,R2,R3,... - boxes per row
		boxes := 0
		rowCounts := strings.Split(params, ",")
		for _, rc := range rowCounts {
			count, err := strconv.Atoi(rc)
			if err != nil {
				return 0, fmt.Errorf("invalid ROW format: %s", format)
			}
			boxes += count
		}
		return boxes, nil

	case "COL":
		// COL\C1,C2,C3,... - boxes per column
		boxes := 0
		colCounts := strings.Split(params, ",")
		for _, cc := range colCounts {
			count, err := strconv.Atoi(cc)
			if err != nil {
				return 0, fmt.Errorf("invalid COL format: %s", format)
			}
			boxes += count
		}
		return boxes, nil

	case "SLIDE", "SUPERSLIDE":
		// Configuration dependent, assume standard counts
		if formatType == "SLIDE" {
			return 20, nil // Typical 35mm slide count
		}
		return 15, nil // Typical 40mm superslide count

	case "CUSTOM":
		// CUSTOM\i - implementation specific
		// Return a default value or require additional configuration
		return 1, nil

	default:
		return 0, fmt.Errorf("unsupported format type: %s", formatType)
	}
}

// InitializeImageBoxes creates and initializes Image Boxes based on the Image Display Format
func (fb *FilmBox) InitializeImageBoxes() error {
	count, err := ParseImageDisplayFormat(fb.ImageDisplayFormat)
	if err != nil {
		return err
	}

	fb.BasicImageBoxes = make([]*ImageBox, count)
	for i := 0; i < count; i++ {
		fb.BasicImageBoxes[i] = NewImageBox(fmt.Sprintf("%s.%d", fb.SOPInstanceUID, i+1), fb.filmSession != nil && fb.filmSession.IsColor)
		fb.BasicImageBoxes[i].filmBox = fb
		fb.BasicImageBoxes[i].ImageBoxPosition = uint16(i + 1)
	}

	return nil
}

// AddImageBox adds an Image Box to the Film Box
func (fb *FilmBox) AddImageBox(imageBox *ImageBox) {
	if imageBox != nil {
		imageBox.filmBox = fb
		fb.BasicImageBoxes = append(fb.BasicImageBoxes, imageBox)
	}
}

// GetImageBox returns the Image Box at the specified position (1-based)
func (fb *FilmBox) GetImageBox(position uint16) *ImageBox {
	for _, ib := range fb.BasicImageBoxes {
		if ib.ImageBoxPosition == position {
			return ib
		}
	}
	return nil
}

// ImageBoxCount returns the number of Image Boxes in the Film Box
func (fb *FilmBox) ImageBoxCount() int {
	return len(fb.BasicImageBoxes)
}

// FilmSession returns the parent Film Session
func (fb *FilmBox) FilmSession() *FilmSession {
	return fb.filmSession
}

// IsValid checks if the Film Box configuration is valid
func (fb *FilmBox) IsValid() bool {
	// Must have a valid Image Display Format
	if fb.ImageDisplayFormat == "" {
		return false
	}

	// Must have at least one Image Box
	if len(fb.BasicImageBoxes) == 0 {
		return false
	}

	// Validate all Image Boxes
	for _, ib := range fb.BasicImageBoxes {
		if !ib.IsValid() {
			return false
		}
	}

	return true
}
