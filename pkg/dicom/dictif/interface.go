// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package dictif defines interfaces for DICOM dictionary lookups.
//
// This package exists to break circular dependencies between the tag and dict packages.
// Instead of having tag depend on dict or using function callbacks, we define interfaces
// that dict implements and tag uses.
package dictif

// Tag represents a minimal interface for DICOM tags.
// This avoids importing the full tag package in the interface definition.
type Tag interface {
	// Group returns the tag's group number
	Group() uint16
	// Element returns the tag's element number
	Element() uint16
	// ToUint32 converts the tag to a 32-bit unsigned integer
	ToUint32() uint32
}

// Entry represents a dictionary entry with tag metadata.
type Entry interface {
	// Name returns the human-readable name
	Name() string
	// Keyword returns the DICOM keyword
	Keyword() string
	// VRs returns the allowed Value Representations
	VRs() []string
	// VM returns the Value Multiplicity
	VM() string
}

// PrivateCreator represents a private tag creator.
type PrivateCreator interface {
	// Creator returns the private creator string
	Creator() string
}

// Lookup defines the interface for DICOM dictionary lookups.
//
// This interface is implemented by the Dictionary type in the dict package
// and used by the tag package to look up tag metadata without creating
// circular dependencies.
type Lookup interface {
	// LookupTag returns the dictionary entry for a tag, or nil if not found
	LookupTag(tag Tag) Entry

	// LookupKeyword returns the tag for a keyword, or nil if not found
	LookupKeyword(keyword string) Tag

	// GetPrivateCreator returns or creates a private creator for the given string
	GetPrivateCreator(creator string) PrivateCreator
}

// globalLookup holds the global dictionary lookup implementation.
// This is set by the dict package during initialization.
var globalLookup Lookup

// SetGlobalLookup sets the global dictionary lookup implementation.
// This should be called once during initialization by the dict package.
func SetGlobalLookup(lookup Lookup) {
	globalLookup = lookup
}

// GlobalLookup returns the global dictionary lookup implementation.
// Returns nil if not yet initialized.
func GlobalLookup() Lookup {
	return globalLookup
}
