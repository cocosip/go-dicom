// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package uid_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

func TestNew(t *testing.T) {
	u := uid.New("1.2.840.10008.1.2", "Implicit VR Little Endian", uid.TypeTransferSyntax, false)

	if u.UID() != "1.2.840.10008.1.2" {
		t.Errorf("UID() = %q, want %q", u.UID(), "1.2.840.10008.1.2")
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
		{"valid simple", "1.2.840.10008.1.2", true},
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
	// Register a UID
	registered := uid.New("1.2.840.10008.1.2", "Implicit VR Little Endian", uid.TypeTransferSyntax, false)
	uid.Register(registered)

	// Parse registered UID
	parsed := uid.Parse("1.2.840.10008.1.2", "Other Name", uid.TypeUnknown)
	if parsed.UID() != "1.2.840.10008.1.2" {
		t.Errorf("Parse() UID = %q, want %q", parsed.UID(), "1.2.840.10008.1.2")
	}
	if parsed.Name() != "Implicit VR Little Endian" {
		t.Errorf("Parse() Name = %q, want %q (should use registered name)", parsed.Name(), "Implicit VR Little Endian")
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
	u1 := uid.New("1.2.840.10008.1.2", "Name1", uid.TypeTransferSyntax, false)
	u2 := uid.New("1.2.840.10008.1.2", "Name2", uid.TypeSOPClass, false)
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
			uid.New("1.2.840.10008.1.2", "Implicit VR Little Endian", uid.TypeTransferSyntax, false),
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

	transferUID := uid.New("1.2.840.10008.1.2", "Transfer Syntax", uid.TypeTransferSyntax, false)
	if transferUID.IsImageStorage() {
		t.Error("IsImageStorage() = true, want false")
	}
}

func TestRegisterAndEnumerate(t *testing.T) {
	// Count initial UIDs
	initialCount := len(uid.Enumerate())

	// Register new UID
	newUID := uid.New("1.2.3.4.5.6", "Test UID", uid.TypeSOPClass, false)
	uid.Register(newUID)

	// Check it's registered
	allUIDs := uid.Enumerate()
	if len(allUIDs) != initialCount+1 {
		t.Errorf("Enumerate() returned %d UIDs, want %d", len(allUIDs), initialCount+1)
	}

	// Verify we can find it
	found := false
	for _, u := range allUIDs {
		if u.UID() == "1.2.3.4.5.6" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Registered UID not found in Enumerate()")
	}
}

func TestMustParse(t *testing.T) {
	// Should not panic for valid UID
	u := uid.MustParse("1.2.840.10008.1.2")
	if u.UID() != "1.2.840.10008.1.2" {
		t.Errorf("MustParse() = %q, want %q", u.UID(), "1.2.840.10008.1.2")
	}

	// Should panic for invalid UID
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParse with invalid UID should panic")
		}
	}()
	uid.MustParse("invalid.uid")
}
