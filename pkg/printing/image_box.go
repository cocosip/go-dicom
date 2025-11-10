// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

// Polarity represents the polarity for image printing
type Polarity string

const (
	// PolarityNormal - pixels printed as specified by Photometric Interpretation
	PolarityNormal Polarity = "NORMAL"
	// PolarityReverse - pixels printed with opposite polarity
	PolarityReverse Polarity = "REVERSE"
)

// ImageBox represents a Basic Image Box in DICOM Print Management
//
// An Image Box defines a single image position on a film sheet and contains
// the image data and display parameters for that position.
//
// Reference: DICOM Part 3, Section C.13.5
type ImageBox struct {
	// filmBox is the parent Film Box
	filmBox *FilmBox

	// SOPClassUID is the Image Box SOP Class UID
	// Different for color vs grayscale
	SOPClassUID string

	// SOPInstanceUID is the Image Box SOP Instance UID
	SOPInstanceUID string

	// ImageBoxPosition specifies the position of the image on the film
	// Based on image display format (1-based index)
	ImageBoxPosition uint16

	// Polarity specifies whether minimum pixel values are printed black or white
	// Enumerated Values: NORMAL, REVERSE
	Polarity Polarity

	// MagnificationType specifies the magnification type for this image box
	// Can override the Film Box setting
	MagnificationType MagnificationType

	// SmoothingType specifies the type of smoothing to apply
	// Implementation-specific values
	SmoothingType string

	// RequestedImageSize specifies the requested image size
	// Format: width\height in mm, or predefined sizes
	RequestedImageSize string

	// PreformattedGrayscaleImageSequence contains the image data for grayscale
	PreformattedGrayscaleImageSequence []byte

	// PreformattedColorImageSequence contains the image data for color
	PreformattedColorImageSequence []byte

	// IsColor indicates whether this is a color image box
	IsColor bool
}

const (
	// SOPClassColorImageBox is the Basic Color Image Box SOP Class UID
	SOPClassColorImageBox = "1.2.840.10008.5.1.1.4.1"

	// SOPClassGrayscaleImageBox is the Basic Grayscale Image Box SOP Class UID
	SOPClassGrayscaleImageBox = "1.2.840.10008.5.1.1.4"
)

// NewImageBox creates a new Image Box
func NewImageBox(sopInstanceUID string, isColor bool) *ImageBox {
	if sopInstanceUID == "" {
		// Generate a UID if not provided
		sopInstanceUID = "1.2.840.10008.5.1.1.4.1.1" // Placeholder
	}

	sopClassUID := SOPClassGrayscaleImageBox
	if isColor {
		sopClassUID = SOPClassColorImageBox
	}

	return &ImageBox{
		SOPClassUID:                        sopClassUID,
		SOPInstanceUID:                     sopInstanceUID,
		ImageBoxPosition:                   1,
		Polarity:                           PolarityNormal,
		MagnificationType:                  MagnificationReplicate,
		SmoothingType:                      "",
		RequestedImageSize:                 "",
		PreformattedGrayscaleImageSequence: nil,
		PreformattedColorImageSequence:     nil,
		IsColor:                            isColor,
	}
}

// SetImageData sets the image data for the Image Box
func (ib *ImageBox) SetImageData(imageData []byte) {
	if ib.IsColor {
		ib.PreformattedColorImageSequence = make([]byte, len(imageData))
		copy(ib.PreformattedColorImageSequence, imageData)
	} else {
		ib.PreformattedGrayscaleImageSequence = make([]byte, len(imageData))
		copy(ib.PreformattedGrayscaleImageSequence, imageData)
	}
}

// GetImageData returns the image data for the Image Box
func (ib *ImageBox) GetImageData() []byte {
	if ib.IsColor {
		return ib.PreformattedColorImageSequence
	}
	return ib.PreformattedGrayscaleImageSequence
}

// HasImageData returns true if the Image Box has image data
func (ib *ImageBox) HasImageData() bool {
	if ib.IsColor {
		return len(ib.PreformattedColorImageSequence) > 0
	}
	return len(ib.PreformattedGrayscaleImageSequence) > 0
}

// ClearImageData clears the image data
func (ib *ImageBox) ClearImageData() {
	ib.PreformattedGrayscaleImageSequence = nil
	ib.PreformattedColorImageSequence = nil
}

// FilmBox returns the parent Film Box
func (ib *ImageBox) FilmBox() *FilmBox {
	return ib.filmBox
}

// IsValid checks if the Image Box configuration is valid
func (ib *ImageBox) IsValid() bool {
	// Must have a valid SOP Class UID
	if ib.SOPClassUID == "" {
		return false
	}

	// Must have a valid SOP Instance UID
	if ib.SOPInstanceUID == "" {
		return false
	}

	// Position must be at least 1
	if ib.ImageBoxPosition < 1 {
		return false
	}

	// For a valid print job, should have image data
	// But we allow empty image boxes for configuration
	return true
}

// GetEffectiveMagnificationType returns the magnification type for this image box
// Uses the image box's setting if specified, otherwise uses the film box's setting
func (ib *ImageBox) GetEffectiveMagnificationType() MagnificationType {
	if ib.MagnificationType != "" {
		return ib.MagnificationType
	}
	if ib.filmBox != nil {
		return ib.filmBox.MagnificationType
	}
	return MagnificationReplicate
}
