// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

type writeConfig struct {
	verifyPositions bool
}

// WriteOption configures DICOMDIR writing.
type WriteOption func(*writeConfig)

// WithWriteVerification controls comparison of first- and second-pass Item positions.
func WithWriteVerification(enabled bool) WriteOption {
	return func(config *writeConfig) {
		config.verifyPositions = enabled
	}
}

// Write writes the DICOMDIR using two passes so all directory offsets are exact.
func (d *Directory) Write(output io.Writer, options ...WriteOption) error {
	if d == nil {
		return fmt.Errorf("directory cannot be nil")
	}
	if output == nil {
		return fmt.Errorf("DICOMDIR writer cannot be nil")
	}
	if d.dataset == nil || d.fileMeta == nil {
		return fmt.Errorf("DICOMDIR Dataset and File Meta Information cannot be nil")
	}
	if err := validateDirectoryTransferSyntax(d.transferSyntax); err != nil {
		return err
	}
	config := writeConfig{verifyPositions: true}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}

	records, err := flattenAndValidateRecords(d)
	if err != nil {
		return err
	}
	items := make([]*dataset.Dataset, len(records))
	for i, record := range records {
		items[i] = record.dataset
	}
	if err := d.dataset.AddOrUpdate(dataset.NewSequenceWithItems(tag.DirectoryRecordSequence, items)); err != nil {
		return fmt.Errorf("set Directory Record Sequence: %w", err)
	}
	if err := zeroDirectoryOffsets(d.dataset, records); err != nil {
		return err
	}

	firstPositions, err := d.writePass(io.Discard, records, true)
	if err != nil {
		return fmt.Errorf("measure DICOMDIR record positions: %w", err)
	}
	for _, record := range records {
		record.offset = firstPositions[record]
	}
	if len(records) == 0 {
		if err := setOffsetValue(d.dataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, 0); err != nil {
			return err
		}
		if err := setOffsetValue(d.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, 0); err != nil {
			return err
		}
	} else if err := writeTreeOffsets(d.dataset, records, d.roots); err != nil {
		return err
	}

	secondPositions, err := d.writePass(output, records, config.verifyPositions)
	if err != nil {
		return fmt.Errorf("write DICOMDIR: %w", err)
	}
	if config.verifyPositions {
		for _, record := range records {
			if secondPositions[record] != firstPositions[record] {
				return fmt.Errorf("DICOMDIR Item position changed between write passes: %d != %d", secondPositions[record], firstPositions[record])
			}
		}
	}
	return nil
}

// Save creates or truncates path and writes the DICOMDIR.
func (d *Directory) Save(path string, options ...WriteOption) error {
	cleanPath := filepath.Clean(path)
	file, err := os.Create(cleanPath)
	if err != nil {
		return fmt.Errorf("create DICOMDIR %s: %w", cleanPath, err)
	}
	writeErr := d.Write(file, options...)
	closeErr := file.Close()
	if writeErr != nil && closeErr != nil {
		return errors.Join(writeErr, closeErr)
	}
	if writeErr != nil {
		return writeErr
	}
	if closeErr != nil {
		return fmt.Errorf("close DICOMDIR %s: %w", cleanPath, closeErr)
	}
	return nil
}

func flattenAndValidateRecords(d *Directory) ([]*Record, error) {
	if _, err := requiredOffset(d.dataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity); err != nil {
		return nil, err
	}
	if _, err := requiredOffset(d.dataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity); err != nil {
		return nil, err
	}
	state := make(map[*Record]uint8)
	records := make([]*Record, 0)
	var visit func([]*Record) error
	visit = func(level []*Record) error {
		for _, record := range level {
			if record == nil {
				return fmt.Errorf("directory record graph contains a nil record")
			}
			switch state[record] {
			case 1:
				return fmt.Errorf("directory record graph contains a cycle")
			case 2:
				return fmt.Errorf("directory record is reachable from more than one parent")
			}
			if record.dataset == nil {
				return fmt.Errorf("%s directory record has a nil Dataset", record.recordType)
			}
			if _, err := requiredOffset(record.dataset, tag.OffsetOfTheNextDirectoryRecord); err != nil {
				return nilRecordError(record, err)
			}
			if _, err := requiredOffset(record.dataset, tag.OffsetOfReferencedLowerLevelDirectoryEntity); err != nil {
				return nilRecordError(record, err)
			}
			state[record] = 1
			records = append(records, record)
			if err := visit(record.children); err != nil {
				return err
			}
			state[record] = 2
		}
		return nil
	}
	if err := visit(d.roots); err != nil {
		return nil, err
	}
	return records, nil
}

func nilRecordError(record *Record, err error) error {
	return fmt.Errorf("%s directory record: %w", record.recordType, err)
}

func zeroDirectoryOffsets(directoryDataset *dataset.Dataset, records []*Record) error {
	if err := setOffsetValue(directoryDataset, tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, 0); err != nil {
		return err
	}
	if err := setOffsetValue(directoryDataset, tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, 0); err != nil {
		return err
	}
	for _, record := range records {
		record.offset = 0
		if err := setOffsetValue(record.dataset, tag.OffsetOfTheNextDirectoryRecord, 0); err != nil {
			return err
		}
		if err := setOffsetValue(record.dataset, tag.OffsetOfReferencedLowerLevelDirectoryEntity, 0); err != nil {
			return err
		}
	}
	return nil
}

func (d *Directory) writePass(output io.Writer, records []*Record, observe bool) (map[*Record]uint32, error) {
	positions := make(map[*Record]uint32, len(records))
	byDataset := make(map[*dataset.Dataset]*Record, len(records))
	for _, record := range records {
		byDataset[record.dataset] = record
	}
	writeOptions := []writer.WriteOption{
		writer.WithTransferSyntax(d.transferSyntax),
		writer.WithFileMetaInfo(d.fileMeta),
	}
	if observe {
		writeOptions = append(writeOptions, writer.WithSequenceItemObserver(func(position writer.SequenceItemPosition) error {
			if !position.SequenceTag.Equals(tag.DirectoryRecordSequence) {
				return nil
			}
			record, ok := byDataset[position.Item]
			if !ok {
				return fmt.Errorf("Directory Record Sequence contains an unexpected item")
			}
			if _, exists := positions[record]; exists {
				return fmt.Errorf("directory record Item position was reported more than once")
			}
			offset, err := checkedItemOffset(position.Offset)
			if err != nil {
				return err
			}
			positions[record] = offset
			return nil
		}))
	}
	if err := writer.Write(output, d.dataset, writeOptions...); err != nil {
		return nil, err
	}
	if observe && len(positions) != len(records) {
		return nil, fmt.Errorf("observed %d directory Item positions, want %d", len(positions), len(records))
	}
	return positions, nil
}

func checkedItemOffset(position uint64) (uint32, error) {
	if position == 0 || position > math.MaxUint32 {
		return 0, fmt.Errorf("directory Item position %d cannot be represented as UL", position)
	}
	return uint32(position), nil
}
