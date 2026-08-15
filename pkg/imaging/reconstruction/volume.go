// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

const volumeGeometryTolerance = 1e-4

// VolumeData is a validated, position-sorted set of immutable source frames.
type VolumeData struct {
	images           []*ImageData
	slicePositions   []float64
	normal           math3d.Vector3
	bounds           math3d.Bounds3
	minSliceSpacing  float64
	maxSliceSpacing  float64
	commonData       *dataset.Dataset
	irregularSpacing bool
}

type volumeOptions struct {
	allowIrregular bool
}

// VolumeOption configures volume validation.
type VolumeOption func(*volumeOptions)

// WithIrregularSpacingAllowed opts into interpolation using actual adjacent
// slice distances instead of requiring a regular source series.
func WithIrregularSpacingAllowed() VolumeOption {
	return func(options *volumeOptions) { options.allowIrregular = true }
}

// NewVolumeData validates, copies, sorts, and indexes source-frame references.
func NewVolumeData(images []*ImageData, options ...VolumeOption) (*VolumeData, error) {
	config := volumeOptions{}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}

	filtered := make([]*ImageData, 0, len(images))
	for _, image := range images {
		if image != nil {
			imageCopy := *image
			if image.geometry != nil {
				geometryCopy := *image.geometry
				imageCopy.geometry = &geometryCopy
			}
			filtered = append(filtered, &imageCopy)
		}
	}
	if len(filtered) < 2 {
		return nil, fmt.Errorf("volume requires at least two source frames")
	}
	first := filtered[0]
	if first.geometry == nil || first.pixelData == nil || first.pixelData.Info == nil {
		return nil, fmt.Errorf("source frame geometry and pixel data must be present")
	}
	for index, image := range filtered[1:] {
		if err := validateCompatibleImage(first, image); err != nil {
			return nil, fmt.Errorf("source frame %d: %w", index+1, err)
		}
	}

	normal := first.geometry.DirectionNormal
	for _, image := range filtered {
		position := image.geometry.TopLeft
		image.sortingPosition = dot3(position, normal)
	}
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].sortingPosition < filtered[j].sortingPosition
	})
	positions := make([]float64, len(filtered))
	for index, image := range filtered {
		positions[index] = image.sortingPosition
		if index > 0 && positions[index]-positions[index-1] <= volumeGeometryTolerance {
			return nil, fmt.Errorf("volume requires at least two unique slice positions")
		}
	}

	minimumSpacing := math.Inf(1)
	maximumSpacing := 0.0
	for index := 1; index < len(positions); index++ {
		spacing := positions[index] - positions[index-1]
		minimumSpacing = math.Min(minimumSpacing, spacing)
		maximumSpacing = math.Max(maximumSpacing, spacing)
	}
	irregular := maximumSpacing-minimumSpacing > math.Max(1e-3, minimumSpacing*1e-3)
	if irregular && !config.allowIrregular {
		return nil, fmt.Errorf("irregular slice spacing %.6g..%.6g requires explicit opt-in", minimumSpacing, maximumSpacing)
	}

	bounds := filtered[0].geometry.BoundingBox()
	for _, image := range filtered[1:] {
		bounds = unionBounds(bounds, image.geometry.BoundingBox())
	}
	return &VolumeData{
		images:           filtered,
		slicePositions:   positions,
		normal:           normal,
		bounds:           bounds,
		minSliceSpacing:  minimumSpacing,
		maxSliceSpacing:  maximumSpacing,
		commonData:       commonDataset(filtered),
		irregularSpacing: irregular,
	}, nil
}

// Len returns the number of source frames in the volume.
func (volume *VolumeData) Len() int {
	if volume == nil {
		return 0
	}
	return len(volume.images)
}

// SlicePositions returns a copy of the sorted projected frame positions.
func (volume *VolumeData) SlicePositions() []float64 {
	if volume == nil {
		return nil
	}
	return append([]float64(nil), volume.slicePositions...)
}

// Normal returns the common source-frame normal.
func (volume *VolumeData) Normal() math3d.Vector3 {
	if volume == nil {
		return math3d.Vector3{}
	}
	return volume.normal
}

// Bounds returns the patient-space volume bounds.
func (volume *VolumeData) Bounds() math3d.Bounds3 {
	if volume == nil {
		return math3d.Bounds3{}
	}
	return volume.bounds
}

// MinSliceSpacing returns the smallest adjacent source-frame spacing.
func (volume *VolumeData) MinSliceSpacing() float64 {
	if volume == nil {
		return 0
	}
	return volume.minSliceSpacing
}

// MaxSliceSpacing returns the largest adjacent source-frame spacing.
func (volume *VolumeData) MaxSliceSpacing() float64 {
	if volume == nil {
		return 0
	}
	return volume.maxSliceSpacing
}

