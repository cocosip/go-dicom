// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestRGBImageParsing tests parsing of RGB DICOM images
func TestRGBImageParsing(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestPattern_RGB.dcm")

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
    defer func() { _ = file.Close() }()

	t.Log("Parsing TestPattern_RGB.dcm...")
	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse TestPattern_RGB.dcm: %v", err)
	}

	if result == nil || result.Dataset == nil {
		t.Fatal("Parse result or dataset is nil")
	}

	t.Logf("✓ Successfully parsed TestPattern_RGB.dcm")
	t.Logf("Dataset contains %d elements", result.Dataset.Count())

	// Check image properties
	t.Log("\n=== Image Properties ===")

	// Photometric Interpretation - must be RGB
	photoInterp, exists := result.Dataset.GetString(tag.PhotometricInterpretation)
	if !exists {
		t.Fatal("PhotometricInterpretation not found")
	}
	t.Logf("Photometric Interpretation: %s", photoInterp)
	if photoInterp != "RGB" && photoInterp != "RGB " {
		t.Errorf("Expected RGB photometric interpretation, got: %s", photoInterp)
	} else {
		t.Log("✓ Confirmed RGB color format")
	}

	// Samples Per Pixel - must be 3 for RGB
	samplesPerPixel, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0)
	if err != nil {
		t.Fatalf("Failed to get SamplesPerPixel: %v", err)
	}
	t.Logf("Samples Per Pixel: %d", samplesPerPixel)
	if samplesPerPixel != 3 {
		t.Errorf("Expected 3 samples per pixel for RGB, got: %d", samplesPerPixel)
	} else {
		t.Log("✓ Confirmed 3 color channels (R, G, B)")
	}

	// Image dimensions
	rows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil {
		t.Fatalf("Failed to get Rows: %v", err)
	}
	t.Logf("Rows: %d", rows)

	cols, err := result.Dataset.GetUInt16(tag.Columns, 0)
	if err != nil {
		t.Fatalf("Failed to get Columns: %v", err)
	}
	t.Logf("Columns: %d", cols)
	t.Logf("Image size: %d × %d pixels", cols, rows)

	// Bits Allocated
	bitsAlloc, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0)
	if err != nil {
		t.Fatalf("Failed to get BitsAllocated: %v", err)
	}
	t.Logf("Bits Allocated: %d", bitsAlloc)

	// Bits Stored
	bitsStored, err := result.Dataset.GetUInt16(tag.BitsStored, 0)
	if err != nil {
		t.Fatalf("Failed to get BitsStored: %v", err)
	}
	t.Logf("Bits Stored: %d", bitsStored)

	// High Bit
	highBit, err := result.Dataset.GetUInt16(tag.HighBit, 0)
	if err != nil {
		t.Fatalf("Failed to get HighBit: %v", err)
	}
	t.Logf("High Bit: %d", highBit)

	// Planar Configuration
	planarConfig, err := result.Dataset.GetUInt16(tag.PlanarConfiguration, 0)
	if err != nil {
		// Planar configuration might not be present
		t.Logf("Planar Configuration: not specified (assuming 0 - interlaced)")
	} else {
		t.Logf("Planar Configuration: %d", planarConfig)
		if planarConfig == 0 {
			t.Log("  → Pixel format: RGBRGBRGB... (interlaced)")
		} else {
			t.Log("  → Pixel format: RRR...GGG...BBB... (planar)")
		}
	}

	// Check for Pixel Data
	t.Log("\n=== Pixel Data ===")
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("❌ PixelData tag not found")
	}

	t.Logf("✓ PixelData element found")
	t.Logf("  PixelData tag: %s", pixelDataElem.Tag())
	t.Logf("  PixelData VR: %s", pixelDataElem.ValueRepresentation())
	t.Logf("  PixelData type: %T", pixelDataElem)

	// Calculate expected pixel data size
	expectedSize := uint32(rows) * uint32(cols) * uint32(samplesPerPixel) * uint32(bitsAlloc/8)
	t.Logf("  Expected pixel data size: %d bytes (%.2f KB)", expectedSize, float64(expectedSize)/1024)

	actualSize := pixelDataElem.Length()
	t.Logf("  Actual pixel data size: %d bytes (%.2f KB)", actualSize, float64(actualSize)/1024)

	if actualSize != expectedSize {
		t.Logf("  ⚠ Size mismatch (expected %d, got %d)", expectedSize, actualSize)
	} else {
		t.Log("  ✓ Pixel data size matches expected")
	}

	// Transfer Syntax
	if result.FileMetaInformation != nil {
		if tsUID, exists := result.FileMetaInformation.TransferSyntaxUID(); exists {
			t.Logf("\nTransfer Syntax: %s", tsUID)
		}
	}

	t.Log("\n✓ RGB DICOM image parsed successfully")
	t.Log("✓ All image metadata accessible")
	t.Log("✓ Pixel data present and accessible")
}

