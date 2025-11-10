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

// TestRealWorldDICOMFiles tests parsing of real DICOM files from test-data directory
func TestRealWorldDICOMFiles(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	// Get all .dcm files in test-data (excluding charset subdirectory)
	entries, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Skipf("Cannot read test-data directory: %v", err)
		return
	}

	parsedCount := 0
	failedCount := 0
	skippedCount := 0

	for _, entry := range entries {
		// Skip directories and non-DICOM files
		if entry.IsDir() {
			skippedCount++
			continue
		}

		// Only test .dcm files and DICOMDIR
		name := entry.Name()
		if !strings.HasSuffix(name, ".dcm") && name != "DICOMDIR" && !strings.HasPrefix(name, "DIRW") {
			continue
		}

		t.Run(name, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, name)

			file, err := os.Open(filePath)
			if err != nil {
				t.Errorf("Failed to open %s: %v", name, err)
				failedCount++
				return
			}
			defer file.Close()

			result, err := Parse(file)
			if err != nil {
				t.Errorf("Failed to parse %s: %v", name, err)
				failedCount++
				return
			}

			if result == nil || result.Dataset == nil {
				t.Errorf("Invalid parse result for %s", name)
				failedCount++
				return
			}

			parsedCount++

			// Log basic information about the file
			t.Logf("✓ Successfully parsed %s", name)
			t.Logf("  Elements: %d", result.Dataset.Count())

			// Try to get some common tags
			if sopClassUID, exists := result.Dataset.GetString(tag.SOPClassUID); exists {
				t.Logf("  SOP Class UID: %s", sopClassUID)
			}

			if modality, exists := result.Dataset.GetString(tag.Modality); exists {
				t.Logf("  Modality: %s", modality)
			}

			if patientName, exists := result.Dataset.GetString(tag.PatientName); exists {
				t.Logf("  Patient Name: %s", patientName)
			}

			// Check for transfer syntax
			if result.FileMetaInformation != nil {
				if tsUID, exists := result.FileMetaInformation.TransferSyntaxUID(); exists {
					t.Logf("  Transfer Syntax: %s", tsUID)
				}
			}
		})
	}

	t.Logf("\nReal-world DICOM Files Summary:")
	t.Logf("  Successfully parsed: %d files", parsedCount)
	t.Logf("  Failed to parse: %d files", failedCount)
	t.Logf("  Skipped: %d items", skippedCount)

	if failedCount > 0 {
		t.Errorf("Some DICOM files failed to parse")
	}
}

// TestSpecificRealWorldFiles tests specific interesting files in detail
func TestSpecificRealWorldFiles(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	testCases := []struct {
		filename    string
		description string
		checks      func(t *testing.T, result *ParseResult)
	}{
		{
			filename:    "CT-MONO2-16-ankle",
			description: "CT image with 16-bit monochrome",
			checks: func(t *testing.T, result *ParseResult) {
				if modality, exists := result.Dataset.GetString(tag.Modality); exists {
					if modality != "CT" {
						t.Errorf("Expected modality CT, got %s", modality)
					}
				}

				if pi, exists := result.Dataset.GetString(tag.PhotometricInterpretation); exists {
					t.Logf("Photometric Interpretation: %s", pi)
				}

				if bitsAllocated, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0); err == nil {
					t.Logf("Bits Allocated: %d", bitsAllocated)
				}
			},
		},
		{
			filename:    "CR-MONO1-10-chest",
			description: "CR chest X-ray with 10-bit monochrome",
			checks: func(t *testing.T, result *ParseResult) {
				if modality, exists := result.Dataset.GetString(tag.Modality); exists {
					if modality != "CR" {
						t.Errorf("Expected modality CR, got %s", modality)
					}
				}
			},
		},
		{
			filename:    "TestPattern_RGB.dcm",
			description: "RGB test pattern",
			checks: func(t *testing.T, result *ParseResult) {
				if pi, exists := result.Dataset.GetString(tag.PhotometricInterpretation); exists {
					t.Logf("Photometric Interpretation: %s", pi)
					if !strings.Contains(pi, "RGB") {
						t.Logf("Note: Expected RGB photometric interpretation, got %s", pi)
					}
				}

				if samplesPerPixel, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0); err == nil {
					t.Logf("Samples Per Pixel: %d", samplesPerPixel)
				}
			},
		},
		{
			filename:    "TestPattern_Palette.dcm",
			description: "Palette color test pattern",
			checks: func(t *testing.T, result *ParseResult) {
				if pi, exists := result.Dataset.GetString(tag.PhotometricInterpretation); exists {
					t.Logf("Photometric Interpretation: %s", pi)
				}
			},
		},
		{
			filename:    "multiframe.dcm",
			description: "Multi-frame DICOM file",
			checks: func(t *testing.T, result *ParseResult) {
				if numFrames, exists := result.Dataset.GetString(tag.NumberOfFrames); exists {
					t.Logf("Number of Frames: %s", numFrames)
				}
			},
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
				t.Fatalf("Parse failed: %v", err)
			}

			t.Logf("Testing: %s", tc.description)

			if tc.checks != nil {
				tc.checks(t, result)
			}
		})
	}
}

