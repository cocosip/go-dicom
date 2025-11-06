// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package element provides DICOM element implementations.
//
// This package contains the core DICOM element types and their implementations.
// Elements are the fundamental building blocks of DICOM datasets, representing
// individual data items with a tag, value representation (VR), and binary data.
//
// # Element Types
//
// Elements are organized by their value representation (VR):
//   - String elements: AE, AS, CS, DA, DS, DT, IS, LO, LT, PN, SH, ST, TM, UC, UI, UR, UT
//   - Numeric elements: FL, FD, SL, SS, UL, US, SV, UV
//   - Binary elements: OB, OD, OF, OL, OV, OW, UN
//   - Sequence elements: SQ
//
// # Design Principles
//
// 1. Type Safety: Each element type provides type-safe accessors for its values
// 2. Lazy Loading: Element data can be loaded from buffers on demand
// 3. Validation: Elements validate their data against DICOM constraints
// 4. Encoding: String elements handle character set encoding properly
//
// # Example Usage
//
//	// Create a string element
//	elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
//
//	// Create a numeric element
//	elem := element.NewUnsignedShort(tag.PixelRows, []uint16{512})
//
//	// Access values with type safety
//	name := elem.GetString(0)
//	rows := elem.GetUInt16(0)
package element
