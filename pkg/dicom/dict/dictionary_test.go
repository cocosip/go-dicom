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
			dict.NewEntry(tag.New(0x0010, 0x0010), "Patient Name", "PatientName", vm.VM1, false, vr.PN),
			tag.New(0x0010, 0x0010),
			true,
		},
		{
			"no match",
			dict.NewEntry(tag.New(0x0010, 0x0010), "Patient Name", "PatientName", vm.VM1, false, vr.PN),
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
	d.Add(dict.NewEntry(tag.New(0x0010, 0x0010), "Patient Name", "PatientName", vm.VM1, false, vr.PN))
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
