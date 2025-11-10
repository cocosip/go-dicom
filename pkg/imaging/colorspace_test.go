// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "testing"

func TestNewComponent(t *testing.T) {
	comp := NewComponent("Red", 1, 1)
	if comp.Name != "Red" {
		t.Errorf("Name = %q, want %q", comp.Name, "Red")
	}
	if comp.SubSampleX != 1 {
		t.Errorf("SubSampleX = %d, want %d", comp.SubSampleX, 1)
	}
	if comp.SubSampleY != 1 {
		t.Errorf("SubSampleY = %d, want %d", comp.SubSampleY, 1)
	}
}

func TestNewColorSpace(t *testing.T) {
	comp1 := NewComponent("Red", 1, 1)
	comp2 := NewComponent("Green", 1, 1)

	cs := NewColorSpace("Test", comp1, comp2)

	if cs.Name != "Test" {
		t.Errorf("Name = %q, want %q", cs.Name, "Test")
	}
	if len(cs.Components) != 2 {
		t.Errorf("Components length = %d, want %d", len(cs.Components), 2)
	}
}

func TestColorSpace_String(t *testing.T) {
	tests := []struct {
		name     string
		cs       *ColorSpace
		expected string
	}{
		{
			name:     "RGB",
			cs:       RGB,
			expected: "RGB",
		},
		{
			name:     "Grayscale",
			cs:       Grayscale,
			expected: "Grayscale",
		},
		{
			name:     "YCbCrJPEG",
			cs:       YCbCrJPEG,
			expected: "Y'CbCr 4:4:4 (JPEG)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cs.String()
			if result != tt.expected {
				t.Errorf("String() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestColorSpace_Equals(t *testing.T) {
	tests := []struct {
		name     string
		cs1      *ColorSpace
		cs2      *ColorSpace
		expected bool
	}{
		{
			name:     "same instance",
			cs1:      RGB,
			cs2:      RGB,
			expected: true,
		},
		{
			name:     "different color spaces",
			cs1:      RGB,
			cs2:      Grayscale,
			expected: false,
		},
		{
			name:     "both nil",
			cs1:      nil,
			cs2:      nil,
			expected: true,
		},
		{
			name:     "one nil",
			cs1:      RGB,
			cs2:      nil,
			expected: false,
		},
		{
			name:     "same name different instance",
			cs1:      RGB,
			cs2:      NewColorSpace("RGB", NewComponent("R", 1, 1)),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cs1.Equals(tt.cs2)
			if result != tt.expected {
				t.Errorf("Equals() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStandardColorSpaces(t *testing.T) {
	tests := []struct {
		name          string
		cs            *ColorSpace
		expectedName  string
		expectedComps int
		firstCompName string
	}{
		{
			name:          "OneBit",
			cs:            OneBit,
			expectedName:  "1-bit",
			expectedComps: 1,
			firstCompName: "Value",
		},
		{
			name:          "Grayscale",
			cs:            Grayscale,
			expectedName:  "Grayscale",
			expectedComps: 1,
			firstCompName: "Value",
		},
		{
			name:          "Indexed",
			cs:            Indexed,
			expectedName:  "Indexed",
			expectedComps: 1,
			firstCompName: "Value",
		},
		{
			name:          "RGB",
			cs:            RGB,
			expectedName:  "RGB",
			expectedComps: 3,
			firstCompName: "Red",
		},
		{
			name:          "BGR",
			cs:            BGR,
			expectedName:  "BGR",
			expectedComps: 3,
			firstCompName: "Blue",
		},
		{
			name:          "RGBA",
			cs:            RGBA,
			expectedName:  "RGBA",
			expectedComps: 4,
			firstCompName: "Red",
		},
		{
			name:          "YCbCrJPEG",
			cs:            YCbCrJPEG,
			expectedName:  "Y'CbCr 4:4:4 (JPEG)",
			expectedComps: 3,
			firstCompName: "Luminance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cs.Name != tt.expectedName {
				t.Errorf("Name = %q, want %q", tt.cs.Name, tt.expectedName)
			}
			if len(tt.cs.Components) != tt.expectedComps {
				t.Errorf("Components length = %d, want %d", len(tt.cs.Components), tt.expectedComps)
			}
			if len(tt.cs.Components) > 0 && tt.cs.Components[0].Name != tt.firstCompName {
				t.Errorf("First component name = %q, want %q",
					tt.cs.Components[0].Name, tt.firstCompName)
			}
		})
	}
}
