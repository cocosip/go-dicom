// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/types"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
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

	// VR code for pixel data (OB/OW) and whether encapsulated
	VRCode       string
	Encapsulated bool

	// Transfer syntax UID
	TransferSyntaxUID string

	// Lossy compression information
	IsLossy                bool
	LossyCompressionMethod string
	LossyCompressionRatio  float64

	// Pixel padding (optional)
	PixelPaddingValue      *int32
	PixelPaddingRangeLimit *int32
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
			info.PhotometricInterpretation.Value == photometricYBRPartial422 ||
			info.PhotometricInterpretation.Value == photometricYBRPartial420) {
		actualWidth++
	}

	// Handle YBR_FULL_422 special case for uncompressed data
	if info.PhotometricInterpretation != nil &&
		(info.PhotometricInterpretation.Value == ybrFull422 ||
			info.PhotometricInterpretation.Value == photometricYBRPartial422 ||
			info.PhotometricInterpretation.Value == photometricYBRPartial420) {
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
	if info.PhotometricInterpretation.IsColor && !info.PhotometricInterpretation.IsPalette && info.SamplesPerPixel < 3 {
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
// This type implements the types.PixelData interface.
type DicomPixelData struct {
	Info             *PixelDataInfo
	frames           [][]byte // Per-frame data (uncompressed for native; compressed for encapsulated)
	basicOffsetTable []uint32 // BOT for encapsulated data
}

// Ensure DicomPixelData implements types.PixelData interface
var _ types.PixelData = (*DicomPixelData)(nil)

// NewDicomPixelData creates a new DicomPixelData instance.
func NewDicomPixelData(info *PixelDataInfo) (*DicomPixelData, error) {
	if err := info.Validate(); err != nil {
		return nil, fmt.Errorf("invalid pixel data info: %w", err)
	}

	return &DicomPixelData{
		Info:             info,
		frames:           make([][]byte, 0, info.NumberOfFrames),
		basicOffsetTable: nil,
	}, nil
}

// NewDicomPixelDataFromBytes creates DicomPixelData from raw pixel bytes.
// The data is assumed to contain all frames concatenated.
func NewDicomPixelDataFromBytes(info *PixelDataInfo, data []byte) (*DicomPixelData, error) {
	info.Encapsulated = false
	if info.VRCode == "" {
		if info.BitsAllocated <= 8 {
			info.VRCode = "OB"
		} else {
			info.VRCode = "OW"
		}
	}
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

// CalculateOptimalWindow computes optimal window center/width from pixel data
// by sampling pixel values and finding min/max range
func (pd *DicomPixelData) CalculateOptimalWindow() (center, width float64) {
	if len(pd.frames) == 0 {
		return 0, 256 // Default fallback
	}

	// Sample first frame for window calculation
	pixelData := pd.frames[0]
	bytesPerPixel := int(pd.Info.BitsAllocated) / 8
	pixelCount := len(pixelData) / bytesPerPixel
	isSigned := pd.Info.PixelRepresentation == SignedPixels

	if pixelCount == 0 {
		return 0, 256
	}

	// Sample pixels (use every Nth pixel for speed, but at least 1000 samples)
	step := pixelCount / 1000
	if step < 1 {
		step = 1
	}

	var minVal, maxVal float64
	firstPixel := true

	for i := 0; i < pixelCount; i += step {
		pixelIndex := i * bytesPerPixel
		if pixelIndex+bytesPerPixel > len(pixelData) {
			break
		}

		var pixelValue float64
		switch bytesPerPixel {
		case 1:
			if isSigned {
				pixelValue = float64(int8(pixelData[pixelIndex]))
			} else {
				pixelValue = float64(pixelData[pixelIndex])
			}
		case 2:
			if isSigned {
				pixelValue = float64(int16(pixelData[pixelIndex]) | int16(pixelData[pixelIndex+1])<<8)
			} else {
				pixelValue = float64(uint16(pixelData[pixelIndex]) | uint16(pixelData[pixelIndex+1])<<8)
			}
		default:
			continue
		}

		if firstPixel {
			minVal = pixelValue
			maxVal = pixelValue
			firstPixel = false
		} else {
			if pixelValue < minVal {
				minVal = pixelValue
			}
			if pixelValue > maxVal {
				maxVal = pixelValue
			}
		}
	}

	// Calculate window center and width from min/max
	center = (minVal + maxVal) / 2
	width = maxVal - minVal

	// Ensure reasonable minimum width
	if width < 1 {
		width = 1
	}

	return center, width
}

// AddFrame appends a new frame to the pixel data.
func (pd *DicomPixelData) AddFrame(frameData []byte) error {
	// For encapsulated (compressed) data, frames can be any size
	// For uncompressed data, validate the frame size
	if !pd.Info.Encapsulated {
		expectedSize := pd.Info.UncompressedFrameSize()
		if len(frameData) < expectedSize {
			return fmt.Errorf("frame data too small: got %d bytes, expected %d bytes",
				len(frameData), expectedSize)
		}

		// Copy the frame data (trim to expected size)
		frame := make([]byte, expectedSize)
		copy(frame, frameData)
		pd.frames = append(pd.frames, frame)
	} else {
		// For encapsulated data, just copy the entire frame as-is
		frame := make([]byte, len(frameData))
		copy(frame, frameData)
		pd.frames = append(pd.frames, frame)
	}

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

// EnsureInterleaved converts planar configuration 1 to interleaved (0) for multi-sample pixels.
// Only applies to uncompressed data; encapsulated data must be decoded first.
func (pd *DicomPixelData) EnsureInterleaved() error {
	if pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return fmt.Errorf("cannot convert planar configuration on encapsulated data")
	}
	if pd.Info.SamplesPerPixel <= 1 || pd.Info.PlanarConfiguration == 0 {
		return nil
	}

	bytesPerSample := pd.Info.BytesAllocated()
	convertedFrames := make([][]byte, len(pd.frames))
	for idx, frame := range pd.frames {
		converted, err := ConvertPlanarToInterleavedGeneric(frame, int(pd.Info.SamplesPerPixel), bytesPerSample)
		if err != nil {
			return fmt.Errorf("frame %d planar->interleaved failed: %w", idx, err)
		}
		convertedFrames[idx] = converted
	}
	pd.frames = convertedFrames

	pd.Info.PlanarConfiguration = InterleavedPlanar
	return nil
}

// ConvertMonochrome1ToMonochrome2 inverts grayscale for MONOCHROME1 data.
// Only applies to uncompressed single-sample images.
func (pd *DicomPixelData) ConvertMonochrome1ToMonochrome2() error {
	if pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return fmt.Errorf("cannot convert photometric on encapsulated data")
	}
	if pd.Info.PhotometricInterpretation == nil || pd.Info.PhotometricInterpretation.Value != photometricMonochrome1 {
		return nil
	}
	if pd.Info.SamplesPerPixel != 1 {
		return fmt.Errorf("%s expected SamplesPerPixel=1, got %d", photometricMonochrome1, pd.Info.SamplesPerPixel)
	}

	bytesPerSample := pd.Info.BytesAllocated()
	convertedFrames := make([][]byte, len(pd.frames))
	for fi, frame := range pd.frames {
		converted, err := ConvertMono1ToMono2(frame, pd.Info.BitsStored, bytesPerSample, pd.Info.PixelRepresentation == SignedPixels)
		if err != nil {
			return fmt.Errorf("frame %d mono1->mono2 failed: %w", fi, err)
		}
		convertedFrames[fi] = converted
	}
	pd.frames = convertedFrames

	pd.Info.PhotometricInterpretation = Monochrome2
	return nil
}

// ConvertYBRToRGB converts uncompressed YBR data to interleaved RGB (Photometric=RGB).
// Supports YBR_FULL and YBR_FULL_422 with 8-bit samples; other variants are not handled here.
func (pd *DicomPixelData) ConvertYBRToRGB() error {
	if pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return fmt.Errorf("cannot convert photometric on encapsulated data")
	}
	if pd.Info.PhotometricInterpretation == nil {
		return fmt.Errorf("photometric interpretation missing")
	}
	if pd.Info.BitsAllocated != 8 && pd.Info.BitsAllocated != 16 {
		return fmt.Errorf("YBR->RGB conversion only implemented for BitsAllocated=8 or 16")
	}
	switch pd.Info.PhotometricInterpretation.Value {
	case "YBR_FULL":
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("YBR_FULL expected SamplesPerPixel=3, got %d", pd.Info.SamplesPerPixel)
		}
		for i, frame := range pd.frames {
			converted, err := ConvertYBRFullToRGB(frame)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case "YBR_FULL_422":
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("YBR_FULL_422 expected SamplesPerPixel=3, got %d", pd.Info.SamplesPerPixel)
		}
		width := int(pd.Info.Width)
		if width == 0 {
			return fmt.Errorf("YBR_FULL_422 requires valid width")
		}
		for i, frame := range pd.frames {
			converted, err := ConvertYBRFull422ToRGB(frame, width)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case photometricYBRPartial422:
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("%s expected SamplesPerPixel=3, got %d", photometricYBRPartial422, pd.Info.SamplesPerPixel)
		}
		width := int(pd.Info.Width)
		if width == 0 {
			return fmt.Errorf("%s requires valid width", photometricYBRPartial422)
		}
		for i, frame := range pd.frames {
			converted, err := ConvertYBRPartial422ToRGB(frame, width)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case "YBR_ICT":
		bytesPerSample := pd.Info.BytesAllocated()
		for i, frame := range pd.frames {
			converted, err := ConvertYBRICTToRGB(frame, bytesPerSample*8)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case "YBR_RCT":
		bytesPerSample := pd.Info.BytesAllocated()
		for i, frame := range pd.frames {
			converted, err := ConvertYBRRCTToRGB(frame, bytesPerSample*8)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	default:
		return fmt.Errorf("photometric %s not supported for YBR->RGB conversion", pd.Info.PhotometricInterpretation.Value)
	}

	pd.Info.PhotometricInterpretation = RGBPhotometric
	pd.Info.PlanarConfiguration = InterleavedPlanar
	pd.Info.SamplesPerPixel = 3
	return nil
}

// WindowTo8bit applies a VOI window (center/width) to pixel data and returns 8-bit frames.
// Only supports uncompressed data with BitsAllocated 8 or 16. For multi-frame, returns one []byte per frame.
// If ignorePadding is true and PixelPaddingValue/(RangeLimit) is set, padding samples are forced to 0.
func (pd *DicomPixelData) WindowTo8bit(center, width float64, ignorePadding bool) ([][]byte, error) {
	return applyWindowTo8bit(pd, center, width, ignorePadding)
}

// MinMax returns the minimum和maximum sample values across all frames.
// If ignorePadding is true and PixelPaddingValue/(RangeLimit) is set, padding samples are skipped.
func (pd *DicomPixelData) MinMax(ignorePadding bool) (minVal float64, maxVal float64, err error) {
	return minMaxSamples(pd, ignorePadding)
}

// MaskPadding returns a copy of frames where padding samples are set to 0 and a mask per frame (true = padding).
// Only applies to uncompressed data; encapsulated must be decoded first.
func (pd *DicomPixelData) MaskPadding() (frames [][]byte, masks [][]bool, err error) {
	if pd.Info == nil {
		return nil, nil, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return nil, nil, fmt.Errorf("cannot mask padding on encapsulated data")
	}
	if pd.Info.PixelPaddingValue == nil {
		return nil, nil, fmt.Errorf("no Pixel Padding Value present")
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, nil, fmt.Errorf("unsupported BytesAllocated=%d for padding mask", bytesPerSample)
	}

	padMin := *pd.Info.PixelPaddingValue
	padMax := padMin
	if pd.Info.PixelPaddingRangeLimit != nil {
		padMax = *pd.Info.PixelPaddingRangeLimit
	}

	for _, frame := range pd.frames {
		out := make([]byte, len(frame))
		copy(out, frame)
		mask := make([]bool, len(frame)/bytesPerSample)

		for idx, off := 0, 0; off+bytesPerSample <= len(frame); off, idx = off+bytesPerSample, idx+1 {
			var val int32
			if bytesPerSample == 1 {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int8(frame[off]))
				} else {
					val = int32(frame[off])
				}
			} else {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int16(binary.LittleEndian.Uint16(frame[off:])))
				} else {
					val = int32(binary.LittleEndian.Uint16(frame[off:]))
				}
			}

			if val >= padMin && val <= padMax {
				mask[idx] = true
				// zero out
				for b := 0; b < bytesPerSample; b++ {
					out[off+b] = 0
				}
			}
		}

		frames = append(frames, out)
		masks = append(masks, mask)
	}

	return frames, masks, nil
}

