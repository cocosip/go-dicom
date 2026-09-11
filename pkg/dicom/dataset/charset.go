// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
	for i, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			return fmt.Errorf("specific character set value[%d] is empty", i)
		}
		if _, ok := charset.GetCharsetInfo(value); !ok {
			return fmt.Errorf("unsupported Specific Character Set %q", value)
		}
		normalized[i] = value
	}

	return ds.AddOrUpdate(element.NewString(tag.SpecificCharacterSet, vr.CS, normalized))
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
