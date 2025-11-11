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
	t.Run("EmptyOffsetTable", func(t *testing.T) {
		// Create dataset with fragment sequence
        ds := dataset.New()

		// Create fragment sequence
        obf := element.NewOtherByteFragment(tag.PixelData)
        obf.AddFragment(buffer.NewMemory([]byte{0x01, 0x02, 0x03, 0x04}))
        obf.AddFragment(buffer.NewMemory([]byte{0x05, 0x06, 0x07, 0x08}))

        if err := ds.Add(obf); err != nil {
            t.Fatalf("Add() error: %v", err)
        }

		// Write to buffer
		buf := &bytes.Buffer{}
		err := Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
		if err != nil {
			t.Fatalf("Write() error = %v", err)
		}

		// Verify we wrote something
		if buf.Len() == 0 {
			t.Fatal("Write produced no output")
		}
	})

	t.Run("WithOffsetTable", func(t *testing.T) {
        ds := dataset.New()

		// Create fragment sequence with offset table
        obf := element.NewOtherByteFragment(tag.PixelData)
        obf.SetOffsetTable([]uint32{0, 8})
        obf.AddFragment(buffer.NewMemory([]byte{0x11, 0x22, 0x33, 0x44}))
        obf.AddFragment(buffer.NewMemory([]byte{0x55, 0x66, 0x77, 0x88}))

        if err := ds.Add(obf); err != nil {
            t.Fatalf("Add() error: %v", err)
        }

		// Write
		buf := &bytes.Buffer{}
		err := Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
		if err != nil {
			t.Fatalf("Write() error = %v", err)
		}

		if buf.Len() == 0 {
			t.Fatal("Write produced no output")
		}
	})
}

func TestFragmentSequenceRoundTrip(t *testing.T) {
	t.Run("RoundTrip", func(t *testing.T) {
		// Create original dataset
        ds := dataset.New()

		// Add some metadata
        if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"})); err != nil {
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
		err := Write(buf, ds, WithTransferSyntax(transfer.ExplicitVRLittleEndian))
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
		if len(readObf.OffsetTable()) != 2 {
			t.Errorf("OffsetTable length = %d, want 2", len(readObf.OffsetTable()))
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
