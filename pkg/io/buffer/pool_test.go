// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package buffer

import (
	"bytes"
	"testing"
)

func TestBytesBufferPool(t *testing.T) {
	t.Run("GetAndPut", func(t *testing.T) {
		buf := GetBytesBuffer()
		if buf == nil {
			t.Fatal("GetBytesBuffer returned nil")
		}

		// Write some data
		buf.WriteString("test data")
		if buf.Len() != 9 {
			t.Errorf("buf.Len() = %d, want 9", buf.Len())
		}

		// Return to pool
		PutBytesBuffer(buf)

		// Get another buffer - might be the same one
		buf2 := GetBytesBuffer()
		// Should be reset (empty)
		if buf2.Len() != 0 {
			t.Errorf("buf2.Len() = %d, want 0 (buffer should be reset)", buf2.Len())
		}
		PutBytesBuffer(buf2)
	})

	t.Run("PutNil", func(_ *testing.T) {
		// Should not panic
		PutBytesBuffer(nil)
	})
}

func BenchmarkGetBytesBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := GetBytesBuffer()
		buf.WriteString("test data")
		PutBytesBuffer(buf)
	}
}

func BenchmarkBytesBufferNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		buf.WriteString("test data")
		_ = buf
	}
}
