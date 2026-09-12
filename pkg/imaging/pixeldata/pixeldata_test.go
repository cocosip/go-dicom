// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/logging"
)

func addLegacyTestElement(ds *dataset.Dataset, elem element.Element) error {
	ds.SetAutoValidate(false)
	defer ds.SetAutoValidate(true)
	return ds.Add(elem)
}

func TestInfo_Validate(t *testing.T) {
	tests := []struct {
		name      string
		info      *Info
		expectErr bool
	}{
		{
			name: "valid grayscale",
			info: &Info{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             16,
				BitsStored:                12,
				HighBit:                   11,
				SamplesPerPixel:           1,
				PixelRepresentation:       pixel.UnsignedPixels,
				PlanarConfiguration:       pixel.InterleavedPlanar,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectErr: false,
		},
		{
			name: "valid RGB",
			info: &Info{
				Width:                     256,
				Height:                    256,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           3,
				PixelRepresentation:       pixel.UnsignedPixels,
				PlanarConfiguration:       pixel.InterleavedPlanar,
				PhotometricInterpretation: pixel.RGBPhotometric,
			},
			expectErr: false,
		},
		{
			name: "zero width",
			info: &Info{
				Width:                     0,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "bits stored exceeds bits allocated",
			info: &Info{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                16,
				HighBit:                   7,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "high bit >= bits allocated",
			info: &Info{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   8,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "color with insufficient samples",
			info: &Info{
				Width:                     256,
				Height:                    256,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           1, // Should be 3 for RGB
				PixelRepresentation:       pixel.UnsignedPixels,
				PlanarConfiguration:       pixel.InterleavedPlanar,
				PhotometricInterpretation: pixel.RGBPhotometric,
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.info.Validate()
			if tt.expectErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestFromDatasetUsesExtendedOffsetTableForMultiFragmentFrames(t *testing.T) {
	ds := newEncapsulatedPixelDataDataset(t, 2)
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, uint64Values(0, 20))); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, uint64Values(20, 20))); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	for _, data := range [][]byte{[]byte("AA"), []byte("BB"), []byte("CC"), []byte("DD")} {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	for i, want := range [][]byte{[]byte("AABB"), []byte("CCDD")} {
		got, err := pd.Frame(context.Background(), i)
		if err != nil {
			t.Fatalf("Frame(%d) error = %v", i, err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("Frame(%d) = %q, want %q", i, got, want)
		}
	}
}

func TestFromDatasetRejectsAmbiguousEmptyOffsetTable(t *testing.T) {
	ds := newEncapsulatedPixelDataDataset(t, 2)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	for _, data := range [][]byte{[]byte("AA"), []byte("BB"), []byte("CC"), []byte("DD")} {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	if _, err := FromDataset(ds); err == nil {
		t.Fatal("FromDataset() error = nil, want indeterminate frame-boundaries error")
	}
}

func TestFromDatasetStandardModeRejectsEncapsulatedOW(t *testing.T) {
	ds := newEncapsulatedPixelDataDataset(t, 1)
	fragments := element.NewOtherWordFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	_, err := FromDatasetWithOptions(ds, WithPixelDataVRMode(PixelDataStandard))
	if err == nil {
		t.Fatal("FromDatasetWithOptions() error = nil, want encapsulated OW rejection")
	}
}

func TestFromDatasetDefaultsToCompatibleEncapsulatedOW(t *testing.T) {
	ds := newEncapsulatedPixelDataDataset(t, 1)
	fragments := element.NewOtherWordFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	if pd.Info.VRCode != "OW" {
		t.Fatalf("VRCode = %q, want preserved OW", pd.Info.VRCode)
	}
}

func TestFromDatasetStandardModeRejectsNativeOBAboveEightBits(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	_, err := FromDatasetWithOptions(ds, WithPixelDataVRMode(PixelDataStandard))
	if err == nil || !strings.Contains(err.Error(), "native Pixel Data uses OB") {
		t.Fatalf("FromDatasetWithOptions() error = %v, want native OB rejection", err)
	}
}

func TestFromDatasetStandardModeRejectsImplicitNativeOB(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ImplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x7f, 0x00}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	_, err := FromDatasetWithOptions(ds, WithPixelDataVRMode(PixelDataStandard))
	if err == nil || !strings.Contains(err.Error(), "implicit VR native Pixel Data uses OB") {
		t.Fatalf("FromDatasetWithOptions() error = %v, want implicit native OB rejection", err)
	}
}

func TestFromDatasetCompatibleModeWarnsForNativeOBAboveEightBits(t *testing.T) {
	var output bytes.Buffer
	logging.Disable()
	if err := logging.Configure(logging.Config{
		Handler: slog.NewJSONHandler(&output, &slog.HandlerOptions{Level: slog.LevelWarn}),
	}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(logging.Disable)

	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	if _, err := FromDataset(ds); err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	if got := output.String(); !strings.Contains(got, `"event":"nonstandard_native_pixel_data_vr"`) {
		t.Fatalf("warning log missing native Pixel Data VR event: %s", got)
	}
}

func newEncapsulatedPixelDataDataset(t *testing.T, frames int) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.JPEG2000Lossless)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewString(tag.NumberOfFrames, vr.IS, []string{fmt.Sprint(frames)}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	return ds
}

func uint64Values(values ...uint64) []byte {
	data := make([]byte, len(values)*8)
	for i, value := range values {
		binary.LittleEndian.PutUint64(data[i*8:], value)
	}
	return data
}

func TestFromDatasetNormalizesBigEndianNativeOW(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{16}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{15}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))
	_ = ds.Add(element.NewOtherWord(tag.PixelData, []byte{0x12, 0x34}))

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	frame, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame(0) error = %v", err)
	}
	if got, want := frame, []byte{0x34, 0x12}; !bytes.Equal(got, want) {
		t.Fatalf("frame = %v, want little-endian normalized %v", got, want)
	}
}

func TestMinMaxSamplesSignExtendsBitsStored(t *testing.T) {
	pd, err := New(&Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                12,
		HighBit:                   11,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.SignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{0xFF, 0x0F, 0x00, 0x00}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	minVal, maxVal, err := minMaxSamples(pd, false)
	if err != nil {
		t.Fatalf("minMaxSamples() error = %v", err)
	}
	if minVal != -1 || maxVal != 0 {
		t.Fatalf("min/max = %v/%v, want -1/0", minVal, maxVal)
	}
}

func TestBuildFragmentSequenceRebuildsOffsetsForEmittedLayout(t *testing.T) {
	elem, err := buildFragmentSequence(
		[][]byte{
			{0xAA, 0xBB, 0xCC},       // len=3 -> padded item payload 4
			{0x11, 0x22, 0x33, 0x44}, // len=4
		},
		[]uint32{0, 999},
		8,
	)
	if err != nil {
		t.Fatalf("buildFragmentSequence() error = %v", err)
	}

	obf, ok := elem.(*element.OtherByteFragment)
	if !ok {
		t.Fatalf("buildFragmentSequence() type = %T, want *element.OtherByteFragment", elem)
	}
	if got := obf.OffsetTable(); len(got) != 2 || got[0] != 0 || got[1] != 12 {
		t.Fatalf("OffsetTable = %v, want [0 12]", got)
	}
}

func TestFromDatasetStripsTrailingPaddingWithBOT(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 12})
	obf.AddFragment(buffer.NewMemory([]byte{0xFF, 0xD9, 0x00}))
	obf.AddFragment(buffer.NewMemory([]byte("B")))
	_ = ds.Add(obf)

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}

	frame, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame(0) error = %v", err)
	}
	if !bytes.Equal(frame, []byte{0xFF, 0xD9}) {
		t.Fatalf("frame 0 = %v, want JPEG EOI without padding", frame)
	}
}

