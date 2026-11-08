// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"fmt"

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
	fileSetID      string
	fileSetIDErr   error
}

// Option configures a newly created Directory.
type Option func(*config)

// WithTransferSyntax selects the DICOMDIR transfer syntax. PS3.10 §8.6 requires
// Explicit VR Little Endian for writing, so any other value is rejected by
// NewDirectory; reading tolerates legacy Implicit VR Little Endian file-sets.
func WithTransferSyntax(ts *transfer.Syntax) Option {
	return func(config *config) {
		config.transferSyntax = ts
	}
}

// WithFileSetID sets File-set ID (0004,1130), limited to 0-16 characters from
// the A-Z, 0-9, underscore repertoire (PS3.10 §8.1). The default is
// zero-length, which is conformant for Type 2.
func WithFileSetID(fileSetID string) Option {
	return func(config *config) {
		if err := validateFileSetID(fileSetID); err != nil {
			config.fileSetIDErr = err
			return
		}
		config.fileSetID = fileSetID
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

const maxFileSetIDLength = 16

// validateFileSetID enforces PS3.10 §8.1/§8.5: at most 16 characters from the
// G0 subset, restricted to the uppercase letters, digits, and underscore
// repertoire shared by the CS value representation.
func validateFileSetID(fileSetID string) error {
	if len(fileSetID) > maxFileSetIDLength {
		return fmt.Errorf("file-set ID %q exceeds %d characters", fileSetID, maxFileSetIDLength)
	}
	for _, character := range fileSetID {
		valid := (character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') || character == '_'
		if !valid {
			return fmt.Errorf("file-set ID character %q is outside the A-Z, 0-9, underscore repertoire", character)
		}
	}
	return nil
}


