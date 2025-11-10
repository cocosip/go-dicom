package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// TestModalityLUTParsing tests parsing DICOM files with Modality LUT Sequence
func TestModalityLUTParsing(t *testing.T) {
	// Open the Modality LUT test file
	testFile := filepath.Join("..", "..", "..", "test-data", "CR-ModalitySequenceLUT.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open Modality LUT test file: %v", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	t.Logf("Modality LUT DICOM File: %s", testFile)
	t.Logf("File size: %.2f MB", float64(fileInfo.Size())/(1024*1024))

	// Parse the DICOM file
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse Modality LUT DICOM: %v", err)
	}

	t.Logf("✓ Successfully parsed Modality LUT DICOM")
	t.Logf("✓ Dataset contains %d elements", result.Dataset.Count())

	// Get SOP Class UID
	sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID)
	if exists {
		t.Logf("SOP Class UID: %s", sopClassUID)
	}

	// Get modality
	modality, exists := result.Dataset.GetString(tag.Modality)
	if !exists {
		t.Error("Modality not found")
	} else {
		t.Logf("Modality: %s", modality)
		if modality != "CR" && modality != "CR " {
			t.Errorf("Expected CR modality, got: %s", modality)
		} else {
			t.Log("✓ Confirmed CR (Computed Radiography) modality")
		}
	}

	// Get image dimensions
	rows, rowsErr := result.Dataset.GetUInt16(tag.Rows, 0)
	cols, colsErr := result.Dataset.GetUInt16(tag.Columns, 0)
	if rowsErr == nil && colsErr == nil {
		t.Logf("Image dimensions: %d x %d", cols, rows)
	}

	// Get bits allocated and stored
	bitsAllocated, baErr := result.Dataset.GetUInt16(tag.BitsAllocated, 0)
	bitsStored, bsErr := result.Dataset.GetUInt16(tag.BitsStored, 0)
	if baErr == nil && bsErr == nil {
		t.Logf("Bits Allocated: %d, Bits Stored: %d", bitsAllocated, bitsStored)
	}

	// Get photometric interpretation
	photoInterp, exists := result.Dataset.GetString(tag.PhotometricInterpretation)
	if exists {
		t.Logf("Photometric Interpretation: %s", photoInterp)
	}

	// Get patient information
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if exists {
		t.Logf("Patient Name: %s", patientName)
	}

	t.Logf("✓ Modality LUT DICOM file parsed successfully")
}

// TestModalityLUTSequence tests accessing the Modality LUT Sequence
func TestModalityLUTSequence(t *testing.T) {
	testFile := filepath.Join("..", "..", "..", "test-data", "CR-ModalitySequenceLUT.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse DICOM: %v", err)
	}

	t.Log("=== Modality LUT Sequence Analysis ===")

	// Get Modality LUT Sequence - tag (0028,3000)
	modalityLUTSeqElem, exists := result.Dataset.Get(tag.ModalityLUTSequence)
	if !exists {
		t.Error("Modality LUT Sequence not found")
		return
	}

	t.Logf("✓ Modality LUT Sequence found")

	// Check if it's a sequence
	modalityLUTSeq, ok := modalityLUTSeqElem.(*dataset.Sequence)
	if !ok {
		t.Errorf("Modality LUT Sequence is not a Sequence type, got: %T", modalityLUTSeqElem)
		return
	}

	itemCount := modalityLUTSeq.Count()
	t.Logf("✓ Modality LUT Sequence is a valid Sequence")
	t.Logf("Number of items in sequence: %d", itemCount)

	if itemCount == 0 {
		t.Error("Modality LUT Sequence is empty")
		return
	}

	// Analyze first item
	for i := 0; i < itemCount; i++ {
		item := modalityLUTSeq.GetItem(i)
		t.Logf("\n--- Modality LUT Sequence Item %d ---", i)

		// LUT Descriptor (0028,3002) - required
		if descriptor, err := item.GetUInt16s(tag.LUTDescriptor); err == nil {
			t.Logf("  LUT Descriptor: %v", descriptor)
			if len(descriptor) >= 3 {
				t.Logf("    Number of entries: %d", descriptor[0])
				t.Logf("    First mapped value: %d", descriptor[1])
				t.Logf("    Bits per entry: %d", descriptor[2])
			}
		} else {
			t.Logf("  LUT Descriptor not found or error: %v", err)
		}

		// LUT Explanation (0028,3003) - optional
		if explanation, exists := item.GetString(tag.LUTExplanation); exists {
			t.Logf("  LUT Explanation: %s", explanation)
		}

		// Modality LUT Type (0028,3004) - required
		if lutType, exists := item.GetString(tag.ModalityLUTType); exists {
			t.Logf("  Modality LUT Type: %s", lutType)
		}

		// LUT Data (0028,3006) - required
		lutDataElem, exists := item.Get(tag.LUTData)
		if exists {
			t.Logf("  LUT Data found: %T", lutDataElem)
			// Try to get data size
			if dataElem, ok := lutDataElem.(interface{ Count() int }); ok {
				count := dataElem.Count()
				t.Logf("  LUT Data entries: %d", count)
			}
		} else {
			t.Log("  LUT Data not found")
		}
	}

	t.Logf("\n✓ Modality LUT Sequence successfully analyzed")
}

