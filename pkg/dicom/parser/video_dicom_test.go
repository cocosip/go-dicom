// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
    "os"
    "path/filepath"
    "strings"
    "testing"

    "github.com/cocosip/go-dicom/pkg/dicom/dataset"
    "github.com/cocosip/go-dicom/pkg/dicom/element"
    "github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestETIAMVideo tests the ETIAM_video_002.dcm video DICOM file
func TestETIAMVideo(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
    filePath := filepath.Join(testDataDir, "ETIAM_video_002.dcm")
    filePath = filepath.Clean(filePath)
    if !strings.HasPrefix(filePath, filepath.Clean(testDataDir)+string(os.PathSeparator)) && filePath != filepath.Clean(testDataDir) {
        t.Skipf("Invalid test file path: %s", filePath)
        return
    }

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
	t.Logf("File size: %.2f MB", float64(fileInfo.Size())/(1024*1024))

	// Open and parse the file
    file, err := os.Open(filePath)
    if err != nil {
        t.Fatalf("Failed to open file: %v", err)
    }
    defer func() { _ = file.Close() }()

	t.Log("Parsing ETIAM_video_002.dcm...")
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse ETIAM_video_002.dcm: %v", err)
	}

	if result == nil || result.Dataset == nil {
		t.Fatal("Parse result or dataset is nil")
	}

	t.Logf("✓ Successfully parsed ETIAM_video_002.dcm")
	t.Logf("Dataset contains %d elements", result.Dataset.Count())

	// Check various video DICOM attributes
	checkSOPClassUID(t, result.Dataset)
	checkModality(t, result.Dataset)
	checkVideoTransferSyntax(t, result.FileMetaInformation)
	checkVideoPixelData(t, result.Dataset)
	checkVideoAttributes(t, result.Dataset)
	checkImageDimensions(t, result.Dataset)
	checkPatientStudyInfo(t, result.Dataset)

	t.Log("\n✓ Video DICOM file parsed successfully")
	t.Log("✓ Pixel data (video stream) is accessible")
}

// TestVideoDICOMPixelDataAccess specifically tests pixel data access
func TestVideoDICOMPixelDataAccess(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
    filePath := filepath.Join(testDataDir, "ETIAM_video_002.dcm")
    filePath = filepath.Clean(filePath)
    if !strings.HasPrefix(filePath, filepath.Clean(testDataDir)+string(os.PathSeparator)) && filePath != filepath.Clean(testDataDir) {
        t.Skipf("Invalid test file path: %s", filePath)
        return
    }

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", filePath)
		return
	}

    file, err := os.Open(filePath)
    if err != nil {
        t.Fatalf("Failed to open file: %v", err)
    }
    defer func() { _ = file.Close() }()

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Test 1: Pixel Data must exist
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("CRITICAL: PixelData element not found - video data is not accessible")
	}
	t.Log("✓ Test 1 PASS: PixelData element exists")

	// Test 2: Pixel Data must have a valid type
	if pixelDataElem == nil {
		t.Fatal("CRITICAL: PixelData element is nil")
	}
	t.Logf("✓ Test 2 PASS: PixelData has type %T", pixelDataElem)

	// Test 3: Check if it's a fragment sequence (typical for compressed video)
	pixelDataStr := pixelDataElem.String()
	t.Logf("PixelData structure: %s", pixelDataStr)

	// Test 4: For fragment sequences (typical for compressed video), access fragments
	if fragSeq, ok := pixelDataElem.(*element.OtherByteFragment); ok {
		fragmentCount := fragSeq.FragmentCount()
		t.Logf("✓ Test 4 PASS: Fragment sequence with %d fragments", fragmentCount)

		if fragmentCount == 0 {
			t.Fatal("CRITICAL: Fragment sequence has no fragments")
		}

		// Access the fragments
		fragments := fragSeq.Fragments()
		if fragments == nil {
			t.Fatal("CRITICAL: Fragments() returned nil")
		}
		t.Logf("✓ Test 5 PASS: Can access %d video fragments", len(fragments))

		// Verify we can access individual fragments
		if fragmentCount > 0 {
			firstFragment, err := fragSeq.GetFragment(0)
			if err != nil {
				t.Fatalf("CRITICAL: Cannot access first fragment: %v", err)
			}
			if firstFragment == nil {
				t.Fatal("CRITICAL: First fragment is nil")
			}
			t.Logf("✓ Test 6 PASS: Can access individual fragments")
		}
	} else if fragSeq, ok := pixelDataElem.(*element.OtherWordFragment); ok {
		// OW fragment sequence
		fragmentCount := fragSeq.FragmentCount()
		t.Logf("✓ Test 4 PASS: OW Fragment sequence with %d fragments", fragmentCount)

		if fragmentCount == 0 {
			t.Fatal("CRITICAL: Fragment sequence has no fragments")
		}
		t.Logf("✓ Test 5 PASS: Can access %d video fragments", len(fragSeq.Fragments()))
	} else {
		// Not a fragment sequence, check regular buffer
		buffer := pixelDataElem.Buffer()
		if buffer == nil {
			t.Fatal("CRITICAL: PixelData buffer is nil - cannot access video data")
		}
		t.Log("✓ Test 4 PASS: PixelData buffer is accessible")
		t.Logf("✓ Test 5 PASS: Buffer type: %T", buffer)
	}

	t.Log("\n✓ ALL TESTS PASSED: Video pixel data is fully accessible")
}