// WindowOrLUTTo8bit applies VOI LUT if present, otherwise window.
func (pd *DicomPixelData) WindowOrLUTTo8bit(ds *dataset.Dataset, center, width float64, ignorePadding bool) ([][]byte, error) {
	if ds != nil {
		if frames, err := applyVOILUT(pd, ds, center, width, ignorePadding); err == nil {
			return frames, nil
		}
	}
	return applyWindowTo8bit(pd, center, width, ignorePadding)
}

// ToElement builds a DICOM pixel data element (OB/OW or encapsulated fragment) from the current frames.
// For encapsulated data, it emits an OB fragment sequence with Basic Offset Table.
func (pd *DicomPixelData) ToElement() (element.Element, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}

	// Encapsulated: create fragment sequence
	// Choose OB/OW based on BitsAllocated (following fo-dicom behavior)
	if pd.Info.Encapsulated {
		return buildFragmentSequence(pd.frames, pd.basicOffsetTable, pd.Info.BitsAllocated)
	}

	// Uncompressed: concatenate frames
	all := pd.GetAllFrames()
	if len(all) == 0 {
		return nil, fmt.Errorf("pixel data is empty")
	}

	if pd.Info.BitsAllocated <= 8 || pd.Info.VRCode == "OB" {
		return element.NewOtherByte(tag.PixelData, all), nil
	}
	return element.NewOtherWord(tag.PixelData, all), nil
}

