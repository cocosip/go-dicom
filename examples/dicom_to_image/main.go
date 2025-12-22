// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to extract pixel data from DICOM files and save as images.
// This example shows how to:
// - Read DICOM files
// - Extract pixel data using the imaging package
// - Apply windowing (Window Center/Width)
// - Convert to standard image formats (PNG, JPEG)
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/imaging"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

// Command-line flags
var (
	inputFile    = flag.String("input", "input.dcm", "Input DICOM file path (required)")
	outputFile   = flag.String("output", "output.jpeg", "Output image file path (default: input filename with .jpeg extension)")
	windowCenter = flag.Float64("window-center", 0, "Window center for display (0 = use DICOM value or auto)")
	windowWidth  = flag.Float64("window-width", 0, "Window width for display (0 = use DICOM value or auto)")
	invert       = flag.Bool("invert", false, "Invert grayscale (swap black/white)")
	frame        = flag.Int("frame", 0, "Frame number to extract (0-indexed)")
	format       = flag.String("format", "jpeg", "Output format: png or jpeg")
	jpegQuality  = flag.Int("jpeg-quality", 90, "JPEG quality (1-100, only for JPEG format)")
)

func main() {
	flag.Parse()

	// Validate input
	if *inputFile == "" {
		fmt.Println("Error: -input is required")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Determine output file path
	output := *outputFile
	if output == "" {
		ext := filepath.Ext(*inputFile)
		baseName := strings.TrimSuffix(*inputFile, ext)
		if *format == "jpeg" || *format == "jpg" {
			output = baseName + ".jpg"
		} else {
			output = baseName + ".png"
		}
	}

	fmt.Printf("=== DICOM to Image Converter ===\n")
	fmt.Printf("Input:  %s\n", *inputFile)
	fmt.Printf("Output: %s\n", output)
	fmt.Printf("Frame:  %d\n", *frame)
	fmt.Println()

	// Parse DICOM file
	fmt.Println("Parsing DICOM file...")
	result, err := parser.ParseFile(*inputFile)
	if err != nil {
		fmt.Printf("Failed to parse DICOM file: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Create DICOM image from dataset
	fmt.Println("Extracting pixel data...")
	pixelData, err := imaging.CreatePixelData(result.Dataset)
	if err != nil {
		fmt.Printf("Failed to extract pixel data: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Create DICOM image (window is auto-calculated from pixel data)
	dicomImage := imaging.NewDicomImage(pixelData)

	// Print image information
	fmt.Printf("\nImage properties:\n")
	fmt.Printf("  Dimensions: %dx%d\n", dicomImage.Width(), dicomImage.Height())
	fmt.Printf("  Number of frames: %d\n", dicomImage.NumberOfFrames())
	fmt.Printf("  Grayscale: %v\n", dicomImage.IsGrayscale())
	fmt.Printf("  Bits Allocated: %d\n", pixelData.Info.BitsAllocated)
	fmt.Printf("  Bits Stored: %d\n", pixelData.Info.BitsStored)
	fmt.Printf("  Pixel Representation: %d (0=unsigned, 1=signed)\n", pixelData.Info.PixelRepresentation)
	fmt.Printf("  Samples Per Pixel: %d\n", pixelData.Info.SamplesPerPixel)
	if pixelData.Info.PhotometricInterpretation != nil {
		fmt.Printf("  Photometric Interpretation: %s\n", pixelData.Info.PhotometricInterpretation.Value)
	}

	// Print some pixel values for debugging
	frameData, err := pixelData.GetFrame(*frame)
	if err == nil && len(frameData) > 0 {
		fmt.Printf("\nPixel data sample (first 10 bytes):\n  ")
		for i := 0; i < 10 && i < len(frameData); i++ {
			fmt.Printf("%02X ", frameData[i])
		}
		fmt.Println()
	}
	fmt.Println()

	// Validate frame number
	if *frame < 0 || *frame >= dicomImage.NumberOfFrames() {
		fmt.Printf("Error: Frame %d out of range (0-%d)\n", *frame, dicomImage.NumberOfFrames()-1)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Apply window settings if provided by user
	if *windowCenter != 0 || *windowWidth != 0 {
		wc := *windowCenter
		ww := *windowWidth

		// If only one is specified, get the other from auto-calculated values
		if wc == 0 {
			wc = dicomImage.WindowCenter()
		}
		if ww == 0 {
			ww = dicomImage.WindowWidth()
		}

		dicomImage.SetWindow(wc, ww)
		fmt.Printf("Using custom window: Center=%.0f, Width=%.0f\n", wc, ww)
	} else {
		// Window is auto-calculated from pixel data in DicomImage
		fmt.Printf("Using auto-calculated window: Center=%.0f, Width=%.0f\n",
			dicomImage.WindowCenter(), dicomImage.WindowWidth())
	}

	// Apply invert if requested
	if *invert {
		dicomImage.SetInvert(true)
		fmt.Println("Grayscale inversion enabled")
	}

	fmt.Println()

	// Create export options
	exportOptions := &render.ExportOptions{
		JPEGQuality: *jpegQuality,
	}

	if *format == "jpeg" || *format == "jpg" {
		exportOptions.Format = render.FormatJPEG
		fmt.Println("Exporting as JPEG...")
	} else {
		exportOptions.Format = render.FormatPNG
		fmt.Println("Exporting as PNG...")
	}

	// Create output file
	outFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Printf("Warning: failed to close output file: %v\n", err)
		}
	}()

	// Render frame to file
	fmt.Printf("Rendering frame %d...\n", *frame)
	if err := dicomImage.RenderFrame(outFile, *frame, exportOptions); err != nil {
		fmt.Printf("Failed to render image: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Flush the file before checking
	if err := outFile.Sync(); err != nil {
		fmt.Printf("Warning: failed to sync file: %v\n", err)
	}

	fmt.Println("Render complete!")

	// Get file info to verify
	if info, err := os.Stat(output); err == nil {
		fmt.Printf("Image saved to: %s (%d bytes)\n", output, info.Size())
	} else {
		fmt.Printf("Warning: file was created but stat failed: %v\n", err)
		fmt.Printf("Expected location: %s\n", output)
	}
	fmt.Println("\nPress Enter to exit...")
	_, _ = fmt.Scanln()
}
