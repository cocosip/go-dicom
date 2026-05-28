// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/daterange"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parseable"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// Dataset represents a DICOM dataset - a collection of data elements.
//
// Datasets maintain elements in sorted order by tag for consistent iteration
// and efficient lookup. Elements are stored in a map indexed by tag.
//
// Dataset is NOT goroutine-safe. All mutations (Add, AddOrUpdate, Remove, Clear,
// Merge, SetInternalTransferSyntax) and any reads that may race with mutations
// must be externally synchronized by the caller. Concurrent read-only access
// (e.g., Get, Contains, Elements, Tags) is safe as long as no concurrent writes
// are in progress.
type Dataset struct {
	// items stores elements indexed by tag
	items map[uint32]element.Element

	// sortedTags caches sorted tag values for deterministic iteration.
	// It is invalidated on structural mutations (add/remove/clear/merge/filter).
	sortedTags []uint32
	cacheDirty bool

	// internalTransferSyntax represents the transfer syntax of this dataset.
	// This is used to track the encoding format of pixel data and other elements.
	// Following fo-dicom pattern, this is internal and can be set by transcoder/parser.
	internalTransferSyntax *transfer.Syntax
}

// New creates a new empty dataset.
func New() *Dataset {
	return &Dataset{
		items:      make(map[uint32]element.Element),
		cacheDirty: true,
	}
}

// NewWithTransferSyntax creates a new empty dataset with the specified transfer syntax.
// This is the recommended way to create datasets when the transfer syntax is known
// (e.g., when transcoding or parsing).
func NewWithTransferSyntax(ts *transfer.Syntax) *Dataset {
	return &Dataset{
		items:                  make(map[uint32]element.Element),
		internalTransferSyntax: ts,
		cacheDirty:             true,
	}
}

// NewWithElements creates a dataset initialized with the given elements.
// Later elements replace earlier ones when the same tag appears multiple times.
// Nil elements are ignored.
func NewWithElements(elements []element.Element) *Dataset {
	ds := New()
	for _, elem := range elements {
		if elem == nil {
			continue
		}
		_ = ds.AddOrUpdate(elem)
	}
	return ds
}

// Add adds an element to the dataset.
// If an element with the same tag already exists, it returns an error.
func (ds *Dataset) Add(elem element.Element) error {
	if ds == nil {
		return fmt.Errorf("cannot add to nil Dataset")
	}
	if elem == nil {
		return fmt.Errorf("cannot add nil element")
	}
	ds.ensureItems()

	tagValue := elem.Tag().ToUint32()
	if _, exists := ds.items[tagValue]; exists {
		return fmt.Errorf("element with tag %s already exists", elem.Tag())
	}

	ds.items[tagValue] = elem
	ds.markDirty()
	return nil
}

// AddOrUpdate adds an element to the dataset, or updates it if it already exists.
func (ds *Dataset) AddOrUpdate(elem element.Element) error {
	if ds == nil {
		return fmt.Errorf("cannot add to nil Dataset")
	}
	if elem == nil {
		return fmt.Errorf("cannot add nil element")
	}
	ds.ensureItems()

	ds.items[elem.Tag().ToUint32()] = elem
	ds.markDirty()
	return nil
}