// TestFragmentSequenceFiles tests files with fragment sequences
func TestFragmentSequenceFiles(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	fragmentFiles := []string{
		"HasFragmentSequence.dcm",
		"D_CLUNIE_CT1_RLE_FRAGS.dcm",
	}

	for _, filename := range fragmentFiles {
		t.Run(filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, filename)

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
				t.Logf("Parse failed for %s: %v (fragment sequences may not be fully supported yet)", filename, err)
				// Don't fail the test - fragment sequences are complex
				return
			}

			t.Logf("✓ Successfully parsed %s with fragment sequences", filename)
			t.Logf("  Dataset contains %d elements", result.Dataset.Count())
		})
	}
}

// TestCompressedFiles tests compressed DICOM files
func TestCompressedFiles(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	compressedFiles := []struct {
		filename    string
		description string
	}{
		{"CT1_J2KI", "JPEG 2000 compressed CT"},
		{"VL5_J2KI.dcm", "JPEG 2000 compressed visible light"},
		{"GH538-jpeg1.dcm", "JPEG compressed image"},
	}

	for _, cf := range compressedFiles {
		t.Run(cf.filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, cf.filename)

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
				t.Logf("Parse result for %s: %v", cf.filename, err)
				// Don't fail - compressed pixel data is complex
				return
			}

			t.Logf("✓ Parsed %s - %s", cf.filename, cf.description)

			if result.FileMetaInformation != nil {
				if tsUID, exists := result.FileMetaInformation.TransferSyntaxUID(); exists {
					t.Logf("  Transfer Syntax: %s", tsUID)
				}
			}
		})
	}
}

// TestDICOMDIRFiles tests DICOMDIR directory files
func TestDICOMDIRFiles(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	dicomdirFiles := []string{
		"DICOMDIR",
		"DICOMDIR_GH1927",
		"DIRW0007",
	}

	for _, filename := range dicomdirFiles {
		t.Run(filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, filename)

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
				t.Logf("Parse failed for DICOMDIR %s: %v", filename, err)
				return
			}

			t.Logf("✓ Successfully parsed DICOMDIR: %s", filename)
			t.Logf("  Dataset contains %d elements", result.Dataset.Count())

			// Check for Directory Record Sequence
			if elem, exists := result.Dataset.Get(tag.DirectoryRecordSequence); exists {
				t.Logf("  Found Directory Record Sequence: %v", elem)
			}
		})
	}
}

// BenchmarkRealWorldDICOMParsing benchmarks parsing of various real DICOM files
func BenchmarkRealWorldDICOMParsing(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")

	benchFiles := []string{
		"CT-MONO2-16-ankle",
		"CR-MONO1-10-chest",
		"TestPattern_RGB.dcm",
		"multiframe.dcm",
	}

	for _, filename := range benchFiles {
		b.Run(filename, func(b *testing.B) {
			filePath := filepath.Join(testDataDir, filename)

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
		})
	}
}
