// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

// Package imaging provides image processing functionality for DICOM images
package imaging

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"io"
	"log/slog"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
	"github.com/cocosip/go-dicom/pkg/imaging/interpolation"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
	"github.com/cocosip/go-dicom/pkg/logging"
	golangdraw "golang.org/x/image/draw"
)

// FrameRenderOptions controls optional final-frame transforms and graphics.
// A zero value preserves RenderFrameImage's existing output.
type FrameRenderOptions struct {
	SpatialTransform *transform.SpatialTransform
	Interpolation    interpolation.Mode
	Background       color.Color
	Viewport         image.Rectangle
	Graphics         []render.Graphic
}

const (
	transferSyntaxExplicitVRBigEndian    = "1.2.840.10008.1.2.2"
	transferSyntaxExplicitVRLittleEndian = "1.2.840.10008.1.2.1"
	transferSyntaxImplicitVRLittleEndian = "1.2.840.10008.1.2"
)

// DicomImage represents a DICOM image with rendering capabilities.
//
// The pixel data pointer is replaced atomically by DecodeIfNeeded, so readers
// such as Width and RenderFrame may run while another goroutine decodes.
// Other render settings (window, LUT selections, pipelines) are guarded by mu;
// configure them from one goroutine and do not mutate a pipeline that another
// goroutine is rendering.
type DicomImage struct {
	mu sync.RWMutex

	// Pixel data; replaced atomically when decoded.
	pixelData atomic.Pointer[pixeldata.Data]
	// decodeMu serializes DecodeIfNeeded calls so a compressed source is decoded
	// once and a second caller observes the already-uncompressed result instead
	// of overwriting the first caller's output.
	decodeMu          sync.Mutex
	dataset           *dataset.Dataset
	presentationState *dataset.Dataset

	// Current frame index
	currentFrame int

	// Rendering pipeline for each frame
	pipelines map[int]render.Pipeline

	// windowRange caches the pixel-value range deriveWindowFromPixels needs.
	// The range spans every frame, so one scan serves all frame pipelines;
	// DecodeIfNeeded drops it when the pixel data is replaced.
	windowRange          [2]float64
	windowRangeComputed  bool
	windowRangeErr       error
	windowRangeScanCount int // pixel scans performed; lets tests prove memoization

	// Scaling factor
	scale float64

	// Whether to show overlays
	showOverlays            bool
	overlays                []*DicomOverlayData
	overlayColor            colorconv.Color32
	autoApplyLUTToAllFrames bool
	grayscaleColorMaps      map[int][256]colorconv.Color32
	windowIndex             int
	voiLUTIndex             int
	lutVRMode               pixeldata.LUTVRMode
}

// NewDicomImage creates a new DicomImage from pixel data.
func NewDicomImage(pixelData *pixeldata.Data) *DicomImage {
	img := &DicomImage{
		currentFrame:            0,
		pipelines:               make(map[int]render.Pipeline),
		scale:                   1.0,
		showOverlays:            true,
		overlayColor:            colorconv.Color32{A: 255, R: 255, G: 255, B: 255},
		autoApplyLUTToAllFrames: false,
		grayscaleColorMaps:      make(map[int][256]colorconv.Color32),
		lutVRMode:               pixeldata.LUTCompatible,
	}
	img.pixelData.Store(pixelData)
	return img
}

// Width returns the image width in pixels
func (img *DicomImage) Width() uint16 {
	if pd := img.pixelData.Load(); pd != nil && pd.Info != nil {
		return pd.Info.Width
	}
	return 0
}

// Height returns the image height in pixels
func (img *DicomImage) Height() uint16 {
	if pd := img.pixelData.Load(); pd != nil && pd.Info != nil {
		return pd.Info.Height
	}
	return 0
}

// NumberOfFrames returns the number of frames in the image
func (img *DicomImage) NumberOfFrames() int {
	if pd := img.pixelData.Load(); pd != nil && pd.Info != nil {
		return pd.Info.NumberOfFrames
	}
	return 0
}

