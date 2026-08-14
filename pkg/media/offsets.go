// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"fmt"
	"math"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type offsetReference struct {
	dataset    *dataset.Dataset
	tag        *tag.Tag
	recordType RecordType
	original   uint32
}

func buildRecordTree(directoryDataset *dataset.Dataset, records []*Record, policy OffsetPolicy) ([]*Record, []Diagnostic, error) {
	references, err := collectOffsetReferences(directoryDataset, records)
	if err != nil {
		return nil, nil, err
	}
	roots, exactErr := buildExactRecordTree(directoryDataset, records)
	if exactErr == nil {
		return roots, nil, nil
	}
	if policy == StrictOffsets || allReferencesResolve(references, records) {
		return nil, nil, exactErr
	}

	stored := make([]uint32, len(references))
	actual := make([]uint32, len(records))
	for i, reference := range references {
		stored[i] = reference.original
	}
	for i, record := range records {
		actual[i] = record.offset
	}
	delta, deltaErr := recoverOffsetDelta(stored, actual)
	if deltaErr == nil && delta != 0 {
		return repairRecordTreeByDelta(directoryDataset, records, references, delta)
	}

	return repairRecordTreeByPhysicalHierarchy(directoryDataset, records, references)
}

func collectOffsetReferences(directoryDataset *dataset.Dataset, records []*Record) ([]offsetReference, error) {
	references := make([]offsetReference, 0, 2+len(records)*2)
	for _, offsetTag := range []*tag.Tag{
		tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity,
		tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity,
	} {
		value, err := requiredOffset(directoryDataset, offsetTag)
		if err != nil {
			return nil, err
		}
		references = append(references, offsetReference{dataset: directoryDataset, tag: offsetTag, original: value})
	}
	for _, record := range records {
		for _, offsetTag := range []*tag.Tag{
			tag.OffsetOfTheNextDirectoryRecord,
			tag.OffsetOfReferencedLowerLevelDirectoryEntity,
		} {
			value, err := requiredOffset(record.dataset, offsetTag)
			if err != nil {
				return nil, fmt.Errorf("record at offset %d: %w", record.offset, err)
			}
			references = append(references, offsetReference{
				dataset:    record.dataset,
				tag:        offsetTag,
				recordType: record.recordType,
				original:   value,
			})
		}
	}
	return references, nil
}

func allReferencesResolve(references []offsetReference, records []*Record) bool {
	actual := make(map[uint32]struct{}, len(records))
	for _, record := range records {
		actual[record.offset] = struct{}{}
	}
	for _, reference := range references {
		if reference.original == 0 {
			continue
		}
		if _, ok := actual[reference.original]; !ok {
			return false
		}
	}
	return true
}

func repairRecordTreeByDelta(
	directoryDataset *dataset.Dataset,
	records []*Record,
	references []offsetReference,
	delta int64,
) ([]*Record, []Diagnostic, error) {
	actual := make(map[uint32]struct{}, len(records))
	for _, record := range records {
		actual[record.offset] = struct{}{}
	}
	repaired := make([]uint32, len(references))
	for i, reference := range references {
		if reference.original == 0 {
			continue
		}
		value := int64(reference.original) - delta
		if value <= 0 || value > math.MaxUint32 {
			return nil, nil, fmt.Errorf("repaired directory offset is outside UL range")
		}
		repaired[i] = uint32(value)
		if _, ok := actual[repaired[i]]; !ok {
			return nil, nil, fmt.Errorf("repaired directory offset %d does not identify a sequence item", repaired[i])
		}
	}
	if err := applyReferenceValues(references, repaired); err != nil {
		return nil, nil, err
	}
	roots, err := buildExactRecordTree(directoryDataset, records)
	if err != nil {
		_ = restoreReferenceValues(references)
		return nil, nil, fmt.Errorf("fixed-offset recovery produced an invalid directory tree: %w", err)
	}
	return roots, repairDiagnostics(references, repaired), nil
}

