// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/logging"
)

type imagingLogContextKey struct{}

func TestRenderFrameImageWritesDebugSlogRecord(t *testing.T) {
	var output bytes.Buffer
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

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
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

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

func TestRenderFrameImageContextPassesContextToLogHandler(t *testing.T) {
	handler := &imagingContextHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	pixelData, err := NewDicomPixelDataFromBytes(&PixelDataInfo{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome2,
	}, []byte{128})
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}
	ctx := context.WithValue(context.Background(), imagingLogContextKey{}, "render-request")
	if _, err := NewDicomImage(pixelData).RenderFrameImageContext(ctx, 0); err != nil {
		t.Fatalf("RenderFrameImageContext() error = %v", err)
	}

	if handler.contextValue != "render-request" {
		t.Fatalf("handler context value = %#v, want render-request", handler.contextValue)
	}
}

type imagingContextHandler struct {
	contextValue any
}

func (h *imagingContextHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *imagingContextHandler) Handle(ctx context.Context, _ slog.Record) error {
	h.contextValue = ctx.Value(imagingLogContextKey{})
	return nil
}

func (h *imagingContextHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *imagingContextHandler) WithGroup(string) slog.Handler { return h }
