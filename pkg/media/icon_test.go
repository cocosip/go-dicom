// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewIconImageSequenceBuildsEightBitMonochromeItem(t *testing.T) {
	pixels := []byte{0, 64, 128, 255, 10, 20}
	sequence, err := NewIconImageSequence(3, 2, pixels)
	if err != nil {
		t.Fatalf("NewIconImageSequence() error = %v", err)
	}
	if !sequence.Tag().Equals(tag.IconImageSequence) || sequence.Count() != 1 {
		t.Fatalf("sequence = %s with %d item(s), want IconImageSequence with one item", sequence.Tag(), sequence.Count())
	}
	item := sequence.GetItem(0)
	checks := []struct {
		name string
		got  uint16
		want uint16
	}{
		{name: "Rows", got: item.TryGetUInt16(tag.Rows, 0), want: 2},
		{name: "Columns", got: item.TryGetUInt16(tag.Columns, 0), want: 3},
		{name: "SamplesPerPixel", got: item.TryGetUInt16(tag.SamplesPerPixel, 0), want: 1},
		{name: "BitsAllocated", got: item.TryGetUInt16(tag.BitsAllocated, 0), want: 8},
		{name: "BitsStored", got: item.TryGetUInt16(tag.BitsStored, 0), want: 8},
		{name: "HighBit", got: item.TryGetUInt16(tag.HighBit, 0), want: 7},
		{name: "PixelRepresentation", got: item.TryGetUInt16(tag.PixelRepresentation, 0), want: 0},
	}
	for _, check := range checks {
		if check.got != check.want {
			t.Errorf("%s = %d, want %d", check.name, check.got, check.want)
		}
	}
	if got := item.TryGetString(tag.PhotometricInterpretation); got != "MONOCHROME2" {
		t.Fatalf("PhotometricInterpretation = %q", got)
	}
	if got := item.TryGetString(tag.NumberOfFrames); got != "1" {
		t.Fatalf("NumberOfFrames = %q, want 1", got)
	}
	pixelElement, ok := item.Get(tag.PixelData)
	if !ok {
		t.Fatal("PixelData is missing")
	}
	if pixelElement.ValueRepresentation() != vr.OB {
		t.Fatalf("PixelData VR = %s, want OB", pixelElement.ValueRepresentation())
	}
	otherByte, ok := pixelElement.(interface{ GetData() []byte })
	if !ok {
		t.Fatalf("PixelData = %T, want byte data element", pixelElement)
	}
	if got := otherByte.GetData(); !bytes.Equal(got, pixels) {
		t.Fatalf("PixelData = %v, want %v", got, pixels)
	}
}

func TestNewIconImageSequenceRejectsInvalidDimensionsAndPixels(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
		pixels        []byte
	}{
		{name: "zero width", width: 0, height: 1, pixels: nil},
		{name: "negative height", width: 1, height: -1, pixels: nil},
		{name: "width above maximum", width: 129, height: 1, pixels: make([]byte, 129)},
		{name: "height above maximum", width: 1, height: 129, pixels: make([]byte, 129)},
		{name: "short pixels", width: 2, height: 2, pixels: make([]byte, 3)},
		{name: "long pixels", width: 2, height: 2, pixels: make([]byte, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewIconImageSequence(tt.width, tt.height, tt.pixels); err == nil {
				t.Fatal("NewIconImageSequence() succeeded")
			}
		})
	}
}
