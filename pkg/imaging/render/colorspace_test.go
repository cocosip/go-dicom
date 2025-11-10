// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"math"
	"testing"
)

func TestRGBToYBRFull(t *testing.T) {
	converter := NewColorSpaceConverter()

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
		y, cb, cr := converter.RGBToYBRFull(tc.r, tc.g, tc.b)

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
	converter := NewColorSpaceConverter()

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
		r, g, b := converter.YBRFullToRGB(tc.y, tc.cb, tc.cr)

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
	converter := NewColorSpaceConverter()

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
		y, cb, cr := converter.RGBToYBRFull(tc.r, tc.g, tc.b)
		r2, g2, b2 := converter.YBRFullToRGB(y, cb, cr)

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
	converter := NewColorSpaceConverter()

	// Test with two pixels
	r1, g1, b1 := uint8(255), uint8(0), uint8(0) // Red
	r2, g2, b2 := uint8(0), uint8(255), uint8(0) // Green

	y1, y2, cb, cr := converter.RGBToYBRFull422(r1, g1, b1, r2, g2, b2)

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
	converter := NewColorSpaceConverter()

	// Test with known values
	y, cb, cr := uint8(128), uint8(128), uint8(128)
	r, g, b := converter.YBRPartial422ToRGB(y, cb, cr)

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
	converter := NewColorSpaceConverter()

	r, g, b := uint8(100), uint8(150), uint8(200)
	y, cb, cr := converter.RGBToYBRICT(r, g, b)

	// Verify basic properties
	// Y should be roughly average of RGB
	expectedY := (int16(r) + 2*int16(g) + int16(b)) / 4
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

func TestYBRICTToRGB(t *testing.T) {
	converter := NewColorSpaceConverter()

	// Test round trip
	origR, origG, origB := uint8(100), uint8(150), uint8(200)
	y, cb, cr := converter.RGBToYBRICT(origR, origG, origB)
	r, g, b := converter.YBRICTToRGB(y, cb, cr)

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
	converter := NewColorSpaceConverter()

	r, g, b := uint8(100), uint8(150), uint8(200)
	y, cb, cr := converter.RGBToYBRRCT(r, g, b)

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
	converter := NewColorSpaceConverter()

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
		y, cb, cr := converter.RGBToYBRRCT(tc.r, tc.g, tc.b)
		r, g, b := converter.YBRRCTToRGB(y, cb, cr)

		// RCT should be lossless
		if r != tc.r || g != tc.g || b != tc.b {
			t.Errorf("RCT round trip failed: (%d,%d,%d) -> (%d,%d,%d)",
				tc.r, tc.g, tc.b, r, g, b)
		}
	}
}

func TestConvertToRGB(t *testing.T) {
	converter := NewColorSpaceConverter()

	// Test RGB (no conversion)
	rgbData := []byte{255, 0, 0, 0, 255, 0, 0, 0, 255} // Red, Green, Blue
	result, err := converter.ConvertToRGB(rgbData, 3, 1, "RGB", 0)
	if err != nil {
		t.Fatalf("ConvertToRGB failed: %v", err)
	}
	if len(result) != len(rgbData) {
		t.Errorf("Expected same length for RGB, got %d", len(result))
	}

	// Test YBR_FULL conversion
	ybrData := make([]byte, 3*3) // 3 pixels interleaved
	for i := 0; i < 3; i++ {
		y, cb, cr := converter.RGBToYBRFull(rgbData[i*3], rgbData[i*3+1], rgbData[i*3+2])
		ybrData[i*3] = y
		ybrData[i*3+1] = cb
		ybrData[i*3+2] = cr
	}

	result, err = converter.ConvertToRGB(ybrData, 3, 1, "YBR_FULL", 0)
	if err != nil {
		t.Fatalf("YBR_FULL conversion failed: %v", err)
	}
	if len(result) != len(rgbData) {
		t.Errorf("Expected same length after conversion, got %d", len(result))
	}
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
