// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package pixeldata

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// TestConvertMonochrome1ToMonochrome2InvertsEveryPackedBit verifies IMG-05:
// 1-bit samples are bit-packed, so MONOCHROME1 -> MONOCHROME2 complements every
// sample bit (0xAA -> 0x55), not only bit 0.
func TestConvertMonochrome1ToMonochrome2InvertsEveryPackedBit(t *testing.T) {
	info := &Info{
		Width: 8, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 1, BitsStored: 1, HighBit: 0,
		SamplesPerPixel:           1,
		PhotometricInterpretation: pixel.Monochrome1,
		VRCode:                    "OB",
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
	}
	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{0xAA}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	if err := pd.ConvertMonochrome1ToMonochrome2(); err != nil {
		t.Fatalf("ConvertMonochrome1ToMonochrome2() error = %v", err)
	}
	frame, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}
	if !bytes.Equal(frame, []byte{0x55}) {
		t.Fatalf("converted frame = %#x, want 0x55", frame)
	}
}

// TestConvertMonochrome1ToMonochrome2MasksPackedPadding verifies that unused
// high bits of the final byte are cleared after complementing.
func TestConvertMonochrome1ToMonochrome2MasksPackedPadding(t *testing.T) {
	info := &Info{
		Width: 5, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 1, BitsStored: 1, HighBit: 0,
		SamplesPerPixel:           1,
		PhotometricInterpretation: pixel.Monochrome1,
		VRCode:                    "OB",
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
	}
	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{0xAA}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	if err := pd.ConvertMonochrome1ToMonochrome2(); err != nil {
		t.Fatalf("ConvertMonochrome1ToMonochrome2() error = %v", err)
	}
	frame, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}
	if !bytes.Equal(frame, []byte{0x15}) {
		t.Fatalf("converted frame = %#x, want 0x15 (5 complemented samples, padding cleared)", frame)
	}
}

// TestConvertYBRToRGBRejectsSixteenBitInput verifies IMG-07: the 8-bit YBR
// kernels must not silently reinterpret 16-bit samples.
func TestConvertYBRToRGBRejectsSixteenBitInput(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15,
		SamplesPerPixel:           3,
		PhotometricInterpretation: pixel.YbrFull,
		VRCode:                    "OW",
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
	}
	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{0, 0, 0, 0, 0, 0}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	err = pd.ConvertYBRToRGB()
	if err == nil || !strings.Contains(err.Error(), "BitsAllocated=8") {
		t.Fatalf("ConvertYBRToRGB() error = %v, want a BitsAllocated=8 rejection", err)
	}
}

// TestConvertPaletteToRGBRejectsEncapsulatedInput verifies IMG-08: a palette
// conversion must not run over compressed frames.
func TestConvertPaletteToRGBRejectsEncapsulatedInput(t *testing.T) {
	pd := &Data{Info: &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7,
		SamplesPerPixel:           1,
		PhotometricInterpretation: pixel.PaletteColor,
		VRCode:                    "OB",
		Encapsulated:              true,
		TransferSyntaxUID:         transfer.JPEG2000Lossless.UID().UID(),
	}}
	err := ConvertPaletteToRGB(nil, pd)
	if err == nil || !strings.Contains(err.Error(), "encapsulated") {
		t.Fatalf("ConvertPaletteToRGB() error = %v, want an encapsulated-data rejection", err)
	}
}
