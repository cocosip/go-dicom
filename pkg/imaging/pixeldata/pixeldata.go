// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/encapsulated"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
	"github.com/cocosip/go-dicom/pkg/imaging/internal/dicomlut"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/logging"
)

// Info contains metadata about DICOM pixel data.
type Info struct {
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
	PixelRepresentation pixel.Representation

	// Planar configuration
	PlanarConfiguration pixel.PlanarConfiguration

	// Photometric interpretation
	PhotometricInterpretation *pixel.PhotometricInterpretation

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
func (info *Info) BytesAllocated() int {
	return int((info.BitsAllocated-1)/8 + 1)
}

// UncompressedFrameSize calculates the uncompressed size of a single frame in bytes.
func (info *Info) UncompressedFrameSize() int {
	if info.BitsAllocated == 1 {
		return (int(info.Width)*int(info.Height)-1)/8 + 1
	}

	// Handle special case for YBR_FULL_422 with uneven width
	actualWidth := int(info.Width)
	if actualWidth%2 != 0 &&
		info.PhotometricInterpretation != nil &&
		(info.PhotometricInterpretation.Value == pixel.YbrFull422.Value ||
			info.PhotometricInterpretation.Value == pixel.YbrPartial422.Value ||
			info.PhotometricInterpretation.Value == pixel.YbrPartial420.Value) {
		actualWidth++
	}

	// Handle YBR_FULL_422 special case for uncompressed data
	if info.PhotometricInterpretation != nil &&
		(info.PhotometricInterpretation.Value == pixel.YbrFull422.Value ||
			info.PhotometricInterpretation.Value == pixel.YbrPartial422.Value ||
			info.PhotometricInterpretation.Value == pixel.YbrPartial420.Value) {
		// For uncompressed transfer syntaxes, chrominance channels are downsampled
		return info.BytesAllocated() * 2 * actualWidth * int(info.Height)
	}

	return info.BytesAllocated() * int(info.SamplesPerPixel) * actualWidth * int(info.Height)
}

// TotalUncompressedSize returns the total size of all frames uncompressed.
func (info *Info) TotalUncompressedSize() int {
	return info.UncompressedFrameSize() * info.NumberOfFrames
}

// Validate checks if the pixel data info is valid.
func (info *Info) Validate() error {
	if info == nil {
		return fmt.Errorf("pixel data info must not be nil")
	}
	var photometric pixel.PhotometricInterpretation
	if info.PhotometricInterpretation != nil {
		photometric = *info.PhotometricInterpretation
	}
	frameInfo := codec.FrameInfo{
		Width:                     info.Width,
		Height:                    info.Height,
		BitDepth:                  *pixel.NewBitDepth(info.BitsAllocated, info.BitsStored, info.HighBit, info.PixelRepresentation.IsSigned()),
		SamplesPerPixel:           info.SamplesPerPixel,
		PixelRepresentation:       info.PixelRepresentation,
		PlanarConfiguration:       info.PlanarConfiguration,
		PhotometricInterpretation: photometric,
	}
	if err := frameInfo.Validate(); err != nil {
		return err
	}
	if info.NumberOfFrames < 1 {
		return fmt.Errorf("number of frames must be at least 1")
	}
	if info.VRCode != "" && info.VRCode != "OB" && info.VRCode != "OW" {
		return fmt.Errorf("pixel data VR must be OB or OW, got %q", info.VRCode)
	}
	if info.PixelPaddingRangeLimit != nil && info.PixelPaddingValue == nil {
		return fmt.Errorf("pixel padding range limit requires pixel padding value")
	}

	return nil
}

// Data manages DICOM pixel data with support for multiple frames and codecs.
type Data struct {
	Info             *Info
	frames           [][]byte // Per-frame data (uncompressed for native; compressed for encapsulated)
	basicOffsetTable []uint32 // BOT for encapsulated data
}

var (
	_ codec.FrameSource = (*Data)(nil)
	_ codec.FrameSink   = (*Data)(nil)
)

// New creates a new Data instance.
func New(info *Info) (*Data, error) {
	if info == nil {
		return nil, fmt.Errorf("pixel data info must not be nil")
	}
	if err := info.Validate(); err != nil {
		return nil, fmt.Errorf("invalid pixel data info: %w", err)
	}

	return &Data{
		Info:             info,
		frames:           make([][]byte, 0, info.NumberOfFrames),
		basicOffsetTable: nil,
	}, nil
}

// NewFromBytes creates Data from raw pixel bytes.
// The data is assumed to contain all frames concatenated.
func NewFromBytes(info *Info, data []byte) (*Data, error) {
	if info == nil {
		return nil, fmt.Errorf("pixel data info must not be nil")
	}
	ownedInfo := *info
	ownedInfo.Encapsulated = false
	if ownedInfo.VRCode == "" {
		ownedInfo.VRCode = nativePixelDataVR(&ownedInfo)
	}
	pd, err := New(&ownedInfo)
	if err != nil {
		return nil, err
	}
	if pixelDataNeedsByteSwap(&ownedInfo) {
		data = swapPixelDataBytes(data, &ownedInfo)
	}

	frameSize := ownedInfo.UncompressedFrameSize()
	expectedSize := frameSize * ownedInfo.NumberOfFrames

	if len(data) < expectedSize {
		return nil, fmt.Errorf("insufficient data: got %d bytes, expected at least %d bytes",
			len(data), expectedSize)
	}

	// Split data into frames
	for i := 0; i < ownedInfo.NumberOfFrames; i++ {
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

// Frame returns an independently owned copy of the specified frame (0-indexed).
func (pd *Data) Frame(ctx context.Context, frameIndex int) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if frameIndex < 0 || frameIndex >= len(pd.frames) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(pd.frames))
	}
	frame := make([]byte, len(pd.frames[frameIndex]))
	if err := copyWithContext(ctx, frame, pd.frames[frameIndex]); err != nil {
		return nil, err
	}
	return frame, nil
}

