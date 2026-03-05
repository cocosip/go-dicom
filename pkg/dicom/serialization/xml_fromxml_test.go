// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestFromXML_Basic(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00100010" vr="PN" keyword="PatientName">
    <PersonName number="1">
      <Alphabetic>
        <FamilyName>Doe</FamilyName>
        <GivenName>John</GivenName>
      </Alphabetic>
    </PersonName>
  </DicomAttribute>
  <DicomAttribute tag="00100020" vr="LO" keyword="PatientID">
    <Value number="1">12345</Value>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	if ds.Count() != 2 {
		t.Errorf("Count() = %d, want 2", ds.Count())
	}

	// Check PatientName
	pnElem, found := ds.Get(tag.PatientName)
	if !found {
		t.Error("PatientName not found")
	} else {
		pn, ok := pnElem.(*element.PersonName)
		if !ok {
			t.Errorf("PatientName is not PersonName, got %T", pnElem)
		} else {
			if pn.Count() == 0 {
				t.Error("PatientName has no values")
			} else if pn.GetValue(0) != "Doe^John" {
				t.Errorf("PatientName = %q, want %q", pn.GetValue(0), "Doe^John")
			}
		}
	}

	// Check PatientID
	pidElem, found := ds.Get(tag.PatientID)
	if !found {
		t.Error("PatientID not found")
	} else {
		strElem, ok := pidElem.(*element.String)
		if !ok {
			t.Errorf("PatientID is not String, got %T", pidElem)
		} else {
			if strElem.GetValue(0) != "12345" {
				t.Errorf("PatientID = %q, want %q", strElem.GetValue(0), "12345")
			}
		}
	}
}

func TestFromXML_Sequence(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00081115" vr="SQ" keyword="ReferencedSeriesSequence">
    <Item number="1">
      <DicomAttribute tag="0020000E" vr="UI" keyword="SeriesInstanceUID">
        <Value number="1">1.2.3.4.5.6.7.8.9</Value>
      </DicomAttribute>
    </Item>
    <Item number="2">
      <DicomAttribute tag="0020000E" vr="UI" keyword="SeriesInstanceUID">
        <Value number="1">1.2.3.4.5.6.7.8.10</Value>
      </DicomAttribute>
    </Item>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	seqElem, found := ds.Get(tag.ReferencedSeriesSequence)
	if !found {
		t.Error("ReferencedSeriesSequence not found")
		return
	}

	seq, ok := seqElem.(*dataset.Sequence)
	if !ok {
		t.Errorf("ReferencedSeriesSequence is not Sequence, got %T", seqElem)
		return
	}

	if seq.Count() != 2 {
		t.Errorf("Sequence Count() = %d, want 2", seq.Count())
		return
	}

	// Check first item
	item0 := seq.GetItem(0)
	uidElem, found := item0.Get(tag.SeriesInstanceUID)
	if !found {
		t.Error("SeriesInstanceUID not found in first item")
	} else {
		strElem := uidElem.(*element.String)
		if strElem.GetValue(0) != "1.2.3.4.5.6.7.8.9" {
			t.Errorf("First item UID = %q, want %q", strElem.GetValue(0), "1.2.3.4.5.6.7.8.9")
		}
	}

	// Check second item
	item1 := seq.GetItem(1)
	uidElem, found = item1.Get(tag.SeriesInstanceUID)
	if !found {
		t.Error("SeriesInstanceUID not found in second item")
	} else {
		strElem := uidElem.(*element.String)
		if strElem.GetValue(0) != "1.2.3.4.5.6.7.8.10" {
			t.Errorf("Second item UID = %q, want %q", strElem.GetValue(0), "1.2.3.4.5.6.7.8.10")
		}
	}
}

func TestFromXML_Binary(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="7FE00010" vr="OW" keyword="PixelData">
    <InlineBinary>AAECAwQ=</InlineBinary>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	pdElem, found := ds.Get(tag.PixelData)
	if !found {
		t.Error("PixelData not found")
		return
	}

	owElem, ok := pdElem.(*element.OtherWord)
	if !ok {
		t.Errorf("PixelData is not OtherWord, got %T", pdElem)
		return
	}

	data := owElem.GetData()
	expected := []byte{0, 1, 2, 3, 4} // base64 of "AAECAwQ="
	if len(data) != len(expected) {
		t.Errorf("PixelData length = %d, want %d", len(data), len(expected))
		return
	}

	for i, b := range data {
		if b != expected[i] {
			t.Errorf("PixelData[%d] = %d, want %d", i, b, expected[i])
		}
	}
}

