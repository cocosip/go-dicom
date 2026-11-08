// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// OffsetPolicy controls whether stale directory record offsets may be repaired.
type OffsetPolicy uint8

const (
	// CompatibleOffsets repairs offsets only when the resulting tree is unique and valid.
	CompatibleOffsets OffsetPolicy = iota
	// StrictOffsets requires every stored offset to match a physical sequence item position.
	StrictOffsets
)

type openConfig struct {
	offsetPolicy OffsetPolicy
}

// OpenOption configures DICOMDIR reading.
type OpenOption func(*openConfig)

// WithOffsetPolicy selects strict validation or bounded compatible recovery.
func WithOffsetPolicy(policy OffsetPolicy) OpenOption {
	return func(config *openConfig) {
		config.offsetPolicy = policy
	}
}

// Open reads a DICOMDIR from path.
func Open(path string, options ...OpenOption) (*Directory, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open DICOMDIR: %w", err)
	}
	dir, readErr := Read(file, options...)
	closeErr := file.Close()
	if readErr != nil {
		return nil, readErr
	}
	if closeErr != nil {
		return nil, fmt.Errorf("close DICOMDIR: %w", closeErr)
	}
	return dir, nil
}

// Read parses a DICOMDIR and constructs its directory record tree.
func Read(r io.Reader, options ...OpenOption) (*Directory, error) {
	if r == nil {
		return nil, fmt.Errorf("DICOMDIR reader cannot be nil")
	}
	config := openConfig{offsetPolicy: CompatibleOffsets}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}
	if config.offsetPolicy != CompatibleOffsets && config.offsetPolicy != StrictOffsets {
		return nil, fmt.Errorf("invalid DICOMDIR offset policy %d", config.offsetPolicy)
	}

	positions := make(map[*dataset.Dataset]uint64)
	result, err := parser.Parse(r,
		parser.WithReadOption(parser.ReadAll),
		parser.WithSequenceItemObserver(func(position parser.SequenceItemPosition) error {
			if position.SequenceTag.Equals(tag.DirectoryRecordSequence) {
				positions[position.Item] = position.Offset
			}
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("parse DICOMDIR: %w", err)
	}
	if result.FileMetaInformation == nil {
		return nil, fmt.Errorf("DICOMDIR File Meta Information is missing")
	}
	sopClassUID, ok := result.FileMetaInformation.MediaStorageSOPClassUID()
	if !ok || sopClassUID != uid.MediaStorageDirectoryStorage.UID() {
		return nil, fmt.Errorf("media storage SOP Class must be Media Storage Directory Storage")
	}
	if err := validateDirectoryTransferSyntax(result.TransferSyntax); err != nil {
		return nil, err
	}

	_, hasDirectoryRecords := result.Dataset.Get(tag.DirectoryRecordSequence)
	records := make([]*Record, 0)
	diagnostics := make([]Diagnostic, 0)
	if hasDirectoryRecords {
		sequence, err := result.Dataset.GetSequence(tag.DirectoryRecordSequence)
		if err != nil {
			return nil, fmt.Errorf("read Directory Record Sequence: %w", err)
		}
		records, diagnostics, err = recordsFromSequence(sequence, positions)
		if err != nil {
			return nil, err
		}
		// PS3.3 F.3.2.2, Table F.3-3: FFFFH "shall never be present"; this
		// reader writes exact offsets, so normalize the consistency flag.
		if err := result.Dataset.AddOrUpdate(element.NewUnsignedShort(tag.FileSetConsistencyFlag, []uint16{0})); err != nil {
			return nil, fmt.Errorf("normalize File-set Consistency Flag: %w", err)
		}
	}

	// PS3.3 F.3.1 marks the Directory Information Module Usage U, so a
	// conformant File-set may contain only File-set Identification. Treat the
	// absent module as an empty directory instead of rejecting the file.
	var roots []*Record
	if hasDirectoryRecords {
		repairDiagnostics := []Diagnostic(nil)
		roots, repairDiagnostics, err = buildRecordTree(result.Dataset, records, config.offsetPolicy)
		if err != nil {
			return nil, err
		}
		diagnostics = append(diagnostics, repairDiagnostics...)
	}

	return &Directory{
		dataset:        result.Dataset,
		fileMeta:       result.FileMetaInformation.Dataset(),
		transferSyntax: result.TransferSyntax,
		roots:          roots,
		diagnostics:    diagnostics,
		fileIDs:        referencedFileIDs(records),
	}, nil
}

// referencedFileIDs collects the File IDs already referenced by records so
// AddFile can reject duplicates (PS3.3 F.2.1 rule 6).
func referencedFileIDs(records []*Record) map[string]struct{} {
	fileIDs := make(map[string]struct{}, len(records))
	for _, record := range records {
		components, ok := record.dataset.GetStrings(tag.ReferencedFileID)
		if !ok || len(components) == 0 {
			continue
		}
		fileIDs[strings.Join(components, "/")] = struct{}{}
	}
	return fileIDs
}

func recordsFromSequence(sequence *dataset.Sequence, positions map[*dataset.Dataset]uint64) ([]*Record, []Diagnostic, error) {
	items := sequence.GetItems()
	records := make([]*Record, 0, len(items))
	diagnostics := make([]Diagnostic, 0)
	seenOffsets := make(map[uint32]struct{}, len(items))
	for index, item := range items {
		position, ok := positions[item]
		if !ok {
			return nil, nil, fmt.Errorf("directory record %d has no physical Item position", index)
		}
		if position == 0 || position > math.MaxUint32 {
			return nil, nil, fmt.Errorf("directory record %d Item position %d cannot be represented as UL", index, position)
		}
		offset := uint32(position)
		if _, exists := seenOffsets[offset]; exists {
			return nil, nil, fmt.Errorf("duplicate directory record Item position %d", offset)
		}
		seenOffsets[offset] = struct{}{}

		recordTypeValue, ok := item.GetString(tag.DirectoryRecordType)
		if !ok || recordTypeValue == "" {
			return nil, nil, fmt.Errorf("directory record at offset %d has no Directory Record Type", offset)
		}
		recordType := RecordType(recordTypeValue)
		// PS3.3 F.3.2.2, Table F.3-3: 0000H "shall never be present"; any other
		// value shall be interpreted as FFFFH, so normalize on read.
		if inUse, err := item.GetUInt16(tag.RecordInUseFlag, 0); err != nil || inUse != 0xFFFF {
			if err := item.AddOrUpdate(element.NewUnsignedShort(tag.RecordInUseFlag, []uint16{0xFFFF})); err != nil {
				return nil, nil, fmt.Errorf("normalize Record In-use Flag at offset %d: %w", offset, err)
			}
		}
		records = append(records, newRecord(recordType, item, offset))
		if !knownRecordType(recordType) {
			diagnostics = append(diagnostics, Diagnostic{
				Code:       DiagnosticUnknownRecordType,
				Severity:   SeverityInfo,
				RecordType: recordType,
				Message:    fmt.Sprintf("unknown directory record type %q was preserved", recordType),
			})
		}
	}
	return records, diagnostics, nil
}

func knownRecordType(recordType RecordType) bool {
	switch recordType {
	case RecordPatient, RecordStudy, RecordSeries, RecordImage, RecordSRDocument, RecordPresentation, RecordKeyObjectDoc:
		return true
	default:
		return false
	}
}
