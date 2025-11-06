// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
)

// EndianByteBuffer is a ByteBuffer implementation that performs endian conversion.
//
// This buffer wraps another buffer and converts the byte order when reading data.
// It's used when the data's endianness doesn't match the local machine's endianness.
type EndianByteBuffer struct {
	internal ByteBuffer
	endian   endian.Endian
	unitSize int // Size of each unit to swap (1=byte, 2=uint16, 4=uint32, 8=uint64)
}

// NewEndian creates an EndianByteBuffer that wraps the given buffer.
// endianness is the byte order of the data in the buffer.
// unitSize is the size of each unit to swap (typically 2, 4, or 8 bytes).
//
// If endianness matches the local machine or unitSize is 1, this returns
// the original buffer without wrapping (optimization).
func NewEndian(buffer ByteBuffer, endianness endian.Endian, unitSize int) ByteBuffer {
	if buffer == nil {
		return nil
	}

	// No conversion needed if endianness matches local machine
	if endianness == endian.LocalMachine {
		return buffer
	}

	// No conversion needed for single bytes
	if unitSize == 1 {
		return buffer
	}

	return &EndianByteBuffer{
		internal: buffer,
		endian:   endianness,
		unitSize: unitSize,
	}
}

// Internal returns the wrapped buffer.
func (e *EndianByteBuffer) Internal() ByteBuffer {
	return e.internal
}

// Endian returns the endianness of the buffer.
func (e *EndianByteBuffer) Endian() endian.Endian {
	return e.endian
}

// UnitSize returns the unit size for byte swapping.
func (e *EndianByteBuffer) UnitSize() int {
	return e.unitSize
}

// IsMemory delegates to the internal buffer.
func (e *EndianByteBuffer) IsMemory() bool {
	return e.internal.IsMemory()
}

// Size delegates to the internal buffer.
func (e *EndianByteBuffer) Size() uint32 {
	return e.internal.Size()
}

// Data returns the data with endian conversion applied.
func (e *EndianByteBuffer) Data() []byte {
	data := e.internal.Data()

	// Make a copy and swap
	result := make([]byte, len(data))
	copy(result, data)

	e.swapBytes(result)
	return result
}

// GetByteRange gets a range of bytes with endian conversion applied.
func (e *EndianByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if err := e.internal.GetByteRange(offset, count, output); err != nil {
		return err
	}

	// Swap bytes in the output buffer
	e.swapBytes(output[:count])
	return nil
}

// WriteTo writes the buffer contents with endian conversion to the writer.
func (e *EndianByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	// Read in chunks and write with swapping
	const bufferSize = 1024 * 1024 // 1MB chunks
	remaining := e.Size()
	var totalWritten int64
	offset := uint32(0)

	for remaining > 0 {
		chunkSize := remaining
		if chunkSize > bufferSize {
			chunkSize = bufferSize
		}

		chunk := make([]byte, chunkSize)
		if err := e.GetByteRange(offset, chunkSize, chunk); err != nil {
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

// swapBytes performs in-place byte swapping based on unitSize.
func (e *EndianByteBuffer) swapBytes(data []byte) {
	switch e.unitSize {
	case 2:
		endian.SwapBytesN(2, data, len(data))
	case 4:
		endian.SwapBytesN(4, data, len(data))
	case 8:
		endian.SwapBytesN(8, data, len(data))
	default:
		// No swapping for unitSize == 1 or other values
	}
}
