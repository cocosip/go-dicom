// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"sync"

	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/types"
)

// Pipeline defines the interface for rendering pipelines
type Pipeline interface {
	// LUT returns the composite LUT of the pipeline
	LUT() LUT
	// ClearCache removes all cached data and only keeps configuration data
	ClearCache()
}

// GrayscalePipeline implements a grayscale rendering pipeline
// The pipeline consists of: Modality LUT -> VOI LUT -> Output LUT -> (optional) Invert LUT
type GrayscalePipeline struct {
	mu sync.RWMutex

	modalityLUT      ModalityLUT
	voiLUT           lut.VOILUT
	invertLUT        LUT
	compositeLUT     LUT
	windowWidth      float64
	windowCenter     float64
	invert           bool
	minInputValue    float64
	maxInputValue    float64
	rescaleSlope     float64
	rescaleIntercept float64
}

// NewGrayscalePipeline creates a new grayscale rendering pipeline
func NewGrayscalePipeline(rescaleSlope, rescaleIntercept, windowCenter, windowWidth, minInput, maxInput float64, invert bool) *GrayscalePipeline {
	p := &GrayscalePipeline{
		windowWidth:      windowWidth,
		windowCenter:     windowCenter,
		invert:           invert,
		minInputValue:    minInput,
		maxInputValue:    maxInput,
		rescaleSlope:     rescaleSlope,
		rescaleIntercept: rescaleIntercept,
	}
	p.buildPipeline()
	return p
}

// WindowWidth returns the current window width
func (p *GrayscalePipeline) WindowWidth() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.windowWidth
}

// SetWindowWidth updates the window width and rebuilds the pipeline
func (p *GrayscalePipeline) SetWindowWidth(width float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.windowWidth != width {
		p.windowWidth = width
		p.compositeLUT = nil
	}
}

// WindowCenter returns the current window center
func (p *GrayscalePipeline) WindowCenter() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.windowCenter
}

// SetWindowCenter updates the window center and rebuilds the pipeline
func (p *GrayscalePipeline) SetWindowCenter(center float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.windowCenter != center {
		p.windowCenter = center
		p.compositeLUT = nil
	}
}

// SetWindow updates both window center and width
func (p *GrayscalePipeline) SetWindow(center, width float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.windowCenter != center || p.windowWidth != width {
		p.windowCenter = center
		p.windowWidth = width
		p.compositeLUT = nil
	}
}

// Invert returns whether the output is inverted
func (p *GrayscalePipeline) Invert() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.invert
}

// SetInvert sets whether to invert the output
func (p *GrayscalePipeline) SetInvert(invert bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.invert != invert {
		p.invert = invert
		p.compositeLUT = nil
	}
}

// LUT returns the composite LUT for the pipeline
func (p *GrayscalePipeline) LUT() LUT {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.compositeLUT == nil {
		p.buildPipeline()
	}

	return p.compositeLUT
}

// ClearCache clears the cached LUT
func (p *GrayscalePipeline) ClearCache() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.compositeLUT = nil
}

func (p *GrayscalePipeline) buildPipeline() {
	// Build the composite LUT chain
	luts := make([]LUT, 0, 4)

	// 1. Modality LUT (rescale)
	if p.rescaleSlope != 1.0 || p.rescaleIntercept != 0.0 {
		p.modalityLUT = NewModalityRescaleLUT(p.rescaleSlope, p.rescaleIntercept, p.minInputValue, p.maxInputValue)
		luts = append(luts, p.modalityLUT)
	}

	// 2. VOI LUT (windowing)
	p.voiLUT = lut.CreateVOILUT(types.VOILUTFunctionLinear, p.windowCenter, p.windowWidth)
	luts = append(luts, p.voiLUT)

	// 3. Output LUT (8-bit normalization) is implicitly handled by VOI LUT

	// 4. Invert LUT (if needed)
	if p.invert {
		p.invertLUT = lut.NewInvertLUT(0, 255)
		luts = append(luts, p.invertLUT)
	}

	// Create composite LUT
	compositeLUT := lut.NewCompositeLUT()
	for _, lutItem := range luts {
		compositeLUT.Add(lutItem)
	}
	p.compositeLUT = compositeLUT
}
