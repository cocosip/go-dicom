// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
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
			base:  newBaseWithByteOrder(elem.Tag(), elem.ValueRepresentation(), cloneBuffer(elem.Buffer()), nil),
			count: elem.Count(),
		}
	}
}

func cloneBase(source *base) *base {
	if source == nil {
		return nil
	}
	return newBaseWithByteOrder(source.tag, source.vr, cloneBuffer(source.buffer), source.getByteOrder())
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
		tag:         source.tag,
		vr:          source.vr,
		offsetTable: append([]uint32(nil), source.offsetTable...),
		fragments:   fragments,
	}
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
