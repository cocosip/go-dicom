// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"fmt"
	"slices"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// GraphicType identifies the geometry stored by a SCOORD content item.
type GraphicType string

const (
	// GraphicTypePoint identifies one two-dimensional point.
	GraphicTypePoint GraphicType = "POINT"
	// GraphicTypeMultipoint identifies a set of independent points.
	GraphicTypeMultipoint GraphicType = "MULTIPOINT"
	// GraphicTypePolyline identifies connected line segments.
	GraphicTypePolyline GraphicType = "POLYLINE"
	// GraphicTypeCircle identifies a center and one point on a circle.
	GraphicTypeCircle GraphicType = "CIRCLE"
	// GraphicTypeEllipse identifies the endpoints of an ellipse's axes.
	GraphicTypeEllipse GraphicType = "ELLIPSE"
)

// SpatialCoordinate contains the two-dimensional coordinates of a SCOORD item.
type SpatialCoordinate struct {
	GraphicType GraphicType
	GraphicData []float32
}

// TemporalRangeType identifies how a TCOORD item interprets its references.
type TemporalRangeType string

const (
	// TemporalRangeTypePoint identifies one temporal position.
	TemporalRangeTypePoint TemporalRangeType = "POINT"
	// TemporalRangeTypeMultipoint identifies independent temporal positions.
	TemporalRangeTypeMultipoint TemporalRangeType = "MULTIPOINT"
	// TemporalRangeTypeSegment identifies one temporal interval.
	TemporalRangeTypeSegment TemporalRangeType = "SEGMENT"
	// TemporalRangeTypeMultisegment identifies multiple temporal intervals.
	TemporalRangeTypeMultisegment TemporalRangeType = "MULTISEGMENT"
	// TemporalRangeTypeBegin identifies the beginning of an event.
	TemporalRangeTypeBegin TemporalRangeType = "BEGIN"
	// TemporalRangeTypeEnd identifies the end of an event.
	TemporalRangeTypeEnd TemporalRangeType = "END"
)

// TemporalCoordinate contains exactly one kind of temporal reference.
type TemporalCoordinate struct {
	RangeType       TemporalRangeType
	SamplePositions []uint32
	TimeOffsets     []float64
	DateTimes       []time.Time
}

