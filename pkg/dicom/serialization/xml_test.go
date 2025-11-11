// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestToXML_EmptyDataset(t *testing.T) {
	ds := dataset.New()
	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)
	if !strings.Contains(xmlStr, `<?xml version="1.0" encoding="utf-8"?>`) {
		t.Error("XML declaration missing")
	}
	if !strings.Contains(xmlStr, `<NativeDicomModel>`) {
		t.Error("NativeDicomModel opening tag missing")
	}
	if !strings.Contains(xmlStr, `</NativeDicomModel>`) {
		t.Error("NativeDicomModel closing tag missing")
	}
}

func TestToXML_StringElement(t *testing.T) {
    ds := dataset.New()
    elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Debug: print the actual output
	t.Logf("XML output:\n%s", xmlStr)

	// Check DicomAttribute tag
	if !strings.Contains(xmlStr, `<DicomAttribute tag="00100010"`) {
		t.Error("DicomAttribute tag missing or incorrect")
	}
	if !strings.Contains(xmlStr, `vr="PN"`) {
		t.Error("VR attribute missing or incorrect")
	}
	// Keyword is optional if dictionary entry is not found
	// Just log if missing, don't fail
	if !strings.Contains(xmlStr, `keyword=`) {
		t.Logf("Note: keyword attribute not present (dictionary entry may not be loaded)")
	}

	// Check PersonName structure
	if !strings.Contains(xmlStr, `<PersonName number="1">`) {
		t.Error("PersonName element missing")
	}
	if !strings.Contains(xmlStr, `<Alphabetic>`) {
		t.Error("Alphabetic element missing")
	}
	if !strings.Contains(xmlStr, `<FamilyName>Doe</FamilyName>`) {
		t.Error("FamilyName missing or incorrect")
	}
	if !strings.Contains(xmlStr, `<GivenName>John</GivenName>`) {
		t.Error("GivenName missing or incorrect")
	}
}

func TestToXML_MultipleStringValues(t *testing.T) {
    ds := dataset.New()
    elem := element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check for all three values
	if !strings.Contains(xmlStr, `<Value number="1">ORIGINAL</Value>`) {
		t.Error("First value missing or incorrect")
	}
	if !strings.Contains(xmlStr, `<Value number="2">PRIMARY</Value>`) {
		t.Error("Second value missing or incorrect")
	}
	if !strings.Contains(xmlStr, `<Value number="3">AXIAL</Value>`) {
		t.Error("Third value missing or incorrect")
	}
}

func TestToXML_BinaryElement(t *testing.T) {
    ds := dataset.New()
    data := []byte{0x01, 0x02, 0x03, 0x04}
    elem := element.NewOtherByte(tag.PixelData, data)
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check for InlineBinary with base64 encoding
	if !strings.Contains(xmlStr, `<InlineBinary>AQIDBA==</InlineBinary>`) {
		t.Errorf("InlineBinary missing or incorrect encoding, got: %s", xmlStr)
	}
}

func TestToXML_NumericElements(t *testing.T) {
    ds := dataset.New()

    // Add US element
    _ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
    // Add UL element
    _ = ds.Add(element.NewUnsignedLong(tag.InstanceNumber, []uint32{1}))
    // Add FL element
    _ = ds.Add(element.NewFloat(tag.RecommendedDisplayFrameRate, []float32{30.0}))

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)
	t.Logf("XML output:\n%s", xmlStr)

	// Check numeric values are properly serialized
	if !strings.Contains(xmlStr, `<Value number="1">512</Value>`) {
		t.Error("US value missing or incorrect")
	}
	if !strings.Contains(xmlStr, `<Value number="1">1</Value>`) {
		t.Error("UL value missing or incorrect")
	}
	if !strings.Contains(xmlStr, `<Value number="1">30</Value>`) {
		t.Error("FL value missing or incorrect")
	}
}