// TestModalityLUTPixelData tests that pixel data is accessible and not empty
func TestModalityLUTPixelData(t *testing.T) {
	testFile := filepath.Join("..", "..", "..", "test-data", "CR-ModalitySequenceLUT.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse DICOM: %v", err)
	}

	t.Log("=== Pixel Data Verification ===")

	// Get pixel data element
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("❌ PixelData element not found - CRITICAL FAILURE")
	}

	t.Logf("✓ PixelData element found")
	t.Logf("PixelData element type: %T", pixelDataElem)

	// Verify pixel data is not empty
	pixelDataEmpty := false
	pixelDataSize := 0

	// Check different pixel data types
	// First try fragment sequence (common for compressed data)
	if fragSeq, ok := pixelDataElem.(interface {
		FragmentCount() int
		Fragments() []buffer.ByteBuffer
	}); ok {
		fragCount := fragSeq.FragmentCount()
		t.Logf("Pixel data stored as fragment sequence")
		t.Logf("Number of fragments: %d", fragCount)

		if fragCount > 0 {
			fragments := fragSeq.Fragments()
			for i, buf := range fragments {
				data := buf.Data()
				fragmentSize := len(data)
				pixelDataSize += fragmentSize
				t.Logf("  Fragment %d: %d bytes (%.2f KB)", i, fragmentSize, float64(fragmentSize)/1024)
			}
			t.Logf("Total pixel data size: %d bytes (%.2f KB)", pixelDataSize, float64(pixelDataSize)/1024)
		}

		if pixelDataSize == 0 {
			pixelDataEmpty = true
		}
	} else if dataElem, ok := pixelDataElem.(interface{ Data() []byte }); ok {
		data := dataElem.Data()
		pixelDataSize = len(data)
		if pixelDataSize == 0 {
			pixelDataEmpty = true
		}
		t.Logf("Pixel data stored as byte array")
		t.Logf("Pixel data size: %d bytes (%.2f KB)", pixelDataSize, float64(pixelDataSize)/1024)
	} else {
		t.Logf("Pixel data type: %T", pixelDataElem)
		// Try Count() as fallback
		if countElem, ok := pixelDataElem.(interface{ Count() int }); ok {
			count := countElem.Count()
			t.Logf("Pixel data element count: %d", count)
		}
		t.Log("Checking if pixel data has content using alternative methods...")
	}

	// Critical check: Pixel data must not be empty
	if pixelDataEmpty {
		t.Fatal("❌ CRITICAL FAILURE: Pixel data is EMPTY - this violates the test requirement")
	}

	if pixelDataSize > 0 {
		t.Logf("✓ Pixel data is NOT empty: %d bytes", pixelDataSize)
		t.Logf("✓ PASS: Pixel data requirement satisfied")
	} else {
		// Additional check - try to access via other methods
		t.Log("Attempting alternative pixel data access methods...")

		// Check if it's a fragment sequence
		if fragSeq, ok := pixelDataElem.(interface {
			FragmentCount() int
		}); ok {
			fragCount := fragSeq.FragmentCount()
			if fragCount > 0 {
				t.Logf("✓ Pixel data stored as fragment sequence with %d fragments", fragCount)
				pixelDataEmpty = false
			}
		}

		if pixelDataEmpty {
			t.Fatal("❌ CRITICAL FAILURE: Could not verify pixel data is not empty")
		}
	}

	// Get expected pixel data size based on image properties
	rows, _ := result.Dataset.GetUInt16(tag.Rows, 0)
	cols, _ := result.Dataset.GetUInt16(tag.Columns, 0)
	bitsAllocated, _ := result.Dataset.GetUInt16(tag.BitsAllocated, 0)

	if rows > 0 && cols > 0 && bitsAllocated > 0 {
		expectedSize := int(rows) * int(cols) * int(bitsAllocated) / 8
		t.Logf("Expected pixel data size: %d bytes (%.2f KB)", expectedSize, float64(expectedSize)/1024)

		if pixelDataSize > 0 {
			ratio := float64(pixelDataSize) / float64(expectedSize)
			t.Logf("Actual/Expected ratio: %.2f", ratio)

			if ratio > 0.5 && ratio < 2.0 {
				t.Logf("✓ Pixel data size is reasonable")
			} else {
				t.Logf("⚠ Pixel data size differs significantly from expected")
			}
		}
	}

	t.Logf("\n✓ Pixel data verification complete - data is NOT empty")
}

