// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// Helper function to create a DICOM file with fragment sequence
func createFragmentSequenceDICOM() *bytes.Buffer {
	buf := bytes.NewBuffer(nil)

	// DICOM Preamble (128 bytes)
	buf.Write(make([]byte, 128))

	// DICOM Prefix "DICM"
	buf.WriteString("DICM")

	// File Meta Information Group Length (0002,0000)
	binary.Write(buf, binary.LittleEndian, uint16(0x0002)) // Group
	binary.Write(buf, binary.LittleEndian, uint16(0x0000)) // Element
	buf.WriteString("UL")                                   // VR
	binary.Write(buf, binary.LittleEndian, uint16(4))      // Length
	binary.Write(buf, binary.LittleEndian, uint32(0))      // Value (placeholder, will calculate later)

	// Transfer Syntax UID (0002,0010) - Explicit VR Little Endian
	binary.Write(buf, binary.LittleEndian, uint16(0x0002))
	binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("UI")
	tsUID := "1.2.840.10008.1.2.1\x00" // Explicit VR Little Endian, padded
	binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)

	// Now add PixelData as Fragment Sequence
	// PixelData tag (7FE0,0010)
	binary.Write(buf, binary.LittleEndian, uint16(0x7FE0))
	binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("OB")                                     // VR
	binary.Write(buf, binary.LittleEndian, uint16(0))        // Reserved
	binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF)) // Undefined length

	// Item 1: Offset Table (FFFE,E000) - empty in this test
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	binary.Write(buf, binary.LittleEndian, uint32(0)) // Length = 0 (empty offset table)

	// Item 2: First fragment (FFFE,E000)
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	fragment1 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	binary.Write(buf, binary.LittleEndian, uint32(len(fragment1)))
	buf.Write(fragment1)

	// Item 3: Second fragment (FFFE,E000)
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	fragment2 := []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}
	binary.Write(buf, binary.LittleEndian, uint32(len(fragment2)))
	buf.Write(fragment2)

	// Sequence Delimitation Item (FFFE,E0DD)
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
	binary.Write(buf, binary.LittleEndian, uint32(0)) // Length = 0

	return buf
}

// Helper function to create fragment sequence with offset table
func createFragmentSequenceWithOffsetTable() *bytes.Buffer {
	buf := bytes.NewBuffer(nil)

	// DICOM Preamble (128 bytes)
	buf.Write(make([]byte, 128))

	// DICOM Prefix "DICM"
	buf.WriteString("DICM")

	// Minimal File Meta Information
	// Transfer Syntax UID (0002,0010)
	binary.Write(buf, binary.LittleEndian, uint16(0x0002))
	binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("UI")
	tsUID := "1.2.840.10008.1.2.1\x00"
	binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)

	// PixelData tag (7FE0,0010) as Fragment Sequence
	binary.Write(buf, binary.LittleEndian, uint16(0x7FE0))
	binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("OB")
	binary.Write(buf, binary.LittleEndian, uint16(0))
	binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF))

	// Item 1: Offset Table with 2 offsets
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	binary.Write(buf, binary.LittleEndian, uint32(8)) // Length = 8 (2 offsets * 4 bytes)
	binary.Write(buf, binary.LittleEndian, uint32(0))  // Offset to frame 1
	binary.Write(buf, binary.LittleEndian, uint32(16)) // Offset to frame 2

	// Item 2: Fragment 1
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	fragment1 := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	binary.Write(buf, binary.LittleEndian, uint32(len(fragment1)))
	buf.Write(fragment1)

	// Item 3: Fragment 2
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	fragment2 := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00, 0x11}
	binary.Write(buf, binary.LittleEndian, uint32(len(fragment2)))
	buf.Write(fragment2)

	// Sequence Delimitation Item
	binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
	binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
	binary.Write(buf, binary.LittleEndian, uint32(0))

	return buf
}

func TestParseFragmentSequence(t *testing.T) {
	t.Run("EmptyOffsetTable", func(t *testing.T) {
		data := createFragmentSequenceDICOM()
		result, err := Parse(data)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		// Get PixelData element
		pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
		if !exists {
			t.Fatal("PixelData element not found")
		}

		// Should be OtherByteFragment
		obf, ok := pixelDataElem.(*element.OtherByteFragment)
		if !ok {
			t.Fatalf("PixelData is not OtherByteFragment, got %T", pixelDataElem)
		}

		// Check offset table (should be empty)
		if len(obf.OffsetTable()) != 0 {
			t.Errorf("OffsetTable length = %d, want 0", len(obf.OffsetTable()))
		}

		// Check fragments
		if obf.FragmentCount() != 2 {
			t.Errorf("FragmentCount() = %d, want 2", obf.FragmentCount())
		}

		// Check fragment 1 data
		frag1, err := obf.GetFragment(0)
		if err != nil {
			t.Errorf("GetFragment(0) error = %v", err)
		}
		if frag1.Size() != 8 {
			t.Errorf("Fragment 0 size = %d, want 8", frag1.Size())
		}

		// Check fragment 2 data
		frag2, err := obf.GetFragment(1)
		if err != nil {
			t.Errorf("GetFragment(1) error = %v", err)
		}
		if frag2.Size() != 8 {
			t.Errorf("Fragment 1 size = %d, want 8", frag2.Size())
		}
	})

	t.Run("WithOffsetTable", func(t *testing.T) {
		data := createFragmentSequenceWithOffsetTable()
		result, err := Parse(data)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		// Get PixelData element
		pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
		if !exists {
			t.Fatal("PixelData element not found")
		}

		obf, ok := pixelDataElem.(*element.OtherByteFragment)
		if !ok {
			t.Fatalf("PixelData is not OtherByteFragment, got %T", pixelDataElem)
		}

		// Check offset table
		offsets := obf.OffsetTable()
		if len(offsets) != 2 {
			t.Errorf("OffsetTable length = %d, want 2", len(offsets))
		}

		if len(offsets) == 2 {
			if offsets[0] != 0 {
				t.Errorf("Offset[0] = %d, want 0", offsets[0])
			}
			if offsets[1] != 16 {
				t.Errorf("Offset[1] = %d, want 16", offsets[1])
			}
		}

		// Check fragments
		if obf.FragmentCount() != 2 {
			t.Errorf("FragmentCount() = %d, want 2", obf.FragmentCount())
		}
	})
}
