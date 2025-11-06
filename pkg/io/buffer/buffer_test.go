// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestMemoryByteBuffer_Basic(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	buf := buffer.NewMemory(data)

	if !buf.IsMemory() {
		t.Error("MemoryByteBuffer should report IsMemory() = true")
	}

	if buf.Size() != 5 {
		t.Errorf("Size() = %d, want 5", buf.Size())
	}

	gotData := buf.Data()
	if len(gotData) != 5 {
		t.Errorf("Data() length = %d, want 5", len(gotData))
	}
	for i := 0; i < 5; i++ {
		if gotData[i] != byte(i+1) {
			t.Errorf("Data()[%d] = %d, want %d", i, gotData[i], i+1)
		}
	}
}

func TestMemoryByteBuffer_NewWithSize(t *testing.T) {
	buf := buffer.NewMemoryWithSize(10)

	if buf.Size() != 10 {
		t.Errorf("Size() = %d, want 10", buf.Size())
	}

	data := buf.Data()
	if len(data) != 10 {
		t.Errorf("Data() length = %d, want 10", len(data))
	}

	// Should be all zeros
	for i, b := range data {
		if b != 0 {
			t.Errorf("Data()[%d] = %d, want 0", i, b)
		}
	}
}

func TestMemoryByteBuffer_GetByteRange(t *testing.T) {
	data := []byte{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	buf := buffer.NewMemory(data)

	tests := []struct {
		name   string
		offset uint32
		count  uint32
		want   []byte
	}{
		{"Full range", 0, 10, []byte{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
		{"First 5", 0, 5, []byte{10, 20, 30, 40, 50}},
		{"Last 5", 5, 5, []byte{60, 70, 80, 90, 100}},
		{"Middle 3", 3, 3, []byte{40, 50, 60}},
		{"Single byte", 7, 1, []byte{80}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := buf.GetByteRange(tt.offset, tt.count, output)
			if err != nil {
				t.Fatalf("GetByteRange() error = %v", err)
			}

			if !bytes.Equal(output, tt.want) {
				t.Errorf("GetByteRange() = %v, want %v", output, tt.want)
			}
		})
	}
}

func TestMemoryByteBuffer_GetByteRange_Errors(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	buf := buffer.NewMemory(data)

	t.Run("Nil output", func(t *testing.T) {
		err := buf.GetByteRange(0, 2, nil)
		if err == nil {
			t.Error("Expected error for nil output")
		}
	})

	t.Run("Output too small", func(t *testing.T) {
		output := make([]byte, 2)
		err := buf.GetByteRange(0, 5, output)
		if err == nil {
			t.Error("Expected error for output too small")
		}
	})

	t.Run("Offset + count exceeds size", func(t *testing.T) {
		output := make([]byte, 10)
		err := buf.GetByteRange(3, 5, output)
		if err == nil {
			t.Error("Expected error for offset + count > size")
		}
	})
}

func TestMemoryByteBuffer_WriteTo(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	buf := buffer.NewMemory(data)

	var w bytes.Buffer
	n, err := buf.WriteTo(&w)
	if err != nil {
		t.Fatalf("WriteTo() error = %v", err)
	}

	if n != 5 {
		t.Errorf("WriteTo() bytes written = %d, want 5", n)
	}

	written := w.Bytes()
	if !bytes.Equal(written, data) {
		t.Errorf("WriteTo() wrote %v, want %v", written, data)
	}
}

func TestEmptyBuffer(t *testing.T) {
	buf := buffer.Empty

	if !buf.IsMemory() {
		t.Error("Empty buffer should report IsMemory() = true")
	}

	if buf.Size() != 0 {
		t.Errorf("Size() = %d, want 0", buf.Size())
	}

	data := buf.Data()
	if len(data) != 0 {
		t.Errorf("Data() length = %d, want 0", len(data))
	}
}

func TestEmptyBuffer_GetByteRange(t *testing.T) {
	buf := buffer.Empty

	t.Run("Zero offset and count", func(t *testing.T) {
		output := make([]byte, 0)
		err := buf.GetByteRange(0, 0, output)
		if err != nil {
			t.Errorf("GetByteRange(0, 0) should not error, got %v", err)
		}
	})

	t.Run("Non-zero offset", func(t *testing.T) {
		output := make([]byte, 1)
		err := buf.GetByteRange(1, 0, output)
		if err == nil {
			t.Error("GetByteRange with non-zero offset should error")
		}
	})

	t.Run("Non-zero count", func(t *testing.T) {
		output := make([]byte, 1)
		err := buf.GetByteRange(0, 1, output)
		if err == nil {
			t.Error("GetByteRange with non-zero count should error")
		}
	})
}

func TestEmptyBuffer_WriteTo(t *testing.T) {
	buf := buffer.Empty

	var w bytes.Buffer
	n, err := buf.WriteTo(&w)
	if err != nil {
		t.Fatalf("WriteTo() error = %v", err)
	}

	if n != 0 {
		t.Errorf("WriteTo() bytes written = %d, want 0", n)
	}

	if w.Len() != 0 {
		t.Errorf("WriteTo() wrote %d bytes, want 0", w.Len())
	}
}
