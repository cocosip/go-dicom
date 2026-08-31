// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/logging"
)

func TestWriteWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	previous := logging.Logger()
	logging.SetLogger(slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { logging.SetLogger(previous) })

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