func TestDataToElementUsesOBForEncapsulated16Bit(t *testing.T) {
	pd, err := New(&Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PhotometricInterpretation: pixel.Monochrome2,
		Encapsulated:              true,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{0xAA, 0xBB}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	elem, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if _, ok := elem.(*element.OtherByteFragment); !ok {
		t.Fatalf("ToElement() = %T, want *element.OtherByteFragment for encapsulated data", elem)
	}
}

func TestInfo_UncompressedFrameSize(t *testing.T) {
	tests := []struct {
		name         string
		info         *Info
		expectedSize int
	}{
		{
			name: "8-bit grayscale 512x512",
			info: &Info{
				Width:                     512,
				Height:                    512,
				BitsAllocated:             8,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectedSize: 512 * 512, // 262,144 bytes
		},
		{
			name: "16-bit grayscale 512x512",
			info: &Info{
				Width:                     512,
				Height:                    512,
				BitsAllocated:             16,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectedSize: 512 * 512 * 2, // 524,288 bytes
		},
		{
			name: "8-bit RGB 256x256",
			info: &Info{
				Width:                     256,
				Height:                    256,
				BitsAllocated:             8,
				SamplesPerPixel:           3,
				PhotometricInterpretation: pixel.RGBPhotometric,
			},
			expectedSize: 256 * 256 * 3, // 196,608 bytes
		},
		{
			name: "1-bit image 100x100",
			info: &Info{
				Width:                     100,
				Height:                    100,
				BitsAllocated:             1,
				SamplesPerPixel:           1,
				PhotometricInterpretation: pixel.Monochrome2,
			},
			expectedSize: (100*100-1)/8 + 1, // 1,250 bytes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := tt.info.UncompressedFrameSize()
			if size != tt.expectedSize {
				t.Errorf("UncompressedFrameSize() = %d, want %d", size, tt.expectedSize)
			}
		})
	}
}

func TestNew(t *testing.T) {
	info := &Info{
		Width:                     256,
		Height:                    256,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if pd.Info != info {
		t.Error("PixelData info not set correctly")
	}

	if pd.FrameCount() != 0 {
		t.Errorf("FrameCount() = %d, want 0", pd.FrameCount())
	}
}

func TestData_AddFrame(t *testing.T) {
	info := &Info{
		Width:                     10,
		Height:                    10,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Create test frame data
	frameData := make([]byte, 100)
	for i := range frameData {
		frameData[i] = byte(i % 256)
	}

	// Add frame
	err = pd.AddFrame(context.Background(), frameData)
	if err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	if pd.FrameCount() != 1 {
		t.Errorf("FrameCount() = %d, want 1", pd.FrameCount())
	}

	// Get frame
	retrievedFrame, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}

	if !bytes.Equal(frameData, retrievedFrame) {
		t.Error("Retrieved frame data does not match original")
	}
}

func TestData_MultiFrame(t *testing.T) {
	info := &Info{
		Width:                     8,
		Height:                    8,
		NumberOfFrames:            3,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add 3 frames
	for i := 0; i < 3; i++ {
		frameData := make([]byte, 64)
		for j := range frameData {
			frameData[j] = byte((i*64 + j) % 256)
		}
		err = pd.AddFrame(context.Background(), frameData)
		if err != nil {
			t.Fatalf("AddFrame(%d) error = %v", i, err)
		}
	}

	if pd.FrameCount() != 3 {
		t.Errorf("FrameCount() = %d, want 3", pd.FrameCount())
	}

	// Verify each frame
	for i := 0; i < 3; i++ {
		frame, err := pd.Frame(context.Background(), i)
		if err != nil {
			t.Fatalf("Frame(%d) error = %v", i, err)
		}
		if len(frame) != 64 {
			t.Errorf("Frame %d size = %d, want 64", i, len(frame))
		}
	}
}

func TestData_EncodeDecodeNative(t *testing.T) {
	// Create 16-bit pixel data
	info := &Info{
		Width:                     8,
		Height:                    8,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add frame
	frameData := make([]byte, 128) // 8x8 x 2 bytes
	for i := 0; i < len(frameData); i += 2 {
		frameData[i] = byte(i & 0xFF)
		frameData[i+1] = byte((i >> 8) & 0xFF)
	}
	err = pd.AddFrame(context.Background(), frameData)
	if err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	// Encode/Decode with Native codec (no compression)
	nativeCodec := codec.NewExplicitVRLittleEndianCodec()

	encoded, err := pd.Encode(context.Background(), nativeCodec, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	decoded, err := encoded.Decode(context.Background(), nativeCodec, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Verify data matches
	originalData := pd.AllFrames()
	decodedData := decoded.AllFrames()

	if !bytes.Equal(originalData, decodedData) {
		t.Error("Decoded data does not match original")
	}
}

func TestDataDecodeSelectsNativeWordVRFor16BitOutput(t *testing.T) {
	source, err := New(&Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                12,
		HighBit:                   11,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		VRCode:                    "OB",
		Encapsulated:              true,
		TransferSyntaxUID:         transfer.JPEG2000Lossless.UID().UID(),
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := source.AddFrame(context.Background(), []byte{0x34, 0x12}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	decoded, err := source.Decode(context.Background(), codec.NewExplicitVRLittleEndianCodec(), nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if decoded.Info.VRCode != "OW" {
		t.Fatalf("decoded VRCode = %q, want OW", decoded.Info.VRCode)
	}
	elem, err := decoded.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if _, ok := elem.(*element.OtherWord); !ok {
		t.Fatalf("decoded ToElement() = %T, want *element.OtherWord", elem)
	}
}

func TestData_EnsureInterleaved(t *testing.T) {
	info := &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.PlanarPlanar,
		PhotometricInterpretation: pixel.RGBPhotometric,
	}

	// Planar: R-plane [1 4], G-plane [2 5], B-plane [3 6]
	data := []byte{1, 4, 2, 5, 3, 6}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.EnsureInterleaved(); err != nil {
		t.Fatalf("EnsureInterleaved() error = %v", err)
	}

	if pd.Info.PlanarConfiguration != pixel.InterleavedPlanar {
		t.Fatalf("PlanarConfiguration not updated, got %d", pd.Info.PlanarConfiguration)
	}

	got := pd.AllFrames()
	expected := []byte{1, 2, 3, 4, 5, 6}
	if !bytes.Equal(got, expected) {
		t.Fatalf("interleaved data mismatch, got %v want %v", got, expected)
	}
}

func TestData_ConvertMonochrome1ToMonochrome2(t *testing.T) {
	info := &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome1,
	}

	data := []byte{0x10, 0x20}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertMonochrome1ToMonochrome2(); err != nil {
		t.Fatalf("ConvertMonochrome1ToMonochrome2() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != pixel.Monochrome2 {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.AllFrames()
	expected := []byte{0xEF, 0xDF} // 0xFF-0x10, 0xFF-0x20
	if !bytes.Equal(got, expected) {
		t.Fatalf("mono inversion mismatch, got %v want %v", got, expected)
	}
}

func TestData_ToElement_EncapsulatedBOT(t *testing.T) {
	info := &Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            2,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		Encapsulated:              true,
		VRCode:                    "OB",
	}

	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	pd.frames = [][]byte{{0xAA}, {0xBB}}

	elem, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}

	obf, ok := elem.(*element.OtherByteFragment)
	if !ok {
		t.Fatalf("expected OtherByteFragment, got %T", elem)
	}

	if obf.FragmentCount() != 2 {
		t.Fatalf("expected 2 fragments, got %d", obf.FragmentCount())
	}

	if len(obf.OffsetTable()) != 2 || obf.OffsetTable()[0] != 0 || obf.OffsetTable()[1] != 10 {
		t.Fatalf("unexpected BOT: %v", obf.OffsetTable())
	}
}

func TestData_MinMaxIgnorePadding(t *testing.T) {
	padding := int32(0)
	info := &Info{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{0, 5, 10}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	minVal, maxVal, err := pd.MinMax(true)
	if err != nil {
		t.Fatalf("MinMax(ignorePadding=true) error = %v", err)
	}
	if minVal != 5 || maxVal != 10 {
		t.Fatalf("expected min=5 max=10, got min=%v max=%v", minVal, maxVal)
	}

	minAll, maxAll, err := pd.MinMax(false)
	if err != nil {
		t.Fatalf("MinMax(ignorePadding=false) error = %v", err)
	}
	if minAll != 0 || maxAll != 10 {
		t.Fatalf("expected min=0 max=10, got min=%v max=%v", minAll, maxAll)
	}
}

func TestData_ConvertYBRFullToRGB(t *testing.T) {
	info := &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.YbrFull,
	}

	// Two pixels: (Y,Cb,Cr) = (16,128,128) and (50,128,128) -> RGB roughly (16,16,16) and (50,50,50)
	data := []byte{16, 128, 128, 50, 128, 128}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != pixel.RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.AllFrames()
	// allow small deviation; here we check exact values for the simple case
	expected := []byte{16, 16, 16, 50, 50, 50}
	if !bytes.Equal(got, expected) {
		t.Fatalf("YBR_FULL->RGB mismatch, got %v want %v", got, expected)
	}
}

func TestData_ConvertYBRFull422ToRGB(t *testing.T) {
	info := &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.YbrFull422,
	}

	// Two pixels packed: Y1=16, Y2=50, Cb=128, Cr=128
	data := []byte{16, 50, 128, 128}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != pixel.RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.AllFrames()
	expected := []byte{16, 16, 16, 50, 50, 50}
	if !bytes.Equal(got, expected) {
		t.Fatalf("YBR_FULL_422->RGB mismatch, got %v want %v", got, expected)
	}
}

func TestFromDataset_PaletteToRGB(t *testing.T) {
	// Palette with 2 entries: index 0 -> black, 1 -> red(255,0,0)
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
		// descriptors: number of entries=2, first=0, bits=8
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		// data: R=[0,255], G=[0,0], B=[0,0]
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{0, 255}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
		// pixel data: indices [0,1]
		element.NewOtherByte(tag.PixelData, []byte{0, 1}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != pixel.RGBPhotometric {
		t.Fatalf("expected photometric RGB, got %v", pd.Info.PhotometricInterpretation)
	}
	if pd.Info.SamplesPerPixel != 3 {
		t.Fatalf("expected SPP=3, got %d", pd.Info.SamplesPerPixel)
	}

	data := pd.AllFrames()
	expected := []byte{0, 0, 0, 255, 0, 0}
	if !bytes.Equal(data, expected) {
		t.Fatalf("palette -> RGB data mismatch, got %v want %v", data, expected)
	}
}

func TestFromDatasetPaletteAppliesFirstMappedValue(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{4}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{10, 20}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{30, 40}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{50, 60}),
		element.NewOtherByte(tag.PixelData, []byte{9, 10, 11, 12}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	want := []byte{
		10, 30, 50,
		10, 30, 50,
		20, 40, 60,
		20, 40, 60,
	}
	if got := pd.AllFrames(); !bytes.Equal(got, want) {
		t.Fatalf("palette RGB data = %v, want %v", got, want)
	}
}

func TestFromDataset_PaletteSegmentedToRGB(t *testing.T) {
	// Standard discrete segment: opcode 0, length 2, values 0 and 255.
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewOtherWord(tag.SegmentedRedPaletteColorLookupTableData, words(binary.LittleEndian, 0, 2, 0, 255)),
		element.NewOtherWord(tag.SegmentedGreenPaletteColorLookupTableData, words(binary.LittleEndian, 0, 2, 0, 0)),
		element.NewOtherWord(tag.SegmentedBluePaletteColorLookupTableData, words(binary.LittleEndian, 0, 2, 0, 0)),
		element.NewOtherByte(tag.PixelData, []byte{0, 1}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	if pd.Info.PhotometricInterpretation != pixel.RGBPhotometric {
		t.Fatalf("expected photometric RGB, got %v", pd.Info.PhotometricInterpretation)
	}
	data := pd.AllFrames()
	expected := []byte{0, 0, 0, 255, 0, 0}
	if !bytes.Equal(data, expected) {
		t.Fatalf("segmented palette -> RGB mismatch, got %v want %v", data, expected)
	}
}

func TestPaletteColorsRejectsDirectAndSegmentedDataForSameChannel(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{0, 255}),
		element.NewOtherWord(tag.SegmentedRedPaletteColorLookupTableData, words(binary.LittleEndian, 0, 2, 0, 255)),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	_, _, err = PaletteColors(ds)
	if err == nil || !strings.Contains(err.Error(), "both direct and segmented") {
		t.Fatalf("PaletteColors() error = %v, want direct/segmented conflict", err)
	}
}

func TestFromDatasetPaletteAlphaToRGBA(t *testing.T) {
	for _, tt := range []struct {
		name      string
		alphaData element.Element
	}{
		{
			name:      "direct alpha",
			alphaData: element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, []byte{70, 80}),
		},
		{
			name: "segmented alpha",
			alphaData: element.NewOtherWord(tag.SegmentedAlphaPaletteColorLookupTableData,
				words(binary.LittleEndian, 0, 2, 70, 80)),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds, err := dataset.NewWithElements([]element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{2}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
				element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{10, 20}),
				element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{30, 40}),
				element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{50, 60}),
				tt.alphaData,
				element.NewOtherByte(tag.PixelData, []byte{0, 1}),
			})
			if err != nil {
				t.Fatalf("NewWithElements() error = %v", err)
			}

			pd, err := FromDataset(ds)
			if err != nil {
				t.Fatalf("FromDataset() error = %v", err)
			}
			if got, want := pd.Info.SamplesPerPixel, uint16(4); got != want {
				t.Fatalf("SamplesPerPixel = %d, want %d", got, want)
			}
			if got, want := pd.AllFrames(), []byte{10, 30, 50, 70, 20, 40, 60, 80}; !bytes.Equal(got, want) {
				t.Fatalf("palette RGBA data = %v, want %v", got, want)
			}
		})
	}
}

func TestBuildPaletteLUTValidatesAlpha(t *testing.T) {
	baseElements := func() []element.Element {
		return []element.Element{
			element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
			element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
			element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
			element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{10, 20}),
			element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{30, 40}),
			element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{50, 60}),
		}
	}
	for _, tt := range []struct {
		name  string
		extra []element.Element
	}{
		{
			name:  "data without descriptor",
			extra: []element.Element{element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, []byte{70, 80})},
		},
		{
			name:  "descriptor without data",
			extra: []element.Element{element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 0, 8})},
		},
		{
			name: "mismatched first mapped value",
			extra: []element.Element{
				element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 1, 8}),
				element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, []byte{70, 80}),
			},
		},
		{
			name: "alpha bits must be eight",
			extra: []element.Element{
				element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 0, 16}),
				element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, words(binary.LittleEndian, 70, 80)),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			elements := append(baseElements(), tt.extra...)
			ds, err := dataset.NewWithElements(elements)
			if err != nil {
				t.Fatalf("NewWithElements() error = %v", err)
			}
			if _, err := buildPaletteLUT(ds); err == nil {
				t.Fatal("buildPaletteLUT() accepted malformed alpha palette")
			}
		})
	}
}

