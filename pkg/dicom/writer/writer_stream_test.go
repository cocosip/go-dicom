// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type streamOnlyBuffer struct {
	data      []byte
	dataCalls int
}

func (b *streamOnlyBuffer) IsMemory() bool {
	return true
}

func (b *streamOnlyBuffer) Size() uint32 {
	return uint32(len(b.data)) //nolint:gosec // test data is bounded
}

func (b *streamOnlyBuffer) Data() []byte {
	b.dataCalls++
	return append([]byte(nil), b.data...)
}

func (b *streamOnlyBuffer) GetByteRange(offset, count uint32, output []byte) error {
	size := uint32(len(b.data)) //nolint:gosec // test data is bounded
	if offset > size || count > size-offset {
		return io.EOF
	}
	copy(output[:int(count)], b.data[int(offset):int(offset+count)])
	return nil
}

func (b *streamOnlyBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.data)
	return int64(n), err
}

func TestWriteLargeObjectUsesWriteTo(t *testing.T) {
	buf := &streamOnlyBuffer{data: bytes.Repeat([]byte{0x5A}, 4096)}

	ds := dataset.New()
	if err := ds.Add(element.NewUnknownFromBuffer(tag.New(0x0011, 0x1010), buf)); err != nil {
		t.Fatalf("failed to add test element: %v", err)
	}

	var out bytes.Buffer
	err := Write(&out, ds,
		WithoutPreamble(),
		WithLargeObjectSize(1), // force stream path
	)
	if err != nil {
		t.Fatalf("Write() error: %v", err)
	}

	if buf.dataCalls != 0 {
		t.Fatalf("Data() should not be called for large stream writes, got %d calls", buf.dataCalls)
	}
}
