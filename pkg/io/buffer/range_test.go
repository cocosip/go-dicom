// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestNewRange(t *testing.T) {
	// Create a base buffer with 10 bytes
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	tests := []struct {
		name      string
		offset    uint32
		length    uint32
		wantError bool
	}{
		{"valid range at start", 0, 5, false},
		{"valid range in middle", 3, 4, false},
		{"valid range at end", 8, 2, false},
		{"valid zero length", 5, 0, false},
		{"valid full range", 0, 10, false},
		{"invalid exceeds size", 5, 10, true},
		{"invalid offset too large", 15, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := buffer.NewRange(base, tt.offset, tt.length)

			if tt.wantError {
				if err == nil {
					t.Error("NewRange() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("NewRange() unexpected error: %v", err)
			}

			if r.Size() != tt.length {
				t.Errorf("Size() = %d, want %d", r.Size(), tt.length)
			}

			if r.Offset() != tt.offset {
				t.Errorf("Offset() = %d, want %d", r.Offset(), tt.offset)
			}

			if r.Length() != tt.length {
				t.Errorf("Length() = %d, want %d", r.Length(), tt.length)
			}
		})
	}
}

func TestNewRange_NilBuffer(t *testing.T) {
	_, err := buffer.NewRange(nil, 0, 10)
	if err == nil {
		t.Error("NewRange(nil, ...) expected error, got nil")
	}
}

func TestRangeByteBuffer_Data(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	tests := []struct {
		name     string
		offset   uint32
		length   uint32
		wantData []byte
	}{
		{"start", 0, 3, []byte{0, 1, 2}},
		{"middle", 3, 4, []byte{3, 4, 5, 6}},
		{"end", 7, 3, []byte{7, 8, 9}},
		{"single byte", 5, 1, []byte{5}},
		{"empty", 5, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := buffer.NewRange(base, tt.offset, tt.length)
			if err != nil {
				t.Fatalf("NewRange() unexpected error: %v", err)
			}

			data := r.Data()
			if !bytes.Equal(data, tt.wantData) {
				t.Errorf("Data() = %v, want %v", data, tt.wantData)
			}
		})
	}
}

func TestRangeByteBuffer_GetByteRange(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	r, err := buffer.NewRange(base, 2, 6) // Range [2:8) = {2, 3, 4, 5, 6, 7}
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
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
			count:    6,
			wantData: []byte{2, 3, 4, 5, 6, 7},
		},
		{
			name:     "partial from start",
			offset:   0,
			count:    3,
			wantData: []byte{2, 3, 4},
		},
		{
			name:     "partial from middle",
			offset:   2,
			count:    3,
			wantData: []byte{4, 5, 6},
		},
		{
			name:     "partial at end",
			offset:   4,
			count:    2,
			wantData: []byte{6, 7},
		},
		{
			name:      "exceeds range",
			offset:    3,
			count:     10,
			wantError: true,
		},
		{
			name:      "offset too large",
			offset:    10,
			count:     1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := r.GetByteRange(tt.offset, tt.count, output)

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

func TestRangeByteBuffer_GetByteRange_Errors(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	r, _ := buffer.NewRange(base, 2, 6)

	t.Run("nil output", func(t *testing.T) {
		err := r.GetByteRange(0, 3, nil)
		if err == nil {
			t.Error("GetByteRange(nil) expected error, got nil")
		}
	})

	t.Run("output too small", func(t *testing.T) {
		output := make([]byte, 2)
		err := r.GetByteRange(0, 5, output)
		if err == nil {
			t.Error("GetByteRange() with small output expected error, got nil")
		}
	})
}

func TestRangeByteBuffer_WriteTo(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	r, err := buffer.NewRange(base, 3, 5) // Range {3, 4, 5, 6, 7}
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := r.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 5 {
		t.Errorf("WriteTo() wrote %d bytes, want 5", n)
	}

	want := []byte{3, 4, 5, 6, 7}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteTo() wrote %v, want %v", buf.Bytes(), want)
	}
}

func TestRangeByteBuffer_WriteTo_Empty(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5})
	r, err := buffer.NewRange(base, 3, 0) // Empty range
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := r.WriteTo(&buf)
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

func TestRangeByteBuffer_WriteTo_NilWriter(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5})
	r, err := buffer.NewRange(base, 0, 3)
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	_, err = r.WriteTo(nil)
	if err == nil {
		t.Error("WriteTo(nil) expected error, got nil")
	}
}

func TestRangeByteBuffer_WriteTo_LargeData(t *testing.T) {
	// Create a large buffer (3MB) to test chunked writing
	largeData := make([]byte, 3*1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	base := buffer.NewMemory(largeData)
	r, err := buffer.NewRange(base, 0, uint32(len(largeData)))
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := r.WriteTo(&buf)
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

func TestRangeByteBuffer_IsMemory(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5})
	r, err := buffer.NewRange(base, 1, 3)
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	if !r.IsMemory() {
		t.Error("IsMemory() = false, want true for memory-backed range")
	}
}

func TestRangeByteBuffer_NestedRanges(t *testing.T) {
	// Create nested ranges: base -> range1 -> range2
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	// First range: [2:8) = {2, 3, 4, 5, 6, 7}
	range1, err := buffer.NewRange(base, 2, 6)
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	// Second range from first range: [1:4) relative to range1
	// = [3:6) in base = {3, 4, 5}
	range2, err := buffer.NewRange(range1, 1, 3)
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	want := []byte{3, 4, 5}
	got := range2.Data()
	if !bytes.Equal(got, want) {
		t.Errorf("Nested range Data() = %v, want %v", got, want)
	}
}

func TestRangeByteBuffer_Internal(t *testing.T) {
	base := buffer.NewMemory([]byte{0, 1, 2, 3, 4, 5})
	r, err := buffer.NewRange(base, 1, 3)
	if err != nil {
		t.Fatalf("NewRange() unexpected error: %v", err)
	}

	internal := r.Internal()
	if internal != base {
		t.Error("Internal() did not return the base buffer")
	}
}
