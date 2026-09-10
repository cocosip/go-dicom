// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package uid

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var (
	// ErrNilUID indicates that a registry operation received a nil UID.
	ErrNilUID = errors.New("UID cannot be nil")
	// ErrInvalidUID indicates that a registry operation received an invalid UID value.
	ErrInvalidUID = errors.New("invalid UID")
	// ErrUIDAlreadyRegistered indicates that Register would replace an effective entry.
	ErrUIDAlreadyRegistered = errors.New("UID already registered")
)

// Registry is an isolated application UID catalog layered over the immutable
// standard catalog. Replace explicitly shadows an entry; Unregister masks it
// only in this Registry.
type Registry struct {
	mu      sync.RWMutex
	overlay map[string]*UID
	masked  map[string]struct{}
}

// NewRegistry creates an isolated UID registry that can resolve every standard UID.
func NewRegistry() *Registry {
	return &Registry{
		overlay: make(map[string]*UID),
		masked:  make(map[string]struct{}),
	}
}

// StandardEntries returns an independent, UID-sorted view of the immutable catalog.
func StandardEntries() []*UID {
	entries := make([]*UID, 0, len(generatedStandardUIDEntries)+len(privateUIDEntries))
	entries = append(entries, generatedStandardUIDEntries...)
	entries = append(entries, privateUIDEntries...)
	sortUIDs(entries)
	return entries
}

// Register adds value without replacing an effective standard or custom entry.
func (r *Registry) Register(value *UID) error {
	key, err := registryKey(value)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, found := r.resolveLocked(key); found {
		return fmt.Errorf("%w: %s", ErrUIDAlreadyRegistered, key)
	}
	r.overlay[key] = value
	delete(r.masked, key)
	return nil
}

// Replace adds or replaces value and returns the previously effective entry.
func (r *Registry) Replace(value *UID) (*UID, error) {
	key, err := registryKey(value)
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	previous, _ := r.resolveLocked(key)
	r.overlay[key] = value
	delete(r.masked, key)
	return previous, nil
}

// Unregister removes the effective entry and masks a standard entry in this Registry.
func (r *Registry) Unregister(value string) (*UID, bool) {
	key := normalize(value)
	r.mu.Lock()
	defer r.mu.Unlock()
	previous, found := r.resolveLocked(key)
	if !found {
		return nil, false
	}
	delete(r.overlay, key)
	r.masked[key] = struct{}{}
	return previous, true
}

// Resolve returns the effective known UID after applying this Registry's overlay and masks.
func (r *Registry) Resolve(value string) (*UID, bool) {
	key := normalize(value)
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.resolveLocked(key)
}

// Parse resolves a known UID or constructs an unknown value with the supplied metadata.
func (r *Registry) Parse(s, name string, uidType Type) *UID {
	s = normalize(s)
	if known, found := r.Resolve(s); found {
		return known
	}
	return New(s, name, uidType, false)
}

// Enumerate returns an independent, UID-sorted view of all effective entries.
func (r *Registry) Enumerate() []*UID {
	r.mu.RLock()
	defer r.mu.RUnlock()
	values := make(map[string]*UID, len(generatedStandardUIDEntries)+len(privateUIDEntries)+len(r.overlay))
	for _, value := range generatedStandardUIDEntries {
		if _, masked := r.masked[value.UID()]; !masked {
			values[value.UID()] = value
		}
	}
	for _, value := range privateUIDEntries {
		if _, masked := r.masked[value.UID()]; !masked {
			values[value.UID()] = value
		}
	}
	for key, value := range r.overlay {
		values[key] = value
	}
	entries := make([]*UID, 0, len(values))
	for _, value := range values {
		entries = append(entries, value)
	}
	sortUIDs(entries)
	return entries
}

func (r *Registry) resolveLocked(key string) (*UID, bool) {
	if value, found := r.overlay[key]; found {
		return value, true
	}
	if _, masked := r.masked[key]; masked {
		return nil, false
	}
	return resolveStandard(key)
}

func registryKey(value *UID) (string, error) {
	if value == nil {
		return "", ErrNilUID
	}
	key := normalize(value.UID())
	if !IsValid(key) {
		return "", fmt.Errorf("%w: %q", ErrInvalidUID, value.UID())
	}
	return key, nil
}

func resolveStandard(value string) (*UID, bool) {
	if standard, found := generatedStandardUIDIndex[value]; found {
		return standard, true
	}
	standard, found := privateUIDIndex[value]
	return standard, found
}

func buildUIDIndex(values []*UID) map[string]*UID {
	index := make(map[string]*UID, len(values))
	for _, value := range values {
		index[value.UID()] = value
	}
	return index
}

func sortUIDs(values []*UID) {
	sort.Slice(values, func(i, j int) bool { return values[i].UID() < values[j].UID() })
}
