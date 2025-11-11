// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"
)

func TestNewFilmBox(t *testing.T) {
	fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")

	if fb == nil {
		t.Fatal("NewFilmBox returned nil")
	}

	if fb.SOPInstanceUID != "1.2.3.4.5" {
		t.Errorf("Expected SOPInstanceUID='1.2.3.4.5', got '%s'", fb.SOPInstanceUID)
	}

	if fb.ImageDisplayFormat != "STANDARD\\2,2" {
		t.Errorf("Expected ImageDisplayFormat='STANDARD\\2,2', got '%s'", fb.ImageDisplayFormat)
	}

	// Check defaults
	if fb.FilmOrientation != FilmOrientationPortrait {
		t.Errorf("Expected default FilmOrientation=PORTRAIT, got '%s'", fb.FilmOrientation)
	}

	if fb.FilmSizeID != FilmSize14INX17IN {
		t.Errorf("Expected default FilmSizeID=14INX17IN, got '%s'", fb.FilmSizeID)
	}

	if fb.MagnificationType != MagnificationReplicate {
		t.Errorf("Expected default MagnificationType=REPLICATE, got '%s'", fb.MagnificationType)
	}

	if fb.Trim != TrimNo {
		t.Errorf("Expected default Trim=NO, got '%s'", fb.Trim)
	}

	if len(fb.BasicImageBoxes) != 0 {
		t.Errorf("Expected empty BasicImageBoxes, got %d", len(fb.BasicImageBoxes))
	}
}

func TestParseImageDisplayFormat(t *testing.T) {
	testCases := []struct {
		format        string
		expectedCount int
		expectError   bool
	}{
		// STANDARD format
		{"STANDARD\\2,2", 4, false},
		{"STANDARD\\3,3", 9, false},
		{"STANDARD\\1,1", 1, false},
		{"STANDARD\\4,2", 8, false},
		{"STANDARD\\2", 0, true},   // Missing dimension
		{"STANDARD\\a,b", 0, true}, // Invalid numbers

		// ROW format
		{"ROW\\1,2,1", 4, false},
		{"ROW\\2,2,2", 6, false},
		{"ROW\\1", 1, false},
		{"ROW\\3,3,3,3", 12, false},
		{"ROW\\a", 0, true}, // Invalid number

		// COL format
		{"COL\\1,2,1", 4, false},
		{"COL\\2,2", 4, false},
		{"COL\\3", 3, false},
		{"COL\\x", 0, true}, // Invalid number

		// SLIDE format
		{"SLIDE\\1", 20, false},

		// SUPERSLIDE format
		{"SUPERSLIDE\\1", 15, false},

		// CUSTOM format
		{"CUSTOM\\1", 1, false},

		// Invalid formats
		{"INVALID", 0, true},
		{"", 0, true},
		{"STANDARD", 0, true}, // Missing separator
	}

	for _, tc := range testCases {
		t.Run(tc.format, func(t *testing.T) {
			count, err := ParseImageDisplayFormat(tc.format)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error for format '%s', got nil", tc.format)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for format '%s': %v", tc.format, err)
				}
				if count != tc.expectedCount {
					t.Errorf("Format '%s': expected count=%d, got %d", tc.format, tc.expectedCount, count)
				}
			}
		})
	}
}

func TestFilmBox_InitializeImageBoxes(t *testing.T) {
	testCases := []struct {
		name          string
		format        string
		expectedCount int
		expectError   bool
	}{
		{"2x2 grid", "STANDARD\\2,2", 4, false},
		{"3x3 grid", "STANDARD\\3,3", 9, false},
		{"row format", "ROW\\1,2,1", 4, false},
		{"invalid format", "INVALID", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fb := NewFilmBox("1.2.3.4.5", tc.format)
			err := fb.InitializeImageBoxes()

			if tc.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if len(fb.BasicImageBoxes) != tc.expectedCount {
					t.Errorf("Expected %d image boxes, got %d", tc.expectedCount, len(fb.BasicImageBoxes))
				}

				// Verify all image boxes are properly initialized
				for i, ib := range fb.BasicImageBoxes {
					if ib == nil {
						t.Errorf("Image box at index %d is nil", i)
						continue
					}

					if int(ib.ImageBoxPosition) != (i + 1) {
						t.Errorf("Image box %d: expected position=%d, got %d", i, i+1, ib.ImageBoxPosition)
					}

					if ib.FilmBox() != fb {
						t.Errorf("Image box %d: does not reference parent film box", i)
					}
				}
			}
		})
	}
}

