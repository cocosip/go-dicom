// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// CodeItem represents a coded entry in DICOM SR.
//
// A code item consists of:
//   - Code Value: The specific code
//   - Coding Scheme Designator: The coding system (e.g., "DCM", "SNOMED", "LN")
//   - Code Meaning: Human-readable meaning
//   - Coding Scheme Version: Optional version of the coding scheme
//
// Reference: DICOM Part 3, Section 8.8
type CodeItem struct {
	dataset  *dataset.Dataset
	buildErr error
}

// NewCodeItem creates a new code item
func NewCodeItem(value, scheme, meaning string) *CodeItem {
	return NewCodeItemWithVersion(value, scheme, meaning, "")
}

// NewCodeItemWithVersion creates a new code item with a specific version
func NewCodeItemWithVersion(value, scheme, meaning, version string) *CodeItem {
	ds := dataset.New()
	item := &CodeItem{dataset: ds}

	lowerValue := strings.ToLower(value)
	switch {
	case strings.HasPrefix(lowerValue, "http") || strings.HasPrefix(lowerValue, "urn:"):
		item.setBuildError("set URN Code Value", ds.AddOrUpdate(element.NewString(tag.URNCodeValue, vr.UR, []string{value})))
	case utf8.RuneCountInString(value) > 16:
		item.setBuildError("set Long Code Value", ds.AddOrUpdate(element.NewString(tag.LongCodeValue, vr.UC, []string{value})))
	default:
		item.setBuildError("set Code Value", ds.AddOrUpdate(element.NewString(tag.CodeValue, vr.SH, []string{value})))
	}

	if scheme != "" {
		item.setBuildError("set Coding Scheme Designator", ds.AddOrUpdate(element.NewString(tag.CodingSchemeDesignator, vr.SH, []string{scheme})))
	}

	item.setBuildError("set Code Meaning", ds.AddOrUpdate(element.NewString(tag.CodeMeaning, vr.LO, []string{meaning})))

	// Add coding scheme version if provided (0008,0103) VR=SH
	if version != "" {
		item.setBuildError("set Coding Scheme Version", ds.AddOrUpdate(element.NewString(tag.CodingSchemeVersion, vr.SH, []string{version})))
	}

	return item
}

func (c *CodeItem) setBuildError(context string, err error) {
	if c.buildErr == nil && err != nil {
		c.buildErr = WrapError(context, err)
	}
}

// NewCodeItemFromDataset creates a CodeItem from an existing dataset
func NewCodeItemFromDataset(ds *dataset.Dataset) *CodeItem {
	if ds == nil {
		return nil
	}
	return &CodeItem{dataset: ds}
}

// NewCodeItemFromSequence creates a CodeItem from the first item in a sequence
func NewCodeItemFromSequence(t *tag.Tag, ds *dataset.Dataset) (*CodeItem, error) {
	seq, err := ds.GetSequence(t)
	if err != nil {
		return nil, WrapError("sequence not found", err)
	}

	if seq.Count() == 0 {
		return nil, ErrEmptySequence
	}

	return NewCodeItemFromDataset(seq.GetItem(0)), nil
}

// Value returns the code value
func (c *CodeItem) Value() string {
	if c == nil || c.dataset == nil {
		return ""
	}
	for _, valueTag := range []*tag.Tag{tag.CodeValue, tag.LongCodeValue, tag.URNCodeValue} {
		if value, ok := c.dataset.GetString(valueTag); ok {
			return value
		}
	}
	return ""
}

// Scheme returns the coding scheme designator
func (c *CodeItem) Scheme() string {
	if c == nil || c.dataset == nil {
		return ""
	}
	val, _ := c.dataset.GetString(tag.CodingSchemeDesignator)
	return val
}

// Meaning returns the code meaning
func (c *CodeItem) Meaning() string {
	if c == nil || c.dataset == nil {
		return ""
	}
	val, _ := c.dataset.GetString(tag.CodeMeaning)
	return val
}

// Version returns the coding scheme version (may be empty)
func (c *CodeItem) Version() string {
	if c == nil || c.dataset == nil {
		return ""
	}
	val, _ := c.dataset.GetString(tag.CodingSchemeVersion)
	return val
}

// Dataset returns the underlying dataset
func (c *CodeItem) Dataset() *dataset.Dataset {
	if c == nil {
		return nil
	}
	return c.dataset
}

// Equals checks if two code items have the same value, scheme, and scheme version.
// Code meaning is descriptive and does not affect equality.
func (c *CodeItem) Equals(other *CodeItem) bool {
	if c == nil && other == nil {
		return true
	}
	if c == nil || other == nil {
		return false
	}
	return c.Value() == other.Value() && c.Scheme() == other.Scheme() && c.Version() == other.Version()
}

// String returns a string representation of the code item
func (c *CodeItem) String() string {
	if c.Version() != "" {
		return fmt.Sprintf("(%s,%s:%s,\"%s\")", c.Value(), c.Scheme(), c.Version(), c.Meaning())
	}
	return fmt.Sprintf("(%s,%s,\"%s\")", c.Value(), c.Scheme(), c.Meaning())
}