// TestVideoDICOMFormats tests various video DICOM formats if available
func TestVideoDICOMFormats(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	// List of potential video DICOM files
	videoFiles := []string{
		"ETIAM_video_002.dcm",
	}

	successCount := 0
	for _, filename := range videoFiles {
		t.Run(filename, func(t *testing.T) {
            filePath := filepath.Join(testDataDir, filename)
            filePath = filepath.Clean(filePath)
            if !strings.HasPrefix(filePath, filepath.Clean(testDataDir)+string(os.PathSeparator)) && filePath != filepath.Clean(testDataDir) {
                t.Skipf("Invalid test file path: %s", filePath)
                return
            }

			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Skipf("Test file not found: %s", filePath)
				return
			}

            file, err := os.Open(filePath)
            if err != nil {
                t.Errorf("Failed to open %s: %v", filename, err)
                return
            }
            defer func() { _ = file.Close() }()

			result, err := Parse(file)
			if err != nil {
				t.Errorf("Failed to parse %s: %v", filename, err)
				return
			}

			// Check for pixel data
			if _, exists := result.Dataset.Get(tag.PixelData); !exists {
				t.Errorf("PixelData not found in %s", filename)
				return
			}

			t.Logf("✓ Successfully parsed video DICOM: %s", filename)
			successCount++
		})
	}

	if successCount == 0 {
		t.Skip("No video DICOM files were tested")
	}
}

// BenchmarkVideoDICOMParsing benchmarks parsing of video DICOM files
func BenchmarkVideoDICOMParsing(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "ETIAM_video_002.dcm")

	// Read file content once
    filePath = filepath.Clean(filePath)
    if !strings.HasPrefix(filePath, filepath.Clean(testDataDir)+string(os.PathSeparator)) && filePath != filepath.Clean(testDataDir) {
        b.Skipf("Invalid test file path: %s", filePath)
        return
    }
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

// checkSOPClassUID checks and logs the SOP Class UID
func checkSOPClassUID(t *testing.T, ds *dataset.Dataset) {
	if sopClassUID, exists := ds.GetString(tag.SOPClassUID); exists {
		t.Logf("SOP Class UID: %s", sopClassUID)
		if sopClassUID == "1.2.840.10008.5.1.4.1.1.77.1.1.1" {
			t.Log("✓ Confirmed: Video Endoscopic Image Storage")
		}
	} else {
		t.Error("SOP Class UID not found")
	}
}

// checkModality checks and logs the modality
func checkModality(t *testing.T, ds *dataset.Dataset) {
	if modality, exists := ds.GetString(tag.Modality); exists {
		t.Logf("Modality: %s", modality)
	}
}

