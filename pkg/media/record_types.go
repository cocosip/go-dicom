// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import "github.com/cocosip/go-dicom/pkg/dicom/tag"

var recordAttributeTags = map[RecordType][]*tag.Tag{
	RecordPatient: {
		tag.PatientID,
		tag.PatientName,
		tag.PatientBirthDate,
		tag.PatientSex,
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
	},
	RecordPresentation: {
		tag.InstanceNumber,
		tag.PresentationCreationDate,
		tag.PresentationCreationTime,
		tag.ReferencedSeriesSequence,
	},
}
