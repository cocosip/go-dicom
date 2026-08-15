// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

// ImageExporter provides functionality to export DICOM images to standard formats
type ImageExporter struct {
	pipeline  Pipeline
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
	bitsAllocated, _ int,
	isSigned bool,
	photometric string,
	options *ExportOptions,
) error {
	if options == nil {
		options = DefaultExportOptions()
	}
	img, err := e.RenderGrayscaleImage(pixelData, width, height, bitsAllocated, isSigned, photometric)
	if err != nil {
		return err
	}
	return e.encodeImage(writer, img, options)
}

// RenderGrayscaleImage renders grayscale pixels into an 8-bit image without encoding it.
func (e *ImageExporter) RenderGrayscaleImage(
	pixelData []byte,
	width, height int,
	bitsAllocated int,
	isSigned bool,
	photometric string,
) (*image.Gray, error) {
	return e.RenderGrayscaleImageWithBitDepth(
		pixelData, width, height, bitsAllocated, bitsAllocated, bitsAllocated-1, isSigned, photometric,
	)
}

// RenderGrayscaleImageWithBitDepth renders grayscale pixels using the stored-bit location.
func (e *ImageExporter) RenderGrayscaleImageWithBitDepth(
	pixelData []byte,
	width, height int,
	bitsAllocated, bitsStored, highBit int,
	isSigned bool,
	photometric string,
) (*image.Gray, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("image dimensions must be positive")
	}
	if bitsAllocated != 1 && bitsAllocated != 8 && bitsAllocated != 16 && bitsAllocated != 32 {
		return nil, fmt.Errorf("unsupported grayscale bits allocated: %d", bitsAllocated)
	}
	if bitsAllocated == 1 {
		required := (width*height + 7) / 8
		if len(pixelData) < required {
			return nil, fmt.Errorf("grayscale pixel data too short: got %d bytes, need %d", len(pixelData), required)
		}
		img := image.NewGray(image.Rect(0, 0, width, height))
		pipelineLUT := e.pipeline.LUT()
		isMonochrome1 := photometric == "MONOCHROME1"
		for index := 0; index < width*height; index++ {
			pixelValue := float64((pixelData[index/8] >> uint(index%8)) & 1)
			grayValue := clampUint8(pipelineLUT.Transform(pixelValue))
			if isMonochrome1 {
				grayValue = 255 - grayValue
			}
			img.SetGray(index%width, index/width, color.Gray{Y: grayValue})
		}
		return img, nil
	}
	bytesPerPixel := bitsAllocated / 8
	required := width * height * bytesPerPixel
	if len(pixelData) < required {
		return nil, fmt.Errorf("grayscale pixel data too short: got %d bytes, need %d", len(pixelData), required)
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	// Apply pipeline LUT if available
	lut := e.pipeline.LUT()

	// Process pixels
	isMonochrome1 := photometric == "MONOCHROME1"

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelIndex := (y*width + x) * bytesPerPixel

			var raw uint64
			switch bytesPerPixel {
			case 1:
				raw = uint64(pixelData[pixelIndex])
			case 2:
				raw = uint64(binary.LittleEndian.Uint16(pixelData[pixelIndex:]))
			case 4:
				raw = uint64(binary.LittleEndian.Uint32(pixelData[pixelIndex:]))
			}
			pixelValue := storedPixelValue(raw, bitsAllocated, bitsStored, highBit, isSigned)

			// Apply LUT transformation
			outputValue := lut.Transform(pixelValue)

			// Clamp to 8-bit range
			grayValue := clampUint8(outputValue)

			// Invert for MONOCHROME1
			if isMonochrome1 {
				grayValue = 255 - grayValue
			}

			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	return img, nil
}

func storedPixelValue(raw uint64, bitsAllocated, bitsStored, highBit int, signed bool) float64 {
	if bitsStored <= 0 || bitsStored > bitsAllocated {
		bitsStored = bitsAllocated
	}
	if highBit+1 < bitsStored || highBit >= bitsAllocated {
		highBit = bitsStored - 1
	}
	shift := highBit + 1 - bitsStored
	mask := uint64(1)<<uint(bitsStored) - 1
	value := (raw >> uint(shift)) & mask
	if signed {
		signBit := uint64(1) << uint(bitsStored-1)
		if value&signBit != 0 {
			return float64(int64(value) - int64(uint64(1)<<uint(bitsStored)))
		}
	}
	return float64(value)
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
	img, err := e.RenderRGBImage(pixelData, width, height, photometric, planarConfig)
	if err != nil {
		return err
	}
	return e.encodeImage(writer, img, options)
}

// RenderRGBImage converts supported color pixels into an RGBA image without encoding it.
func (e *ImageExporter) RenderRGBImage(
	pixelData []byte,
	width, height int,
	photometric string,
	planarConfig int,
) (*image.RGBA, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("image dimensions must be positive")
	}

	// Convert to RGB if needed
	rgbData, err := e.converter.ConvertToRGB(pixelData, width, height, photometric, planarConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to RGB: %w", err)
	}
	required := width * height * 3
	if len(rgbData) < required {
		return nil, fmt.Errorf("RGB pixel data too short: got %d bytes, need %d", len(rgbData), required)
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

	return img, nil
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

// ExportImage encodes an already rendered image using the requested format.
func (e *ImageExporter) ExportImage(writer io.Writer, img image.Image, options *ExportOptions) error {
	if options == nil {
		options = DefaultExportOptions()
	}
	return e.encodeImage(writer, img, options)
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
	switch samplesPerPixel {
	case 1:
		// Grayscale
		return e.ExportGrayscale(writer, frameData, width, height, bitsAllocated, bitsStored, isSigned, photometric, options)
	case 3:
		// RGB/YBR
		return e.ExportRGB(writer, frameData, width, height, photometric, planarConfig, options)
	default:
		return fmt.Errorf("unsupported samples per pixel: %d", samplesPerPixel)
	}
}
