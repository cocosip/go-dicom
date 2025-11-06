// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package buffer implements byte buffer abstractions for DICOM data.
package buffer

import (
	"io"
)

// ByteBuffer is the common interface for byte buffers.
//
// ByteBuffer provides an abstraction for storing and accessing byte data,
// whether in memory, on disk, or from other sources. This is used throughout
// the DICOM library to handle element data efficiently.
type ByteBuffer interface {
	// IsMemory returns whether data is buffered in memory.
	// If false, accessing the data may require I/O operations.
	IsMemory() bool

	// Size returns the size of the buffered data in bytes.
	Size() uint32

	// Data returns the entire buffer as a byte slice.
	// For non-memory buffers, this may load all data into memory.
	Data() []byte

	// GetByteRange copies a subset of the data into the provided output buffer.
	// offset is the starting position in the buffer (0-based).
	// count is the number of bytes to copy.
	// output must have capacity >= count.
	GetByteRange(offset, count uint32, output []byte) error

	// WriteTo writes the buffer contents to the provided writer.
	// This is more efficient than calling Data() for large buffers.
	WriteTo(w io.Writer) (int64, error)
}
