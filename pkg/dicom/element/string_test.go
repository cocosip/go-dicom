// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const testPatientName = "Doe^John"

// TestStringElement_Single tests single-valued string elements
func TestStringElement_Single(t *testing.T) {
	elem := element.NewString(tag.PatientName, vr.PN, []string{testPatientName})

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != 1 {
			t.Errorf("Count() = %d, want 1", elem.Count())
		}
	})

	t.Run("GetString", func(t *testing.T) {
		got := elem.GetString()
		want := testPatientName
		if got != want {
			t.Errorf("GetString() = %q, want %q", got, want)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		got := elem.GetValue(0)
		want := testPatientName
		if got != want {
			t.Errorf("GetValue(0) = %q, want %q", got, want)
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		values := elem.GetValues()
		if len(values) != 1 {
			t.Errorf("len(GetValues()) = %d, want 1", len(values))
		}
		if values[0] != testPatientName {
			t.Errorf("GetValues()[0] = %q, want %q", values[0], testPatientName)
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

// TestStringElement_VRValidation tests VR-specific validation
func TestStringElement_VRValidation(t *testing.T) {
	tests := []struct {
		name      string
		tag       *tag.Tag
		vr        *vr.VR
		values    []string
		wantError bool
		desc      string
	}{
		{
			name:      "CS valid uppercase",
			tag:       tag.Modality,
			vr:        vr.CS,
			values:    []string{"CT"},
			wantError: false,
			desc:      "CS should accept valid uppercase code",
		},
		{
			name:      "CS invalid lowercase",
			tag:       tag.Modality,
			vr:        vr.CS,
			values:    []string{"ct"},
			wantError: true,
			desc:      "CS should reject lowercase",
		},
		{
			name:      "UI valid UID",
			tag:       tag.SOPClassUID,
			vr:        vr.UI,
			values:    []string{"1.2.840.10008.5.1.4.1.1.2"},
			wantError: false,
			desc:      "UI should accept valid UID format",
		},
		{
			name:      "UI invalid UID",
			tag:       tag.SOPClassUID,
			vr:        vr.UI,
			values:    []string{"invalid.uid.format"},
			wantError: true,
			desc:      "UI should reject invalid UID format",
		},
		{
			name:      "AE valid title",
			tag:       tag.New(0x0002, 0x0016),
			vr:        vr.AE,
			values:    []string{"STORESCU"},
			wantError: false,
			desc:      "AE should accept valid AE title",
		},
		{
			name:      "AE too long",
			tag:       tag.New(0x0002, 0x0016),
			vr:        vr.AE,
			values:    []string{"THIS_IS_TOO_LONG_FOR_AE"},
			wantError: true,
			desc:      "AE should reject title exceeding 16 chars",
		},
		{
			name:      "DA valid date",
			tag:       tag.StudyDate,
			vr:        vr.DA,
			values:    []string{"20230515"},
			wantError: false,
			desc:      "DA should accept valid date format",
		},
		{
			name:      "DA invalid date",
			tag:       tag.StudyDate,
			vr:        vr.DA,
			values:    []string{"2023-05-15"},
			wantError: true,
			desc:      "DA should reject invalid date format",
		},
		{
			name:      "TM valid time",
			tag:       tag.StudyTime,
			vr:        vr.TM,
			values:    []string{"143052"},
			wantError: false,
			desc:      "TM should accept valid time format",
		},
		{
			name:      "TM invalid time",
			tag:       tag.StudyTime,
			vr:        vr.TM,
			values:    []string{"14:30:52"},
			wantError: true,
			desc:      "TM should reject invalid time format",
		},
		{
			name:      "AS valid age",
			tag:       tag.PatientAge,
			vr:        vr.AS,
			values:    []string{"045Y"},
			wantError: false,
			desc:      "AS should accept valid age string",
		},
		{
			name:      "AS invalid age",
			tag:       tag.PatientAge,
			vr:        vr.AS,
			values:    []string{"45"},
			wantError: true,
			desc:      "AS should reject invalid age format",
		},
		{
			name:      "PN valid name",
			tag:       tag.PatientName,
			vr:        vr.PN,
			values:    []string{"Doe^John"},
			wantError: false,
			desc:      "PN should accept valid person name",
		},
		{
			name:      "Multiple values mixed validity",
			tag:       tag.ImageType,
			vr:        vr.CS,
			values:    []string{"ORIGINAL", "primary"},
			wantError: true,
			desc:      "Should fail if any value is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elem := element.NewString(tt.tag, tt.vr, tt.values)
			err := elem.Validate()

			if tt.wantError && err == nil {
				t.Errorf("Validate() error = nil, want error: %s", tt.desc)
			}
			if !tt.wantError && err != nil {
				t.Errorf("Validate() error = %v, want nil: %s", err, tt.desc)
			}
		})
	}
}

// TestStringElement_ValidationDisabled tests validation can be disabled
func TestStringElement_ValidationDisabled(t *testing.T) {
	// Save current state
	oldValidation := vr.PerformValidation
	defer func() {
		vr.PerformValidation = oldValidation
	}()

	// Disable validation
	vr.PerformValidation = false

	// Create element with invalid value
	elem := element.NewString(tag.Modality, vr.CS, []string{"invalid_lowercase"})

	// Should not error when validation is disabled
	err := elem.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v, want nil when validation disabled", err)
	}
}
