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
