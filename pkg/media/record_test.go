// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
)

func TestRecordChildrenReturnsDefensiveCopy(t *testing.T) {
	parent := newRecord(RecordPatient, dataset.New(), 100)
	child := newRecord(RecordStudy, dataset.New(), 200)
	parent.children = []*Record{child}

	children := parent.Children()
	children[0] = nil

	got := parent.Children()
	if len(got) != 1 || got[0] != child {
		t.Fatal("caller mutation changed the record tree")
	}
}

func TestRecordPreservesUnknownTypeAndDataset(t *testing.T) {
	ds := dataset.New()
	record := newRecord(RecordType("PRIVATE TYPE"), ds, 1234)

	if got := record.Type(); got != RecordType("PRIVATE TYPE") {
		t.Fatalf("Type() = %q, want PRIVATE TYPE", got)
	}
	if record.Dataset() != ds {
		t.Fatal("Dataset() did not return the original record dataset")
	}
	if got := record.Offset(); got != 1234 {
		t.Fatalf("Offset() = %d, want 1234", got)
	}
}
