// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestReadPreamble tests preamble reading
func TestReadPreamble(t *testing.T) {
	t.Run("ValidPreamble", func(t *testing.T) {
		// Create valid preamble + DICM
		buf := bytes.NewBuffer(nil)
		buf.Write(make([]byte, 128)) // 128 zero bytes
		buf.WriteString("DICM")      // DICM prefix

		ctx := newParseContext()
		ctx.reader = buf
		if err := ctx.readPreamble(); err != nil {
			t.Errorf("readPreamble() error = %v", err)
		}
	})

	t.Run("InvalidDICMPrefix", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		buf.Write(make([]byte, 128))
		buf.WriteString("XXXX") // Invalid prefix

		ctx := newParseContext()
		ctx.reader = buf
		if err := ctx.readPreamble(); err == nil {
			t.Error("readPreamble() should return error for invalid DICM prefix")
		}
	})

	t.Run("ShortPreamble", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		buf.Write(make([]byte, 50)) // Too short

		ctx := newParseContext()
		ctx.reader = buf
		if err := ctx.readPreamble(); err == nil {
			t.Error("readPreamble() should return error for short preamble")
		}
	})
}

// TestReadTag tests tag reading
func TestReadTag(t *testing.T) {
	t.Run("ReadValidTag", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
		binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian

		tag, err := ctx.readTag()
		if err != nil {
			t.Fatalf("readTag() error = %v", err)
		}

		if tag.Group() != 0x0002 {
			t.Errorf("tag.Group() = %04X, want 0002", tag.Group())
		}
		if tag.Element() != 0x0000 {
			t.Errorf("tag.Element() = %04X, want 0000", tag.Element())
		}
	})

	t.Run("ReadTagBigEndian", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.BigEndian, uint16(0x0002))
		binary.Write(buf, binary.BigEndian, uint16(0x0010))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.BigEndian

		tag, err := ctx.readTag()
		if err != nil {
			t.Fatalf("readTag() error = %v", err)
		}

		if tag.Group() != 0x0002 || tag.Element() != 0x0010 {
			t.Errorf("tag = %s, want (0002,0010)", tag)
		}
	})

	t.Run("ReadTagEOF", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{0x00}) // Too short

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian

		_, err := ctx.readTag()
		if err == nil {
			t.Error("readTag() should return error for EOF")
		}
	})
}

// TestReadVR tests VR reading
func TestReadVR(t *testing.T) {
	t.Run("ExplicitVR", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte("PN"))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.isExplicitVR = true

		vr, err := ctx.readVR(tag.PatientName)
		if err != nil {
			t.Fatalf("readVR() error = %v", err)
		}

		if vr.String() != "PN" {
			t.Errorf("vr = %s, want PN", vr)
		}
	})

	t.Run("ImplicitVR", func(t *testing.T) {
		ctx := newParseContext()
		ctx.isExplicitVR = false

		vrResult, err := ctx.readVR(tag.PatientName)
		if err != nil {
			t.Fatalf("readVR() error = %v", err)
		}

		// Should return UN for now (dictionary lookup not implemented)
		if vrResult.Code() != "UN" {
			t.Errorf("vr = %s, want UN", vrResult)
		}
	})
}

// TestReadLength tests length reading
func TestReadLength(t *testing.T) {
	t.Run("ExplicitVR16BitLength", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.LittleEndian, uint16(100))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian
		ctx.isExplicitVR = true

		length, err := ctx.readLength(vr.PN)
		if err != nil {
			t.Fatalf("readLength() error = %v", err)
		}

		if length != 100 {
			t.Errorf("length = %d, want 100", length)
		}
	})

	t.Run("ExplicitVR32BitLength", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.LittleEndian, uint16(0)) // Reserved
		binary.Write(buf, binary.LittleEndian, uint32(1000))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian
		ctx.isExplicitVR = true

		length, err := ctx.readLength(vr.SQ)
		if err != nil {
			t.Fatalf("readLength() error = %v", err)
		}

		if length != 1000 {
			t.Errorf("length = %d, want 1000", length)
		}
	})

	t.Run("ImplicitVR", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.LittleEndian, uint32(2000))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian
		ctx.isExplicitVR = false

		length, err := ctx.readLength(vr.PN)
		if err != nil {
			t.Fatalf("readLength() error = %v", err)
		}

		if length != 2000 {
			t.Errorf("length = %d, want 2000", length)
		}
	})

	t.Run("UndefinedLength", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		binary.Write(buf, binary.LittleEndian, uint16(0)) // Reserved
		binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF))

		ctx := newParseContext()
		ctx.reader = buf
		ctx.byteOrder = binary.LittleEndian
		ctx.isExplicitVR = true

		length, err := ctx.readLength(vr.SQ)
		if err != nil {
			t.Fatalf("readLength() error = %v", err)
		}

		if length != 0xFFFFFFFF {
			t.Errorf("length = %08X, want FFFFFFFF", length)
		}
	})
}

