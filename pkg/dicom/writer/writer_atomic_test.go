// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package writer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestWriteFileKeepsTargetWhenWriteFails verifies that a rejected write does
// not truncate or destroy a file that already exists at the target path.
func TestWriteFileKeepsTargetWhenWriteFails(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output.dcm")

	valid := dataset.New()
	addTestSOPUIDs(t, valid)
	if err := WriteFile(path, valid); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	original, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	// A dataset without SOP UIDs cannot be wrapped in File Meta Information.
	if err := WriteFile(path, dataset.New()); err == nil {
		t.Fatal("WriteFile() accepted a dataset without SOP Class and SOP Instance UID")
	}

	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() after failed write error = %v", err)
	}
	if string(after) != string(original) {
		t.Fatalf("failed write changed the target file (%d -> %d bytes)", len(original), len(after))
	}

	entries, err := os.ReadDir(filepath.Dir(path))
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}
	if len(entries) != 1 {
		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		t.Fatalf("temporary files left behind: %v", names)
	}
}

// TestWritePreservesNonConformantSOPInstanceUID verifies that a value read from
// a non-conformant source (PS3.5 9.1 forbids leading zeros in UID components) is
// copied into the File Meta Information instead of being rejected on write-back.
func TestWritePreservesNonConformantSOPInstanceUID(t *testing.T) {
	const invalidUID = "999.099.1.2.3"

	ds := dataset.New()
	ds.SetAutoValidate(false)
	for _, elem := range []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{invalidUID}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Add() error = %v", err)
		}
	}

	path := filepath.Join(t.TempDir(), "nonconformant.dcm")
	if err := WriteFile(path, ds); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	result, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	got, ok := result.FileMetaInformationDataset().GetString(tag.MediaStorageSOPInstanceUID)
	if !ok {
		t.Fatal("file meta information is missing MediaStorageSOPInstanceUID")
	}
	if strings.TrimRight(got, "\x00 ") != invalidUID {
		t.Fatalf("MediaStorageSOPInstanceUID = %q, want %q", got, invalidUID)
	}
}
