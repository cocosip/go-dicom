// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"
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
			sopInstanceUID: "1.2.3.4.5",
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
			ib := NewImageBox("1.2.3.4.5", tc.isColor)

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
	grayscaleIB := NewImageBox("1.2.3.4.5", false)
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
	ib := NewImageBox("1.2.3.4.5", false)

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
	grayscaleIB := NewImageBox("1.2.3.4.5", false)
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
	fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
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
			setup:    func(ib *ImageBox) {},
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
			ib := NewImageBox("1.2.3.4.5", false)
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
		name               string
		imageBoxMagnif     MagnificationType
		filmBoxMagnif      MagnificationType
		hasFilmBox         bool
		expectedMagnif     MagnificationType
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
			ib := NewImageBox("1.2.3.4.5", false)
			ib.MagnificationType = tc.imageBoxMagnif

			if tc.hasFilmBox {
				fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
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
	ib := NewImageBox("1.2.3.4.5", false)

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
	ib := NewImageBox("1.2.3.4.5", false)

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
	ib := NewImageBox("1.2.3.4.5", false)

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
