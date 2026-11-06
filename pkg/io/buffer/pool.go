// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package buffer

import (
	"bytes"
	"sync"
)

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
