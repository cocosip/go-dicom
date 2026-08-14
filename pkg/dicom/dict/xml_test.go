// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dict_test

import (
	"strings"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewFromXMLLoadsStandardDictionary(t *testing.T) {
	const input = `<?xml version="1.0" encoding="UTF-8"?>
<dictionary version="2026b">
  <tag group="0010" element="0010" keyword="PatientName" vr="PN" vm="1">Patient's Name</tag>
  <tag group="0028" element="0106" keyword="SmallestImagePixelValue" vr="US/SS" vm="1" retired="true">Smallest Image Pixel Value</tag>
  <tag group="FFFE" element="E000" keyword="Item" vm="1">Item</tag>
</dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}

	patientName := d.Lookup(tag.New(0x0010, 0x0010))
	if patientName == nil {
		t.Fatal("Lookup(PatientName) = nil, want entry")
	}
	if patientName.Name() != "Patient's Name" {
		t.Errorf("PatientName.Name() = %q, want %q", patientName.Name(), "Patient's Name")
	}
	if patientName.Keyword() != "PatientName" {
		t.Errorf("PatientName.Keyword() = %q, want PatientName", patientName.Keyword())
	}
	if patientName.VM() != "1" {
		t.Errorf("PatientName.VM() = %q, want 1", patientName.VM())
	}
	if got := patientName.VRs(); len(got) != 1 || got[0] != "PN" {
		t.Errorf("PatientName.VRs() = %v, want [PN]", got)
	}

	pixelValue := d.Lookup(tag.New(0x0028, 0x0106))
	if pixelValue == nil {
		t.Fatal("Lookup(SmallestImagePixelValue) = nil, want entry")
	}
	if !pixelValue.IsRetired() {
		t.Error("SmallestImagePixelValue.IsRetired() = false, want true")
	}
	if got := pixelValue.VRs(); len(got) != 2 || got[0] != "US" || got[1] != "SS" {
		t.Errorf("SmallestImagePixelValue.VRs() = %v, want [US SS]", got)
	}

	item := d.Lookup(tag.New(0xFFFE, 0xE000))
	if item == nil {
		t.Fatal("Lookup(Item) = nil, want entry")
	}
	if got := item.VRs(); len(got) != 1 || got[0] != "NONE" {
		t.Errorf("Item.VRs() = %v, want [NONE]", got)
	}
}

func TestNewFromXMLAcceptsUTF8BOM(t *testing.T) {
	const input = "\ufeff<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		`<dictionary><tag group="0010" element="0010" vr="PN" vm="1">Patient Name</tag></dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}
	if d.Lookup(tag.New(0x0010, 0x0010)) == nil {
		t.Fatal("Lookup(PatientName) = nil, want entry")
	}
}

func TestNewFromXMLIgnoresUIDElementsFromCombinedFODICOMSource(t *testing.T) {
	const input = `<dictionary version="2026b">
  <tag group="0010" element="0010" keyword="PatientName" vr="PN" vm="1">Patient Name</tag>
  <uid uid="1.2.840.10008.1.1" keyword="Verification" type="SOP Class">Verification SOP Class</uid>
</dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}
	if len(d.Entries()) != 1 {
		t.Fatalf("Entries() length = %d, want only the tag entry", len(d.Entries()))
	}
}

func TestNewFromXMLNormalizesAlternativeVM(t *testing.T) {
	const input = `<dictionary>
  <tag group="0028" element="3006" keyword="LUTData" vr="US/OW" vm="1-n or 1">LUT Data</tag>
</dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}

	entry := d.Lookup(tag.New(0x0028, 0x3006))
	if entry == nil {
		t.Fatal("Lookup(LUTData) = nil, want entry")
	}
	if entry.VM() != "1-n" {
		t.Errorf("LUTData.VM() = %q, want 1-n", entry.VM())
	}
}

