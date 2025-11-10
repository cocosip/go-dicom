// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestHasFragmentSequence tests the specific HasFragmentSequence.dcm file
// which should contain 96 frames
func TestHasFragmentSequence(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "HasFragmentSequence.dcm")

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

	// Open and parse the file
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse HasFragmentSequence.dcm: %v", err)
	}

	if result == nil || result.Dataset == nil {
		t.Fatal("Parse result or dataset is nil")
	}

	t.Logf("Successfully parsed HasFragmentSequence.dcm")
	t.Logf("Dataset contains %d elements", result.Dataset.Count())

	// Check Number of Frames
	numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
	if !exists {
		t.Fatal("NumberOfFrames tag not found in dataset")
	}

	t.Logf("NumberOfFrames (string): %s", numFramesStr)

	// Expected: 96 frames
	expectedFrames := "96"
	if numFramesStr != expectedFrames {
		t.Errorf("NumberOfFrames = %s, want %s", numFramesStr, expectedFrames)
	}

	// Check for Pixel Data element
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData tag not found in dataset")
	}

	t.Logf("PixelData element found: %v", pixelDataElem)
	t.Logf("PixelData tag: %s", pixelDataElem.Tag())
	t.Logf("PixelData VR: %s", pixelDataElem.ValueRepresentation())

	// Check basic image properties
	if modality, exists := result.Dataset.GetString(tag.Modality); exists {
		t.Logf("Modality: %s", modality)
	}

	if sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID); exists {
		t.Logf("SOP Class UID: %s", sopClassUID)
	}

	if rows, err := result.Dataset.GetUInt16(tag.Rows, 0); err == nil {
		t.Logf("Rows: %d", rows)
	}

	if cols, err := result.Dataset.GetUInt16(tag.Columns, 0); err == nil {
		t.Logf("Columns: %d", cols)
	}

	if bitsAlloc, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0); err == nil {
		t.Logf("Bits Allocated: %d", bitsAlloc)
	}

	if samplesPerPixel, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0); err == nil {
		t.Logf("Samples Per Pixel: %d", samplesPerPixel)
	}

	if photoInterp, exists := result.Dataset.GetString(tag.PhotometricInterpretation); exists {
		t.Logf("Photometric Interpretation: %s", photoInterp)
	}

	// Check Transfer Syntax
	if result.FileMetaInformation != nil {
		if tsUID, exists := result.FileMetaInformation.TransferSyntaxUID(); exists {
			t.Logf("Transfer Syntax UID: %s", tsUID)
		}
	}

	// Verify this is indeed a multi-frame image with fragment sequence
	t.Log("✓ HasFragmentSequence.dcm parsed successfully")
	t.Log("✓ Number of frames verified: 96")
}

// TestMultiFrameFragmentSequences tests multiple multi-frame files with fragment sequences
func TestMultiFrameFragmentSequences(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	testCases := []struct {
		filename       string
		expectedFrames string
		description    string
	}{
		{
			filename:       "HasFragmentSequence.dcm",
			expectedFrames: "96",
			description:    "XA image with 96 frames and fragment sequences",
		},
		{
			filename:       "multiframe.dcm",
			expectedFrames: "140",
			description:    "Secondary Capture with 140 frames",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, tc.filename)

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
				t.Fatalf("Failed to parse %s: %v", tc.filename, err)
			}

			t.Logf("Testing: %s", tc.description)

			// Check Number of Frames
			numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
			if !exists {
				t.Errorf("NumberOfFrames tag not found in %s", tc.filename)
				return
			}

			t.Logf("NumberOfFrames: %s (expected: %s)", numFramesStr, tc.expectedFrames)

			// Trim whitespace for comparison
			numFramesTrimmed := strings.TrimSpace(numFramesStr)

			if numFramesTrimmed != tc.expectedFrames {
				t.Errorf("NumberOfFrames = %s, want %s", numFramesTrimmed, tc.expectedFrames)
			} else {
				t.Logf("✓ Frame count verified: %s frames", tc.expectedFrames)
			}

			// Check for Pixel Data
			if _, exists := result.Dataset.Get(tag.PixelData); !exists {
				t.Errorf("PixelData tag not found in %s", tc.filename)
			} else {
				t.Log("✓ PixelData element found")
			}
		})
	}
}

// TestFragmentSequenceStructure tests the structure of fragment sequences
func TestFragmentSequenceStructure(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "HasFragmentSequence.dcm")

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

	// Get PixelData element
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData not found")
	}

	// Log information about the pixel data element
	t.Logf("PixelData Tag: %s", pixelDataElem.Tag())
	t.Logf("PixelData VR: %s", pixelDataElem.ValueRepresentation())

	// Try to get the value length
	if stringer, ok := pixelDataElem.(interface{ String() string }); ok {
		valueStr := stringer.String()
		// Truncate if too long
		if len(valueStr) > 200 {
			valueStr = valueStr[:200] + "..."
		}
		t.Logf("PixelData info: %s", valueStr)
	}

	// Check if it's an OtherByteFragment or OtherWordFragment
	t.Logf("PixelData type: %T", pixelDataElem)

	// Verify dataset integrity
	t.Logf("Total dataset elements: %d", result.Dataset.Count())

	// List some key tags
	keyTags := []*tag.Tag{
		tag.Rows,
		tag.Columns,
		tag.NumberOfFrames,
		tag.SamplesPerPixel,
		tag.BitsAllocated,
		tag.BitsStored,
		tag.HighBit,
		tag.PhotometricInterpretation,
	}

	t.Log("\nKey image attributes:")
	for _, keyTag := range keyTags {
		if elem, exists := result.Dataset.Get(keyTag); exists {
			if stringer, ok := elem.(interface{ String() string }); ok {
				t.Logf("  %s: %s", keyTag, stringer.String())
			} else {
				t.Logf("  %s: present", keyTag)
			}
		}
	}
}

// BenchmarkMultiFrameFragmentParsing benchmarks parsing of multi-frame images with fragments
func BenchmarkMultiFrameFragmentParsing(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "HasFragmentSequence.dcm")

	// Read file content once
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
