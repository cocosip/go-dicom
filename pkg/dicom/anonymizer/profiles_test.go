// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

const retainUIDsTestName = "retain UIDs"

func TestNewProfileFromReaderAcceptsGeneralReader(t *testing.T) {
	profile, err := NewProfileFromReader(bytes.NewBufferString("0010,0010;Z\n"), BasicProfile)
	if err != nil {
		t.Fatalf("NewProfileFromReader() error = %v", err)
	}

	action, found := profile.FindAction(mustParseTag("(0010,0010)"))
	if !found || action != ActionZ {
		t.Fatalf("Patient Name action = %v, found = %v, want Z, true", action, found)
	}
}

func TestLoadProfileFromFileMatchesReader(t *testing.T) {
	const content = "0010,0010;Z\n0008,0018;K\n"
	path := filepath.Join(t.TempDir(), "profile.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}

	fromFile, err := LoadProfileFromFile(path, BasicProfile)
	if err != nil {
		t.Fatalf("LoadProfileFromFile() error = %v", err)
	}
	fromReader, err := NewProfileFromReader(strings.NewReader(content), BasicProfile)
	if err != nil {
		t.Fatalf("NewProfileFromReader() error = %v", err)
	}

	for _, tagString := range []string{"(0010,0010)", "(0008,0018)"} {
		tagValue := mustParseTag(tagString)
		fileAction, fileFound := fromFile.FindAction(tagValue)
		readerAction, readerFound := fromReader.FindAction(tagValue)
		if fileAction != readerAction || fileFound != readerFound {
			t.Fatalf("rule %s differs: file = (%v, %v), reader = (%v, %v)",
				tagString, fileAction, fileFound, readerAction, readerFound)
		}
	}
}

func TestNewProfileFromReaderAppliesOptionsLikeDefaultProfile(t *testing.T) {
	const content = "0008,0018;U;;K;;;;;;;;\n"
	tests := []struct {
		name    string
		options SecurityProfileOptions
		want    SecurityProfileAction
	}{
		{name: "basic", options: BasicProfile, want: ActionU},
		{name: retainUIDsTestName, options: RetainUIDs, want: ActionK},
		{name: "later selected option wins", options: BasicProfile | RetainUIDs, want: ActionK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromReader, err := NewProfileFromReader(strings.NewReader(content), tt.options)
			if err != nil {
				t.Fatalf("NewProfileFromReader() error = %v", err)
			}
			fromDefault := NewSecurityProfile(tt.options)

			tagValue := mustParseTag("(0008,0018)")
			readerAction, readerFound := fromReader.FindAction(tagValue)
			defaultAction, defaultFound := fromDefault.FindAction(tagValue)
			if !readerFound || readerAction != tt.want {
				t.Fatalf("reader action = %v, found = %v, want %v, true", readerAction, readerFound, tt.want)
			}
			if !defaultFound || defaultAction != tt.want {
				t.Fatalf("default action = %v, found = %v, want %v, true", defaultAction, defaultFound, tt.want)
			}
		})
	}
}

func TestNewProfileFromReaderMapsEveryOptionColumn(t *testing.T) {
	const content = "0010,0010;D;Z;X;K;C;U;D;Z;X;K;C\n"
	tests := []struct {
		name    string
		options SecurityProfileOptions
		want    SecurityProfileAction
	}{
		{name: "basic", options: BasicProfile, want: ActionD},
		{name: "retain safe private", options: RetainSafePrivate, want: ActionZ},
		{name: retainUIDsTestName, options: RetainUIDs, want: ActionX},
		{name: "retain device identifiers", options: RetainDeviceIdent, want: ActionK},
		{name: "retain institution identifiers", options: RetainInstitutionIdent, want: ActionC},
		{name: "retain patient characteristics", options: RetainPatientChars, want: ActionU},
		{name: "retain full dates", options: RetainLongFullDates, want: ActionD},
		{name: "retain modified dates", options: RetainLongModifDates, want: ActionZ},
		{name: "clean descriptions", options: CleanDesc, want: ActionX},
		{name: "clean structured content", options: CleanStructdCont, want: ActionK},
		{name: "clean graphics", options: CleanGraph, want: ActionC},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, err := NewProfileFromReader(strings.NewReader(content), tt.options)
			if err != nil {
				t.Fatalf("NewProfileFromReader() error = %v", err)
			}
			action, found := profile.FindAction(mustParseTag("(0010,0010)"))
			if !found || action != tt.want {
				t.Fatalf("action = %v, found = %v, want %v, true", action, found, tt.want)
			}
		})
	}
}

