// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

// StackType represents the type of reformatted stack
type StackType int

const (
	// StackTypeAxial represents axial (transverse) plane - perpendicular to patient's long axis
	// Standard CT/MR slices looking from feet to head
	StackTypeAxial StackType = 1

	// StackTypeCoronal represents coronal plane - divides body into anterior/posterior
	// Front-to-back slices
	StackTypeCoronal StackType = 2

	// StackTypeSagittal represents sagittal plane - divides body into left/right
	// Side-to-side slices
	StackTypeSagittal StackType = 3
)

// String returns the string representation of the StackType
func (s StackType) String() string {
	switch s {
	case StackTypeAxial:
		return "Axial"
	case StackTypeCoronal:
		return "Coronal"
	case StackTypeSagittal:
		return "Sagittal"
	default:
		return "Unknown"
	}
}

// TODO: The following types are placeholders for future implementation

// ImageData wraps a DICOM dataset for volume reconstruction
//
// Requirements:
//   - Geometry information (Image Position Patient, Image Orientation Patient)
//   - Pixel data access
//   - Frame of Reference UID for grouping related images
//
// Reference: fo-dicom ImageData.cs
type ImageData struct {
	// TODO: Implement
	// dataset      *dicom.Dataset
	// geometry     *FrameGeometry
	// pixels       IPixelData
	// sortingValue float64
}

// VolumeData represents a reconstructed 3D volume from multiple 2D slices
//
// Capabilities:
//   - Sorts and validates input slices
//   - Builds 3D bounding box
//   - Provides arbitrary plane cuts (MPR)
//   - Handles slice spacing variations
//
// Reference: fo-dicom VolumeData.cs
type VolumeData struct {
	// TODO: Implement
	// slices       []*ImageData
	// boundingMin  Point3D
	// boundingMax  Point3D
	// slicesNormal Vector3D
}

// Slice represents a calculated cut through a volume
//
// A slice is defined by:
//   - Top-left corner point
//   - Row direction vector
//   - Column direction vector
//   - Dimensions (rows, columns)
//   - Pixel spacing
//
// The slice extraction uses trilinear interpolation between volume slices.
//
// Reference: fo-dicom Slice.cs
type Slice struct {
	// TODO: Implement
	// topLeft         Point3D
	// rowDirection    Vector3D
	// columnDirection Vector3D
	// rows            int
	// columns         int
	// spacing         float64
	// output          []float64
}

// Stack represents a calculated stack of slices taken from a volume
//
// A stack is a series of parallel slices through the volume,
// typically in one of the three standard planes (Axial, Coronal, Sagittal).
//
// Reference: fo-dicom Stack.cs
type Stack struct {
	// TODO: Implement
	// slices        []*Slice
	// sliceDistance float64
}

// DicomGenerator helps create DICOM datasets from generated stacks
//
// The generator:
//   - Uses common metadata from source images
//   - Generates new Series and SOP Instance UIDs
//   - Updates geometric tags (Image Position, Orientation, etc.)
//   - Embeds rendered pixel data
//
// Reference: fo-dicom DicomGenerator.cs
type DicomGenerator struct {
	// TODO: Implement
	// commonDataset *dicom.Dataset
}

// NewImageData creates a new ImageData from a DICOM dataset
//
// TODO: Implement when required dependencies are available
func NewImageData(/* dataset *dicom.Dataset */) (*ImageData, error) {
	// Placeholder
	return nil, ErrNotImplemented
}

// NewVolumeData constructs a volume from multiple image slices
//
// TODO: Implement volume reconstruction algorithm
func NewVolumeData(images []*ImageData) (*VolumeData, error) {
	// Placeholder
	return nil, ErrNotImplemented
}

// NewStack creates a reformatted stack from a volume
//
// TODO: Implement MPR (Multi-Planar Reformation)
func NewStack(volume *VolumeData, stackType StackType, spacing, sliceDistance float64) (*Stack, error) {
	// Placeholder
	return nil, ErrNotImplemented
}

// NewDicomGenerator creates a DICOM generator
//
// TODO: Implement DICOM dataset generation
func NewDicomGenerator(/* commonDataset *dicom.Dataset */) *DicomGenerator {
	// Placeholder
	return nil
}

// Common errors
var (
	// ErrNotImplemented indicates the feature is not yet implemented
	ErrNotImplemented = &NotImplementedError{}
)

// NotImplementedError represents a not-yet-implemented feature
type NotImplementedError struct{}

func (e *NotImplementedError) Error() string {
	return "reconstruction: feature not yet implemented - requires 3D geometry library, FrameGeometry, and IPixelData implementation"
}