func TestFromDataset_BigEndianPaletteToRGB(t *testing.T) {
	tests := []struct {
		name  string
		red   element.Element
		green element.Element
		blue  element.Element
	}{
		{
			name:  "standard",
			red:   element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{0, 0, 0x80, 0}),
			green: element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0, 0, 0}),
			blue:  element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0, 0, 0}),
		},
		{
			name:  "segmented",
			red:   element.NewOtherWord(tag.SegmentedRedPaletteColorLookupTableData, words(binary.BigEndian, 0, 2, 0, 0x8000)),
			green: element.NewOtherWord(tag.SegmentedGreenPaletteColorLookupTableData, words(binary.BigEndian, 0, 2, 0, 0)),
			blue:  element.NewOtherWord(tag.SegmentedBluePaletteColorLookupTableData, words(binary.BigEndian, 0, 2, 0, 0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
			element.SetByteOrder(tt.red, binary.BigEndian)
			element.SetByteOrder(tt.green, binary.BigEndian)
			element.SetByteOrder(tt.blue, binary.BigEndian)
			elements := []element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{2}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
				element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 16}),
				element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 16}),
				element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 16}),
				tt.red,
				tt.green,
				tt.blue,
				element.NewOtherByte(tag.PixelData, []byte{0, 1}),
			}
			for _, elem := range elements {
				if err := ds.Add(elem); err != nil {
					t.Fatalf("add %s: %v", elem.Tag(), err)
				}
			}

			pd, err := FromDataset(ds)
			if err != nil {
				t.Fatalf("FromDataset() error = %v", err)
			}
			if got, want := pd.AllFrames(), []byte{0, 0, 0, 128, 0, 0}; !bytes.Equal(got, want) {
				t.Fatalf("palette RGB data = %v, want %v", got, want)
			}
		})
	}
}

