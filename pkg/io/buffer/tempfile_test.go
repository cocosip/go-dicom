// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestNewTempFile(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	if tb.Size() != uint32(len(testData)) {
		t.Errorf("Size() = %d, want %d", tb.Size(), len(testData))
	}

	if tb.FilePath() == "" {
		t.Error("FilePath() returned empty string")
	}

	// Verify file exists
	if _, err := os.Stat(tb.FilePath()); os.IsNotExist(err) {
		t.Error("Temporary file does not exist")
	}
}

func TestNewTempFile_EmptyData(t *testing.T) {
	tb, err := buffer.NewTempFile([]byte{})
	if err != nil {
		t.Fatalf("NewTempFile() with empty data unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	if tb.Size() != 0 {
		t.Errorf("Size() = %d, want 0", tb.Size())
	}
}

func TestNewTempFile_NilData(t *testing.T) {
	tb, err := buffer.NewTempFile(nil)
	if err != nil {
		t.Fatalf("NewTempFile() with nil data unexpected error: %v", err)
	}
            defer func() { _ = tb.Close() }()

	if tb.Size() != 0 {
		t.Errorf("Size() = %d, want 0", tb.Size())
	}
}

func TestTempFileBuffer_IsMemory(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5}
	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	if tb.IsMemory() {
		t.Error("IsMemory() = true, want false for temp file buffer")
	}
}

func TestTempFileBuffer_Data(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{"small data", []byte{1, 2, 3, 4, 5}},
		{"larger data", make([]byte, 1000)},
		{"empty", []byte{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize test data
			if len(tt.data) > 0 && tt.data[0] == 0 && len(tt.data) > 1 {
				for i := range tt.data {
					tt.data[i] = byte(i % 256)
				}
			}

			tb, err := buffer.NewTempFile(tt.data)
			if err != nil {
				t.Fatalf("NewTempFile() unexpected error: %v", err)
			}
    defer func() { _ = tb.Close() }()

			data := tb.Data()
			if !bytes.Equal(data, tt.data) {
				t.Errorf("Data() = %v, want %v", data, tt.data)
			}
		})
	}
}

func TestTempFileBuffer_GetByteRange(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}

	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	tests := []struct {
		name      string
		offset    uint32
		count     uint32
		wantData  []byte
		wantError bool
	}{
		{
			name:     "full range",
			offset:   0,
			count:    100,
			wantData: testData,
		},
		{
			name:     "partial from start",
			offset:   0,
			count:    10,
			wantData: testData[0:10],
		},
		{
			name:     "partial from middle",
			offset:   40,
			count:    20,
			wantData: testData[40:60],
		},
		{
			name:     "partial at end",
			offset:   95,
			count:    5,
			wantData: testData[95:100],
		},
		{
			name:      "exceeds range",
			offset:    90,
			count:     20,
			wantError: true,
		},
		{
			name:      "offset too large",
			offset:    150,
			count:     5,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := tb.GetByteRange(tt.offset, tt.count, output)

			if tt.wantError {
				if err == nil {
					t.Error("GetByteRange() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetByteRange() unexpected error: %v", err)
			}

			if !bytes.Equal(output, tt.wantData) {
				t.Errorf("GetByteRange() = %v, want %v", output, tt.wantData)
			}
		})
	}
}

func TestTempFileBuffer_GetByteRange_Errors(t *testing.T) {
	testData := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	t.Run("nil output", func(t *testing.T) {
		err := tb.GetByteRange(0, 3, nil)
		if err == nil {
			t.Error("GetByteRange(nil) expected error, got nil")
		}
	})

	t.Run("output too small", func(t *testing.T) {
		output := make([]byte, 2)
		err := tb.GetByteRange(0, 5, output)
		if err == nil {
			t.Error("GetByteRange() with small output expected error, got nil")
		}
	})
}

func TestTempFileBuffer_WriteTo(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}

	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	var buf bytes.Buffer
	n, err := tb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != int64(len(testData)) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", n, len(testData))
	}

	if !bytes.Equal(buf.Bytes(), testData) {
		t.Error("WriteTo() data mismatch")
	}
}

func TestTempFileBuffer_WriteTo_Empty(t *testing.T) {
	tb, err := buffer.NewTempFile([]byte{})
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	var buf bytes.Buffer
	n, err := tb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 0 {
		t.Errorf("WriteTo() wrote %d bytes, want 0", n)
	}

	if buf.Len() != 0 {
		t.Errorf("WriteTo() buffer has %d bytes, want 0", buf.Len())
	}
}

func TestTempFileBuffer_WriteTo_NilWriter(t *testing.T) {
	testData := []byte{0, 1, 2, 3, 4, 5}
	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	_, err = tb.WriteTo(nil)
	if err == nil {
		t.Error("WriteTo(nil) expected error, got nil")
	}
}

func TestTempFileBuffer_WriteTo_LargeData(t *testing.T) {
	// Create a large data set (2MB)
	largeData := make([]byte, 2*1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	tb, err := buffer.NewTempFile(largeData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	var buf bytes.Buffer
	n, err := tb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != int64(len(largeData)) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", n, len(largeData))
	}

	if !bytes.Equal(buf.Bytes(), largeData) {
		t.Error("WriteTo() data mismatch")
	}
}

func TestTempFileBuffer_Close(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5}
	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}

	filePath := tb.FilePath()

	// Verify file exists before close
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatal("Temporary file does not exist before Close()")
	}

	// Close the buffer
	err = tb.Close()
	if err != nil {
		t.Fatalf("Close() unexpected error: %v", err)
	}

	// Verify file is removed after close
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Error("Temporary file still exists after Close()")
	}
}

func TestTempFileBuffer_MultipleReads(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}

	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}
    defer func() { _ = tb.Close() }()

	// First read
	data1 := tb.Data()
	if !bytes.Equal(data1, testData) {
		t.Error("First Data() call returned incorrect data")
	}

	// Second read (should work since we seek to beginning)
	data2 := tb.Data()
	if !bytes.Equal(data2, testData) {
		t.Error("Second Data() call returned incorrect data")
	}

	// Read with GetByteRange
	output := make([]byte, 10)
	err = tb.GetByteRange(0, 10, output)
	if err != nil {
		t.Fatalf("GetByteRange() unexpected error: %v", err)
	}
	if !bytes.Equal(output, testData[0:10]) {
		t.Error("GetByteRange() after Data() calls returned incorrect data")
	}
}

func TestTempFileBuffer_CleanupOnError(t *testing.T) {
	// This test ensures that if Close() is called multiple times,
	// it doesn't cause a panic or unexpected error
	testData := []byte{1, 2, 3, 4, 5}
	tb, err := buffer.NewTempFile(testData)
	if err != nil {
		t.Fatalf("NewTempFile() unexpected error: %v", err)
	}

	// First close should succeed
	err = tb.Close()
	if err != nil {
		t.Errorf("First Close() unexpected error: %v", err)
	}

	// Second close should not panic (may return error since file is already closed)
	_ = tb.Close() // Ignore error as file is already closed
}
