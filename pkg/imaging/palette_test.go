// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
)

func TestNewColor32(t *testing.T) {
	color := imagetypes.NewColor32(255, 100, 150, 200)

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

func TestPaletteColorLUTApplyToPixelDataWithByteOrder(t *testing.T) {
	lut, err := NewPaletteColorLUT(
		[]uint16{2, 0, 8},
		[]byte{10, 20},
		[]byte{30, 40},
		[]byte{50, 60},
	)
	if err != nil {
		t.Fatalf("NewPaletteColorLUT() error = %v", err)
	}

	got, err := lut.ApplyToPixelDataWithByteOrder([]byte{0, 0, 0, 1}, 16, binary.BigEndian)
	if err != nil {
		t.Fatalf("ApplyToPixelDataWithByteOrder() error = %v", err)
	}
	if want := []byte{10, 30, 50, 20, 40, 60}; !bytes.Equal(got, want) {
		t.Fatalf("palette output = %v, want %v", got, want)
	}
	if _, err := lut.ApplyToPixelDataWithByteOrder([]byte{0}, 16, binary.BigEndian); err == nil {
		t.Fatal("ApplyToPixelDataWithByteOrder() accepted odd-length 16-bit pixels")
	}
}

func TestPaletteColorLUTWithAlphaOutputsRGBA(t *testing.T) {
	lut, err := NewPaletteColorLUTWithAlpha(
		[]uint16{2, 0, 8},
		[]uint16{2, 0, 8},
		[]byte{10, 20},
		[]byte{30, 40},
		[]byte{50, 60},
		[]byte{70, 80},
	)
	if err != nil {
		t.Fatalf("NewPaletteColorLUTWithAlpha() error = %v", err)
	}
	got, err := lut.ApplyToPixelData([]byte{0, 1}, 8)
	if err != nil {
		t.Fatalf("ApplyToPixelData() error = %v", err)
	}
	if want := []byte{10, 30, 50, 70, 20, 40, 60, 80}; !bytes.Equal(got, want) {
		t.Fatalf("palette RGBA output = %v, want %v", got, want)
	}

	for _, alphaDescriptor := range [][]uint16{
		{2, 1, 8},
		{2, 0, 16},
	} {
		if _, err := NewPaletteColorLUTWithAlpha(
			[]uint16{2, 0, 8}, alphaDescriptor,
			[]byte{10, 20}, []byte{30, 40}, []byte{50, 60}, []byte{70, 80},
		); err == nil {
			t.Fatalf("NewPaletteColorLUTWithAlpha() accepted descriptor %v", alphaDescriptor)
		}
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
	descriptor := []uint16{2, 0, 8}
	red := []byte{10, 20}
	green := []byte{30, 40}
	blue := []byte{50, 60}

	lut, _ := NewPaletteColorLUT(descriptor, red, green, blue)

	// Request color beyond LUT size
	color := lut.GetColor(500)

	// Values above the range use the final LUT entry.
	if color.R != 20 || color.G != 40 || color.B != 60 {
		t.Errorf("Out of range color should clamp to the final entry, got RGB(%d,%d,%d)",
			color.R, color.G, color.B)
	}
}

func TestPaletteColorLUTAppliesFirstMappedValueAndClamps(t *testing.T) {
	lut, err := NewPaletteColorLUT(
		[]uint16{2, 10, 8},
		[]byte{10, 20},
		[]byte{30, 40},
		[]byte{50, 60},
	)
	if err != nil {
		t.Fatalf("NewPaletteColorLUT() error = %v", err)
	}

	tests := []struct {
		input uint16
		want  imagetypes.Color32
	}{
		{input: 9, want: imagetypes.NewColor32(255, 10, 30, 50)},
		{input: 10, want: imagetypes.NewColor32(255, 10, 30, 50)},
		{input: 11, want: imagetypes.NewColor32(255, 20, 40, 60)},
		{input: 12, want: imagetypes.NewColor32(255, 20, 40, 60)},
	}
	for _, tt := range tests {
		if got := lut.GetColor(tt.input); got != tt.want {
			t.Errorf("GetColor(%d) = %#v, want %#v", tt.input, got, tt.want)
		}
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
	for _, descriptor := range [][]uint16{
		{256, 0},
		{256, 0, 8, 0},
	} {
		_, err := NewPaletteColorLUT(descriptor, make([]byte, 256), make([]byte, 256), make([]byte, 256))
		if err == nil {
			t.Errorf("NewPaletteColorLUT() accepted descriptor VM %d", len(descriptor))
		}
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

func TestPaletteColorLUTRejectsNonStandardBitDepth(t *testing.T) {
	data := []byte{0, 0, 0, 0}
	if _, err := NewPaletteColorLUT([]uint16{2, 0, 12}, data, data, data); err == nil {
		t.Fatal("NewPaletteColorLUT() accepted bits per entry other than 8 or 16")
	}
}

func TestPaletteColorLUTRejectsTrailingWordData(t *testing.T) {
	data := []byte{0, 0, 0, 0, 0, 0}
	if _, err := NewPaletteColorLUT([]uint16{2, 0, 16}, data, data, data); err == nil {
		t.Fatal("NewPaletteColorLUT() accepted LUT Data longer than the descriptor")
	}
}

func TestPaletteColorLUTAcceptsEvenLengthPaddingForOddEightBitData(t *testing.T) {
	lut, err := NewPaletteColorLUT(
		[]uint16{3, 0, 8},
		[]byte{10, 20, 30, 0},
		[]byte{40, 50, 60, 0},
		[]byte{70, 80, 90, 0},
	)
	if err != nil {
		t.Fatalf("NewPaletteColorLUT() error = %v", err)
	}
	if got := lut.GetColor(2); got != imagetypes.NewColor32(255, 30, 60, 90) {
		t.Fatalf("GetColor(2) = %#v, want final unpadded LUT entry", got)
	}
}
