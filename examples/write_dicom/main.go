// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
	_ = ds.AddValue(tag.PatientName, "Doe^John")
	_ = ds.AddValue(tag.PatientID, "12345")
	_ = ds.AddValue(tag.PatientBirthDate, "19800101")
	_ = ds.AddValue(tag.PatientSex, "M")
	_ = ds.AddValue(tag.PatientAge, "045Y")

	// === Study Information ===
	_ = ds.AddValue(tag.StudyInstanceUID, "1.2.826.0.1.3680043.10.1142.1")
	_ = ds.AddValue(tag.StudyDate, "20250101")
	_ = ds.AddValue(tag.StudyTime, "120000")
	_ = ds.AddValue(tag.StudyDescription, "Sample Study")
	_ = ds.AddValue(tag.StudyID, "001")
	_ = ds.AddValue(tag.AccessionNumber, "ACC001")
	_ = ds.AddValue(tag.ReferringPhysicianName, "Smith^Jane")

	// === Series Information ===
	_ = ds.AddValue(tag.SeriesInstanceUID, "1.2.826.0.1.3680043.10.1142.1.1")
	_ = ds.AddValue(tag.SeriesNumber, uint16(1))
	_ = ds.AddValue(tag.SeriesDescription, "CT Chest")
	_ = ds.AddValue(tag.Modality, "CT")
	_ = ds.AddValue(tag.SeriesDate, "20250101")
	_ = ds.AddValue(tag.SeriesTime, "120000")

	// === Instance Information ===
	// CT Image Storage SOP Class UID
	_ = ds.AddValue(tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.2")
	_ = ds.AddValue(tag.SOPInstanceUID, "1.2.826.0.1.3680043.10.1142.1.1.1")
	_ = ds.AddValue(tag.InstanceNumber, uint16(1))
	_ = ds.AddValue(tag.ContentDate, "20250101")
	_ = ds.AddValue(tag.ContentTime, "120000")

	// === Image Information ===
	_ = ds.AddValue(tag.ImageType, []string{"ORIGINAL", "PRIMARY", "AXIAL"})

	// Image dimensions - using a smaller size for demo (128x128)
	rows := uint16(128)
	columns := uint16(128)
	_ = ds.AddValue(tag.Rows, rows)
	_ = ds.AddValue(tag.Columns, columns)
	_ = ds.AddValue(tag.BitsAllocated, uint16(16))
	_ = ds.AddValue(tag.BitsStored, uint16(16))
	_ = ds.AddValue(tag.HighBit, uint16(15))
	_ = ds.AddValue(tag.PixelRepresentation, uint16(0)) // 0 = unsigned
	_ = ds.AddValue(tag.SamplesPerPixel, uint16(1))
	_ = ds.AddValue(tag.PhotometricInterpretation, "MONOCHROME2")

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

	_ = ds.AddValueWithVR(tag.PixelData, vr.OW, pixelData)

	// === Equipment Information ===
	_ = ds.AddValue(tag.Manufacturer, "Sample Manufacturer")
	_ = ds.AddValue(tag.ManufacturerModelName, "Sample Model")
	_ = ds.AddValue(tag.SoftwareVersions, "1.0")

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
