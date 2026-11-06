// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package render

import (
	"bytes"
	"image/color"
	"image/png"
	"testing"
)

func TestImageExporterRenderFrameSupportsRGBA(t *testing.T) {
	exporter := NewImageExporter(nil)
	var output bytes.Buffer
	if err := exporter.RenderFrame(
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
		t.Fatalf("RenderFrame() error = %v", err)
	}

	decoded, err := png.Decode(bytes.NewReader(output.Bytes()))
	if err != nil {
		t.Fatalf("png.Decode() error = %v", err)
	}
	if got, want := color.NRGBAModel.Convert(decoded.At(0, 0)).(color.NRGBA), (color.NRGBA{R: 10, G: 20, B: 30, A: 40}); got != want {
		t.Fatalf("decoded pixel = %#v, want %#v", got, want)
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
