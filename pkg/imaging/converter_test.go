// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"testing"
)

func TestInterleavedToPlanar24(t *testing.T) {
	converter := NewPixelDataConverter()

	// Input: RGB RGB RGB (interleaved)
	input := []byte{
		255, 0, 0, // Red pixel
		0, 255, 0, // Green pixel
		0, 0, 255, // Blue pixel
	}

	result := converter.InterleavedToPlanar24(input)

	// Expected: RRR GGG BBB (planar)
	expected := []byte{
		255, 0, 0, // R channel
		0, 255, 0, // G channel
		0, 0, 255, // B channel
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("InterleavedToPlanar24 failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

func TestPlanarToInterleaved24(t *testing.T) {
	converter := NewPixelDataConverter()

	// Input: RRR GGG BBB (planar)
	input := []byte{
		255, 0, 0, // R channel
		0, 255, 0, // G channel
		0, 0, 255, // B channel
	}

	result := converter.PlanarToInterleaved24(input)

	// Expected: RGB RGB RGB (interleaved)
	expected := []byte{
		255, 0, 0, // Red pixel
		0, 255, 0, // Green pixel
		0, 0, 255, // Blue pixel
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("PlanarToInterleaved24 failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

func TestInterleavedPlanarRoundTrip(t *testing.T) {
	converter := NewPixelDataConverter()

	original := []byte{
		100, 150, 200,
		50, 75, 25,
		255, 128, 64,
	}

	planar := converter.InterleavedToPlanar24(original)
	backToInterleaved := converter.PlanarToInterleaved24(planar)

	if !bytes.Equal(original, backToInterleaved) {
		t.Errorf("Round trip failed.\nOriginal: %v\nResult: %v", original, backToInterleaved)
	}
}

func TestYBRFullToRGB(t *testing.T) {
	converter := NewPixelDataConverter()

	// YBR_FULL test data (approximate grayscale values)
	input := []byte{
		128, 128, 128, // Neutral gray
		0, 128, 128,   // Black
		255, 128, 128, // White
	}

	result := converter.YBRFullToRGB(input)

	// Check that results are in valid range
	for i, v := range result {
		if v < 0 || v > 255 {
			t.Errorf("Pixel %d out of range: %d", i, v)
		}
	}

	// Black should stay close to black
	if result[3] > 10 || result[4] > 10 || result[5] > 10 {
		t.Errorf("Black pixel not converted correctly: RGB(%d, %d, %d)", result[3], result[4], result[5])
	}

	// White should stay close to white
	if result[6] < 245 || result[7] < 245 || result[8] < 245 {
		t.Errorf("White pixel not converted correctly: RGB(%d, %d, %d)", result[6], result[7], result[8])
	}
}

func TestRGBToYBRFull(t *testing.T) {
	converter := NewPixelDataConverter()

	input := []byte{
		0, 0, 0,       // Black
		255, 255, 255, // White
		128, 128, 128, // Gray
	}

	result := converter.RGBToYBRFull(input)

	// Check that results are in valid range
	for i, v := range result {
		if v < 0 || v > 255 {
			t.Errorf("Pixel %d out of range: %d", i, v)
		}
	}

	// Black: Y should be ~0, Cb/Cr should be ~128
	if result[0] > 10 {
		t.Errorf("Black Y should be near 0, got %d", result[0])
	}

	// White: Y should be ~255, Cb/Cr should be ~128
	if result[3] < 245 {
		t.Errorf("White Y should be near 255, got %d", result[3])
	}
}

func TestYBRFullRoundTrip(t *testing.T) {
	converter := NewPixelDataConverter()

	original := []byte{
		100, 150, 200,
		50, 75, 25,
		255, 128, 64,
	}

	ybr := converter.RGBToYBRFull(original)
	backToRGB := converter.YBRFullToRGB(ybr)

	// Allow small rounding errors
	for i := 0; i < len(original); i++ {
		diff := int(original[i]) - int(backToRGB[i])
		if diff < -2 || diff > 2 {
			t.Errorf("Round trip failed at index %d: original=%d, result=%d, diff=%d",
				i, original[i], backToRGB[i], diff)
		}
	}
}

func TestSwapBytes16(t *testing.T) {
	converter := NewPixelDataConverter()

	input := []byte{0x12, 0x34, 0x56, 0x78}
	result := converter.SwapBytes16(input)

	expected := []byte{0x34, 0x12, 0x78, 0x56}

	if !bytes.Equal(result, expected) {
		t.Errorf("SwapBytes16 failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

func TestSwapBytes32(t *testing.T) {
	converter := NewPixelDataConverter()

	input := []byte{0x12, 0x34, 0x56, 0x78, 0xAB, 0xCD, 0xEF, 0x01}
	result := converter.SwapBytes32(input)

	expected := []byte{0x78, 0x56, 0x34, 0x12, 0x01, 0xEF, 0xCD, 0xAB}

	if !bytes.Equal(result, expected) {
		t.Errorf("SwapBytes32 failed.\nExpected: %v\nGot: %v", expected, result)
	}
}

func TestClampByte(t *testing.T) {
	testCases := []struct {
		input    int
		expected byte
	}{
		{-10, 0},
		{0, 0},
		{128, 128},
		{255, 255},
		{300, 255},
	}

	for _, tc := range testCases {
		result := clampByte(tc.input)
		if result != tc.expected {
			t.Errorf("clampByte(%d): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkInterleavedToPlanar24(b *testing.B) {
	converter := NewPixelDataConverter()
	data := make([]byte, 1920*1080*3) // Full HD RGB image

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = converter.InterleavedToPlanar24(data)
	}
}

func BenchmarkYBRFullToRGB(b *testing.B) {
	converter := NewPixelDataConverter()
	data := make([]byte, 1920*1080*3) // Full HD image

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = converter.YBRFullToRGB(data)
	}
}