func repairRecordTreeByPhysicalHierarchy(
	directoryDataset *dataset.Dataset,
	records []*Record,
	references []offsetReference,
) ([]*Record, []Diagnostic, error) {
	roots, err := inferPhysicalHierarchy(records)
	if err != nil {
		return nil, nil, fmt.Errorf("DICOMDIR offsets cannot be recovered uniquely: %w", err)
	}
	if err := writeTreeOffsets(directoryDataset, records, roots); err != nil {
		_ = restoreReferenceValues(references)
		return nil, nil, err
	}
	repaired := make([]uint32, len(references))
	for i, reference := range references {
		value, err := requiredOffset(reference.dataset, reference.tag)
		if err != nil {
			_ = restoreReferenceValues(references)
			return nil, nil, err
		}
		repaired[i] = value
	}
	validatedRoots, err := buildExactRecordTree(directoryDataset, records)
	if err != nil {
		_ = restoreReferenceValues(references)
		return nil, nil, fmt.Errorf("physical-order recovery produced an invalid directory tree: %w", err)
	}
	return validatedRoots, repairDiagnostics(references, repaired), nil
}

func inferPhysicalHierarchy(records []*Record) ([]*Record, error) {
	for _, record := range records {
		record.children = nil
	}
	var roots []*Record
	var patient, study, series *Record
	for _, record := range records {
		switch record.recordType {
		case RecordPatient:
			roots = append(roots, record)
			patient, study, series = record, nil, nil
		case RecordStudy:
			if patient == nil {
				return nil, fmt.Errorf("STUDY record precedes a PATIENT record")
			}
			patient.addChild(record)
			study, series = record, nil
		case RecordSeries:
			if study == nil {
				return nil, fmt.Errorf("SERIES record has no preceding STUDY record")
			}
			study.addChild(record)
			series = record
		case RecordImage, RecordSRDocument, RecordPresentation:
			if series == nil {
				return nil, fmt.Errorf("%s record has no preceding SERIES record", record.recordType)
			}
			series.addChild(record)
		default:
			return nil, fmt.Errorf("record type %q has no unique supported hierarchy", record.recordType)
		}
	}
	if len(roots) == 0 {
		return nil, fmt.Errorf("no PATIENT root record found")
	}
	return roots, nil
}

