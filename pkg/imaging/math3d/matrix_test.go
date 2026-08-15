// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package math3d

import "testing"

func TestMatrix4InverseRoundTrip(t *testing.T) {
	matrix := IdentityMatrix4()
	matrix.Set(0, 0, 2)
	matrix.Set(1, 1, 3)
	matrix.Set(2, 2, 4)
	matrix.Set(0, 3, 10)
	matrix.Set(1, 3, -5)
	matrix.Set(2, 3, 7)

	inverse, err := matrix.Inverse()
	if err != nil {
		t.Fatalf("Inverse() error = %v", err)
	}
	point := Point3{X: 2, Y: 3, Z: 4}
	if got := inverse.TransformPoint(matrix.TransformPoint(point)); !pointClose(got, point) {
		t.Fatalf("inverse round trip = %+v, want %+v", got, point)
	}
}

func TestMatrix4InverseRejectsSingularMatrix(t *testing.T) {
	var matrix Matrix4
	if _, err := matrix.Inverse(); err == nil {
		t.Fatal("Inverse() accepted a singular matrix")
	}
}

func TestRectCorners(t *testing.T) {
	rect := Rect{X: 2, Y: 3, Width: 4, Height: 5}
	want := [4]Point2{{X: 2, Y: 3}, {X: 6, Y: 3}, {X: 6, Y: 8}, {X: 2, Y: 8}}
	if got := rect.Corners(); got != want {
		t.Fatalf("Corners() = %+v, want %+v", got, want)
	}
}
