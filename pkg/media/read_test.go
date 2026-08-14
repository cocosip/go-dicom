// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func TestReadBuildsExactOffsetDirectoryTree(t *testing.T) {
	encoded := encodeDirectoryForReadTest(t, nil)

	dir, err := Read(bytes.NewReader(encoded), WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if dir.transferSyntax != transfer.ExplicitVRLittleEndian {
		t.Fatalf("transfer syntax = %v, want Explicit VR Little Endian", dir.transferSyntax)
	}

	roots := dir.RootRecords()
	if len(roots) != 1 || roots[0].Type() != RecordPatient {
		t.Fatalf("roots = %#v, want one PATIENT", recordTypes(roots))
	}
	patient := roots[0]
	study := onlyChild(t, patient, RecordStudy)
	series := onlyChild(t, study, RecordSeries)
	children := series.Children()
	if got, want := recordTypes(children), []RecordType{RecordImage, RecordType("PRIVATE")}; !equalRecordTypes(got, want) {
		t.Fatalf("series children = %v, want %v", got, want)
	}
	if patient.Offset() == 0 || study.Offset() == 0 || series.Offset() == 0 || children[0].Offset() == 0 {
		t.Fatal("parsed records did not retain physical Item offsets")
	}
	if got := patient.Dataset().TryGetString(tag.PatientID); got != "PAT001" {
		t.Fatalf("PatientID = %q, want PAT001", got)
	}
	if got, ok := children[0].Dataset().GetStrings(tag.ReferencedFileID); !ok || len(got) != 2 || got[0] != "IMAGES" || got[1] != "IMG0001" {
		t.Fatalf("ReferencedFileID = %v, %v", got, ok)
	}

	diagnostics := dir.Diagnostics()
	if len(diagnostics) != 1 || diagnostics[0].Code != DiagnosticUnknownRecordType || diagnostics[0].RecordType != RecordType("PRIVATE") {
		t.Fatalf("diagnostics = %#v, want one unknown-record diagnostic", diagnostics)
	}
}

func TestOpenReadsDICOMDIRFromPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "DICOMDIR")
	if err := os.WriteFile(path, encodeDirectoryForReadTest(t, nil), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	dir, err := Open(path, WithOffsetPolicy(StrictOffsets))
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	if len(dir.RootRecords()) != 1 {
		t.Fatalf("root count = %d, want 1", len(dir.RootRecords()))
	}
}

func TestReadRejectsNonDirectoryStorageSOPClass(t *testing.T) {
	encoded := encodeDirectoryForReadTest(t, func(fixture *readTestFixture) {
		if err := fixture.fileMeta.AddOrUpdate(element.NewString(tag.MediaStorageSOPClassUID, vr.UI, []string{"1.2.840.10008.1.3.11"})); err != nil {
			t.Fatalf("replace MediaStorageSOPClassUID: %v", err)
		}
	})

	_, err := Read(bytes.NewReader(encoded))
	if err == nil || !strings.Contains(err.Error(), "Media Storage Directory Storage") {
		t.Fatalf("Read() error = %v, want directory storage SOP Class error", err)
	}
}

func TestReadCompatibleRepairsUniqueFixedOffsetDelta(t *testing.T) {
	encoded := encodeDirectoryForReadTest(t, func(fixture *readTestFixture) {
		shiftReadTestReferences(t, fixture, 2)
	})

	if _, err := Read(bytes.NewReader(encoded), WithOffsetPolicy(StrictOffsets)); err == nil {
		t.Fatal("strict Read() succeeded with shifted offsets")
	}
	dir, err := Read(bytes.NewReader(encoded), WithOffsetPolicy(CompatibleOffsets))
	if err != nil {
		t.Fatalf("compatible Read() error = %v", err)
	}
	series := onlyChild(t, onlyChild(t, dir.RootRecords()[0], RecordStudy), RecordSeries)
	if got := recordTypes(series.Children()); !equalRecordTypes(got, []RecordType{RecordImage, RecordType("PRIVATE")}) {
		t.Fatalf("series children = %v", got)
	}
	if got := dir.dataset.TryGetUInt32(tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, 0); got != dir.roots[0].offset {
		t.Fatalf("repaired root offset = %d, want %d", got, dir.roots[0].offset)
	}

	repairs := 0
	for _, diagnostic := range dir.Diagnostics() {
		if diagnostic.Code != DiagnosticOffsetRepaired {
			continue
		}
		repairs++
		if diagnostic.OriginalOffset != diagnostic.RepairedOffset+2 {
			t.Fatalf("repair = %d -> %d, want fixed delta 2", diagnostic.OriginalOffset, diagnostic.RepairedOffset)
		}
		if strings.Contains(diagnostic.Message, "PAT001") || strings.Contains(diagnostic.Message, "IMG0001") {
			t.Fatalf("repair diagnostic contains identifying data: %q", diagnostic.Message)
		}
	}
	if repairs != 6 {
		t.Fatalf("offset repair diagnostics = %d, want 6", repairs)
	}
}

