// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	customDictionaryTagValue = "77760010"
	customDictionaryValue    = "custom value"
	firstCustomKeyword       = "FirstCustomKeyword"
	secondCustomKeyword      = "SecondCustomKeyword"
)

func TestJSONLookupControlsKeywordWriteAndReadWithoutGlobalState(t *testing.T) {
	customTag := tag.New(0x7776, 0x0010)
	first := newSerializationTestDictionary(customTag, firstCustomKeyword)
	second := newSerializationTestDictionary(customTag, secondCustomKeyword)
	ds := dataset.New()
	if err := ds.Add(element.NewString(customTag, vr.LO, []string{customDictionaryValue})); err != nil {
		t.Fatal(err)
	}

	firstJSON, err := ToJSON(ds, WithWriteTagsAsKeywords(true), WithJSONLookup(first))
	if err != nil {
		t.Fatalf("ToJSON(first lookup) error = %v", err)
	}
	if !strings.Contains(string(firstJSON), `"`+firstCustomKeyword+`"`) {
		t.Fatalf("ToJSON(first lookup) = %s, want keyword %q", firstJSON, firstCustomKeyword)
	}

	secondJSON, err := ToJSON(ds, WithWriteTagsAsKeywords(true), WithJSONLookup(second))
	if err != nil {
		t.Fatalf("ToJSON(second lookup) error = %v", err)
	}
	if !strings.Contains(string(secondJSON), `"`+secondCustomKeyword+`"`) ||
		strings.Contains(string(secondJSON), `"`+firstCustomKeyword+`"`) {
		t.Fatalf("ToJSON(second lookup) = %s, want only keyword %q", secondJSON, secondCustomKeyword)
	}

	defaultJSON, err := ToJSON(ds, WithWriteTagsAsKeywords(true))
	if err != nil {
		t.Fatalf("ToJSON(default lookup) error = %v", err)
	}
	if !strings.Contains(string(defaultJSON), `"`+customDictionaryTagValue+`"`) {
		t.Fatalf("ToJSON(default lookup) = %s, want numeric tag %q", defaultJSON, customDictionaryTagValue)
	}

	restored, err := FromJSON(firstJSON, WithJSONLookup(first))
	if err != nil {
		t.Fatalf("FromJSON(first lookup) error = %v", err)
	}
	if got := restored.TryGetString(customTag); got != customDictionaryValue {
		t.Fatalf("FromJSON(first lookup) value = %q, want %q", got, customDictionaryValue)
	}
}

func TestXMLLookupControlsKeywordWriteAndKeywordOnlyRead(t *testing.T) {
	customTag := tag.New(0x7776, 0x0010)
	first := newSerializationTestDictionary(customTag, firstCustomKeyword)
	second := newSerializationTestDictionary(customTag, secondCustomKeyword)
	ds := dataset.New()
	if err := ds.Add(element.NewString(customTag, vr.LO, []string{customDictionaryValue})); err != nil {
		t.Fatal(err)
	}

	firstXML, err := ToXML(ds, WithXMLLookup(first))
	if err != nil {
		t.Fatalf("ToXML(first lookup) error = %v", err)
	}
	if !strings.Contains(string(firstXML), `keyword="`+firstCustomKeyword+`"`) {
		t.Fatalf("ToXML(first lookup) = %s, want keyword %q", firstXML, firstCustomKeyword)
	}

	secondXML, err := ToXML(ds, WithXMLLookup(second))
	if err != nil {
		t.Fatalf("ToXML(second lookup) error = %v", err)
	}
	if !strings.Contains(string(secondXML), `keyword="`+secondCustomKeyword+`"`) ||
		strings.Contains(string(secondXML), `keyword="`+firstCustomKeyword+`"`) {
		t.Fatalf("ToXML(second lookup) = %s, want only keyword %q", secondXML, secondCustomKeyword)
	}

	keywordOnly := []byte(`<NativeDicomModel><DicomAttribute keyword="` + firstCustomKeyword +
		`" vr="LO"><Value number="1">` + customDictionaryValue + `</Value></DicomAttribute></NativeDicomModel>`)
	restored, err := FromXML(keywordOnly, WithXMLLookup(first))
	if err != nil {
		t.Fatalf("FromXML(keyword-only, first lookup) error = %v", err)
	}
	if got := restored.TryGetString(customTag); got != customDictionaryValue {
		t.Fatalf("FromXML(keyword-only, first lookup) value = %q, want %q", got, customDictionaryValue)
	}
}

func newSerializationTestDictionary(customTag *tag.Tag, keyword string) *dict.Dictionary {
	lookup := dict.New()
	lookup.Add(dict.NewEntry(customTag, "Custom Name", keyword, vm.VM1, false, vr.LO))
	return lookup
}
