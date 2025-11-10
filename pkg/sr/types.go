// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

// ValueType represents the type of value in a content item
type ValueType string

const (
	// ValueTypeContainer represents a container content item
	ValueTypeContainer ValueType = "CONTAINER"
	// ValueTypeText represents a text content item
	ValueTypeText ValueType = "TEXT"
	// ValueTypeCode represents a coded content item
	ValueTypeCode ValueType = "CODE"
	// ValueTypeNumeric represents a numeric measurement
	ValueTypeNumeric ValueType = "NUM"
	// ValueTypePersonName represents a person name
	ValueTypePersonName ValueType = "PNAME"
	// ValueTypeDate represents a date value
	ValueTypeDate ValueType = "DATE"
	// ValueTypeTime represents a time value
	ValueTypeTime ValueType = "TIME"
	// ValueTypeDateTime represents a date-time value
	ValueTypeDateTime ValueType = "DATETIME"
	// ValueTypeUIDReference represents a UID reference
	ValueTypeUIDReference ValueType = "UIDREF"
	// ValueTypeComposite represents a reference to a composite SOP instance
	ValueTypeComposite ValueType = "COMPOSITE"
	// ValueTypeImage represents a reference to an image SOP instance
	ValueTypeImage ValueType = "IMAGE"
	// ValueTypeWaveform represents a reference to a waveform SOP instance
	ValueTypeWaveform ValueType = "WAVEFORM"
	// ValueTypeSpatialCoordinate represents spatial coordinates
	ValueTypeSpatialCoordinate ValueType = "SCOORD"
	// ValueTypeTemporalCoordinate represents temporal coordinates
	ValueTypeTemporalCoordinate ValueType = "TCOORD"
)

// String returns the string representation of the ValueType
func (v ValueType) String() string {
	return string(v)
}

// Continuity represents the continuity of content in a container
type Continuity string

const (
	// ContinuityNone indicates no continuity specified
	ContinuityNone Continuity = ""
	// ContinuitySeparate indicates separate content items
	ContinuitySeparate Continuity = "SEPARATE"
	// ContinuityContinuous indicates continuous content
	ContinuityContinuous Continuity = "CONTINUOUS"
)

// String returns the string representation of the Continuity
func (c Continuity) String() string {
	return string(c)
}

// Relationship represents the relationship between content items
type Relationship string

const (
	// RelationshipContains indicates parent-child containment
	RelationshipContains Relationship = "CONTAINS"
	// RelationshipHasProperties describes properties
	RelationshipHasProperties Relationship = "HAS PROPERTIES"
	// RelationshipInferredFrom indicates inference source
	RelationshipInferredFrom Relationship = "INFERRED FROM"
	// RelationshipSelectedFrom indicates selection from options
	RelationshipSelectedFrom Relationship = "SELECTED FROM"
	// RelationshipHasObservationContext indicates observation context
	RelationshipHasObservationContext Relationship = "HAS OBS CONTEXT"
	// RelationshipHasAcquisitionContext indicates acquisition context
	RelationshipHasAcquisitionContext Relationship = "HAS ACQ CONTEXT"
	// RelationshipHasConceptModifier indicates concept modifier
	RelationshipHasConceptModifier Relationship = "HAS CONCEPT MOD"
)

// String returns the string representation of the Relationship
func (r Relationship) String() string {
	return string(r)
}