// IrregularSpacing reports whether adjacent source-frame spacing varies.
func (volume *VolumeData) IrregularSpacing() bool {
	return volume != nil && volume.irregularSpacing
}

// CommonDataset returns an independent Dataset container containing metadata
// shared by every source frame.
func (volume *VolumeData) CommonDataset() *dataset.Dataset {
	if volume == nil || volume.commonData == nil {
		return dataset.New()
	}
	return volume.commonData.DeepClone()
}

func validateCompatibleImage(first, other *ImageData) error {
	if other == nil || other.geometry == nil || other.pixelData == nil || other.pixelData.Info == nil {
		return fmt.Errorf("source frame geometry and pixel data must be present")
	}
	a, b := first.geometry, other.geometry
	if a.FrameOfReferenceUID == "" || a.FrameOfReferenceUID != b.FrameOfReferenceUID {
		return fmt.Errorf("source Frame of Reference UID does not match")
	}
	if a.Columns != b.Columns || a.Rows != b.Rows {
		return fmt.Errorf("frame dimensions do not match")
	}
	if !floatsClose(a.PixelSpacingColumns, b.PixelSpacingColumns) || !floatsClose(a.PixelSpacingRows, b.PixelSpacingRows) ||
		!vectorClose(a.DirectionRow, b.DirectionRow) || !vectorClose(a.DirectionColumn, b.DirectionColumn) ||
		!vectorClose(a.DirectionNormal, b.DirectionNormal) {
		return fmt.Errorf("frame orientation or pixel spacing does not match")
	}
	pa, pb := first.pixelData.Info, other.pixelData.Info
	if pa.BitsAllocated != pb.BitsAllocated || pa.BitsStored != pb.BitsStored || pa.HighBit != pb.HighBit ||
		pa.SamplesPerPixel != pb.SamplesPerPixel || pa.PixelRepresentation != pb.PixelRepresentation ||
		pa.PhotometricInterpretation == nil || pb.PhotometricInterpretation == nil ||
		pa.PhotometricInterpretation.Value != pb.PhotometricInterpretation.Value {
		return fmt.Errorf("pixel semantics do not match")
	}
	if sourceModality(first.sourceSOPClassUID) != sourceModality(other.sourceSOPClassUID) {
		return fmt.Errorf("source modality does not match")
	}
	return nil
}

func sourceModality(sopClass string) string {
	switch sopClass {
	case uid.CTImageStorage.UID(), uid.EnhancedCTImageStorage.UID():
		return "CT"
	case uid.MRImageStorage.UID(), uid.EnhancedMRImageStorage.UID():
		return "MR"
	default:
		return ""
	}
}

func commonDataset(images []*ImageData) *dataset.Dataset {
	common := dataset.New()
	common.SetInternalTransferSyntax(images[0].dataset.InternalTransferSyntax())
	for _, candidate := range images[0].dataset.Elements() {
		if candidate.Tag() == nil || candidate.Tag().ToUint32() == tag.PixelData.ToUint32() {
			continue
		}
		if _, sequence := candidate.(*dataset.Sequence); sequence {
			continue
		}
		matches := true
		for _, image := range images[1:] {
			other, ok := image.dataset.Get(candidate.Tag())
			if !ok || !elementsEqual(candidate, other) {
				matches = false
				break
			}
		}
		if matches {
			_ = common.AddOrUpdate(candidate)
		}
	}
	return common
}

func elementsEqual(left, right element.Element) bool {
	if left == nil || right == nil || left.ValueRepresentation().Code() != right.ValueRepresentation().Code() {
		return false
	}
	return bytes.Equal(left.Buffer().Data(), right.Buffer().Data())
}

func unionBounds(a, b math3d.Bounds3) math3d.Bounds3 {
	return math3d.Bounds3{
		Min: math3d.Point3{X: math.Min(a.Min.X, b.Min.X), Y: math.Min(a.Min.Y, b.Min.Y), Z: math.Min(a.Min.Z, b.Min.Z)},
		Max: math3d.Point3{X: math.Max(a.Max.X, b.Max.X), Y: math.Max(a.Max.Y, b.Max.Y), Z: math.Max(a.Max.Z, b.Max.Z)},
	}
}

func dot3(point math3d.Point3, vector math3d.Vector3) float64 {
	return point.X*vector.X + point.Y*vector.Y + point.Z*vector.Z
}

func vectorClose(a, b math3d.Vector3) bool {
	return floatsClose(a.X, b.X) && floatsClose(a.Y, b.Y) && floatsClose(a.Z, b.Z)
}

func floatsClose(a, b float64) bool {
	return math.Abs(a-b) <= volumeGeometryTolerance
}
