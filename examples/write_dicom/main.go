// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main demonstrates how to create and write a DICOM file.
package main

import (
	"flag"
	"log"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
	outputPath := flag.String("output", "output.dcm", "Output DICOM file path")
	flag.Parse()

	if err := examplepath.PrepareOutputFile(*outputPath); err != nil {
		log.Fatal(err)
	}

	// Create a new dataset
	ds := dataset.New()

	// === Patient Information ===
	mustAddValue(ds, tag.PatientName, "Doe^John")
	mustAddValue(ds, tag.PatientID, "12345")
	mustAddValue(ds, tag.PatientBirthDate, "19800101")
	mustAddValue(ds, tag.PatientSex, "M")
	mustAddValue(ds, tag.PatientAge, "045Y")

	// === Study Information ===
	mustAddValue(ds, tag.StudyInstanceUID, "1.2.826.0.1.3680043.10.1142.1")
	mustAddValue(ds, tag.StudyDate, "20250101")
	mustAddValue(ds, tag.StudyTime, "120000")
	mustAddValue(ds, tag.StudyDescription, "Sample Study")
	mustAddValue(ds, tag.StudyID, "001")
	mustAddValue(ds, tag.AccessionNumber, "ACC001")
	mustAddValue(ds, tag.ReferringPhysicianName, "Smith^Jane")

	// === Series Information ===
	mustAddValue(ds, tag.SeriesInstanceUID, "1.2.826.0.1.3680043.10.1142.1.1")
	mustAddValue(ds, tag.SeriesNumber, uint16(1))
	mustAddValue(ds, tag.SeriesDescription, "CT Chest")
	mustAddValue(ds, tag.Modality, "CT")
	mustAddValue(ds, tag.SeriesDate, "20250101")
	mustAddValue(ds, tag.SeriesTime, "120000")

	// === Instance Information ===
	// CT Image Storage SOP Class UID
	mustAddValue(ds, tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.2")
	mustAddValue(ds, tag.SOPInstanceUID, "1.2.826.0.1.3680043.10.1142.1.1.1")
	mustAddValue(ds, tag.InstanceNumber, uint16(1))
	mustAddValue(ds, tag.ContentDate, "20250101")
	mustAddValue(ds, tag.ContentTime, "120000")

	// === Image Information ===
	mustAddValue(ds, tag.ImageType, []string{"ORIGINAL", "PRIMARY", "AXIAL"})

	// Image dimensions - using a smaller size for demo (128x128)
	rows := uint16(128)
	columns := uint16(128)
	mustAddValue(ds, tag.Rows, rows)
	mustAddValue(ds, tag.Columns, columns)
	mustAddValue(ds, tag.BitsAllocated, uint16(16))
	mustAddValue(ds, tag.BitsStored, uint16(16))
	mustAddValue(ds, tag.HighBit, uint16(15))
	mustAddValue(ds, tag.PixelRepresentation, uint16(0)) // 0 = unsigned
	mustAddValue(ds, tag.SamplesPerPixel, uint16(1))
	mustAddValue(ds, tag.PhotometricInterpretation, "MONOCHROME2")

	// === Pixel Data ===
	// Create sample pixel data (gradient pattern)
	pixelCount := int(rows) * int(columns)
	pixelData := make([]byte, pixelCount*2) // 2 bytes per pixel (16-bit)

	for i := 0; i < pixelCount; i++ {
		row := i / int(columns)
		col := i % int(columns)

		// Create a gradient pattern
		value := uint16((row*256/int(rows) + col*256/int(columns)) / 2)

		// Write in little-endian format
		pixelData[i*2] = byte(value & 0xFF)
		pixelData[i*2+1] = byte((value >> 8) & 0xFF)
	}

	if err := ds.AddValueWithVR(tag.PixelData, vr.OW, pixelData); err != nil {
		log.Fatalf("Failed to add PixelData: %v", err)
	}

	// === Equipment Information ===
	mustAddValue(ds, tag.Manufacturer, "Sample Manufacturer")
	mustAddValue(ds, tag.ManufacturerModelName, "Sample Model")
	mustAddValue(ds, tag.SoftwareVersions, "1.0")

	// Write to file with explicit VR little endian
	err := writer.WriteFile(*outputPath, ds,
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))

	if err != nil {
		log.Fatalf("Failed to write DICOM file: %v", err)
	}

	log.Printf("DICOM file written successfully to %s", *outputPath)
	log.Println("The file includes:")
	log.Println("  - Complete File Meta Information (Group 0002)")
	log.Println("  - Patient demographics")
	log.Println("  - Study, Series, and Instance information")
	log.Println("  - Image attributes and equipment information")
}

// mustAddValue adds a value and stops the program on the first rejected value,
// so a validation failure cannot silently produce an incomplete file.
func mustAddValue(ds *dataset.Dataset, t *tag.Tag, value any) {
	if err := ds.AddValue(t, value); err != nil {
		log.Fatalf("Failed to add %s: %v", t, err)
	}
}
