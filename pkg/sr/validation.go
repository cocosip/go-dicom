// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

type contentValueTag struct {
	tag  *tag.Tag
	name string
}

var contentValueTags = []contentValueTag{
	{tag.TextValue, "Text Value"},
	{tag.ConceptCodeSequence, "Concept Code Sequence"},
	{tag.MeasuredValueSequence, "Measured Value Sequence"},
	{tag.PersonName, "Person Name"},
	{tag.Date, "Date"},
	{tag.Time, "Time"},
	{tag.DateTime, "Date Time"},
	{tag.UID, "UID"},
	{tag.ReferencedSOPSequence, "Referenced SOP Sequence"},
	{tag.GraphicType, "Graphic Type"},
	{tag.GraphicData, "Graphic Data"},
	{tag.TemporalRangeType, "Temporal Range Type"},
	{tag.ReferencedSamplePositions, "Referenced Sample Positions"},
	{tag.ReferencedTimeOffsets, "Referenced Time Offsets"},
	{tag.ReferencedDateTime, "Referenced Date Time"},
}

var allowedValueTags = map[ValueType]map[uint32]bool{
	ValueTypeContainer:         {},
	ValueTypeText:              tagSet(tag.TextValue),
	ValueTypeCode:              tagSet(tag.ConceptCodeSequence),
	ValueTypeNumeric:           tagSet(tag.MeasuredValueSequence),
	ValueTypePersonName:        tagSet(tag.PersonName),
	ValueTypeDate:              tagSet(tag.Date),
	ValueTypeTime:              tagSet(tag.Time),
	ValueTypeDateTime:          tagSet(tag.DateTime),
	ValueTypeUIDReference:      tagSet(tag.UID),
	ValueTypeComposite:         tagSet(tag.ReferencedSOPSequence),
	ValueTypeImage:             tagSet(tag.ReferencedSOPSequence),
	ValueTypeWaveform:          tagSet(tag.ReferencedSOPSequence),
	ValueTypeSpatialCoordinate: tagSet(tag.GraphicType, tag.GraphicData),
	ValueTypeTemporalCoordinate: tagSet(
		tag.TemporalRangeType,
		tag.ReferencedSamplePositions,
		tag.ReferencedTimeOffsets,
		tag.ReferencedDateTime,
	),
}

func tagSet(tags ...*tag.Tag) map[uint32]bool {
	result := make(map[uint32]bool, len(tags))
	for _, value := range tags {
		result[value.ToUint32()] = true
	}
	return result
}

// Validate verifies that a coded entry is structurally and semantically valid.
func (c *CodeItem) Validate() error {
	if c == nil || c.dataset == nil {
		return NewError("code item is nil")
	}
	if c.buildErr != nil {
		return c.buildErr
	}
	valueTags := []*tag.Tag{tag.CodeValue, tag.LongCodeValue, tag.URNCodeValue}
	present := 0
	var selected *tag.Tag
	for _, valueTag := range valueTags {
		if c.dataset.Contains(valueTag) {
			present++
			selected = valueTag
		}
	}
	if present != 1 {
		return NewErrorf("code item has %d code value fields, want 1", present)
	}
	if c.Value() == "" {
		return NewError("code value is empty")
	}
	if selected != tag.URNCodeValue && c.Scheme() == "" {
		return NewError("coding scheme designator not found")
	}
	if c.Meaning() == "" {
		return NewError("code meaning not found")
	}
	if err := c.dataset.Validate(); err != nil {
		return WrapError("validate code item dataset", err)
	}
	return nil
}

// Validate verifies that a measured value contains a number and one valid unit.
func (m *MeasuredValue) Validate() error {
	if m == nil || m.dataset == nil {
		return NewError("measured value is nil")
	}
	numeric, ok := m.dataset.Get(tag.NumericValue)
	if !ok {
		return NewError("numeric value not found")
	}
	if _, ok := numeric.(*element.DecimalString); !ok {
		if _, ok := numeric.(*element.String); !ok {
			return NewError("numeric value has invalid VR")
		}
	}
	units, err := requiredSingleSequence(tag.MeasurementUnitsCodeSequence, "measurement units code", m.dataset)
	if err != nil {
		return err
	}
	if err := NewCodeItemFromDataset(units).Validate(); err != nil {
		return WrapError("validate measurement units code", err)
	}
	if err := m.dataset.Validate(); err != nil {
		return WrapError("validate measured value dataset", err)
	}
	return nil
}

