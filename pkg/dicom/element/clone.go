// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// DeepClone returns an independent copy of an element. All built-in element
// types retain their concrete type, decoding state, and value bytes.
func DeepClone(elem Element) Element {
	if elem == nil {
		return nil
	}

	switch value := elem.(type) {
	case *String:
		return cloneString(value)
	case *DecimalString:
		return &DecimalString{str: cloneString(value.str)}
	case *IntegerString:
		return &IntegerString{str: cloneString(value.str)}
	case *Date:
		return &Date{str: cloneString(value.str)}
	case *Time:
		return &Time{str: cloneString(value.str)}
	case *DateTime:
		return &DateTime{str: cloneString(value.str)}
	case *PersonName:
		return &PersonName{str: cloneString(value.str)}
	case *AttributeTag:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *UnsignedShort:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *UnsignedLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *SignedShort:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *SignedLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *Float:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *Double:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *SignedVeryLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *UnsignedVeryLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherByte:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherWord:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *Unknown:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherDouble:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherFloat:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherVeryLong:
		clone := *value
		clone.base = cloneBase(value.base)
		return &clone
	case *OtherByteFragment:
		return &OtherByteFragment{FragmentSequence: cloneFragmentSequence(value.FragmentSequence)}
	case *OtherWordFragment:
		return &OtherWordFragment{FragmentSequence: cloneFragmentSequence(value.FragmentSequence)}
	case *FragmentSequence:
		return cloneFragmentSequence(value)
	case *base:
		return cloneBase(value)
	default:
		return &snapshotElement{
			base:  newBaseWithByteOrder(elem.Tag().Clone(), elem.ValueRepresentation(), cloneBuffer(elem.Buffer()), nil),
			count: elem.Count(),
		}
	}
}

// DeepCloneChecked returns an independent copy of an element and reports any
// buffer I/O error encountered while detaching its value bytes.
func DeepCloneChecked(elem Element) (Element, error) {
	if elem == nil {
		return nil, nil
	}

	var clone Element
	var err error
	switch value := elem.(type) {
	case *String:
		clone, err = cloneStringChecked(value)
	case *DecimalString:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &DecimalString{str: str}
	case *IntegerString:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &IntegerString{str: str}
	case *Date:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &Date{str: str}
	case *Time:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &Time{str: str}
	case *DateTime:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &DateTime{str: str}
	case *PersonName:
		var str *String
		str, err = cloneStringChecked(value.str)
		clone = &PersonName{str: str}
	case *AttributeTag:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *UnsignedShort:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *UnsignedLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *SignedShort:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *SignedLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *Float:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *Double:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *SignedVeryLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *UnsignedVeryLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherByte:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherWord:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *Unknown:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherDouble:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherFloat:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherVeryLong:
		clonedValue := *value
		clonedValue.base, err = cloneBaseChecked(value.base)
		clone = &clonedValue
	case *OtherByteFragment:
		var fragments *FragmentSequence
		fragments, err = cloneFragmentSequenceChecked(value.FragmentSequence)
		clone = &OtherByteFragment{FragmentSequence: fragments}
	case *OtherWordFragment:
		var fragments *FragmentSequence
		fragments, err = cloneFragmentSequenceChecked(value.FragmentSequence)
		clone = &OtherWordFragment{FragmentSequence: fragments}
	case *FragmentSequence:
		clone, err = cloneFragmentSequenceChecked(value)
	case *base:
		clone, err = cloneBaseChecked(value)
	default:
		var clonedBuffer buffer.ByteBuffer
		clonedBuffer, err = buffer.Clone(elem.Buffer())
		clone = &snapshotElement{
			base:  newBaseWithByteOrder(elem.Tag().Clone(), elem.ValueRepresentation(), clonedBuffer, nil),
			count: elem.Count(),
		}
	}
	if err != nil {
		return nil, fmt.Errorf("clone element %s: %w", elem.Tag(), err)
	}
	return clone, nil
}

func cloneBase(source *base) *base {
	if source == nil {
		return nil
	}
	return newBaseWithByteOrder(source.tag.Clone(), source.vr, cloneBuffer(source.buffer), source.getByteOrder())
}

func cloneString(source *String) *String {
	if source == nil {
		return nil
	}
	return &String{
		base:      cloneBase(source.base),
		encoding:  source.encoding,
		encodings: append([]encoding.Encoding(nil), source.encodings...),
	}
}

func cloneFragmentSequence(source *FragmentSequence) *FragmentSequence {
	if source == nil {
		return nil
	}
	fragments := make([]buffer.ByteBuffer, len(source.fragments))
	for index, fragment := range source.fragments {
		fragments[index] = cloneBuffer(fragment)
	}
	return &FragmentSequence{
		tag:         source.tag.Clone(),
		vr:          source.vr,
		offsetTable: append([]uint32(nil), source.offsetTable...),
		fragments:   fragments,
	}
}

func cloneBaseChecked(source *base) (*base, error) {
	if source == nil {
		return nil, nil
	}
	clonedBuffer, err := buffer.Clone(source.buffer)
	if err != nil {
		return nil, err
	}
	return newBaseWithByteOrder(source.tag.Clone(), source.vr, clonedBuffer, source.getByteOrder()), nil
}

func cloneStringChecked(source *String) (*String, error) {
	if source == nil {
		return nil, nil
	}
	base, err := cloneBaseChecked(source.base)
	if err != nil {
		return nil, err
	}
	return &String{
		base:      base,
		encoding:  source.encoding,
		encodings: append([]encoding.Encoding(nil), source.encodings...),
	}, nil
}

func cloneFragmentSequenceChecked(source *FragmentSequence) (*FragmentSequence, error) {
	if source == nil {
		return nil, nil
	}
	fragments := make([]buffer.ByteBuffer, len(source.fragments))
	for index, fragment := range source.fragments {
		cloned, err := buffer.Clone(fragment)
		if err != nil {
			return nil, fmt.Errorf("clone fragment %d: %w", index, err)
		}
		fragments[index] = cloned
	}
	return &FragmentSequence{
		tag:         source.tag.Clone(),
		vr:          source.vr,
		offsetTable: append([]uint32(nil), source.offsetTable...),
		fragments:   fragments,
	}, nil
}

func cloneBuffer(source buffer.ByteBuffer) buffer.ByteBuffer {
	if source == nil {
		return nil
	}
	return buffer.NewMemory(append([]byte(nil), source.Data()...))
}

type snapshotElement struct {
	*base
	count int
}

func (e *snapshotElement) Count() int {
	return e.count
}
