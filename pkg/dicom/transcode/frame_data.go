// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package transcode provides Dataset-level transfer syntax conversion.
package transcode

import (
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

func newFrameData(info codec.FrameInfo, encapsulated bool) (*pixeldata.Data, error) {
	photometric := info.PhotometricInterpretation
	if !encapsulated {
		// Native (decoded) data cannot carry an encapsulated-only photometric
		// interpretation (PS3.5 8.2).
		switch photometric.Value {
		case pixel.YbrRct.Value, pixel.YbrIct.Value, pixel.YbrPartial420.Value:
			if info.SamplesPerPixel == 1 {
				photometric = *pixel.Monochrome2
			} else {
				photometric = *pixel.RGBPhotometric
			}
		}
	}
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
