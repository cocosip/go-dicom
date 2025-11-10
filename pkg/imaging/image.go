// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"
	"io"
	"sync"

	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

// DicomImage represents a DICOM image with rendering capabilities
type DicomImage struct {
	mu sync.RWMutex

	// Pixel data
	pixelData *DicomPixelData

	// Current frame index
	currentFrame int

	// Rendering pipeline for each frame
	pipelines map[int]render.Pipeline

	// Scaling factor
	scale float64

	// Whether to show overlays
	showOverlays bool

	// Converter for pixel data format conversion
	converter *PixelDataConverter
}

// NewDicomImage creates a new DicomImage from DicomPixelData
func NewDicomImage(pixelData *DicomPixelData) *DicomImage {
	return &DicomImage{
		pixelData:    pixelData,
		currentFrame: 0,
		pipelines:    make(map[int]render.Pipeline),
		scale:        1.0,
		showOverlays: true,
		converter:    NewPixelDataConverter(),
	}
}

// Width returns the image width in pixels
func (img *DicomImage) Width() uint16 {
	return img.pixelData.Info.Width
}

// Height returns the image height in pixels
func (img *DicomImage) Height() uint16 {
	return img.pixelData.Info.Height
}

// NumberOfFrames returns the number of frames in the image
func (img *DicomImage) NumberOfFrames() int {
	return img.pixelData.Info.NumberOfFrames
}

// IsGrayscale returns true if the image is grayscale
func (img *DicomImage) IsGrayscale() bool {
	pi := img.pixelData.Info.PhotometricInterpretation
	if pi == nil {
		return true // Default to grayscale
	}
	return pi.Value == "MONOCHROME1" || pi.Value == "MONOCHROME2"
}

// CurrentFrame returns the current frame index
func (img *DicomImage) CurrentFrame() int {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.currentFrame
}

// SetCurrentFrame sets the current frame index
func (img *DicomImage) SetCurrentFrame(frame int) error {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return fmt.Errorf("frame index out of range: %d (total: %d)", frame, img.NumberOfFrames())
	}

	img.mu.Lock()
	defer img.mu.Unlock()
	img.currentFrame = frame
	return nil
}

// Scale returns the scaling factor
func (img *DicomImage) Scale() float64 {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.scale
}

// SetScale sets the scaling factor
func (img *DicomImage) SetScale(scale float64) {
	img.mu.Lock()
	defer img.mu.Unlock()
	img.scale = scale
}

// GetOrCreatePipeline returns the rendering pipeline for the specified frame
// Creates a default pipeline if it doesn't exist
func (img *DicomImage) GetOrCreatePipeline(frame int) render.Pipeline {
	img.mu.Lock()
	defer img.mu.Unlock()

	if pipeline, exists := img.pipelines[frame]; exists {
		return pipeline
	}

	// Create default pipeline
	// Use pixel representation to determine min/max input values
	var minInput, maxInput float64
	if img.pixelData.Info.PixelRepresentation == SignedPixels {
		// Signed pixels
		maxVal := int16(1<<(img.pixelData.Info.BitsStored-1) - 1)
		minVal := -maxVal - 1
		minInput = float64(minVal)
		maxInput = float64(maxVal)
	} else {
		// Unsigned pixels
		minInput = 0
		maxInput = float64(uint16(1<<img.pixelData.Info.BitsStored - 1))
	}

	// Default window center and width (full range)
	windowCenter := (maxInput + minInput) / 2
	windowWidth := maxInput - minInput

	// Default rescale: slope=1, intercept=0
	pipeline := render.NewGrayscalePipeline(1.0, 0, windowCenter, windowWidth, minInput, maxInput, false)
	img.pipelines[frame] = pipeline

	return pipeline
}

// SetPipeline sets the rendering pipeline for the specified frame
func (img *DicomImage) SetPipeline(frame int, pipeline render.Pipeline) error {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return fmt.Errorf("frame index out of range: %d", frame)
	}

	img.mu.Lock()
	defer img.mu.Unlock()
	img.pipelines[frame] = pipeline
	return nil
}

// WindowWidth returns the window width for the current frame
func (img *DicomImage) WindowWidth() float64 {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.WindowWidth()
	}
	return 256.0
}

// SetWindowWidth sets the window width for the current frame
func (img *DicomImage) SetWindowWidth(width float64) {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		gp.SetWindowWidth(width)
	}
}

// WindowCenter returns the window center for the current frame
func (img *DicomImage) WindowCenter() float64 {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.WindowCenter()
	}
	return 128.0
}

// SetWindowCenter sets the window center for the current frame
func (img *DicomImage) SetWindowCenter(center float64) {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		gp.SetWindowCenter(center)
	}
}

