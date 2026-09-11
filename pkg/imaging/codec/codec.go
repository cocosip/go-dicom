// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package codec provides frame-level image compression and decompression contracts.
package codec

import (
	"context"
	"errors"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// ErrInvalidParameters identifies parameters that a codec cannot accept.
var ErrInvalidParameters = errors.New("invalid codec parameters")

// FrameInfo contains the pixel format metadata needed by a frame codec.
type FrameInfo struct {
	Width                     uint16
	Height                    uint16
	BitDepth                  pixel.BitDepth
	SamplesPerPixel           uint16
	PixelRepresentation       pixel.Representation
	PlanarConfiguration       pixel.PlanarConfiguration
	PhotometricInterpretation pixel.PhotometricInterpretation
}

// Validate checks whether frame metadata is usable by codecs and sinks.
func (info FrameInfo) Validate() error {
	if info.Width == 0 || info.Height == 0 {
		return fmt.Errorf("frame dimensions must be greater than zero")
	}
	if info.PixelRepresentation != pixel.UnsignedPixels && info.PixelRepresentation != pixel.SignedPixels {
		return fmt.Errorf("invalid pixel representation %d", info.PixelRepresentation)
	}
	if info.BitDepth.BitsStored == 0 || !info.BitDepth.IsValid() {
		return fmt.Errorf(
			"invalid bit depth: allocated=%d stored=%d high-bit=%d",
			info.BitDepth.BitsAllocated,
			info.BitDepth.BitsStored,
			info.BitDepth.HighBit,
		)
	}
	if info.BitDepth.IsSigned != info.PixelRepresentation.IsSigned() {
		return fmt.Errorf("frame bit depth signedness does not match pixel representation")
	}
	if info.SamplesPerPixel == 0 {
		return fmt.Errorf("samples per pixel must be greater than zero")
	}
	if info.PlanarConfiguration != pixel.InterleavedPlanar && info.PlanarConfiguration != pixel.PlanarPlanar {
		return fmt.Errorf("invalid planar configuration %d", info.PlanarConfiguration)
	}
	if info.PhotometricInterpretation.Value == "" {
		return fmt.Errorf("photometric interpretation must be set")
	}
	if info.PhotometricInterpretation.IsColor &&
		!info.PhotometricInterpretation.IsPalette &&
		info.SamplesPerPixel < 3 {
		return fmt.Errorf("color images require at least 3 samples per pixel, got %d", info.SamplesPerPixel)
	}
	if !info.PhotometricInterpretation.IsColor &&
		!info.PhotometricInterpretation.IsPalette &&
		info.SamplesPerPixel != 1 {
		return fmt.Errorf("grayscale images require exactly 1 sample per pixel, got %d", info.SamplesPerPixel)
	}
	return nil
}

// FrameSource provides read-only, independently owned frames to a codec.
type FrameSource interface {
	FrameCount() int
	Frame(context.Context, int) ([]byte, error)
	FrameInfo() FrameInfo
	Encapsulated() bool
}

// FrameSink receives codec output frames and their resulting metadata.
type FrameSink interface {
	AddFrame(context.Context, []byte) error
	SetFrameInfo(FrameInfo) error
}

// Codec represents a DICOM image codec that can encode and decode frames.
type Codec interface {
	// Name returns the codec name.
	Name() string

	// TransferSyntax returns the transfer syntax this codec handles.
	TransferSyntax() *transfer.Syntax

	// DefaultParameters returns a new independent set of default parameters.
	DefaultParameters() Parameters

	// Encode writes encoded source frames to sink.
	Encode(context.Context, FrameSource, FrameSink, Parameters) error

	// Decode writes decoded source frames to sink.
	Decode(context.Context, FrameSource, FrameSink, Parameters) error
}

// Parameters is the ownership and validation contract for codec parameters.
type Parameters interface {
	Clone() Parameters
	Validate() error
}

// NoParameters is the parameter value for codecs without configurable options.
type NoParameters struct{}

// Clone returns an independent NoParameters value.
func (NoParameters) Clone() Parameters { return NoParameters{} }

// Validate reports that NoParameters is always valid.
func (NoParameters) Validate() error { return nil }

// ByteSwapMode controls whether NativeCodec swaps bytes within native OW words.
type ByteSwapMode uint8

const (
	// ByteSwapDefault preserves the codec's transfer-syntax behavior.
	ByteSwapDefault ByteSwapMode = iota
	// ByteSwapDisabled preserves the source byte order.
	ByteSwapDisabled
	// ByteSwapEnabled swaps bytes within every 16-bit OW word.
	ByteSwapEnabled
)

// NativeParameters configures native transfer-syntax byte order conversion.
type NativeParameters struct {
	ByteSwap ByteSwapMode
}

// Clone returns an independent copy of the native parameters.
func (p NativeParameters) Clone() Parameters { return p }

// Validate rejects unknown byte-swap modes.
func (p NativeParameters) Validate() error {
	if p.ByteSwap > ByteSwapEnabled {
		return fmt.Errorf("unknown byte-swap mode %d", p.ByteSwap)
	}
	return nil
}

func nativeParameters(parameters Parameters) (NativeParameters, error) {
	var params NativeParameters
	switch value := parameters.(type) {
	case NativeParameters:
		params = value
	case *NativeParameters:
		if value != nil {
			params = *value
			break
		}
		return NativeParameters{}, fmt.Errorf("%w: NativeCodec requires NativeParameters, got %T", ErrInvalidParameters, parameters)
	default:
		return NativeParameters{}, fmt.Errorf("%w: NativeCodec requires NativeParameters, got %T", ErrInvalidParameters, parameters)
	}
	if err := params.Validate(); err != nil {
		return NativeParameters{}, fmt.Errorf("%w: NativeParameters: %w", ErrInvalidParameters, err)
	}
	return params, nil
}

// PrepareParameters returns an independently owned, validated parameter value
// suitable for one codec invocation.
func PrepareParameters(c Codec, parameters Parameters) (Parameters, error) {
	if interfaceIsNil(c) {
		return nil, fmt.Errorf("%w: codec must not be nil", ErrInvalidParameters)
	}
	if parameters == nil {
		parameters = c.DefaultParameters()
	} else if interfaceIsNil(parameters) {
		return nil, fmt.Errorf("%w: parameters must not be a typed nil", ErrInvalidParameters)
	}
	if interfaceIsNil(parameters) {
		return nil, fmt.Errorf("%w: %s returned nil default parameters", ErrInvalidParameters, c.Name())
	}
	owned := parameters.Clone()
	if interfaceIsNil(owned) {
		return nil, fmt.Errorf("%w: %T returned a nil clone", ErrInvalidParameters, parameters)
	}
	if err := owned.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %T: %w", ErrInvalidParameters, parameters, err)
	}
	return owned, nil
}
