// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
	"strconv"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
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

	// Extract raw bytes from pixel data element
	var pixelData []byte
	switch elem := pixelDataElem.(type) {
	case interface{ GetData() []byte }:
		pixelData = elem.GetData()
	default:
		return nil, fmt.Errorf("unsupported pixel data element type: %T", pixelDataElem)
	}

	if len(pixelData) == 0 {
		return nil, fmt.Errorf("pixel data is empty")
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
	numberOfFrames := 1
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		numberOfFrames = int(nf)
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
	}

	// Create DICOM pixel data from bytes
	return NewDicomPixelDataFromBytes(info, pixelData)
}