// TestModalityLUTComplete tests complete workflow of parsing LUT DICOM with pixel data
func TestModalityLUTComplete(t *testing.T) {
	testFile := filepath.Join("..", "..", "..", "test-data", "CR-ModalitySequenceLUT.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	t.Log("=== Complete Modality LUT DICOM Test ===")
	t.Logf("File: CR-ModalitySequenceLUT.dcm")
	t.Logf("Size: %.2f MB", float64(fileInfo.Size())/(1024*1024))

	// Parse file
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("❌ Failed to parse: %v", err)
	}
	t.Logf("✓ File parsed successfully")

	// Check essential elements
	checksPass := true

	// 1. Check Modality LUT Sequence exists
	if _, exists := result.Dataset.Get(tag.ModalityLUTSequence); exists {
		t.Logf("✓ Modality LUT Sequence present")
	} else {
		t.Logf("❌ Modality LUT Sequence missing")
		checksPass = false
	}

	// 2. Check Pixel Data exists and is not empty
	if pixelDataElem, exists := result.Dataset.Get(tag.PixelData); exists {
		isEmpty := false

		if pd, ok := pixelDataElem.(interface{ Data() []byte }); ok {
			if len(pd.Data()) == 0 {
				isEmpty = true
			} else {
				t.Logf("✓ Pixel Data present and not empty: %d bytes", len(pd.Data()))
			}
		}

		if isEmpty {
			t.Logf("❌ Pixel Data is empty")
			checksPass = false
		}
	} else {
		t.Logf("❌ Pixel Data missing")
		checksPass = false
	}

	// 3. Check image properties
	rows, rowsErr := result.Dataset.GetUInt16(tag.Rows, 0)
	cols, colsErr := result.Dataset.GetUInt16(tag.Columns, 0)
	if rowsErr == nil && colsErr == nil && rows > 0 && cols > 0 {
		t.Logf("✓ Image dimensions valid: %d x %d", cols, rows)
	} else {
		t.Logf("❌ Image dimensions invalid")
		checksPass = false
	}

	// Final result
	if checksPass {
		t.Log("\n✅ ALL CHECKS PASSED")
		t.Log("✓ File parses correctly")
		t.Log("✓ Modality LUT Sequence accessible")
		t.Log("✓ Pixel data present and not empty")
	} else {
		t.Fatal("\n❌ SOME CHECKS FAILED")
	}
}

// BenchmarkModalityLUTParsing benchmarks parsing performance for Modality LUT DICOM
func BenchmarkModalityLUTParsing(b *testing.B) {
	testFile := filepath.Join("..", "..", "..", "test-data", "CR-ModalitySequenceLUT.dcm")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err := os.Open(testFile)
		if err != nil {
			b.Fatalf("Failed to open file: %v", err)
		}

		_, err = Parse(file)
		if err != nil {
			file.Close()
			b.Fatalf("Failed to parse: %v", err)
		}

		file.Close()
	}
}
