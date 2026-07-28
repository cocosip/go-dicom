// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to anonymize a DICOM file.
package main

import (
	"flag"
	"log"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/anonymizer"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
	inputPath := flag.String("input", "", "Input DICOM file path")
	outputPath := flag.String("output", "anonymized.dcm", "Output anonymized DICOM file path")
	flag.Parse()

	if err := examplepath.RequireInputFile(*inputPath); err != nil {
		log.Fatal(err)
	}
	if err := examplepath.PrepareOutputFile(*outputPath); err != nil {
		log.Fatal(err)
	}

	// Parse the input DICOM file
	result, err := parser.ParseFile(*inputPath)
	if err != nil {
		log.Fatalf("Failed to parse DICOM file: %v", err)
	}

	// Create anonymizer with basic profile
	profile := anonymizer.NewSecurityProfile(anonymizer.BasicProfile)

	// Set custom patient information (optional)
	profile.PatientName = "ANONYMOUS^PATIENT"
	profile.PatientID = "ANON-001"

	anon := anonymizer.NewAnonymizer(profile)

	// Anonymize the dataset
	err = anon.AnonymizeInPlace(result.Dataset)
	if err != nil {
		log.Fatalf("Failed to anonymize dataset: %v", err)
	}

	// Write the anonymized file
	err = writer.WriteFile(*outputPath, result.Dataset)
	if err != nil {
		log.Fatalf("Failed to write anonymized file: %v", err)
	}

	log.Printf("DICOM file anonymized successfully: %s", *outputPath)
	log.Printf("Anonymized %d UIDs", len(anon.ReplacedUIDs))
}
