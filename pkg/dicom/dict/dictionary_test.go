// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewEntry(t *testing.T) {
	entry := dict.NewEntry(
		tag.New(0x0010, 0x0010),
		"Patient Name",
		"PatientName",
		vm.VM1,
		false,
		vr.PN,
	)

	if entry.Tag().Group() != 0x0010 {
		t.Errorf("Tag().Group() = 0x%04x, want 0x0010", entry.Tag().Group())
	}
	if entry.Name() != "Patient Name" {
		t.Errorf("Name() = %q, want %q", entry.Name(), "Patient Name")
	}
	if entry.Keyword() != "PatientName" {
		t.Errorf("Keyword() = %q, want %q", entry.Keyword(), "PatientName")
	}
	if entry.IsRetired() {
		t.Error("IsRetired() = true, want false")
	}
	if entry.IsMasked() {
		t.Error("IsMasked() = true, want false")
	}
}

func TestNewEntryWithMask(t *testing.T) {
	mt := tag.MustParseMaskedTag("(xxxx,0000)")
	entry := dict.NewEntryWithMask(
		mt,
		"Group Length",
		"GroupLength",
		vm.VM1,
		false,
		vr.UL,
	)

	if entry.Name() != "Group Length" {
		t.Errorf("Name() = %q, want %q", entry.Name(), "Group Length")
	}
	if !entry.IsMasked() {
		t.Error("IsMasked() = false, want true")
	}
	if entry.MaskTag() == nil {
		t.Error("MaskTag() = nil, want non-nil")
	}
}

func TestEntryMatches(t *testing.T) {
	tests := []struct {
		name      string
		entry     *dict.Entry
		testTag   *tag.Tag
		wantMatch bool
	}{
		{
			"exact match",
			dict.NewEntry(tag.New(0x0010, 0x0010), "Patient's Name", "PatientName", vm.VM1, false, vr.PN),
			tag.New(0x0010, 0x0010),
			true,
		},
		{
			"no match",
			dict.NewEntry(tag.New(0x0010, 0x0010), "Patient's Name", "PatientName", vm.VM1, false, vr.PN),
			tag.New(0x0010, 0x0020),
			false,
		},
		{
			"masked match - group length",
			dict.NewEntryWithMask(tag.MustParseMaskedTag("(xxxx,0000)"), "Group Length", "GroupLength", vm.VM1, false, vr.UL),
			tag.New(0x0010, 0x0000),
			true,
		},
		{
			"masked match - group length 2",
			dict.NewEntryWithMask(tag.MustParseMaskedTag("(xxxx,0000)"), "Group Length", "GroupLength", vm.VM1, false, vr.UL),
			tag.New(0x0028, 0x0000),
			true,
		},
		{
			"masked no match",
			dict.NewEntryWithMask(tag.MustParseMaskedTag("(xxxx,0000)"), "Group Length", "GroupLength", vm.VM1, false, vr.UL),
			tag.New(0x0010, 0x0010),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.entry.Matches(tt.testTag); got != tt.wantMatch {
				t.Errorf("Matches() = %v, want %v", got, tt.wantMatch)
			}
		})
	}
}

func TestDictionaryAdd(t *testing.T) {
	d := dict.New()

	entry := dict.NewEntry(
		tag.New(0x0010, 0x0010),
		"Patient Name",
		"PatientName",
		vm.VM1,
		false,
		vr.PN,
	)

	d.Add(entry)

	// Lookup by tag
	found := d.Lookup(tag.New(0x0010, 0x0010))
	if found == nil {
		t.Fatal("Lookup() returned nil, want entry")
	}
	if found.Name() != "Patient Name" {
		t.Errorf("Lookup().Name() = %q, want %q", found.Name(), "Patient Name")
	}

	// Lookup by keyword
	foundTag := d.LookupKeyword("PatientName")
	if foundTag == nil {
		t.Fatal("LookupKeyword() returned nil, want tag")
	}
	if !foundTag.Equals(tag.New(0x0010, 0x0010)) {
		t.Errorf("LookupKeyword() = %v, want (0010,0010)", foundTag)
	}
}

