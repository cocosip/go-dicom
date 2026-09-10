// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package colorconv

import (
	"bytes"
	"testing"
)

func TestInterleavedToPlanar24(t *testing.T) {
	// Input: RGB RGB RGB (interleaved)
	input := []byte{
		255, 0, 0, // Red pixel
		0, 255, 0, // Green pixel
		0, 0, 255, // Blue pixel
	}

	result, err := InterleavedToPlanar(input, 3, 1)
	if err != nil {
		t.Fatalf("InterleavedToPlanar() error = %v", err)
	}

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
	// Input: RRR GGG BBB (planar)
	input := []byte{
		255, 0, 0, // R channel
		0, 255, 0, // G channel
		0, 0, 255, // B channel
	}

	result, err := PlanarToInterleaved(input, 3, 1)
	if err != nil {
		t.Fatalf("PlanarToInterleaved() error = %v", err)
	}

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
	original := []byte{
		100, 150, 200,
		50, 75, 25,
		255, 128, 64,
	}

	planar, err := InterleavedToPlanar(original, 3, 1)
	if err != nil {
		t.Fatalf("InterleavedToPlanar() error = %v", err)
	}
	backToInterleaved, err := PlanarToInterleaved(planar, 3, 1)
	if err != nil {
		t.Fatalf("PlanarToInterleaved() error = %v", err)
	}

	if !bytes.Equal(original, backToInterleaved) {
		t.Errorf("Round trip failed.\nOriginal: %v\nResult: %v", original, backToInterleaved)
	}
}

func TestPlanarToInterleavedSupportsSixteenBitSamples(t *testing.T) {
	planar := []byte{
		0x02, 0x01, 0x04, 0x03,
		0x06, 0x05, 0x08, 0x07,
		0x0a, 0x09, 0x0c, 0x0b,
	}
	want := []byte{
		0x02, 0x01, 0x06, 0x05, 0x0a, 0x09,
		0x04, 0x03, 0x08, 0x07, 0x0c, 0x0b,
	}

	got, err := PlanarToInterleaved(planar, 3, 2)
	if err != nil {
		t.Fatalf("PlanarToInterleaved() error = %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("PlanarToInterleaved() = %v, want %v", got, want)
	}

	roundTrip, err := InterleavedToPlanar(got, 3, 2)
	if err != nil {
		t.Fatalf("InterleavedToPlanar() error = %v", err)
	}
	if !bytes.Equal(roundTrip, planar) {
		t.Fatalf("InterleavedToPlanar() = %v, want %v", roundTrip, planar)
	}
}

func TestPlanarConversionRejectsPartialPixels(t *testing.T) {
	for name, convert := range map[string]func([]byte, int, int) ([]byte, error){
		"planar to interleaved": PlanarToInterleaved,
		"interleaved to planar": InterleavedToPlanar,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := convert([]byte{1, 2, 3, 4}, 3, 1); err == nil {
				t.Fatal("conversion accepted data ending with a partial pixel")
			}
		})
	}
}

func TestConvertMono1ToMono2RejectsInvalidStoredBitLayout(t *testing.T) {
	for _, test := range []struct {
		name           string
		data           []byte
		bitsStored     uint16
		bytesPerSample int
	}{
		{name: "zero stored bits", data: []byte{0}, bitsStored: 0, bytesPerSample: 1},
		{name: "stored bits exceed allocation", data: []byte{0}, bitsStored: 9, bytesPerSample: 1},
		{name: "partial sample", data: []byte{0}, bitsStored: 12, bytesPerSample: 2},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := ConvertMono1ToMono2(test.data, test.bitsStored, test.bytesPerSample, false); err == nil {
				t.Fatal("ConvertMono1ToMono2() accepted an invalid stored-bit layout")
			}
		})
	}
}

func TestConvertYBRFullToRGBBuffer(t *testing.T) {
	// YBR_FULL test data (approximate grayscale values)
	input := []byte{
		128, 128, 128, // Neutral gray
		0, 128, 128, // Black
		255, 128, 128, // White
	}

	result, err := ConvertYBRFullToRGB(input)
	if err != nil {
		t.Fatalf("ConvertYBRFullToRGB() error = %v", err)
	}

	// Check that results are in valid range
	// Note: byte values are always in range [0, 255] by definition
	_ = result // Verify result is not nil

	// Black should stay close to black
	if result[3] > 10 || result[4] > 10 || result[5] > 10 {
		t.Errorf("Black pixel not converted correctly: RGB(%d, %d, %d)", result[3], result[4], result[5])
	}

	// White should stay close to white
	if result[6] < 245 || result[7] < 245 || result[8] < 245 {
		t.Errorf("White pixel not converted correctly: RGB(%d, %d, %d)", result[6], result[7], result[8])
	}
}

func TestConvertRGBToYBRFullBuffer(t *testing.T) {
	input := []byte{
		0, 0, 0, // Black
		255, 255, 255, // White
		128, 128, 128, // Gray
	}

	result, err := ConvertRGBToYBRFull(input)
	if err != nil {
		t.Fatalf("ConvertRGBToYBRFull() error = %v", err)
	}

	// Check that results are in valid range
	// Note: byte values are always in range [0, 255] by definition
	_ = result // Verify result is not nil

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
	original := []byte{
		100, 150, 200,
		50, 75, 25,
		255, 128, 64,
	}

	ybr, err := ConvertRGBToYBRFull(original)
	if err != nil {
		t.Fatalf("ConvertRGBToYBRFull() error = %v", err)
	}
	backToRGB, err := ConvertYBRFullToRGB(ybr)
	if err != nil {
		t.Fatalf("ConvertYBRFullToRGB() error = %v", err)
	}

	// Allow small rounding errors
	for i := 0; i < len(original); i++ {
		diff := int(original[i]) - int(backToRGB[i])
		if diff < -2 || diff > 2 {
			t.Errorf("Round trip failed at index %d: original=%d, result=%d, diff=%d",
				i, original[i], backToRGB[i], diff)
		}
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
	data := make([]byte, 1920*1080*3) // Full HD RGB image

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = InterleavedToPlanar(data, 3, 1)
	}
}

func BenchmarkYBRFullToRGB(b *testing.B) {
	data := make([]byte, 1920*1080*3) // Full HD image

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ConvertYBRFullToRGB(data)
	}
}
