// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

// Priority represents DICOM message priority.
type Priority uint16

const (
	// PriorityMedium is the default priority.
	PriorityMedium Priority = 0x0000

	// PriorityHigh indicates high priority.
	PriorityHigh Priority = 0x0001

	// PriorityLow indicates low priority.
	PriorityLow Priority = 0x0002
)

// String returns the string representation of the priority.
func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "LOW"
	case PriorityMedium:
		return "MEDIUM"
	case PriorityHigh:
		return "HIGH"
	default:
		return "UNKNOWN"
	}
}
