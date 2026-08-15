// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// ChangeKind identifies the operation that changed an element.
type ChangeKind string

const (
	// ChangeRemove records removal of an existing element.
	ChangeRemove ChangeKind = "remove"
	// ChangeAssign records assignment or replacement of an element.
	ChangeAssign ChangeKind = "set"
	// ChangeMap records an exact value mapping.
	ChangeMap ChangeKind = "map"
	// ChangeCopy records a value copied between tags.
	ChangeCopy ChangeKind = "copy"
	// ChangeEdit records an in-place value edit.
	ChangeEdit ChangeKind = "edit"
	// ChangeUID records UID regeneration.
	ChangeUID ChangeKind = "uid"
)

// Change records one actual Dataset mutation. Before, After, and Path are
// independent snapshots.
type Change struct {
	Kind   ChangeKind
	Path   dataset.Path
	Tag    *tag.Tag
	Before element.Element
	After  element.Element
}

// ChangeSet records mutations in rule execution order.
type ChangeSet []Change

func appendChange(changes *ChangeSet, change Change) error {
	if changes == nil {
		return nil
	}
	change.Path = dataset.ClonePath(change.Path)
	change.Tag = change.Tag.Clone()
	var err error
	change.Before, err = dataset.DeepCloneElementChecked(change.Before)
	if err != nil {
		return fmt.Errorf("clone change Before: %w", err)
	}
	change.After, err = dataset.DeepCloneElementChecked(change.After)
	if err != nil {
		return fmt.Errorf("clone change After: %w", err)
	}
	*changes = append(*changes, change)
	return nil
}
