// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package geometry_test

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/geometry"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

func ExampleNewFrameGeometry() {
	ds := dataset.New()
	for _, value := range []element.Element{
		element.NewUnsignedShort(tag.Columns, []uint16{4}),
		element.NewUnsignedShort(tag.Rows, []uint16{3}),
		element.NewString(tag.FrameOfReferenceUID, vr.UI, []string{testFrameOfReferenceInstanceUID}),
		element.NewDecimalStringFromFloat(tag.ImagePositionPatient, []float64{10, 20, 30}),
		element.NewDecimalStringFromFloat(tag.ImageOrientationPatient, []float64{1, 0, 0, 0, 1, 0}),
		element.NewDecimalStringFromFloat(tag.PixelSpacing, []float64{2, 3}),
	} {
		if err := ds.Add(value); err != nil {
			panic(err)
		}
	}

	frame, err := geometry.NewFrameGeometry(ds, 0)
	if err != nil {
		panic(err)
	}
	patient, err := frame.ImageToPatient(math3d.Point2{X: 2, Y: 1})
	if err != nil {
		panic(err)
	}
	imagePoint, err := frame.PatientToImage(patient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("size: %dx%d\n", frame.Columns, frame.Rows)
	fmt.Printf("axial: %t\n", frame.Orientation == geometry.OrientationAxial)
	fmt.Printf("patient: (%.0f, %.0f, %.0f)\n", patient.X, patient.Y, patient.Z)
	fmt.Printf("image: (%.0f, %.0f)\n", imagePoint.X, imagePoint.Y)
	// Output:
	// size: 4x3
	// axial: true
	// patient: (16, 22, 30)
	// image: (2, 1)
}

func ExampleIntersectionLocalizer() {
	source, err := geometry.NewFrameGeometryFromValues(
		testFrameOfReferenceInstanceUID,
		math3d.Point3{},
		math3d.Vector3{X: 1},
		math3d.Vector3{Y: 1},
		[2]float64{1, 1},
		11,
		11,
	)
	if err != nil {
		panic(err)
	}
	destination, err := geometry.NewFrameGeometryFromValues(
		testFrameOfReferenceInstanceUID,
		math3d.Point3{X: 5, Z: -5},
		math3d.Vector3{Y: 1},
		math3d.Vector3{Z: 1},
		[2]float64{1, 1},
		11,
		11,
	)
	if err != nil {
		panic(err)
	}

	start, end, ok := geometry.IntersectionLocalizer(source, destination)
	if start.X > end.X {
		start, end = end, start
	}
	fmt.Printf("can draw: %t\n", geometry.CanDrawLocalizer(source, destination))
	fmt.Printf("line: (%.0f, %.0f) to (%.0f, %.0f), found=%t\n", start.X, start.Y, end.X, end.Y, ok)
	// Output:
	// can draw: true
	// line: (0, 5) to (10, 5), found=true
}
