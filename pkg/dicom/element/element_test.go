// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// TestElementInterface verifies that all element types implement the Element interface
func TestElementInterface(_ *testing.T) {
	var _ element.Element = element.NewString(tag.PatientName, vr.PN, []string{"Test"})
	var _ element.Element = element.NewUnsignedShort(tag.Rows, []uint16{512})
	var _ element.Element = element.NewUnsignedLong(tag.FileMetaInformationGroupLength, []uint32{100})
	var _ element.Element = element.NewSignedShort(tag.SmallestImagePixelValue, []int16{-100})
	var _ element.Element = element.NewSignedLong(tag.ReferencePixelX0, []int32{-1000})
	var _ element.Element = element.NewFloat(tag.ImagePositionPatient, []float32{1.5, 2.5, 3.5})
	var _ element.Element = element.NewDouble(tag.RealWorldValueSlope, []float64{1.234567890123})
}

// TestBaseElement tests the base element functionality
func TestBaseElement(t *testing.T) {
	data := []byte("Test Patient")
	buf := buffer.NewMemory(data)
	elem := element.NewStringFromBuffer(tag.PatientName, vr.PN, buf, nil)

	t.Run("Tag", func(t *testing.T) {
		if elem.Tag() != tag.PatientName {
			t.Errorf("Tag() = %v, want %v", elem.Tag(), tag.PatientName)
		}
	})

	t.Run("ValueRepresentation", func(t *testing.T) {
		if elem.ValueRepresentation() != vr.PN {
			t.Errorf("ValueRepresentation() = %v, want %v", elem.ValueRepresentation(), vr.PN)
		}
	})

    t.Run("Length", func(t *testing.T) {
        // len(data) is small and within uint32 range; conversion is safe.
        want := uint32(len(data)) // #nosec G115 -- test data size within uint32 range
        if elem.Length() != want {
            t.Errorf("Length() = %d, want %d", elem.Length(), want)
        }
    })

	t.Run("Buffer", func(t *testing.T) {
		if elem.Buffer() != buf {
			t.Error("Buffer() returned different buffer")
		}
	})

	t.Run("String", func(t *testing.T) {
		str := elem.String()
		if str == "" {
			t.Error("String() returned empty string")
		}
	})
}
