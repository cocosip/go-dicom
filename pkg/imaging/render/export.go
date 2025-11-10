// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

// ImageExporter provides functionality to export DICOM images to standard formats
type ImageExporter struct {
	pipeline Pipeline
	converter *ColorSpaceConverter
}

// NewImageExporter creates a new ImageExporter
func NewImageExporter(pipeline Pipeline) *ImageExporter {
	return &ImageExporter{
		pipeline:  pipeline,
		converter: NewColorSpaceConverter(),
	}
}

// ExportFormat represents the output image format
type ExportFormat int

const (
	// FormatPNG exports as PNG format
	FormatPNG ExportFormat = iota
	// FormatJPEG exports as JPEG format
	FormatJPEG
)

// ExportOptions contains options for exporting images
type ExportOptions struct {
	// Format specifies the output format (PNG or JPEG)
	Format ExportFormat
	// JPEGQuality specifies JPEG quality (1-100, only used for JPEG format)
	JPEGQuality int
}

// DefaultExportOptions returns default export options (PNG format)
func DefaultExportOptions() *ExportOptions {
	return &ExportOptions{
		Format:      FormatPNG,
		JPEGQuality: 90,
	}
}

// ExportGrayscale exports grayscale pixel data to an image format
// pixelData: raw pixel data
// width, height: image dimensions
// bitsAllocated, bitsStored: bit depth information
// isSigned: whether pixels are signed
// photometric: photometric interpretation (e.g., "MONOCHROME1", "MONOCHROME2")
func (e *ImageExporter) ExportGrayscale(
	writer io.Writer,
	pixelData []byte,
	width, height int,
	bitsAllocated, bitsStored int,
	isSigned bool,
	photometric string,
	options *ExportOptions,
) error {
	if options == nil {
		options = DefaultExportOptions()
	}

	// Create RGBA image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Apply pipeline LUT if available
	lut := e.pipeline.LUT()

	// Process pixels
	bytesPerPixel := bitsAllocated / 8
	isMonochrome1 := photometric == "MONOCHROME1"

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelIndex := (y*width + x) * bytesPerPixel

			var pixelValue float64
			if bytesPerPixel == 1 {
				pixelValue = float64(pixelData[pixelIndex])
			} else if bytesPerPixel == 2 {
				// 16-bit pixel
				if isSigned {
					pixelValue = float64(int16(pixelData[pixelIndex]) | int16(pixelData[pixelIndex+1])<<8)
				} else {
					pixelValue = float64(uint16(pixelData[pixelIndex]) | uint16(pixelData[pixelIndex+1])<<8)
				}
			}

			// Apply LUT transformation
			outputValue := lut.Transform(pixelValue)

			// Clamp to 8-bit range
			grayValue := uint8(clampUint8(outputValue))

			// Invert for MONOCHROME1
			if isMonochrome1 {
				grayValue = 255 - grayValue
			}

			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	return e.encodeImage(writer, img, options)
}

// ExportRGB exports RGB pixel data to an image format
func (e *ImageExporter) ExportRGB(
	writer io.Writer,
	pixelData []byte,
	width, height int,
	photometric string,
	planarConfig int,
	options *ExportOptions,
) error {
	if options == nil {
		options = DefaultExportOptions()
	}

	// Convert to RGB if needed
	rgbData, err := e.converter.ConvertToRGB(pixelData, width, height, photometric, planarConfig)
	if err != nil {
		return fmt.Errorf("failed to convert to RGB: %w", err)
	}

	// Create RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Copy pixel data
	if planarConfig == 0 || photometric != "RGB" {
		// Interleaved RGB or converted data
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				idx := (y*width + x) * 3
				r := rgbData[idx]
				g := rgbData[idx+1]
				b := rgbData[idx+2]
				img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
			}
		}
	} else {
		// Planar RGB: RRR... GGG... BBB...
		pixelCount := width * height
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				i := y*width + x
				r := rgbData[i]
				g := rgbData[pixelCount+i]
				b := rgbData[2*pixelCount+i]
				img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
			}
		}
	}

	return e.encodeImage(writer, img, options)
}

func (e *ImageExporter) encodeImage(writer io.Writer, img image.Image, options *ExportOptions) error {
	switch options.Format {
	case FormatPNG:
		encoder := png.Encoder{CompressionLevel: png.DefaultCompression}
		return encoder.Encode(writer, img)

	case FormatJPEG:
		quality := options.JPEGQuality
		if quality < 1 || quality > 100 {
			quality = 90
		}
		return jpeg.Encode(writer, img, &jpeg.Options{Quality: quality})

	default:
		return fmt.Errorf("unsupported export format: %d", options.Format)
	}
}

// RenderFrame renders a single frame from DicomPixelData to an image
// This is a high-level convenience function
func (e *ImageExporter) RenderFrame(
	writer io.Writer,
	frameData []byte,
	width, height int,
	bitsAllocated, bitsStored int,
	samplesPerPixel int,
	isSigned bool,
	photometric string,
	planarConfig int,
	options *ExportOptions,
) error {
	if samplesPerPixel == 1 {
		// Grayscale
		return e.ExportGrayscale(writer, frameData, width, height, bitsAllocated, bitsStored, isSigned, photometric, options)
	} else if samplesPerPixel == 3 {
		// RGB/YBR
		return e.ExportRGB(writer, frameData, width, height, photometric, planarConfig, options)
	} else {
		return fmt.Errorf("unsupported samples per pixel: %d", samplesPerPixel)
	}
}
