// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// PersonNameComponents represents the components of a DICOM person name.
// Format: FamilyName^GivenName^MiddleName^NamePrefix^NameSuffix
type PersonNameComponents struct {
	FamilyName string // 姓
	GivenName  string // 名
	MiddleName string // 中间名
	NamePrefix string // 前缀 (如 Dr., Prof.)
	NameSuffix string // 后缀 (如 Jr., III)
}

// String returns the formatted person name string.
func (p *PersonNameComponents) String() string {
	return strings.Join([]string{
		p.FamilyName,
		p.GivenName,
		p.MiddleName,
		p.NamePrefix,
		p.NameSuffix,
	}, "^")
}

// PersonName represents a DICOM element with VR = PN (Person Name).
// Person names have special structure: FamilyName^GivenName^MiddleName^NamePrefix^NameSuffix
type PersonName struct {
	*String
}

// NewPersonName creates a new PN element with the given names.
func NewPersonName(t *tag.Tag, names []string) *PersonName {
	str := NewString(t, vr.PN, names)
	return &PersonName{String: str}
}

// NewPersonNameWithComponents creates a new PN element from person name components.
func NewPersonNameWithComponents(t *tag.Tag, components []*PersonNameComponents) *PersonName {
	names := make([]string, len(components))
	for i, c := range components {
		if c != nil {
			names[i] = c.String()
		}
	}
	return NewPersonName(t, names)
}

// NewPersonNameFromBuffer creates a PN element from an existing buffer.
func NewPersonNameFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *PersonName {
	str := NewStringFromBuffer(t, vr.PN, buf, enc)
	return &PersonName{String: str}
}

// GetComponents returns the person name components at the specified index.
func (p *PersonName) GetComponents(index int) (*PersonNameComponents, error) {
	name := p.GetValue(index)
	if name == "" {
		return nil, fmt.Errorf("person name at index %d is empty", index)
	}

	parts := strings.Split(name, "^")
	components := &PersonNameComponents{}

	if len(parts) > 0 {
		components.FamilyName = strings.TrimSpace(parts[0])
	}
	if len(parts) > 1 {
		components.GivenName = strings.TrimSpace(parts[1])
	}
	if len(parts) > 2 {
		components.MiddleName = strings.TrimSpace(parts[2])
	}
	if len(parts) > 3 {
		components.NamePrefix = strings.TrimSpace(parts[3])
	}
	if len(parts) > 4 {
		components.NameSuffix = strings.TrimSpace(parts[4])
	}

	return components, nil
}

// GetAllComponents returns all person name components.
func (p *PersonName) GetAllComponents() ([]*PersonNameComponents, error) {
	count := p.Count()
	if count == 0 {
		return nil, nil
	}

	components := make([]*PersonNameComponents, count)
	for i := 0; i < count; i++ {
		c, err := p.GetComponents(i)
		if err != nil {
			// If parsing fails, skip this entry
			continue
		}
		components[i] = c
	}

	return components, nil
}

// GetFormattedName returns a formatted name string (Given Middle Family, Prefix Suffix).
func (p *PersonName) GetFormattedName(index int) (string, error) {
	components, err := p.GetComponents(index)
	if err != nil {
		return "", err
	}

	// Format: "Given Middle Family, Prefix Suffix"
	var parts []string

	// First part: Given Middle Family
	var firstName []string
	if components.NamePrefix != "" {
		firstName = append(firstName, components.NamePrefix)
	}
	if components.GivenName != "" {
		firstName = append(firstName, components.GivenName)
	}
	if components.MiddleName != "" {
		firstName = append(firstName, components.MiddleName)
	}
	if components.FamilyName != "" {
		firstName = append(firstName, components.FamilyName)
	}

	if len(firstName) > 0 {
		parts = append(parts, strings.Join(firstName, " "))
	}

	// Add suffix if present
	if components.NameSuffix != "" {
		parts = append(parts, components.NameSuffix)
	}

	return strings.Join(parts, ", "), nil
}
