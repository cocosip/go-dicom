// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package buffer_test

import (
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// failingChildBuffer declares a non-zero size but cannot provide data, modelling
// a child whose backing storage fails to read.
type failingChildBuffer struct{ size uint32 }

func (b failingChildBuffer) IsMemory() bool                           { return false }
func (b failingChildBuffer) Size() uint32                             { return b.size }
func (b failingChildBuffer) Data() []byte                             { return nil }
func (b failingChildBuffer) GetByteRange(_, _ uint32, _ []byte) error { return io.ErrUnexpectedEOF }
func (b failingChildBuffer) WriteTo(_ io.Writer) (int64, error)       { return 0, io.ErrUnexpectedEOF }

// Composite.Data must report failure (nil) instead of returning a zero-filled
// span that looks like real data.
func TestCompositeDataReturnsNilOnChildFailure(t *testing.T) {
	composite := buffer.NewComposite(buffer.NewMemory([]byte{1, 2}), failingChildBuffer{size: 4})
	if got := composite.Data(); got != nil {
		t.Fatalf("Data() = %v, want nil when a child buffer read fails", got)
	}
}
