// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package tag implements DICOM Tag types and operations.
//
// A DICOM tag uniquely identifies a data element. Tags consist of a group
// number and an element number, each represented as a 16-bit unsigned integer.
//
// Example tags:
//   - (0008,0005): Specific Character Set
//   - (0010,0010): Patient Name
//   - (0028,0010): Rows
package tag

import (
	"fmt"
	"strconv"
	"strings"
)

// Tag represents a DICOM tag, which uniquely identifies a data element.
//
// A tag consists of:
//   - Group: 16-bit group number
//   - Element: 16-bit element number
//   - PrivateCreator: Optional private creator for private tags
//
// Tags with odd group numbers are private tags and require a private creator.
type Tag struct {
	group          uint16
	element        uint16
	privateCreator *PrivateCreator
}

// New creates a new DICOM tag with the specified group and element.
func New(group, element uint16) *Tag {
	return &Tag{
		group:   group,
		element: element,
	}
}

// NewWithPrivateCreator creates a new private DICOM tag with a private creator.
func NewWithPrivateCreator(group, element uint16, privateCreator *PrivateCreator) *Tag {
	return &Tag{
		group:          group,
		element:        element,
		privateCreator: privateCreator,
	}
}

// FromUint32 creates a tag from a 32-bit unsigned integer.
// The upper 16 bits represent the group, and the lower 16 bits represent the element.
func FromUint32(value uint32) *Tag {
	return &Tag{
		group:   uint16((value >> 16) & 0xFFFF),
		element: uint16(value & 0xFFFF),
	}
}

// Group returns the tag's group number.
func (t *Tag) Group() uint16 {
	return t.group
}

// Element returns the tag's element number.
func (t *Tag) Element() uint16 {
	return t.element
}

// PrivateCreator returns the tag's private creator, or nil for public tags.
func (t *Tag) PrivateCreator() *PrivateCreator {
	return t.privateCreator
}

// SetPrivateCreator sets the tag's private creator.
func (t *Tag) SetPrivateCreator(pc *PrivateCreator) {
	t.privateCreator = pc
}

// IsPrivate returns true if this is a private tag (odd group number).
func (t *Tag) IsPrivate() bool {
	return (t.group & 1) == 1
}

// ToUint32 converts the tag to a 32-bit unsigned integer.
// The upper 16 bits represent the group, and the lower 16 bits represent the element.
func (t *Tag) ToUint32() uint32 {
	return (uint32(t.group) << 16) | uint32(t.element)
}

// String returns the default string representation of the tag.
// Format: "(GGGG,EEEE)" for public tags, "(GGGG,EEEE:CREATOR)" for private tags.
func (t *Tag) String() string {
	return t.Format("G")
}

// Format returns a formatted string representation of the tag.
//
// Supported formats:
//   - "G" (default): (0028,0010) for public, (0029,1001:MYPRIVATE) for private
//   - "X": (0028,0010) for public, (0029,xx01:MYPRIVATE) for private (masks private element)
//   - "g": 0028,0010 (no parentheses)
//   - "J": 00280010 (no separators)
func (t *Tag) Format(format string) string {
	switch format {
	case "X":
		if t.privateCreator != nil {
			return fmt.Sprintf("(%04x,xx%02x:%s)", t.group, t.element&0xFF, t.privateCreator.Creator())
		}
		return fmt.Sprintf("(%04x,%04x)", t.group, t.element)

	case "g":
		if t.privateCreator != nil {
			return fmt.Sprintf("%04x,%04x:%s", t.group, t.element, t.privateCreator.Creator())
		}
		return fmt.Sprintf("%04x,%04x", t.group, t.element)

	case "J":
		return fmt.Sprintf("%04X%04X", t.group, t.element)

	case "G":
		fallthrough
	default:
		if t.privateCreator != nil {
			return fmt.Sprintf("(%04x,%04x:%s)", t.group, t.element, t.privateCreator.Creator())
		}
		return fmt.Sprintf("(%04x,%04x)", t.group, t.element)
	}
}

// Compare compares this tag to another tag.
// Returns:
//   - negative if this tag < other
//   - zero if this tag == other
//   - positive if this tag > other
//
// Tags are compared by group first, then element, then private creator.
func (t *Tag) Compare(other *Tag) int {
	if t.group != other.group {
		if t.group < other.group {
			return -1
		}
		return 1
	}

	if t.element != other.element {
		if t.element < other.element {
			return -1
		}
		return 1
	}

	// Sort by private creator only if element values are equal
	if t.privateCreator != nil || other.privateCreator != nil {
		if t.privateCreator == nil {
			return -1
		}
		if other.privateCreator == nil {
			return 1
		}
		return t.privateCreator.Compare(other.privateCreator)
	}

	return 0
}

// Equals checks if this tag equals another tag.
// For private tags, only the lower 8 bits of the element are compared.
func (t *Tag) Equals(other *Tag) bool {
	if other == nil {
		return false
	}

	if t.group != other.group {
		return false
	}

	// Handle private tags
	if t.privateCreator != nil || other.privateCreator != nil {
		if t.privateCreator == nil || other.privateCreator == nil {
			return false
		}

		if t.privateCreator.Creator() != other.privateCreator.Creator() {
			return false
		}

		// For private tags, only compare lower 8 bits of element
		return (t.element & 0xFF) == (other.element & 0xFF)
	}

	return t.element == other.element
}

// Hash returns a hash code for the tag.
func (t *Tag) Hash() uint32 {
	if t.privateCreator == nil {
		return t.ToUint32()
	}

	// For private tags, only use lower 8 bits of element
	tagValue := (uint32(t.group) << 16) | uint32(t.element&0xFF)
	// XOR with private creator hash
	return tagValue ^ t.privateCreator.Hash()
}

