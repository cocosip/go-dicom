// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main demonstrates how to extract pixel data from DICOM files and save as images.
// This example shows how to:
// - Read DICOM files
// - Extract pixel data using the imaging package
// - Apply windowing (Window Center/Width)
// - Convert to standard image formats (PNG, JPEG)
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/imaging"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

// Command-line flags
var (
	inputFile    = flag.String("input", "", "Input DICOM file path")
	outputFile   = flag.String("output", "", "Output image file path (default: input filename with format extension)")
	windowCenter = flag.Float64("window-center", 0, "Window center for display (0 = use DICOM value or auto)")
	windowWidth  = flag.Float64("window-width", 0, "Window width for display (0 = use DICOM value or auto)")
	invert       = flag.Bool("invert", false, "Invert grayscale (swap black/white)")
	frame        = flag.Int("frame", 0, "Frame number to extract (0-indexed)")
	format       = flag.String("format", "jpeg", "Output format: png or jpeg")
	jpegQuality  = flag.Int("jpeg-quality", 90, "JPEG quality (1-100, only for JPEG format)")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	// Validate input
	if err := examplepath.RequireInputFile(*inputFile); err != nil {
		return err
	}

	// Validate output format before doing any work.
	exportFormat, extension, err := parseFormat(*format)
	if err != nil {
		return err
	}

	// Determine output file path
	output := *outputFile
	if output == "" {
		baseName := strings.TrimSuffix(*inputFile, filepath.Ext(*inputFile))
		output = baseName + extension
	}
	if err := examplepath.PrepareOutputFile(output); err != nil {
		return err
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
		return fmt.Errorf("parse DICOM file: %w", err)
	}

	// Create DICOM image from dataset
	fmt.Println("Extracting pixel data...")
	pixelData, err := pixeldata.FromDataset(result.Dataset)
	if err != nil {
		return fmt.Errorf("extract pixel data: %w", err)
	}

	// Create a Dataset-driven image so DICOM window, LUT, and overlay metadata is applied.
	dicomImage, err := imaging.NewDicomImageFromParseResult(result)
	if err != nil {
		return fmt.Errorf("create DICOM image: %w", err)
	}

	// Validate frame number before using it.
	if *frame < 0 || *frame >= dicomImage.NumberOfFrames() {
		return fmt.Errorf("frame %d out of range (0-%d)", *frame, dicomImage.NumberOfFrames()-1)
	}

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
	frameData, err := pixelData.Frame(context.Background(), *frame)
	if err != nil {
		return fmt.Errorf("read frame %d: %w", *frame, err)
	}
	if len(frameData) > 0 {
		fmt.Printf("\nPixel data sample (first 10 bytes):\n  ")
		for i := 0; i < 10 && i < len(frameData); i++ {
			fmt.Printf("%02X ", frameData[i])
		}
		fmt.Println()
	}
	fmt.Println()

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
		fmt.Printf("Using DICOM or auto-calculated window: Center=%.0f, Width=%.0f\n",
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
		Format:      exportFormat,
		JPEGQuality: *jpegQuality,
	}

	// Create output file
	outFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", output, err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close output file: %v\n", err)
		}
	}()

	// Render frame to file
	fmt.Printf("Rendering frame %d...\n", *frame)
	if err := dicomImage.RenderFrame(outFile, *frame, exportOptions); err != nil {
		return fmt.Errorf("render image: %w", err)
	}

	// Flush the file before checking
	if err := outFile.Sync(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to sync file: %v\n", err)
	}

	fmt.Println("Render complete!")

	// Get file info to verify
	if info, err := os.Stat(output); err == nil {
		fmt.Printf("Image saved to: %s (%d bytes)\n", output, info.Size())
	} else {
		fmt.Fprintf(os.Stderr, "Warning: file was created but stat failed: %v\n", err)
		return fmt.Errorf("verify output file %q: %w", output, err)
	}
	return nil
}

// parseFormat maps the -format flag to a render export format and file
// extension. Matching is case-insensitive; any other value is an error.
func parseFormat(raw string) (render.ExportFormat, string, error) {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "png":
		return render.FormatPNG, ".png", nil
	case "jpeg", "jpg":
		return render.FormatJPEG, ".jpg", nil
	default:
		return 0, "", fmt.Errorf("invalid -format %q: expected png, jpeg, or jpg", raw)
	}
}
