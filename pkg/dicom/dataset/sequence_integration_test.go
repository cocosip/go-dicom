// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestAddSequenceToDataset tests adding a Sequence to a Dataset
func TestAddSequenceToDataset(t *testing.T) {
	// Create a dataset
	ds := dataset.New()

	// Create a sequence
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Create item datasets
	item1 := dataset.New()
	item1.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
	item1.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.6"}))

	item2 := dataset.New()
	item2.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.7"}))
	item2.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.4.8"}))

	// Add items to sequence
	seq.AddItem(item1)
	seq.AddItem(item2)

	// Try to add sequence to dataset
	err := ds.Add(seq)
	if err != nil {
		t.Fatalf("Failed to add sequence to dataset: %v", err)
	}

	// Verify sequence was added
	elem, exists := ds.Get(tag.ReferencedImageSequence)
	if !exists {
		t.Fatal("Sequence not found in dataset")
	}

	// Verify it's a sequence
	retrievedSeq, ok := elem.(*dataset.Sequence)
	if !ok {
		t.Fatalf("Element is not a Sequence, got type: %T", elem)
	}

	// Verify sequence contents
	if retrievedSeq.Count() != 2 {
		t.Errorf("Expected 2 items in sequence, got %d", retrievedSeq.Count())
	}

	t.Logf("✓ Successfully added Sequence to Dataset with %d items", retrievedSeq.Count())
}

// TestSequenceAsElement verifies Sequence implements Element interface
func TestSequenceAsElement(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedStudySequence)

	// Test Element interface methods
	if seq.Tag() == nil {
		t.Error("Tag() returned nil")
	}

	if seq.ValueRepresentation() != vr.SQ {
		t.Errorf("Expected VR to be SQ, got %v", seq.ValueRepresentation())
	}

	if seq.Buffer() == nil {
		t.Error("Buffer() returned nil")
	}

	if seq.String() == "" {
		t.Error("String() returned empty string")
	}

	if err := seq.Validate(); err != nil {
		t.Errorf("Validate() failed: %v", err)
	}

	t.Log("✓ Sequence correctly implements Element interface")
}

// TestDatasetWithMultipleSequences tests adding multiple sequences
func TestDatasetWithMultipleSequences(t *testing.T) {
	ds := dataset.New()

	// Add regular elements
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	// Create and add first sequence
	seq1 := dataset.NewSequence(tag.ReferencedStudySequence)
	item1 := dataset.New()
	item1.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
	seq1.AddItem(item1)
	ds.Add(seq1)

	// Create and add second sequence
	seq2 := dataset.NewSequence(tag.ReferencedImageSequence)
	item2 := dataset.New()
	item2.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.4"}))
	seq2.AddItem(item2)
	ds.Add(seq2)

	// Verify dataset structure
	if ds.Count() != 4 {
		t.Errorf("Expected 4 elements (2 regular + 2 sequences), got %d", ds.Count())
	}

	// Verify sequences can be retrieved
	elem1, exists1 := ds.Get(tag.ReferencedStudySequence)
	if !exists1 {
		t.Error("First sequence not found")
	}
	if _, ok := elem1.(*dataset.Sequence); !ok {
		t.Error("First sequence is not of type Sequence")
	}

	elem2, exists2 := ds.Get(tag.ReferencedImageSequence)
	if !exists2 {
		t.Error("Second sequence not found")
	}
	if _, ok := elem2.(*dataset.Sequence); !ok {
		t.Error("Second sequence is not of type Sequence")
	}

	t.Logf("✓ Successfully added %d elements including multiple sequences", ds.Count())
}

// TestNestedSequences tests sequences containing sequences
func TestNestedSequences(t *testing.T) {
	// Create outer dataset
	ds := dataset.New()

	// Create outer sequence
	outerSeq := dataset.NewSequence(tag.ReferencedStudySequence)

	// Create item with nested sequence
	item := dataset.New()
	item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))

	// Create nested sequence
	nestedSeq := dataset.NewSequence(tag.ReferencedImageSequence)
	nestedItem := dataset.New()
	nestedItem.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.4"}))
	nestedSeq.AddItem(nestedItem)

	// Add nested sequence to item
	err := item.Add(nestedSeq)
	if err != nil {
		t.Fatalf("Failed to add nested sequence: %v", err)
	}

	// Add item to outer sequence
	outerSeq.AddItem(item)

	// Add outer sequence to dataset
	err = ds.Add(outerSeq)
	if err != nil {
		t.Fatalf("Failed to add outer sequence: %v", err)
	}

	// Verify nested structure
	outerElem, exists := ds.Get(tag.ReferencedStudySequence)
	if !exists {
		t.Fatal("Outer sequence not found")
	}

	outerSeqRetrieved := outerElem.(*dataset.Sequence)
	retrievedItem := outerSeqRetrieved.GetItem(0)
	if retrievedItem == nil {
		t.Fatal("Item not found in outer sequence")
	}

	nestedElem, exists := retrievedItem.Get(tag.ReferencedImageSequence)
	if !exists {
		t.Fatal("Nested sequence not found in item")
	}

	nestedSeqRetrieved, ok := nestedElem.(*dataset.Sequence)
	if !ok {
		t.Fatal("Nested element is not a Sequence")
	}

	if nestedSeqRetrieved.Count() != 1 {
		t.Errorf("Expected 1 item in nested sequence, got %d", nestedSeqRetrieved.Count())
	}

	t.Log("✓ Successfully created and verified nested sequence structure")
}
