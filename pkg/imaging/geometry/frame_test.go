// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package geometry

import (
	"math"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const geometryTolerance = 1e-9

func TestClassicFrameGeometryCoordinateRoundTrip(t *testing.T) {
	ds := classicGeometryDataset(
		[]float64{10, 20, 30},
		[]float64{1, 0, 0, 0, 1, 0},
		[]float64{2, 3},
		4,
		3,
	)

	geometry, err := NewFrameGeometry(ds, 0)
	if err != nil {
		t.Fatalf("NewFrameGeometry() error = %v", err)
	}
	if geometry.Type != GeometryVolume || geometry.Orientation != OrientationAxial {
		t.Fatalf("geometry type/orientation = %v/%v", geometry.Type, geometry.Orientation)
	}
	if geometry.DirectionNormal != (math3d.Vector3{Z: 1}) {
		t.Fatalf("DirectionNormal = %+v", geometry.DirectionNormal)
	}
	if geometry.TopRight != (math3d.Point3{X: 19, Y: 20, Z: 30}) {
		t.Fatalf("TopRight = %+v, want x=19", geometry.TopRight)
	}
	if geometry.BottomLeft != (math3d.Point3{X: 10, Y: 24, Z: 30}) {
		t.Fatalf("BottomLeft = %+v, want y=24", geometry.BottomLeft)
	}

	patient, err := geometry.ImageToPatient(math3d.Point2{X: 2, Y: 1})
	if err != nil {
		t.Fatalf("ImageToPatient() error = %v", err)
	}
	if !point3Close(patient, math3d.Point3{X: 16, Y: 22, Z: 30}) {
		t.Fatalf("ImageToPatient() = %+v", patient)
	}
	imagePoint, err := geometry.PatientToImage(patient)
	if err != nil {
		t.Fatalf("PatientToImage() error = %v", err)
	}
	if !point2Close(imagePoint, math3d.Point2{X: 2, Y: 1}) {
		t.Fatalf("PatientToImage() = %+v", imagePoint)
	}
}

func TestEnhancedMultiFrameGeometryUsesSharedAndPerFrameMacros(t *testing.T) {
	ds := dataset.New()
	mustAdd(t, ds,
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.Rows, []uint16{2}),
		element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{2}),
		element.NewString(tag.FrameOfReferenceUID, vr.UI, []string{"1.2.3"}),
	)

	shared := dataset.New()
	addNestedItem(t, shared, tag.PixelMeasuresSequence, decimalDataset(tag.PixelSpacing, []float64{0.5, 0.75}))
	addNestedItem(t, shared, tag.PlaneOrientationSequence, decimalDataset(tag.ImageOrientationPatient, []float64{1, 0, 0, 0, 1, 0}))
	addNestedItem(t, shared, tag.PlanePositionSequence, decimalDataset(tag.ImagePositionPatient, []float64{0, 0, 1}))
	mustAdd(t, ds, dataset.NewSequenceWithItems(tag.SharedFunctionalGroupsSequence, []*dataset.Dataset{shared}))

	frames := make([]*dataset.Dataset, 2)
	for frame, z := range []float64{2, 7} {
		frames[frame] = dataset.New()
		addNestedItem(t, frames[frame], tag.PlanePositionSequence, decimalDataset(tag.ImagePositionPatient, []float64{0, 0, z}))
	}
	mustAdd(t, ds, dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, frames))

	geometry, err := NewFrameGeometry(ds, 1)
	if err != nil {
		t.Fatalf("NewFrameGeometry(frame 1) error = %v", err)
	}
	if geometry.TopLeft.Z != 7 {
		t.Fatalf("TopLeft.Z = %v, want per-frame value 7", geometry.TopLeft.Z)
	}
	if geometry.PixelSpacingRows != 0.5 || geometry.PixelSpacingColumns != 0.75 {
		t.Fatalf("pixel spacing = %v/%v", geometry.PixelSpacingRows, geometry.PixelSpacingColumns)
	}
	if _, err := NewFrameGeometry(ds, 2); err == nil {
		t.Fatal("NewFrameGeometry() accepted an out-of-range enhanced frame")
	}
}

func TestFrameGeometryUsesFODICOMSpacingPrecedence(t *testing.T) {
	ds := classicGeometryDataset(
		[]float64{0, 0, 0},
		[]float64{1, 0, 0, 0, 1, 0},
		[]float64{2, 3},
		2,
		2,
	)
	mustAdd(t, ds, element.NewDecimalStringFromFloat(tag.ImagerPixelSpacing, []float64{4, 5}))
	shared := dataset.New()
	addNestedItem(t, shared, tag.PixelMeasuresSequence, decimalDataset(tag.PixelSpacing, []float64{6, 7}))
	mustAdd(t, ds, dataset.NewSequenceWithItems(tag.SharedFunctionalGroupsSequence, []*dataset.Dataset{shared}))

	geometry, err := NewFrameGeometry(ds, 0)
	if err != nil {
		t.Fatalf("NewFrameGeometry() error = %v", err)
	}
	if geometry.PixelSpacingRows != 4 || geometry.PixelSpacingColumns != 5 {
		t.Fatalf("pixel spacing = %v/%v, want top-level Imager Pixel Spacing 4/5", geometry.PixelSpacingRows, geometry.PixelSpacingColumns)
	}
}

