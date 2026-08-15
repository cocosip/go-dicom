// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

// StructuredReport represents a DICOM Structured Report.
//
// A Structured Report is a hierarchical document consisting of content items
// arranged in a tree structure. The root is typically a CONTAINER item.
//
// Reference: DICOM Part 3, Annex A.35
type StructuredReport struct {
	*ContentItem
	fileMeta       *dataset.Dataset
	transferSyntax *transfer.Syntax
}

// NewStructuredReport creates a new structured report with the given root code
func NewStructuredReport(rootCode *CodeItem, items ...*ContentItem) (*StructuredReport, error) {
	// Create root container (relationship type is not needed for root)
	ds := dataset.New()

	// Set concept name code sequence
	if err := setConceptNameCode(ds, rootCode); err != nil {
		return nil, err
	}

	// Set value type to CONTAINER (0040,A040) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ValueType, vr.CS, []string{string(ValueTypeContainer)}))

	// Set continuity to SEPARATE (typical for SR root) (0040,A050) VR=CS
	_ = ds.AddOrUpdate(element.NewString(tag.ContinuityOfContent, vr.CS, []string{string(ContinuitySeparate)}))

	// Note: RelationshipType is intentionally NOT set for the root element

	// Add content sequence if items provided (0040,A730) VR=SQ
	if len(items) > 0 {
		datasets, err := contentItemDatasets(items)
		if err != nil {
			return nil, err
		}

		seq := dataset.NewSequenceWithItems(tag.ContentSequence, datasets)
		_ = ds.AddOrUpdate(seq)
	}

	contentItem := NewContentItemFromDataset(ds)
	return &StructuredReport{ContentItem: contentItem}, nil
}

// NewStructuredReportFromDataset creates a StructuredReport from an existing dataset
func NewStructuredReportFromDataset(ds *dataset.Dataset) *StructuredReport {
	return &StructuredReport{
		ContentItem: NewContentItemFromDataset(ds),
	}
}

// Add adds a content item to the report
func (sr *StructuredReport) Add(item *ContentItem) error {
	if sr == nil || sr.ContentItem == nil {
		return NewError("structured report is nil")
	}
	if item == nil {
		return NewError("cannot add nil item")
	}

	// Get the existing content sequence or create a new one. A present but
	// malformed sequence must not be silently replaced.
	if !sr.dataset.Contains(tag.ContentSequence) {
		// Create new sequence with this item
		seq := dataset.NewSequenceWithItems(tag.ContentSequence, []*dataset.Dataset{item.Dataset()})
		return sr.dataset.AddOrUpdate(seq)
	}
	seq, err := sr.dataset.GetSequence(tag.ContentSequence)
	if err != nil {
		return WrapError("get Content Sequence", err)
	}
	seq.AddItem(item.Dataset())

	return nil
}

// AddText is a convenience method to add a text content item
func (sr *StructuredReport) AddText(code *CodeItem, relationship Relationship, text string) error {
	item, err := NewContentItemText(code, relationship, text)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddCode is a convenience method to add a code content item
func (sr *StructuredReport) AddCode(code *CodeItem, relationship Relationship, value *CodeItem) error {
	item, err := NewContentItemCode(code, relationship, value)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddNumeric is a convenience method to add a numeric content item
func (sr *StructuredReport) AddNumeric(code *CodeItem, relationship Relationship, value *MeasuredValue) error {
	item, err := NewContentItemNumeric(code, relationship, value)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// AddContainer is a convenience method to add a container content item
func (sr *StructuredReport) AddContainer(code *CodeItem, relationship Relationship, continuity Continuity, items ...*ContentItem) error {
	item, err := NewContentItemContainer(code, relationship, continuity, items...)
	if err != nil {
		return err
	}
	return sr.Add(item)
}

// Open reads and validates a structured report from a file.
func Open(path string, options ...parser.Option) (*StructuredReport, error) {
	result, err := parser.ParseFile(path, options...)
	if err != nil {
		return nil, WrapError("open structured report", err)
	}
	return structuredReportFromParseResult(result)
}

// Read parses and validates a structured report from a stream.
func Read(input io.Reader, options ...parser.Option) (*StructuredReport, error) {
	if input == nil {
		return nil, NewError("structured report reader is nil")
	}
	result, err := parser.Parse(input, options...)
	if err != nil {
		return nil, WrapError("read structured report", err)
	}
	return structuredReportFromParseResult(result)
}

func structuredReportFromParseResult(result *parser.ParseResult) (*StructuredReport, error) {
	if result == nil || result.Dataset == nil {
		return nil, NewError("parsed structured report dataset is nil")
	}
	if result.IsPartial {
		return nil, NewError("partial structured report parse is not accepted")
	}
	report := NewStructuredReportFromDataset(result.Dataset)
	if result.FileMetaInformation != nil {
		report.fileMeta = result.FileMetaInformation.Dataset()
	}
	report.transferSyntax = result.TransferSyntax
	if err := report.Validate(); err != nil {
		return nil, WrapError("validate structured report", err)
	}
	return report, nil
}

// Write validates and writes the structured report to a stream.
// Sequences and sequence items always use explicit lengths for fo-dicom parity.
func (sr *StructuredReport) Write(output io.Writer, options ...writer.WriteOption) error {
	if output == nil {
		return NewError("structured report writer is nil")
	}
	if err := sr.Validate(); err != nil {
		return WrapError("validate structured report", err)
	}
	if err := writer.Write(output, sr.Dataset(), sr.writeOptions(options)...); err != nil {
		return WrapError("write structured report", err)
	}
	return nil
}

// Save validates and writes the structured report to a file.
func (sr *StructuredReport) Save(path string, options ...writer.WriteOption) error {
	if err := sr.Validate(); err != nil {
		return WrapError("validate structured report", err)
	}
	if err := writer.WriteFile(path, sr.Dataset(), sr.writeOptions(options)...); err != nil {
		return WrapError("save structured report", err)
	}
	return nil
}

func (sr *StructuredReport) writeOptions(options []writer.WriteOption) []writer.WriteOption {
	result := make([]writer.WriteOption, 0, len(options)+4)
	if sr.fileMeta != nil {
		result = append(result, writer.WithFileMetaInfo(sr.fileMeta))
	}
	if sr.transferSyntax != nil {
		result = append(result, writer.WithTransferSyntax(sr.transferSyntax))
	}
	result = append(result, options...)
	result = append(result,
		writer.WithExplicitLengthSequences(true),
		writer.WithExplicitLengthSequenceItems(),
	)
	return result
}