// Sample returns one decoded native pixel sample. Coordinates and frame and
// sample indexes are zero-based. Encapsulated data must be decoded first.
func (pd *Data) Sample(frame, x, y, sample int) (int64, error) {
	if pd == nil || pd.Info == nil {
		return 0, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return 0, fmt.Errorf("cannot read a scalar sample from encapsulated pixel data")
	}
	if frame < 0 || frame >= len(pd.frames) {
		return 0, fmt.Errorf("frame index %d out of range [0, %d)", frame, len(pd.frames))
	}
	if x < 0 || x >= int(pd.Info.Width) {
		return 0, fmt.Errorf("x coordinate %d out of range [0, %d)", x, pd.Info.Width)
	}
	if y < 0 || y >= int(pd.Info.Height) {
		return 0, fmt.Errorf("y coordinate %d out of range [0, %d)", y, pd.Info.Height)
	}
	if sample < 0 || sample >= int(pd.Info.SamplesPerPixel) {
		return 0, fmt.Errorf("sample index %d out of range [0, %d)", sample, pd.Info.SamplesPerPixel)
	}

	pixelIndex := y*int(pd.Info.Width) + x
	sampleIndex := pixelIndex*int(pd.Info.SamplesPerPixel) + sample
	if pd.Info.PlanarConfiguration == pixel.PlanarPlanar && pd.Info.SamplesPerPixel > 1 {
		sampleIndex = sample*int(pd.Info.Width)*int(pd.Info.Height) + pixelIndex
	}
	offset := sampleIndex * pd.Info.BytesAllocated()
	if pd.Info.BitsAllocated == 1 {
		offset = sampleIndex
	}
	value, ok := decodePixelSampleLE(pd.frames[frame], offset, pd.Info)
	if !ok {
		return 0, fmt.Errorf("cannot decode sample at byte offset %d with BitsAllocated=%d", offset, pd.Info.BitsAllocated)
	}
	return value, nil
}

// IsPaddingSample reports whether value is inside the inclusive DICOM pixel
// padding interval. A reversed Pixel Padding Range Limit is normalized.
func (pd *Data) IsPaddingSample(value int64) bool {
	if pd == nil {
		return false
	}
	minimum, maximum, ok := pixelPaddingRange(pd.Info)
	if !ok {
		return false
	}
	return value >= minimum && value <= maximum
}

// CalculateOptimalWindow computes optimal window center/width from pixel data
// by sampling pixel values and finding min/max range
func (pd *Data) CalculateOptimalWindow() (center, width float64) {
	if len(pd.frames) == 0 {
		return 0, 256 // Default fallback
	}

	// Sample first frame for window calculation
	pixelData := pd.frames[0]
	bytesPerPixel := pd.Info.BytesAllocated()
	pixelCount := frameSampleCount(pixelData, pd.Info)

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
		pixelIndex := sampleOffset(i, pd.Info)
		if pd.Info.BitsAllocated != 1 && pixelIndex+bytesPerPixel > len(pixelData) {
			break
		}

		val, ok := decodePixelSampleLE(pixelData, pixelIndex, pd.Info)
		if !ok {
			continue
		}
		pixelValue := float64(val)

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
func (pd *Data) AddFrame(ctx context.Context, frameData []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
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
		if err := copyWithContext(ctx, frame, frameData[:expectedSize]); err != nil {
			return err
		}
		pd.frames = append(pd.frames, frame)
	} else {
		// For encapsulated data, just copy the entire frame as-is
		frame := make([]byte, len(frameData))
		if err := copyWithContext(ctx, frame, frameData); err != nil {
			return err
		}
		pd.frames = append(pd.frames, frame)
	}

	// Update frame count
	pd.Info.NumberOfFrames = len(pd.frames)

	return nil
}

// AllFrames returns all pixel data as a single byte slice.
func (pd *Data) AllFrames() []byte {
	if pd == nil {
		return nil
	}
	totalSize := 0
	for _, frame := range pd.frames {
		totalSize += len(frame)
	}
	result := make([]byte, totalSize)

	offset := 0
	for _, frame := range pd.frames {
		copy(result[offset:], frame)
		offset += len(frame)
	}

	return result
}

// FrameCount returns the number of frames in the pixel data.
func (pd *Data) FrameCount() int {
	return len(pd.frames)
}

// Clone returns an independent copy of the pixel metadata, frames, and offset table.
func (pd *Data) Clone() *Data {
	if pd == nil {
		return nil
	}
	clone := &Data{
		frames:           make([][]byte, len(pd.frames)),
		basicOffsetTable: append([]uint32(nil), pd.basicOffsetTable...),
	}
	if pd.Info != nil {
		info := *pd.Info
		if pd.Info.PhotometricInterpretation != nil {
			photometric := *pd.Info.PhotometricInterpretation
			info.PhotometricInterpretation = &photometric
		}
		if pd.Info.PixelPaddingValue != nil {
			padding := *pd.Info.PixelPaddingValue
			info.PixelPaddingValue = &padding
		}
		if pd.Info.PixelPaddingRangeLimit != nil {
			limit := *pd.Info.PixelPaddingRangeLimit
			info.PixelPaddingRangeLimit = &limit
		}
		clone.Info = &info
	}
	for index, frame := range pd.frames {
		clone.frames[index] = append([]byte(nil), frame...)
	}
	return clone
}

// EnsureInterleaved converts planar configuration 1 to interleaved (0) for multi-sample pixels.
// Only applies to uncompressed data; encapsulated data must be decoded first.
func (pd *Data) EnsureInterleaved() error {
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
		converted, err := colorconv.PlanarToInterleaved(frame, int(pd.Info.SamplesPerPixel), bytesPerSample)
		if err != nil {
			return fmt.Errorf("frame %d planar->interleaved failed: %w", idx, err)
		}
		convertedFrames[idx] = converted
	}
	pd.frames = convertedFrames

	pd.Info.PlanarConfiguration = pixel.InterleavedPlanar
	return nil
}

// ConvertMonochrome1ToMonochrome2 inverts grayscale for MONOCHROME1 data.
// Only applies to uncompressed single-sample images.
func (pd *Data) ConvertMonochrome1ToMonochrome2() error {
	if pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return fmt.Errorf("cannot convert photometric on encapsulated data")
	}
	if pd.Info.PhotometricInterpretation == nil || pd.Info.PhotometricInterpretation.Value != pixel.Monochrome1.Value {
		return nil
	}
	if pd.Info.SamplesPerPixel != 1 {
		return fmt.Errorf("%s expected SamplesPerPixel=1, got %d", pixel.Monochrome1.Value, pd.Info.SamplesPerPixel)
	}

	bytesPerSample := pd.Info.BytesAllocated()
	convertedFrames := make([][]byte, len(pd.frames))
	for fi, frame := range pd.frames {
		converted, err := colorconv.ConvertMono1ToMono2(frame, pd.Info.BitsStored, bytesPerSample, pd.Info.PixelRepresentation == pixel.SignedPixels)
		if err != nil {
			return fmt.Errorf("frame %d mono1->mono2 failed: %w", fi, err)
		}
		convertedFrames[fi] = converted
	}
	pd.frames = convertedFrames

	pd.Info.PhotometricInterpretation = pixel.Monochrome2
	return nil
}

