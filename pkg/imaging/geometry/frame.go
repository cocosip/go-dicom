// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package geometry extracts DICOM frame geometry and converts between image
// pixel coordinates and the patient coordinate system.
package geometry

import (
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const (
	directionTolerance     = 1e-9
	orthogonalityTolerance = 1e-4
)

// Orientation identifies the patient axis most nearly perpendicular to a
// frame. Oblique frames are classified by their nearest cardinal normal.
type Orientation int

const (
	// OrientationNone indicates that no patient-space orientation is available.
	OrientationNone Orientation = iota
	// OrientationAxial identifies a frame whose normal is nearest the patient Z axis.
	OrientationAxial
	// OrientationSagittal identifies a frame whose normal is nearest the patient X axis.
	OrientationSagittal
	// OrientationCoronal identifies a frame whose normal is nearest the patient Y axis.
	OrientationCoronal
)

// Type describes how much spatial information a frame contains.
type Type int

const (
	// GeometryNone indicates that pixel spacing is unavailable.
	GeometryNone Type = iota
	// GeometryPlane indicates two-dimensional pixel spacing without patient orientation.
	GeometryPlane
	// GeometryVolume indicates complete patient-space position and orientation.
	GeometryVolume
)

// FrameGeometry contains the spatial metadata for one zero-based DICOM frame.
// Image coordinates use pixel centers: (0,0) is the center of the first pixel.
type FrameGeometry struct {
	Type Type

	FrameOfReferenceUID string
	Columns             int
	Rows                int

	PixelSpacingColumns float64
	PixelSpacingRows    float64

	DirectionRow    math3d.Vector3
	DirectionColumn math3d.Vector3
	DirectionNormal math3d.Vector3
	Orientation     Orientation

	TopLeft     math3d.Point3
	TopRight    math3d.Point3
	BottomLeft  math3d.Point3
	BottomRight math3d.Point3
}

// NewFrameGeometry extracts classic or enhanced multi-frame geometry from ds.
func NewFrameGeometry(ds *dataset.Dataset, frame int) (*FrameGeometry, error) {
	if ds == nil {
		return nil, fmt.Errorf("frame geometry dataset is nil")
	}
	if frame < 0 {
		return nil, fmt.Errorf("frame index %d is negative", frame)
	}
	if err := validateFrameIndex(ds, frame); err != nil {
		return nil, err
	}
	columns, err := ds.GetUInt16(tag.Columns, 0)
	if err != nil || columns == 0 {
		return nil, fmt.Errorf("read Columns: %w", nonNilError(err, "value must be positive"))
	}
	rows, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil || rows == 0 {
		return nil, fmt.Errorf("read Rows: %w", nonNilError(err, "value must be positive"))
	}

	functional := functionalGroupValues(ds, frame)
	spacing, spacingFound, err := findSpacing(ds, functional)
	if err != nil {
		return nil, err
	}
	position, positionFound, err := findDecimalValues(ds, functional, tag.ImagePositionPatient)
	if err != nil {
		return nil, err
	}
	orientation, orientationFound, err := findDecimalValues(ds, functional, tag.ImageOrientationPatient)
	if err != nil {
		return nil, err
	}

	geometry := &FrameGeometry{
		FrameOfReferenceUID: ds.TryGetString(tag.FrameOfReferenceUID),
		Columns:             int(columns),
		Rows:                int(rows),
		DirectionRow:        math3d.Vector3{X: 1},
		DirectionColumn:     math3d.Vector3{Y: 1},
	}
	if !spacingFound {
		geometry.updateCorners()
		return geometry, nil
	}
	geometry.PixelSpacingRows = spacing[0]
	geometry.PixelSpacingColumns = spacing[1]
	geometry.Type = GeometryPlane

	if positionFound != orientationFound {
		return nil, fmt.Errorf("image position and image orientation must both be present")
	}
	if positionFound {
		if len(position) != 3 {
			return nil, fmt.Errorf("image position patient has %d values, want 3", len(position))
		}
		if len(orientation) != 6 {
			return nil, fmt.Errorf("image orientation patient has %d values, want 6", len(orientation))
		}
		return NewFrameGeometryFromValues(
			geometry.FrameOfReferenceUID,
			math3d.Point3{X: position[0], Y: position[1], Z: position[2]},
			math3d.Vector3{X: orientation[0], Y: orientation[1], Z: orientation[2]},
			math3d.Vector3{X: orientation[3], Y: orientation[4], Z: orientation[5]},
			spacing,
			geometry.Columns,
			geometry.Rows,
		)
	}
	geometry.updateCorners()
	return geometry, nil
}

// NewFrameGeometryFromValues constructs volume geometry from explicit DICOM
// patient-space values. spacing is ordered as row spacing, column spacing.
func NewFrameGeometryFromValues(
	frameOfReferenceUID string,
	position math3d.Point3,
	rowDirection math3d.Vector3,
	columnDirection math3d.Vector3,
	spacing [2]float64,
	columns int,
	rows int,
) (*FrameGeometry, error) {
	if columns <= 0 || rows <= 0 {
		return nil, fmt.Errorf("frame dimensions must be positive")
	}
	if !finitePoint(position) {
		return nil, fmt.Errorf("patient position must contain finite values")
	}
	if !positiveFinite(spacing[0]) || !positiveFinite(spacing[1]) {
		return nil, fmt.Errorf("pixel spacing must contain positive finite values")
	}
	row, err := rowDirection.Normalize()
	if err != nil {
		return nil, fmt.Errorf("row direction: %w", err)
	}
	column, err := columnDirection.Normalize()
	if err != nil {
		return nil, fmt.Errorf("column direction: %w", err)
	}
	projection := dot(row, column)
	if math.Abs(projection) > orthogonalityTolerance {
		return nil, fmt.Errorf("row and column directions are not perpendicular")
	}
	column, err = subtractVector(column, scale(row, projection)).Normalize()
	if err != nil {
		return nil, fmt.Errorf("orthogonalize column direction: %w", err)
	}
	normal, err := row.Cross(column).Normalize()
	if err != nil {
		return nil, fmt.Errorf("frame normal: %w", err)
	}

	geometry := &FrameGeometry{
		Type:                GeometryVolume,
		FrameOfReferenceUID: frameOfReferenceUID,
		Columns:             columns,
		Rows:                rows,
		PixelSpacingColumns: spacing[1],
		PixelSpacingRows:    spacing[0],
		DirectionRow:        row,
		DirectionColumn:     column,
		DirectionNormal:     normal,
		Orientation:         orientationFromNormal(normal),
		TopLeft:             position,
	}
	geometry.updateCorners()
	return geometry, nil
}

// HasGeometry reports whether pixel spacing is available.
func (g *FrameGeometry) HasGeometry() bool {
	return g != nil && g.Type != GeometryNone
}

// ImageToPatient maps a zero-based pixel-center coordinate into patient space.
func (g *FrameGeometry) ImageToPatient(point math3d.Point2) (math3d.Point3, error) {
	if !g.HasGeometry() {
		return math3d.Point3{}, fmt.Errorf("cannot transform an image point without geometry")
	}
	return addPoint(
		g.TopLeft,
		scale(g.DirectionRow, point.X*g.PixelSpacingColumns),
		scale(g.DirectionColumn, point.Y*g.PixelSpacingRows),
	), nil
}

// PatientToImage projects a patient-space point onto the frame and returns a
// zero-based pixel-center coordinate.
func (g *FrameGeometry) PatientToImage(point math3d.Point3) (math3d.Point2, error) {
	if !g.HasGeometry() {
		return math3d.Point2{}, fmt.Errorf("cannot transform a patient point without geometry")
	}
	delta := subtractPoint(point, g.TopLeft)
	return math3d.Point2{
		X: dot(delta, g.DirectionRow) / g.PixelSpacingColumns,
		Y: dot(delta, g.DirectionColumn) / g.PixelSpacingRows,
	}, nil
}

func (g *FrameGeometry) updateCorners() {
	g.TopRight = addPoint(g.TopLeft, scale(g.DirectionRow, float64(max(g.Columns-1, 0))*g.PixelSpacingColumns))
	g.BottomLeft = addPoint(g.TopLeft, scale(g.DirectionColumn, float64(max(g.Rows-1, 0))*g.PixelSpacingRows))
	g.BottomRight = addPoint(g.BottomLeft, subtractPoint(g.TopRight, g.TopLeft))
}

func validateFrameIndex(ds *dataset.Dataset, frame int) error {
	if sequence, err := ds.GetSequence(tag.PerFrameFunctionalGroupsSequence); err == nil {
		if frame >= sequence.Count() {
			return fmt.Errorf("frame index %d is outside Per-Frame Functional Groups Sequence with %d items", frame, sequence.Count())
		}
		return nil
	}
	if value, ok := ds.Get(tag.NumberOfFrames); ok {
		frames, valid := value.(*element.IntegerString)
		if !valid {
			return fmt.Errorf("number of frames is not an IntegerString")
		}
		count, err := frames.GetInt(0)
		if err != nil {
			return fmt.Errorf("read Number of Frames: %w", err)
		}
		if frame >= count {
			return fmt.Errorf("frame index %d is outside Number of Frames %d", frame, count)
		}
		return nil
	}
	if frame != 0 {
		return fmt.Errorf("frame index %d is invalid for a single-frame dataset", frame)
	}
	return nil
}

func findSpacing(primary, functional *dataset.Dataset) ([2]float64, bool, error) {
	for _, candidate := range []struct {
		dataset *dataset.Dataset
		tag     *tag.Tag
	}{
		{primary, tag.ImagerPixelSpacing},
		{primary, tag.PixelSpacing},
		{primary, tag.NominalScannedPixelSpacing},
		{functional, tag.PixelSpacing},
		{functional, tag.ImagerPixelSpacing},
	} {
		values, found, err := decimalValues(candidate.dataset, candidate.tag)
		if err != nil {
			return [2]float64{}, false, err
		}
		if !found {
			continue
		}
		if len(values) != 2 || !positiveFinite(values[0]) || !positiveFinite(values[1]) {
			return [2]float64{}, false, fmt.Errorf("%s must contain two positive finite values", candidate.tag)
		}
		return [2]float64{values[0], values[1]}, true, nil
	}
	return [2]float64{}, false, nil
}

func findDecimalValues(primary, functional *dataset.Dataset, target *tag.Tag) ([]float64, bool, error) {
	if values, found, err := decimalValues(primary, target); found || err != nil {
		return values, found, err
	}
	return decimalValues(functional, target)
}

func decimalValues(ds *dataset.Dataset, target *tag.Tag) ([]float64, bool, error) {
	if ds == nil {
		return nil, false, nil
	}
	value, found := ds.Get(target)
	if !found {
		return nil, false, nil
	}
	decimal, ok := value.(*element.DecimalString)
	if !ok {
		return nil, true, fmt.Errorf("element %s is not DecimalString", target)
	}
	values, err := decimal.GetFloats()
	if err != nil {
		return nil, true, fmt.Errorf("read %s: %w", target, err)
	}
	for _, entry := range values {
		if math.IsNaN(entry) || math.IsInf(entry, 0) {
			return nil, true, fmt.Errorf("element %s contains a non-finite value", target)
		}
	}
	return values, true, nil
}

func functionalGroupValues(ds *dataset.Dataset, frame int) *dataset.Dataset {
	values := dataset.New()
	mergeFunctionalGroup(values, ds, tag.SharedFunctionalGroupsSequence, 0)
	mergeFunctionalGroup(values, ds, tag.PerFrameFunctionalGroupsSequence, frame)
	return values
}

func mergeFunctionalGroup(values, source *dataset.Dataset, target *tag.Tag, itemIndex int) {
	sequence, err := source.GetSequence(target)
	if err != nil || itemIndex < 0 || itemIndex >= sequence.Count() {
		return
	}
	item := sequence.GetItem(itemIndex)
	if item == nil {
		return
	}
	for _, value := range item.Elements() {
		nested, ok := value.(*dataset.Sequence)
		if !ok || nested.Count() == 0 || nested.GetItem(0) == nil {
			continue
		}
		for _, nestedValue := range nested.GetItem(0).Elements() {
			_ = values.AddOrUpdate(nestedValue)
		}
	}
}

func orientationFromNormal(normal math3d.Vector3) Orientation {
	axis := normal.NearestAxis()
	switch {
	case axis.X != 0:
		return OrientationSagittal
	case axis.Y != 0:
		return OrientationCoronal
	case axis.Z != 0:
		return OrientationAxial
	default:
		return OrientationNone
	}
}

func positiveFinite(value float64) bool {
	return value > 0 && !math.IsNaN(value) && !math.IsInf(value, 0)
}

func finitePoint(point math3d.Point3) bool {
	return !math.IsNaN(point.X) && !math.IsNaN(point.Y) && !math.IsNaN(point.Z) &&
		!math.IsInf(point.X, 0) && !math.IsInf(point.Y, 0) && !math.IsInf(point.Z, 0)
}

func dot(a, b math3d.Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func scale(vector math3d.Vector3, factor float64) math3d.Vector3 {
	return math3d.Vector3{X: vector.X * factor, Y: vector.Y * factor, Z: vector.Z * factor}
}

func subtractVector(a, b math3d.Vector3) math3d.Vector3 {
	return math3d.Vector3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func addPoint(point math3d.Point3, vectors ...math3d.Vector3) math3d.Point3 {
	for _, vector := range vectors {
		point.X += vector.X
		point.Y += vector.Y
		point.Z += vector.Z
	}
	return point
}

func subtractPoint(a, b math3d.Point3) math3d.Vector3 {
	return math3d.Vector3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func nonNilError(err error, message string) error {
	if err != nil {
		return err
	}
	return fmt.Errorf("%s", message)
}
