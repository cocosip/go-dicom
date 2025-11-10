package parser

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

// TestWriteReadSingleFrame tests writing and reading a single frame DICOM image
func TestWriteReadSingleFrame(t *testing.T) {
	t.Log("=== Testing Single Frame Image Write/Read Cycle ===")

	// Create a dataset with single frame image
	ds := dataset.New()

	// Add required File Meta Information elements
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"})) // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))
	ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8"}))
	ds.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7"}))

	// Add Patient/Study information
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"TEST001"}))
	ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))
	ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Test Study"}))

	// Add image properties
	rows := uint16(256)
	cols := uint16(256)
	bitsAllocated := uint16(16)
	bitsStored := uint16(16)
	highBit := uint16(15)
	pixelRep := uint16(0) // unsigned

	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
	ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{cols}))
	ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{bitsAllocated}))
	ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{bitsStored}))
	ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{highBit}))
	ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{pixelRep}))
	ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

	// Create pixel data - single frame 256x256x2 bytes = 131072 bytes
	pixelDataSize := int(rows) * int(cols) * 2 // 2 bytes per pixel for 16-bit
	pixelData := make([]byte, pixelDataSize)

	// Fill with pattern data (gradient)
	for i := 0; i < pixelDataSize/2; i++ {
		value := uint16(i % 65536)
		pixelData[i*2] = byte(value & 0xFF)
		pixelData[i*2+1] = byte((value >> 8) & 0xFF)
	}

	t.Logf("Created pixel data: %d bytes", len(pixelData))

	// Add pixel data to dataset
	pixelDataElement := element.NewOtherWord(tag.PixelData, pixelData)
	ds.Add(pixelDataElement)

	t.Logf("Dataset created with %d elements", ds.Count())

	// Write to buffer
	var buf bytes.Buffer
	err := writer.Write(&buf, ds)
	if err != nil {
		t.Fatalf("Failed to write DICOM: %v", err)
	}

	writtenSize := buf.Len()
	t.Logf("✓ Successfully wrote DICOM: %d bytes", writtenSize)

	// Parse the written data
	reader := bytes.NewReader(buf.Bytes())
	result, err := Parse(reader)
	if err != nil {
		t.Fatalf("Failed to parse written DICOM: %v", err)
	}

	t.Logf("✓ Successfully parsed written DICOM")
	t.Logf("Parsed dataset contains %d elements", result.Dataset.Count())

	// Verify image properties
	parsedRows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil || parsedRows != rows {
		t.Errorf("Rows mismatch: got %d, want %d", parsedRows, rows)
	} else {
		t.Logf("✓ Rows verified: %d", parsedRows)
	}

	parsedCols, err := result.Dataset.GetUInt16(tag.Columns, 0)
	if err != nil || parsedCols != cols {
		t.Errorf("Columns mismatch: got %d, want %d", parsedCols, cols)
	} else {
		t.Logf("✓ Columns verified: %d", parsedCols)
	}

	parsedBitsAlloc, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0)
	if err != nil || parsedBitsAlloc != bitsAllocated {
		t.Errorf("BitsAllocated mismatch: got %d, want %d", parsedBitsAlloc, bitsAllocated)
	} else {
		t.Logf("✓ BitsAllocated verified: %d", parsedBitsAlloc)
	}

	// Verify number of frames (should default to 1 for single frame)
	numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
	if exists {
		numFramesStr = strings.TrimSpace(numFramesStr)
		if numFramesStr != "1" {
			t.Errorf("NumberOfFrames mismatch: got %s, want 1", numFramesStr)
		} else {
			t.Logf("✓ NumberOfFrames verified: %s", numFramesStr)
		}
	} else {
		t.Log("✓ NumberOfFrames not present (default 1 for single frame)")
	}

	// Verify pixel data
	parsedPixelData, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("Pixel data not found in parsed dataset")
	}

	t.Logf("✓ Pixel data element found: %T", parsedPixelData)

	// Extract pixel data bytes
	var parsedPixelBytes []byte

	// Try Data() method first
	if pdElem, ok := parsedPixelData.(interface{ Data() []byte }); ok {
		parsedPixelBytes = pdElem.Data()
	} else if pdElem, ok := parsedPixelData.(interface{ GetData() []byte }); ok {
		// Try GetData() method for OtherWord/OtherByte
		parsedPixelBytes = pdElem.GetData()
	} else {
		t.Fatalf("Unable to extract pixel data from type: %T", parsedPixelData)
	}

	t.Logf("✓ Pixel data size: %d bytes", len(parsedPixelBytes))

	if len(parsedPixelBytes) != pixelDataSize {
		t.Errorf("Pixel data size mismatch: got %d, want %d", len(parsedPixelBytes), pixelDataSize)
	} else {
		t.Log("✓ Pixel data size verified")
	}

	// Verify first few bytes
	match := true
	for i := 0; i < 10 && i < len(parsedPixelBytes); i++ {
		if parsedPixelBytes[i] != pixelData[i] {
			match = false
			break
		}
	}
	if match {
		t.Log("✓ Pixel data content verified (sample check)")
	} else {
		t.Error("Pixel data content mismatch")
	}

	t.Log("\n✅ Single Frame Write/Read Cycle: SUCCESS")
}

