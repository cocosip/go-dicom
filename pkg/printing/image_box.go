// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

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

	// MaxDensity overrides the Film Box maximum density when non-nil.
	MaxDensity *uint16

	// MinDensity overrides the Film Box minimum density when non-nil.
	MinDensity *uint16

	// ConfigurationInformation overrides the Film Box configuration when non-nil.
	ConfigurationInformation *string

	// RequestedDecimateCropBehavior specifies DECIMATE, CROP, or FAIL.
	RequestedDecimateCropBehavior string

	// PreformattedGrayscaleImageSequence contains the image data for grayscale
	PreformattedGrayscaleImageSequence []byte

	// PreformattedColorImageSequence contains the image data for color
	PreformattedColorImageSequence []byte

	// IsColor indicates whether this is a color image box
	IsColor bool

	imageSequence *dataset.Dataset
}

// SetImageSequence stores an independent copy of a complete preformatted image
// sequence item Dataset.
func (ib *ImageBox) SetImageSequence(value *dataset.Dataset) error {
	if ib == nil {
		return fmt.Errorf("printing: nil ImageBox")
	}
	if value == nil {
		ib.imageSequence = nil
		ib.PreformattedGrayscaleImageSequence = nil
		ib.PreformattedColorImageSequence = nil
		return nil
	}
	clone, err := value.DeepCloneChecked()
	if err != nil {
		return fmt.Errorf("printing: clone Image Box sequence: %w", err)
	}
	ib.imageSequence = clone
	ib.setLegacyImageData(pixelDataBytes(clone))
	return nil
}