func TestReadCompatibleRecoversUniquePhysicalTypeHierarchy(t *testing.T) {
	encoded := encodeDirectoryForReadTest(t, func(fixture *readTestFixture) {
		if err := fixture.items["private"].AddOrUpdate(element.NewString(tag.DirectoryRecordType, vr.CS, []string{string(RecordImage)})); err != nil {
			t.Fatalf("replace DirectoryRecordType: %v", err)
		}
		setReadTestOffset(t, fixture.dataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, fixture.offsets["patient"]+2)
		setReadTestOffset(t, fixture.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, fixture.offsets["patient"]+4)
		setReadTestOffset(t, fixture.items["patient"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, fixture.offsets["study"]+6)
		setReadTestOffset(t, fixture.items["study"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, fixture.offsets["series"]+8)
		setReadTestOffset(t, fixture.items["series"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, fixture.offsets["image"]+10)
		setReadTestOffset(t, fixture.items["image"], tag.OffsetOfTheNextDirectoryRecord, fixture.offsets["private"]+12)
	})

	dir, err := Read(bytes.NewReader(encoded), WithOffsetPolicy(CompatibleOffsets))
	if err != nil {
		t.Fatalf("compatible Read() error = %v", err)
	}
	series := onlyChild(t, onlyChild(t, dir.RootRecords()[0], RecordStudy), RecordSeries)
	if got := recordTypes(series.Children()); !equalRecordTypes(got, []RecordType{RecordImage, RecordImage}) {
		t.Fatalf("series children = %v, want two IMAGE records", got)
	}
	if len(dir.Diagnostics()) != 6 {
		t.Fatalf("diagnostics = %d, want 6 repairs", len(dir.Diagnostics()))
	}
}

func TestReadRejectsStructuralCorruptionInBothPolicies(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*readTestFixture)
	}{
		{
			name: "cycle",
			mutate: func(fixture *readTestFixture) {
				setReadTestOffset(t, fixture.items["image"], tag.OffsetOfTheNextDirectoryRecord, fixture.offsets["image"])
			},
		},
		{
			name: "duplicate reference",
			mutate: func(fixture *readTestFixture) {
				setReadTestOffset(t, fixture.items["study"], tag.OffsetOfTheNextDirectoryRecord, fixture.offsets["series"])
			},
		},
		{
			name: "unreachable item",
			mutate: func(fixture *readTestFixture) {
				setReadTestOffset(t, fixture.items["image"], tag.OffsetOfTheNextDirectoryRecord, 0)
			},
		},
		{
			name: "invalid last root",
			mutate: func(fixture *readTestFixture) {
				setReadTestOffset(t, fixture.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, fixture.offsets["image"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := encodeDirectoryForReadTest(t, tt.mutate)
			for _, policy := range []OffsetPolicy{StrictOffsets, CompatibleOffsets} {
				if _, err := Read(bytes.NewReader(encoded), WithOffsetPolicy(policy)); err == nil {
					t.Fatalf("Read(policy=%d) succeeded", policy)
				}
			}
		})
	}
}

type readTestFixture struct {
	dataset  *dataset.Dataset
	fileMeta *dataset.Dataset
	items    map[string]*dataset.Dataset
	offsets  map[string]uint32
}

func encodeDirectoryForReadTest(t *testing.T, mutate func(*readTestFixture)) []byte {
	t.Helper()

	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}
	items := map[string]*dataset.Dataset{
		"patient": readTestRecord(t, RecordPatient,
			element.NewString(tag.PatientID, vr.LO, []string{"PAT001"})),
		"study": readTestRecord(t, RecordStudy,
			element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.1"})),
		"series": readTestRecord(t, RecordSeries,
			element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.1.1"})),
		"image": readTestRecord(t, RecordImage,
			element.NewString(tag.ReferencedFileID, vr.CS, []string{"IMAGES", "IMG0001"})),
		"private": readTestRecord(t, RecordType("PRIVATE")),
	}
	sequence := dataset.NewSequenceWithItems(tag.DirectoryRecordSequence, []*dataset.Dataset{
		items["patient"], items["study"], items["series"], items["image"], items["private"],
	})
	if err := dir.dataset.AddOrUpdate(sequence); err != nil {
		t.Fatalf("replace DirectoryRecordSequence: %v", err)
	}

	offsets := observeDirectoryOffsets(t, io.Discard, dir.dataset, dir.fileMeta, items)
	setReadTestOffset(t, dir.dataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, offsets["patient"])
	setReadTestOffset(t, dir.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, offsets["patient"])
	setReadTestOffset(t, items["patient"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, offsets["study"])
	setReadTestOffset(t, items["study"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, offsets["series"])
	setReadTestOffset(t, items["series"], tag.OffsetOfReferencedLowerLevelDirectoryEntity, offsets["image"])
	setReadTestOffset(t, items["image"], tag.OffsetOfTheNextDirectoryRecord, offsets["private"])

	fixture := &readTestFixture{dataset: dir.dataset, fileMeta: dir.fileMeta, items: items, offsets: offsets}
	if mutate != nil {
		mutate(fixture)
	}

	var encoded bytes.Buffer
	actual := observeDirectoryOffsets(t, &encoded, dir.dataset, dir.fileMeta, items)
	for name, offset := range offsets {
		if actual[name] != offset {
			t.Fatalf("record %s moved from %d to %d after offset update", name, offset, actual[name])
		}
	}
	return encoded.Bytes()
}

func readTestRecord(t *testing.T, recordType RecordType, extra ...element.Element) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	items := []element.Element{
		element.NewUnsignedLong(tag.OffsetOfTheNextDirectoryRecord, []uint32{0}),
		element.NewUnsignedShort(tag.RecordInUseFlag, []uint16{0xFFFF}),
		element.NewUnsignedLong(tag.OffsetOfReferencedLowerLevelDirectoryEntity, []uint32{0}),
		element.NewString(tag.DirectoryRecordType, vr.CS, []string{string(recordType)}),
	}
	items = append(items, extra...)
	for _, item := range items {
		if err := ds.Add(item); err != nil {
			t.Fatalf("record Add(%s): %v", item.Tag(), err)
		}
	}
	return ds
}

func observeDirectoryOffsets(t *testing.T, output io.Writer, ds, fileMeta *dataset.Dataset, items map[string]*dataset.Dataset) map[string]uint32 {
	t.Helper()
	names := make(map[*dataset.Dataset]string, len(items))
	for name, item := range items {
		names[item] = name
	}
	offsets := make(map[string]uint32, len(items))
	err := writer.Write(output, ds,
		writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian),
		writer.WithFileMetaInfo(fileMeta),
		writer.WithSequenceItemObserver(func(position writer.SequenceItemPosition) error {
			if position.SequenceTag.Equals(tag.DirectoryRecordSequence) {
				offsets[names[position.Item]] = uint32(position.Offset)
			}
			return nil
		}),
	)
	if err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	return offsets
}

func setReadTestOffset(t *testing.T, ds *dataset.Dataset, offsetTag *tag.Tag, value uint32) {
	t.Helper()
	if err := ds.AddOrUpdate(element.NewUnsignedLong(offsetTag, []uint32{value})); err != nil {
		t.Fatalf("set %s: %v", offsetTag, err)
	}
}

func shiftReadTestReferences(t *testing.T, fixture *readTestFixture, delta uint32) {
	t.Helper()
	for _, reference := range []struct {
		ds  *dataset.Dataset
		tag *tag.Tag
	}{
		{fixture.dataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity},
		{fixture.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity},
		{fixture.items["patient"], tag.OffsetOfReferencedLowerLevelDirectoryEntity},
		{fixture.items["study"], tag.OffsetOfReferencedLowerLevelDirectoryEntity},
		{fixture.items["series"], tag.OffsetOfReferencedLowerLevelDirectoryEntity},
		{fixture.items["image"], tag.OffsetOfTheNextDirectoryRecord},
	} {
		value := reference.ds.TryGetUInt32(reference.tag, 0)
		setReadTestOffset(t, reference.ds, reference.tag, value+delta)
	}
}

func onlyChild(t *testing.T, parent *Record, wantType RecordType) *Record {
	t.Helper()
	children := parent.Children()
	if len(children) != 1 || children[0].Type() != wantType {
		t.Fatalf("%s children = %v, want one %s", parent.Type(), recordTypes(children), wantType)
	}
	return children[0]
}

func recordTypes(records []*Record) []RecordType {
	types := make([]RecordType, len(records))
	for i, record := range records {
		types[i] = record.Type()
	}
	return types
}

func equalRecordTypes(a, b []RecordType) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
