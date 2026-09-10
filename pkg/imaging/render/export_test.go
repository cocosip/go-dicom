// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
		8, 8,
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
