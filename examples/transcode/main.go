// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to transcode DICOM files between different transfer syntaxes.
// This example shows how to:
// - Read DICOM files with any transfer syntax
// - Compress uncompressed pixel data to various formats (JPEG, JPEG-LS, JPEG 2000, RLE Lossless)
// - Decompress compressed pixel data
// - Convert between different compression formats
//
// To enable compression/decompression, import codec packages from go-dicom-codec:
//
//	import (
//	    _ "github.com/cocosip/go-dicom-codec/jpeg/baseline"        // JPEG Baseline
//	    _ "github.com/cocosip/go-dicom-codec/jpeg/lossless"        // JPEG Lossless
//	    _ "github.com/cocosip/go-dicom-codec/jpegls/lossless"      // JPEG-LS Lossless
//	    _ "github.com/cocosip/go-dicom-codec/jpeg2000/lossless"    // JPEG 2000 Lossless
//	)
//
// The blank imports will automatically register the codecs with the global registry.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"

	// Import codec implementations to auto-register them
	_ "github.com/cocosip/go-dicom-codec/jpeg/baseline"
	_ "github.com/cocosip/go-dicom-codec/jpeg/extended"
	_ "github.com/cocosip/go-dicom-codec/jpeg/lossless"
	_ "github.com/cocosip/go-dicom-codec/jpeg/lossless14sv1"
	_ "github.com/cocosip/go-dicom-codec/jpeg2000/lossless"
	_ "github.com/cocosip/go-dicom-codec/jpeg2000/lossy"
	_ "github.com/cocosip/go-dicom-codec/jpegls/lossless"
	_ "github.com/cocosip/go-dicom-codec/jpegls/nearlossless"
)

// Command-line flags
var (
	inputFile   = flag.String("input", "D:\\1.dcm", "Input DICOM file path (required)")
	outputFile  = flag.String("output", "output.dcm", "Output DICOM file path (default: input filename with _transcoded suffix)")
	outputTS    = flag.String("ts", "jpeg", "Output transfer syntax: uncompressed, jpeg, jpeg-extended, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle")
	jpegQuality = flag.Int("jpeg-quality", 90, "JPEG quality for lossy compression (1-100)")
	verbose     = flag.Bool("verbose", true, "Show verbose output")
	allFormats  = flag.Bool("all", false, "Generate output files for all compatible formats")
)

func showUsage() {
	fmt.Println("Error: -input is required")
	fmt.Println("\nUsage:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  # Compress to JPEG Baseline (lossy)")
	fmt.Println("  transcode -input input.dcm -ts jpeg -jpeg-quality 85")
	fmt.Println()
	fmt.Println("  # Compress to JPEG Extended (lossy, 12-bit)")
	fmt.Println("  transcode -input input.dcm -ts jpeg-extended -jpeg-quality 85")
	fmt.Println()
	fmt.Println("  # Compress to JPEG Lossless")
	fmt.Println("  transcode -input input.dcm -ts jpeg-lossless")
	fmt.Println()
	fmt.Println("  # Compress to JPEG 2000 Lossless")
	fmt.Println("  transcode -input input.dcm -ts jpeg2000-lossless")
	fmt.Println()
	fmt.Println("  # Compress to RLE Lossless")
	fmt.Println("  transcode -input input.dcm -ts rle")
	fmt.Println()
	fmt.Println("  # Decompress to uncompressed")
	fmt.Println("  transcode -input compressed.dcm -ts uncompressed")
	fmt.Println()
	fmt.Println("  # Generate all compatible formats")
	fmt.Println("  transcode -input input.dcm -all")
	fmt.Println("\nPress Enter to exit...")
	_, _ = fmt.Scanln()
}

func selectTransferSyntax(tsName string, quality int) (*transfer.Syntax, codec.Parameters, error) {
	var targetTS *transfer.Syntax
	var codecParams codec.Parameters

	switch strings.ToLower(tsName) {
	case "uncompressed", "explicit", "explicit-vr":
		targetTS = transfer.ExplicitVRLittleEndian
	case "implicit", "implicit-vr":
		targetTS = transfer.ImplicitVRLittleEndian
	case "jpeg", "jpeg-baseline":
		targetTS = transfer.JPEGBaseline8Bit
		params := codec.NewBaseParameters()
		params.SetParameter("Quality", quality)
		codecParams = params
	case "jpeg-extended", "jpeg-12bit":
		targetTS = transfer.JPEGExtended12Bit
		params := codec.NewBaseParameters()
		params.SetParameter("Quality", quality)
		codecParams = params
	case "jpeg-lossless":
		targetTS = transfer.JPEGLossless
	case "jpeg-ls", "jpegls":
		targetTS = transfer.JPEGLSLossless
	case "jpeg-ls-lossy", "jpegls-lossy":
		targetTS = transfer.JPEGLSNearLossless
	case "jpeg2000", "j2k":
		targetTS = transfer.JPEG2000
	case "jpeg2000-lossless", "j2k-lossless":
		targetTS = transfer.JPEG2000Lossless
	case "rle":
		targetTS = transfer.RLELossless
	default:
		return nil, nil, fmt.Errorf("unknown transfer syntax '%s'", tsName)
	}

	return targetTS, codecParams, nil
}

