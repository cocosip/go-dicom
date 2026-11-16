// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package render

import (
	"bytes"
	"encoding/binary"
	"image/color"
	"image/png"
	"math"
	"testing"
)

func TestImageExporterRenderFrameSupportsRGBA(t *testing.T) {
	exporter := NewImageExporter(nil)
	var output bytes.Buffer
	if err := exporter.RenderFrameWithHighBit(
		&output,
		[]byte{10, 20, 30, 40},
		1, 1,
		8, 8, 7,
		4,
		false,
		"RGB",
		0,
		&ExportOptions{Format: FormatPNG},
	); err != nil {
		t.Fatalf("RenderFrameWithHighBit() error = %v", err)
	}

	decoded, err := png.Decode(bytes.NewReader(output.Bytes()))
	if err != nil {
		t.Fatalf("png.Decode() error = %v", err)
	}
	if got, want := color.NRGBAModel.Convert(decoded.At(0, 0)).(color.NRGBA), (color.NRGBA{R: 10, G: 20, B: 30, A: 40}); got != want {
		t.Fatalf("decoded pixel = %#v, want %#v", got, want)
	}
}

func TestImageExporterRenderGrayscaleOneBitHonorsSignedRepresentation(t *testing.T) {
	// Signed one-bit samples are two's complement: 0 and -1. With this window,
	// -1 is the low end and 0 the high end of the display range.
	pipeline := NewGrayscalePipeline(1, 0, -0.5, 1, -1, 0, false)
	img, err := NewImageExporter(pipeline).RenderGrayscaleImageWithBitDepth(
		[]byte{0b10}, 2, 1, 1, 1, 0, true, "MONOCHROME2",
	)
	if err != nil {
		t.Fatalf("RenderGrayscaleImageWithBitDepth() error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := img.GrayAt(x, 0).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestImageExporterRenderRGBImageSixteenBit(t *testing.T) {
	// Little-endian 16-bit RGB: 0xFFFF, 0x8000, 0x0000 then 0x1234, 0x00FF, 0x0100.
	data := []byte{
		0xFF, 0xFF, 0x00, 0x80, 0x00, 0x00,
		0x34, 0x12, 0xFF, 0x00, 0x00, 0x01,
	}
	img, err := NewImageExporter(nil).RenderRGBImageWithBytesPerSample(data, 2, 1, "RGB", 0, 2)
	if err != nil {
		t.Fatalf("RenderRGBImageWithBytesPerSample() error = %v", err)
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

func TestImageExporterRenderRGBAImageARGB(t *testing.T) {
	tests := []struct {
		name         string
		data         []byte
		planarConfig int
	}{
		{
			name:         "interleaved",
			data:         []byte{40, 10, 20, 30, 255, 100, 110, 120},
			planarConfig: 0,
		},
		{
			name:         "planar",
			data:         []byte{40, 255, 10, 100, 20, 110, 30, 120},
			planarConfig: 1,
		},
	}
	want := []color.NRGBA{
		{R: 10, G: 20, B: 30, A: 40},
		{R: 100, G: 110, B: 120, A: 255},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			img, err := NewImageExporter(nil).RenderRGBAImageWithConfig(tc.data, 2, 1, "ARGB", tc.planarConfig)
			if err != nil {
				t.Fatalf("RenderRGBAImageWithConfig() error = %v", err)
			}
			for x, expected := range want {
				if got := color.NRGBAModel.Convert(img.At(x, 0)).(color.NRGBA); got != expected {
					t.Fatalf("pixel %d = %#v, want %#v", x, got, expected)
				}
			}
		})
	}
}

func TestGrayscalePipelineCachesPrecalculatedLUT(t *testing.T) {
	pipeline := NewGrayscalePipeline(1, 0, 128, 256, 0, 65535, false)
	exporter := NewImageExporter(pipeline)

	first := exporter.precalculatedLUT(16, false)
	if first == nil {
		t.Fatal("precalculatedLUT(16) = nil, want a table")
	}
	if got := exporter.precalculatedLUT(16, false); got != first {
		t.Fatal("expected the pipeline to reuse the cached table")
	}
	if got := exporter.precalculatedLUT(8, false); got == first {
		t.Fatal("a different stored-value domain must rebuild the table")
	}

	pipeline.SetWindow(64, 512)
	afterWindow := exporter.precalculatedLUT(16, false)
	if afterWindow == first {
		t.Fatal("a window change must drop the cached table")
	}
	want := int(math.Round(pipeline.LUT().Transform(400)))
	if got := int(afterWindow.Transform(400)); got != want {
		t.Fatalf("cached table Transform(400) = %d, want %d", got, want)
	}
}

func TestGrayscaleRenderingUsesCurrentWindowAfterCaching(t *testing.T) {
	const size = 256 // 65536 pixels, enough to build the 16-bit table
	pipeline := NewGrayscalePipeline(1, 0, 128, 256, 0, 65535, false)
	exporter := NewImageExporter(pipeline)
	pixelData := make([]byte, size*size*2)
	binary.LittleEndian.PutUint16(pixelData, 400)

	before, err := exporter.RenderGrayscaleImageWithBitDepth(
		pixelData, size, size, 16, 16, 15, false, "MONOCHROME2",
	)
	if err != nil {
		t.Fatalf("first RenderGrayscaleImageWithBitDepth() error = %v", err)
	}
	pipeline.SetWindow(512, 1024)
	after, err := exporter.RenderGrayscaleImageWithBitDepth(
		pixelData, size, size, 16, 16, 15, false, "MONOCHROME2",
	)
	if err != nil {
		t.Fatalf("second RenderGrayscaleImageWithBitDepth() error = %v", err)
	}
	if before.GrayAt(0, 0).Y == after.GrayAt(0, 0).Y {
		t.Fatalf("window change did not affect the rendered value %d", before.GrayAt(0, 0).Y)
	}
}
