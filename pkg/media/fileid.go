// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"fmt"
	"regexp"
	"strings"
)

const maxFileIDComponents = 8

var fileIDComponentPattern = regexp.MustCompile(`^[A-Z0-9_]{1,8}$`)

// FileID is a DICOM Referenced File ID containing one to eight components.
type FileID struct {
	components []string
}

// ParseFileID parses a slash- or backslash-separated DICOM Referenced File ID.
func ParseFileID(path string) (FileID, error) {
	normalized := strings.ReplaceAll(path, `\`, "/")
	return NewFileID(strings.Split(normalized, "/")...)
}

// NewFileID creates a validated DICOM Referenced File ID.
func NewFileID(components ...string) (FileID, error) {
	if len(components) == 0 || len(components) > maxFileIDComponents {
		return FileID{}, fmt.Errorf("file ID must contain 1 to %d components, got %d", maxFileIDComponents, len(components))
	}

	validated := make([]string, len(components))
	for i, component := range components {
		if !fileIDComponentPattern.MatchString(component) {
			return FileID{}, fmt.Errorf("file ID component %d must match [A-Z0-9_]{1,8}", i+1)
		}
		validated[i] = component
	}
	return FileID{components: validated}, nil
}

// Components returns a defensive copy of the File ID components.
func (id FileID) Components() []string {
	return append([]string(nil), id.components...)
}

// String returns the DICOM multi-value representation of the File ID.
func (id FileID) String() string {
	return strings.Join(id.components, `\`)
}

func (id FileID) validate() error {
	_, err := NewFileID(id.components...)
	return err
}
