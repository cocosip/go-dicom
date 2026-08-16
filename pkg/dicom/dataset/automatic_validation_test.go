// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestDatasetAddRejectsNilTagWithoutMutation(t *testing.T) {
	ds := dataset.New()
	err := ds.Add(element.NewString(nil, vr.UI, []string{testStudyInstanceUID}))
	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) || validationErr.Kind != dataset.ValidationStructural {
		t.Fatalf("Add() error = %v, want structural validation error", err)
	}
	if ds.Count() != 0 {
		t.Fatalf("Count() = %d after failed Add(), want 0", ds.Count())
	}
}

func invalidUIDElement(t *tag.Tag) element.Element {
	return element.NewString(t, vr.UI, []string{testInvalidUID})
}

func TestDatasetAutomaticValidationDefaultsToEnabled(t *testing.T) {
	for _, ds := range []*dataset.Dataset{dataset.New(), &dataset.Dataset{}} {
		if !ds.AutoValidate() {
			t.Fatal("AutoValidate() = false, want true")
		}
		if err := ds.Add(invalidUIDElement(tag.StudyInstanceUID)); err == nil {
			t.Fatal("Add() should reject an invalid UID by default")
		}
		if ds.Count() != 0 {
			t.Fatalf("Count() = %d after failed Add(), want 0", ds.Count())
		}
	}
}

func TestDatasetAutomaticValidationIgnoresLegacyGlobalSwitch(t *testing.T) {
	oldValidation := vr.PerformValidation
	t.Cleanup(func() { vr.PerformValidation = oldValidation })
	vr.PerformValidation = false

	ds := dataset.New()
	if err := ds.Add(invalidUIDElement(tag.StudyInstanceUID)); err == nil {
		t.Fatal("Add() should reject an invalid UID regardless of the legacy global switch")
	}
}

func TestDatasetDisabledAutomaticValidationAllowsInsertButNotExplicitValidation(t *testing.T) {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	if ds.AutoValidate() {
		t.Fatal("AutoValidate() = true after disabling")
	}
	if err := ds.Add(invalidUIDElement(tag.StudyInstanceUID)); err != nil {
		t.Fatalf("Add() error = %v with automatic validation disabled", err)
	}
	if err := ds.Validate(); err == nil {
		t.Fatal("explicit Validate() should reject the invalid UID")
	}
}

func TestDatasetAddOrUpdatePreservesExistingElementOnValidationFailure(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyInstanceUID})); err != nil {
		t.Fatal(err)
	}
	if err := ds.AddOrUpdate(invalidUIDElement(tag.StudyInstanceUID)); err == nil {
		t.Fatal("AddOrUpdate() should reject an invalid UID")
	}
	got, ok := ds.GetString(tag.StudyInstanceUID)
	if !ok || got != testStudyInstanceUID {
		t.Fatalf("StudyInstanceUID = %q, %v; want original value", got, ok)
	}
}

func TestNewWithElementsReturnsNoPartialDatasetOnValidationFailure(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{testStudyInstanceUID}),
		invalidUIDElement(tag.SOPInstanceUID),
	})
	if err == nil {
		t.Fatal("NewWithElements() should reject an invalid UID")
	}
	if ds != nil {
		t.Fatalf("NewWithElements() dataset = %v, want nil on failure", ds)
	}
}

func TestDatasetMergeIsAtomicOnValidationFailure(t *testing.T) {
	destination := dataset.New()
	if err := destination.Add(element.NewString(tag.PatientID, vr.LO, []string{"existing"})); err != nil {
		t.Fatal(err)
	}

	source := newValidationDisabledDataset()
	if err := source.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testStudyInstanceUID})); err != nil {
		t.Fatal(err)
	}
	if err := source.Add(invalidUIDElement(tag.SOPInstanceUID)); err != nil {
		t.Fatal(err)
	}

	if err := destination.Merge(source, true); err == nil {
		t.Fatal("Merge() should reject the invalid UID")
	}
	if destination.Contains(tag.SOPClassUID) || destination.Contains(tag.SOPInstanceUID) {
		t.Fatal("Merge() partially changed the destination after validation failure")
	}
}

func TestDatasetCloneAndFilterPreserveValidationModeAndInvalidContent(t *testing.T) {
	source := newValidationDisabledDataset()
	if err := source.Add(invalidUIDElement(tag.StudyInstanceUID)); err != nil {
		t.Fatal(err)
	}

	clone := source.Clone()
	filtered := source.Filter(func(element.Element) bool { return true })
	for name, got := range map[string]*dataset.Dataset{"clone": clone, "filter": filtered} {
		if got.AutoValidate() {
			t.Fatalf("%s AutoValidate() = true, want false", name)
		}
		if !got.Contains(tag.StudyInstanceUID) {
			t.Fatalf("%s dropped invalid readable content", name)
		}
	}
}
