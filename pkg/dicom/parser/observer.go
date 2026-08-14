// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// SequenceItemPosition describes the physical location of a parsed sequence item.
// Offset is the absolute byte position of the item's (FFFE,E000) tag.
type SequenceItemPosition struct {
	SequenceTag *tag.Tag
	Item        *dataset.Dataset
	Offset      uint64
}

// SequenceItemObserver receives sequence item positions as items are parsed.
type SequenceItemObserver func(SequenceItemPosition) error

// WithSequenceItemObserver reports the absolute position of each parsed sequence item.
// Deflated transfer syntaxes are not supported because compressed stream positions do
// not identify item offsets in the decoded dataset.
func WithSequenceItemObserver(observer SequenceItemObserver) Option {
	return func(ctx *parseContext) {
		ctx.sequenceItemObserver = observer
	}
}
