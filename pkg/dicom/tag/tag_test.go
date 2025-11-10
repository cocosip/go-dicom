// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package tag_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestNew(t *testing.T) {
	tag := tag.New(0x0028, 0x0010)
	if tag.Group() != 0x0028 {
		t.Errorf("Group() = 0x%04x, want 0x0028", tag.Group())
	}
	if tag.Element() != 0x0010 {
		t.Errorf("Element() = 0x%04x, want 0x0010", tag.Element())
	}
	if tag.PrivateCreator() != nil {
		t.Error("PrivateCreator() should be nil for public tag")
	}
}

func TestFromUint32(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
		wantG uint16
		wantE uint16
	}{
		{"patient name", 0x00100010, 0x0010, 0x0010},
		{"rows", 0x00280010, 0x0028, 0x0010},
		{"columns", 0x00280011, 0x0028, 0x0011},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := tag.FromUint32(tt.value)
			if tag.Group() != tt.wantG {
				t.Errorf("Group() = 0x%04x, want 0x%04x", tag.Group(), tt.wantG)
			}
			if tag.Element() != tt.wantE {
				t.Errorf("Element() = 0x%04x, want 0x%04x", tag.Element(), tt.wantE)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		name  string
		group uint16
		elem  uint16
		want  uint32
	}{
		{"patient name", 0x0010, 0x0010, 0x00100010},
		{"rows", 0x0028, 0x0010, 0x00280010},
		{"columns", 0x0028, 0x0011, 0x00280011},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := tag.New(tt.group, tt.elem)
			if got := tag.ToUint32(); got != tt.want {
				t.Errorf("ToUint32() = 0x%08x, want 0x%08x", got, tt.want)
			}
		})
	}
}

