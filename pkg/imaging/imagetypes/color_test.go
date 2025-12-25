// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imagetypes

import (
	"testing"
)

func TestNewColor32(t *testing.T) {
	color := NewColor32(255, 100, 150, 200)

	if color.A != 255 {
		t.Errorf("Expected A=255, got %d", color.A)
	}
	if color.R != 100 {
		t.Errorf("Expected R=100, got %d", color.R)
	}
	if color.G != 150 {
		t.Errorf("Expected G=150, got %d", color.G)
	}
	if color.B != 200 {
		t.Errorf("Expected B=200, got %d", color.B)
	}
}

func TestColor32_ToInt32(t *testing.T) {
	testCases := []struct {
		name     string
		color    Color32
		expected int32
	}{
		{
			name:     "white",
			color:    NewColor32(255, 255, 255, 255),
			expected: -1, // 0xFFFFFFFF as signed int32
		},
		{
			name:     "black",
			color:    NewColor32(255, 0, 0, 0),
			expected: -16777216, // 0xFF000000 as signed int32
		},
		{
			name:     "red",
			color:    NewColor32(255, 255, 0, 0),
			expected: -65536, // 0xFFFF0000 as signed int32
		},
		{
			name:     "green",
			color:    NewColor32(255, 0, 255, 0),
			expected: -16711936, // 0xFF00FF00 as signed int32
		},
		{
			name:     "blue",
			color:    NewColor32(255, 0, 0, 255),
			expected: -16776961, // 0xFF0000FF as signed int32
		},
		{
			name:     "custom color",
			color:    NewColor32(255, 100, 150, 200),
			expected: -10185016, // 0xFF6496C8 as signed int32
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.color.ToInt32()
			if result != tc.expected {
				t.Errorf("ToInt32() = %d (0x%08X), expected %d (0x%08X)",
					result, uint32(result), tc.expected, uint32(tc.expected))
			}
		})
	}
}
