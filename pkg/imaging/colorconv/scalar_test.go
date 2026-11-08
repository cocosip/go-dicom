// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package colorconv

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

func TestRGBToYBRFull(t *testing.T) {
	testCases := []struct {
		r, g, b     uint8
		expectedY   uint8
		expectedCb  uint8
		expectedCr  uint8
		description string
	}{
		{0, 0, 0, 0, 128, 128, "Black"},
		{255, 255, 255, 255, 128, 128, "White"},
		{255, 0, 0, 76, 85, 255, "Red"},
		{0, 255, 0, 150, 44, 21, "Green"},
		{0, 0, 255, 29, 255, 107, "Blue"},
	}

	for _, tc := range testCases {
		y, cb, cr := RGBToYBRFull(tc.r, tc.g, tc.b)

		// Allow small rounding errors
		if math.Abs(float64(y)-float64(tc.expectedY)) > 2 {
			t.Errorf("%s: Y mismatch: expected %d, got %d", tc.description, tc.expectedY, y)
		}
		if math.Abs(float64(cb)-float64(tc.expectedCb)) > 2 {
			t.Errorf("%s: Cb mismatch: expected %d, got %d", tc.description, tc.expectedCb, cb)
		}
		if math.Abs(float64(cr)-float64(tc.expectedCr)) > 2 {
			t.Errorf("%s: Cr mismatch: expected %d, got %d", tc.description, tc.expectedCr, cr)
		}
	}
}

func TestYBRFullToRGB(t *testing.T) {
	testCases := []struct {
		y, cb, cr   uint8
		expectedR   uint8
		expectedG   uint8
		expectedB   uint8
		description string
	}{
		{0, 128, 128, 0, 0, 0, "Black"},
		{255, 128, 128, 255, 255, 255, "White"},
		{128, 128, 128, 128, 128, 128, "Gray"},
	}

	for _, tc := range testCases {
		r, g, b := YBRFullToRGB(tc.y, tc.cb, tc.cr)

		// Allow small rounding errors
		if math.Abs(float64(r)-float64(tc.expectedR)) > 2 {
			t.Errorf("%s: R mismatch: expected %d, got %d", tc.description, tc.expectedR, r)
		}
		if math.Abs(float64(g)-float64(tc.expectedG)) > 2 {
			t.Errorf("%s: G mismatch: expected %d, got %d", tc.description, tc.expectedG, g)
		}
		if math.Abs(float64(b)-float64(tc.expectedB)) > 2 {
			t.Errorf("%s: B mismatch: expected %d, got %d", tc.description, tc.expectedB, b)
		}
	}
}

func TestRGBToYBRFullRoundTrip(t *testing.T) {
	testCases := []struct {
		r, g, b uint8
	}{
		{100, 150, 200},
		{0, 0, 0},
		{255, 255, 255},
		{128, 64, 192},
	}

	for _, tc := range testCases {
		// Convert RGB -> YBR -> RGB
		y, cb, cr := RGBToYBRFull(tc.r, tc.g, tc.b)
		r2, g2, b2 := YBRFullToRGB(y, cb, cr)

		// Allow rounding errors up to 2
		if math.Abs(float64(tc.r)-float64(r2)) > 2 {
			t.Errorf("RGB->YBR->RGB round trip failed for R: %d -> %d", tc.r, r2)
		}
		if math.Abs(float64(tc.g)-float64(g2)) > 2 {
			t.Errorf("RGB->YBR->RGB round trip failed for G: %d -> %d", tc.g, g2)
		}
		if math.Abs(float64(tc.b)-float64(b2)) > 2 {
			t.Errorf("RGB->YBR->RGB round trip failed for B: %d -> %d", tc.b, b2)
		}
	}
}

func TestRGBToYBRFull422(t *testing.T) {
	// Test with two pixels
	r1, g1, b1 := uint8(255), uint8(0), uint8(0) // Red
	r2, g2, b2 := uint8(0), uint8(255), uint8(0) // Green

	y1, y2, cb, cr := RGBToYBRFull422(r1, g1, b1, r2, g2, b2)

	// Y values should be different for different colors
	if y1 == y2 {
		t.Error("Y values should differ for red and green")
	}

	// Cb and Cr should be averaged
	if cb == 0 || cr == 0 {
		t.Logf("Cb=%d, Cr=%d (averaged chroma values)", cb, cr)
	}
}