func TestDictionaryMaskedLookup(t *testing.T) {
	d := dict.New()

	// Add group length masked entry
	groupLengthEntry := dict.NewEntryWithMask(
		tag.MustParseMaskedTag("(xxxx,0000)"),
		"Group Length",
		"GroupLength",
		vm.VM1,
		false,
		vr.UL,
	)
	d.Add(groupLengthEntry)

	// Lookup various group length tags
	tests := []struct {
		name      string
		lookupTag *tag.Tag
		wantFound bool
	}{
		{"group 0x0000", tag.New(0x0000, 0x0000), true},
		{"group 0x0010", tag.New(0x0010, 0x0000), true},
		{"group 0x0028", tag.New(0x0028, 0x0000), true},
		{"not group length", tag.New(0x0010, 0x0010), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := d.Lookup(tt.lookupTag)
			if tt.wantFound && found == nil {
				t.Error("Lookup() returned nil, want entry")
			}
			if !tt.wantFound && found != nil {
				t.Errorf("Lookup() returned %v, want nil", found.Name())
			}
			if tt.wantFound && found != nil && found.Name() != "Group Length" {
				t.Errorf("Lookup().Name() = %q, want %q", found.Name(), "Group Length")
			}
		})
	}
}

func TestDictionaryEntries(t *testing.T) {
	d := dict.New()

	// Add some entries
	d.Add(dict.NewEntry(tag.New(0x0010, 0x0010), "Patient's Name", "PatientName", vm.VM1, false, vr.PN))
	d.Add(dict.NewEntry(tag.New(0x0010, 0x0020), "Patient ID", "PatientID", vm.VM1, false, vr.LO))
	d.Add(dict.NewEntryWithMask(tag.MustParseMaskedTag("(xxxx,0000)"), "Group Length", "GroupLength", vm.VM1, false, vr.UL))

	entries := d.Entries()
	if len(entries) != 3 {
		t.Errorf("Entries() returned %d entries, want 3", len(entries))
	}
}

func TestNewPrivateDictionary(t *testing.T) {
	pc := tag.NewPrivateCreator("MYCOMPANY")
	d := dict.NewPrivate(pc)

	if d.PrivateCreator() == nil {
		t.Error("PrivateCreator() = nil, want non-nil")
	}
	if d.PrivateCreator().Creator() != "MYCOMPANY" {
		t.Errorf("PrivateCreator().Creator() = %q, want %q", d.PrivateCreator().Creator(), "MYCOMPANY")
	}
}

func TestDefaultDictionary(t *testing.T) {
	// Get the default dictionary
	d := dict.Default()
	if d == nil {
		t.Fatal("Default() returned nil")
	}

	// Should be the same instance on repeated calls
	d2 := dict.Default()
	if d != d2 {
		t.Error("Default() should return the same instance")
	}
}

func TestTagDictionaryEntryIntegration(t *testing.T) {
	// The default dictionary should already have PatientName entry loaded from generated data
	testTag := tag.New(0x0010, 0x0010)

	// Now use Tag.DictionaryEntry() method to look it up
	dictEntry := testTag.DictionaryEntry()

	if dictEntry == nil {
		t.Fatal("Tag.DictionaryEntry() returned nil")
	}

	// Type assert to *dict.Entry
	e, ok := dictEntry.(*dict.Entry)
	if !ok {
		t.Fatalf("Tag.DictionaryEntry() returned %T, want *dict.Entry", dictEntry)
	}

	if e.Name() != "Patient's Name" {
		t.Errorf("Entry.Name() = %q, want %q", e.Name(), "Patient's Name")
	}
	if e.Keyword() != "PatientName" {
		t.Errorf("Entry.Keyword() = %q, want %q", e.Keyword(), "PatientName")
	}
}

func TestTagDictionaryEntryNotFound(t *testing.T) {
	// Create a tag that's not in the dictionary
	unknownTag := tag.New(0xEEEE, 0xEEEE)

	dictEntry := unknownTag.DictionaryEntry()

	// Should return nil for tags not in the dictionary
	// Note: dictEntry might be a typed nil (*dict.Entry)(nil), so we need to check carefully
	if dictEntry != nil {
		// Try to type assert - if successful but nil, that's OK
		if e, ok := dictEntry.(*dict.Entry); ok && e != nil {
			t.Errorf("Tag.DictionaryEntry() for unknown tag returned %v, want nil", dictEntry)
		}
	}
}

func TestGetPrivateCreator(t *testing.T) {
	d := dict.New()

	// Get a private creator for the first time
	pc1 := d.GetPrivateCreator("MYCOMPANY")
	if pc1 == nil {
		t.Fatal("GetPrivateCreator() returned nil")
	}
	if pc1.Creator() != "MYCOMPANY" {
		t.Errorf("GetPrivateCreator().Creator() = %q, want %q", pc1.Creator(), "MYCOMPANY")
	}

	// Get the same private creator again - should return the same instance
	pc2 := d.GetPrivateCreator("MYCOMPANY")
	if pc1 != pc2 {
		t.Error("GetPrivateCreator() should return the same instance for the same creator string")
	}

	// Get a different private creator - should return a different instance
	pc3 := d.GetPrivateCreator("OTHERCOMPANY")
	if pc1 == pc3 {
		t.Error("GetPrivateCreator() should return different instances for different creator strings")
	}
	if pc3.Creator() != "OTHERCOMPANY" {
		t.Errorf("GetPrivateCreator().Creator() = %q, want %q", pc3.Creator(), "OTHERCOMPANY")
	}
}