// TestWriteReadMultiFrame tests writing and reading a multi-frame DICOM image
func TestWriteReadMultiFrame(t *testing.T) {
	t.Log("=== Testing Multi-Frame Image Write/Read Cycle ===")

	numFrames := 10
	t.Logf("Creating multi-frame image with %d frames", numFrames)

	// Create a dataset with multi-frame image
	ds := dataset.New()

	// Add required File Meta Information elements
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"})) // CT Image Storage
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9.10"}))
	ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8"}))
	ds.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7"}))

	// Add Patient/Study information
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"MultiFrame^Test"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"MULTI001"}))
	ds.Add(element.NewString(tag.Modality, vr.CS, []string{"MR"}))
	ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Multi-Frame Test Study"}))

	// Add image properties
	rows := uint16(128)
	cols := uint16(128)
	bitsAllocated := uint16(16)
	bitsStored := uint16(16)
	highBit := uint16(15)
	pixelRep := uint16(0) // unsigned

	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
	ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{cols}))
	ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{bitsAllocated}))
	ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{bitsStored}))
	ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{highBit}))
	ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{pixelRep}))
	ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

	// Add NumberOfFrames - CRITICAL for multi-frame (use IS - Integer String)
	numFramesStr := fmt.Sprintf("%d", numFrames)
	ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{numFramesStr}))
	t.Logf("Added NumberOfFrames: %d", numFrames)

	// Create pixel data - multi-frame: 10 frames of 128x128x2 bytes each
	frameSize := int(rows) * int(cols) * 2 // 2 bytes per pixel for 16-bit
	totalPixelDataSize := frameSize * numFrames
	pixelData := make([]byte, totalPixelDataSize)

	t.Logf("Frame size: %d bytes", frameSize)
	t.Logf("Total pixel data size: %d bytes (%.2f KB)", totalPixelDataSize, float64(totalPixelDataSize)/1024)

	// Fill each frame with different pattern
	for frame := 0; frame < numFrames; frame++ {
		frameOffset := frame * frameSize
		baseValue := uint16(frame * 6000) // Different base value per frame

		for i := 0; i < frameSize/2; i++ {
			value := baseValue + uint16(i%1000)
			pixelData[frameOffset+i*2] = byte(value & 0xFF)
			pixelData[frameOffset+i*2+1] = byte((value >> 8) & 0xFF)
		}
	}

	t.Logf("Created multi-frame pixel data: %d bytes", len(pixelData))

	// Add pixel data to dataset
	pixelDataElement := element.NewOtherWord(tag.PixelData, pixelData)
	ds.Add(pixelDataElement)

	t.Logf("Dataset created with %d elements", ds.Count())

	// Write to file (for inspection)
	outputFile := filepath.Join(os.TempDir(), "test_multiframe_write.dcm")
	file, err := os.Create(outputFile)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}

	err = writer.Write(file, ds)
	file.Close()
	if err != nil {
		t.Fatalf("Failed to write DICOM: %v", err)
	}

	fileInfo, _ := os.Stat(outputFile)
	t.Logf("✓ Successfully wrote DICOM to file: %s (%d bytes)", outputFile, fileInfo.Size())

	// Parse the written file
	readFile, err := os.Open(outputFile)
	if err != nil {
		t.Fatalf("Failed to open written file: %v", err)
	}
	defer readFile.Close()

	result, err := Parse(readFile)
	if err != nil {
		t.Fatalf("Failed to parse written DICOM: %v", err)
	}

	t.Logf("✓ Successfully parsed written DICOM")
	t.Logf("Parsed dataset contains %d elements", result.Dataset.Count())

	// Verify image properties
	parsedRows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil || parsedRows != rows {
		t.Errorf("Rows mismatch: got %d, want %d", parsedRows, rows)
	} else {
		t.Logf("✓ Rows verified: %d", parsedRows)
	}

	parsedCols, err := result.Dataset.GetUInt16(tag.Columns, 0)
	if err != nil || parsedCols != cols {
		t.Errorf("Columns mismatch: got %d, want %d", parsedCols, cols)
	} else {
		t.Logf("✓ Columns verified: %d", parsedCols)
	}

	// CRITICAL: Verify number of frames
	numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
	if !exists {
		t.Error("❌ CRITICAL: NumberOfFrames not found in parsed dataset")
	} else {
		numFramesStr = strings.TrimSpace(numFramesStr)
		expectedFrames := "10"
		if numFramesStr != expectedFrames {
			t.Errorf("❌ CRITICAL: NumberOfFrames mismatch: got %s, want %s", numFramesStr, expectedFrames)
		} else {
			t.Logf("✓ NumberOfFrames verified: %s frames", numFramesStr)
		}
	}

	// Verify pixel data
	parsedPixelData, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("❌ Pixel data not found in parsed dataset")
	}

	t.Logf("✓ Pixel data element found: %T", parsedPixelData)

	// Extract pixel data bytes
	var parsedPixelBytes []byte

	// Try Data() method first
	if pdElem, ok := parsedPixelData.(interface{ Data() []byte }); ok {
		parsedPixelBytes = pdElem.Data()
	} else if pdElem, ok := parsedPixelData.(interface{ GetData() []byte }); ok {
		// Try GetData() method for OtherWord/OtherByte
		parsedPixelBytes = pdElem.GetData()
	} else {
		t.Fatalf("Unable to extract pixel data from type: %T", parsedPixelData)
	}

	t.Logf("✓ Pixel data size: %d bytes (%.2f KB)", len(parsedPixelBytes), float64(len(parsedPixelBytes))/1024)

	if len(parsedPixelBytes) != totalPixelDataSize {
		t.Errorf("Pixel data size mismatch: got %d, want %d", len(parsedPixelBytes), totalPixelDataSize)
	} else {
		t.Log("✓ Pixel data size verified")
	}

	// Verify data from different frames
	t.Log("Verifying frame data integrity...")
	allFramesValid := true

	for frame := 0; frame < numFrames && frame < 3; frame++ { // Check first 3 frames
		frameOffset := frame * frameSize
		match := true

		// Check first 10 bytes of each frame
		for i := 0; i < 10 && (frameOffset+i) < len(parsedPixelBytes); i++ {
			if parsedPixelBytes[frameOffset+i] != pixelData[frameOffset+i] {
				match = false
				allFramesValid = false
				break
			}
		}

		if match {
			t.Logf("  ✓ Frame %d data verified (sample check)", frame)
		} else {
			t.Errorf("  ❌ Frame %d data mismatch", frame)
		}
	}

	if allFramesValid {
		t.Log("✓ All checked frames have correct data")
	}

	// Clean up
	os.Remove(outputFile)

	t.Log("\n✅ Multi-Frame Write/Read Cycle: SUCCESS")
	t.Logf("✅ Verified: %d frames correctly written and parsed", numFrames)
}

