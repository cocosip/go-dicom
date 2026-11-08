// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewImageBox(t *testing.T) {
	testCases := []struct {
		name           string
		sopInstanceUID string
		isColor        bool
		expectedClass  string
	}{
		{
			name:           "grayscale image box",
			sopInstanceUID: testSOPInstanceUID,
			isColor:        false,
			expectedClass:  SOPClassGrayscaleImageBox,
		},
		{
			name:           "color image box",
			sopInstanceUID: "1.2.3.4.6",
			isColor:        true,
			expectedClass:  SOPClassColorImageBox,
		},
		{
			name:           "empty UID generates placeholder",
			sopInstanceUID: "",
			isColor:        false,
			expectedClass:  SOPClassGrayscaleImageBox,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ib := NewImageBox(tc.sopInstanceUID, tc.isColor)

			if ib == nil {
				t.Fatal("NewImageBox returned nil")
			}

			if tc.sopInstanceUID != "" && ib.SOPInstanceUID != tc.sopInstanceUID {
				t.Errorf("Expected SOPInstanceUID='%s', got '%s'", tc.sopInstanceUID, ib.SOPInstanceUID)
			}

			if ib.SOPClassUID != tc.expectedClass {
				t.Errorf("Expected SOPClassUID='%s', got '%s'", tc.expectedClass, ib.SOPClassUID)
			}

			if ib.IsColor != tc.isColor {
				t.Errorf("Expected IsColor=%v, got %v", tc.isColor, ib.IsColor)
			}

			// Check defaults
			if ib.ImageBoxPosition != 1 {
				t.Errorf("Expected default position=1, got %d", ib.ImageBoxPosition)
			}

			if ib.Polarity != PolarityNormal {
				t.Errorf("Expected default Polarity=NORMAL, got '%s'", ib.Polarity)
			}

			if ib.MagnificationType != MagnificationReplicate {
				t.Errorf("Expected default MagnificationType=REPLICATE, got '%s'", ib.MagnificationType)
			}
		})
	}
}

func TestImageBox_SetImageData(t *testing.T) {
	testCases := []struct {
		name    string
		isColor bool
	}{
		{"grayscale", false},
		{"color", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ib := NewImageBox(testSOPInstanceUID, tc.isColor)

			testData := []byte{1, 2, 3, 4, 5, 6, 7, 8}
			ib.SetImageData(testData)

			if !ib.HasImageData() {
				t.Error("HasImageData should return true after SetImageData")
			}

			retrievedData := ib.GetImageData()
			if len(retrievedData) != len(testData) {
				t.Errorf("Expected data length=%d, got %d", len(testData), len(retrievedData))
			}

			for i, b := range testData {
				if retrievedData[i] != b {
					t.Errorf("Data mismatch at index %d: expected %d, got %d", i, b, retrievedData[i])
				}
			}

			// Verify data was copied, not referenced
			testData[0] = 99
			if ib.GetImageData()[0] == 99 {
				t.Error("SetImageData should copy data, not reference it")
			}
		})
	}
}

func TestImageBox_GetImageData(t *testing.T) {
	// Grayscale
	grayscaleIB := NewImageBox(testSOPInstanceUID, false)
	grayscaleData := []byte{10, 20, 30}
	grayscaleIB.SetImageData(grayscaleData)

	retrieved := grayscaleIB.GetImageData()
	if len(retrieved) != 3 || retrieved[0] != 10 || retrieved[1] != 20 || retrieved[2] != 30 {
		t.Error("Grayscale GetImageData returned incorrect data")
	}

	// Color
	colorIB := NewImageBox("1.2.3.4.6", true)
	colorData := []byte{100, 200}
	colorIB.SetImageData(colorData)

	retrieved = colorIB.GetImageData()
	if len(retrieved) != 2 || retrieved[0] != 100 || retrieved[1] != 200 {
		t.Error("Color GetImageData returned incorrect data")
	}

	// Empty
	emptyIB := NewImageBox("1.2.3.4.7", false)
	if emptyIB.GetImageData() != nil {
		t.Error("GetImageData should return nil for empty image box")
	}
}

