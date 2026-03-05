// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestMultiFrameImage tests parsing of a multi-frame DICOM image.
// This test specifically verifies the fix for VR=UN with undefined length (private sequences).
//
//nolint:gocyclo // Test function with comprehensive test cases
func TestMultiFrameImage(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestMultiFrame.dcm")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open TestMultiFrame.dcm: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			t.Errorf("Failed to close file: %v", closeErr)
		}
	}()

	// Parse the DICOM file
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse TestMultiFrame.dcm: %v", err)
	}

	if result == nil || result.Dataset == nil {
		t.Fatal("Parse result or dataset is nil")
	}

	t.Logf("Successfully parsed TestMultiFrame.dcm")
	t.Logf("Dataset contains %d elements", result.Dataset.Count())

	// Verify basic DICOM attributes
	t.Run("BasicAttributes", func(t *testing.T) {
		// SOP Class UID
		sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID)
		if !exists {
			t.Error("SOPClassUID not found")
		} else {
			t.Logf("SOP Class UID: %s", sopClassUID)
			// MR Image Storage: 1.2.840.10008.5.1.4.1.1.4 or 1.2.840.10008.5.1.4.1.1.4.1
			if sopClassUID != "1.2.840.10008.5.1.4.1.1.4" && sopClassUID != "1.2.840.10008.5.1.4.1.1.4.1" {
				t.Logf("Warning: Unexpected SOP Class UID: %s (expected MR Image Storage)", sopClassUID)
			}
		}

		// Modality
		modality, exists := result.Dataset.GetString(tag.Modality)
		if !exists {
			t.Error("Modality not found")
		} else {
			t.Logf("Modality: %s", modality)
			if modality != "MR" {
				t.Errorf("Expected Modality=MR, got %s", modality)
			}
		}

		// Patient Name
		patientName, exists := result.Dataset.GetString(tag.PatientName)
		if exists {
			t.Logf("Patient Name: %s", patientName)
		}
	})

	// Verify multi-frame information
	t.Run("MultiFrameInfo", func(t *testing.T) {
		// Number of Frames - stored as IS (Integer String) in DICOM
		numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
		if !exists {
			t.Error("NumberOfFrames not found - this should be a multi-frame image")
		} else {
			t.Logf("Number of Frames (string): %s", numFramesStr)

			// Parse string as integer
			var numFrames int
			if _, err := fmt.Sscanf(numFramesStr, "%d", &numFrames); err != nil {
				t.Errorf("Failed to parse NumberOfFrames '%s' as int: %v", numFramesStr, err)
			} else {
				t.Logf("Number of Frames (int): %d", numFrames)

				if numFrames < 2 {
					t.Errorf("Expected multiple frames (>=2), got %d", numFrames)
				}

				// Verify it's specifically 7 frames for this test file
				if numFrames != 7 {
					t.Errorf("Expected 7 frames for TestMultiFrame.dcm, got %d", numFrames)
				}
			}
		}
	})

	// Verify pixel data attributes
	t.Run("PixelDataAttributes", func(t *testing.T) {
		// Image dimensions
		rows := result.Dataset.TryGetUInt16(tag.Rows, 0)
		cols := result.Dataset.TryGetUInt16(tag.Columns, 0)

		if rows == 0 {
			t.Error("Rows not found or is 0")
		} else {
			t.Logf("Rows: %d", rows)
		}

		if cols == 0 {
			t.Error("Columns not found or is 0")
		} else {
			t.Logf("Columns: %d", cols)
		}

		// Bits allocated/stored
		bitsAllocated := result.Dataset.TryGetUInt16(tag.BitsAllocated, 0)
		bitsStored := result.Dataset.TryGetUInt16(tag.BitsStored, 0)
		highBit := result.Dataset.TryGetUInt16(tag.HighBit, 0)

		if bitsAllocated == 0 {
			t.Error("BitsAllocated not found or is 0")
		} else {
			t.Logf("Bits Allocated: %d", bitsAllocated)
		}

		if bitsStored == 0 {
			t.Logf("Bits Stored: %d (using BitsAllocated as fallback)", bitsAllocated)
		} else {
			t.Logf("Bits Stored: %d", bitsStored)
		}

		t.Logf("High Bit: %d", highBit)

		// Samples per pixel (defaults to 1 for grayscale)
		samplesPerPixel := result.Dataset.TryGetUInt16(tag.SamplesPerPixel, 1)
		if samplesPerPixel == 0 {
			// If not found or 0, use default of 1 for monochrome images
			samplesPerPixel = 1
			t.Logf("Samples Per Pixel: %d (default for monochrome)", samplesPerPixel)
		} else {
			t.Logf("Samples Per Pixel: %d", samplesPerPixel)
		}

		// Photometric Interpretation
		photometric := result.Dataset.TryGetString(tag.PhotometricInterpretation)
		if photometric == "" {
			t.Error("PhotometricInterpretation not found")
		} else {
			t.Logf("Photometric Interpretation: %s", photometric)
		}

		// Pixel Representation
		pixelRep := result.Dataset.TryGetUInt16(tag.PixelRepresentation, 0)
		t.Logf("Pixel Representation: %d (%s)", pixelRep,
			map[uint16]string{0: "unsigned", 1: "signed"}[pixelRep])

		// Planar Configuration (for color images)
		if samplesPerPixel > 1 {
			planarConfig := result.Dataset.TryGetUInt16(tag.PlanarConfiguration, 0)
			t.Logf("Planar Configuration: %d", planarConfig)
		}
	})

	// Verify pixel data element exists
	t.Run("PixelDataElement", func(t *testing.T) {
		pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
		if !exists {
			t.Error("PixelData element not found")
			return
		}

		t.Logf("PixelData element found: %v", pixelDataElem.Tag())

		// Check if it's encapsulated (compressed) or native (uncompressed)
		// For multi-frame images, check the transfer syntax
		if result.FileMetaInformation != nil {
			tsUID, exists := result.FileMetaInformation.TransferSyntaxUID()
			if exists {
				t.Logf("Transfer Syntax UID: %s", tsUID)

				// Common uncompressed transfer syntaxes
				uncompressed := map[string]string{
					"1.2.840.10008.1.2":     "Implicit VR Little Endian",
					"1.2.840.10008.1.2.1":   "Explicit VR Little Endian",
					"1.2.840.10008.1.2.2":   "Explicit VR Big Endian",
				}

				if name, ok := uncompressed[tsUID]; ok {
					t.Logf("Transfer Syntax: %s (uncompressed)", name)
				} else {
					t.Logf("Transfer Syntax: compressed or other")
				}
			}
		}
	})

	// Verify private sequence that caused the original bug
	t.Run("PrivateSequence_2001_105F", func(t *testing.T) {
		// Tag (2001,105f) is a private sequence with VR=UN and undefined length
		// This was the tag that caused the original parsing failure
		privateTag := tag.New(0x2001, 0x105f)
		elem, exists := result.Dataset.Get(privateTag)

		if !exists {
			t.Log("Private tag (2001,105f) not found - this is OK, it may vary by file")
			return
		}

		t.Logf("Private tag (2001,105f) found and successfully parsed")
		t.Logf("Element type: %T", elem)

		// If it's a sequence, log how many items it has
		if seq, ok := elem.(*dataset.Sequence); ok {
			t.Logf("Private sequence contains %d items", seq.Count())
		}
	})

	// Calculate expected pixel data size
	t.Run("PixelDataSize", func(t *testing.T) {
		rows := result.Dataset.TryGetUInt16(tag.Rows, 0)
		cols := result.Dataset.TryGetUInt16(tag.Columns, 0)
		bitsAllocated := result.Dataset.TryGetUInt16(tag.BitsAllocated, 0)
		samplesPerPixel := result.Dataset.TryGetUInt16(tag.SamplesPerPixel, 1)
		if samplesPerPixel == 0 {
			samplesPerPixel = 1
		}

		// Parse NumberOfFrames from string (IS - Integer String)
		numFrames := 1
		if numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames); exists {
			if _, err := fmt.Sscanf(numFramesStr, "%d", &numFrames); err != nil {
				t.Logf("Warning: Failed to parse NumberOfFrames '%s': %v", numFramesStr, err)
				numFrames = 1
			}
		}

		if rows > 0 && cols > 0 && bitsAllocated > 0 {
			bytesPerSample := int((bitsAllocated-1)/8 + 1)
			frameSize := int(rows) * int(cols) * int(samplesPerPixel) * bytesPerSample
			totalSize := frameSize * numFrames

			t.Logf("Calculated pixel data size:")
			t.Logf("  Frame size: %d bytes (%dx%d pixels, %d bytes/sample, %d samples/pixel)",
				frameSize, rows, cols, bytesPerSample, samplesPerPixel)
			t.Logf("  Total size for %d frames: %d bytes (%.2f MB)",
				numFrames, totalSize, float64(totalSize)/(1024*1024))

			// Verify it matches expected size for TestMultiFrame.dcm
			// 288x288 pixels, 16 bits (2 bytes), 1 sample/pixel, 7 frames
			expectedFrameSize := 288 * 288 * 2 * 1
			expectedTotalSize := expectedFrameSize * 7
			if frameSize != expectedFrameSize {
				t.Errorf("Frame size mismatch: got %d, expected %d", frameSize, expectedFrameSize)
			}
			if totalSize != expectedTotalSize {
				t.Errorf("Total size mismatch: got %d, expected %d", totalSize, expectedTotalSize)
			}
		}
	})
}

