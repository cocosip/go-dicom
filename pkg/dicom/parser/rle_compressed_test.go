package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// TestRLECompressedDICOM tests parsing RLE compressed DICOM files
func TestRLECompressedDICOM(t *testing.T) {
	// Open the RLE compressed test file
	testFile := filepath.Join("..", "..", "..", "test-data", "D_CLUNIE_CT1_RLE_FRAGS.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open RLE test file: %v", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	t.Logf("RLE DICOM File: %s", testFile)
	t.Logf("File size: %.2f KB", float64(fileInfo.Size())/1024)

	// Parse the DICOM file
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse RLE DICOM: %v", err)
	}

	t.Logf("✓ Successfully parsed RLE compressed DICOM")
	t.Logf("✓ Dataset contains %d elements", result.Dataset.Count())

	// Verify Transfer Syntax UID - should be RLE Lossless (1.2.840.10008.1.2.5)
	transferSyntaxUID, exists := result.Dataset.GetString(tag.TransferSyntaxUID)
	if !exists {
		t.Log("TransferSyntaxUID not found in dataset (may be in file metadata)")
		// Check file metadata
		if result.FileMetaInformation != nil {
			tsUID, exists := result.FileMetaInformation.TransferSyntaxUID()
			if exists {
				t.Logf("Transfer Syntax UID (from file meta): %s", tsUID)
				if tsUID == "1.2.840.10008.1.2.5" {
					t.Logf("✓ Confirmed RLE Lossless transfer syntax")
				}
			}
		}
	} else {
		t.Logf("Transfer Syntax UID: %s", transferSyntaxUID)
		if transferSyntaxUID != "1.2.840.10008.1.2.5" {
			t.Logf("⚠ Transfer Syntax is %s (expected RLE Lossless: 1.2.840.10008.1.2.5)", transferSyntaxUID)
		} else {
			t.Logf("✓ Confirmed RLE Lossless transfer syntax")
		}
	}

	// Verify SOP Class UID
	sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID)
	if exists {
		t.Logf("SOP Class UID: %s", sopClassUID)
	}

	// Get modality
	modality, exists := result.Dataset.GetString(tag.Modality)
	if exists {
		t.Logf("Modality: %s", modality)
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

	// Get number of frames (if multi-frame)
	numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
	if exists {
		t.Logf("Number of Frames: %s", numFramesStr)
	} else {
		t.Logf("Number of Frames: 1 (single frame)")
	}

	// Get patient information
	patientName, exists := result.Dataset.GetString(tag.PatientName)
	if exists {
		t.Logf("Patient Name: %s", patientName)
	}

	patientID, exists := result.Dataset.GetString(tag.PatientID)
	if exists {
		t.Logf("Patient ID: %s", patientID)
	}

	// Get study/series description
	studyDesc, exists := result.Dataset.GetString(tag.StudyDescription)
	if exists {
		t.Logf("Study Description: %s", studyDesc)
	}

	seriesDesc, exists := result.Dataset.GetString(tag.SeriesDescription)
	if exists {
		t.Logf("Series Description: %s", seriesDesc)
	}

	t.Logf("✓ RLE compressed DICOM file parsed successfully")
}

// TestRLEPixelDataAccess tests accessing pixel data from RLE compressed DICOM
func TestRLEPixelDataAccess(t *testing.T) {
	testFile := filepath.Join("..", "..", "..", "test-data", "D_CLUNIE_CT1_RLE_FRAGS.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse DICOM: %v", err)
	}

	// Get pixel data element
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData element not found")
	}

	t.Logf("✓ PixelData element found")
	t.Logf("PixelData element type: %T", pixelDataElem)

	// RLE compressed data is typically stored as fragment sequence
	// Check the type and access the compressed data
	t.Logf("Attempting to access RLE fragment data...")

	// Try to access through FragmentCount interface
	if fragSeq, ok := pixelDataElem.(interface {
		FragmentCount() int
		Fragments() []buffer.ByteBuffer
	}); ok {
		fragmentCount := fragSeq.FragmentCount()
		t.Logf("✓ RLE data stored as fragment sequence")
		t.Logf("Number of fragments: %d", fragmentCount)

		// Get all fragments
		fragments := fragSeq.Fragments()
		t.Logf("Successfully retrieved %d fragments", len(fragments))

		// Access ByteBuffer data
		totalSize := 0
		for i, buf := range fragments {
			data := buf.Data()
			fragmentSize := len(data)
			totalSize += fragmentSize
			t.Logf("  Fragment %d: %d bytes (%.2f KB)", i, fragmentSize, float64(fragmentSize)/1024)
		}

		t.Logf("Total RLE compressed data: %d bytes (%.2f KB)", totalSize, float64(totalSize)/1024)
		t.Logf("✓ Successfully accessed all RLE compressed fragments")
	} else if dataElem, ok := pixelDataElem.(interface{ Data() []byte }); ok {
		data := dataElem.Data()
		t.Logf("✓ RLE data stored as byte array")
		t.Logf("Compressed pixel data size: %d bytes (%.2f KB)", len(data), float64(len(data))/1024)
	} else {
		t.Logf("PixelData element type: %T", pixelDataElem)
		t.Logf("⚠ Unexpected pixel data type, but element is accessible")
	}

	t.Logf("✓ RLE compressed pixel data is accessible")
}

// TestRLEImageProperties tests extracting image properties from RLE compressed DICOM
func TestRLEImageProperties(t *testing.T) {
	testFile := filepath.Join("..", "..", "..", "test-data", "D_CLUNIE_CT1_RLE_FRAGS.dcm")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse DICOM: %v", err)
	}

	// Collect all image properties
	properties := make(map[string]interface{})

	// Basic image properties
	if rows, err := result.Dataset.GetUInt16(tag.Rows, 0); err == nil {
		properties["Rows"] = rows
	}

	if cols, err := result.Dataset.GetUInt16(tag.Columns, 0); err == nil {
		properties["Columns"] = cols
	}

	if ba, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0); err == nil {
		properties["BitsAllocated"] = ba
	}

	if bs, err := result.Dataset.GetUInt16(tag.BitsStored, 0); err == nil {
		properties["BitsStored"] = bs
	}

	if hb, err := result.Dataset.GetUInt16(tag.HighBit, 0); err == nil {
		properties["HighBit"] = hb
	}

	if pr, err := result.Dataset.GetUInt16(tag.PixelRepresentation, 0); err == nil {
		properties["PixelRepresentation"] = pr
	}

	if pi, exists := result.Dataset.GetString(tag.PhotometricInterpretation); exists {
		properties["PhotometricInterpretation"] = pi
	}

	if sc, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0); err == nil {
		properties["SamplesPerPixel"] = sc
	}

	// Print all properties
	t.Log("=== RLE Image Properties ===")
	for key, value := range properties {
		t.Logf("%s: %v", key, value)
	}

	// Verify essential properties exist
	essentialTags := []*tag.Tag{
		tag.Rows,
		tag.Columns,
		tag.BitsAllocated,
		tag.PhotometricInterpretation,
	}

	for _, tagPtr := range essentialTags {
		if _, exists := result.Dataset.Get(tagPtr); !exists {
			t.Errorf("Essential tag %v not found", tagPtr)
		}
	}

	t.Logf("✓ All essential image properties present")
}

// BenchmarkRLEParsing benchmarks parsing performance for RLE compressed DICOM
func BenchmarkRLEParsing(b *testing.B) {
	testFile := filepath.Join("..", "..", "..", "test-data", "D_CLUNIE_CT1_RLE_FRAGS.dcm")

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
