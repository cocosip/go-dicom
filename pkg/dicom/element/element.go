// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package element provides DICOM element types and interfaces.
package element

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

var _ Element = (*base)(nil)

// Element is the interface for all DICOM elements.
//
// A DICOM element represents a single data element in a DICOM dataset.
// Each element has a tag, a value representation (VR), and binary data
// stored in a ByteBuffer.
type Element interface {
	// Tag returns the DICOM tag associated with this element.
	Tag() *tag.Tag

	// ValueRepresentation returns the VR of this element.
	ValueRepresentation() *vr.VR

	// Buffer returns the binary data buffer for this element.
	Buffer() buffer.ByteBuffer

	// Length returns the length of the element's data in bytes.
	Length() uint32

	// Count returns the number of values in this element.
	// For example, a multi-valued string element returns the number of strings.
	Count() int

	// String returns a string representation of the element for debugging.
	String() string

	// Validate checks if the element's data is valid according to DICOM rules.
	Validate() error
}

// base provides common functionality for all DICOM elements.
type base struct {
	tag    *tag.Tag
	vr     *vr.VR
	buffer buffer.ByteBuffer
	// byteOrder is used by numeric-like VRs when decoding values from Buffer().
	// Most DICOM datasets are little-endian; parser overrides this for big-endian inputs.
	byteOrder binary.ByteOrder
}

// newBase creates a new base element with the given tag, VR, and buffer.
func newBase(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer) *base {
	return newBaseWithByteOrder(t, v, buf, binary.LittleEndian)
}

// newBaseWithByteOrder creates a new base element with explicit byte order.
func newBaseWithByteOrder(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer, order binary.ByteOrder) *base {
	if buf == nil {
		buf = buffer.Empty
	}
	if order == nil {
		order = binary.LittleEndian
	}
	return &base{
		tag:       t,
		vr:        v,
		buffer:    buf,
		byteOrder: order,
	}
}

// Tag returns the tag associated with this element.
func (e *base) Tag() *tag.Tag {
	return e.tag
}

// ValueRepresentation returns the VR of this element.
func (e *base) ValueRepresentation() *vr.VR {
	return e.vr
}

// Buffer returns the binary data buffer for this element.
func (e *base) Buffer() buffer.ByteBuffer {
	return e.buffer
}

// Length returns the length of the element's data in bytes.
func (e *base) Length() uint32 {
	if e.buffer == nil {
		return 0
	}
	return e.buffer.Size()
}

// String returns a string representation of the element.
func (e *base) String() string {
	return fmt.Sprintf("Tag: %s, VR: %s, Length: %d", e.tag.String(), e.vr.Code(), e.Length())
}

// Validate performs basic validation.
// Subclasses should override this to add type-specific validation.
func (e *base) Validate() error {
	// Check if VR matches the tag's expected VR
	// (This would require dictionary lookup, which we'll add later)
	return nil
}

// Count returns the number of values in this element.
// For example, a multi-valued string element returns the number of strings.
func (e *base) Count() int {
	return 1
}

func (e *base) setByteOrder(order binary.ByteOrder) {
	if order == nil {
		order = binary.LittleEndian
	}
	e.byteOrder = order
}

func (e *base) getByteOrder() binary.ByteOrder {
	if e.byteOrder == nil {
		return binary.LittleEndian
	}
	return e.byteOrder
}

// SetByteOrder sets byte order for an element when decoding numeric data.
// This is primarily used by parser for big-endian transfer syntaxes.
func SetByteOrder(elem Element, order binary.ByteOrder) {
	if setter, ok := elem.(interface{ setByteOrder(binary.ByteOrder) }); ok {
		setter.setByteOrder(order)
	}
}
