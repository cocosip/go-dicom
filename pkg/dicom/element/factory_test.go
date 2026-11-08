// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package element

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
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

func TestNewElementFromBufferRejectsMisalignedNumericValue(t *testing.T) {
	if _, err := NewElementFromBuffer(tag.Rows, vr.US, buffer.NewMemory([]byte{1, 2, 3}), nil); err == nil {
		t.Fatal("NewElementFromBuffer() accepted a 3-byte US value; want error")
	}
	if _, err := NewElementFromBuffer(tag.Rows, vr.US, buffer.NewMemory([]byte{1, 2}), nil); err != nil {
		t.Fatalf("NewElementFromBuffer() rejected a valid US value: %v", err)
	}
}
