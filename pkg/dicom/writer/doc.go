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
//	ds := dataset.New()
//	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
//	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
//
//	// Write to file (simplest form)
//	err := writer.WriteFile("output.dcm", ds)
//
// # With Options
//
//	err := writer.WriteFile("output.dcm", ds,
//	    writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
//	    writer.WithImplementationClassUID("1.2.840.12345.1"),
//	    writer.WithImplementationVersionName("MyApp_1.0"))
//
// # Advanced Usage
//
//	// Write to stream
//	file, err := os.Create("output.dcm")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	err = writer.Write(file, ds,
//	    writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))
//
//	// Write without preamble (for network transmission)
//	err := writer.Write(w, ds, writer.WithoutPreamble())
//
//	// Custom File Meta Information
//	fileMetaInfo := dataset.New()
//	fileMetaInfo.Add(element.NewString(tag.SourceApplicationEntityTitle, vr.AE,
//	    []string{"MY_WORKSTATION"}))
//	err := writer.Write(w, ds, writer.WithFileMetaInfo(fileMetaInfo))
//
// # Transfer Syntax
//
// The writer automatically handles:
//   - Explicit VR vs Implicit VR encoding
//   - Little Endian vs Big Endian byte order
//   - File Meta Information (always Explicit VR Little Endian)
//
// # Implementation Information (RECOMMENDED APPROACH)
//
// Configure implementation information globally at application startup:
//
//	func main() {
//	    // Set once at startup
//	    writer.SetDefaultImplementationClassUID("1.2.840.12345.1.2.3")
//	    writer.SetDefaultImplementationVersionName("MyDicomApp_2.1.0")
//
//	    // Now all files automatically use these settings
//	    writer.WriteFile("file1.dcm", ds1)
//	    writer.WriteFile("file2.dcm", ds2)
//	    // etc...
//	}
//
// You can also override on a per-file basis if needed (rare):
//
//	writer.WriteFile("special.dcm", ds,
//	    writer.WithImplementationClassUID("1.2.840.99999.1"))
//
// This is useful for identifying files created by your application and for
// compliance with your organization's UID scheme.
package writer