func TestYBRPartial422ToRGB(t *testing.T) {
	// Test with known values
	y, cb, cr := uint8(128), uint8(128), uint8(128)
	r, g, b := YBRPartial422ToRGB(y, cb, cr)

	// For Y=128, Cb=Cr=128 (neutral), RGB should be close to 128
	// Allow some tolerance for BT.601 conversion
	if math.Abs(float64(r)-128) > 10 {
		t.Errorf("Expected R near 128, got %d", r)
	}
	if math.Abs(float64(g)-128) > 10 {
		t.Errorf("Expected G near 128, got %d", g)
	}
	if math.Abs(float64(b)-128) > 10 {
		t.Errorf("Expected B near 128, got %d", b)
	}
}

func TestRGBToYBRICT(t *testing.T) {
	r, g, b := uint8(100), uint8(150), uint8(200)
	y, cb, cr := RGBToYBRICT(r, g, b)

	// JPEG 2000 irreversible color transform (PS3.3 C.7.6.3.1.3).
	rf, gf, bf := float64(r), float64(g), float64(b)
	if want := int16(math.Round(0.299*rf + 0.587*gf + 0.114*bf)); y != want {
		t.Errorf("Y mismatch: want %d, got %d", want, y)
	}
	if want := int16(math.Round(-0.1687*rf - 0.3313*gf + 0.5*bf)); cb != want {
		t.Errorf("Cb mismatch: want %d, got %d", want, cb)
	}
	if want := int16(math.Round(0.5*rf - 0.4187*gf - 0.0813*bf)); cr != want {
		t.Errorf("Cr mismatch: want %d, got %d", want, cr)
	}
}

func TestYBRICTToRGB(t *testing.T) {
	// Test round trip
	origR, origG, origB := uint8(100), uint8(150), uint8(200)
	y, cb, cr := RGBToYBRICT(origR, origG, origB)
	r, g, b := YBRICTToRGB(y, cb, cr)

	// ICT is not perfectly reversible due to rounding
	// Allow small errors
	if math.Abs(float64(origR)-float64(r)) > 2 {
		t.Errorf("R round trip: %d -> %d", origR, r)
	}
	if math.Abs(float64(origG)-float64(g)) > 2 {
		t.Errorf("G round trip: %d -> %d", origG, g)
	}
	if math.Abs(float64(origB)-float64(b)) > 2 {
		t.Errorf("B round trip: %d -> %d", origB, b)
	}
}

func TestRGBToYBRRCT(t *testing.T) {
	r, g, b := uint8(100), uint8(150), uint8(200)
	y, cb, cr := RGBToYBRRCT(r, g, b)

	// RCT uses floor division
	expectedY := (int16(r) + 2*int16(g) + int16(b)) >> 2
	if y != expectedY {
		t.Errorf("Y mismatch: expected %d, got %d", expectedY, y)
	}

	// Cb = B - G
	expectedCb := int16(b) - int16(g)
	if cb != expectedCb {
		t.Errorf("Cb mismatch: expected %d, got %d", expectedCb, cb)
	}

	// Cr = R - G
	expectedCr := int16(r) - int16(g)
	if cr != expectedCr {
		t.Errorf("Cr mismatch: expected %d, got %d", expectedCr, cr)
	}
}

func TestYBRRCTToRGB(t *testing.T) {
	// RCT is reversible (lossless)
	testCases := []struct {
		r, g, b uint8
	}{
		{0, 0, 0},
		{255, 255, 255},
		{100, 150, 200},
		{50, 100, 150},
	}

	for _, tc := range testCases {
		y, cb, cr := RGBToYBRRCT(tc.r, tc.g, tc.b)
		r, g, b := YBRRCTToRGB(y, cb, cr)

		// RCT should be lossless
		if r != tc.r || g != tc.g || b != tc.b {
			t.Errorf("RCT round trip failed: (%d,%d,%d) -> (%d,%d,%d)",
				tc.r, tc.g, tc.b, r, g, b)
		}
	}
}

func TestConvertToRGB(t *testing.T) {
	// Test RGB (no conversion)
	rgbData := []byte{255, 0, 0, 0, 255, 0, 0, 0, 255} // Red, Green, Blue
	result, err := ConvertToRGB(rgbData, 3, 1, *pixel.RGBPhotometric, pixel.InterleavedPlanar)
	if err != nil {
		t.Fatalf("ConvertToRGB failed: %v", err)
	}
	if len(result) != len(rgbData) {
		t.Errorf("Expected same length for RGB, got %d", len(result))
	}

	// Test YBR_FULL conversion
	ybrData := make([]byte, 3*3) // 3 pixels interleaved
	for i := 0; i < 3; i++ {
		y, cb, cr := RGBToYBRFull(rgbData[i*3], rgbData[i*3+1], rgbData[i*3+2])
		ybrData[i*3] = y
		ybrData[i*3+1] = cb
		ybrData[i*3+2] = cr
	}

	result, err = ConvertToRGB(ybrData, 3, 1, *pixel.YbrFull, pixel.InterleavedPlanar)
	if err != nil {
		t.Fatalf("YBR_FULL conversion failed: %v", err)
	}
	if len(result) != len(rgbData) {
		t.Errorf("Expected same length after conversion, got %d", len(result))
	}
}

