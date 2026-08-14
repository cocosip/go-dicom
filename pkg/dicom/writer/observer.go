// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// SequenceItemPosition describes the physical location of a written sequence item.
// Offset is the absolute byte position of the item's (FFFE,E000) tag.
type SequenceItemPosition struct {
	SequenceTag *tag.Tag
	Item        *dataset.Dataset
	Offset      uint64
}

// SequenceItemObserver receives sequence item positions as items are written.
type SequenceItemObserver func(SequenceItemPosition) error

// WithSequenceItemObserver reports the absolute position of each written sequence item.
// Deflated transfer syntaxes are not supported because compressed stream positions do
// not identify item offsets in the decoded dataset.
func WithSequenceItemObserver(observer SequenceItemObserver) WriteOption {
	return func(config *writeConfig) {
		config.sequenceItemObserver = observer
	}
}

type positionWriter struct {
	writer io.Writer
	offset uint64
}

func (w *positionWriter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.offset += uint64(n) // #nosec G115 -- Write never returns a negative count
	return n, err
}
