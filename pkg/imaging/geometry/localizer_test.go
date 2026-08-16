// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package geometry

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

func TestFrameGeometryBoundingBox(t *testing.T) {
	geometry := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{X: 10, Y: 20, Z: 30}, math3d.Vector3{X: 1}, math3d.Vector3{Y: 1}, 4, 3)

	box := geometry.BoundingBox()
	if box.Min != (math3d.Point3{X: 10, Y: 20, Z: 30}) || box.Max != (math3d.Point3{X: 13, Y: 22, Z: 30}) {
		t.Fatalf("BoundingBox() = %+v", box)
	}
}

func TestCanDrawLocalizer(t *testing.T) {
	axial := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{}, math3d.Vector3{X: 1}, math3d.Vector3{Y: 1}, 11, 11)
	sagittal := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{X: 5, Z: -5}, math3d.Vector3{Y: 1}, math3d.Vector3{Z: 1}, 11, 11)
	otherReference := mustFrameGeometry(t, "9.8.7", math3d.Point3{X: 5, Z: -5}, math3d.Vector3{Y: 1}, math3d.Vector3{Z: 1}, 11, 11)

	if !CanDrawLocalizer(axial, sagittal) {
		t.Fatal("CanDrawLocalizer() rejected orthogonal frames in one frame of reference")
	}
	if CanDrawLocalizer(axial, axial) {
		t.Fatal("CanDrawLocalizer() accepted frames with the same orientation")
	}
	if CanDrawLocalizer(axial, otherReference) {
		t.Fatal("CanDrawLocalizer() accepted different frame-of-reference UIDs")
	}
	if CanDrawLocalizer(nil, sagittal) {
		t.Fatal("CanDrawLocalizer() accepted a nil frame")
	}
}

func TestIntersectionLocalizer(t *testing.T) {
	axial := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{}, math3d.Vector3{X: 1}, math3d.Vector3{Y: 1}, 11, 11)
	sagittal := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{X: 5, Z: -5}, math3d.Vector3{Y: 1}, math3d.Vector3{Z: 1}, 11, 11)

	start, end, ok := IntersectionLocalizer(axial, sagittal)
	if !ok {
		t.Fatal("IntersectionLocalizer() reported no intersection")
	}
	wantA := math3d.Point2{X: 0, Y: 5}
	wantB := math3d.Point2{X: 10, Y: 5}
	forward := point2Close(start, wantA) && point2Close(end, wantB)
	reverse := point2Close(start, wantB) && point2Close(end, wantA)
	if forward || reverse {
		return
	}
	t.Fatalf("IntersectionLocalizer() = %+v to %+v", start, end)
}

func TestProjectionLocalizer(t *testing.T) {
	axial := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{}, math3d.Vector3{X: 1}, math3d.Vector3{Y: 1}, 11, 11)
	sagittal := mustFrameGeometry(t, testFrameOfReferenceInstanceUID, math3d.Point3{X: 5, Z: -5}, math3d.Vector3{Y: 1}, math3d.Vector3{Z: 1}, 11, 11)

	points, err := ProjectionLocalizer(axial, sagittal)
	if err != nil {
		t.Fatalf("ProjectionLocalizer() error = %v", err)
	}
	want := []math3d.Point2{{X: 0, Y: 5}, {X: 0, Y: 5}, {X: 10, Y: 5}, {X: 10, Y: 5}}
	if len(points) != len(want) {
		t.Fatalf("ProjectionLocalizer() returned %d points", len(points))
	}
	for index := range want {
		if !point2Close(points[index], want[index]) {
			t.Fatalf("point %d = %+v, want %+v", index, points[index], want[index])
		}
	}
}

func mustFrameGeometry(t *testing.T, uid string, position math3d.Point3, row, column math3d.Vector3, columns, rows int) *FrameGeometry {
	t.Helper()
	geometry, err := NewFrameGeometryFromValues(uid, position, row, column, [2]float64{1, 1}, columns, rows)
	if err != nil {
		t.Fatalf("NewFrameGeometryFromValues() error = %v", err)
	}
	return geometry
}
