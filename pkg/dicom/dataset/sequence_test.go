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

// TestSequenceBasics tests basic sequence operations
func TestSequenceBasics(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	t.Run("EmptySequence", func(t *testing.T) {
		if !seq.IsEmpty() {
			t.Error("New sequence should be empty")
		}
		if seq.Count() != 0 {
			t.Errorf("Count() = %d, want 0", seq.Count())
		}
	})

	t.Run("AddItem", func(t *testing.T) {
    item := dataset.New()
    _ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
		seq.AddItem(item)

		if seq.IsEmpty() {
			t.Error("Sequence should not be empty after adding item")
		}
		if seq.Count() != 1 {
			t.Errorf("Count() = %d, want 1", seq.Count())
		}
	})

	t.Run("GetItem", func(t *testing.T) {
		item := seq.GetItem(0)
		if item == nil {
			t.Fatal("GetItem(0) should return non-nil item")
		}

		uid, exists := item.GetString(tag.StudyInstanceUID)
		if !exists {
			t.Fatal("Item should contain StudyInstanceUID")
		}
		if uid != "1.2.3" {
			t.Errorf("StudyInstanceUID = %q, want %q", uid, "1.2.3")
		}
	})

	t.Run("GetItemOutOfRange", func(t *testing.T) {
		item := seq.GetItem(10)
		if item != nil {
			t.Error("GetItem() should return nil for out of range index")
		}
	})
}

// TestSequenceMultipleItems tests sequence with multiple items
func TestSequenceMultipleItems(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Add multiple items
    for i := 0; i < 3; i++ {
        item := dataset.New()
        _ = item.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3." + string(rune('0'+i))}))
        seq.AddItem(item)
    }

	if seq.Count() != 3 {
		t.Errorf("Count() = %d, want 3", seq.Count())
	}

	// Verify items
	items := seq.GetItems()
	if len(items) != 3 {
		t.Errorf("len(GetItems()) = %d, want 3", len(items))
	}

	// Check first item
	uid, _ := items[0].GetString(tag.SOPInstanceUID)
	if uid != "1.2.3.0" {
		t.Errorf("First item UID = %q, want %q", uid, "1.2.3.0")
	}
}

// TestSequenceRemoveItem tests RemoveItem
func TestSequenceRemoveItem(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Add items
	for i := 0; i < 3; i++ {
		item := dataset.New()
		seq.AddItem(item)
	}

	t.Run("RemoveMiddleItem", func(t *testing.T) {
		removed := seq.RemoveItem(1)
		if !removed {
			t.Error("RemoveItem() should return true")
		}
		if seq.Count() != 2 {
			t.Errorf("Count() = %d, want 2", seq.Count())
		}
	})

	t.Run("RemoveOutOfRange", func(t *testing.T) {
		removed := seq.RemoveItem(10)
		if removed {
			t.Error("RemoveItem() should return false for out of range index")
		}
	})
}

// TestSequenceClear tests Clear method
func TestSequenceClear(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Add items
	for i := 0; i < 3; i++ {
		seq.AddItem(dataset.New())
	}

	seq.Clear()

	if !seq.IsEmpty() {
		t.Error("Sequence should be empty after Clear()")
	}
	if seq.Count() != 0 {
		t.Errorf("Count() = %d, want 0", seq.Count())
	}
}

// TestSequenceClone tests Clone method
func TestSequenceClone(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Add items
    item := dataset.New()
    _ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
	seq.AddItem(item)

	clone := seq.Clone()

	if clone.Count() != seq.Count() {
		t.Errorf("Clone count = %d, want %d", clone.Count(), seq.Count())
	}

	// Modify clone shouldn't affect original
	clone.Clear()
	if seq.Count() != 1 {
		t.Error("Original sequence should still have 1 item")
	}
}

// TestSequenceInDataset tests adding sequence to dataset
func TestSequenceInDataset(t *testing.T) {
	ds := dataset.New()

	// Create sequence
	seq := dataset.NewSequence(tag.ReferencedImageSequence)
    item := dataset.New()
    _ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3"}))
	seq.AddItem(item)

	// Add sequence to dataset
	err := ds.Add(seq)
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	// Retrieve sequence from dataset
	retrievedSeq, err := ds.GetSequence(tag.ReferencedImageSequence)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}

	if retrievedSeq.Count() != 1 {
		t.Errorf("Retrieved sequence count = %d, want 1", retrievedSeq.Count())
	}

	// Verify item content
	retrievedItem := retrievedSeq.GetItem(0)
	uid, _ := retrievedItem.GetString(tag.StudyInstanceUID)
	if uid != "1.2.3" {
		t.Errorf("StudyInstanceUID = %q, want %q", uid, "1.2.3")
	}
}

// TestSequenceFilter tests Filter method
func TestSequenceFilter(t *testing.T) {
	seq := dataset.NewSequence(tag.ReferencedImageSequence)

	// Add items with different modalities
	for i, modality := range []string{"CT", "MR", "CT"} {
        item := dataset.New()
        _ = item.Add(element.NewString(tag.Modality, vr.CS, []string{modality}))
        _ = item.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3." + string(rune('0'+i))}))
		seq.AddItem(item)
	}

	// Filter to keep only CT items
	filtered := seq.Filter(func(ds *dataset.Dataset) bool {
		modality, _ := ds.GetString(tag.Modality)
		return modality == "CT"
	})

	if filtered.Count() != 2 {
		t.Errorf("Filtered count = %d, want 2", filtered.Count())
	}

	// Verify filtered items are CT
	for i := 0; i < filtered.Count(); i++ {
		item := filtered.GetItem(i)
		modality, _ := item.GetString(tag.Modality)
		if modality != "CT" {
			t.Errorf("Filtered item %d modality = %q, want CT", i, modality)
		}
	}
}
