// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

//revive:disable:var-naming // package name must match public import path (pkg/dicom/dict)
package dict_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestDefaultContainsGenerated2026bStandardAndPrivateDictionaries(t *testing.T) {
	dictionary := dict.Default()
	if got := len(dictionary.Entries()); got != 5356 {
		t.Errorf("default standard dictionary entries = %d, want 5356", got)
	}

	creator := dictionary.GetPrivateCreator("MED NM")
	entry := dictionary.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1010, creator))
	if entry == nil || entry.Name() != "Unknown" {
		t.Fatalf("default MED NM (0011,xx10) entry = %#v, want Unknown", entry)
	}
}

func TestDefaultContainsSupplemental2026dEntries(t *testing.T) {
	dictionary := dict.Default()

	displayURI := dictionary.Lookup(tag.New(0x0040, 0xE012))
	if displayURI == nil || displayURI.Name() != "Display URI" || displayURI.Keyword() != "DisplayURI" {
		t.Fatalf("Lookup((0040,E012)) = %#v, want Display URI", displayURI)
	}
	if got := displayURI.VRs(); len(got) != 1 || got[0] != "UR" {
		t.Errorf("DisplayURI.VRs() = %v, want [UR]", got)
	}

	doseModel := dictionary.Lookup(tag.New(0x3004, 0x007F))
	if doseModel == nil || doseModel.Keyword() != "DoseCalculationModelName" {
		t.Fatalf("Lookup((3004,007F)) = %#v, want DoseCalculationModelName", doseModel)
	}
	if got := doseModel.VRs(); len(got) != 1 || got[0] != "LO" {
		t.Errorf("DoseCalculationModelName.VRs() = %v, want [LO]", got)
	}
}
