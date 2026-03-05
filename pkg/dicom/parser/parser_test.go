// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

const (
	testExplicitVRLittleLE = "1.2.840.10008.1.2.1"
	testPatientName        = "Doe^John"
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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element

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
		_ = binary.Write(buf, binary.BigEndian, uint16(0x0002))
		_ = binary.Write(buf, binary.BigEndian, uint16(0x0010))

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

		// Should automatically use default dictionary and return PN
		if vrResult.Code() != "PN" {
			t.Errorf("vr = %s, want PN", vrResult)
		}
	})
}

// TestReadLength tests length reading
func TestReadLength(t *testing.T) {
	t.Run("ExplicitVR16BitLength", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		_ = binary.Write(buf, binary.LittleEndian, uint16(100))

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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0)) // Reserved
		_ = binary.Write(buf, binary.LittleEndian, uint32(1000))

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
		_ = binary.Write(buf, binary.LittleEndian, uint32(2000))

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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0)) // Reserved
		_ = binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF))

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
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element
	buf.WriteString("UL")                                      // VR
	_ = binary.Write(buf, binary.LittleEndian, uint16(4))      // Length
	_ = binary.Write(buf, binary.LittleEndian, uint32(100))    // Value (placeholder)

	// Transfer Syntax UID (0002,0010) - UI, Explicit VR
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Element
	buf.WriteString("UI")                                      // VR
	tsUID := testExplicitVRLittleLE                            // Explicit VR Little Endian
	_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)

	// Patient Name (0010,0010) - PN, Explicit VR
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Group
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010)) // Element
	buf.WriteString("PN")                                      // VR
	name := testPatientName
	_ = binary.Write(buf, binary.LittleEndian, uint16(len(name)))
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
	if tsUID != testExplicitVRLittleLE {
		t.Errorf("TransferSyntaxUID = %q, want %q", tsUID, testExplicitVRLittleLE)
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
	if name != testPatientName {
		t.Errorf("PatientName = %q, want %q", name, "Doe^John")
	}
}

func TestParseNoPreambleDICOM(t *testing.T) {
	full := createMiniDICOM().Bytes()
	if len(full) <= 132 {
		t.Fatal("test fixture too short")
	}

	noPreamble := bytes.NewReader(full[132:])
	result, err := Parse(noPreamble)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if result.Format != FormatDICOM3NoPreamble {
		t.Fatalf("Format = %s, want %s", result.Format, FormatDICOM3NoPreamble)
	}

	name, exists := result.Dataset.GetString(tag.PatientName)
	if !exists {
		t.Fatal("Dataset should contain PatientName")
	}
	if name != testPatientName {
		t.Fatalf("PatientName = %q, want %q", name, testPatientName)
	}
}

type countingReader struct {
	r io.Reader
	n int64
}

func (c *countingReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	c.n += int64(n)
	return n, err
}

func TestParseStopAtTagStopsBeforePayload(t *testing.T) {
	buf := createMiniDICOM()

	// Add a large Pixel Data element after Patient Name.
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x7FE0))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("OB")
	_ = binary.Write(buf, binary.LittleEndian, uint16(0)) // Reserved

	pixelData := make([]byte, 1024*1024)
	_ = binary.Write(buf, binary.LittleEndian, uint32(len(pixelData)))
	buf.Write(pixelData)

	cr := &countingReader{r: bytes.NewReader(buf.Bytes())}
	result, err := Parse(cr, WithStopAtTag(tag.PixelData))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if !result.IsPartial {
		t.Fatal("IsPartial should be true when stopAtTag stops parsing early")
	}

	if _, exists := result.Dataset.Get(tag.PixelData); exists {
		t.Fatal("PixelData should not be present when stopping at PixelData")
	}

	if cr.n >= int64(len(buf.Bytes())) {
		t.Fatalf("expected early termination to avoid full read, read=%d total=%d", cr.n, len(buf.Bytes()))
	}
}

