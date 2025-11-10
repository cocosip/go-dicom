// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package dicom implements DICOM data structures and operations.
package dicom

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Item is the interface for all DICOM items (elements, sequences, etc.).
//
// A DICOM item represents a single data element in a DICOM dataset.
// Each item has a tag that identifies it and a value representation (VR)
// that defines how its data is encoded.
type Item interface {
	// Tag returns the DICOM tag associated with this item.
	Tag() *tag.Tag

	// ValueRepresentation returns the VR of this item.
	ValueRepresentation() *vr.VR

	// String returns a string representation of the item for debugging.
	String() string
}

// CompareItems compares two items by their tags.
// Returns -1 if a < b, 0 if a == b, 1 if a > b.
func CompareItems(a, b Item) int {
	if a == nil && b == nil {
		return 0
	}
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}
	return a.Tag().Compare(b.Tag())
}
