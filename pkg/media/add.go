// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Entry identifies the patient, study, series, and instance records used by AddFile.
type Entry struct {
	Patient  *Record
	Study    *Record
	Series   *Record
	Instance *Record
}

// AddFile adds a parsed DICOM file to the directory without modifying the source file.
func (d *Directory) AddFile(file *parser.ParseResult, fileID FileID) (*Entry, error) {
	if d == nil {
		return nil, fmt.Errorf("directory cannot be nil")
	}
	if file == nil || file.Dataset == nil || file.FileMetaInformation == nil {
		return nil, fmt.Errorf("parsed DICOM file and metadata cannot be nil")
	}
	if err := fileID.validate(); err != nil {
		return nil, fmt.Errorf("invalid Referenced File ID: %w", err)
	}

	studyUID, err := requiredString(file.Dataset, tag.StudyInstanceUID)
	if err != nil {
		return nil, err
	}
	seriesUID, err := requiredString(file.Dataset, tag.SeriesInstanceUID)
	if err != nil {
		return nil, err
	}
	instanceUID, err := requiredString(file.Dataset, tag.SOPInstanceUID)
	if err != nil {
		return nil, err
	}

	fmi := file.FileMetaInformation
	sopClassUID, ok := fmi.MediaStorageSOPClassUID()
	if !ok || sopClassUID == "" {
		return nil, fmt.Errorf("missing MediaStorageSOPClassUID")
	}
	referencedInstanceUID, ok := fmi.MediaStorageSOPInstanceUID()
	if !ok || referencedInstanceUID == "" {
		return nil, fmt.Errorf("missing MediaStorageSOPInstanceUID")
	}
	transferSyntaxUID, ok := fmi.TransferSyntaxUID()
	if !ok || transferSyntaxUID == "" {
		return nil, fmt.Errorf("missing TransferSyntaxUID")
	}

	patientID, _ := stringValue(file.Dataset, tag.PatientID)
	patientName, _ := stringValue(file.Dataset, tag.PatientName)
	patientKey := patientID + "\x00" + normalizePersonName(patientName)

	patient, err := d.findOrCreate(nil, d.roots, RecordPatient, patientKey, file.Dataset)
	if err != nil {
		return nil, err
	}
	if !containsRecord(d.roots, patient) {
		d.roots = append(d.roots, patient)
	}

	study, err := d.findOrCreate(patient, patient.children, RecordStudy, studyUID, file.Dataset)
	if err != nil {
		return nil, err
	}
	series, err := d.findOrCreate(study, study.children, RecordSeries, seriesUID, file.Dataset)
	if err != nil {
		return nil, err
	}

	for _, current := range series.children {
		if current.key == instanceUID {
			return &Entry{Patient: patient, Study: study, Series: series, Instance: current}, nil
		}
	}

	instanceType := instanceRecordType(sopClassUID)
	instance, err := d.newRecord(instanceType, instanceUID, file.Dataset)
	if err != nil {
		return nil, err
	}
	for _, item := range []element.Element{
		element.NewString(tag.ReferencedFileID, vr.CS, fileID.Components()),
		element.NewString(tag.ReferencedSOPClassUIDInFile, vr.UI, []string{sopClassUID}),
		element.NewString(tag.ReferencedSOPInstanceUIDInFile, vr.UI, []string{referencedInstanceUID}),
		element.NewString(tag.ReferencedTransferSyntaxUIDInFile, vr.UI, []string{transferSyntaxUID}),
	} {
		if err := instance.dataset.AddOrUpdate(item); err != nil {
			return nil, fmt.Errorf("add instance reference %s: %w", item.Tag(), err)
		}
	}
	if err := d.addInstanceIcon(file.Dataset, instance); err != nil {
		return nil, err
	}
	series.addChild(instance)

	return &Entry{Patient: patient, Study: study, Series: series, Instance: instance}, nil
}

func (d *Directory) addInstanceIcon(source *dataset.Dataset, instance *Record) error {
	if !d.imageIcons {
		return nil
	}
	if d.iconGenerator == nil {
		d.addIconFailureDiagnostic(instance.recordType)
		return nil
	}
	frame := representativeIconFrame(source)
	width, height, pixels, err := d.iconGenerator.GenerateDirectoryIcon(source, frame)
	if err != nil {
		d.addIconFailureDiagnostic(instance.recordType)
		return nil
	}
	sequence, err := NewIconImageSequence(width, height, pixels)
	if err != nil {
		d.addIconFailureDiagnostic(instance.recordType)
		return nil
	}
	if err := instance.dataset.AddOrUpdate(sequence); err != nil {
		return fmt.Errorf("attach Icon Image Sequence: %w", err)
	}
	return nil
}

