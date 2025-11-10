// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"testing"
)

func TestCodeItem_Creation(t *testing.T) {
	code := NewCodeItem("113704", "DCM", "SR Document")

	if code.Value() != "113704" {
		t.Errorf("Expected value '113704', got '%s'", code.Value())
	}

	if code.Scheme() != "DCM" {
		t.Errorf("Expected scheme 'DCM', got '%s'", code.Scheme())
	}

	if code.Meaning() != "SR Document" {
		t.Errorf("Expected meaning 'SR Document', got '%s'", code.Meaning())
	}

	if code.Version() != "" {
		t.Errorf("Expected empty version, got '%s'", code.Version())
	}
}

func TestCodeItem_WithVersion(t *testing.T) {
	code := NewCodeItemWithVersion("12345", "SNOMED", "Finding", "2021")

	if code.Version() != "2021" {
		t.Errorf("Expected version '2021', got '%s'", code.Version())
	}
}

func TestCodeItem_String(t *testing.T) {
	testCases := []struct {
		name     string
		code     *CodeItem
		expected string
	}{
		{
			name:     "without version",
			code:     NewCodeItem("113704", "DCM", "SR Document"),
			expected: "(113704,DCM,\"SR Document\")",
		},
		{
			name:     "with version",
			code:     NewCodeItemWithVersion("12345", "SNOMED", "Finding", "2021"),
			expected: "(12345,SNOMED:2021,\"Finding\")",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.code.String()
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

func TestCodeItem_Equals(t *testing.T) {
	code1 := NewCodeItem("113704", "DCM", "SR Document")
	code2 := NewCodeItem("113704", "DCM", "SR Document")
	code3 := NewCodeItem("113704", "DCM", "Different Meaning")
	code4 := NewCodeItem("12345", "DCM", "SR Document")

	if !code1.Equals(code2) {
		t.Error("Identical codes should be equal")
	}

	if !code1.Equals(code3) {
		t.Error("Codes with same value and scheme should be equal (meaning doesn't matter)")
	}

	if code1.Equals(code4) {
		t.Error("Codes with different values should not be equal")
	}

	// Test nil cases
	var nilCode *CodeItem
	if code1.Equals(nilCode) {
		t.Error("Code should not equal nil")
	}

	if !nilCode.Equals(nilCode) {
		t.Error("Nil should equal nil")
	}
}

func TestMeasuredValue_Creation(t *testing.T) {
	units := NewCodeItem("mm", "UCUM", "millimeter")
	mv := NewMeasuredValue(25.5, units)

	if mv.Value() != 25.5 {
		t.Errorf("Expected value 25.5, got %f", mv.Value())
	}

	retrievedUnits, err := mv.Units()
	if err != nil {
		t.Fatalf("Failed to get units: %v", err)
	}

	if !retrievedUnits.Equals(units) {
		t.Error("Retrieved units do not match original units")
	}
}

func TestMeasuredValue_String(t *testing.T) {
	units := NewCodeItem("kg", "UCUM", "kilogram")
	mv := NewMeasuredValue(75.3, units)

	expected := "75.300000 kg"
	result := mv.String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestReferencedSOP_Creation(t *testing.T) {
	instanceUID := "1.2.840.10008.5.1.4.1.1.2.1.12345"
	classUID := "1.2.840.10008.5.1.4.1.1.2"

	ref := NewReferencedSOP(instanceUID, classUID)

	if ref.InstanceUID() != instanceUID {
		t.Errorf("Expected instance UID '%s', got '%s'", instanceUID, ref.InstanceUID())
	}

	if ref.ClassUID() != classUID {
		t.Errorf("Expected class UID '%s', got '%s'", classUID, ref.ClassUID())
	}
}

func TestValueType_String(t *testing.T) {
	testCases := []struct {
		vt       ValueType
		expected string
	}{
		{ValueTypeText, "TEXT"},
		{ValueTypeCode, "CODE"},
		{ValueTypeNumeric, "NUM"},
		{ValueTypeContainer, "CONTAINER"},
	}

	for _, tc := range testCases {
		if tc.vt.String() != tc.expected {
			t.Errorf("ValueType.String(): expected '%s', got '%s'", tc.expected, tc.vt.String())
		}
	}
}

func TestRelationship_String(t *testing.T) {
	testCases := []struct {
		rel      Relationship
		expected string
	}{
		{RelationshipContains, "CONTAINS"},
		{RelationshipHasProperties, "HAS PROPERTIES"},
		{RelationshipInferredFrom, "INFERRED FROM"},
	}

	for _, tc := range testCases {
		if tc.rel.String() != tc.expected {
			t.Errorf("Relationship.String(): expected '%s', got '%s'", tc.expected, tc.rel.String())
		}
	}
}

func TestError_Error(t *testing.T) {
	err := NewError("test error")
	expected := "sr: test error"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestError_Wrap(t *testing.T) {
	innerErr := NewError("inner error")
	wrappedErr := WrapError("wrapper", innerErr)

	expected := "sr: wrapper: sr: inner error"
	if wrappedErr.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, wrappedErr.Error())
	}
}

func TestContentItem_Text(t *testing.T) {
	code := NewCodeItem("121071", "DCM", "Finding")
	item, err := NewContentItemText(code, RelationshipContains, "No abnormalities detected")

	if err != nil {
		t.Fatalf("Failed to create text content item: %v", err)
	}

	vt, err := item.ValueType()
	if err != nil {
		t.Fatalf("Failed to get value type: %v", err)
	}

	if vt != ValueTypeText {
		t.Errorf("Expected value type TEXT, got %s", vt)
	}

	text, err := item.GetText()
	if err != nil {
		t.Fatalf("Failed to get text: %v", err)
	}

	if text != "No abnormalities detected" {
		t.Errorf("Expected text 'No abnormalities detected', got '%s'", text)
	}
}

func TestContentItem_Numeric(t *testing.T) {
	code := NewCodeItem("33728-7", "LN", "Size")
	units := NewCodeItem("mm", "UCUM", "millimeter")
	value := NewMeasuredValue(25.5, units)

	item, err := NewContentItemNumeric(code, RelationshipContains, value)
	if err != nil {
		t.Fatalf("Failed to create numeric content item: %v", err)
	}

	vt, err := item.ValueType()
	if err != nil {
		t.Fatalf("Failed to get value type: %v", err)
	}

	if vt != ValueTypeNumeric {
		t.Errorf("Expected value type NUM, got %s", vt)
	}

	mv, err := item.GetNumeric()
	if err != nil {
		t.Fatalf("Failed to get numeric value: %v", err)
	}

	if mv.Value() != 25.5 {
		t.Errorf("Expected numeric value 25.5, got %f", mv.Value())
	}
}

func TestStructuredReport_Creation(t *testing.T) {
	rootCode := NewCodeItem("113704", "DCM", "SR Document")
	report, err := NewStructuredReport(rootCode)

	if err != nil {
		t.Fatalf("Failed to create structured report: %v", err)
	}

	vt, err := report.ValueType()
	if err != nil {
		t.Fatalf("Failed to get value type: %v", err)
	}

	if vt != ValueTypeContainer {
		t.Errorf("Expected root to be CONTAINER type, got %s", vt)
	}

	// Root should not have a relationship type
	_, err = report.Relationship()
	if err == nil {
		t.Error("Root element should not have a relationship type")
	}
}

func TestStructuredReport_AddItems(t *testing.T) {
	rootCode := NewCodeItem("113704", "DCM", "SR Document")
	report, err := NewStructuredReport(rootCode)
	if err != nil {
		t.Fatalf("Failed to create structured report: %v", err)
	}

	// Add text item
	err = report.AddText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		"No abnormalities detected",
	)
	if err != nil {
		t.Fatalf("Failed to add text item: %v", err)
	}

	// Add numeric item
	units := NewCodeItem("mm", "UCUM", "millimeter")
	value := NewMeasuredValue(25.5, units)
	err = report.AddNumeric(
		NewCodeItem("33728-7", "LN", "Size"),
		RelationshipContains,
		value,
	)
	if err != nil {
		t.Fatalf("Failed to add numeric item: %v", err)
	}

	// Get children
	children, err := report.Children()
	if err != nil {
		t.Fatalf("Failed to get children: %v", err)
	}

	if len(children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(children))
	}
}
