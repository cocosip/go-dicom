// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
)

const unknownStr = "Unknown"

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
		return unknownStr
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

	// MultiFrameOverlay reports whether Number of Frames in Overlay (60xx,0015)
	// or Image Frame Origin (60xx,0051) was present. When neither is present the
	// overlay applies to every frame (PS3.3 C.9.2.1.4).
	MultiFrameOverlay bool

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
	if len(o.Data) == 0 {
		return nil, fmt.Errorf("failed to unpack overlay data")
	}

	// Unpack only the requested frame; unpacking every frame here made a
	// multi-frame render cost O(frames) per frame.
	frameSize := o.Rows * o.Columns
	startBit := frameIndex * frameSize
	frame := make([]byte, frameSize)
	for index := 0; index < frameSize; index++ {
		bitIndex := startBit + index
		byteIndex := bitIndex / 8
		if byteIndex >= len(o.Data) {
			break
		}
		if (o.Data[byteIndex]>>uint(bitIndex%8))&0x01 != 0 {
			frame[index] = 255
		}
	}

	return frame, nil
}

// ApplyToImage composites the first overlay frame (index 0) onto imageData.
//
// imageData holds one byte per pixel sample: interleaved RGB when
// samplesPerPixel is 3, a single gray sample when it is 1. The image is one
// frame, so only overlay frame 0 is used; select any other overlay frame with
// GetFrame before compositing.
//
// overlayColor is blended over each marked pixel with straight-alpha
// semantics: alpha 0 leaves the pixel unchanged and alpha 255 replaces it.
func (o *DicomOverlayData) ApplyToImage(imageData []byte, width, height int, samplesPerPixel int, overlayColor colorconv.Color32) ([]byte, error) {
	if width != o.Columns || height != o.Rows {
		return nil, fmt.Errorf("image dimensions (%dx%d) don't match overlay dimensions (%dx%d)",
			width, height, o.Columns, o.Rows)
	}
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("image dimensions must be positive, got %dx%d", width, height)
	}
	if samplesPerPixel != 1 && samplesPerPixel != 3 {
		return nil, fmt.Errorf("unsupported samples per pixel: %d", samplesPerPixel)
	}
	required := int64(width) * int64(height) * int64(samplesPerPixel)
	if int64(len(imageData)) < required {
		return nil, fmt.Errorf("image data too short: got %d bytes, need %d for %dx%d with %d sample(s) per pixel",
			len(imageData), required, width, height, samplesPerPixel)
	}

	unpacked, err := o.GetFrame(0)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(imageData))
	copy(result, imageData)

	// Apply overlay with blending
	for i, marked := range unpacked {
		if marked == 0 {
			continue
		}
		switch samplesPerPixel {
		case 1:
			// Grayscale: blend with the average of the overlay color channels.
			gray := uint8((uint16(overlayColor.R) + uint16(overlayColor.G) + uint16(overlayColor.B)) / 3)
			result[i] = compositeOverlayChannel(result[i], gray, overlayColor.A)
		case 3:
			// RGB: apply overlay color with alpha blending
			idx := i * 3

			result[idx] = compositeOverlayChannel(result[idx], overlayColor.R, overlayColor.A)
			result[idx+1] = compositeOverlayChannel(result[idx+1], overlayColor.G, overlayColor.A)
			result[idx+2] = compositeOverlayChannel(result[idx+2], overlayColor.B, overlayColor.A)
		}
	}

	return result, nil
}

// compositeOverlayChannel blends overlay into base with straight-alpha
// semantics: alpha 0 keeps base and alpha 255 replaces it. It is the single
// compositing implementation shared by ApplyToImage and DicomImage overlays.
func compositeOverlayChannel(base, overlay, alpha uint8) uint8 {
	if alpha == 0 {
		return base
	}
	if alpha == 255 {
		return overlay
	}
	inverse := 255 - uint32(alpha)
	return uint8((uint32(base)*inverse + uint32(overlay)*uint32(alpha) + 127) / 255)
}

// OverlayGroupNumbers returns all valid overlay group numbers (0x6000, 0x6002, ..., 0x60FE)
func OverlayGroupNumbers() []uint16 {
	groups := make([]uint16, 128) // (0x60FF - 0x6000) / 2 + 1 = 128
	for i := 0; i < 128; i++ {
		groups[i] = uint16(0x6000 + i*2) // #nosec G115 -- loop bound ensures result within uint16 range
	}
	return groups
}
