// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"
)

func TestNewFilmSession(t *testing.T) {
	testCases := []struct {
		name           string
		sopClassUID    string
		sopInstanceUID string
		isColor        bool
		expectDefaults bool
	}{
		{
			name:           "create grayscale session with UIDs",
			sopClassUID:    "1.2.840.10008.5.1.1.1",
			sopInstanceUID: "1.2.3.4.5",
			isColor:        false,
			expectDefaults: true,
		},
		{
			name:           "create color session with empty UIDs",
			sopClassUID:    "",
			sopInstanceUID: "",
			isColor:        true,
			expectDefaults: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFilmSession(tc.sopClassUID, tc.sopInstanceUID, tc.isColor)

			if fs == nil {
				t.Fatal("NewFilmSession returned nil")
			}

			if tc.sopClassUID != "" && fs.SOPClassUID != tc.sopClassUID {
				t.Errorf("Expected SOPClassUID='%s', got '%s'", tc.sopClassUID, fs.SOPClassUID)
			}

			if tc.sopInstanceUID != "" && fs.SOPInstanceUID != tc.sopInstanceUID {
				t.Errorf("Expected SOPInstanceUID='%s', got '%s'", tc.sopInstanceUID, fs.SOPInstanceUID)
			}

			if fs.IsColor != tc.isColor {
				t.Errorf("Expected IsColor=%v, got %v", tc.isColor, fs.IsColor)
			}

			if tc.expectDefaults {
				if fs.FilmDestination != FilmDestinationProcessor {
					t.Errorf("Expected default FilmDestination=PROCESSOR, got '%s'", fs.FilmDestination)
				}

				if fs.MediumType != MediumTypeClearFilm {
					t.Errorf("Expected default MediumType=CLEAR FILM, got '%s'", fs.MediumType)
				}

				if fs.PrintPriority != PrintPriorityMed {
					t.Errorf("Expected default PrintPriority=MED, got '%s'", fs.PrintPriority)
				}

				if fs.NumberOfCopies != 1 {
					t.Errorf("Expected default NumberOfCopies=1, got %d", fs.NumberOfCopies)
				}

				if len(fs.BasicFilmBoxes) != 0 {
					t.Errorf("Expected empty BasicFilmBoxes, got %d", len(fs.BasicFilmBoxes))
				}

				if len(fs.PresentationLUTs) != 0 {
					t.Errorf("Expected empty PresentationLUTs, got %d", len(fs.PresentationLUTs))
				}
			}
		})
	}
}

func TestFilmSession_AddFilmBox(t *testing.T) {
	fs := NewFilmSession("", "", false)

	if fs.FilmBoxCount() != 0 {
		t.Errorf("Expected initial count=0, got %d", fs.FilmBoxCount())
	}

	// Add first film box
	fb1 := NewFilmBox("1.2.3.4.5.1", "STANDARD\\2,2")
	fs.AddFilmBox(fb1)

	if fs.FilmBoxCount() != 1 {
		t.Errorf("Expected count=1 after adding, got %d", fs.FilmBoxCount())
	}

	if fb1.FilmSession() != fs {
		t.Error("FilmBox should reference parent FilmSession")
	}

	// Add second film box
	fb2 := NewFilmBox("1.2.3.4.5.2", "STANDARD\\1,1")
	fs.AddFilmBox(fb2)

	if fs.FilmBoxCount() != 2 {
		t.Errorf("Expected count=2 after adding second, got %d", fs.FilmBoxCount())
	}

	// Test adding nil
	fs.AddFilmBox(nil)
	if fs.FilmBoxCount() != 2 {
		t.Errorf("Adding nil should not change count, got %d", fs.FilmBoxCount())
	}
}

func TestFilmSession_AddPresentationLUT(t *testing.T) {
	fs := NewFilmSession("", "", false)

	if fs.PresentationLUTCount() != 0 {
		t.Errorf("Expected initial count=0, got %d", fs.PresentationLUTCount())
	}

	// Add LUT
	lut := NewPresentationLUT("1.2.3.4.5")
	fs.AddPresentationLUT(lut)

	if fs.PresentationLUTCount() != 1 {
		t.Errorf("Expected count=1 after adding, got %d", fs.PresentationLUTCount())
	}

	// Test adding nil
	fs.AddPresentationLUT(nil)
	if fs.PresentationLUTCount() != 1 {
		t.Errorf("Adding nil should not change count, got %d", fs.PresentationLUTCount())
	}
}

