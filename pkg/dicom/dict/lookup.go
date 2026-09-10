// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict

import "github.com/cocosip/go-dicom/pkg/dicom/tag"

// Lookup provides typed DICOM dictionary queries to serializers and other consumers.
type Lookup interface {
	Lookup(*tag.Tag) *Entry
	LookupKeyword(string) *tag.Tag
	GetPrivateCreator(string) *tag.PrivateCreator
}