// IsEncapsulated returns true if pixel data is encapsulated.
func (pd *DicomPixelData) IsEncapsulated() bool {
	return pd.Info != nil && pd.Info.Encapsulated
}

// BasicOffsetTable returns the BOT for encapsulated data.
func (pd *DicomPixelData) BasicOffsetTable() []uint32 {
	return pd.basicOffsetTable
}

// GetFrameInfo returns frame metadata for codec operations.
func (pd *DicomPixelData) GetFrameInfo() *types.FrameInfo {
	piValue := ""
	if pd.Info.PhotometricInterpretation != nil {
		piValue = pd.Info.PhotometricInterpretation.Value
	}

	return &types.FrameInfo{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       uint16(pd.Info.PixelRepresentation),
		PlanarConfiguration:       uint16(pd.Info.PlanarConfiguration),
		PhotometricInterpretation: piValue,
	}
}

// Encode encodes the pixel data using the specified codec and returns a new DicomPixelData.
func (pd *DicomPixelData) Encode(c codec.Codec, params codec.Parameters) (*DicomPixelData, error) {
	if c == nil {
		return nil, fmt.Errorf("codec must not be nil")
	}

	// Create new pixel data info for encoded data
	newInfo := &PixelDataInfo{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       pd.Info.PixelRepresentation,
		PlanarConfiguration:       pd.Info.PlanarConfiguration,
		PhotometricInterpretation: pd.Info.PhotometricInterpretation,
		VRCode:                    "OB", // Encoded data typically uses OB
		Encapsulated:              c.TransferSyntax().IsEncapsulated(),
		TransferSyntaxUID:         c.TransferSyntax().UID().UID(),
		IsLossy:                   pd.Info.IsLossy,
		LossyCompressionMethod:    pd.Info.LossyCompressionMethod,
		LossyCompressionRatio:     pd.Info.LossyCompressionRatio,
		PixelPaddingValue:         pd.Info.PixelPaddingValue,
		PixelPaddingRangeLimit:    pd.Info.PixelPaddingRangeLimit,
	}

	newPD, err := NewDicomPixelData(newInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pixel data: %w", err)
	}

	// Use the codec to encode
	if err := c.Encode(pd, newPD, params); err != nil {
		return nil, fmt.Errorf("failed to encode pixel data: %w", err)
	}

	return newPD, nil
}