// TestRGBImageToImage tests converting RGB DICOM to Go image.Image
func TestRGBImageToImage(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestPattern_RGB.dcm")

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

	// Get image dimensions
	rows, _ := result.Dataset.GetUInt16(tag.Rows, 0)
	cols, _ := result.Dataset.GetUInt16(tag.Columns, 0)
	samplesPerPixel, _ := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0)
	bitsAlloc, _ := result.Dataset.GetUInt16(tag.BitsAllocated, 0)

	// Get planar configuration (0 = interlaced, 1 = planar)
	planarConfig, err := result.Dataset.GetUInt16(tag.PlanarConfiguration, 0)
	if err != nil {
		planarConfig = 0 // Default to interlaced
	}

	t.Logf("Converting %dx%d RGB image to image.Image", cols, rows)

	// Get pixel data
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData not found")
	}

	// Extract pixel data bytes
	var pixelData []byte

	// Check if it's a regular element or fragment
	if otherByte, ok := pixelDataElem.(*element.OtherByte); ok {
		// Regular OB element
		buffer := otherByte.Buffer()
		if buffer == nil {
			t.Fatal("PixelData buffer is nil")
		}
		pixelData = buffer.Data()
	} else if otherWord, ok := pixelDataElem.(*element.OtherWord); ok {
		// OW element
		buffer := otherWord.Buffer()
		if buffer == nil {
			t.Fatal("PixelData buffer is nil")
		}
		pixelData = buffer.Data()
	} else {
		t.Fatalf("Unexpected PixelData type: %T", pixelDataElem)
	}

	t.Logf("Pixel data extracted: %d bytes", len(pixelData))

	// Create Go image
	img := image.NewRGBA(image.Rect(0, 0, int(cols), int(rows)))

	// Convert pixel data to image
	if samplesPerPixel == 3 && bitsAlloc == 8 {
		if planarConfig == 0 {
			// Interlaced: RGBRGBRGB...
			t.Log("Converting interlaced RGB data")
			convertInterlacedRGB(img, pixelData, int(cols), int(rows))
		} else {
			// Planar: RRR...GGG...BBB...
			t.Log("Converting planar RGB data")
			convertPlanarRGB(img, pixelData, int(cols), int(rows))
		}
	} else {
		t.Fatalf("Unsupported format: %d samples/pixel, %d bits", samplesPerPixel, bitsAlloc)
	}

	t.Logf("✓ Image converted successfully")
	t.Logf("  Image size: %dx%d", img.Bounds().Dx(), img.Bounds().Dy())

	// Verify image has data (check a few pixels are not black)
	hasColor := false
	for y := 0; y < img.Bounds().Dy() && !hasColor; y++ {
		for x := 0; x < img.Bounds().Dx() && !hasColor; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r > 0 || g > 0 || b > 0 {
				hasColor = true
			}
		}
	}

	if !hasColor {
		t.Error("Warning: Converted image appears to be completely black")
	} else {
		t.Log("✓ Image contains color data")
	}

	t.Log("\n✓ RGB DICOM successfully converted to image.Image")
}