// Validate verifies both identifiers in a referenced SOP entry.
func (r *ReferencedSOP) Validate() error {
	if r == nil || r.dataset == nil {
		return NewError("referenced SOP is nil")
	}
	if !uid.IsValid(r.InstanceUID()) {
		return NewErrorf("invalid Referenced SOP Instance UID %q", r.InstanceUID())
	}
	if !uid.IsValid(r.ClassUID()) {
		return NewErrorf("invalid Referenced SOP Class UID %q", r.ClassUID())
	}
	if err := r.dataset.Validate(); err != nil {
		return WrapError("validate referenced SOP dataset", err)
	}
	return nil
}

// Validate verifies a non-root content item and all of its descendants.
func (c *ContentItem) Validate() error {
	return validateContentItem(c, false, "")
}

// Validate verifies the SR root and recursively validates its content tree.
func (sr *StructuredReport) Validate() error {
	if sr == nil || sr.ContentItem == nil || sr.Dataset() == nil {
		return NewError("structured report is nil")
	}
	valueType, err := sr.ValueType()
	if err != nil {
		return err
	}
	if valueType != ValueTypeContainer {
		return NewErrorf("root must have Value Type CONTAINER, got %s", valueType)
	}
	if sr.Dataset().Contains(tag.RelationshipType) {
		return NewError("root must not have Relationship Type")
	}
	return validateContentItem(sr.ContentItem, true, "")
}

func validateContentItem(item *ContentItem, root bool, path string) error {
	if item == nil || item.dataset == nil {
		return validationPathError(path, NewError("content item is nil"))
	}
	if item.dataset.Contains(tag.ReferencedContentItemIdentifier) {
		return validateByReferenceContentItem(item, root, path)
	}
	valueType, err := item.ValueType()
	if err != nil {
		return validationPathError(path, err)
	}
	allowed, recognized := allowedValueTags[valueType]
	if !recognized {
		return validationPathError(path, NewErrorf("unknown value type %q", valueType))
	}
	if !root {
		relationship, err := item.Relationship()
		if err != nil {
			return validationPathError(path, err)
		}
		if !validRelationship(relationship) {
			return validationPathError(path, NewErrorf("unknown relationship type %q", relationship))
		}
	}
	for _, candidate := range contentValueTags {
		if item.dataset.Contains(candidate.tag) && !allowed[candidate.tag.ToUint32()] {
			return validationPathError(path, NewErrorf("%s contradicts Value Type %s", candidate.name, valueType))
		}
	}
	if err := validateConceptName(item.dataset); err != nil {
		return validationPathError(path, err)
	}
	if err := validateContentValue(item, valueType); err != nil {
		return validationPathError(path, err)
	}
	if err := item.dataset.Validate(); err != nil {
		return validationPathError(path, WrapError("validate content item dataset", err))
	}
	children, err := item.Children()
	if err != nil {
		return validationPathError(path, WrapError("read Content Sequence", err))
	}
	for index, child := range children {
		childPath := fmt.Sprintf("ContentSequence[%d]", index)
		if path != "" {
			childPath = path + "." + childPath
		}
		if err := validateContentItem(child, false, childPath); err != nil {
			return err
		}
	}
	return nil
}

func validateByReferenceContentItem(item *ContentItem, root bool, path string) error {
	if root {
		return validationPathError(path, NewError("root cannot be a by-reference content item"))
	}
	relationship, err := item.Relationship()
	if err != nil {
		return validationPathError(path, err)
	}
	if !validRelationship(relationship) {
		return validationPathError(path, NewErrorf("unknown relationship type %q", relationship))
	}
	if item.dataset.Contains(tag.ValueType) {
		return validationPathError(path, NewError("by-reference content item must not have Value Type"))
	}
	if item.dataset.Contains(tag.ConceptNameCodeSequence) {
		return validationPathError(path, NewError("by-reference content item must not have Concept Name Code Sequence"))
	}
	if item.dataset.Contains(tag.ContentSequence) {
		return validationPathError(path, NewError("by-reference content item must not have Content Sequence"))
	}
	for _, candidate := range contentValueTags {
		if item.dataset.Contains(candidate.tag) {
			return validationPathError(path, NewErrorf("by-reference content item must not have %s", candidate.name))
		}
	}
	identifier, err := item.dataset.GetUInt32s(tag.ReferencedContentItemIdentifier)
	if err != nil {
		return validationPathError(path, WrapError("read Referenced Content Item Identifier", err))
	}
	if len(identifier) == 0 {
		return validationPathError(path, NewError("Referenced Content Item Identifier is empty"))
	}
	for _, component := range identifier {
		if component == 0 {
			return validationPathError(path, NewError("Referenced Content Item Identifier contains zero"))
		}
	}
	if err := item.dataset.Validate(); err != nil {
		return validationPathError(path, WrapError("validate by-reference content item dataset", err))
	}
	return nil
}

