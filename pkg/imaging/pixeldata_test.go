// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/codec"
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
			expectedSize: (100*100 - 1) / 8 + 1, // 1,250 bytes
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
			frameData[j] = byte((i * 64 + j) % 256)
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
