// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package dataset

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"golang.org/x/text/encoding"
)

// SetSpecificCharacterSet validates and writes the DICOM Specific Character
// Set value used when subsequently creating text elements in this Dataset.
// Passing no values removes the element and restores the DICOM default.
func (ds *Dataset) SetSpecificCharacterSet(values ...string) error {
	if ds == nil {
		return fmt.Errorf("cannot set Specific Character Set on nil Dataset")
	}
	if len(values) == 0 {
		ds.Remove(tag.SpecificCharacterSet)
		return nil
	}

	normalized := make([]string, len(values))
	nonEmpty := make([]string, 0, len(values))
	for i, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			normalized[i] = value
			continue
		}
		if _, ok := charset.GetCharsetInfo(value); !ok {
			return fmt.Errorf("unsupported Specific Character Set %q", value)
		}
		normalized[i] = value
		nonEmpty = append(nonEmpty, value)
	}
	if len(nonEmpty) > 1 {
		for _, value := range nonEmpty {
			if isSingleValueCharacterSet(value) {
				return fmt.Errorf("specific character set %q cannot be combined with extensions", value)
			}
		}
	}

	return ds.AddOrUpdate(element.NewString(tag.SpecificCharacterSet, vr.CS, normalized))
}

// EnsureUTF8CharacterSet declares ISO_IR 192 (UTF-8) when the Dataset, or a
// nested sequence item, holds a string value that element.NewString had to
// encode as UTF-8 fallback bytes. Writing those bytes without a matching
// (0008,0005) is non-conformant (PS3.3 C.12.1, Table C.12-1) and corrupts the
// value for every conformant reader.
func (ds *Dataset) EnsureUTF8CharacterSet() error {
	if ds == nil || !ds.needsUTF8() {
		return nil
	}
	if declared := ds.SpecificCharacterSet(); len(declared) == 1 && declared[0] == "ISO_IR 192" {
		return nil
	}
	return ds.SetSpecificCharacterSet("ISO_IR 192")
}

func (ds *Dataset) needsUTF8() bool {
	for _, elem := range ds.Elements() {
		if element.NeedsCharacterSet(elem) {
			return true
		}
		sequence, ok := elem.(*Sequence)
		if !ok {
			continue
		}
		for i := 0; i < sequence.Count(); i++ {
			if item := sequence.GetItem(i); item != nil && item.needsUTF8() {
				return true
			}
		}
	}
	return false
}

func isSingleValueCharacterSet(value string) bool {
	switch value {
	case "ISO_IR 192", "GB18030", "GBK":
		return true
	default:
		return false
	}
}

// SpecificCharacterSet returns the declared DICOM Specific Character Set
// values. An empty result means the DICOM default character repertoire.
func (ds *Dataset) SpecificCharacterSet() []string {
	values, ok := ds.GetStrings(tag.SpecificCharacterSet)
	if !ok {
		return nil
	}
	return values
}

func (ds *Dataset) effectiveTextEncodings() []encoding.Encoding {
	if declared := ds.SpecificCharacterSet(); len(declared) > 0 {
		return charset.GetEncodings(declared)
	}
	if inherited := ds.InternalTextEncodings(); len(inherited) > 0 {
		return inherited
	}
	return charset.GetEncodings(nil)
}

func (ds *Dataset) inheritElementContext(elem element.Element) {
	sequence, ok := elem.(*Sequence)
	if !ok {
		return
	}
	encodings := ds.effectiveTextEncodings()
	visited := map[*Dataset]struct{}{ds: {}}
	for i := 0; i < sequence.Count(); i++ {
		item := sequence.GetItem(i)
		if item == nil {
			continue
		}
		item.setInternalTextEncodings(encodings, visited)
	}
}

func (ds *Dataset) refreshChildTextEncodings() {
	visited := map[*Dataset]struct{}{ds: {}}
	ds.refreshChildTextEncodingsWithVisited(visited)
}

func (ds *Dataset) refreshChildTextEncodingsWithVisited(visited map[*Dataset]struct{}) {
	encodings := ds.effectiveTextEncodings()
	for _, elem := range ds.items {
		if sequence, ok := elem.(*Sequence); ok {
			for i := 0; i < sequence.Count(); i++ {
				if item := sequence.GetItem(i); item != nil {
					item.setInternalTextEncodings(encodings, visited)
				}
			}
		}
	}
}
