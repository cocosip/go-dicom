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
		// PS3.3 C.12.1.1: Value 1 is the replacement character repertoire and
		// must be a single-byte character set; the multi-byte sets can only be
		// used as code extensions.
		if isMultiByteOnlyCharacterSet(nonEmpty[0]) {
			return fmt.Errorf("specific character set %q cannot be value 1 when code extensions are used", nonEmpty[0])
		}
		for _, value := range nonEmpty[1:] {
			if !strings.HasPrefix(value, "ISO 2022 ") {
				return fmt.Errorf("specific character set %q cannot follow value 1; only ISO 2022 code extensions may", value)
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
	if declared := ds.SpecificCharacterSet(); len(declared) > 0 {
		if len(declared) == 1 && declared[0] == "ISO_IR 192" {
			return nil
		}
		// Re-labeling would reinterpret every value that is still encoded with
		// the declared character set (PS3.5 6.1.2.2), so the caller has to
		// resolve the mixed encoding instead.
		return fmt.Errorf("dataset declares Specific Character Set %v but holds values encoded as UTF-8; convert those values or declare ISO_IR 192", declared)
	}
	return ds.SetSpecificCharacterSet("ISO_IR 192")
}

func (ds *Dataset) needsUTF8() bool {
	return ds.needsUTF8Visited(make(map[*Dataset]struct{}))
}

func (ds *Dataset) needsUTF8Visited(visited map[*Dataset]struct{}) bool {
	if ds == nil {
		return false
	}
	if _, ok := visited[ds]; ok {
		return false
	}
	visited[ds] = struct{}{}
	for _, elem := range ds.Elements() {
		if element.NeedsCharacterSet(elem) {
			return true
		}
		sequence, ok := elem.(*Sequence)
		if !ok {
			continue
		}
		for i := 0; i < sequence.Count(); i++ {
			if item := sequence.GetItem(i); item != nil && item.needsUTF8Visited(visited) {
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

// isMultiByteOnlyCharacterSet reports the code extension sets that cannot be
// used as Value 1 of Specific Character Set (PS3.3 C.12.1.1).
func isMultiByteOnlyCharacterSet(value string) bool {
	switch value {
	case "ISO 2022 IR 87", "ISO 2022 IR 149", "ISO 2022 IR 58", "ISO 2022 IR 159":
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
