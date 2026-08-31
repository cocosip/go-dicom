// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/logging"
)

type writerLogContextKey struct{}

func TestWriteWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Private^Patient"}))
	var encoded bytes.Buffer
	if err := Write(&encoded, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	got := output.String()
	for _, want := range []string{
		`"component":"dicom.writer"`,
		`"event":"write_completed"`,
		`"transfer_syntax":"1.2.840.10008.1.2.1"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("slog output missing %q: %s", want, got)
		}
	}
	if strings.Contains(got, "Private^Patient") {
		t.Fatalf("slog output leaked patient data: %s", got)
	}
}

func TestWriteContextPassesContextToLogHandler(t *testing.T) {
	handler := &writerContextHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	ctx := context.WithValue(context.Background(), writerLogContextKey{}, "writer-request")
	if err := WriteContext(ctx, &bytes.Buffer{}, ds); err != nil {
		t.Fatalf("WriteContext() error = %v", err)
	}

	if handler.contextValue != "writer-request" {
		t.Fatalf("handler context value = %#v, want writer-request", handler.contextValue)
	}
}

type writerContextHandler struct {
	contextValue any
}

func (h *writerContextHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *writerContextHandler) Handle(ctx context.Context, _ slog.Record) error {
	h.contextValue = ctx.Value(writerLogContextKey{})
	return nil
}

func (h *writerContextHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *writerContextHandler) WithGroup(string) slog.Handler { return h }
