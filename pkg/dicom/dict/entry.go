// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package dict implements DICOM data dictionary functionality.
package dict

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Entry represents a DICOM dictionary entry.
//
// Each entry contains information about a DICOM tag including its name,
// keyword, value representation(s), value multiplicity, and retirement status.
type Entry struct {
	tag                   *tag.Tag
	maskTag               *tag.MaskedTag
	name                  string
	keyword               string
	valueRepresentations  []*vr.VR
	valueMultiplicity     *vm.VM
	isRetired             bool
}

// NewEntry creates a new dictionary entry with a regular tag.
func NewEntry(
	t *tag.Tag,
	name string,
	keyword string,
	vmSpec *vm.VM,
	retired bool,
	vrs ...*vr.VR,
) *Entry {
	// Use tag string if name is empty
	if name == "" {
		name = t.String()
	}
	// Use name as keyword if keyword is empty
	if keyword == "" {
		keyword = name
	}

	return &Entry{
		tag:                  t,
		name:                 name,
		keyword:              keyword,
		valueMultiplicity:    vmSpec,
		valueRepresentations: vrs,
		isRetired:            retired,
	}
}

// NewEntryWithMask creates a new dictionary entry with a masked tag.
func NewEntryWithMask(
	mt *tag.MaskedTag,
	name string,
	keyword string,
	vmSpec *vm.VM,
	retired bool,
	vrs ...*vr.VR,
) *Entry {
	t := mt.Tag()

	// Use tag string if name is empty
	if name == "" {
		name = t.String()
	}
	// Use name as keyword if keyword is empty
	if keyword == "" {
		keyword = name
	}

	return &Entry{
		tag:                  t,
		maskTag:              mt,
		name:                 name,
		keyword:              keyword,
		valueMultiplicity:    vmSpec,
		valueRepresentations: vrs,
		isRetired:            retired,
	}
}

// Tag returns the DICOM tag.
func (e *Entry) Tag() *tag.Tag {
	return e.tag
}

// MaskTag returns the masked tag, or nil if this is not a masked entry.
func (e *Entry) MaskTag() *tag.MaskedTag {
	return e.maskTag
}

// Name returns the descriptive name of the tag.
func (e *Entry) Name() string {
	return e.name
}

// Keyword returns the keyword identifier for the tag.
func (e *Entry) Keyword() string {
	return e.keyword
}

// ValueRepresentations returns the allowed VR types for this tag.
func (e *Entry) ValueRepresentations() []*vr.VR {
	return e.valueRepresentations
}

// ValueMultiplicity returns the VM specification for this tag.
func (e *Entry) ValueMultiplicity() *vm.VM {
	return e.valueMultiplicity
}

// IsRetired returns true if this tag is retired.
func (e *Entry) IsRetired() bool {
	return e.isRetired
}

// IsMasked returns true if this is a masked entry.
func (e *Entry) IsMasked() bool {
	return e.maskTag != nil
}

// Matches checks if the given tag matches this entry.
// For masked entries, uses wildcard matching. For regular entries, uses exact matching.
func (e *Entry) Matches(t *tag.Tag) bool {
	if e.maskTag != nil {
		return e.maskTag.IsMatch(t)
	}
	return e.tag.Equals(t)
}
