// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

// VRMode controls how non-standard Pixel Data VRs are handled.
type VRMode uint8

const (
	// PixelDataCompatible accepts recoverable non-standard Pixel Data VRs.
	PixelDataCompatible VRMode = iota
	// PixelDataStandard rejects Pixel Data VRs that violate DICOM PS3.5.
	PixelDataStandard
)

type readConfig struct {
	vrMode VRMode
}

// ReadOption configures Dataset Pixel Data extraction.
type ReadOption func(*readConfig)

// WithPixelDataVRMode selects standard validation or compatibility recovery.
func WithPixelDataVRMode(mode VRMode) ReadOption {
	return func(config *readConfig) {
		config.vrMode = mode
	}
}
