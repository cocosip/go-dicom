// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// AttributeTag represents a DICOM element with VR = AT (Attribute Tag).
// Each value is a 32-bit tag (group + element).
type AttributeTag struct {
	*base
}

// NewAttributeTag creates a new AT element with the given tag values.
func NewAttributeTag(t *tag.Tag, values []*tag.Tag) *AttributeTag {
	return NewAttributeTagWithEndian(t, values, endian.LocalMachine)
}

// NewAttributeTagWithEndian creates a new AT element with specific byte order.
func NewAttributeTagWithEndian(t *tag.Tag, values []*tag.Tag, e endian.Endian) *AttributeTag {
	data := make([]byte, len(values)*4)
	for i, v := range values {
		if v == nil {
			continue
		}
		// Tag is stored as: group (2 bytes) + element (2 bytes)
		if e == endian.Little {
			binary.LittleEndian.PutUint16(data[i*4:], v.Group())
			binary.LittleEndian.PutUint16(data[i*4+2:], v.Element())
		} else {
			binary.BigEndian.PutUint16(data[i*4:], v.Group())
			binary.BigEndian.PutUint16(data[i*4+2:], v.Element())
		}
	}

	buf := buffer.NewMemory(data)
	return &AttributeTag{
		base: NewBase(t, vr.AT, buf),
	}
}

// NewAttributeTagFromBuffer creates an AT element from an existing buffer.
func NewAttributeTagFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *AttributeTag {
	return &AttributeTag{
		base: NewBase(t, vr.AT, buf),
	}
}

// Count returns the number of tags (length / 4).
func (a *AttributeTag) Count() int {
	return int(a.Length() / 4)
}

// GetValue returns the tag at the specified index.
func (a *AttributeTag) GetValue(index int) (*tag.Tag, error) {
	if index < 0 || index >= a.Count() {
		return nil, fmt.Errorf("index %d out of range [0, %d)", index, a.Count())
	}

	data := make([]byte, 4)
	offset := uint32(index * 4)
	if err := a.buffer.GetByteRange(offset, 4, data); err != nil {
		return nil, err
	}

	// Assume little endian by default
	group := binary.LittleEndian.Uint16(data[0:2])
	element := binary.LittleEndian.Uint16(data[2:4])

	return tag.New(group, element), nil
}

// GetValues returns all tags as a slice.
func (a *AttributeTag) GetValues() ([]*tag.Tag, error) {
	count := a.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]*tag.Tag, count)
	for i := 0; i < count; i++ {
		v, err := a.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}