// ConvertYBRToRGB converts uncompressed YBR data to interleaved RGB (Photometric=RGB).
// Supports YBR_FULL and YBR_FULL_422 with 8-bit samples; other variants are not handled here.
func (pd *Data) ConvertYBRToRGB() error {
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
	case pixel.YbrFull.Value:
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("YBR_FULL expected SamplesPerPixel=3, got %d", pd.Info.SamplesPerPixel)
		}
		for i, frame := range pd.frames {
			converted, err := colorconv.ConvertYBRFullToRGB(frame)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case pixel.YbrFull422.Value:
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("YBR_FULL_422 expected SamplesPerPixel=3, got %d", pd.Info.SamplesPerPixel)
		}
		width := int(pd.Info.Width)
		if width == 0 {
			return fmt.Errorf("YBR_FULL_422 requires valid width")
		}
		for i, frame := range pd.frames {
			converted, err := colorconv.ConvertYBRFull422ToRGB(frame, width)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case pixel.YbrPartial422.Value:
		if pd.Info.SamplesPerPixel != 3 {
			return fmt.Errorf("%s expected SamplesPerPixel=3, got %d", pixel.YbrPartial422.Value, pd.Info.SamplesPerPixel)
		}
		width := int(pd.Info.Width)
		if width == 0 {
			return fmt.Errorf("%s requires valid width", pixel.YbrPartial422.Value)
		}
		for i, frame := range pd.frames {
			converted, err := colorconv.ConvertYBRPartial422ToRGB(frame, width)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case pixel.YbrIct.Value:
		bytesPerSample := pd.Info.BytesAllocated()
		for i, frame := range pd.frames {
			converted, err := colorconv.ConvertYBRICTToRGB(frame, bytesPerSample*8)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	case pixel.YbrRct.Value:
		bytesPerSample := pd.Info.BytesAllocated()
		for i, frame := range pd.frames {
			converted, err := colorconv.ConvertYBRRCTToRGB(frame, bytesPerSample*8)
			if err != nil {
				return fmt.Errorf("frame %d: %w", i, err)
			}
			pd.frames[i] = converted
		}
	default:
		return fmt.Errorf("photometric %s not supported for YBR->RGB conversion", pd.Info.PhotometricInterpretation.Value)
	}

	pd.Info.PhotometricInterpretation = pixel.RGBPhotometric
	pd.Info.PlanarConfiguration = pixel.InterleavedPlanar
	pd.Info.SamplesPerPixel = 3
	return nil
}

// WindowTo8bit applies a VOI window (center/width) to pixel data and returns 8-bit frames.
// Only supports uncompressed data with BitsAllocated 8 or 16. For multi-frame, returns one []byte per frame.
// If ignorePadding is true and PixelPaddingValue/(RangeLimit) is set, padding samples are forced to 0.
func (pd *Data) WindowTo8bit(center, width float64, ignorePadding bool) ([][]byte, error) {
	return applyWindowTo8bit(pd, center, width, ignorePadding)
}

// MinMax returns the minimum and maximum sample values across all frames.
// If ignorePadding is true and PixelPaddingValue/(RangeLimit) is set, padding samples are skipped.
func (pd *Data) MinMax(ignorePadding bool) (minVal float64, maxVal float64, err error) {
	return minMaxSamples(pd, ignorePadding)
}

// MaskPadding returns a copy of frames where padding samples are set to 0 and a mask per frame (true = padding).
// Only applies to uncompressed data; encapsulated must be decoded first.
func (pd *Data) MaskPadding() (frames [][]byte, masks [][]bool, err error) {
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

	padMin, padMax, _ := pixelPaddingRange(pd.Info)

	for _, frame := range pd.frames {
		out := make([]byte, len(frame))
		copy(out, frame)
		mask := make([]bool, frameSampleCount(frame, pd.Info))

		for idx := 0; idx < len(mask); idx++ {
			off := sampleOffset(idx, pd.Info)
			if pd.Info.BitsAllocated != 1 && off+bytesPerSample > len(frame) {
				break
			}
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			if val >= padMin && val <= padMax {
				mask[idx] = true
				// zero out
				if pd.Info.BitsAllocated == 1 {
					out[off/8] &^= 1 << uint(off%8)
				} else {
					for b := 0; b < bytesPerSample; b++ {
						out[off+b] = 0
					}
				}
			}
		}

		frames = append(frames, out)
		masks = append(masks, mask)
	}

	return frames, masks, nil
}

// WindowOrLUTTo8bit applies VOI LUT if present, otherwise window.
func (pd *Data) WindowOrLUTTo8bit(ds *dataset.Dataset, center, width float64, ignorePadding bool) ([][]byte, error) {
	result := make([][]byte, len(pd.frames))
	for frameIndex, frame := range pd.frames {
		functional := dicomlut.FunctionalGroupValues(ds, frameIndex)
		modality, err := modalityTransformForDatasets(functional, ds, pd)
		if err != nil {
			return nil, fmt.Errorf("frame %d: %w", frameIndex, err)
		}
		frameData := &Data{Info: pd.Info, frames: [][]byte{frame}}
		var converted [][]byte
		if voiSource := datasetContaining(functional, ds, tag.VOILUTSequence); voiSource != nil {
			converted, err = applyVOILUTWithModality(frameData, voiSource, center, width, ignorePadding, modality)
		} else {
			converted, err = applyWindowTo8bitWithModality(frameData, center, width, ignorePadding, modality)
		}
		if err != nil {
			return nil, fmt.Errorf("frame %d: %w", frameIndex, err)
		}
		result[frameIndex] = converted[0]
	}
	return result, nil
}

func modalityTransformForDatasets(primary, fallback *dataset.Dataset, pd *Data) (lut.LUT, error) {
	if primary == nil && fallback == nil {
		return nil, nil
	}
	signed := pd != nil && pd.Info != nil && pd.Info.PixelRepresentation == pixel.SignedPixels
	// Select one precedence source for the complete modality transform. A
	// Functional Group transform overrides top-level attributes instead of
	// being combined with the other representation.
	source := datasetContaining(
		primary,
		nil,
		tag.ModalityLUTSequence,
		tag.RescaleSlope,
		tag.RescaleIntercept,
	)
	if source == nil {
		source = datasetContaining(
			fallback,
			nil,
			tag.ModalityLUTSequence,
			tag.RescaleSlope,
			tag.RescaleIntercept,
		)
	}
	if source == nil {
		return nil, nil
	}
	if source.Contains(tag.ModalityLUTSequence) {
		if source.Contains(tag.RescaleSlope) || source.Contains(tag.RescaleIntercept) {
			return nil, fmt.Errorf("modality LUT sequence cannot coexist with rescale slope/intercept")
		}
		return ModalityLUT(source, signed)
	}
	hasSlope, hasIntercept := source.Contains(tag.RescaleSlope), source.Contains(tag.RescaleIntercept)
	if !hasSlope && !hasIntercept {
		return nil, nil
	}
	if !hasSlope || !hasIntercept {
		return nil, fmt.Errorf("rescale slope and rescale intercept must be present together")
	}
	if source.Contains(tag.RescaleType) {
		if _, err := dicomlut.RequiredLongString(source, tag.RescaleType, "Rescale Type"); err != nil {
			return nil, err
		}
	}
	readDecimal := func(t *tag.Tag, name string) (float64, error) {
		values, exists := source.GetStrings(t)
		if !exists {
			return 0, fmt.Errorf("read %s: value is missing", name)
		}
		if len(values) != 1 {
			return 0, fmt.Errorf("read %s: want one DS value, got %d", name, len(values))
		}
		value, err := strconv.ParseFloat(strings.TrimSpace(values[0]), 64)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", name, err)
		}
		return value, nil
	}
	slope, err := readDecimal(tag.RescaleSlope, "Rescale Slope")
	if err != nil {
		return nil, err
	}
	intercept, err := readDecimal(tag.RescaleIntercept, "Rescale Intercept")
	if err != nil {
		return nil, err
	}
	if slope == 0 {
		return nil, fmt.Errorf("rescale slope must not be zero")
	}
	return lut.NewModalityRescaleLUT(slope, intercept, math.NaN(), math.NaN()), nil
}