// ImageSequence returns an independent copy of the complete preformatted image
// sequence item Dataset.
func (ib *ImageBox) ImageSequence() (*dataset.Dataset, error) {
	if ib == nil {
		return nil, fmt.Errorf("printing: nil ImageBox")
	}
	if ib.imageSequence != nil {
		result, err := ib.imageSequence.DeepCloneChecked()
		if err != nil {
			return nil, fmt.Errorf("printing: clone Image Box sequence: %w", err)
		}
		return result, nil
	}
	imageData := ib.GetImageData()
	if len(imageData) == 0 {
		return nil, nil
	}
	result := dataset.New()
	if err := result.Add(element.NewOtherByte(tag.PixelData, append([]byte(nil), imageData...))); err != nil {
		return nil, fmt.Errorf("printing: add Image Box Pixel Data: %w", err)
	}
	return result, nil
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
		sopInstanceUID = newSOPInstanceUID()
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

// ValidateImageAttributes checks that the preformatted image sequence carries
// the Image Box Pixel Presentation Module attributes with values permitted by
// PS3.3 C.13.5, Table C.13-5. Pixel conversion is the caller's responsibility:
// an invalid image is rejected instead of being sent unconverted.
func (ib *ImageBox) ValidateImageAttributes() error {
	if ib == nil {
		return fmt.Errorf("printing: nil ImageBox")
	}
	sequence, err := ib.ImageSequence()
	if err != nil {
		return err
	}
	if sequence == nil {
		return fmt.Errorf("printing: ImageBox has no image sequence")
	}
	required := []*tag.Tag{
		tag.Rows, tag.Columns, tag.BitsAllocated, tag.BitsStored, tag.HighBit,
		tag.SamplesPerPixel, tag.PhotometricInterpretation, tag.PixelRepresentation,
		tag.PixelData,
	}
	if ib.IsColor {
		required = append(required, tag.PlanarConfiguration)
	}
	for _, requiredTag := range required {
		if _, ok := sequence.Get(requiredTag); !ok {
			return fmt.Errorf("printing: image box sequence is missing required attribute %s", requiredTag)
		}
	}
	if ib.IsColor {
		return validateColorImageAttributes(sequence)
	}
	return validateGrayscaleImageAttributes(sequence)
}

func validateGrayscaleImageAttributes(sequence *dataset.Dataset) error {
	samplesPerPixel, err := uint16ImageAttribute(sequence, tag.SamplesPerPixel)
	if err != nil {
		return err
	}
	if samplesPerPixel != 1 {
		return fmt.Errorf("printing: grayscale image box requires Samples per Pixel 1, got %d", samplesPerPixel)
	}
	photometric, err := stringImageAttribute(sequence, tag.PhotometricInterpretation)
	if err != nil {
		return err
	}
	if photometric != "MONOCHROME1" && photometric != "MONOCHROME2" {
		return fmt.Errorf("printing: grayscale image box requires MONOCHROME1 or MONOCHROME2, got %q", photometric)
	}
	bitsStored, err := uint16ImageAttribute(sequence, tag.BitsStored)
	if err != nil {
		return err
	}
	var bitsAllocated, highBit uint16
	switch bitsStored {
	case 8:
		bitsAllocated, highBit = 8, 7
	case 12:
		bitsAllocated, highBit = 16, 11
	default:
		return fmt.Errorf("printing: grayscale image box requires Bits Stored 8 or 12, got %d", bitsStored)
	}
	if value, err := uint16ImageAttribute(sequence, tag.BitsAllocated); err != nil {
		return err
	} else if value != bitsAllocated {
		return fmt.Errorf("printing: grayscale image box with Bits Stored %d requires Bits Allocated %d, got %d", bitsStored, bitsAllocated, value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.HighBit); err != nil {
		return err
	} else if value != highBit {
		return fmt.Errorf("printing: grayscale image box with Bits Stored %d requires High Bit %d, got %d", bitsStored, highBit, value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.PixelRepresentation); err != nil {
		return err
	} else if value != 0 {
		return fmt.Errorf("printing: image box requires Pixel Representation 0, got %d", value)
	}
	if _, ok := sequence.Get(tag.PlanarConfiguration); ok {
		return fmt.Errorf("printing: grayscale image box must not carry Planar Configuration")
	}
	return nil
}

func validateColorImageAttributes(sequence *dataset.Dataset) error {
	samplesPerPixel, err := uint16ImageAttribute(sequence, tag.SamplesPerPixel)
	if err != nil {
		return err
	}
	if samplesPerPixel != 3 {
		return fmt.Errorf("printing: color image box requires Samples per Pixel 3, got %d", samplesPerPixel)
	}
	photometric, err := stringImageAttribute(sequence, tag.PhotometricInterpretation)
	if err != nil {
		return err
	}
	if photometric != "RGB" {
		return fmt.Errorf("printing: color image box requires Photometric Interpretation RGB, got %q", photometric)
	}
	if value, err := uint16ImageAttribute(sequence, tag.PlanarConfiguration); err != nil {
		return err
	} else if value != 1 {
		return fmt.Errorf("printing: color image box requires Planar Configuration 1, got %d", value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.BitsAllocated); err != nil {
		return err
	} else if value != 8 {
		return fmt.Errorf("printing: color image box requires Bits Allocated 8, got %d", value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.BitsStored); err != nil {
		return err
	} else if value != 8 {
		return fmt.Errorf("printing: color image box requires Bits Stored 8, got %d", value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.HighBit); err != nil {
		return err
	} else if value != 7 {
		return fmt.Errorf("printing: color image box requires High Bit 7, got %d", value)
	}
	if value, err := uint16ImageAttribute(sequence, tag.PixelRepresentation); err != nil {
		return err
	} else if value != 0 {
		return fmt.Errorf("printing: image box requires Pixel Representation 0, got %d", value)
	}
	return nil
}

func uint16ImageAttribute(sequence *dataset.Dataset, attributeTag *tag.Tag) (uint16, error) {
	value, err := sequence.GetUInt16(attributeTag, 0)
	if err != nil {
		return 0, fmt.Errorf("printing: image box sequence attribute %s: %w", attributeTag, err)
	}
	return value, nil
}

func stringImageAttribute(sequence *dataset.Dataset, attributeTag *tag.Tag) (string, error) {
	value, ok := sequence.GetString(attributeTag)
	if !ok {
		return "", fmt.Errorf("printing: image box sequence attribute %s is not a string", attributeTag)
	}
	return strings.TrimSpace(value), nil
}

// SetImageData sets the image data for the Image Box
func (ib *ImageBox) SetImageData(imageData []byte) {
	ib.setLegacyImageData(imageData)
	if ib.imageSequence != nil {
		if len(imageData) == 0 {
			ib.imageSequence.Remove(tag.PixelData)
		} else {
			pixelData := element.Element(element.NewOtherByte(tag.PixelData, append([]byte(nil), imageData...)))
			if existing := ib.imageSequence.GetOrNil(tag.PixelData); existing != nil && existing.ValueRepresentation() == vr.OW {
				pixelData = element.NewOtherWord(tag.PixelData, append([]byte(nil), imageData...))
			}
			_ = ib.imageSequence.AddOrUpdate(pixelData)
		}
	}
}

func (ib *ImageBox) setLegacyImageData(imageData []byte) {
	if ib.IsColor {
		ib.PreformattedColorImageSequence = append([]byte(nil), imageData...)
		ib.PreformattedGrayscaleImageSequence = nil
	} else {
		ib.PreformattedGrayscaleImageSequence = append([]byte(nil), imageData...)
		ib.PreformattedColorImageSequence = nil
	}
}

func pixelDataBytes(ds *dataset.Dataset) []byte {
	if ds == nil {
		return nil
	}
	pixelData := ds.GetOrNil(tag.PixelData)
	if pixelData == nil || pixelData.Buffer() == nil {
		return nil
	}
	return append([]byte(nil), pixelData.Buffer().Data()...)
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
	if ib.imageSequence != nil {
		ib.imageSequence.Remove(tag.PixelData)
	}
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
