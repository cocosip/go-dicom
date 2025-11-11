// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
	"os"
)

// FileByteBuffer represents a buffer backed by a file on disk.
// It provides efficient access to a specific range within a file without loading
// the entire file into memory.
type FileByteBuffer struct {
	filePath string // Path to the file
	position uint32 // Starting position within the file
	size     uint32 // Size of the data range
}

// NewFile creates a new file-backed buffer.
//
// Parameters:
//   - filePath: Path to the file on disk
//   - position: Starting position within the file (offset)
//   - size: Number of bytes to represent from the file
//
// Returns an error if the file doesn't exist or if the range is invalid.
func NewFile(filePath string, position, size uint32) (*FileByteBuffer, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path cannot be empty")
	}

	// Verify file exists and check bounds
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot access file: %w", err)
	}

    fileSize := uint32(info.Size()) //nolint:gosec // file size for DICOM buffer
	if position > fileSize {
		return nil, fmt.Errorf("position %d exceeds file size %d", position, fileSize)
	}

	if position+size > fileSize {
		return nil, fmt.Errorf("range [%d:%d) exceeds file size %d", position, position+size, fileSize)
	}

	return &FileByteBuffer{
		filePath: filePath,
		position: position,
		size:     size,
	}, nil
}

// IsMemory returns false since this buffer is file-backed, not memory-backed.
func (f *FileByteBuffer) IsMemory() bool {
	return false
}

// Size returns the size of the buffer in bytes.
func (f *FileByteBuffer) Size() uint32 {
	return f.size
}

// Data reads and returns all data from the file range.
// For large files, consider using GetByteRange or WriteTo to avoid loading
// everything into memory at once.
func (f *FileByteBuffer) Data() []byte {
	if f.size == 0 {
		return nil
	}

	file, err := os.Open(f.filePath)
	if err != nil {
		return nil
	}
    defer func() { _ = file.Close() }()

	data := make([]byte, f.size)
	_, err = file.ReadAt(data, int64(f.position))
	if err != nil && err != io.EOF {
		return nil
	}

	return data
}

// GetByteRange reads a sub-range of bytes from the file into the output buffer.
//
// Parameters:
//   - offset: Offset relative to the buffer's start position
//   - count: Number of bytes to read
//   - output: Destination buffer (must be at least 'count' bytes long)
//
// Returns an error if the range is invalid or if reading fails.
func (f *FileByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}

    if uint32(len(output)) < count { //nolint:gosec // buffer size check
		return fmt.Errorf("output buffer with %d bytes cannot fit %d bytes of data", len(output), count)
	}

	if offset+count > f.size {
		return fmt.Errorf("range [%d:%d) exceeds buffer size %d", offset, offset+count, f.size)
	}

	if count == 0 {
		return nil
	}

	file, err := os.Open(f.filePath)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
    defer func() { _ = file.Close() }()

	// Read from the absolute position: buffer position + offset
	_, err = file.ReadAt(output[:count], int64(f.position+offset))
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read from file: %w", err)
	}

	return nil
}

// WriteTo writes the buffer contents to the given writer.
// This uses chunked reading to avoid loading the entire buffer into memory,
// making it efficient for large files.
//
// Returns the number of bytes written and any error encountered.
func (f *FileByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	if f.size == 0 {
		return 0, nil
	}

	file, err := os.Open(f.filePath)
	if err != nil {
		return 0, fmt.Errorf("cannot open file: %w", err)
	}
    defer func() { _ = file.Close() }()

	// Seek to the starting position
	_, err = file.Seek(int64(f.position), io.SeekStart)
	if err != nil {
		return 0, fmt.Errorf("cannot seek to position %d: %w", f.position, err)
	}

	// Use a limited reader to read only 'size' bytes
	limitedReader := io.LimitReader(file, int64(f.size))

	// Copy data in chunks (1MB at a time for efficiency)
	const chunkSize = 1024 * 1024
	buffer := make([]byte, chunkSize)
	var totalWritten int64

	for {
		n, err := limitedReader.Read(buffer)
		if n > 0 {
			written, writeErr := w.Write(buffer[:n])
			totalWritten += int64(written)
			if writeErr != nil {
				return totalWritten, fmt.Errorf("write failed: %w", writeErr)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return totalWritten, fmt.Errorf("read failed: %w", err)
		}
	}

	return totalWritten, nil
}

// FilePath returns the path to the underlying file.
func (f *FileByteBuffer) FilePath() string {
	return f.filePath
}

// Position returns the starting position within the file.
func (f *FileByteBuffer) Position() uint32 {
	return f.position
}
