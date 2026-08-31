// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/logging"
)

func TestParseWritesSafeSlogRecord(t *testing.T) {
	var output bytes.Buffer
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelDebug}),
	}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	var file bytes.Buffer
	file.Write(make([]byte, 128))
	file.WriteString("DICM")
	writeExplicitStringElement(&file, tag.TransferSyntaxUID, "UI", []byte(testExplicitVRLittleLE+"\x00"))
	writeExplicitStringElement(&file, tag.PatientName, "PN", []byte("Private^Patient"))
	if _, err := Parse(bytes.NewReader(file.Bytes())); err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	got := output.String()
	for _, want := range []string{
		`"component":"dicom.parser"`,
		`"event":"parse_completed"`,
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
