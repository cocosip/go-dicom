// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"testing"
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

	src := &PixelData{
		Data:                      pixelData,
		Width:                     10,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	// Encode (should be a simple copy for 8-bit)
	encoded := &PixelData{
		Width:                     10,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err := codec.Encode(src, encoded, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	if !bytes.Equal(pixelData, encoded.Data) {
		t.Error("Encoded 8-bit data does not match original")
	}

	// Decode
	decoded := &PixelData{
		Width:                     10,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err = codec.Decode(encoded, decoded, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if !bytes.Equal(pixelData, decoded.Data) {
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

	src := &PixelData{
		Data:                      pixelData,
		Width:                     4,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	// Encode without swapping
	encoded := &PixelData{
		Width:                     4,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err := codec.Encode(src, encoded, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	if !bytes.Equal(pixelData, encoded.Data) {
		t.Error("Encoded 16-bit data does not match original")
	}

	// Decode
	decoded := &PixelData{
		Width:                     4,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err = codec.Decode(encoded, decoded, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if !bytes.Equal(pixelData, decoded.Data) {
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

	src := &PixelData{
		Data:                      pixelData,
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	// Encode with byte swapping
	params := NewBaseParameters()
	params.SetParameter("swap_bytes", true)

	encoded := &PixelData{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err := codec.Encode(src, encoded, params)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	// Check that bytes are swapped
	expected := []byte{0x34, 0x12, 0x78, 0x56}
	if !bytes.Equal(expected, encoded.Data) {
		t.Errorf("Swapped data = %v, want %v", encoded.Data, expected)
	}

	// Decode with byte swapping should restore original
	decoded := &PixelData{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err = codec.Decode(encoded, decoded, params)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if !bytes.Equal(pixelData, decoded.Data) {
		t.Errorf("Decoded data = %v, want %v", decoded.Data, pixelData)
	}
}

func TestNativeCodec_BigEndian(t *testing.T) {
	codec := NewExplicitVRBigEndianCodec()

	// 16-bit data in big endian
	pixelData := []byte{
		0x12, 0x34, // 0x1234 in BE
		0x56, 0x78, // 0x5678 in BE
	}

	src := &PixelData{
		Data:                      pixelData,
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	// Decode (should swap to little endian by default for big endian codec)
	decoded := &PixelData{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	err := codec.Decode(src, decoded, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Should be swapped to little endian
	expected := []byte{0x34, 0x12, 0x78, 0x56}
	if !bytes.Equal(expected, decoded.Data) {
		t.Errorf("Decoded BE data = %v, want %v", decoded.Data, expected)
	}
}

func TestConvertEndianness(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		bytesPerSample int
		expected       []byte
		expectErr      bool
	}{
		{
			name:           "16-bit swap",
			input:          []byte{0x12, 0x34, 0x56, 0x78},
			bytesPerSample: 2,
			expected:       []byte{0x34, 0x12, 0x78, 0x56},
		},
		{
			name:           "32-bit swap",
			input:          []byte{0x12, 0x34, 0x56, 0x78},
			bytesPerSample: 4,
			expected:       []byte{0x78, 0x56, 0x34, 0x12},
		},
		{
			name:           "invalid length",
			input:          []byte{0x12, 0x34, 0x56},
			bytesPerSample: 2,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertEndianness(tt.input, tt.bytesPerSample)

			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !bytes.Equal(tt.expected, result) {
				t.Errorf("ConvertEndianness() = %v, want %v", result, tt.expected)
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
