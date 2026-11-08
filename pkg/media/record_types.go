// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// recordAttributeTags lists the attributes copied from the source instance into
// each directory record. Attributes absent from the source are reported; the
// Type 1 subset in recordRequiredTags is a hard error instead.
var recordAttributeTags = map[RecordType][]*tag.Tag{
	RecordPatient: {
		tag.PatientID,
		tag.PatientName,
		tag.PatientBirthDate,
		tag.PatientSex,
	},
	RecordKeyObjectDoc: {
		tag.InstanceNumber,
		tag.ContentDate,
		tag.ContentTime,
		tag.ConceptNameCodeSequence,
		tag.ContentSequence,
	},
	RecordStudy: {
		tag.StudyInstanceUID,
		tag.StudyID,
		tag.StudyDate,
		tag.StudyTime,
		tag.AccessionNumber,
		tag.StudyDescription,
	},
	RecordSeries: {
		tag.SeriesInstanceUID,
		tag.Modality,
		tag.SeriesDate,
		tag.SeriesTime,
		tag.SeriesNumber,
		tag.SeriesDescription,
	},
	RecordImage: {
		tag.InstanceNumber,
	},
	RecordSRDocument: {
		tag.InstanceNumber,
		tag.CompletionFlag,
		tag.VerificationFlag,
		tag.ContentDate,
		tag.ContentTime,
		tag.VerificationDateTime,
		tag.ConceptNameCodeSequence,
		tag.ContentSequence,
	},
	RecordPresentation: {
		tag.InstanceNumber,
		tag.ContentLabel,
		tag.ContentDescription,
		tag.PresentationCreationDate,
		tag.PresentationCreationTime,
		tag.ReferencedSeriesSequence,
	},
}

// recordRequiredTags lists the directory record keys that must be present
// (PS3.3 F.5.1-F.5.26). A record missing one of these cannot be matched during
// File-set traversal, so AddFile reports an error instead of silently dropping
// the value. Other Type 1/2 attributes are reported as diagnostics; they are
// frequently absent from real instances and are not part of the record key.
var recordRequiredTags = map[RecordType][]*tag.Tag{
	RecordPatient: {
		tag.PatientID,
	},
	RecordStudy: {
		tag.StudyDate,
		tag.StudyTime,
		tag.StudyID,
		tag.StudyInstanceUID,
	},
	RecordSeries: {
		tag.Modality,
		tag.SeriesInstanceUID,
		tag.SeriesNumber,
	},
	RecordImage: {
		tag.InstanceNumber,
	},
	RecordSRDocument: {
		tag.InstanceNumber,
		tag.CompletionFlag,
		tag.VerificationFlag,
		tag.ContentDate,
		tag.ContentTime,
		tag.ConceptNameCodeSequence,
	},
	RecordPresentation: {
		tag.InstanceNumber,
		tag.ContentLabel,
		tag.PresentationCreationDate,
		tag.PresentationCreationTime,
	},
	RecordKeyObjectDoc: {
		tag.InstanceNumber,
		tag.ContentDate,
		tag.ContentTime,
		tag.ConceptNameCodeSequence,
	},
}

// type2Attribute is a Type 2 directory record key written as a zero-length
// element when the source instance does not provide it (PS3.3 F.5).
type type2Attribute struct {
	tag *tag.Tag
	vr  *vr.VR
}

// recordType2Tags lists the Type 2 keys of each record type. They are written
// with zero length when absent so the record remains conformant.
var recordType2Tags = map[RecordType][]type2Attribute{
	RecordPatient: {
		{tag.PatientName, vr.PN},
	},
	RecordStudy: {
		{tag.StudyDescription, vr.LO},
		{tag.AccessionNumber, vr.SH},
	},
	RecordPresentation: {
		{tag.ContentDescription, vr.LO},
	},
}