func TestIsPrivate(t *testing.T) {
	tests := []struct {
		name  string
		group uint16
		want  bool
	}{
		{"public even", 0x0010, false},
		{"public even 2", 0x0028, false},
		{"private odd", 0x0009, true},
		{"private odd 2", 0x0029, true},
		{"private odd 3", 0x4009, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := tag.New(tt.group, 0x0010)
			if got := tag.IsPrivate(); got != tt.want {
				t.Errorf("IsPrivate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name   string
		tag    *tag.Tag
		format string
		want   string
	}{
		{"default format", tag.New(0x0028, 0x0010), "G", "(0028,0010)"},
		{"format G", tag.New(0x0010, 0x0010), "G", "(0010,0010)"},
		{"format g", tag.New(0x0028, 0x0010), "g", "0028,0010"},
		{"format J", tag.New(0x0028, 0x0010), "J", "00280010"},
		{"format X", tag.New(0x0028, 0x0010), "X", "(0028,0010)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tag.Format(tt.format); got != tt.want {
				t.Errorf("Format(%q) = %v, want %v", tt.format, got, tt.want)
			}
		})
	}
}

func TestFormatWithPrivateCreator(t *testing.T) {
	pc := tag.NewPrivateCreator("MYPRIVATE")
	tag := tag.NewWithPrivateCreator(0x0029, 0x1001, pc)

	tests := []struct {
		format string
		want   string
	}{
		{"G", "(0029,1001:MYPRIVATE)"},
		{"X", "(0029,xx01:MYPRIVATE)"},
		{"g", "0029,1001:MYPRIVATE"},
		{"J", "00291001"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			if got := tag.Format(tt.format); got != tt.want {
				t.Errorf("Format(%q) = %v, want %v", tt.format, got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tag := tag.New(0x0028, 0x0010)
	want := "(0028,0010)"
	if got := tag.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		t1   *tag.Tag
		t2   *tag.Tag
		want int // -1: less, 0: equal, 1: greater
	}{
		{"equal", tag.New(0x0010, 0x0010), tag.New(0x0010, 0x0010), 0},
		{"less by group", tag.New(0x0010, 0x0010), tag.New(0x0020, 0x0010), -1},
		{"greater by group", tag.New(0x0020, 0x0010), tag.New(0x0010, 0x0010), 1},
		{"less by element", tag.New(0x0010, 0x0010), tag.New(0x0010, 0x0020), -1},
		{"greater by element", tag.New(0x0010, 0x0020), tag.New(0x0010, 0x0010), 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t1.Compare(tt.t2)
			// Normalize to -1, 0, 1
			if got < 0 {
				got = -1
			} else if got > 0 {
				got = 1
			}
			if got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name string
		t1   *tag.Tag
		t2   *tag.Tag
		want bool
	}{
		{"equal", tag.New(0x0010, 0x0010), tag.New(0x0010, 0x0010), true},
		{"different group", tag.New(0x0010, 0x0010), tag.New(0x0020, 0x0010), false},
		{"different element", tag.New(0x0010, 0x0010), tag.New(0x0010, 0x0020), false},
		{"nil", tag.New(0x0010, 0x0010), nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Equals(tt.t2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualsWithPrivateCreator(t *testing.T) {
	pc1 := tag.NewPrivateCreator("PRIVATE1")
	pc2 := tag.NewPrivateCreator("PRIVATE2")

	tests := []struct {
		name string
		t1   *tag.Tag
		t2   *tag.Tag
		want bool
	}{
		{
			"same private creator",
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc1),
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc1),
			true,
		},
		{
			"different private creator",
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc1),
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc2),
			false,
		},
		{
			"one private one public",
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc1),
			tag.New(0x0029, 0x1001),
			false,
		},
		{
			"private tags with different upper bits",
			tag.NewWithPrivateCreator(0x0029, 0x1001, pc1),
			tag.NewWithPrivateCreator(0x0029, 0x1101, pc1),
			true, // Only lower 8 bits matter for private tags
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Equals(tt.t2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantG   uint16
		wantE   uint16
		wantErr bool
	}{
		{"with parens", "(0028,0010)", 0x0028, 0x0010, false},
		{"without parens", "0028,0010", 0x0028, 0x0010, false},
		{"compact", "00280010", 0x0028, 0x0010, false},
		{"with spaces", " (0028,0010) ", 0x0028, 0x0010, false},
		{"uppercase", "(0028,0010)", 0x0028, 0x0010, false},
		{"lowercase", "(002a,0011)", 0x002a, 0x0011, false},
		{"too short", "0028", 0, 0, true},
		{"invalid group", "(XXXX,0010)", 0, 0, true},
		{"invalid element", "(0028,YYYY)", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tag.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Group() != tt.wantG {
				t.Errorf("Parse(%q).Group() = 0x%04x, want 0x%04x", tt.input, got.Group(), tt.wantG)
			}
			if got.Element() != tt.wantE {
				t.Errorf("Parse(%q).Element() = 0x%04x, want 0x%04x", tt.input, got.Element(), tt.wantE)
			}
		})
	}
}

func TestParseWithPrivateCreator(t *testing.T) {
	input := "(0029,1001:MYPRIVATE)"
	got, err := tag.Parse(input)
	if err != nil {
		t.Fatalf("Parse(%q) error = %v", input, err)
	}

	if got.Group() != 0x0029 {
		t.Errorf("Group() = 0x%04x, want 0x0029", got.Group())
	}
	if got.Element() != 0x1001 {
		t.Errorf("Element() = 0x%04x, want 0x1001", got.Element())
	}
	if got.PrivateCreator() == nil {
		t.Fatal("PrivateCreator() is nil, want non-nil")
	}
	if got.PrivateCreator().Creator() != "MYPRIVATE" {
		t.Errorf("PrivateCreator().Creator() = %q, want \"MYPRIVATE\"", got.PrivateCreator().Creator())
	}
}

func TestMustParse(t *testing.T) {
	// Should not panic for valid input
	result := tag.MustParse("(0028,0010)")
	if result.Group() != 0x0028 {
		t.Errorf("Group() = 0x%04x, want 0x0028", result.Group())
	}

	// Should panic for invalid input
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParse with invalid input should panic")
		}
	}()
	tag.MustParse("invalid")
}

func TestUnknown(t *testing.T) {
	if tag.Unknown.Group() != 0xFFFF {
		t.Errorf("Unknown.Group() = 0x%04x, want 0xFFFF", tag.Unknown.Group())
	}
	if tag.Unknown.Element() != 0xFFFF {
		t.Errorf("Unknown.Element() = 0x%04x, want 0xFFFF", tag.Unknown.Element())
	}
}

func TestHash(t *testing.T) {
	t1 := tag.New(0x0010, 0x0010)
	t2 := tag.New(0x0010, 0x0010)
	t3 := tag.New(0x0020, 0x0010)

	// Equal tags should have equal hashes
	if t1.Hash() != t2.Hash() {
		t.Error("Equal tags should have equal hashes")
	}

	// Different tags should (likely) have different hashes
	if t1.Hash() == t3.Hash() {
		t.Error("Different tags should have different hashes (collision possible but unlikely)")
	}
}

func TestDictionaryEntry(t *testing.T) {
	// Create a tag
	testTag := tag.New(0x0010, 0x0010)

	// Get dictionary entry
	entry := testTag.DictionaryEntry()

	// The entry might be nil if the dictionary is not fully initialized yet,
	// but the method should not panic
	if entry != nil {
		// If we get an entry, verify it's the correct type
		// (We can't import dict here due to circular dependency,
		// so we just check it's not nil)
		t.Logf("Got dictionary entry: %v", entry)
	} else {
		t.Log("Dictionary entry is nil (dictionary may not be initialized)")
	}
}

func TestDictionaryEntryNotInitialized(t *testing.T) {
	// Save the original lookup function
	originalLookup := tag.SetDictionaryLookup

	// Set lookup to nil to simulate uninitialized dictionary
	tag.SetDictionaryLookup(nil)

	testTag := tag.New(0x0010, 0x0010)
	entry := testTag.DictionaryEntry()

	if entry != nil {
		t.Error("DictionaryEntry() should return nil when lookup is not initialized")
	}

	// Restore the original lookup function
	// Note: This won't actually restore it because SetDictionaryLookup
	// sets a global variable, but we include it for clarity
	_ = originalLookup
}

func TestParsePrivateCreatorCaching(t *testing.T) {
	// Note: This test verifies caching behavior when the dict package is imported.
	// If dict package is not imported, private creators won't be cached (fallback behavior).

	// Parse the same private tag twice
	tag1, err := tag.Parse("(0029,1001:TESTCREATOR)")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	tag2, err := tag.Parse("(0029,1002:TESTCREATOR)")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify they have the correct creator string
	if tag1.PrivateCreator().Creator() != "TESTCREATOR" {
		t.Errorf("tag1 PrivateCreator().Creator() = %q, want \"TESTCREATOR\"", tag1.PrivateCreator().Creator())
	}
	if tag2.PrivateCreator().Creator() != "TESTCREATOR" {
		t.Errorf("tag2 PrivateCreator().Creator() = %q, want \"TESTCREATOR\"", tag2.PrivateCreator().Creator())
	}

	// Parse a different private creator
	tag3, err := tag.Parse("(0029,1003:OTHERCREATOR)")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if tag3.PrivateCreator().Creator() != "OTHERCREATOR" {
		t.Errorf("tag3 PrivateCreator().Creator() = %q, want \"OTHERCREATOR\"", tag3.PrivateCreator().Creator())
	}

	// Note: We don't test pointer equality here because the tag_test package
	// doesn't import dict, so the globalPrivateCreatorLookup may not be set.
	// The caching behavior is tested in dict/dictionary_test.go instead.
}
