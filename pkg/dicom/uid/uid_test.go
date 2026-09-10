// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package uid_test

import (
	"sort"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

const testImplicitVRLittleLE = "1.2.840.10008.1.2"

func TestNew(t *testing.T) {
	u := uid.New(testImplicitVRLittleLE, "Implicit VR Little Endian", uid.TypeTransferSyntax, false)

	if u.UID() != testImplicitVRLittleLE {
		t.Errorf("UID() = %q, want %q", u.UID(), testImplicitVRLittleLE)
	}
	if u.Name() != "Implicit VR Little Endian" {
		t.Errorf("Name() = %q, want %q", u.Name(), "Implicit VR Little Endian")
	}
	if u.Type() != uid.TypeTransferSyntax {
		t.Errorf("Type() = %v, want %v", u.Type(), uid.TypeTransferSyntax)
	}
	if u.IsRetired() {
		t.Error("IsRetired() = true, want false")
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		uid   string
		valid bool
	}{
		{"valid simple", testImplicitVRLittleLE, true},
		{"valid complex", "1.2.840.10008.5.1.4.1.1.2", true},
		{"valid single component", "1", true},
		{"valid zero component", "0", true},
		{"invalid empty", "", false},
		{"invalid too long", "1.2.3.4.5.6.7.8.9.10.11.12.13.14.15.16.17.18.19.20.21.22.23.24.25", false},
		{"invalid leading dot", ".1.2.3", false},
		{"invalid trailing dot", "1.2.3.", false},
		{"invalid double dot", "1..2.3", false},
		{"invalid characters", "1.2.ABC.3", false},
		{"invalid leading zero", "1.02.3", false},
		{"valid zero only", "1.0.3", true},
		{"invalid spaces", "1.2. 3", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uid.IsValid(tt.uid); got != tt.valid {
				t.Errorf("IsValid(%q) = %v, want %v", tt.uid, got, tt.valid)
			}
		})
	}
}

func TestParse(t *testing.T) {
	// Parse a standard UID.
	parsed := uid.Parse(testImplicitVRLittleLE, "Other Name", uid.TypeUnknown)
	if parsed != uid.ImplicitVRLittleEndian {
		t.Fatal("Parse() did not return the canonical standard UID")
	}

	// Parse unregistered UID
	unregistered := uid.Parse("1.2.3.4.5", "Custom UID", uid.TypeSOPInstance)
	if unregistered.UID() != "1.2.3.4.5" {
		t.Errorf("Parse() UID = %q, want %q", unregistered.UID(), "1.2.3.4.5")
	}
	if unregistered.Name() != "Custom UID" {
		t.Errorf("Parse() Name = %q, want %q", unregistered.Name(), "Custom UID")
	}
}

func TestAppend(t *testing.T) {
	base := uid.New("1.2.840.10008", "Base UID", uid.TypeUnknown, false)
	appended := uid.Append(base, 123)

	want := "1.2.840.10008.123"
	if appended.UID() != want {
		t.Errorf("Append() = %q, want %q", appended.UID(), want)
	}
	if appended.Type() != uid.TypeSOPInstance {
		t.Errorf("Append() Type = %v, want %v", appended.Type(), uid.TypeSOPInstance)
	}
}

func TestEquals(t *testing.T) {
	u1 := uid.New(testImplicitVRLittleLE, "Name1", uid.TypeTransferSyntax, false)
	u2 := uid.New(testImplicitVRLittleLE, "Name2", uid.TypeSOPClass, false)
	u3 := uid.New("1.2.840.10008.1.3", "Name3", uid.TypeTransferSyntax, false)

	if !u1.Equals(u2) {
		t.Error("Equals() = false, want true (same UID string)")
	}
	if u1.Equals(u3) {
		t.Error("Equals() = true, want false (different UID string)")
	}
	if u1.Equals(nil) {
		t.Error("Equals(nil) = true, want false")
	}
}

