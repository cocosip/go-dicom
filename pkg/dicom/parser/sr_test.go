// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

const textValueType = "TEXT"

// TestStructuredReportParsing tests parsing of Structured Report DICOM files
func TestStructuredReportParsing(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "test_SR.dcm")

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	t.Logf("File size: %.2f KB", float64(fileInfo.Size())/1024)

	// Open and parse the file
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	t.Log("Parsing test_SR.dcm (Structured Report)...")
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse test_SR.dcm: %v", err)
	}

	if result == nil || result.Dataset == nil {
		t.Fatal("Parse result or dataset is nil")
	}

	t.Logf("✓ Successfully parsed test_SR.dcm")
	t.Logf("Dataset contains %d elements", result.Dataset.Count())

	// Check SOP Class UID - should be a SR-related class
	t.Log("\n=== SR Identification ===")
	if sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID); exists {
		t.Logf("SOP Class UID: %s", sopClassUID)

		// Common SR SOP Classes:
		// Basic Text SR: 1.2.840.10008.5.1.4.1.1.88.11
		// Enhanced SR: 1.2.840.10008.5.1.4.1.1.88.22
		// Comprehensive SR: 1.2.840.10008.5.1.4.1.1.88.33
		// Mammography CAD SR: 1.2.840.10008.5.1.4.1.1.88.50
		switch sopClassUID {
		case "1.2.840.10008.5.1.4.1.1.88.11":
			t.Log("✓ SR Type: Basic Text SR")
		case "1.2.840.10008.5.1.4.1.1.88.22":
			t.Log("✓ SR Type: Enhanced SR")
		case "1.2.840.10008.5.1.4.1.1.88.33":
			t.Log("✓ SR Type: Comprehensive SR")
		case "1.2.840.10008.5.1.4.1.1.88.50":
			t.Log("✓ SR Type: Mammography CAD SR")
		default:
			if len(sopClassUID) > 20 && sopClassUID[:20] == "1.2.840.10008.5.1.4." {
				t.Logf("✓ SR Type: Structured Report (UID: %s)", sopClassUID)
			} else {
				t.Logf("SOP Class: %s", sopClassUID)
			}
		}
	} else {
		t.Error("SOP Class UID not found")
	}

	// Check Modality
	if modality, exists := result.Dataset.GetString(tag.Modality); exists {
		t.Logf("Modality: %s", modality)
		if modality == "SR" || modality == "SR " {
			t.Log("✓ Confirmed: Structured Report modality")
		}
	}

	// Patient and Study Information
	t.Log("\n=== Patient/Study Information ===")
	if patientName, exists := result.Dataset.GetString(tag.PatientName); exists {
		t.Logf("Patient Name: %s", patientName)
	}

	if patientID, exists := result.Dataset.GetString(tag.PatientID); exists {
		t.Logf("Patient ID: %s", patientID)
	}

	if studyDesc, exists := result.Dataset.GetString(tag.StudyDescription); exists {
		t.Logf("Study Description: %s", studyDesc)
	}

	if seriesDesc, exists := result.Dataset.GetString(tag.SeriesDescription); exists {
		t.Logf("Series Description: %s", seriesDesc)
	}

	// SR-specific attributes
	t.Log("\n=== SR-Specific Attributes ===")

	// Verification Flag
	if verificationFlag, exists := result.Dataset.GetString(tag.VerificationFlag); exists {
		t.Logf("Verification Flag: %s", verificationFlag)
	}

	// Completion Flag
	if completionFlag, exists := result.Dataset.GetString(tag.CompletionFlag); exists {
		t.Logf("Completion Flag: %s", completionFlag)
	}

	// Content Date and Time
	if contentDate, exists := result.Dataset.GetString(tag.ContentDate); exists {
		t.Logf("Content Date: %s", contentDate)
	}

	if contentTime, exists := result.Dataset.GetString(tag.ContentTime); exists {
		t.Logf("Content Time: %s", contentTime)
	}

	// Check for Content Sequence - the main SR content
	t.Log("\n=== SR Content Sequence ===")
	contentSeqElem, exists := result.Dataset.Get(tag.ContentSequence)
	if !exists {
		t.Fatal("❌ ContentSequence not found - SR content not accessible")
	}

	t.Log("✓ ContentSequence found")

	// Cast to Sequence to access items
	contentSeq, ok := contentSeqElem.(*dataset.Sequence)
	if !ok {
		t.Fatalf("ContentSequence is not a Sequence type, got: %T", contentSeqElem)
	}

	itemCount := contentSeq.Count()
	t.Logf("✓ Content Sequence contains %d items", itemCount)

	if itemCount == 0 {
		t.Error("Warning: Content Sequence is empty")
	} else {
		t.Log("✓ SR has content items")
	}

	// Transfer Syntax
	if result.FileMetaInformation != nil {
		if tsUID, exists := result.FileMetaInformation.TransferSyntaxUID(); exists {
			t.Logf("\nTransfer Syntax: %s", tsUID)
		}
	}

	t.Log("\n✓ Structured Report DICOM parsed successfully")
	t.Log("✓ SR content accessible")
}