// Decode decodes the pixel data using the specified codec and returns a new DicomPixelData.
func (pd *DicomPixelData) Decode(c codec.Codec, params codec.Parameters) (*DicomPixelData, error) {
	if c == nil {
		return nil, fmt.Errorf("codec must not be nil")
	}

	// Create new pixel data info for decoded data
	newInfo := &PixelDataInfo{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		NumberOfFrames:            pd.Info.NumberOfFrames,
		BitsAllocated:             pd.Info.BitsAllocated,
		BitsStored:                pd.Info.BitsStored,
		HighBit:                   pd.Info.HighBit,
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       pd.Info.PixelRepresentation,
		PlanarConfiguration:       pd.Info.PlanarConfiguration,
		PhotometricInterpretation: pd.Info.PhotometricInterpretation,
		VRCode:                    pd.Info.VRCode, // Keep original VR
		Encapsulated:              false,          // Decoded data is not encapsulated
		TransferSyntaxUID:         "1.2.840.10008.1.2.1", // Explicit VR Little Endian
		IsLossy:                   pd.Info.IsLossy,
		LossyCompressionMethod:    pd.Info.LossyCompressionMethod,
		LossyCompressionRatio:     pd.Info.LossyCompressionRatio,
		PixelPaddingValue:         pd.Info.PixelPaddingValue,
		PixelPaddingRangeLimit:    pd.Info.PixelPaddingRangeLimit,
	}

	newPD, err := NewDicomPixelData(newInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pixel data: %w", err)
	}

	// Use the codec to decode
	if err := c.Decode(pd, newPD, params); err != nil {
		return nil, fmt.Errorf("failed to decode pixel data: %w", err)
	}

	return newPD, nil
}