func datasetContaining(primary, fallback *dataset.Dataset, tags ...*tag.Tag) *dataset.Dataset {
	for _, ds := range []*dataset.Dataset{primary, fallback} {
		if ds == nil {
			continue
		}
		for _, t := range tags {
			if ds.Contains(t) {
				return ds
			}
		}
	}
	return nil
}

// ToElement builds a DICOM pixel data element (OB/OW or encapsulated fragment) from the current frames.
// For encapsulated data, it emits an OB fragment sequence with Basic Offset Table.
func (pd *Data) ToElement() (element.Element, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}

	// Encapsulated pixel data is always encoded as OB per DICOM Part 5.
	if pd.Info.Encapsulated {
		return buildFragmentSequence(pd.frames, pd.basicOffsetTable, pd.Info.BitsAllocated)
	}

	// Uncompressed: concatenate frames
	all := pd.AllFrames()
	if len(all) == 0 {
		return nil, fmt.Errorf("pixel data is empty")
	}

	if nativePixelDataVR(pd.Info) == "OW" {
		// Frames are normalized to little-endian internally. Restore the
		// requested native transfer syntax when materializing the element.
		if pixelDataNeedsByteSwap(pd.Info) {
			all = swapPixelDataBytes(all, pd.Info)
		}
		return element.NewOtherWord(tag.PixelData, all), nil
	}
	return element.NewOtherByte(tag.PixelData, all), nil
}

func nativePixelDataVR(info *Info) string {
	if info.TransferSyntaxUID == transfer.ImplicitVRLittleEndian.UID().UID() || info.BitsAllocated > 8 || info.VRCode == "OW" {
		return "OW"
	}
	return "OB"
}

// Encapsulated returns true if pixel data is encapsulated.
func (pd *Data) Encapsulated() bool {
	return pd.Info != nil && pd.Info.Encapsulated
}

// BasicOffsetTable returns the BOT for encapsulated data.
func (pd *Data) BasicOffsetTable() []uint32 {
	return append([]uint32(nil), pd.basicOffsetTable...)
}

// FrameInfo returns frame metadata for codec operations.
func (pd *Data) FrameInfo() codec.FrameInfo {
	if pd == nil || pd.Info == nil {
		return codec.FrameInfo{}
	}
	var photometric pixel.PhotometricInterpretation
	if pd.Info.PhotometricInterpretation != nil {
		photometric = *pd.Info.PhotometricInterpretation
	}

	return codec.FrameInfo{
		Width:                     pd.Info.Width,
		Height:                    pd.Info.Height,
		BitDepth:                  *pixel.NewBitDepth(pd.Info.BitsAllocated, pd.Info.BitsStored, pd.Info.HighBit, pd.Info.PixelRepresentation.IsSigned()),
		SamplesPerPixel:           pd.Info.SamplesPerPixel,
		PixelRepresentation:       pd.Info.PixelRepresentation,
		PlanarConfiguration:       pd.Info.PlanarConfiguration,
		PhotometricInterpretation: photometric,
	}
}

// SetFrameInfo validates and applies codec output metadata.
func (pd *Data) SetFrameInfo(info codec.FrameInfo) error {
	if pd == nil || pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if err := info.Validate(); err != nil {
		return fmt.Errorf("invalid frame info: %w", err)
	}
	photometric := info.PhotometricInterpretation
	next := *pd.Info
	next.Width = info.Width
	next.Height = info.Height
	next.BitsAllocated = info.BitDepth.BitsAllocated
	next.BitsStored = info.BitDepth.BitsStored
	next.HighBit = info.BitDepth.HighBit
	next.SamplesPerPixel = info.SamplesPerPixel
	next.PixelRepresentation = info.PixelRepresentation
	next.PlanarConfiguration = info.PlanarConfiguration
	next.PhotometricInterpretation = &photometric
	if err := next.Validate(); err != nil {
		return fmt.Errorf("invalid frame info: %w", err)
	}
	*pd.Info = next
	return nil
}

// Encode encodes the pixel data using the specified codec and returns a new Data.
func (pd *Data) Encode(ctx context.Context, c codec.Codec, params codec.Parameters) (*Data, error) {
	if pd == nil || pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("codec must not be nil")
	}

	// Create new pixel data info for encoded data
	newInfo := &Info{
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

	newPD, err := New(newInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pixel data: %w", err)
	}

	ownedParams, err := codec.PrepareParameters(c, params)
	if err != nil {
		return nil, err
	}
	if err := c.Encode(ctx, pd, newPD, ownedParams); err != nil {
		return nil, fmt.Errorf("failed to encode pixel data: %w", err)
	}

	return newPD, nil
}

