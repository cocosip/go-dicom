// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
	"os"
)

// TempFileBuffer represents a buffer backed by a temporary file.
// This is useful when data is too large to keep in memory.
// The temporary file is created on initialization and cleaned up when Close() is called.
type TempFileBuffer struct {
	file     *os.File // The temporary file
	size     uint32   // Size of the data
	filePath string   // Path to the temporary file (for reference)
}

// NewTempFile creates a new temporary file buffer and writes the provided data to it.
//
// Parameters:
//   - data: The byte data to write to the temporary file
//
// Returns a TempFileBuffer and any error encountered.
// The caller is responsible for calling Close() to clean up the temporary file.
func NewTempFile(data []byte) (*TempFileBuffer, error) {
	if data == nil {
		data = []byte{}
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "godicom_*.tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}

	// Write data to the temporary file
	n, err := tmpFile.Write(data)
	if err != nil {
        _ = tmpFile.Close()
        _ = os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to write to temporary file: %w", err)
	}

	if n != len(data) {
        _ = tmpFile.Close()
        _ = os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to write all data to temporary file: wrote %d of %d bytes", n, len(data))
	}

	// Sync to ensure data is written to disk
	if err := tmpFile.Sync(); err != nil {
        _ = tmpFile.Close()
        _ = os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to sync temporary file: %w", err)
	}

	return &TempFileBuffer{
		file:     tmpFile,
    size:     uint32(len(data)), //nolint:gosec // DICOM buffer size within uint32 range
		filePath: tmpFile.Name(),
	}, nil
}

// IsMemory returns false since this buffer is file-backed, not memory-backed.
func (t *TempFileBuffer) IsMemory() bool {
	return false
}

// Size returns the size of the buffer in bytes.
func (t *TempFileBuffer) Size() uint32 {
	return t.size
}

// Data reads and returns all data from the temporary file.
// For large files, consider using GetByteRange or WriteTo to avoid loading
// everything into memory at once.
func (t *TempFileBuffer) Data() []byte {
	if t.size == 0 {
		return nil
	}

	data := make([]byte, t.size)

	// Seek to the beginning of the file
	_, err := t.file.Seek(0, io.SeekStart)
	if err != nil {
		return nil
	}

	// Read all data
	n, err := t.file.Read(data)
	if err != nil && err != io.EOF {
		return nil
	}

    if uint32(n) != t.size { //nolint:gosec // bytes read check
		return nil
	}

	return data
}

// GetByteRange reads a sub-range of bytes from the temporary file into the output buffer.
//
// Parameters:
//   - offset: Offset from the start of the buffer
//   - count: Number of bytes to read
//   - output: Destination buffer (must be at least 'count' bytes long)
//
// Returns an error if the range is invalid or if reading fails.
func (t *TempFileBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if output == nil {
		return fmt.Errorf("output buffer cannot be nil")
	}

    if uint32(len(output)) < count { //nolint:gosec // buffer size check
		return fmt.Errorf("output buffer with %d bytes cannot fit %d bytes of data", len(output), count)
	}

	if offset+count > t.size {
		return fmt.Errorf("range [%d:%d) exceeds buffer size %d", offset, offset+count, t.size)
	}

	if count == 0 {
		return nil
	}

	// Seek to the desired offset
	_, err := t.file.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to position %d: %w", offset, err)
	}

	// Read the requested number of bytes
	n, err := t.file.Read(output[:count])
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read from temporary file: %w", err)
	}

    if uint32(n) != count { //nolint:gosec // bytes read check
		return fmt.Errorf("read %d bytes, expected %d", n, count)
	}

	return nil
}

// WriteTo writes the buffer contents to the given writer.
// This efficiently copies data from the temporary file without loading
// everything into memory.
//
// Returns the number of bytes written and any error encountered.
func (t *TempFileBuffer) WriteTo(w io.Writer) (int64, error) {
	if w == nil {
		return 0, fmt.Errorf("writer cannot be nil")
	}

	if t.size == 0 {
		return 0, nil
	}

	// Seek to the beginning of the file
	_, err := t.file.Seek(0, io.SeekStart)
	if err != nil {
		return 0, fmt.Errorf("failed to seek to beginning: %w", err)
	}

	// Use io.CopyN to copy exactly 'size' bytes
	n, err := io.CopyN(w, t.file, int64(t.size))
	if err != nil {
		return n, fmt.Errorf("failed to copy data: %w", err)
	}

	return n, nil
}

// FilePath returns the path to the temporary file.
func (t *TempFileBuffer) FilePath() string {
	return t.filePath
}

// Close closes the temporary file and removes it from the filesystem.
// It's important to call this method to clean up temporary files.
func (t *TempFileBuffer) Close() error {
	if t.file == nil {
		return nil
	}

	// Close the file first
	closeErr := t.file.Close()

	// Then try to remove it
	removeErr := os.Remove(t.filePath)

	// Return the first error encountered
	if closeErr != nil {
		return fmt.Errorf("failed to close temporary file: %w", closeErr)
	}
	if removeErr != nil {
		return fmt.Errorf("failed to remove temporary file: %w", removeErr)
	}

	return nil
}
