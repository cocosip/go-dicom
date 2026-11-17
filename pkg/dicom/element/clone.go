// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package element

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// bufferCloner copies one value buffer and reports I/O errors for buffers that
// are backed by a file.
type bufferCloner func(buffer.ByteBuffer) (buffer.ByteBuffer, error)

// strictClone propagates buffer I/O errors.
func strictClone(source buffer.ByteBuffer) (buffer.ByteBuffer, error) {
	return buffer.Clone(source)
}

// lenientClone keeps the original buffer when it cannot be detached, so
// DeepClone always returns an element (DeepCloneChecked is the strict variant).
func lenientClone(source buffer.ByteBuffer) (buffer.ByteBuffer, error) {
	if source == nil {
		return nil, nil
	}
	cloned, err := buffer.Clone(source)
	if err != nil {
		return source, nil
	}
	return cloned, nil
}

// DeepClone returns an independent copy of an element. All built-in element
// types retain their concrete type, decoding state, and value bytes. Use
// DeepCloneChecked when a value buffer may be backed by a file and detaching its
// bytes can fail.
func DeepClone(elem Element) Element {
	clone, _ := cloneElement(elem, lenientClone)
	return clone
}

// DeepCloneChecked returns an independent copy of an element and reports any
// buffer I/O error encountered while detaching its value bytes.
func DeepCloneChecked(elem Element) (Element, error) {
	clone, err := cloneElement(elem, strictClone)
	if err != nil {
		return nil, fmt.Errorf("clone element %s: %w", elem.Tag(), err)
	}
	return clone, nil
}

// cloneElement is the single clone implementation behind both exported
// variants; cloneBuffer decides whether a buffer that cannot be read is an
// error or is simply shared with the source.
func cloneElement(elem Element, cloneBuffer bufferCloner) (Element, error) {
	if elem == nil {
		return nil, nil
	}

	cloneBase := func(source *base) (*base, error) {
		if source == nil {
			return nil, nil
		}
		clonedBuffer, err := cloneBuffer(source.buffer)
		if err != nil {
			return nil, err
		}
		return newBaseWithByteOrder(source.tag.Clone(), source.vr, clonedBuffer, source.byteOrder), nil
	}

	if clone, matched, err := cloneStringElement(elem, cloneBase); matched {
		return clone, err
	}
	if clone, matched, err := cloneBaseElement(elem, cloneBase); matched {
		return clone, err
	}

	switch value := elem.(type) {
	case *OtherByteFragment:
		fragments, err := cloneFragmentSequence(value.FragmentSequence, cloneBuffer)
		return &OtherByteFragment{FragmentSequence: fragments}, err
	case *OtherWordFragment:
		fragments, err := cloneFragmentSequence(value.FragmentSequence, cloneBuffer)
		return &OtherWordFragment{FragmentSequence: fragments}, err
	case *FragmentSequence:
		return cloneFragmentSequence(value, cloneBuffer)
	case *base:
		return cloneBase(value)
	default:
		clonedBuffer, err := cloneBuffer(elem.Buffer())
		return &snapshotElement{
			base:  newBaseWithByteOrder(elem.Tag().Clone(), elem.ValueRepresentation(), clonedBuffer, nil),
			count: elem.Count(),
		}, err
	}
}

// cloneStringElement clones the string elements that wrap a *String; matched
// reports whether elem is one of them.
func cloneStringElement(elem Element, cloneBase func(*base) (*base, error)) (clone Element, matched bool, err error) {
	switch value := elem.(type) {
	case *String:
		clone, err = cloneString(value, cloneBase)
	case *DecimalString:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &DecimalString{str: str}
	case *IntegerString:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &IntegerString{str: str}
	case *Date:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &Date{str: str}
	case *Time:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &Time{str: str}
	case *DateTime:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &DateTime{str: str}
	case *PersonName:
		var str *String
		str, err = cloneString(value.str, cloneBase)
		clone = &PersonName{str: str}
	default:
		return nil, false, nil
	}
	return clone, true, err
}

// cloneBaseElement clones the elements whose only field is the shared *base;
// matched reports whether elem is one of them.
func cloneBaseElement(elem Element, cloneBase func(*base) (*base, error)) (clone Element, matched bool, err error) {
	switch value := elem.(type) {
	case *AttributeTag:
		clonedBase, err := cloneBase(value.base)
		return &AttributeTag{base: clonedBase}, true, err
	case *UnsignedShort:
		clonedBase, err := cloneBase(value.base)
		return &UnsignedShort{base: clonedBase}, true, err
	case *UnsignedLong:
		clonedBase, err := cloneBase(value.base)
		return &UnsignedLong{base: clonedBase}, true, err
	case *SignedShort:
		clonedBase, err := cloneBase(value.base)
		return &SignedShort{base: clonedBase}, true, err
	case *SignedLong:
		clonedBase, err := cloneBase(value.base)
		return &SignedLong{base: clonedBase}, true, err
	case *Float:
		clonedBase, err := cloneBase(value.base)
		return &Float{base: clonedBase}, true, err
	case *Double:
		clonedBase, err := cloneBase(value.base)
		return &Double{base: clonedBase}, true, err
	case *SignedVeryLong:
		clonedBase, err := cloneBase(value.base)
		return &SignedVeryLong{base: clonedBase}, true, err
	case *UnsignedVeryLong:
		clonedBase, err := cloneBase(value.base)
		return &UnsignedVeryLong{base: clonedBase}, true, err
	case *OtherByte:
		clonedBase, err := cloneBase(value.base)
		return &OtherByte{base: clonedBase}, true, err
	case *OtherWord:
		clonedBase, err := cloneBase(value.base)
		return &OtherWord{base: clonedBase}, true, err
	case *Unknown:
		clonedBase, err := cloneBase(value.base)
		return &Unknown{base: clonedBase}, true, err
	case *OtherDouble:
		clonedBase, err := cloneBase(value.base)
		return &OtherDouble{base: clonedBase}, true, err
	case *OtherFloat:
		clonedBase, err := cloneBase(value.base)
		return &OtherFloat{base: clonedBase}, true, err
	case *OtherLong:
		clonedBase, err := cloneBase(value.base)
		return &OtherLong{base: clonedBase}, true, err
	case *OtherVeryLong:
		clonedBase, err := cloneBase(value.base)
		return &OtherVeryLong{base: clonedBase}, true, err
	default:
		return nil, false, nil
	}
}
func cloneString(source *String, cloneBase func(*base) (*base, error)) (*String, error) {
	if source == nil {
		return nil, nil
	}
	clonedBase, err := cloneBase(source.base)
	if err != nil {
		return nil, err
	}
	return &String{
		base:                 clonedBase,
		encoding:             source.encoding,
		encodings:            append([]encoding.Encoding(nil), source.encodings...),
		encodeErr:            source.encodeErr,
		requiresCharacterSet: source.requiresCharacterSet,
	}, nil
}

func cloneFragmentSequence(source *FragmentSequence, cloneBuffer bufferCloner) (*FragmentSequence, error) {
	if source == nil {
		return nil, nil
	}
	fragments := make([]buffer.ByteBuffer, len(source.fragments))
	for index, fragment := range source.fragments {
		cloned, err := cloneBuffer(fragment)
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

type snapshotElement struct {
	*base
	count int
}

func (e *snapshotElement) Count() int {
	return e.count
}
