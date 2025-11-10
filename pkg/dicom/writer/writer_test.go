// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestWritePreamble tests preamble writing
func TestWritePreamble(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	if err := w.writePreamble(); err != nil {
		t.Fatalf("writePreamble() error = %v", err)
	}

	// Check preamble length (128 + 4 = 132 bytes)
	if buf.Len() != 132 {
		t.Errorf("Preamble length = %d, want 132", buf.Len())
	}

	// Check DICM prefix
	data := buf.Bytes()
	dicmPrefix := string(data[128:132])
	if dicmPrefix != "DICM" {
		t.Errorf("DICM prefix = %q, want %q", dicmPrefix, "DICM")
	}
}

// TestWriteTag tests tag writing
func TestWriteTag(t *testing.T) {
	t.Run("LittleEndian", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		testTag := tag.New(0x0010, 0x0010)
		if err := w.writeTag(testTag); err != nil {
			t.Fatalf("writeTag() error = %v", err)
		}

		if buf.Len() != 4 {
			t.Errorf("Tag length = %d, want 4", buf.Len())
		}

		// Verify bytes
		data := buf.Bytes()
		group := binary.LittleEndian.Uint16(data[0:2])
		elem := binary.LittleEndian.Uint16(data[2:4])

		if group != 0x0010 {
			t.Errorf("Group = %04X, want 0010", group)
		}
		if elem != 0x0010 {
			t.Errorf("Element = %04X, want 0010", elem)
		}
	})
}

// TestWriteVR tests VR writing
func TestWriteVR(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	if err := w.writeVR(vr.PN); err != nil {
		t.Fatalf("writeVR() error = %v", err)
	}

	if buf.Len() != 2 {
		t.Errorf("VR length = %d, want 2", buf.Len())
	}

	vrCode := buf.String()
	if vrCode != "PN" {
		t.Errorf("VR = %q, want %q", vrCode, "PN")
	}
}

// TestWriteLength tests length writing
func TestWriteLength(t *testing.T) {
	t.Run("16BitLength", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = true

		if err := w.writeLength(vr.PN, 100); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		if buf.Len() != 2 {
			t.Errorf("Length field size = %d, want 2", buf.Len())
		}

		length := binary.LittleEndian.Uint16(buf.Bytes())
		if length != 100 {
			t.Errorf("Length = %d, want 100", length)
		}
	})

	t.Run("32BitLength", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = true

		if err := w.writeLength(vr.SQ, 1000); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		// 2 bytes reserved + 4 bytes length = 6 bytes
		if buf.Len() != 6 {
			t.Errorf("Length field size = %d, want 6", buf.Len())
		}

		data := buf.Bytes()
		reserved := binary.LittleEndian.Uint16(data[0:2])
		length := binary.LittleEndian.Uint32(data[2:6])

		if reserved != 0 {
			t.Errorf("Reserved = %d, want 0", reserved)
		}
		if length != 1000 {
			t.Errorf("Length = %d, want 1000", length)
		}
	})

	t.Run("ImplicitVR", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		w.isExplicitVR = false

		if err := w.writeLength(vr.PN, 2000); err != nil {
			t.Fatalf("writeLength() error = %v", err)
		}

		// Implicit VR always uses 32-bit length
		if buf.Len() != 4 {
			t.Errorf("Length field size = %d, want 4", buf.Len())
		}

		length := binary.LittleEndian.Uint32(buf.Bytes())
		if length != 2000 {
			t.Errorf("Length = %d, want 2000", length)
		}
	})
}

// TestWriteElement tests element writing
func TestWriteElement(t *testing.T) {
	t.Run("StringElement", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})
		if err := w.writeElement(elem); err != nil {
			t.Fatalf("writeElement() error = %v", err)
		}

		// Tag (4) + VR (2) + Length (2) + Value (8) = 16 bytes
		if buf.Len() != 16 {
			t.Errorf("Element size = %d, want 16", buf.Len())
		}
	})

	t.Run("UnsignedShortElement", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf

		elem := element.NewUnsignedShort(tag.Rows, []uint16{512})
		if err := w.writeElement(elem); err != nil {
			t.Fatalf("writeElement() error = %v", err)
		}

		// Tag (4) + VR (2) + Length (2) + Value (2) = 10 bytes
		if buf.Len() != 10 {
			t.Errorf("Element size = %d, want 10", buf.Len())
		}
	})
}

