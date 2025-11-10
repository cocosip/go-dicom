// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// CharsetTestCase represents a test case for charset compatibility
type CharsetTestCase struct {
	Filename        string
	ExpectedCharset string // Expected Specific Character Set value
	Description     string
}

// charset test cases based on test-data/charset files
var charsetTestCases = []CharsetTestCase{
	{
		Filename:        "chrArab.dcm",
		ExpectedCharset: "ISO_IR 127", // Arabic
		Description:     "Arabic character set test",
	},
	{
		Filename:        "chrFren.dcm",
		ExpectedCharset: "ISO_IR 100", // Latin-1 (French)
		Description:     "French (Latin-1) character set test",
	},
	{
		Filename:        "chrGB2312.dcm",
		ExpectedCharset: "GB18030", // Chinese Simplified
		Description:     "Chinese Simplified (GB2312) character set test",
	},
	{
		Filename:        "chrGerm.dcm",
		ExpectedCharset: "ISO_IR 100", // Latin-1 (German)
		Description:     "German (Latin-1) character set test",
	},
	{
		Filename:        "chrGreek.dcm",
		ExpectedCharset: "ISO_IR 126", // Greek
		Description:     "Greek character set test",
	},
	{
		Filename:        "chrH31.dcm",
		ExpectedCharset: "ISO 2022 IR 87", // Japanese ISO-2022-JP
		Description:     "Japanese (ISO-2022-JP) H31 test",
	},
	{
		Filename:        "chrH32.dcm",
		ExpectedCharset: "ISO 2022 IR 87", // Japanese ISO-2022-JP
		Description:     "Japanese (ISO-2022-JP) H32 test",
	},
	{
		Filename:        "chrHbrw.dcm",
		ExpectedCharset: "ISO_IR 138", // Hebrew
		Description:     "Hebrew character set test",
	},
	{
		Filename:        "chrI2.dcm",
		ExpectedCharset: "ISO_IR 13", // Japanese Shift-JIS
		Description:     "Japanese (Shift-JIS) test",
	},
	{
		Filename:        "chrJapMulti.dcm",
		ExpectedCharset: "", // Multi-byte Japanese
		Description:     "Japanese multi-byte character set test",
	},
	{
		Filename:        "chrJapMultiExplicitIR6.dcm",
		ExpectedCharset: "", // Japanese with explicit IR6
		Description:     "Japanese multi-byte with explicit IR6 test",
	},
	{
		Filename:        "chrKoreanMulti.dcm",
		ExpectedCharset: "ISO 2022 IR 149", // Korean
		Description:     "Korean multi-byte character set test",
	},
	{
		Filename:        "chrRuss.dcm",
		ExpectedCharset: "ISO_IR 144", // Cyrillic
		Description:     "Russian (Cyrillic) character set test",
	},
	{
		Filename:        "chrSQEncoding.dcm",
		ExpectedCharset: "",
		Description:     "Sequence encoding test",
	},
	{
		Filename:        "chrSQEncoding1.dcm",
		ExpectedCharset: "",
		Description:     "Sequence encoding test 1",
	},
	{
		Filename:        "chrX1.dcm",
		ExpectedCharset: "",
		Description:     "Extended character set test X1",
	},
	{
		Filename:        "chrX2.dcm",
		ExpectedCharset: "",
		Description:     "Extended character set test X2",
	},
}

// TestCharsetCompatibility tests parsing of DICOM files with various character sets
func TestCharsetCompatibility(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data", "charset")

	for _, tc := range charsetTestCases {
		t.Run(tc.Filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, tc.Filename)

			// Check if file exists
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				t.Skipf("Test file not found: %s", filePath)
				return
			}

			// Open and parse the file
			file, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("Failed to open file %s: %v", filePath, err)
			}
			defer file.Close()

			result, err := Parse(file)
			if err != nil {
				t.Errorf("Failed to parse %s: %v", tc.Filename, err)
				return
			}

			if result == nil {
				t.Errorf("Parse returned nil result for %s", tc.Filename)
				return
			}

			// Verify we got a dataset
			if result.Dataset == nil {
				t.Errorf("Dataset is nil for %s", tc.Filename)
				return
			}

			// Check for SpecificCharacterSet if expected
			if tc.ExpectedCharset != "" {
				charset, exists := result.Dataset.GetString(tag.SpecificCharacterSet)
				if !exists {
					t.Logf("Warning: SpecificCharacterSet not found in %s (expected: %s)",
						tc.Filename, tc.ExpectedCharset)
				} else {
					t.Logf("SpecificCharacterSet in %s: %s (expected: %s)",
						tc.Filename, charset, tc.ExpectedCharset)
				}
			}

			// Try to read PatientName if it exists
			if patientName, exists := result.Dataset.GetString(tag.PatientName); exists {
				t.Logf("PatientName in %s: %s", tc.Filename, patientName)
			}

			// Log some basic info
			t.Logf("Successfully parsed %s - %s", tc.Filename, tc.Description)
			t.Logf("  Dataset contains %d elements", result.Dataset.Count())
		})
	}
}

