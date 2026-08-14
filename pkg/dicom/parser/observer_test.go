// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestParseSequenceItemObserverReportsAbsoluteItemOffset(t *testing.T) {
	encoded, wantOffset := observedSequenceFile(t)

	tests := []struct {
		name   string
		reader io.Reader
	}{
		{name: "seekable", reader: bytes.NewReader(encoded)},
		{name: "non-seekable", reader: &readOnly{reader: bytes.NewReader(encoded)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var positions []SequenceItemPosition
			result, err := Parse(tt.reader, WithSequenceItemObserver(func(position SequenceItemPosition) error {
				positions = append(positions, position)
				return nil
			}))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			if len(positions) != 1 {
				t.Fatalf("observer calls = %d, want 1", len(positions))
			}

			position := positions[0]
			if !position.SequenceTag.Equals(tag.ReferencedSeriesSequence) {
				t.Fatalf("SequenceTag = %s, want %s", position.SequenceTag, tag.ReferencedSeriesSequence)
			}
			if position.Offset != wantOffset {
				t.Fatalf("Offset = %d, want %d", position.Offset, wantOffset)
			}
			if got := binary.LittleEndian.Uint32(encoded[position.Offset : position.Offset+4]); got != 0xE000FFFE {
				t.Fatalf("offset points to %#08x, want Item Tag", got)
			}

			sequence, err := result.Dataset.GetSequence(tag.ReferencedSeriesSequence)
			if err != nil {
				t.Fatalf("GetSequence() error = %v", err)
			}
			if position.Item != sequence.GetItem(0) {
				t.Fatal("observer Item does not match the parsed sequence item")
			}
		})
	}
}

func observedSequenceFile(t *testing.T) ([]byte, uint64) {
	t.Helper()

	var file bytes.Buffer
	file.Write(make([]byte, 128))
	file.WriteString("DICM")
	writeExplicitStringElement(&file, tag.TransferSyntaxUID, "UI", []byte(testExplicitVRLittleLE+"\x00"))

	writeTagForObserverTest(t, &file, tag.ReferencedSeriesSequence)
	file.WriteString("SQ")
	writeUint16ForObserverTest(t, &file, 0)
	writeUint32ForObserverTest(t, &file, 0xFFFFFFFF)

	wantOffset := uint64(file.Len())
	writeTagForObserverTest(t, &file, tag.New(0xFFFE, 0xE000))
	writeUint32ForObserverTest(t, &file, 0xFFFFFFFF)
	writeExplicitStringElement(&file, tag.SeriesInstanceUID, "UI", []byte("1.2.3\x00"))
	writeTagForObserverTest(t, &file, tag.New(0xFFFE, 0xE00D))
	writeUint32ForObserverTest(t, &file, 0)
	writeTagForObserverTest(t, &file, tag.New(0xFFFE, 0xE0DD))
	writeUint32ForObserverTest(t, &file, 0)

	return file.Bytes(), wantOffset
}

func writeTagForObserverTest(t *testing.T, w io.Writer, value *tag.Tag) {
	t.Helper()
	writeUint16ForObserverTest(t, w, value.Group())
	writeUint16ForObserverTest(t, w, value.Element())
}

func writeUint16ForObserverTest(t *testing.T, w io.Writer, value uint16) {
	t.Helper()
	if err := binary.Write(w, binary.LittleEndian, value); err != nil {
		t.Fatalf("binary.Write(uint16) error = %v", err)
	}
}

func writeUint32ForObserverTest(t *testing.T, w io.Writer, value uint32) {
	t.Helper()
	if err := binary.Write(w, binary.LittleEndian, value); err != nil {
		t.Fatalf("binary.Write(uint32) error = %v", err)
	}
}

type readOnly struct {
	reader io.Reader
}

func (r *readOnly) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}
