// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

const ybrFull422 = "YBR_FULL_422"

// PixelDataInfo contains metadata about DICOM pixel data.
type PixelDataInfo struct {
	// Image dimensions
	Width  uint16
	Height uint16

	// Number of frames (1 for single frame)
	NumberOfFrames int

	// Bit depth information
	BitsAllocated uint16
	BitsStored    uint16
	HighBit       uint16

	// Sampling information
	SamplesPerPixel uint16

	// Pixel representation
	PixelRepresentation PixelRepresentation

	// Planar configuration
	PlanarConfiguration PlanarConfiguration

	// Photometric interpretation
	PhotometricInterpretation *PhotometricInterpretation

	// Transfer syntax UID
	TransferSyntaxUID string

	// Lossy compression information
	IsLossy                bool
	LossyCompressionMethod string
	LossyCompressionRatio  float64
}

// BytesAllocated returns the number of bytes allocated per pixel sample.
func (info *PixelDataInfo) BytesAllocated() int {
	return int((info.BitsAllocated-1)/8 + 1)
}

// UncompressedFrameSize calculates the uncompressed size of a single frame in bytes.
func (info *PixelDataInfo) UncompressedFrameSize() int {
	if info.BitsAllocated == 1 {
		return (int(info.Width)*int(info.Height)-1)/8 + 1
	}

	// Handle special case for YBR_FULL_422 with uneven width
	actualWidth := int(info.Width)
	if actualWidth%2 != 0 &&
		info.PhotometricInterpretation != nil &&
		(info.PhotometricInterpretation.Value == ybrFull422 ||
			info.PhotometricInterpretation.Value == "YBR_PARTIAL_422" ||
			info.PhotometricInterpretation.Value == "YBR_PARTIAL_420") {
		actualWidth++
	}

	// Handle YBR_FULL_422 special case for uncompressed data
	if info.PhotometricInterpretation != nil && info.PhotometricInterpretation.Value == ybrFull422 {
		// For uncompressed transfer syntaxes, chrominance channels are downsampled
		return info.BytesAllocated() * 2 * actualWidth * int(info.Height)
	}

	return info.BytesAllocated() * int(info.SamplesPerPixel) * actualWidth * int(info.Height)
}

// TotalUncompressedSize returns the total size of all frames uncompressed.
func (info *PixelDataInfo) TotalUncompressedSize() int {
	return info.UncompressedFrameSize() * info.NumberOfFrames
}

// Validate checks if the pixel data info is valid.
func (info *PixelDataInfo) Validate() error {
	if info.Width == 0 {
		return fmt.Errorf("width must be greater than 0")
	}
	if info.Height == 0 {
		return fmt.Errorf("height must be greater than 0")
	}
	if info.NumberOfFrames < 1 {
		return fmt.Errorf("number of frames must be at least 1")
	}
	if info.BitsAllocated == 0 {
		return fmt.Errorf("bits allocated must be greater than 0")
	}
	if info.BitsStored == 0 {
		return fmt.Errorf("bits stored must be greater than 0")
	}
	if info.BitsStored > info.BitsAllocated {
		return fmt.Errorf("bits stored (%d) cannot exceed bits allocated (%d)",
			info.BitsStored, info.BitsAllocated)
	}
	if info.HighBit >= info.BitsAllocated {
		return fmt.Errorf("high bit (%d) must be less than bits allocated (%d)",
			info.HighBit, info.BitsAllocated)
	}
	if info.SamplesPerPixel == 0 {
		return fmt.Errorf("samples per pixel must be greater than 0")
	}
	if info.PhotometricInterpretation == nil {
		return fmt.Errorf("photometric interpretation must be set")
	}

	// Validate samples per pixel based on photometric interpretation
	if info.PhotometricInterpretation.IsColor && info.SamplesPerPixel < 3 {
		return fmt.Errorf("color images require at least 3 samples per pixel, got %d",
			info.SamplesPerPixel)
	}
	if !info.PhotometricInterpretation.IsColor && !info.PhotometricInterpretation.IsPalette && info.SamplesPerPixel != 1 {
		return fmt.Errorf("grayscale images require exactly 1 sample per pixel, got %d",
			info.SamplesPerPixel)
	}

	return nil
}

// DicomPixelData manages DICOM pixel data with support for multiple frames and codecs.
type DicomPixelData struct {
	Info   *PixelDataInfo
	frames [][]byte // Pixel data for each frame
}

// NewDicomPixelData creates a new DicomPixelData instance.
func NewDicomPixelData(info *PixelDataInfo) (*DicomPixelData, error) {
	if err := info.Validate(); err != nil {
		return nil, fmt.Errorf("invalid pixel data info: %w", err)
	}

	return &DicomPixelData{
		Info:   info,
		frames: make([][]byte, 0, info.NumberOfFrames),
	}, nil
}

// NewDicomPixelDataFromBytes creates DicomPixelData from raw pixel bytes.
// The data is assumed to contain all frames concatenated.
func NewDicomPixelDataFromBytes(info *PixelDataInfo, data []byte) (*DicomPixelData, error) {
	pd, err := NewDicomPixelData(info)
	if err != nil {
		return nil, err
	}

	frameSize := info.UncompressedFrameSize()
	expectedSize := frameSize * info.NumberOfFrames

	if len(data) < expectedSize {
		return nil, fmt.Errorf("insufficient data: got %d bytes, expected at least %d bytes",
			len(data), expectedSize)
	}

	// Split data into frames
	for i := 0; i < info.NumberOfFrames; i++ {
		start := i * frameSize
		end := start + frameSize
		if end > len(data) {
			end = len(data)
		}
		frameData := make([]byte, frameSize)
		copy(frameData, data[start:end])
		pd.frames = append(pd.frames, frameData)
	}

	return pd, nil
}

