// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestNewSwap(t *testing.T) {
	base := buffer.NewMemory([]byte{1, 2, 3, 4})

	tests := []struct {
		name      string
		buffer    buffer.ByteBuffer
		unitSize  int
		wantError bool
	}{
		{"valid unitSize 2", base, 2, false},
		{"valid unitSize 4", base, 4, false},
		{"valid unitSize 1", base, 1, false},
		{"nil buffer", nil, 2, true},
		{"invalid unitSize 0", base, 0, true},
		{"invalid unitSize -1", base, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb, err := buffer.NewSwap(tt.buffer, tt.unitSize)

			if tt.wantError {
				if err == nil {
					t.Error("NewSwap() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("NewSwap() unexpected error: %v", err)
			}

			if sb.UnitSize() != tt.unitSize {
				t.Errorf("UnitSize() = %d, want %d", sb.UnitSize(), tt.unitSize)
			}
		})
	}
}

func TestSwapByteBuffer_IsMemory(t *testing.T) {
	base := buffer.NewMemory([]byte{1, 2, 3, 4})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	if !sb.IsMemory() {
		t.Error("IsMemory() = false, want true for memory-backed buffer")
	}
}

func TestSwapByteBuffer_Size(t *testing.T) {
	base := buffer.NewMemory([]byte{1, 2, 3, 4, 5, 6, 7, 8})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	if sb.Size() != 8 {
		t.Errorf("Size() = %d, want 8", sb.Size())
	}
}

func TestSwapByteBuffer_Data_UnitSize2(t *testing.T) {
	// Test swapping in groups of 2 bytes
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	data := sb.Data()
	// Each pair should be swapped
	want := []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05, 0x08, 0x07}

	if !bytes.Equal(data, want) {
		t.Errorf("Data() = %v, want %v", data, want)
	}

	// Verify original data is not modified
	originalData := base.Data()
	originalWant := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	if !bytes.Equal(originalData, originalWant) {
		t.Error("Original buffer was modified")
	}
}

func TestSwapByteBuffer_Data_UnitSize4(t *testing.T) {
	// Test swapping in groups of 4 bytes
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	sb, err := buffer.NewSwap(base, 4)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	data := sb.Data()
	// Each group of 4 should be reversed
	want := []byte{0x04, 0x03, 0x02, 0x01, 0x08, 0x07, 0x06, 0x05}

	if !bytes.Equal(data, want) {
		t.Errorf("Data() = %v, want %v", data, want)
	}
}

func TestSwapByteBuffer_Data_UnitSize1(t *testing.T) {
	// Test with unitSize 1 (no swapping should occur)
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04})
	sb, err := buffer.NewSwap(base, 1)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	data := sb.Data()
	want := []byte{0x01, 0x02, 0x03, 0x04}

	if !bytes.Equal(data, want) {
		t.Errorf("Data() = %v, want %v (no swap should occur)", data, want)
	}
}

func TestSwapByteBuffer_Data_OddLength(t *testing.T) {
	// Test with odd length data (last byte shouldn't be swapped)
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	data := sb.Data()
	// Pairs are swapped, last byte remains
	want := []byte{0x02, 0x01, 0x04, 0x03, 0x05}

	if !bytes.Equal(data, want) {
		t.Errorf("Data() = %v, want %v", data, want)
	}
}

func TestSwapByteBuffer_GetByteRange(t *testing.T) {
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
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
			count:    8,
			wantData: []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05, 0x08, 0x07},
		},
		{
			name:     "first 4 bytes",
			offset:   0,
			count:    4,
			wantData: []byte{0x02, 0x01, 0x04, 0x03},
		},
		{
			name:     "middle 4 bytes",
			offset:   2,
			count:    4,
			wantData: []byte{0x04, 0x03, 0x06, 0x05},
		},
		{
			name:     "last 2 bytes",
			offset:   6,
			count:    2,
			wantData: []byte{0x08, 0x07},
		},
		{
			name:      "exceeds buffer",
			offset:    6,
			count:     10,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := sb.GetByteRange(tt.offset, tt.count, output)

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

func TestSwapByteBuffer_GetByteRange_Errors(t *testing.T) {
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	t.Run("nil output", func(t *testing.T) {
		err := sb.GetByteRange(0, 2, nil)
		if err == nil {
			t.Error("GetByteRange(nil) expected error, got nil")
		}
	})

	t.Run("output too small", func(t *testing.T) {
		output := make([]byte, 1)
		err := sb.GetByteRange(0, 2, output)
		if err == nil {
			t.Error("GetByteRange() with small output expected error, got nil")
		}
	})
}

func TestSwapByteBuffer_WriteTo(t *testing.T) {
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := sb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 8 {
		t.Errorf("WriteTo() wrote %d bytes, want 8", n)
	}

	want := []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05, 0x08, 0x07}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteTo() wrote %v, want %v", buf.Bytes(), want)
	}
}

func TestSwapByteBuffer_WriteTo_LargeData(t *testing.T) {
	// Create 2MB of test data
	largeData := make([]byte, 2*1024*1024)
	for i := 0; i < len(largeData); i += 2 {
		largeData[i] = byte(i % 256)
		if i+1 < len(largeData) {
			largeData[i+1] = byte((i + 1) % 256)
		}
	}

	base := buffer.NewMemory(largeData)
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := sb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != int64(len(largeData)) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", n, len(largeData))
	}

	// Verify data is correctly swapped
	result := buf.Bytes()
	for i := 0; i < len(largeData); i += 2 {
		if i+1 < len(largeData) {
			// Each pair should be swapped
			if result[i] != largeData[i+1] || result[i+1] != largeData[i] {
				t.Errorf("At position %d: got [%d,%d], want [%d,%d]",
					i, result[i], result[i+1], largeData[i+1], largeData[i])
				break
			}
		}
	}
}

func TestSwapByteBuffer_WriteTo_NilWriter(t *testing.T) {
	base := buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	_, err = sb.WriteTo(nil)
	if err == nil {
		t.Error("WriteTo(nil) expected error, got nil")
	}
}

func TestSwapByteBuffer_WriteTo_Empty(t *testing.T) {
	base := buffer.NewMemory([]byte{})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := sb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 0 {
		t.Errorf("WriteTo() wrote %d bytes, want 0", n)
	}
}

func TestSwapByteBuffer_Internal(t *testing.T) {
	base := buffer.NewMemory([]byte{1, 2, 3, 4})
	sb, err := buffer.NewSwap(base, 2)
	if err != nil {
		t.Fatalf("NewSwap() unexpected error: %v", err)
	}

	internal := sb.Internal()
	if internal != base {
		t.Error("Internal() did not return the base buffer")
	}
}

func TestSwapByteBuffer_MultipleSwapSizes(t *testing.T) {
	// Test different swap sizes on the same data
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	tests := []struct {
		unitSize int
		want     []byte
	}{
		{1, []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}, // no swap
		{2, []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05, 0x08, 0x07}}, // pair swap
		{4, []byte{0x04, 0x03, 0x02, 0x01, 0x08, 0x07, 0x06, 0x05}}, // quad swap
		{8, []byte{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}}, // full reverse
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("unitSize=%d", tt.unitSize), func(t *testing.T) {
			base := buffer.NewMemory(data)
			sb, err := buffer.NewSwap(base, tt.unitSize)
			if err != nil {
				t.Fatalf("NewSwap() unexpected error: %v", err)
			}

			result := sb.Data()
			if !bytes.Equal(result, tt.want) {
				t.Errorf("Data() = %v, want %v", result, tt.want)
			}
		})
	}
}
