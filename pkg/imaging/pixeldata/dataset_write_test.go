// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"bytes"
	"context"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func TestNewForDatasetAndWriteToDatasetInferPixelDataRepresentation(t *testing.T) {
	tests := []struct {
		name          string
		syntax        *transfer.Syntax
		bitsAllocated uint16
		wantVR        string
		wantElement   any
	}{
		{"explicit 8-bit", transfer.ExplicitVRLittleEndian, 8, "OB", (*element.OtherByte)(nil)},
		{"explicit 16-bit", transfer.ExplicitVRLittleEndian, 16, "OW", (*element.OtherWord)(nil)},
		{"implicit 8-bit", transfer.ImplicitVRLittleEndian, 8, "OW", (*element.OtherWord)(nil)},
		{"encapsulated", transfer.JPEG2000Lossless, 16, "OB", (*element.OtherByteFragment)(nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := newWritablePixelDataset(t, tt.syntax, tt.bitsAllocated)
			pixels, err := NewForDataset(ds)
			if err != nil {
				t.Fatalf("NewForDataset() error = %v", err)
			}
			if pixels.Info.VRCode != tt.wantVR {
				t.Fatalf("Info.VRCode = %q, want %q", pixels.Info.VRCode, tt.wantVR)
			}

			frameSize := 2
			if tt.bitsAllocated > 8 {
				frameSize = 4
			}
			frame := make([]byte, frameSize)
			if err := pixels.AddFrame(context.Background(), frame); err != nil {
				t.Fatalf("AddFrame() error = %v", err)
			}
			if err := pixels.WriteToDataset(ds); err != nil {
				t.Fatalf("WriteToDataset() error = %v", err)
			}

			got, ok := ds.Get(tag.PixelData)
			if !ok {
				t.Fatal("WriteToDataset() did not add Pixel Data")
			}
			switch tt.wantElement.(type) {
			case *element.OtherByte:
				if _, ok := got.(*element.OtherByte); !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherByte", got)
				}
			case *element.OtherWord:
				if _, ok := got.(*element.OtherWord); !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherWord", got)
				}
			case *element.OtherByteFragment:
				if _, ok := got.(*element.OtherByteFragment); !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherByteFragment", got)
				}
			}
			if gotFrames, ok := ds.GetString(tag.NumberOfFrames); !ok || gotFrames != "1" {
				t.Fatalf("NumberOfFrames = %q, %v; want 1, true", gotFrames, ok)
			}
		})
	}
}

func TestNewForDatasetRequiresTransferSyntaxAndBitsAllocated(t *testing.T) {
	t.Run("transfer syntax", func(t *testing.T) {
		ds := newWritablePixelDataset(t, transfer.ExplicitVRLittleEndian, 8)
		ds.SetInternalTransferSyntax(nil)
		if _, err := NewForDataset(ds); err == nil {
			t.Fatal("NewForDataset() error = nil, want missing Transfer Syntax error")
		}
	})

	t.Run("bits allocated", func(t *testing.T) {
		ds := newWritablePixelDataset(t, transfer.ExplicitVRLittleEndian, 8)
		ds.Remove(tag.BitsAllocated)
		if _, err := NewForDataset(ds); err == nil {
			t.Fatal("NewForDataset() error = nil, want missing Bits Allocated error")
		}
	})
}

