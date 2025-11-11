// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to create and write a DICOM file.
package main

import (
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func main() {
	// Create a new dataset
	ds := dataset.New()

	// Add patient information
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	_ = ds.Add(element.NewString(tag.PatientBirthDate, vr.DA, []string{"19800101"}))
	_ = ds.Add(element.NewString(tag.PatientSex, vr.CS, []string{"M"}))

	// Add study information
	_ = ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
	_ = ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250101"}))
	_ = ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Sample Study"}))

	// Add series information
	_ = ds.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.5.1"}))
	_ = ds.Add(element.NewUnsignedShort(tag.SeriesNumber, []uint16{1}))
	_ = ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))

	// Add instance information
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.1.1"}))
	_ = ds.Add(element.NewUnsignedShort(tag.InstanceNumber, []uint16{1}))

	// Add image dimensions
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))

	// Write to file with explicit VR little endian
	err := writer.WriteFile("output.dcm", ds,
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))

	if err != nil {
		log.Fatalf("Failed to write DICOM file: %v", err)
	}

	log.Println("DICOM file written successfully to output.dcm")
}