// Get retrieves an element by tag.
// Returns the element and true if found, nil and false otherwise.
func (ds *Dataset) Get(t *tag.Tag) (element.Element, bool) {
	if ds == nil || t == nil {
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
	if ds == nil || t == nil {
		return false
	}
	_, exists := ds.items[t.ToUint32()]
	return exists
}

// Remove removes an element by tag.
// Returns true if the element was removed, false if it didn't exist.
func (ds *Dataset) Remove(t *tag.Tag) bool {
	if ds == nil || t == nil {
		return false
	}
	if ds.items == nil {
		return false
	}
	tagValue := t.ToUint32()
	if _, exists := ds.items[tagValue]; exists {
		delete(ds.items, tagValue)
		ds.markDirty()
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
	if ds == nil {
		return
	}
	ds.items = make(map[uint32]element.Element)
	ds.markDirty()
}

// Count returns the number of elements in the dataset.
func (ds *Dataset) Count() int {
	if ds == nil {
		return 0
	}
	return len(ds.items)
}

// IsEmpty returns true if the dataset contains no elements.
func (ds *Dataset) IsEmpty() bool {
	return ds == nil || len(ds.items) == 0
}

// Elements returns all elements in the dataset, sorted by tag.
func (ds *Dataset) Elements() []element.Element {
	if ds == nil || len(ds.items) == 0 {
		return nil
	}

	tagValues := ds.sortedTagValues()

	// Build sorted element list
	elements := make([]element.Element, len(tagValues))
	for i, tagValue := range tagValues {
		elements[i] = ds.items[tagValue]
	}

	return elements
}

// Tags returns all tags in the dataset, sorted.
func (ds *Dataset) Tags() []*tag.Tag {
	if ds == nil || len(ds.items) == 0 {
		return nil
	}

	tagValues := ds.sortedTagValues()

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
	if ds == nil {
		return New()
	}
	clone := New()
	clone.internalTransferSyntax = ds.internalTransferSyntax // Preserve transfer syntax
	for tagValue, elem := range ds.items {
		clone.items[tagValue] = elem
	}
	if !ds.cacheDirty && len(ds.sortedTags) > 0 {
		clone.sortedTags = append([]uint32(nil), ds.sortedTags...)
		clone.cacheDirty = false
	}
	return clone
}

// Merge merges elements from another dataset into this one.
// If overwrite is true, existing elements are replaced.
// If overwrite is false, only new elements are added.
func (ds *Dataset) Merge(other *Dataset, overwrite bool) {
	if ds == nil || other == nil {
		return
	}
	ds.ensureItems()

	for tagValue, elem := range other.items {
		if overwrite || !ds.Contains(elem.Tag()) {
			ds.items[tagValue] = elem
			ds.markDirty()
		}
	}
}

// Filter returns a new dataset containing only elements that match the predicate.
func (ds *Dataset) Filter(predicate func(element.Element) bool) *Dataset {
	if ds == nil {
		return New()
	}
	filtered := New()
	for _, elem := range ds.items {
		if predicate(elem) {
			if err := filtered.Add(elem); err != nil {
				continue
			}
		}
	}
	return filtered
}

// String returns a string representation of the dataset.
func (ds *Dataset) String() string {
	if ds == nil {
		return "Dataset{nil}"
	}
	return fmt.Sprintf("Dataset{%d elements}", len(ds.items))
}

// InternalTransferSyntax returns the transfer syntax of this dataset.
// This represents the encoding format of pixel data and other elements.
// Returns nil if no transfer syntax has been set.
func (ds *Dataset) InternalTransferSyntax() *transfer.Syntax {
	if ds == nil {
		return nil
	}
	return ds.internalTransferSyntax
}

// SetInternalTransferSyntax sets the transfer syntax of this dataset.
// This is an internal method that should primarily be used by transcoder and parser.
// Following fo-dicom pattern, this also propagates to nested sequence items.
//
// Note: In most cases, you should create a new dataset with NewWithTransferSyntax()
// rather than modifying an existing one. However, this setter is provided for
// scenarios where in-place modification is necessary (e.g., after cloning).
func (ds *Dataset) SetInternalTransferSyntax(ts *transfer.Syntax) {
	if ds == nil {
		return
	}
	ds.internalTransferSyntax = ts

	// Update transfer syntax for sequence items (following fo-dicom pattern)
	for _, elem := range ds.items {
		if seq, ok := elem.(*Sequence); ok {
			for i := 0; i < seq.Count(); i++ {
				item := seq.GetItem(i)
				if item != nil {
					item.SetInternalTransferSyntax(ts)
				}
			}
		}
	}
}

// ensureItems lazily initializes the items map if it is nil.
// This allows zero-value Datasets (created via &Dataset{} instead of New())
// to be safely used without panicking on first write.
func (ds *Dataset) ensureItems() {
	if ds.items == nil {
		ds.items = make(map[uint32]element.Element)
	}
}

func (ds *Dataset) markDirty() {
	ds.cacheDirty = true
	ds.sortedTags = nil
}

func (ds *Dataset) sortedTagValues() []uint32 {
	if !ds.cacheDirty && len(ds.sortedTags) == len(ds.items) {
		return ds.sortedTags
	}

	tagValues := make([]uint32, 0, len(ds.items))
	for tagValue := range ds.items {
		tagValues = append(tagValues, tagValue)
	}
	sort.Slice(tagValues, func(i, j int) bool {
		return tagValues[i] < tagValues[j]
	})
	ds.sortedTags = tagValues
	ds.cacheDirty = false
	return ds.sortedTags
}

// GetParseable retrieves an element by tag and parses it using the parseable.Parseable interface.
// This is useful for types that implement Parseable, such as uid.UID or transfer.Syntax.
//
// Example:
//
//	ts, err := ds.GetParseable(tag.TransferSyntaxUID, parseable.ParserFor(func() *transfer.Syntax { return &transfer.Syntax{} }))
//	sopClass, err := ds.GetParseable(tag.SOPClassUID, parseable.ParserFor(func() *uid.UID { return &uid.UID{} }))
func GetParseable[T parseable.Parseable](ds *Dataset, t *tag.Tag, parser parseable.Parser[T]) (T, error) {
	var zero T
	elem, exists := ds.Get(t)
	if !exists {
		return zero, fmt.Errorf("element %s not found", t)
	}

	strElem, ok := elem.(*element.String)
	if !ok {
		return zero, fmt.Errorf("element %s is not a string type", t)
	}

	value := strElem.GetString()
	if value == "" {
		return zero, fmt.Errorf("element %s is empty", t)
	}

	return parser(value)
}

// GetUID retrieves a UID element and parses it.
func (ds *Dataset) GetUID(t *tag.Tag) (*uid.UID, error) {
	return GetParseable(ds, t, parseable.ParserFor(func() *uid.UID { return &uid.UID{} }))
}

// GetTransferSyntax retrieves the Transfer Syntax UID and returns the TransferSyntax.
func (ds *Dataset) GetTransferSyntax() (*transfer.Syntax, error) {
	return GetParseable(ds, tag.TransferSyntaxUID, parseable.ParserFor(func() *transfer.Syntax { return &transfer.Syntax{} }))
}

// GetDateRange retrieves a date element and parses it as a DateRange.
// This supports DICOM query format: "YYYYMMDD-YYYYMMDD", "YYYYMMDD-", "-YYYYMMDD", or single date.
func (ds *Dataset) GetDateRange(t *tag.Tag) (*daterange.DateRange, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	dateElem, ok := elem.(*element.Date)
	if !ok {
		return nil, fmt.Errorf("element %s is not a Date type", t)
	}

	return dateElem.GetDateRange()
}

// GetTimeRange retrieves a time element and parses it as a TimeRange.
func (ds *Dataset) GetTimeRange(t *tag.Tag) (*daterange.TimeRange, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	timeElem, ok := elem.(*element.Time)
	if !ok {
		return nil, fmt.Errorf("element %s is not a Time type", t)
	}

	return timeElem.GetTimeRange()
}

// GetDateTimeRange retrieves a datetime element and parses it as a DateTimeRange.
func (ds *Dataset) GetDateTimeRange(t *tag.Tag) (*daterange.DateTimeRange, error) {
	elem, exists := ds.Get(t)
	if !exists {
		return nil, fmt.Errorf("element %s not found", t)
	}

	dtElem, ok := elem.(*element.DateTime)
	if !ok {
		return nil, fmt.Errorf("element %s is not a DateTime type", t)
	}

	return dtElem.GetDateTimeRange()
}
