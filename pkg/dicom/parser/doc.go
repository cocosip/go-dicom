// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package parser provides DICOM file parsing functionality.
//
// This package implements reading and parsing of DICOM files (Part 10 format).
// It handles:
//   - File preamble and DICM prefix validation
//   - Transfer syntax detection
//   - Element parsing (Tag, VR, Length, Value)
//   - Dataset construction
//   - Sequence and nested dataset handling
//
// # DICOM File Format
//
// A DICOM Part 10 file consists of:
//  1. 128-byte preamble (usually zeros)
//  2. 4-byte DICM prefix
//  3. File Meta Information (Group 0002, Explicit VR Little Endian)
//  4. Dataset (encoding depends on Transfer Syntax)
//
// # Basic Usage
//
//	// Open and parse a DICOM file
//	file, err := os.Open("image.dcm")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	// Parse the file
//	ds, err := parser.Parse(file)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Access parsed data
//	patientName := ds.TryGetString(tag.PatientName)
//	fmt.Printf("Patient: %s\n", patientName)
//
// # Advanced Usage
//
//	// Create parser with options
//	p := parser.New(
//	    parser.WithMaxElementSize(100*1024*1024), // 100MB max
//	    parser.WithStopAtTag(tag.PixelData),       // Stop before pixel data
//	)
//
//	ds, err := p.Parse(file)
package parser
