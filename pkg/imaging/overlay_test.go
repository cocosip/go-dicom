// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"testing"
)

func TestOverlayType_String(t *testing.T) {
	testCases := []struct {
		overlayType OverlayType
		expected    string
	}{
		{OverlayGraphics, "Graphics"},
		{OverlayROI, "ROI"},
	}

	for _, tc := range testCases {
		result := tc.overlayType.String()
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}

func TestNewDicomOverlayData(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)

	if overlay.Group != 0x6000 {
		t.Errorf("Expected Group=0x6000, got 0x%X", overlay.Group)
	}

	if overlay.BitsAllocated != 1 {
		t.Errorf("Expected BitsAllocated=1, got %d", overlay.BitsAllocated)
	}

	if overlay.NumberOfFrames != 1 {
		t.Errorf("Expected NumberOfFrames=1, got %d", overlay.NumberOfFrames)
	}
}

func TestDicomOverlayData_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		group    uint16
		expected bool
	}{
		{"valid 0x6000", 0x6000, true},
		{"valid 0x6002", 0x6002, true},
		{"valid 0x60FE", 0x60FE, true},
		{"invalid: too low", 0x5FFE, false},
		{"invalid: too high", 0x6100, false},
		{"invalid: odd", 0x6001, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			overlay := NewDicomOverlayData(tc.group)
			result := overlay.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid for group 0x%X: expected %v, got %v",
					tc.group, tc.expected, result)
			}
		})
	}
}

func TestDicomOverlayData_UnpackedDataSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 512
	overlay.Columns = 512
	overlay.NumberOfFrames = 1

	expected := 512 * 512 * 1
	result := overlay.UnpackedDataSize()

	if result != expected {
		t.Errorf("UnpackedDataSize: expected %d, got %d", expected, result)
	}

	// Multi-frame
	overlay.NumberOfFrames = 10
	expected = 512 * 512 * 10
	result = overlay.UnpackedDataSize()

	if result != expected {
		t.Errorf("UnpackedDataSize (multi-frame): expected %d, got %d", expected, result)
	}
}

func TestDicomOverlayData_PackedDataSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 100
	overlay.Columns = 100
	overlay.NumberOfFrames = 1

	// 100*100 = 10000 bits = 1250 bytes
	expected := 1250
	result := overlay.PackedDataSize()

	if result != expected {
		t.Errorf("PackedDataSize: expected %d, got %d", expected, result)
	}

	// Test with size not divisible by 8
	overlay.Rows = 101
	overlay.Columns = 101
	// 101*101 = 10201 bits = 1276 bytes (rounded up)
	expected = 1276
	result = overlay.PackedDataSize()

	if result != expected {
		t.Errorf("PackedDataSize (rounded): expected %d, got %d", expected, result)
	}
}

func TestDicomOverlayData_PackUnpack(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Create test data (4x4 = 16 pixels)
	unpacked := []byte{
		255, 0, 255, 0,
		0, 255, 0, 255,
		255, 0, 255, 0,
		0, 255, 0, 255,
	}

	// Pack
	err := overlay.PackData(unpacked)
	if err != nil {
		t.Fatalf("PackData failed: %v", err)
	}

	// Check packed size (16 bits = 2 bytes)
	if len(overlay.Data) != 2 {
		t.Errorf("Packed size: expected 2 bytes, got %d", len(overlay.Data))
	}

	// Unpack
	result := overlay.UnpackData()

	// Verify
	if len(result) != len(unpacked) {
		t.Fatalf("Unpacked size: expected %d, got %d", len(unpacked), len(result))
	}

	for i := 0; i < len(unpacked); i++ {
		expected := byte(0)
		if unpacked[i] != 0 {
			expected = 255
		}
		if result[i] != expected {
			t.Errorf("Pixel %d: expected %d, got %d", i, expected, result[i])
		}
	}
}

func TestDicomOverlayData_PackData_InvalidSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Wrong size data
	unpacked := []byte{1, 2, 3} // Should be 16

	err := overlay.PackData(unpacked)
	if err == nil {
		t.Error("Expected error for invalid data size, got nil")
	}
}

