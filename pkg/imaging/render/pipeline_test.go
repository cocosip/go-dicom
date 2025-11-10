// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package render

import (
	"testing"
)

func TestGrayscalePipeline(t *testing.T) {
	// Create a pipeline with rescale slope=2, intercept=10, window=128/256
	pipeline := NewGrayscalePipeline(2.0, 10.0, 128.0, 256.0, 0, 100, false)

	if pipeline.WindowCenter() != 128.0 {
		t.Errorf("Expected window center 128, got %f", pipeline.WindowCenter())
	}

	if pipeline.WindowWidth() != 256.0 {
		t.Errorf("Expected window width 256, got %f", pipeline.WindowWidth())
	}

	if pipeline.Invert() {
		t.Error("Expected invert to be false")
	}

	// Test LUT
	lut := pipeline.LUT()
	if lut == nil {
		t.Fatal("Expected non-nil LUT")
	}

	// Test transformation
	// Input=0 -> Modality: 0*2+10=10 -> VOI: window transform
	result := lut.Transform(0)
	if result < 0 || result > 255 {
		t.Errorf("Transform result should be in [0,255], got %f", result)
	}
}

func TestGrayscalePipelineSetWindow(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	// Change window
	pipeline.SetWindow(100.0, 200.0)

	if pipeline.WindowCenter() != 100.0 {
		t.Errorf("Expected window center 100, got %f", pipeline.WindowCenter())
	}

	if pipeline.WindowWidth() != 200.0 {
		t.Errorf("Expected window width 200, got %f", pipeline.WindowWidth())
	}

	// LUT should be rebuilt
	lut := pipeline.LUT()
	if lut == nil {
		t.Fatal("Expected non-nil LUT after window change")
	}
}

func TestGrayscalePipelineSetWindowWidth(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	pipeline.SetWindowWidth(512.0)

	if pipeline.WindowWidth() != 512.0 {
		t.Errorf("Expected window width 512, got %f", pipeline.WindowWidth())
	}
}

func TestGrayscalePipelineSetWindowCenter(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	pipeline.SetWindowCenter(200.0)

	if pipeline.WindowCenter() != 200.0 {
		t.Errorf("Expected window center 200, got %f", pipeline.WindowCenter())
	}
}

func TestGrayscalePipelineWithInvert(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, true)

	if !pipeline.Invert() {
		t.Error("Expected invert to be true")
	}

	lut := pipeline.LUT()
	if lut == nil {
		t.Fatal("Expected non-nil LUT")
	}

	// Test that inversion is applied
	// Low input should give high output after inversion
	lowResult := lut.Transform(0)
	highResult := lut.Transform(255)

	if lowResult <= highResult {
		t.Errorf("Expected inversion: low input should give higher output than high input, got low=%f, high=%f",
			lowResult, highResult)
	}
}

func TestGrayscalePipelineSetInvert(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	pipeline.SetInvert(true)

	if !pipeline.Invert() {
		t.Error("Expected invert to be true")
	}

	// LUT should be rebuilt
	lut := pipeline.LUT()
	if lut == nil {
		t.Fatal("Expected non-nil LUT after invert change")
	}
}

func TestGrayscalePipelineClearCache(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	// Get LUT to build it
	lut1 := pipeline.LUT()
	if lut1 == nil {
		t.Fatal("Expected non-nil LUT")
	}

	// Clear cache
	pipeline.ClearCache()

	// LUT should be rebuilt on next access
	lut2 := pipeline.LUT()
	if lut2 == nil {
		t.Fatal("Expected non-nil LUT after clear cache")
	}
}

func TestGrayscalePipelineConcurrency(t *testing.T) {
	pipeline := NewGrayscalePipeline(1.0, 0, 128.0, 256.0, 0, 255, false)

	// Test concurrent access
	done := make(chan bool)

	// Reader goroutines
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				lut := pipeline.LUT()
				_ = lut.Transform(float64(j))
			}
			done <- true
		}()
	}

	// Writer goroutine
	go func() {
		for j := 0; j < 50; j++ {
			pipeline.SetWindow(float64(100+j), float64(200+j))
		}
		done <- true
	}()

	// Wait for all goroutines
	for i := 0; i < 6; i++ {
		<-done
	}
}

func TestCompositeLUTEmpty(t *testing.T) {
	lut := NewCompositeLUT()

	if lut == nil {
		t.Fatal("Expected non-nil composite LUT")
	}

	if lut.MinimumOutputValue() != 0 {
		t.Errorf("Expected min output 0, got %f", lut.MinimumOutputValue())
	}

	if lut.MaximumOutputValue() != 255 {
		t.Errorf("Expected max output 255, got %f", lut.MaximumOutputValue())
	}

	// Empty composite should just pass through
	result := lut.Transform(100)
	if result != 100 {
		t.Errorf("Empty composite should pass through, expected 100, got %f", result)
	}
}

func TestCompositeLUTIsValid(t *testing.T) {
	// Create composite with modality LUT (always valid) and invert LUT (always valid)
	lut1 := NewModalityRescaleLUT(1.0, 0, 0, 100)
	lut2 := NewInvertLUT(0, 255)
	composite := NewCompositeLUT(lut1, lut2)

	if !composite.IsValid() {
		t.Error("Expected composite to be valid when all components are valid")
	}

	// Create composite with VOI LUT (always returns false for IsValid to force recalculation)
	lut3 := NewLinearVOILUT(128, 256)
	composite2 := NewCompositeLUT(lut1, lut3)

	if composite2.IsValid() {
		t.Error("Expected composite to be invalid when VOI LUT is included (VOI LUT always returns false)")
	}
}

func TestCompositeLUTRecalculate(t *testing.T) {
	lut1 := NewModalityRescaleLUT(1.0, 0, 0, 100)
	lut2 := NewLinearVOILUT(128, 256)
	composite := NewCompositeLUT(lut1, lut2)

	// Recalculate should not panic
	composite.Recalculate()
}
