// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestStructuredReportValidateAcceptsNestedReport(t *testing.T) {
	leaf, err := NewContentItemText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		"No abnormalities detected",
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	section, err := NewContentItemContainer(
		NewCodeItem("121070", "DCM", "Findings"),
		RelationshipContains,
		ContinuitySeparate,
		leaf,
	)
	if err != nil {
		t.Fatalf("NewContentItemContainer() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), section)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}

	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}

func TestStructuredReportValidateRejectsInvalidRoot(t *testing.T) {
	tests := []struct {
		name string
		edit func(*dataset.Dataset) error
		want string
	}{
		{
			name: "non-container value type",
			edit: func(ds *dataset.Dataset) error {
				return ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeText)}))
			},
			want: "root must have Value Type CONTAINER",
		},
		{
			name: "relationship type present",
			edit: func(ds *dataset.Dataset) error {
				return ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(RelationshipContains)}))
			},
			want: "root must not have Relationship Type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"))
			if err != nil {
				t.Fatalf("NewStructuredReport() error = %v", err)
			}
			if err := tt.edit(report.Dataset()); err != nil {
				t.Fatalf("edit root dataset: %v", err)
			}

			err = report.Validate()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func TestStructuredReportValidateRejectsMissingAndUnknownChildRelationships(t *testing.T) {
	tests := []struct {
		name         string
		relationship *string
		want         string
	}{
		{name: "missing", want: "relationship type not found"},
		{name: "unknown", relationship: stringPointer("INVALID"), want: "unknown relationship type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child, err := NewContentItemText(NewCodeItem("121071", "DCM", "Finding"), RelationshipContains, "normal")
			if err != nil {
				t.Fatalf("NewContentItemText() error = %v", err)
			}
			child.Dataset().Remove(tag.RelationshipType)
			if tt.relationship != nil {
				if err := child.Dataset().AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{*tt.relationship})); err != nil {
					t.Fatalf("set relationship: %v", err)
				}
			}
			report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), child)
			if err != nil {
				t.Fatalf("NewStructuredReport() error = %v", err)
			}

			err = report.Validate()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func TestStructuredReportValidateReportsNestedContentPath(t *testing.T) {
	leaf, err := NewContentItemText(NewCodeItem("121071", "DCM", "Finding"), RelationshipContains, "normal")
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	leaf.Dataset().Remove(tag.TextValue)
	section, err := NewContentItemContainer(
		NewCodeItem("121070", "DCM", "Findings"),
		RelationshipContains,
		ContinuitySeparate,
		leaf,
	)
	if err != nil {
		t.Fatalf("NewContentItemContainer() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), section)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}

	err = report.Validate()
	for _, want := range []string{"ContentSequence[0]", "ContentSequence[0].ContentSequence[0]", "text value not found"} {
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("Validate() error = %v, want substring %q", err, want)
		}
	}
}

func TestContentItemValidateRejectsValueTypeContradiction(t *testing.T) {
	item, err := NewContentItemText(NewCodeItem("121071", "DCM", "Finding"), RelationshipContains, "normal")
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	code := NewCodeItem("F-01775", "SRT", "Findings")
	if err := item.Dataset().AddOrUpdate(dataset.NewSequenceWithItems(tag.ConceptCodeSequence, []*dataset.Dataset{code.Dataset()})); err != nil {
		t.Fatalf("add contradictory value: %v", err)
	}

	err = item.Validate()
	if err == nil || !strings.Contains(err.Error(), "Concept Code Sequence contradicts Value Type TEXT") {
		t.Fatalf("Validate() error = %v, want value type contradiction", err)
	}
}

func TestContentItemValidateAcceptsByReferenceItem(t *testing.T) {
	ds := dataset.New()
	if err := ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(RelationshipSelectedFrom)})); err != nil {
		t.Fatalf("set relationship: %v", err)
	}
	if err := ds.AddOrUpdate(element.NewUnsignedLong(tag.ReferencedContentItemIdentifier, []uint32{1, 3, 2})); err != nil {
		t.Fatalf("set Referenced Content Item Identifier: %v", err)
	}

	if err := NewContentItemFromDataset(ds).Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}

func TestStructuredReportAddPreservesMalformedContentSequence(t *testing.T) {
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"))
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	report.Dataset().SetAutoValidate(false)
	if err := report.Dataset().AddOrUpdate(element.NewString(tag.ContentSequence, vr.LO, []string{"malformed"})); err != nil {
		t.Fatalf("add malformed Content Sequence: %v", err)
	}
	report.Dataset().SetAutoValidate(true)
	child, err := NewContentItemText(NewCodeItem("121071", "DCM", "Finding"), RelationshipContains, "normal")
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}

	err = report.Add(child)
	if err == nil {
		t.Fatal("Add() error = nil, want malformed Content Sequence error")
	}
	value, ok := report.Dataset().Get(tag.ContentSequence)
	if !ok {
		t.Fatal("malformed Content Sequence was removed")
	}
	if _, ok := value.(*element.String); !ok {
		t.Fatalf("Content Sequence type = %T, want original *element.String", value)
	}
}

func TestCodeItemValidateReportsConstructionError(t *testing.T) {
	code := NewCodeItem("121071", "DCM", strings.Repeat("x", 65))

	err := code.Validate()
	if err == nil || !strings.Contains(err.Error(), "Code Meaning") {
		t.Fatalf("Validate() error = %v, want Code Meaning construction error", err)
	}
}

func TestMeasuredValueValidateRequiresNumericValueAndUnits(t *testing.T) {
	tests := []struct {
		name string
		edit func(*dataset.Dataset)
		want string
	}{
		{name: "numeric value", edit: func(ds *dataset.Dataset) { ds.Remove(tag.NumericValue) }, want: "numeric value not found"},
		{name: "units", edit: func(ds *dataset.Dataset) { ds.Remove(tag.MeasurementUnitsCodeSequence) }, want: "measurement units code not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := NewMeasuredValue(12.5, NewCodeItem("mm", "UCUM", "millimeter"))
			tt.edit(value.Dataset())
			err := value.Validate()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func TestReferencedSOPValidateRequiresValidUIDs(t *testing.T) {
	tests := []struct {
		name        string
		instanceUID string
		classUID    string
		want        string
	}{
		{name: "instance", instanceUID: "not-a-uid", classUID: "1.2.840.10008.5.1.4.1.1.2", want: "invalid Referenced SOP Instance UID"},
		{name: "class", instanceUID: "1.2.826.0.1.3680043.10.543.1", classUID: "not-a-uid", want: "invalid Referenced SOP Class UID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			ds.SetAutoValidate(false)
			if err := ds.AddOrUpdate(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{tt.instanceUID})); err != nil {
				t.Fatalf("set instance UID: %v", err)
			}
			if err := ds.AddOrUpdate(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{tt.classUID})); err != nil {
				t.Fatalf("set class UID: %v", err)
			}

			err := NewReferencedSOPFromDataset(ds).Validate()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func stringPointer(value string) *string {
	return &value
}
