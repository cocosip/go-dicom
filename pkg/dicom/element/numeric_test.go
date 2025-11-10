// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// numericElement is an interface for testing numeric elements
type numericElement[T comparable] interface {
	Count() int
	Length() uint32
	GetValue(index int) (T, error)
	GetValues() ([]T, error)
}

// testNumericElement is a helper function to test numeric elements
func testNumericElement[T comparable](t *testing.T, elem numericElement[T], values []T, bytesPerValue int) {
	t.Helper()

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != len(values) {
			t.Errorf("Count() = %d, want %d", elem.Count(), len(values))
		}
	})

	t.Run("Length", func(t *testing.T) {
		want := uint32(len(values) * bytesPerValue)
		if elem.Length() != want {
			t.Errorf("Length() = %d, want %d", elem.Length(), want)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		for i, want := range values {
			got, err := elem.GetValue(i)
			if err != nil {
				t.Fatalf("GetValue(%d) error = %v", i, err)
			}
			if got != want {
				t.Errorf("GetValue(%d) = %v, want %v", i, got, want)
			}
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		got, err := elem.GetValues()
		if err != nil {
			t.Fatalf("GetValues() error = %v", err)
		}
		if len(got) != len(values) {
			t.Errorf("len(GetValues()) = %d, want %d", len(got), len(values))
		}
		for i, want := range values {
			if got[i] != want {
				t.Errorf("GetValues()[%d] = %v, want %v", i, got[i], want)
			}
		}
	})
}

// TestUnsignedShort tests US (Unsigned Short) elements
func TestUnsignedShort(t *testing.T) {
	values := []uint16{512, 512, 1}
	elem := element.NewUnsignedShort(tag.Rows, values)
	testNumericElement(t, elem, values, 2)

	t.Run("GetValue out of range", func(t *testing.T) {
		_, err := elem.GetValue(10)
		if err == nil {
			t.Error("GetValue(10) should return error for out of range index")
		}
	})
}

// TestUnsignedLong tests UL (Unsigned Long) elements
func TestUnsignedLong(t *testing.T) {
	values := []uint32{100, 200, 300}
	elem := element.NewUnsignedLong(tag.FileMetaInformationGroupLength, values)
	testNumericElement(t, elem, values, 4)
}

// TestSignedShort tests SS (Signed Short) elements
func TestSignedShort(t *testing.T) {
	values := []int16{-100, 0, 100}
	elem := element.NewSignedShort(tag.SmallestImagePixelValue, values)

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != len(values) {
			t.Errorf("Count() = %d, want %d", elem.Count(), len(values))
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		for i, want := range values {
			got, err := elem.GetValue(i)
			if err != nil {
				t.Fatalf("GetValue(%d) error = %v", i, err)
			}
			if got != want {
				t.Errorf("GetValue(%d) = %d, want %d", i, got, want)
			}
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		got, err := elem.GetValues()
		if err != nil {
			t.Fatalf("GetValues() error = %v", err)
		}
		if len(got) != len(values) {
			t.Errorf("len(GetValues()) = %d, want %d", len(got), len(values))
		}
		for i, want := range values {
			if got[i] != want {
				t.Errorf("GetValues()[%d] = %d, want %d", i, got[i], want)
			}
		}
	})
}

// TestSignedLong tests SL (Signed Long) elements
func TestSignedLong(t *testing.T) {
	values := []int32{-1000, 0, 1000}
	elem := element.NewSignedLong(tag.ReferencePixelX0, values)

	t.Run("Count", func(t *testing.T) {
		if elem.Count() != len(values) {
			t.Errorf("Count() = %d, want %d", elem.Count(), len(values))
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		for i, want := range values {
			got, err := elem.GetValue(i)
			if err != nil {
				t.Fatalf("GetValue(%d) error = %v", i, err)
			}
			if got != want {
				t.Errorf("GetValue(%d) = %d, want %d", i, got, want)
			}
		}
	})
}

// TestFloat tests FL (Floating Point Single) elements
func TestFloat(t *testing.T) {
	values := []float32{1.5, 2.5, 3.5}
	elem := element.NewFloat(tag.ImagePositionPatient, values)
	testNumericElement(t, elem, values, 4)
}

// TestDouble tests FD (Floating Point Double) elements
func TestDouble(t *testing.T) {
	values := []float64{1.234567890123, -2.345678901234, 0.0}
	elem := element.NewDouble(tag.RealWorldValueSlope, values)
	testNumericElement(t, elem, values, 8)
}

// TestSignedVeryLong tests SV (Signed 64-bit Very Long) elements
func TestSignedVeryLong(t *testing.T) {
	values := []int64{-9223372036854775808, 0, 9223372036854775807}
	elem := element.NewSignedVeryLong(tag.ReferencePixelX0, values)
	testNumericElement(t, elem, values, 8)
}

// TestUnsignedVeryLong tests UV (Unsigned 64-bit Very Long) elements
func TestUnsignedVeryLong(t *testing.T) {
	values := []uint64{0, 100, 18446744073709551615}
	elem := element.NewUnsignedVeryLong(tag.FileMetaInformationGroupLength, values)
	testNumericElement(t, elem, values, 8)
}

// TestNumericElement_Empty tests empty numeric elements
func TestNumericElement_Empty(t *testing.T) {
	t.Run("UnsignedShort", func(t *testing.T) {
		elem := element.NewUnsignedShort(tag.Rows, []uint16{})
		if elem.Count() != 0 {
			t.Errorf("Count() = %d, want 0", elem.Count())
		}
		values, err := elem.GetValues()
		if err != nil {
			t.Fatalf("GetValues() error = %v", err)
		}
		if values != nil {
			t.Errorf("GetValues() = %v, want nil", values)
		}
	})

	t.Run("Float", func(t *testing.T) {
		elem := element.NewFloat(tag.ImagePositionPatient, []float32{})
		if elem.Count() != 0 {
			t.Errorf("Count() = %d, want 0", elem.Count())
		}
	})
}
