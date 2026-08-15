// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package interpolation provides scalar nearest-neighbor and bilinear image
// sampling shared by rendering and volume reconstruction.
package interpolation

import (
	"fmt"
	"math"
)

// Mode selects an interpolation algorithm.
type Mode int

const (
	// ModeNearestNeighbor samples the nearest pixel center.
	ModeNearestNeighbor Mode = iota
	// ModeBilinear blends the four neighboring pixel centers.
	ModeBilinear
)

// Grid is a read-only view over row-major scalar image data. Stride is counted
// in float64 values rather than bytes.
type Grid struct {
	data   []float64
	width  int
	height int
	stride int
}

// NewGrid creates a scalar grid view without copying data.
func NewGrid(data []float64, width, height, stride int) (*Grid, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("grid dimensions must be positive")
	}
	if stride < width {
		return nil, fmt.Errorf("grid stride %d is below width %d", stride, width)
	}
	required := (height-1)*stride + width
	if len(data) < required {
		return nil, fmt.Errorf("grid data has %d values, want at least %d", len(data), required)
	}
	return &Grid{data: data, width: width, height: height, stride: stride}, nil
}

// Nearest samples the nearest pixel center. Half coordinates round toward the
// next positive pixel.
func (g *Grid) Nearest(x, y float64) (float64, bool) {
	if !g.contains(x, y) {
		return 0, false
	}
	column := min(int(math.Floor(x+0.5)), g.width-1)
	row := min(int(math.Floor(y+0.5)), g.height-1)
	return g.data[row*g.stride+column], true
}

// Bilinear samples the four neighboring pixel centers.
func (g *Grid) Bilinear(x, y float64) (float64, bool) {
	if !g.contains(x, y) {
		return 0, false
	}
	x0, y0 := int(math.Floor(x)), int(math.Floor(y))
	x1, y1 := min(x0+1, g.width-1), min(y0+1, g.height-1)
	dx, dy := x-float64(x0), y-float64(y0)
	top := g.data[y0*g.stride+x0]*(1-dx) + g.data[y0*g.stride+x1]*dx
	bottom := g.data[y1*g.stride+x0]*(1-dx) + g.data[y1*g.stride+x1]*dx
	return top*(1-dy) + bottom*dy, true
}

// Resize returns tightly packed output with source and destination endpoints
// aligned in both dimensions.
func (g *Grid) Resize(width, height int, mode Mode) ([]float64, error) {
	if g == nil {
		return nil, fmt.Errorf("grid is nil")
	}
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("output dimensions must be positive")
	}
	if mode != ModeNearestNeighbor && mode != ModeBilinear {
		return nil, fmt.Errorf("unknown interpolation mode %d", mode)
	}
	output := make([]float64, width*height)
	for row := 0; row < height; row++ {
		y := resizeCoordinate(row, height, g.height)
		for column := 0; column < width; column++ {
			x := resizeCoordinate(column, width, g.width)
			var value float64
			if mode == ModeNearestNeighbor {
				value, _ = g.Nearest(x, y)
			} else {
				value, _ = g.Bilinear(x, y)
			}
			output[row*width+column] = value
		}
	}
	return output, nil
}

func (g *Grid) contains(x, y float64) bool {
	return g != nil && !math.IsNaN(x) && !math.IsNaN(y) &&
		!math.IsInf(x, 0) && !math.IsInf(y, 0) &&
		x >= 0 && y >= 0 && x <= float64(g.width-1) && y <= float64(g.height-1)
}

func resizeCoordinate(index, outputSize, inputSize int) float64 {
	if outputSize == 1 || inputSize == 1 {
		return 0
	}
	return float64(index) * float64(inputSize-1) / float64(outputSize-1)
}
