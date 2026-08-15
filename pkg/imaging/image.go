// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

// Package imaging provides image processing functionality for DICOM images
package imaging

import (
	"fmt"
	"image"
	"io"
	"math"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
	golangdraw "golang.org/x/image/draw"
)

const (
	monochrome2                          = "MONOCHROME2"
	transferSyntaxExplicitVRBigEndian    = "1.2.840.10008.1.2.2"
	transferSyntaxExplicitVRLittleEndian = "1.2.840.10008.1.2.1"
	transferSyntaxImplicitVRLittleEndian = "1.2.840.10008.1.2"
)

// DicomImage represents a DICOM image with rendering capabilities
type DicomImage struct {
	mu sync.RWMutex

	// Pixel data
	pixelData *DicomPixelData
	dataset   *dataset.Dataset

	// Current frame index
	currentFrame int

	// Rendering pipeline for each frame
	pipelines map[int]render.Pipeline

	// Scaling factor
	scale float64

	// Whether to show overlays
	showOverlays            bool
	overlays                []*DicomOverlayData
	overlayColor            imagetypes.Color32
	autoApplyLUTToAllFrames bool
	grayscaleColorMaps      map[int][256]imagetypes.Color32

	// Converter for pixel data format conversion
	converter *PixelDataConverter
}

// NewDicomImage creates a new DicomImage from DicomPixelData
func NewDicomImage(pixelData *DicomPixelData) *DicomImage {
	return &DicomImage{
		pixelData:               pixelData,
		currentFrame:            0,
		pipelines:               make(map[int]render.Pipeline),
		scale:                   1.0,
		showOverlays:            true,
		overlayColor:            imagetypes.Color32{A: 255, R: 255, B: 255},
		autoApplyLUTToAllFrames: false,
		grayscaleColorMaps:      make(map[int][256]imagetypes.Color32),
		converter:               NewPixelDataConverter(),
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
	return pi.Value == photometricMonochrome1 || pi.Value == monochrome2
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
	if scale <= 0 {
		return
	}
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
	if img.pixelData.Info.BitsStored == 0 {
		// DICOM requires BitsStored to be at least 1
		return render.NewGrayscalePipeline(1.0, 0, 256, 256, 0, 255, false)
	}

	// Use pixel representation to determine min/max input values
	var minInput, maxInput float64
	if img.pixelData.Info.PixelRepresentation == SignedPixels {
		rangeSize := math.Pow(2, float64(img.pixelData.Info.BitsStored-1))
		minInput = -rangeSize
		maxInput = rangeSize - 1
	} else {
		minInput = 0
		maxInput = math.Pow(2, float64(img.pixelData.Info.BitsStored)) - 1
	}

	// Calculate a fallback window from actual pixel data.
	windowCenter, windowWidth := 1.0, 1.0
	if img.pixelData.Info.BitsStored != 1 {
		windowCenter, windowWidth = img.pixelData.CalculateOptimalWindow()
	}
	rescaleSlope, rescaleIntercept := 1.0, 0.0
	var modalityLUT render.ModalityLUT
	hasExplicitWindow := false
	if img.dataset != nil {
		functional := imageFunctionalGroupValues(img.dataset, frame)
		if value, err := imageDecimalFrom(img.dataset, functional, tag.RescaleSlope); err == nil {
			rescaleSlope = value
		}
		if value, err := imageDecimalFrom(img.dataset, functional, tag.RescaleIntercept); err == nil {
			rescaleIntercept = value
		}
		if center, width, err := imageWindowPair(img.dataset, functional); err == nil {
			windowCenter, windowWidth = center, width
			hasExplicitWindow = true
		}
		modalityLUT, _ = imageModalityLUT(img.dataset, img.pixelData.Info.PixelRepresentation == SignedPixels)
	}
	if !hasExplicitWindow && img.dataset != nil && img.pixelData.Info.BitsStored != 1 {
		minimum, maximum, err := 0.0, 0.0, fmt.Errorf("image pixel value range is unavailable")
		if img.dataset != nil {
			minimum, maximum, err = imagePixelValueRange(img.dataset)
		}
		if err != nil {
			minimum, maximum, err = img.pixelData.MinMax(true)
		}
		if err == nil {
			if modalityLUT != nil {
				minimum = modalityLUT.Transform(minimum)
				maximum = modalityLUT.Transform(maximum)
			} else {
				minimum = minimum*rescaleSlope + rescaleIntercept
				maximum = maximum*rescaleSlope + rescaleIntercept
			}
			windowCenter = (minimum + maximum) / 2
			windowWidth = math.Max(1, math.Abs(maximum-minimum))
		}
	}

	pipeline := render.NewGrayscalePipeline(rescaleSlope, rescaleIntercept, windowCenter, windowWidth, minInput, maxInput, false)
	if img.dataset != nil {
		functional := imageFunctionalGroupValues(img.dataset, frame)
		if modalityLUT != nil {
			pipeline.SetModalityLUT(modalityLUT)
		}
		if voiLUT, err := imageVOILUTFrom(img.dataset, functional, img.pixelData.Info.PixelRepresentation == SignedPixels); err == nil {
			pipeline.SetVOILUT(voiLUT)
		}
		if function, ok := imageStringFrom(img.dataset, functional, tag.VOILUTFunction); ok {
			pipeline.SetVOILUTFunction(imagetypes.VOILUTFunction(function))
		}
	}
	img.pipelines[frame] = pipeline

	return pipeline
}

// SetPipeline sets the rendering pipeline for the specified frame.
// Custom pipelines should implement render.PipelineCloner when Clone must own
// an independent copy; other custom pipeline instances are retained by reference.
func (img *DicomImage) SetPipeline(frame int, pipeline render.Pipeline) error {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return fmt.Errorf("frame index out of range: %d", frame)
	}

	img.mu.Lock()
	defer img.mu.Unlock()
	img.pipelines[frame] = pipeline
	return nil
}

