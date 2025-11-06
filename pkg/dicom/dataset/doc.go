// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package dataset provides DICOM dataset implementation.
//
// A DICOM dataset is a collection of data elements (items) organized by tag.
// Datasets are the fundamental container for DICOM data, used in:
//   - DICOM files
//   - Network messages (C-STORE, C-FIND, etc.)
//   - Sequence items (nested datasets)
//
// # Basic Usage
//
//	// Create a new dataset
//	ds := dataset.New()
//
//	// Add elements
//	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
//	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
//
//	// Retrieve elements
//	elem, exists := ds.Get(tag.PatientName)
//	if exists {
//	    str := elem.(*element.String)
//	    name := str.GetString()
//	}
//
//	// Remove elements
//	ds.Remove(tag.PatientName)
//
//	// Iterate over all elements
//	for _, elem := range ds.Elements() {
//	    fmt.Println(elem.Tag(), elem.ValueRepresentation())
//	}
//
// # Thread Safety
//
// Dataset is NOT thread-safe by default. If concurrent access is needed,
// use external synchronization (e.g., sync.RWMutex).
//
// # Sequences
//
// Sequence elements (VR=SQ) contain nested datasets:
//
//	seq := dataset.NewSequence(tag.ReferencedImageSequence)
//	item1 := dataset.New()
//	item1.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
//	seq.AddItem(item1)
//	ds.Add(seq)
package dataset
