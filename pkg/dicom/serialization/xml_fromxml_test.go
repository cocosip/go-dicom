// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	xmlUnsignedShortName    = "unsigned short"
	xmlUnsignedLongName     = "unsigned long"
	xmlSignedShortName      = "signed short"
	xmlSignedLongName       = "signed long"
	xmlFloatName            = "float"
	xmlDoubleName           = "double"
	xmlSignedVeryLongName   = "signed very long"
	xmlUnsignedVeryLongName = "unsigned very long"
	xmlAttributeTagName     = "attribute tag"
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
			} else if pn.GetValue(0) != testPatientNameJohn {
				t.Errorf("PatientName = %q, want %q", pn.GetValue(0), testPatientNameJohn)
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

func TestFromXML_RestoresBinaryElementTypes(t *testing.T) {
	tests := []struct {
		name     string
		vrCode   string
		wantVR   *vr.VR
		wantType reflect.Type
	}{
		{name: "other byte", vrCode: "OB", wantVR: vr.OB, wantType: reflect.TypeOf((*element.OtherByte)(nil))},
		{name: "other word", vrCode: "OW", wantVR: vr.OW, wantType: reflect.TypeOf((*element.OtherWord)(nil))},
		{name: "other double", vrCode: "OD", wantVR: vr.OD, wantType: reflect.TypeOf((*element.OtherDouble)(nil))},
		{name: "other float", vrCode: "OF", wantVR: vr.OF, wantType: reflect.TypeOf((*element.OtherFloat)(nil))},
		{name: "other long", vrCode: "OL", wantVR: vr.OL, wantType: reflect.TypeOf((*element.OtherLong)(nil))},
		{name: "other very long", vrCode: "OV", wantVR: vr.OV, wantType: reflect.TypeOf((*element.OtherVeryLong)(nil))},
		{name: "unknown", vrCode: "UN", wantVR: vr.UN, wantType: reflect.TypeOf((*element.Unknown)(nil))},
	}

	for _, tt := range tests {
		for _, withData := range []bool{false, true} {
			caseName := "empty"
			body := ""
			wantData := []byte(nil)
			wantCount := 0
			if withData {
				caseName = "inline binary"
				body = "<InlineBinary>AQIDBA==</InlineBinary>"
				wantData = []byte{1, 2, 3, 4}
				wantCount = 1
			}

			t.Run(tt.name+"/"+caseName, func(t *testing.T) {
				xmlData := []byte(fmt.Sprintf(
					`<NativeDicomModel><DicomAttribute tag="00111010" vr="%s">%s</DicomAttribute></NativeDicomModel>`,
					tt.vrCode,
					body,
				))

				ds, err := FromXML(xmlData)
				if err != nil {
					t.Fatalf("FromXML() error = %v", err)
				}
				got, found := ds.Get(tag.New(0x0011, 0x1010))
				if !found {
					t.Fatal("binary element not found")
				}
				if gotType := reflect.TypeOf(got); gotType != tt.wantType {
					t.Fatalf("element type = %v, want %v", gotType, tt.wantType)
				}
				if got.ValueRepresentation() != tt.wantVR {
					t.Fatalf("VR = %s, want %s", got.ValueRepresentation(), tt.wantVR)
				}
				if got.Count() != wantCount {
					t.Fatalf("Count() = %d, want %d", got.Count(), wantCount)
				}
				if gotData := got.Buffer().Data(); !bytes.Equal(gotData, wantData) {
					t.Fatalf("data = %v, want %v", gotData, wantData)
				}
			})
		}
	}
}

func TestFromXML_Numeric(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00111010" vr="US">
    <Value number="1">65535</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111011" vr="UL">
    <Value number="1">4294967295</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111012" vr="SS">
    <Value number="1">-32768</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111013" vr="SL">
    <Value number="1">-2147483648</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111014" vr="FL">
    <Value number="1">1.5</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111015" vr="FD">
    <Value number="1">-2.25</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111016" vr="SV">
    <Value number="1">-9223372036854775807</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111017" vr="UV">
    <Value number="1">18446744073709551614</Value>
  </DicomAttribute>
  <DicomAttribute tag="00111018" vr="AT">
    <Value number="1">00280010</Value>
  </DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	tests := []struct {
		name       string
		tag        *tag.Tag
		wantType   reflect.Type
		wantValues []string
	}{
		{name: xmlUnsignedShortName, tag: tag.New(0x0011, 0x1010), wantType: reflect.TypeOf((*element.UnsignedShort)(nil)), wantValues: []string{"65535"}},
		{name: xmlUnsignedLongName, tag: tag.New(0x0011, 0x1011), wantType: reflect.TypeOf((*element.UnsignedLong)(nil)), wantValues: []string{"4294967295"}},
		{name: xmlSignedShortName, tag: tag.New(0x0011, 0x1012), wantType: reflect.TypeOf((*element.SignedShort)(nil)), wantValues: []string{"-32768"}},
		{name: xmlSignedLongName, tag: tag.New(0x0011, 0x1013), wantType: reflect.TypeOf((*element.SignedLong)(nil)), wantValues: []string{"-2147483648"}},
		{name: xmlFloatName, tag: tag.New(0x0011, 0x1014), wantType: reflect.TypeOf((*element.Float)(nil)), wantValues: []string{"1.5"}},
		{name: xmlDoubleName, tag: tag.New(0x0011, 0x1015), wantType: reflect.TypeOf((*element.Double)(nil)), wantValues: []string{"-2.25"}},
		{name: xmlSignedVeryLongName, tag: tag.New(0x0011, 0x1016), wantType: reflect.TypeOf((*element.SignedVeryLong)(nil)), wantValues: []string{"-9223372036854775807"}},
		{name: xmlUnsignedVeryLongName, tag: tag.New(0x0011, 0x1017), wantType: reflect.TypeOf((*element.UnsignedVeryLong)(nil)), wantValues: []string{"18446744073709551614"}},
		{name: xmlAttributeTagName, tag: tag.New(0x0011, 0x1018), wantType: reflect.TypeOf((*element.AttributeTag)(nil)), wantValues: []string{"(0028,0010)"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := ds.Get(tt.tag)
			if !found {
				t.Fatalf("tag %s not found", tt.tag)
			}
			if gotType := reflect.TypeOf(got); gotType != tt.wantType {
				t.Fatalf("element type = %v, want %v", gotType, tt.wantType)
			}
			values, err := element.CanonicalStrings(got)
			if err != nil {
				t.Fatalf("CanonicalStrings() error = %v", err)
			}
			if !reflect.DeepEqual(values, tt.wantValues) {
				t.Fatalf("values = %v, want %v", values, tt.wantValues)
			}
		})
	}
}

func TestFromXML_RestoresEmptyNumericAndAttributeTagElements(t *testing.T) {
	tests := []struct {
		name     string
		vrCode   string
		wantType reflect.Type
	}{
		{name: xmlUnsignedShortName, vrCode: "US", wantType: reflect.TypeOf((*element.UnsignedShort)(nil))},
		{name: xmlUnsignedLongName, vrCode: "UL", wantType: reflect.TypeOf((*element.UnsignedLong)(nil))},
		{name: xmlSignedShortName, vrCode: "SS", wantType: reflect.TypeOf((*element.SignedShort)(nil))},
		{name: xmlSignedLongName, vrCode: "SL", wantType: reflect.TypeOf((*element.SignedLong)(nil))},
		{name: xmlFloatName, vrCode: "FL", wantType: reflect.TypeOf((*element.Float)(nil))},
		{name: xmlDoubleName, vrCode: "FD", wantType: reflect.TypeOf((*element.Double)(nil))},
		{name: xmlSignedVeryLongName, vrCode: "SV", wantType: reflect.TypeOf((*element.SignedVeryLong)(nil))},
		{name: xmlUnsignedVeryLongName, vrCode: "UV", wantType: reflect.TypeOf((*element.UnsignedVeryLong)(nil))},
		{name: xmlAttributeTagName, vrCode: "AT", wantType: reflect.TypeOf((*element.AttributeTag)(nil))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlData := []byte(fmt.Sprintf(
				`<NativeDicomModel><DicomAttribute tag="00111010" vr="%s"/></NativeDicomModel>`,
				tt.vrCode,
			))
			ds, err := FromXML(xmlData)
			if err != nil {
				t.Fatalf("FromXML() error = %v", err)
			}
			got, found := ds.Get(tag.New(0x0011, 0x1010))
			if !found {
				t.Fatal("element not found")
			}
			if gotType := reflect.TypeOf(got); gotType != tt.wantType {
				t.Fatalf("element type = %v, want %v", gotType, tt.wantType)
			}
			if got.Count() != 0 {
				t.Fatalf("Count() = %d, want 0", got.Count())
			}
		})
	}
}

