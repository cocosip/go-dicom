// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// EvenLengthBuffer is a ByteBuffer wrapper that ensures even length.
//
// DICOM requires that all data elements have even length. This buffer
// adds a padding byte (0x00) to odd-length buffers to make them even.
type EvenLengthBuffer struct {
	internal ByteBuffer
}

// NewEvenLength creates an EvenLengthBuffer that wraps the given buffer.
// If the buffer is already even length, returns the original buffer without wrapping.
func NewEvenLength(buffer ByteBuffer) ByteBuffer {
	if buffer == nil {
		return nil
	}

	// If already even length, no wrapping needed
	if buffer.Size()%2 == 0 {
		return buffer
	}

	return &EvenLengthBuffer{internal: buffer}
}

// Internal returns the wrapped buffer.
func (e *EvenLengthBuffer) Internal() ByteBuffer {
	return e.internal
}

// IsMemory delegates to the internal buffer.
func (e *EvenLengthBuffer) IsMemory() bool {
	return e.internal.IsMemory()
}

// Size returns the size of the buffer plus 1 (to make it even).
func (e *EvenLengthBuffer) Size() uint32 {
	return e.internal.Size() + 1
}

// Data returns the internal data with a padding byte added at the end.
func (e *EvenLengthBuffer) Data() []byte {
	internalData := e.internal.Data()
	data := make([]byte, len(internalData)+1)
	copy(data, internalData)
	// Last byte is already zero from make()
	return data
}

// GetByteRange gets a range of bytes from the buffer, including the padding byte.
func (e *EvenLengthBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}
	if uint32(len(output)) < count { // #nosec G115 -- buffer size check
		return fmt.Errorf("output buffer length %d is less than requested count %d", len(output), count)
	}
	if offset+count > e.Size() {
		return fmt.Errorf("offset %d + count %d exceeds buffer size %d", offset, count, e.Size())
	}

	internalSize := e.internal.Size()

	// If the range is entirely within the internal buffer
	if offset+count <= internalSize {
		return e.internal.GetByteRange(offset, count, output)
	}

	// If the range starts within the internal buffer but extends to the padding
	if offset < internalSize {
		// Read what we can from the internal buffer
		bytesToRead := internalSize - offset
		if err := e.internal.GetByteRange(offset, bytesToRead, output); err != nil {
			return err
		}
		// Fill the rest with zeros (padding)
		for i := bytesToRead; i < count; i++ {
			output[i] = 0
		}
		return nil
	}

	// If the range is entirely in the padding (just the last byte)
	for i := uint32(0); i < count; i++ {
		output[i] = 0
	}
	return nil
}

// WriteTo writes the buffer contents with padding byte to the writer.
func (e *EvenLengthBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	// Write the internal buffer
	n, err := e.internal.WriteTo(w)
	if err != nil {
		return n, err
	}

	// Write the padding byte
	if _, err := w.Write([]byte{0}); err != nil {
		return n, err
	}

	return n + 1, nil
}