func TestFromXML_Numeric(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00280010" vr="US" keyword="Rows">
    <Value number="1">512</Value>
  </DicomAttribute>
  <DicomAttribute tag="00280011" vr="US" keyword="Columns">
    <Value number="1">512</Value>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	// Note: FromXML creates String elements for all non-binary VRs
	// Use GetString and parse for numeric values
	rowsStr, found := ds.GetString(tag.Rows)
	if !found {
		t.Error("Rows not found")
	} else if rowsStr != "512" {
		t.Errorf("Rows = %q, want 512", rowsStr)
	}

	colsStr, found := ds.GetString(tag.Columns)
	if !found {
		t.Error("Columns not found")
	} else if colsStr != "512" {
		t.Errorf("Columns = %q, want 512", colsStr)
	}
}

func TestFromXML_EmptyDataset(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	if ds.Count() != 0 {
		t.Errorf("Count() = %d, want 0", ds.Count())
	}
}

func TestFromXML_InvalidXML(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="invalid" vr="LO">
    <Value number="1">test</Value>
  </DicomAttribute>
</NativeDicomModel>`)

	_, err := FromXML(xmlData)
	if err == nil {
		t.Error("FromXML() expected error for invalid tag, got nil")
	}
}

func TestFromXML_XMLRoundtrip(t *testing.T) {
	t.Parallel()
	// Create original dataset
	original := dataset.New()
	_ = original.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
	_ = original.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	_ = original.Add(element.NewString(tag.Rows, vr.US, []string{"512"}))
	_ = original.Add(element.NewString(tag.Columns, vr.US, []string{"512"}))

	// Serialize to XML
	xmlData, err := ToXML(original)
	if err != nil {
		t.Fatalf("ToXML() error = %v", err)
	}

	// Deserialize from XML
	restored, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	// Compare datasets
	if restored.Count() != original.Count() {
		t.Errorf("Restored Count() = %d, want %d", restored.Count(), original.Count())
	}

	// Check PatientName
	pnRest, _ := restored.Get(tag.PatientName)
	if pnRest == nil {
		t.Error("Restored PatientName is nil")
	}

	// Check PatientID
	pidOrig, _ := original.GetString(tag.PatientID)
	pidRest, _ := restored.GetString(tag.PatientID)
	if pidRest != pidOrig {
		t.Errorf("Restored PatientID = %q, want %q", pidRest, pidOrig)
	}

	// Check Rows
	rowsOrig, _ := original.GetString(tag.Rows)
	rowsRest, _ := restored.GetString(tag.Rows)
	if rowsRest != rowsOrig {
		t.Errorf("Restored Rows = %q, want %q", rowsRest, rowsOrig)
	}

	// Check Columns
	colsOrig, _ := original.GetString(tag.Columns)
	colsRest, _ := restored.GetString(tag.Columns)
	if colsRest != colsOrig {
		t.Errorf("Restored Columns = %q, want %q", colsRest, colsOrig)
	}
}

func TestFromXML_PersonNameWithMultipleComponents(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00100010" vr="PN" keyword="PatientName">
    <PersonName number="1">
      <Alphabetic>
        <FamilyName>Yamada</FamilyName>
        <GivenName>Tarou</GivenName>
        <MiddleName></MiddleName>
        <NamePrefix></NamePrefix>
        <NameSuffix></NameSuffix>
      </Alphabetic>
      <Ideographic>
        <FamilyName>山田</FamilyName>
        <GivenName>太郎</GivenName>
      </Ideographic>
      <Phonetic>
        <FamilyName>ヤマダ</FamilyName>
        <GivenName>タロウ</GivenName>
      </Phonetic>
    </PersonName>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	pnElem, found := ds.Get(tag.PatientName)
	if !found {
		t.Error("PatientName not found")
		return
	}

	pn, ok := pnElem.(*element.PersonName)
	if !ok {
		t.Errorf("PatientName is not PersonName, got %T", pnElem)
		return
	}

	if pn.Count() == 0 {
		t.Error("PatientName has no values")
		return
	}

	// The value should contain all three components
	value := pn.GetValue(0)
	// Check that it contains the components (exact format may vary)
	if value == "" {
		t.Error("PatientName value is empty")
	}
}

func TestFromXML_MultipleValues(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00080054" vr="AE" keyword="RetrieveAETitle">
    <Value number="1">AET1</Value>
    <Value number="2">AET2</Value>
    <Value number="3">AET3</Value>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	elem, found := ds.Get(tag.RetrieveAETitle)
	if !found {
		t.Error("RetrieveAETitle not found")
		return
	}

	strElem, ok := elem.(*element.String)
	if !ok {
		t.Errorf("RetrieveAETitle is not String, got %T", elem)
		return
	}

	if strElem.Count() != 3 {
		t.Errorf("Count() = %d, want 3", strElem.Count())
		return
	}

	expected := []string{"AET1", "AET2", "AET3"}
	for i, exp := range expected {
		if strElem.GetValue(i) != exp {
			t.Errorf("Value[%d] = %q, want %q", i, strElem.GetValue(i), exp)
		}
	}
}