// TestParseWithOptions tests parser options
func TestParseWithOptions(t *testing.T) {
	t.Run("WithMaxElementSize", func(t *testing.T) {
		ctx := newParseContext(WithMaxElementSize(1000))
		if ctx.maxElementSize != 1000 {
			t.Errorf("maxElementSize = %d, want 1000", ctx.maxElementSize)
		}
	})

	t.Run("WithStopAtTag", func(t *testing.T) {
		stopTag := tag.PixelData
		ctx := newParseContext(WithStopAtTag(stopTag))
		if ctx.stopAtTag != stopTag {
			t.Errorf("stopAtTag = %v, want %v", ctx.stopAtTag, stopTag)
		}
	})
}

// createMiniDICOM creates a minimal valid DICOM file for testing
func createMiniDICOM() *bytes.Buffer {
	buf := bytes.NewBuffer(nil)

	// Preamble + DICM
	buf.Write(make([]byte, 128))
	buf.WriteString("DICM")

	// File Meta Information Group Length (0002,0000) - UL, Explicit VR
	binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
	binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element
	buf.WriteString("UL")                                   // VR
	binary.Write(buf, binary.LittleEndian, uint16(4))      // Length
	binary.Write(buf, binary.LittleEndian, uint32(100))    // Value (placeholder)

	// Transfer Syntax UID (0002,0010) - UI, Explicit VR
	binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
	binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Element
	buf.WriteString("UI")                                   // VR
	tsUID := "1.2.840.10008.1.2.1"                          // Explicit VR Little Endian
	binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)

	// Patient Name (0010,0010) - PN, Explicit VR
	binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Group
	binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Element
	buf.WriteString("PN")                                   // VR
	name := "Doe^John"
	binary.Write(buf, binary.LittleEndian, uint16(len(name)))
	buf.WriteString(name)

	return buf
}

// TestParseMiniDICOM tests parsing a minimal DICOM file
func TestParseMiniDICOM(t *testing.T) {
	buf := createMiniDICOM()

	result, err := Parse(buf)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if result == nil {
		t.Fatal("Parse() returned nil result")
	}

	// Check FileMetaInformation
	if result.FileMetaInformation == nil {
		t.Fatal("FileMetaInformation should not be nil")
	}

	// Check for Transfer Syntax UID in meta info
	tsUID, exists := result.FileMetaInformation.TransferSyntaxUID()
	if !exists {
		t.Error("FileMetaInformation should contain TransferSyntaxUID")
	}
	if tsUID != "1.2.840.10008.1.2.1" {
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, "1.2.840.10008.1.2.1")
	}

	// Check Dataset
	if result.Dataset == nil {
		t.Fatal("Dataset should not be nil")
	}

	// Check for Patient Name in main dataset
	name, exists := result.Dataset.GetString(tag.PatientName)
	if !exists {
		t.Error("Dataset should contain PatientName")
	}
	if name != "Doe^John" {
		t.Errorf("PatientName = %q, want %q", name, "Doe^John")
	}
}

// TestParseInvalidFile tests error handling
func TestParseInvalidFile(t *testing.T) {
	t.Run("EmptyFile", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		_, err := Parse(buf)
		if err == nil {
			t.Error("Parse() should return error for empty file")
		}
	})

	t.Run("NoDICMPrefix", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		buf.Write(make([]byte, 128))
		buf.WriteString("XXXX")

		_, err := Parse(buf)
		if err == nil {
			t.Error("Parse() should return error for invalid DICM prefix")
		}
	})

	t.Run("TruncatedFile", func(t *testing.T) {
		// Create file with preamble and DICM but incomplete element
		buf := bytes.NewBuffer(nil)
		buf.Write(make([]byte, 128))
		buf.WriteString("DICM")

		// Start of Group 0002 element, but incomplete (no VR or value)
		binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
		binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element
		// Missing VR and value - EOF will occur when trying to read VR

		_, err := Parse(buf)
		if err == nil {
			t.Error("Parse() should return error for truncated file with incomplete element")
		}
	})
}

