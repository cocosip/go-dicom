// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to transcode DICOM files between different transfer syntaxes.
// This example shows how to:
// - Read DICOM files with any transfer syntax
// - Compress uncompressed pixel data to various formats (JPEG, JPEG-LS, JPEG 2000, RLE)
// - Decompress compressed pixel data
// - Convert between different compression formats
//
// To enable compression/decompression, import codec packages from go-dicom-codec:
//
//	import (
//	    _ "github.com/cocosip/go-dicom-codec/jpeg"     // JPEG codecs
//	    _ "github.com/cocosip/go-dicom-codec/jpegls"   // JPEG-LS codecs
//	    _ "github.com/cocosip/go-dicom-codec/jpeg2000" // JPEG 2000 codecs
//	    _ "github.com/cocosip/go-dicom-codec/rle"      // RLE codec
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

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

// Command-line flags
var (
	inputFile   = flag.String("input", "", "Input DICOM file path (required)")
	outputFile  = flag.String("output", "", "Output DICOM file path (default: input filename with _transcoded suffix)")
	outputTS    = flag.String("ts", "jpeg", "Output transfer syntax: uncompressed, jpeg, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle")
	jpegQuality = flag.Int("jpeg-quality", 90, "JPEG quality for lossy compression (1-100)")
	verbose     = flag.Bool("verbose", true, "Show verbose output")
)

func showUsage() {
	fmt.Println("Error: -input is required")
	fmt.Println("\nUsage:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  # Compress to JPEG Baseline (lossy)")
	fmt.Println("  transcode -input input.dcm -ts jpeg -jpeg-quality 85")
	fmt.Println()
	fmt.Println("  # Compress to JPEG Lossless")
	fmt.Println("  transcode -input input.dcm -ts jpeg-lossless")
	fmt.Println()
	fmt.Println("  # Compress to RLE (lossless)")
	fmt.Println("  transcode -input input.dcm -ts rle")
	fmt.Println()
	fmt.Println("  # Decompress to uncompressed")
	fmt.Println("  transcode -input compressed.dcm -ts uncompressed")
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
	}

	// Select output transfer syntax
	targetTS, codecParams, err := selectTransferSyntax(*outputTS, *jpegQuality)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Supported: uncompressed, jpeg, jpeg-lossless, jpeg-ls, jpeg2000, jpeg2000-lossless, rle")
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