// IsGrayscale returns true if the image is grayscale
func (img *DicomImage) IsGrayscale() bool {
	pd := img.pixelData.Load()
	if pd == nil || pd.Info == nil {
		return true
	}
	pi := pd.Info.PhotometricInterpretation
	if pi == nil {
		return true // Default to grayscale
	}
	return pi.Value == pixel.Monochrome1.Value || pi.Value == pixel.Monochrome2.Value
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

// GetOrCreatePipeline returns the rendering pipeline for the specified frame.
//
// It creates a default pipeline when none exists and falls back to an identity
// pipeline when the dataset's attributes cannot be validated, so it always
// returns a usable pipeline. Dataset rendering uses
// GetOrCreatePipelineWithError to surface malformed attributes instead.
func (img *DicomImage) GetOrCreatePipeline(frame int) render.Pipeline {
	pipeline, err := img.GetOrCreatePipelineWithError(frame)
	if err == nil {
		return pipeline
	}
	return img.newFallbackPipeline()
}

// GetOrCreatePipelineWithError returns a standards-validated rendering
// pipeline. Dataset rendering uses this path so malformed modality and VOI
// attributes are surfaced instead of silently falling back.
func (img *DicomImage) GetOrCreatePipelineWithError(frame int) (render.Pipeline, error) {
	if frame < 0 || frame >= img.NumberOfFrames() {
		return nil, fmt.Errorf("frame index out of range: %d", frame)
	}
	img.mu.Lock()
	defer img.mu.Unlock()

	if pipeline, exists := img.pipelines[frame]; exists {
		return pipeline, nil
	}
	pipeline, err := img.createGrayscalePipeline(frame)
	if err != nil {
		return nil, err
	}
	img.pipelines[frame] = pipeline
	return pipeline, nil
}

func (img *DicomImage) newFallbackPipeline() render.Pipeline {
	pd := img.pixelData.Load()
	// An image constructed without pixel data has no attributes to derive a
	// window from; a neutral pipeline keeps the public setters usable.
	if pd == nil || pd.Info == nil || pd.Info.BitsStored == 0 {
		return render.NewGrayscalePipeline(1.0, 0, 256, 256, 0, 255, false)
	}
	pixelSigned := pd.Info.PixelRepresentation == pixel.SignedPixels
	minInput, maxInput := imageStoredValueRange(pd.Info.BitsStored, pixelSigned)
	windowCenter, windowWidth := defaultWindowCenter(pd.Info.BitsStored, pixelSigned), 1.0
	if pd.Info.BitsStored != 1 {
		windowCenter, windowWidth = pd.CalculateOptimalWindow()
	}
	return render.NewGrayscalePipeline(1, 0, windowCenter, windowWidth, minInput, maxInput, false)
}

// defaultWindowCenter returns the fallback window center used when no explicit
// VOI is available. One-bit signed pixels hold -1 and 0; a window centered at 1
// clips both to the minimum, so center it on their midpoint instead.
func defaultWindowCenter(bitsStored uint16, signed bool) float64 {
	if bitsStored == 1 && signed {
		return -0.5
	}
	return 1
}

func imageStoredValueRange(bitsStored uint16, signed bool) (float64, float64) {
	if signed {
		rangeSize := math.Pow(2, float64(bitsStored-1))
		return -rangeSize, rangeSize - 1
	}
	return 0, math.Pow(2, float64(bitsStored)) - 1
}

// grayscaleTransforms holds the modality and VOI transformation inputs resolved
// for one frame.
type grayscaleTransforms struct {
	modalityLUT         lut.LUT
	rescaleSlope        float64
	rescaleIntercept    float64
	voiDescriptorSigned bool
	voiPrimary          *dataset.Dataset
	voiFallback         *dataset.Dataset
	windowCenter        float64
	windowWidth         float64
	hasExplicitWindow   bool
}

func (img *DicomImage) createGrayscalePipeline(frame int) (*render.GrayscalePipeline, error) {
	if img.pixelData.Load().Info.BitsStored == 0 {
		return render.NewGrayscalePipeline(1.0, 0, 256, 256, 0, 255, false), nil
	}

	pixelSigned := img.pixelData.Load().Info.PixelRepresentation == pixel.SignedPixels
	minInput, maxInput := imageStoredValueRange(img.pixelData.Load().Info.BitsStored, pixelSigned)
	presentationStateApplied := img.presentationState != nil

	transforms, err := img.resolveGrayscaleTransforms(frame, presentationStateApplied, pixelSigned, minInput, maxInput)
	if err != nil {
		return nil, err
	}

	pipeline := render.NewGrayscalePipeline(
		transforms.rescaleSlope, transforms.rescaleIntercept,
		transforms.windowCenter, transforms.windowWidth, minInput, maxInput, false,
	)
	if img.dataset != nil {
		if err := img.configurePipelineLUTs(pipeline, transforms, presentationStateApplied); err != nil {
			return nil, err
		}
	}
	function := lut.VOILUTFunctionLinear
	if img.dataset != nil {
		if value, ok := imageStringFrom(transforms.voiPrimary, transforms.voiFallback, tag.VOILUTFunction); ok {
			function = lut.VOILUTFunction(value)
		}
	}
	if _, err := lut.CreateValidatedVOILUT(function, transforms.windowCenter, transforms.windowWidth); err != nil {
		return nil, err
	}
	return pipeline, nil
}

// resolveGrayscaleTransforms resolves the modality and VOI transformations for
// one frame, including the window that the render pipeline should use.
func (img *DicomImage) resolveGrayscaleTransforms(
	frame int, presentationStateApplied, pixelSigned bool, minInput, maxInput float64,
) (grayscaleTransforms, error) {
	bitsStored := img.pixelData.Load().Info.BitsStored
	windowCenter, windowWidth := defaultWindowCenter(bitsStored, pixelSigned), 1.0
	transforms := grayscaleTransforms{
		rescaleSlope:        1.0,
		rescaleIntercept:    0.0,
		windowCenter:        windowCenter,
		windowWidth:         windowWidth,
		voiDescriptorSigned: pixelSigned,
	}

	// Without a Presentation State, fall back to a window derived from the
	// actual pixel data. With one, PS3.4 N.2 replaces that heuristic: an absent
	// VOI transformation is identity, so the image data must not be consulted.
	// Dataset-backed images resolve their window from Dataset attributes below,
	// so scanning the pixels here would only produce a discarded value.
	if !presentationStateApplied && bitsStored != 1 && img.dataset == nil {
		transforms.windowCenter, transforms.windowWidth = img.pixelData.Load().CalculateOptimalWindow()
	}

	var functional *dataset.Dataset
	if img.dataset != nil {
		functional = imageFunctionalGroupValues(img.dataset, frame)
	}
	// Enhanced multi-frame functional groups override top-level image
	// attributes for the selected frame. The top-level Dataset remains the
	// fallback for legacy images and attributes omitted from the macro.
	transforms.voiPrimary, transforms.voiFallback = functional, img.dataset
	switch {
	case presentationStateApplied:
		// PS3.4 N.2: transformations defined by the Presentation State replace
		// those of the referenced image, and absent ones are treated as
		// identity instead of falling back to the image attributes.
		modalityLUT, slope, intercept, descriptorSigned, err := imageModalityTransformWithLUTMode(
			img.presentationState, nil, pixelSigned, minInput, maxInput, img.lutVRMode,
		)
		if err != nil {
			return transforms, err
		}
		transforms.modalityLUT = modalityLUT
		transforms.rescaleSlope = slope
		transforms.rescaleIntercept = intercept
		transforms.voiDescriptorSigned = descriptorSigned
		transforms.voiPrimary, transforms.voiFallback = nil, nil
		if img.dataset != nil {
			softcopyVOI, err := imageSoftcopyVOI(img.presentationState, img.dataset, frame)
			if err != nil {
				return transforms, err
			}
			if softcopyVOI != nil {
				transforms.voiPrimary = softcopyVOI
			}
		}
		if err := transforms.applyExplicitWindow(nil, img.windowIndex); err != nil {
			return transforms, err
		}
	case img.dataset != nil:
		modalityLUT, slope, intercept, descriptorSigned, err := imageModalityTransformWithLUTMode(
			functional, img.dataset, pixelSigned, minInput, maxInput, img.lutVRMode,
		)
		if err != nil {
			return transforms, err
		}
		transforms.modalityLUT = modalityLUT
		transforms.rescaleSlope = slope
		transforms.rescaleIntercept = intercept
		transforms.voiDescriptorSigned = descriptorSigned
		softcopyVOI, err := imageSoftcopyVOI(img.presentationState, img.dataset, frame)
		if err != nil {
			return transforms, err
		}
		if softcopyVOI != nil {
			transforms.voiPrimary, transforms.voiFallback = softcopyVOI, nil
		}
		if err := transforms.applyExplicitWindow(img.dataset, img.windowIndex); err != nil {
			return transforms, err
		}
	}

	if img.dataset == nil {
		return transforms, nil
	}
	if !transforms.hasExplicitWindow && !presentationStateApplied && img.pixelData.Load().Info.BitsStored != 1 {
		img.deriveWindowFromPixels(&transforms)
	}
	if presentationStateApplied && !transforms.hasExplicitWindow {
		transforms.applyIdentityWindow(minInput, maxInput)
	}
	return transforms, nil
}

// applyExplicitWindow adopts an explicitly encoded Window Center/Width pair.
func (t *grayscaleTransforms) applyExplicitWindow(fallback *dataset.Dataset, windowIndex int) error {
	if pixeldata.DatasetContaining(t.voiPrimary, fallback, tag.WindowCenter, tag.WindowWidth) == nil {
		return nil
	}
	center, width, err := imageWindowPairAt(t.voiPrimary, fallback, windowIndex)
	if err != nil {
		return err
	}
	t.windowCenter, t.windowWidth = center, width
	t.hasExplicitWindow = true
	return nil
}

// deriveWindowFromPixels sets the window from the actual pixel value range.
func (img *DicomImage) deriveWindowFromPixels(t *grayscaleTransforms) {
	minimum, maximum, err := pixeldata.DatasetValueRange(img.dataset)
	if err != nil {
		minimum, maximum, err = img.windowRangeFromPixels()
	}
	if err != nil {
		// Neither the Dataset attributes nor MinMax supplied a range; fall back
		// to the sampled pixel range used for images without a Dataset.
		if img.pixelData.Load().Info.BitsStored != 1 {
			t.windowCenter, t.windowWidth = img.pixelData.Load().CalculateOptimalWindow()
		}
		return
	}
	if t.modalityLUT != nil {
		minimum = t.modalityLUT.Transform(minimum)
		maximum = t.modalityLUT.Transform(maximum)
	} else {
		minimum = minimum*t.rescaleSlope + t.rescaleIntercept
		maximum = maximum*t.rescaleSlope + t.rescaleIntercept
	}
	t.windowCenter = (minimum + maximum) / 2
	t.windowWidth = math.Max(1, math.Abs(maximum-minimum))
}

// windowRangeFromPixels returns the cached MinMax range of the pixel data.
// The range spans every frame and depends only on the pixel data, so scanning
// once serves every frame pipeline instead of rescanning all frames per frame.
// The caller must hold img.mu.
func (img *DicomImage) windowRangeFromPixels() (float64, float64, error) {
	if !img.windowRangeComputed {
		img.windowRangeScanCount++
		pd := img.pixelData.Load()
		if pd == nil || pd.Info == nil {
			img.windowRange[0], img.windowRange[1], img.windowRangeErr = 0, 0, fmt.Errorf("image has no pixel data")
		} else {
			img.windowRange[0], img.windowRange[1], img.windowRangeErr = pd.MinMax(true)
		}
		img.windowRangeComputed = true
	}
	return img.windowRange[0], img.windowRange[1], img.windowRangeErr
}

// applyIdentityWindow applies the identity VOI transformation that PS3.3
// C.11.2.1.2.1 Note 4 defines over the possible input range. It also provides
// the implicit scaling to the 0..255 intermediate requested by Note 9 and by
// the Presentation LUT.
func (t *grayscaleTransforms) applyIdentityWindow(minInput, maxInput float64) {
	minimum, maximum := minInput, maxInput
	if t.modalityLUT != nil {
		minimum = t.modalityLUT.MinimumOutputValue()
		maximum = t.modalityLUT.MaximumOutputValue()
	} else {
		minimum = minimum*t.rescaleSlope + t.rescaleIntercept
		maximum = maximum*t.rescaleSlope + t.rescaleIntercept
	}
	if minimum > maximum {
		minimum, maximum = maximum, minimum
	}
	t.windowCenter = (minimum + maximum + 1) / 2
	t.windowWidth = math.Max(1, maximum-minimum+1)
}

// configurePipelineLUTs attaches the modality, VOI and presentation LUTs.
func (img *DicomImage) configurePipelineLUTs(
	pipeline *render.GrayscalePipeline, t grayscaleTransforms, presentationStateApplied bool,
) error {
	if t.modalityLUT != nil {
		pipeline.SetModalityLUT(t.modalityLUT)
	}
	if pixeldata.DatasetContaining(t.voiPrimary, t.voiFallback, tag.VOILUTSequence) != nil {
		var voiLUT lut.LUT
		var err error
		if presentationStateApplied {
			voiLUT, err = pixeldata.SoftcopyVOILUTFromWithVRMode(t.voiPrimary, t.voiFallback, t.voiDescriptorSigned, img.voiLUTIndex, img.lutVRMode)
		} else {
			voiLUT, err = pixeldata.VOILUTFromWithVRMode(t.voiPrimary, t.voiFallback, t.voiDescriptorSigned, img.voiLUTIndex, img.lutVRMode)
		}
		if err != nil {
			return err
		}
		pipeline.SetVOILUT(voiLUT)
	}
	if function, ok := imageStringFrom(t.voiPrimary, t.voiFallback, tag.VOILUTFunction); ok {
		pipeline.SetVOILUTFunction(lut.VOILUTFunction(function))
	}
	// PS3.4 N.2: when a Presentation State is applied, its Presentation LUT
	// (including the absence of one, which means identity) replaces any
	// Presentation LUT carried by the referenced image.
	var presentationSource *dataset.Dataset
	if presentationStateApplied {
		presentationSource = pixeldata.DatasetContaining(
			img.presentationState, nil, tag.PresentationLUTSequence, tag.PresentationLUTShape,
		)
	} else {
		presentationSource = pixeldata.DatasetContaining(
			nil, img.dataset, tag.PresentationLUTSequence, tag.PresentationLUTShape,
		)
	}
	presentationLUT, present, err := pixeldata.PresentationLUTWithVRMode(presentationSource, img.lutVRMode)
	if err != nil {
		return err
	}
	if present {
		pipeline.SetPresentationLUT(presentationLUT)
	}
	return nil
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
	width = clampRequestedWindowWidth(width)
	from, to := img.lutFrameRange()
	for frame := from; frame <= to; frame++ {
		if gp, ok := img.GetOrCreatePipeline(frame).(*render.GrayscalePipeline); ok {
			gp.SetWindowWidth(width)
		}
	}
}

// clampRequestedWindowWidth enforces the PS3.3 C.11.2.1.2 LINEAR rule that
// Window Width is at least 1 on caller-supplied values; a smaller width inverts
// the display mapping. The dataset path is validated separately.
func clampRequestedWindowWidth(width float64) float64 {
	if math.IsNaN(width) || width < 1 {
		return 1
	}
	return width
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
	width = clampRequestedWindowWidth(width)
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
	return img.RenderFrameContext(context.Background(), writer, frame, options)
}

// RenderFrameContext renders a frame and associates logging with ctx.
func (img *DicomImage) RenderFrameContext(ctx context.Context, writer io.Writer, frame int, options *render.ExportOptions) error {
	rendered, err := img.RenderFrameImageContext(ctx, frame)
	if err != nil {
		return err
	}
	return render.NewImageExporter(img.GetOrCreatePipeline(frame)).ExportImage(writer, rendered, options)
}

// RenderFrameImage renders the specified frame and returns the unencoded Go image.
func (img *DicomImage) RenderFrameImage(frame int) (image.Image, error) {
	return img.RenderFrameImageContext(context.Background(), frame)
}

// RenderFrameImageContext renders a frame and associates logging with ctx.
func (img *DicomImage) RenderFrameImageContext(ctx context.Context, frame int) (image.Image, error) {
	return img.renderFrameImage(ctx, frame, true)
}

// RenderFrameImageWithOptions renders a frame with an optional spatial
// transform, viewport, and final-coordinate graphics. When SpatialTransform is
// supplied it replaces legacy SetScale; otherwise legacy scale remains active.
func (img *DicomImage) RenderFrameImageWithOptions(frame int, options FrameRenderOptions) (image.Image, error) {
	return img.RenderFrameImageWithOptionsContext(context.Background(), frame, options)
}

// RenderFrameImageWithOptionsContext renders a frame with spatial options and
// associates logging with ctx.
func (img *DicomImage) RenderFrameImageWithOptionsContext(ctx context.Context, frame int, options FrameRenderOptions) (image.Image, error) {
	if options.SpatialTransform == nil && options.Viewport.Empty() && len(options.Graphics) == 0 && options.Background == nil && options.Interpolation == interpolation.ModeNearestNeighbor {
		return img.RenderFrameImageContext(ctx, frame)
	}
	rendered, err := img.renderFrameImage(ctx, frame, options.SpatialTransform == nil)
	if err != nil {
		return nil, err
	}
	if options.SpatialTransform != nil || !options.Viewport.Empty() {
		matrix := transform.Identity()
		if options.SpatialTransform != nil {
			matrix, err = options.SpatialTransform.Affine(transformRect(rendered.Bounds()))
			if err != nil {
				return nil, err
			}
		}
		rendered, err = render.ApplyAffine(rendered, matrix, options.Viewport, options.Interpolation, options.Background)
		if err != nil {
			return nil, err
		}
	}
	if len(options.Graphics) != 0 {
		return render.DrawGraphics(rendered, options.Graphics)
	}
	return rendered, nil
}

func transformRect(bounds image.Rectangle) math3d.Rect {
	return math3d.Rect{Width: float64(bounds.Dx()), Height: float64(bounds.Dy())}
}

func (img *DicomImage) renderFrameImage(ctx context.Context, frame int, applyLegacyScale bool) (rendered image.Image, err error) {
	// A DicomImage built without pixel data has nothing to render; report that
	// instead of dereferencing a nil snapshot below.
	if pd := img.pixelData.Load(); pd == nil || pd.Info == nil {
		return nil, fmt.Errorf("image has no pixel data")
	}
	// Snapshot the pixel data once: a concurrent DecodeIfNeeded may replace it
	// between loads, and mixing frame bytes with post-decode metadata produces
	// misleading size errors or corrupt pixels.
	var pd *pixeldata.Data
	logDebug := logging.Enabled(ctx, slog.LevelDebug)
	logError := logging.Enabled(ctx, slog.LevelError)
	if logDebug || logError {
		started := time.Now()
		defer func() {
			level := slog.LevelDebug
			event := "render_completed"
			message := "DICOM render completed"
			if err != nil {
				level = slog.LevelError
				event = "render_failed"
				message = "DICOM render failed"
			}
			if !logging.Enabled(ctx, level) {
				return
			}
			attrs := []slog.Attr{
				slog.Int("frame", frame),
				slog.Int("width", int(pd.Info.Width)),
				slog.Int("height", int(pd.Info.Height)),
				slog.Int("samples_per_pixel", int(pd.Info.SamplesPerPixel)),
				slog.Int("bits_allocated", int(pd.Info.BitsAllocated)),
				slog.Duration("duration", time.Since(started)),
			}
			if err != nil {
				attrs = append(attrs,
					slog.String("failure_stage", "render"),
					slog.String("error_type", fmt.Sprintf("%T", err)),
				)
			}
			logging.Emit(ctx, logging.Record{
				Level: level, Component: "imaging.render", Event: event, Message: message, Attrs: attrs,
			})
		}()
	}

	var frameData []byte
	var pipeline render.Pipeline
	// Retry once if a decode swapped the pixel data while the pipeline was
	// being created, so frame bytes and rendering metadata come from the same
	// snapshot.
	for attempt := 0; ; attempt++ {
		pd = img.pixelData.Load()
		if frame < 0 || frame >= pd.Info.NumberOfFrames {
			return nil, fmt.Errorf("frame index out of range: %d", frame)
		}
		// Rendering only reads the frame, so borrow the snapshot's buffer
		// instead of copying it; pd stays alive for the whole render.
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		frameData, err = pd.FrameView(frame)
		if err != nil {
			return nil, fmt.Errorf("failed to get frame data: %w", err)
		}
		pipeline, err = img.GetOrCreatePipelineWithError(frame)
		if err != nil {
			return nil, fmt.Errorf("create rendering pipeline: %w", err)
		}
		if img.pixelData.Load() == pd || attempt > 0 {
			break
		}
	}
	exporter := render.NewImageExporter(pipeline)
	width := int(pd.Info.Width)
	height := int(pd.Info.Height)
	switch pd.Info.SamplesPerPixel {
	case 1:
		photometric := pixel.Monochrome2.Value
		// PS3.4 N.2: the referenced image's MONOCHROME1/MONOCHROME2 value is
		// ignored when a Presentation State is applied, because the presentation
		// transformations own display polarity.
		if img.presentationState == nil && pd.Info.PhotometricInterpretation != nil {
			photometric = pd.Info.PhotometricInterpretation.Value
		}
		rendered, err = exporter.RenderGrayscaleImageWithBitDepth(
			frameData,
			width,
			height,
			int(pd.Info.BitsAllocated),
			int(pd.Info.BitsStored),
			int(pd.Info.HighBit),
			pd.Info.PixelRepresentation == pixel.SignedPixels,
			photometric,
		)
		if err == nil {
			rendered = maskGrayscalePadding(rendered, pd, frame)
		}
		if err == nil && img.supplementalPaletteEligible(pd) {
			rendered, err = img.applySupplementalPalette(rendered, frame, pd)
		}
	case 3:
		photometric := pixel.RGBPhotometric.Value
		if pd.Info.PhotometricInterpretation != nil {
			photometric = pd.Info.PhotometricInterpretation.Value
		}
		rendered, err = exporter.RenderRGBImageWithBitDepth(
			frameData,
			width,
			height,
			photometric,
			int(pd.Info.PlanarConfiguration),
			int(pd.Info.BitsAllocated),
			int(pd.Info.BitsStored),
			int(pd.Info.HighBit),
		)
	case 4:
		photometric := pixel.ARGBPhotometric.Value
		if pd.Info.PhotometricInterpretation != nil {
			photometric = pd.Info.PhotometricInterpretation.Value
		}
		// ARGB/RGBA with 4 samples are defined with 8-bit samples; rendering a
		// wider allocation as 8-bit would emit silently wrong pixels.
		if pd.Info.BitsAllocated != 8 {
			return nil, fmt.Errorf("unsupported BitsAllocated %d for %s: only 8-bit samples are supported",
				pd.Info.BitsAllocated, photometric)
		}
		rendered, err = exporter.RenderRGBAImageWithConfig(
			frameData,
			width,
			height,
			photometric,
			int(pd.Info.PlanarConfiguration),
		)
	default:
		return nil, fmt.Errorf("unsupported samples per pixel: %d", pd.Info.SamplesPerPixel)
	}
	if err != nil {
		return nil, err
	}
	rendered = img.applyGrayscaleColorMap(rendered, frame)
	rendered = img.applyOverlays(rendered, frame)
	if applyLegacyScale {
		rendered = img.scaleImage(rendered, pd)
	}
	return rendered, nil
}

// maskGrayscalePadding forces Pixel Padding samples to display black after the
// grayscale pipeline. PS3.3 C.7.5.1.1.2 says padding is not image data, and
// applying the mask after windowing keeps MONOCHROME1 inversion and the
// pipeline LUT from mapping padding to a visible gray.
func maskGrayscalePadding(rendered image.Image, pd *pixeldata.Data, frame int) image.Image {
	gray, ok := rendered.(*image.Gray)
	if !ok || pd.Info.PixelPaddingValue == nil || pd.Info.SamplesPerPixel != 1 {
		return rendered
	}
	mask, err := pd.PaddingMask(frame)
	if err != nil {
		return rendered
	}
	width, height := gray.Bounds().Dx(), gray.Bounds().Dy()
	if width <= 0 || height <= 0 {
		return rendered
	}
	for index, padded := range mask {
		if !padded {
			continue
		}
		x, y := index%width, index/width
		if x >= width || y >= height {
			break
		}
		gray.Pix[y*gray.Stride+x] = 0
	}
	return gray
}

func (img *DicomImage) supplementalPaletteEligible(pd *pixeldata.Data) bool {
	if img.dataset == nil || pd.Info.SamplesPerPixel != 1 {
		return false
	}
	photometric := pixel.Monochrome2.Value
	if pd.Info.PhotometricInterpretation != nil {
		photometric = pd.Info.PhotometricInterpretation.Value
	}
	if photometric != pixel.Monochrome1.Value && photometric != pixel.Monochrome2.Value {
		return false
	}
	presentation, ok := img.dataset.GetString(tag.PixelPresentation)
	if !ok || strings.ToUpper(strings.TrimSpace(presentation)) != "COLOR" {
		return false
	}
	return pixeldata.DatasetContaining(img.dataset, nil,
		tag.RedPaletteColorLookupTableDescriptor,
		tag.EnhancedPaletteColorLookupTableSequence,
		tag.PaletteColorLookupTableSequence,
	) != nil
}

func (img *DicomImage) applySupplementalPalette(grayscale image.Image, frame int, pd *pixeldata.Data) (image.Image, error) {
	first, entries, err := pixeldata.PaletteColors(img.dataset)
	if err != nil {
		return nil, fmt.Errorf("build Supplemental Palette Color LUT: %w", err)
	}
	bounds := grayscale.Bounds()
	width := int(pd.Info.Width)
	height := int(pd.Info.Height)
	result := image.NewNRGBA(bounds)

	// Read the rendered gray value directly for the pipeline's common outputs so
	// the per-pixel color model conversion is not repeated 4 bytes at a time.
	var grayAt func(x, y int) uint8
	switch source := grayscale.(type) {
	case *image.Gray:
		grayAt = func(x, y int) uint8 { return source.Pix[source.PixOffset(x, y)] }
	case *image.Gray16:
		// Gray16.RGBA reports the 16-bit value and color.GrayModel keeps its high
		// byte, which is the first byte of the big-endian pixel.
		grayAt = func(x, y int) uint8 { return source.Pix[source.PixOffset(x, y)] }
	default:
		grayAt = func(x, y int) uint8 {
			return color.GrayModel.Convert(grayscale.At(x, y)).(color.Gray).Y
		}
	}

	for y := 0; y < height; y++ {
		rowOffset := result.PixOffset(bounds.Min.X, bounds.Min.Y+y)
		for x := 0; x < width; x++ {
			gray := grayAt(bounds.Min.X+x, bounds.Min.Y+y)
			red, green, blue, alpha := gray, gray, gray, uint8(255)
			value, sampleErr := pd.Sample(frame, x, y, 0)
			if sampleErr != nil {
				return nil, sampleErr
			}
			if value >= int64(first) {
				paletteIndex := int(value - int64(first))
				if paletteIndex >= len(entries) {
					paletteIndex = len(entries) - 1
				}
				entry := entries[paletteIndex]
				red, green, blue, alpha = entry.R, entry.G, entry.B, entry.A
			}
			offset := rowOffset + x*4
			result.Pix[offset] = red
			result.Pix[offset+1] = green
			result.Pix[offset+2] = blue
			result.Pix[offset+3] = alpha
		}
	}
	return result, nil
}

func (img *DicomImage) scaleImage(source image.Image, pd *pixeldata.Data) image.Image {
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
		if pd.Info.BitsStored == 1 {
			golangdraw.NearestNeighbor.Scale(gray, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
		} else {
			golangdraw.BiLinear.Scale(gray, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
		}
		return gray
	}
	if pd.Info.BitsStored == 1 {
		// An overlay can composite a 1-bit image into RGBA; interpolation would
		// invent gray levels between the two stored values, so scale nearest.
		rgba := image.NewRGBA(destinationBounds)
		golangdraw.NearestNeighbor.Scale(rgba, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
		return rgba
	}
	rgba := image.NewRGBA(destinationBounds)
	golangdraw.BiLinear.Scale(rgba, destinationBounds, source, source.Bounds(), golangdraw.Src, nil)
	return rgba
}

// RenderCurrentFrame renders the current frame to a writer
func (img *DicomImage) RenderCurrentFrame(writer io.Writer, options *render.ExportOptions) error {
	return img.RenderCurrentFrameContext(context.Background(), writer, options)
}

// RenderCurrentFrameContext renders the current frame and associates logging with ctx.
func (img *DicomImage) RenderCurrentFrameContext(ctx context.Context, writer io.Writer, options *render.ExportOptions) error {
	img.mu.RLock()
	frame := img.currentFrame
	img.mu.RUnlock()
	return img.RenderFrameContext(ctx, writer, frame, options)
}

// DecodeIfNeeded decodes the pixel data if it's in a compressed format
func (img *DicomImage) DecodeIfNeeded(c codec.Codec, params codec.Parameters) error {
	return img.DecodeIfNeededContext(context.Background(), c, params)
}

// DecodeIfNeededContext decodes compressed pixel data and associates logging with ctx.
func (img *DicomImage) DecodeIfNeededContext(ctx context.Context, c codec.Codec, params codec.Parameters) (err error) {
	img.decodeMu.Lock()
	defer img.decodeMu.Unlock()

	// Check if already uncompressed
	if img.pixelData.Load().Info.TransferSyntaxUID == transferSyntaxImplicitVRLittleEndian ||
		img.pixelData.Load().Info.TransferSyntaxUID == transferSyntaxExplicitVRLittleEndian ||
		img.pixelData.Load().Info.TransferSyntaxUID == transferSyntaxExplicitVRBigEndian {
		return nil // Already uncompressed
	}
	logDebug := logging.Enabled(ctx, slog.LevelDebug)
	logError := logging.Enabled(ctx, slog.LevelError)
	if logDebug || logError {
		started := time.Now()
		inputTransferSyntax := img.pixelData.Load().Info.TransferSyntaxUID
		frameCount := img.pixelData.Load().Info.NumberOfFrames
		codecType := fmt.Sprintf("%T", c)
		defer func() {
			level := slog.LevelDebug
			event := "pixel_decode_completed"
			message := "DICOM pixel decode completed"
			if err != nil {
				level = slog.LevelError
				event = "pixel_decode_failed"
				message = "DICOM pixel decode failed"
			}
			if !logging.Enabled(ctx, level) {
				return
			}
			attrs := []slog.Attr{
				slog.String("input_transfer_syntax", inputTransferSyntax),
				slog.String("output_transfer_syntax", transferSyntaxExplicitVRLittleEndian),
				slog.String("codec_type", codecType),
				slog.Int("frame_count", frameCount),
				slog.Duration("duration", time.Since(started)),
			}
			if err != nil {
				attrs = append(attrs,
					slog.String("failure_stage", "pixel_decode"),
					slog.String("error_type", fmt.Sprintf("%T", err)),
				)
			}
			logging.Emit(ctx, logging.Record{
				Level: level, Component: "imaging.decode", Event: event, Message: message, Attrs: attrs,
			})
		}()
	}

	// Decode
	decoded, err := img.pixelData.Load().Decode(ctx, c, params)
	if err != nil {
		return fmt.Errorf("failed to decode pixel data: %w", err)
	}

	img.mu.Lock()
	img.pixelData.Store(decoded)
	// Clear cached pipelines as pixel data has changed. The derived window
	// range comes from the same samples and must be rescanned as well.
	img.pipelines = make(map[int]render.Pipeline)
	img.windowRangeComputed = false
	img.windowRangeErr = nil
	img.mu.Unlock()

	return nil
}

// Clone copies image data and configuration. Pipelines implementing
// render.PipelineCloner are copied independently; other custom pipelines are
// retained by reference because the Pipeline interface has no general clone operation.
func (img *DicomImage) Clone() *DicomImage {
	img.mu.RLock()

	clonedPixelData := img.pixelData.Load().Clone()

	// Create new image
	cloned := &DicomImage{
		currentFrame:            img.currentFrame,
		pipelines:               make(map[int]render.Pipeline),
		scale:                   img.scale,
		showOverlays:            img.showOverlays,
		overlays:                append([]*DicomOverlayData(nil), img.overlays...),
		overlayColor:            img.overlayColor,
		autoApplyLUTToAllFrames: img.autoApplyLUTToAllFrames,
		grayscaleColorMaps:      make(map[int][256]colorconv.Color32, len(img.grayscaleColorMaps)),
		windowIndex:             img.windowIndex,
		voiLUTIndex:             img.voiLUTIndex,
		lutVRMode:               img.lutVRMode,
	}
	cloned.pixelData.Store(clonedPixelData)
	if img.dataset != nil {
		cloned.dataset = img.dataset.Clone()
	}
	if img.presentationState != nil {
		cloned.presentationState = img.presentationState.Clone()
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
