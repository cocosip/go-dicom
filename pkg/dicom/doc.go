// Package dicom provides core DICOM data types and functionality.
//
// This package implements the fundamental DICOM (Digital Imaging and
// Communications in Medicine) data structures including:
//   - DICOM Tags (unique identifiers for data elements)
//   - Value Representations (VR - data types)
//   - DICOM Elements (individual data items)
//   - DICOM Datasets (collections of elements)
//   - DICOM Dictionary (tag definitions and metadata)
//
// Example usage:
//
//	file, err := dicom.Open("example.dcm")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	patientName, _ := file.Dataset.GetString(dicom.TagPatientName)
//	fmt.Printf("Patient: %s\n", patientName)
package dicom
