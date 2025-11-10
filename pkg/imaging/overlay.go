// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
)

// OverlayType represents DICOM overlay types
type OverlayType int

const (
	// OverlayGraphics represents a graphic overlay
	OverlayGraphics OverlayType = iota
	// OverlayROI represents a region of interest overlay
	OverlayROI
)

// String returns the string representation of OverlayType
func (t OverlayType) String() string {
	switch t {
	case OverlayGraphics:
		return "Graphics"
	case OverlayROI:
		return "ROI"
	default:
		return "Unknown"
	}
}

// DicomOverlayData represents a DICOM image overlay
type DicomOverlayData struct {
	// Group is the overlay group number (0x6000-0x60FF, even numbers only)
	Group uint16

	// Rows is the number of rows in the overlay
	Rows int

	// Columns is the number of columns in the overlay
	Columns int

	// Type is the overlay type (Graphics or ROI)
	Type OverlayType

	// OriginRow is the row origin of the overlay
	OriginRow int

	// OriginColumn is the column origin of the overlay
	OriginColumn int

	// BitsAllocated is the number of bits allocated (always 1 for overlays)
	BitsAllocated int

	// BitPosition is the bit position for embedded overlays
	BitPosition int

	// Description is the overlay description
	Description string

	// Subtype is the overlay subtype
	Subtype string

	// Label is the overlay label
	Label string

	// NumberOfFrames is the number of frames in the overlay
	NumberOfFrames int

	// ImageFrameOrigin is the frame origin for multi-frame images
	ImageFrameOrigin int

	// Data is the raw overlay data (bit-packed)
	Data []byte
}

// NewDicomOverlayData creates a new overlay data instance
func NewDicomOverlayData(group uint16) *DicomOverlayData {
	return &DicomOverlayData{
		Group:          group,
		BitsAllocated:  1,
		NumberOfFrames: 1,
	}
}

// IsValid checks if the overlay group number is valid
func (o *DicomOverlayData) IsValid() bool {
	// Overlay groups must be 0x6000-0x60FF and even
	if o.Group < 0x6000 || o.Group > 0x60FF {
		return false
	}
	if o.Group%2 != 0 {
		return false
	}
	return true
}

// UnpackedDataSize returns the size of unpacked overlay data in bytes
func (o *DicomOverlayData) UnpackedDataSize() int {
	return o.Rows * o.Columns * o.NumberOfFrames
}

// PackedDataSize returns the size of packed (bit-packed) overlay data in bytes
func (o *DicomOverlayData) PackedDataSize() int {
	bits := o.Rows * o.Columns * o.NumberOfFrames
	return (bits + 7) / 8 // Round up to nearest byte
}

// UnpackData unpacks bit-packed overlay data to byte array (0 or 255)
func (o *DicomOverlayData) UnpackData() []byte {
	if len(o.Data) == 0 {
		return nil
	}

	totalPixels := o.Rows * o.Columns * o.NumberOfFrames
	unpacked := make([]byte, totalPixels)

	bitIndex := 0
	for i := 0; i < totalPixels; i++ {
		byteIndex := bitIndex / 8
		bitOffset := bitIndex % 8

		if byteIndex >= len(o.Data) {
			break
		}

		// Get bit value
		bit := (o.Data[byteIndex] >> bitOffset) & 0x01
		if bit != 0 {
			unpacked[i] = 255 // White
		} else {
			unpacked[i] = 0 // Black
		}

		bitIndex++
	}

	return unpacked
}

// PackData packs byte array (0 or non-zero) to bit-packed overlay data
func (o *DicomOverlayData) PackData(unpacked []byte) error {
	expectedSize := o.Rows * o.Columns * o.NumberOfFrames
	if len(unpacked) != expectedSize {
		return fmt.Errorf("unpacked data size mismatch: expected %d, got %d", expectedSize, len(unpacked))
	}

	packedSize := (len(unpacked) + 7) / 8
	o.Data = make([]byte, packedSize)

	bitIndex := 0
	for i := 0; i < len(unpacked); i++ {
		byteIndex := bitIndex / 8
		bitOffset := bitIndex % 8

		if unpacked[i] != 0 {
			o.Data[byteIndex] |= (1 << bitOffset)
		}

		bitIndex++
	}

	return nil
}

// GetFrame extracts a single frame from multi-frame overlay data
func (o *DicomOverlayData) GetFrame(frameIndex int) ([]byte, error) {
	if frameIndex < 0 || frameIndex >= o.NumberOfFrames {
		return nil, fmt.Errorf("frame index out of range: %d (total: %d)", frameIndex, o.NumberOfFrames)
	}

	unpacked := o.UnpackData()
	if unpacked == nil {
		return nil, fmt.Errorf("failed to unpack overlay data")
	}

	frameSize := o.Rows * o.Columns
	start := frameIndex * frameSize
	end := start + frameSize

	if end > len(unpacked) {
		return nil, fmt.Errorf("frame data out of range")
	}

	frame := make([]byte, frameSize)
	copy(frame, unpacked[start:end])

	return frame, nil
}

// ApplyToImage applies the overlay to an image
// imageData: RGB or grayscale image data
// overlayColor: color to use for overlay (ARGB format)
// Returns: modified image data with overlay applied
func (o *DicomOverlayData) ApplyToImage(imageData []byte, width, height int, samplesPerPixel int, overlayColor Color32) ([]byte, error) {
	if width != o.Columns || height != o.Rows {
		return nil, fmt.Errorf("image dimensions (%dx%d) don't match overlay dimensions (%dx%d)",
			width, height, o.Columns, o.Rows)
	}

	unpacked := o.UnpackData()
	if unpacked == nil {
		return nil, fmt.Errorf("failed to unpack overlay data")
	}

	result := make([]byte, len(imageData))
	copy(result, imageData)

	// Apply overlay with blending
	for i := 0; i < width*height; i++ {
		if unpacked[i] != 0 {
			if samplesPerPixel == 1 {
				// Grayscale: blend with gray value
				gray := (uint16(overlayColor.R) + uint16(overlayColor.G) + uint16(overlayColor.B)) / 3
				result[i] = uint8(gray)
			} else if samplesPerPixel == 3 {
				// RGB: apply overlay color with alpha blending
				alpha := float64(overlayColor.A) / 255.0
				idx := i * 3

				result[idx] = uint8(float64(result[idx])*(1-alpha) + float64(overlayColor.R)*alpha)
				result[idx+1] = uint8(float64(result[idx+1])*(1-alpha) + float64(overlayColor.G)*alpha)
				result[idx+2] = uint8(float64(result[idx+2])*(1-alpha) + float64(overlayColor.B)*alpha)
			}
		}
	}

	return result, nil
}

// OverlayGroupNumbers returns all valid overlay group numbers (0x6000, 0x6002, ..., 0x60FE)
func OverlayGroupNumbers() []uint16 {
	groups := make([]uint16, 128) // (0x60FF - 0x6000) / 2 + 1 = 128
	for i := 0; i < 128; i++ {
		groups[i] = uint16(0x6000 + i*2)
	}
	return groups
}
