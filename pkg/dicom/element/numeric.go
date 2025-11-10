// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// UnsignedShort represents a DICOM element with VR = US (Unsigned Short).
// Stores 16-bit unsigned integers.
type UnsignedShort struct {
	*base
}

// NewUnsignedShort creates a new US element with the given values.
func NewUnsignedShort(t *tag.Tag, values []uint16) *UnsignedShort {
	return NewUnsignedShortWithEndian(t, values, endian.LocalMachine)
}

// NewUnsignedShortWithEndian creates a new US element with specific byte order.
func NewUnsignedShortWithEndian(t *tag.Tag, values []uint16, e endian.Endian) *UnsignedShort {
	data := make([]byte, len(values)*2)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint16(data[i*2:], v)
		} else {
			binary.BigEndian.PutUint16(data[i*2:], v)
		}
	}

	buf := buffer.NewMemory(data)
	return &UnsignedShort{
		base: newBase(t, vr.US, buf),
	}
}

// NewUnsignedShortFromBuffer creates a US element from an existing buffer.
func NewUnsignedShortFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *UnsignedShort {
	return &UnsignedShort{
		base: newBase(t, vr.US, buf),
	}
}

// Count returns the number of values (length / 2).
func (u *UnsignedShort) Count() int {
	return int(u.Length() / 2)
}

// GetValue returns the value at the specified index.
func (u *UnsignedShort) GetValue(index int) (uint16, error) {
	if index < 0 || index >= u.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, u.Count())
	}

	data := make([]byte, 2)
	offset := uint32(index * 2)
	if err := u.buffer.GetByteRange(offset, 2, data); err != nil {
		return 0, err
	}

	// Assume little endian by default (most common in DICOM)
	return binary.LittleEndian.Uint16(data), nil
}

// GetValues returns all values as a slice.
func (u *UnsignedShort) GetValues() ([]uint16, error) {
	count := u.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]uint16, count)
	for i := 0; i < count; i++ {
		v, err := u.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// UnsignedLong represents a DICOM element with VR = UL (Unsigned Long).
// Stores 32-bit unsigned integers.
type UnsignedLong struct {
	*base
}

// NewUnsignedLong creates a new UL element with the given values.
func NewUnsignedLong(t *tag.Tag, values []uint32) *UnsignedLong {
	return NewUnsignedLongWithEndian(t, values, endian.LocalMachine)
}

// NewUnsignedLongWithEndian creates a new UL element with specific byte order.
func NewUnsignedLongWithEndian(t *tag.Tag, values []uint32, e endian.Endian) *UnsignedLong {
	data := make([]byte, len(values)*4)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint32(data[i*4:], v)
		} else {
			binary.BigEndian.PutUint32(data[i*4:], v)
		}
	}

	buf := buffer.NewMemory(data)
	return &UnsignedLong{
		base: newBase(t, vr.UL, buf),
	}
}

// NewUnsignedLongFromBuffer creates a UL element from an existing buffer.
func NewUnsignedLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *UnsignedLong {
	return &UnsignedLong{
		base: newBase(t, vr.UL, buf),
	}
}

// Count returns the number of values (length / 4).
func (u *UnsignedLong) Count() int {
	return int(u.Length() / 4)
}

// GetValue returns the value at the specified index.
func (u *UnsignedLong) GetValue(index int) (uint32, error) {
	if index < 0 || index >= u.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, u.Count())
	}

	data := make([]byte, 4)
	offset := uint32(index * 4)
	if err := u.buffer.GetByteRange(offset, 4, data); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(data), nil
}

// GetValues returns all values as a slice.
func (u *UnsignedLong) GetValues() ([]uint32, error) {
	count := u.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]uint32, count)
	for i := 0; i < count; i++ {
		v, err := u.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// SignedShort represents a DICOM element with VR = SS (Signed Short).
// Stores 16-bit signed integers.
type SignedShort struct {
	*base
}

// NewSignedShort creates a new SS element with the given values.
func NewSignedShort(t *tag.Tag, values []int16) *SignedShort {
	return NewSignedShortWithEndian(t, values, endian.LocalMachine)
}

// NewSignedShortWithEndian creates a new SS element with specific byte order.
func NewSignedShortWithEndian(t *tag.Tag, values []int16, e endian.Endian) *SignedShort {
	data := make([]byte, len(values)*2)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint16(data[i*2:], uint16(v))
		} else {
			binary.BigEndian.PutUint16(data[i*2:], uint16(v))
		}
	}

	buf := buffer.NewMemory(data)
	return &SignedShort{
		base: newBase(t, vr.SS, buf),
	}
}

// NewSignedShortFromBuffer creates a SS element from an existing buffer.
func NewSignedShortFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *SignedShort {
	return &SignedShort{
		base: newBase(t, vr.SS, buf),
	}
}

// Count returns the number of values (length / 2).
func (s *SignedShort) Count() int {
	return int(s.Length() / 2)
}

