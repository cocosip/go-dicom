// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser_test

import (
	"fmt"
	"log"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// ExampleParse demonstrates basic DICOM file parsing.
func ExampleParse() {
	file, err := os.Open("testdata/sample.dcm")
	if err != nil {
		log.Fatal(err)
	}
    defer func() { _ = file.Close() }()

	result, err := parser.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Format: %s\n", result.Format)
	fmt.Printf("Transfer Syntax: %s\n", result.TransferSyntax.UID())
}

// ExampleParse_withOptions demonstrates parsing with custom options.
func ExampleParse_withOptions() {
	file, err := os.Open("testdata/large.dcm")
	if err != nil {
		log.Fatal(err)
	}
    defer func() { _ = file.Close() }()

	// Parse with custom options:
	// - Skip large tags (e.g., pixel data) to save memory
	// - Set large object threshold to 128KB
	result, err := parser.Parse(file,
		parser.WithReadOption(parser.SkipLargeTags),
		parser.WithLargeObjectSize(128*1024),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Access patient information (small tags are still read)
	patientName, _ := result.Dataset.GetString(tag.PatientName)
	fmt.Printf("Patient: %s\n", patientName)
}

// ExampleParse_stopEarly demonstrates stopping parsing at a specific tag.
func ExampleParse_stopEarly() {
	file, err := os.Open("testdata/sample.dcm")
	if err != nil {
		log.Fatal(err)
	}
    defer func() { _ = file.Close() }()

	// Stop parsing when we reach pixel data
	result, err := parser.Parse(file,
		parser.WithStopAtTag(tag.PixelData),
	)
	if err != nil {
		log.Fatal(err)
	}

	if result.IsPartial {
		fmt.Println("Parsing stopped early as requested")
	}
}

// ExampleParseFile demonstrates parsing a DICOM file by path.
func ExampleParseFile() {
	result, err := parser.ParseFile("testdata/sample.dcm",
		parser.WithReadOption(parser.ReadAll),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transfer Syntax: %s\n", result.TransferSyntax.UID())
}
