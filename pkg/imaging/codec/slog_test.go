// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
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

func TestTranscodeWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	previous := logging.Logger()
	logging.SetLogger(slog.New(slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { logging.SetLogger(previous) })

	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Private^Patient"}))
	transcoder := NewTranscoder(transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian)
	if _, err := transcoder.Transcode(ds); err != nil {
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
