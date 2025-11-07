// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to anonymize a DICOM file.
package main

import (
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/anonymizer"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
	// Parse the input DICOM file
	result, err := parser.ParseFile("input.dcm")
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
	err = writer.WriteFile("anonymized.dcm", result.Dataset)
	if err != nil {
		log.Fatalf("Failed to write anonymized file: %v", err)
	}

	log.Println("DICOM file anonymized successfully")
	log.Printf("Anonymized %d UIDs", len(anon.ReplacedUIDs))
}
