// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/logging"
)

type transcoderLogContextKey struct{}

func TestTranscodeWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Private^Patient"}))
	transcoder := newTestTranscoder(t, transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian)
	if _, err := transcoder.Transcode(context.Background(), ds); err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	got := output.String()
	for _, want := range []string{
		`"component":"imaging.codec"`,
		`"event":"transcode_completed"`,
		`"input_transfer_syntax":"1.2.840.10008.1.2.1"`,
		`"output_transfer_syntax":"1.2.840.10008.1.2"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("slog output missing %q: %s", want, got)
		}
	}
	if strings.Contains(got, "Private^Patient") {
		t.Fatalf("slog output leaked patient data: %s", got)
	}
}

func TestTranscodePassesContextToLogHandler(t *testing.T) {
	handler := &transcoderContextHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	ctx := context.WithValue(context.Background(), transcoderLogContextKey{}, "transcode-request")
	transcoder := newTestTranscoder(t, transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian)
	if _, err := transcoder.Transcode(ctx, dataset.New()); err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if handler.contextValue != "transcode-request" {
		t.Fatalf("handler context value = %#v, want transcode-request", handler.contextValue)
	}
}

func TestTranscodeCompatibleNativeOBWritesWarning(t *testing.T) {
	var output bytes.Buffer
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelWarn}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16})); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12})); err != nil {
		t.Fatal(err)
	}

	transcoder := newTestTranscoder(t, transfer.ExplicitVRLittleEndian, transfer.ExplicitVRLittleEndian)
	if _, err := transcoder.Transcode(context.Background(), ds); err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	got := output.String()
	if count := strings.Count(got, `"event":"nonstandard_native_pixel_data_vr"`); count != 1 {
		t.Fatalf("native Pixel Data VR warnings = %d, want 1: %s", count, got)
	}
	if !strings.Contains(got, `"component":"imaging.codec"`) {
		t.Fatalf("warning log missing component: %s", got)
	}
}

type transcoderContextHandler struct {
	contextValue any
}

func (h *transcoderContextHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *transcoderContextHandler) Handle(ctx context.Context, _ slog.Record) error {
	h.contextValue = ctx.Value(transcoderLogContextKey{})
	return nil
}

func (h *transcoderContextHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *transcoderContextHandler) WithGroup(string) slog.Handler { return h }