// checkVideoTransferSyntax checks and logs the transfer syntax
func checkVideoTransferSyntax(t *testing.T, fmi *dataset.FileMetaInformation) {
	if fmi == nil {
		return
	}

	if tsUID, exists := fmi.TransferSyntaxUID(); exists {
		t.Logf("Transfer Syntax UID: %s", tsUID)

		switch tsUID {
		case "1.2.840.10008.1.2.4.100":
			t.Log("✓ Video format: MPEG2 Main Profile @ Main Level")
		case "1.2.840.10008.1.2.4.101":
			t.Log("✓ Video format: MPEG2 Main Profile @ High Level")
		case "1.2.840.10008.1.2.4.102":
			t.Log("✓ Video format: MPEG4 AVC/H.264 High Profile")
		case "1.2.840.10008.1.2.4.103":
			t.Log("✓ Video format: MPEG4 AVC/H.264 BD-compatible")
		default:
			t.Logf("Video format: Other (%s)", tsUID)
		}
	}
}

// checkVideoPixelData checks and logs pixel data information
func checkVideoPixelData(t *testing.T, ds *dataset.Dataset) {
	pixelDataElem, exists := ds.Get(tag.PixelData)
	if !exists {
		t.Fatal("❌ PixelData tag not found - cannot get video data")
	}

	t.Logf("✓ PixelData element found")
	t.Logf("  PixelData tag: %s", pixelDataElem.Tag())
	t.Logf("  PixelData VR: %s", pixelDataElem.ValueRepresentation())
	t.Logf("  PixelData type: %T", pixelDataElem)

	if pixelDataElem.Length() > 0 {
		t.Logf("  PixelData length: %.2f MB", float64(pixelDataElem.Length())/(1024*1024))
	} else {
		t.Logf("  PixelData length: undefined (fragment sequence)")
	}
}

// checkVideoAttributes checks and logs video-specific attributes
func checkVideoAttributes(t *testing.T, ds *dataset.Dataset) {
	t.Log("\nVideo-specific attributes:")

	if numFrames, exists := ds.GetString(tag.NumberOfFrames); exists {
		t.Logf("  Number of Frames: %s", numFrames)
	} else {
		t.Log("  Number of Frames: not specified (continuous video)")
	}

	if frameTime, exists := ds.GetString(tag.FrameTime); exists {
		t.Logf("  Frame Time: %s ms", frameTime)
	}

	if frameRate, exists := ds.GetString(tag.RecommendedDisplayFrameRate); exists {
		t.Logf("  Recommended Display Frame Rate: %s fps", frameRate)
	}

	if cineRate, exists := ds.GetString(tag.CineRate); exists {
		t.Logf("  Cine Rate: %s", cineRate)
	}
}

// checkImageDimensions checks and logs image dimensions
func checkImageDimensions(t *testing.T, ds *dataset.Dataset) {
	if rows, err := ds.GetUInt16(tag.Rows, 0); err == nil {
		t.Logf("  Rows: %d", rows)
	}

	if cols, err := ds.GetUInt16(tag.Columns, 0); err == nil {
		t.Logf("  Columns: %d", cols)
	}

	if samplesPerPixel, err := ds.GetUInt16(tag.SamplesPerPixel, 0); err == nil {
		t.Logf("  Samples Per Pixel: %d", samplesPerPixel)
		if samplesPerPixel == 3 {
			t.Log("  ✓ Color video (RGB or YBR)")
		}
	}

	if photoInterp, exists := ds.GetString(tag.PhotometricInterpretation); exists {
		t.Logf("  Photometric Interpretation: %s", photoInterp)
	}

	if bitsAlloc, err := ds.GetUInt16(tag.BitsAllocated, 0); err == nil {
		t.Logf("  Bits Allocated: %d", bitsAlloc)
	}
}

// checkPatientStudyInfo checks and logs patient and study information
func checkPatientStudyInfo(t *testing.T, ds *dataset.Dataset) {
	t.Log("\nPatient/Study information:")

	if patientName, exists := ds.GetString(tag.PatientName); exists {
		t.Logf("  Patient Name: %s", patientName)
	}

	if studyDesc, exists := ds.GetString(tag.StudyDescription); exists {
		t.Logf("  Study Description: %s", studyDesc)
	}

	if seriesDesc, exists := ds.GetString(tag.SeriesDescription); exists {
		t.Logf("  Series Description: %s", seriesDesc)
	}
}