func (d *Directory) addIconFailureDiagnostic(recordType RecordType) {
	d.diagnostics = append(d.diagnostics, Diagnostic{
		Code:       DiagnosticIconGenerationFailed,
		Severity:   SeverityWarning,
		RecordType: recordType,
		Message:    "directory icon generation failed; record was added without an icon",
	})
}

func representativeIconFrame(ds *dataset.Dataset) int {
	frameCount := 1
	if value, ok := ds.GetString(tag.NumberOfFrames); ok {
		if parsed, err := strconv.Atoi(strings.TrimSpace(value)); err == nil && parsed > 0 {
			frameCount = parsed
		}
	}
	if representative, err := ds.GetUInt16(tag.RepresentativeFrameNumber, 0); err == nil {
		frame := int(representative) - 1
		if frame >= 0 && frame < frameCount {
			return frame
		}
	}
	frame := frameCount / 3
	if frame >= frameCount {
		return frameCount - 1
	}
	return frame
}

func (d *Directory) findOrCreate(parent *Record, records []*Record, recordType RecordType, key string, source *dataset.Dataset) (*Record, error) {
	for _, record := range records {
		if record.recordType == recordType && record.key == key {
			return record, nil
		}
	}
	record, err := d.newRecord(recordType, key, source)
	if err != nil {
		return nil, err
	}
	if parent != nil {
		parent.addChild(record)
	}
	return record, nil
}

func (d *Directory) newRecord(recordType RecordType, key string, source *dataset.Dataset) (*Record, error) {
	recordDataset := dataset.NewWithTransferSyntax(d.transferSyntax)
	for _, item := range []element.Element{
		element.NewUnsignedLong(tag.OffsetOfTheNextDirectoryRecord, []uint32{0}),
		element.NewUnsignedShort(tag.RecordInUseFlag, []uint16{0xFFFF}),
		element.NewUnsignedLong(tag.OffsetOfReferencedLowerLevelDirectoryEntity, []uint32{0}),
		element.NewString(tag.DirectoryRecordType, vr.CS, []string{string(recordType)}),
	} {
		if err := recordDataset.Add(item); err != nil {
			return nil, fmt.Errorf("add base %s record element %s: %w", recordType, item.Tag(), err)
		}
	}
	if specificCharacterSet, ok := source.Get(tag.SpecificCharacterSet); ok {
		if err := recordDataset.Add(specificCharacterSet); err != nil {
			return nil, fmt.Errorf("copy Specific Character Set: %w", err)
		}
	}

	for _, attributeTag := range recordAttributeTags[recordType] {
		item, ok := source.Get(attributeTag)
		if !ok {
			d.diagnostics = append(d.diagnostics, Diagnostic{
				Code:       DiagnosticOptionalAttributeMissing,
				Severity:   SeverityInfo,
				RecordType: recordType,
				Message:    fmt.Sprintf("optional attribute %s is missing", attributeTag),
			})
			continue
		}
		if err := recordDataset.Add(item); err != nil {
			return nil, fmt.Errorf("copy %s record attribute %s: %w", recordType, attributeTag, err)
		}
	}

	record := newRecord(recordType, recordDataset, 0)
	record.key = key
	return record, nil
}

func requiredString(ds *dataset.Dataset, valueTag *tag.Tag) (string, error) {
	value, ok := stringValue(ds, valueTag)
	if !ok || value == "" {
		return "", fmt.Errorf("missing required %s", valueTag)
	}
	return value, nil
}

func stringValue(ds *dataset.Dataset, valueTag *tag.Tag) (string, bool) {
	if value, ok := ds.GetString(valueTag); ok {
		return value, true
	}
	item, ok := ds.Get(valueTag)
	if !ok {
		return "", false
	}
	personName, ok := item.(*element.PersonName)
	if !ok || personName.Count() == 0 {
		return "", false
	}
	return personName.GetValue(0), true
}

func normalizePersonName(value string) string {
	representations := strings.Split(value, "=")
	for i, representation := range representations {
		components := strings.Split(representation, "^")
		for len(components) > 0 && components[len(components)-1] == "" {
			components = components[:len(components)-1]
		}
		representations[i] = strings.Join(components, "^")
	}
	for len(representations) > 0 && representations[len(representations)-1] == "" {
		representations = representations[:len(representations)-1]
	}
	return strings.Join(representations, "=")
}

func instanceRecordType(sopClassUID string) RecordType {
	class := uid.Parse(sopClassUID, "", uid.TypeSOPClass)
	switch class.StorageCategory() {
	case uid.StorageCategoryStructuredReport:
		return RecordSRDocument
	case uid.StorageCategoryPresentationState:
		return RecordPresentation
	default:
		return RecordImage
	}
}

func containsRecord(records []*Record, target *Record) bool {
	for _, record := range records {
		if record == target {
			return true
		}
	}
	return false
}