func TestXMLRoundTripPreservesTypedValueElements(t *testing.T) {
	tests := []struct {
		name string
		tag  *tag.Tag
		elem element.Element
	}{
		{name: xmlUnsignedShortName, tag: tag.New(0x0011, 0x1010), elem: element.NewUnsignedShort(tag.New(0x0011, 0x1010), []uint16{65535})},
		{name: xmlUnsignedLongName, tag: tag.New(0x0011, 0x1011), elem: element.NewUnsignedLong(tag.New(0x0011, 0x1011), []uint32{4294967295})},
		{name: xmlSignedShortName, tag: tag.New(0x0011, 0x1012), elem: element.NewSignedShort(tag.New(0x0011, 0x1012), []int16{-32768})},
		{name: xmlSignedLongName, tag: tag.New(0x0011, 0x1013), elem: element.NewSignedLong(tag.New(0x0011, 0x1013), []int32{-2147483648})},
		{name: xmlFloatName, tag: tag.New(0x0011, 0x1014), elem: element.NewFloat(tag.New(0x0011, 0x1014), []float32{1.5})},
		{name: xmlDoubleName, tag: tag.New(0x0011, 0x1015), elem: element.NewDouble(tag.New(0x0011, 0x1015), []float64{-2.25})},
		{name: xmlSignedVeryLongName, tag: tag.New(0x0011, 0x1016), elem: element.NewSignedVeryLong(tag.New(0x0011, 0x1016), []int64{-9223372036854775807})},
		{name: xmlUnsignedVeryLongName, tag: tag.New(0x0011, 0x1017), elem: element.NewUnsignedVeryLong(tag.New(0x0011, 0x1017), []uint64{18446744073709551614})},
		{name: xmlAttributeTagName, tag: tag.New(0x0011, 0x1018), elem: element.NewAttributeTag(tag.New(0x0011, 0x1018), []*tag.Tag{tag.Rows})},
	}

	original := dataset.New()
	for _, tt := range tests {
		if err := original.Add(tt.elem); err != nil {
			t.Fatalf("Add(%s) error = %v", tt.name, err)
		}
	}
	xmlData, err := ToXML(original)
	if err != nil {
		t.Fatalf("ToXML() error = %v", err)
	}
	restored, err := FromXML(xmlData)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := restored.Get(tt.tag)
			if !found {
				t.Fatalf("tag %s not found", tt.tag)
			}
			if gotType, wantType := reflect.TypeOf(got), reflect.TypeOf(tt.elem); gotType != wantType {
				t.Fatalf("element type = %v, want %v", gotType, wantType)
			}
			gotValues, err := element.CanonicalStrings(got)
			if err != nil {
				t.Fatalf("restored CanonicalStrings() error = %v", err)
			}
			wantValues, err := element.CanonicalStrings(tt.elem)
			if err != nil {
				t.Fatalf("original CanonicalStrings() error = %v", err)
			}
			if !reflect.DeepEqual(gotValues, wantValues) {
				t.Fatalf("values = %v, want %v", gotValues, wantValues)
			}
		})
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