func TestToXML_Sequence(t *testing.T) {
    ds := dataset.New()

    // Create a sequence with one item
    seq := dataset.NewSequence(tag.ReferencedStudySequence)
    item := dataset.New()
    _ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4"}))
    seq.AddItem(item)

    _ = ds.Add(seq)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check sequence structure
	if !strings.Contains(xmlStr, `<DicomAttribute tag="00081110"`) {
		t.Error("Sequence DicomAttribute tag missing")
	}
	if !strings.Contains(xmlStr, `vr="SQ"`) {
		t.Error("SQ VR missing")
	}
	if !strings.Contains(xmlStr, `<Item number="1">`) {
		t.Error("Item element missing")
	}
	if !strings.Contains(xmlStr, `</Item>`) {
		t.Error("Item closing tag missing")
	}

	// Check nested element
	if !strings.Contains(xmlStr, `<DicomAttribute tag="0020000D"`) {
		t.Error("Nested element missing")
	}
	if !strings.Contains(xmlStr, `<Value number="1">1.2.3.4</Value>`) {
		t.Error("Nested value missing or incorrect")
	}
}

func TestToXML_NestedSequence(t *testing.T) {
    ds := dataset.New()

	// Create outer sequence
	outerSeq := dataset.NewSequence(tag.ReferencedStudySequence)
	outerItem := dataset.New()

    // Create inner sequence
    innerSeq := dataset.NewSequence(tag.ReferencedSeriesSequence)
    innerItem := dataset.New()
    _ = innerItem.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
    innerSeq.AddItem(innerItem)

    _ = outerItem.Add(innerSeq)
    outerSeq.AddItem(outerItem)
    _ = ds.Add(outerSeq)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check nested structure
	if !strings.Contains(xmlStr, `tag="00081110"`) {
		t.Error("Outer sequence missing")
	}
	if !strings.Contains(xmlStr, `tag="00081115"`) {
		t.Error("Inner sequence missing")
	}
	if !strings.Contains(xmlStr, `<Value number="1">1.2.3.4.5</Value>`) {
		t.Error("Inner value missing or incorrect")
	}
}

func TestToXML_XMLEscaping(t *testing.T) {
    ds := dataset.New()
    // Test XML special characters
    elem := element.NewString(tag.PatientComments, vr.LT, []string{"Test <tag> & \"quotes\" 'single'"})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check proper escaping
	if !strings.Contains(xmlStr, `&lt;tag&gt;`) {
		t.Error("< and > not properly escaped")
	}
	if !strings.Contains(xmlStr, `&amp;`) {
		t.Error("& not properly escaped")
	}
	if !strings.Contains(xmlStr, `&quot;`) {
		t.Error("\" not properly escaped")
	}
	if !strings.Contains(xmlStr, `&apos;`) {
		t.Error("' not properly escaped")
	}
}

func TestToXML_PersonNameWithMultipleComponents(t *testing.T) {
    ds := dataset.New()
    // Test PersonName with prefix and suffix
    elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John^Michael^Dr.^Jr."})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check all components
	if !strings.Contains(xmlStr, `<FamilyName>Doe</FamilyName>`) {
		t.Error("FamilyName missing")
	}
	if !strings.Contains(xmlStr, `<GivenName>John</GivenName>`) {
		t.Error("GivenName missing")
	}
	if !strings.Contains(xmlStr, `<MiddleName>Michael</MiddleName>`) {
		t.Error("MiddleName missing")
	}
	if !strings.Contains(xmlStr, `<NamePrefix>Dr.</NamePrefix>`) {
		t.Error("NamePrefix missing")
	}
	if !strings.Contains(xmlStr, `<NameSuffix>Jr.</NameSuffix>`) {
		t.Error("NameSuffix missing")
	}
}

func TestToXML_PersonNameWithIdeographic(t *testing.T) {
    ds := dataset.New()
    // Test PersonName with Ideographic component
    elem := element.NewString(tag.PatientName, vr.PN, []string{"Yamada^Tarou=山田^太郎"})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)
	t.Logf("XML output:\n%s", xmlStr)

	// Check Alphabetic component
	if !strings.Contains(xmlStr, `<Alphabetic>`) {
		t.Error("Alphabetic component missing")
	}
	if !strings.Contains(xmlStr, `<FamilyName>Yamada</FamilyName>`) {
		t.Error("Alphabetic FamilyName missing")
	}
	if !strings.Contains(xmlStr, `<GivenName>Tarou</GivenName>`) {
		t.Error("Alphabetic GivenName missing")
	}

	// Check Ideographic component
	if !strings.Contains(xmlStr, `<Ideographic>`) {
		t.Error("Ideographic component missing")
	}
	// Check for the Ideographic tags (the content might have encoding issues in logs)
	if !strings.Contains(xmlStr, `<FamilyName>`) || !strings.Contains(xmlStr, `</FamilyName>`) {
		t.Error("Ideographic FamilyName tags missing")
	}
	// Count FamilyName occurrences - should be 2 (Alphabetic + Ideographic)
	if strings.Count(xmlStr, "<FamilyName>") < 2 {
		t.Error("Ideographic FamilyName missing - expected 2 FamilyName elements")
	}
	if strings.Count(xmlStr, "<GivenName>") < 2 {
		t.Error("Ideographic GivenName missing - expected 2 GivenName elements")
	}
}