// currentFrameSafe returns the current frame index under read lock.
func (img *DicomImage) currentFrameSafe() int {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.currentFrame
}

// WindowWidth returns the window width for the current frame
func (img *DicomImage) WindowWidth() float64 {
	pipeline := img.GetOrCreatePipeline(img.currentFrameSafe())
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.WindowWidth()
	}
	return 256.0
}

// SetWindowWidth sets the window width for the current frame
func (img *DicomImage) SetWindowWidth(width float64) {
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if gp, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			gp.SetWindowWidth(width)
		}
	}
}

// WindowCenter returns the window center for the current frame
func (img *DicomImage) WindowCenter() float64 {
	pipeline := img.GetOrCreatePipeline(img.currentFrameSafe())
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.WindowCenter()
	}
	return 128.0
}

// SetWindowCenter sets the window center for the current frame
func (img *DicomImage) SetWindowCenter(center float64) {
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if gp, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			gp.SetWindowCenter(center)
		}
	}
}

// SetWindow sets both window center and width for the current frame
func (img *DicomImage) SetWindow(center, width float64) {
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if gp, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			gp.SetWindow(center, width)
		}
	}
}

// Invert returns whether the current frame is inverted
func (img *DicomImage) Invert() bool {
	pipeline := img.GetOrCreatePipeline(img.currentFrameSafe())
	if gp, ok := pipeline.(*render.GrayscalePipeline); ok {
		return gp.Invert()
	}
	return false
}

// SetInvert sets whether to invert the current frame
func (img *DicomImage) SetInvert(invert bool) {
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if gp, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			gp.SetInvert(invert)
		}
	}
}

// UseVOILUT reports whether the current frame uses its Dataset VOI LUT Sequence.
func (img *DicomImage) UseVOILUT() bool {
	if pipeline, ok := img.GetOrCreatePipeline(img.currentFrameSafe()).(*render.GrayscalePipeline); ok {
		return pipeline.UseVOILUT()
	}
	return false
}

// SetUseVOILUT selects Dataset VOI LUT Sequence rendering when available.
func (img *DicomImage) SetUseVOILUT(value bool) {
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if pipeline, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			pipeline.SetUseVOILUT(value)
		}
	}
}

// AutoApplyLUTToAllFrames reports whether caller LUT changes affect every frame.
func (img *DicomImage) AutoApplyLUTToAllFrames() bool {
	img.mu.RLock()
	defer img.mu.RUnlock()
	return img.autoApplyLUTToAllFrames
}

// SetAutoApplyLUTToAllFrames controls whether caller LUT changes affect every frame.
func (img *DicomImage) SetAutoApplyLUTToAllFrames(value bool) {
	img.mu.Lock()
	defer img.mu.Unlock()
	img.autoApplyLUTToAllFrames = value
}

func (img *DicomImage) lutFrameRange() (int, int) {
	img.mu.RLock()
	defer img.mu.RUnlock()
	if img.autoApplyLUTToAllFrames {
		return 0, img.NumberOfFrames() - 1
	}
	return img.currentFrame, img.currentFrame
}

// RenderFrame renders the specified frame to a writer in the specified format
func (img *DicomImage) RenderFrame(writer io.Writer, frame int, options *render.ExportOptions) error {
	rendered, err := img.RenderFrameImage(frame)
	if err != nil {
		return err
	}
	return render.NewImageExporter(img.GetOrCreatePipeline(frame)).ExportImage(writer, rendered, options)
}

