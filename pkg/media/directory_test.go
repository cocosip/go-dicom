// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

func TestNewDirectoryCreatesRequiredMediaStorageMetadata(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}

	fmi := dir.FileMetaInformation()
	if got, ok := fmi.GetString(tag.MediaStorageSOPClassUID); !ok || got != uid.MediaStorageDirectoryStorage.UID() {
		t.Fatalf("MediaStorageSOPClassUID = %q, %v", got, ok)
	}
	if got, ok := fmi.GetString(tag.MediaStorageSOPInstanceUID); !ok || got == "" {
		t.Fatalf("MediaStorageSOPInstanceUID = %q, %v", got, ok)
	}
	if got, ok := fmi.GetString(tag.TransferSyntaxUID); !ok || got != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Fatalf("TransferSyntaxUID = %q, %v", got, ok)
	}

	ds := dir.Dataset()
	if got := ds.TryGetUInt32(tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, 0); got != 0 {
		t.Fatalf("first root offset = %d, want 0", got)
	}
	if got := ds.TryGetUInt32(tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, 0); got != 0 {
		t.Fatalf("last root offset = %d, want 0", got)
	}
	if got := ds.TryGetUInt16(tag.FileSetConsistencyFlag, 0); got != 0 {
		t.Fatalf("FileSetConsistencyFlag = %d, want 0", got)
	}
	sequence, err := ds.GetSequence(tag.DirectoryRecordSequence)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}
	if sequence.Count() != 0 {
		t.Fatalf("DirectoryRecordSequence count = %d, want 0", sequence.Count())
	}
	if got := ds.InternalTransferSyntax(); got != transfer.ExplicitVRLittleEndian {
		t.Fatalf("InternalTransferSyntax = %v, want Explicit VR Little Endian", got)
	}
}

func TestNewDirectoryAcceptsOnlyFoDicomDirectoryTransferSyntaxes(t *testing.T) {
	for _, ts := range []*transfer.Syntax{transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian} {
		if _, err := NewDirectory(WithTransferSyntax(ts)); err != nil {
			t.Fatalf("NewDirectory(%s) error = %v", ts, err)
		}
	}

	for _, ts := range []*transfer.Syntax{
		transfer.ExplicitVRBigEndian,
		transfer.DeflatedExplicitVRLittleEndian,
		transfer.JPEGBaseline8Bit,
	} {
		if _, err := NewDirectory(WithTransferSyntax(ts)); err == nil {
			t.Fatalf("NewDirectory(%s) succeeded, want error", ts)
		}
	}
}

func TestDirectoryRootRecordsAndDiagnosticsReturnDefensiveCopies(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}
	dir.roots = []*Record{newRecord(RecordPatient, nil, 0)}
	dir.diagnostics = []Diagnostic{{Code: DiagnosticOptionalAttributeMissing}}

	roots := dir.RootRecords()
	roots[0] = nil
	diagnostics := dir.Diagnostics()
	diagnostics[0].Code = DiagnosticIconGenerationFailed

	if dir.RootRecords()[0] == nil {
		t.Fatal("caller mutation changed root records")
	}
	if got := dir.Diagnostics()[0].Code; got != DiagnosticOptionalAttributeMissing {
		t.Fatalf("caller mutation changed diagnostic to %q", got)
	}
}
