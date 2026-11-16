// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main demonstrates how to transcode DICOM files while preserving
// all metadata including FileMetaInformation tags and private tags.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transcode"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	inputPath := flag.String("input", "", "Input DICOM file path")
	outputPath := flag.String("output", "transcoded.dcm", "Output transcoded DICOM file path")
	flag.Parse()

	if err := examplepath.RequireInputFile(*inputPath); err != nil {
		return err
	}
	if err := examplepath.PrepareOutputFile(*outputPath); err != nil {
		return err
	}

	fmt.Println("=== DICOM Transcoding with Metadata Preservation ===")

	// Step 1: Parse the input file
	fmt.Printf("Reading input file: %s\n", *inputPath)
	result, err := parser.ParseFile(*inputPath)
	if err != nil {
		return fmt.Errorf("parse input file: %w", err)
	}

	fmt.Printf("Input Transfer Syntax: %s\n", result.TransferSyntax.UID())

	orig := snapshotMetadata(result)
	fmt.Printf("Original Dataset: %d elements (%d private)\n", orig.datasetCount, orig.privateCount)
	fmt.Printf("Original FileMetaInformation: %d elements\n", orig.fmiCount)

	// Step 2: Define output transfer syntax. Native (uncompressed) transcoding
	// only needs a byte-order/VR change, so pick the opposite little-endian VR.
	outputTS := transfer.ImplicitVRLittleEndian
	if result.TransferSyntax.UID().UID() == transfer.ImplicitVRLittleEndian.UID().UID() {
		outputTS = transfer.ExplicitVRLittleEndian
	}
	fmt.Printf("\nOutput Transfer Syntax: %s\n", outputTS.UID())

	// Step 3: Create transcoder and transcode with metadata preservation
	fmt.Println("\nTranscoding...")
	manager, err := transcode.NewManager(codec.GlobalRegistry())
	if err != nil {
		return fmt.Errorf("create transcode manager: %w", err)
	}
	transcoder, err := manager.NewTranscoder(result.TransferSyntax, outputTS)
	if err != nil {
		return fmt.Errorf("create transcoder: %w", err)
	}

	// Use the new TranscodeWithMetadata method with FileMetaInformation type
	newDS, newMeta, err := transcoder.TranscodeWithMetadata(
		context.Background(),
		result.Dataset,
		result.FileMetaInformation) // Pass FileMetaInformation directly
	if err != nil {
		return fmt.Errorf("transcode: %w", err)
	}

	// Step 4: Write output file with preserved metadata
	fmt.Printf("\nWriting output file: %s\n", *outputPath)
	err = writer.WriteFile(*outputPath, newDS,
		writer.WithFileMetaInfo(newMeta.Dataset())) // Use Dataset() to get the underlying dataset
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	// Step 5: Verify the result
	fmt.Println("\nVerifying output file...")
	verifyResult, err := parser.ParseFile(*outputPath)
	if err != nil {
		return fmt.Errorf("parse output file: %w", err)
	}

	out := snapshotMetadata(verifyResult)
	fmt.Printf("Output Dataset: %d elements (%d private)\n", out.datasetCount, out.privateCount)
	fmt.Printf("Output FileMetaInformation: %d elements\n", out.fmiCount)

	fmt.Println("\n=== Verification Results ===")
	failures := verifyPreservation(orig, out)
	if len(failures) > 0 {
		for _, failure := range failures {
			fmt.Printf("❌ %s\n", failure)
		}
		return fmt.Errorf("metadata preservation failed: %d check(s) failed", len(failures))
	}

	fmt.Println("✓ Dataset element count preserved")
	fmt.Println("✓ All private tags preserved")
	fmt.Println("✓ All FileMetaInformation elements preserved")
	if orig.hasSourceAE {
		fmt.Println("✓ SourceApplicationEntityTitle (0002,0016) preserved")
	} else {
		fmt.Println("✓ SourceApplicationEntityTitle (0002,0016) absent from source, nothing to preserve")
	}
	fmt.Println("\n=== Transcoding Complete ===")
	fmt.Println("All metadata has been successfully preserved!")
	return nil
}

// metadataSnapshot captures the metadata counts compared before and after
// transcoding. Group-length elements are excluded because the writer drops them
// by default, which would otherwise look like metadata loss.
type metadataSnapshot struct {
	datasetCount int
	privateCount int
	fmiCount     int
	hasSourceAE  bool
	sourceAE     string
}

func snapshotMetadata(result *parser.ParseResult) metadataSnapshot {
	var snapshot metadataSnapshot
	if result.Dataset != nil {
		snapshot.datasetCount, snapshot.privateCount = countElements(result.Dataset)
	}
	if result.FileMetaInformation != nil {
		snapshot.fmiCount, _ = countElements(result.FileMetaInformation.Dataset())
		snapshot.sourceAE, snapshot.hasSourceAE = result.FileMetaInformation.SourceApplicationEntityTitle()
	}
	return snapshot
}

// verifyPreservation returns one message per metadata check that failed. An
// empty result means everything comparable was preserved.
func verifyPreservation(orig, out metadataSnapshot) []string {
	var failures []string
	if out.datasetCount != orig.datasetCount {
		failures = append(failures, fmt.Sprintf("Dataset element count changed: %d → %d", orig.datasetCount, out.datasetCount))
	}
	if out.privateCount != orig.privateCount {
		failures = append(failures, fmt.Sprintf("Private tags changed: %d → %d", orig.privateCount, out.privateCount))
	}
	if out.fmiCount != orig.fmiCount {
		failures = append(failures, fmt.Sprintf("FileMetaInformation changed: %d → %d elements", orig.fmiCount, out.fmiCount))
	}
	// SourceApplicationEntityTitle (0002,0016) is optional: only require it when
	// the input actually has one.
	if orig.hasSourceAE && (!out.hasSourceAE || out.sourceAE != orig.sourceAE) {
		failures = append(failures, fmt.Sprintf("SourceApplicationEntityTitle (0002,0016) lost or changed: %q → %q", orig.sourceAE, out.sourceAE))
	}
	return failures
}

// countElements returns the total and private element counts, ignoring
// group-length (GGGG,0000) elements.
func countElements(ds *dataset.Dataset) (total, private int) {
	for _, elem := range ds.Elements() {
		if elem == nil || elem.Tag() == nil || elem.Tag().Element() == 0x0000 {
			continue
		}
		total++
		if elem.Tag().Group()%2 == 1 {
			private++
		}
	}
	return total, private
}
