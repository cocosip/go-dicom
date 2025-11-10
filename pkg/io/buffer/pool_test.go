// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"bytes"
	"testing"
)

func TestBufferPool(t *testing.T) {
	// Reset pool before testing
	ResetPool()

	t.Run("GetSmallBuffer", func(t *testing.T) {
		buf := GetBuffer(1024)
		if len(buf) != 1024 {
			t.Errorf("len(buf) = %d, want 1024", len(buf))
		}
		if cap(buf) < 1024 {
			t.Errorf("cap(buf) = %d, want >= 1024", cap(buf))
		}
		PutBuffer(buf)
	})

	t.Run("GetMediumBuffer", func(t *testing.T) {
		buf := GetBuffer(32768)
		if len(buf) != 32768 {
			t.Errorf("len(buf) = %d, want 32768", len(buf))
		}
		PutBuffer(buf)
	})

	t.Run("GetLargeBuffer", func(t *testing.T) {
		buf := GetBuffer(524288)
		if len(buf) != 524288 {
			t.Errorf("len(buf) = %d, want 524288", len(buf))
		}
		PutBuffer(buf)
	})

	t.Run("GetVeryLargeBuffer", func(t *testing.T) {
		buf := GetBuffer(2097152) // 2MB - not pooled
		if len(buf) != 2097152 {
			t.Errorf("len(buf) = %d, want 2097152", len(buf))
		}
		PutBuffer(buf) // Should be no-op for very large buffers
	})

	t.Run("BufferReuse", func(t *testing.T) {
		// Get and return a buffer
		buf1 := GetBuffer(2048)
		buf1[0] = 42
		PutBuffer(buf1)

		// Get another buffer - might be the same one
		buf2 := GetBuffer(2048)
		// Should be zeroed
		if buf2[0] != 0 {
			t.Errorf("buf2[0] = %d, want 0 (buffer should be cleared)", buf2[0])
		}
		PutBuffer(buf2)
	})
}

func TestNewMemoryPooled(t *testing.T) {
	ResetPool()

	t.Run("CreateAndRelease", func(t *testing.T) {
		buf := NewMemoryPooled(1024)
		if buf.Size() != 1024 {
			t.Errorf("buf.Size() = %d, want 1024", buf.Size())
		}
		ReleaseMemoryBuffer(buf)

		// After release, data should be nil
		if buf.data != nil {
			t.Error("buf.data should be nil after release")
		}
	})

	t.Run("ReleaseNil", func(t *testing.T) {
		// Should not panic
		ReleaseMemoryBuffer(nil)
	})
}

func TestPooledBuffer(t *testing.T) {
	ResetPool()

	t.Run("AutoRelease", func(t *testing.T) {
		pb := NewPooledBuffer(2048)
		if pb.Size() != 2048 {
			t.Errorf("pb.Size() = %d, want 2048", pb.Size())
		}
		pb.Release()

		// Multiple releases should be safe
		pb.Release()
		pb.Release()
	})

	t.Run("WriteData", func(t *testing.T) {
		pb := NewPooledBuffer(100)
		data := pb.Data()
		data[0] = 123
		data[99] = 45

		if pb.Data()[0] != 123 {
			t.Error("Failed to write to pooled buffer")
		}

		pb.Release()
	})
}

// Benchmark tests

func BenchmarkGetBufferSmall(b *testing.B) {
	ResetPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := GetBuffer(1024)
		PutBuffer(buf)
	}
}

func BenchmarkGetBufferMedium(b *testing.B) {
	ResetPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := GetBuffer(32768)
		PutBuffer(buf)
	}
}

func BenchmarkGetBufferLarge(b *testing.B) {
	ResetPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := GetBuffer(524288)
		PutBuffer(buf)
	}
}

func BenchmarkAllocateSmallNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 1024)
		_ = buf
	}
}

func BenchmarkAllocateMediumNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 32768)
		_ = buf
	}
}

func BenchmarkAllocateLargeNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 524288)
		_ = buf
	}
}

func BenchmarkNewMemoryPooled(b *testing.B) {
	ResetPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewMemoryPooled(4096)
		ReleaseMemoryBuffer(buf)
	}
}

func BenchmarkNewMemoryNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := NewMemory(make([]byte, 4096))
		_ = buf
	}
}

func BenchmarkPooledBuffer(b *testing.B) {
	ResetPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pb := NewPooledBuffer(4096)
		pb.Release()
	}
}

func TestBytesBufferPool(t *testing.T) {
	ResetPool()

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

	t.Run("PutNil", func(t *testing.T) {
		// Should not panic
		PutBytesBuffer(nil)
	})
}

func BenchmarkGetBytesBuffer(b *testing.B) {
	ResetPool()
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
