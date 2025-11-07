// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to read a DICOM file.
package main

import (
	"fmt"
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func main() {
	// Parse a DICOM file
	result, err := parser.ParseFile("sample.dcm")
	if err != nil {
		log.Fatalf("Failed to parse DICOM file: %v", err)
	}

	fmt.Println("=== DICOM File Parsed Successfully ===")
	fmt.Printf("Transfer Syntax: %s\n", result.TransferSyntax.UID().Name())
	fmt.Printf("File Format: %s\n", result.Format)
	fmt.Println()

	// Access patient information
	if patientName, ok := result.Dataset.GetString(tag.PatientName); ok {
		fmt.Printf("Patient Name: %s\n", patientName)
	}

	if patientID, ok := result.Dataset.GetString(tag.PatientID); ok {
		fmt.Printf("Patient ID: %s\n", patientID)
	}

	// Access study information
	if studyDescription, ok := result.Dataset.GetString(tag.StudyDescription); ok {
		fmt.Printf("Study Description: %s\n", studyDescription)
	}

	// Access image dimensions
	if rows, err := result.Dataset.GetUInt16(tag.Rows, 0); err == nil {
		if columns, err := result.Dataset.GetUInt16(tag.Columns, 0); err == nil {
			fmt.Printf("Image Size: %d x %d\n", columns, rows)
		}
	}

	// Count total elements
	fmt.Printf("\nTotal elements: %d\n", len(result.Dataset.Elements()))
}