func TestFromDatasetReadsEightBitPaletteEntriesFromWordData(t *testing.T) {
	tests := []struct {
		name   string
		syntax *transfer.Syntax
		order  binary.ByteOrder
	}{
		{name: "little endian", order: binary.LittleEndian},
		{name: "big endian", syntax: transfer.ExplicitVRBigEndian, order: binary.BigEndian},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			if tt.syntax != nil {
				ds.SetInternalTransferSyntax(tt.syntax)
			}
			red := element.NewOtherWord(tag.RedPaletteColorLookupTableData, words(tt.order, 17, 34))
			green := element.NewOtherWord(tag.GreenPaletteColorLookupTableData, words(tt.order, 0, 0))
			blue := element.NewOtherWord(tag.BluePaletteColorLookupTableData, words(tt.order, 0, 0))
			element.SetByteOrder(red, tt.order)
			element.SetByteOrder(green, tt.order)
			element.SetByteOrder(blue, tt.order)
			for _, elem := range []element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{2}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
				element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
				red,
				green,
				blue,
				element.NewOtherByte(tag.PixelData, []byte{0, 1}),
			} {
				if err := ds.Add(elem); err != nil {
					t.Fatalf("add %s: %v", elem.Tag(), err)
				}
			}

			pd, err := FromDataset(ds)
			if err != nil {
				t.Fatalf("FromDataset() error = %v", err)
			}
			if got, want := pd.AllFrames(), []byte{17, 0, 0, 34, 0, 0}; !bytes.Equal(got, want) {
				t.Fatalf("palette RGB data = %v, want %v", got, want)
			}
		})
	}
}

func TestExpandSegmentedLUTStandardSegments(t *testing.T) {
	tests := []struct {
		name string
		raw  []byte
		want []uint16
	}{
		{name: "discrete", raw: words(binary.LittleEndian, 0, 2, 0, 255), want: []uint16{0, 255}},
		{name: "linear", raw: words(binary.LittleEndian, 0, 1, 0, 1, 3, 300), want: []uint16{0, 100, 200, 300}},
		{name: "indirect", raw: words(binary.LittleEndian, 0, 2, 10, 20, 2, 1, 0, 0), want: []uint16{10, 20, 10, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandSegmentedLUT(tt.raw, len(tt.want), binary.LittleEndian)
			if err != nil {
				t.Fatalf("expandSegmentedLUT() error = %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Fatalf("expandSegmentedLUT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpandSegmentedLUTBigEndianIndirectOffsetUsesLowWordFirst(t *testing.T) {
	raw := words(binary.BigEndian,
		0, 1, 5,
		0, 1, 10,
		2, 1, 6, 0,
	)

	got, err := expandSegmentedLUT(raw, 3, binary.BigEndian)
	if err != nil {
		t.Fatalf("expandSegmentedLUT() error = %v", err)
	}
	if want := []uint16{5, 10, 10}; !slices.Equal(got, want) {
		t.Fatalf("expandSegmentedLUT() = %v, want %v", got, want)
	}
}

func TestExpandSegmentedLUTRejectsIndirectToIndirectReference(t *testing.T) {
	raw := words(binary.LittleEndian,
		0, 1, 5,
		2, 1, 0, 0,
		2, 1, 6, 0,
	)

	if _, err := expandSegmentedLUT(raw, 3, binary.LittleEndian); err == nil {
		t.Fatal("expandSegmentedLUT() accepted an indirect segment that references another indirect segment")
	}
}

func TestExpandSegmentedLUTRejectsIndirectFirstSegment(t *testing.T) {
	raw := words(binary.LittleEndian,
		2, 1, 8, 0,
		0, 1, 5,
	)

	if _, err := expandSegmentedLUT(raw, 2, binary.LittleEndian); err == nil {
		t.Fatal("expandSegmentedLUT() accepted an indirect first segment")
	}
}

func TestBuildPaletteLUTRejectsNonStandardBitDepth(t *testing.T) {
	ds := dataset.New()
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 12}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 12}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 12}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, words(binary.LittleEndian, 0, 4095)),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, words(binary.LittleEndian, 0, 4095)),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, words(binary.LittleEndian, 0, 4095)),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	if _, err := buildPaletteLUT(ds); err == nil {
		t.Fatal("buildPaletteLUT() accepted a Palette LUT bit depth other than 8 or 16")
	}
}

func TestBuildPaletteLUTDoesNotFallBackFromMalformedSequence(t *testing.T) {
	ds := dataset.New()
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{1, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{1, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{1, 0, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{10}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{20}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{30}),
	} {
		_ = ds.Add(elem)
	}
	malformedItem := dataset.New()
	_ = malformedItem.Add(element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{1, 0, 8}))
	_ = ds.Add(dataset.NewSequenceWithItems(
		tag.EnhancedPaletteColorLookupTableSequence,
		[]*dataset.Dataset{malformedItem},
	))

	if _, err := buildPaletteLUT(ds); err == nil {
		t.Fatal("buildPaletteLUT() silently fell back from a malformed Enhanced Palette LUT Sequence")
	}
}

func TestExpandSegmentedLUTRejectsMalformedData(t *testing.T) {
	tests := []struct {
		name string
		raw  []byte
		size int
	}{
		{name: "empty", size: 1},
		{name: "odd byte length", raw: []byte{0}, size: 1},
		{name: "truncated discrete", raw: words(binary.LittleEndian, 0, 2, 10), size: 2},
		{name: "linear without prior value", raw: words(binary.LittleEndian, 1, 2, 10), size: 2},
		{name: "truncated linear", raw: words(binary.LittleEndian, 0, 1, 10, 1, 2), size: 3},
		{name: "indirect offset out of range", raw: words(binary.LittleEndian, 0, 1, 10, 2, 1, 100, 0), size: 2},
		{name: "indirect cycle", raw: words(binary.LittleEndian, 2, 1, 0, 0), size: 1},
		{name: "unknown opcode", raw: words(binary.LittleEndian, 3, 0), size: 1},
		{name: "output count mismatch", raw: words(binary.LittleEndian, 0, 1, 10), size: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Errorf("expandSegmentedLUT() panicked: %v", recovered)
				}
			}()
			if _, err := expandSegmentedLUT(tt.raw, tt.size, binary.LittleEndian); err == nil {
				t.Fatal("expandSegmentedLUT() accepted malformed data")
			}
		})
	}
}

func words(order binary.ByteOrder, values ...uint16) []byte {
	raw := make([]byte, len(values)*2)
	for index, value := range values {
		order.PutUint16(raw[index*2:], value)
	}
	return raw
}

func TestData_ConvertYBRPartial422ToRGB(t *testing.T) {
	info := &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.YbrPartial422,
	}

	// Two pixels packed: Y1=16, Y2=50, Cb=128, Cr=128 (limited range)
	data := []byte{16, 50, 128, 128}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != pixel.RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.AllFrames()
	// Expected approximate values: first pixel near 0, second near 39
	if got[0] > 5 || got[1] > 5 || got[2] > 5 {
		t.Fatalf("first pixel not near black: %v", got[:3])
	}
	if got[3] < 35 || got[3] > 45 || got[4] < 35 || got[4] > 45 || got[5] < 35 || got[5] > 45 {
		t.Fatalf("second pixel not near gray: %v", got[3:6])
	}
}

