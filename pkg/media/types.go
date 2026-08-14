// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

// RecordType identifies a DICOM directory record type.
type RecordType string

const (
	// RecordPatient identifies a PATIENT directory record.
	RecordPatient RecordType = "PATIENT"
	// RecordStudy identifies a STUDY directory record.
	RecordStudy RecordType = "STUDY"
	// RecordSeries identifies a SERIES directory record.
	RecordSeries RecordType = "SERIES"
	// RecordImage identifies an IMAGE directory record.
	RecordImage RecordType = "IMAGE"
	// RecordSRDocument identifies an SR DOCUMENT directory record.
	RecordSRDocument RecordType = "SR DOCUMENT"
	// RecordPresentation identifies a PRESENTATION directory record.
	RecordPresentation RecordType = "PRESENTATION"
)
