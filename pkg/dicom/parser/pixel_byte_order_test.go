// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

// encodeGEPrivateBigEndianPixelDataset builds an implicit VR little-endian
// dataset whose native Pixel Data words are big-endian, as GE's private
// transfer syntax encodes them. The single sample value is 0x1234.
func encodeGEPrivateBigEndianPixelDataset() []byte {
	var raw bytes.Buffer
	writeImplicitElement(&raw, tag.Rows.Group(), tag.Rows.Element(), []byte{1, 0})
	writeImplicitElement(&raw, tag.Columns.Group(), tag.Columns.Element(), []byte{1, 0})
	writeImplicitElement(&raw, tag.BitsAllocated.Group(), tag.BitsAllocated.Element(), []byte{16, 0})
	writeImplicitElement(&raw, tag.BitsStored.Group(), tag.BitsStored.Element(), []byte{16, 0})
	writeImplicitElement(&raw, tag.HighBit.Group(), tag.HighBit.Element(), []byte{15, 0})
	writeImplicitElement(&raw, tag.SamplesPerPixel.Group(), tag.SamplesPerPixel.Element(), []byte{1, 0})
	writeImplicitElement(&raw, tag.PixelRepresentation.Group(), tag.PixelRepresentation.Element(), []byte{0, 0})
	writeImplicitElement(&raw, tag.PhotometricInterpretation.Group(),
		tag.PhotometricInterpretation.Element(), []byte(testPhotometricMonoChrome+" "))
	writeImplicitElement(&raw, tag.PixelData.Group(), tag.PixelData.Element(), []byte{0x12, 0x34})
	return raw.Bytes()
}

func parseGEPrivateBigEndianPixelDataset(t *testing.T) *ParseResult {
	t.Helper()
	result, err := Parse(
		bytes.NewReader(encodeGEPrivateBigEndianPixelDataset()),
		WithAssumedTransferSyntax(transfer.GEPrivateImplicitVRBigEndian),
	)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	return result
}

func assertParsedPixelDataBytes(t *testing.T, result *ParseResult, want []byte) {
	t.Helper()
	elem, ok := result.Dataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("dataset has no PixelData")
	}
	if got := elem.Buffer().Data(); !bytes.Equal(got, want) {
		t.Fatalf("PixelData bytes = %x, want %x", got, want)
	}
}

func assertParsedSample(t *testing.T, result *ParseResult) {
	t.Helper()
	pd, err := pixeldata.FromDataset(result.Dataset)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	sample, err := pd.Sample(0, 0, 0, 0)
	if err != nil {
		t.Fatalf("Sample() error = %v", err)
	}
	if sample != 0x1234 {
		t.Fatalf("sample = %#x, want 0x1234", sample)
	}
}

func TestParseGEPrivateBigEndianPixelDataRecordsBigEndianOrder(t *testing.T) {
	result := parseGEPrivateBigEndianPixelDataset(t)

	elem, ok := result.Dataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("dataset has no PixelData")
	}
	if _, ok := elem.(*element.OtherWord); !ok {
		t.Fatalf("PixelData = %T, want *element.OtherWord", elem)
	}
	// The buffered bytes remain the encoded big-endian words.
	if got := elem.Buffer().Data(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("PixelData bytes = %x, want big-endian 1234", got)
	}
	order, known := element.NumericByteOrder(elem)
	if !known || order.Uint16([]byte{0x01, 0x02}) != binary.BigEndian.Uint16([]byte{0x01, 0x02}) {
		t.Fatalf("PixelData numeric byte order = %T (known %v), want big-endian", order, known)
	}
	assertParsedSample(t, result)

	pd, err := pixeldata.FromDataset(result.Dataset)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	written, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if got := written.Buffer().Data(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("ToElement() bytes = %x, want big-endian 1234", got)
	}
}

func TestWriteGEPrivateBigEndianPixelDataAsLittleEndian(t *testing.T) {
	result := parseGEPrivateBigEndianPixelDataset(t)
	if err := result.Dataset.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatal(err)
	}
	if err := result.Dataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testCycleSOPInstanceUID})); err != nil {
		t.Fatal(err)
	}

	var encoded bytes.Buffer
	if err := writer.Write(&encoded, result.Dataset, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	written, err := Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	assertParsedPixelDataBytes(t, written, []byte{0x34, 0x12})
	assertParsedSample(t, result)
}

func TestWriteGEPrivateBigEndianPixelDataAsGESyntax(t *testing.T) {
	result := parseGEPrivateBigEndianPixelDataset(t)
	if err := result.Dataset.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatal(err)
	}
	if err := result.Dataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testCycleSOPInstanceUID})); err != nil {
		t.Fatal(err)
	}

	var encoded bytes.Buffer
	if err := writer.Write(&encoded, result.Dataset, writer.WithTransferSyntax(transfer.GEPrivateImplicitVRBigEndian)); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	written, err := Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	assertParsedPixelDataBytes(t, written, []byte{0x12, 0x34})
	assertParsedSample(t, result)
}

func TestExplicitVRBigEndianPixelDataRoundTrip(t *testing.T) {
	var raw bytes.Buffer
	writeExplicitBigEndianElement(&raw, tag.Rows, "US", []byte{0, 1})
	writeExplicitBigEndianElement(&raw, tag.Columns, "US", []byte{0, 1})
	writeExplicitBigEndianElement(&raw, tag.BitsAllocated, "US", []byte{0, 16})
	writeExplicitBigEndianElement(&raw, tag.BitsStored, "US", []byte{0, 16})
	writeExplicitBigEndianElement(&raw, tag.HighBit, "US", []byte{0, 15})
	writeExplicitBigEndianElement(&raw, tag.SamplesPerPixel, "US", []byte{0, 1})
	writeExplicitBigEndianElement(&raw, tag.PixelRepresentation, "US", []byte{0, 0})
	writeExplicitBigEndianElement(&raw, tag.PhotometricInterpretation, "CS", []byte(testPhotometricMonoChrome+" "))
	writeExplicitBigEndianElement(&raw, tag.PixelData, "OW", []byte{0x12, 0x34})

	result, err := Parse(bytes.NewReader(raw.Bytes()), WithAssumedTransferSyntax(transfer.ExplicitVRBigEndian))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	assertParsedPixelDataBytes(t, result, []byte{0x12, 0x34})
	assertParsedSample(t, result)
	if err := result.Dataset.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatal(err)
	}
	if err := result.Dataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testCycleSOPInstanceUID})); err != nil {
		t.Fatal(err)
	}

	var encoded bytes.Buffer
	if err := writer.Write(&encoded, result.Dataset); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	written, err := Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	assertParsedPixelDataBytes(t, written, []byte{0x12, 0x34})
	assertParsedSample(t, result)
}

func writeExplicitBigEndianElement(buf *bytes.Buffer, t *tag.Tag, vrCode string, value []byte) {
	_ = binary.Write(buf, binary.BigEndian, t.Group())
	_ = binary.Write(buf, binary.BigEndian, t.Element())
	buf.WriteString(vrCode)
	switch vrCode {
	case "OB", "OW", "OF", "OD", "OL", "OV", "SQ", "UC", "UR", "UT", "UN":
		_ = binary.Write(buf, binary.BigEndian, uint16(0))
		_ = binary.Write(buf, binary.BigEndian, uint32(len(value)))
	default:
		_ = binary.Write(buf, binary.BigEndian, uint16(len(value))) //nolint:gosec // test values are tiny
	}
	_, _ = buf.Write(value)
}