// Decode decodes the pixel data using the specified codec and returns a new Data.
func (pd *Data) Decode(ctx context.Context, c codec.Codec, params codec.Parameters) (*Data, error) {
	if pd == nil || pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("codec must not be nil")
	}

	// Create new pixel data info for decoded data
	newInfo := &Info{
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
		VRCode:                    "",
		Encapsulated:              false, // Decoded data is not encapsulated
		TransferSyntaxUID:         transfer.ExplicitVRLittleEndian.UID().UID(),
		IsLossy:                   pd.Info.IsLossy,
		LossyCompressionMethod:    pd.Info.LossyCompressionMethod,
		LossyCompressionRatio:     pd.Info.LossyCompressionRatio,
		PixelPaddingValue:         pd.Info.PixelPaddingValue,
		PixelPaddingRangeLimit:    pd.Info.PixelPaddingRangeLimit,
	}
	newInfo.VRCode = nativePixelDataVR(newInfo)

	newPD, err := New(newInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pixel data: %w", err)
	}

	ownedParams, err := codec.PrepareParameters(c, params)
	if err != nil {
		return nil, err
	}
	if err := c.Decode(ctx, pd, newPD, ownedParams); err != nil {
		return nil, fmt.Errorf("failed to decode pixel data: %w", err)
	}

	return newPD, nil
}

// FromDataset creates a new Data from a DICOM dataset.
// This function extracts all necessary image information from the dataset
// including pixel data, image dimensions, bit depth, and photometric interpretation.
//
// Example:
//
//	result, err := parser.ParseFile("image.dcm")
//	if err != nil {
//	    return err
//	}
//	pixelData, err := pixeldata.FromDataset(result.Dataset)
//	if err != nil {
//	    return err
//	}
//	image := imaging.NewDicomImage(pixelData)
//
//nolint:gocyclo // Complex function handling many DICOM variations
func FromDataset(ds *dataset.Dataset) (*Data, error) {
	return FromDatasetWithOptions(ds)
}

// FromDatasetWithOptions creates pixel data from a DICOM dataset using the
// requested VR compatibility policy.
//
//nolint:gocyclo // Complex function handling many DICOM variations
func FromDatasetWithOptions(ds *dataset.Dataset, options ...ReadOption) (*Data, error) {
	config := readConfig{vrMode: PixelDataCompatible}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}
	if config.vrMode != PixelDataCompatible && config.vrMode != PixelDataStandard {
		return nil, fmt.Errorf("unsupported Pixel Data VR mode: %d", config.vrMode)
	}
	if config.lutVRMode != LUTCompatible && config.lutVRMode != LUTStandard {
		return nil, fmt.Errorf("unsupported LUT VR mode: %d", config.lutVRMode)
	}
	if ds == nil {
		return nil, fmt.Errorf("dataset cannot be nil")
	}

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

	numberOfFrames := frameCountFromDataset(ds)

	planarConfig := ds.TryGetUInt16(tag.PlanarConfiguration, 0)

	// Get photometric interpretation
	photoInterp, _ := ds.GetString(tag.PhotometricInterpretation)
	if photoInterp == "" {
		photoInterp = pixel.Monochrome2.Value
	}

	pi, err := pixel.ParsePhotometricInterpretation(photoInterp)
	if err != nil {
		return nil, fmt.Errorf("invalid photometric interpretation %q: %w", photoInterp, err)
	}

	// Get transfer syntax UID from dataset
	// Priority: 1) InternalTransferSyntax 2) TransferSyntaxUID tag 3) Default
	transferSyntaxUID := transfer.ExplicitVRLittleEndian.UID().UID()
	transferSyntaxExplicitVR := true
	if ts := ds.InternalTransferSyntax(); ts != nil {
		transferSyntaxUID = ts.UID().UID()
		transferSyntaxExplicitVR = ts.IsExplicitVR()
	} else if tsUID, ok := ds.GetString(tag.TransferSyntaxUID); ok {
		transferSyntaxUID = tsUID
		if ts, parseErr := transfer.Parse(tsUID); parseErr == nil {
			transferSyntaxExplicitVR = ts.IsExplicitVR()
		}
	}

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
	if pv, err := dicomlut.ShortValue(ds, tag.PixelPaddingValue); err == nil {
		value := int32(pv)
		paddingVal = &value
	}
	var paddingRange *int32
	if pr, err := dicomlut.ShortValue(ds, tag.PixelPaddingRangeLimit); err == nil {
		value := int32(pr)
		paddingRange = &value
	}

	// Create pixel data info
	info := &Info{
		Width:                     cols,
		Height:                    rows,
		NumberOfFrames:            numberOfFrames,
		BitsAllocated:             bitsAllocated,
		BitsStored:                bitsStored,
		HighBit:                   highBit,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       pixel.Representation(pixelRepr),
		PlanarConfiguration:       pixel.PlanarConfiguration(planarConfig),
		PhotometricInterpretation: pi,
		TransferSyntaxUID:         transferSyntaxUID,
		IsLossy:                   isLossy,
		LossyCompressionMethod:    lossyMethod,
		LossyCompressionRatio:     lossyRatio,
		PixelPaddingValue:         paddingVal,
		PixelPaddingRangeLimit:    paddingRange,
	}

	// Create DICOM pixel data from bytes
	var pd *Data
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		var violation string
		switch {
		case !transferSyntaxExplicitVR:
			violation = "implicit VR native Pixel Data uses OB; DICOM PS3.5 requires OW"
		case bitsAllocated > 8:
			violation = fmt.Sprintf("native Pixel Data uses OB with Bits Allocated=%d; DICOM PS3.5 requires OW", bitsAllocated)
		}
		if violation != "" {
			if config.vrMode == PixelDataStandard {
				return nil, fmt.Errorf("%s", violation)
			}
			logging.Emit(context.Background(), logging.Record{
				Level:     slog.LevelWarn,
				Component: "imaging.pixeldata",
				Event:     "nonstandard_native_pixel_data_vr",
				Message:   "accepting non-standard native OB Pixel Data in compatibility mode",
				Attrs: []slog.Attr{
					slog.String("transfer_syntax", transferSyntaxUID),
					slog.Int("bits_allocated", int(bitsAllocated)),
				},
			})
		}
		info.VRCode = "OB"
		info.Encapsulated = false
		data := elem.GetData()
		if len(data) == 0 {
			return nil, fmt.Errorf("pixel data is empty")
		}
		pd, err = NewFromBytes(info, data)
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
		pd, err = NewFromBytes(info, data)
		if err != nil {
			return nil, err
		}
	case *element.OtherByteFragment:
		info.VRCode = "OB"
		info.Encapsulated = true
		encapsulatedData, ferr := encapsulated.OpenDataset(ds)
		if ferr != nil {
			return nil, ferr
		}
		frames, ferr := encapsulatedData.Frames(numberOfFrames)
		if ferr != nil {
			return nil, ferr
		}
		stripFramePadding(frames)
		pd, err = New(info)
		if err != nil {
			return nil, err
		}
		pd.frames = append(pd.frames, frames...)
		pd.basicOffsetTable = append(pd.basicOffsetTable, elem.OffsetTable()...)
		pd.Info.NumberOfFrames = len(pd.frames)
	case *element.OtherWordFragment:
		if config.vrMode == PixelDataStandard {
			return nil, fmt.Errorf("encapsulated Pixel Data uses OW; DICOM PS3.5 Section 8.2 requires OB")
		}
		logging.Emit(context.Background(), logging.Record{
			Level:     slog.LevelWarn,
			Component: "imaging.pixeldata",
			Event:     "nonstandard_encapsulated_pixel_data_vr",
			Message:   "accepting non-standard encapsulated OW Pixel Data in compatibility mode",
		})
		info.VRCode = "OW"
		info.Encapsulated = true
		encapsulatedData, ferr := encapsulated.OpenDataset(ds)
		if ferr != nil {
			return nil, ferr
		}
		frames, ferr := encapsulatedData.Frames(numberOfFrames)
		if ferr != nil {
			return nil, ferr
		}
		stripFramePadding(frames)
		pd, err = New(info)
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
	if pi.Value == pixel.PaletteColor.Value && !pd.Info.Encapsulated {
		if err := ConvertPaletteToRGBWithVRMode(ds, pd, config.lutVRMode); err != nil {
			return nil, fmt.Errorf("palette conversion failed: %w", err)
		}
	}

	return pd, nil
}

