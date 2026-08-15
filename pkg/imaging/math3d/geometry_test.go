// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package math3d

import (
	"math"
	"testing"
)

const testTolerance = 1e-9

func TestVector3Operations(t *testing.T) {
	x := Vector3{X: 1}
	y := Vector3{Y: 1}

	if got := x.Cross(y); !vectorClose(got, Vector3{Z: 1}) {
		t.Fatalf("Cross() = %+v, want {Z:1}", got)
	}
	if got := (Vector3{X: 3, Y: 4}).Length(); math.Abs(got-5) > testTolerance {
		t.Fatalf("Length() = %v, want 5", got)
	}
	if got, err := (Vector3{X: 3, Y: 4}).Normalize(); err != nil || !vectorClose(got, Vector3{X: 0.6, Y: 0.8}) {
		t.Fatalf("Normalize() = %+v, %v", got, err)
	}
	if _, err := (Vector3{}).Normalize(); err == nil {
		t.Fatal("Normalize() accepted a zero vector")
	}
	if got := (Vector3{X: -0.2, Y: 0.7, Z: 0.1}).NearestAxis(); got != (Vector3{Y: 1}) {
		t.Fatalf("NearestAxis() = %+v, want +Y", got)
	}
}

func TestPlaneIntersectSegment(t *testing.T) {
	plane, err := NewPlane(Vector3{Z: 1}, Point3{Z: 5})
	if err != nil {
		t.Fatalf("NewPlane() error = %v", err)
	}

	point, ok := plane.IntersectSegment(Segment3{
		A: Point3{X: 1, Y: 2, Z: 0},
		B: Point3{X: 1, Y: 2, Z: 10},
	})
	if !ok || !pointClose(point, Point3{X: 1, Y: 2, Z: 5}) {
		t.Fatalf("IntersectSegment() = %+v, %v", point, ok)
	}

	if _, ok := plane.IntersectSegment(Segment3{A: Point3{}, B: Point3{X: 1}}); ok {
		t.Fatal("IntersectSegment() reported an intersection for a parallel segment")
	}
}

func TestBoundingBox(t *testing.T) {
	box, err := BoundingBox([]Point3{
		{X: 4, Y: -2, Z: 8},
		{X: -1, Y: 7, Z: 3},
		{X: 2, Y: 1, Z: -5},
	})
	if err != nil {
		t.Fatalf("BoundingBox() error = %v", err)
	}
	if box.Min != (Point3{X: -1, Y: -2, Z: -5}) || box.Max != (Point3{X: 4, Y: 7, Z: 8}) {
		t.Fatalf("BoundingBox() = %+v", box)
	}
	if _, err := BoundingBox(nil); err == nil {
		t.Fatal("BoundingBox(nil) succeeded")
	}
}

func vectorClose(a, b Vector3) bool {
	return math.Abs(a.X-b.X) <= testTolerance && math.Abs(a.Y-b.Y) <= testTolerance && math.Abs(a.Z-b.Z) <= testTolerance
}

func pointClose(a, b Point3) bool {
	return math.Abs(a.X-b.X) <= testTolerance && math.Abs(a.Y-b.Y) <= testTolerance && math.Abs(a.Z-b.Z) <= testTolerance
}
