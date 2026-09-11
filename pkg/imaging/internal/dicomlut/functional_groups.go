// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dicomlut

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// FunctionalGroupValues flattens the first item of each shared and per-frame
// functional-group macro. Per-frame values replace shared values.
func FunctionalGroupValues(ds *dataset.Dataset, frame int) *dataset.Dataset {
	values := dataset.New()
	// Functional-group flattening is a read operation. Preserve elements from
	// parsed or legacy datasets even when their nested VR is non-standard; the
	// consumer's strict/compatible policy performs the semantic check later.
	values.SetAutoValidate(false)
	if ds == nil {
		return values
	}
	values.SetInternalTransferSyntax(ds.InternalTransferSyntax())
	mergeFunctionalGroupValues(values, ds, tag.SharedFunctionalGroupsSequence, 0)
	mergeFunctionalGroupValues(values, ds, tag.PerFrameFunctionalGroupsSequence, frame)
	return values
}

func mergeFunctionalGroupValues(values, ds *dataset.Dataset, sequenceTag *tag.Tag, itemIndex int) {
	sequence, err := ds.GetSequence(sequenceTag)
	if err != nil || itemIndex < 0 || itemIndex >= sequence.Count() {
		return
	}
	item := sequence.GetItem(itemIndex)
	if item == nil {
		return
	}
	for _, elem := range item.Elements() {
		nested, ok := elem.(*dataset.Sequence)
		if !ok || nested.Count() == 0 || nested.GetItem(0) == nil {
			continue
		}
		for _, value := range nested.GetItem(0).Elements() {
			_ = values.AddOrUpdate(value)
		}
	}
}
