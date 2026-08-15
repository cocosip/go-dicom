// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

// NewDicomImageFromDataset creates a renderable image from a DICOM dataset.
func NewDicomImageFromDataset(ds *dataset.Dataset, options ...DicomImageOption) (*DicomImage, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset is nil")
	}
	config := &dicomImageConfig{codecRegistry: codec.GetGlobalRegistry()}
	for _, option := range options {
		if option != nil {
			option(config)
		}
	}
	pixelData, err := CreatePixelData(ds)
	if err != nil {
		return nil, fmt.Errorf("create pixel data: %w", err)
	}
	image := NewDicomImage(pixelData)
	image.dataset = ds.Clone()
	image.autoApplyLUTToAllFrames = true
	image.overlays = imageOverlays(image.dataset, pixelData)
	if pixelData.IsEncapsulated() {
		syntax, err := transfer.Parse(pixelData.Info.TransferSyntaxUID)
		if err != nil {
			return nil, fmt.Errorf("parse pixel data transfer syntax: %w", err)
		}
		decoder, ok := config.codecRegistry.GetCodec(syntax)
		if !ok {
			return nil, fmt.Errorf("no codec available for transfer syntax %s", pixelData.Info.TransferSyntaxUID)
		}
		parameters := config.codecParameters
		if parameters == nil {
			parameters = decoder.GetDefaultParameters()
		}
		if err := image.DecodeIfNeeded(decoder, parameters); err != nil {
			return nil, err
		}
		if image.pixelData.Info.PhotometricInterpretation != nil &&
			image.pixelData.Info.PhotometricInterpretation.Value == photometricPaletteColor {
			if err := convertPaletteToRGB(image.dataset, image.pixelData); err != nil {
				return nil, fmt.Errorf("palette conversion failed: %w", err)
			}
		}
	}
	return image, nil
}

// NewDicomImageFromParseResult creates a renderable image from parsed DICOM data.
func NewDicomImageFromParseResult(result *parser.ParseResult, options ...DicomImageOption) (*DicomImage, error) {
	if result == nil || result.Dataset == nil {
		return nil, fmt.Errorf("parse result dataset is nil")
	}
	return NewDicomImageFromDataset(result.Dataset, options...)
}

// OpenDicomImage parses a DICOM file and creates a renderable image.
func OpenDicomImage(path string, options ...DicomImageOption) (*DicomImage, error) {
	result, err := parser.ParseFile(path)
	if err != nil {
		return nil, fmt.Errorf("open DICOM image: %w", err)
	}
	return NewDicomImageFromParseResult(result, options...)
}
