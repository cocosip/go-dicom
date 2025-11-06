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

// TestDatasetGetString tests GetString accessor
func TestDatasetGetString(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))

	str, exists := ds.GetString(tag.PatientName)
	if !exists {
		t.Fatal("GetString() should find existing element")
	}
	if str != "Doe^John" {
		t.Errorf("GetString() = %q, want %q", str, "Doe^John")
	}

	// Non-existing tag
	str, exists = ds.GetString(tag.StudyDate)
	if exists {
		t.Error("GetString() should return false for non-existing tag")
	}
	if str != "" {
		t.Errorf("GetString() = %q, want empty string", str)
	}
}

// TestDatasetGetStrings tests GetStrings accessor
func TestDatasetGetStrings(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY", "AXIAL"}))

	strs, exists := ds.GetStrings(tag.ImageType)
	if !exists {
		t.Fatal("GetStrings() should find existing element")
	}
	if len(strs) != 3 {
		t.Errorf("len(GetStrings()) = %d, want 3", len(strs))
	}
	if strs[0] != "ORIGINAL" || strs[1] != "PRIMARY" || strs[2] != "AXIAL" {
		t.Errorf("GetStrings() = %v, want [ORIGINAL PRIMARY AXIAL]", strs)
	}
}

// TestDatasetGetUInt16 tests GetUInt16 accessor
func TestDatasetGetUInt16(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512, 256}))

	val, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil {
		t.Fatalf("GetUInt16() error = %v", err)
	}
	if val != 512 {
		t.Errorf("GetUInt16() = %d, want 512", val)
	}

	val, err = ds.GetUInt16(tag.Rows, 1)
	if err != nil {
		t.Fatalf("GetUInt16() error = %v", err)
	}
	if val != 256 {
		t.Errorf("GetUInt16() = %d, want 256", val)
	}

	// Non-existing tag
	_, err = ds.GetUInt16(tag.Columns, 0)
	if err == nil {
		t.Error("GetUInt16() should return error for non-existing tag")
	}
}

// TestDatasetGetUInt16s tests GetUInt16s accessor
func TestDatasetGetUInt16s(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512, 256}))

	vals, err := ds.GetUInt16s(tag.Rows)
	if err != nil {
		t.Fatalf("GetUInt16s() error = %v", err)
	}
	if len(vals) != 2 {
		t.Errorf("len(GetUInt16s()) = %d, want 2", len(vals))
	}
	if vals[0] != 512 || vals[1] != 256 {
		t.Errorf("GetUInt16s() = %v, want [512 256]", vals)
	}
}

// TestDatasetGetUInt32 tests GetUInt32 accessor
func TestDatasetGetUInt32(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewUnsignedLong(tag.FileMetaInformationGroupLength, []uint32{100}))

	val, err := ds.GetUInt32(tag.FileMetaInformationGroupLength, 0)
	if err != nil {
		t.Fatalf("GetUInt32() error = %v", err)
	}
	if val != 100 {
		t.Errorf("GetUInt32() = %d, want 100", val)
	}
}

// TestDatasetGetInt16 tests GetInt16 accessor
func TestDatasetGetInt16(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewSignedShort(tag.SmallestImagePixelValue, []int16{-100}))

	val, err := ds.GetInt16(tag.SmallestImagePixelValue, 0)
	if err != nil {
		t.Fatalf("GetInt16() error = %v", err)
	}
	if val != -100 {
		t.Errorf("GetInt16() = %d, want -100", val)
	}
}

// TestDatasetGetFloat32 tests GetFloat32 accessor
func TestDatasetGetFloat32(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewFloat(tag.ImagePositionPatient, []float32{1.5, 2.5, 3.5}))

	val, err := ds.GetFloat32(tag.ImagePositionPatient, 0)
	if err != nil {
		t.Fatalf("GetFloat32() error = %v", err)
	}
	if val != 1.5 {
		t.Errorf("GetFloat32() = %f, want 1.5", val)
	}
}

// TestDatasetGetFloat64 tests GetFloat64 accessor
func TestDatasetGetFloat64(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewDouble(tag.RealWorldValueSlope, []float64{1.234567}))

	val, err := ds.GetFloat64(tag.RealWorldValueSlope, 0)
	if err != nil {
		t.Fatalf("GetFloat64() error = %v", err)
	}
	if val != 1.234567 {
		t.Errorf("GetFloat64() = %f, want 1.234567", val)
	}
}

// TestDatasetTryGetters tests Try* accessor methods
func TestDatasetTryGetters(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	ds.Add(element.NewUnsignedLong(tag.FileMetaInformationGroupLength, []uint32{100}))

	// TryGetString
	str := ds.TryGetString(tag.PatientName)
	if str != "Doe^John" {
		t.Errorf("TryGetString() = %q, want %q", str, "Doe^John")
	}

	// TryGetString for non-existing tag
	str = ds.TryGetString(tag.StudyDate)
	if str != "" {
		t.Errorf("TryGetString() = %q, want empty string", str)
	}

	// TryGetUInt16
	val16 := ds.TryGetUInt16(tag.Rows, 0)
	if val16 != 512 {
		t.Errorf("TryGetUInt16() = %d, want 512", val16)
	}

	// TryGetUInt16 for non-existing tag
	val16 = ds.TryGetUInt16(tag.Columns, 0)
	if val16 != 0 {
		t.Errorf("TryGetUInt16() = %d, want 0", val16)
	}

	// TryGetUInt32
	val32 := ds.TryGetUInt32(tag.FileMetaInformationGroupLength, 0)
	if val32 != 100 {
		t.Errorf("TryGetUInt32() = %d, want 100", val32)
	}
}
