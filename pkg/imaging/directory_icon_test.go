// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"image/color"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestDirectoryIconGeneratorPreservesAspectRatioWithin128Pixels(t *testing.T) {
	tests := []struct {
		name                  string
		width, height         int
		wantWidth, wantHeight int
	}{
		{name: "landscape", width: 200, height: 100, wantWidth: 128, wantHeight: 64},
		{name: "portrait", width: 100, height: 200, wantWidth: 64, wantHeight: 128},
		{name: "square", width: 200, height: 200, wantWidth: 128, wantHeight: 128},
		{name: "small image", width: 40, height: 20, wantWidth: 40, wantHeight: 20},
		{name: "thin image", width: 300, height: 1, wantWidth: 128, wantHeight: 1},
	}
	generator := NewDirectoryIconGenerator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixels := make([]byte, tt.width*tt.height)
			for i := range pixels {
				pixels[i] = byte(i % 256)
			}
			ds := iconTestDataset(t, tt.width, tt.height, 1, "MONOCHROME2", pixels)
			width, height, got, err := generator.GenerateDirectoryIcon(ds, 0)
			if err != nil {
				t.Fatalf("GenerateDirectoryIcon() error = %v", err)
			}
			if width != tt.wantWidth || height != tt.wantHeight {
				t.Fatalf("dimensions = %dx%d, want %dx%d", width, height, tt.wantWidth, tt.wantHeight)
			}
			if len(got) != width*height {
				t.Fatalf("pixel length = %d, want %d", len(got), width*height)
			}
			_, _, again, err := generator.GenerateDirectoryIcon(ds, 0)
			if err != nil {
				t.Fatalf("second GenerateDirectoryIcon() error = %v", err)
			}
			if string(got) != string(again) {
				t.Fatal("icon generation is not deterministic")
			}
		})
	}
}

func TestDirectoryIconGeneratorRendersMonochromeAndColor(t *testing.T) {
	generator := NewDirectoryIconGenerator()

	mono := iconTestDataset(t, 2, 1, 1, "MONOCHROME1", []byte{0, 255})
	_, _, monoPixels, err := generator.GenerateDirectoryIcon(mono, 0)
	if err != nil {
		t.Fatalf("MONOCHROME1 GenerateDirectoryIcon() error = %v", err)
	}
	if len(monoPixels) != 2 || monoPixels[0] != 255 || monoPixels[1] != 0 {
		t.Fatalf("MONOCHROME1 pixels = %v, want [255 0]", monoPixels)
	}

	rgb := iconTestDataset(t, 1, 1, 3, "RGB", []byte{255, 0, 0})
	_, _, rgbPixels, err := generator.GenerateDirectoryIcon(rgb, 0)
	if err != nil {
		t.Fatalf("RGB GenerateDirectoryIcon() error = %v", err)
	}
	wantGray := color.GrayModel.Convert(color.RGBA{R: 255, A: 255}).(color.Gray).Y
	if len(rgbPixels) != 1 || rgbPixels[0] != wantGray {
		t.Fatalf("RGB grayscale pixels = %v, want [%d]", rgbPixels, wantGray)
	}
}

func TestDirectoryIconGeneratorUsesRequestedFrame(t *testing.T) {
	ds := iconTestDataset(t, 2, 1, 1, "MONOCHROME2", []byte{0, 255, 255, 0})
	if err := ds.AddOrUpdate(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
		t.Fatalf("set NumberOfFrames: %v", err)
	}

	_, _, pixels, err := NewDirectoryIconGenerator().GenerateDirectoryIcon(ds, 1)
	if err != nil {
		t.Fatalf("GenerateDirectoryIcon(frame=1) error = %v", err)
	}
	if len(pixels) != 2 || pixels[0] != 255 || pixels[1] != 0 {
		t.Fatalf("frame 1 pixels = %v, want [255 0]", pixels)
	}
}

func TestDirectoryIconGeneratorRejectsInvalidFrame(t *testing.T) {
	ds := iconTestDataset(t, 1, 1, 1, "MONOCHROME2", []byte{0})
	for _, frame := range []int{-1, 1} {
		if _, _, _, err := NewDirectoryIconGenerator().GenerateDirectoryIcon(ds, frame); err == nil {
			t.Fatalf("GenerateDirectoryIcon(frame=%d) succeeded", frame)
		}
	}
}

func iconTestDataset(t *testing.T, width, height, samples int, photometric string, pixels []byte) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	items := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{uint16(height)}),             // #nosec G115 -- test dimensions are bounded
		element.NewUnsignedShort(tag.Columns, []uint16{uint16(width)}),           // #nosec G115 -- test dimensions are bounded
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{uint16(samples)}), // #nosec G115 -- test values are bounded
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometric}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewOtherByte(tag.PixelData, append([]byte(nil), pixels...)),
	}
	if samples > 1 {
		items = append(items, element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}))
	}
	for _, item := range items {
		if err := ds.Add(item); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", item.Tag(), err)
		}
	}
	return ds
}
