// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// Sequence represents a DICOM element with VR = SQ (Sequence of Items).
//
// A sequence contains zero or more items (datasets), allowing hierarchical
// nesting of DICOM data. Each item in the sequence is itself a complete dataset.
//
// Example:
//
//	seq := dataset.NewSequence(tag.ReferencedImageSequence)
//	item1 := dataset.New()
//	item1.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
//	seq.AddItem(item1)
var _ element.Element = (*Sequence)(nil)

// Sequence represents a DICOM sequence element.
type Sequence struct {
	tag   *tag.Tag
	items []*Dataset
}

// NewSequence creates a new empty sequence with the given tag.
func NewSequence(t *tag.Tag) *Sequence {
	return &Sequence{
		tag:   t,
		items: make([]*Dataset, 0),
	}
}

// NewSequenceWithItems creates a sequence initialized with the given items.
func NewSequenceWithItems(t *tag.Tag, items []*Dataset) *Sequence {
	return &Sequence{
		tag:   t,
		items: items,
	}
}

// Tag returns the tag associated with this sequence.
func (s *Sequence) Tag() *tag.Tag {
	return s.tag
}

// ValueRepresentation returns the VR for this sequence (always SQ).
func (s *Sequence) ValueRepresentation() *vr.VR {
	return vr.SQ
}

// Buffer returns an empty buffer (sequences don't have binary data directly).
func (s *Sequence) Buffer() buffer.ByteBuffer {
	return buffer.Empty
}

// Length returns 0 (sequences use undefined length encoding).
func (s *Sequence) Length() uint32 {
	return 0 // Sequences use undefined length
}

// Count returns the number of items in the sequence.
func (s *Sequence) Count() int {
	return len(s.items)
}

// String returns a string representation of the sequence.
func (s *Sequence) String() string {
	return fmt.Sprintf("Sequence{Tag: %s, Items: %d}", s.tag.String(), len(s.items))
}

// Validate validates the sequence (placeholder for future validation logic).
func (s *Sequence) Validate() error {
	return nil
}

// AddItem adds a dataset item to the sequence.
func (s *Sequence) AddItem(ds *Dataset) {
	if ds != nil {
		s.items = append(s.items, ds)
	}
}

// GetItem retrieves the item at the specified index.
// Returns nil if index is out of range.
func (s *Sequence) GetItem(index int) *Dataset {
	if index < 0 || index >= len(s.items) {
		return nil
	}
	return s.items[index]
}

// GetItems returns all items in the sequence.
func (s *Sequence) GetItems() []*Dataset {
	return s.items
}

// RemoveItem removes the item at the specified index.
// Returns true if the item was removed, false if index was out of range.
func (s *Sequence) RemoveItem(index int) bool {
	if index < 0 || index >= len(s.items) {
		return false
	}

	// Remove item by slicing
	s.items = append(s.items[:index], s.items[index+1:]...)
	return true
}

// Clear removes all items from the sequence.
func (s *Sequence) Clear() {
	s.items = make([]*Dataset, 0)
}

// IsEmpty returns true if the sequence contains no items.
func (s *Sequence) IsEmpty() bool {
	return len(s.items) == 0
}

// Clone creates a shallow copy of the sequence.
// Items (datasets) are not cloned.
func (s *Sequence) Clone() *Sequence {
	items := make([]*Dataset, len(s.items))
	copy(items, s.items)
	return &Sequence{
		tag:   s.tag,
		items: items,
	}
}

// Filter returns a new sequence containing only items that match the predicate.
func (s *Sequence) Filter(predicate func(*Dataset) bool) *Sequence {
	filtered := NewSequence(s.tag)
	for _, item := range s.items {
		if predicate(item) {
			filtered.AddItem(item)
		}
	}
	return filtered
}
