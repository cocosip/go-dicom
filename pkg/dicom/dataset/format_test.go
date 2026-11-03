// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestFormatSingleLineSummarizesDataset(t *testing.T) {
	ds := newFormatTestDataset(t)

	got := ds.Format(dataset.FormatOptions{})
	want := "(0008,0005) CS [ISO_IR 100] | (0008,1140) SQ (1 item) | (0028,0010) US 512 | (7fe0,0010) OB (4 bytes)"
	if got != want {
		t.Errorf("Format() = %q, want %q", got, want)
	}
}

func TestFormatRowsRendersSequenceItems(t *testing.T) {
	ds := newFormatTestDataset(t)

	got := ds.Format(dataset.FormatOptions{Layout: dataset.LayoutRows, MaxDepth: -1})
	for _, want := range []string{
		"(0008,1140) SQ (1 item)",
		"  Item #1:",
		"  (0008,1150) UI [1.2.3]",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("Format() = %q, want line %q", got, want)
		}
	}
}

func TestFormatDumpAlignsLengthAndKeyword(t *testing.T) {
	ds := newFormatTestDataset(t)

	got := ds.Format(dataset.FormatOptions{Layout: dataset.LayoutDump, MaxDepth: -1})
	for _, want := range []string{
		", SpecificCharacterSet",
		"(0008,1140) SQ (1 item) ReferencedImageSequence",
		"> (0008,1150) UI [1.2.3]",
		", ReferencedSOPClassUID",
		", PixelData",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("Format() = %q, want %q", got, want)
		}
	}
}

func TestFormatRespectsDepthAndElementLimits(t *testing.T) {
	ds := newFormatTestDataset(t)

	shallow := ds.Format(dataset.FormatOptions{Layout: dataset.LayoutRows, MaxDepth: 0})
	if strings.Contains(shallow, "Item #1:") || strings.Contains(shallow, "1.2.3") {
		t.Errorf("depth-limited Format() = %q, want no sequence children", shallow)
	}

	limited := ds.Format(dataset.FormatOptions{Layout: dataset.LayoutRows, MaxDepth: -1, MaxElements: 1})
	if !strings.HasPrefix(limited, "(0008,0005) CS [ISO_IR 100]") || !strings.HasSuffix(limited, "\n...") {
		t.Errorf("element-limited Format() = %q, want first element and truncation marker", limited)
	}
}

func TestFormatElementLimitCountsElementsNotItemLines(t *testing.T) {
	ds := newFormatTestDataset(t)

	// (0008,0005), the sequence, and the nested (0008,1150) are three element
	// rows; the "Item #1:" marker must not consume the budget.
	limited := ds.Format(dataset.FormatOptions{Layout: dataset.LayoutRows, MaxDepth: -1, MaxElements: 3})
	if !strings.Contains(limited, "  (0008,1150) UI [1.2.3]") || !strings.HasSuffix(limited, "\n...") {
		t.Errorf("element-limited Format() = %q, want nested element and truncation marker", limited)
	}
}

func TestFormatHideValuesReturnsTagAndVR(t *testing.T) {
	ds := newFormatTestDataset(t)

	got := ds.Format(dataset.FormatOptions{HideValues: true})
	want := "(0008,0005) CS | (0008,1140) SQ (1 item) | (0028,0010) US | (7fe0,0010) OB"
	if got != want {
		t.Errorf("Format() = %q, want %q", got, want)
	}
}

func TestFormatNilDataset(t *testing.T) {
	var ds *dataset.Dataset
	if got := ds.Format(dataset.FormatOptions{}); got != "" {
		t.Errorf("Format() = %q, want empty string", got)
	}
}

func newFormatTestDataset(t *testing.T) *dataset.Dataset {
	t.Helper()
	ds := dataset.New()
	addFormatTestElement(t, ds, element.NewString(tag.SpecificCharacterSet, vr.CS, []string{"ISO_IR 100"}))
	addFormatTestElement(t, ds, element.NewUnsignedShort(tag.Rows, []uint16{512}))
	addFormatTestElement(t, ds, element.NewOtherByte(tag.PixelData, []byte{1, 2, 3, 4}))

	item := dataset.New()
	addFormatTestElement(t, item, element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{"1.2.3"}))
	addFormatTestElement(t, ds, dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{item}))
	return ds
}

func addFormatTestElement(t *testing.T, ds *dataset.Dataset, elem element.Element) {
	t.Helper()
	if err := ds.Add(elem); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
}
