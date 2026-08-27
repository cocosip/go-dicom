// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package imagetypes provides shared types and interfaces for the imaging package.
// This package defines interfaces that can be used across different packages
// without creating circular dependencies.
package imagetypes

// PixelData represents the interface for DICOM pixel data operations.
// This mirrors fo-dicom's abstract DicomPixelData class pattern.
//
// In fo-dicom, DicomPixelData is an abstract class with three implementations:
// - OtherBytePixelData (uncompressed, OB VR)
// - OtherWordPixelData (uncompressed, OW VR)
// - EncapsulatedPixelData (compressed/encapsulated)
//
// This Go interface provides the same abstraction for use by codecs.
type PixelData interface {
	// GetFrame returns the pixel data for the specified frame (0-indexed).
	// In fo-dicom: public abstract IByteBuffer GetFrame(int frame);
	GetFrame(frameIndex int) ([]byte, error)

	// AddFrame appends a new frame to the pixel data.
	// In fo-dicom: public abstract void AddFrame(IByteBuffer data);
	AddFrame(frameData []byte) error

	// FrameCount returns the number of frames in the pixel data.
	FrameCount() int

	// GetFrameInfo returns frame metadata for codec operations.
	GetFrameInfo() *FrameInfo

	// IsEncapsulated returns true if pixel data is encapsulated (compressed).
	IsEncapsulated() bool
}

// FrameInfoSetter optionally allows codecs to report the metadata of their
// decoded or encoded pixel data without coupling to a DICOM dataset.
//
// PixelData implementations are not required to implement this interface.
// Codecs should continue to work with read-only PixelData implementations.
type FrameInfoSetter interface {
	SetFrameInfo(info *FrameInfo)
}

// SetFrameInfo updates pixel metadata when pixelData supports FrameInfoSetter.
// It returns false when the implementation exposes read-only metadata.
func SetFrameInfo(pixelData PixelData, info *FrameInfo) bool {
	if pixelData == nil || info == nil {
		return false
	}
	setter, ok := pixelData.(FrameInfoSetter)
	if !ok {
		return false
	}
	setter.SetFrameInfo(info)
	return true
}

// FrameInfo contains the metadata needed for encoding/decoding a frame.
// This is a lightweight struct used by codecs for frame-level operations.
type FrameInfo struct {
	// Image dimensions
	Width  uint16
	Height uint16

	// Bit depth information
	BitsAllocated   uint16
	BitsStored      uint16
	HighBit         uint16
	SamplesPerPixel uint16

	// Pixel representation (0 = unsigned, 1 = signed)
	PixelRepresentation uint16

	// Planar configuration (0 = interleaved, 1 = planar)
	PlanarConfiguration uint16

	// Photometric interpretation
	PhotometricInterpretation string
}
