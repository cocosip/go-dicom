// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestPixelDataInfo_Validate(t *testing.T) {
	tests := []struct {
		name      string
		info      *PixelDataInfo
		expectErr bool
	}{
		{
			name: "valid grayscale",
			info: &PixelDataInfo{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             16,
				BitsStored:                12,
				HighBit:                   11,
				SamplesPerPixel:           1,
				PixelRepresentation:       UnsignedPixels,
				PlanarConfiguration:       InterleavedPlanar,
				PhotometricInterpretation: Monochrome2,
			},
			expectErr: false,
		},
		{
			name: "valid RGB",
			info: &PixelDataInfo{
				Width:                     256,
				Height:                    256,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           3,
				PixelRepresentation:       UnsignedPixels,
				PlanarConfiguration:       InterleavedPlanar,
				PhotometricInterpretation: RGBPhotometric,
			},
			expectErr: false,
		},
		{
			name: "zero width",
			info: &PixelDataInfo{
				Width:                     0,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "bits stored exceeds bits allocated",
			info: &PixelDataInfo{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                16,
				HighBit:                   7,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "high bit >= bits allocated",
			info: &PixelDataInfo{
				Width:                     512,
				Height:                    512,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   8,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
			},
			expectErr: true,
		},
		{
			name: "color with insufficient samples",
			info: &PixelDataInfo{
				Width:                     256,
				Height:                    256,
				NumberOfFrames:            1,
				BitsAllocated:             8,
				BitsStored:                8,
				HighBit:                   7,
				SamplesPerPixel:           1, // Should be 3 for RGB
				PixelRepresentation:       UnsignedPixels,
				PlanarConfiguration:       InterleavedPlanar,
				PhotometricInterpretation: RGBPhotometric,
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

func TestFramesFromFragmentsUsesEncodedItemOffsets(t *testing.T) {
	fragments := []buffer.ByteBuffer{
		buffer.NewMemory([]byte{0xAA, 0xBB, 0xCC}),       // len=3 -> padded item payload 4
		buffer.NewMemory([]byte{0x11, 0x22, 0x33, 0x44}), // len=4
		buffer.NewMemory([]byte{0x55, 0x66, 0x77}),       // len=3
	}

	frames, err := framesFromFragments(fragments, []uint32{0, 24}, 2)
	if err != nil {
		t.Fatalf("framesFromFragments() error = %v", err)
	}
	if len(frames) != 2 {
		t.Fatalf("framesFromFragments() returned %d frames, want 2", len(frames))
	}
	if got, want := frames[0], []byte{0xAA, 0xBB, 0xCC, 0x11, 0x22, 0x33, 0x44}; !bytes.Equal(got, want) {
		t.Fatalf("frame 0 = %v, want %v", got, want)
	}
	if got, want := frames[1], []byte{0x55, 0x66, 0x77}; !bytes.Equal(got, want) {
		t.Fatalf("frame 1 = %v, want %v", got, want)
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

func TestPixelDataInfo_UncompressedFrameSize(t *testing.T) {
	tests := []struct {
		name         string
		info         *PixelDataInfo
		expectedSize int
	}{
		{
			name: "8-bit grayscale 512x512",
			info: &PixelDataInfo{
				Width:                     512,
				Height:                    512,
				BitsAllocated:             8,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
			},
			expectedSize: 512 * 512, // 262,144 bytes
		},
		{
			name: "16-bit grayscale 512x512",
			info: &PixelDataInfo{
				Width:                     512,
				Height:                    512,
				BitsAllocated:             16,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
			},
			expectedSize: 512 * 512 * 2, // 524,288 bytes
		},
		{
			name: "8-bit RGB 256x256",
			info: &PixelDataInfo{
				Width:                     256,
				Height:                    256,
				BitsAllocated:             8,
				SamplesPerPixel:           3,
				PhotometricInterpretation: RGBPhotometric,
			},
			expectedSize: 256 * 256 * 3, // 196,608 bytes
		},
		{
			name: "1-bit image 100x100",
			info: &PixelDataInfo{
				Width:                     100,
				Height:                    100,
				BitsAllocated:             1,
				SamplesPerPixel:           1,
				PhotometricInterpretation: Monochrome2,
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

func TestNewDicomPixelData(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     256,
		Height:                    256,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
	}

	if pd.Info != info {
		t.Error("PixelData info not set correctly")
	}

	if pd.FrameCount() != 0 {
		t.Errorf("FrameCount() = %d, want 0", pd.FrameCount())
	}
}

func TestDicomPixelData_AddGetFrame(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     10,
		Height:                    10,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
	}

	// Create test frame data
	frameData := make([]byte, 100)
	for i := range frameData {
		frameData[i] = byte(i % 256)
	}

	// Add frame
	err = pd.AddFrame(frameData)
	if err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	if pd.FrameCount() != 1 {
		t.Errorf("FrameCount() = %d, want 1", pd.FrameCount())
	}

	// Get frame
	retrievedFrame, err := pd.GetFrame(0)
	if err != nil {
		t.Fatalf("GetFrame() error = %v", err)
	}

	if !bytes.Equal(frameData, retrievedFrame) {
		t.Error("Retrieved frame data does not match original")
	}
}

func TestDicomPixelData_MultiFrame(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     8,
		Height:                    8,
		NumberOfFrames:            3,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
	}

	// Add 3 frames
	for i := 0; i < 3; i++ {
		frameData := make([]byte, 64)
		for j := range frameData {
			frameData[j] = byte((i*64 + j) % 256)
		}
		err = pd.AddFrame(frameData)
		if err != nil {
			t.Fatalf("AddFrame(%d) error = %v", i, err)
		}
	}

	if pd.FrameCount() != 3 {
		t.Errorf("FrameCount() = %d, want 3", pd.FrameCount())
	}

	// Verify each frame
	for i := 0; i < 3; i++ {
		frame, err := pd.GetFrame(i)
		if err != nil {
			t.Fatalf("GetFrame(%d) error = %v", i, err)
		}
		if len(frame) != 64 {
			t.Errorf("Frame %d size = %d, want 64", i, len(frame))
		}
	}
}

func TestDicomPixelData_EncodeDecodeRLE(t *testing.T) {
	// Create pixel data
	info := &PixelDataInfo{
		Width:                     16,
		Height:                    16,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
	}

	// Add frame with pattern
	frameData := make([]byte, 256)
	for i := range frameData {
		frameData[i] = byte(i / 16) // Repeating pattern, good for RLE
	}
	err = pd.AddFrame(frameData)
	if err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	// Encode with RLE
	rleCodec := codec.NewRLECodec()
	encoded, err := pd.Encode(rleCodec, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	// Decode
	decoded, err := encoded.Decode(rleCodec, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Verify decoded data matches original
	originalData := pd.GetAllFrames()
	decodedData := decoded.GetAllFrames()

	if !bytes.Equal(originalData, decodedData) {
		t.Error("Decoded data does not match original")
	}
}

func TestDicomPixelData_EncodeDecodeNative(t *testing.T) {
	// Create 16-bit pixel data
	info := &PixelDataInfo{
		Width:                     8,
		Height:                    8,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
	}

	// Add frame
	frameData := make([]byte, 128) // 8x8 x 2 bytes
	for i := 0; i < len(frameData); i += 2 {
		frameData[i] = byte(i & 0xFF)
		frameData[i+1] = byte((i >> 8) & 0xFF)
	}
	err = pd.AddFrame(frameData)
	if err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}

	// Encode/Decode with Native codec (no compression)
	nativeCodec := codec.NewExplicitVRLittleEndianCodec()

	encoded, err := pd.Encode(nativeCodec, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	decoded, err := encoded.Decode(nativeCodec, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Verify data matches
	originalData := pd.GetAllFrames()
	decodedData := decoded.GetAllFrames()

	if !bytes.Equal(originalData, decodedData) {
		t.Error("Decoded data does not match original")
	}
}

func TestDicomPixelData_EnsureInterleaved(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       PlanarPlanar,
		PhotometricInterpretation: RGBPhotometric,
	}

	// Planar: R-plane [1 4], G-plane [2 5], B-plane [3 6]
	data := []byte{1, 4, 2, 5, 3, 6}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.EnsureInterleaved(); err != nil {
		t.Fatalf("EnsureInterleaved() error = %v", err)
	}

	if pd.Info.PlanarConfiguration != InterleavedPlanar {
		t.Fatalf("PlanarConfiguration not updated, got %d", pd.Info.PlanarConfiguration)
	}

	got := pd.GetAllFrames()
	expected := []byte{1, 2, 3, 4, 5, 6}
	if !bytes.Equal(got, expected) {
		t.Fatalf("interleaved data mismatch, got %v want %v", got, expected)
	}
}

func TestDicomPixelData_ConvertMonochrome1ToMonochrome2(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome1,
	}

	data := []byte{0x10, 0x20}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertMonochrome1ToMonochrome2(); err != nil {
		t.Fatalf("ConvertMonochrome1ToMonochrome2() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != Monochrome2 {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.GetAllFrames()
	expected := []byte{0xEF, 0xDF} // 0xFF-0x10, 0xFF-0x20
	if !bytes.Equal(got, expected) {
		t.Fatalf("mono inversion mismatch, got %v want %v", got, expected)
	}
}

func TestDicomPixelData_ToElement_EncapsulatedBOT(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            2,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
		Encapsulated:              true,
		VRCode:                    "OB",
	}

	pd, err := NewDicomPixelData(info)
	if err != nil {
		t.Fatalf("NewDicomPixelData() error = %v", err)
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

func TestDicomPixelData_MinMaxIgnorePadding(t *testing.T) {
	padding := int32(0)
	info := &PixelDataInfo{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{0, 5, 10}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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

func TestDicomPixelData_ConvertYBRFullToRGB(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: YbrFull,
	}

	// Two pixels: (Y,Cb,Cr) = (16,128,128) and (50,128,128) -> RGB roughly (16,16,16) and (50,50,50)
	data := []byte{16, 128, 128, 50, 128, 128}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.GetAllFrames()
	// allow small deviation; here we check exact values for the simple case
	expected := []byte{16, 16, 16, 50, 50, 50}
	if !bytes.Equal(got, expected) {
		t.Fatalf("YBR_FULL->RGB mismatch, got %v want %v", got, expected)
	}
}

func TestDicomPixelData_ConvertYBRFull422ToRGB(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: YbrFull422,
	}

	// Two pixels packed: Y1=16, Y2=50, Cb=128, Cr=128
	data := []byte{16, 50, 128, 128}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.GetAllFrames()
	expected := []byte{16, 16, 16, 50, 50, 50}
	if !bytes.Equal(got, expected) {
		t.Fatalf("YBR_FULL_422->RGB mismatch, got %v want %v", got, expected)
	}
}

func TestCreatePixelData_PaletteToRGB(t *testing.T) {
	// Palette with 2 entries: index 0 -> black, 1 -> red(255,0,0)
	ds := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, nil, []string{"PALETTE COLOR"}),
		// descriptors: number of entries=2, first=0, bits=8
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		// data: R=[0,255], G=[0,0], B=[0,0]
		element.NewOtherByte(tag.RedPaletteColorLookupTableData, []byte{0, 255}),
		element.NewOtherByte(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherByte(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
		// pixel data: indices [0,1]
		element.NewOtherByte(tag.PixelData, []byte{0, 1}),
	})

	pd, err := CreatePixelData(ds)
	if err != nil {
		t.Fatalf("CreatePixelData() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != RGBPhotometric {
		t.Fatalf("expected photometric RGB, got %v", pd.Info.PhotometricInterpretation)
	}
	if pd.Info.SamplesPerPixel != 3 {
		t.Fatalf("expected SPP=3, got %d", pd.Info.SamplesPerPixel)
	}

	data := pd.GetAllFrames()
	expected := []byte{0, 0, 0, 255, 0, 0}
	if !bytes.Equal(data, expected) {
		t.Fatalf("palette -> RGB data mismatch, got %v want %v", data, expected)
	}
}

func TestCreatePixelData_PaletteSegmentedToRGB(t *testing.T) {
	// Use segmented LUT: two segments, values 0 and 255
	ds := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, nil, []string{"PALETTE COLOR"}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		// segmented data: segment type 0, count=1 (2 values), values 0,255 => desc=0x0001
		element.NewOtherWord(tag.SegmentedRedPaletteColorLookupTableData, []byte{0x01, 0x00, 0x00, 0x00, 0xFF, 0x00}),
		element.NewOtherWord(tag.SegmentedGreenPaletteColorLookupTableData, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00}),
		element.NewOtherWord(tag.SegmentedBluePaletteColorLookupTableData, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00}),
		element.NewOtherByte(tag.PixelData, []byte{0, 1}),
	})

	pd, err := CreatePixelData(ds)
	if err != nil {
		t.Fatalf("CreatePixelData() error = %v", err)
	}
	if pd.Info.PhotometricInterpretation != RGBPhotometric {
		t.Fatalf("expected photometric RGB, got %v", pd.Info.PhotometricInterpretation)
	}
	data := pd.GetAllFrames()
	expected := []byte{0, 0, 0, 255, 0, 0}
	if !bytes.Equal(data, expected) {
		t.Fatalf("segmented palette -> RGB mismatch, got %v want %v", data, expected)
	}
}

func TestDicomPixelData_ConvertYBRPartial422ToRGB(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: YbrPartial422,
	}

	// Two pixels packed: Y1=16, Y2=50, Cb=128, Cr=128 (limited range)
	data := []byte{16, 50, 128, 128}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	if pd.Info.PhotometricInterpretation != RGBPhotometric {
		t.Fatalf("PhotometricInterpretation not updated, got %v", pd.Info.PhotometricInterpretation)
	}

	got := pd.GetAllFrames()
	// Expected approximate values: first pixel near 0, second near 39
	if got[0] > 5 || got[1] > 5 || got[2] > 5 {
		t.Fatalf("first pixel not near black: %v", got[:3])
	}
	if got[3] < 35 || got[3] > 45 || got[4] < 35 || got[4] > 45 || got[5] < 35 || got[5] > 45 {
		t.Fatalf("second pixel not near gray: %v", got[3:6])
	}
}

func TestDicomPixelData_ConvertYBRICTToRGB(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           3,
		PixelRepresentation:       SignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: YbrIct,
	}

	// Y=50, Cb=0, Cr=0 should map to roughly RGB(50,50,50)
	data := []byte{50, 0, 0}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	got := pd.GetAllFrames()
	if got[0] < 48 || got[0] > 52 || got[1] < 48 || got[1] > 52 || got[2] < 48 || got[2] > 52 {
		t.Fatalf("YBR_ICT->RGB not near expected 50s: %v", got)
	}
}

func TestDicomPixelData_ConvertYBRRCTToRGB(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     1,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           3,
		PixelRepresentation:       SignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: YbrRct,
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

	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if err := pd.ConvertYBRToRGB(); err != nil {
		t.Fatalf("ConvertYBRToRGB() error = %v", err)
	}

	got := pd.GetAllFrames()
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

func TestDicomPixelData_VOILUTSequence(t *testing.T) {
	// Build dataset with LUT Descriptor [3 entries, first=0, bits=8], LUT data [0,128,255]
	ds := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, nil, []string{"MONOCHROME2"}),
	})

	// VOI LUT Sequence with one item
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
	_ = lutItem.Add(element.NewOtherByte(tag.LUTData, []byte{0, 128, 255}))
	voiSeq := dataset.NewSequence(tag.VOILUTSequence)
	voiSeq.AddItem(lutItem)
	_ = ds.Add(voiSeq)

	// PixelData: 0,1,2
	_ = ds.Add(element.NewOtherByte(tag.PixelData, []byte{0, 1, 2}))

	pd, err := CreatePixelData(ds)
	if err != nil {
		t.Fatalf("CreatePixelData() error = %v", err)
	}

	out, err := pd.WindowOrLUTTo8bit(ds, 0, 0, false)
	if err != nil {
		t.Fatalf("WindowOrLUTTo8bit() error = %v", err)
	}
	got := out[0]
	expected := []byte{0, 128, 255}
	if !bytes.Equal(got, expected) {
		t.Fatalf("VOI LUT mapping mismatch, got %v want %v", got, expected)
	}
}

func TestDicomPixelData_WindowTo8bit(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       SignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	// values: -1000, 0, 1000
	data := []byte{
		0x18, 0xFC, // -1000
		0x00, 0x00, // 0
		0xE8, 0x03, // 1000
	}

	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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

func TestDicomPixelData_WindowTo8bit_Padding(t *testing.T) {
	padding := int32(0)
	info := &PixelDataInfo{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{0, 10, 20} // first is padding
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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

func TestDicomPixelData_MaskPadding(t *testing.T) {
	padding := int32(5)
	info := &PixelDataInfo{
		Width:                     3,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
		PixelPaddingValue:         &padding,
	}

	data := []byte{5, 10, 5}
	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
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

func TestNewDicomPixelDataFromBytes(t *testing.T) {
	info := &PixelDataInfo{
		Width:                     10,
		Height:                    10,
		NumberOfFrames:            2,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       UnsignedPixels,
		PlanarConfiguration:       InterleavedPlanar,
		PhotometricInterpretation: Monochrome2,
	}

	// Create 2 frames of data
	data := make([]byte, 200) // 2 frames x 100 bytes
	for i := range data {
		data[i] = byte(i % 256)
	}

	pd, err := NewDicomPixelDataFromBytes(info, data)
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}

	if pd.FrameCount() != 2 {
		t.Errorf("FrameCount() = %d, want 2", pd.FrameCount())
	}

	// Verify frame 0
	frame0, err := pd.GetFrame(0)
	if err != nil {
		t.Fatalf("GetFrame(0) error = %v", err)
	}
	if !bytes.Equal(data[:100], frame0) {
		t.Error("Frame 0 data mismatch")
	}

	// Verify frame 1
	frame1, err := pd.GetFrame(1)
	if err != nil {
		t.Fatalf("GetFrame(1) error = %v", err)
	}
	if !bytes.Equal(data[100:200], frame1) {
		t.Error("Frame 1 data mismatch")
	}
}

// TestDicomPixelData_ToElement_VRSelection tests that ToElement correctly chooses
// OB vs OW based on BitsAllocated for both encapsulated and native formats.
func TestDicomPixelData_ToElement_VRSelection(t *testing.T) {
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
			name:           "Encapsulated 16-bit should use OW",
			bitsAllocated:  16,
			encapsulated:   true,
			expectedVRType: "OW",
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
			info := &PixelDataInfo{
				Width:                     10,
				Height:                    10,
				NumberOfFrames:            1,
				BitsAllocated:             tt.bitsAllocated,
				BitsStored:                tt.bitsAllocated,
				HighBit:                   tt.bitsAllocated - 1,
				SamplesPerPixel:           1,
				PixelRepresentation:       UnsignedPixels,
				PlanarConfiguration:       InterleavedPlanar,
				PhotometricInterpretation: Monochrome2,
				Encapsulated:              tt.encapsulated,
			}

			pd, err := NewDicomPixelData(info)
			if err != nil {
				t.Fatalf("NewDicomPixelData error = %v", err)
			}

			// Add dummy frame data
			// Calculate correct frame size based on BitsAllocated
			bytesPerPixel := (tt.bitsAllocated + 7) / 8
			frameSize := int(10 * 10 * bytesPerPixel)
			frameData := make([]byte, frameSize)
			for i := range frameData {
				frameData[i] = byte(i % 256)
			}
			if err := pd.AddFrame(frameData); err != nil {
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

// TestCreatePixelData_TransferSyntax verifies that CreatePixelData correctly reads
// transfer syntax from the dataset
func TestCreatePixelData_TransferSyntax(t *testing.T) {
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
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

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
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

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
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

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
				_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

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

			pd, err := CreatePixelData(ds)
			if err != nil {
				t.Fatalf("CreatePixelData() error = %v", err)
			}

			if pd.Info.TransferSyntaxUID != tt.expectedTransferSyntax {
				t.Errorf("TransferSyntaxUID = %s, want %s", pd.Info.TransferSyntaxUID, tt.expectedTransferSyntax)
			}
		})
	}
}
