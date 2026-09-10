// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"context"
	"encoding/binary"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

func TestDataSampleDecodesStoredBits(t *testing.T) {
	tests := []struct {
		name string
		info Info
		data []byte
		want int64
	}{
		{
			name: "unsigned 8 bit",
			info: grayscaleSampleInfo(8, 8, 7, pixel.UnsignedPixels),
			data: []byte{0xfe},
			want: 254,
		},
		{
			name: "unsigned 12 stored bits aligned at high bit 14",
			info: grayscaleSampleInfo(16, 12, 14, pixel.UnsignedPixels),
			data: littleEndian16(0x5a3 << 3),
			want: 0x5a3,
		},
		{
			name: "signed 12 stored bits",
			info: grayscaleSampleInfo(16, 12, 11, pixel.SignedPixels),
			data: littleEndian16(0x0ffe),
			want: -2,
		},
		{
			name: "signed 32 bit",
			info: grayscaleSampleInfo(32, 32, 31, pixel.SignedPixels),
			data: littleEndian32(0xfffffffe),
			want: -2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pixels, err := NewFromBytes(&test.info, test.data)
			if err != nil {
				t.Fatalf("NewFromBytes() error = %v", err)
			}
			got, err := pixels.Sample(0, 0, 0, 0)
			if err != nil {
				t.Fatalf("Sample() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("Sample() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestDataSampleNormalizesBigEndianInput(t *testing.T) {
	info := grayscaleSampleInfo(16, 16, 15, pixel.UnsignedPixels)
	info.TransferSyntaxUID = transfer.ExplicitVRBigEndian.UID().UID()
	pixels, err := NewFromBytes(&info, []byte{0x12, 0x34})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	got, err := pixels.Sample(0, 0, 0, 0)
	if err != nil {
		t.Fatalf("Sample() error = %v", err)
	}
	if got != 0x1234 {
		t.Fatalf("Sample() = %#x, want 0x1234", got)
	}
}

func TestDataSampleSupportsInterleavedAndPlanarSamples(t *testing.T) {
	for _, test := range []struct {
		name   string
		planar pixel.PlanarConfiguration
		data   []byte
	}{
		{name: "interleaved", planar: pixel.InterleavedPlanar, data: []byte{10, 20, 30, 11, 21, 31}},
		{name: "planar", planar: pixel.PlanarPlanar, data: []byte{10, 11, 20, 21, 30, 31}},
	} {
		t.Run(test.name, func(t *testing.T) {
			info := Info{
				Width: 2, Height: 1, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7,
				SamplesPerPixel: 3, PixelRepresentation: pixel.UnsignedPixels,
				PlanarConfiguration: test.planar, PhotometricInterpretation: pixel.RGBPhotometric,
				TransferSyntaxUID: transfer.ExplicitVRLittleEndian.UID().UID(),
			}
			pixels, err := NewFromBytes(&info, test.data)
			if err != nil {
				t.Fatalf("NewFromBytes() error = %v", err)
			}
			got, err := pixels.Sample(0, 1, 0, 2)
			if err != nil {
				t.Fatalf("Sample() error = %v", err)
			}
			if got != 31 {
				t.Fatalf("Sample(pixel 1, sample 2) = %d, want 31", got)
			}
		})
	}
}

func TestDataSampleRejectsInvalidAccess(t *testing.T) {
	info := grayscaleSampleInfo(16, 16, 15, pixel.UnsignedPixels)
	pixels, err := NewFromBytes(&info, littleEndian16(1))
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	for _, test := range []struct {
		name                string
		frame, x, y, sample int
		want                string
	}{
		{name: "frame", frame: 1, want: "frame index"},
		{name: "negative x", x: -1, want: "x coordinate"},
		{name: "x", x: 1, want: "x coordinate"},
		{name: "negative y", y: -1, want: "y coordinate"},
		{name: "y", y: 1, want: "y coordinate"},
		{name: "negative sample", sample: -1, want: "sample index"},
		{name: "sample", sample: 1, want: "sample index"},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := pixels.Sample(test.frame, test.x, test.y, test.sample)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Sample() error = %v, want containing %q", err, test.want)
			}
		})
	}

	encapsulatedInfo := info
	encapsulatedInfo.Encapsulated = true
	encapsulated, err := New(&encapsulatedInfo)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := encapsulated.AddFrame(context.Background(), []byte{1, 2, 3}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	if _, err := encapsulated.Sample(0, 0, 0, 0); err == nil || !strings.Contains(err.Error(), "encapsulated") {
		t.Fatalf("Sample() error = %v, want encapsulated rejection", err)
	}
}

func TestDataIsPaddingSampleUsesInclusiveUnorderedRange(t *testing.T) {
	paddingValue := int32(20)
	rangeLimit := int32(10)
	info := grayscaleSampleInfo(16, 16, 15, pixel.SignedPixels)
	info.PixelPaddingValue = &paddingValue
	info.PixelPaddingRangeLimit = &rangeLimit
	pixels, err := NewFromBytes(&info, littleEndian16(15))
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if !pixels.IsPaddingSample(10) || !pixels.IsPaddingSample(15) || !pixels.IsPaddingSample(20) {
		t.Fatal("IsPaddingSample() excluded a value in the inclusive padding range")
	}
	if pixels.IsPaddingSample(9) || pixels.IsPaddingSample(21) {
		t.Fatal("IsPaddingSample() included a value outside the padding range")
	}
}

func grayscaleSampleInfo(bitsAllocated, bitsStored, highBit uint16, representation pixel.Representation) Info {
	return Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: bitsAllocated, BitsStored: bitsStored, HighBit: highBit,
		SamplesPerPixel: 1, PixelRepresentation: representation,
		PhotometricInterpretation: pixel.Monochrome2,
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
	}
}

func littleEndian16(value uint16) []byte {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, value)
	return data
}

func littleEndian32(value uint32) []byte {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)
	return data
}