func TestImageBox_HasImageData(t *testing.T) {
	ib := NewImageBox(testSOPInstanceUID, false)

	// Initially empty
	if ib.HasImageData() {
		t.Error("New image box should not have image data")
	}

	// After setting data
	ib.SetImageData([]byte{1, 2, 3})
	if !ib.HasImageData() {
		t.Error("Should have image data after SetImageData")
	}

	// After clearing
	ib.ClearImageData()
	if ib.HasImageData() {
		t.Error("Should not have image data after ClearImageData")
	}
}

func TestImageBox_ClearImageData(t *testing.T) {
	// Test grayscale
	grayscaleIB := NewImageBox(testSOPInstanceUID, false)
	grayscaleIB.SetImageData([]byte{1, 2, 3})
	grayscaleIB.ClearImageData()

	if grayscaleIB.HasImageData() {
		t.Error("Grayscale image box should not have data after clear")
	}
	if grayscaleIB.PreformattedGrayscaleImageSequence != nil {
		t.Error("PreformattedGrayscaleImageSequence should be nil after clear")
	}

	// Test color
	colorIB := NewImageBox("1.2.3.4.6", true)
	colorIB.SetImageData([]byte{1, 2, 3})
	colorIB.ClearImageData()

	if colorIB.HasImageData() {
		t.Error("Color image box should not have data after clear")
	}
	if colorIB.PreformattedColorImageSequence != nil {
		t.Error("PreformattedColorImageSequence should be nil after clear")
	}
}

func TestImageBox_FilmBox(t *testing.T) {
	fb := NewFilmBox(testSOPInstanceUID, testImageDisplayFormat2x2)
	ib := NewImageBox("1.2.3.4.5.1", false)

	// Initially no parent
	if ib.FilmBox() != nil {
		t.Error("New ImageBox should have no parent FilmBox")
	}

	// After adding to film box
	fb.AddImageBox(ib)
	if ib.FilmBox() != fb {
		t.Error("ImageBox should reference parent FilmBox after AddImageBox")
	}
}