// GetValue returns the value at the specified index.
func (s *SignedShort) GetValue(index int) (int16, error) {
	if index < 0 || index >= s.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, s.Count())
	}

	data := make([]byte, 2)
	offset := uint32(index * 2)
	if err := s.buffer.GetByteRange(offset, 2, data); err != nil {
		return 0, err
	}

	return int16(binary.LittleEndian.Uint16(data)), nil
}

// GetValues returns all values as a slice.
func (s *SignedShort) GetValues() ([]int16, error) {
	count := s.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]int16, count)
	for i := 0; i < count; i++ {
		v, err := s.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// SignedLong represents a DICOM element with VR = SL (Signed Long).
// Stores 32-bit signed integers.
type SignedLong struct {
	*base
}

// NewSignedLong creates a new SL element with the given values.
func NewSignedLong(t *tag.Tag, values []int32) *SignedLong {
	return NewSignedLongWithEndian(t, values, endian.LocalMachine)
}

// NewSignedLongWithEndian creates a new SL element with specific byte order.
func NewSignedLongWithEndian(t *tag.Tag, values []int32, e endian.Endian) *SignedLong {
	data := make([]byte, len(values)*4)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint32(data[i*4:], uint32(v))
		} else {
			binary.BigEndian.PutUint32(data[i*4:], uint32(v))
		}
	}

	buf := buffer.NewMemory(data)
	return &SignedLong{
		base: newBase(t, vr.SL, buf),
	}
}

// NewSignedLongFromBuffer creates a SL element from an existing buffer.
func NewSignedLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *SignedLong {
	return &SignedLong{
		base: newBase(t, vr.SL, buf),
	}
}

// Count returns the number of values (length / 4).
func (s *SignedLong) Count() int {
	return int(s.Length() / 4)
}

// GetValue returns the value at the specified index.
func (s *SignedLong) GetValue(index int) (int32, error) {
	if index < 0 || index >= s.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, s.Count())
	}

	data := make([]byte, 4)
	offset := uint32(index * 4)
	if err := s.buffer.GetByteRange(offset, 4, data); err != nil {
		return 0, err
	}

	return int32(binary.LittleEndian.Uint32(data)), nil
}

// GetValues returns all values as a slice.
func (s *SignedLong) GetValues() ([]int32, error) {
	count := s.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]int32, count)
	for i := 0; i < count; i++ {
		v, err := s.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// Float represents a DICOM element with VR = FL (Floating Point Single).
// Stores 32-bit floating point values.
type Float struct {
	*base
}

// NewFloat creates a new FL element with the given values.
func NewFloat(t *tag.Tag, values []float32) *Float {
	return NewFloatWithEndian(t, values, endian.LocalMachine)
}

// NewFloatWithEndian creates a new FL element with specific byte order.
func NewFloatWithEndian(t *tag.Tag, values []float32, e endian.Endian) *Float {
	data := make([]byte, len(values)*4)
	for i, v := range values {
		bits := math.Float32bits(v)
		if e == endian.Little {
			binary.LittleEndian.PutUint32(data[i*4:], bits)
		} else {
			binary.BigEndian.PutUint32(data[i*4:], bits)
		}
	}

	buf := buffer.NewMemory(data)
	return &Float{
		base: newBase(t, vr.FL, buf),
	}
}

// NewFloatFromBuffer creates a FL element from an existing buffer.
func NewFloatFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *Float {
	return &Float{
		base: newBase(t, vr.FL, buf),
	}
}

// Count returns the number of values (length / 4).
func (f *Float) Count() int {
	return int(f.Length() / 4)
}

// GetValue returns the value at the specified index.
func (f *Float) GetValue(index int) (float32, error) {
	if index < 0 || index >= f.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, f.Count())
	}

	data := make([]byte, 4)
	offset := uint32(index * 4)
	if err := f.buffer.GetByteRange(offset, 4, data); err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint32(data)
	return math.Float32frombits(bits), nil
}

// GetValues returns all values as a slice.
func (f *Float) GetValues() ([]float32, error) {
	count := f.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]float32, count)
	for i := 0; i < count; i++ {
		v, err := f.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// Double represents a DICOM element with VR = FD (Floating Point Double).
// Stores 64-bit floating point values.
type Double struct {
	*base
}

// NewDouble creates a new FD element with the given values.
func NewDouble(t *tag.Tag, values []float64) *Double {
	return NewDoubleWithEndian(t, values, endian.LocalMachine)
}

// NewDoubleWithEndian creates a new FD element with specific byte order.
func NewDoubleWithEndian(t *tag.Tag, values []float64, e endian.Endian) *Double {
	data := make([]byte, len(values)*8)
	for i, v := range values {
		bits := math.Float64bits(v)
		if e == endian.Little {
			binary.LittleEndian.PutUint64(data[i*8:], bits)
		} else {
			binary.BigEndian.PutUint64(data[i*8:], bits)
		}
	}

	buf := buffer.NewMemory(data)
	return &Double{
		base: newBase(t, vr.FD, buf),
	}
}

// NewDoubleFromBuffer creates a FD element from an existing buffer.
func NewDoubleFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *Double {
	return &Double{
		base: newBase(t, vr.FD, buf),
	}
}

// Count returns the number of values (length / 8).
func (d *Double) Count() int {
	return int(d.Length() / 8)
}

// GetValue returns the value at the specified index.
func (d *Double) GetValue(index int) (float64, error) {
	if index < 0 || index >= d.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, d.Count())
	}

	data := make([]byte, 8)
	offset := uint32(index * 8)
	if err := d.buffer.GetByteRange(offset, 8, data); err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint64(data)
	return math.Float64frombits(bits), nil
}

