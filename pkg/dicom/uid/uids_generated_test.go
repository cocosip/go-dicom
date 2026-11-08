// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

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
			testImplicitVRLittleLE,
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

// Test that generated UIDs are available from the immutable standard catalog.
func TestGeneratedUIDsInStandardCatalog(t *testing.T) {
	parsed := uid.Parse(testImplicitVRLittleLE, "Implicit VR Little Endian", uid.TypeUnknown)

	if parsed != uid.ImplicitVRLittleEndian {
		t.Error("parsed UID should be the generated canonical object")
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

func TestGeneratedStandardUIDCount(t *testing.T) {
	if got := len(uid.StandardEntries()); got != 1990 {
		t.Errorf("len(StandardEntries()) = %d, want 1990", got)
	}
}

func TestGeneratedSupplemental2026dUIDs(t *testing.T) {
	forProcessing := uid.Parse("1.2.840.10008.5.1.4.1.1.2.3", "ignored", uid.TypeUnknown)
	if forProcessing.Name() != "CT Image Storage - For Processing" {
		t.Errorf("CTImageStorageForProcessing.Name() = %q", forProcessing.Name())
	}
	if !forProcessing.IsImageStorage() {
		t.Error("CTImageStorageForProcessing.IsImageStorage() = false, want true")
	}

	frameOfReference := uid.Parse("1.2.840.10008.1.4.1.1", "ignored", uid.TypeUnknown)
	if frameOfReference.Name() != "Talairach Brain Atlas Frame of Reference" {
		t.Errorf("TalairachBrainAtlas.Name() = %q", frameOfReference.Name())
	}
	if frameOfReference.Type() != uid.TypeFrameOfReference {
		t.Errorf("TalairachBrainAtlas.Type() = %v, want TypeFrameOfReference", frameOfReference.Type())
	}

	template := uid.Parse("1.2.840.10008.9.1", "ignored", uid.TypeUnknown)
	if template.Type() != uid.TypeTemplate {
		t.Errorf("DocumentTemplateID91.Type() = %v, want TypeTemplate", template.Type())
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
