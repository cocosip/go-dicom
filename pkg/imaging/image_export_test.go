// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

// TestExportPaletteToJPEG tests the complete workflow of reading a palette DICOM file
// and exporting it to a JPEG image.
func TestExportPaletteToJPEG(t *testing.T) {
	// Path to test DICOM file
	inputFile := filepath.Join("..", "..", "test-data", "TestPattern_Palette.dcm")

	// Check if test file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", inputFile)
	}

	// Parse DICOM file
	t.Logf("Parsing DICOM file: %s", inputFile)
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to parse DICOM file: %v", err)
	}

	// Create PixelData from dataset
	t.Log("Extracting pixel data...")
	pixelData, err := CreatePixelData(result.Dataset)
	if err != nil {
		t.Fatalf("Failed to extract pixel data: %v", err)
	}

	// Log photometric interpretation
	// Note: CreatePixelData automatically converts PALETTE COLOR to RGB
	if pixelData.Info.PhotometricInterpretation == nil {
		t.Fatal("PhotometricInterpretation is nil")
	}

	photometric := pixelData.Info.PhotometricInterpretation.Value
	t.Logf("Photometric Interpretation: %s", photometric)
	t.Logf("Note: Original PALETTE COLOR was automatically converted to RGB by CreatePixelData")

	// Create DICOM image
	t.Log("Creating DICOM image...")
	dicomImage := NewDicomImage(pixelData)

	// Verify image properties
	t.Logf("Image properties:")
	t.Logf("  Dimensions: %dx%d", dicomImage.Width(), dicomImage.Height())
	t.Logf("  Number of frames: %d", dicomImage.NumberOfFrames())
	t.Logf("  Grayscale: %v", dicomImage.IsGrayscale())
	t.Logf("  Bits Allocated: %d", pixelData.Info.BitsAllocated)
	t.Logf("  Bits Stored: %d", pixelData.Info.BitsStored)
	t.Logf("  Samples Per Pixel: %d", pixelData.Info.SamplesPerPixel)

	// Validate expected properties for palette color
	if dicomImage.IsGrayscale() {
		t.Error("Expected color image, got grayscale")
	}

	if dicomImage.Width() == 0 || dicomImage.Height() == 0 {
		t.Fatalf("Invalid image dimensions: %dx%d", dicomImage.Width(), dicomImage.Height())
	}

	if dicomImage.NumberOfFrames() == 0 {
		t.Fatal("Expected at least 1 frame")
	}

	// Create temporary output file
	outputFile := filepath.Join(t.TempDir(), "test_output.jpeg")
	t.Logf("Output file: %s", outputFile)

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			t.Errorf("Failed to close output file: %v", err)
		}
	}()

	// Create export options
	exportOptions := &render.ExportOptions{
		Format:      render.FormatJPEG,
		JPEGQuality: 90,
	}

	// Render frame to file
	frameIndex := 0
	t.Logf("Rendering frame %d as JPEG...", frameIndex)
	if err := dicomImage.RenderFrame(outFile, frameIndex, exportOptions); err != nil {
		t.Fatalf("Failed to render image: %v", err)
	}

	// Flush the file
	if err := outFile.Sync(); err != nil {
		t.Errorf("Failed to sync file: %v", err)
	}

	// Verify output file was created and has content
	info, err := os.Stat(outputFile)
	if err != nil {
		t.Fatalf("Failed to stat output file: %v", err)
	}

	t.Logf("Output file size: %d bytes", info.Size())

	// JPEG files should have reasonable size (at least a few hundred bytes)
	if info.Size() < 100 {
		t.Errorf("Output file seems too small: %d bytes", info.Size())
	}

	// Verify JPEG magic number (FF D8 FF)
	_, err = outFile.Seek(0, 0)
	if err != nil {
		t.Fatalf("Failed to seek to start of file: %v", err)
	}
	header := make([]byte, 3)
	n, err := outFile.Read(header)
	if err != nil || n != 3 {
		t.Fatalf("Failed to read file header: %v", err)
	}

	if header[0] != 0xFF || header[1] != 0xD8 || header[2] != 0xFF {
		t.Errorf("Invalid JPEG header: %02X %02X %02X (expected FF D8 FF)",
			header[0], header[1], header[2])
	}

	t.Log("Successfully exported palette DICOM to JPEG!")
}

