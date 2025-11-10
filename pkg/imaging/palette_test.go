// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"testing"
)

func TestNewColor32(t *testing.T) {
	color := NewColor32(255, 100, 150, 200)

	if color.A != 255 {
		t.Errorf("Expected A=255, got %d", color.A)
	}
	if color.R != 100 {
		t.Errorf("Expected R=100, got %d", color.R)
	}
	if color.G != 150 {
		t.Errorf("Expected G=150, got %d", color.G)
	}
	if color.B != 200 {
		t.Errorf("Expected B=200, got %d", color.B)
	}
}

func TestNewPaletteColorLUT_8bit(t *testing.T) {
	// Create 8-bit palette LUT
	descriptor := []uint16{256, 0, 8} // size=256, firstMapped=0, bits=8
	red := make([]byte, 256)
	green := make([]byte, 256)
	blue := make([]byte, 256)

	// Create a simple gradient
	for i := 0; i < 256; i++ {
		red[i] = byte(i)
		green[i] = byte(255 - i)
		blue[i] = 128
	}

	lut, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err != nil {
		t.Fatalf("Failed to create palette LUT: %v", err)
	}

	if lut.Size != 256 {
		t.Errorf("Expected size=256, got %d", lut.Size)
	}

	if lut.Bits != 8 {
		t.Errorf("Expected bits=8, got %d", lut.Bits)
	}

	// Check some entries
	color := lut.GetColor(0)
	if color.R != 0 || color.G != 255 || color.B != 128 {
		t.Errorf("Color[0]: expected RGB(0,255,128), got RGB(%d,%d,%d)",
			color.R, color.G, color.B)
	}

	color = lut.GetColor(255)
	if color.R != 255 || color.G != 0 || color.B != 128 {
		t.Errorf("Color[255]: expected RGB(255,0,128), got RGB(%d,%d,%d)",
			color.R, color.G, color.B)
	}
}

func TestNewPaletteColorLUT_16bit(t *testing.T) {
	// Create 16-bit palette LUT with 8-bit entries
	descriptor := []uint16{256, 0, 16} // size=256, firstMapped=0, bits=16
	red := make([]byte, 512)           // 256 * 2 bytes
	green := make([]byte, 512)
	blue := make([]byte, 512)

	// 16-bit entries with 8-bits stored (stored in high byte, offset=1)
	for i := 0; i < 256; i++ {
		red[i*2+1] = byte(i)
		green[i*2+1] = byte(255 - i)
		blue[i*2+1] = 128
	}

	lut, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err != nil {
		t.Fatalf("Failed to create 16-bit palette LUT: %v", err)
	}

	// Check entries
	color := lut.GetColor(0)
	if color.R != 0 || color.G != 255 || color.B != 128 {
		t.Errorf("Color[0]: expected RGB(0,255,128), got RGB(%d,%d,%d)",
			color.R, color.G, color.B)
	}
}

func TestPaletteColorLUT_SizeZero(t *testing.T) {
	// Size=0 means 65536
	descriptor := []uint16{0, 0, 8}
	red := make([]byte, 65536)
	green := make([]byte, 65536)
	blue := make([]byte, 65536)

	lut, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err != nil {
		t.Fatalf("Failed to create LUT with size=0: %v", err)
	}

	if lut.Size != 65536 {
		t.Errorf("Expected size=65536 for descriptor size=0, got %d", lut.Size)
	}
}

func TestPaletteColorLUT_GetColor_OutOfRange(t *testing.T) {
	descriptor := []uint16{256, 0, 8}
	red := make([]byte, 256)
	green := make([]byte, 256)
	blue := make([]byte, 256)

	lut, _ := NewPaletteColorLUT(descriptor, red, green, blue)

	// Request color beyond LUT size
	color := lut.GetColor(500)

	// Should return black
	if color.R != 0 || color.G != 0 || color.B != 0 {
		t.Errorf("Out of range color should be black, got RGB(%d,%d,%d)",
			color.R, color.G, color.B)
	}
}

func TestPaletteColorLUT_ApplyToPixelData_8bit(t *testing.T) {
	descriptor := []uint16{4, 0, 8} // Small palette for testing
	red := []byte{255, 0, 128, 64}
	green := []byte{0, 255, 128, 192}
	blue := []byte{0, 0, 255, 255}

	lut, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err != nil {
		t.Fatalf("Failed to create palette LUT: %v", err)
	}

	// Apply to 8-bit pixel data
	pixelData := []byte{0, 1, 2, 3} // 4 pixels
	rgbData, err := lut.ApplyToPixelData(pixelData, 8)
	if err != nil {
		t.Fatalf("ApplyToPixelData failed: %v", err)
	}

	expected := []byte{
		255, 0, 0, // Pixel 0: red
		0, 255, 0, // Pixel 1: green
		128, 128, 255, // Pixel 2
		64, 192, 255, // Pixel 3
	}

	if len(rgbData) != len(expected) {
		t.Fatalf("Expected %d bytes, got %d", len(expected), len(rgbData))
	}

	for i, v := range expected {
		if rgbData[i] != v {
			t.Errorf("Byte %d: expected %d, got %d", i, v, rgbData[i])
		}
	}
}

func TestPaletteColorLUT_ApplyToPixelData_16bit(t *testing.T) {
	descriptor := []uint16{256, 0, 8}
	red := make([]byte, 256)
	green := make([]byte, 256)
	blue := make([]byte, 256)

	// Simple gradient
	for i := 0; i < 256; i++ {
		red[i] = byte(i)
		green[i] = 0
		blue[i] = 0
	}

	lut, _ := NewPaletteColorLUT(descriptor, red, green, blue)

	// 16-bit pixel data (little endian)
	pixelData := []byte{
		0x00, 0x00, // Pixel value 0
		0xFF, 0x00, // Pixel value 255
	}

	rgbData, err := lut.ApplyToPixelData(pixelData, 16)
	if err != nil {
		t.Fatalf("ApplyToPixelData failed: %v", err)
	}

	// Check first pixel (value 0)
	if rgbData[0] != 0 {
		t.Errorf("Pixel 0 R: expected 0, got %d", rgbData[0])
	}

	// Check second pixel (value 255)
	if rgbData[3] != 255 {
		t.Errorf("Pixel 1 R: expected 255, got %d", rgbData[3])
	}
}

func TestPaletteColorLUT_InvalidDescriptor(t *testing.T) {
	descriptor := []uint16{256, 0} // Missing bits field
	red := make([]byte, 256)
	green := make([]byte, 256)
	blue := make([]byte, 256)

	_, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err == nil {
		t.Error("Expected error for invalid descriptor, got nil")
	}
}

func TestPaletteColorLUT_InvalidDataSize(t *testing.T) {
	descriptor := []uint16{256, 0, 8}
	red := make([]byte, 100) // Wrong size
	green := make([]byte, 256)
	blue := make([]byte, 256)

	_, err := NewPaletteColorLUT(descriptor, red, green, blue)
	if err == nil {
		t.Error("Expected error for invalid data size, got nil")
	}
}