func TestFilmBox_AddImageBox(t *testing.T) {
	fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")

	if fb.ImageBoxCount() != 0 {
		t.Errorf("Expected initial count=0, got %d", fb.ImageBoxCount())
	}

	// Add first image box
	ib1 := NewImageBox("1.2.3.4.5.1", false)
	fb.AddImageBox(ib1)

	if fb.ImageBoxCount() != 1 {
		t.Errorf("Expected count=1 after adding, got %d", fb.ImageBoxCount())
	}

	if ib1.FilmBox() != fb {
		t.Error("ImageBox should reference parent FilmBox")
	}

	// Add second image box
	ib2 := NewImageBox("1.2.3.4.5.2", false)
	fb.AddImageBox(ib2)

	if fb.ImageBoxCount() != 2 {
		t.Errorf("Expected count=2 after adding second, got %d", fb.ImageBoxCount())
	}

	// Test adding nil
	fb.AddImageBox(nil)
	if fb.ImageBoxCount() != 2 {
		t.Errorf("Adding nil should not change count, got %d", fb.ImageBoxCount())
	}
}

func TestFilmBox_GetImageBox(t *testing.T) {
    fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
    if err := fb.InitializeImageBoxes(); err != nil {
        t.Fatalf("InitializeImageBoxes error: %v", err)
    }

	// Test valid positions (1-based)
	ib1 := fb.GetImageBox(1)
	if ib1 == nil {
		t.Error("GetImageBox(1) should return first image box")
	}
	if ib1 != nil && ib1.ImageBoxPosition != 1 {
		t.Errorf("Expected position=1, got %d", ib1.ImageBoxPosition)
	}

	ib4 := fb.GetImageBox(4)
	if ib4 == nil {
		t.Error("GetImageBox(4) should return fourth image box")
	}
	if ib4 != nil && ib4.ImageBoxPosition != 4 {
		t.Errorf("Expected position=4, got %d", ib4.ImageBoxPosition)
	}

	// Test invalid positions
	if fb.GetImageBox(0) != nil {
		t.Error("GetImageBox(0) should return nil")
	}

	if fb.GetImageBox(5) != nil {
		t.Error("GetImageBox(5) should return nil for out of range")
	}
}

func TestFilmBox_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(*FilmBox)
		expected bool
	}{
		{
			name: "valid film box",
            setup: func(fb *FilmBox) {
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
            },
			expected: true,
		},
		{
			name: "invalid: empty image display format",
            setup: func(fb *FilmBox) {
                fb.ImageDisplayFormat = ""
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
            },
			expected: false,
		},
		{
			name: "invalid: no image boxes",
			setup: func(_ *FilmBox) {
				// Don't initialize image boxes
			},
			expected: false,
		},
		{
			name: "invalid: image box is invalid",
            setup: func(fb *FilmBox) {
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
            	// Corrupt an image box
                if len(fb.BasicImageBoxes) > 0 {
                    fb.BasicImageBoxes[0].SOPInstanceUID = ""
                }
            },
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
			tc.setup(fb)

			result := fb.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFilmBox_FilmSession(t *testing.T) {
	fs := NewFilmSession("", "", false)
	fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")

	// Initially no parent
	if fb.FilmSession() != nil {
		t.Error("New FilmBox should have no parent FilmSession")
	}

	// After adding to session
	fs.AddFilmBox(fb)
	if fb.FilmSession() != fs {
		t.Error("FilmBox should reference parent FilmSession after AddFilmBox")
	}
}

func TestFilmBox_ColorInheritance(t *testing.T) {
	// Grayscale session
	grayscaleFS := NewFilmSession("", "", false)
	grayscaleFB := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
	grayscaleFS.AddFilmBox(grayscaleFB)
    if err := grayscaleFB.InitializeImageBoxes(); err != nil {
        t.Fatalf("InitializeImageBoxes error: %v", err)
    }

	for _, ib := range grayscaleFB.BasicImageBoxes {
		if ib.IsColor {
			t.Error("Image box in grayscale film box should have IsColor=false")
		}
		if ib.SOPClassUID != SOPClassGrayscaleImageBox {
			t.Error("Image box in grayscale film box should use grayscale SOP class")
		}
	}

	// Color session
	colorFS := NewFilmSession("", "", true)
	colorFB := NewFilmBox("1.2.3.4.6", "STANDARD\\2,2")
	colorFS.AddFilmBox(colorFB)
    if err := colorFB.InitializeImageBoxes(); err != nil {
        t.Fatalf("InitializeImageBoxes error: %v", err)
    }

	for _, ib := range colorFB.BasicImageBoxes {
		if !ib.IsColor {
			t.Error("Image box in color film box should have IsColor=true")
		}
		if ib.SOPClassUID != SOPClassColorImageBox {
			t.Error("Image box in color film box should use color SOP class")
		}
	}
}