func TestData_ConvertYBRICTToRGB(t *testing.T) {
	info := &Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.SignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.YbrIct,
	}

	// Y=50, Cb=0, Cr=0 should map to roughly RGB(50,50,50)
	data := []byte{50, 0, 0}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	got := pd.AllFrames()
	if got[0] < 48 || got[0] > 52 || got[1] < 48 || got[1] > 52 || got[2] < 48 || got[2] > 52 {
		t.Fatalf("YBR_ICT->RGB not near expected 50s: %v", got)
	}
}

func TestData_ConvertYBRRCTToRGB(t *testing.T) {
	info := &Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.SignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.YbrRct,
	}

	// From R=100,G=110,B=120 -> Y=floor((R+2G+B)/4)=110, Cb=B-G=10, Cr=R-G=-10
	y := int16(110)
	cb := int16(10)
	cr := int16(-10)
	data := []byte{
		byte(y & 0xFF), byte((y >> 8) & 0xFF),
		byte(cb & 0xFF), byte((cb >> 8) & 0xFF),
		byte(cr & 0xFF), byte((cr >> 8) & 0xFF),
	}

	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	got := pd.AllFrames()
	if len(got) < 6 {
		t.Fatalf("expected 6 bytes, got %d", len(got))
	}
	r := int16(got[0]) | int16(got[1])<<8
	g := int16(got[2]) | int16(got[3])<<8
	b := int16(got[4]) | int16(got[5])<<8
	if r != 100 || g != 110 || b != 120 {
		t.Fatalf("YBR_RCT->RGB mismatch: got R=%d G=%d B=%d", r, g, b)
	}
}

func TestData_VOILUTSequence(t *testing.T) {
	// Descriptor bit depth defines the output range; entries are not stretched by their extrema.
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	// VOI LUT Sequence with one item
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
	_ = addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{0, 100, 200}))
	voiSeq := dataset.NewSequence(tag.VOILUTSequence)
	voiSeq.AddItem(lutItem)
	_ = addLegacyTestElement(ds, voiSeq)

	// PixelData: 0,1,2
	_ = ds.Add(element.NewOtherByte(tag.PixelData, []byte{0, 1, 2}))

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}

	out, err := pd.WindowOrLUTTo8bit(ds, 0, 0, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	got := out[0]
	expected := []byte{0, 100, 200}
	if !bytes.Equal(got, expected) {
		t.Fatalf("VOI LUT mapping mismatch, got %v want %v", got, expected)
	}
}

func TestApplyVOILUTReadsEightBitEntriesFromWordData(t *testing.T) {
	tests := []struct {
		name   string
		syntax *transfer.Syntax
		data   []byte
	}{
		{name: "little endian", data: []byte{10, 0, 20, 0, 30, 0}},
		{name: "big endian", syntax: transfer.ExplicitVRBigEndian, data: []byte{0, 10, 0, 20, 0, 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			if tt.syntax != nil {
				ds.SetInternalTransferSyntax(tt.syntax)
			}
			item := dataset.New()
			_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
			lutData := element.NewOtherWord(tag.LUTData, tt.data)
			if tt.syntax == transfer.ExplicitVRBigEndian {
				element.SetByteOrder(lutData, binary.BigEndian)
			}
			_ = item.Add(lutData)
			_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
			pd, err := NewFromBytes(&Info{
				Width: 3, Height: 1, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
				PhotometricInterpretation: pixel.Monochrome2,
			}, []byte{0, 1, 2})
			if err != nil {
				t.Fatalf("NewFromBytes() error = %v", err)
			}

			got, err := applyVOILUT(pd, ds, 0, 0, false)
			if err != nil {
				t.Fatalf("applyVOILUT() error = %v", err)
			}
			if want := []byte{10, 20, 30}; !bytes.Equal(got[0], want) {
				t.Fatalf("applyVOILUT() = %v, want %v", got[0], want)
			}
		})
	}
}

func TestApplyVOILUTKeepsUnsignedFirstMappedValue(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 40000, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10, 20}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
	pd, err := NewFromBytes(&Info{
		Width: 2, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{0x40, 0x9c, 0x41, 0x9c})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	got, err := applyVOILUT(pd, ds, 0, 0, false)
	if err != nil {
		t.Fatalf("applyVOILUT() error = %v", err)
	}
	if want := []byte{10, 20}; !bytes.Equal(got[0], want) {
		t.Fatalf("applyVOILUT() = %v, want %v", got[0], want)
	}
}

func TestApplyVOILUTRejectsShortData(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10, 20}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
	pd, err := NewFromBytes(&Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{0})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if _, err := applyVOILUT(pd, ds, 0, 0, false); err == nil {
		t.Fatal("applyVOILUT() accepted LUT Data shorter than the descriptor")
	}
}

func TestWindowOrLUTTo8bitSurfacesMalformedVOILUT(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
	pd, err := NewFromBytes(&Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{0})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if _, err := pd.WindowOrLUTTo8bit(ds, 0, 1, false); err == nil {
		t.Fatal("WindowOrLUTTo8bit() silently fell back from malformed VOI LUT")
	}
}

func TestData_VOILUTSequenceNormalizes16BitEntries(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 32768, 65535} {
		binary.LittleEndian.PutUint16(lutBytes[index*2:], value)
	}
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	_ = lutItem.Add(element.NewOtherWord(tag.LUTData, lutBytes))
	voiSeq := dataset.NewSequence(tag.VOILUTSequence)
	voiSeq.AddItem(lutItem)
	_ = ds.Add(voiSeq)
	_ = ds.Add(element.NewOtherByte(tag.PixelData, []byte{0, 1, 2}))

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	out, err := pd.WindowOrLUTTo8bit(ds, 0, 0, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	if got, want := out[0], []byte{0, 128, 255}; !bytes.Equal(got, want) {
		t.Fatalf("VOI LUT mapping mismatch, got %v want %v", got, want)
	}
}

func TestData_VOILUTSequenceReadsBigEndianData(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
	}
	for _, elem := range elements {
		_ = ds.Add(elem)
	}
	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 32768, 65535} {
		binary.BigEndian.PutUint16(lutBytes[index*2:], value)
	}
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	lutData := element.NewOtherWord(tag.LUTData, lutBytes)
	element.SetByteOrder(lutData, binary.BigEndian)
	_ = lutItem.Add(lutData)
	voiSeq := dataset.NewSequence(tag.VOILUTSequence)
	voiSeq.AddItem(lutItem)
	_ = ds.Add(voiSeq)
	_ = ds.Add(element.NewOtherByte(tag.PixelData, []byte{0, 1, 2}))

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	out, err := pd.WindowOrLUTTo8bit(ds, 0, 0, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	if got, want := out[0], []byte{0, 128, 255}; !bytes.Equal(got, want) {
		t.Fatalf("VOI LUT mapping mismatch, got %v want %v", got, want)
	}
}

func TestApplyVOILUTRejectsInvalidBitsPerEntry(t *testing.T) {
	for _, bitsPerEntry := range []uint16{9, 12, 15, 17, 32} {
		t.Run(fmt.Sprintf("bits_%d", bitsPerEntry), func(t *testing.T) {
			ds := dataset.New()
			lutItem := dataset.New()
			_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, bitsPerEntry}))
			_ = lutItem.Add(element.NewOtherWord(tag.LUTData, []byte{0, 0}))
			_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))
			pd, err := NewFromBytes(&Info{
				Width: 1, Height: 1, NumberOfFrames: 1,
				BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
				PhotometricInterpretation: pixel.Monochrome2,
			}, []byte{0})
			if err != nil {
				t.Fatalf("NewFromBytes() error = %v", err)
			}

			if _, err := applyVOILUT(pd, ds, 0, 0, false); err == nil {
				t.Fatalf("applyVOILUT() accepted bitsPerEntry=%d", bitsPerEntry)
			}
		})
	}
}

