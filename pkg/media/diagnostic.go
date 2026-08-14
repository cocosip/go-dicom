// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

// Severity describes the impact of a non-fatal DICOMDIR diagnostic.
type Severity string

const (
	// SeverityInfo describes informational recovery details.
	SeverityInfo Severity = "info"
	// SeverityWarning describes a recovered or optional-data problem.
	SeverityWarning Severity = "warning"
)

// DiagnosticCode identifies a stable class of DICOMDIR diagnostic.
type DiagnosticCode string

const (
	// DiagnosticOffsetRepaired indicates that a stale record offset was repaired.
	DiagnosticOffsetRepaired DiagnosticCode = "offset_repaired"
	// DiagnosticOptionalAttributeMissing indicates that an optional record attribute was absent.
	DiagnosticOptionalAttributeMissing DiagnosticCode = "optional_attribute_missing"
	// DiagnosticIconGenerationFailed indicates that optional icon generation failed.
	DiagnosticIconGenerationFailed DiagnosticCode = "icon_generation_failed"
	// DiagnosticUnknownRecordType indicates that an unknown record type was preserved.
	DiagnosticUnknownRecordType DiagnosticCode = "unknown_record_type"
	// DiagnosticUnreachableRecord indicates that a sequence item was not referenced by the recovered tree.
	DiagnosticUnreachableRecord DiagnosticCode = "unreachable_record"
)

// Diagnostic reports a non-fatal DICOMDIR issue without patient or pixel data.
type Diagnostic struct {
	Code           DiagnosticCode
	Severity       Severity
	RecordType     RecordType
	OriginalOffset uint32
	RepairedOffset uint32
	Message        string
}
