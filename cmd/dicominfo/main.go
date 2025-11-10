// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main provides a command-line tool for displaying information about DICOM files.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

const version = "0.1.0"

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help information")
	verbose := flag.Bool("v", false, "Verbose output")

	flag.Parse()

	// Show version
	if *showVersion {
		fmt.Printf("dicominfo version %s\n", version)
		os.Exit(0)
	}

	// Show help
	if *showHelp || flag.NArg() == 0 {
		printUsage()
		os.Exit(0)
	}

	// Get input file
	inputFile := flag.Arg(0)

	// Parse DICOM file
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse DICOM file: %v\n", err)
		os.Exit(1)
	}

	// Print file information
	printFileInfo(result, *verbose)
}

func printUsage() {
	fmt.Println("Usage: dicominfo [options] <dicom-file>")
	fmt.Println()
	fmt.Println("Display information about a DICOM file")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -version    Show version information")
	fmt.Println("  -help       Show this help message")
	fmt.Println("  -v          Verbose output")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  dicominfo image.dcm")
	fmt.Println("  dicominfo -v image.dcm")
}

func printFileInfo(result *parser.ParseResult, verbose bool) {
	fmt.Println("=== DICOM File Information ===")
	fmt.Println()

	// File format
	fmt.Printf("File Format: %s\n", formatFileFormat(result.Format))
	fmt.Printf("Transfer Syntax: %s (%s)\n",
		result.TransferSyntax.UID().Name(),
		result.TransferSyntax.UID().UID())

	// File meta information
	if result.FileMetaInformation != nil {
		fmt.Println()
		fmt.Println("--- File Meta Information ---")
		printFileMetaInfo(result.FileMetaInformation, verbose)
	}

	// Main dataset information
	fmt.Println()
	fmt.Println("--- Main Dataset ---")
	printMainDatasetInfo(result.Dataset, verbose)

	// Statistics
	fmt.Println()
	fmt.Println("--- Statistics ---")
	fmt.Printf("Total Elements: %d\n", len(result.Dataset.Elements()))
	if result.IsPartial {
		fmt.Println("Note: File was partially read (some large elements may be skipped)")
	}
}

func formatFileFormat(format parser.FileFormat) string {
	switch format {
	case parser.FormatDICOM3:
		return "DICOM Part 10 (with preamble)"
	case parser.FormatDICOM3NoPreamble:
		return "DICOM Part 10 (no preamble)"
	case parser.FormatDICOM3NoFileMetaInfo:
		return "DICOM Part 10 (no file meta info)"
	default:
		return "Unknown"
	}
}

func printFileMetaInfo(fmi *dataset.FileMetaInformation, verbose bool) {
	// Media Storage SOP Class UID
	if sopClassUID, ok := fmi.MediaStorageSOPClassUID(); ok {
		fmt.Printf("  Media Storage SOP Class UID: %s\n", sopClassUID)
	}

	// Media Storage SOP Instance UID
	if sopInstanceUID, ok := fmi.MediaStorageSOPInstanceUID(); ok {
		fmt.Printf("  Media Storage SOP Instance UID: %s\n", sopInstanceUID)
	}

	// Transfer Syntax UID
	if tsUID, ok := fmi.TransferSyntaxUID(); ok {
		fmt.Printf("  Transfer Syntax UID: %s\n", tsUID)
	}

	// Implementation Class UID
	if implClassUID, ok := fmi.ImplementationClassUID(); ok {
		fmt.Printf("  Implementation Class UID: %s\n", implClassUID)
	}

	// Implementation Version Name
	if implVersionName, ok := fmi.ImplementationVersionName(); ok {
		fmt.Printf("  Implementation Version Name: %s\n", implVersionName)
	}

	if verbose {
		// Source Application Entity Title
		if sourceAET, ok := fmi.SourceApplicationEntityTitle(); ok {
			fmt.Printf("  Source Application Entity Title: %s\n", sourceAET)
		}

		// Sending Application Entity Title
		if sendingAET, ok := fmi.SendingApplicationEntityTitle(); ok {
			fmt.Printf("  Sending Application Entity Title: %s\n", sendingAET)
		}

		// Receiving Application Entity Title
		if receivingAET, ok := fmi.ReceivingApplicationEntityTitle(); ok {
			fmt.Printf("  Receiving Application Entity Title: %s\n", receivingAET)
		}
	}
}

