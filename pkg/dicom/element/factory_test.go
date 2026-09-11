// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewElementFromValuePreservesVRAndVM(t *testing.T) {
	elem, err := NewElementFromValue(tag.WindowCenter, vr.DS, []float64{0, 100})
	if err != nil {
		t.Fatalf("NewElementFromValue() error = %v", err)
	}
	if _, ok := elem.(*DecimalString); !ok {
		t.Fatalf("element type = %T, want *DecimalString", elem)
	}
	if elem.Count() != 2 {
		t.Fatalf("element count = %d, want 2", elem.Count())
	}
}

func TestNewElementFromValueRejectsStringForBinaryVR(t *testing.T) {
	if _, err := NewElementFromValue(tag.SamplesPerPixel, vr.US, "1"); err == nil {
		t.Fatal("NewElementFromValue() accepted string for US")
	}
}
