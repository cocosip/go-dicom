// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package tag_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// Test a few standard tags to ensure they were generated correctly
func TestGeneratedTags(t *testing.T) {
	tests := []struct {
		name    string
		tag     *tag.Tag
		wantG   uint16
		wantE   uint16
		wantStr string
	}{
		{
			"CommandGroupLength",
			tag.CommandGroupLength,
			0x0000,
			0x0000,
			"(0000,0000)",
		},
		{
			"AffectedSOPClassUID",
			tag.AffectedSOPClassUID,
			0x0000,
			0x0002,
			"(0000,0002)",
		},
		{
			"PatientName",
			tag.PatientName,
			0x0010,
			0x0010,
			"(0010,0010)",
		},
		{
			"Rows",
			tag.Rows,
			0x0028,
			0x0010,
			"(0028,0010)",
		},
		{
			"Columns",
			tag.Columns,
			0x0028,
			0x0011,
			"(0028,0011)",
		},
		{
			"PixelData",
			tag.PixelData,
			0x7FE0,
			0x0010,
			"(7fe0,0010)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tag.Group() != tt.wantG {
				t.Errorf("Group() = 0x%04x, want 0x%04x", tt.tag.Group(), tt.wantG)
			}
			if tt.tag.Element() != tt.wantE {
				t.Errorf("Element() = 0x%04x, want 0x%04x", tt.tag.Element(), tt.wantE)
			}
			if got := tt.tag.String(); got != tt.wantStr {
				t.Errorf("String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

// Test that Unknown tag is defined
func TestUnknownTag(t *testing.T) {
	if tag.Unknown == nil {
		t.Fatal("Unknown tag should not be nil")
	}
	if tag.Unknown.Group() != 0xFFFF {
		t.Errorf("Unknown.Group() = 0x%04x, want 0xFFFF", tag.Unknown.Group())
	}
	if tag.Unknown.Element() != 0xFFFF {
		t.Errorf("Unknown.Element() = 0x%04x, want 0xFFFF", tag.Unknown.Element())
	}
}