func writeTreeOffsets(directoryDataset *dataset.Dataset, records, roots []*Record) error {
	for _, record := range records {
		if err := setOffsetValue(record.dataset, tag.OffsetOfTheNextDirectoryRecord, 0); err != nil {
			return err
		}
		if err := setOffsetValue(record.dataset, tag.OffsetOfReferencedLowerLevelDirectoryEntity, 0); err != nil {
			return err
		}
	}
	if err := setOffsetValue(directoryDataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, roots[0].offset); err != nil {
		return err
	}
	if err := setOffsetValue(directoryDataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, roots[len(roots)-1].offset); err != nil {
		return err
	}
	var writeLevel func([]*Record) error
	writeLevel = func(level []*Record) error {
		for i, record := range level {
			if i+1 < len(level) {
				if err := setOffsetValue(record.dataset, tag.OffsetOfTheNextDirectoryRecord, level[i+1].offset); err != nil {
					return err
				}
			}
			if len(record.children) != 0 {
				if err := setOffsetValue(record.dataset, tag.OffsetOfReferencedLowerLevelDirectoryEntity, record.children[0].offset); err != nil {
					return err
				}
				if err := writeLevel(record.children); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return writeLevel(roots)
}

func applyReferenceValues(references []offsetReference, values []uint32) error {
	if len(references) != len(values) {
		return fmt.Errorf("offset reference/value count mismatch")
	}
	for i, reference := range references {
		if err := setOffsetValue(reference.dataset, reference.tag, values[i]); err != nil {
			return err
		}
	}
	return nil
}

func restoreReferenceValues(references []offsetReference) error {
	values := make([]uint32, len(references))
	for i, reference := range references {
		values[i] = reference.original
	}
	return applyReferenceValues(references, values)
}

func setOffsetValue(ds *dataset.Dataset, offsetTag *tag.Tag, value uint32) error {
	if err := ds.AddOrUpdate(element.NewUnsignedLong(offsetTag, []uint32{value})); err != nil {
		return fmt.Errorf("set directory offset %s: %w", offsetTag, err)
	}
	return nil
}

func repairDiagnostics(references []offsetReference, repaired []uint32) []Diagnostic {
	diagnostics := make([]Diagnostic, 0)
	for i, reference := range references {
		if reference.original == repaired[i] {
			continue
		}
		diagnostics = append(diagnostics, Diagnostic{
			Code:           DiagnosticOffsetRepaired,
			Severity:       SeverityWarning,
			RecordType:     reference.recordType,
			OriginalOffset: reference.original,
			RepairedOffset: repaired[i],
			Message:        fmt.Sprintf("repaired stale directory reference %s", reference.tag),
		})
	}
	return diagnostics
}

func buildExactRecordTree(directoryDataset *dataset.Dataset, records []*Record) ([]*Record, error) {
	first, err := requiredOffset(directoryDataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity)
	if err != nil {
		return nil, err
	}
	last, err := requiredOffset(directoryDataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		if first != 0 || last != 0 {
			return nil, fmt.Errorf("empty Directory Record Sequence has non-zero root offsets")
		}
		return nil, nil
	}
	if first == 0 || last == 0 {
		return nil, fmt.Errorf("non-empty Directory Record Sequence has zero root offset")
	}

	byOffset := make(map[uint32]*Record, len(records))
	for _, record := range records {
		byOffset[record.offset] = record
		record.children = nil
	}
	visited := make(map[uint32]struct{}, len(records))
	var buildLevel func(uint32) ([]*Record, error)
	buildLevel = func(start uint32) ([]*Record, error) {
		level := make([]*Record, 0)
		for offset := start; offset != 0; {
			record, ok := byOffset[offset]
			if !ok {
				return nil, fmt.Errorf("directory record offset %d does not identify a sequence item", offset)
			}
			if _, exists := visited[offset]; exists {
				return nil, fmt.Errorf("directory record offset %d is referenced more than once or forms a cycle", offset)
			}
			visited[offset] = struct{}{}
			level = append(level, record)

			lower, err := requiredOffset(record.dataset, tag.OffsetOfReferencedLowerLevelDirectoryEntity)
			if err != nil {
				return nil, fmt.Errorf("record at offset %d: %w", offset, err)
			}
			if lower != 0 {
				children, err := buildLevel(lower)
				if err != nil {
					return nil, err
				}
				record.children = children
			}

			next, err := requiredOffset(record.dataset, tag.OffsetOfTheNextDirectoryRecord)
			if err != nil {
				return nil, fmt.Errorf("record at offset %d: %w", offset, err)
			}
			offset = next
		}
		return level, nil
	}

	roots, err := buildLevel(first)
	if err != nil {
		return nil, err
	}
	if len(roots) == 0 || roots[len(roots)-1].offset != last {
		return nil, fmt.Errorf("last root offset %d does not match the root sibling chain", last)
	}
	if len(visited) != len(records) {
		return nil, fmt.Errorf("Directory Record Sequence contains %d unreachable item(s)", len(records)-len(visited))
	}
	return roots, nil
}

func requiredOffset(ds *dataset.Dataset, offsetTag *tag.Tag) (uint32, error) {
	value, err := ds.GetUInt32(offsetTag, 0)
	if err != nil {
		return 0, fmt.Errorf("read required offset %s: %w", offsetTag, err)
	}
	values, err := ds.GetUInt32s(offsetTag)
	if err != nil {
		return 0, fmt.Errorf("read required offset %s: %w", offsetTag, err)
	}
	if len(values) != 1 {
		return 0, fmt.Errorf("required offset %s has VM %d, want 1", offsetTag, len(values))
	}
	return value, nil
}

// recoverOffsetDelta returns stored-actual when exactly one fixed delta maps
// every non-zero stored reference to a physical item offset.
func recoverOffsetDelta(references, actualOffsets []uint32) (int64, error) {
	filtered := make([]uint32, 0, len(references))
	for _, reference := range references {
		if reference != 0 {
			filtered = append(filtered, reference)
		}
	}
	if len(filtered) == 0 || len(actualOffsets) == 0 {
		return 0, fmt.Errorf("no non-zero offsets are available for recovery")
	}
	actual := make(map[int64]struct{}, len(actualOffsets))
	for _, offset := range actualOffsets {
		actual[int64(offset)] = struct{}{}
	}

	candidates := make(map[int64]struct{})
	for _, offset := range actualOffsets {
		delta := int64(filtered[0]) - int64(offset)
		matches := true
		for _, reference := range filtered[1:] {
			if _, ok := actual[int64(reference)-delta]; !ok {
				matches = false
				break
			}
		}
		if matches {
			candidates[delta] = struct{}{}
		}
	}
	if len(candidates) != 1 {
		values := make([]int64, 0, len(candidates))
		for candidate := range candidates {
			values = append(values, candidate)
		}
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
		return 0, fmt.Errorf("offset recovery is ambiguous: candidate deltas %v", values)
	}
	for candidate := range candidates {
		return candidate, nil
	}
	panic("unreachable")
}