func TestCreateElementSupportsAttributeAndBinaryVRs(t *testing.T) {
	ctx := newParseContext()
	testTag := tag.New(0x0008, 0x0001)

	tests := []struct {
		name string
		vr   *vr.VR
		data []byte
		want any
	}{
		{"AT", vr.AT, []byte{0x10, 0x00, 0x10, 0x00}, (*element.AttributeTag)(nil)},
		{"OD", vr.OD, make([]byte, 8), (*element.OtherDouble)(nil)},
		{"OF", vr.OF, make([]byte, 4), (*element.OtherFloat)(nil)},
		{"OL", vr.OL, make([]byte, 4), (*element.OtherLong)(nil)},
		{"OV", vr.OV, make([]byte, 8), (*element.OtherVeryLong)(nil)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			el, err := ctx.createElement(testTag, tc.vr, buffer.NewMemory(tc.data))
			if err != nil {
				t.Fatalf("createElement() error = %v", err)
			}

			switch tc.want.(type) {
			case *element.AttributeTag:
				if _, ok := el.(*element.AttributeTag); !ok {
					t.Fatalf("got %T, want *element.AttributeTag", el)
				}
			case *element.OtherDouble:
				if _, ok := el.(*element.OtherDouble); !ok {
					t.Fatalf("got %T, want *element.OtherDouble", el)
				}
			case *element.OtherFloat:
				if _, ok := el.(*element.OtherFloat); !ok {
					t.Fatalf("got %T, want *element.OtherFloat", el)
				}
			case *element.OtherLong:
				if _, ok := el.(*element.OtherLong); !ok {
					t.Fatalf("got %T, want *element.OtherLong", el)
				}
			case *element.OtherVeryLong:
				if _, ok := el.(*element.OtherVeryLong); !ok {
					t.Fatalf("got %T, want *element.OtherVeryLong", el)
				}
			}
		})
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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element
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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("UI")
		tsUID := testExplicitVRLittleLE
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
		buf.WriteString(tsUID)

		// Sequence element with undefined length (SQ)
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0008))     // Group
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x1140))     // Element (Referenced Image Sequence)
		buf.WriteString("SQ")                                          // VR
		_ = binary.Write(buf, binary.LittleEndian, uint16(0))          // Reserved
		_ = binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF)) // Undefined length

		// Item 1 (FFFE,E000) with defined length
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xE000))
		itemData := bytes.NewBuffer(nil)
		// Add SOP Instance UID in item
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x0008))
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x1155))
		itemData.WriteString("UI")
		sopUID := "1.2.3.4"
		_ = binary.Write(itemData, binary.LittleEndian, uint16(len(sopUID)))
		itemData.WriteString(sopUID)
		_ = binary.Write(buf, binary.LittleEndian, uint32(itemData.Len()))
		buf.Write(itemData.Bytes())

		// Sequence Delimitation Item (FFFE,E0DD)
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
		_ = binary.Write(buf, binary.LittleEndian, uint32(0))

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
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("UI")
		tsUID := testExplicitVRLittleLE
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
		buf.WriteString(tsUID)

		// Build sequence data first to know its length
		seqData := bytes.NewBuffer(nil)

		// Item 1 (FFFE,E000) with defined length
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xE000))
		itemData := bytes.NewBuffer(nil)
		// Add SOP Instance UID in item
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x0008))
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x1155))
		itemData.WriteString("UI")
		sopUID := "1.2.3.4"
		_ = binary.Write(itemData, binary.LittleEndian, uint16(len(sopUID)))
		itemData.WriteString(sopUID)
		_ = binary.Write(seqData, binary.LittleEndian, uint32(itemData.Len()))
		seqData.Write(itemData.Bytes())

		// Now write sequence element with defined length
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0008))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x1140))
		buf.WriteString("SQ")
		_ = binary.Write(buf, binary.LittleEndian, uint16(0))
		_ = binary.Write(buf, binary.LittleEndian, uint32(seqData.Len()))
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

	t.Run("DefinedLengthSequenceWithLargeItemValue_StreamParsing", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		// Preamble + DICM
		buf.Write(make([]byte, 128))
		buf.WriteString("DICM")

		// Transfer Syntax UID (0002,0010)
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("UI")
		tsUID := testExplicitVRLittleLE
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
		buf.WriteString(tsUID)

		// Build sequence payload (defined length).
		seqData := bytes.NewBuffer(nil)
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xE000))

		itemData := bytes.NewBuffer(nil)
		largeValue := bytes.Repeat([]byte{0xAB}, 70*1024)
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x7FE1))
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x0010))
		itemData.WriteString("OB")
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0))
		_ = binary.Write(itemData, binary.LittleEndian, uint32(len(largeValue)))
		_, _ = itemData.Write(largeValue)

		_ = binary.Write(seqData, binary.LittleEndian, uint32(itemData.Len()))
		seqData.Write(itemData.Bytes())

		// Sequence element (0008,1140) with defined length.
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0008))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x1140))
		buf.WriteString("SQ")
		_ = binary.Write(buf, binary.LittleEndian, uint16(0))
		_ = binary.Write(buf, binary.LittleEndian, uint32(seqData.Len()))
		buf.Write(seqData.Bytes())

		// Tail element after sequence to verify stream alignment.
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("PN")
		tailName := "Tail^Marker"
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(tailName)))
		buf.WriteString(tailName)

		result, err := Parse(buf,
			WithReadOption(ReadLargeOnDemand),
			WithLargeObjectSize(1024), // force large-object branch
		)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		elem, exists := result.Dataset.Get(tag.New(0x0008, 0x1140))
		if !exists {
			t.Fatal("Sequence element should exist")
		}

		seq, ok := elem.(*dataset.Sequence)
		if !ok {
			t.Fatalf("Element should be a Sequence, got %T", elem)
		}
		if seq.Count() != 1 {
			t.Fatalf("Sequence should have 1 item, got %d", seq.Count())
		}

		item := seq.GetItem(0)
		if item == nil {
			t.Fatal("Item should not be nil")
		}
		largeElem, exists := item.Get(tag.New(0x7FE1, 0x0010))
		if !exists {
			t.Fatal("Large OB element should exist in item")
		}
		if largeElem.Length() != uint32(len(largeValue)) {
			t.Fatalf("Large OB length = %d, want %d", largeElem.Length(), len(largeValue))
		}
		got := largeElem.Buffer().Data()
		if len(got) != len(largeValue) || got[0] != 0xAB || got[len(got)-1] != 0xAB {
			t.Fatalf("Large OB payload mismatch: len=%d first=%d last=%d", len(got), got[0], got[len(got)-1])
		}

		if name, ok := result.Dataset.GetString(tag.PatientName); !ok || name != tailName {
			t.Fatalf("tail element parse failed: got=%q exists=%v", name, ok)
		}
	})
}

