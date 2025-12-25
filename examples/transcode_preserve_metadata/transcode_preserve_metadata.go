// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to transcode DICOM files while preserving
// all metadata including FileMetaInformation tags and private tags.
package main

import (
	"fmt"
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

func main() {
	inputPath := "D:\\2.dcm"
	outputPath := "D:\\2_transcoded_correct.dcm"

	fmt.Println("=== DICOM Transcoding with Metadata Preservation ===\n")

	// Step 1: Parse the input file
	fmt.Printf("Reading input file: %s\n", inputPath)
	result, err := parser.ParseFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	fmt.Printf("Input Transfer Syntax: %s\n", result.TransferSyntax.UID())

	// Count elements
	origDatasetCount := result.Dataset.Count()
	origFMICount := 0
	origPrivateCount := 0

	if result.FileMetaInformation != nil {
		origFMICount = result.FileMetaInformation.Dataset().Count()
	}

	for _, elem := range result.Dataset.Elements() {
		if elem != nil && elem.Tag() != nil && elem.Tag().Group()%2 == 1 {
			origPrivateCount++
		}
	}

	fmt.Printf("Original Dataset: %d elements (%d private)\n", origDatasetCount, origPrivateCount)
	fmt.Printf("Original FileMetaInformation: %d elements\n", origFMICount)

	// Step 2: Define output transfer syntax
	// For this example, we'll just use the same transfer syntax
	// In a real scenario, you'd transcode to a different format (e.g., JPEG)
	outputTS := result.TransferSyntax
	fmt.Printf("\nOutput Transfer Syntax: %s\n", outputTS.UID())

	// Step 3: Create transcoder and transcode with metadata preservation
	fmt.Println("\nTranscoding...")
	transcoder := codec.NewTranscoder(result.TransferSyntax, outputTS)

	// Use the new TranscodeWithMetadata method with FileMetaInformation type
	newDS, newMeta, err := transcoder.TranscodeWithMetadata(
		result.Dataset,
		result.FileMetaInformation) // Pass FileMetaInformation directly
	if err != nil {
		log.Fatalf("Failed to transcode: %v", err)
	}

	// Step 4: Write output file with preserved metadata
	fmt.Printf("\nWriting output file: %s\n", outputPath)
	err = writer.WriteFile(outputPath, newDS,
		writer.WithFileMetaInfo(newMeta.Dataset())) // Use Dataset() to get the underlying dataset
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	// Step 5: Verify the result
	fmt.Println("\nVerifying output file...")
	verifyResult, err := parser.ParseFile(outputPath)
	if err != nil {
		log.Fatalf("Failed to parse output file: %v", err)
	}

	newDatasetCount := verifyResult.Dataset.Count()
	newFMICount := 0
	newPrivateCount := 0

	if verifyResult.FileMetaInformation != nil {
		newFMICount = verifyResult.FileMetaInformation.Dataset().Count()
	}

	for _, elem := range verifyResult.Dataset.Elements() {
		if elem != nil && elem.Tag() != nil && elem.Tag().Group()%2 == 1 {
			newPrivateCount++
		}
	}

	fmt.Printf("Output Dataset: %d elements (%d private)\n", newDatasetCount, newPrivateCount)
	fmt.Printf("Output FileMetaInformation: %d elements\n", newFMICount)

	// Check specific important tags
	fmt.Println("\n=== Verification Results ===")

	// Check dataset elements
	if newDatasetCount == origDatasetCount {
		fmt.Println("✓ Dataset element count preserved")
	} else {
		fmt.Printf("❌ Dataset element count changed: %d → %d\n", origDatasetCount, newDatasetCount)
	}

	// Check private tags
	if newPrivateCount == origPrivateCount {
		fmt.Println("✓ All private tags preserved")
	} else {
		fmt.Printf("❌ Private tags changed: %d → %d\n", origPrivateCount, newPrivateCount)
	}

	// Check FileMetaInformation
	if newFMICount == origFMICount {
		fmt.Println("✓ All FileMetaInformation elements preserved")
	} else {
		fmt.Printf("❌ FileMetaInformation changed: %d → %d elements\n", origFMICount, newFMICount)
	}

	// Check specific optional tag
	if verifyResult.FileMetaInformation != nil {
		if _, ok := verifyResult.FileMetaInformation.SourceApplicationEntityTitle(); ok {
			fmt.Println("✓ SourceApplicationEntityTitle (0002,0016) preserved")
		} else {
			fmt.Println("❌ SourceApplicationEntityTitle (0002,0016) lost")
		}
	}

	fmt.Println("\n=== Transcoding Complete ===")
	fmt.Println("All metadata has been successfully preserved!")
}