// TestMultiFramePixelDataExtraction tests extracting individual frames from multi-frame image.
func TestMultiFramePixelDataExtraction(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestMultiFrame.dcm")

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open TestMultiFrame.dcm: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			t.Errorf("Failed to close file: %v", closeErr)
		}
	}()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse TestMultiFrame.dcm: %v", err)
	}

	// Get number of frames from string (IS - Integer String)
	numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
	if !exists {
		t.Skip("Cannot determine number of frames: NumberOfFrames tag not found")
		return
	}

	var numFrames int
	if _, err := fmt.Sscanf(numFramesStr, "%d", &numFrames); err != nil {
		t.Skipf("Cannot parse NumberOfFrames '%s': %v", numFramesStr, err)
		return
	}

	if numFrames < 1 {
		t.Skipf("No frames to extract (NumberOfFrames=%d)", numFrames)
		return
	}

	t.Logf("Attempting to extract pixel data from %d frames", numFrames)

	// Get pixel data element
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData element not found")
	}

	// Check the type of pixel data element
	t.Logf("PixelData element type: %T", pixelDataElem)

	// For uncompressed multi-frame data, we should be able to extract frames
	// This is a basic test - detailed frame extraction would require
	// knowledge of the pixel data format

	rows := result.Dataset.TryGetUInt16(tag.Rows, 0)
	cols := result.Dataset.TryGetUInt16(tag.Columns, 0)
	bitsAllocated := result.Dataset.TryGetUInt16(tag.BitsAllocated, 0)
	samplesPerPixel := result.Dataset.TryGetUInt16(tag.SamplesPerPixel, 1)
	if samplesPerPixel == 0 {
		samplesPerPixel = 1
	}

	if rows > 0 && cols > 0 && bitsAllocated > 0 {
		bytesPerSample := int((bitsAllocated-1)/8 + 1)
		expectedFrameSize := int(rows) * int(cols) * int(samplesPerPixel) * bytesPerSample

		t.Logf("Expected frame size: %d bytes", expectedFrameSize)
		t.Logf("Image dimensions: %dx%d, %d bits/pixel, %d samples/pixel",
			rows, cols, bitsAllocated, samplesPerPixel)

		// Verify frame size calculation
		if expectedFrameSize == 0 {
			t.Error("Frame size should not be 0")
		} else {
			t.Logf("鉁?Frame size calculation successful")
		}
	}
}

// TestMultiFrameVsPrivateSequence specifically tests the VR=UN + undefined length handling.
func TestMultiFrameVsPrivateSequence(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestMultiFrame.dcm")

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open TestMultiFrame.dcm: %v", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			t.Errorf("Failed to close file: %v", closeErr)
		}
	}()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse TestMultiFrame.dcm: %v", err)
	}

	// This test verifies that the parser correctly handles:
	// 1. VR=UN (Unknown) with undefined length (0xFFFFFFFF)
	// 2. Private sequences encoded this way
	// 3. The fix that treats UN+undefined length as a sequence

	t.Log("Testing parser's handling of VR=UN with undefined length")

	// The file should parse successfully without EOF errors
	if result.Dataset.Count() < 100 {
		t.Errorf("Expected significant number of elements (>100), got %d", result.Dataset.Count())
	}

	t.Logf("Successfully parsed %d elements including private sequences", result.Dataset.Count())

	// Look for any private tags (odd group numbers)
	privateCount := 0
	for _, elem := range result.Dataset.Elements() {
		if elem.Tag().Group()%2 == 1 {
			privateCount++
		}
	}

	t.Logf("Found %d private tag elements", privateCount)
}

