// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package dataset_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	testCharsetLatin1      = "ISO_IR 100"
	testCharsetUTF8        = "ISO_IR 192"
	testCharsetISO2022IR87 = "ISO 2022 IR 87"
)

// TestSetSpecificCharacterSetRejectsInvalidCombinations verifies PS3.3 C.12.1.1:
// Value 1 is the replacement repertoire, so a multi-byte-only set cannot be
// Value 1 and only ISO 2022 code extensions may follow it.
func TestSetSpecificCharacterSetRejectsInvalidCombinations(t *testing.T) {
	tests := []struct {
		name   string
		values []string
	}{
		{"multi-byte set as value 1", []string{testCharsetISO2022IR87, "ISO 2022 IR 6"}},
		{"single-value set with extension", []string{testCharsetUTF8, testCharsetISO2022IR87}},
		{"non ISO 2022 extension", []string{testCharsetLatin1, "ISO_IR 144"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := dataset.New().SetSpecificCharacterSet(test.values...); err == nil {
				t.Fatalf("SetSpecificCharacterSet(%v) accepted an invalid combination", test.values)
			}
		})
	}

	if err := dataset.New().SetSpecificCharacterSet(testCharsetLatin1, testCharsetISO2022IR87); err != nil {
		t.Fatalf("SetSpecificCharacterSet() rejected a standard combination: %v", err)
	}
}

// TestEnsureUTF8CharacterSetDoesNotRelabelExistingDeclaration verifies that a
// dataset which already declares another character set is not silently
// relabeled: the existing values would be reinterpreted (PS3.5 6.1.2.2).
func TestEnsureUTF8CharacterSetDoesNotRelabelExistingDeclaration(t *testing.T) {
	ds := dataset.New()
	if err := ds.SetSpecificCharacterSet(testCharsetLatin1); err != nil {
		t.Fatalf("SetSpecificCharacterSet() error = %v", err)
	}
	// The legacy constructor falls back to UTF-8 bytes for text the default
	// repertoire cannot represent.
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"张三"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	if err := ds.EnsureUTF8CharacterSet(); err == nil {
		t.Fatal("EnsureUTF8CharacterSet() relabeled a dataset that declares ISO_IR 100")
	}
	if declared := ds.SpecificCharacterSet(); len(declared) != 1 || declared[0] != testCharsetLatin1 {
		t.Fatalf("Specific Character Set = %v, want ISO_IR 100 unchanged", declared)
	}

	undeclared := dataset.New()
	if err := undeclared.Add(element.NewString(tag.PatientName, vr.PN, []string{"张三"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := undeclared.EnsureUTF8CharacterSet(); err != nil {
		t.Fatalf("EnsureUTF8CharacterSet() rejected an undeclared dataset: %v", err)
	}
	if declared := undeclared.SpecificCharacterSet(); len(declared) != 1 || declared[0] != testCharsetUTF8 {
		t.Fatalf("Specific Character Set = %v, want ISO_IR 192", declared)
	}
}
