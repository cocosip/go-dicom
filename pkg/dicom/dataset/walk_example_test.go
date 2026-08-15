// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func ExampleWalk() {
	item := New()
	_ = item.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3"}))
	ds := New()
	_ = ds.Add(NewSequenceWithItems(tag.ReferencedImageSequence, []*Dataset{item}))
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"}))

	_ = Walk(ds, func(event WalkEvent) (WalkAction, error) {
		fmt.Printf("%s %s\n", event.Kind, FormatPath(event.Path))
		if event.Kind == WalkSequenceBegin {
			return WalkSkipChildren, nil
		}
		return WalkContinue, nil
	})

	// Output:
	// sequence-begin (0008,1140)
	// sequence-end (0008,1140)
	// element (0010,0010)
}
