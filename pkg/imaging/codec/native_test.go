// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"context"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

func TestNativeCodec_Name(t *testing.T) {
	tests := []struct {
		name         string
		codec        *NativeCodec
		expectedName string
	}{
		{
			name:         "little endian",
			codec:        NewExplicitVRLittleEndianCodec(),
			expectedName: "Native Little Endian",
		},
		{
			name:         "big endian",
			codec:        NewExplicitVRBigEndianCodec(),
			expectedName: "Native Big Endian",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.codec.Name() != tt.expectedName {
				t.Errorf("Name() = %q, want %q", tt.codec.Name(), tt.expectedName)
			}
		})
	}
}

func TestNativeCodec_EncodeDecode8Bit(t *testing.T) {
	codec := NewExplicitVRLittleEndianCodec()

	// 8-bit grayscale image data (no swapping needed)
	pixelData := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	frameInfo := FrameInfo{
		Width:                     10,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(8, 8, 7, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	src := newTestPixelData(frameInfo)
	_ = src.AddFrame(context.Background(), pixelData)

	// Encode (should be a simple copy for 8-bit)
	encoded := newTestPixelData(frameInfo)

	err := codec.Encode(context.Background(), src, encoded, NativeParameters{})
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	encodedData, _ := encoded.Frame(context.Background(), 0)
	if !bytes.Equal(pixelData, encodedData) {
		t.Error("Encoded 8-bit data does not match original")
	}

	// Decode
	decoded := newTestPixelData(frameInfo)

	err = codec.Decode(context.Background(), encoded, decoded, NativeParameters{})
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	decodedData, _ := decoded.Frame(context.Background(), 0)
	if !bytes.Equal(pixelData, decodedData) {
		t.Error("Decoded 8-bit data does not match original")
	}
}

func TestNativeCodec_EncodeDecode16Bit(t *testing.T) {
	codec := NewExplicitVRLittleEndianCodec()

	// 16-bit grayscale image data (little endian)
	pixelData := []byte{
		0x00, 0x01, // 256
		0x00, 0x02, // 512
		0x00, 0x03, // 768
		0x00, 0x04, // 1024
	}

	frameInfo := FrameInfo{
		Width:                     4,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(16, 16, 15, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	src := newTestPixelData(frameInfo)
	_ = src.AddFrame(context.Background(), pixelData)

	// Encode without swapping
	encoded := newTestPixelData(frameInfo)

	err := codec.Encode(context.Background(), src, encoded, NativeParameters{})
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	encodedData, _ := encoded.Frame(context.Background(), 0)
	if !bytes.Equal(pixelData, encodedData) {
		t.Error("Encoded 16-bit data does not match original")
	}

	// Decode
	decoded := newTestPixelData(frameInfo)

	err = codec.Decode(context.Background(), encoded, decoded, NativeParameters{})
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	decodedData, _ := decoded.Frame(context.Background(), 0)
	if !bytes.Equal(pixelData, decodedData) {
		t.Error("Decoded 16-bit data does not match original")
	}
}

func TestNativeCodec_ByteSwapping16Bit(t *testing.T) {
	codec := NewExplicitVRLittleEndianCodec()

	// 16-bit data in little endian
	pixelData := []byte{
		0x12, 0x34, // 0x3412 in LE
		0x56, 0x78, // 0x7856 in LE
	}

	frameInfo := FrameInfo{
		Width:                     2,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(16, 16, 15, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	src := newTestPixelData(frameInfo)
	_ = src.AddFrame(context.Background(), pixelData)

	// Encode with byte swapping
	params := NativeParameters{ByteSwap: ByteSwapEnabled}

	encoded := newTestPixelData(frameInfo)

	err := codec.Encode(context.Background(), src, encoded, params)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	// Check that bytes are swapped
	expected := []byte{0x34, 0x12, 0x78, 0x56}
	encodedData, _ := encoded.Frame(context.Background(), 0)
	if !bytes.Equal(expected, encodedData) {
		t.Errorf("Swapped data = %v, want %v", encodedData, expected)
	}

	// Decode with byte swapping should restore original
	decoded := newTestPixelData(frameInfo)

	err = codec.Decode(context.Background(), encoded, decoded, params)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	decodedData, _ := decoded.Frame(context.Background(), 0)
	if !bytes.Equal(pixelData, decodedData) {
		t.Errorf("Decoded data = %v, want %v", decodedData, pixelData)
	}
}

func TestNativeCodec_BigEndian(t *testing.T) {
	codec := NewExplicitVRBigEndianCodec()

	// 16-bit data in big endian
	pixelData := []byte{
		0x12, 0x34, // 0x1234 in BE
		0x56, 0x78, // 0x5678 in BE
	}

	frameInfo := FrameInfo{
		Width:                     2,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(16, 16, 15, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	src := newTestPixelData(frameInfo)
	_ = src.AddFrame(context.Background(), pixelData)

	// Decode (should swap to little endian by default for big endian codec)
	decoded := newTestPixelData(frameInfo)

	err := codec.Decode(context.Background(), src, decoded, NativeParameters{})
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Should be swapped to little endian
	expected := []byte{0x34, 0x12, 0x78, 0x56}
	decodedData, _ := decoded.Frame(context.Background(), 0)
	if !bytes.Equal(expected, decodedData) {
		t.Errorf("Decoded BE data = %v, want %v", decodedData, expected)
	}
}

func TestNativeCodec_BigEndianDecodeHonorsByteSwapMode(t *testing.T) {
	frameInfo := FrameInfo{
		Width:                     1,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(16, 16, 15, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	tests := []struct {
		name string
		mode ByteSwapMode
		want []byte
	}{
		{name: "default follows transfer syntax", mode: ByteSwapDefault, want: []byte{0x34, 0x12}},
		{name: "disabled preserves source bytes", mode: ByteSwapDisabled, want: []byte{0x12, 0x34}},
		{name: "enabled swaps source bytes", mode: ByteSwapEnabled, want: []byte{0x34, 0x12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := newTestPixelData(frameInfo)
			if err := source.AddFrame(context.Background(), []byte{0x12, 0x34}); err != nil {
				t.Fatal(err)
			}
			sink := newTestPixelData(frameInfo)

			err := NewExplicitVRBigEndianCodec().Decode(
				context.Background(),
				source,
				sink,
				NativeParameters{ByteSwap: tt.mode},
			)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			got, err := sink.Frame(context.Background(), 0)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Fatalf("Decode() frame = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNativeCodec_ReadWriteUint16(t *testing.T) {
	leCodec := NewExplicitVRLittleEndianCodec()
	beCodec := NewExplicitVRBigEndianCodec()

	data := make([]byte, 4)

	// Test little endian write
	err := leCodec.WriteUint16(data, 0, 0x1234)
	if err != nil {
		t.Fatalf("WriteUint16 LE error: %v", err)
	}
	if data[0] != 0x34 || data[1] != 0x12 {
		t.Errorf("LE write got %v, want [0x34 0x12 ...]", data[:2])
	}

	// Test little endian read
	val, err := leCodec.ReadUint16(data, 0)
	if err != nil {
		t.Fatalf("ReadUint16 LE error: %v", err)
	}
	if val != 0x1234 {
		t.Errorf("LE read got 0x%04X, want 0x1234", val)
	}

	// Test big endian write
	err = beCodec.WriteUint16(data, 2, 0x5678)
	if err != nil {
		t.Fatalf("WriteUint16 BE error: %v", err)
	}
	if data[2] != 0x56 || data[3] != 0x78 {
		t.Errorf("BE write got %v, want [... 0x56 0x78]", data[2:])
	}

	// Test big endian read
	val, err = beCodec.ReadUint16(data, 2)
	if err != nil {
		t.Fatalf("ReadUint16 BE error: %v", err)
	}
	if val != 0x5678 {
		t.Errorf("BE read got 0x%04X, want 0x5678", val)
	}
}

func TestNativeCodec_ReadWriteUint32(t *testing.T) {
	leCodec := NewExplicitVRLittleEndianCodec()
	beCodec := NewExplicitVRBigEndianCodec()

	data := make([]byte, 8)

	// Test little endian write
	err := leCodec.WriteUint32(data, 0, 0x12345678)
	if err != nil {
		t.Fatalf("WriteUint32 LE error: %v", err)
	}
	if data[0] != 0x78 || data[1] != 0x56 || data[2] != 0x34 || data[3] != 0x12 {
		t.Errorf("LE write got %v", data[:4])
	}

	// Test little endian read
	val, err := leCodec.ReadUint32(data, 0)
	if err != nil {
		t.Fatalf("ReadUint32 LE error: %v", err)
	}
	if val != 0x12345678 {
		t.Errorf("LE read got 0x%08X, want 0x12345678", val)
	}

	// Test big endian write
	err = beCodec.WriteUint32(data, 4, 0x9ABCDEF0)
	if err != nil {
		t.Fatalf("WriteUint32 BE error: %v", err)
	}
	if data[4] != 0x9A || data[5] != 0xBC || data[6] != 0xDE || data[7] != 0xF0 {
		t.Errorf("BE write got %v", data[4:])
	}

	// Test big endian read
	val, err = beCodec.ReadUint32(data, 4)
	if err != nil {
		t.Fatalf("ReadUint32 BE error: %v", err)
	}
	if val != 0x9ABCDEF0 {
		t.Errorf("BE read got 0x%08X, want 0x9ABCDEF0", val)
	}
}

func TestNativeCodec_TransferSyntax(t *testing.T) {
	tests := []struct {
		name        string
		codec       *NativeCodec
		expectedUID string
	}{
		{
			name:        "Implicit VR Little Endian",
			codec:       NewImplicitVRLittleEndianCodec(),
			expectedUID: "1.2.840.10008.1.2",
		},
		{
			name:        "Explicit VR Little Endian",
			codec:       NewExplicitVRLittleEndianCodec(),
			expectedUID: "1.2.840.10008.1.2.1",
		},
		{
			name:        "Explicit VR Big Endian",
			codec:       NewExplicitVRBigEndianCodec(),
			expectedUID: "1.2.840.10008.1.2.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := tt.codec.TransferSyntax()
			if ts == nil {
				t.Fatal("TransferSyntax() returned nil")
			}

			if ts.UID().UID() != tt.expectedUID {
				t.Errorf("UID = %s, want %s", ts.UID().UID(), tt.expectedUID)
			}
		})
	}
}
