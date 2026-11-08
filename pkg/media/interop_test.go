// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFoDicomInteropFiles(t *testing.T) {
	outputDirectory := os.Getenv("GO_DICOM_INTEROP_DIR")
	if outputDirectory == "" {
		t.Skip("GO_DICOM_INTEROP_DIR is not set")
	}
	if err := os.MkdirAll(outputDirectory, 0o700); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}

	dir, _ := roundTripDirectory(t)
	path := filepath.Join(outputDirectory, "explicit.dicomdir")
	if err := dir.Save(path); err != nil {
		t.Fatalf("Save(explicit.dicomdir) error = %v", err)
	}
	if _, err := Open(path, WithOffsetPolicy(StrictOffsets)); err != nil {
		t.Fatalf("strict Open(explicit.dicomdir) error = %v", err)
	}
}
