// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dictif_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dictif"
)

// mockTag implements dictif.Tag for testing
type mockTag struct {
	group   uint16
	element uint16
}

func (m *mockTag) Group() uint16        { return m.group }
func (m *mockTag) Element() uint16      { return m.element }
func (m *mockTag) ToUint32() uint32     { return (uint32(m.group) << 16) | uint32(m.element) }

// mockEntry implements dictif.Entry for testing
type mockEntry struct {
	name    string
	keyword string
	vrs     []string
	vm      string
}

func (m *mockEntry) Name() string     { return m.name }
func (m *mockEntry) Keyword() string  { return m.keyword }
func (m *mockEntry) VRs() []string    { return m.vrs }
func (m *mockEntry) VM() string       { return m.vm }

// mockPrivateCreator implements dictif.PrivateCreator for testing
type mockPrivateCreator struct {
	creator string
}

func (m *mockPrivateCreator) Creator() string { return m.creator }

// mockLookup implements dictif.Lookup for testing
type mockLookup struct {
	entries map[uint32]dictif.Entry
	tags    map[string]dictif.Tag
	creators map[string]dictif.PrivateCreator
}

func newMockLookup() *mockLookup {
	return &mockLookup{
		entries:  make(map[uint32]dictif.Entry),
		tags:     make(map[string]dictif.Tag),
		creators: make(map[string]dictif.PrivateCreator),
	}
}

func (m *mockLookup) LookupTag(tag dictif.Tag) dictif.Entry {
	return m.entries[tag.ToUint32()]
}

func (m *mockLookup) LookupKeyword(keyword string) dictif.Tag {
	return m.tags[keyword]
}

func (m *mockLookup) GetPrivateCreator(creator string) dictif.PrivateCreator {
	if pc, ok := m.creators[creator]; ok {
		return pc
	}
	pc := &mockPrivateCreator{creator: creator}
	m.creators[creator] = pc
	return pc
}

func TestMockLookup(t *testing.T) {
	lookup := newMockLookup()

	// Add test data
	tag := &mockTag{group: 0x0010, element: 0x0010}
	entry := &mockEntry{
		name:    "Patient Name",
		keyword: "PatientName",
		vrs:     []string{"PN"},
		vm:      "1",
	}
	lookup.entries[tag.ToUint32()] = entry
	lookup.tags["PatientName"] = tag

	// Test LookupTag
	foundEntry := lookup.LookupTag(tag)
	if foundEntry == nil {
		t.Fatal("LookupTag returned nil")
	}
	if foundEntry.Name() != "Patient Name" {
		t.Errorf("Expected 'Patient Name', got '%s'", foundEntry.Name())
	}

	// Test LookupKeyword
	foundTag := lookup.LookupKeyword("PatientName")
	if foundTag == nil {
		t.Fatal("LookupKeyword returned nil")
	}
	if foundTag.Group() != 0x0010 || foundTag.Element() != 0x0010 {
		t.Errorf("Expected (0010,0010), got (%04X,%04X)", foundTag.Group(), foundTag.Element())
	}

	// Test GetPrivateCreator
	pc := lookup.GetPrivateCreator("TEST_CREATOR")
	if pc == nil {
		t.Fatal("GetPrivateCreator returned nil")
	}
	if pc.Creator() != "TEST_CREATOR" {
		t.Errorf("Expected 'TEST_CREATOR', got '%s'", pc.Creator())
	}

	// Verify caching
	pc2 := lookup.GetPrivateCreator("TEST_CREATOR")
	if pc != pc2 {
		t.Error("GetPrivateCreator should return cached instance")
	}
}

func TestGlobalLookup(t *testing.T) {
	// Save original
	original := dictif.GlobalLookup()
	defer dictif.SetGlobalLookup(original)

	// Test with nil
	dictif.SetGlobalLookup(nil)
	if dictif.GlobalLookup() != nil {
		t.Error("GlobalLookup should return nil when not set")
	}

	// Test with mock
	mock := newMockLookup()
	dictif.SetGlobalLookup(mock)
	if dictif.GlobalLookup() != mock {
		t.Error("GlobalLookup should return the set mock")
	}
}

func TestInterfaceImplementations(_ *testing.T) {
	// Verify that our mock types implement the interfaces
	var _ dictif.Tag = (*mockTag)(nil)
	var _ dictif.Entry = (*mockEntry)(nil)
	var _ dictif.PrivateCreator = (*mockPrivateCreator)(nil)
	var _ dictif.Lookup = (*mockLookup)(nil)
}
