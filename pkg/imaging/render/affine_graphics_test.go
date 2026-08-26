// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"image"
	"image/color"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/interpolation"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
)

func TestApplyAffineRotateAndGraphicsComposition(t *testing.T) {
	source := image.NewRGBA(image.Rect(0, 0, 2, 1))
	source.SetRGBA(0, 0, color.RGBA{R: 255, A: 255})
	source.SetRGBA(1, 0, color.RGBA{G: 255, A: 255})

	rotated, err := ApplyAffine(source, transform.Rotate(90), image.Rectangle{}, interpolation.ModeNearestNeighbor, color.Transparent)
	if err != nil {
		t.Fatalf("ApplyAffine() error = %v", err)
	}
	if got, want := rotated.Bounds(), image.Rect(0, 0, 1, 2); got != want {
		t.Fatalf("bounds = %v, want %v", got, want)
	}
	if got := rotated.At(0, 0); got != (color.RGBA{R: 255, A: 255}) {
		t.Fatalf("top pixel = %#v, want red", got)
	}
	if got := rotated.At(0, 1); got != (color.RGBA{G: 255, A: 255}) {
		t.Fatalf("bottom pixel = %#v, want green", got)
	}

	overlay := image.NewRGBA(image.Rect(0, 0, 1, 1))
	overlay.SetRGBA(0, 0, color.RGBA{B: 255, A: 255})
	composed, err := DrawGraphics(rotated, []Graphic{ImageGraphic{Image: overlay, Position: image.Pt(0, 1)}})
	if err != nil {
		t.Fatalf("DrawGraphics() error = %v", err)
	}
	if got := composed.At(0, 1); got != (color.RGBA{B: 255, A: 255}) {
		t.Fatalf("graphic pixel = %#v, want blue", got)
	}
}

func TestCompositeGraphicClipsOutOfBoundsLayers(t *testing.T) {
	destination := image.NewRGBA(image.Rect(0, 0, 1, 1))
	overlay := image.NewUniform(color.RGBA{R: 255, A: 255})
	graphic := CompositeGraphic{Graphics: []Graphic{ImageGraphic{Image: overlay, Position: image.Pt(10, 10)}}}
	if err := graphic.Draw(destination); err != nil {
		t.Fatalf("Draw() error = %v", err)
	}
	if got := destination.At(0, 0); got != (color.RGBA{}) {
		t.Fatalf("destination changed to %#v", got)
	}
}

func TestApplyAffineUsesExplicitViewportCoordinatesForCropping(t *testing.T) {
	source := image.NewRGBA(image.Rect(0, 0, 2, 1))
	source.SetRGBA(0, 0, color.RGBA{R: 255, A: 255})
	source.SetRGBA(1, 0, color.RGBA{G: 255, A: 255})

	cropped, err := ApplyAffine(
		source,
		transform.Identity(),
		image.Rect(1, 0, 2, 1),
		interpolation.ModeNearestNeighbor,
		color.Transparent,
	)
	if err != nil {
		t.Fatalf("ApplyAffine() error = %v", err)
	}
	if got, want := cropped.Bounds(), image.Rect(1, 0, 2, 1); got != want {
		t.Fatalf("bounds = %v, want %v", got, want)
	}
	if got, want := color.RGBAModel.Convert(cropped.At(1, 0)).(color.RGBA), (color.RGBA{G: 255, A: 255}); got != want {
		t.Fatalf("cropped pixel = %#v, want %#v", got, want)
	}
}

func BenchmarkApplyAffine512Gray16(b *testing.B) {
	source := image.NewGray16(image.Rect(0, 0, 512, 512))
	matrix := transform.Rotate(90)
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_, _ = ApplyAffine(source, matrix, image.Rectangle{}, interpolation.ModeBilinear, color.Black)
	}
}

func BenchmarkApplyAffine512RGB(b *testing.B) {
	source := image.NewRGBA(image.Rect(0, 0, 512, 512))
	matrix := transform.Rotate(90)
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_, _ = ApplyAffine(source, matrix, image.Rectangle{}, interpolation.ModeBilinear, color.Black)
	}
}
