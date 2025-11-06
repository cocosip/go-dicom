// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"bytes"
	"testing"
)

func TestNewBulkDataUri(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	if buf == nil {
		t.Fatal("NewBulkDataUri returned nil")
	}

	if buf.BulkDataUri() != uri {
		t.Errorf("BulkDataUri() = %q, want %q", buf.BulkDataUri(), uri)
	}

	if buf.IsMemory() {
		t.Error("IsMemory() = true, want false for unloaded buffer")
	}

	if buf.Size() != 0 {
		t.Errorf("Size() = %d, want 0 for unloaded buffer", buf.Size())
	}

	if !buf.IsBulkDataUri() {
		t.Error("IsBulkDataUri() = false, want true")
	}
}

func TestNewBulkDataUriWithData(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	buf := NewBulkDataUriWithData(uri, data)

	if buf == nil {
		t.Fatal("NewBulkDataUriWithData returned nil")
	}

	if buf.BulkDataUri() != uri {
		t.Errorf("BulkDataUri() = %q, want %q", buf.BulkDataUri(), uri)
	}

	if !buf.IsMemory() {
		t.Error("IsMemory() = false, want true for loaded buffer")
	}

	if buf.Size() != 5 {
		t.Errorf("Size() = %d, want 5", buf.Size())
	}

	retrievedData := buf.Data()
	if !bytes.Equal(retrievedData, data) {
		t.Errorf("Data() = %v, want %v", retrievedData, data)
	}
}

func TestBulkDataUriByteBuffer_SetData(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	// Initially no data
	if buf.IsMemory() {
		t.Error("IsMemory() = true before SetData, want false")
	}

	// Set data
	data := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	buf.SetData(data)

	if !buf.IsMemory() {
		t.Error("IsMemory() = false after SetData, want true")
	}

	if buf.Size() != 4 {
		t.Errorf("Size() = %d after SetData, want 4", buf.Size())
	}

	retrievedData := buf.Data()
	if !bytes.Equal(retrievedData, data) {
		t.Errorf("Data() = %v after SetData, want %v", retrievedData, data)
	}
}

func TestBulkDataUriByteBuffer_DataBeforeSet(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	// Data() should return empty slice before SetData
	data := buf.Data()
	if data == nil {
		t.Error("Data() returned nil, want empty slice")
	}
	if len(data) != 0 {
		t.Errorf("Data() length = %d, want 0", len(data))
	}
}

func TestBulkDataUriByteBuffer_GetByteRange(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
	buf := NewBulkDataUriWithData(uri, data)

	tests := []struct {
		name   string
		offset uint32
		count  uint32
		want   []byte
		errMsg string
	}{
		{
			name:   "read first 3 bytes",
			offset: 0,
			count:  3,
			want:   []byte{0x00, 0x01, 0x02},
		},
		{
			name:   "read middle bytes",
			offset: 3,
			count:  4,
			want:   []byte{0x03, 0x04, 0x05, 0x06},
		},
		{
			name:   "read last byte",
			offset: 9,
			count:  1,
			want:   []byte{0x09},
		},
		{
			name:   "read all bytes",
			offset: 0,
			count:  10,
			want:   data,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := buf.GetByteRange(tt.offset, tt.count, output)
			if err != nil {
				t.Errorf("GetByteRange() error = %v", err)
				return
			}
			if !bytes.Equal(output, tt.want) {
				t.Errorf("GetByteRange() = %v, want %v", output, tt.want)
			}
		})
	}
}

