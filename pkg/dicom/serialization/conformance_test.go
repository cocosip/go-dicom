// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestFromJSONDeclaresUTF8ForUnicodeText(t *testing.T) {
	data := []byte(`{
		"00080005":{"vr":"CS","Value":["ISO_IR 100"]},
		"00100010":{"vr":"PN","Value":[{"Alphabetic":"Müller^Jörg"}]}
	}`)
	ds, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}
	charset := ds.SpecificCharacterSet()
	if len(charset) != 1 || charset[0] != "ISO_IR 192" {
		t.Fatalf("SpecificCharacterSet = %v, want [ISO_IR 192]", charset)
	}
	value, ok := ds.GetString(tag.PatientName)
	if !ok || value != "Müller^Jörg" {
		t.Fatalf("PatientName = %q, %v, want Müller^Jörg", value, ok)
	}

	roundTrip, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}
	restored, err := FromJSON(roundTrip)
	if err != nil {
		t.Fatalf("FromJSON(round trip) error = %v", err)
	}
	if value, _ := restored.GetString(tag.PatientName); value != "Müller^Jörg" {
		t.Fatalf("round-tripped PatientName = %q, want Müller^Jörg", value)
	}
}

func TestFromXMLDeclaresUTF8ForUnicodeText(t *testing.T) {
	data := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel xmlns="http://dicom.nema.org/PS3.19/models/NativeDICOM" xml:space="preserve">
  <DicomAttribute tag="00100010" vr="PN">
    <PersonName number="1">
      <Alphabetic><FamilyName>Yamada</FamilyName><GivenName>Tarou</GivenName></Alphabetic>
      <Ideographic><FamilyName>山田</FamilyName><GivenName>太郎</GivenName></Ideographic>
    </PersonName>
  </DicomAttribute>
</NativeDicomModel>`)
	ds, err := FromXML(data)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	charset := ds.SpecificCharacterSet()
	if len(charset) != 1 || charset[0] != "ISO_IR 192" {
		t.Fatalf("SpecificCharacterSet = %v, want [ISO_IR 192]", charset)
	}
	if value, _ := ds.GetString(tag.PatientName); !strings.Contains(value, "山田") {
		t.Fatalf("PatientName = %q, want ideographic component preserved", value)
	}
}

func TestFromJSONPreservesDSNumberText(t *testing.T) {
	data := []byte(`{
		"00101030":{"vr":"DS","Value":[9999999999999999]},
		"00101020":{"vr":"DS","Value":[1E+20]}
	}`)
	ds, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}
	weight, ok := ds.GetString(tag.PatientWeight)
	if !ok || weight != "9999999999999999" {
		t.Fatalf("PatientWeight = %q, %v, want 9999999999999999", weight, ok)
	}
	size, ok := ds.GetString(tag.PatientSize)
	if !ok || size != "1E+20" {
		t.Fatalf("PatientSize = %q, %v, want 1E+20", size, ok)
	}

	encoded, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}
	if !strings.Contains(string(encoded), "9999999999999999") {
		t.Fatalf("ToJSON() lost the exact DS text: %s", encoded)
	}
}

func TestFromXMLInfersOptionalVR(t *testing.T) {
	data := []byte(`<NativeDicomModel>
  <DicomAttribute tag="00100010">
    <PersonName number="1"><Alphabetic><FamilyName>Doe</FamilyName></Alphabetic></PersonName>
  </DicomAttribute>
</NativeDicomModel>`)
	ds, err := FromXML(data)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	if _, ok := ds.Get(tag.PatientName); !ok {
		t.Fatal("PatientName element not found")
	}
}

func TestXMLBulkDataStringVR(t *testing.T) {
	uri := "http://example.com/report/text"
	elem, err := element.NewElementFromBuffer(tag.TextValue, vr.UT, buffer.NewBulkDataURI(uri), nil)
	if err != nil {
		t.Fatalf("NewElementFromBuffer() error = %v", err)
	}
	ds := dataset.New()
	if err := ds.Add(elem); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	encoded, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML() error = %v", err)
	}
	if !strings.Contains(string(encoded), `<BulkData uri="`+uri+`"/>`) {
		t.Fatalf("ToXML() did not emit BulkData for UT: %s", encoded)
	}

	xmlData := []byte(`<NativeDicomModel>
  <DicomAttribute tag="0040A160" vr="UT"><BulkData uri="` + uri + `"/></DicomAttribute>
</NativeDicomModel>`)
	restored, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	got, ok := restored.Get(tag.TextValue)
	if !ok {
		t.Fatal("TextValue element not found")
	}
	bulk, ok := got.Buffer().(*buffer.BulkDataURIByteBuffer)
	if !ok || bulk.BulkDataURI() != uri {
		t.Fatalf("TextValue buffer = %T, want BulkDataURI %q", got.Buffer(), uri)
	}
}

func TestXMLBulkDataUUIDReportsMissingDataStore(t *testing.T) {
	xmlData := []byte(`<NativeDicomModel>
  <DicomAttribute tag="0040A160" vr="UT"><BulkData uuid="d3b2a1e0-0000-0000-0000-000000000000"/></DicomAttribute>
</NativeDicomModel>`)
	_, err := FromXML(xmlData)
	if err == nil || !strings.Contains(err.Error(), "data store") {
		t.Fatalf("FromXML() error = %v, want missing data store error", err)
	}
}

func TestXMLPrivateCreatorWithoutCreatorElement(t *testing.T) {
	xmlData := []byte(`<NativeDicomModel>
  <DicomAttribute tag="00090010" vr="LO" privateCreator="ACME"><Value number="1">value</Value></DicomAttribute>
</NativeDicomModel>`)
	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	privateTag := tag.NewWithPrivateCreator(0x0009, 0x1010, tag.NewPrivateCreator("ACME"))
	if _, ok := ds.Get(privateTag); !ok {
		t.Fatalf("private element %s not found", privateTag)
	}
	creatorTag := tag.New(0x0009, 0x0010)
	if value, ok := ds.GetString(creatorTag); !ok || value != "ACME" {
		t.Fatalf("Private Creator = %q, %v, want ACME", value, ok)
	}

	if _, err := ToXML(ds); err != nil {
		t.Fatalf("ToXML() after allocating the block error = %v", err)
	}
}

func TestToXMLRejectsPrivateElementWithoutCreator(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.New(0x0009, 0x1010), vr.LO, []string{"value"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	_, err := ToXML(ds)
	if err == nil || !strings.Contains(err.Error(), "Private Creator") {
		t.Fatalf("ToXML() error = %v, want Private Creator error", err)
	}
}

func TestPersonNameConcreteTypeIsConsistent(t *testing.T) {
	jsonData := []byte(`{"00100010":{"vr":"PN","Value":[{"Alphabetic":"Doe^Jane"}]}}`)
	fromJSON, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}
	jsonElem, _ := fromJSON.Get(tag.PatientName)
	if _, ok := jsonElem.(*element.PersonName); !ok {
		t.Fatalf("JSON PatientName type = %T, want *element.PersonName", jsonElem)
	}

	xmlData := []byte(`<NativeDicomModel>
  <DicomAttribute tag="00100010" vr="PN"><PersonName number="1"><Alphabetic><FamilyName>Doe</FamilyName></Alphabetic></PersonName></DicomAttribute>
</NativeDicomModel>`)
	fromXML, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	xmlElem, _ := fromXML.Get(tag.PatientName)
	if reflect.TypeOf(jsonElem) != reflect.TypeOf(xmlElem) {
		t.Fatalf("PN types differ: JSON %T, XML %T", jsonElem, xmlElem)
	}
}
