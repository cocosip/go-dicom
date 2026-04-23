// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewWithElementsDuplicateTagUsesLatestValue(t *testing.T) {
	ds := NewWithElements([]element.Element{
		element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}),
		element.NewString(tag.PatientName, vr.PN, []string{"Smith^Jane"}),
	})

	got, ok := ds.GetString(tag.PatientName)
	if !ok {
		t.Fatal("expected PatientName to be present")
	}
	if got != "Smith^Jane" {
		t.Fatalf("expected latest PatientName to win, got %q", got)
	}
}

func TestNewWithElementsIgnoresNilElement(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("expected nil element to be ignored, got panic: %v", r)
		}
	}()

	ds := NewWithElements([]element.Element{nil})
	if ds == nil {
		t.Fatal("expected dataset to be created")
	}
	if ds.Count() != 0 {
		t.Fatalf("expected nil element to be ignored, got %d elements", ds.Count())
	}
}