// NewContentItemSpatialCoordinate creates a SCOORD content item.
func NewContentItemSpatialCoordinate(code *CodeItem, relationship Relationship, value SpatialCoordinate) (*ContentItem, error) {
	if err := validateSpatialCoordinate(value); err != nil {
		return nil, err
	}
	ds, err := newContentItemDataset(code, relationship, ValueTypeSpatialCoordinate)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewString(tag.GraphicType, vr.CS, []string{string(value.GraphicType)})); err != nil {
		return nil, WrapError("set graphic type", err)
	}
	if err := ds.AddOrUpdate(element.NewFloat(tag.GraphicData, slices.Clone(value.GraphicData))); err != nil {
		return nil, WrapError("set graphic data", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// GetSpatialCoordinate returns the value of a SCOORD content item.
func (c *ContentItem) GetSpatialCoordinate() (*SpatialCoordinate, error) {
	if err := c.requireValueType(ValueTypeSpatialCoordinate); err != nil {
		return nil, err
	}
	graphicType, ok := c.dataset.GetString(tag.GraphicType)
	if !ok {
		return nil, NewError("graphic type not found")
	}
	value, ok := c.dataset.Get(tag.GraphicData)
	if !ok {
		return nil, NewError("graphic data not found")
	}
	graphicData, ok := value.(*element.Float)
	if !ok {
		return nil, NewError("graphic data has invalid VR")
	}
	values, err := graphicData.GetValues()
	if err != nil {
		return nil, WrapError("read graphic data", err)
	}
	result := &SpatialCoordinate{GraphicType: GraphicType(graphicType), GraphicData: slices.Clone(values)}
	if err := validateSpatialCoordinate(*result); err != nil {
		return nil, err
	}
	return result, nil
}

func validateSpatialCoordinate(value SpatialCoordinate) error {
	count := len(value.GraphicData)
	switch value.GraphicType {
	case GraphicTypePoint:
		if count != 2 {
			return NewErrorf("POINT graphic data has %d coordinates, want 2", count)
		}
	case GraphicTypeMultipoint, GraphicTypePolyline:
		if count == 0 || count%2 != 0 {
			return NewErrorf("%s graphic data has %d coordinates, want a non-empty even count", value.GraphicType, count)
		}
	case GraphicTypeCircle:
		if count != 4 {
			return NewErrorf("CIRCLE graphic data has %d coordinates, want 4", count)
		}
	case GraphicTypeEllipse:
		if count != 8 {
			return NewErrorf("ELLIPSE graphic data has %d coordinates, want 8", count)
		}
	default:
		return NewErrorf("unknown graphic type %q", value.GraphicType)
	}
	return nil
}

// NewContentItemTemporalCoordinate creates a TCOORD content item.
func NewContentItemTemporalCoordinate(code *CodeItem, relationship Relationship, value TemporalCoordinate) (*ContentItem, error) {
	if err := validateTemporalCoordinate(value); err != nil {
		return nil, err
	}
	ds, err := newContentItemDataset(code, relationship, ValueTypeTemporalCoordinate)
	if err != nil {
		return nil, err
	}
	if err := ds.AddOrUpdate(element.NewString(tag.TemporalRangeType, vr.CS, []string{string(value.RangeType)})); err != nil {
		return nil, WrapError("set temporal range type", err)
	}
	switch {
	case len(value.SamplePositions) > 0:
		err = ds.AddOrUpdate(element.NewUnsignedLong(tag.ReferencedSamplePositions, slices.Clone(value.SamplePositions)))
	case len(value.TimeOffsets) > 0:
		err = ds.AddOrUpdate(element.NewDecimalStringFromFloat(tag.ReferencedTimeOffsets, slices.Clone(value.TimeOffsets)))
	case len(value.DateTimes) > 0:
		err = ds.AddOrUpdate(element.NewDateTimeFromTime(tag.ReferencedDateTime, slices.Clone(value.DateTimes)))
	}
	if err != nil {
		return nil, WrapError("set temporal reference", err)
	}
	return &ContentItem{dataset: ds}, nil
}

// GetTemporalCoordinate returns the value of a TCOORD content item.
func (c *ContentItem) GetTemporalCoordinate() (*TemporalCoordinate, error) {
	if err := c.requireValueType(ValueTypeTemporalCoordinate); err != nil {
		return nil, err
	}
	rangeType, ok := c.dataset.GetString(tag.TemporalRangeType)
	if !ok {
		return nil, NewError("temporal range type not found")
	}
	result := &TemporalCoordinate{RangeType: TemporalRangeType(rangeType)}
	present := 0
	if c.dataset.Contains(tag.ReferencedSamplePositions) {
		values, err := c.dataset.GetUInt32s(tag.ReferencedSamplePositions)
		if err != nil {
			return nil, WrapError("read referenced sample positions", err)
		}
		result.SamplePositions = slices.Clone(values)
		present++
	}
	if c.dataset.Contains(tag.ReferencedTimeOffsets) {
		value, _ := c.dataset.Get(tag.ReferencedTimeOffsets)
		var offsets *element.DecimalString
		switch typed := value.(type) {
		case *element.DecimalString:
			offsets = typed
		case *element.String:
			offsets = element.NewDecimalStringFromBuffer(tag.ReferencedTimeOffsets, typed.Buffer(), nil)
		default:
			return nil, NewError("referenced time offsets have invalid VR")
		}
		values, err := offsets.GetFloats()
		if err != nil {
			return nil, WrapError("read referenced time offsets", err)
		}
		result.TimeOffsets = slices.Clone(values)
		present++
	}
	if c.dataset.Contains(tag.ReferencedDateTime) {
		value, _ := c.dataset.Get(tag.ReferencedDateTime)
		var dateTimes *element.DateTime
		switch typed := value.(type) {
		case *element.DateTime:
			dateTimes = typed
		case *element.String:
			dateTimes = element.NewDateTimeFromBuffer(tag.ReferencedDateTime, typed.Buffer(), nil)
		default:
			return nil, NewError("referenced date times have invalid VR")
		}
		values, err := dateTimes.GetDateTimes()
		if err != nil {
			return nil, WrapError("read referenced date times", err)
		}
		result.DateTimes = slices.Clone(values)
		present++
	}
	if present != 1 {
		return nil, NewErrorf("TCOORD has %d temporal reference types, want 1", present)
	}
	if err := validateTemporalCoordinate(*result); err != nil {
		return nil, err
	}
	return result, nil
}

func validateTemporalCoordinate(value TemporalCoordinate) error {
	referenceKinds := 0
	count := 0
	for _, length := range []int{len(value.SamplePositions), len(value.TimeOffsets), len(value.DateTimes)} {
		if length > 0 {
			referenceKinds++
			count = length
		}
	}
	if referenceKinds != 1 {
		return NewErrorf("TCOORD has %d temporal reference types, want 1", referenceKinds)
	}

	validCount := false
	switch value.RangeType {
	case TemporalRangeTypePoint, TemporalRangeTypeBegin, TemporalRangeTypeEnd:
		validCount = count == 1
	case TemporalRangeTypeMultipoint:
		validCount = count > 0
	case TemporalRangeTypeSegment:
		validCount = count == 2
	case TemporalRangeTypeMultisegment:
		validCount = count > 0 && count%2 == 0
	default:
		return NewErrorf("unknown temporal range type %q", value.RangeType)
	}
	if !validCount {
		return NewError(fmt.Sprintf("%s temporal range has invalid value count %d", value.RangeType, count))
	}
	return nil
}
