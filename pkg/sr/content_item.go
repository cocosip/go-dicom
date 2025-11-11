// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// ContentItem represents a content item in a DICOM Structured Report.
//
// A content item is a node in the SR tree structure, containing:
//   - Concept Name Code Sequence: What the item represents
//   - Value Type: The type of value (TEXT, NUM, CODE, CONTAINER, etc.)
//   - Relationship Type: How it relates to its parent
//   - Value: The actual content (depends on value type)
//   - Content Sequence: Child items (for CONTAINER type)
//
// Reference: DICOM Part 3, Section C.17.3
type ContentItem struct {
	dataset *dataset.Dataset
}

// NewContentItemText creates a text content item
func NewContentItemText(code *CodeItem, relationship Relationship, text string) (*ContentItem, error) {
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, code); err != nil {
		return nil, err
	}

	// Set value type (0040,A040) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeText)}))

	// Set relationship type (0040,A010) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(relationship)}))

	// Set text value (0040,A160) VR=UT
	_ = ds.AddOrUpdate(element.NewString(tag.TextValue, vr.UT, []string{text}))

	return &ContentItem{dataset: ds}, nil
}

// NewContentItemCode creates a code content item
func NewContentItemCode(code *CodeItem, relationship Relationship, value *CodeItem) (*ContentItem, error) {
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, code); err != nil {
		return nil, err
	}

	// Set value type (0040,A040) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeCode)}))

	// Set relationship type (0040,A010) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(relationship)}))

	// Set concept code sequence (0040,A168) VR=SQ
	if value != nil {
		seq := dataset.NewSequenceWithItems(tag.ConceptCodeSequence, []*dataset.Dataset{value.Dataset()})
		_ = ds.AddOrUpdate(seq)
	}

	return &ContentItem{dataset: ds}, nil
}

// NewContentItemNumeric creates a numeric content item
func NewContentItemNumeric(code *CodeItem, relationship Relationship, value *MeasuredValue) (*ContentItem, error) {
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, code); err != nil {
		return nil, err
	}

	// Set value type (0040,A040) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeNumeric)}))

	// Set relationship type (0040,A010) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(relationship)}))

	// Set measured value sequence (0040,A300) VR=SQ
	if value != nil {
		seq := dataset.NewSequenceWithItems(tag.MeasuredValueSequence, []*dataset.Dataset{value.Dataset()})
		_ = ds.AddOrUpdate(seq)
	}

	return &ContentItem{dataset: ds}, nil
}

// NewContentItemContainer creates a container content item
func NewContentItemContainer(code *CodeItem, relationship Relationship, continuity Continuity, items ...*ContentItem) (*ContentItem, error) {
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, code); err != nil {
		return nil, err
	}

	// Set value type (0040,A040) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeContainer)}))

	// Set relationship type (0040,A010) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(relationship)}))

	// Set continuity (0040,A050) VR=CS
	if continuity != ContinuityNone {
		_ = ds.AddOrUpdate(element.NewString(tag.ContinuityOfContent, vr.CS, []string{string(continuity)}))
	}

	// Add content sequence (0040,A730) VR=SQ
	if len(items) > 0 {
		datasets := make([]*dataset.Dataset, len(items))
		for i, item := range items {
			datasets[i] = item.Dataset()
		}

		seq := dataset.NewSequenceWithItems(tag.ContentSequence, datasets)
		_ = ds.AddOrUpdate(seq)
	}

	return &ContentItem{dataset: ds}, nil
}

// NewContentItemFromDataset creates a ContentItem from an existing dataset
func NewContentItemFromDataset(ds *dataset.Dataset) *ContentItem {
	if ds == nil {
		return nil
	}
	return &ContentItem{dataset: ds}
}

// Dataset returns the underlying dataset
func (c *ContentItem) Dataset() *dataset.Dataset {
	if c == nil {
		return nil
	}
	return c.dataset
}

// Code returns the concept name code
func (c *ContentItem) Code() (*CodeItem, error) {
	if c == nil || c.dataset == nil {
		return nil, NewError("content item is nil")
	}
	return NewCodeItemFromSequence(tag.ConceptNameCodeSequence, c.dataset)
}

// ValueType returns the value type
func (c *ContentItem) ValueType() (ValueType, error) {
	if c == nil || c.dataset == nil {
		return "", NewError("content item is nil")
	}
	s, ok := c.dataset.GetString(tag.ValueType)
	if !ok {
		return "", NewError("value type not found")
	}
	return ValueType(s), nil
}

// Relationship returns the relationship type
func (c *ContentItem) Relationship() (Relationship, error) {
	if c == nil || c.dataset == nil {
		return "", NewError("content item is nil")
	}
	s, ok := c.dataset.GetString(tag.RelationshipType)
	if !ok {
		return "", NewError("relationship type not found")
	}
	return Relationship(s), nil
}

// Continuity returns the continuity of content (for containers)
func (c *ContentItem) Continuity() Continuity {
	if c == nil || c.dataset == nil {
		return ContinuityNone
	}
	s, _ := c.dataset.GetString(tag.ContinuityOfContent)
	return Continuity(s)
}

// Children returns the child content items (for containers)
func (c *ContentItem) Children() ([]*ContentItem, error) {
	if c == nil || c.dataset == nil {
		return nil, nil
	}

	seq, err := c.dataset.GetSequence(tag.ContentSequence)
	if err != nil {
		return nil, nil // No children
	}

	items := make([]*ContentItem, seq.Count())
	for i := 0; i < seq.Count(); i++ {
		items[i] = NewContentItemFromDataset(seq.GetItem(i))
	}

	return items, nil
}

// GetText returns the text value (for TEXT type)
func (c *ContentItem) GetText() (string, error) {
	if c == nil || c.dataset == nil {
		return "", NewError("content item is nil")
	}
	vt, err := c.ValueType()
	if err != nil {
		return "", err
	}
	if vt != ValueTypeText {
		return "", NewErrorf("content item is not TEXT type, got %s", vt)
	}
	s, ok := c.dataset.GetString(tag.TextValue)
	if !ok {
		return "", NewError("text value not found")
	}
	return s, nil
}

// GetNumeric returns the measured value (for NUM type)
func (c *ContentItem) GetNumeric() (*MeasuredValue, error) {
	if c == nil || c.dataset == nil {
		return nil, NewError("content item is nil")
	}
	vt, err := c.ValueType()
	if err != nil {
		return nil, err
	}
	if vt != ValueTypeNumeric {
		return nil, NewErrorf("content item is not NUM type, got %s", vt)
	}
	return NewMeasuredValueFromSequence(tag.MeasuredValueSequence, c.dataset)
}

// GetCode returns the code value (for CODE type)
func (c *ContentItem) GetCode() (*CodeItem, error) {
	if c == nil || c.dataset == nil {
		return nil, NewError("content item is nil")
	}
	vt, err := c.ValueType()
	if err != nil {
		return nil, err
	}
	if vt != ValueTypeCode {
		return nil, NewErrorf("content item is not CODE type, got %s", vt)
	}
	return NewCodeItemFromSequence(tag.ConceptCodeSequence, c.dataset)
}

// Helper function to set concept name code sequence
func setConceptNameCode(ds *dataset.Dataset, code *CodeItem) error {
	if code == nil {
		return ErrMissingCode
	}

	// Set concept name code sequence (0040,A043) VR=SQ
	seq := dataset.NewSequenceWithItems(tag.ConceptNameCodeSequence, []*dataset.Dataset{code.Dataset()})
	_ = ds.AddOrUpdate(seq)

	return nil
}