func TestNewFromXMLAcceptsFODICOMVRSeparators(t *testing.T) {
	const input = `<dictionary>
  <tag group="0028" element="0106" vr="US_SS\OW,OB|UN" vm="1">Pixel Value</tag>
</dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}
	entry := d.Lookup(tag.New(0x0028, 0x0106))
	if entry == nil {
		t.Fatal("Lookup(PixelValue) = nil, want entry")
	}
	want := []string{"US", "SS", "OW", "OB", "UN"}
	got := entry.VRs()
	if len(got) != len(want) {
		t.Fatalf("VRs() = %v, want %v", got, want)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("VRs() = %v, want %v", got, want)
		}
	}
}

func TestNewFromXMLLoadsPrivateDictionariesByCreator(t *testing.T) {
	const (
		creatorAName = "ACME A"
		input        = `<dictionaries>
  <dictionary creator="ACME A">
    <tag group="0011" element="xx10" keyword="AcmeValue" vr="LO" vm="1">ACME A Value</tag>
    <tag group="0011" element="xx11" vr="UN" vm="1">Unnamed Keyword Value</tag>
  </dictionary>
  <dictionary creator="ACME B">
    <tag group="0011" element="xx10" keyword="OtherValue" vr="US" vm="1">ACME B Value</tag>
  </dictionary>
</dictionaries>`
	)

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}

	creatorA := d.GetPrivateCreator(creatorAName)
	creatorB := d.GetPrivateCreator("ACME B")
	entryA := d.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1010, creatorA))
	entryB := d.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1010, creatorB))
	if entryA == nil || entryA.Name() != "ACME A Value" {
		t.Fatalf("Lookup(ACME A tag) = %#v, want ACME A Value", entryA)
	}
	if entryB == nil || entryB.Name() != "ACME B Value" {
		t.Fatalf("Lookup(ACME B tag) = %#v, want ACME B Value", entryB)
	}
	if got := entryA.Tag().PrivateCreator(); got == nil || got.Creator() != creatorAName {
		t.Errorf("ACME A entry private creator = %v, want ACME A", got)
	}

	keywordTag := d.LookupKeyword("AcmeValue")
	if keywordTag == nil {
		t.Fatal("LookupKeyword(AcmeValue) = nil, want private tag")
	}
	if got := keywordTag.PrivateCreator(); got == nil || got.Creator() != creatorAName {
		t.Errorf("LookupKeyword(AcmeValue).PrivateCreator() = %v, want ACME A", got)
	}
	if d.LookupKeyword("Unnamed Keyword Value") != nil {
		t.Error("missing private keyword was derived from the name, want no keyword mapping")
	}

	privateA := d.GetPrivateDictionary(creatorAName)
	if privateA.PrivateCreator() != creatorA {
		t.Error("GetPrivateDictionary() does not share the dictionary's cached creator")
	}
	if privateA != d.GetPrivateDictionary(creatorAName) {
		t.Error("GetPrivateDictionary() returned different dictionaries for the same creator")
	}
	if got := privateA.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1010, creatorB)); got != nil {
		t.Errorf("private dictionary lookup with another creator = %#v, want nil", got)
	}
}

func TestPrivateExactEntryIgnoresAllocatedCreatorBlock(t *testing.T) {
	const input = `<dictionary creator="ACME">
  <tag group="0019" element="1000" keyword="FieldOfView" vr="DS" vm="1">Field Of View</tag>
</dictionary>`

	d, err := dict.NewFromXML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}
	creator := d.GetPrivateCreator("ACME")
	entry := d.Lookup(tag.NewWithPrivateCreator(0x0019, 0x1100, creator))
	if entry == nil || entry.Name() != "Field Of View" {
		t.Fatalf("lookup in allocated block 0x11 = %#v, want Field Of View", entry)
	}
	if d.Lookup(tag.NewWithPrivateCreator(0x0019, 0x1100, d.GetPrivateCreator("OTHER"))) != nil {
		t.Fatal("private exact entry resolved for a different creator")
	}
}

func TestPrivateDictionaryAddDoesNotMutateCallerEntry(t *testing.T) {
	entry := dict.NewEntry(
		tag.New(0x0019, 0x1000),
		"Field Of View",
		"FieldOfView",
		vm.VM1,
		false,
		vr.DS,
	)
	creatorA := tag.NewPrivateCreator("ACME A")
	creatorB := tag.NewPrivateCreator("ACME B")
	dictionaryA := dict.NewPrivate(creatorA)
	dictionaryB := dict.NewPrivate(creatorB)

	dictionaryA.Add(entry)
	dictionaryB.Add(entry)

	if entry.Tag().PrivateCreator() != nil {
		t.Fatalf("Add() changed caller entry creator to %v", entry.Tag().PrivateCreator())
	}
	entryA := dictionaryA.Lookup(tag.NewWithPrivateCreator(0x0019, 0x1100, creatorA))
	entryB := dictionaryB.Lookup(tag.NewWithPrivateCreator(0x0019, 0x1200, creatorB))
	if entryA == nil || entryA.Tag().PrivateCreator().Creator() != "ACME A" {
		t.Fatalf("first private dictionary entry = %#v, want creator ACME A", entryA)
	}
	if entryB == nil || entryB.Tag().PrivateCreator().Creator() != "ACME B" {
		t.Fatalf("second private dictionary entry = %#v, want creator ACME B", entryB)
	}
}

func TestDictionaryAddRoutesProgrammaticPrivateEntry(t *testing.T) {
	d := dict.New()
	creator := d.GetPrivateCreator("ACME")
	d.Add(dict.NewEntry(
		tag.NewWithPrivateCreator(0x0019, 0x1000, creator),
		"Field Of View",
		"FieldOfView",
		vm.VM1,
		false,
		vr.DS,
	))

	entry := d.Lookup(tag.NewWithPrivateCreator(0x0019, 0x1100, creator))
	if entry == nil || entry.Name() != "Field Of View" {
		t.Fatalf("programmatic private lookup = %#v, want Field Of View", entry)
	}
}

func TestLoadXMLUsesExactThenLatestMaskedEntry(t *testing.T) {
	const input = `<dictionary>
  <tag group="60xx" element="3000" keyword="FirstOverlayData" vr="OW" vm="1">First Overlay Data</tag>
  <tag group="60xx" element="3000" keyword="LatestOverlayData" vr="OB/OW" vm="1">Latest Overlay Data</tag>
  <tag group="6002" element="3000" keyword="ExactOverlayData" vr="OB" vm="1">Exact Overlay Data</tag>
</dictionary>`

	d := dict.New()
	if err := d.LoadXML(strings.NewReader(input)); err != nil {
		t.Fatalf("LoadXML() error = %v", err)
	}

	exact := d.Lookup(tag.New(0x6002, 0x3000))
	if exact == nil || exact.Name() != "Exact Overlay Data" {
		t.Fatalf("exact lookup = %#v, want Exact Overlay Data", exact)
	}
	masked := d.Lookup(tag.New(0x6004, 0x3000))
	if masked == nil || masked.Name() != "Latest Overlay Data" {
		t.Fatalf("masked lookup = %#v, want Latest Overlay Data", masked)
	}
}

func TestLoadXMLReplacesDuplicateExactEntry(t *testing.T) {
	d := dict.New()
	d.Add(dict.NewEntry(
		tag.New(0x0010, 0x0010),
		"Old Patient Name",
		"OldPatientName",
		vm.VM1,
		false,
		vr.PN,
	))

	const input = `<dictionary>
  <tag group="0010" element="0010" keyword="PatientName" vr="PN" vm="1">Patient's Name</tag>
</dictionary>`
	if err := d.LoadXML(strings.NewReader(input)); err != nil {
		t.Fatalf("LoadXML() error = %v", err)
	}

	entry := d.Lookup(tag.New(0x0010, 0x0010))
	if entry == nil || entry.Name() != "Patient's Name" {
		t.Fatalf("duplicate lookup = %#v, want replacement entry", entry)
	}
	if d.LookupKeyword("OldPatientName") != nil {
		t.Error("LookupKeyword(OldPatientName) still resolves after its entry was replaced")
	}
}

func TestLoadXMLRejectsInvalidInputWithEntryContext(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantDetail string
	}{
		{
			name:       "unexpected root",
			input:      `<dicomDictionary/>`,
			wantDetail: `expected <dictionary> or <dictionaries> root`,
		},
		{
			name:       "missing group",
			input:      `<dictionary><tag element="0010" vr="PN" vm="1">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1: missing group`,
		},
		{
			name:       "invalid masked tag",
			input:      `<dictionary><tag group="001x" element="001Z" vr="PN" vm="1">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (001x,001Z): invalid element`,
		},
		{
			name:       "missing VM",
			input:      `<dictionary><tag group="0010" element="0010" vr="PN">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): missing VM`,
		},
		{
			name:       "unknown VR",
			input:      `<dictionary><tag group="0010" element="0010" vr="ZZ" vm="1">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): invalid VR "ZZ"`,
		},
		{
			name:       "empty composite VR component",
			input:      `<dictionary><tag group="0010" element="0010" vr="US//SS" vm="1">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): invalid VR list "US//SS"`,
		},
		{
			name:       "only VR separator",
			input:      `<dictionary><tag group="0010" element="0010" vr="/" vm="1">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): invalid VR list "/"`,
		},
		{
			name:       "invalid VM",
			input:      `<dictionary><tag group="0010" element="0010" vr="PN" vm="many">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): invalid VM "many"`,
		},
		{
			name:       "invalid retired",
			input:      `<dictionary><tag group="0010" element="0010" vr="PN" vm="1" retired="yes">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): invalid retired value "yes"`,
		},
		{
			name:       "missing name",
			input:      `<dictionary><tag group="0010" element="0010" vr="PN" vm="1"></tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): missing name`,
		},
		{
			name:       "unknown tag attribute",
			input:      `<dictionary><tag group="0010" element="0010" vr="PN" vm="1" custom="x">Patient Name</tag></dictionary>`,
			wantDetail: `tag 1 (0010,0010): unexpected attribute "custom"`,
		},
		{
			name:       "unexpected child",
			input:      `<dictionary><entry group="0010" element="0010" vr="PN" vm="1">Patient Name</entry></dictionary>`,
			wantDetail: `unexpected <entry> element`,
		},
		{
			name:       "malformed XML",
			input:      `<dictionary><tag></dictionary>`,
			wantDetail: `XML syntax error`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dict.New()
			err := d.LoadXML(strings.NewReader(tt.input))
			if err == nil {
				t.Fatal("LoadXML() error = nil, want validation error")
			}
			if !strings.Contains(err.Error(), tt.wantDetail) {
				t.Fatalf("LoadXML() error = %q, want detail %q", err, tt.wantDetail)
			}
		})
	}
}

func TestLoadXMLDoesNotPartiallyMergeInvalidDocument(t *testing.T) {
	d := dict.New()
	d.Add(dict.NewEntry(
		tag.New(0x0010, 0x0010),
		"Original Patient Name",
		"PatientName",
		vm.VM1,
		false,
		vr.PN,
	))

	const input = `<dictionary>
  <tag group="0010" element="0010" keyword="PatientName" vr="LO" vm="1">Replacement Patient Name</tag>
  <tag group="0010" element="0020" keyword="PatientID" vr="ZZ" vm="1">Patient ID</tag>
