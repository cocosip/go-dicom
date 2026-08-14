// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict

import (
	"sort"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/dictif"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// dictionaryAdapter wraps Dictionary to implement dictif.Lookup interface.
// This adapter is needed because Go's interface matching requires exact method signatures,
// and we can't change existing Dictionary methods without breaking compatibility.
type dictionaryAdapter struct {
	*Dictionary
}

// LookupKeyword implements dictif.Lookup by wrapping Dictionary.LookupKeyword.
func (da *dictionaryAdapter) LookupKeyword(keyword string) dictif.Tag {
	tag := da.Dictionary.LookupKeyword(keyword)
	if tag == nil {
		return nil
	}
	return tag // *tag.Tag implements dictif.Tag
}

// GetPrivateCreator implements dictif.Lookup by wrapping Dictionary.GetPrivateCreator.
func (da *dictionaryAdapter) GetPrivateCreator(creator string) dictif.PrivateCreator {
	pc := da.Dictionary.GetPrivateCreator(creator)
	if pc == nil {
		return nil
	}
	return pc // *tag.PrivateCreator implements dictif.PrivateCreator
}

// init registers the default dictionary as the global lookup implementation.
func init() {
	dictif.SetGlobalLookup(&dictionaryAdapter{Dictionary: Default()})
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

	// privateDictionaries keeps entries isolated by private creator.
	privateDictionaries map[string]*Dictionary

	mu sync.RWMutex
}

// New creates a new empty DICOM dictionary.
func New() *Dictionary {
	return &Dictionary{
		entries:             make(map[uint32]*Entry),
		keywords:            make(map[string]*tag.Tag),
		masked:              make([]*Entry, 0),
		creators:            make(map[string]*tag.PrivateCreator),
		privateDictionaries: make(map[string]*Dictionary),
	}
}

// NewPrivate creates a new private dictionary for a specific creator.
func NewPrivate(creator *tag.PrivateCreator) *Dictionary {
	return &Dictionary{
		entries:             make(map[uint32]*Entry),
		keywords:            make(map[string]*tag.Tag),
		masked:              make([]*Entry, 0),
		privateCreator:      creator,
		creators:            make(map[string]*tag.PrivateCreator),
		privateDictionaries: make(map[string]*Dictionary),
	}
}

// Add adds a dictionary entry.
func (d *Dictionary) Add(entry *Entry) {
	if d.privateCreator == nil {
		if creator := entry.Tag().PrivateCreator(); creator != nil {
			d.GetPrivateDictionary(creator.Creator()).Add(entry)
			return
		}
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if d.privateCreator != nil {
		entry = cloneEntryWithPrivateCreator(entry, d.privateCreator)
	}

	// If this is a masked entry, add to masked list
	if entry.IsMasked() {
		for i, existing := range d.masked {
			if existing.MaskTag().Card() == entry.MaskTag().Card() &&
				existing.MaskTag().Mask() == entry.MaskTag().Mask() {
				d.deleteKeywordLocked(existing)
				d.masked = append(d.masked[:i], d.masked[i+1:]...)
				break
			}
		}

		// Masked entries use last-added-wins precedence, matching fo-dicom.
		d.masked = append(d.masked, nil)
		copy(d.masked[1:], d.masked)
		d.masked[0] = entry
	} else {
		// Add to exact match map
		key := d.entryKey(entry.Tag())
		if existing, ok := d.entries[key]; ok {
			d.deleteKeywordLocked(existing)
		}
		d.entries[key] = entry
	}

	// Add keyword mapping
	if entry.Keyword() != "" {
		d.keywords[entry.Keyword()] = entry.Tag()
	}
}

func cloneEntryWithPrivateCreator(entry *Entry, creator *tag.PrivateCreator) *Entry {
	cloned := *entry
	clonedTag := tag.NewWithPrivateCreator(entry.Tag().Group(), entry.Tag().Element(), creator)
	cloned.tag = clonedTag
	if entry.IsMasked() {
		cloned.maskTag = tag.NewMaskedTagWithMask(clonedTag, entry.MaskTag().Mask())
	}
	cloned.valueRepresentations = append(
		cloned.valueRepresentations[:0:0],
		cloned.valueRepresentations...,
	)
	return &cloned
}

func (d *Dictionary) entryKey(t *tag.Tag) uint32 {
	if d.privateCreator != nil {
		return (uint32(t.Group()) << 16) | uint32(t.Element()&0x00FF)
	}
	return t.ToUint32()
}

func (d *Dictionary) deleteKeywordLocked(entry *Entry) {
	keyword := entry.Keyword()
	if keyword == "" {
		return
	}
	mapped, ok := d.keywords[keyword]
	if ok && mapped.Equals(entry.Tag()) {
		delete(d.keywords, keyword)
	}
}

// Lookup finds a dictionary entry for the given tag.
//
// Returns nil if the tag is not found.
func (d *Dictionary) Lookup(t *tag.Tag) *Entry {
	if t == nil {
		return nil
	}
	if creator := t.PrivateCreator(); creator != nil {
		if d.privateCreator != nil {
			if creator.Creator() != d.privateCreator.Creator() {
				return nil
			}
		} else {
			d.mu.RLock()
			privateDictionary := d.privateDictionaries[creator.Creator()]
			d.mu.RUnlock()
			if privateDictionary == nil {
				return nil
			}
			return privateDictionary.Lookup(t)
		}
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	// Try exact match first
	if entry, ok := d.entries[d.entryKey(t)]; ok {
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
	if t, ok := d.keywords[keyword]; ok {
		d.mu.RUnlock()
		return t
	}

	creators := make([]string, 0, len(d.privateDictionaries))
	for creator := range d.privateDictionaries {
		creators = append(creators, creator)
	}
	privateDictionaries := make(map[string]*Dictionary, len(d.privateDictionaries))
	for creator, privateDictionary := range d.privateDictionaries {
		privateDictionaries[creator] = privateDictionary
	}
	d.mu.RUnlock()

	sort.Strings(creators)
	for _, creator := range creators {
		if t := privateDictionaries[creator].LookupKeyword(keyword); t != nil {
			return t
		}
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

// GetPrivateDictionary returns or creates the dictionary for a private creator.
// Repeated calls with the same creator return the same dictionary.
func (d *Dictionary) GetPrivateDictionary(creator string) *Dictionary {
	privateCreator := d.GetPrivateCreator(creator)

	d.mu.RLock()
	if privateDictionary, ok := d.privateDictionaries[creator]; ok {
		d.mu.RUnlock()
		return privateDictionary
	}
	d.mu.RUnlock()

	d.mu.Lock()
	defer d.mu.Unlock()
	if privateDictionary, ok := d.privateDictionaries[creator]; ok {
		return privateDictionary
	}

	privateDictionary := NewPrivate(privateCreator)
	d.privateDictionaries[creator] = privateDictionary
	return privateDictionary
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

// initializeDefaultDictionary loads the bundled standard and private entries.
func initializeDefaultDictionary(d *Dictionary) {
	loadStandardEntries(d)
	loadPrivateEntries(d)
}

// LookupTag implements dictif.Lookup interface.
// It converts dictif.Tag to *tag.Tag and returns dictif.Entry.
func (d *Dictionary) LookupTag(t dictif.Tag) dictif.Entry {
	tagPtr, ok := t.(*tag.Tag)
	if !ok {
		// Convert an interface-only tag when no private creator information exists.
		tagPtr = tag.New(t.Group(), t.Element())
	}

	// Use existing Lookup method
	entry := d.Lookup(tagPtr)
	if entry == nil {
		return nil
	}

	// Entry already implements dictif.Entry
	return entry
}