type paletteLUT struct {
	first    int32
	entries  []colorconv.Color32
	hasAlpha bool
}

// ConvertPaletteToRGB loads a Dataset palette and converts frames to RGB or RGBA.
func ConvertPaletteToRGB(ds *dataset.Dataset, pd *Data) error {
	return ConvertPaletteToRGBWithVRMode(ds, pd, LUTCompatible)
}

// ConvertPaletteToRGBWithVRMode loads a Dataset palette and converts frames to
// RGB or RGBA using the requested LUT Data VR policy.
func ConvertPaletteToRGBWithVRMode(ds *dataset.Dataset, pd *Data, mode LUTVRMode) error {
	lut, err := buildPaletteLUTWithVRMode(ds, mode)
	if err != nil {
		return err
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return fmt.Errorf("unsupported BytesAllocated=%d for palette conversion", bytesPerSample)
	}

	for fi, frame := range pd.frames {
		pixelCount := frameSampleCount(frame, pd.Info)
		channels := 3
		if lut.hasAlpha {
			channels = 4
		}
		out := make([]byte, pixelCount*channels)

		for idx := 0; idx < pixelCount; idx++ {
			off := sampleOffset(idx, pd.Info)
			if pd.Info.BitsAllocated != 1 && off+bytesPerSample > len(frame) {
				break
			}
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			idxLUT := int(val - int64(lut.first))
			if idxLUT < 0 {
				idxLUT = 0
			}
			if idxLUT >= len(lut.entries) {
				idxLUT = len(lut.entries) - 1
			}

			color := lut.entries[idxLUT]
			base := idx * channels
			out[base] = color.R
			out[base+1] = color.G
			out[base+2] = color.B
			if lut.hasAlpha {
				out[base+3] = color.A
			}
		}

		pd.frames[fi] = out
	}

	// Update metadata to RGB
	pd.Info.PhotometricInterpretation = pixel.RGBPhotometric
	pd.Info.SamplesPerPixel = uint16(3)
	if lut.hasAlpha {
		pd.Info.SamplesPerPixel = 4
	}
	pd.Info.PlanarConfiguration = pixel.InterleavedPlanar
	pd.Info.BitsAllocated = 8
	pd.Info.BitsStored = 8
	pd.Info.HighBit = 7

	return nil
}

func buildPaletteLUT(ds *dataset.Dataset) (*paletteLUT, error) {
	return buildPaletteLUTWithVRMode(ds, LUTCompatible)
}

func buildPaletteLUTWithVRMode(ds *dataset.Dataset, mode LUTVRMode) (*paletteLUT, error) {
	if mode != LUTCompatible && mode != LUTStandard {
		return nil, fmt.Errorf("unsupported LUT VR mode: %d", mode)
	}
	byteOrder := dicomlut.ByteOrder(ds)
	signed := ds.TryGetUInt16(tag.PixelRepresentation, 0) == uint16(pixel.SignedPixels)
	for _, sequenceTag := range []*tag.Tag{
		tag.EnhancedPaletteColorLookupTableSequence,
		tag.PaletteColorLookupTableSequence,
	} {
		lut, present, err := buildPaletteLUTFromSequence(ds, sequenceTag, byteOrder, signed, mode)
		if err != nil {
			return nil, err
		}
		if present {
			return lut, nil
		}
	}

	// Fall back to top-level descriptors/data
	return buildPaletteLUTFromDataset(ds, byteOrder, signed, mode)
}

// PaletteColors returns the first mapped value and an owned copy of Dataset palette colors.
func PaletteColors(ds *dataset.Dataset) (int32, []colorconv.Color32, error) {
	palette, err := buildPaletteLUT(ds)
	if err != nil {
		return 0, nil, err
	}
	return palette.first, append([]colorconv.Color32(nil), palette.entries...), nil
}

func buildPaletteLUTFromSequence(
	ds *dataset.Dataset,
	sequenceTag *tag.Tag,
	byteOrder binary.ByteOrder,
	signed bool,
	mode LUTVRMode,
) (*paletteLUT, bool, error) {
	sequenceElement, present := ds.Get(sequenceTag)
	if !present {
		return nil, false, nil
	}
	sequence, ok := sequenceElement.(*dataset.Sequence)
	if !ok {
		return nil, true, fmt.Errorf("%s must use SQ VR", sequenceTag)
	}
	if sequence.Count() == 0 || sequence.GetItem(0) == nil {
		return nil, true, fmt.Errorf("%s must contain a non-nil item", sequenceTag)
	}
	lut, err := buildPaletteLUTFromDataset(sequence.GetItem(0), byteOrder, signed, mode)
	if err != nil {
		return nil, true, fmt.Errorf("read %s item 0: %w", sequenceTag, err)
	}
	return lut, true, nil
}

