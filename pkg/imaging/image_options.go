// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

type dicomImageConfig struct {
	codecRegistry     *codec.Registry
	codecParameters   codec.Parameters
	windowIndex       int
	voiLUTIndex       int
	lutVRMode         pixeldata.LUTVRMode
	presentationState *dataset.Dataset
}

// WithPresentationState supplies a Grayscale Softcopy Presentation State
// Dataset used for Softcopy VOI and Presentation LUT selection.
func WithPresentationState(ds *dataset.Dataset) DicomImageOption {
	return func(config *dicomImageConfig) {
		config.presentationState = ds
	}
}

// WithWindowIndex selects one alternative Window Center/Width pair. The
// default is the first pair.
func WithWindowIndex(index int) DicomImageOption {
	return func(config *dicomImageConfig) {
		config.windowIndex = index
	}
}

// WithVOILUTIndex selects one alternative VOI LUT Sequence item. The default
// is the first item.
func WithVOILUTIndex(index int) DicomImageOption {
	return func(config *dicomImageConfig) {
		config.voiLUTIndex = index
	}
}

// WithLUTVRMode selects compatibility or strict DICOM LUT Data VR handling
// during image rendering. The default is LUTCompatible for interoperability.
func WithLUTVRMode(mode pixeldata.LUTVRMode) DicomImageOption {
	return func(config *dicomImageConfig) {
		config.lutVRMode = mode
	}
}

// DicomImageOption configures Dataset and file based DicomImage constructors.
type DicomImageOption func(*dicomImageConfig)

// WithImageCodecRegistry selects the registry used for automatic frame decoding.
func WithImageCodecRegistry(registry *codec.Registry) DicomImageOption {
	return func(config *dicomImageConfig) {
		if registry != nil {
			config.codecRegistry = registry
		}
	}
}

// WithImageCodecParameters supplies parameters for automatic frame decoding.
func WithImageCodecParameters(parameters codec.Parameters) DicomImageOption {
	return func(config *dicomImageConfig) {
		config.codecParameters = parameters
	}
}
