// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to convert DICOM to/from JSON.
package main

import (
	"fmt"
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/serialization"
)

func main() {
	// Parse a DICOM file
	result, err := parser.ParseFile("sample.dcm")
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

	fmt.Println("=== DICOM JSON Output ===")
	fmt.Println(string(jsonData))

	// Convert back from JSON
	ds, err := serialization.FromJSON(jsonData)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("\n=== Converted back to Dataset ===\n")
	fmt.Printf("Total elements: %d\n", len(ds.Elements()))
}
