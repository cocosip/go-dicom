// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"fmt"
	"io"
)

// BulkDataURIByteBuffer represents a byte buffer that references bulk data via a URI.
// This is used in DICOM JSON serialization according to PS3.18 Chapter F.2.2.
//
// The actual data may be loaded lazily from the URI, or set explicitly.
type BulkDataURIByteBuffer struct {
	// bulkDataURI is the URI for retrieving the referenced bulk data
	bulkDataURI string

	// data holds the actual bulk data once loaded
	data []byte
}

// NewBulkDataURI creates a new BulkDataUriByteBuffer with the given URI.
// The data is not loaded until it's accessed or explicitly set.
func NewBulkDataURI(uri string) *BulkDataURIByteBuffer {
	return &BulkDataURIByteBuffer{
		bulkDataURI: uri,
		data:        nil,
	}
}

// NewBulkDataURIWithData creates a new BulkDataUriByteBuffer with both URI and data.
func NewBulkDataURIWithData(uri string, data []byte) *BulkDataURIByteBuffer {
	return &BulkDataURIByteBuffer{
		bulkDataURI: uri,
		data:        data,
	}
}

// IsMemory returns true if the data is buffered in memory.
func (b *BulkDataURIByteBuffer) IsMemory() bool {
	return b.data != nil
}

// Size returns the size of the buffered data in bytes.
// Returns an error if the data has not been set yet.
func (b *BulkDataURIByteBuffer) Size() uint32 {
	if b.data == nil {
		// In Go, we can't panic in a method that returns uint32 only
		// So we return 0, but callers should check IsMemory() first
		return 0
	}
	return uint32(len(b.data))
}

// Data returns the buffered data.
// Returns an error if the data has not been set yet.
func (b *BulkDataURIByteBuffer) Data() []byte {
	if b.data == nil {
		// Return empty slice instead of nil to avoid nil pointer errors
		// Callers should check IsMemory() or BulkDataUri() first
		return []byte{}
	}
	return b.data
}

// SetData sets the bulk data explicitly.
// This is typically called after fetching data from the BulkDataUri.
func (b *BulkDataURIByteBuffer) SetData(data []byte) {
	b.data = data
}

// BulkDataURI returns the URI for retrieving the referenced bulk data.
func (b *BulkDataURIByteBuffer) BulkDataURI() string {
	return b.bulkDataURI
}

// GetByteRange copies a range of bytes from the buffer to the output slice.
// Returns an error if the data has not been set yet.
func (b *BulkDataURIByteBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if b.data == nil {
		return fmt.Errorf("BulkDataUriByteBuffer cannot provide data until SetData() has been called")
	}

	if uint32(len(output)) < count {
		return fmt.Errorf("output slice with %d bytes cannot fit %d bytes of data", len(output), count)
	}

	if offset+count > uint32(len(b.data)) {
		return fmt.Errorf("range [%d:%d] exceeds buffer size %d", offset, offset+count, len(b.data))
	}

	copy(output, b.data[offset:offset+count])
	return nil
}

// WriteTo writes the buffered data to the writer.
// Returns an error if the data has not been set yet.
func (b *BulkDataURIByteBuffer) WriteTo(w io.Writer) (int64, error) {
	if b.data == nil {
		return 0, fmt.Errorf("BulkDataUriByteBuffer cannot provide data until SetData() has been called")
	}

	n, err := w.Write(b.data)
	return int64(n), err
}

// IsBulkDataURI returns true, indicating this buffer represents bulk data referenced by URI.
func (b *BulkDataURIByteBuffer) IsBulkDataURI() bool {
	return true
}
