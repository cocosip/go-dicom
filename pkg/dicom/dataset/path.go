// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// PathSegment identifies an element, Sequence item, or fragment in a Dataset
// path. ItemIndex and FragmentIndex are mutually exclusive.
type PathSegment struct {
	Tag           *tag.Tag
	ItemIndex     *int
	FragmentIndex *int
}

// Path identifies a location in a Dataset hierarchy.
type Path []PathSegment

// ValidationPathSegment is retained as an alias for source compatibility.
type ValidationPathSegment = PathSegment

// ClonePath returns an independent path snapshot, including mutable Tags and
// index values.
func ClonePath(path Path) Path {
	if path == nil {
		return nil
	}
	clone := make(Path, len(path))
	for index, segment := range path {
		clone[index].Tag = segment.Tag.Clone()
		if segment.ItemIndex != nil {
			itemIndex := *segment.ItemIndex
			clone[index].ItemIndex = &itemIndex
		}
		if segment.FragmentIndex != nil {
			fragmentIndex := *segment.FragmentIndex
			clone[index].FragmentIndex = &fragmentIndex
		}
	}
	return clone
}

// FormatPath returns a stable human-readable Dataset path.
func FormatPath(path Path) string {
	if len(path) == 0 {
		return "<dataset>"
	}
	parts := make([]string, 0, len(path))
	for _, segment := range path {
		part := "<unknown-tag>"
		if segment.Tag != nil {
			part = segment.Tag.String()
		}
		if segment.ItemIndex != nil {
			part += fmt.Sprintf("[%d]", *segment.ItemIndex)
		}
		if segment.FragmentIndex != nil {
			part += fmt.Sprintf("#%d", *segment.FragmentIndex)
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, "/")
}

func appendPath(path Path, segment PathSegment) Path {
	result := ClonePath(path)
	return append(result, ClonePath(Path{segment})[0])
}