// buildPaletteLUTFromDataset builds palette LUT using descriptors/data in the provided dataset (no sequence recursion).
//
//nolint:gocyclo // Complex function handling palette LUT variations
func buildPaletteLUTFromDataset(ds *dataset.Dataset, byteOrder binary.ByteOrder, signed bool, mode LUTVRMode) (*paletteLUT, error) {
	rDescriptor, err := dicomlut.ReadDescriptor(ds, tag.RedPaletteColorLookupTableDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("missing Red Palette LUT Descriptor: %w", err)
	}
	gDescriptor, err := dicomlut.ReadDescriptor(ds, tag.GreenPaletteColorLookupTableDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("missing Green Palette LUT Descriptor: %w", err)
	}
	bDescriptor, err := dicomlut.ReadDescriptor(ds, tag.BluePaletteColorLookupTableDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("missing Blue Palette LUT Descriptor: %w", err)
	}
	if rDescriptor != gDescriptor || rDescriptor != bDescriptor {
		return nil, fmt.Errorf("palette LUT descriptors do not match")
	}
	if rDescriptor.BitsPerEntry != 8 && rDescriptor.BitsPerEntry != 16 {
		return nil, fmt.Errorf("palette LUT bits per entry must be 8 or 16, got %d", rDescriptor.BitsPerEntry)
	}

	rLUT, err := readPaletteLUTChannel(ds,
		tag.RedPaletteColorLookupTableData,
		tag.SegmentedRedPaletteColorLookupTableData,
		rDescriptor,
		byteOrder,
		mode,
	)
	if err != nil {
		return nil, fmt.Errorf("read Red Palette LUT Data: %w", err)
	}
	gLUT, err := readPaletteLUTChannel(ds,
		tag.GreenPaletteColorLookupTableData,
		tag.SegmentedGreenPaletteColorLookupTableData,
		gDescriptor,
		byteOrder,
		mode,
	)
	if err != nil {
		return nil, fmt.Errorf("read Green Palette LUT Data: %w", err)
	}
	bLUT, err := readPaletteLUTChannel(ds,
		tag.BluePaletteColorLookupTableData,
		tag.SegmentedBluePaletteColorLookupTableData,
		bDescriptor,
		byteOrder,
		mode,
	)
	if err != nil {
		return nil, fmt.Errorf("read Blue Palette LUT Data: %w", err)
	}

	alphaDescriptorPresent := ds.Contains(tag.AlphaPaletteColorLookupTableDescriptor)
	alphaDataPresent := ds.Contains(tag.AlphaPaletteColorLookupTableData) ||
		ds.Contains(tag.SegmentedAlphaPaletteColorLookupTableData)
	if alphaDescriptorPresent != alphaDataPresent {
		return nil, fmt.Errorf("alpha palette descriptor and data must both be present")
	}
	var alphaLUT []uint16
	if alphaDescriptorPresent {
		descriptorElement, _ := ds.Get(tag.AlphaPaletteColorLookupTableDescriptor)
		if _, ok := descriptorElement.(*element.UnsignedShort); !ok {
			return nil, fmt.Errorf("alpha palette LUT descriptor must use US VR")
		}
		alphaDescriptor, err := dicomlut.ReadDescriptor(ds, tag.AlphaPaletteColorLookupTableDescriptor, false)
		if err != nil {
			return nil, fmt.Errorf("read Alpha Palette LUT Descriptor: %w", err)
		}
		if alphaDescriptor.EntryCount != rDescriptor.EntryCount ||
			alphaDescriptor.FirstMappedValue != rDescriptor.FirstMappedValue {
			return nil, fmt.Errorf("alpha palette descriptor must match RGB entry count and first mapped value")
		}
		if alphaDescriptor.BitsPerEntry != 8 {
			return nil, fmt.Errorf("alpha palette LUT bits per entry must be 8, got %d", alphaDescriptor.BitsPerEntry)
		}
		alphaLUT, err = readPaletteLUTChannel(ds,
			tag.AlphaPaletteColorLookupTableData,
			tag.SegmentedAlphaPaletteColorLookupTableData,
			alphaDescriptor,
			byteOrder,
			mode,
		)
		if err != nil {
			return nil, fmt.Errorf("read Alpha Palette LUT Data: %w", err)
		}
	}

	return &paletteLUT{
		first:    int32(rDescriptor.FirstMappedValue),
		entries:  buildPaletteEntries(int(rDescriptor.BitsPerEntry), rLUT, gLUT, bLUT, alphaLUT),
		hasAlpha: alphaDescriptorPresent,
	}, nil
}

func readPaletteLUTChannel(
	ds *dataset.Dataset,
	directTag, segmentedTag *tag.Tag,
	descriptor dicomlut.Descriptor,
	byteOrder binary.ByteOrder,
	mode LUTVRMode,
) ([]uint16, error) {
	hasDirect := ds.Contains(directTag)
	hasSegmented := ds.Contains(segmentedTag)
	if hasDirect && hasSegmented {
		return nil, fmt.Errorf(
			"palette channel contains both direct and segmented Palette LUT data (%s and %s)",
			directTag,
			segmentedTag,
		)
	}
	var values []uint16
	var err error
	if mode == LUTStandard {
		values, err = dicomlut.ReadDataStrict(ds, directTag, descriptor, byteOrder)
	} else {
		values, err = dicomlut.ReadData(ds, directTag, descriptor, byteOrder)
	}
	if err == nil {
		return values, nil
	}
	segmented, ok := ds.Get(segmentedTag)
	if !ok {
		return nil, err
	}
	switch value := segmented.(type) {
	case *element.OtherByte:
		if mode == LUTStandard {
			return nil, fmt.Errorf("segmented palette LUT %s uses non-standard OB VR; expected OW", segmentedTag)
		}
		return expandSegmentedLUT(value.GetData(), descriptor.EntryCount, byteOrder)
	case *element.OtherWord:
		return expandSegmentedLUT(value.GetData(), descriptor.EntryCount, dicomlut.NumericByteOrderOr(value, byteOrder))
	default:
		return nil, fmt.Errorf("unsupported segmented palette element type %T", segmented)
	}
}

func buildPaletteEntries(bits int, rLUT, gLUT, bLUT, alphaLUT []uint16) []colorconv.Color32 {
	shift := 0
	if bits > 8 {
		shift = bits - 8
	}

	entries := make([]colorconv.Color32, len(rLUT))
	for i := 0; i < len(rLUT); i++ {
		alpha := uint8(255)
		if alphaLUT != nil {
			alpha = clampByte(int(alphaLUT[i]))
		}
		entries[i] = colorconv.Color32{
			A: alpha,
			R: clampByte(int(rLUT[i] >> shift)),
			G: clampByte(int(gLUT[i] >> shift)),
			B: clampByte(int(bLUT[i] >> shift)),
		}
	}
	return entries
}

