// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestWriteFragmentSequence(t *testing.T) {
	t.Run("SingleFrame", testWriteFragmentSequenceSingleFrame)
	t.Run("EmptyOffsetTable", testWriteFragmentSequenceEmptyOffsetTable)
	t.Run("WithOffsetTable", testWriteFragmentSequenceWithOffsetTable)
	t.Run("AutoOffsetTableWithMultiFragmentFrames", testWriteFragmentSequenceAutoOffsetTable)
}

func roundTripOtherByteFragment(t *testing.T, obf *element.OtherByteFragment) *element.OtherByteFragment {
	t.Helper()

	ds := dataset.New()
	addTestSOPUIDs(t, ds)
	if err := ds.Add(obf); err != nil {
		t.Fatalf("Add() error: %v", err)
	}

	buf := &bytes.Buffer{}
	if err := Write(buf, ds, WithTransferSyntax(transfer.JPEG2000Lossless)); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if buf.Len() == 0 {
		t.Fatal("Write produced no output")
	}

	result, err := parser.Parse(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
	if !exists {
		t.Fatal("PixelData not found after round-trip")
	}

	readObf, ok := pixelDataElem.(*element.OtherByteFragment)
	if !ok {
		t.Fatalf("PixelData is not OtherByteFragment, got %T", pixelDataElem)
	}
	return readObf
}

func testWriteFragmentSequenceSingleFrame(t *testing.T) {
	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.AddFragment(buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04}))

	offsetTable := roundTripOtherByteFragment(t, obf).OffsetTable()
	if len(offsetTable) != 1 {
		t.Errorf("Single-frame offset table length = %d, want 1", len(offsetTable))
	}
	if len(offsetTable) > 0 && offsetTable[0] != 0 {
		t.Errorf("Single-frame offset table[0] = %d, want 0", offsetTable[0])
	}
}

func testWriteFragmentSequenceEmptyOffsetTable(t *testing.T) {
	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.AddFragment(buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04}))
	obf.AddFragment(buffer.NewMemory([]byte{0x05, 0x06, 0x07, 0x08}))

	if got := len(roundTripOtherByteFragment(t, obf).OffsetTable()); got != 0 {
		t.Errorf("OffsetTable length = %d, want 0 (empty BOT when not provided)", got)
	}
}

func testWriteFragmentSequenceWithOffsetTable(t *testing.T) {
	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 8})
	obf.AddFragment(buffer.NewMemory([]byte{0x11, 0x22, 0x33, 0x44}))
	obf.AddFragment(buffer.NewMemory([]byte{0x55, 0x66, 0x77, 0x88}))

	if got := roundTripOtherByteFragment(t, obf).OffsetTable(); len(got) != 2 || got[0] != 0 || got[1] != 8 {
		t.Errorf("OffsetTable = %v, want [0 8]", got)
	}
}

func testWriteFragmentSequenceAutoOffsetTable(t *testing.T) {
	// Two frames: Frame0 uses fragments 0-1, Frame1 uses fragment 2.
	obf := element.NewOtherByteFragment(tag.PixelData)
	// Fragment lengths intentionally odd/even to exercise padding.
	obf.AddFragment(buffer.NewMemory([]byte{0xAA, 0xBB, 0xCC}))       // len=3 -> padded 4
	obf.AddFragment(buffer.NewMemory([]byte{0x11, 0x22, 0x33, 0x44})) // len=4 -> padded 4
	obf.AddFragment(buffer.NewMemory([]byte{0x55, 0x66, 0x77}))       // len=3 -> padded 4

	if err := SetOffsetTableForFrames(obf.FragmentSequence, []int{0, 2}); err != nil {
		t.Fatalf("SetOffsetTableForFrames error: %v", err)
	}

	// Expect offsets: frame0 at 0, frame1 after two encoded fragment items:
	// (8-byte item header + padded 4-byte payload) * 2 = 24.
	if got := roundTripOtherByteFragment(t, obf).OffsetTable(); len(got) != 2 || got[0] != 0 || got[1] != 24 {
		t.Errorf("OffsetTable = %v, want [0 24]", got)
	}
}

func TestFragmentSequenceRoundTrip(t *testing.T) {
	t.Run("RoundTrip", func(t *testing.T) {
		// Create original dataset
		ds := dataset.New()
		addTestSOPUIDs(t, ds)

		// Add some metadata
		if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName})); err != nil {
			t.Fatalf("Add() error: %v", err)
		}
		if err := ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"})); err != nil {
			t.Fatalf("Add() error: %v", err)
		}

		// Create fragment sequence
		obf := element.NewOtherByteFragment(tag.PixelData)
		obf.SetOffsetTable([]uint32{0, 16})
		obf.AddFragment(buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}))
		obf.AddFragment(buffer.NewMemory([]byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}))

		if err := ds.Add(obf); err != nil {
			t.Fatalf("Add() error: %v", err)
		}

		// Write to buffer
		buf := &bytes.Buffer{}
		err := Write(buf, ds, WithTransferSyntax(transfer.JPEG2000Lossless))
		if err != nil {
			t.Fatalf("Write() error = %v", err)
		}

		// Read back
		result, err := parser.Parse(buf)
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}

		// Verify PixelData
		pixelDataElem, exists := result.Dataset.Get(tag.PixelData)
		if !exists {
			t.Fatal("PixelData not found after round-trip")
		}

		readObf, ok := pixelDataElem.(*element.OtherByteFragment)
		if !ok {
			t.Fatalf("PixelData is not OtherByteFragment after round-trip, got %T", pixelDataElem)
		}

		// Check offset table
		if offsets := readObf.OffsetTable(); len(offsets) != 2 || offsets[0] != 0 || offsets[1] != 16 {
			t.Errorf("OffsetTable = %v, want [0 16]", offsets)
		}

		// Check fragments
		if readObf.FragmentCount() != 2 {
			t.Errorf("FragmentCount = %d, want 2", readObf.FragmentCount())
		}

		// Check fragment data
		frag1, err := readObf.GetFragment(0)
		if err != nil {
			t.Errorf("GetFragment(0) error = %v", err)
		} else if frag1.Size() != 8 {
			t.Errorf("Fragment 0 size = %d, want 8", frag1.Size())
		}

		frag2, err := readObf.GetFragment(1)
		if err != nil {
			t.Errorf("GetFragment(1) error = %v", err)
		} else if frag2.Size() != 8 {
			t.Errorf("Fragment 1 size = %d, want 8", frag2.Size())
		}
	})
}
