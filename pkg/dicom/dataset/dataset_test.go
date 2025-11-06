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

// TestDatasetBasics tests basic dataset operations
func TestDatasetBasics(t *testing.T) {
	ds := dataset.New()

	t.Run("EmptyDataset", func(t *testing.T) {
		if !ds.IsEmpty() {
			t.Error("New dataset should be empty")
		}
		if ds.Count() != 0 {
			t.Errorf("Count() = %d, want 0", ds.Count())
		}
	})

	t.Run("Add", func(t *testing.T) {
		elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
		err := ds.Add(elem)
		if err != nil {
			t.Fatalf("Add() error = %v", err)
		}

		if ds.IsEmpty() {
			t.Error("Dataset should not be empty after adding element")
		}
		if ds.Count() != 1 {
			t.Errorf("Count() = %d, want 1", ds.Count())
		}
	})

	t.Run("AddDuplicate", func(t *testing.T) {
		elem := element.NewString(tag.PatientName, vr.PN, []string{"Smith^Jane"})
		err := ds.Add(elem)
		if err == nil {
			t.Error("Add() should return error for duplicate tag")
		}
	})

	t.Run("Get", func(t *testing.T) {
		elem, exists := ds.Get(tag.PatientName)
		if !exists {
			t.Fatal("Get() should find existing element")
		}

		str, ok := elem.(*element.String)
		if !ok {
			t.Fatal("Element should be *element.String")
		}

		name := str.GetString()
		if name != "Doe^John" {
			t.Errorf("GetString() = %q, want %q", name, "Doe^John")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		if !ds.Contains(tag.PatientName) {
			t.Error("Contains() should return true for existing tag")
		}

		if ds.Contains(tag.StudyDate) {
			t.Error("Contains() should return false for non-existing tag")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		removed := ds.Remove(tag.PatientName)
		if !removed {
			t.Error("Remove() should return true for existing element")
		}

		if ds.Contains(tag.PatientName) {
			t.Error("Element should be removed")
		}

		removed = ds.Remove(tag.PatientName)
		if removed {
			t.Error("Remove() should return false for non-existing element")
		}
	})
}

// TestDatasetAddOrUpdate tests the AddOrUpdate method
func TestDatasetAddOrUpdate(t *testing.T) {
	ds := dataset.New()

	// Add initial element
	elem1 := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
	err := ds.AddOrUpdate(elem1)
	if err != nil {
		t.Fatalf("AddOrUpdate() error = %v", err)
	}

	// Update existing element
	elem2 := element.NewString(tag.PatientName, vr.PN, []string{"Smith^Jane"})
	err = ds.AddOrUpdate(elem2)
	if err != nil {
		t.Fatalf("AddOrUpdate() error = %v", err)
	}

	// Verify update
	elem, _ := ds.Get(tag.PatientName)
	str := elem.(*element.String)
	name := str.GetString()
	if name != "Smith^Jane" {
		t.Errorf("GetString() = %q, want %q", name, "Smith^Jane")
	}

	if ds.Count() != 1 {
		t.Errorf("Count() = %d, want 1", ds.Count())
	}
}

// TestDatasetMultipleElements tests multiple elements
func TestDatasetMultipleElements(t *testing.T) {
	ds := dataset.New()

	// Add multiple elements
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

	if ds.Count() != 3 {
		t.Errorf("Count() = %d, want 3", ds.Count())
	}

	// Test Elements() - should be sorted by tag
	elements := ds.Elements()
	if len(elements) != 3 {
		t.Errorf("len(Elements()) = %d, want 3", len(elements))
	}

	// Verify sorting (tags should be in ascending order)
	for i := 1; i < len(elements); i++ {
		if elements[i-1].Tag().ToUint32() >= elements[i].Tag().ToUint32() {
			t.Error("Elements should be sorted by tag")
		}
	}
}

// TestDatasetClear tests Clear method
func TestDatasetClear(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))

	ds.Clear()

	if !ds.IsEmpty() {
		t.Error("Dataset should be empty after Clear()")
	}
	if ds.Count() != 0 {
		t.Errorf("Count() = %d, want 0", ds.Count())
	}
}

// TestDatasetClone tests Clone method
func TestDatasetClone(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))

	clone := ds.Clone()

	if clone.Count() != ds.Count() {
		t.Errorf("Clone count = %d, want %d", clone.Count(), ds.Count())
	}

	// Verify elements are the same
	if !clone.Contains(tag.PatientName) {
		t.Error("Clone should contain PatientName")
	}
	if !clone.Contains(tag.StudyDate) {
		t.Error("Clone should contain StudyDate")
	}

	// Modify clone shouldn't affect original
	clone.Remove(tag.PatientName)
	if !ds.Contains(tag.PatientName) {
		t.Error("Original dataset should still contain PatientName")
	}
}

// TestDatasetMerge tests Merge method
func TestDatasetMerge(t *testing.T) {
	ds1 := dataset.New()
	ds1.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds1.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))

	ds2 := dataset.New()
	ds2.Add(element.NewString(tag.PatientName, vr.PN, []string{"Smith^Jane"}))
	ds2.Add(element.NewString(tag.StudyTime, vr.TM, []string{"120000"}))

	t.Run("MergeWithoutOverwrite", func(t *testing.T) {
		merged := ds1.Clone()
		merged.Merge(ds2, false)

		// Should have 3 elements (PatientName not overwritten)
		if merged.Count() != 3 {
			t.Errorf("Count() = %d, want 3", merged.Count())
		}

		// PatientName should still be "Doe^John"
		name, _ := merged.GetString(tag.PatientName)
		if name != "Doe^John" {
			t.Errorf("PatientName = %q, want %q", name, "Doe^John")
		}
	})

	t.Run("MergeWithOverwrite", func(t *testing.T) {
		merged := ds1.Clone()
		merged.Merge(ds2, true)

		// Should have 3 elements
		if merged.Count() != 3 {
			t.Errorf("Count() = %d, want 3", merged.Count())
		}

		// PatientName should be overwritten to "Smith^Jane"
		name, _ := merged.GetString(tag.PatientName)
		if name != "Smith^Jane" {
			t.Errorf("PatientName = %q, want %q", name, "Smith^Jane")
		}
	})
}

// TestDatasetFilter tests Filter method
func TestDatasetFilter(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

	// Filter to keep only string elements
	filtered := ds.Filter(func(elem element.Element) bool {
		_, ok := elem.(*element.String)
		return ok
	})

	if filtered.Count() != 2 {
		t.Errorf("Filtered count = %d, want 2", filtered.Count())
	}

	if filtered.Contains(tag.Rows) {
		t.Error("Filtered dataset should not contain Rows (UnsignedShort)")
	}
}
