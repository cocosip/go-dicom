// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset_test

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestDatasetAddValuePixelDataErrorPointsToHighLevelAPI(t *testing.T) {
	ds := dataset.New()
	err := ds.AddValue(tag.PixelData, []byte{1, 2})
	if err == nil {
		t.Fatal("AddValue(PixelData) error = nil, want ambiguous VR error")
	}
	if !strings.Contains(err.Error(), "pixeldata.NewForDataset") || !strings.Contains(err.Error(), "AddValueWithVR") {
		t.Fatalf("AddValue(PixelData) error = %q, want both high-level and explicit VR guidance", err)
	}
}

func TestDatasetAddValueCreatesTypedElements(t *testing.T) {
	ds := dataset.New()
	if err := ds.AddValue(tag.RescaleSlope, 1.0); err != nil {
		t.Fatalf("AddValue(RescaleSlope) error = %v", err)
	}
	elem, ok := ds.Get(tag.RescaleSlope)
	if !ok {
		t.Fatal("RescaleSlope was not added")
	}
	if _, ok := elem.(*element.DecimalString); !ok {
		t.Fatalf("RescaleSlope type = %T, want *element.DecimalString", elem)
	}

	if err := ds.AddValue(tag.WindowCenter, []float64{0, 100}); err != nil {
		t.Fatalf("AddValue(WindowCenter) error = %v", err)
	}
	elem, ok = ds.Get(tag.WindowCenter)
	if !ok {
		t.Fatal("WindowCenter was not added")
	}
	decimal, ok := elem.(*element.DecimalString)
	if !ok {
		t.Fatalf("WindowCenter type = %T, want *element.DecimalString", elem)
	}
	if got := decimal.Count(); got != 2 {
		t.Fatalf("WindowCenter VM = %d, want 2", got)
	}
	if got, ok := ds.GetString(tag.WindowCenter); !ok || got != "0\\100" {
		t.Fatalf("GetString(WindowCenter) = %q, %v, want 0\\100, true", got, ok)
	}
	if got, ok := ds.GetStrings(tag.WindowCenter); !ok || len(got) != 2 || got[1] != "100" {
		t.Fatalf("GetStrings(WindowCenter) = %v, %v, want [0 100], true", got, ok)
	}
	if err := ds.AddValue(tag.RecommendedDisplayFrameRateInFloat, float64(30)); err != nil {
		t.Fatalf("AddValue(FL) error = %v", err)
	}
	if _, ok := ds.Get(tag.RecommendedDisplayFrameRateInFloat); !ok {
		t.Fatal("RecommendedDisplayFrameRateInFloat was not added")
	}
}

func TestDatasetAddOrUpdateValueReplacesElement(t *testing.T) {
	ds := dataset.New()
	if err := ds.AddOrUpdateValue(tag.WindowWidth, 500.0); err != nil {
		t.Fatalf("AddOrUpdateValue(initial) error = %v", err)
	}
	if err := ds.AddOrUpdateValue(tag.WindowWidth, []float64{400, 800}); err != nil {
		t.Fatalf("AddOrUpdateValue(replacement) error = %v", err)
	}
	elem, ok := ds.Get(tag.WindowWidth)
	if !ok {
		t.Fatal("WindowWidth was not added")
	}
	decimal, ok := elem.(*element.DecimalString)
	if !ok {
		t.Fatalf("WindowWidth type = %T, want *element.DecimalString", elem)
	}
	if got := decimal.GetValues(); len(got) != 2 || got[0] != "400" || got[1] != "800" {
		t.Fatalf("WindowWidth values = %v, want [400 800]", got)
	}
}

func TestDatasetAddRejectsTagVRMismatch(t *testing.T) {
	ds := dataset.New()
	err := ds.Add(element.NewDecimalString(tag.SamplesPerPixel, []string{"1"}))
	if err == nil {
		t.Fatal("Add() accepted DS element for SamplesPerPixel, want VR mismatch error")
	}
}

func TestDatasetAddValueWithVRRejectsDisallowedVR(t *testing.T) {
	ds := dataset.New()
	if err := ds.AddValueWithVR(tag.SamplesPerPixel, vr.DS, uint16(1)); err == nil {
		t.Fatal("AddValueWithVR() accepted disallowed DS VR for SamplesPerPixel")
	}
}

func TestDatasetAutomaticValidationCanBeDisabledGlobally(t *testing.T) {
	previous := dataset.AutoValidate()
	t.Cleanup(func() { dataset.SetAutoValidate(previous) })
	dataset.SetAutoValidate(false)

	ds := dataset.New()
	if err := ds.Add(element.NewDecimalString(tag.SamplesPerPixel, []string{"1"})); err != nil {
		t.Fatalf("Add() should honor the global validation switch: %v", err)
	}
	if err := ds.AddValueWithVR(tag.Rows, vr.DS, "1"); err != nil {
		t.Fatalf("AddValueWithVR() should honor the global validation switch: %v", err)
	}
	if err := ds.Validate(); err == nil {
		t.Fatal("explicit Validate() should still reject invalid Tag/VR combinations")
	}
}