func TestConvertToRGBInterleavesPlanarRGB(t *testing.T) {
	planar := []byte{10, 20, 30, 40, 50, 60}
	want := []byte{10, 30, 50, 20, 40, 60}

	got, err := ConvertToRGB(planar, 2, 1, *pixel.RGBPhotometric, pixel.PlanarPlanar)
	if err != nil {
		t.Fatalf("ConvertToRGB() error = %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("ConvertToRGB() = %v, want interleaved %v", got, want)
	}
}

func TestConvertToRGBWithBitDepthHonorsStoredBits(t *testing.T) {
	// Two RGB pixels of 12-bit samples 4095 and 2048, right-justified in
	// 16-bit words (High Bit = 11).
	data := make([]byte, 12)
	for sample := 0; sample < 3; sample++ {
		binary.LittleEndian.PutUint16(data[sample*2:], 4095)
		binary.LittleEndian.PutUint16(data[(sample+3)*2:], 2048)
	}
	got, err := ConvertToRGBWithBitDepth(
		data, 2, 1, *pixel.RGBPhotometric, pixel.InterleavedPlanar, 16, 12, 11,
	)
	if err != nil {
		t.Fatalf("ConvertToRGBWithBitDepth() error = %v", err)
	}
	want := []byte{255, 255, 255, 128, 128, 128}
	if !bytes.Equal(got, want) {
		t.Fatalf("ConvertToRGBWithBitDepth() = %v, want %v", got, want)
	}
}

func TestConvertToRGBWithBitDepthHonorsHighBit(t *testing.T) {
	// The same 12-bit samples left-justified in the 16-bit word (High Bit = 15).
	data := make([]byte, 6)
	for sample := 0; sample < 3; sample++ {
		binary.LittleEndian.PutUint16(data[sample*2:], 4095<<4)
	}
	got, err := ConvertToRGBWithBitDepth(
		data, 1, 1, *pixel.RGBPhotometric, pixel.InterleavedPlanar, 16, 12, 15,
	)
	if err != nil {
		t.Fatalf("ConvertToRGBWithBitDepth() error = %v", err)
	}
	if want := []byte{255, 255, 255}; !bytes.Equal(got, want) {
		t.Fatalf("ConvertToRGBWithBitDepth() = %v, want %v", got, want)
	}
}

func TestConvertToRGBWithBytesPerSampleSixteenBit(t *testing.T) {
	t.Run("interleaved", func(t *testing.T) {
		data := []byte{0xFF, 0xFF, 0x00, 0x80, 0x00, 0x00}
		got, err := ConvertToRGBWithBytesPerSample(data, 1, 1, *pixel.RGBPhotometric, pixel.InterleavedPlanar, 2)
		if err != nil {
			t.Fatalf("ConvertToRGBWithBytesPerSample() error = %v", err)
		}
		if want := []byte{0xFF, 0x80, 0x00}; !bytes.Equal(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("planar", func(t *testing.T) {
		planar := []byte{
			0xFF, 0xFF, 0x34, 0x12, // R plane
			0x00, 0x80, 0xFF, 0x00, // G plane
			0x00, 0x00, 0x00, 0x01, // B plane
		}
		got, err := ConvertToRGBWithBytesPerSample(planar, 2, 1, *pixel.RGBPhotometric, pixel.PlanarPlanar, 2)
		if err != nil {
			t.Fatalf("ConvertToRGBWithBytesPerSample() error = %v", err)
		}
		if want := []byte{0xFF, 0x80, 0x00, 0x12, 0x00, 0x01}; !bytes.Equal(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestClampUint8(t *testing.T) {
	testCases := []struct {
		input    float64
		expected uint8
	}{
		{-10, 0},
		{0, 0},
		{127.4, 127},
		{127.6, 128},
		{255, 255},
		{300, 255},
	}

	for _, tc := range testCases {
		result := clampUint8(tc.input)
		if result != tc.expected {
			t.Errorf("clampUint8(%f): expected %d, got %d", tc.input, tc.expected, result)
		}
	}
}
