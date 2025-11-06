// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"strconv"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// DecimalString represents a DICOM element with VR = DS (Decimal String).
// Stores numeric values as strings that can be converted to float64.
type DecimalString struct {
	*String
}

// NewDecimalString creates a new DS element with the given string values.
func NewDecimalString(t *tag.Tag, values []string) *DecimalString {
	str := NewString(t, vr.DS, values)
	return &DecimalString{String: str}
}

// NewDecimalStringFromFloat creates a new DS element from float64 values.
func NewDecimalStringFromFloat(t *tag.Tag, values []float64) *DecimalString {
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return NewDecimalString(t, strs)
}

// NewDecimalStringFromBuffer creates a DS element from an existing buffer.
func NewDecimalStringFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *DecimalString {
	str := NewStringFromBuffer(t, vr.DS, buf, enc)
	return &DecimalString{String: str}
}

// GetFloat parses and returns the float value at the specified index.
func (ds *DecimalString) GetFloat(index int) (float64, error) {
	strVal := ds.GetValue(index)
	if strVal == "" {
		return 0, fmt.Errorf("decimal string at index %d is empty", index)
	}

	f, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse decimal string %q: %w", strVal, err)
	}
	return f, nil
}

// GetFloats parses and returns all float values.
func (ds *DecimalString) GetFloats() ([]float64, error) {
	count := ds.Count()
	if count == 0 {
		return nil, nil
	}

	floats := make([]float64, count)
	for i := 0; i < count; i++ {
		f, err := ds.GetFloat(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float at index %d: %w", i, err)
		}
		floats[i] = f
	}
	return floats, nil
}

// GetFloat32 parses and returns the float32 value at the specified index.
func (ds *DecimalString) GetFloat32(index int) (float32, error) {
	f64, err := ds.GetFloat(index)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}

// GetFloat32s parses and returns all float32 values.
func (ds *DecimalString) GetFloat32s() ([]float32, error) {
	floats64, err := ds.GetFloats()
	if err != nil {
		return nil, err
	}

	floats32 := make([]float32, len(floats64))
	for i, f := range floats64 {
		floats32[i] = float32(f)
	}
	return floats32, nil
}

// IntegerString represents a DICOM element with VR = IS (Integer String).
// Stores integer values as strings that can be converted to int.
type IntegerString struct {
	*String
}

// NewIntegerString creates a new IS element with the given string values.
func NewIntegerString(t *tag.Tag, values []string) *IntegerString {
	str := NewString(t, vr.IS, values)
	return &IntegerString{String: str}
}

// NewIntegerStringFromInt creates a new IS element from int values.
func NewIntegerStringFromInt(t *tag.Tag, values []int) *IntegerString {
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = strconv.Itoa(v)
	}
	return NewIntegerString(t, strs)
}

// NewIntegerStringFromInt32 creates a new IS element from int32 values.
func NewIntegerStringFromInt32(t *tag.Tag, values []int32) *IntegerString {
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = strconv.FormatInt(int64(v), 10)
	}
	return NewIntegerString(t, strs)
}

// NewIntegerStringFromBuffer creates an IS element from an existing buffer.
func NewIntegerStringFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *IntegerString {
	str := NewStringFromBuffer(t, vr.IS, buf, enc)
	return &IntegerString{String: str}
}

// GetInt parses and returns the int value at the specified index.
func (is *IntegerString) GetInt(index int) (int, error) {
	strVal := is.GetValue(index)
	if strVal == "" {
		return 0, fmt.Errorf("integer string at index %d is empty", index)
	}

	i, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer string %q: %w", strVal, err)
	}
	return i, nil
}

// GetInts parses and returns all int values.
func (is *IntegerString) GetInts() ([]int, error) {
	count := is.Count()
	if count == 0 {
		return nil, nil
	}

	ints := make([]int, count)
	for i := 0; i < count; i++ {
		val, err := is.GetInt(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int at index %d: %w", i, err)
		}
		ints[i] = val
	}
	return ints, nil
}

// GetInt32 parses and returns the int32 value at the specified index.
func (is *IntegerString) GetInt32(index int) (int32, error) {
	strVal := is.GetValue(index)
	if strVal == "" {
		return 0, fmt.Errorf("integer string at index %d is empty", index)
	}

	i, err := strconv.ParseInt(strVal, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer string %q: %w", strVal, err)
	}
	return int32(i), nil
}

// GetInt32s parses and returns all int32 values.
func (is *IntegerString) GetInt32s() ([]int32, error) {
	count := is.Count()
	if count == 0 {
		return nil, nil
	}

	ints := make([]int32, count)
	for i := 0; i < count; i++ {
		val, err := is.GetInt32(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int32 at index %d: %w", i, err)
		}
		ints[i] = val
	}
	return ints, nil
}

// GetInt64 parses and returns the int64 value at the specified index.
func (is *IntegerString) GetInt64(index int) (int64, error) {
	strVal := is.GetValue(index)
	if strVal == "" {
		return 0, fmt.Errorf("integer string at index %d is empty", index)
	}

	i, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer string %q: %w", strVal, err)
	}
	return i, nil
}

// GetInt64s parses and returns all int64 values.
func (is *IntegerString) GetInt64s() ([]int64, error) {
	count := is.Count()
	if count == 0 {
		return nil, nil
	}

	ints := make([]int64, count)
	for i := 0; i < count; i++ {
		val, err := is.GetInt64(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int64 at index %d: %w", i, err)
		}
		ints[i] = val
	}
	return ints, nil
}
