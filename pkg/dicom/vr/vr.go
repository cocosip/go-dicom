// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package vr implements DICOM Value Representation (VR) types and validation.
package vr

import (
	"fmt"
	"reflect"
)

// VR represents a DICOM Value Representation (VR).
// A VR defines the data type and encoding of a DICOM element value.
type VR struct {
	// code is the two-character VR code (e.g., "PN", "LO", "US")
	code string

	// name is the descriptive name of the VR
	name string

	// isString indicates if the value is a string type
	isString bool

	// isStringEncoded indicates if the string uses the specific character set
	isStringEncoded bool

	// is16bitLength indicates if the length field is 16-bit (true) or 32-bit (false)
	is16bitLength bool

	// isMultiValue indicates if the value can contain multiple items
	isMultiValue bool

	// paddingValue is the byte used to pad values to even length
	paddingValue byte

	// maximumLength is the maximum length of a single value (0 = unlimited)
	maximumLength uint32

	// unitSize is the size of each individual value unit for fixed length types
	unitSize int

	// byteSwap is the number of bytes to swap when changing endianness
	byteSwap int

	// valueType is the Go type used to represent the VR value
	valueType reflect.Type

	// stringValidator is an optional validation function for string values
	stringValidator func(string) error
}

// Code returns the two-character VR code.
func (v *VR) Code() string {
	return v.code
}

// Name returns the descriptive name of the VR.
func (v *VR) Name() string {
	return v.name
}

// IsString returns true if the value is a string type.
func (v *VR) IsString() bool {
	return v.isString
}

// IsStringEncoded returns true if the string uses a specific character set.
func (v *VR) IsStringEncoded() bool {
	return v.isStringEncoded
}

// Is16bitLength returns true if the length field is 16-bit.
func (v *VR) Is16bitLength() bool {
	return v.is16bitLength
}

// IsMultiValue returns true if the value can contain multiple items.
func (v *VR) IsMultiValue() bool {
	return v.isMultiValue
}

// PaddingValue returns the byte used to pad values to even length.
func (v *VR) PaddingValue() byte {
	return v.paddingValue
}

// MaximumLength returns the maximum length of a single value.
// Returns 0 for unlimited length.
func (v *VR) MaximumLength() uint32 {
	return v.maximumLength
}

// UnitSize returns the size of each individual value unit.
func (v *VR) UnitSize() int {
	return v.unitSize
}

// ByteSwap returns the number of bytes to swap when changing endianness.
func (v *VR) ByteSwap() int {
	return v.byteSwap
}

// ValueType returns the Go type used to represent the VR value.
func (v *VR) ValueType() reflect.Type {
	return v.valueType
}

// ValidateString validates a string value according to VR rules.
// Returns nil if valid, or an error describing the validation failure.
func (v *VR) ValidateString(value string) error {
	if !PerformValidation {
		return nil
	}
	if v.stringValidator != nil {
		return v.stringValidator(value)
	}
	return nil
}

// String returns the VR code as a string.
func (v *VR) String() string {
	return v.code
}

// PerformValidation controls whether VR validation is performed.
// Set to false to disable validation (e.g., for performance or when reading non-compliant files).
var PerformValidation = true

// Padding constants
const (
	padZero  byte = 0x00
	padSpace byte = 0x20
)

// Parse parses a VR from its string code.
// Returns an error if the VR code is unknown.
func Parse(code string) (*VR, error) {
	vr, ok := registry[code]
	if !ok {
		return nil, fmt.Errorf("unknown VR code: '%s'", code)
	}
	return vr, nil
}

// ParseBytes parses a VR from its two-byte representation.
// Returns an error if the VR code is unknown or the byte slice is invalid.
func ParseBytes(b []byte) (*VR, error) {
	if len(b) < 2 {
		return nil, fmt.Errorf("VR byte slice must be at least 2 bytes, got %d", len(b))
	}

	// Convert bytes to string and lookup in registry
	code := string(b[0:2])
	if vr, ok := registry[code]; ok {
		return vr, nil
	}

	return nil, fmt.Errorf("unknown VR code: '%s'", code)
}

// TryParse attempts to parse a VR from its string code.
// Returns the VR and true if successful, or nil and false if unknown.
func TryParse(code string) (*VR, bool) {
	vr, ok := registry[code]
	return vr, ok
}
