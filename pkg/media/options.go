// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// IconGenerator renders an 8-bit grayscale DICOMDIR icon without depending on media types.
type IconGenerator interface {
	GenerateDirectoryIcon(ds *dataset.Dataset, frame int) (width, height int, pixels []byte, err error)
}

type config struct {
	transferSyntax *transfer.Syntax
	iconGenerator  IconGenerator
	imageIcons     bool
}

// Option configures a newly created Directory.
type Option func(*config)

// WithTransferSyntax selects Explicit or Implicit VR Little Endian for a new DICOMDIR.
func WithTransferSyntax(ts *transfer.Syntax) Option {
	return func(config *config) {
		config.transferSyntax = ts
	}
}

// WithIconGenerator sets the optional DICOMDIR icon generator.
func WithIconGenerator(generator IconGenerator) Option {
	return func(config *config) {
		config.iconGenerator = generator
	}
}

// WithImageIcons enables or disables icon generation while adding image records.
func WithImageIcons(enabled bool) Option {
	return func(config *config) {
		config.imageIcons = enabled
	}
}
