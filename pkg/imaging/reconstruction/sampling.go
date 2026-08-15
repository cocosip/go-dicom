// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

// CutSpec describes one output plane in DICOM patient coordinates.
type CutSpec struct {
	TopLeft             math3d.Point3
	RowDirection        math3d.Vector3
	ColumnDirection     math3d.Vector3
	Rows                int
	Columns             int
	PixelSpacingRows    float64
	PixelSpacingColumns float64
}

// CutOptions controls deterministic cut materialization.
type CutOptions struct {
	// Workers defaults to one. Values above the pixel count are capped.
	Workers int
}

// Slice is a materialized patient-space cut. Valid distinguishes samples that
// are outside the volume or depend on padding from genuine numeric zero.
type Slice struct {
	Spec   CutSpec
	Values []float64
	Valid  []bool
}

// Sample performs bilinear interpolation in source frames and linear
// interpolation between the two adjacent, actual slice positions.
func (volume *VolumeData) Sample(point math3d.Point3) (float64, bool, error) {
	if volume == nil || len(volume.images) < 2 || len(volume.slicePositions) != len(volume.images) {
		return 0, false, fmt.Errorf("volume is not initialized")
	}
	position := dot3(point, volume.normal)
	first := volume.slicePositions[0]
	last := volume.slicePositions[len(volume.slicePositions)-1]
	if position < first-volumeGeometryTolerance || position > last+volumeGeometryTolerance {
		return 0, false, nil
	}
	if math.Abs(position-first) <= volumeGeometryTolerance {
		return bilinearImageSample(volume.images[0], point)
	}
	if math.Abs(position-last) <= volumeGeometryTolerance {
		return bilinearImageSample(volume.images[len(volume.images)-1], point)
	}

	upper := sort.Search(len(volume.slicePositions), func(index int) bool {
		return volume.slicePositions[index] >= position
	})
	if upper <= 0 || upper >= len(volume.images) {
		return 0, false, nil
	}
	lower := upper - 1
	if math.Abs(position-volume.slicePositions[lower]) <= volumeGeometryTolerance {
		return bilinearImageSample(volume.images[lower], point)
	}
	if math.Abs(position-volume.slicePositions[upper]) <= volumeGeometryTolerance {
		return bilinearImageSample(volume.images[upper], point)
	}
	lowerValue, lowerValid, err := bilinearImageSample(volume.images[lower], point)
	if err != nil || !lowerValid {
		return 0, false, err
	}
	upperValue, upperValid, err := bilinearImageSample(volume.images[upper], point)
	if err != nil || !upperValid {
		return 0, false, err
	}
	span := volume.slicePositions[upper] - volume.slicePositions[lower]
	weight := (position - volume.slicePositions[lower]) / span
	return lowerValue*(1-weight) + upperValue*weight, true, nil
}

// Cut materializes an arbitrary plane with deterministic, bounded parallelism.
func (volume *VolumeData) Cut(ctx context.Context, spec CutSpec, options CutOptions) (*Slice, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	normalized, err := normalizeCutSpec(spec)
	if err != nil {
		return nil, err
	}
	count := normalized.Rows * normalized.Columns
	workers := options.Workers
	if workers <= 0 {
		workers = 1
	}
	if workers > count {
		workers = count
	}
	slice := &Slice{
		Spec:   normalized,
		Values: make([]float64, count),
		Valid:  make([]bool, count),
	}
	errorsByIndex := make([]error, count)

	var wait sync.WaitGroup
	wait.Add(workers)
	for worker := 0; worker < workers; worker++ {
		go func(worker int) {
			defer wait.Done()
			for index := worker; index < count; index += workers {
				if err := ctx.Err(); err != nil {
					errorsByIndex[index] = err
					return
				}
				row := index / normalized.Columns
				column := index % normalized.Columns
				point := addPatientPoint(
					normalized.TopLeft,
					normalized.RowDirection,
					float64(column)*normalized.PixelSpacingColumns,
					normalized.ColumnDirection,
					float64(row)*normalized.PixelSpacingRows,
				)
				value, valid, err := volume.Sample(point)
				if err != nil {
					errorsByIndex[index] = err
					continue
				}
				slice.Values[index] = value
				slice.Valid[index] = valid
			}
		}(worker)
	}
	wait.Wait()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	for index, err := range errorsByIndex {
		if err != nil {
			return nil, fmt.Errorf("calculate output pixel %d: %w", index, err)
		}
	}
	return slice, nil
}

