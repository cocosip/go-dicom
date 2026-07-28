// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to convert DICOM to/from JSON.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/serialization"
)

func main() {
	inputPath := flag.String("input", "", "Input DICOM file path")
	outputPath := flag.String("output", "output.json", "Output JSON file path")
	flag.Parse()

	if err := examplepath.RequireInputFile(*inputPath); err != nil {
		log.Fatal(err)
	}
	if err := examplepath.PrepareOutputFile(*outputPath); err != nil {
		log.Fatal(err)
	}

	// Parse a DICOM file
	result, err := parser.ParseFile(*inputPath)
	if err != nil {
		log.Fatalf("Failed to parse DICOM file: %v", err)
	}

	// Convert to JSON with keywords
	jsonData, err := serialization.ToJSON(result.Dataset,
		serialization.WithWriteTagsAsKeywords(true),
		serialization.WithIndent("  "))

	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}
	if err := os.WriteFile(*outputPath, jsonData, 0644); err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	fmt.Printf("DICOM JSON written to: %s\n", *outputPath)

	// Convert back from JSON
	ds, err := serialization.FromJSON(jsonData)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("\n=== Converted back to Dataset ===\n")
	fmt.Printf("Total elements: %d\n", len(ds.Elements()))
}