// RenderFrameImage renders the specified frame and returns the unencoded Go image.
func (img *DicomImage) RenderFrameImage(frame int) (image.Image, error) {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return nil, fmt.Errorf("frame index out of range: %d", frame)
	}
	frameData, err := img.pixelData.GetFrame(frame)
	if err != nil {
		return nil, fmt.Errorf("failed to get frame data: %w", err)
	}
	exporter := render.NewImageExporter(img.GetOrCreatePipeline(frame))
	var rendered image.Image
	switch img.pixelData.Info.SamplesPerPixel {
	case 1:
		photometric := monochrome2
		if img.pixelData.Info.PhotometricInterpretation != nil {
			photometric = img.pixelData.Info.PhotometricInterpretation.Value
		}
		rendered, err = exporter.RenderGrayscaleImageWithBitDepth(
			frameData,
			int(img.Width()),
			int(img.Height()),
			int(img.pixelData.Info.BitsAllocated),
			int(img.pixelData.Info.BitsStored),
			int(img.pixelData.Info.HighBit),
			img.pixelData.Info.PixelRepresentation == SignedPixels,
			photometric,
		)
	case 3:
		photometric := photometricRGB
		if img.pixelData.Info.PhotometricInterpretation != nil {
			photometric = img.pixelData.Info.PhotometricInterpretation.Value
		}
		rendered, err = exporter.RenderRGBImage(
			frameData,
			int(img.Width()),
			int(img.Height()),
			photometric,
			int(img.pixelData.Info.PlanarConfiguration),
		)
	default:
		return nil, fmt.Errorf("unsupported samples per pixel: %d", img.pixelData.Info.SamplesPerPixel)
	}
	if err != nil {
		return nil, err
	}
	rendered = img.applyGrayscaleColorMap(rendered, frame)
	return img.scaleImage(img.applyOverlays(rendered, frame)), nil
}

func (img *DicomImage) scaleImage(source image.Image) image.Image {
	img.mu.RLock()
	scale := img.scale
	img.mu.RUnlock()
	if scale == 1 {
		return source
	}
	width := max(1, int(float64(source.Bounds().Dx())*scale))
	height := max(1, int(float64(source.Bounds().Dy())*scale))
	destinationBounds := image.Rect(0, 0, width, height)
	if _, ok := source.(*image.Gray); ok {
		gray := image.NewGray(destinationBounds)
		if img.pixelData.Info.BitsStored == 1 {
			golangdraw.NearestNeighbor.Scale(gray, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
		} else {
			golangdraw.BiLinear.Scale(gray, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
		}
		return gray
	}
	rgba := image.NewRGBA(destinationBounds)
	golangdraw.BiLinear.Scale(rgba, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
	return rgba
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
	if img.pixelData.Info.TransferSyntaxUID == transferSyntaxImplicitVRLittleEndian ||
		img.pixelData.Info.TransferSyntaxUID == transferSyntaxExplicitVRLittleEndian ||
		img.pixelData.Info.TransferSyntaxUID == transferSyntaxExplicitVRBigEndian {
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

// Clone copies image data and configuration. Pipelines implementing
// render.PipelineCloner are copied independently; other custom pipelines are
// retained by reference because the Pipeline interface has no general clone operation.
func (img *DicomImage) Clone() *DicomImage {
	img.mu.RLock()

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
		pixelData:               clonedPixelData,
		currentFrame:            img.currentFrame,
		pipelines:               make(map[int]render.Pipeline),
		scale:                   img.scale,
		showOverlays:            img.showOverlays,
		overlays:                append([]*DicomOverlayData(nil), img.overlays...),
		overlayColor:            img.overlayColor,
		autoApplyLUTToAllFrames: img.autoApplyLUTToAllFrames,
		grayscaleColorMaps:      make(map[int][256]imagetypes.Color32, len(img.grayscaleColorMaps)),
		converter:               NewPixelDataConverter(),
	}
	if img.dataset != nil {
		cloned.dataset = img.dataset.Clone()
	}
	for frame, colorMap := range img.grayscaleColorMaps {
		cloned.grayscaleColorMaps[frame] = colorMap
	}
	pipelines := make(map[int]render.Pipeline, len(img.pipelines))
	for frame, pipeline := range img.pipelines {
		pipelines[frame] = pipeline
	}
	img.mu.RUnlock()

	for frame, pipeline := range pipelines {
		if cloneable, ok := pipeline.(render.PipelineCloner); ok {
			cloned.pipelines[frame] = cloneable.ClonePipeline()
		} else {
			// Preserve custom pipeline behavior when it does not expose a clone contract.
			cloned.pipelines[frame] = pipeline
		}
	}

	return cloned
}