func validateConceptName(ds *dataset.Dataset) error {
	if !ds.Contains(tag.ConceptNameCodeSequence) {
		return nil
	}
	item, err := requiredSingleSequence(tag.ConceptNameCodeSequence, "concept name code", ds)
	if err != nil {
		return err
	}
	if err := NewCodeItemFromDataset(item).Validate(); err != nil {
		return WrapError("validate concept name code", err)
	}
	return nil
}

func validateContentValue(item *ContentItem, valueType ValueType) error {
	switch valueType {
	case ValueTypeContainer:
		continuity := item.Continuity()
		if continuity != ContinuityNone && continuity != ContinuitySeparate && continuity != ContinuityContinuous {
			return NewErrorf("unknown continuity of content %q", continuity)
		}
	case ValueTypeText:
		_, err := item.GetText()
		return err
	case ValueTypeCode:
		value, err := requiredSingleSequence(tag.ConceptCodeSequence, "concept code", item.dataset)
		if err != nil {
			return err
		}
		return NewCodeItemFromDataset(value).Validate()
	case ValueTypeNumeric:
		value, err := requiredSingleSequence(tag.MeasuredValueSequence, "measured value", item.dataset)
		if err != nil {
			return err
		}
		return NewMeasuredValueFromDataset(value).Validate()
	case ValueTypePersonName:
		_, err := item.GetPersonName()
		return err
	case ValueTypeDate:
		_, err := item.GetDate()
		return err
	case ValueTypeTime:
		_, err := item.GetTime()
		return err
	case ValueTypeDateTime:
		_, err := item.GetDateTime()
		return err
	case ValueTypeUIDReference:
		value, err := item.GetUIDReference()
		if err != nil {
			return err
		}
		if !uid.IsValid(value) {
			return NewErrorf("invalid UID reference %q", value)
		}
	case ValueTypeComposite, ValueTypeImage, ValueTypeWaveform:
		value, err := requiredSingleSequence(tag.ReferencedSOPSequence, "referenced SOP", item.dataset)
		if err != nil {
			return err
		}
		return NewReferencedSOPFromDataset(value).Validate()
	case ValueTypeSpatialCoordinate:
		_, err := item.GetSpatialCoordinate()
		return err
	case ValueTypeTemporalCoordinate:
		_, err := item.GetTemporalCoordinate()
		return err
	}
	return nil
}

func requiredSingleSequence(sequenceTag *tag.Tag, name string, ds *dataset.Dataset) (*dataset.Dataset, error) {
	if !ds.Contains(sequenceTag) {
		return nil, NewError(name + " not found")
	}
	sequence, err := ds.GetSequence(sequenceTag)
	if err != nil {
		return nil, WrapError(name+" has invalid VR", err)
	}
	if sequence.Count() != 1 {
		return nil, NewErrorf("%s sequence has %d items, want 1", name, sequence.Count())
	}
	item := sequence.GetItem(0)
	if item == nil {
		return nil, NewError(name + " sequence item is nil")
	}
	return item, nil
}

func validRelationship(value Relationship) bool {
	switch value {
	case RelationshipContains,
		RelationshipHasProperties,
		RelationshipInferredFrom,
		RelationshipSelectedFrom,
		RelationshipHasObservationContext,
		RelationshipHasAcquisitionContext,
		RelationshipHasConceptModifier:
		return true
	default:
		return false
	}
}

func validationPathError(path string, err error) error {
	if err == nil || strings.TrimSpace(path) == "" {
		return err
	}
	return WrapError(path, err)
}