// CreatePixelData creates a new DicomPixelData from a DICOM dataset.
// This function extracts all necessary image information from the dataset
// including pixel data, image dimensions, bit depth, and photometric interpretation.
//
// Example:
//
//	result, err := parser.ParseFile("image.dcm")
//	if err != nil {
//	    return err
//	}
//	pixelData, err := imaging.CreatePixelData(result.Dataset)
//	if err != nil {
//	    return err
//	}
//	image := imaging.NewDicomImage(pixelData)
//
//nolint:gocyclo // Complex function handling many DICOM variations
func CreatePixelData(ds *dataset.Dataset) (*DicomPixelData, error) {
	// Extract required tags
	rows, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil {
		return nil, fmt.Errorf("missing or invalid Rows tag: %w", err)
	}

	cols, err := ds.GetUInt16(tag.Columns, 0)
	if err != nil {
		return nil, fmt.Errorf("missing or invalid Columns tag: %w", err)
	}

	// Get pixel data element
	pixelDataElem, ok := ds.Get(tag.PixelData)
	if !ok {
		return nil, fmt.Errorf("missing Pixel Data tag")
	}

	// Get optional parameters with defaults
	bitsAllocated := ds.TryGetUInt16(tag.BitsAllocated, 0)
	if bitsAllocated == 0 {
		bitsAllocated = 16
	}

	bitsStored := ds.TryGetUInt16(tag.BitsStored, 0)
	if bitsStored == 0 {
		bitsStored = bitsAllocated
	}

	highBit := ds.TryGetUInt16(tag.HighBit, 0)
	if highBit == 0 {
		highBit = bitsStored - 1
	}

	samplesPerPixel := ds.TryGetUInt16(tag.SamplesPerPixel, 0)
	if samplesPerPixel == 0 {
		samplesPerPixel = 1
	}

	pixelRepr := ds.TryGetUInt16(tag.PixelRepresentation, 0)

	// Get number of frames
	// Note: NumberOfFrames has VR of IS (Integer String) per DICOM standard
	numberOfFrames := 1
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		numberOfFrames = int(nf)
	} else if nfStr, ok := ds.GetString(tag.NumberOfFrames); ok {
		// Try parsing as string (IS VR type)
		if parsed, err := strconv.Atoi(strings.TrimSpace(nfStr)); err == nil && parsed > 0 {
			numberOfFrames = parsed
		}
	}

	planarConfig := ds.TryGetUInt16(tag.PlanarConfiguration, 0)

	// Get photometric interpretation
	photoInterp, _ := ds.GetString(tag.PhotometricInterpretation)
	if photoInterp == "" {
		photoInterp = "MONOCHROME2"
	}

	pi, err := ParsePhotometricInterpretation(photoInterp)
	if err != nil {
		return nil, fmt.Errorf("invalid photometric interpretation %q: %w", photoInterp, err)
	}

	// Get transfer syntax UID
	transferSyntaxUID := "1.2.840.10008.1.2" // Default: Implicit VR Little Endian
	// Note: In a full implementation, this should be extracted from the file meta information

	// Get lossy compression information if present
	isLossy := false
	if lossyComp, ok := ds.GetString(tag.LossyImageCompression); ok && lossyComp == "01" {
		isLossy = true
	}

	lossyMethod := ""
	if method, ok := ds.GetString(tag.LossyImageCompressionMethod); ok {
		lossyMethod = method
	}

	lossyRatio := 0.0
	if ratioStr, ok := ds.GetString(tag.LossyImageCompressionRatio); ok {
		if ratio, err := strconv.ParseFloat(ratioStr, 64); err == nil {
			lossyRatio = ratio
		}
	}

	// Pixel padding (optional)
	var paddingVal *int32
	if pv, err := ds.GetInt32(tag.PixelPaddingValue, 0); err == nil {
		paddingVal = &pv
	}
	var paddingRange *int32
	if pr, err := ds.GetInt32(tag.PixelPaddingRangeLimit, 0); err == nil {
		paddingRange = &pr
	}

	// Create pixel data info
	info := &PixelDataInfo{
		Width:                     cols,
		Height:                    rows,
		NumberOfFrames:            numberOfFrames,
		BitsAllocated:             bitsAllocated,
		BitsStored:                bitsStored,
		HighBit:                   highBit,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       PixelRepresentation(pixelRepr),
		PlanarConfiguration:       PlanarConfiguration(planarConfig),
		PhotometricInterpretation: pi,
		TransferSyntaxUID:         transferSyntaxUID,
		IsLossy:                   isLossy,
		LossyCompressionMethod:    lossyMethod,
		LossyCompressionRatio:     lossyRatio,
		PixelPaddingValue:         paddingVal,
		PixelPaddingRangeLimit:    paddingRange,
	}

	// Create DICOM pixel data from bytes
	var pd *DicomPixelData
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		info.VRCode = "OB"
		info.Encapsulated = false
		data := elem.GetData()
		if len(data) == 0 {
			return nil, fmt.Errorf("pixel data is empty")
		}
		pd, err = NewDicomPixelDataFromBytes(info, data)
		if err != nil {
			return nil, err
		}
	case *element.OtherWord:
		info.VRCode = "OW"
		info.Encapsulated = false
		data := elem.GetData()
		if len(data) == 0 {
			return nil, fmt.Errorf("pixel data is empty")
		}
		pd, err = NewDicomPixelDataFromBytes(info, data)
		if err != nil {
			return nil, err
		}
	case *element.OtherByteFragment:
		info.VRCode = "OB"
		info.Encapsulated = true
		frames, ferr := framesFromFragments(elem.Fragments(), elem.OffsetTable(), numberOfFrames)
		if ferr != nil {
			return nil, ferr
		}
		pd, err = NewDicomPixelData(info)
		if err != nil {
			return nil, err
		}
		pd.frames = append(pd.frames, frames...)
		pd.basicOffsetTable = append(pd.basicOffsetTable, elem.OffsetTable()...)
		pd.Info.NumberOfFrames = len(pd.frames)
	case *element.OtherWordFragment:
		info.VRCode = "OW"
		info.Encapsulated = true
		frames, ferr := framesFromFragments(elem.Fragments(), elem.OffsetTable(), numberOfFrames)
		if ferr != nil {
			return nil, ferr
		}
		pd, err = NewDicomPixelData(info)
		if err != nil {
			return nil, err
		}
		pd.frames = append(pd.frames, frames...)
		pd.basicOffsetTable = append(pd.basicOffsetTable, elem.OffsetTable()...)
		pd.Info.NumberOfFrames = len(pd.frames)
	default:
		return nil, fmt.Errorf("unsupported pixel data element type: %T", pixelDataElem)
	}

	// Palette Color handling: convert to RGB if palette LUT present
	if pi.Value == "PALETTE COLOR" {
		if err := convertPaletteToRGB(ds, pd); err != nil {
			return nil, fmt.Errorf("palette conversion failed: %w", err)
		}
	}

	return pd, nil
}

