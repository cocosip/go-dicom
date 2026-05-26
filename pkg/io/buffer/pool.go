// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"bytes"
	"sync"
)

// bufferPool is a pool for reusing byte slices to reduce allocations.
// It maintains multiple pools for different size classes to optimize memory usage.
type bufferPool struct {
	// Pools for common buffer sizes
	small  sync.Pool // <= 4KB
	medium sync.Pool // <= 64KB
	large  sync.Pool // <= 1MB
}

// Global buffer pool instance
var defaultBufferPool = &bufferPool{
	small: sync.Pool{
		New: func() interface{} {
			b := make([]byte, 4096)
			return &b
		},
	},
	medium: sync.Pool{
		New: func() interface{} {
			b := make([]byte, 65536)
			return &b
		},
	},
	large: sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1048576)
			return &b
		},
	},
}

// GetBuffer returns a buffer from the pool with at least the specified size.
// The returned slice may be larger than requested.
// Caller should call PutBuffer when done to return it to the pool.
func GetBuffer(size uint32) []byte {
	switch {
	case size <= 4096:
		bufPtr, ok := defaultBufferPool.small.Get().(*[]byte)
		if !ok {
			panic("buffer pool: small pool returned non-*[]byte value")
		}
		return (*bufPtr)[:size]
	case size <= 65536:
		bufPtr, ok := defaultBufferPool.medium.Get().(*[]byte)
		if !ok {
			panic("buffer pool: medium pool returned non-*[]byte value")
		}
		return (*bufPtr)[:size]
	case size <= 1048576:
		bufPtr, ok := defaultBufferPool.large.Get().(*[]byte)
		if !ok {
			panic("buffer pool: large pool returned non-*[]byte value")
		}
		return (*bufPtr)[:size]
	default:
		// For very large buffers, allocate directly without pooling
		return make([]byte, size)
	}
}

// PutBuffer returns a buffer to the pool for reuse.
// The buffer should not be used after calling this function.
// Only buffers originally obtained from GetBuffer should be returned.
// Buffers with capacities that don't match a pool size class are discarded
// to avoid putting incorrectly-sized slices into the pool.
func PutBuffer(buf []byte) {
	if buf == nil {
		return
	}

	capacity := cap(buf)

	// Only return to pool if capacity exactly matches one of the pool sizes.
	// Sub-slices and externally-allocated slices may have mismatched capacities
	// that would corrupt pool invariants.
	var pool *sync.Pool
	switch capacity {
	case 4096:
		pool = &defaultBufferPool.small
	case 65536:
		pool = &defaultBufferPool.medium
	case 1048576:
		pool = &defaultBufferPool.large
	default:
		// Capacity doesn't match any pool size — don't pool, let GC handle it
		return
	}

	// Clear the buffer before returning to pool
	full := buf[:capacity]
	clear(full)
	pool.Put(&full)
}

// NewMemoryPooled creates a new MemoryByteBuffer using a pooled buffer.
// The buffer should be returned to the pool by calling ReleaseMemoryBuffer.
func NewMemoryPooled(size uint32) *MemoryByteBuffer {
	data := GetBuffer(size)
	return &MemoryByteBuffer{data: data}
}

// ReleaseMemoryBuffer returns the buffer's memory to the pool.
// The buffer should not be used after calling this function.
func ReleaseMemoryBuffer(buf *MemoryByteBuffer) {
	if buf != nil && buf.data != nil {
		PutBuffer(buf.data)
		buf.data = nil
	}
}

// PooledBuffer is a wrapper that automatically returns buffer to pool when done.
type PooledBuffer struct {
	*MemoryByteBuffer
	released bool
}

// NewPooledBuffer creates a new pooled buffer that will automatically
// return to pool when released.
func NewPooledBuffer(size uint32) *PooledBuffer {
	return &PooledBuffer{
		MemoryByteBuffer: NewMemoryPooled(size),
		released:         false,
	}
}

// Release returns the buffer to the pool. Safe to call multiple times.
func (p *PooledBuffer) Release() {
	if !p.released && p.MemoryByteBuffer != nil {
		ReleaseMemoryBuffer(p.MemoryByteBuffer)
		p.released = true
	}
}

// ResetPool clears all buffers from the pool.
// This is mainly useful for testing.
func ResetPool() {
	defaultBufferPool.small = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 4096)
			return &b
		},
	}
	defaultBufferPool.medium = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 65536)
			return &b
		},
	}
	defaultBufferPool.large = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1048576)
			return &b
		},
	}
	bytesBufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// bytes.Buffer pooling for temporary buffers

// bytesBufferPool is a pool for temporary bytes.Buffer instances.
var bytesBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// GetBytesBuffer returns a bytes.Buffer from the pool.
// The caller should call PutBytesBuffer when done.
func GetBytesBuffer() *bytes.Buffer {
	buf, ok := bytesBufferPool.Get().(*bytes.Buffer)
	if !ok {
		panic("buffer pool: bytes buffer pool returned non-*bytes.Buffer value")
	}
	return buf
}

// PutBytesBuffer returns a bytes.Buffer to the pool for reuse.
// The buffer is reset before returning to the pool.
func PutBytesBuffer(buf *bytes.Buffer) {
	if buf != nil {
		buf.Reset()
		bytesBufferPool.Put(buf)
	}
}
