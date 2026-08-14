// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package dicomxml reads the merged DICOM dictionary XML used by generators.
package dicomxml

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Dictionary is the merged standard dictionary document.
type Dictionary struct {
	XMLName xml.Name `xml:"dictionary"`
	Version string   `xml:"version,attr"`
	Tags    []Tag    `xml:"tag"`
	UIDs    []UID    `xml:"uid"`
}

// PrivateDictionaries is the private dictionary collection document.
type PrivateDictionaries struct {
	XMLName      xml.Name            `xml:"dictionaries"`
	Dictionaries []PrivateDictionary `xml:"dictionary"`
}

// PrivateDictionary contains entries owned by one Private Creator.
type PrivateDictionary struct {
	Creator string `xml:"creator,attr"`
	Tags    []Tag  `xml:"tag"`
}

// Tag is a standard DICOM data element definition.
type Tag struct {
	Group   string `xml:"group,attr"`
	Element string `xml:"element,attr"`
	Keyword string `xml:"keyword,attr"`
	VR      string `xml:"vr,attr"`
	VM      string `xml:"vm,attr"`
	Retired string `xml:"retired,attr"`
	Name    string `xml:",chardata"`
}

// UID is a standard DICOM unique identifier definition.
type UID struct {
	Value   string `xml:"uid,attr"`
	Keyword string `xml:"keyword,attr"`
	Type    string `xml:"type,attr"`
	Retired string `xml:"retired,attr"`
	Name    string `xml:",chardata"`
}

// Load reads a merged DICOM dictionary XML document from path.
func Load(path string) (*Dictionary, error) {
	file, err := os.Open(path) // #nosec G304 -- path is supplied by the generator user.
	if err != nil {
		return nil, fmt.Errorf("open dictionary XML %q: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	var dictionary Dictionary
	if err := xml.NewDecoder(file).Decode(&dictionary); err != nil {
		return nil, fmt.Errorf("parse dictionary XML %q: %w", path, err)
	}
	if dictionary.XMLName.Local != "dictionary" {
		return nil, fmt.Errorf("parse dictionary XML %q: expected <dictionary> root", path)
	}
	return &dictionary, nil
}

// LoadPrivate reads a private dictionary collection XML document from path.
func LoadPrivate(path string) (*PrivateDictionaries, error) {
	file, err := os.Open(path) // #nosec G304 -- path is supplied by the generator user.
	if err != nil {
		return nil, fmt.Errorf("open private dictionary XML %q: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	var dictionaries PrivateDictionaries
	if err := xml.NewDecoder(file).Decode(&dictionaries); err != nil {
		return nil, fmt.Errorf("parse private dictionary XML %q: %w", path, err)
	}
	if dictionaries.XMLName.Local != "dictionaries" {
		return nil, fmt.Errorf("parse private dictionary XML %q: expected <dictionaries> root", path)
	}
	return &dictionaries, nil
}
