// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const (
	maxStackSlices  = 65535
	maxCutDimension = 65535
	maxCutPixels    = 4096 * 4096
)

// Stack is a lazy series of parallel patient-space cuts.
type Stack struct {
	Type          StackType
	PixelSpacing  float64
	SliceDistance float64

	volume *VolumeData
	specs  []CutSpec
}

// NewStack builds lazy standard-plane specifications covering volume bounds.
func NewStack(volume *VolumeData, stackType StackType, spacing, sliceDistance float64) (*Stack, error) {
	if volume == nil || len(volume.images) < 2 {
		return nil, fmt.Errorf("stack volume is nil or incomplete")
	}
	if stackType != StackTypeAxial && stackType != StackTypeCoronal && stackType != StackTypeSagittal {
		return nil, fmt.Errorf("unknown stack type %d", stackType)
	}
	if !positiveFiniteStack(spacing) || !positiveFiniteStack(sliceDistance) {
		return nil, fmt.Errorf("stack spacing and slice distance must be positive and finite")
	}
	bounds := volume.bounds
	extentX := math.Max(0, bounds.Max.X-bounds.Min.X)
	extentY := math.Max(0, bounds.Max.Y-bounds.Min.Y)
	extentZ := math.Max(0, bounds.Max.Z-bounds.Min.Z)

	var count, rows, columns int
	var topLeft func(float64) math3d.Point3
	var rowDirection, columnDirection math3d.Vector3
	var minimum, extent float64
	switch stackType {
	case StackTypeAxial:
		minimum, extent = bounds.Min.Z, extentZ
		count, rows, columns = boundedInclusiveCount(extentZ, sliceDistance),
			boundedInclusiveCount(extentY, spacing), boundedInclusiveCount(extentX, spacing)
		rowDirection, columnDirection = math3d.Vector3{X: 1}, math3d.Vector3{Y: 1}
		topLeft = func(position float64) math3d.Point3 {
			return math3d.Point3{X: bounds.Min.X, Y: bounds.Min.Y, Z: position}
		}
	case StackTypeCoronal:
		minimum, extent = bounds.Min.Y, extentY
		count, rows, columns = boundedInclusiveCount(extentY, sliceDistance),
			boundedInclusiveCount(extentZ, spacing), boundedInclusiveCount(extentX, spacing)
		rowDirection, columnDirection = math3d.Vector3{X: 1}, math3d.Vector3{Z: 1}
		topLeft = func(position float64) math3d.Point3 {
			return math3d.Point3{X: bounds.Min.X, Y: position, Z: bounds.Min.Z}
		}
	case StackTypeSagittal:
		minimum, extent = bounds.Min.X, extentX
		count, rows, columns = boundedInclusiveCount(extentX, sliceDistance),
			boundedInclusiveCount(extentZ, spacing), boundedInclusiveCount(extentY, spacing)
		rowDirection, columnDirection = math3d.Vector3{Y: 1}, math3d.Vector3{Z: 1}
		topLeft = func(position float64) math3d.Point3 {
			return math3d.Point3{X: position, Y: bounds.Min.Y, Z: bounds.Min.Z}
		}
	}
	if count == 0 {
		return nil, fmt.Errorf("stack slice count exceeds limit %d", maxStackSlices)
	}
	if rows == 0 || columns == 0 {
		return nil, fmt.Errorf("stack output dimensions exceed DICOM limit %d", maxCutDimension)
	}
	if int64(rows)*int64(columns) > maxCutPixels {
		return nil, fmt.Errorf("stack output dimensions %dx%d exceed cut pixel limit %d", rows, columns, maxCutPixels)
	}

	specs := make([]CutSpec, count)
	for index := range count {
		position := minimum + math.Min(float64(index)*sliceDistance, extent)
		specs[index] = CutSpec{
			TopLeft:             topLeft(position),
			RowDirection:        rowDirection,
			ColumnDirection:     columnDirection,
			Rows:                rows,
			Columns:             columns,
			PixelSpacingRows:    spacing,
			PixelSpacingColumns: spacing,
		}
	}
	return &Stack{Type: stackType, PixelSpacing: spacing, SliceDistance: sliceDistance, volume: volume, specs: specs}, nil
}

// Len returns the number of lazy slice specifications.
func (stack *Stack) Len() int {
	if stack == nil {
		return 0
	}
	return len(stack.specs)
}

// Spec returns a copy of one slice specification.
func (stack *Stack) Spec(index int) (CutSpec, error) {
	if stack == nil || index < 0 || index >= len(stack.specs) {
		return CutSpec{}, fmt.Errorf("stack index %d out of range [0, %d)", index, stack.Len())
	}
	return stack.specs[index], nil
}

// Materialize calculates one stack slice on demand.
func (stack *Stack) Materialize(ctx context.Context, index int, options CutOptions) (*Slice, error) {
	spec, err := stack.Spec(index)
	if err != nil {
		return nil, err
	}
	return stack.volume.Cut(ctx, spec, options)
}

// MinMax returns the range of valid slice values.
func (slice *Slice) MinMax() (float64, float64, error) {
	if slice == nil || len(slice.Values) == 0 || len(slice.Valid) != len(slice.Values) {
		return 0, 0, fmt.Errorf("slice values and validity mask are incomplete")
	}
	found := false
	minimum, maximum := 0.0, 0.0
	for index, value := range slice.Values {
		if !slice.Valid[index] {
			continue
		}
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return 0, 0, fmt.Errorf("slice sample %d is not finite", index)
		}
		if !found {
			minimum, maximum, found = value, value, true
			continue
		}
		minimum = math.Min(minimum, value)
		maximum = math.Max(maximum, value)
	}
	if !found {
		return 0, 0, fmt.Errorf("slice has no valid samples")
	}
	return minimum, maximum, nil
}

// Render8Bit applies the DICOM linear window function to valid samples.
func (slice *Slice) Render8Bit(center, width float64, invalid byte) ([]byte, error) {
	if slice == nil || len(slice.Valid) != len(slice.Values) {
		return nil, fmt.Errorf("slice values and validity mask are incomplete")
	}
	if width < 1 || math.IsNaN(center) || math.IsNaN(width) || math.IsInf(center, 0) || math.IsInf(width, 0) {
		return nil, fmt.Errorf("window center and width must be finite and width at least one")
	}
	output := make([]byte, len(slice.Values))
	for index, value := range slice.Values {
		if !slice.Valid[index] {
			output[index] = invalid
			continue
		}
		output[index] = window8(value, center, width)
	}
	return output, nil
}

func window8(value, center, width float64) byte {
	if width == 1 {
		if value > center-0.5 {
			return 255
		}
		return 0
	}
	lower := center - 0.5 - (width-1)/2
	upper := center - 0.5 + (width-1)/2
	if value <= lower {
		return 0
	}
	if value >= upper {
		return 255
	}
	mapped := ((value-(center-0.5))/(width-1) + 0.5) * 255
	return byte(math.Round(mapped))
}

func boundedInclusiveCount(extent, spacing float64) int {
	if extent <= 0 {
		return 1
	}
	intervals := math.Ceil(extent / spacing)
	if math.IsInf(intervals, 0) || math.IsNaN(intervals) || intervals >= float64(maxCutDimension) {
		return 0
	}
	return int(intervals) + 1
}

func positiveFiniteStack(value float64) bool {
	return value > 0 && !math.IsNaN(value) && !math.IsInf(value, 0)
}
