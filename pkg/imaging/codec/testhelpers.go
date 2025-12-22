// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/types"
)

// testPixelData is a simple in-memory implementation of types.PixelData for testing.
type testPixelData struct {
	frames                    [][]byte
	width                     uint16
	height                    uint16
	bitsAllocated             uint16
	bitsStored                uint16
	highBit                   uint16
	samplesPerPixel           uint16
	pixelRepresentation       uint16
	planarConfiguration       uint16
	photometricInterpretation string
	encapsulated              bool
}

// newTestPixelData creates a new test pixel data instance.
func newTestPixelData(info *types.FrameInfo) *testPixelData {
	return &testPixelData{
		frames:                    make([][]byte, 0),
		width:                     info.Width,
		height:                    info.Height,
		bitsAllocated:             info.BitsAllocated,
		bitsStored:                info.BitsStored,
		highBit:                   info.HighBit,
		samplesPerPixel:           info.SamplesPerPixel,
		pixelRepresentation:       info.PixelRepresentation,
		planarConfiguration:       info.PlanarConfiguration,
		photometricInterpretation: info.PhotometricInterpretation,
	}
}

// GetFrame returns the pixel data for the specified frame.
func (pd *testPixelData) GetFrame(frameIndex int) ([]byte, error) {
	if frameIndex < 0 || frameIndex >= len(pd.frames) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(pd.frames))
	}
	return pd.frames[frameIndex], nil
}

// AddFrame appends a new frame to the pixel data.
func (pd *testPixelData) AddFrame(frameData []byte) error {
	pd.frames = append(pd.frames, frameData)
	return nil
}

// FrameCount returns the number of frames.
func (pd *testPixelData) FrameCount() int {
	return len(pd.frames)
}

// GetFrameInfo returns frame metadata.
func (pd *testPixelData) GetFrameInfo() *types.FrameInfo {
	return &types.FrameInfo{
		Width:                     pd.width,
		Height:                    pd.height,
		BitsAllocated:             pd.bitsAllocated,
		BitsStored:                pd.bitsStored,
		HighBit:                   pd.highBit,
		SamplesPerPixel:           pd.samplesPerPixel,
		PixelRepresentation:       pd.pixelRepresentation,
		PlanarConfiguration:       pd.planarConfiguration,
		PhotometricInterpretation: pd.photometricInterpretation,
	}
}

// IsEncapsulated returns whether the pixel data is encapsulated.
func (pd *testPixelData) IsEncapsulated() bool {
	return pd.encapsulated
}
