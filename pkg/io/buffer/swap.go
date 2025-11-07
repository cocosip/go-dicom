// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
)

// SwapByteBuffer is a buffer wrapper that swaps bytes in groups.
// This is used for specific DICOM transfer syntaxes (particularly GE private syntax)
// where bytes need to be swapped differently than endianness conversion.
//
// Note: This is different from EndianByteBuffer:
// - EndianByteBuffer converts between different endianness (little <-> big)
// - SwapByteBuffer always swaps bytes regardless of machine endianness
type SwapByteBuffer struct {
	internal ByteBuffer // The underlying buffer
	unitSize int        // Number of bytes to swap as a unit (2, 4, etc.)
}

// NewSwap creates a new swap buffer that wraps another buffer.
//
// Parameters:
//   - buffer: The underlying buffer to wrap
//   - unitSize: The number of bytes to swap as a unit (e.g., 2 for 16-bit swaps, 4 for 32-bit swaps)
//
// Returns an error if unitSize is invalid.
func NewSwap(buffer ByteBuffer, unitSize int) (*SwapByteBuffer, error) {
	if buffer == nil {
		return nil, fmt.Errorf("buffer cannot be nil")
	}

	if unitSize < 1 {
		return nil, fmt.Errorf("unitSize must be positive, got %d", unitSize)
	}

	return &SwapByteBuffer{
		internal: buffer,
		unitSize: unitSize,
	}, nil
}

// IsMemory returns whether the underlying buffer is memory-backed.
func (s *SwapByteBuffer) IsMemory() bool {
	return s.internal.IsMemory()
}

// Size returns the size of the buffer.
func (s *SwapByteBuffer) Size() uint32 {
	return s.internal.Size()
}

// Data returns all data with bytes swapped.
// If the underlying buffer is memory-backed, this makes a copy before swapping.
// Otherwise, it gets the data from the internal buffer and swaps it in place.
func (s *SwapByteBuffer) Data() []byte {
	var data []byte

	// If memory-backed, make a copy to avoid modifying the original
	if s.internal.IsMemory() {
		originalData := s.internal.Data()
		if originalData == nil {
			return nil
		}
		data = make([]byte, len(originalData))
		copy(data, originalData)
	} else {
		// For non-memory buffers (file, stream, etc.), get a fresh copy
		data = s.internal.Data()
		if data == nil {
			return nil
		}
	}

	// Swap bytes in place
	endian.SwapBytes(s.unitSize, data)

	return data
}

// GetByteRange reads a range of bytes and swaps them.
//
// Parameters:
//   - offset: Offset from the start of the buffer
//   - count: Number of bytes to read
//   - output: Destination buffer (must be at least 'count' bytes long)
//
// Returns an error if the operation fails.
func (s *SwapByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}

	if uint32(len(output)) < count {
		return fmt.Errorf("output buffer with %d bytes cannot fit %d bytes of data", len(output), count)
	}

	// Get the bytes from the internal buffer
	err := s.internal.GetByteRange(offset, count, output)
	if err != nil {
		return err
	}

	// Swap bytes in the output
	endian.SwapBytesN(s.unitSize, output, int(count))

	return nil
}

// WriteTo writes the buffer contents to the given writer with bytes swapped.
// This uses chunked reading to handle large buffers efficiently.
//
// Returns the number of bytes written and any error encountered.
func (s *SwapByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	size := s.Size()
	if size == 0 {
		return 0, nil
	}

	// Use 1MB chunks for efficiency
	const chunkSize = 1024 * 1024
	buffer := make([]byte, chunkSize)
	var totalWritten int64

	offset := uint32(0)
	remaining := size

	for remaining > 0 {
		// Determine how much to read in this iteration
		toRead := remaining
		if toRead > chunkSize {
			toRead = chunkSize
		}

		// Read and swap the chunk
		err := s.GetByteRange(offset, toRead, buffer)
		if err != nil {
			return totalWritten, fmt.Errorf("failed to read chunk at offset %d: %w", offset, err)
		}

		// Write the chunk
		n, err := w.Write(buffer[:toRead])
		totalWritten += int64(n)
		if err != nil {
			return totalWritten, fmt.Errorf("failed to write chunk: %w", err)
		}

		offset += toRead
		remaining -= toRead
	}

	return totalWritten, nil
}

// Internal returns the underlying buffer.
func (s *SwapByteBuffer) Internal() ByteBuffer {
	return s.internal
}

// UnitSize returns the swap unit size.
func (s *SwapByteBuffer) UnitSize() int {
	return s.unitSize
}
