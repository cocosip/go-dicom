// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// RangeByteBuffer is a ByteBuffer that represents a sub-range of another buffer.
//
// This is useful for efficiently representing a portion of a larger buffer
// without copying the data. It's commonly used when parsing DICOM elements
// to create views into the file data.
type RangeByteBuffer struct {
	internal ByteBuffer
	offset   uint32
	length   uint32
}

// NewRange creates a new RangeByteBuffer that represents a sub-range of the given buffer.
// offset is the starting position in the internal buffer.
// length is the number of bytes in the range.
func NewRange(buffer ByteBuffer, offset, length uint32) (*RangeByteBuffer, error) {
	if buffer == nil {
		return nil, fmt.Errorf("buffer cannot be nil")
	}
	if offset+length > buffer.Size() {
		return nil, fmt.Errorf("offset %d + length %d exceeds buffer size %d", offset, length, buffer.Size())
	}

	return &RangeByteBuffer{
		internal: buffer,
		offset:   offset,
		length:   length,
	}, nil
}

// Internal returns the underlying buffer.
func (r *RangeByteBuffer) Internal() ByteBuffer {
	return r.internal
}

// Offset returns the starting offset in the internal buffer.
func (r *RangeByteBuffer) Offset() uint32 {
	return r.offset
}

// Length returns the length of this range.
func (r *RangeByteBuffer) Length() uint32 {
	return r.length
}

// IsMemory delegates to the internal buffer.
func (r *RangeByteBuffer) IsMemory() bool {
	return r.internal.IsMemory()
}

// Size returns the size of the range (same as Length).
func (r *RangeByteBuffer) Size() uint32 {
	return r.length
}

// Data returns the data in the range.
func (r *RangeByteBuffer) Data() []byte {
	data := make([]byte, r.length)
	if r.length == 0 {
		return data
	}

	// Read from internal buffer at offset
	if err := r.internal.GetByteRange(r.offset, r.length, data); err != nil {
		// Return empty data on error (best effort)
		return make([]byte, r.length)
	}

	return data
}

// GetByteRange gets a sub-range within this range.
// The offset is relative to this range's start, not the internal buffer's start.
func (r *RangeByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}
	if uint32(len(output)) < count { // #nosec G115 -- buffer size check
		return fmt.Errorf("output buffer length %d is less than requested count %d", len(output), count)
	}
	if offset+count > r.length {
		return fmt.Errorf("offset %d + count %d exceeds range size %d", offset, count, r.length)
	}

	// Translate to internal buffer's coordinate space
	internalOffset := r.offset + offset
	return r.internal.GetByteRange(internalOffset, count, output)
}

// WriteTo writes the range contents to the writer.
func (r *RangeByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	// Write in chunks to avoid loading all data into memory at once
	const bufferSize = 1024 * 1024 // 1MB chunks
	remaining := r.length
	offset := uint32(0)
	var totalWritten int64

	for remaining > 0 {
		chunkSize := remaining
		if chunkSize > bufferSize {
			chunkSize = bufferSize
		}

		chunk := make([]byte, chunkSize)
		if err := r.GetByteRange(offset, chunkSize, chunk); err != nil {
			return totalWritten, err
		}

		n, err := w.Write(chunk)
		totalWritten += int64(n)
		if err != nil {
			return totalWritten, err
		}

		offset += chunkSize
		remaining -= chunkSize
	}

	return totalWritten, nil
}
