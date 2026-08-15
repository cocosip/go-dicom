// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package geometry

import (
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

// BoundingBox returns the axis-aligned patient-space bounds of the frame's
// four pixel-center corners.
func (g *FrameGeometry) BoundingBox() math3d.Bounds3 {
	box, _ := math3d.BoundingBox(g.corners())
	return box
}

// CanDrawLocalizer reports whether source can be localized on destination.
func CanDrawLocalizer(source, destination *FrameGeometry) bool {
	if source == nil || destination == nil || source.Type != GeometryVolume || destination.Type != GeometryVolume {
		return false
	}
	if source.Orientation == OrientationNone || destination.Orientation == OrientationNone || source.Orientation == destination.Orientation {
		return false
	}
	return source.FrameOfReferenceUID != "" && source.FrameOfReferenceUID == destination.FrameOfReferenceUID
}

// ProjectionLocalizer projects all four source frame corners onto destination
// and returns destination image coordinates in clockwise source-corner order.
func ProjectionLocalizer(source, destination *FrameGeometry) ([]math3d.Point2, error) {
	if source == nil || destination == nil {
		return nil, fmt.Errorf("localizer frames must not be nil")
	}
	if source.Type != GeometryVolume || destination.Type != GeometryVolume {
		return nil, fmt.Errorf("projection localizer requires volume geometry")
	}
	result := make([]math3d.Point2, 0, 4)
	for _, corner := range source.corners() {
		point, err := destination.PatientToImage(corner)
		if err != nil {
			return nil, err
		}
		result = append(result, point)
	}
	return result, nil
}

// IntersectionLocalizer returns the line where the source rectangle crosses
// the destination plane, expressed in destination image coordinates.
func IntersectionLocalizer(source, destination *FrameGeometry) (math3d.Point2, math3d.Point2, bool) {
	if !CanDrawLocalizer(source, destination) {
		return math3d.Point2{}, math3d.Point2{}, false
	}
	plane, err := math3d.NewPlane(destination.DirectionNormal, destination.TopLeft)
	if err != nil {
		return math3d.Point2{}, math3d.Point2{}, false
	}
	corners := source.corners()
	edges := [4]math3d.Segment3{
		{A: corners[0], B: corners[1]},
		{A: corners[1], B: corners[2]},
		{A: corners[2], B: corners[3]},
		{A: corners[3], B: corners[0]},
	}
	intersections := make([]math3d.Point3, 0, 2)
	for _, edge := range edges {
		point, ok := plane.IntersectSegment(edge)
		if ok && !containsPoint(intersections, point) {
			intersections = append(intersections, point)
		}
	}
	if len(intersections) != 2 {
		return math3d.Point2{}, math3d.Point2{}, false
	}
	start, err := destination.PatientToImage(intersections[0])
	if err != nil {
		return math3d.Point2{}, math3d.Point2{}, false
	}
	end, err := destination.PatientToImage(intersections[1])
	if err != nil {
		return math3d.Point2{}, math3d.Point2{}, false
	}
	return start, end, true
}

func (g *FrameGeometry) corners() []math3d.Point3 {
	return []math3d.Point3{g.TopLeft, g.TopRight, g.BottomRight, g.BottomLeft}
}

func containsPoint(points []math3d.Point3, candidate math3d.Point3) bool {
	for _, point := range points {
		if math.Abs(point.X-candidate.X) <= directionTolerance &&
			math.Abs(point.Y-candidate.Y) <= directionTolerance &&
			math.Abs(point.Z-candidate.Z) <= directionTolerance {
			return true
		}
	}
	return false
}
