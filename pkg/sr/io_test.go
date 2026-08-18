// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func TestOpenAndSaveRoundTripSharedSRFixture(t *testing.T) {
	sourcePath := filepath.Join("..", "..", "test-data", "test_SR.dcm")
	report, err := Open(sourcePath)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	outputPath := filepath.Join(t.TempDir(), "round-trip.dcm")
	if err := report.Save(outputPath); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	roundTripped, err := Open(outputPath)
	if err != nil {
		t.Fatalf("Open(round trip) error = %v", err)
	}
	children, err := roundTripped.Children()
	if err != nil {
		t.Fatalf("Children() error = %v", err)
	}
	if len(children) == 0 {
		t.Fatal("round-tripped report has no content items")
	}

	seen := map[string]bool{}
	maxDepth := 0
	var walk func(*ContentItem, int)
	walk = func(item *ContentItem, depth int) {
		t.Helper()
		if depth > maxDepth {
			maxDepth = depth
		}
		if item.Dataset().Contains(tag.ReferencedContentItemIdentifier) {
			seen["by-reference"] = true
			return
		}
		valueType, err := item.ValueType()
		if err != nil {
			t.Fatalf("ValueType() at depth %d: %v", depth, err)
		}
		switch valueType {
		case ValueTypeUIDReference:
			if _, err := item.GetUIDReference(); err != nil {
				t.Fatalf("GetUIDReference() error = %v", err)
			}
			seen["UIDREF"] = true
		case ValueTypeSpatialCoordinate:
			if _, err := item.GetSpatialCoordinate(); err != nil {
				t.Fatalf("GetSpatialCoordinate() error = %v", err)
			}
			seen["SCOORD"] = true
		case ValueTypeTemporalCoordinate:
			if _, err := item.GetTemporalCoordinate(); err != nil {
				t.Fatalf("GetTemporalCoordinate() error = %v", err)
			}
			seen["TCOORD"] = true
		}
		children, err := item.Children()
		if err != nil {
			t.Fatalf("Children() at depth %d: %v", depth, err)
		}
		for _, child := range children {
			walk(child, depth+1)
		}
	}
	walk(roundTripped.ContentItem, 0)
	for _, required := range []string{"UIDREF", "SCOORD", "TCOORD", "by-reference"} {
		if !seen[required] {
			t.Errorf("round-tripped fo-dicom fixture did not expose %s content", required)
		}
	}
	if maxDepth < 3 {
		t.Errorf("round-tripped report maximum depth = %d, want at least 3", maxDepth)
	}
}

func TestReadRejectsPartialParse(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "test-data", "test_SR.dcm"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	_, err = Read(bytes.NewReader(data), parser.WithStopAtTag(tag.ContentSequence))
	if err == nil || !strings.Contains(err.Error(), "partial") {
		t.Fatalf("Read() error = %v, want partial parse error", err)
	}
}

func TestWriteValidatesBeforeWriting(t *testing.T) {
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"))
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	report.Dataset().Remove(tag.ValueType)
	var output bytes.Buffer

	err = report.Write(&output)
	if err == nil || !strings.Contains(err.Error(), "value type not found") {
		t.Fatalf("Write() error = %v, want validation error", err)
	}
	if output.Len() != 0 {
		t.Fatalf("Write() wrote %d bytes before validation failed", output.Len())
	}
}

func TestWritePreservesFileMetaAndTransferSyntax(t *testing.T) {
	report, err := Open(filepath.Join("..", "..", "test-data", "test_SR.dcm"))
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	var output bytes.Buffer
	if err := report.Write(&output); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	parsed, err := parser.Parse(bytes.NewReader(output.Bytes()))
	if err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if got := parsed.TransferSyntax.UID().String(); got != "1.2.840.10008.1.2.1" {
		t.Fatalf("Transfer Syntax UID = %q, want Explicit VR Little Endian", got)
	}
	version, ok := parsed.FileMetaInformation.ImplementationVersionName()
	if !ok || version != "OFFIS_DCMTK_342" {
		t.Fatalf("Implementation Version Name = %q, %v, want preserved OFFIS_DCMTK_342", version, ok)
	}
}

func TestWriteForcesExplicitSequenceAndItemLengths(t *testing.T) {
	report, err := Open(filepath.Join("..", "..", "test-data", "test_SR.dcm"))
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	var output bytes.Buffer
	if err := report.Write(&output, writer.WithExplicitLengthSequences(false)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	data := output.Bytes()
	contentSequenceTag := []byte{0x40, 0x00, 0x30, 0xa7}
	sequenceOffset := bytes.Index(data, contentSequenceTag)
	if sequenceOffset < 0 {
		t.Fatal("Content Sequence tag not found in output")
	}
	if sequenceOffset+20 > len(data) {
		t.Fatal("Content Sequence header is truncated")
	}
	if got := binary.LittleEndian.Uint32(data[sequenceOffset+8 : sequenceOffset+12]); got == 0xffffffff {
		t.Fatal("Content Sequence used undefined length")
	}
	itemTag := []byte{0xfe, 0xff, 0x00, 0xe0}
	if !bytes.Equal(data[sequenceOffset+12:sequenceOffset+16], itemTag) {
		t.Fatalf("first Content Sequence item tag = % x, want % x", data[sequenceOffset+12:sequenceOffset+16], itemTag)
	}
	if got := binary.LittleEndian.Uint32(data[sequenceOffset+16 : sequenceOffset+20]); got == 0xffffffff {
		t.Fatal("Content Sequence item used undefined length")
	}
}

func TestReadRoundTripsStringBackedTypedValues(t *testing.T) {
	code := NewCodeItem("121071", "DCM", "Finding")
	personName, err := NewContentItemPersonName(code, RelationshipContains, "Doe^Jane")
	if err != nil {
		t.Fatalf("NewContentItemPersonName() error = %v", err)
	}
	dateTimes := []time.Time{
		time.Date(2026, time.August, 15, 8, 0, 0, 0, time.UTC),
		time.Date(2026, time.August, 15, 9, 0, 0, 0, time.UTC),
	}
	temporal, err := NewContentItemTemporalCoordinate(
		NewCodeItem("111034", "DCM", "Temporal Region"),
		RelationshipSelectedFrom,
		TemporalCoordinate{RangeType: TemporalRangeTypeSegment, DateTimes: dateTimes},
	)
	if err != nil {
		t.Fatalf("NewContentItemTemporalCoordinate() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), personName, temporal)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	var output bytes.Buffer
	if err := report.Write(&output, writer.WithoutPreamble()); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	parsed, err := Read(bytes.NewReader(output.Bytes()), parser.WithAssumedTransferSyntax(transfer.ExplicitVRLittleEndian))
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	children, err := parsed.Children()
	if err != nil {
		t.Fatalf("Children() error = %v", err)
	}
	if got, err := children[0].GetPersonName(); err != nil || got != "Doe^Jane" {
		t.Fatalf("GetPersonName() = %q, %v, want Doe^Jane, nil", got, err)
	}
	gotTemporal, err := children[1].GetTemporalCoordinate()
	if err != nil {
		t.Fatalf("GetTemporalCoordinate() error = %v", err)
	}
	if !equalTimes(gotTemporal.DateTimes, dateTimes) {
		t.Fatalf("DateTimes = %v, want %v", gotTemporal.DateTimes, dateTimes)
	}
}
