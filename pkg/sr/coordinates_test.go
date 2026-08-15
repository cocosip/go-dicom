// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"slices"
	"testing"
	"time"
)

func TestSpatialCoordinateRoundTripCopiesGraphicData(t *testing.T) {
	code := NewCodeItem("111030", "DCM", "Image Region")
	input := SpatialCoordinate{
		GraphicType: GraphicTypeCircle,
		GraphicData: []float32{10, 20, 15, 20},
	}

	item, err := NewContentItemSpatialCoordinate(code, RelationshipContains, input)
	if err != nil {
		t.Fatalf("NewContentItemSpatialCoordinate() error = %v", err)
	}
	input.GraphicData[0] = 999

	got, err := item.GetSpatialCoordinate()
	if err != nil {
		t.Fatalf("GetSpatialCoordinate() error = %v", err)
	}
	if got.GraphicType != GraphicTypeCircle || !slices.Equal(got.GraphicData, []float32{10, 20, 15, 20}) {
		t.Fatalf("GetSpatialCoordinate() = %#v, want CIRCLE [10 20 15 20]", got)
	}
	got.GraphicData[0] = 777
	second, err := item.GetSpatialCoordinate()
	if err != nil {
		t.Fatalf("second GetSpatialCoordinate() error = %v", err)
	}
	if second.GraphicData[0] != 10 {
		t.Fatalf("GetSpatialCoordinate() returned shared data: first value = %v, want 10", second.GraphicData[0])
	}
}

func TestSpatialCoordinateRejectsInvalidCardinality(t *testing.T) {
	code := NewCodeItem("111030", "DCM", "Image Region")
	tests := []SpatialCoordinate{
		{GraphicType: GraphicTypePoint, GraphicData: []float32{1}},
		{GraphicType: GraphicTypeMultipoint, GraphicData: []float32{1, 2, 3}},
		{GraphicType: GraphicTypePolyline, GraphicData: nil},
		{GraphicType: GraphicTypeCircle, GraphicData: []float32{1, 2, 3, 4, 5, 6}},
		{GraphicType: GraphicTypeEllipse, GraphicData: []float32{1, 2, 3, 4}},
		{GraphicType: GraphicType("TRIANGLE"), GraphicData: []float32{1, 2, 3, 4, 5, 6}},
	}

	for _, value := range tests {
		if _, err := NewContentItemSpatialCoordinate(code, RelationshipContains, value); err == nil {
			t.Errorf("NewContentItemSpatialCoordinate(%#v) succeeded, want error", value)
		}
	}
}

func TestTemporalCoordinateRoundTripsEachReferenceKind(t *testing.T) {
	code := NewCodeItem("111034", "DCM", "Temporal Region")
	dates := []time.Time{
		time.Date(2026, time.August, 15, 8, 0, 0, 0, time.UTC),
		time.Date(2026, time.August, 15, 9, 0, 0, 0, time.UTC),
	}
	tests := []struct {
		name  string
		input TemporalCoordinate
	}{
		{name: "sample positions", input: TemporalCoordinate{RangeType: TemporalRangeTypeSegment, SamplePositions: []uint32{3, 9}}},
		{name: "time offsets", input: TemporalCoordinate{RangeType: TemporalRangeTypeMultipoint, TimeOffsets: []float64{0.25, 1.5, 2.75}}},
		{name: "date times", input: TemporalCoordinate{RangeType: TemporalRangeTypeSegment, DateTimes: dates}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := NewContentItemTemporalCoordinate(code, RelationshipSelectedFrom, tt.input)
			if err != nil {
				t.Fatalf("NewContentItemTemporalCoordinate() error = %v", err)
			}
			got, err := item.GetTemporalCoordinate()
			if err != nil {
				t.Fatalf("GetTemporalCoordinate() error = %v", err)
			}
			if got.RangeType != tt.input.RangeType ||
				!slices.Equal(got.SamplePositions, tt.input.SamplePositions) ||
				!slices.Equal(got.TimeOffsets, tt.input.TimeOffsets) ||
				!equalTimes(got.DateTimes, tt.input.DateTimes) {
				t.Fatalf("GetTemporalCoordinate() = %#v, want %#v", got, tt.input)
			}
		})
	}
}

func TestTemporalCoordinateRejectsInvalidReferenceChoiceAndCardinality(t *testing.T) {
	code := NewCodeItem("111034", "DCM", "Temporal Region")
	tests := []TemporalCoordinate{
		{RangeType: TemporalRangeTypePoint},
		{RangeType: TemporalRangeTypePoint, SamplePositions: []uint32{1}, TimeOffsets: []float64{1}},
		{RangeType: TemporalRangeTypePoint, SamplePositions: []uint32{1, 2}},
		{RangeType: TemporalRangeTypeSegment, TimeOffsets: []float64{1}},
		{RangeType: TemporalRangeTypeMultisegment, SamplePositions: []uint32{1, 2, 3}},
		{RangeType: TemporalRangeType("INTERVAL"), SamplePositions: []uint32{1}},
	}

	for _, value := range tests {
		if _, err := NewContentItemTemporalCoordinate(code, RelationshipContains, value); err == nil {
			t.Errorf("NewContentItemTemporalCoordinate(%#v) succeeded, want error", value)
		}
	}
}

func equalTimes(a, b []time.Time) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}
