// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// CompositeByteBuffer is a ByteBuffer implementation that combines multiple buffers.
//
// This is useful for efficiently concatenating buffers without copying data.
// The composite buffer presents a unified view over all contained buffers.
type CompositeByteBuffer struct {
	buffers []ByteBuffer
}

// NewComposite creates a new CompositeByteBuffer from the given buffers.
func NewComposite(buffers ...ByteBuffer) *CompositeByteBuffer {
	if buffers == nil {
		buffers = []ByteBuffer{}
	}
	return &CompositeByteBuffer{buffers: buffers}
}

// Buffers returns the slice of buffers that make up this composite.
func (c *CompositeByteBuffer) Buffers() []ByteBuffer {
	return c.buffers
}

// Add appends a buffer to the composite.
func (c *CompositeByteBuffer) Add(buf ByteBuffer) {
	c.buffers = append(c.buffers, buf)
}

// IsMemory returns true only if all contained buffers are in memory.
func (c *CompositeByteBuffer) IsMemory() bool {
	for _, buf := range c.buffers {
		if !buf.IsMemory() {
			return false
		}
	}
	return true
}

// Size returns the total size of all contained buffers.
func (c *CompositeByteBuffer) Size() uint32 {
	var total uint32
	for _, buf := range c.buffers {
		total += buf.Size()
	}
	return total
}

// Data returns all data concatenated into a single byte slice.
// This will allocate memory for the entire composite buffer.
func (c *CompositeByteBuffer) Data() []byte {
	size := c.Size()
	if size == 0 {
		return []byte{}
	}

	data := make([]byte, size)
	offset := uint32(0)
	for _, buf := range c.buffers {
		bufData := buf.Data()
		copy(data[offset:], bufData)
		offset += uint32(len(bufData))
	}
	return data
}

// GetByteRange copies a range of bytes from the composite buffer to the output slice.
// This handles spanning across multiple contained buffers.
func (c *CompositeByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}
	if uint32(len(output)) < count {
		return fmt.Errorf("output buffer length %d is less than requested count %d", len(output), count)
	}
	if offset+count > c.Size() {
		return fmt.Errorf("offset %d + count %d exceeds buffer size %d", offset, count, c.Size())
	}

	// Find the starting buffer
	bufIdx := 0
	bufOffset := offset
	for bufIdx < len(c.buffers) && bufOffset >= c.buffers[bufIdx].Size() {
		bufOffset -= c.buffers[bufIdx].Size()
		bufIdx++
	}

	// Copy data from buffers
	outputOffset := uint32(0)
	remaining := count

	for bufIdx < len(c.buffers) && remaining > 0 {
		currentBuf := c.buffers[bufIdx]
		available := currentBuf.Size() - bufOffset
		toCopy := remaining
		if toCopy > available {
			toCopy = available
		}

		// Get data from current buffer
		temp := make([]byte, toCopy)
		if err := currentBuf.GetByteRange(bufOffset, toCopy, temp); err != nil {
			return fmt.Errorf("error reading from buffer %d: %w", bufIdx, err)
		}

		copy(output[outputOffset:], temp)

		remaining -= toCopy
		outputOffset += toCopy
		bufOffset = 0 // Reset offset for next buffer
		bufIdx++
	}

	return nil
}

// WriteTo writes all contained buffers to the writer in sequence.
func (c *CompositeByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	var totalWritten int64
	for i, buf := range c.buffers {
		n, err := buf.WriteTo(w)
		totalWritten += n
		if err != nil {
			return totalWritten, fmt.Errorf("error writing buffer %d: %w", i, err)
		}
	}

	return totalWritten, nil
}
