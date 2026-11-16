// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// readBufferData returns the complete contents of a buffer and surfaces a
// deferred load failure. ByteBuffer.Data cannot report errors, so a lazy buffer
// whose loader failed would otherwise serialize as an empty value.
func readBufferData(buf buffer.ByteBuffer) ([]byte, error) {
	if lazy, ok := buf.(interface{ LoadError() error }); ok {
		if err := lazy.LoadError(); err != nil {
			return nil, fmt.Errorf("load deferred value: %w", err)
		}
	}
	return buf.Data(), nil
}
