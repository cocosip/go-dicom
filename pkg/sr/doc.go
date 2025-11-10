// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package sr provides support for DICOM Structured Reports (SR).
//
// DICOM Structured Reports are used to encode clinical findings, measurements,
// and observations in a structured, machine-readable format. They represent
// information as a tree of content items with defined relationships.
//
// # Key Concepts
//
// **Content Items**: The building blocks of an SR. Each item has:
//   - A concept name (coded entry describing what the item represents)
//   - A value type (TEXT, NUM, CODE, IMAGE, CONTAINER, etc.)
//   - A relationship type (CONTAINS, HAS PROPERTIES, INFERRED FROM, etc.)
//   - An optional value
//
// **Value Types**:
//   - TEXT: Text values
//   - NUM: Numeric measurements with units
//   - CODE: Coded concepts
//   - CONTAINER: Grouping of other content items
//   - IMAGE: Reference to an image SOP instance
//   - COMPOSITE: Reference to a composite SOP instance
//   - PNAME: Person name
//   - DATE, TIME, DATETIME: Temporal values
//   - UIDREF: UID reference
//
// **Relationship Types**:
//   - CONTAINS: Parent-child containment
//   - HAS PROPERTIES: Describes properties
//   - INFERRED FROM: Indicates source
//   - SELECTED FROM: Selection from options
//   - HAS OBS CONTEXT: Observation context
//   - HAS ACQ CONTEXT: Acquisition context
//   - HAS CONCEPT MOD: Concept modifier
//
// # Basic Usage
//
//	// Create a new structured report
//	rootCode := sr.NewCodeItem("113704", "DCM", "SR Document")
//	report := sr.NewStructuredReport(rootCode)
//
//	// Add a text finding
//	findingCode := sr.NewCodeItem("121071", "DCM", "Finding")
//	report.Add(findingCode, sr.RelationshipContains, sr.ValueTypeText, "No abnormalities detected")
//
//	// Add a numeric measurement
//	measureCode := sr.NewCodeItem("33728-7", "LN", "Size")
//	units := sr.NewCodeItem("mm", "UCUM", "millimeter")
//	value := sr.NewMeasuredValue(25.5, units)
//	report.Add(measureCode, sr.RelationshipContains, value)
//
//	// Save to file
//	err := report.Save("report.dcm")
//
// # Reference
//
// Based on:
//   - DICOM Part 3: Information Object Definitions, Annex A (SR IOD)
//   - fo-dicom StructuredReport package
//   - https://dicom.nema.org/medical/dicom/current/output/chtml/part03/sect_A.35.html
//
package sr
