// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// MeasuredValue represents a numeric measurement with units in DICOM SR.
//
// A measured value consists of:
//   - Numeric Value: The actual measurement
//   - Measurement Units Code Sequence: Coded units (e.g., "mm", "kg", "mg/dL")
//
// Reference: DICOM Part 3, Section C.18.1
type MeasuredValue struct {
	dataset *dataset.Dataset
}

// NewMeasuredValue creates a new measured value
func NewMeasuredValue(value float64, units *CodeItem) *MeasuredValue {
	ds := dataset.New()

	// Add numeric value (0040,A30A) VR=DS
	// Use the embedded String from DecimalString to satisfy Element interface
	dsElem := element.NewDecimalStringFromFloat(tag.NumericValue, []float64{value})
	ds.AddOrUpdate(dsElem.String)

	// Add measurement units code sequence (0040,08EA) VR=SQ
	if units != nil {
		seq := dataset.NewSequenceWithItems(tag.MeasurementUnitsCodeSequence, []*dataset.Dataset{units.Dataset()})
		ds.AddOrUpdate(seq)
	}

	return &MeasuredValue{dataset: ds}
}

// NewMeasuredValueFromDataset creates a MeasuredValue from an existing dataset
func NewMeasuredValueFromDataset(ds *dataset.Dataset) *MeasuredValue {
	if ds == nil {
		return nil
	}
	return &MeasuredValue{dataset: ds}
}

// NewMeasuredValueFromSequence creates a MeasuredValue from the first item in a sequence
func NewMeasuredValueFromSequence(t *tag.Tag, ds *dataset.Dataset) (*MeasuredValue, error) {
	seq, err := ds.GetSequence(t)
	if err != nil {
		return nil, WrapError("sequence not found", err)
	}

	if seq.Count() == 0 {
		return nil, ErrEmptySequence
	}

	return NewMeasuredValueFromDataset(seq.GetItem(0)), nil
}

// Value returns the numeric value
func (m *MeasuredValue) Value() float64 {
	if m == nil || m.dataset == nil {
		return 0.0
	}

	// Get as string element (DS VR stores as string)
	elem, exists := m.dataset.Get(tag.NumericValue)
	if !exists {
		return 0.0
	}

	// Cast to String element and parse as decimal
	if str, ok := elem.(*element.String); ok {
		// Create DecimalString wrapper to parse
		ds := &element.DecimalString{String: str}
		if val, err := ds.GetFloat(0); err == nil {
			return val
		}
	}

	return 0.0
}

// Units returns the measurement units code item
func (m *MeasuredValue) Units() (*CodeItem, error) {
	if m == nil || m.dataset == nil {
		return nil, NewError("measured value is nil")
	}
	return NewCodeItemFromSequence(tag.MeasurementUnitsCodeSequence, m.dataset)
}

// Dataset returns the underlying dataset
func (m *MeasuredValue) Dataset() *dataset.Dataset {
	if m == nil {
		return nil
	}
	return m.dataset
}

// String returns a string representation of the measured value
func (m *MeasuredValue) String() string {
	units, err := m.Units()
	if err != nil {
		return fmt.Sprintf("%f (unknown units)", m.Value())
	}
	return fmt.Sprintf("%f %s", m.Value(), units.Value())
}
