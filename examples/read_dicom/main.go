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
	result, err := parser.ParseFile("D:\\2.dcm")
	if err != nil {
		log.Fatalf("Failed to parse DICOM file: %v", err)
	}

	fmt.Println("=== DICOM File Parsed Successfully ===")
	fmt.Printf("Transfer Syntax: %s\n", result.TransferSyntax.UID().Name())
	fmt.Printf("File Format: %s\n", result.Format)
	fmt.Println()

	// === Patient Information ===
	fmt.Println("=== Patient Information ===")
	if patientName, ok := result.Dataset.GetString(tag.PatientName); ok {
		fmt.Printf("Patient Name: %s\n", patientName)
	}

	if patientID, ok := result.Dataset.GetString(tag.PatientID); ok {
		fmt.Printf("Patient ID: %s\n", patientID)
	}

	if patientBirthDate, ok := result.Dataset.GetString(tag.PatientBirthDate); ok {
		fmt.Printf("Patient Birth Date: %s\n", patientBirthDate)
	}

	if patientSex, ok := result.Dataset.GetString(tag.PatientSex); ok {
		fmt.Printf("Patient Sex: %s\n", patientSex)
	}

	if patientAge, ok := result.Dataset.GetString(tag.PatientAge); ok {
		fmt.Printf("Patient Age: %s\n", patientAge)
	}

	if patientWeight, ok := result.Dataset.GetString(tag.PatientWeight); ok {
		fmt.Printf("Patient Weight: %s kg\n", patientWeight)
	}

	if patientSize, ok := result.Dataset.GetString(tag.PatientSize); ok {
		fmt.Printf("Patient Size: %s m\n", patientSize)
	}

	// === Study Information ===
	fmt.Println("\n=== Study Information ===")
	if studyInstanceUID, ok := result.Dataset.GetString(tag.StudyInstanceUID); ok {
		fmt.Printf("Study Instance UID: %s\n", studyInstanceUID)
	}

	if studyDate, ok := result.Dataset.GetString(tag.StudyDate); ok {
		fmt.Printf("Study Date: %s\n", studyDate)
	}

	if studyTime, ok := result.Dataset.GetString(tag.StudyTime); ok {
		fmt.Printf("Study Time: %s\n", studyTime)
	}

	if studyDescription, ok := result.Dataset.GetString(tag.StudyDescription); ok {
		fmt.Printf("Study Description: %s\n", studyDescription)
	}

	if studyID, ok := result.Dataset.GetString(tag.StudyID); ok {
		fmt.Printf("Study ID: %s\n", studyID)
	}

	if accessionNumber, ok := result.Dataset.GetString(tag.AccessionNumber); ok {
		fmt.Printf("Accession Number: %s\n", accessionNumber)
	}

	if referringPhysician, ok := result.Dataset.GetString(tag.ReferringPhysicianName); ok {
		fmt.Printf("Referring Physician: %s\n", referringPhysician)
	}

	// === Series Information ===
	fmt.Println("\n=== Series Information ===")
	if seriesInstanceUID, ok := result.Dataset.GetString(tag.SeriesInstanceUID); ok {
		fmt.Printf("Series Instance UID: %s\n", seriesInstanceUID)
	}

	if seriesNumber, err := result.Dataset.GetInt32(tag.SeriesNumber, 0); err == nil {
		fmt.Printf("Series Number: %d\n", seriesNumber)
	}

	if seriesDescription, ok := result.Dataset.GetString(tag.SeriesDescription); ok {
		fmt.Printf("Series Description: %s\n", seriesDescription)
	}

	if modality, ok := result.Dataset.GetString(tag.Modality); ok {
		fmt.Printf("Modality: %s\n", modality)
	}

	if seriesDate, ok := result.Dataset.GetString(tag.SeriesDate); ok {
		fmt.Printf("Series Date: %s\n", seriesDate)
	}

	if seriesTime, ok := result.Dataset.GetString(tag.SeriesTime); ok {
		fmt.Printf("Series Time: %s\n", seriesTime)
	}

	// === Image Information ===
	fmt.Println("\n=== Image Information ===")
	if sopInstanceUID, ok := result.Dataset.GetString(tag.SOPInstanceUID); ok {
		fmt.Printf("SOP Instance UID: %s\n", sopInstanceUID)
	}

	if instanceNumber, err := result.Dataset.GetInt32(tag.InstanceNumber, 0); err == nil {
		fmt.Printf("Instance Number: %d\n", instanceNumber)
	}

	if imageType, ok := result.Dataset.GetString(tag.ImageType); ok {
		fmt.Printf("Image Type: %s\n", imageType)
	}

	// Access image dimensions
	if rows, err := result.Dataset.GetUInt16(tag.Rows, 0); err == nil {
		if columns, err := result.Dataset.GetUInt16(tag.Columns, 0); err == nil {
			fmt.Printf("Image Size: %d x %d\n", columns, rows)
		}
	}

	if bitsAllocated, err := result.Dataset.GetUInt16(tag.BitsAllocated, 0); err == nil {
		fmt.Printf("Bits Allocated: %d\n", bitsAllocated)
	}

	if bitsStored, err := result.Dataset.GetUInt16(tag.BitsStored, 0); err == nil {
		fmt.Printf("Bits Stored: %d\n", bitsStored)
	}

	if samplesPerPixel, err := result.Dataset.GetUInt16(tag.SamplesPerPixel, 0); err == nil {
		fmt.Printf("Samples Per Pixel: %d\n", samplesPerPixel)
	}

	// === Summary ===
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total elements: %d\n", len(result.Dataset.Elements()))
}