// TestSRContentSequenceDetails tests detailed parsing of SR content
func TestSRContentSequenceDetails(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "test_SR.dcm")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Get Content Sequence
	contentSeqElem, exists := result.Dataset.Get(tag.ContentSequence)
	if !exists {
		t.Fatal("ContentSequence not found")
	}

	contentSeq, ok := contentSeqElem.(*dataset.Sequence)
	if !ok {
		t.Fatalf("ContentSequence is not a Sequence")
	}

	t.Logf("Analyzing Content Sequence with %d items...\n", contentSeq.Count())

	// Analyze each content item
	for i := 0; i < contentSeq.Count(); i++ {
		item := contentSeq.GetItem(i)
		if item == nil {
			t.Logf("Item %d: nil", i)
			continue
		}
		analyzeSRContentItem(t, i, item)
	}

	t.Logf("✓ Analyzed %d content items", contentSeq.Count())
}

// TestSRContentExtraction tests extracting structured content from SR
func TestSRContentExtraction(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "test_SR.dcm")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Extract SR content as structured data
	t.Log("=== Extracting SR Content ===\n")

	content := extractSRContent(result.Dataset)
	if len(content) == 0 {
		t.Error("No SR content extracted")
		return
	}

	t.Logf("✓ Extracted %d content items\n", len(content))

	// Display extracted content
	for i, item := range content {
		t.Logf("Item %d:", i+1)
		for key, value := range item {
			t.Logf("  %s: %s", key, value)
		}
		t.Log("")
	}

	t.Log("✓ SR content successfully extracted and accessible")
}

// TestSRHierarchy tests the hierarchical structure of SR
func TestSRHierarchy(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "test_SR.dcm")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	t.Log("=== SR Hierarchical Structure ===\n")

	// Print the SR tree structure
	printSRTree(t, result.Dataset, 0)

	t.Log("\n✓ SR hierarchy accessible")
}

