// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "testing"

func TestPixelRepresentation(t *testing.T) {
	tests := []struct {
		name     string
		pr       PixelRepresentation
		isSigned bool
		str      string
	}{
		{
			name:     "unsigned",
			pr:       UnsignedPixels,
			isSigned: false,
			str:      "Unsigned",
		},
		{
			name:     "signed",
			pr:       SignedPixels,
			isSigned: true,
			str:      "Signed",
		},
		{
			name:     "invalid",
			pr:       PixelRepresentation(99),
			isSigned: false,
			str:      "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pr.IsSigned() != tt.isSigned {
				t.Errorf("IsSigned() = %v, want %v", tt.pr.IsSigned(), tt.isSigned)
			}
			if tt.pr.String() != tt.str {
				t.Errorf("String() = %q, want %q", tt.pr.String(), tt.str)
			}
		})
	}
}

func TestPlanarConfiguration(t *testing.T) {
	tests := []struct {
		name          string
		pc            PlanarConfiguration
		isInterleaved bool
		isPlanar      bool
		str           string
	}{
		{
			name:          "interleaved",
			pc:            InterleavedPlanar,
			isInterleaved: true,
			isPlanar:      false,
			str:           "Interleaved",
		},
		{
			name:          "planar",
			pc:            PlanarPlanar,
			isInterleaved: false,
			isPlanar:      true,
			str:           "Planar",
		},
		{
			name:          "invalid",
			pc:            PlanarConfiguration(99),
			isInterleaved: false,
			isPlanar:      false,
			str:           "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pc.IsInterleaved() != tt.isInterleaved {
				t.Errorf("IsInterleaved() = %v, want %v", tt.pc.IsInterleaved(), tt.isInterleaved)
			}
			if tt.pc.IsPlanar() != tt.isPlanar {
				t.Errorf("IsPlanar() = %v, want %v", tt.pc.IsPlanar(), tt.isPlanar)
			}
			if tt.pc.String() != tt.str {
				t.Errorf("String() = %q, want %q", tt.pc.String(), tt.str)
			}
		})
	}
}
