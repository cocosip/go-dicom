// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

func TestWriteFoDicomInteropFiles(t *testing.T) {
	outputDirectory := os.Getenv("GO_DICOM_INTEROP_DIR")
	if outputDirectory == "" {
		t.Skip("GO_DICOM_INTEROP_DIR is not set")
	}
	if err := os.MkdirAll(outputDirectory, 0o700); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}

	for _, fixture := range []struct {
		name   string
		syntax *transfer.Syntax
	}{
		{name: "explicit.dicomdir", syntax: transfer.ExplicitVRLittleEndian},
		{name: "implicit.dicomdir", syntax: transfer.ImplicitVRLittleEndian},
	} {
		dir, _ := roundTripDirectory(t, fixture.syntax)
		path := filepath.Join(outputDirectory, fixture.name)
		if err := dir.Save(path); err != nil {
			t.Fatalf("Save(%s) error = %v", fixture.name, err)
		}
		if _, err := Open(path, WithOffsetPolicy(StrictOffsets)); err != nil {
			t.Fatalf("strict Open(%s) error = %v", fixture.name, err)
		}
	}
}