// TestSequenceReading tests reading sequences with defined and undefined lengths
func TestSequenceReading(t *testing.T) {
	t.Run("UndefinedLengthSequence", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		// Preamble + DICM
		buf.Write(make([]byte, 128))
		buf.WriteString("DICM")

		// Transfer Syntax UID (0002,0010) - required meta element
		binary.Write(buf, binary.LittleEndian, uint16(0x0002))
		binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("UI")
		tsUID := "1.2.840.10008.1.2.1"
		binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
		buf.WriteString(tsUID)

		// Sequence element with undefined length (SQ)
		binary.Write(buf, binary.LittleEndian, uint16(0x0008)) // Group
		binary.Write(buf, binary.LittleEndian, uint16(0x1140)) // Element (Referenced Image Sequence)
		buf.WriteString("SQ")                                   // VR
		binary.Write(buf, binary.LittleEndian, uint16(0))      // Reserved
		binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF)) // Undefined length

		// Item 1 (FFFE,E000) with defined length
		binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
		binary.Write(buf, binary.LittleEndian, uint16(0xE000))
		itemData := bytes.NewBuffer(nil)
		// Add SOP Instance UID in item
		binary.Write(itemData, binary.LittleEndian, uint16(0x0008))
		binary.Write(itemData, binary.LittleEndian, uint16(0x1155))
		itemData.WriteString("UI")
		sopUID := "1.2.3.4"
		binary.Write(itemData, binary.LittleEndian, uint16(len(sopUID)))
		itemData.WriteString(sopUID)
		binary.Write(buf, binary.LittleEndian, uint32(itemData.Len()))
		buf.Write(itemData.Bytes())

		// Sequence Delimitation Item (FFFE,E0DD)
		binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
		binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
		binary.Write(buf, binary.LittleEndian, uint32(0))

		result, err := Parse(buf)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		// Check sequence exists in main dataset
		elem, exists := result.Dataset.Get(tag.New(0x0008, 0x1140))
		if !exists {
			t.Fatal("Sequence element should exist")
		}

		seq, ok := elem.(*dataset.Sequence)
		if !ok {
			t.Fatal("Element should be a Sequence")
		}

		if seq.Count() != 1 {
			t.Errorf("Sequence should have 1 item, got %d", seq.Count())
		}
	})

	t.Run("DefinedLengthSequence", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		// Preamble + DICM
		buf.Write(make([]byte, 128))
		buf.WriteString("DICM")

		// Transfer Syntax UID (0002,0010)
		binary.Write(buf, binary.LittleEndian, uint16(0x0002))
		binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("UI")
		tsUID := "1.2.840.10008.1.2.1"
		binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
		buf.WriteString(tsUID)

		// Build sequence data first to know its length
		seqData := bytes.NewBuffer(nil)

		// Item 1 (FFFE,E000) with defined length
		binary.Write(seqData, binary.LittleEndian, uint16(0xFFFE))
		binary.Write(seqData, binary.LittleEndian, uint16(0xE000))
		itemData := bytes.NewBuffer(nil)
		// Add SOP Instance UID in item
		binary.Write(itemData, binary.LittleEndian, uint16(0x0008))
		binary.Write(itemData, binary.LittleEndian, uint16(0x1155))
		itemData.WriteString("UI")
		sopUID := "1.2.3.4"
		binary.Write(itemData, binary.LittleEndian, uint16(len(sopUID)))
		itemData.WriteString(sopUID)
		binary.Write(seqData, binary.LittleEndian, uint32(itemData.Len()))
		seqData.Write(itemData.Bytes())

		// Now write sequence element with defined length
		binary.Write(buf, binary.LittleEndian, uint16(0x0008))
		binary.Write(buf, binary.LittleEndian, uint16(0x1140))
		buf.WriteString("SQ")
		binary.Write(buf, binary.LittleEndian, uint16(0))
		binary.Write(buf, binary.LittleEndian, uint32(seqData.Len()))
		buf.Write(seqData.Bytes())

		result, err := Parse(buf)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		// Check sequence exists in main dataset
		elem, exists := result.Dataset.Get(tag.New(0x0008, 0x1140))
		if !exists {
			t.Fatal("Sequence element should exist")
		}

		seq, ok := elem.(*dataset.Sequence)
		if !ok {
			t.Fatal("Element should be a Sequence")
		}

		if seq.Count() != 1 {
			t.Errorf("Sequence should have 1 item, got %d", seq.Count())
		}

		// Verify item content
		item := seq.GetItem(0)
		if item == nil {
			t.Fatal("Item should not be nil")
		}

		sopUIDElem, exists := item.Get(tag.New(0x0008, 0x1155))
		if !exists {
			t.Fatal("SOP Instance UID should exist in item")
		}

		if sopUIDElem.Tag().ToUint32() != 0x00081155 {
			t.Errorf("Wrong tag: %s", sopUIDElem.Tag())
		}
	})
}
