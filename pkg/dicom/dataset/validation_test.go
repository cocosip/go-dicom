// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func newValidationDisabledDataset() *dataset.Dataset {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	return ds
}

func TestDatasetValidateReturnsNestedPathAndCause(t *testing.T) {
	leaf := newValidationDisabledDataset()
	if err := leaf.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{testInvalidUID})); err != nil {
		t.Fatal(err)
	}

	inner := dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{leaf})
	middle := newValidationDisabledDataset()
	if err := middle.Add(inner); err != nil {
		t.Fatal(err)
	}

	outer := dataset.NewSequenceWithItems(tag.ReferencedStudySequence, []*dataset.Dataset{middle})
	root := newValidationDisabledDataset()
	if err := root.Add(outer); err != nil {
		t.Fatal(err)
	}

	err := root.Validate()
	if err == nil {
		t.Fatal("Validate() should reject the nested invalid UID")
	}
	wantPath := "(0008,1110)[0]/(0008,1140)[0]/(0008,1155)"
	if !strings.Contains(err.Error(), wantPath) {
		t.Fatalf("Validate() error = %q, want path %q", err, wantPath)
	}

	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("Validate() error type = %T, want *dataset.ValidationError", err)
	}
	if validationErr.Kind != dataset.ValidationValue {
		t.Fatalf("validation kind = %q, want %q", validationErr.Kind, dataset.ValidationValue)
	}
	var vrErr *vr.ValidationError
	if !errors.As(err, &vrErr) {
		t.Fatalf("Validate() should unwrap to *vr.ValidationError, got %T", err)
	}
}

func TestDatasetValidateUsesSortedFailFastOrder(t *testing.T) {
	ds := newValidationDisabledDataset()
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{strings.Repeat("A", 65)})); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testInvalidUID})); err != nil {
		t.Fatal(err)
	}

	err := ds.Validate()
	if err == nil {
		t.Fatal("Validate() should reject invalid values")
	}
	if !strings.Contains(err.Error(), "(0008,0016)") {
		t.Fatalf("Validate() error = %q, want first sorted tag (0008,0016)", err)
	}
}

func TestDatasetValidateChecksPublicVM(t *testing.T) {
	ds := newValidationDisabledDataset()
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName, "Doe^John"})); err != nil {
		t.Fatal(err)
	}

	err := ds.Validate()
	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("Validate() error = %v, want *dataset.ValidationError", err)
	}
	if validationErr.Kind != dataset.ValidationVM {
		t.Fatalf("validation kind = %q, want %q", validationErr.Kind, dataset.ValidationVM)
	}
}

func TestDatasetValidateSkipsVMForEmptyPrivateUnknownAndBulkValues(t *testing.T) {
	tests := []struct {
		name string
		elem element.Element
	}{
		{name: "empty public value", elem: element.NewString(tag.PatientName, vr.PN, nil)},
		{name: "private explicit VR", elem: element.NewString(tag.New(0x0011, 0x1010), vr.UI, []string{testStudyInstanceUID, testAlternateUID})},
		{name: "unknown public tag", elem: element.NewString(tag.New(0x7776, 0x0010), vr.UI, []string{testStudyInstanceUID, testAlternateUID})},
		{name: "other byte", elem: element.NewOtherByte(tag.PixelData, []byte{1, 2, 3})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := newValidationDisabledDataset()
			if err := ds.Add(tt.elem); err != nil {
				t.Fatal(err)
			}
			if err := ds.Validate(); err != nil {
				t.Fatalf("Validate() error = %v, want nil", err)
			}
		})
	}
}

func TestDatasetValidateStillChecksPrivateValueSyntax(t *testing.T) {
	ds := newValidationDisabledDataset()
	if err := ds.Add(element.NewString(tag.New(0x0011, 0x1010), vr.UI, []string{testInvalidUID})); err != nil {
		t.Fatal(err)
	}

	err := ds.Validate()
	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) || validationErr.Kind != dataset.ValidationValue {
		t.Fatalf("Validate() error = %v, want value validation error", err)
	}
}

func TestSequenceValidateSkipsNilAndEmptyItems(t *testing.T) {
	seq := dataset.NewSequenceWithItems(tag.ReferencedStudySequence, []*dataset.Dataset{nil, dataset.New()})
	if err := seq.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}

func TestSequenceValidateReportsDatasetCycleAsStructuralError(t *testing.T) {
	item := newValidationDisabledDataset()
	sequence := dataset.NewSequence(tag.ReferencedImageSequence)
	if err := item.Add(sequence); err != nil {
		t.Fatal(err)
	}
	sequence.AddItem(item)

	err := sequence.Validate()
	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) || validationErr.Kind != dataset.ValidationStructural {
		t.Fatalf("Validate() error = %v, want structural cycle error", err)
	}
	if !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("Validate() error = %v, want cycle cause", err)
	}
}

func TestDatasetAddReportsSequenceCycleAsStructuralError(t *testing.T) {
	item := newValidationDisabledDataset()
	sequence := dataset.NewSequence(tag.ReferencedImageSequence)
	if err := item.Add(sequence); err != nil {
		t.Fatal(err)
	}
	sequence.AddItem(item)

	err := dataset.New().Add(sequence)
	var validationErr *dataset.ValidationError
	if !errors.As(err, &validationErr) || validationErr.Kind != dataset.ValidationStructural {
		t.Fatalf("Add() error = %v, want structural cycle error", err)
	}
}