// Helper function: Extract SR content as map
func extractSRContent(ds *dataset.Dataset) []map[string]string {
	var content []map[string]string

	contentSeqElem, exists := ds.Get(tag.ContentSequence)
	if !exists {
		return content
	}

	contentSeq, ok := contentSeqElem.(*dataset.Sequence)
	if !ok {
		return content
	}

	for i := 0; i < contentSeq.Count(); i++ {
		item := contentSeq.GetItem(i)
		if item == nil {
			continue
		}

		itemData := make(map[string]string)

		// Extract common fields
		if relType, exists := item.GetString(tag.RelationshipType); exists {
			itemData["RelationshipType"] = relType
		}

		if valueType, exists := item.GetString(tag.ValueType); exists {
			itemData["ValueType"] = valueType

			// Extract value based on type
			switch valueType {
			case textValueType:
				if textValue, exists := item.GetString(tag.TextValue); exists {
					itemData["Value"] = textValue
				}
			case "CODE":
				if codeSeqElem, exists := item.Get(tag.ConceptCodeSequence); exists {
					if codeSeq, ok := codeSeqElem.(*dataset.Sequence); ok && codeSeq.Count() > 0 {
						codeItem := codeSeq.GetItem(0)
						if codeMeaning, exists := codeItem.GetString(tag.CodeMeaning); exists {
							itemData["Value"] = codeMeaning
						}
					}
				}
			}
		}

		// Get concept name
		if conceptNameSeqElem, exists := item.Get(tag.ConceptNameCodeSequence); exists {
			if conceptSeq, ok := conceptNameSeqElem.(*dataset.Sequence); ok && conceptSeq.Count() > 0 {
				conceptItem := conceptSeq.GetItem(0)
				if codeMeaning, exists := conceptItem.GetString(tag.CodeMeaning); exists {
					itemData["Concept"] = codeMeaning
				}
			}
		}

		if len(itemData) > 0 {
			content = append(content, itemData)
		}
	}

	return content
}

// Helper function: Print SR tree structure
func printSRTree(t *testing.T, ds *dataset.Dataset, level int) {
	contentSeqElem, exists := ds.Get(tag.ContentSequence)
	if !exists {
		return
	}

	contentSeq, ok := contentSeqElem.(*dataset.Sequence)
	if !ok {
		return
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	for i := 0; i < contentSeq.Count(); i++ {
		item := contentSeq.GetItem(i)
		if item == nil {
			continue
		}

		// Get relationship and value type
		relType := "?"
		if rt, exists := item.GetString(tag.RelationshipType); exists {
			relType = rt
		}

		valueType := "?"
		if vt, exists := item.GetString(tag.ValueType); exists {
			valueType = vt
		}

		// Get concept name
		conceptName := ""
		if conceptNameSeqElem, exists := item.Get(tag.ConceptNameCodeSequence); exists {
			if conceptSeq, ok := conceptNameSeqElem.(*dataset.Sequence); ok && conceptSeq.Count() > 0 {
				conceptItem := conceptSeq.GetItem(0)
				if codeMeaning, exists := conceptItem.GetString(tag.CodeMeaning); exists {
					conceptName = codeMeaning
				}
			}
		}

		// Print this item
		t.Logf("%s[%s] %s: %s", indent, relType, valueType, conceptName)

		// Get value if TEXT
		if valueType == textValueType {
			if textValue, exists := item.GetString(tag.TextValue); exists {
				t.Logf("%s  → %s", indent, textValue)
			}
		}

		// Recursively print nested content
		if valueType == "CONTAINER" {
			printSRTree(t, item, level+1)
		}
	}
}

// BenchmarkSRParsing benchmarks parsing of SR DICOM files
func BenchmarkSRParsing(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "test_SR.dcm")

	data, err := os.ReadFile(filePath)
	if err != nil {
		b.Skipf("Cannot read file: %v", err)
		return
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		reader := &fileReader{data: data, pos: 0}
		_, _ = Parse(reader)
	}
}

// analyzeSRContentItem analyzes a single SR content item
func analyzeSRContentItem(t *testing.T, index int, item *dataset.Dataset) {
	t.Logf("=== Content Item %d ===", index)

	// Relationship Type (required)
	if relType, exists := item.GetString(tag.RelationshipType); exists {
		t.Logf("  Relationship Type: %s", relType)
	}

	// Value Type (required)
	if valueType, exists := item.GetString(tag.ValueType); exists {
		t.Logf("  Value Type: %s", valueType)
		logContentByValueType(t, item, valueType)
	}

	// Concept Name Code Sequence (describes what this item represents)
	logConceptName(t, item)

	t.Log("")
}

