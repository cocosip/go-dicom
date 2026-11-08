// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package imaging

import (
	"image/color"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// TestMonochrome1WithPresentationLUTShapeInverseInvertsOnce verifies that a
// MONOCHROME1 image whose Presentation LUT Shape is INVERSE is not inverted
// twice. The shape and the photometric interpretation encode the same display
// polarity, so the frame must match the MONOCHROME2 + INVERSE rendering.
func TestMonochrome1WithPresentationLUTShapeInverseInvertsOnce(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	if err := ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome1.Value})); err != nil {
		t.Fatalf("set PhotometricInterpretation: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	if err := ds.Add(element.NewString(tag.PresentationLUTShape, vr.CS, []string{lutShapeInverse})); err != nil {
		t.Fatalf("add PresentationLUTShape: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d (exactly one inversion)", x, got, want)
		}
	}
}

// TestMonochrome1WithoutPresentationLUTShapeStillInverts confirms the default
// MONOCHROME1 display polarity is unchanged when no presentation LUT is present.
func TestMonochrome1WithoutPresentationLUTShapeStillInverts(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	if err := ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome1.Value})); err != nil {
		t.Fatalf("set PhotometricInterpretation: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