// Benchmark tests for Parser

// createBenchmarkDICOMData creates test DICOM data for benchmarking
func createBenchmarkDICOMData(numElements int) []byte {
	buf := bytes.NewBuffer(nil)

	// Write preamble
	buf.Write(make([]byte, 128))
	buf.WriteString("DICM")

	// Write File Meta Information Group Length (0002,0000)
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0000))
	buf.WriteString("UL")
	_ = binary.Write(buf, binary.LittleEndian, uint16(4))
	_ = binary.Write(buf, binary.LittleEndian, uint32(0))

	// Write Transfer Syntax UID (0002,0010) - Explicit VR Little Endian
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("UI")
	tsUID := testExplicitVRLittleLE
	_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)
	if len(tsUID)%2 != 0 {
		buf.WriteByte(0)
	}

	// Write some test elements
	for i := 0; i < numElements; i++ {
		// Write tag
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		_ = binary.Write(buf, binary.LittleEndian, uint16(i))

		// Write VR (LO)
		buf.WriteString("LO")

		// Write value
		value := "TestValue"
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(value)))
		buf.WriteString(value)
		if len(value)%2 != 0 {
			buf.WriteByte(0)
		}
	}

	return buf.Bytes()
}

func BenchmarkParseSmallDataset(b *testing.B) {
	data := createBenchmarkDICOMData(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		_, _ = Parse(reader)
	}
}

func BenchmarkParseMediumDataset(b *testing.B) {
	data := createBenchmarkDICOMData(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		_, _ = Parse(reader)
	}
}

func BenchmarkParseLargeDataset(b *testing.B) {
	data := createBenchmarkDICOMData(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		_, _ = Parse(reader)
	}
}

func BenchmarkReadTag(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < 1000; i++ {
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		_ = binary.Write(buf, binary.LittleEndian, uint16(i))
	}
	data := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := newParseContext()
		ctx.reader = bytes.NewReader(data)
		ctx.byteOrder = binary.LittleEndian
		_, _ = ctx.readTag()
	}
}

func BenchmarkReadElement(b *testing.B) {
	// Create element data (PatientName with value "Doe^John")
	buf := bytes.NewBuffer(nil)
	for i := 0; i < 1000; i++ {
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
		buf.WriteString("PN")
		value := testPatientName
		_ = binary.Write(buf, binary.LittleEndian, uint16(len(value)))
		buf.WriteString(value)
	}
	data := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := newParseContext()
		ctx.reader = bytes.NewReader(data)
		ctx.byteOrder = binary.LittleEndian
		ctx.isExplicitVR = true
		_, _ = ctx.readElement()
	}
}
