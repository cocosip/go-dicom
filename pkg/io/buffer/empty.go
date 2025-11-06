// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// emptyBuffer is a ByteBuffer with no data.
// This is a singleton - use Empty to access it.
type emptyBuffer struct{}

// Empty is a shared instance of an empty buffer.
var Empty ByteBuffer = &emptyBuffer{}

// IsMemory returns true (empty buffer is conceptually in memory).
func (e *emptyBuffer) IsMemory() bool {
	return true
}

// Size returns 0 (empty buffer has no data).
func (e *emptyBuffer) Size() uint32 {
	return 0
}

// Data returns an empty byte slice.
func (e *emptyBuffer) Data() []byte {
	return []byte{}
}

// GetByteRange returns an error if offset or count is non-zero.
func (e *emptyBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if offset != 0 || count != 0 {
		return fmt.Errorf("offset and count must be 0 for empty buffer, got offset=%d count=%d", offset, count)
	}
	return nil
}

// WriteTo writes nothing to the writer (returns 0 bytes written).
func (e *emptyBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}
	return 0, nil
}
