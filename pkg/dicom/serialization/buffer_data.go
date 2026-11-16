// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// readBufferData returns the complete contents of a buffer and surfaces a read
// failure. ByteBuffer.Data cannot report errors, so a lazy loader failure is
// returned as an error and a file, stream, range, or wrapped buffer that cannot
// produce its bytes is reported instead of being serialized as an empty value.
func readBufferData(buf buffer.ByteBuffer) ([]byte, error) {
	if buf == nil {
		return nil, fmt.Errorf("element has no value buffer")
	}
	if lazy, ok := buf.(interface{ LoadError() error }); ok {
		if err := lazy.LoadError(); err != nil {
			return nil, fmt.Errorf("load deferred value: %w", err)
		}
	}
	data := buf.Data()
	if data == nil && buf.Size() > 0 {
		return nil, fmt.Errorf("read %d-byte value: %w", buf.Size(), io.ErrUnexpectedEOF)
	}
	return data, nil
}
