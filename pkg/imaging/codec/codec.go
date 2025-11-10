// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package codec provides image compression and decompression interfaces and implementations.
package codec

import (
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// Codec represents a DICOM image codec that can encode and decode pixel data.
type Codec interface {
	// Name returns the codec name.
	Name() string

	// TransferSyntax returns the transfer syntax this codec handles.
	TransferSyntax() *transfer.Syntax

	// Encode encodes pixel data from source to destination with the given parameters.
	Encode(src *PixelData, dst *PixelData, params Parameters) error

	// Decode decodes pixel data from source to destination with the given parameters.
	Decode(src *PixelData, dst *PixelData, params Parameters) error
}

// Parameters represents codec-specific parameters.
// Different codec implementations may provide their own parameter types.
type Parameters interface {
	// GetParameter retrieves a parameter by name.
	GetParameter(name string) interface{}

	// SetParameter sets a parameter value.
	SetParameter(name string, value interface{})
}

// BaseParameters provides a basic implementation of Parameters.
type BaseParameters struct {
	params map[string]interface{}
}

// NewBaseParameters creates a new BaseParameters instance.
func NewBaseParameters() *BaseParameters {
	return &BaseParameters{
		params: make(map[string]interface{}),
	}
}

// GetParameter retrieves a parameter by name.
func (p *BaseParameters) GetParameter(name string) interface{} {
	return p.params[name]
}

// SetParameter sets a parameter value.
func (p *BaseParameters) SetParameter(name string, value interface{}) {
	p.params[name] = value
}

// PixelData represents pixel data for encoding/decoding operations.
// This is a simplified version that holds only the essential information
// for codec operations, avoiding circular dependencies with the full dataset package.
type PixelData struct {
	// Raw pixel data bytes
	Data []byte

	// Image dimensions
	Width  uint16
	Height uint16

	// Number of frames
	NumberOfFrames int

	// Bit depth information
	BitsAllocated uint16
	BitsStored    uint16
	HighBit       uint16

	// Sampling information
	SamplesPerPixel uint16

	// Pixel representation (0 = unsigned, 1 = signed)
	PixelRepresentation uint16

	// Planar configuration (0 = interleaved, 1 = planar)
	PlanarConfiguration uint16

	// Photometric interpretation value
	PhotometricInterpretation string

	// Transfer syntax
	TransferSyntaxUID string
}

// BytesAllocated returns the number of bytes allocated per pixel sample.
func (pd *PixelData) BytesAllocated() int {
	return int((pd.BitsAllocated-1)/8 + 1)
}

// UncompressedFrameSize calculates the uncompressed size of a single frame.
func (pd *PixelData) UncompressedFrameSize() int {
	if pd.BitsAllocated == 1 {
		return (int(pd.Width)*int(pd.Height)-1)/8 + 1
	}

	// Handle special case for YBR_FULL_422 with uneven width
	actualWidth := int(pd.Width)
	if actualWidth%2 != 0 &&
		(pd.PhotometricInterpretation == "YBR_FULL_422" ||
			pd.PhotometricInterpretation == "YBR_PARTIAL_422" ||
			pd.PhotometricInterpretation == "YBR_PARTIAL_420") {
		actualWidth++
	}

	// Handle YBR_FULL_422 special case for uncompressed data
	if pd.PhotometricInterpretation == "YBR_FULL_422" {
		// For uncompressed transfer syntaxes, chrominance channels are downsampled
		// Check if this is an uncompressed transfer syntax
		// This is a simplified check - in a full implementation, you would
		// check against the actual transfer syntax object
		return pd.BytesAllocated() * 2 * actualWidth * int(pd.Height)
	}

	return pd.BytesAllocated() * int(pd.SamplesPerPixel) * actualWidth * int(pd.Height)
}

// IsSigned returns true if the pixel data is signed.
func (pd *PixelData) IsSigned() bool {
	return pd.PixelRepresentation == 1
}

// IsInterleaved returns true if the planar configuration is interleaved.
func (pd *PixelData) IsInterleaved() bool {
	return pd.PlanarConfiguration == 0
}
