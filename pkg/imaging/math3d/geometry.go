// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package math3d provides the geometry primitives shared by DICOM frame
// geometry, spatial transforms, and volume reconstruction.
package math3d

import (
	"fmt"
	"math"
)

const epsilon = 1e-12

// Point2 is a point in a two-dimensional coordinate system.
type Point2 struct {
	X float64
	Y float64
}

// Rect is an axis-aligned rectangle.
type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Corners returns the rectangle corners in clockwise order, starting at the
// top-left corner.
func (r Rect) Corners() [4]Point2 {
	return [4]Point2{
		{X: r.X, Y: r.Y},
		{X: r.X + r.Width, Y: r.Y},
		{X: r.X + r.Width, Y: r.Y + r.Height},
		{X: r.X, Y: r.Y + r.Height},
	}
}

// Point3 is a point in the DICOM patient coordinate system.
type Point3 struct {
	X float64
	Y float64
	Z float64
}

// Vector3 is a three-dimensional vector.
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Add returns v + other.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Sub returns v - other.
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// Scale returns v multiplied by factor.
func (v Vector3) Scale(factor float64) Vector3 {
	return Vector3{X: v.X * factor, Y: v.Y * factor, Z: v.Z * factor}
}

// Dot returns the dot product of v and other.
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Distance returns the distance from v to point.
func (v Vector3) Distance(point Point3) float64 {
	return math.Sqrt((v.X-point.X)*(v.X-point.X) + (v.Y-point.Y)*(v.Y-point.Y) + (v.Z-point.Z)*(v.Z-point.Z))
}

// Distance returns the distance between p and other.
func (p Point3) Distance(other Point3) float64 {
	return p.Sub(other).Length()
}

// Add returns p translated by v.
func (p Point3) Add(v Vector3) Point3 {
	return Point3{X: p.X + v.X, Y: p.Y + v.Y, Z: p.Z + v.Z}
}

// Sub returns the vector pointing from other to p.
func (p Point3) Sub(other Point3) Vector3 {
	return Vector3{X: p.X - other.X, Y: p.Y - other.Y, Z: p.Z - other.Z}
}

// Dot returns the dot product of the position vector of p and v.
func (p Point3) Dot(v Vector3) float64 {
	return p.X*v.X + p.Y*v.Y + p.Z*v.Z
}

// Cross returns the cross product of v and other.
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

// Length returns the Euclidean length of v.
func (v Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a unit vector in the direction of v.
func (v Vector3) Normalize() (Vector3, error) {
	if !finite(v.X) || !finite(v.Y) || !finite(v.Z) {
		return Vector3{}, fmt.Errorf("cannot normalize a vector with non-finite components")
	}
	length := v.Length()
	if length <= epsilon {
		return Vector3{}, fmt.Errorf("cannot normalize a zero vector")
	}
	return Vector3{X: v.X / length, Y: v.Y / length, Z: v.Z / length}, nil
}

func finite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

// NearestAxis returns the signed cardinal axis nearest to v.
func (v Vector3) NearestAxis() Vector3 {
	ax, ay, az := math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)
	switch {
	case ax >= ay && ax >= az:
		return Vector3{X: math.Copysign(1, v.X)}
	case ay >= az:
		return Vector3{Y: math.Copysign(1, v.Y)}
	default:
		return Vector3{Z: math.Copysign(1, v.Z)}
	}
}

// Segment3 is a finite line segment between A and B.
type Segment3 struct {
	A Point3
	B Point3
}

// Plane is a plane represented by a unit normal and a point on the plane.
type Plane struct {
	Normal Vector3
	Point  Point3
}

// NewPlane constructs a plane from a normal and a point on the plane.
func NewPlane(normal Vector3, point Point3) (Plane, error) {
	unit, err := normal.Normalize()
	if err != nil {
		return Plane{}, fmt.Errorf("plane normal: %w", err)
	}
	return Plane{Normal: unit, Point: point}, nil
}

// IntersectSegment returns the intersection of the plane and segment. A
// segment lying entirely in the plane is treated as ambiguous and returns no
// single intersection.
func (p Plane) IntersectSegment(segment Segment3) (Point3, bool) {
	da := p.signedDistance(segment.A)
	db := p.signedDistance(segment.B)
	if math.Abs(da) <= epsilon && math.Abs(db) <= epsilon {
		return Point3{}, false
	}
	denominator := da - db
	if math.Abs(denominator) <= epsilon {
		return Point3{}, false
	}
	t := da / denominator
	if t < -epsilon || t > 1+epsilon {
		return Point3{}, false
	}
	return Point3{
		X: segment.A.X + (segment.B.X-segment.A.X)*t,
		Y: segment.A.Y + (segment.B.Y-segment.A.Y)*t,
		Z: segment.A.Z + (segment.B.Z-segment.A.Z)*t,
	}, true
}

func (p Plane) signedDistance(point Point3) float64 {
	return p.Normal.X*(point.X-p.Point.X) +
		p.Normal.Y*(point.Y-p.Point.Y) +
		p.Normal.Z*(point.Z-p.Point.Z)
}

// Bounds3 is an axis-aligned three-dimensional bounding box.
type Bounds3 struct {
	Min Point3
	Max Point3
}

// Union returns the smallest box containing both b and other.
func (b Bounds3) Union(other Bounds3) Bounds3 {
	return Bounds3{
		Min: Point3{X: math.Min(b.Min.X, other.Min.X), Y: math.Min(b.Min.Y, other.Min.Y), Z: math.Min(b.Min.Z, other.Min.Z)},
		Max: Point3{X: math.Max(b.Max.X, other.Max.X), Y: math.Max(b.Max.Y, other.Max.Y), Z: math.Max(b.Max.Z, other.Max.Z)},
	}
}

// BoundingBox returns the smallest axis-aligned box containing points.
func BoundingBox(points []Point3) (Bounds3, error) {
	if len(points) == 0 {
		return Bounds3{}, fmt.Errorf("cannot calculate a bounding box without points")
	}
	minimum, maximum := points[0], points[0]
	for _, point := range points[1:] {
		minimum.X = math.Min(minimum.X, point.X)
		minimum.Y = math.Min(minimum.Y, point.Y)
		minimum.Z = math.Min(minimum.Z, point.Z)
		maximum.X = math.Max(maximum.X, point.X)
		maximum.Y = math.Max(maximum.Y, point.Y)
		maximum.Z = math.Max(maximum.Z, point.Z)
	}
	return Bounds3{Min: minimum, Max: maximum}, nil
}