func TestPaletteLUTStandardRejectsOBData(t *testing.T) {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	for _, descriptorTag := range []*tag.Tag{
		tag.RedPaletteColorLookupTableDescriptor,
		tag.GreenPaletteColorLookupTableDescriptor,
		tag.BluePaletteColorLookupTableDescriptor,
	} {
		if err := ds.Add(element.NewUnsignedShort(descriptorTag, []uint16{1, 0, 8})); err != nil {
			t.Fatal(err)
		}
	}
	for _, dataTag := range []*tag.Tag{
		tag.RedPaletteColorLookupTableData,
		tag.GreenPaletteColorLookupTableData,
		tag.BluePaletteColorLookupTableData,
	} {
		if err := ds.Add(element.NewOtherByte(dataTag, []byte{1})); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := buildPaletteLUTWithVRMode(ds, LUTStandard); err == nil {
		t.Fatal("buildPaletteLUTWithVRMode() accepted non-standard OB Palette LUT data")
	}
}

func TestPaletteLUTStandardRequiresOWForSegmentedData(t *testing.T) {
	ds := dataset.New()
	for _, descriptorTag := range []*tag.Tag{
		tag.RedPaletteColorLookupTableDescriptor,
		tag.GreenPaletteColorLookupTableDescriptor,
		tag.BluePaletteColorLookupTableDescriptor,
	} {
		if err := ds.Add(element.NewUnsignedShort(descriptorTag, []uint16{2, 0, 8})); err != nil {
			t.Fatal(err)
		}
	}
	segment := words(binary.LittleEndian, 0, 2, 0, 255)
	for _, dataTag := range []*tag.Tag{
		tag.SegmentedRedPaletteColorLookupTableData,
		tag.SegmentedGreenPaletteColorLookupTableData,
		tag.SegmentedBluePaletteColorLookupTableData,
	} {
		if err := ds.Add(element.NewOtherWord(dataTag, segment)); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := buildPaletteLUTWithVRMode(ds, LUTStandard); err != nil {
		t.Fatalf("buildPaletteLUTWithVRMode() rejected standard OW segmented data: %v", err)
	}

	ds.SetAutoValidate(false)
	if err := ds.AddOrUpdate(element.NewOtherByte(tag.SegmentedRedPaletteColorLookupTableData, segment)); err != nil {
		t.Fatal(err)
	}
	if _, err := buildPaletteLUTWithVRMode(ds, LUTStandard); err == nil {
		t.Fatal("buildPaletteLUTWithVRMode() accepted non-standard OB segmented data")
	}
}

func TestData_WindowTo8bit(t *testing.T) {
	info := &Info{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.SignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	// values: -1000, 0, 1000
	data := []byte{
		0x18, 0xFC, // -1000
		0x00, 0x00, // 0
		0xE8, 0x03, // 1000
	}

	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	frames, err := pd.WindowTo8bit(0, 2000, false)
	if err != nil {
		t.Fatalf("WindowTo8bit() error = %v", err)
	}
	if len(frames) != 1 {
		t.Fatalf("expected 1 frame, got %d", len(frames))
	}

	got := frames[0]
	if len(got) != 3 {
		t.Fatalf("expected 3 samples, got %d", len(got))
	}

	if got[0] != 0 {
		t.Fatalf("expected first sample 0, got %d", got[0])
	}
	// middle should be around mid-gray
	if got[1] < 125 || got[1] > 130 {
		t.Fatalf("expected mid sample ~127, got %d", got[1])
	}
	if got[2] != 255 {
		t.Fatalf("expected last sample 255, got %d", got[2])
	}
}

func TestData_WindowTo8bit_Padding(t *testing.T) {
	padding := int32(0)
	info := &Info{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{0, 10, 20} // first is padding
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	frames, err := pd.WindowTo8bit(10, 10, true)
	if err != nil {
		t.Fatalf("WindowTo8bit() error = %v", err)
	}
	out := frames[0]
	if out[0] != 0 {
		t.Fatalf("padding sample expected 0, got %d", out[0])
	}
	if out[1] < 140 || out[1] > 142 {
		t.Fatalf("first real sample expected ~141, got %d", out[1])
	}
	if out[2] != 255 {
		t.Fatalf("second real sample expected 255, got %d", out[2])
	}
}

func TestData_MaskPadding(t *testing.T) {
	padding := int32(5)
	info := &Info{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{5, 10, 5}
	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	frames, masks, err := pd.MaskPadding()
	if err != nil {
		t.Fatalf("MaskPadding() error = %v", err)
	}
	if len(frames) != 1 || len(masks) != 1 {
		t.Fatalf("expected 1 frame/mask, got %d/%d", len(frames), len(masks))
	}
	out := frames[0]
	mask := masks[0]

	expected := []byte{0, 10, 0}
	if !bytes.Equal(out, expected) {
		t.Fatalf("masked data mismatch, got %v want %v", out, expected)
	}
	if len(mask) != 3 || mask[0] != true || mask[1] != false || mask[2] != true {
		t.Fatalf("mask mismatch: %v", mask)
	}
}

func TestDataNormalizesReversedPixelPaddingRange(t *testing.T) {
	padding := int32(20)
	rangeLimit := int32(10)
	info := &Info{
		Width: 4, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7,
		SamplesPerPixel: 1, PixelRepresentation: pixel.UnsignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
		PixelPaddingValue:         &padding, PixelPaddingRangeLimit: &rangeLimit,
	}
	pd, err := NewFromBytes(info, []byte{5, 10, 15, 20})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	for _, value := range []int64{10, 15, 20} {
		if !pd.IsPaddingSample(value) {
			t.Errorf("IsPaddingSample(%d) = false, want true", value)
		}
	}
	if pd.IsPaddingSample(5) {
		t.Fatal("IsPaddingSample(5) = true, want false")
	}
	minValue, maxValue, err := pd.MinMax(true)
	if err != nil {
		t.Fatalf("MinMax(ignorePadding=true) error = %v", err)
	}
	if minValue != 5 || maxValue != 5 {
		t.Fatalf("MinMax(ignorePadding=true) = (%v, %v), want (5, 5)", minValue, maxValue)
	}
	masked, masks, err := pd.MaskPadding()
	if err != nil {
		t.Fatalf("MaskPadding() error = %v", err)
	}
	if got, want := masked[0], []byte{5, 0, 0, 0}; !bytes.Equal(got, want) {
		t.Fatalf("MaskPadding() data = %v, want %v", got, want)
	}
	if got, want := masks[0], []bool{false, true, true, true}; !slices.Equal(got, want) {
		t.Fatalf("MaskPadding() mask = %v, want %v", got, want)
	}
	windowed, err := pd.WindowTo8bit(5, 10, true)
	if err != nil {
		t.Fatalf("WindowTo8bit() error = %v", err)
	}
	if got, want := windowed[0][1:], []byte{0, 0, 0}; !bytes.Equal(got, want) {
		t.Fatalf("WindowTo8bit() padding = %v, want %v", got, want)
	}
}

func TestNewFromBytes(t *testing.T) {
	info := &Info{
		Width:                     10,
		Height:                    10,
		NumberOfFrames:            2,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}

	// Create 2 frames of data
	data := make([]byte, 200) // 2 frames x 100 bytes
	for i := range data {
		data[i] = byte(i % 256)
	}

	pd, err := NewFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	if pd.FrameCount() != 2 {
		t.Errorf("FrameCount() = %d, want 2", pd.FrameCount())
	}

	// Verify frame 0
	frame0, err := pd.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame(0) error = %v", err)
	}
	if !bytes.Equal(data[:100], frame0) {
		t.Error("Frame 0 data mismatch")
	}

	// Verify frame 1
	frame1, err := pd.Frame(context.Background(), 1)
	if err != nil {
		t.Fatalf("Frame(1) error = %v", err)
	}
	if !bytes.Equal(data[100:200], frame1) {
		t.Error("Frame 1 data mismatch")
	}
}

func TestNewFromBytesRejectsNilInfo(t *testing.T) {
	if _, err := NewFromBytes(nil, nil); err == nil {
		t.Fatal("NewFromBytes(nil, nil) error = nil, want an error")
	}
}

func TestEncodeAndDecodeRejectNilInfo(t *testing.T) {
	var pd *Data
	if _, err := pd.Encode(context.Background(), nil, nil); err == nil {
		t.Fatal("Encode() error = nil, want an error")
	}
	if _, err := pd.Decode(context.Background(), nil, nil); err == nil {
		t.Fatal("Decode() error = nil, want an error")
	}
}

func TestOneBitSampleAndOptimalWindow(t *testing.T) {
	info := &Info{
		Width: 8, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 1, BitsStored: 1, HighBit: 0,
		SamplesPerPixel: 1, PhotometricInterpretation: pixel.Monochrome2,
	}
	pd, err := NewFromBytes(info, []byte{0x81})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	for x, want := range []int64{1, 0, 0, 0, 0, 0, 0, 1} {
		got, err := pd.Sample(0, x, 0, 0)
		if err != nil {
			t.Fatalf("Sample(%d) error = %v", x, err)
		}
		if got != want {
			t.Errorf("Sample(%d) = %d, want %d", x, got, want)
		}
	}
	center, width := pd.CalculateOptimalWindow()
	if center != 0.5 || width != 1 {
		t.Fatalf("CalculateOptimalWindow() = (%v, %v), want (0.5, 1)", center, width)
	}
	frames, err := pd.WindowTo8bit(0.5, 1, false)
	if err != nil {
		t.Fatalf("WindowTo8bit() error = %v", err)
	}
	if got, want := frames[0], []byte{255, 0, 0, 0, 0, 0, 0, 255}; !bytes.Equal(got, want) {
		t.Fatalf("WindowTo8bit() = %v, want %v", got, want)
	}
	minValue, maxValue, err := pd.MinMax(false)
	if err != nil || minValue != 0 || maxValue != 1 {
		t.Fatalf("MinMax() = (%v, %v), %v; want (0, 1), nil", minValue, maxValue, err)
	}
}

func TestEncapsulatedAllFramesUsesPayloadSize(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7,
		SamplesPerPixel: 1, PhotometricInterpretation: pixel.Monochrome2,
		Encapsulated: true, VRCode: "OB",
	}
	pd, err := New(info)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if err := pd.AddFrame(context.Background(), []byte{1, 2, 3}); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	if got, want := pd.AllFrames(), []byte{1, 2, 3}; !bytes.Equal(got, want) {
		t.Fatalf("AllFrames() = %v, want %v", got, want)
	}
}

func TestNewFromBytesUsesWordVRForImplicitSyntax(t *testing.T) {
	pd, err := NewFromBytes(&Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		TransferSyntaxUID:         transfer.ImplicitVRLittleEndian.UID().UID(),
	}, []byte{0x7f})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	if pd.Info.VRCode != "OW" {
		t.Fatalf("VRCode = %q, want OW for Implicit VR Little Endian", pd.Info.VRCode)
	}
	elem, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if _, ok := elem.(*element.OtherWord); !ok {
		t.Fatalf("ToElement() = %T, want *element.OtherWord", elem)
	}
}

func TestDataToElementPreservesLegalNativeEightBitWordVR(t *testing.T) {
	pd, err := NewFromBytes(&Info{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
		VRCode:                    "OW",
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
	}, []byte{0x7f})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	elem, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if _, ok := elem.(*element.OtherWord); !ok {
		t.Fatalf("ToElement() = %T, want *element.OtherWord", elem)
	}
}

func TestWindowOrLUTTo8bitAppliesRescaleBeforeWindow(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15,
		SamplesPerPixel: 1, PixelRepresentation: pixel.UnsignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
	}
	pd, err := NewFromBytes(info, []byte{50, 0})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	ds := dataset.New()
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
		t.Fatalf("add RescaleIntercept: %v", err)
	}
	if err := ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add RescaleType: %v", err)
	}

	frames, err := pd.WindowOrLUTTo8bit(ds, 100, 100, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	if got := frames[0][0]; got < 127 || got > 129 {
		t.Fatalf("rescaled sample mapped to %d, want approximately 128", got)
	}
}

