// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// OtherByte represents a DICOM element with VR = OB (Other Byte).
// Stores binary data as a byte sequence.
type OtherByte struct {
	*base
}

// NewOtherByte creates a new OB element with the given data.
func NewOtherByte(t *tag.Tag, data []byte) *OtherByte {
	buf := buffer.NewMemory(data)
	return &OtherByte{
		base: newBase(t, vr.OB, buf),
	}
}

// NewOtherByteFromBuffer creates an OB element from an existing buffer.
func NewOtherByteFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherByte {
	return &OtherByte{
		base: newBase(t, vr.OB, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherByte) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherByte) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}

// OtherWord represents a DICOM element with VR = OW (Other Word).
// Stores binary data as a sequence of 16-bit words.
type OtherWord struct {
	*base
}

// NewOtherWord creates a new OW element with the given data.
func NewOtherWord(t *tag.Tag, data []byte) *OtherWord {
	buf := buffer.NewMemory(data)
	return &OtherWord{
		base: newBase(t, vr.OW, buf),
	}
}

// NewOtherWordFromBuffer creates an OW element from an existing buffer.
func NewOtherWordFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherWord {
	return &OtherWord{
		base: newBase(t, vr.OW, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherWord) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherWord) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}

// Unknown represents a DICOM element with VR = UN (Unknown).
// Stores data of unknown type.
type Unknown struct {
	*base
}

// NewUnknown creates a new UN element with the given data.
func NewUnknown(t *tag.Tag, data []byte) *Unknown {
	buf := buffer.NewMemory(data)
	return &Unknown{
		base: newBase(t, vr.UN, buf),
	}
}

// NewUnknownFromBuffer creates a UN element from an existing buffer.
func NewUnknownFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *Unknown {
	return &Unknown{
		base: newBase(t, vr.UN, buf),
	}
}

// Count returns 1 (unknown data is treated as a single value).
func (u *Unknown) Count() int {
	return 1
}

// GetData returns the binary data.
func (u *Unknown) GetData() []byte {
	if u.buffer == nil {
		return nil
	}
	return u.buffer.Data()
}

// OtherDouble represents a DICOM element with VR = OD (Other Double).
// Stores binary data as a sequence of 64-bit floating point values.
type OtherDouble struct {
	*base
}

// NewOtherDouble creates a new OD element with the given data.
func NewOtherDouble(t *tag.Tag, data []byte) *OtherDouble {
	buf := buffer.NewMemory(data)
	return &OtherDouble{
		base: newBase(t, vr.OD, buf),
	}
}

// NewOtherDoubleFromBuffer creates an OD element from an existing buffer.
func NewOtherDoubleFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherDouble {
	return &OtherDouble{
		base: newBase(t, vr.OD, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherDouble) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherDouble) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}

// OtherFloat represents a DICOM element with VR = OF (Other Float).
// Stores binary data as a sequence of 32-bit floating point values.
type OtherFloat struct {
	*base
}

// NewOtherFloat creates a new OF element with the given data.
func NewOtherFloat(t *tag.Tag, data []byte) *OtherFloat {
	buf := buffer.NewMemory(data)
	return &OtherFloat{
		base: newBase(t, vr.OF, buf),
	}
}

// NewOtherFloatFromBuffer creates an OF element from an existing buffer.
func NewOtherFloatFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherFloat {
	return &OtherFloat{
		base: newBase(t, vr.OF, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherFloat) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherFloat) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}

// OtherLong represents a DICOM element with VR = OL (Other Long).
// Stores binary data as a sequence of 32-bit integers.
type OtherLong struct {
	*base
}

// NewOtherLong creates a new OL element with the given data.
func NewOtherLong(t *tag.Tag, data []byte) *OtherLong {
	buf := buffer.NewMemory(data)
	return &OtherLong{
		base: newBase(t, vr.OL, buf),
	}
}

// NewOtherLongFromBuffer creates an OL element from an existing buffer.
func NewOtherLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherLong {
	return &OtherLong{
		base: newBase(t, vr.OL, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherLong) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherLong) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}

// OtherVeryLong represents a DICOM element with VR = OV (Other 64-bit Very Long).
// Stores binary data as a sequence of 64-bit integers.
type OtherVeryLong struct {
	*base
}

// NewOtherVeryLong creates a new OV element with the given data.
func NewOtherVeryLong(t *tag.Tag, data []byte) *OtherVeryLong {
	buf := buffer.NewMemory(data)
	return &OtherVeryLong{
		base: newBase(t, vr.OV, buf),
	}
}

// NewOtherVeryLongFromBuffer creates an OV element from an existing buffer.
func NewOtherVeryLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *OtherVeryLong {
	return &OtherVeryLong{
		base: newBase(t, vr.OV, buf),
	}
}

// Count returns 1 (binary data is treated as a single value).
func (o *OtherVeryLong) Count() int {
	return 1
}

// GetData returns the binary data.
func (o *OtherVeryLong) GetData() []byte {
	if o.buffer == nil {
		return nil
	}
	return o.buffer.Data()
}