// SetWindow sets both window center and width for the current frame
func (img *DicomImage) SetWindow(center, width float64) {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		gp.SetWindow(center, width)
	}
}

// Invert returns whether the current frame is inverted
func (img *DicomImage) Invert() bool {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.Invert()
	}
	return false
}

// SetInvert sets whether to invert the current frame
func (img *DicomImage) SetInvert(invert bool) {
	pipeline := img.GetOrCreatePipeline(img.currentFrame)
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		gp.SetInvert(invert)
	}
}

// RenderFrame renders the specified frame to a writer in the specified format
func (img *DicomImage) RenderFrame(writer io.Writer, frame int, options *render.ExportOptions) error {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return fmt.Errorf("frame index out of range: %d", frame)
	}

	// Get frame data
	frameData, err := img.pixelData.GetFrame(frame)
	if err != nil {
		return fmt.Errorf("failed to get frame data: %w", err)
	}

	// Get pipeline for this frame
	pipeline := img.GetOrCreatePipeline(frame)

	// Create exporter
	exporter := render.NewImageExporter(pipeline)

	// Render based on samples per pixel
	switch img.pixelData.Info.SamplesPerPixel {
	case 1:
		// Grayscale
		isSigned := img.pixelData.Info.PixelRepresentation == SignedPixels
		photometric := "MONOCHROME2"
		if img.pixelData.Info.PhotometricInterpretation != nil {
			photometric = img.pixelData.Info.PhotometricInterpretation.Value
		}

		return exporter.ExportGrayscale(
			writer,
			frameData,
			int(img.Width()),
			int(img.Height()),
			int(img.pixelData.Info.BitsAllocated),
			int(img.pixelData.Info.BitsStored),
			isSigned,
			photometric,
			options,
		)
	case 3:
		// RGB/YBR
		photometric := "RGB"
		if img.pixelData.Info.PhotometricInterpretation != nil {
			photometric = img.pixelData.Info.PhotometricInterpretation.Value
		}

		return exporter.ExportRGB(
			writer,
			frameData,
			int(img.Width()),
			int(img.Height()),
			photometric,
			int(img.pixelData.Info.PlanarConfiguration),
			options,
		)
	default:
		return fmt.Errorf("unsupported samples per pixel: %d", img.pixelData.Info.SamplesPerPixel)
	}
}

// RenderCurrentFrame renders the current frame to a writer
func (img *DicomImage) RenderCurrentFrame(writer io.Writer, options *render.ExportOptions) error {
	img.mu.RLock()
	frame := img.currentFrame
	img.mu.RUnlock()
	return img.RenderFrame(writer, frame, options)
}

// DecodeIfNeeded decodes the pixel data if it's in a compressed format
func (img *DicomImage) DecodeIfNeeded(c codec.Codec, params codec.Parameters) error {
	// Check if already uncompressed
	if img.pixelData.Info.TransferSyntaxUID == "1.2.840.10008.1.2" || // Implicit VR Little Endian
		img.pixelData.Info.TransferSyntaxUID == "1.2.840.10008.1.2.1" || // Explicit VR Little Endian
		img.pixelData.Info.TransferSyntaxUID == "1.2.840.10008.1.2.2" { // Explicit VR Big Endian
		return nil // Already uncompressed
	}

	// Decode
	decoded, err := img.pixelData.Decode(c, params)
	if err != nil {
		return fmt.Errorf("failed to decode pixel data: %w", err)
	}

	img.mu.Lock()
	img.pixelData = decoded
	// Clear cached pipelines as pixel data has changed
	img.pipelines = make(map[int]render.Pipeline)
	img.mu.Unlock()

	return nil
}

// Clone creates a deep copy of the DicomImage
func (img *DicomImage) Clone() *DicomImage {
	img.mu.RLock()
	defer img.mu.RUnlock()

	// Clone pixel data
	clonedPixelData := &DicomPixelData{
		Info:   img.pixelData.Info, // Info can be shared (read-only)
		frames: make([][]byte, len(img.pixelData.frames)),
	}

	for i, frame := range img.pixelData.frames {
		clonedFrame := make([]byte, len(frame))
		copy(clonedFrame, frame)
		clonedPixelData.frames[i] = clonedFrame
	}

	// Create new image
	cloned := &DicomImage{
		pixelData:    clonedPixelData,
		currentFrame: img.currentFrame,
		pipelines:    make(map[int]render.Pipeline),
		scale:        img.scale,
		showOverlays: img.showOverlays,
		converter:    NewPixelDataConverter(),
	}

	// Note: Pipelines are not cloned, they will be recreated on demand

	return cloned
}
