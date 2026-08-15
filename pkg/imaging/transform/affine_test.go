// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transform

import (
	"math"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const transformTolerance = 1e-9

func TestAffineCompositionOrder(t *testing.T) {
	matrix := Identity().
		Then(Scale(2, 3)).
		Then(Rotate(90)).
		Then(Translate(10, 20))

	if got := matrix.Apply(math3d.Point2{X: 1, Y: 2}); !pointClose(got, math3d.Point2{X: 4, Y: 22}) {
		t.Fatalf("Apply() = %+v, want {4,22}", got)
	}
}

func TestAffineFlipAndBounds(t *testing.T) {
	point := math3d.Point2{X: 3, Y: 4}
	if got := FlipX().Apply(point); got != (math3d.Point2{X: -3, Y: 4}) {
		t.Fatalf("FlipX().Apply() = %+v", got)
	}
	if got := FlipY().Apply(point); got != (math3d.Point2{X: 3, Y: -4}) {
		t.Fatalf("FlipY().Apply() = %+v", got)
	}

	bounds, err := Rotate(90).Bounds(math3d.Rect{Width: 2, Height: 1})
	if err != nil {
		t.Fatalf("Bounds() error = %v", err)
	}
	want := math3d.Rect{X: -1, Width: 1, Height: 2}
	if !rectClose(bounds, want) {
		t.Fatalf("Bounds() = %+v, want %+v", bounds, want)
	}
}

func TestBestFitCentersAndPreservesAspectRatio(t *testing.T) {
	matrix, err := BestFit(
		math3d.Rect{Width: 100, Height: 50},
		math3d.Rect{X: 10, Y: 20, Width: 200, Height: 200},
	)
	if err != nil {
		t.Fatalf("BestFit() error = %v", err)
	}
	if got := matrix.Apply(math3d.Point2{}); !pointClose(got, math3d.Point2{X: 10, Y: 70}) {
		t.Fatalf("top-left = %+v, want {10,70}", got)
	}
	if got := matrix.Apply(math3d.Point2{X: 100, Y: 50}); !pointClose(got, math3d.Point2{X: 210, Y: 170}) {
		t.Fatalf("bottom-right = %+v, want {210,170}", got)
	}
	if _, err := BestFit(math3d.Rect{}, math3d.Rect{Width: 1, Height: 1}); err == nil {
		t.Fatal("BestFit() accepted an empty source")
	}
}

func TestSpatialTransformStateAndCenteredAffine(t *testing.T) {
	spatial := NewSpatialTransform()
	if spatial.IsTransformed() {
		t.Fatal("new SpatialTransform is transformed")
	}
	spatial.FlipX = true
	if !spatial.IsTransformed() {
		t.Fatal("IsTransformed() ignored FlipX")
	}
	spatial.Reset()
	spatial.Rotate(450)
	if spatial.Rotation != 90 {
		t.Fatalf("Rotation = %d, want 90", spatial.Rotation)
	}
	spatial.Rotation = 0
	spatial.Scale = 2
	spatial.Pan = math3d.Point2{X: 3, Y: -1}
	matrix, err := spatial.Affine(math3d.Rect{Width: 2, Height: 2})
	if err != nil {
		t.Fatalf("Affine() error = %v", err)
	}
	if got := matrix.Apply(math3d.Point2{X: 2, Y: 1}); !pointClose(got, math3d.Point2{X: 6, Y: 0}) {
		t.Fatalf("centered transform = %+v, want {6,0}", got)
	}
}

func pointClose(a, b math3d.Point2) bool {
	return math.Abs(a.X-b.X) <= transformTolerance && math.Abs(a.Y-b.Y) <= transformTolerance
}

func rectClose(a, b math3d.Rect) bool {
	return math.Abs(a.X-b.X) <= transformTolerance && math.Abs(a.Y-b.Y) <= transformTolerance &&
		math.Abs(a.Width-b.Width) <= transformTolerance && math.Abs(a.Height-b.Height) <= transformTolerance
}
