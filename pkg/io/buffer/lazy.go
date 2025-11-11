// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
	"sync"
)

// LazyByteBuffer is a buffer that delays loading data until it's actually needed.
// The data is loaded once and cached for subsequent accesses.
// This is useful when you want to defer expensive data loading operations.
type LazyByteBuffer struct {
	loader func() []byte // Function that loads the data
	data   []byte        // Cached data after first load
	once   sync.Once     // Ensures loader is called only once
	mu     sync.RWMutex  // Protects data access
}

// NewLazy creates a new lazy-loading buffer.
//
// Parameters:
//   - loader: A function that returns the byte data when called.
//     This function will be called at most once, when the data is first needed.
//
// Returns an error if loader is nil.
func NewLazy(loader func() []byte) (*LazyByteBuffer, error) {
	if loader == nil {
		return nil, fmt.Errorf("loader function cannot be nil")
	}

	return &LazyByteBuffer{
		loader: loader,
	}, nil
}

// load ensures the data is loaded exactly once.
func (l *LazyByteBuffer) load() []byte {
	l.once.Do(func() {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.data = l.loader()
	})

	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.data
}

// IsMemory returns true since lazy buffers ultimately load data into memory.
func (l *LazyByteBuffer) IsMemory() bool {
	return true
}

// Size returns the size of the buffer.
// Note: This will trigger data loading if not already loaded.
func (l *LazyByteBuffer) Size() uint32 {
	data := l.load()
	if data == nil {
		return 0
	}
	return uint32(len(data)) // #nosec G115 -- DICOM buffer size within uint32 range
}

// Data returns all the buffer data.
// Note: This will trigger data loading if not already loaded.
func (l *LazyByteBuffer) Data() []byte {
	return l.load()
}

// GetByteRange reads a range of bytes from the buffer.
//
// Parameters:
//   - offset: Offset from the start of the buffer
//   - count: Number of bytes to read
//   - output: Destination buffer (must be at least 'count' bytes long)
//
// Returns an error if the operation fails.
// Note: This will trigger data loading if not already loaded.
func (l *LazyByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}

	if uint32(len(output)) < count { // #nosec G115 -- buffer size check
		return fmt.Errorf("output buffer with %d bytes cannot fit %d bytes of data", len(output), count)
	}

	data := l.load()
	if data == nil {
		return fmt.Errorf("lazy loader returned nil data")
	}

	if offset+count > uint32(len(data)) { // #nosec G115 -- buffer size check
		return fmt.Errorf("range [%d:%d) exceeds buffer size %d", offset, offset+count, len(data))
	}

	copy(output[:count], data[offset:offset+count])
	return nil
}

// WriteTo writes the buffer contents to the given writer.
//
// Returns the number of bytes written and any error encountered.
// Note: This will trigger data loading if not already loaded.
func (l *LazyByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	data := l.load()
	if data == nil {
		return 0, nil
	}

	n, err := w.Write(data)
	return int64(n), err
}

// IsLoaded returns true if the data has been loaded.
// This is useful for checking if the lazy loading has been triggered.
func (l *LazyByteBuffer) IsLoaded() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.data != nil
}
