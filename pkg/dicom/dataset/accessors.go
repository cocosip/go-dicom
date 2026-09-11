// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"strings"

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

	values, ok := stringValues(elem)
	if !ok {
		return "", false
	}

	return strings.Join(values, "\\"), true
}

// GetStrings retrieves all string values from the dataset.
func (ds *Dataset) GetStrings(t *tag.Tag) ([]string, bool) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, false
	}

	values, ok := stringValues(elem)
	if !ok {
		return nil, false
	}

	return values, true
}

// GetBytes retrieves the element's contiguous encoded value bytes. The
// returned slice follows ByteBuffer.Data ownership semantics and must not be
// modified by the caller.
func (ds *Dataset) GetBytes(t *tag.Tag) ([]byte, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}
	buf := elem.Buffer()
	if buf == nil {
		return nil, fmt.Errorf("element %s does not have a contiguous value buffer", t)
	}
	return buf.Data(), nil
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

// GetInt16s retrieves all int16 values from the dataset.
func (ds *Dataset) GetInt16s(t *tag.Tag) ([]int16, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	ss, ok := elem.(*element.SignedShort)
	if !ok {
		return nil, fmt.Errorf("element %s is not SignedShort", t)
	}

	return ss.GetValues()
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

// GetInt32s retrieves all int32 values from the dataset.
func (ds *Dataset) GetInt32s(t *tag.Tag) ([]int32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	sl, ok := elem.(*element.SignedLong)
	if !ok {
		return nil, fmt.Errorf("element %s is not SignedLong", t)
	}

	return sl.GetValues()
}

// GetInt64 retrieves an int64 value from the dataset.
func (ds *Dataset) GetInt64(t *tag.Tag, index int) (int64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	sv, ok := elem.(*element.SignedVeryLong)
	if !ok {
		return 0, fmt.Errorf("element %s is not SignedVeryLong", t)
	}

	return sv.GetValue(index)
}

// GetInt64s retrieves all int64 values from the dataset.
func (ds *Dataset) GetInt64s(t *tag.Tag) ([]int64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	sv, ok := elem.(*element.SignedVeryLong)
	if !ok {
		return nil, fmt.Errorf("element %s is not SignedVeryLong", t)
	}

	return sv.GetValues()
}

// GetUInt64 retrieves a uint64 value from the dataset.
func (ds *Dataset) GetUInt64(t *tag.Tag, index int) (uint64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return 0, fmt.Errorf("element %s not found", t)
	}

	uv, ok := elem.(*element.UnsignedVeryLong)
	if !ok {
		return 0, fmt.Errorf("element %s is not UnsignedVeryLong", t)
	}

	return uv.GetValue(index)
}

// GetUInt64s retrieves all uint64 values from the dataset.
func (ds *Dataset) GetUInt64s(t *tag.Tag) ([]uint64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	uv, ok := elem.(*element.UnsignedVeryLong)
	if !ok {
		return nil, fmt.Errorf("element %s is not UnsignedVeryLong", t)
	}

	return uv.GetValues()
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

// GetFloat32s retrieves all float32 values from the dataset.
func (ds *Dataset) GetFloat32s(t *tag.Tag) ([]float32, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	fl, ok := elem.(*element.Float)
	if !ok {
		return nil, fmt.Errorf("element %s is not Float", t)
	}

	return fl.GetValues()
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

// GetFloat64s retrieves all float64 values from the dataset.
func (ds *Dataset) GetFloat64s(t *tag.Tag) ([]float64, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	fd, ok := elem.(*element.Double)
	if !ok {
		return nil, fmt.Errorf("element %s is not Double", t)
	}

	return fd.GetValues()
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

func stringValues(elem element.Element) ([]string, bool) {
	values, err := element.CanonicalStrings(elem)
	if err != nil {
		return nil, false
	}
	return values, true
}