func TestFrameGeometryOrientationAndRoundTrip(t *testing.T) {
	tests := []struct {
		name        string
		row         math3d.Vector3
		column      math3d.Vector3
		orientation Orientation
	}{
		{name: "axial", row: math3d.Vector3{X: 1}, column: math3d.Vector3{Y: 1}, orientation: OrientationAxial},
		{name: "sagittal", row: math3d.Vector3{Y: 1}, column: math3d.Vector3{Z: 1}, orientation: OrientationSagittal},
		{name: "coronal", row: math3d.Vector3{X: 1}, column: math3d.Vector3{Z: 1}, orientation: OrientationCoronal},
		{name: "oblique", row: math3d.Vector3{X: math.Sqrt(0.5), Y: math.Sqrt(0.5)}, column: math3d.Vector3{Z: 1}, orientation: OrientationSagittal},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			geometry, err := NewFrameGeometryFromValues(
				"1.2.3",
				math3d.Point3{X: 5, Y: 6, Z: 7},
				tc.row,
				tc.column,
				[2]float64{0.8, 0.6},
				16,
				12,
			)
			if err != nil {
				t.Fatalf("NewFrameGeometryFromValues() error = %v", err)
			}
			if geometry.Orientation != tc.orientation {
				t.Fatalf("Orientation = %v, want %v", geometry.Orientation, tc.orientation)
			}
			input := math3d.Point2{X: 3.25, Y: 4.5}
			patient, err := geometry.ImageToPatient(input)
			if err != nil {
				t.Fatalf("ImageToPatient() error = %v", err)
			}
			output, err := geometry.PatientToImage(patient)
			if err != nil {
				t.Fatalf("PatientToImage() error = %v", err)
			}
			if !point2Close(output, input) {
				t.Fatalf("round trip = %+v, want %+v", output, input)
			}
		})
	}
}

func TestFrameGeometryToleratesRoundedDirectionCosines(t *testing.T) {
	geometry, err := NewFrameGeometryFromValues(
		"1.2.3",
		math3d.Point3{},
		math3d.Vector3{X: 1, Y: 0.000001},
		math3d.Vector3{Y: 1, Z: 0.000001},
		[2]float64{1, 1},
		4,
		4,
	)
	if err != nil {
		t.Fatalf("rounded direction cosines were rejected: %v", err)
	}
	input := math3d.Point2{X: 2.5, Y: 1.25}
	patient, err := geometry.ImageToPatient(input)
	if err != nil {
		t.Fatalf("ImageToPatient() error = %v", err)
	}
	output, err := geometry.PatientToImage(patient)
	if err != nil {
		t.Fatalf("PatientToImage() error = %v", err)
	}
	if !point2Close(output, input) {
		t.Fatalf("round trip = %+v, want %+v", output, input)
	}
}

func TestFrameGeometryMissingAndInvalidGeometry(t *testing.T) {
	if _, err := NewFrameGeometry(nil, 0); err == nil {
		t.Fatal("NewFrameGeometry(nil) succeeded")
	}

	ds := dataset.New()
	mustAdd(t, ds,
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.Rows, []uint16{2}),
	)
	geometry, err := NewFrameGeometry(ds, 0)
	if err != nil {
		t.Fatalf("missing optional geometry returned error: %v", err)
	}
	if geometry.Type != GeometryNone {
		t.Fatalf("Type = %v, want GeometryNone", geometry.Type)
	}
	if _, err := geometry.ImageToPatient(math3d.Point2{}); err == nil {
		t.Fatal("ImageToPatient() succeeded without spacing")
	}

	if _, err := NewFrameGeometryFromValues(
		"1.2.3",
		math3d.Point3{},
		math3d.Vector3{X: 1},
		math3d.Vector3{X: 1},
		[2]float64{1, 1},
		2,
		2,
	); err == nil {
		t.Fatal("parallel row and column directions were accepted")
	}

	if _, err := NewFrameGeometryFromValues(
		"1.2.3",
		math3d.Point3{X: math.NaN()},
		math3d.Vector3{X: 1},
		math3d.Vector3{Y: 1},
		[2]float64{1, 1},
		2,
		2,
	); err == nil {
		t.Fatal("non-finite patient position was accepted")
	}
	if _, err := NewFrameGeometryFromValues(
		"1.2.3",
		math3d.Point3{},
		math3d.Vector3{X: math.Inf(1)},
		math3d.Vector3{Y: 1},
		[2]float64{1, 1},
		2,
		2,
	); err == nil {
		t.Fatal("non-finite row direction was accepted")
	}
}

func classicGeometryDataset(position, orientation, spacing []float64, columns, rows uint16) *dataset.Dataset {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{columns}))
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{rows}))
	_ = ds.Add(element.NewString(tag.FrameOfReferenceUID, vr.UI, []string{"1.2.3"}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.ImagePositionPatient, position))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.ImageOrientationPatient, orientation))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.PixelSpacing, spacing))
	return ds
}

func decimalDataset(target *tag.Tag, values []float64) *dataset.Dataset {
	ds := dataset.New()
	_ = ds.Add(element.NewDecimalStringFromFloat(target, values))
	return ds
}

func addNestedItem(t *testing.T, ds *dataset.Dataset, sequenceTag *tag.Tag, item *dataset.Dataset) {
	t.Helper()
	mustAdd(t, ds, dataset.NewSequenceWithItems(sequenceTag, []*dataset.Dataset{item}))
}

func mustAdd(t *testing.T, ds *dataset.Dataset, elements ...element.Element) {
	t.Helper()
	for _, value := range elements {
		if err := ds.Add(value); err != nil {
			t.Fatalf("add %s: %v", value.Tag(), err)
		}
	}
}

func point2Close(a, b math3d.Point2) bool {
	return math.Abs(a.X-b.X) <= geometryTolerance && math.Abs(a.Y-b.Y) <= geometryTolerance
}

func point3Close(a, b math3d.Point3) bool {
	return math.Abs(a.X-b.X) <= geometryTolerance && math.Abs(a.Y-b.Y) <= geometryTolerance && math.Abs(a.Z-b.Z) <= geometryTolerance
}
