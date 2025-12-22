// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package codec provides image compression and decompression interfaces and implementations.
//
// Following fo-dicom's architecture, codecs work with per-frame encoding/decoding operations.
// The codec package uses the types package to avoid circular dependencies.
package codec

import (
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/types"
)

// Codec represents a DICOM image codec that can encode and decode pixel data.
// This interface exactly mirrors fo-dicom's IDicomCodec interface:
//
//	public interface IDicomCodec {
//	    string Name { get; }
//	    DicomTransferSyntax TransferSyntax { get; }
//	    DicomCodecParams GetDefaultParameters();
//	    void Encode(DicomPixelData oldPixelData, DicomPixelData newPixelData, DicomCodecParams parameters);
//	    void Decode(DicomPixelData oldPixelData, DicomPixelData newPixelData, DicomCodecParams parameters);
//	}
type Codec interface {
	// Name returns the codec name.
	Name() string

	// TransferSyntax returns the transfer syntax this codec handles.
	TransferSyntax() *transfer.Syntax

	// GetDefaultParameters returns the default codec parameters.
	GetDefaultParameters() Parameters

	// Encode encodes pixel data from oldPixelData to newPixelData.
	// This mirrors fo-dicom's: void Encode(DicomPixelData oldPixelData, DicomPixelData newPixelData, DicomCodecParams parameters)
	Encode(oldPixelData types.PixelData, newPixelData types.PixelData, parameters Parameters) error

	// Decode decodes pixel data from oldPixelData to newPixelData.
	// This mirrors fo-dicom's: void Decode(DicomPixelData oldPixelData, DicomPixelData newPixelData, DicomCodecParams parameters)
	Decode(oldPixelData types.PixelData, newPixelData types.PixelData, parameters Parameters) error
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
