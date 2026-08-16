// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type byteCountingReadSeeker struct {
	*bytes.Reader
	read int
}

func (r *byteCountingReadSeeker) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.read += n
	return n, err
}

func TestParseStopBeforePixelDataStopsBeforeTopLevelValue(t *testing.T) {
	tests := []struct {
		name string
		tag  *tag.Tag
		vr   string
	}{
		{name: "retired variable pixel data", tag: tag.New(0x7F01, 0x0010), vr: "OB"},
		{name: "float pixel data", tag: tag.FloatPixelData, vr: "OF"},
		{name: "double float pixel data", tag: tag.DoubleFloatPixelData, vr: "OD"},
		{name: "pixel data", tag: tag.PixelData, vr: "OB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := createMiniDICOM()
			stopOffset := file.Len() + 4
			writeExplicitLongValueElement(file, tt.tag, tt.vr, bytes.Repeat([]byte{0x5A}, 1024))

			reader := &byteCountingReadSeeker{Reader: bytes.NewReader(file.Bytes())}
			result, err := Parse(reader, WithStopBeforePixelData())
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			if !result.IsPartial {
				t.Fatal("IsPartial = false, want true")
			}
			if got := result.Dataset.TryGetString(tag.PatientName); got != testPatientName {
				t.Fatalf("PatientName = %q, want %q", got, testPatientName)
			}
			if _, ok := result.Dataset.Get(tt.tag); ok {
				t.Fatalf("Dataset contains stopped tag %s", tt.tag)
			}
			if reader.read != stopOffset {
				t.Fatalf("bytes read = %d, want %d before the pixel VR and value", reader.read, stopOffset)
			}
		})
	}
}

func TestParseStopBeforePixelDataStopsInsideSequenceItem(t *testing.T) {
	file := createMiniDICOM()
	sequenceTag := tag.New(0x7776, 0x0010)
	writeTestTag(file, sequenceTag)
	file.WriteString("SQ")
	_ = binary.Write(file, binary.LittleEndian, uint16(0))
	_ = binary.Write(file, binary.LittleEndian, uint32(0xFFFFFFFF))

	writeTestTag(file, tag.New(0xFFFE, 0xE000))
	_ = binary.Write(file, binary.LittleEndian, uint32(0xFFFFFFFF))
	stopOffset := file.Len() + 4
	writeExplicitLongValueElement(file, tag.PixelData, "OB", bytes.Repeat([]byte{0xA5}, 1024))
	writeTestTag(file, tag.New(0xFFFE, 0xE00D))
	_ = binary.Write(file, binary.LittleEndian, uint32(0))
	writeTestTag(file, tag.New(0xFFFE, 0xE0DD))
	_ = binary.Write(file, binary.LittleEndian, uint32(0))

	reader := &byteCountingReadSeeker{Reader: bytes.NewReader(file.Bytes())}
	result, err := Parse(reader, WithStopBeforePixelData())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if !result.IsPartial {
		t.Fatal("IsPartial = false, want true")
	}
	if got := result.Dataset.TryGetString(tag.PatientName); got != testPatientName {
		t.Fatalf("PatientName = %q, want %q", got, testPatientName)
	}
	if _, ok := result.Dataset.Get(sequenceTag); ok {
		t.Fatalf("Dataset contains incomplete sequence %s", sequenceTag)
	}
	if reader.read != stopOffset {
		t.Fatalf("bytes read = %d, want %d before nested pixel VR and value", reader.read, stopOffset)
	}
}

func writeTestTag(buf *bytes.Buffer, value *tag.Tag) {
	_ = binary.Write(buf, binary.LittleEndian, value.Group())
	_ = binary.Write(buf, binary.LittleEndian, value.Element())
}
