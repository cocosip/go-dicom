// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"image"
	"image/png"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

func TestRenderFrameImageMatchesExistingPNGExport(t *testing.T) {
	tests := []struct {
		name string
		info *PixelDataInfo
		data []byte
		kind any
	}{
		{
			name: photometricMonochrome1,
			info: &PixelDataInfo{
				Width: 2, Height: 2, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
				PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome1,
			},
			data: []byte{0, 64, 128, 255},
			kind: (*image.Gray)(nil),
		},
		{
			name: "RGB",
			info: &PixelDataInfo{
				Width: 2, Height: 1, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 3,
				PixelRepresentation: UnsignedPixels, PhotometricInterpretation: RGBPhotometric,
				PlanarConfiguration: InterleavedPlanar,
			},
			data: []byte{255, 0, 0, 0, 128, 255},
			kind: (*image.RGBA)(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixelData, err := NewDicomPixelDataFromBytes(tt.info, tt.data)
			if err != nil {
				t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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

func TestRenderFrameImageRejectsInvalidFrame(t *testing.T) {
	pixelData, err := NewDicomPixelDataFromBytes(&PixelDataInfo{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome2,
	}, []byte{1})
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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
