// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestNewRetainUIDsProfile(t *testing.T) {
	profile := NewRetainUIDsProfile()

	if profile == nil {
		t.Fatal("NewRetainUIDsProfile() returned nil")
	}

	// Check that SOP Instance UID is kept
	action, found := profile.FindAction(mustParseTag("(0008,0018)"))
	if !found {
		t.Error("SOP Instance UID rule not found")
	}
	if action != ActionK {
		t.Errorf("SOP Instance UID action = %v, want ActionK", action)
	}

	// Check that Study Instance UID is kept
	action, found = profile.FindAction(mustParseTag("(0020,000D)"))
	if !found {
		t.Error("Study Instance UID rule not found")
	}
	if action != ActionK {
		t.Errorf("Study Instance UID action = %v, want ActionK", action)
	}
}

func TestNewRetainInstitutionProfile(t *testing.T) {
	profile := NewRetainInstitutionProfile()

	if profile == nil {
		t.Fatal("NewRetainInstitutionProfile() returned nil")
	}

	// Check that Institution Name is kept
	action, found := profile.FindAction(mustParseTag("(0008,0080)"))
	if !found {
		t.Error("Institution Name rule not found")
	}
	if action != ActionK {
		t.Errorf("Institution Name action = %v, want ActionK", action)
	}

	// Check that Referring Physician's Name is kept
	action, found = profile.FindAction(mustParseTag("(0008,0090)"))
	if !found {
		t.Error("Referring Physician's Name rule not found")
	}
	if action != ActionK {
		t.Errorf("Referring Physician's Name action = %v, want ActionK", action)
	}
}

func TestNewRetainDatesProfile(t *testing.T) {
	profile := NewRetainDatesProfile()

	if profile == nil {
		t.Fatal("NewRetainDatesProfile() returned nil")
	}

	// Check that Study Date is kept
	action, found := profile.FindAction(mustParseTag("(0008,0020)"))
	if !found {
		t.Error("Study Date rule not found")
	}
	if action != ActionK {
		t.Errorf("Study Date action = %v, want ActionK", action)
	}

	// Check that Patient's Birth Date is kept
	action, found = profile.FindAction(mustParseTag("(0010,0030)"))
	if !found {
		t.Error("Patient's Birth Date rule not found")
	}
	if action != ActionK {
		t.Errorf("Patient's Birth Date action = %v, want ActionK", action)
	}
}

func TestNewMinimalProfile(t *testing.T) {
	profile := NewMinimalProfile()

	if profile == nil {
		t.Fatal("NewMinimalProfile() returned nil")
	}

	// Check that Patient's Name is zeroed
	action, found := profile.FindAction(mustParseTag("(0010,0010)"))
	if !found {
		t.Error("Patient's Name rule not found")
	}
	if action != ActionZ {
		t.Errorf("Patient's Name action = %v, want ActionZ", action)
	}

	// Check that Patient ID is zeroed
	action, found = profile.FindAction(mustParseTag("(0010,0020)"))
	if !found {
		t.Error("Patient ID rule not found")
	}
	if action != ActionZ {
		t.Errorf("Patient ID action = %v, want ActionZ", action)
	}
}

func TestNewResearchProfile(t *testing.T) {
	profile := NewResearchProfile()

	if profile == nil {
		t.Fatal("NewResearchProfile() returned nil")
	}

	// Check that Patient's Name is zeroed
	action, found := profile.FindAction(mustParseTag("(0010,0010)"))
	if !found {
		t.Error("Patient's Name rule not found")
	}
	if action != ActionZ {
		t.Errorf("Patient's Name action = %v, want ActionZ", action)
	}

	// Check that Patient's Age is kept
	action, found = profile.FindAction(mustParseTag("(0010,1010)"))
	if !found {
		t.Error("Patient's Age rule not found")
	}
	if action != ActionK {
		t.Errorf("Patient's Age action = %v, want ActionK", action)
	}

	// Check that SOP Instance UID is kept
	action, found = profile.FindAction(mustParseTag("(0008,0018)"))
	if !found {
		t.Error("SOP Instance UID rule not found")
	}
	if action != ActionK {
		t.Errorf("SOP Instance UID action = %v, want ActionK", action)
	}
}

func TestNewProfileFromReader(t *testing.T) {
	content := `# Custom profile
0010,0010;Z
0010,0020;Z
0008,0018;K
0020,000D;K
# This is a comment
0010,0030;X

invalid line without semicolon
0010,0040;C
`
	reader := strings.NewReader(content)
	profile, err := NewProfileFromReader(reader, BasicProfile)

	if err != nil {
		t.Errorf("NewProfileFromReader() error = %v", err)
	}

	if profile == nil {
		t.Fatal("NewProfileFromReader() returned nil profile")
	}

	// Test rules were loaded
	tests := []struct {
		tag      string
		expected SecurityProfileAction
	}{
		{"(0010,0010)", ActionZ}, // Patient's Name
		{"(0010,0020)", ActionZ}, // Patient ID
		{"(0008,0018)", ActionK}, // SOP Instance UID
		{"(0020,000D)", ActionK}, // Study Instance UID
		{"(0010,0030)", ActionX}, // Patient's Birth Date
		{"(0010,0040)", ActionC}, // Patient's Sex
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			action, found := profile.FindAction(mustParseTag(tt.tag))
			if !found {
				t.Errorf("Rule for %s not found", tt.tag)
				return
			}
			if action != tt.expected {
				t.Errorf("Action for %s = %v, want %v", tt.tag, action, tt.expected)
			}
		})
	}
}

func TestNewProfileFromReaderEmpty(t *testing.T) {
	reader := strings.NewReader("")
	profile, err := NewProfileFromReader(reader, BasicProfile)

	if err != nil {
		t.Errorf("NewProfileFromReader() error = %v", err)
	}

	if profile == nil {
		t.Fatal("NewProfileFromReader() returned nil profile")
	}
}

func TestNewProfileFromReaderComments(t *testing.T) {
	content := `# Only comments
# Another comment
`
	reader := strings.NewReader(content)
	profile, err := NewProfileFromReader(reader, BasicProfile)

	if err != nil {
		t.Errorf("NewProfileFromReader() error = %v", err)
	}

	if profile == nil {
		t.Fatal("NewProfileFromReader() returned nil profile")
	}
}

func mustParseTag(s string) *tag.Tag {
	t, err := tag.Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}