func TestWriteToDatasetRejectsTransferSyntaxMismatch(t *testing.T) {
	source := newWritablePixelDataset(t, transfer.ExplicitVRLittleEndian, 8)
	pixels, err := NewForDataset(source)
	if err != nil {
		t.Fatalf("NewForDataset() error = %v", err)
	}
	if err := pixels.AddFrame(context.Background(), []byte{1, 2}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	target := newWritablePixelDataset(t, transfer.ExplicitVRBigEndian, 8)
	if err := pixels.WriteToDataset(target); err == nil {
		t.Fatal("WriteToDataset() error = nil, want Transfer Syntax mismatch error")
	}
	if target.Contains(tag.PixelData) {
		t.Fatal("WriteToDataset() modified target after Transfer Syntax mismatch")
	}
}

func TestWriteToDatasetRejectsImageMetadataMismatch(t *testing.T) {
	source := newWritablePixelDataset(t, transfer.ExplicitVRLittleEndian, 8)
	pixels, err := NewForDataset(source)
	if err != nil {
		t.Fatalf("NewForDataset() error = %v", err)
	}
	if err := pixels.AddFrame(context.Background(), []byte{1, 2}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	target := newWritablePixelDataset(t, transfer.ExplicitVRLittleEndian, 8)
	if err := target.AddOrUpdateValueWithVR(tag.Columns, vr.US, uint16(3)); err != nil {
		t.Fatalf("update Columns: %v", err)
	}
	if err := pixels.WriteToDataset(target); err == nil {
		t.Fatal("WriteToDataset() error = nil, want image metadata mismatch error")
	}
	if target.Contains(tag.PixelData) {
		t.Fatal("WriteToDataset() modified target after rejecting image metadata")
	}
}

func TestWriteToDatasetRoundTripsInferredRepresentation(t *testing.T) {
	tests := []struct {
		name          string
		syntax        *transfer.Syntax
		bitsAllocated uint16
		frame         []byte
		assertElement func(*testing.T, element.Element)
	}{
		{
			name:   "native OB",
			syntax: transfer.ExplicitVRLittleEndian, bitsAllocated: 8,
			frame: []byte{1, 2},
			assertElement: func(t *testing.T, elem element.Element) {
				if _, ok := elem.(*element.OtherByte); !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherByte", elem)
				}
			},
		},
		{
			name:   "native OW big endian",
			syntax: transfer.ExplicitVRBigEndian, bitsAllocated: 16,
			frame: []byte{0x34, 0x12, 0x78, 0x56},
			assertElement: func(t *testing.T, elem element.Element) {
				word, ok := elem.(*element.OtherWord)
				if !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherWord", elem)
				}
				if got := word.GetData(); !bytes.Equal(got, []byte{0x12, 0x34, 0x56, 0x78}) {
					t.Fatalf("big-endian Pixel Data bytes = %x, want 12345678", got)
				}
			},
		},
		{
			name:   "encapsulated OB fragments",
			syntax: transfer.JPEG2000Lossless, bitsAllocated: 16,
			frame: []byte{0xff, 0x4f, 0xff, 0x51},
			assertElement: func(t *testing.T, elem element.Element) {
				fragments, ok := elem.(*element.OtherByteFragment)
				if !ok {
					t.Fatalf("Pixel Data = %T, want *element.OtherByteFragment", elem)
				}
				if fragments.FragmentCount() != 1 {
					t.Fatalf("FragmentCount() = %d, want 1", fragments.FragmentCount())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := newWritablePixelDataset(t, tt.syntax, tt.bitsAllocated)
			if err := ds.AddValue(tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.2"); err != nil {
				t.Fatalf("AddValue(SOPClassUID) error = %v", err)
			}
			if err := ds.AddValue(tag.SOPInstanceUID, "1.2.826.0.1.3680043.10.543.1"); err != nil {
				t.Fatalf("AddValue(SOPInstanceUID) error = %v", err)
			}
			pixels, err := NewForDataset(ds)
			if err != nil {
				t.Fatalf("NewForDataset() error = %v", err)
			}
			if err := pixels.AddFrame(context.Background(), tt.frame); err != nil {
				t.Fatalf("AddFrame() error = %v", err)
			}
			if err := pixels.WriteToDataset(ds); err != nil {
				t.Fatalf("WriteToDataset() error = %v", err)
			}

			var encoded bytes.Buffer
			if err := writer.Write(&encoded, ds, writer.WithTransferSyntax(tt.syntax)); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			result, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			pixelElement, ok := result.Dataset.Get(tag.PixelData)
			if !ok {
				t.Fatal("round-tripped Dataset is missing Pixel Data")
			}
			tt.assertElement(t, pixelElement)
		})
	}
}

func TestWriteToDatasetRoundTripsMultipleFrames(t *testing.T) {
	for _, tt := range []struct {
		name   string
		syntax *transfer.Syntax
		frames [][]byte
	}{
		{
			name:   "native",
			syntax: transfer.ExplicitVRLittleEndian,
			frames: [][]byte{{1, 2}, {3, 4}},
		},
		{
			name:   "encapsulated",
			syntax: transfer.JPEG2000Lossless,
			frames: [][]byte{{0xff, 0x4f, 1, 2}, {0xff, 0x4f, 3, 4}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := newWritablePixelDataset(t, tt.syntax, 8)
			if err := ds.AddValue(tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.2"); err != nil {
				t.Fatalf("AddValue(SOPClassUID) error = %v", err)
			}
			if err := ds.AddValue(tag.SOPInstanceUID, "1.2.826.0.1.3680043.10.543.2"); err != nil {
				t.Fatalf("AddValue(SOPInstanceUID) error = %v", err)
			}
			pixels, err := NewForDataset(ds)
			if err != nil {
				t.Fatalf("NewForDataset() error = %v", err)
			}
			for _, frame := range tt.frames {
				if err := pixels.AddFrame(context.Background(), frame); err != nil {
					t.Fatalf("AddFrame() error = %v", err)
				}
			}
			if err := pixels.WriteToDataset(ds); err != nil {
				t.Fatalf("WriteToDataset() error = %v", err)
			}
			if got, ok := ds.GetString(tag.NumberOfFrames); !ok || got != "2" {
				t.Fatalf("NumberOfFrames = %q, %v; want 2, true", got, ok)
			}

			var encoded bytes.Buffer
			if err := writer.Write(&encoded, ds, writer.WithTransferSyntax(tt.syntax)); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			result, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			roundTripped, err := FromDataset(result.Dataset)
			if err != nil {
				t.Fatalf("FromDataset() error = %v", err)
			}
			if roundTripped.FrameCount() != len(tt.frames) {
				t.Fatalf("FrameCount() = %d, want %d", roundTripped.FrameCount(), len(tt.frames))
			}
			for i, want := range tt.frames {
				got, err := roundTripped.Frame(context.Background(), i)
				if err != nil {
					t.Fatalf("Frame(%d) error = %v", i, err)
				}
				if !bytes.Equal(got, want) {
					t.Fatalf("Frame(%d) = %x, want %x", i, got, want)
				}
			}
		})
	}
}

func newWritablePixelDataset(t *testing.T, syntax *transfer.Syntax, bitsAllocated uint16) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(syntax)
	values := []struct {
		tag   *tag.Tag
		value any
	}{
		{tag.Rows, uint16(1)},
		{tag.Columns, uint16(2)},
		{tag.BitsAllocated, bitsAllocated},
		{tag.BitsStored, bitsAllocated},
		{tag.HighBit, bitsAllocated - 1},
		{tag.SamplesPerPixel, uint16(1)},
		{tag.PixelRepresentation, uint16(0)},
		{tag.PhotometricInterpretation, "MONOCHROME2"},
	}
	for _, value := range values {
		if err := ds.AddValue(value.tag, value.value); err != nil {
			t.Fatalf("AddValue(%s) error = %v", value.tag, err)
		}
	}
	return ds
}