// logContentByValueType logs content based on value type
func logContentByValueType(t *testing.T, item *dataset.Dataset, valueType string) {
	switch valueType {
	case textValueType:
		if textValue, exists := item.GetString(tag.TextValue); exists {
			t.Logf("  Text Value: %s", textValue)
		}

	case "CODE":
		logCodeSequence(t, item)

	case "NUM":
		logNumericValue(t, item)

	case "DATETIME", "DATE", "TIME":
		logDateTimeValues(t, item)

	case "PNAME":
		if personName, exists := item.GetString(tag.PersonName); exists {
			t.Logf("  Person Name: %s", personName)
		}

	case "UIDREF":
		if uidRef, exists := item.GetString(tag.UID); exists {
			t.Logf("  UID Reference: %s", uidRef)
		}

	case "CONTAINER":
		logContainerInfo(t, item)

	case "IMAGE":
		logImageReference(t, item)

	default:
		t.Logf("  [Other value type: %s]", valueType)
	}
}

// logCodeSequence logs code sequence information
func logCodeSequence(t *testing.T, item *dataset.Dataset) {
	if codeSeqElem, exists := item.Get(tag.ConceptCodeSequence); exists {
		if codeSeq, ok := codeSeqElem.(*dataset.Sequence); ok && codeSeq.Count() > 0 {
			codeItem := codeSeq.GetItem(0)
			if codeValue, exists := codeItem.GetString(tag.CodeValue); exists {
				t.Logf("  Code Value: %s", codeValue)
			}
			if codeMeaning, exists := codeItem.GetString(tag.CodeMeaning); exists {
				t.Logf("  Code Meaning: %s", codeMeaning)
			}
		}
	}
}

// logNumericValue logs numeric value information
func logNumericValue(t *testing.T, item *dataset.Dataset) {
	if measuredValueSeqElem, exists := item.Get(tag.MeasuredValueSequence); exists {
		if measuredSeq, ok := measuredValueSeqElem.(*dataset.Sequence); ok && measuredSeq.Count() > 0 {
			measuredItem := measuredSeq.GetItem(0)
			if numValue, exists := measuredItem.GetString(tag.NumericValue); exists {
				t.Logf("  Numeric Value: %s", numValue)
			}
		}
	}
}

// logDateTimeValues logs date/time values
func logDateTimeValues(t *testing.T, item *dataset.Dataset) {
	if dateTime, exists := item.GetString(tag.DateTime); exists {
		t.Logf("  DateTime: %s", dateTime)
	}
	if date, exists := item.GetString(tag.Date); exists {
		t.Logf("  Date: %s", date)
	}
	if time, exists := item.GetString(tag.Time); exists {
		t.Logf("  Time: %s", time)
	}
}

// logContainerInfo logs container information
func logContainerInfo(t *testing.T, item *dataset.Dataset) {
	t.Log("  [Container - may have nested content]")
	if nestedSeqElem, exists := item.Get(tag.ContentSequence); exists {
		if nestedSeq, ok := nestedSeqElem.(*dataset.Sequence); ok {
			t.Logf("  Nested items: %d", nestedSeq.Count())
		}
	}
}

// logImageReference logs image reference information
func logImageReference(t *testing.T, item *dataset.Dataset) {
	if refSOPSeqElem, exists := item.Get(tag.ReferencedSOPSequence); exists {
		if refSeq, ok := refSOPSeqElem.(*dataset.Sequence); ok && refSeq.Count() > 0 {
			t.Logf("  Referenced Image: %d items", refSeq.Count())
		}
	}
}

// logConceptName logs concept name information
func logConceptName(t *testing.T, item *dataset.Dataset) {
	if conceptNameSeqElem, exists := item.Get(tag.ConceptNameCodeSequence); exists {
		if conceptSeq, ok := conceptNameSeqElem.(*dataset.Sequence); ok && conceptSeq.Count() > 0 {
			conceptItem := conceptSeq.GetItem(0)
			if codeMeaning, exists := conceptItem.GetString(tag.CodeMeaning); exists {
				t.Logf("  Concept Name: %s", codeMeaning)
			}
		}
	}
}