// GetValues returns all values as a slice.
func (d *Double) GetValues() ([]float64, error) {
	count := d.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]float64, count)
	for i := 0; i < count; i++ {
		v, err := d.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// SignedVeryLong represents a DICOM element with VR = SV (Signed 64-bit Very Long).
// Stores 64-bit signed integers.
type SignedVeryLong struct {
	*base
}

// NewSignedVeryLong creates a new SV element with the given values.
func NewSignedVeryLong(t *tag.Tag, values []int64) *SignedVeryLong {
	return NewSignedVeryLongWithEndian(t, values, endian.LocalMachine)
}

// NewSignedVeryLongWithEndian creates a new SV element with specific byte order.
func NewSignedVeryLongWithEndian(t *tag.Tag, values []int64, e endian.Endian) *SignedVeryLong {
	data := make([]byte, len(values)*8)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint64(data[i*8:], uint64(v))
		} else {
			binary.BigEndian.PutUint64(data[i*8:], uint64(v))
		}
	}

	buf := buffer.NewMemory(data)
	return &SignedVeryLong{
		base: newBase(t, vr.SV, buf),
	}
}

// NewSignedVeryLongFromBuffer creates a SV element from an existing buffer.
func NewSignedVeryLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *SignedVeryLong {
	return &SignedVeryLong{
		base: newBase(t, vr.SV, buf),
	}
}

// Count returns the number of values (length / 8).
func (s *SignedVeryLong) Count() int {
	return int(s.Length() / 8)
}

// GetValue returns the value at the specified index.
func (s *SignedVeryLong) GetValue(index int) (int64, error) {
	if index < 0 || index >= s.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, s.Count())
	}

	data := make([]byte, 8)
	offset := uint32(index * 8)
	if err := s.buffer.GetByteRange(offset, 8, data); err != nil {
		return 0, err
	}

	return int64(binary.LittleEndian.Uint64(data)), nil
}

// GetValues returns all values as a slice.
func (s *SignedVeryLong) GetValues() ([]int64, error) {
	count := s.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]int64, count)
	for i := 0; i < count; i++ {
		v, err := s.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}

// UnsignedVeryLong represents a DICOM element with VR = UV (Unsigned 64-bit Very Long).
// Stores 64-bit unsigned integers.
type UnsignedVeryLong struct {
	*base
}

// NewUnsignedVeryLong creates a new UV element with the given values.
func NewUnsignedVeryLong(t *tag.Tag, values []uint64) *UnsignedVeryLong {
	return NewUnsignedVeryLongWithEndian(t, values, endian.LocalMachine)
}

// NewUnsignedVeryLongWithEndian creates a new UV element with specific byte order.
func NewUnsignedVeryLongWithEndian(t *tag.Tag, values []uint64, e endian.Endian) *UnsignedVeryLong {
	data := make([]byte, len(values)*8)
	for i, v := range values {
		if e == endian.Little {
			binary.LittleEndian.PutUint64(data[i*8:], v)
		} else {
			binary.BigEndian.PutUint64(data[i*8:], v)
		}
	}

	buf := buffer.NewMemory(data)
	return &UnsignedVeryLong{
		base: newBase(t, vr.UV, buf),
	}
}

// NewUnsignedVeryLongFromBuffer creates a UV element from an existing buffer.
func NewUnsignedVeryLongFromBuffer(t *tag.Tag, buf buffer.ByteBuffer) *UnsignedVeryLong {
	return &UnsignedVeryLong{
		base: newBase(t, vr.UV, buf),
	}
}

// Count returns the number of values (length / 8).
func (u *UnsignedVeryLong) Count() int {
	return int(u.Length() / 8)
}

// GetValue returns the value at the specified index.
func (u *UnsignedVeryLong) GetValue(index int) (uint64, error) {
	if index < 0 || index >= u.Count() {
		return 0, fmt.Errorf("index %d out of range [0, %d)", index, u.Count())
	}

	data := make([]byte, 8)
	offset := uint32(index * 8)
	if err := u.buffer.GetByteRange(offset, 8, data); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(data), nil
}

// GetValues returns all values as a slice.
func (u *UnsignedVeryLong) GetValues() ([]uint64, error) {
	count := u.Count()
	if count == 0 {
		return nil, nil
	}

	values := make([]uint64, count)
	for i := 0; i < count; i++ {
		v, err := u.GetValue(i)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}
	return values, nil
}