func printTransferSyntaxInfo(ts *transfer.Syntax, label string) {
	fmt.Printf("%s: %s (%s)\n", label, ts.UID().Name(), ts.UID().UID())
	fmt.Printf("  Explicit VR: %v\n", ts.IsExplicitVR())
	fmt.Printf("  Encapsulated: %v\n", ts.IsEncapsulated())
	if ts.IsLossy() {
		fmt.Printf("  Lossy: %v (method: %s)\n", ts.IsLossy(), ts.LossyCompressionMethod())
	} else {
		fmt.Printf("  Lossy: false\n")
	}
	fmt.Println()
}

func printImageInfo(ds *dataset.Dataset) {
	if !ds.Contains(tag.PixelData) {
		return
	}

	rows := ds.TryGetUInt16(tag.Rows, 0)
	columns := ds.TryGetUInt16(tag.Columns, 0)
	bitsAlloc := ds.TryGetUInt16(tag.BitsAllocated, 0)
	bitsStored := ds.TryGetUInt16(tag.BitsStored, 0)
	samplesPerPixel := ds.TryGetUInt16(tag.SamplesPerPixel, 0)

	frames := int32(1)
	if val, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		frames = val
	}

	fmt.Printf("Image Properties:\n")
	fmt.Printf("  Dimensions: %dx%d\n", columns, rows)
	fmt.Printf("  Frames: %d\n", frames)
	fmt.Printf("  Bits Allocated: %d\n", bitsAlloc)
	fmt.Printf("  Bits Stored: %d\n", bitsStored)
	fmt.Printf("  Samples Per Pixel: %d\n", samplesPerPixel)
	fmt.Println()
}

func validateTranscodeCompatibility(ds *dataset.Dataset, targetTS *transfer.Syntax) error {
	// Check if dataset contains pixel data
	if !ds.Contains(tag.PixelData) {
		return nil // No pixel data, no compatibility issues
	}

	// Get bits allocated
	bitsAllocated := ds.TryGetUInt16(tag.BitsAllocated, 0)
	if bitsAllocated == 0 {
		return nil // Can't determine, let transcoder handle it
	}

	// Check JPEG Baseline constraints
	if targetTS.UID().UID() == transfer.JPEGBaseline8Bit.UID().UID() {
		if bitsAllocated != 8 {
			return fmt.Errorf("jpeg Baseline only supports 8-bit images, your image has %d bits allocated, auto-switching to compatible format", bitsAllocated)
		}
	}

	// Check JPEG Extended constraints
	if targetTS.UID().UID() == transfer.JPEGExtended12Bit.UID().UID() {
		if bitsAllocated > 12 {
			return fmt.Errorf("jpeg Extended only supports up to 12-bit images, your image has %d bits allocated, auto-switching to compatible format", bitsAllocated)
		}
	}

	return nil
}

func findCompatibleAlternative(ds *dataset.Dataset, requestedTS *transfer.Syntax) *transfer.Syntax {
	// Get bits allocated
	bitsAllocated := ds.TryGetUInt16(tag.BitsAllocated, 0)
	if bitsAllocated == 0 {
		return nil
	}

	// Determine if lossy or lossless was requested
	wantedLossy := requestedTS.IsLossy()

	// Map from requested format to compatible alternative
	requestedUID := requestedTS.UID().UID()

	switch requestedUID {
	case transfer.JPEGBaseline8Bit.UID().UID():
		// User wanted JPEG Baseline (lossy), but image is not 8-bit
		if wantedLossy {
			// Try JPEG Extended for 12-bit
			if bitsAllocated <= 12 {
				fmt.Println("  → Switching to JPEG Extended (supports up to 12-bit, lossy)")
				return transfer.JPEGExtended12Bit
			}
			// For >12-bit, use JPEG 2000 lossy
			fmt.Println("  → Switching to JPEG 2000 Lossy (supports all bit depths)")
			return transfer.JPEG2000
		}
		// User wanted lossy but we'll suggest lossless for safety
		fmt.Println("  → Switching to JPEG Lossless (supports all bit depths, lossless)")
		return transfer.JPEGLossless

	case transfer.JPEGExtended12Bit.UID().UID():
		// User wanted JPEG Extended, but image is >12-bit
		if wantedLossy {
			fmt.Println("  → Switching to JPEG 2000 Lossy (supports all bit depths)")
			return transfer.JPEG2000
		}
		fmt.Println("  → Switching to JPEG Lossless (supports all bit depths, lossless)")
		return transfer.JPEGLossless

	default:
		// For other incompatible formats, suggest JPEG Lossless as safe default
		return transfer.JPEGLossless
	}
}