func TestImageBox_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(*ImageBox)
		expected bool
	}{
		{
			name:     "valid image box",
			setup:    func(_ *ImageBox) {},
			expected: true,
		},
		{
			name: "valid with image data",
			setup: func(ib *ImageBox) {
				ib.SetImageData([]byte{1, 2, 3})
			},
			expected: true,
		},
		{
			name: "invalid: empty SOP Class UID",
			setup: func(ib *ImageBox) {
				ib.SOPClassUID = ""
			},
			expected: false,
		},
		{
			name: "invalid: empty SOP Instance UID",
			setup: func(ib *ImageBox) {
				ib.SOPInstanceUID = ""
			},
			expected: false,
		},
		{
			name: "invalid: position = 0",
			setup: func(ib *ImageBox) {
				ib.ImageBoxPosition = 0
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ib := NewImageBox(testSOPInstanceUID, false)
			tc.setup(ib)

			result := ib.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestImageBox_GetEffectiveMagnificationType(t *testing.T) {
	testCases := []struct {
		name           string
		imageBoxMagnif MagnificationType
		filmBoxMagnif  MagnificationType
		hasFilmBox     bool
		expectedMagnif MagnificationType
	}{
		{
			name:           "image box has its own setting",
			imageBoxMagnif: MagnificationBilinear,
			filmBoxMagnif:  MagnificationReplicate,
			hasFilmBox:     true,
			expectedMagnif: MagnificationBilinear,
		},
		{
			name:           "inherit from film box",
			imageBoxMagnif: "",
			filmBoxMagnif:  MagnificationCubic,
			hasFilmBox:     true,
			expectedMagnif: MagnificationCubic,
		},
		{
			name:           "no film box, use default",
			imageBoxMagnif: "",
			filmBoxMagnif:  "",
			hasFilmBox:     false,
			expectedMagnif: MagnificationReplicate,
		},
		{
			name:           "empty setting with film box",
			imageBoxMagnif: "",
			filmBoxMagnif:  MagnificationNone,
			hasFilmBox:     true,
			expectedMagnif: MagnificationNone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ib := NewImageBox(testSOPInstanceUID, false)
			ib.MagnificationType = tc.imageBoxMagnif

			if tc.hasFilmBox {
				fb := NewFilmBox(testSOPInstanceUID, testImageDisplayFormat2x2)
				fb.MagnificationType = tc.filmBoxMagnif
				fb.AddImageBox(ib)
			}

			result := ib.GetEffectiveMagnificationType()
			if result != tc.expectedMagnif {
				t.Errorf("Expected magnification=%s, got %s", tc.expectedMagnif, result)
			}
		})
	}
}

func TestImageBox_PolaritySettings(t *testing.T) {
	ib := NewImageBox(testSOPInstanceUID, false)

	// Test default
	if ib.Polarity != PolarityNormal {
		t.Errorf("Expected default Polarity=NORMAL, got %s", ib.Polarity)
	}

	// Test reverse
	ib.Polarity = PolarityReverse
	if ib.Polarity != PolarityReverse {
		t.Error("Failed to set Polarity to REVERSE")
	}
}

func TestImageBox_RequestedImageSize(t *testing.T) {
	ib := NewImageBox(testSOPInstanceUID, false)

	// Test default
	if ib.RequestedImageSize != "" {
		t.Errorf("Expected empty RequestedImageSize, got '%s'", ib.RequestedImageSize)
	}

	// Test setting
	ib.RequestedImageSize = "100\\150"
	if ib.RequestedImageSize != "100\\150" {
		t.Error("Failed to set RequestedImageSize")
	}
}

func TestImageBox_SmoothingType(t *testing.T) {
	ib := NewImageBox(testSOPInstanceUID, false)

	// Test default
	if ib.SmoothingType != "" {
		t.Errorf("Expected empty SmoothingType, got '%s'", ib.SmoothingType)
	}

	// Test setting (implementation-specific values)
	ib.SmoothingType = "MEDIUM"
	if ib.SmoothingType != "MEDIUM" {
		t.Error("Failed to set SmoothingType")
	}
}

func TestImageBoxValidateImageAttributesRequiresPixelMacro(t *testing.T) {
	ib := NewImageBox("2.25.700", false)
	ib.SetImageData([]byte{1, 2, 3, 4})
	if err := ib.ValidateImageAttributes(); err == nil {
		t.Fatal("ValidateImageAttributes() = nil, want error for PixelData-only sequence")
	}
}

type imageAttributeSet struct {
	samplesPerPixel     uint16
	photometric         string
	bitsAllocated       uint16
	bitsStored          uint16
	highBit             uint16
	pixelRepresentation uint16
	planarConfiguration *uint16
}

func uint16Pointer(value uint16) *uint16 {
	return &value
}

func imageSequenceDataset(t *testing.T, attributes imageAttributeSet) *dataset.Dataset {
	t.Helper()
	item := dataset.New()
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{2}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{attributes.samplesPerPixel}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{attributes.photometric}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{attributes.bitsAllocated}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{attributes.bitsStored}),
		element.NewUnsignedShort(tag.HighBit, []uint16{attributes.highBit}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{attributes.pixelRepresentation}),
		element.NewOtherByte(tag.PixelData, []byte{1, 2, 3, 4}),
	}
	if attributes.planarConfiguration != nil {
		elements = append(elements, element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{*attributes.planarConfiguration}))
	}
	for _, elem := range elements {
		if err := item.Add(elem); err != nil {
			t.Fatalf("build image item %s: %v", elem.Tag(), err)
		}
	}
	return item
}