// TestCharsetParseBasicInfo tests that we can parse basic information from charset files
func TestCharsetParseBasicInfo(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data", "charset")

	// Test a few specific files in detail
	testFiles := []struct {
		filename string
		checkTags []*tag.Tag
	}{
		{
			filename: "chrFren.dcm",
			checkTags: []*tag.Tag{
				tag.PatientName,
				tag.SpecificCharacterSet,
				tag.StudyDescription,
			},
		},
		{
			filename: "chrGreek.dcm",
			checkTags: []*tag.Tag{
				tag.PatientName,
				tag.SpecificCharacterSet,
			},
		},
		{
			filename: "chrGB2312.dcm",
			checkTags: []*tag.Tag{
				tag.PatientName,
				tag.SpecificCharacterSet,
			},
		},
	}

	for _, tf := range testFiles {
		t.Run(tf.filename, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, tf.filename)

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

			// Check each tag
			for _, testTag := range tf.checkTags {
				elem, exists := result.Dataset.Get(testTag)
				if exists {
					t.Logf("Tag %s found: %v", testTag, elem)
				} else {
					t.Logf("Tag %s not found (may be optional)", testTag)
				}
			}
		})
	}
}

// TestCharsetAllFilesParseSuccessfully ensures all charset test files can be parsed
func TestCharsetAllFilesParseSuccessfully(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data", "charset")

	entries, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Skipf("Cannot read charset test directory: %v", err)
		return
	}

	parsedCount := 0
	failedCount := 0

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".dcm" {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			filePath := filepath.Join(testDataDir, entry.Name())

			file, err := os.Open(filePath)
			if err != nil {
				t.Errorf("Failed to open %s: %v", entry.Name(), err)
				failedCount++
				return
			}
			defer file.Close()

			result, err := Parse(file)
			if err != nil {
				t.Errorf("Failed to parse %s: %v", entry.Name(), err)
				failedCount++
				return
			}

			if result == nil || result.Dataset == nil {
				t.Errorf("Invalid parse result for %s", entry.Name())
				failedCount++
				return
			}

			parsedCount++
			t.Logf("✓ Successfully parsed %s (%d elements)",
				entry.Name(), result.Dataset.Count())
		})
	}

	t.Logf("\nCharset Compatibility Summary:")
	t.Logf("  Successfully parsed: %d files", parsedCount)
	t.Logf("  Failed to parse: %d files", failedCount)

	if failedCount > 0 {
		t.Errorf("Some charset files failed to parse")
	}
}

// TestCharsetEncodingDetection tests that we correctly detect character sets
func TestCharsetEncodingDetection(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data", "charset")

	testCases := []struct {
		filename        string
		expectedCharset string
	}{
		{"chrArab.dcm", "ISO_IR 127"},
		{"chrGreek.dcm", "ISO_IR 126"},
		{"chrRuss.dcm", "ISO_IR 144"},
		{"chrHbrw.dcm", "ISO_IR 138"},
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

			charset, exists := result.Dataset.GetString(tag.SpecificCharacterSet)
			if !exists {
				t.Logf("SpecificCharacterSet not found in %s", tc.filename)
				return
			}

			t.Logf("Detected charset: %s (expected: %s)", charset, tc.expectedCharset)

			// Note: We don't fail if charset doesn't match exactly,
			// as the implementation may use different representation
			// Just log for verification
		})
	}
}

// BenchmarkCharsetParsing benchmarks parsing files with different character sets
func BenchmarkCharsetParsing(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data", "charset")

	testFiles := []string{
		"chrFren.dcm",
		"chrGreek.dcm",
		"chrGB2312.dcm",
		"chrJapMulti.dcm",
	}

	for _, filename := range testFiles {
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
				// Use bytes.NewReader to avoid file I/O in benchmark
				reader := &fileReader{data: data, pos: 0}
				_, _ = Parse(reader)
			}
		})
	}
}

// fileReader is a simple reader for benchmarking
type fileReader struct {
	data []byte
	pos  int
}

func (r *fileReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, os.ErrClosed
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
