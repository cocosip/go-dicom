// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// Dataset represents a DICOM dataset - a collection of data elements.
//
// Datasets maintain elements in sorted order by tag for consistent iteration
// and efficient lookup. Elements are stored in a map indexed by tag.
type Dataset struct {
	// items stores elements indexed by tag
	items map[uint32]element.Element
}

// New creates a new empty dataset.
func New() *Dataset {
	return &Dataset{
		items: make(map[uint32]element.Element),
	}
}

// NewWithElements creates a dataset initialized with the given elements.
func NewWithElements(elements []element.Element) *Dataset {
	ds := New()
	for _, elem := range elements {
		ds.Add(elem)
	}
	return ds
}

// Add adds an element to the dataset.
// If an element with the same tag already exists, it returns an error.
func (ds *Dataset) Add(elem element.Element) error {
	if elem == nil {
		return fmt.Errorf("cannot add nil element")
	}

	tagValue := elem.Tag().ToUint32()
	if _, exists := ds.items[tagValue]; exists {
		return fmt.Errorf("element with tag %s already exists", elem.Tag())
	}

	ds.items[tagValue] = elem
	return nil
}

// AddOrUpdate adds an element to the dataset, or updates it if it already exists.
func (ds *Dataset) AddOrUpdate(elem element.Element) error {
	if elem == nil {
		return fmt.Errorf("cannot add nil element")
	}

	ds.items[elem.Tag().ToUint32()] = elem
	return nil
}

// Get retrieves an element by tag.
// Returns the element and true if found, nil and false otherwise.
func (ds *Dataset) Get(t *tag.Tag) (element.Element, bool) {
	if t == nil {
		return nil, false
	}
	elem, exists := ds.items[t.ToUint32()]
	return elem, exists
}

// GetOrNil retrieves an element by tag, returning nil if not found.
func (ds *Dataset) GetOrNil(t *tag.Tag) element.Element {
	elem, _ := ds.Get(t)
	return elem
}

// Contains checks if an element with the given tag exists.
func (ds *Dataset) Contains(t *tag.Tag) bool {
	if t == nil {
		return false
	}
	_, exists := ds.items[t.ToUint32()]
	return exists
}

// Remove removes an element by tag.
// Returns true if the element was removed, false if it didn't exist.
func (ds *Dataset) Remove(t *tag.Tag) bool {
	if t == nil {
		return false
	}
	tagValue := t.ToUint32()
	if _, exists := ds.items[tagValue]; exists {
		delete(ds.items, tagValue)
		return true
	}
	return false
}

// RemoveAll removes all elements with the specified tags.
// Returns the number of elements removed.
func (ds *Dataset) RemoveAll(tags ...*tag.Tag) int {
	count := 0
	for _, t := range tags {
		if ds.Remove(t) {
			count++
		}
	}
	return count
}

// Clear removes all elements from the dataset.
func (ds *Dataset) Clear() {
	ds.items = make(map[uint32]element.Element)
}

// Count returns the number of elements in the dataset.
func (ds *Dataset) Count() int {
	return len(ds.items)
}

// IsEmpty returns true if the dataset contains no elements.
func (ds *Dataset) IsEmpty() bool {
	return len(ds.items) == 0
}

// Elements returns all elements in the dataset, sorted by tag.
func (ds *Dataset) Elements() []element.Element {
	if len(ds.items) == 0 {
		return nil
	}

	// Get sorted tags
	tags := make([]uint32, 0, len(ds.items))
	for tagValue := range ds.items {
		tags = append(tags, tagValue)
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	// Build sorted element list
	elements := make([]element.Element, len(tags))
	for i, tagValue := range tags {
		elements[i] = ds.items[tagValue]
	}

	return elements
}

// Tags returns all tags in the dataset, sorted.
func (ds *Dataset) Tags() []*tag.Tag {
	if len(ds.items) == 0 {
		return nil
	}

	// Get sorted tags
	tagValues := make([]uint32, 0, len(ds.items))
	for tagValue := range ds.items {
		tagValues = append(tagValues, tagValue)
	}
	sort.Slice(tagValues, func(i, j int) bool {
		return tagValues[i] < tagValues[j]
	})

	// Convert to Tag objects
	tags := make([]*tag.Tag, len(tagValues))
	for i, tagValue := range tagValues {
		tags[i] = tag.FromUint32(tagValue)
	}

	return tags
}

// Clone creates a deep copy of the dataset.
// Note: Elements themselves are not cloned, only the dataset structure.
func (ds *Dataset) Clone() *Dataset {
	clone := New()
	for tagValue, elem := range ds.items {
		clone.items[tagValue] = elem
	}
	return clone
}

// Merge merges elements from another dataset into this one.
// If overwrite is true, existing elements are replaced.
// If overwrite is false, only new elements are added.
func (ds *Dataset) Merge(other *Dataset, overwrite bool) {
	if other == nil {
		return
	}

	for tagValue, elem := range other.items {
		if overwrite || !ds.Contains(elem.Tag()) {
			ds.items[tagValue] = elem
		}
	}
}

// Filter returns a new dataset containing only elements that match the predicate.
func (ds *Dataset) Filter(predicate func(element.Element) bool) *Dataset {
	filtered := New()
	for _, elem := range ds.items {
		if predicate(elem) {
			filtered.items[elem.Tag().ToUint32()] = elem
		}
	}
	return filtered
}

// String returns a string representation of the dataset.
func (ds *Dataset) String() string {
	return fmt.Sprintf("Dataset{%d elements}", len(ds.items))
}
