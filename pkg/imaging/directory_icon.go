// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
	"image"
	"math"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	golangdraw "golang.org/x/image/draw"
)

const maximumDirectoryIconDimension = 128

// DirectoryIconGenerator renders DICOM datasets as DICOMDIR-compatible icons.
type DirectoryIconGenerator struct{}

// NewDirectoryIconGenerator creates a pure-Go DICOMDIR icon generator.
func NewDirectoryIconGenerator() *DirectoryIconGenerator {
	return &DirectoryIconGenerator{}
}

// GenerateDirectoryIcon renders frame as an 8-bit grayscale image no larger than 128x128.
func (g *DirectoryIconGenerator) GenerateDirectoryIcon(ds *dataset.Dataset, frame int) (width, height int, pixels []byte, err error) {
	if g == nil {
		return 0, 0, nil, fmt.Errorf("directory icon generator cannot be nil")
	}
	if ds == nil {
		return 0, 0, nil, fmt.Errorf("DICOM Dataset cannot be nil")
	}
	pixelData, err := CreatePixelData(ds)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("create DICOM pixel data: %w", err)
	}
	if frame < 0 || frame >= pixelData.FrameCount() {
		return 0, 0, nil, fmt.Errorf("frame index out of range: %d (total: %d)", frame, pixelData.FrameCount())
	}

	dicomImage := NewDicomImage(pixelData)
	if pixelData.Info.Encapsulated {
		syntax, err := transfer.Parse(pixelData.Info.TransferSyntaxUID)
		if err != nil {
			return 0, 0, nil, fmt.Errorf("parse pixel transfer syntax: %w", err)
		}
		decoder, ok := codec.GetGlobalRegistry().GetCodec(syntax)
		if !ok {
			return 0, 0, nil, fmt.Errorf("no codec registered for transfer syntax %s", syntax.UID().UID())
		}
		if err := dicomImage.DecodeIfNeeded(decoder, decoder.GetDefaultParameters()); err != nil {
			return 0, 0, nil, err
		}
	}

	rendered, err := dicomImage.RenderFrameImage(frame)
	if err != nil {
		return 0, 0, nil, err
	}
	width, height = directoryIconDimensions(rendered.Bounds().Dx(), rendered.Bounds().Dy())
	gray := image.NewGray(image.Rect(0, 0, width, height))
	golangdraw.CatmullRom.Scale(gray, gray.Bounds(), rendered, rendered.Bounds(), golangdraw.Src, nil)

	pixels = make([]byte, width*height)
	for y := 0; y < height; y++ {
		start := gray.PixOffset(0, y)
		copy(pixels[y*width:(y+1)*width], gray.Pix[start:start+width])
	}
	return width, height, pixels, nil
}

func directoryIconDimensions(width, height int) (int, int) {
	if width <= maximumDirectoryIconDimension && height <= maximumDirectoryIconDimension {
		return width, height
	}
	scale := float64(maximumDirectoryIconDimension) / float64(max(width, height))
	resizedWidth := max(1, int(math.Round(float64(width)*scale)))
	resizedHeight := max(1, int(math.Round(float64(height)*scale)))
	return resizedWidth, resizedHeight
}
