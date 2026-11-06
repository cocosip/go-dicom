// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package imaging

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

func TestRenderFrameImageMatchesExistingPNGExport(t *testing.T) {
	tests := []struct {
		name string
		info *pixeldata.Info
		data []byte
		kind any
	}{
		{
			name: pixel.Monochrome1.Value,
			info: &pixeldata.Info{
				Width: 2, Height: 2, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
				PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome1,
			},
			data: []byte{0, 64, 128, 255},
			kind: (*image.Gray)(nil),
		},
		{
			name: "RGB",
			info: &pixeldata.Info{
				Width: 2, Height: 1, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 3,
				PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.RGBPhotometric,
				PlanarConfiguration: pixel.InterleavedPlanar,
			},
			data: []byte{255, 0, 0, 0, 128, 255},
			kind: (*image.RGBA)(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixelData, err := pixeldata.NewFromBytes(tt.info, tt.data)
			if err != nil {
				t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
			}
			dicomImage := NewDicomImage(pixelData)
			got, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage() error = %v", err)
			}
			switch tt.kind.(type) {
			case *image.Gray:
				if _, ok := got.(*image.Gray); !ok {
					t.Fatalf("RenderFrameImage() = %T, want *image.Gray", got)
				}
			case *image.RGBA:
				if _, ok := got.(*image.RGBA); !ok {
					t.Fatalf("RenderFrameImage() = %T, want *image.RGBA", got)
				}
			}

			var encoded bytes.Buffer
			if err := dicomImage.RenderFrame(&encoded, 0, &render.ExportOptions{Format: render.FormatPNG}); err != nil {
				t.Fatalf("RenderFrame() error = %v", err)
			}
			want, err := png.Decode(bytes.NewReader(encoded.Bytes()))
			if err != nil {
				t.Fatalf("png.Decode() error = %v", err)
			}
			assertImagesEqual(t, got, want)
		})
	}
}

func TestRenderFrameImageARGB(t *testing.T) {
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: 2, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 4,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.ARGBPhotometric,
		PlanarConfiguration: pixel.InterleavedPlanar,
	}, []byte{40, 10, 20, 30, 255, 100, 110, 120})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}
	rendered, err := NewDicomImage(pixelData).RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}
	img, ok := rendered.(*image.NRGBA)
	if !ok {
		t.Fatalf("RenderFrameImage() = %T, want *image.NRGBA", rendered)
	}
	want := []color.NRGBA{
		{R: 10, G: 20, B: 30, A: 40},
		{R: 100, G: 110, B: 120, A: 255},
	}
	for x, expected := range want {
		if got := color.NRGBAModel.Convert(img.At(x, 0)).(color.NRGBA); got != expected {
			t.Fatalf("pixel %d = %#v, want %#v", x, got, expected)
		}
	}
}

func TestRenderFrameImageRGBSixteenBit(t *testing.T) {
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: 2, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15, SamplesPerPixel: 3,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.RGBPhotometric,
		PlanarConfiguration: pixel.InterleavedPlanar,
	}, []byte{
		0xFF, 0xFF, 0x00, 0x80, 0x00, 0x00,
		0x34, 0x12, 0xFF, 0x00, 0x00, 0x01,
	})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}
	rendered, err := NewDicomImage(pixelData).RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}
	img, ok := rendered.(*image.RGBA)
	if !ok {
		t.Fatalf("RenderFrameImage() = %T, want *image.RGBA", rendered)
	}
	want := []color.RGBA{
		{R: 0xFF, G: 0x80, B: 0x00, A: 0xFF},
		{R: 0x12, G: 0x00, B: 0x01, A: 0xFF},
	}
	for x, expected := range want {
		if got := color.RGBAModel.Convert(img.At(x, 0)).(color.RGBA); got != expected {
			t.Fatalf("pixel %d = %#v, want %#v", x, got, expected)
		}
	}
}

func TestRenderFrameImageRejectsInvalidFrame(t *testing.T) {
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{1})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}
	if _, err := NewDicomImage(pixelData).RenderFrameImage(1); err == nil {
		t.Fatal("RenderFrameImage() accepted an out-of-range frame")
	}
}

func assertImagesEqual(t *testing.T, got, want image.Image) {
	t.Helper()
	if got.Bounds() != want.Bounds() {
		t.Fatalf("bounds = %v, want %v", got.Bounds(), want.Bounds())
	}
	for y := got.Bounds().Min.Y; y < got.Bounds().Max.Y; y++ {
		for x := got.Bounds().Min.X; x < got.Bounds().Max.X; x++ {
			gr, gg, gb, ga := got.At(x, y).RGBA()
			wr, wg, wb, wa := want.At(x, y).RGBA()
			if gr != wr || gg != wg || gb != wb || ga != wa {
				t.Fatalf("pixel (%d,%d) = (%d,%d,%d,%d), want (%d,%d,%d,%d)", x, y, gr, gg, gb, ga, wr, wg, wb, wa)
			}
		}
	}
}