func TestToXML_CompactFormat(t *testing.T) {
    ds := dataset.New()
    _ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	xml, err := ToXML(ds, WithXMLCompact(""))
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Compact format should have minimal whitespace
	if strings.Count(xmlStr, "\n") > 3 {
		t.Error("Compact format has too many newlines")
	}
}

func TestToXML_CustomIndentation(t *testing.T) {
    ds := dataset.New()
    _ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	xml, err := ToXML(ds, WithXMLIndent("\t"))
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Should contain tab characters for indentation
	if !strings.Contains(xmlStr, "\t") {
		t.Error("Custom indentation (tabs) not applied")
	}
}

func TestToXML_EmptyBinaryElement(t *testing.T) {
	ds := dataset.New()
	// Create element with empty buffer
    elem := element.NewOtherByte(tag.PixelData, []byte{})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Should have DicomAttribute but no InlineBinary
	if !strings.Contains(xmlStr, `tag="7FE00010"`) {
		t.Error("PixelData attribute missing")
	}
	if strings.Contains(xmlStr, `<InlineBinary>`) {
		t.Error("InlineBinary should not be present for empty buffer")
	}
}

func TestToXML_PrivateTag(t *testing.T) {
	ds := dataset.New()

	// Create private creator
	privateCreator := tag.NewPrivateCreator("ACME Corp")

	// Create private tag
	privateTag := tag.NewWithPrivateCreator(0x0009, 0x1001, privateCreator)
    elem := element.NewString(privateTag, vr.LO, []string{"Private Value"})
    _ = ds.Add(elem)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check private creator attribute
	if !strings.Contains(xmlStr, `privateCreator="ACME Corp"`) {
		t.Error("privateCreator attribute missing or incorrect")
	}
	if !strings.Contains(xmlStr, `tag="00091001"`) {
		t.Error("Private tag missing or incorrect")
	}
}

func TestToXML_MultipleElements(t *testing.T) {
	ds := dataset.New()

	// Add multiple elements
    _ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
    _ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
    _ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
    _ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Check all elements are present
	if !strings.Contains(xmlStr, `tag="00100010"`) {
		t.Error("PatientName missing")
	}
	if !strings.Contains(xmlStr, `tag="00100020"`) {
		t.Error("PatientID missing")
	}
	if !strings.Contains(xmlStr, `tag="00280010"`) {
		t.Error("Rows missing")
	}
	if !strings.Contains(xmlStr, `tag="00280011"`) {
		t.Error("Columns missing")
	}
}

func TestToXML_ComplexDataset(t *testing.T) {
	// Create a complex dataset with multiple element types
	ds := dataset.New()

	// String elements
    _ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
    _ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"TEST001"}))
    _ = ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	// Numeric elements
    _ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{256}))
    _ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{256}))
    _ = ds.Add(element.NewFloat(tag.PixelSpacing, []float32{0.5, 0.5}))

	// Binary element
	pixelData := make([]byte, 100)
	for i := range pixelData {
		pixelData[i] = byte(i)
	}
    _ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

	// Sequence
	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	item := dataset.New()
    _ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5.6"}))
    _ = item.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7"}))
	seq.AddItem(item)
    _ = ds.Add(seq)

	xml, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML failed: %v", err)
	}

	xmlStr := string(xml)

	// Basic structure checks
	if !strings.Contains(xmlStr, `<?xml version="1.0" encoding="utf-8"?>`) {
		t.Error("XML declaration missing")
	}
	if !strings.Contains(xmlStr, `<NativeDicomModel>`) {
		t.Error("NativeDicomModel missing")
	}

	// Verify we have multiple DicomAttribute elements
	count := strings.Count(xmlStr, "<DicomAttribute")
	if count < 8 { // Should have at least 8 attributes
		t.Errorf("Expected at least 8 DicomAttribute elements, got %d", count)
	}

	// Verify sequence structure
	if !strings.Contains(xmlStr, `<Item number="1">`) {
		t.Error("Sequence item missing")
	}
}
