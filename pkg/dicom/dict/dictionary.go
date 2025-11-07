// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict

import (
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// init registers the dictionary lookup functions with the tag package.
func init() {
	tag.SetDictionaryLookup(func(t *tag.Tag) interface{} {
		return Default().Lookup(t)
	})
	tag.SetKeywordLookup(func(keyword string) (*tag.Tag, error) {
		t := Default().LookupKeyword(keyword)
		if t == nil {
			return nil, nil
		}
		return t, nil
	})
	tag.SetPrivateCreatorLookup(func(creator string) *tag.PrivateCreator {
		return Default().GetPrivateCreator(creator)
	})
}

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

	// creators maintains a cache of private creators by creator string
	creators map[string]*tag.PrivateCreator

	mu sync.RWMutex
}

// New creates a new empty DICOM dictionary.
func New() *Dictionary {
	return &Dictionary{
		entries:  make(map[uint32]*Entry),
		keywords: make(map[string]*tag.Tag),
		masked:   make([]*Entry, 0),
		creators: make(map[string]*tag.PrivateCreator),
	}
}

// NewPrivate creates a new private dictionary for a specific creator.
func NewPrivate(creator *tag.PrivateCreator) *Dictionary {
	return &Dictionary{
		entries:        make(map[uint32]*Entry),
		keywords:       make(map[string]*tag.Tag),
		masked:         make([]*Entry, 0),
		privateCreator: creator,
		creators:       make(map[string]*tag.PrivateCreator),
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

// GetPrivateCreator returns or creates a private creator for the given creator string.
//
// This method maintains a cache of private creators to ensure that the same
// PrivateCreator instance is returned for identical creator strings.
// This is important for proper equality comparison of tags with private creators.
//
// The method is thread-safe and handles concurrent access.
func (d *Dictionary) GetPrivateCreator(creator string) *tag.PrivateCreator {
	// Try to read from cache first (read lock)
	d.mu.RLock()
	if pc, ok := d.creators[creator]; ok {
		d.mu.RUnlock()
		return pc
	}
	d.mu.RUnlock()

	// Not found, acquire write lock to add
	d.mu.Lock()
	defer d.mu.Unlock()

	// Double-check after acquiring write lock (another goroutine might have added it)
	if pc, ok := d.creators[creator]; ok {
		return pc
	}

	// Create new private creator and cache it
	pc := tag.NewPrivateCreator(creator)
	d.creators[creator] = pc
	return pc
}

// Global default dictionary instance
var (
	defaultDictionary     *Dictionary
	defaultDictionaryOnce sync.Once
)

// Default returns the default DICOM dictionary.
//
// The default dictionary is lazily initialized on first access.
// It contains standard DICOM tags and can be extended with custom entries.
func Default() *Dictionary {
	defaultDictionaryOnce.Do(func() {
		defaultDictionary = New()
		// Load all standard DICOM dictionary entries from generated data
		initializeDefaultDictionary(defaultDictionary)
	})
	return defaultDictionary
}

// initializeDefaultDictionary adds basic DICOM dictionary entries.
// This function loads all standard DICOM tags from the generated dictionary data.
func initializeDefaultDictionary(d *Dictionary) {
	// Load all standard DICOM dictionary entries from generated data
	loadStandardEntries(d)
}
