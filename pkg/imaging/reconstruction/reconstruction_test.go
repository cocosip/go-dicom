// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package reconstruction

import "testing"

func TestStackType_String(t *testing.T) {
	testCases := []struct {
		stackType StackType
		expected  string
	}{
		{StackTypeAxial, "Axial"},
		{StackTypeCoronal, "Coronal"},
		{StackTypeSagittal, "Sagittal"},
		{StackType(99), "Unknown"},
	}

	for _, tc := range testCases {
		result := tc.stackType.String()
		if result != tc.expected {
			t.Errorf("StackType(%d).String(): expected %s, got %s", tc.stackType, tc.expected, result)
		}
	}
}