func TestDicomOverlayData_GetFrame(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.NumberOfFrames = 2

	// Create data for 2 frames (2x2x2 = 8 pixels)
	unpacked := []byte{
		255, 0, // Frame 0
		0, 255,
		0, 255, // Frame 1
		255, 0,
	}

    _ = overlay.PackData(unpacked)

	// Get frame 0
	frame0, err := overlay.GetFrame(0)
	if err != nil {
		t.Fatalf("GetFrame(0) failed: %v", err)
	}

	if len(frame0) != 4 {
		t.Fatalf("Frame 0 size: expected 4, got %d", len(frame0))
	}

	// Check frame 0 values
	if frame0[0] != 255 || frame0[1] != 0 || frame0[2] != 0 || frame0[3] != 255 {
		t.Errorf("Frame 0 values incorrect: %v", frame0)
	}

	// Get frame 1
	frame1, err := overlay.GetFrame(1)
	if err != nil {
		t.Fatalf("GetFrame(1) failed: %v", err)
	}

	if frame1[0] != 0 || frame1[1] != 255 || frame1[2] != 255 || frame1[3] != 0 {
		t.Errorf("Frame 1 values incorrect: %v", frame1)
	}

	// Invalid frame index
	_, err = overlay.GetFrame(2)
	if err == nil {
		t.Error("Expected error for invalid frame index, got nil")
	}
}

func TestDicomOverlayData_ApplyToImage_Grayscale(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Create overlay data (checkerboard)
	unpacked := []byte{
		255, 0, 255, 0,
		0, 255, 0, 255,
		255, 0, 255, 0,
		0, 255, 0, 255,
	}
    _ = overlay.PackData(unpacked)

	// Create grayscale image (all 128)
	imageData := make([]byte, 16)
	for i := range imageData {
		imageData[i] = 128
	}

	// Apply overlay
	overlayColor := NewColor32(255, 200, 200, 200) // Light gray
	result, err := overlay.ApplyToImage(imageData, 4, 4, 1, overlayColor)
	if err != nil {
		t.Fatalf("ApplyToImage failed: %v", err)
	}

	// Check that overlay pixels were modified
	if result[0] == 128 {
		t.Error("Overlay pixel should have been modified")
	}

	// Check that non-overlay pixels remained the same
	if result[1] != 128 {
		t.Error("Non-overlay pixel should remain unchanged")
	}
}

func TestDicomOverlayData_ApplyToImage_RGB(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.NumberOfFrames = 1

	// Simple overlay (top-left pixel only)
	unpacked := []byte{255, 0, 0, 0}
    _ = overlay.PackData(unpacked)

	// Create RGB image (all black)
	imageData := make([]byte, 2*2*3)

	// Apply red overlay
	overlayColor := NewColor32(255, 255, 0, 0) // Red
	result, err := overlay.ApplyToImage(imageData, 2, 2, 3, overlayColor)
	if err != nil {
		t.Fatalf("ApplyToImage failed: %v", err)
	}

	// First pixel should be blended with red
	if result[0] == 0 {
		t.Error("Overlay pixel R should have been modified")
	}
}

func TestDicomOverlayData_ApplyToImage_DimensionMismatch(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4

	imageData := make([]byte, 16)

	overlayColor := NewColor32(255, 255, 255, 255)

	// Mismatched dimensions
	_, err := overlay.ApplyToImage(imageData, 8, 8, 1, overlayColor)
	if err == nil {
		t.Error("Expected error for dimension mismatch, got nil")
	}
}

func TestOverlayGroupNumbers(t *testing.T) {
	groups := OverlayGroupNumbers()

	if len(groups) != 128 {
		t.Errorf("Expected 128 overlay groups, got %d", len(groups))
	}

	// Check first and last
	if groups[0] != 0x6000 {
		t.Errorf("First group: expected 0x6000, got 0x%X", groups[0])
	}

	if groups[127] != 0x60FE {
		t.Errorf("Last group: expected 0x60FE, got 0x%X", groups[127])
	}

	// Check that all are even
	for i, g := range groups {
		if g%2 != 0 {
			t.Errorf("Group[%d]=0x%X is odd", i, g)
		}
	}
}
