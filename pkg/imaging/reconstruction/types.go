// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

// StackType identifies one of the three standard patient-space MPR planes.
type StackType int

const (
	// StackTypeAxial generates planes perpendicular to the patient Z axis.
	StackTypeAxial StackType = iota + 1
	// StackTypeCoronal generates planes perpendicular to the patient Y axis.
	StackTypeCoronal
	// StackTypeSagittal generates planes perpendicular to the patient X axis.
	StackTypeSagittal
)

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
