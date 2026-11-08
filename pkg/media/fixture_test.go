// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// The fixture is copied from fo-dicom Tests/FO-DICOM.Tests/Test Data/DICOMDIR
// at SHA-256 7628497A488A8AFC14FEEE30888F914380AC7A23CA3E07E0341058F6D898D287.
// fo-dicom is distributed under the MS-PL; this repository is licensed under the MIT License.
func TestFoDicomFixtureOpensStrictly(t *testing.T) {
	dir, err := Open(foDicomFixturePath(), WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Open(fo-dicom fixture) error = %v", err)
	}
	if got := countDirectoryRecords(t, dir); got != 80 {
		t.Fatalf("record count = %d, want 80", got)
	}
	if len(dir.RootRecords()) == 0 {
		t.Fatal("fo-dicom fixture has no root records")
	}
}

func TestFoDicomFixtureRoundTripsConformantDirectoryTransferSyntax(t *testing.T) {
	dir, err := Open(foDicomFixturePath(), WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Open(fo-dicom fixture) error = %v", err)
	}

	var encoded bytes.Buffer
	if err := dir.Write(&encoded); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	opened, err := Read(bytes.NewReader(encoded.Bytes()), WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("strict Read(round trip) error = %v", err)
	}
	if got := countDirectoryRecords(t, opened); got != 80 {
		t.Fatalf("round-trip record count = %d, want 80", got)
	}
}

func TestFoDicomFixtureWriteRejectsImplicitVR(t *testing.T) {
	dir, err := Open(foDicomFixturePath(), WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Open(fo-dicom fixture) error = %v", err)
	}
	dir.transferSyntax = transfer.ImplicitVRLittleEndian
	dir.dataset.SetInternalTransferSyntax(transfer.ImplicitVRLittleEndian)
	fileMeta := dataset.NewFileMetaInformationFromDataset(dir.fileMeta)
	if err := fileMeta.SetTransferSyntax(transfer.ImplicitVRLittleEndian); err != nil {
		t.Fatalf("SetTransferSyntax() error = %v", err)
	}

	if err := dir.Write(&bytes.Buffer{}); err == nil {
		t.Fatal("Write() encoded an Implicit VR Little Endian DICOMDIR (PS3.10 §8.6)")
	}
}

func foDicomFixturePath() string {
	return filepath.Join("..", "..", "test-data", "DICOMDIR")
}

func countDirectoryRecords(t *testing.T, dir *Directory) int {
	t.Helper()
	count := 0
	if err := dir.Walk(func(*Record) error {
		count++
		return nil
	}); err != nil {
		t.Fatalf("Walk() error = %v", err)
	}
	return count
}
