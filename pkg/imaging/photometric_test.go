// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"testing"
)

func TestParsePhotometricInterpretation(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  *PhotometricInterpretation
		expectErr bool
	}{
		{
			name:     "MONOCHROME1",
			input:    "MONOCHROME1",
			expected: Monochrome1,
		},
		{
			name:     "MONOCHROME2",
			input:    "MONOCHROME2",
			expected: Monochrome2,
		},
		{
			name:     "MONOCHROME (alias)",
			input:    "MONOCHROME",
			expected: Monochrome2,
		},
		{
			name:     "PALETTE COLOR",
			input:    "PALETTE COLOR",
			expected: PaletteColor,
		},
		{
			name:     "RGB",
			input:    "RGB",
			expected: RGBPhotometric,
		},
		{
			name:     "YBR_FULL",
			input:    "YBR_FULL",
			expected: YbrFull,
		},
		{
			name:     "YBR_FULL_422",
			input:    "YBR_FULL_422",
			expected: YbrFull422,
		},
		{
			name:     "YBR_PARTIAL_422",
			input:    "YBR_PARTIAL_422",
			expected: YbrPartial422,
		},
		{
			name:     "YBR_PARTIAL_420",
			input:    "YBR_PARTIAL_420",
			expected: YbrPartial420,
		},
		{
			name:     "YBR_ICT",
			input:    "YBR_ICT",
			expected: YbrIct,
		},
		{
			name:     "YBR_RCT",
			input:    "YBR_RCT",
			expected: YbrRct,
		},
		{
			name:     "with spaces",
			input:    "  RGB  ",
			expected: RGBPhotometric,
		},
		{
			name:     "with null characters",
			input:    "RGB\x00",
			expected: RGBPhotometric,
		},
		{
			name:      "invalid",
			input:     "INVALID",
			expectErr: true,
		},
		{
			name:      "empty",
			input:     "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParsePhotometricInterpretation(tt.input)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error for input %q: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("ParsePhotometricInterpretation(%q) = %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestPhotometricInterpretation_Properties(t *testing.T) {
	tests := []struct {
		name        string
		pi          *PhotometricInterpretation
		isColor     bool
		isPalette   bool
		isYBR       bool
		description string
	}{
		{
			name:        "Monochrome1",
			pi:          Monochrome1,
			isColor:     false,
			isPalette:   false,
			isYBR:       false,
			description: "Monochrome 1",
		},
		{
			name:        "Monochrome2",
			pi:          Monochrome2,
			isColor:     false,
			isPalette:   false,
			isYBR:       false,
			description: "Monochrome 2",
		},
		{
			name:        "PaletteColor",
			pi:          PaletteColor,
			isColor:     true,
			isPalette:   true,
			isYBR:       false,
			description: "Palette Color",
		},
		{
			name:        "RGB",
			pi:          RGBPhotometric,
			isColor:     true,
			isPalette:   false,
			isYBR:       false,
			description: "RGB",
		},
		{
			name:        "YbrFull",
			pi:          YbrFull,
			isColor:     true,
			isPalette:   false,
			isYBR:       true,
			description: "YBR Full",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pi.IsColor != tt.isColor {
				t.Errorf("%s.IsColor = %v, want %v", tt.name, tt.pi.IsColor, tt.isColor)
			}
			if tt.pi.IsPalette != tt.isPalette {
				t.Errorf("%s.IsPalette = %v, want %v", tt.name, tt.pi.IsPalette, tt.isPalette)
			}
			if tt.pi.IsYBR != tt.isYBR {
				t.Errorf("%s.IsYBR = %v, want %v", tt.name, tt.pi.IsYBR, tt.isYBR)
			}
			if tt.pi.String() != tt.description {
				t.Errorf("%s.String() = %q, want %q", tt.name, tt.pi.String(), tt.description)
			}
		})
	}
}

func TestPhotometricInterpretation_Equals(t *testing.T) {
	tests := []struct {
		name     string
		pi1      *PhotometricInterpretation
		pi2      *PhotometricInterpretation
		expected bool
	}{
		{
			name:     "same instance",
			pi1:      Monochrome1,
			pi2:      Monochrome1,
			expected: true,
		},
		{
			name:     "different values",
			pi1:      Monochrome1,
			pi2:      Monochrome2,
			expected: false,
		},
		{
			name:     "both nil",
			pi1:      nil,
			pi2:      nil,
			expected: true,
		},
		{
			name:     "one nil",
			pi1:      Monochrome1,
			pi2:      nil,
			expected: false,
		},
		{
			name:     "RGB equality",
			pi1:      RGBPhotometric,
			pi2:      RGBPhotometric,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pi1.Equals(tt.pi2)
			if result != tt.expected {
				t.Errorf("Equals() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMustParsePhotometricInterpretation(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		result := MustParsePhotometricInterpretation("RGB")
		if result != RGBPhotometric {
			t.Errorf("MustParsePhotometricInterpretation(\"RGB\") = %v, want %v",
				result, RGBPhotometric)
		}
	})

	t.Run("invalid input panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustParsePhotometricInterpretation did not panic on invalid input")
			}
		}()
		MustParsePhotometricInterpretation("INVALID")
	})
}
