// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package sr

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewContentItemContainerRejectsNilChild(t *testing.T) {
	code := NewCodeItem("121071", "DCM", "Finding")

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic: %v", r)
		}
	}()

	_, err := NewContentItemContainer(code, RelationshipContains, ContinuitySeparate, nil)
	if err == nil {
		t.Fatal("expected error when child item is nil")
	}
}

func TestNewStructuredReportRejectsNilChild(t *testing.T) {
	rootCode := NewCodeItem("113704", "DCM", "SR Document")

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic: %v", r)
		}
	}()

	_, err := NewStructuredReport(rootCode, nil)
	if err == nil {
		t.Fatal("expected error when root child item is nil")
	}
}

func TestContentItemChildrenReturnsErrorForInvalidContentSequence(t *testing.T) {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	if err := ds.Add(element.NewString(tag.ContentSequence, vr.LO, []string{"not-a-sequence"})); err != nil {
		t.Fatalf("failed to add invalid ContentSequence: %v", err)
	}

	item := NewContentItemFromDataset(ds)
	children, err := item.Children()
	if err == nil {
		t.Fatal("expected error for invalid ContentSequence element")
	}
	if children != nil {
		t.Fatalf("expected nil children on error, got %v", children)
	}
	if !strings.Contains(err.Error(), "not Sequence") {
		t.Fatalf("expected sequence type error, got %v", err)
	}
}

func TestNewContentItemCodeRejectsNilValue(t *testing.T) {
	code := NewCodeItem("121032", "DCM", "Subject Sex")

	item, err := NewContentItemCode(code, RelationshipContains, nil)
	if err == nil {
		t.Fatalf("NewContentItemCode() = %v, want an error for a nil code value", item)
	}
	if item != nil {
		t.Fatalf("NewContentItemCode() returned an item with a nil code value")
	}
}

func TestNewContentItemNumericAlwaysEmitsMeasuredValueSequence(t *testing.T) {
	code := NewCodeItem("121033", "DCM", "Subject Age")
	tests := []struct {
		name      string
		value     *MeasuredValue
		wantItems int
	}{
		{name: "with value", value: NewMeasuredValue(42, NewCodeItem("a", "UCUM", "year")), wantItems: 1},
		{name: "without value", value: nil, wantItems: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := NewContentItemNumeric(code, RelationshipContains, tt.value)
			if err != nil {
				t.Fatalf("NewContentItemNumeric() error = %v", err)
			}
			sequence, err := item.Dataset().GetSequence(tag.MeasuredValueSequence)
			if err != nil {
				t.Fatalf("Measured Value Sequence missing: %v", err)
			}
			if sequence.Count() != tt.wantItems {
				t.Fatalf("Measured Value Sequence items = %d, want %d", sequence.Count(), tt.wantItems)
			}
		})
	}
}

func TestNewStructuredReportPopulatesMandatoryModules(t *testing.T) {
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"))
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}

	path := filepath.Join(t.TempDir(), "standalone-report.dcm")
	if err := report.Save(path); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	opened, err := Open(path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	ds := opened.Dataset()
	// Type 2 attributes must be present; an empty value still counts as present.
	for _, required := range []*tag.Tag{
		tag.PatientName,
		tag.PatientID,
		tag.PatientBirthDate,
		tag.PatientSex,
		tag.StudyDate,
		tag.StudyTime,
		tag.ReferringPhysicianName,
		tag.StudyID,
		tag.AccessionNumber,
		tag.SeriesNumber,
		tag.Manufacturer,
	} {
		if !ds.Contains(required) {
			t.Errorf("reopened report is missing Type 2 attribute %s", required)
		}
	}
	// Type 1 attributes must carry a value after the round trip.
	for _, required := range []*tag.Tag{tag.StudyInstanceUID, tag.SeriesInstanceUID} {
		value, ok := ds.GetString(required)
		if !ok || !uid.IsValid(value) {
			t.Errorf("%s = %q, %v, want a generated UID", required, value, ok)
		}
	}
	if modality, _ := ds.GetString(tag.Modality); modality != "SR" {
		t.Errorf("Modality = %q, want SR", modality)
	}
}
