// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

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

// genericImage routes every pixel read through the generic image.Image path, so
// tests can compare the direct Pix fast paths against it.
type genericImage struct{ image.Image }

// TestApplyAffineFastSamplersMatchGenericPath verifies the direct Pix readers
// produce exactly the same pixels as color.RGBAModel on image.Image.At.
func TestApplyAffineFastSamplersMatchGenericPath(t *testing.T) {
	images := []struct {
		name   string
		source image.Image
	}{
		{"Gray", gradientGray()},
		{"Gray16", gradientGray16()},
		{"RGBA", gradientRGBA()},
		{"NRGBA", gradientNRGBA()},
	}
	matrices := []struct {
		name   string
		matrix transform.Affine2D
	}{
		{"rotate", transform.Rotate(90)},
		{"scale", transform.Scale(1.5, 0.75)},
	}
	modes := []interpolation.Mode{interpolation.ModeNearestNeighbor, interpolation.ModeBilinear}
	for _, imageCase := range images {
		for _, matrixCase := range matrices {
			for _, mode := range modes {
				t.Run(imageCase.name+"/"+matrixCase.name, func(t *testing.T) {
					fast, err := ApplyAffine(imageCase.source, matrixCase.matrix, image.Rectangle{}, mode, color.Black)
					if err != nil {
						t.Fatalf("ApplyAffine(fast) error = %v", err)
					}
					generic, err := ApplyAffine(genericImage{imageCase.source}, matrixCase.matrix, image.Rectangle{}, mode, color.Black)
					if err != nil {
						t.Fatalf("ApplyAffine(generic) error = %v", err)
					}
					if fast.Bounds() != generic.Bounds() {
						t.Fatalf("bounds = %v, want %v", fast.Bounds(), generic.Bounds())
					}
					for y := fast.Bounds().Min.Y; y < fast.Bounds().Max.Y; y++ {
						for x := fast.Bounds().Min.X; x < fast.Bounds().Max.X; x++ {
							if got, want := fast.At(x, y), generic.At(x, y); got != want {
								t.Fatalf("pixel (%d,%d) = %v, want %v", x, y, got, want)
							}
						}
					}
				})
			}
		}
	}
}

func gradientGray() *image.Gray {
	img := image.NewGray(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.SetGray(x, y, color.Gray{Y: uint8(x*16 + y*3)})
		}
	}
	return img
}

func gradientGray16() *image.Gray16 {
	img := image.NewGray16(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.SetGray16(x, y, color.Gray16{Y: uint16(x*4096 + y*777)})
		}
	}
	return img
}

func gradientRGBA() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			alpha := uint8(64 + x*20)
			img.SetRGBA(x, y, color.RGBA{R: alpha, G: alpha / 2, B: alpha / 3, A: alpha})
		}
	}
	return img
}

func gradientNRGBA() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.SetNRGBA(x, y, color.NRGBA{
				R: uint8(x * 30),
				G: uint8(y * 30),
				B: 128,
				A: uint8(64 + x*20),
			})
		}
	}
	return img
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
