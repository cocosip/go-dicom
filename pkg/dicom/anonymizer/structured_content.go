// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package anonymizer

import (
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// isStructuredContentSequence reports whether the sequence carries Content
// Items that PS3.15 E.3.4 requires to be cleaned rather than deleted.
func isStructuredContentSequence(t *tag.Tag) bool {
	if t == nil {
		return false
	}
	return t.Equals(tag.ContentSequence) ||
		t.Equals(tag.AcquisitionContextSequence) ||
		t.Equals(tag.SpecimenPreparationSequence)
}

// cleanStructuredContent implements the Clean Structured Content Option
// (PS3.15 E.3.4) and the sequence dummy ("D") action from Table E.1-1 on a
// sequence of Content Items. The structure is preserved: only Content Items
// whose Concept Name Code appears in Table E.3.4-1 are removed or dummied.
//
// When dummyText is set (Basic Profile "D"), free text and person names of all
// remaining Content Items are replaced with dummy values, so the report
// structure survives without leaking unlisted free text.
func (a *Anonymizer) cleanStructuredContent(seq *dataset.Sequence, dummyText bool, state *anonymizationState) error {
	items := seq.GetItems()
	keep := make([]bool, len(items))
	for i, item := range items {
		if item == nil {
			continue
		}
		// Apply the attribute-level profile to the Content Item first.
		if err := a.anonymizeInPlace(item, false, state); err != nil {
			return err
		}
		action, listed := structuredContentAction(item)
		if listed && action == ActionX {
			continue
		}
		keep[i] = true
		if dummyText || (listed && action == ActionD) {
			a.dummyContentValue(item)
		}
	}
	for i := len(items) - 1; i >= 0; i-- {
		if !keep[i] {
			seq.RemoveItem(i)
		}
	}
	return nil
}

// structuredContentAction resolves the Table E.3.4-1 action for a Content Item
// from its Concept Name Code (Coding Scheme Designator, Code Value) and Value
// Type. Content Items whose code is not listed are retained.
func structuredContentAction(item *dataset.Dataset) (SecurityProfileAction, bool) {
	valueType, _ := item.GetString(tag.ValueType)
	valueType = strings.TrimSpace(valueType)

	names, err := item.GetSequence(tag.ConceptNameCodeSequence)
	if err != nil || names == nil || names.Count() == 0 {
		return 0, false
	}
	code := names.GetItem(0)
	value, _ := code.GetString(tag.CodeValue)
	scheme, _ := code.GetString(tag.CodingSchemeDesignator)
	key := strings.TrimSpace(scheme) + "|" + strings.TrimSpace(value) + "|" + valueType
	action, ok := structuredContentActions[key]
	return action, ok
}

// dummyContentValue replaces the free-text value of a Content Item with a
// non-zero dummy that is valid for its VR (Table E.1-1a "D").
func (a *Anonymizer) dummyContentValue(item *dataset.Dataset) {
	for _, t := range []*tag.Tag{tag.TextValue, tag.PersonName} {
		elem, ok := item.Get(t)
		if !ok {
			continue
		}
		if err := a.replaceString(item, elem, dummyStringValue(elem.ValueRepresentation())); err != nil {
			return
		}
	}
}
