// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict

import (
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// Dictionary manages DICOM dictionary entries.
//
// The dictionary provides lookup functionality for DICOM tags to retrieve
// their metadata including VR, VM, name, and keyword.
type Dictionary struct {
	// entries maps tags to their dictionary entries
	entries map[uint32]*Entry

	// keywords maps keywords to tags
	keywords map[string]*tag.Tag

	// masked contains entries with wildcard tags (e.g., group length)
	masked []*Entry

	// privateCreator is set for private dictionaries
	privateCreator *tag.PrivateCreator

	mu sync.RWMutex
}

// New creates a new empty DICOM dictionary.
func New() *Dictionary {
	return &Dictionary{
		entries:  make(map[uint32]*Entry),
		keywords: make(map[string]*tag.Tag),
		masked:   make([]*Entry, 0),
	}
}

// NewPrivate creates a new private dictionary for a specific creator.
func NewPrivate(creator *tag.PrivateCreator) *Dictionary {
	return &Dictionary{
		entries:        make(map[uint32]*Entry),
		keywords:       make(map[string]*tag.Tag),
		masked:         make([]*Entry, 0),
		privateCreator: creator,
	}
}

// Add adds a dictionary entry.
func (d *Dictionary) Add(entry *Entry) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// If this is a masked entry, add to masked list
	if entry.IsMasked() {
		d.masked = append(d.masked, entry)
	} else {
		// Add to exact match map
		d.entries[entry.Tag().ToUint32()] = entry
	}

	// Add keyword mapping
	if entry.Keyword() != "" {
		d.keywords[entry.Keyword()] = entry.Tag()
	}
}

// Lookup finds a dictionary entry for the given tag.
//
// Returns nil if the tag is not found.
func (d *Dictionary) Lookup(t *tag.Tag) *Entry {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Try exact match first
	if entry, ok := d.entries[t.ToUint32()]; ok {
		return entry
	}

	// Try masked entries
	for _, entry := range d.masked {
		if entry.Matches(t) {
			return entry
		}
	}

	return nil
}

// LookupKeyword finds a tag by its keyword.
//
// Returns nil if the keyword is not found.
func (d *Dictionary) LookupKeyword(keyword string) *tag.Tag {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if t, ok := d.keywords[keyword]; ok {
		return t
	}
	return nil
}

// PrivateCreator returns the private creator, or nil for public dictionaries.
func (d *Dictionary) PrivateCreator() *tag.PrivateCreator {
	return d.privateCreator
}

// Entries returns all entries in the dictionary.
func (d *Dictionary) Entries() []*Entry {
	d.mu.RLock()
	defer d.mu.RUnlock()

	result := make([]*Entry, 0, len(d.entries)+len(d.masked))
	for _, entry := range d.entries {
		result = append(result, entry)
	}
	result = append(result, d.masked...)
	return result
}
