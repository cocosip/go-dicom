// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"time"

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

func newContentItemDataset(code *CodeItem, relationship Relationship, valueType ValueType) (*dataset.Dataset, error) {
	if code == nil || code.Dataset() == nil {
		return nil, ErrMissingCode
	}
	ds := dataset.New()
	if err := setConceptNameCode(ds, code); err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(valueType)})); err != nil {
		return nil, WrapError("set value type", err)
	}
	if err := ds.AddOrUpdate(element.NewString(tag.RelationshipType, vr.CS, []string{string(relationship)})); err != nil {
		return nil, WrapError("set relationship type", err)
	}
	return ds, nil
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
		datasets, err := contentItemDatasets(items)
		if err != nil {
			return nil, err
		}

		seq := dataset.NewSequenceWithItems(tag.ContentSequence, datasets)
		_ = ds.AddOrUpdate(seq)
	}

	return &ContentItem{dataset: ds}, nil
}

// NewContentItemPersonName creates a person-name content item.
func NewContentItemPersonName(code *CodeItem, relationship Relationship, value string) (*ContentItem, error) {
	ds, err := newContentItemDataset(code, relationship, ValueTypePersonName)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewPersonName(tag.PersonName, []string{value})); err != nil {
		return nil, WrapError("set person name", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// NewContentItemDate creates a date content item.
func NewContentItemDate(code *CodeItem, relationship Relationship, value time.Time) (*ContentItem, error) {
	ds, err := newContentItemDataset(code, relationship, ValueTypeDate)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewDateFromTime(tag.Date, []time.Time{value})); err != nil {
		return nil, WrapError("set date", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// NewContentItemTime creates a time content item.
func NewContentItemTime(code *CodeItem, relationship Relationship, value time.Time) (*ContentItem, error) {
	ds, err := newContentItemDataset(code, relationship, ValueTypeTime)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewTimeFromTime(tag.Time, []time.Time{value})); err != nil {
		return nil, WrapError("set time", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// NewContentItemDateTime creates a date-time content item.
func NewContentItemDateTime(code *CodeItem, relationship Relationship, value time.Time) (*ContentItem, error) {
	ds, err := newContentItemDataset(code, relationship, ValueTypeDateTime)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewDateTimeFromTime(tag.DateTime, []time.Time{value})); err != nil {
		return nil, WrapError("set date time", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// NewContentItemUIDReference creates a UID-reference content item.
func NewContentItemUIDReference(code *CodeItem, relationship Relationship, value string) (*ContentItem, error) {
	ds, err := newContentItemDataset(code, relationship, ValueTypeUIDReference)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewString(tag.UID, vr.UI, []string{value})); err != nil {
		return nil, WrapError("set UID reference", err)
	}
	return &ContentItem{dataset: ds}, nil
}

func newContentItemReferencedSOP(code *CodeItem, relationship Relationship, valueType ValueType, value *ReferencedSOP) (*ContentItem, error) {
	if value == nil || value.Dataset() == nil {
		return nil, NewError("referenced SOP is nil")
	}
	ds, err := newContentItemDataset(code, relationship, valueType)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(dataset.NewSequenceWithItems(tag.ReferencedSOPSequence, []*dataset.Dataset{value.Dataset()})); err != nil {
		return nil, WrapError("set referenced SOP", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// NewContentItemComposite creates a composite-reference content item.
func NewContentItemComposite(code *CodeItem, relationship Relationship, value *ReferencedSOP) (*ContentItem, error) {
	return newContentItemReferencedSOP(code, relationship, ValueTypeComposite, value)
}

// NewContentItemImage creates an image-reference content item.
func NewContentItemImage(code *CodeItem, relationship Relationship, value *ReferencedSOP) (*ContentItem, error) {
	return newContentItemReferencedSOP(code, relationship, ValueTypeImage, value)
}

// NewContentItemWaveform creates a waveform-reference content item.
func NewContentItemWaveform(code *CodeItem, relationship Relationship, value *ReferencedSOP) (*ContentItem, error) {
	return newContentItemReferencedSOP(code, relationship, ValueTypeWaveform, value)
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
		if !c.dataset.Contains(tag.ContentSequence) {
			return nil, nil
		}
		return nil, err
	}

	items := make([]*ContentItem, seq.Count())
	for i := 0; i < seq.Count(); i++ {
		items[i] = NewContentItemFromDataset(seq.GetItem(i))
	}

	return items, nil
}

func contentItemDatasets(items []*ContentItem) ([]*dataset.Dataset, error) {
	datasets := make([]*dataset.Dataset, len(items))
	for i, item := range items {
		if item == nil {
			return nil, NewError("cannot add nil item")
		}
		if item.Dataset() == nil {
			return nil, NewError("content item dataset is nil")
		}
		datasets[i] = item.Dataset()
	}
	return datasets, nil
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

// GetPersonName returns the person name for a PNAME content item.
func (c *ContentItem) GetPersonName() (string, error) {
	if err := c.requireValueType(ValueTypePersonName); err != nil {
		return "", err
	}
	value, ok := c.dataset.Get(tag.PersonName)
	if !ok {
		return "", NewError("person name not found")
	}
	var personName *element.PersonName
	switch typed := value.(type) {
	case *element.PersonName:
		personName = typed
	case *element.String:
		personName = element.NewPersonNameFromBuffer(tag.PersonName, typed.Buffer(), nil)
	default:
		return "", NewError("person name has invalid VR")
	}
	return personName.GetValue(0), nil
}

// GetDate returns the date for a DATE content item.
func (c *ContentItem) GetDate() (time.Time, error) {
	if err := c.requireValueType(ValueTypeDate); err != nil {
		return time.Time{}, err
	}
	value, ok := c.dataset.Get(tag.Date)
	if !ok {
		return time.Time{}, NewError("date not found")
	}
	var date *element.Date
	switch typed := value.(type) {
	case *element.Date:
		date = typed
	case *element.String:
		date = element.NewDateFromBuffer(tag.Date, typed.Buffer(), nil)
	default:
		return time.Time{}, NewError("date has invalid VR")
	}
	return date.GetDate(0)
}

// GetTime returns the time for a TIME content item.
func (c *ContentItem) GetTime() (time.Time, error) {
	if err := c.requireValueType(ValueTypeTime); err != nil {
		return time.Time{}, err
	}
	value, ok := c.dataset.Get(tag.Time)
	if !ok {
		return time.Time{}, NewError("time not found")
	}
	var timeValue *element.Time
	switch typed := value.(type) {
	case *element.Time:
		timeValue = typed
	case *element.String:
		timeValue = element.NewTimeFromBuffer(tag.Time, typed.Buffer(), nil)
	default:
		return time.Time{}, NewError("time has invalid VR")
	}
	return timeValue.GetTime(0)
}

// GetDateTime returns the date-time for a DATETIME content item.
func (c *ContentItem) GetDateTime() (time.Time, error) {
	if err := c.requireValueType(ValueTypeDateTime); err != nil {
		return time.Time{}, err
	}
	value, ok := c.dataset.Get(tag.DateTime)
	if !ok {
		return time.Time{}, NewError("date time not found")
	}
	var dateTime *element.DateTime
	switch typed := value.(type) {
	case *element.DateTime:
		dateTime = typed
	case *element.String:
		dateTime = element.NewDateTimeFromBuffer(tag.DateTime, typed.Buffer(), nil)
	default:
		return time.Time{}, NewError("date time has invalid VR")
	}
	return dateTime.GetDateTime(0)
}

// GetUIDReference returns the UID for a UIDREF content item.
func (c *ContentItem) GetUIDReference() (string, error) {
	if err := c.requireValueType(ValueTypeUIDReference); err != nil {
		return "", err
	}
	value, ok := c.dataset.GetString(tag.UID)
	if !ok {
		return "", NewError("UID reference not found")
	}
	return value, nil
}

// GetReferencedSOP returns the reference for COMPOSITE, IMAGE, or WAVEFORM.
func (c *ContentItem) GetReferencedSOP() (*ReferencedSOP, error) {
	valueType, err := c.ValueType()
	if err != nil {
		return nil, err
	}
	if valueType != ValueTypeComposite && valueType != ValueTypeImage && valueType != ValueTypeWaveform {
		return nil, NewErrorf("content item is not a referenced SOP type, got %s", valueType)
	}
	return NewReferencedSOPFromSequence(tag.ReferencedSOPSequence, c.dataset)
}

func (c *ContentItem) requireValueType(expected ValueType) error {
	if c == nil || c.dataset == nil {
		return NewError("content item is nil")
	}
	actual, err := c.ValueType()
	if err != nil {
		return err
	}
	if actual != expected {
		return NewErrorf("content item is not %s type, got %s", expected, actual)
	}
	return nil
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
