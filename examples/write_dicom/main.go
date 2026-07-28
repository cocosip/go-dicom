// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to create and write a DICOM file.
package main

import (
	"flag"
	"log"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
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
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	_ = ds.Add(element.NewString(tag.PatientBirthDate, vr.DA, []string{"19800101"}))
	_ = ds.Add(element.NewString(tag.PatientSex, vr.CS, []string{"M"}))
	_ = ds.Add(element.NewString(tag.PatientAge, vr.AS, []string{"045Y"}))

	// === Study Information ===
	_ = ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.1142.1"}))
	_ = ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250101"}))
	_ = ds.Add(element.NewString(tag.StudyTime, vr.TM, []string{"120000"}))
	_ = ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Sample Study"}))
	_ = ds.Add(element.NewString(tag.StudyID, vr.SH, []string{"001"}))
	_ = ds.Add(element.NewString(tag.AccessionNumber, vr.SH, []string{"ACC001"}))
	_ = ds.Add(element.NewString(tag.ReferringPhysicianName, vr.PN, []string{"Smith^Jane"}))

	// === Series Information ===
	_ = ds.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.1142.1.1"}))
	_ = ds.Add(element.NewUnsignedShort(tag.SeriesNumber, []uint16{1}))
	_ = ds.Add(element.NewString(tag.SeriesDescription, vr.LO, []string{"CT Chest"}))
	_ = ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))
	_ = ds.Add(element.NewString(tag.SeriesDate, vr.DA, []string{"20250101"}))
	_ = ds.Add(element.NewString(tag.SeriesTime, vr.TM, []string{"120000"}))

	// === Instance Information ===
	// CT Image Storage SOP Class UID
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.1142.1.1.1"}))
	_ = ds.Add(element.NewUnsignedShort(tag.InstanceNumber, []uint16{1}))
	_ = ds.Add(element.NewString(tag.ContentDate, vr.DA, []string{"20250101"}))
	_ = ds.Add(element.NewString(tag.ContentTime, vr.TM, []string{"120000"}))

	// === Image Information ===
	_ = ds.Add(element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"}))

	// Image dimensions - using a smaller size for demo (128x128)
	rows := uint16(128)
	columns := uint16(128)
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{columns}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{16}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{15}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0})) // 0 = unsigned
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

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

	_ = ds.Add(element.NewOtherWord(tag.PixelData, pixelData))

	// === Equipment Information ===
	_ = ds.Add(element.NewString(tag.Manufacturer, vr.LO, []string{"Sample Manufacturer"}))
	_ = ds.Add(element.NewString(tag.ManufacturerModelName, vr.LO, []string{"Sample Model"}))
	_ = ds.Add(element.NewString(tag.SoftwareVersions, vr.LO, []string{"1.0"}))

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
