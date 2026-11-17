// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package imaging

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"runtime"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
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

// TestRenderFrameImageScansDerivedWindowRangeOnce verifies the fallback window
// range is derived from all pixels once and shared by every frame pipeline,
// instead of rescanning every frame for each of the N frames.
func TestRenderFrameImageScansDerivedWindowRangeOnce(t *testing.T) {
	const frames = 3
	pixels := make([]byte, 0, frames*4)
	for frame := 0; frame < frames; frame++ {
		pixels = append(pixels, byte(frame*10), 0, 0, 0)
	}
	ds := newNativeMonochromeDataset(t, 4, 1, pixels)
	if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"3"})); err != nil {
		t.Fatalf("add NumberOfFrames: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	for frame := 0; frame < frames; frame++ {
		if _, err := dicomImage.RenderFrameImage(frame); err != nil {
			t.Fatalf("RenderFrameImage(%d) error = %v", frame, err)
		}
	}
	if got := dicomImage.windowRangeScanCount; got != 1 {
		t.Fatalf("pixel range was scanned %d times for %d frames, want once", got, frames)
	}
}

// TestRenderFrameImageBorrowsFrameWithoutCopy verifies rendering reads the
// stored frame instead of allocating an equally sized copy per render.
func TestRenderFrameImageBorrowsFrameWithoutCopy(t *testing.T) {
	const width, height = 256, 256
	frameBytes := int64(width * height * 2)
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: width, Height: height, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}, make([]byte, frameBytes))
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}
	dicomImage := NewDicomImage(pixelData)
	// Warm the pipeline so one-time table construction is not measured.
	if _, err := dicomImage.RenderFrameImage(0); err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}

	const runs = 10
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	for i := 0; i < runs; i++ {
		if _, err := dicomImage.RenderFrameImage(0); err != nil {
			t.Fatalf("RenderFrameImage(0) error = %v", err)
		}
	}
	runtime.ReadMemStats(&after)
	if perRender := int64(after.TotalAlloc-before.TotalAlloc) / runs; perRender >= frameBytes {
		t.Fatalf("render allocated %d bytes per frame, want less than one frame copy (%d bytes)", perRender, frameBytes)
	}
}

// TestRenderFrameImageMasksPixelPadding verifies padding samples display black
// after the pipeline, including after MONOCHROME1 inversion.
func TestRenderFrameImageMasksPixelPadding(t *testing.T) {
	tests := []struct {
		name        string
		photometric string
		want        []uint8
	}{
		{"MONOCHROME2", pixel.Monochrome2.Value, []uint8{0, 0, 255}},
		{"MONOCHROME1", pixel.Monochrome1.Value, []uint8{255, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := paddingTestDataset(t, tt.photometric, []byte{0, 128, 255}, true)
			dicomImage, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage(0) error = %v", err)
			}
			for x, want := range tt.want {
				if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
					t.Fatalf("pixel %d = %d, want %d", x, got, want)
				}
			}
		})
	}

	t.Run("no padding value", func(t *testing.T) {
		ds := paddingTestDataset(t, pixel.Monochrome2.Value, []byte{0, 128, 255}, false)
		dicomImage, err := NewDicomImageFromDataset(ds)
		if err != nil {
			t.Fatalf("NewDicomImageFromDataset() error = %v", err)
		}
		rendered, err := dicomImage.RenderFrameImage(0)
		if err != nil {
			t.Fatalf("RenderFrameImage(0) error = %v", err)
		}
		for x, want := range []uint8{0, 128, 255} {
			if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
				t.Fatalf("pixel %d = %d, want %d", x, got, want)
			}
		}
	})
}

func paddingTestDataset(t *testing.T, photometric string, pixels []byte, withPadding bool) *dataset.Dataset {
	t.Helper()
	ds := newNativeMonochromeDataset(t, uint16(len(pixels)), 1, pixels)
	if photometric != pixel.Monochrome2.Value {
		ds.Remove(tag.PhotometricInterpretation)
		if err := ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometric})); err != nil {
			t.Fatalf("add PhotometricInterpretation: %v", err)
		}
	}
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{128}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	if withPadding {
		if err := ds.Add(element.NewUnsignedShort(tag.PixelPaddingValue, []uint16{128})); err != nil {
			t.Fatalf("add PixelPaddingValue: %v", err)
		}
	}
	return ds
}

// TestSetWindowClampsWidthBelowOne verifies PS3.3 C.11.2.1.2: a LINEAR window
// width below 1 is invalid, and accepting it inverts the display mapping.
func TestSetWindowClampsWidthBelowOne(t *testing.T) {
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: 2, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{0, 255})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}
	dicomImage := NewDicomImage(pixelData)

	dicomImage.SetWindow(128, 0.5)
	if got := dicomImage.WindowWidth(); got != 1 {
		t.Fatalf("WindowWidth() after SetWindow(..., 0.5) = %v, want 1", got)
	}
	clamped, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.GrayModel.Convert(clamped.At(0, 0)).(color.Gray).Y; got != 0 {
		t.Fatalf("darkest pixel = %d, want 0 (an inverted mapping would be 255)", got)
	}
	if got := color.GrayModel.Convert(clamped.At(1, 0)).(color.Gray).Y; got != 255 {
		t.Fatalf("brightest pixel = %d, want 255", got)
	}

	dicomImage.SetWindow(128, 1)
	reference, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	assertImagesEqual(t, clamped, reference)

	dicomImage.SetWindowWidth(0.25)
	if got := dicomImage.WindowWidth(); got != 1 {
		t.Fatalf("WindowWidth() after SetWindowWidth(0.25) = %v, want 1", got)
	}
}

// TestNilPixelDataDoesNotPanic verifies an image built without pixel data
// reports empty accessors and render errors instead of dereferencing nil.
func TestNilPixelDataDoesNotPanic(t *testing.T) {
	dicomImage := NewDicomImage(nil)
	if got := dicomImage.Width(); got != 0 {
		t.Fatalf("Width() = %d, want 0", got)
	}
	if got := dicomImage.Height(); got != 0 {
		t.Fatalf("Height() = %d, want 0", got)
	}
	if got := dicomImage.NumberOfFrames(); got != 0 {
		t.Fatalf("NumberOfFrames() = %d, want 0", got)
	}
	if _, err := dicomImage.RenderFrameImage(0); err == nil {
		t.Fatal("RenderFrameImage(0) = nil error, want a missing-pixel-data error")
	}
	if err := dicomImage.RenderFrame(io.Discard, 0, nil); err == nil {
		t.Fatal("RenderFrame() = nil error, want a missing-pixel-data error")
	}
	// Window setters use the fallback pipeline when no frame exists, so they
	// must not dereference the absent pixel data either (the value is not
	// retained because there is no frame pipeline to cache it in).
	dicomImage.SetWindow(128, 256)
	dicomImage.SetWindowCenter(128)
	dicomImage.SetWindowWidth(256)
	if got := dicomImage.WindowCenter(); got != got { // NaN
		t.Fatal("WindowCenter() = NaN after SetWindow on an image without pixel data")
	}
}
