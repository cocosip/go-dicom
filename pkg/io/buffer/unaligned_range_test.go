// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package buffer_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/io/endian"
)

// TestUnalignedGetByteRangeRejected verifies that byte-swapping buffers reject
// offsets that are not a multiple of the swap unit size, and that aligned
// offsets continue to swap correctly relative to the buffer start.
func TestUnalignedGetByteRangeRejected(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	wantSwapped := func(offset, count uint32) []byte {
		// Full-buffer pair swap, then slice the requested window.
		swapped := make([]byte, len(data))
		for i := 0; i < len(data); i += 2 {
			swapped[i], swapped[i+1] = data[i+1], data[i]
		}
		return swapped[offset : offset+count]
	}

	t.Run("SwapByteBuffer", func(t *testing.T) {
		sb, err := buffer.NewSwap(buffer.NewMemory(data), 2)
		if err != nil {
			t.Fatalf("NewSwap() error = %v", err)
		}

		if err := sb.GetByteRange(1, 2, make([]byte, 2)); err == nil {
			t.Error("GetByteRange(offset=1) with unit size 2: want error, got nil")
		}

		aligned := make([]byte, 4)
		if err := sb.GetByteRange(2, 4, aligned); err != nil {
			t.Fatalf("aligned GetByteRange() error = %v", err)
		}
		if want := wantSwapped(2, 4); !bytes.Equal(aligned, want) {
			t.Errorf("aligned GetByteRange() = %v, want %v", aligned, want)
		}
	})

	t.Run("EndianByteBuffer", func(t *testing.T) {
		// Choose an endianness opposite the native one so a wrapper is created.
		other := endian.Big
		if endian.Native() == endian.Big {
			other = endian.Little
		}
		eb := buffer.NewEndian(buffer.NewMemory(data), other, 2)

		if err := eb.GetByteRange(1, 2, make([]byte, 2)); err == nil {
			t.Error("GetByteRange(offset=1) with unit size 2: want error, got nil")
		}

		aligned := make([]byte, 4)
		if err := eb.GetByteRange(2, 4, aligned); err != nil {
			t.Fatalf("aligned GetByteRange() error = %v", err)
		}
		if want := wantSwapped(2, 4); !bytes.Equal(aligned, want) {
			t.Errorf("aligned GetByteRange() = %v, want %v", aligned, want)
		}
	})
}