func TestGetPrivateCreatorThreadSafety(t *testing.T) {
	d := dict.New()

	// Test concurrent access to GetPrivateCreator
	const numGoroutines = 100
	const creatorName = "CONCURRENT"

	done := make(chan *tag.PrivateCreator, numGoroutines)

	// Launch multiple goroutines that all try to get the same private creator
	for i := 0; i < numGoroutines; i++ {
		go func() {
			pc := d.GetPrivateCreator(creatorName)
			done <- pc
		}()
	}

	// Collect all results
	var creators []*tag.PrivateCreator
	for i := 0; i < numGoroutines; i++ {
		creators = append(creators, <-done)
	}

	// All should be the same instance (pointer equality)
	first := creators[0]
	for i, pc := range creators {
		if pc != first {
			t.Errorf("Creator at index %d is different from first creator (not the same instance)", i)
		}
		if pc.Creator() != creatorName {
			t.Errorf("Creator at index %d has name %q, want %q", i, pc.Creator(), creatorName)
		}
	}
}

func TestDefaultDictionaryGetPrivateCreator(t *testing.T) {
	// Get private creator from default dictionary
	pc1 := dict.Default().GetPrivateCreator("DEFAULTTEST")
	if pc1 == nil {
		t.Fatal("Default().GetPrivateCreator() returned nil")
	}

	// Should get the same instance again
	pc2 := dict.Default().GetPrivateCreator("DEFAULTTEST")
	if pc1 != pc2 {
		t.Error("Default().GetPrivateCreator() should return the same instance for the same creator")
	}
}

func TestDefaultDictionaryInitialization(t *testing.T) {
	// Get the default dictionary
	d := dict.Default()
	if d == nil {
		t.Fatal("Default() returned nil")
	}

	// Verify that standard DICOM tags are loaded
	testCases := []struct {
		name     string
		tag      *tag.Tag
		wantKW   string
		wantName string
	}{
		{
			"PatientName",
			tag.New(0x0010, 0x0010),
			"PatientName",
			"Patient's Name",
		},
		{
			"StudyInstanceUID",
			tag.New(0x0020, 0x000D),
			"StudyInstanceUID",
			"Study Instance UID",
		},
		{
			"Rows",
			tag.New(0x0028, 0x0010),
			"Rows",
			"Rows",
		},
		{
			"Columns",
			tag.New(0x0028, 0x0011),
			"Columns",
			"Columns",
		},
		{
			"PixelData",
			tag.New(0x7FE0, 0x0010),
			"PixelData",
			"Pixel Data",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entry := d.Lookup(tc.tag)
			if entry == nil {
				t.Fatalf("Lookup(%s) returned nil, expected entry to be found", tc.tag)
			}

			if entry.Keyword() != tc.wantKW {
				t.Errorf("Keyword() = %q, want %q", entry.Keyword(), tc.wantKW)
			}

			if entry.Name() != tc.wantName {
				t.Errorf("Name() = %q, want %q", entry.Name(), tc.wantName)
			}
		})
	}
}

func TestDefaultDictionaryKeywordLookup(t *testing.T) {
	d := dict.Default()

	testCases := []struct {
		keyword     string
		wantGroup   uint16
		wantElement uint16
	}{
		{"PatientName", 0x0010, 0x0010},
		{"StudyInstanceUID", 0x0020, 0x000D},
		{"Rows", 0x0028, 0x0010},
		{"Columns", 0x0028, 0x0011},
		{"PixelData", 0x7FE0, 0x0010},
	}

	for _, tc := range testCases {
		t.Run(tc.keyword, func(t *testing.T) {
			tag := d.LookupKeyword(tc.keyword)
			if tag == nil {
				t.Fatalf("LookupKeyword(%q) returned nil", tc.keyword)
			}

			if tag.Group() != tc.wantGroup {
				t.Errorf("Group() = 0x%04X, want 0x%04X", tag.Group(), tc.wantGroup)
			}

			if tag.Element() != tc.wantElement {
				t.Errorf("Element() = 0x%04X, want 0x%04X", tag.Element(), tc.wantElement)
			}
		})
	}
}

func TestDefaultDictionarySize(t *testing.T) {
	d := dict.Default()
	entries := d.Entries()

	// The DICOM dictionary should have thousands of entries
	if len(entries) < 1000 {
		t.Errorf("Entries() returned %d entries, expected at least 1000", len(entries))
	}

	t.Logf("Dictionary loaded with %d entries", len(entries))
}