func bilinearImageSample(image *ImageData, patient math3d.Point3) (float64, bool, error) {
	point, err := image.geometry.PatientToImage(patient)
	if err != nil {
		return 0, false, err
	}
	maximumX := float64(image.geometry.Columns - 1)
	maximumY := float64(image.geometry.Rows - 1)
	if point.X < -volumeGeometryTolerance || point.Y < -volumeGeometryTolerance ||
		point.X > maximumX+volumeGeometryTolerance || point.Y > maximumY+volumeGeometryTolerance {
		return 0, false, nil
	}
	point.X = math.Max(0, math.Min(maximumX, point.X))
	point.Y = math.Max(0, math.Min(maximumY, point.Y))
	x0, y0 := int(math.Floor(point.X)), int(math.Floor(point.Y))
	x1 := min(x0+1, image.geometry.Columns-1)
	y1 := min(y0+1, image.geometry.Rows-1)
	dx, dy := point.X-float64(x0), point.Y-float64(y0)
	samples := []struct {
		x, y   int
		weight float64
	}{
		{x0, y0, (1 - dx) * (1 - dy)},
		{x1, y0, dx * (1 - dy)},
		{x0, y1, (1 - dx) * dy},
		{x1, y1, dx * dy},
	}
	value := 0.0
	for _, sample := range samples {
		if sample.weight <= 0 {
			continue
		}
		entry, valid, err := image.ValueAt(sample.x, sample.y)
		if err != nil {
			return 0, false, err
		}
		if !valid {
			return 0, false, nil
		}
		value += entry * sample.weight
	}
	return value, true, nil
}

func normalizeCutSpec(spec CutSpec) (CutSpec, error) {
	if spec.Rows <= 0 || spec.Columns <= 0 {
		return CutSpec{}, fmt.Errorf("cut dimensions must be positive")
	}
	if spec.Rows > maxCutDimension || spec.Columns > maxCutDimension {
		return CutSpec{}, fmt.Errorf("cut dimensions %dx%d exceed DICOM limit %d", spec.Rows, spec.Columns, maxCutDimension)
	}
	if int64(spec.Rows)*int64(spec.Columns) > maxCutPixels {
		return CutSpec{}, fmt.Errorf("cut dimensions %dx%d exceed pixel limit %d", spec.Rows, spec.Columns, maxCutPixels)
	}
	if spec.PixelSpacingRows <= 0 || spec.PixelSpacingColumns <= 0 ||
		math.IsNaN(spec.PixelSpacingRows) || math.IsNaN(spec.PixelSpacingColumns) ||
		math.IsInf(spec.PixelSpacingRows, 0) || math.IsInf(spec.PixelSpacingColumns, 0) {
		return CutSpec{}, fmt.Errorf("cut pixel spacing must be positive and finite")
	}
	row, err := spec.RowDirection.Normalize()
	if err != nil {
		return CutSpec{}, fmt.Errorf("cut row direction: %w", err)
	}
	column, err := spec.ColumnDirection.Normalize()
	if err != nil {
		return CutSpec{}, fmt.Errorf("cut column direction: %w", err)
	}
	if math.Abs(row.X*column.X+row.Y*column.Y+row.Z*column.Z) > volumeGeometryTolerance {
		return CutSpec{}, fmt.Errorf("cut row and column directions must be perpendicular")
	}
	spec.RowDirection = row
	spec.ColumnDirection = column
	return spec, nil
}

func addPatientPoint(origin math3d.Point3, row math3d.Vector3, rowScale float64, column math3d.Vector3, columnScale float64) math3d.Point3 {
	return math3d.Point3{
		X: origin.X + row.X*rowScale + column.X*columnScale,
		Y: origin.Y + row.Y*rowScale + column.Y*columnScale,
		Z: origin.Z + row.Z*rowScale + column.Z*columnScale,
	}
}