type formatInfo struct {
	name         string
	transferSyntax *transfer.Syntax
	params       codec.Parameters
	outputFile   string
}

func generateAllFormats(result *parser.ParseResult, inputPath string, quality int, _ bool) {
	fmt.Println("\n=== Generating All Compatible Formats ===")
	fmt.Println()

	// Get image properties
	bitsAllocated := result.Dataset.TryGetUInt16(tag.BitsAllocated, 0)
	if bitsAllocated == 0 {
		bitsAllocated = 8 // Default assumption
	}

	// Determine compatible formats based on bit depth
	formats := getCompatibleFormats(inputPath, bitsAllocated, quality)

	if len(formats) == 0 {
		fmt.Println("No compatible formats found for this image.")
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	fmt.Printf("Found %d compatible formats for %d-bit image:\n\n", len(formats), bitsAllocated)

	// Transcode to each format
	var results []transcodeResult
	inputInfo, _ := os.Stat(inputPath)

	for i, format := range formats {
		fmt.Printf("[%d/%d] Processing %s...\n", i+1, len(formats), format.name)

		// Create transcoder
		var transcoderOpts []codec.TranscoderOption
		if format.params != nil {
			transcoderOpts = append(transcoderOpts, codec.WithOutputParameters(format.params))
		}

		transcoder := codec.NewTranscoder(
			result.TransferSyntax,
			format.transferSyntax,
			transcoderOpts...,
		)

		// Transcode
		newDataset, err := transcoder.Transcode(result.Dataset)
		if err != nil {
			fmt.Printf("  ❌ Failed: %v\n", err)
			continue
		}

		// Write output
		err = writer.WriteFile(format.outputFile, newDataset, writer.WithTransferSyntax(format.transferSyntax))
		if err != nil {
			fmt.Printf("  ❌ Write failed: %v\n", err)
			continue
		}

		// Get file info
		outputInfo, err := os.Stat(format.outputFile)
		if err != nil {
			fmt.Printf("  ❌ Stat failed: %v\n", err)
			continue
		}

		ratio := float64(inputInfo.Size()) / float64(outputInfo.Size())
		reduction := (1.0 - 1.0/ratio) * 100

		results = append(results, transcodeResult{
			format:     format.name,
			outputPath: format.outputFile,
			inputSize:  inputInfo.Size(),
			outputSize: outputInfo.Size(),
			ratio:      ratio,
			reduction:  reduction,
		})

		fmt.Printf("  ✅ Success: %s (%d bytes, %.1f%% reduction)\n",
			format.outputFile, outputInfo.Size(), reduction)
	}

	// Print summary table
	printComparisonTable(inputPath, results)

	fmt.Println("\nPress Enter to exit...")
	_, _ = fmt.Scanln()
}

func getCompatibleFormats(inputPath string, bitsAllocated uint16, quality int) []formatInfo {
	// Extract base name from input file path
	ext := filepath.Ext(inputPath)
	baseName := strings.TrimSuffix(filepath.Base(inputPath), ext)
	baseDir := filepath.Dir(inputPath)

	var formats []formatInfo

	// Uncompressed - always compatible
	formats = append(formats, formatInfo{
		name:           "Uncompressed (Explicit VR)",
		transferSyntax: transfer.ExplicitVRLittleEndian,
		outputFile:     filepath.Join(baseDir, baseName+"_uncompressed.dcm"),
	})

	// RLE - always compatible
	formats = append(formats, formatInfo{
		name:           "RLE Lossless",
		transferSyntax: transfer.RLELossless,
		outputFile:     filepath.Join(baseDir, baseName+"_rle.dcm"),
	})

	// JPEG Baseline - only for 8-bit
	if bitsAllocated == 8 {
		params := codec.NewBaseParameters()
		params.SetParameter("Quality", quality)
		formats = append(formats, formatInfo{
			name:           "JPEG Baseline (Lossy)",
			transferSyntax: transfer.JPEGBaseline8Bit,
			params:         params,
			outputFile:     filepath.Join(baseDir, baseName+"_jpeg_baseline.dcm"),
		})
	}

	// JPEG Extended - for 8-bit and 12-bit
	if bitsAllocated <= 12 {
		params := codec.NewBaseParameters()
		params.SetParameter("Quality", quality)
		formats = append(formats, formatInfo{
			name:           "JPEG Extended (Lossy)",
			transferSyntax: transfer.JPEGExtended12Bit,
			params:         params,
			outputFile:     filepath.Join(baseDir, baseName+"_jpeg_extended.dcm"),
		})
	}

	// JPEG Lossless - always compatible
	formats = append(formats, formatInfo{
		name:           "JPEG Lossless",
		transferSyntax: transfer.JPEGLossless,
		outputFile:     filepath.Join(baseDir, baseName+"_jpeg_lossless.dcm"),
	})

	// JPEG-LS Lossless - always compatible
	formats = append(formats, formatInfo{
		name:           "JPEG-LS Lossless",
		transferSyntax: transfer.JPEGLSLossless,
		outputFile:     filepath.Join(baseDir, baseName+"_jpegls_lossless.dcm"),
	})

	// JPEG 2000 Lossless - always compatible
	formats = append(formats, formatInfo{
		name:           "JPEG 2000 Lossless",
		transferSyntax: transfer.JPEG2000Lossless,
		outputFile:     filepath.Join(baseDir, baseName+"_jpeg2000_lossless.dcm"),
	})

	// JPEG 2000 Lossy - always compatible
	formats = append(formats, formatInfo{
		name:           "JPEG 2000 Lossy",
		transferSyntax: transfer.JPEG2000,
		outputFile:     filepath.Join(baseDir, baseName+"_jpeg2000_lossy.dcm"),
	})

	return formats
}

type transcodeResult struct {
	format     string
	outputPath string
	inputSize  int64
	outputSize int64
	ratio      float64
	reduction  float64
}

func printComparisonTable(inputPath string, results []transcodeResult) {
	if len(results) == 0 {
		return
	}

	fmt.Println()
	fmt.Println("=== Compression Comparison ===")
	fmt.Println()
	fmt.Printf("Input File: %s\n\n", inputPath)

	// Print table header
	fmt.Printf("%-30s %-20s %-15s %-15s\n", "Format", "Output File", "Size", "Compression")
	fmt.Println(strings.Repeat("-", 80))

	// Print each result
	for _, r := range results {
		sizeStr := formatSize(r.outputSize)
		var compressionStr string
		if r.ratio > 1.0 {
			compressionStr = fmt.Sprintf("%.2f:1 (%.1f%%↓)", r.ratio, r.reduction)
		} else if r.ratio < 1.0 {
			compressionStr = fmt.Sprintf("%.1f%%↑", (1.0/r.ratio-1.0)*100)
		} else {
			compressionStr = "same"
		}

		fmt.Printf("%-30s %-20s %-15s %-15s\n",
			r.format,
			filepath.Base(r.outputPath),
			sizeStr,
			compressionStr)
	}

	// Find best compression
	var bestRatio transcodeResult
	for _, r := range results {
		if r.ratio > bestRatio.ratio {
			bestRatio = r
		}
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("\nBest Compression: %s (%.2f:1, %.1f%% reduction)\n",
		bestRatio.format, bestRatio.ratio, bestRatio.reduction)
	fmt.Printf("Input Size:       %s\n", formatSize(results[0].inputSize))
	fmt.Printf("Smallest Output:  %s\n", formatSize(bestRatio.outputSize))
}

func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(size)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
}

func printSummary(inputPath, outputPath string) {
	info, err := os.Stat(outputPath)
	if err != nil {
		fmt.Printf("Warning: file was created but stat failed: %v\n", err)
		fmt.Printf("Expected location: %s\n", outputPath)
		return
	}

	inputInfo, _ := os.Stat(inputPath)
	fmt.Println()
	fmt.Printf("=== Summary ===\n")
	fmt.Printf("Input file:  %s (%d bytes)\n", inputPath, inputInfo.Size())
	fmt.Printf("Output file: %s (%d bytes)\n", outputPath, info.Size())

	compressionRatio := float64(inputInfo.Size()) / float64(info.Size())
	if info.Size() < inputInfo.Size() {
		fmt.Printf("Compression ratio: %.2f:1 (%.1f%% smaller)\n",
			compressionRatio,
			(1.0-1.0/compressionRatio)*100)
	} else if info.Size() > inputInfo.Size() {
		fmt.Printf("File size increased by %.1f%%\n",
			(float64(info.Size())/float64(inputInfo.Size())-1.0)*100)
	} else {
		fmt.Println("File size unchanged")
	}
}

func main() {
	flag.Parse()

	// Validate input
	if *inputFile == "" {
		showUsage()
		return
	}

	// Determine output file path
	output := *outputFile
	if output == "" {
		ext := filepath.Ext(*inputFile)
		baseName := strings.TrimSuffix(*inputFile, ext)
		output = baseName + "_transcoded" + ext
	}

	if *verbose {
		fmt.Printf("=== DICOM Transcoder ===\n")
		fmt.Printf("Input:  %s\n", *inputFile)
		fmt.Printf("Output: %s\n", output)
		fmt.Printf("Target Transfer Syntax: %s\n", *outputTS)
		fmt.Println()
	}

	// Parse input DICOM file
	if *verbose {
		fmt.Println("Parsing input DICOM file...")
	}
	result, err := parser.ParseFile(*inputFile)
	if err != nil {
		fmt.Printf("Failed to parse DICOM file: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	if *verbose {
		printTransferSyntaxInfo(result.TransferSyntax, "Input Transfer Syntax")
		printImageInfo(result.Dataset)
	}

	// Check if user wants all compatible formats
	if *allFormats {
		generateAllFormats(result, *inputFile, *jpegQuality, *verbose)
		return
	}

	// Select output transfer syntax
	targetTS, codecParams, err := selectTransferSyntax(*outputTS, *jpegQuality)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Supported: uncompressed, jpeg, jpeg-extended, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle")
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	if *verbose {
		printTransferSyntaxInfo(targetTS, "Output Transfer Syntax")
	}

	// Check if transcoding is needed
	if result.TransferSyntax.UID().UID() == targetTS.UID().UID() {
		fmt.Println("Input and output transfer syntaxes are the same. No transcoding needed.")
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	// Validate compatibility and auto-adjust if needed
	if err := validateTranscodeCompatibility(result.Dataset, targetTS); err != nil {
		fmt.Printf("\nCompatibility Issue: %v\n", err)

		// Try to find a compatible alternative
		alternative := findCompatibleAlternative(result.Dataset, targetTS)
		if alternative != nil {
			fmt.Printf("\nAuto-switching to compatible format: %s\n", alternative.UID().Name())
			targetTS = alternative
			// Update codec parameters if needed
			if targetTS.UID().UID() == transfer.JPEGBaseline8Bit.UID().UID() {
				params := codec.NewBaseParameters()
				params.SetParameter("Quality", *jpegQuality)
				codecParams = params
			}
		} else {
			fmt.Println("\nNo compatible alternative found.")
			fmt.Println("\nPress Enter to exit...")
			_, _ = fmt.Scanln()
			return
		}
	}

	// Create transcoder
	if *verbose {
		fmt.Println("Creating transcoder...")
	}

	var transcoderOpts []codec.TranscoderOption
	if codecParams != nil {
		transcoderOpts = append(transcoderOpts, codec.WithOutputParameters(codecParams))
	}

	transcoder := codec.NewTranscoder(
		result.TransferSyntax,
		targetTS,
		transcoderOpts...,
	)

	// Perform transcoding
	if *verbose {
		fmt.Println("Transcoding dataset...")
	}
	newDataset, err := transcoder.Transcode(result.Dataset)
	if err != nil {
		fmt.Printf("Failed to transcode: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	if *verbose {
		fmt.Println("Transcoding complete!")
		fmt.Println()
	}

	// Write output file
	if *verbose {
		fmt.Printf("Writing output file: %s\n", output)
	}

	// Write DICOM file with new transfer syntax
	err = writer.WriteFile(output, newDataset, writer.WithTransferSyntax(targetTS))
	if err != nil {
		fmt.Printf("Failed to write DICOM file: %v\n", err)
		fmt.Println("\nPress Enter to exit...")
		_, _ = fmt.Scanln()
		return
	}

	if *verbose {
		fmt.Println("Write complete!")
	}

	// Print summary
	printSummary(*inputFile, output)

	fmt.Println("\nPress Enter to exit...")
	_, _ = fmt.Scanln()
}