// GetFrame returns the pixel data for the specified frame (0-indexed).
func (pd *DicomPixelData) GetFrame(frameIndex int) ([]byte, error) {
	if frameIndex < 0 || frameIndex >= len(pd.frames) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(pd.frames))
	}
	return pd.frames[frameIndex], nil
}

// AddFrame appends a new frame to the pixel data.
func (pd *DicomPixelData) AddFrame(frameData []byte) error {
	expectedSize := pd.Info.UncompressedFrameSize()
	if len(frameData) < expectedSize {
		return fmt.Errorf("frame data too small: got %d bytes, expected %d bytes",
			len(frameData), expectedSize)
	}

	// Copy the frame data
	frame := make([]byte, expectedSize)
	copy(frame, frameData)
	pd.frames = append(pd.frames, frame)

	// Update frame count
	pd.Info.NumberOfFrames = len(pd.frames)

	return nil
}

// GetAllFrames returns all pixel data as a single byte slice.
func (pd *DicomPixelData) GetAllFrames() []byte {
	totalSize := pd.Info.TotalUncompressedSize()
	result := make([]byte, totalSize)

	offset := 0
	for _, frame := range pd.frames {
		copy(result[offset:], frame)
		offset += len(frame)
	}

	return result
}

// FrameCount returns the number of frames in the pixel data.
func (pd *DicomPixelData) FrameCount() int {
	return len(pd.frames)
}

// ToCodecPixelData converts DicomPixelData to codec.PixelData for codec operations.
func (pd *DicomPixelData) ToCodecPixelData() *codec.PixelData {
	piValue := ""
	if pd.Info.PhotometricInterpretation != nil {
		piValue = pd.Info.PhotometricInterpretation.Value
	}

	return &codec.PixelData{
		Data:                      pd.GetAllFrames(),
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       uint16(pd.Info.PixelRepresentation),
		PlanarConfiguration:       uint16(pd.Info.PlanarConfiguration),
		PhotometricInterpretation: piValue,
		TransferSyntaxUID:         pd.Info.TransferSyntaxUID,
	}
}

// FromCodecPixelData updates DicomPixelData from codec.PixelData after codec operations.
func (pd *DicomPixelData) FromCodecPixelData(codecData *codec.PixelData) error {
	// Clear existing frames
	pd.frames = make([][]byte, 0, pd.Info.NumberOfFrames)

	// Split data into frames
	frameSize := pd.Info.UncompressedFrameSize()
	data := codecData.Data

	for i := 0; i < pd.Info.NumberOfFrames; i++ {
		start := i * frameSize
		end := start + frameSize
		if end > len(data) {
			// Handle partial last frame
			end = len(data)
		}
		if start >= len(data) {
			break
		}

		frameData := make([]byte, frameSize)
		copy(frameData, data[start:end])
		pd.frames = append(pd.frames, frameData)
	}

	return nil
}

// Encode compresses the pixel data using the specified codec.
func (pd *DicomPixelData) Encode(c codec.Codec, params codec.Parameters) (*DicomPixelData, error) {
	src := pd.ToCodecPixelData()

	// Create destination with same metadata
	dst := &codec.PixelData{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       uint16(pd.Info.PixelRepresentation),
		PlanarConfiguration:       uint16(pd.Info.PlanarConfiguration),
		PhotometricInterpretation: pd.Info.PhotometricInterpretation.Value,
	}

	// Encode
	err := c.Encode(src, dst, params)
	if err != nil {
		return nil, fmt.Errorf("encode failed: %w", err)
	}

	// Create new DicomPixelData from encoded result
	// Note: For compressed data, we can't easily split into frames,
	// so we store the entire compressed data as a single blob
	encodedInfo := *pd.Info
	result, err := NewDicomPixelData(&encodedInfo)
	if err != nil {
		return nil, err
	}

	// For compressed data, add all data as a single "frame"
	// This is a simplification - in reality, compressed data needs special handling
	result.frames = [][]byte{dst.Data}

	return result, nil
}

// Decode decompresses the pixel data using the specified codec.
func (pd *DicomPixelData) Decode(c codec.Codec, params codec.Parameters) (*DicomPixelData, error) {
	src := &codec.PixelData{
		Data:                      pd.GetAllFrames(),
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       uint16(pd.Info.PixelRepresentation),
		PlanarConfiguration:       uint16(pd.Info.PlanarConfiguration),
		PhotometricInterpretation: pd.Info.PhotometricInterpretation.Value,
	}

	dst := &codec.PixelData{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       uint16(pd.Info.PixelRepresentation),
		PlanarConfiguration:       uint16(pd.Info.PlanarConfiguration),
		PhotometricInterpretation: pd.Info.PhotometricInterpretation.Value,
	}

	// Decode
	err := c.Decode(src, dst, params)
	if err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	// Create new DicomPixelData from decoded result
	decodedInfo := *pd.Info
	result, err := NewDicomPixelDataFromBytes(&decodedInfo, dst.Data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