type paletteLUT struct {
	first   int32
	entries []types.Color32
}

// convertPaletteToRGB loads palette LUT from dataset and converts frames to RGB using shared LUT Color32 type.
func convertPaletteToRGB(ds *dataset.Dataset, pd *DicomPixelData) error {
	lut, err := buildPaletteLUT(ds)
	if err != nil {
		return err
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return fmt.Errorf("unsupported BytesAllocated=%d for palette conversion", bytesPerSample)
	}

	for fi, frame := range pd.frames {
		pixelCount := len(frame) / bytesPerSample
		out := make([]byte, pixelCount*3)

		for idx, off := 0, 0; idx < pixelCount; idx, off = idx+1, off+bytesPerSample {
			var val int32
			if bytesPerSample == 1 {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int8(frame[off]))
				} else {
					val = int32(frame[off])
				}
			} else {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int16(binary.LittleEndian.Uint16(frame[off:])))
				} else {
					val = int32(binary.LittleEndian.Uint16(frame[off:]))
				}
			}

			idxLUT := int(val - lut.first)
			if idxLUT < 0 {
				idxLUT = 0
			}
			if idxLUT >= len(lut.entries) {
				idxLUT = len(lut.entries) - 1
			}

			color := lut.entries[idxLUT]
			base := idx * 3
			out[base] = color.R
			out[base+1] = color.G
			out[base+2] = color.B
		}

		pd.frames[fi] = out
	}

	// Update metadata to RGB
	pd.Info.PhotometricInterpretation = RGBPhotometric
	pd.Info.SamplesPerPixel = 3
	pd.Info.PlanarConfiguration = InterleavedPlanar
	pd.Info.BitsAllocated = 8
	pd.Info.BitsStored = 8
	pd.Info.HighBit = 7

	return nil
}

func buildPaletteLUT(ds *dataset.Dataset) (*paletteLUT, error) {
	// Enhanced Palette Color LUT Sequence (0028,140B)
	if seqElem, ok := ds.Get(tag.EnhancedPaletteColorLookupTableSequence); ok {
		if seq, ok2 := seqElem.(*dataset.Sequence); ok2 && seq.Count() > 0 {
			for i := 0; i < seq.Count(); i++ {
				if lut, err := buildPaletteLUTFromDataset(seq.GetItem(i)); err == nil {
					return lut, nil
				}
			}
		}
	}

	// Palette Color LUT Sequence (0048,0120)
	if seqElem, ok := ds.Get(tag.PaletteColorLookupTableSequence); ok {
		if seq, ok2 := seqElem.(*dataset.Sequence); ok2 && seq.Count() > 0 {
			for i := 0; i < seq.Count(); i++ {
				if lut, err := buildPaletteLUTFromDataset(seq.GetItem(i)); err == nil {
					return lut, nil
				}
			}
		}
	}

	// Fall back to top-level descriptors/data
	return buildPaletteLUTFromDataset(ds)
}