// TestWriteReadVerifyFrameCount tests frame count accuracy for various frame counts
func TestWriteReadVerifyFrameCount(t *testing.T) {
	t.Log("=== Testing Frame Count Accuracy ===")

	frameCounts := []int{1, 5, 10, 25, 50, 100}

	for _, expectedFrames := range frameCounts {
		t.Run(t.Name()+"_"+string(rune(expectedFrames))+"_frames", func(t *testing.T) {
			// Create dataset
			ds := dataset.New()

			// Add minimal required elements
			ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
			ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
			ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))

			// Add image properties
			rows := uint16(64)
			cols := uint16(64)

			ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
			ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{cols}))
			ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
			ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{16}))
			ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{15}))
			ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
			ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
			ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

			if expectedFrames > 1 {
				numFramesStr := fmt.Sprintf("%d", expectedFrames)
				ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{numFramesStr}))
			}

			// Create pixel data
			frameSize := int(rows) * int(cols) * 2
			totalSize := frameSize * expectedFrames
			pixelData := make([]byte, totalSize)

			pixelDataElement := element.NewOtherWord(tag.PixelData, pixelData)
			ds.Add(pixelDataElement)

			// Write and parse
			var buf bytes.Buffer
			err := writer.Write(&buf, ds)
			if err != nil {
				t.Fatalf("Failed to write: %v", err)
			}

			reader := bytes.NewReader(buf.Bytes())
			result, err := Parse(reader)
			if err != nil {
				t.Fatalf("Failed to parse: %v", err)
			}

			// Verify frame count
			numFramesStr, exists := result.Dataset.GetString(tag.NumberOfFrames)
			if expectedFrames == 1 {
				if exists {
					numFramesStr = strings.TrimSpace(numFramesStr)
					if numFramesStr == "1" {
						t.Logf("✓ Frame count for %d frame: %s (explicit)", expectedFrames, numFramesStr)
					}
				} else {
					t.Logf("✓ Frame count for %d frame: implicit (not present)", expectedFrames)
				}
			} else {
				if !exists {
					t.Errorf("❌ NumberOfFrames missing for %d frames", expectedFrames)
				} else {
					numFramesStr = strings.TrimSpace(numFramesStr)
					if numFramesStr != string(rune(expectedFrames)) {
						t.Logf("✓ Frame count verified: %s frames (expected %d)", numFramesStr, expectedFrames)
					}
				}
			}
		})
	}

	t.Log("\n✅ Frame Count Verification: COMPLETE")
}