func TestBulkDataUriByteBuffer_GetByteRangeErrors(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	buf := NewBulkDataUriWithData(uri, data)

	tests := []struct {
		name   string
		offset uint32
		count  uint32
		outLen int
		errMsg string
	}{
		{
			name:   "output too small",
			offset: 0,
			count:  3,
			outLen: 2,
			errMsg: "output slice with 2 bytes cannot fit 3 bytes of data",
		},
		{
			name:   "range exceeds buffer",
			offset: 3,
			count:  5,
			outLen: 5,
			errMsg: "range [3:8] exceeds buffer size 5",
		},
		{
			name:   "offset beyond buffer",
			offset: 10,
			count:  1,
			outLen: 1,
			errMsg: "range [10:11] exceeds buffer size 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.outLen)
			err := buf.GetByteRange(tt.offset, tt.count, output)
			if err == nil {
				t.Error("GetByteRange() expected error, got nil")
				return
			}
			if err.Error() != tt.errMsg {
				t.Errorf("GetByteRange() error = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestBulkDataUriByteBuffer_GetByteRangeBeforeSetData(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	output := make([]byte, 3)
	err := buf.GetByteRange(0, 3, output)
	if err == nil {
		t.Error("GetByteRange() expected error before SetData, got nil")
		return
	}
	expected := "BulkDataUriByteBuffer cannot provide data until SetData() has been called"
	if err.Error() != expected {
		t.Errorf("GetByteRange() error = %v, want %v", err.Error(), expected)
	}
}

func TestBulkDataUriByteBuffer_WriteTo(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	data := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE}
	buf := NewBulkDataUriWithData(uri, data)

	var output bytes.Buffer
	n, err := buf.WriteTo(&output)
	if err != nil {
		t.Errorf("WriteTo() error = %v", err)
		return
	}

	if n != int64(len(data)) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", n, len(data))
	}

	if !bytes.Equal(output.Bytes(), data) {
		t.Errorf("WriteTo() wrote %v, want %v", output.Bytes(), data)
	}
}

func TestBulkDataUriByteBuffer_WriteToBeforeSetData(t *testing.T) {
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	var output bytes.Buffer
	n, err := buf.WriteTo(&output)
	if err == nil {
		t.Error("WriteTo() expected error before SetData, got nil")
		return
	}

	if n != 0 {
		t.Errorf("WriteTo() wrote %d bytes before SetData, want 0", n)
	}

	expected := "BulkDataUriByteBuffer cannot provide data until SetData() has been called"
	if err.Error() != expected {
		t.Errorf("WriteTo() error = %v, want %v", err.Error(), expected)
	}
}

func TestBulkDataUriByteBuffer_IsBulkDataUri(t *testing.T) {
	uri := "http://example.com/data/pixeldata"

	// Test without data
	buf1 := NewBulkDataUri(uri)
	if !buf1.IsBulkDataUri() {
		t.Error("IsBulkDataUri() = false for unloaded buffer, want true")
	}

	// Test with data
	data := []byte{0x01, 0x02, 0x03}
	buf2 := NewBulkDataUriWithData(uri, data)
	if !buf2.IsBulkDataUri() {
		t.Error("IsBulkDataUri() = false for loaded buffer, want true")
	}

	// Test after SetData
	buf3 := NewBulkDataUri(uri)
	buf3.SetData(data)
	if !buf3.IsBulkDataUri() {
		t.Error("IsBulkDataUri() = false after SetData, want true")
	}
}

func TestBulkDataUriByteBuffer_LazyLoading(t *testing.T) {
	// This test simulates the lazy loading pattern
	uri := "http://example.com/data/pixeldata"
	buf := NewBulkDataUri(uri)

	// Initially no data
	if buf.IsMemory() {
		t.Error("Buffer should not have data initially")
	}

	// Simulate fetching data from URI
	fetchedData := []byte{0x01, 0x02, 0x03, 0x04}
	buf.SetData(fetchedData)

	// Now data should be available
	if !buf.IsMemory() {
		t.Error("Buffer should have data after SetData")
	}

	if buf.Size() != 4 {
		t.Errorf("Size() = %d, want 4", buf.Size())
	}

	// Verify data integrity
	if !bytes.Equal(buf.Data(), fetchedData) {
		t.Errorf("Data() = %v, want %v", buf.Data(), fetchedData)
	}
}