func TestStorageCategory(t *testing.T) {
	tests := []struct {
		name     string
		uid      *uid.UID
		category uid.StorageCategory
	}{
		{
			"image storage",
			uid.New("1.2.840.10008.5.1.4.1.1.2", "CT Image Storage", uid.TypeSOPClass, false),
			uid.StorageCategoryImage,
		},
		{
			"volume storage",
			uid.New("1.2.840.10008.5.1.4.1.1.9.1.1", "Ultrasound Multi-frame Volume Storage", uid.TypeSOPClass, false),
			uid.StorageCategoryVolume,
		},
		{
			"presentation state",
			uid.New("1.2.840.10008.5.1.4.1.1.11.1", "Grayscale Softcopy Presentation State Storage", uid.TypeSOPClass, false),
			uid.StorageCategoryPresentationState,
		},
		{
			"not storage",
			uid.New(testImplicitVRLittleLE, "Implicit VR Little Endian", uid.TypeTransferSyntax, false),
			uid.StorageCategoryNone,
		},
		{
			"private storage",
			uid.New("1.3.12.2.1107.5.2.32.35119", "Private Storage", uid.TypeSOPClass, false),
			uid.StorageCategoryPrivate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uid.StorageCategory(); got != tt.category {
				t.Errorf("StorageCategory() = %v, want %v", got, tt.category)
			}
		})
	}
}

func TestIsImageStorage(t *testing.T) {
	imageUID := uid.New("1.2.840.10008.5.1.4.1.1.2", "CT Image Storage", uid.TypeSOPClass, false)
	if !imageUID.IsImageStorage() {
		t.Error("IsImageStorage() = false, want true")
	}

	transferUID := uid.New(testImplicitVRLittleLE, "Transfer Syntax", uid.TypeTransferSyntax, false)
	if transferUID.IsImageStorage() {
		t.Error("IsImageStorage() = true, want false")
	}
}

func TestStandardEntriesAreSortedCanonicalSnapshot(t *testing.T) {
	entries := uid.StandardEntries()
	if len(entries) == 0 {
		t.Fatal("StandardEntries() returned no UIDs")
	}
	values := make([]string, len(entries))
	for i, value := range entries {
		values[i] = value.UID()
	}
	if !sort.StringsAreSorted(values) {
		t.Fatal("StandardEntries() is not sorted by UID")
	}
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		if _, duplicate := seen[value]; duplicate {
			t.Fatalf("StandardEntries() contains duplicate UID %q", value)
		}
		seen[value] = struct{}{}
	}
	parsed := uid.Parse(testImplicitVRLittleLE, "ignored", uid.TypeUnknown)
	if parsed != uid.ImplicitVRLittleEndian {
		t.Fatal("standard catalog did not preserve canonical object identity")
	}
	entries[0] = nil
	if uid.StandardEntries()[0] == nil {
		t.Fatal("StandardEntries() returned shared slice storage")
	}
}

func TestStandardUIDParseIsImmutable(t *testing.T) {
	standard := uid.ImplicitVRLittleEndian
	original := standard.UID()
	if err := standard.Parse("1.2.3.4.5"); err == nil {
		t.Fatal("Parse() on a standard UID should fail")
	}
	if standard.UID() != original {
		t.Fatal("Parse() changed the standard UID")
	}

	custom := uid.New("1.2.3.4.6", "Custom", uid.TypeUnknown, false)
	if err := custom.Parse(testImplicitVRLittleLE); err != nil {
		t.Fatalf("Parse() on a non-standard UID = %v", err)
	}
	if custom.UID() != original {
		t.Fatal("Parse() did not apply the canonical UID value")
	}
}

