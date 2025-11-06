// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestStringElement_Single tests single-valued string elements
func TestStringElement_Single(t *testing.T) {
	elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != 1 {
			t.Errorf("Count() = %d, want 1", elem.Count())
		}
	})

	t.Run("GetString", func(t *testing.T) {
		got := elem.GetString()
		want := "Doe^John"
		if got != want {
			t.Errorf("GetString() = %q, want %q", got, want)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		got := elem.GetValue(0)
		want := "Doe^John"
		if got != want {
			t.Errorf("GetValue(0) = %q, want %q", got, want)
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		values := elem.GetValues()
		if len(values) != 1 {
			t.Errorf("len(GetValues()) = %d, want 1", len(values))
		}
		if values[0] != "Doe^John" {
			t.Errorf("GetValues()[0] = %q, want %q", values[0], "Doe^John")
		}
	})
}

// TestStringElement_Multiple tests multi-valued string elements
func TestStringElement_Multiple(t *testing.T) {
	elem := element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"})

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != 3 {
			t.Errorf("Count() = %d, want 3", elem.Count())
		}
	})

	t.Run("GetString", func(t *testing.T) {
		got := elem.GetString()
		want := "ORIGINAL\\PRIMARY\\AXIAL"
		if got != want {
			t.Errorf("GetString() = %q, want %q", got, want)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		tests := []struct {
			index int
			want  string
		}{
			{0, "ORIGINAL"},
			{1, "PRIMARY"},
			{2, "AXIAL"},
		}
		for _, tt := range tests {
			got := elem.GetValue(tt.index)
			if got != tt.want {
				t.Errorf("GetValue(%d) = %q, want %q", tt.index, got, tt.want)
			}
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		values := elem.GetValues()
		want := []string{"ORIGINAL", "PRIMARY", "AXIAL"}
		if len(values) != len(want) {
			t.Errorf("len(GetValues()) = %d, want %d", len(values), len(want))
		}
		for i, v := range values {
			if v != want[i] {
				t.Errorf("GetValues()[%d] = %q, want %q", i, v, want[i])
			}
		}
	})
}

// TestStringElement_Empty tests empty string elements
func TestStringElement_Empty(t *testing.T) {
	elem := element.NewString(tag.PatientComments, vr.LT, []string{})

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != 0 {
			t.Errorf("Count() = %d, want 0", elem.Count())
		}
	})

	t.Run("GetString", func(t *testing.T) {
		if elem.GetString() != "" {
			t.Errorf("GetString() = %q, want empty string", elem.GetString())
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		values := elem.GetValues()
		if values != nil {
			t.Errorf("GetValues() = %v, want nil", values)
		}
	})
}

// TestStringElement_WithSpaces tests string with trailing spaces
func TestStringElement_WithSpaces(t *testing.T) {
	// DICOM removes trailing spaces
	elem := element.NewString(tag.StudyDescription, vr.LO, []string{"Brain Study   "})

	got := elem.GetString()
	want := "Brain Study"
	if got != want {
		t.Errorf("GetString() = %q, want %q (trailing spaces should be removed)", got, want)
	}
}

// TestStringElement_Validation tests string validation
func TestStringElement_Validation(t *testing.T) {
	// Create a string that's too long for CS (Code String, max 16 chars)
	longValue := "THIS_IS_A_VERY_LONG_CODE_STRING"
	elem := element.NewString(tag.Modality, vr.CS, []string{longValue})

	err := elem.Validate()
	if err == nil {
		t.Error("Validate() should return error for string exceeding max length")
	}
}