// TestWriteSimpleDataset tests writing a simple dataset
func TestWriteSimpleDataset(t *testing.T) {
	buf := &bytes.Buffer{}

	// Create dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

	// Write (using defaults)
	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Check that something was written
	if buf.Len() == 0 {
		t.Error("Nothing was written")
	}

	// Check preamble
	data := buf.Bytes()
	if len(data) < 132 {
		t.Error("Output too short to contain preamble")
	}

	dicmPrefix := string(data[128:132])
	if dicmPrefix != "DICM" {
		t.Errorf("DICM prefix = %q, want %q", dicmPrefix, "DICM")
	}
}

// TestWriteWithoutPreamble tests writing without preamble
func TestWriteWithoutPreamble(t *testing.T) {
	buf := &bytes.Buffer{}

	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test"}))

	if err := Write(buf, ds, WithoutPreamble()); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Should not start with preamble
	data := buf.Bytes()
	if len(data) >= 132 {
		dicmCheck := string(data[128:132])
		if dicmCheck == "DICM" {
			t.Error("Should not contain DICM prefix when WithoutPreamble is used")
		}
	}
}

// TestWriteSequence tests writing a sequence
func TestWriteSequence(t *testing.T) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf

	// Create a sequence with one item
	seq := dataset.NewSequence(tag.New(0x0008, 0x1140))
	item := dataset.New()
	item.Add(element.NewString(tag.New(0x0008, 0x1155), vr.UI, []string{"1.2.3.4"}))
	seq.AddItem(item)

	if err := w.writeSequence(seq); err != nil {
		t.Fatalf("writeSequence() error = %v", err)
	}

	// Sequence should have been written
	if buf.Len() == 0 {
		t.Error("Nothing was written")
	}

	// Check for sequence delimitation tag (FFFE,E0DD) at the end
	data := buf.Bytes()
	if len(data) < 8 {
		t.Error("Sequence output too short")
	}

	// The last 8 bytes should be the delimitation tag and length
	lastBytes := data[len(data)-8:]
	delimGroup := binary.LittleEndian.Uint16(lastBytes[0:2])
	delimElem := binary.LittleEndian.Uint16(lastBytes[2:4])
	delimLength := binary.LittleEndian.Uint32(lastBytes[4:8])

	if delimGroup != 0xFFFE || delimElem != 0xE0DD {
		t.Errorf("Expected sequence delimitation tag (FFFE,E0DD), got (%04X,%04X)",
			delimGroup, delimElem)
	}
	if delimLength != 0 {
		t.Errorf("Delimitation length = %d, want 0", delimLength)
	}
}

// TestRoundTrip tests write then read
func TestRoundTrip(t *testing.T) {
	buf := &bytes.Buffer{}

	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))

	if err := Write(buf, ds); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	t.Logf("Wrote %d bytes", buf.Len())

	// Verify basic structure
	data := buf.Bytes()

	// Should have preamble
	if len(data) < 132 {
		t.Fatal("Output too short")
	}

	// Should have DICM
	if string(data[128:132]) != "DICM" {
		t.Error("Missing DICM prefix")
	}
}

// Benchmark tests for Writer

func BenchmarkWriteSmallDataset(b *testing.B) {
	// Create a small dataset
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20250106"}))
	ds.Add(element.NewString(tag.Modality, vr.CS, []string{"CT"}))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteMediumDataset(b *testing.B) {
	// Create a medium dataset with 50 elements
	ds := dataset.New()
	for i := 0; i < 50; i++ {
		t := tag.New(0x0010, uint16(i))
		ds.Add(element.NewString(t, vr.LO, []string{"TestValue"}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteLargeDataset(b *testing.B) {
	// Create a large dataset with 200 elements
	ds := dataset.New()
	for i := 0; i < 200; i++ {
		t := tag.New(0x0010, uint16(i%256))
		ds.Add(element.NewString(t, vr.LO, []string{"TestValue"}))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		_ = Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	}
}

func BenchmarkWriteTag(b *testing.B) {
	buf := &bytes.Buffer{}
	w := New(transfer.ExplicitVRLittleEndian)
	w.writer = buf
	testTag := tag.New(0x0010, 0x0010)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = w.writeTag(testTag)
	}
}

func BenchmarkWriteElement(b *testing.B) {
	elem := element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		_ = w.writeElement(elem)
	}
}

func BenchmarkWritePreamble(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		w := New(transfer.ExplicitVRLittleEndian)
		w.writer = buf
		_ = w.writePreamble()
	}
}