func TestModalityLUTRequiresModalityLUTType(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 8})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	if err := addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0, 255})); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	ds := dataset.New()
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add ModalityLUTSequence: %v", err)
	}

	_, err := ModalityLUT(ds, false)
	if err == nil || !strings.Contains(err.Error(), "Modality LUT Type") {
		t.Fatalf("ModalityLUT() error = %v, want missing Modality LUT Type", err)
	}
}

func TestFunctionalGroupModalityLUTOverridesTopLevelRescale(t *testing.T) {
	primary := dataset.New()
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add Modality LUT Type: %v", err)
	}
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 8})); err != nil {
		t.Fatalf("add LUT Descriptor: %v", err)
	}
	if err := addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{10, 20})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}
	if err := addLegacyTestElement(primary, dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add Modality LUT Sequence: %v", err)
	}

	fallback := dataset.New()
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{100}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}),
		element.NewString(tag.RescaleType, vr.LO, []string{"US"}),
	} {
		if err := fallback.Add(elem); err != nil {
			t.Fatalf("add top-level rescale element: %v", err)
		}
	}

	modality, err := modalityTransformForDatasets(primary, fallback, &Data{
		Info: &Info{PixelRepresentation: pixel.UnsignedPixels},
	})
	if err != nil {
		t.Fatalf("modalityTransformForDatasets() error = %v", err)
	}
	if modality == nil {
		t.Fatal("modalityTransformForDatasets() returned no Functional Group LUT")
	}
	if got := modality.Transform(1); got != 20 {
		t.Fatalf("Functional Group LUT Transform(1) = %v, want 20", got)
	}
}

func TestWindowOrLUTTo8bitAllowsMissingRescaleType(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15,
		SamplesPerPixel: 1, PixelRepresentation: pixel.UnsignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
	}
	pd, err := NewFromBytes(info, []byte{50, 0})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}

	frames, err := pd.WindowOrLUTTo8bit(ds, 100, 100, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	if len(frames) != 1 || len(frames[0]) != 1 {
		t.Fatalf("WindowOrLUTTo8bit() returned %d frames with lengths %v, want one 1-byte frame", len(frames), frameLengths(frames))
	}
}

func frameLengths(frames [][]byte) []int {
	lengths := make([]int, len(frames))
	for index, frame := range frames {
		lengths[index] = len(frame)
	}
	return lengths
}

func TestWindowOrLUTTo8bitAppliesPerFrameFunctionalGroupRescale(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 2,
		BitsAllocated: 16, BitsStored: 16, HighBit: 15,
		SamplesPerPixel: 1, PixelRepresentation: pixel.UnsignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
	}
	pd, err := NewFromBytes(info, []byte{50, 0, 50, 0})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}

	perFrameItems := make([]*dataset.Dataset, 2)
	for frame, slope := range []float64{1, 2} {
		transform := dataset.New()
		if err := transform.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{slope})); err != nil {
			t.Fatalf("frame %d add RescaleSlope: %v", frame, err)
		}
		if err := transform.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
			t.Fatalf("frame %d add RescaleIntercept: %v", frame, err)
		}
		if err := transform.Add(element.NewString(tag.RescaleType, vr.LO, []string{"HU"})); err != nil {
			t.Fatalf("frame %d add RescaleType: %v", frame, err)
		}
		frameItem := dataset.New()
		if err := frameItem.Add(dataset.NewSequenceWithItems(
			tag.PixelValueTransformationSequence,
			[]*dataset.Dataset{transform},
		)); err != nil {
			t.Fatalf("frame %d add PixelValueTransformationSequence: %v", frame, err)
		}
		perFrameItems[frame] = frameItem
	}
	ds := dataset.New()
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, perFrameItems)); err != nil {
		t.Fatalf("add PerFrameFunctionalGroupsSequence: %v", err)
	}

	frames, err := pd.WindowOrLUTTo8bit(ds, 100, 100, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	if got := frames[0][0]; got != 0 {
		t.Fatalf("frame 0 = %d, want 0", got)
	}
	if got := frames[1][0]; got < 127 || got > 129 {
		t.Fatalf("frame 1 = %d, want approximately 128", got)
	}
}