func printMainDatasetInfo(ds *dataset.Dataset, verbose bool) {
	printPatientInfo(ds)
	printStudyInfo(ds)
	printSeriesInfo(ds)
	printImageInfo(ds)

	if verbose {
		printDeviceInfo(ds)
	}
}

func printPatientInfo(ds *dataset.Dataset) {
	if patientName, ok := ds.GetString(tag.PatientName); ok {
		fmt.Printf("  Patient Name: %s\n", patientName)
	}
	if patientID, ok := ds.GetString(tag.PatientID); ok {
		fmt.Printf("  Patient ID: %s\n", patientID)
	}
	if patientBirthDate, ok := ds.GetString(tag.PatientBirthDate); ok {
		fmt.Printf("  Patient Birth Date: %s\n", patientBirthDate)
	}
	if patientSex, ok := ds.GetString(tag.PatientSex); ok {
		fmt.Printf("  Patient Sex: %s\n", patientSex)
	}
}

func printStudyInfo(ds *dataset.Dataset) {
	fmt.Println()
	if studyInstanceUID, ok := ds.GetString(tag.StudyInstanceUID); ok {
		fmt.Printf("  Study Instance UID: %s\n", studyInstanceUID)
	}
	if studyDate, ok := ds.GetString(tag.StudyDate); ok {
		fmt.Printf("  Study Date: %s\n", studyDate)
	}
	if studyTime, ok := ds.GetString(tag.StudyTime); ok {
		fmt.Printf("  Study Time: %s\n", studyTime)
	}
	if studyDescription, ok := ds.GetString(tag.StudyDescription); ok {
		fmt.Printf("  Study Description: %s\n", studyDescription)
	}
}

func printSeriesInfo(ds *dataset.Dataset) {
	fmt.Println()
	if seriesInstanceUID, ok := ds.GetString(tag.SeriesInstanceUID); ok {
		fmt.Printf("  Series Instance UID: %s\n", seriesInstanceUID)
	}
	if seriesNumber, err := ds.GetUInt16(tag.SeriesNumber, 0); err == nil {
		fmt.Printf("  Series Number: %d\n", seriesNumber)
	}
	if seriesDescription, ok := ds.GetString(tag.SeriesDescription); ok {
		fmt.Printf("  Series Description: %s\n", seriesDescription)
	}
	if modality, ok := ds.GetString(tag.Modality); ok {
		fmt.Printf("  Modality: %s\n", modality)
	}
}

func printImageInfo(ds *dataset.Dataset) {
	fmt.Println()
	if sopClassUID, ok := ds.GetString(tag.SOPClassUID); ok {
		fmt.Printf("  SOP Class UID: %s\n", sopClassUID)
	}
	if sopInstanceUID, ok := ds.GetString(tag.SOPInstanceUID); ok {
		fmt.Printf("  SOP Instance UID: %s\n", sopInstanceUID)
	}
	if instanceNumber, err := ds.GetUInt16(tag.InstanceNumber, 0); err == nil {
		fmt.Printf("  Instance Number: %d\n", instanceNumber)
	}

	// Image dimensions (if present)
	if rows, err := ds.GetUInt16(tag.Rows, 0); err == nil {
		if columns, err := ds.GetUInt16(tag.Columns, 0); err == nil {
			fmt.Printf("  Image Size: %d x %d\n", columns, rows)
		}
	}
}

func printDeviceInfo(ds *dataset.Dataset) {
	if manufacturer, ok := ds.GetString(tag.Manufacturer); ok {
		fmt.Printf("  Manufacturer: %s\n", manufacturer)
	}
	if manufacturerModelName, ok := ds.GetString(tag.ManufacturerModelName); ok {
		fmt.Printf("  Manufacturer Model Name: %s\n", manufacturerModelName)
	}
	if stationName, ok := ds.GetString(tag.StationName); ok {
		fmt.Printf("  Station Name: %s\n", stationName)
	}
}
