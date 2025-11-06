// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestDecimalString tests DS (Decimal String) elements
func TestDecimalString(t *testing.T) {
	t.Run("FromString", func(t *testing.T) {
		elem := element.NewDecimalString(tag.SliceThickness, []string{"1.5", "2.0", "3.14159"})

		f, err := elem.GetFloat(0)
		if err != nil {
			t.Fatalf("GetFloat(0) error = %v", err)
		}
		if f != 1.5 {
			t.Errorf("GetFloat(0) = %f, want 1.5", f)
		}

		f, err = elem.GetFloat(2)
		if err != nil {
			t.Fatalf("GetFloat(2) error = %v", err)
		}
		if f != 3.14159 {
			t.Errorf("GetFloat(2) = %f, want 3.14159", f)
		}
	})

	t.Run("FromFloat", func(t *testing.T) {
		elem := element.NewDecimalStringFromFloat(tag.SliceThickness, []float64{1.5, 2.0, 3.14159})

		floats, err := elem.GetFloats()
		if err != nil {
			t.Fatalf("GetFloats() error = %v", err)
		}

		if len(floats) != 3 {
			t.Errorf("len(GetFloats()) = %d, want 3", len(floats))
		}
		if floats[0] != 1.5 {
			t.Errorf("GetFloats()[0] = %f, want 1.5", floats[0])
		}
	})

	t.Run("GetFloat32", func(t *testing.T) {
		elem := element.NewDecimalString(tag.SliceThickness, []string{"1.5"})

		f, err := elem.GetFloat32(0)
		if err != nil {
			t.Fatalf("GetFloat32() error = %v", err)
		}
		if f != 1.5 {
			t.Errorf("GetFloat32() = %f, want 1.5", f)
		}
	})

	t.Run("InvalidFloat", func(t *testing.T) {
		elem := element.NewDecimalString(tag.SliceThickness, []string{"invalid"})

		_, err := elem.GetFloat(0)
		if err == nil {
			t.Error("GetFloat() should return error for invalid float")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		elem := element.NewDecimalString(tag.SliceThickness, []string{})

		floats, err := elem.GetFloats()
		if err != nil {
			t.Fatalf("GetFloats() error = %v", err)
		}
		if floats != nil {
			t.Errorf("GetFloats() = %v, want nil", floats)
		}
	})
}

// TestIntegerString tests IS (Integer String) elements
func TestIntegerString(t *testing.T) {
	t.Run("FromString", func(t *testing.T) {
		elem := element.NewIntegerString(tag.NumberOfFrames, []string{"1", "100", "-50"})

		i, err := elem.GetInt(0)
		if err != nil {
			t.Fatalf("GetInt(0) error = %v", err)
		}
		if i != 1 {
			t.Errorf("GetInt(0) = %d, want 1", i)
		}

		i, err = elem.GetInt(1)
		if err != nil {
			t.Fatalf("GetInt(1) error = %v", err)
		}
		if i != 100 {
			t.Errorf("GetInt(1) = %d, want 100", i)
		}

		i, err = elem.GetInt(2)
		if err != nil {
			t.Fatalf("GetInt(2) error = %v", err)
		}
		if i != -50 {
			t.Errorf("GetInt(2) = %d, want -50", i)
		}
	})

	t.Run("FromInt", func(t *testing.T) {
		elem := element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{1, 100, -50})

		ints, err := elem.GetInts()
		if err != nil {
			t.Fatalf("GetInts() error = %v", err)
		}

		if len(ints) != 3 {
			t.Errorf("len(GetInts()) = %d, want 3", len(ints))
		}
		if ints[0] != 1 {
			t.Errorf("GetInts()[0] = %d, want 1", ints[0])
		}
	})

	t.Run("GetInt32", func(t *testing.T) {
		elem := element.NewIntegerStringFromInt32(tag.NumberOfFrames, []int32{100})

		i, err := elem.GetInt32(0)
		if err != nil {
			t.Fatalf("GetInt32() error = %v", err)
		}
		if i != 100 {
			t.Errorf("GetInt32() = %d, want 100", i)
		}
	})

	t.Run("GetInt64", func(t *testing.T) {
		elem := element.NewIntegerString(tag.NumberOfFrames, []string{"9223372036854775807"})

		i, err := elem.GetInt64(0)
		if err != nil {
			t.Fatalf("GetInt64() error = %v", err)
		}
		if i != 9223372036854775807 {
			t.Errorf("GetInt64() = %d, want 9223372036854775807", i)
		}
	})

	t.Run("InvalidInt", func(t *testing.T) {
		elem := element.NewIntegerString(tag.NumberOfFrames, []string{"invalid"})

		_, err := elem.GetInt(0)
		if err == nil {
			t.Error("GetInt() should return error for invalid int")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		elem := element.NewIntegerString(tag.NumberOfFrames, []string{})

		ints, err := elem.GetInts()
		if err != nil {
			t.Fatalf("GetInts() error = %v", err)
		}
		if ints != nil {
			t.Errorf("GetInts() = %v, want nil", ints)
		}
	})
}
