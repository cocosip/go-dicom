// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package geometry extracts DICOM frame geometry and converts between image
// pixel coordinates and the patient coordinate system.
package geometry

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/internal/dicomlut"
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

	functional := dicomlut.FunctionalGroupValues(ds, frame)
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
	projection := row.Dot(column)
	if math.Abs(projection) > orthogonalityTolerance {
		return nil, fmt.Errorf("row and column directions are not perpendicular")
	}
	column, err = column.Sub(row.Scale(projection)).Normalize()
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
	return g.TopLeft.
		Add(g.DirectionRow.Scale(point.X * g.PixelSpacingColumns)).
		Add(g.DirectionColumn.Scale(point.Y * g.PixelSpacingRows)), nil
}

// PatientToImage projects a patient-space point onto the frame and returns a
// zero-based pixel-center coordinate.
func (g *FrameGeometry) PatientToImage(point math3d.Point3) (math3d.Point2, error) {
	if !g.HasGeometry() {
		return math3d.Point2{}, fmt.Errorf("cannot transform a patient point without geometry")
	}
	delta := point.Sub(g.TopLeft)
	return math3d.Point2{
		X: delta.Dot(g.DirectionRow) / g.PixelSpacingColumns,
		Y: delta.Dot(g.DirectionColumn) / g.PixelSpacingRows,
	}, nil
}

func (g *FrameGeometry) updateCorners() {
	g.TopRight = g.TopLeft.Add(g.DirectionRow.Scale(float64(max(g.Columns-1, 0)) * g.PixelSpacingColumns))
	g.BottomLeft = g.TopLeft.Add(g.DirectionColumn.Scale(float64(max(g.Rows-1, 0)) * g.PixelSpacingRows))
	g.BottomRight = g.BottomLeft.Add(g.TopRight.Sub(g.TopLeft))
}

func validateFrameIndex(ds *dataset.Dataset, frame int) error {
	if sequence, err := ds.GetSequence(tag.PerFrameFunctionalGroupsSequence); err == nil {
		if frame >= sequence.Count() {
			return fmt.Errorf("frame index %d is outside Per-Frame Functional Groups Sequence with %d items", frame, sequence.Count())
		}
		return nil
	}
	if value, ok := ds.Get(tag.NumberOfFrames); ok {
		count, err := frameCountValue(value)
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
		// PS3.3 C.7.6.2.1.1: Pixel Spacing is measured in the patient (and
		// already corrected for magnification), while Imager Pixel Spacing is
		// measured at the detector. Prefer the patient-relative attribute so
		// DX/MG distances are physical; Imager Pixel Spacing is the last resort.
		{primary, tag.PixelSpacing},
		{primary, tag.NominalScannedPixelSpacing},
		{primary, tag.ImagerPixelSpacing},
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
	values, err := geometryDecimalValues(value)
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

func frameCountValue(value element.Element) (int, error) {
	switch typed := value.(type) {
	case *element.IntegerString:
		return typed.GetInt(0)
	case *element.String:
		if typed.ValueRepresentation().Code() != vr.CodeIS {
			return 0, fmt.Errorf("value representation is %s, want IS", typed.ValueRepresentation().Code())
		}
		return strconv.Atoi(strings.TrimSpace(typed.GetValue(0)))
	default:
		return 0, fmt.Errorf("element type is %T, want Integer String", value)
	}
}

func geometryDecimalValues(value element.Element) ([]float64, error) {
	switch typed := value.(type) {
	case *element.DecimalString:
		return typed.GetFloats()
	case *element.String:
		if typed.ValueRepresentation().Code() != vr.CodeDS {
			return nil, fmt.Errorf("value representation is %s, want DS", typed.ValueRepresentation().Code())
		}
		stringsValues := typed.GetValues()
		values := make([]float64, len(stringsValues))
		for index, entry := range stringsValues {
			parsed, err := strconv.ParseFloat(strings.TrimSpace(entry), 64)
			if err != nil {
				return nil, fmt.Errorf("parse value %d: %w", index, err)
			}
			values[index] = parsed
		}
		return values, nil
	default:
		return nil, fmt.Errorf("element type is %T, want Decimal String", value)
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

func nonNilError(err error, message string) error {
	if err != nil {
		return err
	}
	return fmt.Errorf("%s", message)
}
