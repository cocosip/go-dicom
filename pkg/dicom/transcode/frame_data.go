// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package transcode provides Dataset-level transfer syntax conversion.
package transcode

import (
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

func newFrameData(info codec.FrameInfo, encapsulated bool) (*pixeldata.Data, error) {
	photometric := info.PhotometricInterpretation
	vrCode := "OB"
	if !encapsulated && info.BitDepth.BitsAllocated > 8 {
		vrCode = "OW"
	}
	return pixeldata.New(&pixeldata.Info{
		Width:                     info.Width,
		Height:                    info.Height,
		NumberOfFrames:            1,
		BitsAllocated:             info.BitDepth.BitsAllocated,
		BitsStored:                info.BitDepth.BitsStored,
		HighBit:                   info.BitDepth.HighBit,
		SamplesPerPixel:           info.SamplesPerPixel,
		PixelRepresentation:       info.PixelRepresentation,
		PlanarConfiguration:       info.PlanarConfiguration,
		PhotometricInterpretation: &photometric,
		VRCode:                    vrCode,
		Encapsulated:              encapsulated,
	})
}
