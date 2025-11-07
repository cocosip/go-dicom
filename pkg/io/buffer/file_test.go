// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// createTempFile creates a temporary file with the given content for testing
func createTempFile(t *testing.T, content []byte) string {
	t.Helper()

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.dat")

	err := os.WriteFile(tmpFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	return tmpFile
}

func TestNewFile(t *testing.T) {
	// Create a test file with 100 bytes
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}
	tmpFile := createTempFile(t, testData)

	tests := []struct {
		name      string
		position  uint32
		size      uint32
		wantError bool
	}{
		{"valid range at start", 0, 50, false},
		{"valid range in middle", 25, 50, false},
		{"valid range at end", 90, 10, false},
		{"valid zero size", 50, 0, false},
		{"valid full file", 0, 100, false},
		{"invalid position too large", 150, 10, true},
		{"invalid range exceeds file", 90, 20, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb, err := buffer.NewFile(tmpFile, tt.position, tt.size)

			if tt.wantError {
				if err == nil {
					t.Error("NewFile() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("NewFile() unexpected error: %v", err)
			}

			if fb.Size() != tt.size {
				t.Errorf("Size() = %d, want %d", fb.Size(), tt.size)
			}

			if fb.Position() != tt.position {
				t.Errorf("Position() = %d, want %d", fb.Position(), tt.position)
			}
		})
	}
}

func TestNewFile_EmptyPath(t *testing.T) {
	_, err := buffer.NewFile("", 0, 10)
	if err == nil {
		t.Error("NewFile(\"\", ...) expected error, got nil")
	}
}

func TestNewFile_NonExistentFile(t *testing.T) {
	_, err := buffer.NewFile("nonexistent_file_xyz123.dat", 0, 10)
	if err == nil {
		t.Error("NewFile() with nonexistent file expected error, got nil")
	}
}

func TestFileByteBuffer_IsMemory(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 0, 5)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	if fb.IsMemory() {
		t.Error("IsMemory() = true, want false for file-backed buffer")
	}
}

func TestFileByteBuffer_Data(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}
	tmpFile := createTempFile(t, testData)

	tests := []struct {
		name     string
		position uint32
		size     uint32
		wantData []byte
	}{
		{"start", 0, 10, testData[0:10]},
		{"middle", 40, 20, testData[40:60]},
		{"end", 90, 10, testData[90:100]},
		{"single byte", 50, 1, testData[50:51]},
		{"empty", 50, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb, err := buffer.NewFile(tmpFile, tt.position, tt.size)
			if err != nil {
				t.Fatalf("NewFile() unexpected error: %v", err)
			}

			data := fb.Data()
			if !bytes.Equal(data, tt.wantData) {
				t.Errorf("Data() = %v, want %v", data, tt.wantData)
			}
		})
	}
}

func TestFileByteBuffer_GetByteRange(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}
	tmpFile := createTempFile(t, testData)

	// Create a file buffer representing bytes [20:70) = 50 bytes
	fb, err := buffer.NewFile(tmpFile, 20, 50)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

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
			count:    50,
			wantData: testData[20:70],
		},
		{
			name:     "partial from start",
			offset:   0,
			count:    10,
			wantData: testData[20:30],
		},
		{
			name:     "partial from middle",
			offset:   10,
			count:    20,
			wantData: testData[30:50],
		},
		{
			name:     "partial at end",
			offset:   45,
			count:    5,
			wantData: testData[65:70],
		},
		{
			name:      "exceeds range",
			offset:    40,
			count:     20,
			wantError: true,
		},
		{
			name:      "offset too large",
			offset:    60,
			count:     5,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := fb.GetByteRange(tt.offset, tt.count, output)

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

func TestFileByteBuffer_GetByteRange_Errors(t *testing.T) {
	testData := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 2, 6)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	t.Run("nil output", func(t *testing.T) {
		err := fb.GetByteRange(0, 3, nil)
		if err == nil {
			t.Error("GetByteRange(nil) expected error, got nil")
		}
	})

	t.Run("output too small", func(t *testing.T) {
		output := make([]byte, 2)
		err := fb.GetByteRange(0, 5, output)
		if err == nil {
			t.Error("GetByteRange() with small output expected error, got nil")
		}
	})
}

func TestFileByteBuffer_WriteTo(t *testing.T) {
	testData := make([]byte, 100)
	for i := range testData {
		testData[i] = byte(i)
	}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 30, 40)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := fb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 40 {
		t.Errorf("WriteTo() wrote %d bytes, want 40", n)
	}

	want := testData[30:70]
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteTo() wrote %v, want %v", buf.Bytes(), want)
	}
}

func TestFileByteBuffer_WriteTo_Empty(t *testing.T) {
	testData := []byte{0, 1, 2, 3, 4, 5}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 3, 0)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := fb.WriteTo(&buf)
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

func TestFileByteBuffer_WriteTo_NilWriter(t *testing.T) {
	testData := []byte{0, 1, 2, 3, 4, 5}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 0, 3)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	_, err = fb.WriteTo(nil)
	if err == nil {
		t.Error("WriteTo(nil) expected error, got nil")
	}
}

func TestFileByteBuffer_WriteTo_LargeData(t *testing.T) {
	// Create a large file (3MB) to test chunked writing
	largeData := make([]byte, 3*1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}
	tmpFile := createTempFile(t, largeData)

	fb, err := buffer.NewFile(tmpFile, 0, uint32(len(largeData)))
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := fb.WriteTo(&buf)
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

func TestFileByteBuffer_FilePath(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5}
	tmpFile := createTempFile(t, testData)

	fb, err := buffer.NewFile(tmpFile, 0, 5)
	if err != nil {
		t.Fatalf("NewFile() unexpected error: %v", err)
	}

	if fb.FilePath() != tmpFile {
		t.Errorf("FilePath() = %q, want %q", fb.FilePath(), tmpFile)
	}
}

func TestFileByteBuffer_MultipleRangesFromSameFile(t *testing.T) {
	// Test creating multiple buffers from different ranges of the same file
	testData := make([]byte, 1000)
	for i := range testData {
		testData[i] = byte(i % 256)
	}
	tmpFile := createTempFile(t, testData)

	// Create three different buffers from the same file
	fb1, err := buffer.NewFile(tmpFile, 0, 100)
	if err != nil {
		t.Fatalf("NewFile() #1 unexpected error: %v", err)
	}

	fb2, err := buffer.NewFile(tmpFile, 100, 200)
	if err != nil {
		t.Fatalf("NewFile() #2 unexpected error: %v", err)
	}

	fb3, err := buffer.NewFile(tmpFile, 800, 200)
	if err != nil {
		t.Fatalf("NewFile() #3 unexpected error: %v", err)
	}

	// Verify each buffer reads the correct data
	data1 := fb1.Data()
	data2 := fb2.Data()
	data3 := fb3.Data()

	if !bytes.Equal(data1, testData[0:100]) {
		t.Error("Buffer #1 data mismatch")
	}

	if !bytes.Equal(data2, testData[100:300]) {
		t.Error("Buffer #2 data mismatch")
	}

	if !bytes.Equal(data3, testData[800:1000]) {
		t.Error("Buffer #3 data mismatch")
	}
}