// TestExportPaletteToPNG tests exporting a palette DICOM file to PNG format.
func TestExportPaletteToPNG(t *testing.T) {
	// Path to test DICOM file
	inputFile := filepath.Join("..", "..", "test-data", "TestPattern_Palette.dcm")

	// Check if test file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", inputFile)
	}

	// Parse DICOM file
	t.Logf("Parsing DICOM file: %s", inputFile)
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to parse DICOM file: %v", err)
	}

	// Create PixelData from dataset
	pixelData, err := CreatePixelData(result.Dataset)
	if err != nil {
		t.Fatalf("Failed to extract pixel data: %v", err)
	}

	// Create DICOM image
	dicomImage := NewDicomImage(pixelData)

	// Create temporary output file
	outputFile := filepath.Join(t.TempDir(), "test_output.png")
	t.Logf("Output file: %s", outputFile)

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			t.Errorf("Failed to close output file: %v", err)
		}
	}()

	// Create export options for PNG
	exportOptions := &render.ExportOptions{
		Format: render.FormatPNG,
	}

	// Render frame to file
	frameIndex := 0
	t.Logf("Rendering frame %d as PNG...", frameIndex)
	if err := dicomImage.RenderFrame(outFile, frameIndex, exportOptions); err != nil {
		t.Fatalf("Failed to render image: %v", err)
	}

	// Flush the file
	if err := outFile.Sync(); err != nil {
		t.Errorf("Failed to sync file: %v", err)
	}

	// Verify output file was created and has content
	info, err := os.Stat(outputFile)
	if err != nil {
		t.Fatalf("Failed to stat output file: %v", err)
	}

	t.Logf("Output file size: %d bytes", info.Size())

	if info.Size() < 100 {
		t.Errorf("Output file seems too small: %d bytes", info.Size())
	}

	// Verify PNG magic number (89 50 4E 47 0D 0A 1A 0A)
	_, err = outFile.Seek(0, 0)
	if err != nil {
		t.Fatalf("Failed to seek to start of file: %v", err)
	}
	header := make([]byte, 8)
	n, err := outFile.Read(header)
	if err != nil || n != 8 {
		t.Fatalf("Failed to read file header: %v", err)
	}

	pngMagic := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	for i := 0; i < 8; i++ {
		if header[i] != pngMagic[i] {
			t.Errorf("Invalid PNG header at byte %d: %02X (expected %02X)",
				i, header[i], pngMagic[i])
		}
	}

	t.Log("Successfully exported palette DICOM to PNG!")
}

// TestExportMultiplePaletteFormats tests exporting with different quality settings
// and validates the basic functionality across formats.
func TestExportMultiplePaletteFormats(t *testing.T) {
	// Path to test DICOM file
	inputFile := filepath.Join("..", "..", "test-data", "TestPattern_Palette.dcm")

	// Check if test file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		t.Skipf("Test file not found: %s", inputFile)
	}

	// Parse DICOM file once
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to parse DICOM file: %v", err)
	}

	pixelData, err := CreatePixelData(result.Dataset)
	if err != nil {
		t.Fatalf("Failed to extract pixel data: %v", err)
	}

	dicomImage := NewDicomImage(pixelData)

	// Test cases for different export formats and settings
	testCases := []struct {
		name        string
		format      render.ExportFormat
		jpegQuality int
		extension   string
		magicBytes  []byte
	}{
		{
			name:        "JPEG Quality 75",
			format:      render.FormatJPEG,
			jpegQuality: 75,
			extension:   ".jpeg",
			magicBytes:  []byte{0xFF, 0xD8, 0xFF},
		},
		{
			name:        "JPEG Quality 100",
			format:      render.FormatJPEG,
			jpegQuality: 100,
			extension:   ".jpeg",
			magicBytes:  []byte{0xFF, 0xD8, 0xFF},
		},
		{
			name:       "PNG",
			format:     render.FormatPNG,
			extension:  ".png",
			magicBytes: []byte{0x89, 0x50, 0x4E, 0x47},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create temporary output file
			outputFile := filepath.Join(t.TempDir(), "test_output"+tc.extension)

			// Create output file
			outFile, err := os.Create(outputFile)
			if err != nil {
				t.Fatalf("Failed to create output file: %v", err)
			}
			defer func() {
				if closeErr := outFile.Close(); closeErr != nil {
					t.Errorf("Failed to close output file: %v", closeErr)
				}
			}()

			// Create export options
			exportOptions := &render.ExportOptions{
				Format:      tc.format,
				JPEGQuality: tc.jpegQuality,
			}

			// Render frame
			if err := dicomImage.RenderFrame(outFile, 0, exportOptions); err != nil {
				t.Fatalf("Failed to render image: %v", err)
			}

			// Flush
			if err := outFile.Sync(); err != nil {
				t.Errorf("Failed to sync file: %v", err)
			}

			// Verify file size
			info, err := os.Stat(outputFile)
			if err != nil {
				t.Fatalf("Failed to stat output file: %v", err)
			}

			if info.Size() < 100 {
				t.Errorf("Output file seems too small: %d bytes", info.Size())
			}

			// Verify magic bytes
			_, err = outFile.Seek(0, 0)
			if err != nil {
				t.Fatalf("Failed to seek to start of file: %v", err)
			}
			header := make([]byte, len(tc.magicBytes))
			n, err := outFile.Read(header)
			if err != nil || n != len(tc.magicBytes) {
				t.Fatalf("Failed to read file header: %v", err)
			}

			for i := 0; i < len(tc.magicBytes); i++ {
				if header[i] != tc.magicBytes[i] {
					t.Errorf("Invalid header at byte %d: %02X (expected %02X)",
						i, header[i], tc.magicBytes[i])
				}
			}

			t.Logf("Successfully exported to %s (size: %d bytes)", tc.name, info.Size())
		})
	}
}