func TestWindowOrLUTTo8bitRejectsNonPositiveWindowWidth(t *testing.T) {
	info := &Info{
		Width: 1, Height: 1, NumberOfFrames: 1,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7,
		SamplesPerPixel: 1, PixelRepresentation: pixel.UnsignedPixels,
		PhotometricInterpretation: pixel.Monochrome2,
	}
	pd, err := NewFromBytes(info, []byte{1})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	if _, err := pd.WindowOrLUTTo8bit(nil, 0, 0, false); err == nil {
		t.Fatal("WindowOrLUTTo8bit() accepted a non-positive window width")
	}
}

func TestFromDatasetNormalizesBigEndianPixelDataOnce(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherWord(tag.PixelData, []byte{0x12, 0x34}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	got, err := pd.Sample(0, 0, 0, 0)
	if err != nil {
		t.Fatalf("Sample() error = %v", err)
	}
	if got != 0x1234 {
		t.Fatalf("Sample() = %#x, want 0x1234", got)
	}
}

func TestFromDatasetToleratesNativeOBForSixteenBitPixelDataAndNormalizesOnWrite(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	pd, err := FromDataset(ds)
	if err != nil {
		t.Fatalf("FromDataset() error = %v", err)
	}
	elem, err := pd.ToElement()
	if err != nil {
		t.Fatalf("ToElement() error = %v", err)
	}
	if _, ok := elem.(*element.OtherWord); !ok {
		t.Fatalf("ToElement() = %T, want *element.OtherWord", elem)
	}
}

// TestData_ToElement_VRSelection tests that ToElement correctly chooses
// OB vs OW based on BitsAllocated for both encapsulated and native formats.
func TestData_ToElement_VRSelection(t *testing.T) {
	tests := []struct {
		name           string
		bitsAllocated  uint16
		encapsulated   bool
		expectedVRType string // "OB" or "OW"
	}{
		{
			name:           "Encapsulated 8-bit should use OB",
			bitsAllocated:  8,
			encapsulated:   true,
			expectedVRType: "OB",
		},
		{
			name:           "Encapsulated 16-bit should use OB",
			bitsAllocated:  16,
			encapsulated:   true,
			expectedVRType: "OB",
		},
		{
			name:           "Native 8-bit should use OB",
			bitsAllocated:  8,
			encapsulated:   false,
			expectedVRType: "OB",
		},
		{
			name:           "Native 16-bit should use OW",
			bitsAllocated:  16,
			encapsulated:   false,
			expectedVRType: "OW",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &Info{
				Width:                     10,
				Height:                    10,
				NumberOfFrames:            1,
				BitsAllocated:             tt.bitsAllocated,
				BitsStored:                tt.bitsAllocated,
				HighBit:                   tt.bitsAllocated - 1,
				SamplesPerPixel:           1,
				PixelRepresentation:       pixel.UnsignedPixels,
				PlanarConfiguration:       pixel.InterleavedPlanar,
				PhotometricInterpretation: pixel.Monochrome2,
				Encapsulated:              tt.encapsulated,
			}

			pd, err := New(info)
			if err != nil {
				t.Fatalf("New error = %v", err)
			}

			// Add dummy frame data
			// Calculate correct frame size based on BitsAllocated
			bytesPerPixel := (tt.bitsAllocated + 7) / 8
			frameSize := int(10 * 10 * bytesPerPixel)
			frameData := make([]byte, frameSize)
			for i := range frameData {
				frameData[i] = byte(i % 256)
			}
			if err := pd.AddFrame(context.Background(), frameData); err != nil {
				t.Fatalf("AddFrame error = %v", err)
			}

			// Convert to element
			elem, err := pd.ToElement()
			if err != nil {
				t.Fatalf("ToElement error = %v", err)
			}

			// Check VR type
			var actualVR string
			switch elem.(type) {
			case *element.OtherByte:
				actualVR = "OB"
			case *element.OtherWord:
				actualVR = "OW"
			case *element.OtherByteFragment:
				actualVR = "OB"
			case *element.OtherWordFragment:
				actualVR = "OW"
			default:
				t.Fatalf("Unexpected element type: %T", elem)
			}

			if actualVR != tt.expectedVRType {
				t.Errorf("VR type = %s, want %s (BitsAllocated=%d, Encapsulated=%v)",
					actualVR, tt.expectedVRType, tt.bitsAllocated, tt.encapsulated)
			}
		})
	}
}

// TestFromDataset_TransferSyntax verifies that FromDataset correctly reads
// transfer syntax from the dataset
func TestFromDataset_TransferSyntax(t *testing.T) {
	tests := []struct {
		name                   string
		setupDataset           func() *dataset.Dataset
		expectedTransferSyntax string
	}{
		{
			name: "reads from InternalTransferSyntax",
			setupDataset: func() *dataset.Dataset {
				ts, err := transfer.Parse("1.2.840.10008.1.2.4.50") // JPEG Baseline
				if err != nil {
					t.Fatalf("failed to parse transfer syntax: %v", err)
				}
				ds := dataset.NewWithTransferSyntax(ts)

				// Add required image attributes
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Rows, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Columns, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

				// Add pixel data
				pixelData := make([]byte, 256*256)
				_ = ds.AddOrUpdate(element.NewOtherByte(tag.PixelData, pixelData))

				return ds
			},
			expectedTransferSyntax: "1.2.840.10008.1.2.4.50",
		},
		{
			name: "reads from TransferSyntaxUID tag",
			setupDataset: func() *dataset.Dataset {
				ds := dataset.New()

				// Add transfer syntax as a tag (e.g., from file meta information)
				_ = ds.AddOrUpdate(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{"1.2.840.10008.1.2.2"})) // Explicit VR Big Endian

				// Add required image attributes
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Rows, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Columns, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsStored, []uint16{12}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.HighBit, []uint16{11}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

				// Add pixel data
				pixelData := make([]byte, 256*256*2)
				_ = ds.AddOrUpdate(element.NewOtherWord(tag.PixelData, pixelData))

				return ds
			},
			expectedTransferSyntax: "1.2.840.10008.1.2.2",
		},
		{
			name: "uses default when no transfer syntax specified",
			setupDataset: func() *dataset.Dataset {
				ds := dataset.New()

				// Add required image attributes
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Rows, []uint16{128}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Columns, []uint16{128}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

				// Add pixel data
				pixelData := make([]byte, 128*128)
				_ = ds.AddOrUpdate(element.NewOtherByte(tag.PixelData, pixelData))

				return ds
			},
			expectedTransferSyntax: "1.2.840.10008.1.2.1", // Default: Explicit VR Little Endian
		},
		{
			name: "InternalTransferSyntax takes priority over tag",
			setupDataset: func() *dataset.Dataset {
				ts, err := transfer.Parse("1.2.840.10008.1.2.5") // RLE Lossless
				if err != nil {
					t.Fatalf("failed to parse transfer syntax: %v", err)
				}
				ds := dataset.NewWithTransferSyntax(ts)

				// Add a different transfer syntax as a tag (should be ignored)
				_ = ds.AddOrUpdate(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{"1.2.840.10008.1.2"}))

				// Add required image attributes
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Rows, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.Columns, []uint16{256}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
				_ = ds.AddOrUpdate(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

				// Add pixel data
				pixelData := make([]byte, 256*256)
				_ = ds.AddOrUpdate(element.NewOtherByte(tag.PixelData, pixelData))

				return ds
			},
			expectedTransferSyntax: "1.2.840.10008.1.2.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.setupDataset()

			pd, err := FromDataset(ds)
			if err != nil {
				t.Fatalf("FromDataset() error = %v", err)
			}

			if pd.Info.TransferSyntaxUID != tt.expectedTransferSyntax {
				t.Errorf("TransferSyntaxUID = %s, want %s", pd.Info.TransferSyntaxUID, tt.expectedTransferSyntax)
			}
		})
	}
}
