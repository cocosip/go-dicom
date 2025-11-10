// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// StructuredReport represents a DICOM Structured Report.
//
// A Structured Report is a hierarchical document consisting of content items
// arranged in a tree structure. The root is typically a CONTAINER item.
//
// Reference: DICOM Part 3, Annex A.35
type StructuredReport struct {
	*ContentItem
}

// NewStructuredReport creates a new structured report with the given root code
func NewStructuredReport(rootCode *CodeItem, items ...*ContentItem) (*StructuredReport, error) {
	// Create root container (relationship type is not needed for root)
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, rootCode); err != nil {
		return nil, err
	}

	// Set value type to CONTAINER (0040,A040) VR=CS
	ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeContainer)}))

	// Set continuity to SEPARATE (typical for SR root) (0040,A050) VR=CS
	ds.AddOrUpdate(element.NewString(tag.ContinuityOfContent, vr.CS, []string{string(ContinuitySeparate)}))

	// Note: RelationshipType is intentionally NOT set for the root element

	// Add content sequence if items provided (0040,A730) VR=SQ
	if len(items) > 0 {
		datasets := make([]*dataset.Dataset, len(items))
		for i, item := range items {
			datasets[i] = item.Dataset()
		}

		seq := dataset.NewSequenceWithItems(tag.ContentSequence, datasets)
		ds.AddOrUpdate(seq)
	}

	contentItem := NewContentItemFromDataset(ds)
	return &StructuredReport{ContentItem: contentItem}, nil
}

// NewStructuredReportFromDataset creates a StructuredReport from an existing dataset
func NewStructuredReportFromDataset(ds *dataset.Dataset) *StructuredReport {
	return &StructuredReport{
		ContentItem: NewContentItemFromDataset(ds),
	}
}

// Add adds a content item to the report
func (sr *StructuredReport) Add(item *ContentItem) error {
	if item == nil {
		return NewError("cannot add nil item")
	}

	// Get existing content sequence or create new one
	seq, err := sr.dataset.GetSequence(tag.ContentSequence)
	if err != nil {
		// Create new sequence with this item
		seq = dataset.NewSequenceWithItems(tag.ContentSequence, []*dataset.Dataset{item.Dataset()})
		sr.dataset.AddOrUpdate(seq)
	} else {
		// Append to existing sequence
		seq.AddItem(item.Dataset())
	}

	return nil
}

// AddText is a convenience method to add a text content item
func (sr *StructuredReport) AddText(code *CodeItem, relationship Relationship, text string) error {
	item, err := NewContentItemText(code, relationship, text)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddCode is a convenience method to add a code content item
func (sr *StructuredReport) AddCode(code *CodeItem, relationship Relationship, value *CodeItem) error {
	item, err := NewContentItemCode(code, relationship, value)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddNumeric is a convenience method to add a numeric content item
func (sr *StructuredReport) AddNumeric(code *CodeItem, relationship Relationship, value *MeasuredValue) error {
	item, err := NewContentItemNumeric(code, relationship, value)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddContainer is a convenience method to add a container content item
func (sr *StructuredReport) AddContainer(code *CodeItem, relationship Relationship, continuity Continuity, items ...*ContentItem) error {
	item, err := NewContentItemContainer(code, relationship, continuity, items...)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// TODO: The following methods require full dicom.File implementation

// Open opens a structured report from a file
// func Open(filename string) (*StructuredReport, error) {
// 	file, err := dicom.ReadFile(filename)
// 	if err != nil {
// 		return nil, WrapError("failed to read DICOM file", err)
// 	}
// 	return NewStructuredReportFromDataset(file.Dataset()), nil
// }

// Save saves the structured report to a file
// func (sr *StructuredReport) Save(filename string) error {
// 	file := dicom.NewFile(sr.dataset)
// 	return file.Save(filename)
// }
