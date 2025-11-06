// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// GetString retrieves a string value from the dataset.
// Returns the string value and true if found, empty string and false otherwise.
func (ds *Dataset) GetString(t *tag.Tag) (string, bool) {
	elem, exists := ds.Get(t)
	if !exists {
		return "", false
	}

	str, ok := elem.(*element.String)
	if !ok {
		return "", false
	}

	return str.GetString(), true
}

// GetStrings retrieves all string values from the dataset.
func (ds *Dataset) GetStrings(t *tag.Tag) ([]string, bool) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, false
	}

	str, ok := elem.(*element.String)
	if !ok {
		return nil, false
	}

	return str.GetValues(), true
}

// GetUInt16 retrieves a uint16 value from the dataset.
func (ds *Dataset) GetUInt16(t *tag.Tag, index int) (uint16, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	us, ok := elem.(*element.UnsignedShort)
	if !ok {
		return 0, fmt.Errorf("element %s is not UnsignedShort", t)
	}

	return us.GetValue(index)
}

// GetUInt16s retrieves all uint16 values from the dataset.
func (ds *Dataset) GetUInt16s(t *tag.Tag) ([]uint16, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	us, ok := elem.(*element.UnsignedShort)
	if !ok {
		return nil, fmt.Errorf("element %s is not UnsignedShort", t)
	}

	return us.GetValues()
}

// GetUInt32 retrieves a uint32 value from the dataset.
func (ds *Dataset) GetUInt32(t *tag.Tag, index int) (uint32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	ul, ok := elem.(*element.UnsignedLong)
	if !ok {
		return 0, fmt.Errorf("element %s is not UnsignedLong", t)
	}

	return ul.GetValue(index)
}

// GetUInt32s retrieves all uint32 values from the dataset.
func (ds *Dataset) GetUInt32s(t *tag.Tag) ([]uint32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	ul, ok := elem.(*element.UnsignedLong)
	if !ok {
		return nil, fmt.Errorf("element %s is not UnsignedLong", t)
	}

	return ul.GetValues()
}

// GetInt16 retrieves an int16 value from the dataset.
func (ds *Dataset) GetInt16(t *tag.Tag, index int) (int16, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	ss, ok := elem.(*element.SignedShort)
	if !ok {
		return 0, fmt.Errorf("element %s is not SignedShort", t)
	}

	return ss.GetValue(index)
}

// GetInt32 retrieves an int32 value from the dataset.
func (ds *Dataset) GetInt32(t *tag.Tag, index int) (int32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	sl, ok := elem.(*element.SignedLong)
	if !ok {
		return 0, fmt.Errorf("element %s is not SignedLong", t)
	}

	return sl.GetValue(index)
}

// GetFloat32 retrieves a float32 value from the dataset.
func (ds *Dataset) GetFloat32(t *tag.Tag, index int) (float32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	fl, ok := elem.(*element.Float)
	if !ok {
		return 0, fmt.Errorf("element %s is not Float", t)
	}

	return fl.GetValue(index)
}

// GetFloat64 retrieves a float64 value from the dataset.
func (ds *Dataset) GetFloat64(t *tag.Tag, index int) (float64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	fd, ok := elem.(*element.Double)
	if !ok {
		return 0, fmt.Errorf("element %s is not Double", t)
	}

	return fd.GetValue(index)
}

// GetSequence retrieves a sequence element from the dataset.
func (ds *Dataset) GetSequence(t *tag.Tag) (*Sequence, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	seq, ok := elem.(*Sequence)
	if !ok {
		return nil, fmt.Errorf("element %s is not Sequence", t)
	}

	return seq, nil
}

// TryGetString tries to retrieve a string value, returning empty string if not found.
func (ds *Dataset) TryGetString(t *tag.Tag) string {
	str, _ := ds.GetString(t)
	return str
}

// TryGetUInt16 tries to retrieve a uint16 value, returning 0 if not found or error.
func (ds *Dataset) TryGetUInt16(t *tag.Tag, index int) uint16 {
	val, _ := ds.GetUInt16(t, index)
	return val
}

// TryGetUInt32 tries to retrieve a uint32 value, returning 0 if not found or error.
func (ds *Dataset) TryGetUInt32(t *tag.Tag, index int) uint32 {
	val, _ := ds.GetUInt32(t, index)
	return val
}
