// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transform_test

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
)

func ExampleAffine2D_Then() {
	matrix := transform.Identity().
		Then(transform.Scale(2, 3)).
		Then(transform.Rotate(90)).
		Then(transform.Translate(10, 20))

	point := matrix.Apply(math3d.Point2{X: 1, Y: 2})
	fmt.Printf("(%.0f, %.0f)\n", point.X, point.Y)
	// Output:
	// (4, 22)
}

func ExampleBestFit() {
	matrix, err := transform.BestFit(
		math3d.Rect{Width: 100, Height: 50},
		math3d.Rect{X: 10, Y: 20, Width: 200, Height: 200},
	)
	if err != nil {
		panic(err)
	}

	topLeft := matrix.Apply(math3d.Point2{})
	bottomRight := matrix.Apply(math3d.Point2{X: 100, Y: 50})
	fmt.Printf("top-left: (%.0f, %.0f)\n", topLeft.X, topLeft.Y)
	fmt.Printf("bottom-right: (%.0f, %.0f)\n", bottomRight.X, bottomRight.Y)
	// Output:
	// top-left: (10, 70)
	// bottom-right: (210, 170)
}

func ExampleSpatialTransform_Affine() {
	viewer := transform.NewSpatialTransform()
	viewer.Scale = 2
	viewer.Pan = math3d.Point2{X: 3, Y: -1}

	matrix, err := viewer.Affine(math3d.Rect{Width: 2, Height: 2})
	if err != nil {
		panic(err)
	}
	point := matrix.Apply(math3d.Point2{X: 2, Y: 1})
	fmt.Printf("(%.0f, %.0f)\n", point.X, point.Y)
	// Output:
	// (6, 0)
}
