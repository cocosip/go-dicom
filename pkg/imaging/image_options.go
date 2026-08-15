// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "github.com/cocosip/go-dicom/pkg/imaging/codec"

type dicomImageConfig struct {
	codecRegistry   *codec.Registry
	codecParameters codec.Parameters
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