func TestRegistryOverlayMaskAndIsolation(t *testing.T) {
	first := uid.NewRegistry()
	second := uid.NewRegistry()
	if got, found := first.Resolve(testImplicitVRLittleLE); !found || got != uid.ImplicitVRLittleEndian {
		t.Fatalf("Resolve(standard) = %v, %v", got, found)
	}
	if err := first.Register(uid.ImplicitVRLittleEndian); err == nil {
		t.Fatal("Register() accepted an effective standard duplicate")
	}

	replacement := uid.New(testImplicitVRLittleLE, "Private implicit syntax", uid.TypeTransferSyntax, false)
	previous, err := first.Replace(replacement)
	if err != nil || previous != uid.ImplicitVRLittleEndian {
		t.Fatalf("Replace() = %v, %v", previous, err)
	}
	if got, found := first.Resolve(testImplicitVRLittleLE); !found || got != replacement {
		t.Fatalf("first.Resolve() = %v, %v", got, found)
	}
	if got, found := second.Resolve(testImplicitVRLittleLE); !found || got != uid.ImplicitVRLittleEndian {
		t.Fatalf("second.Resolve() = %v, %v", got, found)
	}
	if uid.Parse(testImplicitVRLittleLE, "ignored", uid.TypeUnknown) != uid.ImplicitVRLittleEndian {
		t.Fatal("registry replacement changed the standard catalog")
	}

	removed, found := first.Unregister(testImplicitVRLittleLE)
	if !found || removed != replacement {
		t.Fatalf("Unregister() = %v, %v", removed, found)
	}
	if got, found := first.Resolve(testImplicitVRLittleLE); found || got != nil {
		t.Fatalf("masked Resolve() = %v, %v", got, found)
	}
	if parsed := first.Parse(testImplicitVRLittleLE, "Fallback", uid.TypeUnknown); parsed.Name() != "Fallback" {
		t.Fatalf("Parse(masked standard).Name() = %q", parsed.Name())
	}

	custom := uid.New("1.2.3.4.5.6", "Test UID", uid.TypeSOPClass, false)
	if err := first.Register(custom); err != nil {
		t.Fatalf("Register(custom) error = %v", err)
	}
	if err := first.Register(custom); err == nil {
		t.Fatal("Register() accepted a custom duplicate")
	}
	if err := first.Register(nil); err == nil {
		t.Fatal("Register(nil) succeeded")
	}
	if err := first.Register(uid.New("invalid", "Invalid", uid.TypeUnknown, false)); err == nil {
		t.Fatal("Register(invalid) succeeded")
	}

	entries := first.Enumerate()
	values := make([]string, len(entries))
	for i, value := range entries {
		values[i] = value.UID()
	}
	if !sort.StringsAreSorted(values) {
		t.Fatal("Enumerate() is not sorted by UID")
	}
	entries[0] = nil
	if first.Enumerate()[0] == nil {
		t.Fatal("Enumerate() returned shared slice storage")
	}
}

func TestRegistryConcurrentAccess(t *testing.T) {
	registry := uid.NewRegistry()
	custom := uid.New("1.2.3.4.8", "Concurrent UID", uid.TypeSOPClass, false)
	if err := registry.Register(custom); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	var wg sync.WaitGroup
	for range 16 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			registry.Resolve(custom.UID())
		}()
		go func() {
			defer wg.Done()
			registry.Enumerate()
		}()
		go func() {
			defer wg.Done()
			_, _ = registry.Replace(uid.New(custom.UID(), "Replacement", uid.TypeSOPClass, false))
		}()
	}
	wg.Wait()

	if got, found := registry.Resolve(custom.UID()); !found || got.UID() != custom.UID() {
		t.Fatalf("Resolve() after concurrent access = %v, %v", got, found)
	}
}

func TestMustParse(t *testing.T) {
	// Should not panic for valid UID
	u := uid.MustParse(testImplicitVRLittleLE)
	if u.UID() != testImplicitVRLittleLE {
		t.Errorf("MustParse() = %q, want %q", u.UID(), testImplicitVRLittleLE)
	}

	// Should panic for invalid UID
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParse with invalid UID should panic")
		}
	}()
	uid.MustParse("invalid.uid")
}