// Parse parses a tag from a string.
//
// Supported formats:
//   - "(0028,0010)" - standard format with parentheses
//   - "0028,0010" - without parentheses
//   - "00280010" - compact format
//   - "(0029,1001:MYPRIVATE)" - private tag with creator
//
// Note: Private creator lookup requires a dictionary, which is not available yet.
// For now, private tags are parsed but the creator is stored as-is.
func Parse(s string) (*Tag, error) {
	if len(s) < 8 {
		return nil, fmt.Errorf("tag string too short: '%s' (expected at least 8 characters)", s)
	}

	// Remove spaces
	s = strings.TrimSpace(s)

	pos := 0

	// Skip opening parenthesis if present
	if s[pos] == '(' {
		pos++
	}

	// Parse group (4 hex digits)
	if pos+4 > len(s) {
		return nil, fmt.Errorf("invalid tag format: '%s'", s)
	}
	group, err := strconv.ParseUint(s[pos:pos+4], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid group number: %w", err)
	}
	pos += 4

	// Skip comma if present
	if pos < len(s) && s[pos] == ',' {
		pos++
	}

	// Parse element (4 hex digits)
	if pos+4 > len(s) {
		return nil, fmt.Errorf("invalid tag format: '%s'", s)
	}
	element, err := strconv.ParseUint(s[pos:pos+4], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid element number: %w", err)
	}
	pos += 4

	tag := &Tag{
		group:   uint16(group),
		element: uint16(element),
	}

	// Check for private creator
	if pos < len(s) && s[pos] == ':' {
		pos++

		// Extract creator string
		creatorStr := s[pos:]

		creatorStr = strings.TrimSuffix(creatorStr, ")")

		// Look up private creator in dictionary if available
		if globalPrivateCreatorLookup != nil {
			tag.privateCreator = globalPrivateCreatorLookup(creatorStr)
		} else {
			// Fallback: create a simple private creator if dictionary not initialized
			tag.privateCreator = NewPrivateCreator(creatorStr)
		}
	}

	return tag, nil
}

// MustParse parses a tag from a string and panics if there is an error.
// Use this only for static/constant tag values.
func MustParse(s string) *Tag {
	tag, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("invalid tag string '%s': %v", s, err))
	}
	return tag
}

// Standard tag constants
var (
	// Unknown represents an unknown tag (FFFF,FFFF)
	Unknown = New(0xFFFF, 0xFFFF)
)

// DictionaryEntryLookup is a function type for looking up dictionary entries.
// This is used to avoid circular dependencies between tag and dict packages.
type DictionaryEntryLookup func(*Tag) interface{}

// globalDictionaryLookup holds the function to lookup dictionary entries.
// This is set by the dict package during initialization.
var globalDictionaryLookup DictionaryEntryLookup

// SetDictionaryLookup sets the global dictionary lookup function.
// This is called by the dict package to register the lookup function.
func SetDictionaryLookup(lookup DictionaryEntryLookup) {
	globalDictionaryLookup = lookup
}

// DictionaryEntry returns the dictionary entry for this tag.
//
// The dictionary entry contains metadata about the tag including its name,
// keyword, value representation(s), value multiplicity, and retirement status.
//
// Returns nil if the tag is not found in the dictionary or if the dictionary
// has not been initialized.
//
// Note: The returned interface{} should be type-asserted to *dict.Entry.
// This design avoids circular dependencies between tag and dict packages.
func (t *Tag) DictionaryEntry() interface{} {
	if globalDictionaryLookup == nil {
		return nil
	}
	return globalDictionaryLookup(t)
}

// Uint32 returns the tag as a 32-bit unsigned integer.
// This is an alias for ToUint32() for convenience.
func (t *Tag) Uint32() uint32 {
	return t.ToUint32()
}

// KeywordLookup is a function type for looking up tags by keyword.
// This is used to avoid circular dependencies between tag and dict packages.
type KeywordLookup func(keyword string) (*Tag, error)

// globalKeywordLookup holds the function to lookup tags by keyword.
// This is set by the dict package during initialization.
var globalKeywordLookup KeywordLookup

// SetKeywordLookup sets the global keyword lookup function.
// This is called by the dict package to register the lookup function.
func SetKeywordLookup(lookup KeywordLookup) {
	globalKeywordLookup = lookup
}

// PrivateCreatorLookup is a function type for looking up or creating private creators.
// This is used to avoid circular dependencies between tag and dict packages.
type PrivateCreatorLookup func(creator string) *PrivateCreator

// globalPrivateCreatorLookup holds the function to lookup/create private creators.
// This is set by the dict package during initialization.
var globalPrivateCreatorLookup PrivateCreatorLookup

// SetPrivateCreatorLookup sets the global private creator lookup function.
// This is called by the dict package to register the lookup function.
func SetPrivateCreatorLookup(lookup PrivateCreatorLookup) {
	globalPrivateCreatorLookup = lookup
}

// ParseKeyword parses a tag from its DICOM keyword.
//
// Examples:
//   - "PatientName" -> (0010,0010)
//   - "Rows" -> (0028,0010)
//   - "PixelData" -> (7FE0,0010)
//
// Returns an error if the keyword is not found in the dictionary or if the
// dictionary has not been initialized.
func ParseKeyword(keyword string) (Tag, error) {
	if globalKeywordLookup == nil {
		return Tag{}, fmt.Errorf("keyword lookup not available (dictionary not initialized)")
	}

	t, err := globalKeywordLookup(keyword)
	if err != nil {
		return Tag{}, err
	}
	if t == nil {
		return Tag{}, fmt.Errorf("keyword not found: %s", keyword)
	}

	return *t, nil
}