func TestNewProfileFromReaderUsesStructurePreservingCompositeAction(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    SecurityProfileAction
	}{
		{name: "explicit X/Z", content: "0010,0010;X/Z\n", want: ActionZ},
		{name: "option X/Z/D", content: "0010,0010;X/Z/D;;;;;;;;;;\n", want: ActionD},
		{name: "explicit X/U", content: "0010,0010;X/U\n", want: ActionU},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, err := NewProfileFromReader(strings.NewReader(tt.content), BasicProfile)
			if err != nil {
				t.Fatalf("NewProfileFromReader() error = %v", err)
			}
			action, found := profile.FindAction(mustParseTag("(0010,0010)"))
			if !found || action != tt.want {
				t.Fatalf("action = %v, found = %v, want %v, true", action, found, tt.want)
			}
		})
	}
}

func TestNewSecurityProfileRetainUIDsKeepsBasicPatientNameAction(t *testing.T) {
	action, found := NewSecurityProfile(BasicProfile | RetainUIDs).FindAction(tag.PatientName)
	if !found || action != ActionZ {
		t.Fatalf("Patient Name action = %v, found = %v, want Z, true", action, found)
	}
}

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

	for _, retainedByBusinessPolicy := range []*tag.Tag{
		tag.StudyDate,
		tag.SeriesDate,
		tag.StudyTime,
		tag.AccessionNumber,
		tag.PatientID,
		tag.StudyID,
		tag.ModifyingDeviceIDRETIRED,
	} {
		action, found = profile.FindAction(retainedByBusinessPolicy)
		if found && action == ActionK {
			t.Errorf("%s action = K, want the standard profile action", retainedByBusinessPolicy)
		}
	}
}

func TestPredefinedRetainProfilesIncludeBasicProfile(t *testing.T) {
	tests := []struct {
		name       string
		newProfile func() *SecurityProfile
	}{
		{name: retainUIDsTestName, newProfile: NewRetainUIDsProfile},
		{name: "retain institution", newProfile: NewRetainInstitutionProfile},
		{name: "retain dates", newProfile: NewRetainDatesProfile},
		{name: "research", newProfile: NewResearchProfile},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			action, found := tt.newProfile().FindAction(mustParseTag("(0010,1005)"))
			if !found || action != ActionX {
				t.Fatalf("Patient Birth Name action = %v, found = %v, want X, true", action, found)
			}
		})
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

func TestNewProfileFromReaderRejectsInvalidLines(t *testing.T) {
	const (
		actionDetail  = "action"
		patternDetail = "pattern"
	)
	tests := []struct {
		name       string
		content    string
		options    SecurityProfileOptions
		wantDetail string
	}{
		{name: "wrong column count", content: "# comment\n0010,0010;Z;extra\n", wantDetail: "column"},
		{name: "unknown direct action", content: "# comment\n0010,0010;Q\n", wantDetail: actionDetail},
		{name: "unknown table action", content: "# comment\n0010,0010;Q;;;;;;;;;;\n", wantDetail: actionDetail},
		{name: "unknown composite direct action", content: "# comment\n0010,0010;X/Q\n", wantDetail: actionDetail},
		{name: "unknown composite table action", content: "# comment\n0010,0010;X/Q;;;;;;;;;;\n", wantDetail: actionDetail},
		{name: "empty pattern", content: "# comment\n;Z\n", wantDetail: patternDetail},
		{name: "invalid pattern", content: "# comment\n[;Z\n", wantDetail: patternDetail},
		{
			name:       "invalid pattern with inactive option",
			content:    "# comment\n[;Z;;;;;;;;;;\n",
			options:    RetainUIDs,
			wantDetail: patternDetail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := tt.options
			if options == 0 {
				options = BasicProfile
			}
			_, err := NewProfileFromReader(strings.NewReader(tt.content), options)
			if err == nil {
				t.Fatal("NewProfileFromReader() error = nil, want validation error")
			}
			if !strings.Contains(err.Error(), "line 2") || !strings.Contains(err.Error(), tt.wantDetail) {
				t.Fatalf("NewProfileFromReader() error = %q, want line number and %q", err, tt.wantDetail)
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
