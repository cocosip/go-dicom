// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// This file demonstrates using global configuration for Implementation information.
// This is the recommended approach - set it once at app startup.
// To run: go run global_config_example.go

//go:build ignore

package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
	outputDir := flag.String("output-dir", ".", "Directory to write generated DICOM files")
	flag.Parse()
	if err := examplepath.PrepareOutputDir(*outputDir); err != nil {
		log.Fatal(err)
	}

	// ========================================================================
	// RECOMMENDED USAGE: Configure implementation information at app startup
	// This is typically done once in your main() or init() function
	// ========================================================================
	log.Println("=== Global DICOM Configuration Demo ===")
	log.Println("Step 1: Setting global implementation information...")
	writer.SetDefaultImplementationClassUID("1.2.840.99999.5.6.7")
	writer.SetDefaultImplementationVersionName("MyHospitalApp_3.2.1")

	log.Printf("  ✓ Implementation Class UID: %s", writer.GetDefaultImplementationClassUID())
	log.Printf("  ✓ Implementation Version:   %s\n", writer.GetDefaultImplementationVersionName())

	// ========================================================================
	// Now all files written will automatically use these settings
	// ========================================================================

	log.Println("\nStep 2: Writing multiple DICOM files...")
	log.Println("  (All files will automatically use the global configuration)")

	// Create and write first file
	ds1 := createSampleDataset("Smith^John", "P001")
	if err := writer.WriteFile(filepath.Join(*outputDir, "test_file1.dcm"), ds1); err != nil {
		log.Fatalf("Failed to write file1: %v", err)
	}
	log.Println("  ✓ test_file1.dcm written")

	// Create and write second file
	ds2 := createSampleDataset("Doe^Jane", "P002")
	if err := writer.WriteFile(filepath.Join(*outputDir, "test_file2.dcm"), ds2); err != nil {
		log.Fatalf("Failed to write file2: %v", err)
	}
	log.Println("  ✓ test_file2.dcm written")

	// Create and write third file
	ds3 := createSampleDataset("Brown^Robert", "P003")
	if err := writer.WriteFile(filepath.Join(*outputDir, "test_file3.dcm"), ds3); err != nil {
		log.Fatalf("Failed to write file3: %v", err)
	}
	log.Println("  ✓ test_file3.dcm written")

	log.Println("\n=== Success! ===")
	log.Println("\nAll files have been created with:")
	log.Printf("  Implementation Class UID: %s", writer.GetDefaultImplementationClassUID())
	log.Printf("  Implementation Version:   %s", writer.GetDefaultImplementationVersionName())

	log.Println("\n✓ BENEFIT: You only need to configure this ONCE at app startup,")
	log.Println("  and ALL DICOM files will have consistent implementation identification!")

	log.Println("\nYou can verify the files using:")
	log.Println("  go run ../read_dicom/main.go  (update path in main.go)")
	log.Println("  or: dcmdump test_file1.dcm")
}

func createSampleDataset(patientName, patientID string) *dataset.Dataset {
	ds := dataset.New()
	_ = ds.AddValue(tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.2") // CT Image Storage
	_ = ds.AddValue(tag.SOPInstanceUID, "1.2.3.4.5."+strings.TrimLeft(strings.TrimPrefix(patientID, "P"), "0"))
	_ = ds.AddValue(tag.PatientName, patientName)
	_ = ds.AddValue(tag.PatientID, patientID)
	return ds
}
