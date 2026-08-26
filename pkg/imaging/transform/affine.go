// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package transform provides composable two-dimensional spatial transforms.
package transform

import (
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const affineTolerance = 1e-12

// Affine2D is a two-dimensional affine matrix:
//
//	x' = a*x + c*y + e
//	y' = b*x + d*y + f
type Affine2D struct {
	a float64
	b float64
	c float64
	d float64
	e float64
	f float64
}

// Identity returns an identity transform.
func Identity() Affine2D {
	return Affine2D{a: 1, d: 1}
}

// Translate returns a translation transform.
func Translate(x, y float64) Affine2D {
	return Affine2D{a: 1, d: 1, e: x, f: y}
}

// Scale returns a scale transform around the origin.
func Scale(x, y float64) Affine2D {
	return Affine2D{a: x, d: y}
}

// Rotate returns a counter-clockwise rotation around the origin, in degrees.
func Rotate(degrees float64) Affine2D {
	radians := degrees * math.Pi / 180
	sine, cosine := clean(math.Sin(radians)), clean(math.Cos(radians))
	return Affine2D{a: cosine, b: sine, c: -sine, d: cosine}
}

// FlipX reflects coordinates across the Y axis.
func FlipX() Affine2D { return Scale(-1, 1) }

// FlipY reflects coordinates across the X axis.
func FlipY() Affine2D { return Scale(1, -1) }

// Then composes transforms in execution order: m.Then(next) applies m first
// and next second.
func (m Affine2D) Then(next Affine2D) Affine2D {
	return Affine2D{
		a: next.a*m.a + next.c*m.b,
		b: next.b*m.a + next.d*m.b,
		c: next.a*m.c + next.c*m.d,
		d: next.b*m.c + next.d*m.d,
		e: next.a*m.e + next.c*m.f + next.e,
		f: next.b*m.e + next.d*m.f + next.f,
	}
}

// Apply transforms point.
func (m Affine2D) Apply(point math3d.Point2) math3d.Point2 {
	return math3d.Point2{
		X: m.a*point.X + m.c*point.Y + m.e,
		Y: m.b*point.X + m.d*point.Y + m.f,
	}
}

// Determinant returns the determinant of the linear portion of the affine
// transform. A transform is invertible exactly when this value is non-zero.
func (m Affine2D) Determinant() float64 {
	return m.a*m.d - m.b*m.c
}

// Inverse returns the affine transform that reverses m.
func (m Affine2D) Inverse() (Affine2D, error) {
	determinant := m.Determinant()
	if math.Abs(determinant) <= affineTolerance || math.IsNaN(determinant) || math.IsInf(determinant, 0) {
		return Affine2D{}, fmt.Errorf("affine transform is singular")
	}
	return Affine2D{
		a: m.d / determinant,
		b: -m.b / determinant,
		c: -m.c / determinant,
		d: m.a / determinant,
		e: (m.c*m.f - m.d*m.e) / determinant,
		f: (m.b*m.e - m.a*m.f) / determinant,
	}, nil
}

// Bounds returns the axis-aligned bounds of a transformed rectangle.
func (m Affine2D) Bounds(rect math3d.Rect) (math3d.Rect, error) {
	if !validRect(rect) {
		return math3d.Rect{}, fmt.Errorf("rectangle dimensions must be non-negative and finite")
	}
	corners := rect.Corners()
	first := m.Apply(corners[0])
	minimumX, maximumX := first.X, first.X
	minimumY, maximumY := first.Y, first.Y
	for _, corner := range corners[1:] {
		point := m.Apply(corner)
		minimumX, maximumX = math.Min(minimumX, point.X), math.Max(maximumX, point.X)
		minimumY, maximumY = math.Min(minimumY, point.Y), math.Max(maximumY, point.Y)
	}
	return math3d.Rect{X: clean(minimumX), Y: clean(minimumY), Width: clean(maximumX - minimumX), Height: clean(maximumY - minimumY)}, nil
}

// BestFit returns an aspect-preserving transform that centers source inside
// destination.
func BestFit(source, destination math3d.Rect) (Affine2D, error) {
	if !validRect(source) || !validRect(destination) || source.Width <= 0 || source.Height <= 0 || destination.Width <= 0 || destination.Height <= 0 {
		return Affine2D{}, fmt.Errorf("best-fit rectangles must have positive finite dimensions")
	}
	scale := math.Min(destination.Width/source.Width, destination.Height/source.Height)
	sourceCenter := math3d.Point2{X: source.X + source.Width/2, Y: source.Y + source.Height/2}
	destinationCenter := math3d.Point2{X: destination.X + destination.Width/2, Y: destination.Y + destination.Height/2}
	return Translate(-sourceCenter.X, -sourceCenter.Y).
		Then(Scale(scale, scale)).
		Then(Translate(destinationCenter.X, destinationCenter.Y)), nil
}

// SpatialTransform stores the user-facing scale, rotation, flip, and pan
// state used by DICOM viewers.
type SpatialTransform struct {
	Scale    float64
	Rotation int
	FlipX    bool
	FlipY    bool
	Pan      math3d.Point2
}

// NewSpatialTransform creates a reset spatial transform.
func NewSpatialTransform() *SpatialTransform {
	transform := &SpatialTransform{}
	transform.Reset()
	return transform
}

// Reset restores identity transform state.
func (s *SpatialTransform) Reset() {
	s.Scale = 1
	s.Rotation = 0
	s.FlipX = false
	s.FlipY = false
	s.Pan = math3d.Point2{}
}

// Rotate adds degrees and normalizes the stored rotation to [0,360).
func (s *SpatialTransform) Rotate(degrees int) {
	s.Rotation = (s.Rotation + degrees) % 360
	if s.Rotation < 0 {
		s.Rotation += 360
	}
}

// IsTransformed reports whether any non-identity state is active.
func (s *SpatialTransform) IsTransformed() bool {
	return s != nil && (math.Abs(s.Scale-1) > affineTolerance || s.Rotation%360 != 0 || s.FlipX || s.FlipY || s.Pan != (math3d.Point2{}))
}

// Affine builds a transform around the center of source, followed by pan.
func (s *SpatialTransform) Affine(source math3d.Rect) (Affine2D, error) {
	if s == nil {
		return Affine2D{}, fmt.Errorf("spatial transform is nil")
	}
	if !validRect(source) || source.Width <= 0 || source.Height <= 0 {
		return Affine2D{}, fmt.Errorf("source rectangle must have positive finite dimensions")
	}
	if s.Scale <= 0 || math.IsNaN(s.Scale) || math.IsInf(s.Scale, 0) {
		return Affine2D{}, fmt.Errorf("spatial scale must be positive and finite")
	}
	center := math3d.Point2{X: source.X + source.Width/2, Y: source.Y + source.Height/2}
	matrix := Translate(-center.X, -center.Y)
	if s.FlipX {
		matrix = matrix.Then(FlipX())
	}
	if s.FlipY {
		matrix = matrix.Then(FlipY())
	}
	matrix = matrix.Then(Scale(s.Scale, s.Scale)).Then(Rotate(float64(s.Rotation)))
	return matrix.Then(Translate(center.X+s.Pan.X, center.Y+s.Pan.Y)), nil
}

func validRect(rect math3d.Rect) bool {
	return rect.Width >= 0 && rect.Height >= 0 &&
		!math.IsNaN(rect.X) && !math.IsNaN(rect.Y) && !math.IsNaN(rect.Width) && !math.IsNaN(rect.Height) &&
		!math.IsInf(rect.X, 0) && !math.IsInf(rect.Y, 0) && !math.IsInf(rect.Width, 0) && !math.IsInf(rect.Height, 0)
}

func clean(value float64) float64 {
	if math.Abs(value) <= affineTolerance {
		return 0
	}
	return value
}
