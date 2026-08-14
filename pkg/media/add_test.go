// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func TestAddFileGroupsRecordsAndReturnsExistingDuplicate(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}

	first := testParseResult(t, testFileValues{
		PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
		SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
	})
	second := testParseResult(t, testFileValues{
		PatientID: "PAT1", PatientName: "Pat^Name^^^", StudyUID: "1.2.3.1",
		SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.2",
	})
	differentName := testParseResult(t, testFileValues{
		PatientID: "PAT1", PatientName: "PAT^Name^", StudyUID: "1.2.3.2",
		SeriesUID: "1.2.3.2.1", SOPInstanceUID: "1.2.3.2.1.1",
	})

	entry1 := addTestFile(t, dir, first, "DIR1/IMAGE001")
	entry2 := addTestFile(t, dir, second, "DIR1/IMAGE002")
	entry3 := addTestFile(t, dir, differentName, "DIR2/IMAGE001")
	duplicate := addTestFile(t, dir, first, "DIR1/IMAGE001")

	if entry1.Patient != entry2.Patient || entry1.Study != entry2.Study || entry1.Series != entry2.Series {
		t.Fatal("equivalent patient name and identical UIDs did not reuse records")
	}
	if entry1.Patient == entry3.Patient {
		t.Fatal("different patient name content reused the patient record")
	}
	if duplicate.Instance != entry1.Instance {
		t.Fatal("duplicate SOP Instance UID created another instance record")
	}
	if got := len(dir.RootRecords()); got != 2 {
		t.Fatalf("root records = %d, want 2", got)
	}
	if got := len(entry1.Series.Children()); got != 2 {
		t.Fatalf("instance records = %d, want 2", got)
	}

	values, ok := entry1.Instance.Dataset().GetStrings(tag.ReferencedFileID)
	if !ok || len(values) != 2 || values[0] != "DIR1" || values[1] != "IMAGE001" {
		t.Fatalf("ReferencedFileID = %#v, %v", values, ok)
	}
}

func TestAddFileSelectsFoDicomInstanceRecordTypes(t *testing.T) {
	tests := []struct {
		name     string
		sopClass *uid.UID
		wantType RecordType
	}{
		{name: "image", sopClass: uid.CTImageStorage, wantType: RecordImage},
		{name: "structured report", sopClass: uid.BasicTextSRStorage, wantType: RecordSRDocument},
		{name: "presentation state", sopClass: uid.GrayscaleSoftcopyPresentationStateStorage, wantType: RecordPresentation},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := NewDirectory()
			if err != nil {
				t.Fatalf("NewDirectory() error = %v", err)
			}
			file := testParseResult(t, testFileValues{
				PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
				SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1." + string(rune('1'+i)),
				SOPClass: tt.sopClass,
			})
			entry := addTestFile(t, dir, file, "DIR1/FILE0001")
			if got := entry.Instance.Type(); got != tt.wantType {
				t.Fatalf("instance type = %q, want %q", got, tt.wantType)
			}
		})
	}
}

func TestAddFileRejectsMissingRequiredUIDs(t *testing.T) {
	tests := []struct {
		name string
		tag  *tag.Tag
	}{
		{name: "study", tag: tag.StudyInstanceUID},
		{name: "series", tag: tag.SeriesInstanceUID},
		{name: "instance", tag: tag.SOPInstanceUID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := NewDirectory()
			if err != nil {
				t.Fatalf("NewDirectory() error = %v", err)
			}
			file := testParseResult(t, testFileValues{
				PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
				SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
			})
			file.Dataset.Remove(tt.tag)
			id, idErr := ParseFileID("DIR1/IMAGE001")
			if idErr != nil {
				t.Fatalf("ParseFileID() error = %v", idErr)
			}
			if _, err := dir.AddFile(file, id); err == nil {
				t.Fatalf("AddFile() without %s succeeded", tt.tag)
			}
		})
	}
}

func TestAddFileDoesNotMutateSourceFile(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}
	file := testParseResult(t, testFileValues{
		PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
		SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
	})
	before := serializeTestFile(t, file)

	addTestFile(t, dir, file, "DIR1/IMAGE001")

	after := serializeTestFile(t, file)
	if !bytes.Equal(before, after) {
		t.Fatal("AddFile() changed the source DICOM file")
	}
}

func TestAddFileGeneratesIconUsingRepresentativeFrame(t *testing.T) {
	tests := []struct {
		name           string
		representative uint16
		wantFrame      int
	}{
		{name: "valid representative frame", representative: 4, wantFrame: 3},
		{name: "invalid representative frame", representative: 9, wantFrame: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := &iconGeneratorStub{width: 2, height: 1, pixels: []byte{10, 20}}
			dir, err := NewDirectory(WithImageIcons(true), WithIconGenerator(generator))
			if err != nil {
				t.Fatalf("NewDirectory() error = %v", err)
			}
			file := testParseResult(t, testFileValues{
				PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
				SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
			})
			if err := file.Dataset.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"6"})); err != nil {
				t.Fatalf("add NumberOfFrames: %v", err)
			}
			if err := file.Dataset.Add(element.NewUnsignedShort(tag.RepresentativeFrameNumber, []uint16{tt.representative})); err != nil {
				t.Fatalf("add RepresentativeFrameNumber: %v", err)
			}
			before := serializeTestFile(t, file)

			entry := addTestFile(t, dir, file, "DIR1/IMAGE001")
			if generator.frame != tt.wantFrame {
				t.Fatalf("generator frame = %d, want %d", generator.frame, tt.wantFrame)
			}
			sequence, err := entry.Instance.Dataset().GetSequence(tag.IconImageSequence)
			if err != nil {
				t.Fatalf("GetSequence(IconImageSequence) error = %v", err)
			}
			if sequence.GetItem(0).TryGetUInt16(tag.Columns, 0) != 2 || sequence.GetItem(0).TryGetUInt16(tag.Rows, 0) != 1 {
				t.Fatal("generated icon dimensions were not attached")
			}
			if after := serializeTestFile(t, file); !bytes.Equal(before, after) {
				t.Fatal("icon generation changed the source DICOM file")
			}
		})
	}
}