// buildPaletteLUTFromDataset builds palette LUT using descriptors/data in the provided dataset (no sequence recursion).
//
//nolint:gocyclo // Complex function handling palette LUT variations
func buildPaletteLUTFromDataset(ds *dataset.Dataset) (*paletteLUT, error) {
	// Descriptors
	rDesc, err := ds.GetUInt16(tag.RedPaletteColorLookupTableDescriptor, 0)
	if err != nil {
		return nil, fmt.Errorf("missing Red Palette LUT Descriptor: %w", err)
	}
	_, _ = ds.GetUInt16(tag.GreenPaletteColorLookupTableDescriptor, 0) // for validation, but not used directly
	_, _ = ds.GetUInt16(tag.BluePaletteColorLookupTableDescriptor, 0)

	rFirst := ds.TryGetUInt16(tag.RedPaletteColorLookupTableDescriptor, 1)
	gFirst := ds.TryGetUInt16(tag.GreenPaletteColorLookupTableDescriptor, 1)
	bFirst := ds.TryGetUInt16(tag.BluePaletteColorLookupTableDescriptor, 1)

	rBits := ds.TryGetUInt16(tag.RedPaletteColorLookupTableDescriptor, 2)
	gBits := ds.TryGetUInt16(tag.GreenPaletteColorLookupTableDescriptor, 2)
	bBits := ds.TryGetUInt16(tag.BluePaletteColorLookupTableDescriptor, 2)

	size := int(rDesc)
	if size == 0 {
		size = 65536
	}

	// use max bits among channels
	bits := int(rBits)
	if int(gBits) > bits {
		bits = int(gBits)
	}
	if int(bBits) > bits {
		bits = int(bBits)
	}
	if bits == 0 {
		bits = 8
	}

	loadLUT := func(dataTag *tag.Tag, expectedBits uint16) ([]uint16, error) {
		elem, ok := ds.Get(dataTag)
		if !ok {
			return nil, fmt.Errorf("missing palette data %s", dataTag)
		}
		var raw []byte
		switch v := elem.(type) {
		case *element.OtherByte:
			raw = v.GetData()
		case *element.OtherWord:
			raw = v.GetData()
		default:
			return nil, fmt.Errorf("unsupported palette data element type %T for %s", elem, dataTag)
		}

		out := make([]uint16, size)
		if expectedBits <= 8 {
			for i := 0; i < size && i < len(raw); i++ {
				out[i] = uint16(raw[i])
			}
		} else {
			for i := 0; i < size && (i*2+1) < len(raw); i++ {
				out[i] = binary.LittleEndian.Uint16(raw[i*2:])
			}
		}
		return out, nil
	}

	// Prefer standard LUT data; if missing, try segmented LUT data
	rLUT, err := loadLUT(tag.RedPaletteColorLookupTableData, rBits)
	if err != nil {
		if seg, ok := ds.Get(tag.SegmentedRedPaletteColorLookupTableData); ok {
			if ob, ok2 := seg.(*element.OtherByte); ok2 {
				rLUT, err = expandSegmentedLUT(ob.GetData(), size)
			} else if ow, ok2 := seg.(*element.OtherWord); ok2 {
				rLUT, err = expandSegmentedLUT(ow.GetData(), size)
			} else {
				err = fmt.Errorf("unsupported segmented palette element type %T", seg)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	gLUT, err := loadLUT(tag.GreenPaletteColorLookupTableData, gBits)
	if err != nil {
		if seg, ok := ds.Get(tag.SegmentedGreenPaletteColorLookupTableData); ok {
			if ob, ok2 := seg.(*element.OtherByte); ok2 {
				gLUT, err = expandSegmentedLUT(ob.GetData(), size)
			} else if ow, ok2 := seg.(*element.OtherWord); ok2 {
				gLUT, err = expandSegmentedLUT(ow.GetData(), size)
			} else {
				err = fmt.Errorf("unsupported segmented palette element type %T", seg)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	bLUT, err := loadLUT(tag.BluePaletteColorLookupTableData, bBits)
	if err != nil {
		if seg, ok := ds.Get(tag.SegmentedBluePaletteColorLookupTableData); ok {
			if ob, ok2 := seg.(*element.OtherByte); ok2 {
				bLUT, err = expandSegmentedLUT(ob.GetData(), size)
			} else if ow, ok2 := seg.(*element.OtherWord); ok2 {
				bLUT, err = expandSegmentedLUT(ow.GetData(), size)
			} else {
				err = fmt.Errorf("unsupported segmented palette element type %T", seg)
			}
		}
		if err != nil {
			return nil, err
		}
	}

	// Align sizes
	minLen := len(rLUT)
	if len(gLUT) < minLen {
		minLen = len(gLUT)
	}
	if len(bLUT) < minLen {
		minLen = len(bLUT)
	}
	rLUT = rLUT[:minLen]
	gLUT = gLUT[:minLen]
	bLUT = bLUT[:minLen]

	first := int32(rFirst)
	// if first differs among channels, use min
	if int32(gFirst) < first {
		first = int32(gFirst)
	}
	if int32(bFirst) < first {
		first = int32(bFirst)
	}

	return &paletteLUT{
		first:   first,
		entries: buildPaletteEntries(bits, rLUT, gLUT, bLUT),
	}, nil
}

func buildPaletteEntries(bits int, rLUT, gLUT, bLUT []uint16) []types.Color32 {
	shift := 0
	if bits > 8 {
		shift = bits - 8
	}

	entries := make([]types.Color32, len(rLUT))
	for i := 0; i < len(rLUT); i++ {
		entries[i] = types.Color32{
			A: 255,
			R: clampByte(int(rLUT[i] >> shift)),
			G: clampByte(int(gLUT[i] >> shift)),
			B: clampByte(int(bLUT[i] >> shift)),
		}
	}
	return entries
}

// expandSegmentedLUT expands DICOM segmented palette LUT data (Type 0/1 segments).
// Supports discrete and linear segments; skips unsupported types.
func expandSegmentedLUT(raw []byte, expectedSize int) ([]uint16, error) {
	var out []uint16
	for i := 0; i+1 < len(raw); {
		desc := binary.LittleEndian.Uint16(raw[i:])
		i += 2
		segType := desc >> 14
		count := int(desc & 0x3FFF)

		switch segType {
		case 0: // discrete: count+1 values follow
			n := count + 1
			for j := 0; j < n && i+1 < len(raw); j++ {
				val := binary.LittleEndian.Uint16(raw[i:])
				out = append(out, val)
				i += 2
			}
		case 1: // linear: count+1 values generated from start/end
			if i+3 >= len(raw) {
				return nil, fmt.Errorf("segmented LUT linear segment truncated")
			}
			start := binary.LittleEndian.Uint16(raw[i:])
			end := binary.LittleEndian.Uint16(raw[i+2:])
			i += 4
			n := count + 1
			for k := 0; k < n; k++ {
				val := uint16(float64(start) + float64(end-start)*float64(k)/float64(n-1))
				out = append(out, val)
			}
		default:
			// unsupported segment types -> skip safely
			return nil, fmt.Errorf("unsupported segmented LUT segment type %d", segType)
		}
		if expectedSize > 0 && len(out) >= expectedSize {
			break
		}
	}

	if expectedSize > 0 && len(out) < expectedSize {
		// pad if short
		for len(out) < expectedSize {
			out = append(out, out[len(out)-1])
		}
	}

	return out, nil
}

// framesFromFragments builds per-frame compressed data using fragments and an optional BOT.
// If offsetTable is present, it slices the concatenated fragments using offsets.
// Otherwise, it assumes one fragment per frame (best-effort fallback).
func framesFromFragments(fragments []buffer.ByteBuffer, offsetTable []uint32, frameCount int) ([][]byte, error) {
	if len(fragments) == 0 {
		return nil, fmt.Errorf("no fragments available")
	}

	for i, frag := range fragments {
		if len(frag.Data()) == 0 {
			return nil, fmt.Errorf("fragment %d is empty", i)
		}
	}

	if frameCount < 1 {
		frameCount = len(offsetTable)
	}
	if frameCount < 1 {
		frameCount = len(fragments)
	}
	if frameCount < 1 {
		frameCount = 1
	}

	// BOT present: slice concatenated stream by offsets.
	if len(offsetTable) > 0 {
		if frameCount != len(offsetTable) {
			return nil, fmt.Errorf("offset table frames mismatch: expected %d, got %d", frameCount, len(offsetTable))
		}
		var concat []byte
		for _, frag := range fragments {
			concat = append(concat, frag.Data()...)
		}

		var frames [][]byte
		for i := 0; i < frameCount; i++ {
			start := int(offsetTable[i])
			end := len(concat)
			if i+1 < len(offsetTable) {
				end = int(offsetTable[i+1])
			}
			if start < 0 || end < 0 || start > end || end > len(concat) {
				return nil, fmt.Errorf("invalid BOT slice for frame %d: start %d end %d total %d", i, start, end, len(concat))
			}
			if start == end {
				return nil, fmt.Errorf("frame %d derived from BOT is empty", i)
			}
			frames = append(frames, concat[start:end])
		}
		return frames, nil
	}

	// Fallback: require one fragment per frame and consistent sizes.
	if frameCount > len(fragments) {
		return nil, fmt.Errorf("frame count %d exceeds available fragments %d without BOT", frameCount, len(fragments))
	}
	framesToUse := frameCount
	var frames [][]byte
	var expectedSize int
	for i := 0; i < framesToUse; i++ {
		data := fragments[i].Data()
		if i == 0 {
			expectedSize = len(data)
		} else if len(data) != expectedSize {
			return nil, fmt.Errorf("fragment size mismatch at index %d: got %d, expected %d", i, len(data), expectedSize)
		}
		frames = append(frames, data)
	}
	return frames, nil
}

// buildFragmentSequence creates an OB fragment sequence from per-frame compressed data,
// populating the Basic Offset Table for multi-frame images.
// If an existing BOT is provided and matches frames length, it is used; otherwise BOT is rebuilt.
func buildFragmentSequence(frames [][]byte, existingBOT []uint32, bitsAllocated uint16) (element.Element, error) {
	if len(frames) == 0 {
		return nil, fmt.Errorf("no frame data provided for fragment sequence")
	}

	// Rebuild BOT if missing or length mismatch
	// According to DICOM standard, Basic Offset Table should contain at least one offset (0x00000000) for single-frame,
	// and all frame offsets for multi-frame images.
	var offsets []uint32
	useExisting := len(existingBOT) == len(frames)
	if useExisting {
		offsets = append(offsets, existingBOT...)
	} else {
		var runningOffset uint32
		for i, frame := range frames {
			offsets = append(offsets, runningOffset)
			if len(frame) > int(math.MaxUint32-runningOffset) {
				return nil, fmt.Errorf("fragment too large to represent in BOT at frame %d", i)
			}
			runningOffset += uint32(len(frame))
		}
	}

	// Choose OB/OW based on BitsAllocated (following fo-dicom behavior)
	// - BitsAllocated > 8: use OW (Other Word)
	// - BitsAllocated <= 8: use OB (Other Byte)
	if bitsAllocated > 8 {
		// Create OtherWordFragment
		owf := element.NewOtherWordFragment(tag.PixelData)
		for _, frame := range frames {
			owf.AddFragment(buffer.NewMemory(frame))
		}
		owf.SetOffsetTable(offsets)
		return owf, nil
	}

	// Create OtherByteFragment
	obf := element.NewOtherByteFragment(tag.PixelData)
	for _, frame := range frames {
		obf.AddFragment(buffer.NewMemory(frame))
	}
	obf.SetOffsetTable(offsets)
	return obf, nil
}
