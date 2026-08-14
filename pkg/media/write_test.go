// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"bytes"
	"errors"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

func TestDirectoryWriteRoundTripsStrictly(t *testing.T) {
	for _, syntax := range []*transfer.Syntax{transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian} {
		t.Run(syntax.UID().UID(), func(t *testing.T) {
			dir, sources := roundTripDirectory(t, syntax)
			before := make([][]byte, len(sources))
			for i, source := range sources {
				before[i] = serializeTestFile(t, source)
			}

			var encoded bytes.Buffer
			if err := dir.Write(&encoded); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			for i, source := range sources {
				if after := serializeTestFile(t, source); !bytes.Equal(before[i], after) {
					t.Fatalf("Write() changed source DICOM file %d", i)
				}
			}

			opened, err := Read(bytes.NewReader(encoded.Bytes()), WithOffsetPolicy(StrictOffsets))
			if err != nil {
				t.Fatalf("strict Read() error = %v", err)
			}
			if opened.transferSyntax != syntax {
				t.Fatalf("transfer syntax = %v, want %v", opened.transferSyntax, syntax)
			}
			if got := recordTypes(opened.RootRecords()); !equalRecordTypes(got, []RecordType{RecordPatient, RecordPatient}) {
				t.Fatalf("root types = %v, want two PATIENT records", got)
			}
			firstSeries := onlyChild(t, onlyChild(t, opened.RootRecords()[0], RecordStudy), RecordSeries)
			if got := recordTypes(firstSeries.Children()); !equalRecordTypes(got, []RecordType{RecordImage, RecordImage}) {
				t.Fatalf("first series children = %v, want two IMAGE records", got)
			}
			count := 0
			if err := opened.Walk(func(record *Record) error {
				count++
				if record.Offset() == 0 {
					t.Fatalf("record %s has zero physical offset", record.Type())
				}
				return nil
			}); err != nil {
				t.Fatalf("Walk() error = %v", err)
			}
			if count != 9 {
				t.Fatalf("record count = %d, want 9", count)
			}
		})
	}
}

func TestDirectorySaveTruncatesAndReopens(t *testing.T) {
	dir, _ := roundTripDirectory(t, transfer.ExplicitVRLittleEndian)
	path := filepath.Join(t.TempDir(), "DICOMDIR")
	if err := os.WriteFile(path, bytes.Repeat([]byte{0xCC}, 64*1024), 0o600); err != nil {
		t.Fatalf("seed file: %v", err)
	}
	if err := dir.Save(path); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	opened, err := Open(path, WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	if len(opened.RootRecords()) != 2 {
		t.Fatalf("root count = %d, want 2", len(opened.RootRecords()))
	}
}

func TestDirectoryWriteRejectsInvalidRecordGraph(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Directory)
	}{
		{
			name: "duplicate root",
			mutate: func(dir *Directory) {
				dir.roots = append(dir.roots, dir.roots[0])
			},
		},
		{
			name: "cycle",
			mutate: func(dir *Directory) {
				dir.roots[0].children = append(dir.roots[0].children, dir.roots[0])
			},
		},
		{
			name: "nil record",
			mutate: func(dir *Directory) {
				dir.roots = append(dir.roots, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, _ := roundTripDirectory(t, transfer.ExplicitVRLittleEndian)
			tt.mutate(dir)
			if err := dir.Write(&bytes.Buffer{}); err == nil {
				t.Fatal("Write() succeeded with invalid record graph")
			}
		})
	}
}

func TestDirectoryWritePropagatesTargetWriterError(t *testing.T) {
	dir, _ := roundTripDirectory(t, transfer.ExplicitVRLittleEndian)
	want := errors.New("target write failed")
	if err := dir.Write(&errorAfterWriter{remaining: 200, err: want}); !errors.Is(err, want) {
		t.Fatalf("Write() error = %v, want %v", err, want)
	}
}

func TestCheckedItemOffsetRejectsULOverflow(t *testing.T) {
	if _, err := checkedItemOffset(uint64(math.MaxUint32) + 1); err == nil {
		t.Fatal("checkedItemOffset() accepted an offset larger than UL")
	}
}

func roundTripDirectory(t *testing.T, syntax *transfer.Syntax) (*Directory, []*parser.ParseResult) {
	t.Helper()
	dir, err := NewDirectory(WithTransferSyntax(syntax))
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}
	sources := []*parser.ParseResult{
		testParseResult(t, testFileValues{
			PatientID: "PAT1", PatientName: "One^Patient", StudyUID: "1.2.3.1",
			SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
		}),
		testParseResult(t, testFileValues{
			PatientID: "PAT1", PatientName: "One^Patient", StudyUID: "1.2.3.1",
			SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.2",
		}),
		testParseResult(t, testFileValues{
			PatientID: "PAT2", PatientName: "Two^Patient", StudyUID: "1.2.3.2",
			SeriesUID: "1.2.3.2.1", SOPInstanceUID: "1.2.3.2.1.1",
		}),
	}
	addTestFile(t, dir, sources[0], "DIR1/IMAGE001")
	addTestFile(t, dir, sources[1], "DIR1/IMAGE002")
	addTestFile(t, dir, sources[2], "DIR2/IMAGE001")
	return dir, sources
}

type errorAfterWriter struct {
	remaining int
	err       error
}

func (w *errorAfterWriter) Write(p []byte) (int, error) {
	if w.remaining <= 0 {
		return 0, w.err
	}
	if len(p) > w.remaining {
		n := w.remaining
		w.remaining = 0
		return n, w.err
	}
	w.remaining -= len(p)
	return len(p), nil
}