// TestRGBImageToPNG tests converting RGB DICOM to PNG file
func TestRGBImageToPNG(t *testing.T) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestPattern_RGB.dcm")

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

	// Convert to image
	img, err := convertDICOMToImage(result)
	if err != nil {
		t.Fatalf("Failed to convert to image: %v", err)
	}

	t.Logf("✓ Converted to image: %dx%d", img.Bounds().Dx(), img.Bounds().Dy())

	// Save as PNG to temporary file
	outputPath := filepath.Join(os.TempDir(), "test_rgb_output.png")
	outFile, err := os.Create(outputPath)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer func() { _ = outFile.Close() }()

	err = png.Encode(outFile, img)
	if err != nil {
		t.Fatalf("Failed to encode PNG: %v", err)
	}

	t.Logf("✓ PNG saved to: %s", outputPath)

	// Verify the PNG file was created
	pngInfo, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("Failed to stat PNG file: %v", err)
	}

	t.Logf("✓ PNG file size: %.2f KB", float64(pngInfo.Size())/1024)

	// Verify we can read it back
	pngFile, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open PNG file: %v", err)
	}
	defer func() { _ = pngFile.Close() }()

	decodedImg, err := png.Decode(pngFile)
	if err != nil {
		t.Fatalf("Failed to decode PNG: %v", err)
	}

	t.Logf("✓ PNG successfully read back: %dx%d",
		decodedImg.Bounds().Dx(), decodedImg.Bounds().Dy())

	t.Log("\n✓ RGB DICOM successfully converted to PNG file")
	t.Logf("✓ Output file: %s", outputPath)
}

// Helper function: Convert interlaced RGB data to image
func convertInterlacedRGB(img *image.RGBA, pixelData []byte, width, height int) {
	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if idx+2 >= len(pixelData) {
				return
			}
			r := pixelData[idx]
			g := pixelData[idx+1]
			b := pixelData[idx+2]
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
			idx += 3
		}
	}
}

// Helper function: Convert planar RGB data to image
func convertPlanarRGB(img *image.RGBA, pixelData []byte, width, height int) {
	planeSize := width * height
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*width + x
			if pos >= planeSize || pos+2*planeSize >= len(pixelData) {
				return
			}
			r := pixelData[pos]
			g := pixelData[pos+planeSize]
			b := pixelData[pos+2*planeSize]
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
}

// Helper function: Convert DICOM to Go image
func convertDICOMToImage(result *ParseResult) (image.Image, error) {
	// Get image properties
	rows, err := result.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil {
		return nil, err
	}

	cols, err := result.Dataset.GetUInt16(tag.Columns, 0)
	if err != nil {
		return nil, err
	}

	samplesPerPixel, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0)
	if err != nil {
		return nil, err
	}

	bitsAlloc, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0)
	if err != nil {
		return nil, err
	}

	planarConfig, err := result.Dataset.GetUInt16(tag.PlanarConfiguration, 0)
	if err != nil {
		planarConfig = 0
	}

	// Get pixel data
	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		return nil, image.ErrFormat
	}

	var pixelData []byte
	if otherByte, ok := pixelDataElem.(*element.OtherByte); ok {
		buffer := otherByte.Buffer()
		pixelData = buffer.Data()
	} else if otherWord, ok := pixelDataElem.(*element.OtherWord); ok {
		buffer := otherWord.Buffer()
		pixelData = buffer.Data()
	} else {
		return nil, image.ErrFormat
	}

	// Create image
	img := image.NewRGBA(image.Rect(0, 0, int(cols), int(rows)))

	// Convert based on format
	if samplesPerPixel == 3 && bitsAlloc == 8 {
		if planarConfig == 0 {
			convertInterlacedRGB(img, pixelData, int(cols), int(rows))
		} else {
			convertPlanarRGB(img, pixelData, int(cols), int(rows))
		}
	}

	return img, nil
}

// BenchmarkRGBImageConversion benchmarks RGB DICOM to image conversion
func BenchmarkRGBImageConversion(b *testing.B) {
	testDataDir := filepath.Join("..", "..", "..", "test-data")
	filePath := filepath.Join(testDataDir, "TestPattern_RGB.dcm")

	data, err := os.ReadFile(filePath)
	if err != nil {
		b.Skipf("Cannot read file: %v", err)
		return
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		result, _ := Parse(reader)
		_, _ = convertDICOMToImage(result)
	}
}
