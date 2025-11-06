// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package writer provides DICOM file writing functionality.
//
// This package implements writing DICOM files in Part 10 format.
// It handles:
//   - File preamble and DICM prefix
//   - File Meta Information (Group 0002)
//   - Dataset encoding based on Transfer Syntax
//   - Element writing (Tag, VR, Length, Value)
//   - Sequence and nested dataset handling
//
// # Basic Usage
//
//	// Create datasets
//	fileMetaInfo := dataset.New()
//	fileMetaInfo.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
//	    []string{"1.2.840.10008.1.2.1"}))
//
//	ds := dataset.New()
//	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
//	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
//
//	// Write to file
//	file, err := os.Create("output.dcm")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	ts := transfer.ExplicitVRLittleEndian
//	err = writer.Write(file, fileMetaInfo, ds, ts)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Advanced Usage
//
//	// Write without preamble (for dataset-only output)
//	err := writer.Write(w, fileMetaInfo, ds, ts, writer.WithoutPreamble())
//
// # Transfer Syntax
//
// The writer automatically handles:
//   - Explicit VR vs Implicit VR encoding
//   - Little Endian vs Big Endian byte order
//   - File Meta Information (always Explicit VR Little Endian)
package writer
