// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// MemoryByteBuffer is a ByteBuffer implementation backed by an in-memory byte slice.
//
// This is the most common buffer type, storing data directly in memory.
// It takes ownership of the provided byte slice, so callers should not
// modify the slice after passing it to NewMemory.
type MemoryByteBuffer struct {
	data []byte
}

// NewMemory creates a new MemoryByteBuffer from the given byte slice.
// The buffer takes ownership of the slice - do not modify it after this call.
func NewMemory(data []byte) *MemoryByteBuffer {
	if data == nil {
		data = []byte{}
	}
	return &MemoryByteBuffer{data: data}
}

// NewMemoryWithSize creates a new MemoryByteBuffer with the specified size,
// initialized to zero bytes.
func NewMemoryWithSize(size uint32) *MemoryByteBuffer {
	return &MemoryByteBuffer{data: make([]byte, size)}
}

// IsMemory returns true (this buffer is in memory).
func (m *MemoryByteBuffer) IsMemory() bool {
	return true
}

// Size returns the size of the buffer in bytes.
func (m *MemoryByteBuffer) Size() uint32 {
	return uint32(len(m.data)) // #nosec G115 -- DICOM buffer size within uint32 range
}

// Data returns the underlying byte slice.
func (m *MemoryByteBuffer) Data() []byte {
	return m.data
}

// GetByteRange copies a range of bytes from the buffer to the output slice.
func (m *MemoryByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}
	if uint32(len(output)) < count { // #nosec G115 -- buffer size check
		return fmt.Errorf("output buffer length %d is less than requested count %d", len(output), count)
	}
	if offset+count > m.Size() {
		return fmt.Errorf("offset %d + count %d exceeds buffer size %d", offset, count, m.Size())
	}

	copy(output, m.data[offset:offset+count])
	return nil
}

// WriteTo writes the buffer contents to the provided writer.
func (m *MemoryByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	n, err := w.Write(m.data)
	return int64(n), err
}
