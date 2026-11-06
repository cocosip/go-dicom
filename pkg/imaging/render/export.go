// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package render

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// ImageExporter provides functionality to export DICOM images to standard formats
type ImageExporter struct {
	pipeline Pipeline
}

// NewImageExporter creates a new ImageExporter
func NewImageExporter(pipeline Pipeline) *ImageExporter {
	return &ImageExporter{
		pipeline: pipeline,
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
//
// Conformant stored values are right-justified, so the high bit is bitsStored-1.
// Use ExportGrayscaleWithHighBit for files that declare a different High Bit.
func (e *ImageExporter) ExportGrayscale(
	writer io.Writer,
	pixelData []byte,
	width, height int,
	bitsAllocated, bitsStored int,
	isSigned bool,
	photometric string,
	options *ExportOptions,
) error {
	return e.ExportGrayscaleWithHighBit(writer, pixelData, width, height,
		bitsAllocated, bitsStored, bitsStored-1, isSigned, photometric, options)
}

// ExportGrayscaleWithHighBit exports grayscale pixel data, honoring an explicit
// High Bit (0028,0102) position.
func (e *ImageExporter) ExportGrayscaleWithHighBit(
	writer io.Writer,
	pixelData []byte,
	width, height int,
	bitsAllocated, bitsStored, highBit int,
	isSigned bool,
	photometric string,
	options *ExportOptions,
) error {
	if options == nil {
		options = DefaultExportOptions()
	}
	img, err := e.RenderGrayscaleImageWithBitDepth(
		pixelData, width, height, bitsAllocated, bitsStored, highBit, isSigned, photometric,
	)
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
		invertOutput := photometric == "MONOCHROME1" && !pipelineInverts(e.pipeline)
		for index := 0; index < width*height; index++ {
			pixelValue := float64((pixelData[index/8] >> uint(index%8)) & 1)
			grayValue := clampRenderByte(pipelineLUT.Transform(pixelValue))
			if invertOutput {
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

	// Apply pipeline LUT if available. Precalculate it over the stored value
	// domain so each pixel becomes a direct table lookup instead of a per-pixel
	// interface call; very wide samples fall back to the source LUT.
	pipelineLUT := e.pipeline.LUT()
	// MONOCHROME1 displays minimum sample values as white. A Presentation LUT
	// Shape of INVERSE (or an explicit SetInvert) encodes that same display
	// polarity, so the photometric inversion must not be applied on top of it.
	invertOutput := photometric == "MONOCHROME1" && !pipelineInverts(e.pipeline)
	effectiveStored := bitsStored
	if effectiveStored <= 0 || effectiveStored > bitsAllocated {
		effectiveStored = bitsAllocated
	}
	precalc := newPrecalculatedLUT(pipelineLUT, effectiveStored, isSigned)
	// Building the table costs one transform per stored value; it only pays off
	// once the frame has at least that many pixels.
	if precalc != nil && width*height < 1<<uint(effectiveStored) {
		precalc = nil
	}

	for y := 0; y < height; y++ {
		row := y * img.Stride
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

			var grayValue uint8
			if precalc != nil {
				grayValue = clampRenderByte(precalc.Transform(pixelValue))
			} else {
				grayValue = clampRenderByte(pipelineLUT.Transform(pixelValue))
			}

			// Invert for MONOCHROME1
			if invertOutput {
				grayValue = 255 - grayValue
			}

			img.Pix[row+x] = grayValue
		}
	}

	return img, nil
}

// pipelineInverts reports whether the pipeline output is already inverted.
func pipelineInverts(p Pipeline) bool {
	if inverting, ok := p.(interface{ Inverts() bool }); ok {
		return inverting.Inverts()
	}
	return false
}

// newPrecalculatedLUT builds a table-backed LUT over the stored value domain.
// It returns nil for very wide samples, where the table would be too large, so
// the caller falls back to per-pixel transforms.
func newPrecalculatedLUT(source lut.LUT, bitsStored int, isSigned bool) *lut.PrecalculatedLUT {
	if source == nil || bitsStored <= 0 || bitsStored > 16 {
		return nil
	}
	if isSigned {
		half := 1 << uint(bitsStored-1)
		return lut.NewPrecalculatedLUT(source, -half, half-1)
	}
	return lut.NewPrecalculatedLUT(source, 0, (1<<uint(bitsStored))-1)
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

// RenderRGBImage converts supported 8-bit color pixels into an RGBA image without encoding it.
func (e *ImageExporter) RenderRGBImage(
	pixelData []byte,
	width, height int,
	photometric string,
	planarConfig int,
) (*image.RGBA, error) {
	return e.RenderRGBImageWithBytesPerSample(pixelData, width, height, photometric, planarConfig, 1)
}

// RenderRGBImageWithBytesPerSample converts supported color pixels with 1- or
// 2-byte samples into an RGBA image without encoding it. Two-byte samples are
// reduced to 8-bit (see colorconv.ConvertToRGBWithBytesPerSample).
func (e *ImageExporter) RenderRGBImageWithBytesPerSample(
	pixelData []byte,
	width, height int,
	photometric string,
	planarConfig int,
	bytesPerSample int,
) (*image.RGBA, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("image dimensions must be positive")
	}

	// Convert to 8-bit RGB if needed
	rgbData, err := colorconv.ConvertToRGBWithBytesPerSample(
		pixelData,
		width,
		height,
		pixel.PhotometricInterpretation{Value: photometric},
		pixel.PlanarConfiguration(planarConfig),
		bytesPerSample,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to RGB: %w", err)
	}
	required := width * height * 3
	if len(rgbData) < required {
		return nil, fmt.Errorf("RGB pixel data too short: got %d bytes, need %d", len(rgbData), required)
	}

	// colorconv.ConvertToRGBWithBytesPerSample always returns interleaved 8-bit RGB.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for index := 0; index < width*height; index++ {
		src := index * 3
		dst := index * 4
		img.Pix[dst] = rgbData[src]
		img.Pix[dst+1] = rgbData[src+1]
		img.Pix[dst+2] = rgbData[src+2]
		img.Pix[dst+3] = 255
	}

	return img, nil
}

func clampRenderByte(value float64) uint8 {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return uint8(math.Round(value))
}

// RenderRGBAImage renders interleaved straight-alpha RGBA pixel data.
func (e *ImageExporter) RenderRGBAImage(pixelData []byte, width, height int) (*image.NRGBA, error) {
	return e.RenderRGBAImageWithConfig(pixelData, width, height, "RGBA", int(pixel.InterleavedPlanar))
}

// RenderRGBAImageWithConfig renders 8-bit four-sample pixel data, honoring the
// sample order from photometric ("ARGB" is A,R,G,B; anything else is treated as
// straight-alpha R,G,B,A) and planar configuration. The result is NRGBA
// (straight alpha, R,G,B,A order).
func (e *ImageExporter) RenderRGBAImageWithConfig(
	pixelData []byte,
	width, height int,
	photometric string,
	planarConfig int,
) (*image.NRGBA, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("image dimensions must be positive")
	}
	required := width * height * 4
	if len(pixelData) < required {
		return nil, fmt.Errorf("RGBA pixel data too short: got %d bytes, need %d", len(pixelData), required)
	}

	interleaved := pixelData[:required]
	if planarConfig == int(pixel.PlanarPlanar) {
		converted, err := colorconv.PlanarToInterleaved(interleaved, 4, 1)
		if err != nil {
			return nil, fmt.Errorf("failed to interleave planar RGBA: %w", err)
		}
		interleaved = converted
	} else if planarConfig != int(pixel.InterleavedPlanar) {
		return nil, fmt.Errorf("invalid planar configuration %d", planarConfig)
	}

	result := image.NewNRGBA(image.Rect(0, 0, width, height))
	if photometric == pixel.ARGBPhotometric.Value {
		// Store as A,R,G,B; reorder to R,G,B,A for NRGBA.
		for index := 0; index < width*height; index++ {
			src := index * 4
			dst := index * 4
			result.Pix[dst] = interleaved[src+1]
			result.Pix[dst+1] = interleaved[src+2]
			result.Pix[dst+2] = interleaved[src+3]
			result.Pix[dst+3] = interleaved[src]
		}
		return result, nil
	}
	copy(result.Pix, interleaved)
	return result, nil
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

// RenderFrame renders a single pixel-data frame to an image.
// This is a high-level convenience function.
//
// Conformant stored values are right-justified, so the high bit is bitsStored-1.
// Use RenderFrameWithHighBit for files that declare a different High Bit.
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
	return e.RenderFrameWithHighBit(writer, frameData, width, height,
		bitsAllocated, bitsStored, bitsStored-1, samplesPerPixel, isSigned, photometric, planarConfig, options)
}

// RenderFrameWithHighBit renders a single pixel-data frame, honoring an explicit
// High Bit (0028,0102) position.
func (e *ImageExporter) RenderFrameWithHighBit(
	writer io.Writer,
	frameData []byte,
	width, height int,
	bitsAllocated, bitsStored, highBit int,
	samplesPerPixel int,
	isSigned bool,
	photometric string,
	planarConfig int,
	options *ExportOptions,
) error {
	switch samplesPerPixel {
	case 1:
		// Grayscale
		return e.ExportGrayscaleWithHighBit(writer, frameData, width, height, bitsAllocated, bitsStored, highBit, isSigned, photometric, options)
	case 3:
		// RGB/YBR
		img, err := e.RenderRGBImageWithBytesPerSample(
			frameData, width, height, photometric, planarConfig, (bitsAllocated+7)/8,
		)
		if err != nil {
			return err
		}
		return e.ExportImage(writer, img, options)
	case 4:
		img, err := e.RenderRGBAImageWithConfig(frameData, width, height, photometric, planarConfig)
		if err != nil {
			return err
		}
		return e.ExportImage(writer, img, options)
	default:
		return fmt.Errorf("unsupported samples per pixel: %d", samplesPerPixel)
	}
}
