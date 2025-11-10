// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"testing"
)

func TestRLECodec_Name(t *testing.T) {
	codec := NewRLECodec()
	if codec.Name() != "RLE Lossless" {
		t.Errorf("Name() = %q, want %q", codec.Name(), "RLE Lossless")
	}
}

func TestRLECodec_EncodeDecodeSimple(t *testing.T) {
	codec := NewRLECodec()

	// Create simple test data: 10x10 grayscale image with 8-bit pixels
	width := uint16(10)
	height := uint16(10)
	pixelData := make([]byte, width*height)

	// Fill with a simple pattern
	for i := range pixelData {
		pixelData[i] = byte(i % 256)
	}

	src := &PixelData{
		Data:                      pixelData,
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       0, // unsigned
		PlanarConfiguration:       0, // interleaved
		PhotometricInterpretation: "MONOCHROME2",
	}

	// Encode
	encoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	if len(encoded.Data) == 0 {
		t.Fatal("Encode() produced no data")
	}

	// Decode
	decoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	// Verify decoded data matches original
	if !bytes.Equal(pixelData, decoded.Data[:len(pixelData)]) {
		t.Error("Decoded data does not match original")
	}
}

func TestRLECodec_EncodeDecodeRepeating(t *testing.T) {
	codec := NewRLECodec()

	// Create test data with repeating pattern (good for RLE)
	width := uint16(20)
	height := uint16(20)
	pixelData := make([]byte, width*height)

	// Fill with repeating values
	for i := range pixelData {
		pixelData[i] = byte(i/10) % 16
	}

	src := &PixelData{
		Data:                      pixelData,
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	encoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	// RLE should compress repeating data
	if len(encoded.Data) >= len(pixelData) {
		t.Logf("RLE compression: %d bytes -> %d bytes (%.1f%%)",
			len(pixelData), len(encoded.Data),
			float64(len(encoded.Data))/float64(len(pixelData))*100)
	}

	decoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	if !bytes.Equal(pixelData, decoded.Data[:len(pixelData)]) {
		t.Error("Decoded data does not match original")
	}
}

func TestRLECodec_EncodeDecodeRGB(t *testing.T) {
	codec := NewRLECodec()

	// Create RGB test data: 10x10 RGB image
	width := uint16(10)
	height := uint16(10)
	samplesPerPixel := uint16(3)
	pixelData := make([]byte, int(width)*int(height)*int(samplesPerPixel))

	// Fill with RGB pattern
	for i := 0; i < len(pixelData); i += 3 {
		pixelData[i] = byte((i / 3) % 256)   // R
		pixelData[i+1] = byte((i/3 + 50) % 256) // G
		pixelData[i+2] = byte((i/3 + 100) % 256) // B
	}

	src := &PixelData{
		Data:                      pixelData,
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       0,
		PlanarConfiguration:       0, // interleaved
		PhotometricInterpretation: "RGB",
	}

	encoded := &PixelData{
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "RGB",
	}

	err := codec.Encode(src, encoded, nil)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	decoded := &PixelData{
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "RGB",
	}

	err = codec.Decode(encoded, decoded, nil)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if !bytes.Equal(pixelData, decoded.Data[:len(pixelData)]) {
		t.Error("Decoded RGB data does not match original")
	}
}

func TestRLECodec_Encode16Bit(t *testing.T) {
	codec := NewRLECodec()

	// Create 16-bit grayscale test data
	width := uint16(8)
	height := uint16(8)
	pixelData := make([]byte, int(width)*int(height)*2) // 2 bytes per pixel

	// Fill with 16-bit pattern
	for i := 0; i < len(pixelData); i += 2 {
		val := uint16((i / 2) * 257) // Spread across 16-bit range
		pixelData[i] = byte(val & 0xFF)
		pixelData[i+1] = byte((val >> 8) & 0xFF)
	}

	src := &PixelData{
		Data:                      pixelData,
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            1,
		BitsAllocated:             16,
		BitsStored:                16,
		HighBit:                   15,
		SamplesPerPixel:           1,
		PixelRepresentation:       0,
		PlanarConfiguration:       0,
		PhotometricInterpretation: "MONOCHROME2",
	}

	encoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	decoded := &PixelData{
		Width:                     width,
		Height:                    height,
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

	if !bytes.Equal(pixelData, decoded.Data[:len(pixelData)]) {
		t.Error("Decoded 16-bit data does not match original")
	}
}

func TestRLEEncoder_AllSameValue(t *testing.T) {
	// Test encoding data with all same values
	encoder := newRLEEncoder()
	encoder.NextSegment()

	// Encode 200 bytes of the same value
	for i := 0; i < 200; i++ {
		encoder.Encode(0xAA)
	}
	encoder.Flush()

	data := encoder.GetBuffer()
	if len(data) == 0 {
		t.Fatal("Encoder produced no data")
	}

	// The encoded data should be much smaller than 200 bytes
	if len(data) > 100 {
		t.Errorf("RLE compression inefficient for repeating data: %d bytes", len(data))
	}
}

func TestRLEEncoder_NoRepetition(t *testing.T) {
	// Test encoding data with no repetition
	encoder := newRLEEncoder()
	encoder.NextSegment()

	// Encode sequential bytes
	for i := 0; i < 100; i++ {
		encoder.Encode(byte(i))
	}
	encoder.Flush()

	data := encoder.GetBuffer()
	if len(data) == 0 {
		t.Fatal("Encoder produced no data")
	}
}

func TestBaseParameters(t *testing.T) {
	params := NewBaseParameters()

	// Test setting and getting parameters
	params.SetParameter("quality", 90)
	params.SetParameter("name", "test")

	if val := params.GetParameter("quality"); val != 90 {
		t.Errorf("GetParameter(quality) = %v, want 90", val)
	}

	if val := params.GetParameter("name"); val != "test" {
		t.Errorf("GetParameter(name) = %v, want \"test\"", val)
	}

	if val := params.GetParameter("nonexistent"); val != nil {
		t.Errorf("GetParameter(nonexistent) = %v, want nil", val)
	}
}