func TestAddFileContinuesWhenIconGenerationFails(t *testing.T) {
	want := errors.New("sensitive generator detail")
	dir, err := NewDirectory(
		WithImageIcons(true),
		WithIconGenerator(&iconGeneratorStub{err: want}),
	)
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}
	file := testParseResult(t, testFileValues{
		PatientID: "PAT1", PatientName: "Pat^Name", StudyUID: "1.2.3.1",
		SeriesUID: "1.2.3.1.1", SOPInstanceUID: "1.2.3.1.1.1",
	})
	before := serializeTestFile(t, file)

	entry := addTestFile(t, dir, file, "DIR1/IMAGE001")
	if entry.Instance.Dataset().Contains(tag.IconImageSequence) {
		t.Fatal("failed icon generation attached IconImageSequence")
	}
	diagnostics := dir.Diagnostics()
	if len(diagnostics) == 0 || diagnostics[len(diagnostics)-1].Code != DiagnosticIconGenerationFailed {
		t.Fatalf("diagnostics = %#v, want icon generation failure", diagnostics)
	}
	if bytes.Contains([]byte(diagnostics[len(diagnostics)-1].Message), []byte(want.Error())) {
		t.Fatal("icon diagnostic exposed generator error details")
	}
	if after := serializeTestFile(t, file); !bytes.Equal(before, after) {
		t.Fatal("failed icon generation changed the source DICOM file")
	}
}

type testFileValues struct {
	PatientID      string
	PatientName    string
	StudyUID       string
	SeriesUID      string
	SOPInstanceUID string
	SOPClass       *uid.UID
}

func testParseResult(t *testing.T, values testFileValues) *parser.ParseResult {
	t.Helper()
	if values.SOPClass == nil {
		values.SOPClass = uid.CTImageStorage
	}

	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	elements := []element.Element{
		element.NewString(tag.PatientID, vr.LO, []string{values.PatientID}),
		element.NewPersonName(tag.PatientName, []string{values.PatientName}),
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{values.StudyUID}),
		element.NewString(tag.StudyID, vr.SH, []string{"STUDY1"}),
		element.NewString(tag.SeriesInstanceUID, vr.UI, []string{values.SeriesUID}),
		element.NewString(tag.Modality, vr.CS, []string{"CT"}),
		element.NewString(tag.SOPClassUID, vr.UI, []string{values.SOPClass.UID()}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{values.SOPInstanceUID}),
		element.NewString(tag.InstanceNumber, vr.IS, []string{"1"}),
	}
	for _, item := range elements {
		if err := ds.Add(item); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", item.Tag(), err)
		}
	}

	fmi := dataset.NewDefaultFileMetaInformation()
	if err := fmi.SetMediaStorageSOPClassUID(values.SOPClass.UID()); err != nil {
		t.Fatalf("SetMediaStorageSOPClassUID() error = %v", err)
	}
	if err := fmi.SetMediaStorageSOPInstanceUID(values.SOPInstanceUID); err != nil {
		t.Fatalf("SetMediaStorageSOPInstanceUID() error = %v", err)
	}
	if err := fmi.SetTransferSyntax(transfer.ExplicitVRLittleEndian); err != nil {
		t.Fatalf("SetTransferSyntax() error = %v", err)
	}

	return &parser.ParseResult{
		FileMetaInformation: fmi,
		Dataset:             ds,
		TransferSyntax:      transfer.ExplicitVRLittleEndian,
		Format:              parser.FormatDICOM3,
	}
}

func addTestFile(t *testing.T, dir *Directory, file *parser.ParseResult, path string) *Entry {
	t.Helper()
	id, err := ParseFileID(path)
	if err != nil {
		t.Fatalf("ParseFileID() error = %v", err)
	}
	entry, err := dir.AddFile(file, id)
	if err != nil {
		t.Fatalf("AddFile() error = %v", err)
	}
	return entry
}

func serializeTestFile(t *testing.T, file *parser.ParseResult) []byte {
	t.Helper()
	var encoded bytes.Buffer
	if err := writer.Write(&encoded, file.Dataset,
		writer.WithTransferSyntax(file.TransferSyntax),
		writer.WithFileMetaInfo(file.FileMetaInformationDataset()),
	); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	return append([]byte(nil), encoded.Bytes()...)
}

type iconGeneratorStub struct {
	width, height int
	pixels        []byte
	frame         int
	err           error
}

func (g *iconGeneratorStub) GenerateDirectoryIcon(_ *dataset.Dataset, frame int) (int, int, []byte, error) {
	g.frame = frame
	return g.width, g.height, append([]byte(nil), g.pixels...), g.err
}
