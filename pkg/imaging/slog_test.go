// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/logging"
)

func TestRenderFrameImageWritesDebugSlogRecord(t *testing.T) {
	var output bytes.Buffer
	previous := logging.Logger()
	logging.SetLogger(slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { logging.SetLogger(previous) })

	pixelData, err := NewDicomPixelDataFromBytes(&PixelDataInfo{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome2,
	}, []byte{128})
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}
	if _, err := NewDicomImage(pixelData).RenderFrameImage(0); err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}

	got := output.String()
	for _, want := range []string{
		`"component":"imaging.render"`,
		`"event":"render_completed"`,
		`"frame":0`,
		`"width":1`,
		`"height":1`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("slog output missing %q: %s", want, got)
		}
	}
}

func TestDecodeIfNeededWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	previous := logging.Logger()
	logging.SetLogger(slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { logging.SetLogger(previous) })

	pixelData, err := NewDicomPixelDataFromBytes(&PixelDataInfo{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome2,
		TransferSyntaxUID: transfer.JPEG2000Lossless.UID().UID(),
	}, []byte{128})
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}
	if err := NewDicomImage(pixelData).DecodeIfNeeded(imagePassthroughCodec{}, codec.NewBaseParameters()); err != nil {
		t.Fatalf("DecodeIfNeeded() error = %v", err)
	}

	got := output.String()
	for _, want := range []string{
		`"component":"imaging.decode"`,
		`"event":"pixel_decode_completed"`,
		`"input_transfer_syntax":"1.2.840.10008.1.2.4.90"`,
		`"frame_count":1`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("slog output missing %q: %s", want, got)
		}
	}
}
