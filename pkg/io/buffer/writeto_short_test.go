// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package buffer_test

import (
	"errors"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/io/endian"
)

// shortWriter accepts every write but reports fewer bytes written with a nil
// error, which io.Writer allows only as a contract violation; callers must
// translate it to io.ErrShortWrite instead of silently dropping the tail.
type shortWriter struct{ written int }

func (w *shortWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	w.written += len(p) - 1
	return len(p) - 1, nil
}

func TestChunkedWriteToReportsShortWrite(t *testing.T) {
	internal := buffer.NewMemory([]byte("abcdefgh"))
	ranged, err := buffer.NewRange(internal, 0, 8)
	if err != nil {
		t.Fatalf("NewRange() error = %v", err)
	}
	big := buffer.NewMemory(make([]byte, 2*1024*1024+3)) // forces multi-chunk writes
	endianWrapped := buffer.NewEndian(big, endian.Big, 2)
	swapped, err := buffer.NewSwap(big, 2)
	if err != nil {
		t.Fatalf("NewSwap() error = %v", err)
	}

	tests := []struct {
		name string
		buf  buffer.ByteBuffer
	}{
		{"range", ranged},
		{"endian", endianWrapped},
		{"swap", swapped},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &shortWriter{}
			n, err := tt.buf.WriteTo(w)
			if !errors.Is(err, io.ErrShortWrite) {
				t.Fatalf("WriteTo() error = %v, want io.ErrShortWrite", err)
			}
			if n != int64(w.written) {
				t.Fatalf("WriteTo() n = %d, want %d", n, w.written)
			}
		})
	}
}