func expandSegmentedLUT(raw []byte, expectedSize int, byteOrder binary.ByteOrder) ([]uint16, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("segmented LUT data is empty")
	}
	if len(raw)%2 != 0 {
		return nil, fmt.Errorf("segmented LUT data has odd byte length %d", len(raw))
	}
	if byteOrder.Uint16(raw) == 2 {
		return nil, fmt.Errorf("segmented LUT first segment must not be indirect")
	}
	decoder := segmentedLUTDecoder{
		raw:          raw,
		byteOrder:    byteOrder,
		expectedSize: expectedSize,
		active:       make(map[int]bool),
	}
	position := 0
	for position < len(raw) {
		var err error
		position, err = decoder.decodeSegment(position)
		if err != nil {
			return nil, err
		}
	}
	if len(decoder.output) != expectedSize {
		return nil, fmt.Errorf("segmented LUT produced %d entries, want %d", len(decoder.output), expectedSize)
	}
	return decoder.output, nil
}

type segmentedLUTDecoder struct {
	raw          []byte
	byteOrder    binary.ByteOrder
	expectedSize int
	output       []uint16
	active       map[int]bool
}

func (d *segmentedLUTDecoder) decodeSegment(position int) (int, error) {
	if position < 0 || position%2 != 0 || position+4 > len(d.raw) {
		return 0, fmt.Errorf("segmented LUT segment at byte offset %d is truncated or unaligned", position)
	}
	if d.active[position] {
		return 0, fmt.Errorf("segmented LUT indirect segment cycle at byte offset %d", position)
	}
	d.active[position] = true
	defer delete(d.active, position)

	opcode := d.byteOrder.Uint16(d.raw[position:])
	length := int(d.byteOrder.Uint16(d.raw[position+2:]))
	if length == 0 {
		return 0, fmt.Errorf("segmented LUT opcode %d has zero length", opcode)
	}
	payload := position + 4

	switch opcode {
	case 0:
		end := payload + length*2
		if end > len(d.raw) {
			return 0, fmt.Errorf("segmented LUT discrete segment at byte offset %d is truncated", position)
		}
		if err := d.reserve(length); err != nil {
			return 0, err
		}
		for offset := payload; offset < end; offset += 2 {
			d.output = append(d.output, d.byteOrder.Uint16(d.raw[offset:]))
		}
		return end, nil
	case 1:
		if len(d.output) == 0 {
			return 0, fmt.Errorf("segmented LUT linear segment at byte offset %d has no prior value", position)
		}
		if payload+2 > len(d.raw) {
			return 0, fmt.Errorf("segmented LUT linear segment at byte offset %d is truncated", position)
		}
		if err := d.reserve(length); err != nil {
			return 0, err
		}
		start := int(d.output[len(d.output)-1])
		end := int(d.byteOrder.Uint16(d.raw[payload:]))
		for step := 1; step <= length; step++ {
			value := math.Round(float64(start) + float64(end-start)*float64(step)/float64(length))
			d.output = append(d.output, uint16(value))
		}
		return payload + 2, nil
	case 2:
		if payload+4 > len(d.raw) {
			return 0, fmt.Errorf("segmented LUT indirect segment at byte offset %d is truncated", position)
		}
		lowWord := d.byteOrder.Uint16(d.raw[payload:])
		highWord := d.byteOrder.Uint16(d.raw[payload+2:])
		offset := int(uint32(lowWord) | uint32(highWord)<<16)
		if offset < 0 || offset%2 != 0 || offset >= len(d.raw) {
			return 0, fmt.Errorf("segmented LUT indirect byte offset %d is out of range", offset)
		}
		referencedPosition := offset
		for segment := 0; segment < length; segment++ {
			if referencedPosition+2 > len(d.raw) {
				return 0, fmt.Errorf("segmented LUT indirect reference at byte offset %d is truncated", referencedPosition)
			}
			if d.byteOrder.Uint16(d.raw[referencedPosition:]) == 2 {
				return 0, fmt.Errorf("segmented LUT indirect segment at byte offset %d references another indirect segment", position)
			}
			var err error
			referencedPosition, err = d.decodeSegment(referencedPosition)
			if err != nil {
				return 0, err
			}
		}
		return payload + 4, nil
	default:
		return 0, fmt.Errorf("unsupported segmented LUT opcode %d", opcode)
	}
}

func (d *segmentedLUTDecoder) reserve(count int) error {
	if d.expectedSize >= 0 && len(d.output)+count > d.expectedSize {
		return fmt.Errorf("segmented LUT produces more than %d entries", d.expectedSize)
	}
	return nil
}

func stripFramePadding(frames [][]byte) {
	for i := range frames {
		frames[i] = codec.StripTrailingPadding(frames[i])
	}
}

// buildFragmentSequence creates an OB fragment sequence from per-frame compressed data,
// populating the Basic Offset Table for multi-frame images.
// If an existing BOT is provided and matches frames length, it is used; otherwise BOT is rebuilt.
func buildFragmentSequence(frames [][]byte, _ []uint32, _ uint16) (element.Element, error) {
	if len(frames) == 0 {
		return nil, fmt.Errorf("no frame data provided for fragment sequence")
	}

	// Rebuild BOT for the fragment layout emitted below: one fragment item per
	// frame. Existing BOT values may refer to a different source fragment layout.
	offsets := make([]uint32, 0, len(frames))
	var runningOffset uint32
	for i, frame := range frames {
		offsets = append(offsets, runningOffset)
		paddedSize := len(frame)
		if paddedSize%2 != 0 {
			paddedSize++
		}
		if paddedSize > int(math.MaxUint32-8) {
			return nil, fmt.Errorf("fragment too large to represent in BOT at frame %d", i)
		}
		padded := uint32(paddedSize)
		if runningOffset > math.MaxUint32-8-padded {
			return nil, fmt.Errorf("fragment too large to represent in BOT at frame %d", i)
		}
		runningOffset += 8 + padded
	}

	// Encapsulated pixel data uses OB regardless of BitsAllocated.
	obf := element.NewOtherByteFragment(tag.PixelData)
	for _, frame := range frames {
		obf.AddFragment(buffer.NewMemory(frame))
	}
	obf.SetOffsetTable(offsets)
	return obf, nil
}

func frameCountFromDataset(ds *dataset.Dataset) int {
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil && nf > 0 {
		return int(nf)
	}
	if nfStr, ok := ds.GetString(tag.NumberOfFrames); ok {
		if parsed, err := strconv.Atoi(strings.TrimSpace(nfStr)); err == nil && parsed > 0 {
			return parsed
		}
	}
	return 1
}
