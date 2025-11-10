// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package uid_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// Test a few standard UIDs to ensure they were generated correctly
func TestGeneratedUIDs(t *testing.T) {
	tests := []struct {
		name     string
		uid      *uid.UID
		wantUID  string
		wantType uid.Type
	}{
		{
			"Verification",
			uid.Verification,
			"1.2.840.10008.1.1",
			uid.TypeSOPClass,
		},
		{
			"ImplicitVRLittleEndian",
			uid.ImplicitVRLittleEndian,
			"1.2.840.10008.1.2",
			uid.TypeTransferSyntax,
		},
		{
			"ExplicitVRLittleEndian",
			uid.ExplicitVRLittleEndian,
			"1.2.840.10008.1.2.1",
			uid.TypeTransferSyntax,
		},
		{
			"JPEGBaseline8Bit",
			uid.JPEGBaseline8Bit,
			"1.2.840.10008.1.2.4.50",
			uid.TypeTransferSyntax,
		},
		{
			"CTImageStorage",
			uid.CTImageStorage,
			"1.2.840.10008.5.1.4.1.1.2",
			uid.TypeSOPClass,
		},
		{
			"MRImageStorage",
			uid.MRImageStorage,
			"1.2.840.10008.5.1.4.1.1.4",
			uid.TypeSOPClass,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.uid.UID() != tt.wantUID {
				t.Errorf("UID() = %q, want %q", tt.uid.UID(), tt.wantUID)
			}
			if tt.uid.Type() != tt.wantType {
				t.Errorf("Type() = %v, want %v", tt.uid.Type(), tt.wantType)
			}
		})
	}
}

// Test that UIDs are registered in the global registry
func TestGeneratedUIDsRegistered(t *testing.T) {
	// Parse a standard UID - should return the registered instance
	parsed := uid.Parse("1.2.840.10008.1.2", "Implicit VR Little Endian", uid.TypeUnknown)

	// Verify it equals the registered constant
	if !parsed.Equals(uid.ImplicitVRLittleEndian) {
		t.Error("Parsed UID should equal registered constant")
	}

	// Verify the full name from the registered instance
	if uid.ImplicitVRLittleEndian.Name() != "Implicit VR Little Endian: Default Transfer Syntax for DICOM" {
		t.Errorf("ImplicitVRLittleEndian.Name() = %q, want full name", uid.ImplicitVRLittleEndian.Name())
	}
}

// Test transfer syntax UIDs
func TestTransferSyntaxUIDs(t *testing.T) {
	transferSyntaxes := []*uid.UID{
		uid.ImplicitVRLittleEndian,
		uid.ExplicitVRLittleEndian,
		uid.DeflatedExplicitVRLittleEndian,
		uid.JPEGBaseline8Bit,
		uid.JPEG2000Lossless,
		uid.RLELossless,
	}

	for _, ts := range transferSyntaxes {
		if ts.Type() != uid.TypeTransferSyntax {
			t.Errorf("%s: Type() = %v, want TypeTransferSyntax", ts.Name(), ts.Type())
		}
	}
}

// Test image storage UIDs
func TestImageStorageUIDs(t *testing.T) {
	imageStorages := []*uid.UID{
		uid.CTImageStorage,
		uid.MRImageStorage,
		uid.UltrasoundImageStorage,
		uid.SecondaryCaptureImageStorage,
		uid.XRayAngiographicImageStorage,
	}

	for _, storage := range imageStorages {
		if storage.Type() != uid.TypeSOPClass {
			t.Errorf("%s: Type() = %v, want TypeSOPClass", storage.Name(), storage.Type())
		}
		if !storage.IsImageStorage() {
			t.Errorf("%s: IsImageStorage() = false, want true", storage.Name())
		}
	}
}

// Test retired UIDs
func TestRetiredUIDs(t *testing.T) {
	if !uid.ExplicitVRBigEndianRETIRED.IsRetired() {
		t.Error("ExplicitVRBigEndianRETIRED: IsRetired() = false, want true")
	}

	if uid.ImplicitVRLittleEndian.IsRetired() {
		t.Error("ImplicitVRLittleEndian: IsRetired() = true, want false")
	}
}