</dictionary>`
	err := d.LoadXML(strings.NewReader(input))
	if err == nil {
		t.Fatal("LoadXML() error = nil, want validation error")
	}
	if !strings.Contains(err.Error(), `tag 2 (0010,0020)`) {
		t.Fatalf("LoadXML() error = %q, want second tag context", err)
	}

	entry := d.Lookup(tag.New(0x0010, 0x0010))
	if entry == nil || entry.Name() != "Original Patient Name" {
		t.Fatalf("dictionary changed after failed load: entry = %#v", entry)
	}
	if d.Lookup(tag.New(0x0010, 0x0020)) != nil {
		t.Error("invalid document partially added PatientID")
	}
}

func TestLoadXMLReportsContextForMalformedTagXML(t *testing.T) {
	const input = `<dictionaries>
  <dictionary creator="ACME">
    <tag group="0011" element="xx10" vr="LO" vm="1">Broken</dictionary>
  </dictionaries>`

	err := dict.New().LoadXML(strings.NewReader(input))
	if err == nil {
		t.Fatal("LoadXML() error = nil, want XML syntax error")
	}
	want := `dictionary "ACME" tag 1 (0011,xx10): XML syntax error`
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("LoadXML() error = %q, want context %q", err, want)
	}
}

func TestGetPrivateDictionaryThreadSafety(t *testing.T) {
	const goroutines = 32
	d := dict.New()
	start := make(chan struct{})
	results := make(chan *dict.Dictionary, goroutines)

	var workers sync.WaitGroup
	workers.Add(goroutines)
	for range goroutines {
		go func() {
			defer workers.Done()
			<-start
			results <- d.GetPrivateDictionary("ACME")
		}()
	}
	close(start)
	workers.Wait()
	close(results)

	var first *dict.Dictionary
	for privateDictionary := range results {
		if first == nil {
			first = privateDictionary
			continue
		}
		if privateDictionary != first {
			t.Fatal("GetPrivateDictionary() returned different instances concurrently")
		}
	}
}
