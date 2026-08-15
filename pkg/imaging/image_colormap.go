// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"image"
	"image/color"

	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
)

// SetGrayscaleColorMap applies a 256-entry color map to rendered grayscale values.
func (img *DicomImage) SetGrayscaleColorMap(colorMap [256]imagetypes.Color32) {
	from, to := img.lutFrameRange()
	img.mu.Lock()
	defer img.mu.Unlock()
	for frame := from; frame <= to; frame++ {
		img.grayscaleColorMaps[frame] = colorMap
	}
}

// GrayscaleColorMap returns the current frame color map and whether one is set.
func (img *DicomImage) GrayscaleColorMap() ([256]imagetypes.Color32, bool) {
	img.mu.RLock()
	defer img.mu.RUnlock()
	colorMap, ok := img.grayscaleColorMaps[img.currentFrame]
	return colorMap, ok
}

// ClearGrayscaleColorMap restores normal grayscale rendering.
func (img *DicomImage) ClearGrayscaleColorMap() {
	from, to := img.lutFrameRange()
	img.mu.Lock()
	defer img.mu.Unlock()
	for frame := from; frame <= to; frame++ {
		delete(img.grayscaleColorMaps, frame)
	}
}

func (img *DicomImage) applyGrayscaleColorMap(source image.Image, frame int) image.Image {
	gray, ok := source.(*image.Gray)
	if !ok {
		return source
	}
	img.mu.RLock()
	colorMap, ok := img.grayscaleColorMaps[frame]
	img.mu.RUnlock()
	if !ok {
		return source
	}
	destination := image.NewRGBA(gray.Bounds())
	for y := gray.Bounds().Min.Y; y < gray.Bounds().Max.Y; y++ {
		for x := gray.Bounds().Min.X; x < gray.Bounds().Max.X; x++ {
			mapped := colorMap[gray.GrayAt(x, y).Y]
			destination.Set(x, y, color.NRGBA{R: mapped.R, G: mapped.G, B: mapped.B, A: mapped.A})
		}
	}
	return destination
}