// TestImageBoxValidateImageAttributesValues covers PS3.3 C.13.5, Table C.13-5.
func TestImageBoxValidateImageAttributesValues(t *testing.T) {
	grayscale := imageAttributeSet{
		samplesPerPixel: 1, photometric: "MONOCHROME2",
		bitsAllocated: 8, bitsStored: 8, highBit: 7,
	}
	color := imageAttributeSet{
		samplesPerPixel: 3, photometric: testPhotometricRGB,
		bitsAllocated: 8, bitsStored: 8, highBit: 7,
		planarConfiguration: uint16Pointer(1),
	}

	tests := []struct {
		name       string
		isColor    bool
		attributes imageAttributeSet
		wantErr    bool
	}{
		{name: "valid grayscale 8-bit", attributes: grayscale},
		{name: "valid grayscale MONOCHROME1", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.photometric = "MONOCHROME1"
			return attributes
		}()},
		{name: "valid grayscale 12-bit", attributes: imageAttributeSet{
			samplesPerPixel: 1, photometric: "MONOCHROME2",
			bitsAllocated: 16, bitsStored: 12, highBit: 11,
		}},
		{name: "valid color", isColor: true, attributes: color},
		{name: "grayscale Bits Stored 16", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.bitsAllocated, attributes.bitsStored, attributes.highBit = 16, 16, 15
			return attributes
		}(), wantErr: true},
		{name: "grayscale Bits Allocated mismatch", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.bitsAllocated = 16
			return attributes
		}(), wantErr: true},
		{name: "grayscale High Bit mismatch", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.highBit = 11
			return attributes
		}(), wantErr: true},
		{name: "grayscale Pixel Representation signed", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.pixelRepresentation = 1
			return attributes
		}(), wantErr: true},
		{name: "grayscale RGB photometric", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.photometric = testPhotometricRGB
			return attributes
		}(), wantErr: true},
		{name: "grayscale Samples per Pixel 3", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.samplesPerPixel = 3
			return attributes
		}(), wantErr: true},
		{name: "grayscale Planar Configuration present", attributes: func() imageAttributeSet {
			attributes := grayscale
			attributes.planarConfiguration = uint16Pointer(0)
			return attributes
		}(), wantErr: true},
		{name: "color Planar Configuration 0", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.planarConfiguration = uint16Pointer(0)
			return attributes
		}(), wantErr: true},
		{name: "color Planar Configuration missing", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.planarConfiguration = nil
			return attributes
		}(), wantErr: true},
		{name: "color Samples per Pixel 1", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.samplesPerPixel = 1
			return attributes
		}(), wantErr: true},
		{name: "color Photometric Interpretation YBR_FULL", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.photometric = "YBR_FULL"
			return attributes
		}(), wantErr: true},
		{name: "color 16-bit", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.bitsAllocated, attributes.bitsStored, attributes.highBit = 16, 12, 11
			return attributes
		}(), wantErr: true},
		{name: "color Pixel Representation signed", isColor: true, attributes: func() imageAttributeSet {
			attributes := color
			attributes.pixelRepresentation = 1
			return attributes
		}(), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := NewImageBox("2.25.701", tt.isColor)
			if err := box.SetImageSequence(imageSequenceDataset(t, tt.attributes)); err != nil {
				t.Fatalf("SetImageSequence() error = %v", err)
			}
			err := box.ValidateImageAttributes()
			if tt.wantErr && err == nil {
				t.Fatal("ValidateImageAttributes() = nil, want error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("ValidateImageAttributes() error = %v", err)
			}
		})
	}
}

// TestImageBoxColorDatasetOmitsGrayscaleOnlyAttributes covers PS3.4 Table H.4-11,
// which does not list Minimum/Maximum Density or Configuration Information.
func TestImageBoxColorDatasetOmitsGrayscaleOnlyAttributes(t *testing.T) {
	maxDensity, minDensity := uint16(200), uint16(10)
	configuration := "CS1"

	color := NewImageBox("2.25.702", true)
	color.MaxDensity = &maxDensity
	color.MinDensity = &minDensity
	color.ConfigurationInformation = &configuration
	color.SetImageData([]byte{1, 2, 3, 4, 5, 6})
	colorDataset, err := color.ToDataset()
	if err != nil {
		t.Fatalf("color ToDataset() error = %v", err)
	}
	for _, forbidden := range []*tag.Tag{tag.MaxDensity, tag.MinDensity, tag.ConfigurationInformation} {
		if _, ok := colorDataset.Get(forbidden); ok {
			t.Errorf("color Image Box N-SET dataset contains %s", forbidden)
		}
	}

	grayscale := NewImageBox("2.25.703", false)
	grayscale.MaxDensity = &maxDensity
	grayscale.MinDensity = &minDensity
	grayscale.ConfigurationInformation = &configuration
	grayscale.SetImageData([]byte{1, 2, 3, 4})
	grayscaleDataset, err := grayscale.ToDataset()
	if err != nil {
		t.Fatalf("grayscale ToDataset() error = %v", err)
	}
	for _, required := range []*tag.Tag{tag.MaxDensity, tag.MinDensity, tag.ConfigurationInformation} {
		if _, ok := grayscaleDataset.Get(required); !ok {
			t.Errorf("grayscale Image Box N-SET dataset is missing %s", required)
		}
	}
}

