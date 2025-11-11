// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
	"sync"
)

// StreamByteBuffer is a ByteBuffer implementation backed by an io.ReadSeeker.
//
// This buffer reads data from a stream (e.g., file) on demand rather than
// loading it all into memory. It's useful for large DICOM files.
type StreamByteBuffer struct {
	stream   io.ReadSeeker
	position int64
	size     uint32
	mu       sync.Mutex // Protects concurrent access to stream
}

// NewStream creates a new StreamByteBuffer from a stream.
// position is the starting position in the stream.
// size is the number of bytes to read from the stream.
func NewStream(stream io.ReadSeeker, position int64, size uint32) (*StreamByteBuffer, error) {
	if stream == nil {
		return nil, fmt.Errorf("stream cannot be nil")
	}
	if position < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	return &StreamByteBuffer{
		stream:   stream,
		position: position,
		size:     size,
	}, nil
}

// IsMemory returns false (this buffer is stream-based).
func (s *StreamByteBuffer) IsMemory() bool {
	return false
}

// Size returns the size of the buffer in bytes.
func (s *StreamByteBuffer) Size() uint32 {
	return s.size
}

// Data reads all data from the stream into memory and returns it.
// This may be expensive for large buffers.
func (s *StreamByteBuffer) Data() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := make([]byte, s.size)
	if s.size == 0 {
		return data
	}

	// Seek to position
	if _, err := s.stream.Seek(s.position, io.SeekStart); err != nil {
		// Return empty data on error (best effort)
		return data
	}

	// Read data
	_, _ = io.ReadFull(s.stream, data)
	return data
}

// GetByteRange reads a range of bytes from the stream into the output slice.
func (s *StreamByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}
    if uint32(len(output)) < count { //nolint:gosec // buffer size check
		return fmt.Errorf("output buffer length %d is less than requested count %d", len(output), count)
	}
	if offset+count > s.size {
		return fmt.Errorf("offset %d + count %d exceeds buffer size %d", offset, count, s.size)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Seek to position + offset
	seekPos := s.position + int64(offset)
	if _, err := s.stream.Seek(seekPos, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to position %d: %w", seekPos, err)
	}

	// Read data
	n, err := io.ReadFull(s.stream, output[:count])
	if err != nil {
		return fmt.Errorf("failed to read %d bytes (read %d): %w", count, n, err)
	}

	return nil
}

// WriteTo copies the buffer contents from the stream to the writer.
func (s *StreamByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Seek to position
	if _, err := s.stream.Seek(s.position, io.SeekStart); err != nil {
		return 0, fmt.Errorf("failed to seek to position %d: %w", s.position, err)
	}

	// Copy limited amount
	n, err := io.CopyN(w, s.stream, int64(s.size))
	return n, err
}