func TestFilmSession_GetFilmBox(t *testing.T) {
	fs := NewFilmSession("", "", false)

	fb1 := NewFilmBox("1.2.3.4.5.1", "STANDARD\\2,2")
	fb2 := NewFilmBox("1.2.3.4.5.2", "STANDARD\\1,1")

	fs.AddFilmBox(fb1)
	fs.AddFilmBox(fb2)

	// Test valid indices
	retrieved1 := fs.GetFilmBox(0)
	if retrieved1 != fb1 {
		t.Error("GetFilmBox(0) should return first film box")
	}

	retrieved2 := fs.GetFilmBox(1)
	if retrieved2 != fb2 {
		t.Error("GetFilmBox(1) should return second film box")
	}

	// Test invalid indices
	if fs.GetFilmBox(-1) != nil {
		t.Error("GetFilmBox(-1) should return nil")
	}

	if fs.GetFilmBox(2) != nil {
		t.Error("GetFilmBox(2) should return nil for out of range")
	}
}

func TestFilmSession_GetPresentationLUT(t *testing.T) {
	fs := NewFilmSession("", "", false)

	lut1 := NewPresentationLUT("1.2.3.4.5.1")
	lut2 := NewPresentationLUT("1.2.3.4.5.2")

	fs.AddPresentationLUT(lut1)
	fs.AddPresentationLUT(lut2)

	// Test valid indices
	retrieved1 := fs.GetPresentationLUT(0)
	if retrieved1 != lut1 {
		t.Error("GetPresentationLUT(0) should return first LUT")
	}

	retrieved2 := fs.GetPresentationLUT(1)
	if retrieved2 != lut2 {
		t.Error("GetPresentationLUT(1) should return second LUT")
	}

	// Test invalid indices
	if fs.GetPresentationLUT(-1) != nil {
		t.Error("GetPresentationLUT(-1) should return nil")
	}

	if fs.GetPresentationLUT(2) != nil {
		t.Error("GetPresentationLUT(2) should return nil for out of range")
	}
}

func TestFilmSession_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(*FilmSession)
		expected bool
	}{
		{
			name: "valid session with one film box",
			setup: func(fs *FilmSession) {
                fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
				fs.AddFilmBox(fb)
			},
			expected: true,
		},
		{
			name: "invalid: no film boxes",
			setup: func(_ *FilmSession) {
				// Don't add any film boxes
			},
			expected: false,
		},
		{
			name: "invalid: number of copies = 0",
			setup: func(fs *FilmSession) {
                fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
				fs.AddFilmBox(fb)
				fs.NumberOfCopies = 0
			},
			expected: false,
		},
		{
			name: "invalid: number of copies < 0",
            setup: func(fs *FilmSession) {
                fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
                if err := fb.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
                fs.AddFilmBox(fb)
                fs.NumberOfCopies = -1
            },
			expected: false,
		},
		{
			name: "invalid: film box is invalid",
			setup: func(fs *FilmSession) {
				fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
				// Don't initialize image boxes - makes it invalid
				fs.AddFilmBox(fb)
			},
			expected: false,
		},
		{
			name: "valid: multiple film boxes",
			setup: func(fs *FilmSession) {
                fb1 := NewFilmBox("1.2.3.4.5.1", "STANDARD\\2,2")
                if err := fb1.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
				fs.AddFilmBox(fb1)

                fb2 := NewFilmBox("1.2.3.4.5.2", "STANDARD\\1,1")
                if err := fb2.InitializeImageBoxes(); err != nil {
                    t.Fatalf("InitializeImageBoxes error: %v", err)
                }
				fs.AddFilmBox(fb2)
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFilmSession("", "", false)
			tc.setup(fs)

			result := fs.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid: expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFilmSession_ColorSupport(t *testing.T) {
	// Test grayscale session
	grayscaleFS := NewFilmSession("", "", false)
	if grayscaleFS.IsColor {
		t.Error("Grayscale session should have IsColor=false")
	}

	// Test color session
	colorFS := NewFilmSession("", "", true)
	if !colorFS.IsColor {
		t.Error("Color session should have IsColor=true")
	}

	// Verify film boxes inherit color setting
	fb := NewFilmBox("1.2.3.4.5", "STANDARD\\2,2")
	colorFS.AddFilmBox(fb)
    if err := fb.InitializeImageBoxes(); err != nil {
        t.Fatalf("InitializeImageBoxes error: %v", err)
    }

	if len(fb.BasicImageBoxes) > 0 {
		ib := fb.BasicImageBoxes[0]
		if !ib.IsColor {
			t.Error("Image box in color session should have IsColor=true")
		}
		if ib.SOPClassUID != SOPClassColorImageBox {
			t.Error("Image box in color session should use color SOP class")
		}
	}
}
