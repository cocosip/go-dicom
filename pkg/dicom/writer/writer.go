// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Writer writes DICOM files.
type Writer struct {
	writer         io.Writer
	byteOrder      binary.ByteOrder
	isExplicitVR   bool
	transferSyntax *transfer.TransferSyntax

	// Write options
	includePreamble             bool   // Whether to include 128-byte preamble + DICM
	explicitLengthSequences     bool   // Use explicit length for sequences
	explicitLengthSequenceItems bool   // Use explicit length for sequence items
	keepGroupLengths            bool   // Keep group length tags (0xGGGG,0x0000)
	largeObjectSize             uint32 // Threshold for large objects
}

// WriteOption is a functional option for Write function.
type WriteOption func(*writeConfig)

// writeConfig holds the configuration for a write operation.
type writeConfig struct {
	transferSyntax              *transfer.TransferSyntax
	fileMetaInfo                *dataset.Dataset
	includePreamble             bool
	explicitLengthSequences     bool   // Use explicit length for sequences (default: false, use undefined)
	explicitLengthSequenceItems bool   // Use explicit length for sequence items (default: false, use undefined)
	keepGroupLengths            bool   // Keep group length tags (0xGGGG,0x0000) (default: false)
	largeObjectSize             uint32 // Threshold for large objects (default: 1MB)
}

// WithTransferSyntax specifies the transfer syntax to use.
// If not specified, defaults to Explicit VR Little Endian.
func WithTransferSyntax(ts *transfer.TransferSyntax) WriteOption {
	return func(c *writeConfig) {
		c.transferSyntax = ts
	}
}

// WithFileMetaInfo specifies the File Meta Information to write.
// If not specified, minimal File Meta Information will be auto-generated.
func WithFileMetaInfo(fmi *dataset.Dataset) WriteOption {
	return func(c *writeConfig) {
		c.fileMetaInfo = fmi
	}
}

// WithoutPreamble configures the writer to skip the preamble and DICM prefix.
// This is useful for writing datasets that are not standalone files.
func WithoutPreamble() WriteOption {
	return func(c *writeConfig) {
		c.includePreamble = false
	}
}

// WithExplicitLengthSequences configures the writer to use explicit lengths for sequences.
// By default, sequences are written with undefined length (0xFFFFFFFF) and delimited by
// Sequence Delimitation Items.
func WithExplicitLengthSequences() WriteOption {
	return func(c *writeConfig) {
		c.explicitLengthSequences = true
	}
}

// WithExplicitLengthSequenceItems configures the writer to use explicit lengths for sequence items.
// By default, sequence items are written with undefined length (0xFFFFFFFF) and delimited by
// Item Delimitation Items.
func WithExplicitLengthSequenceItems() WriteOption {
	return func(c *writeConfig) {
		c.explicitLengthSequenceItems = true
	}
}

// WithKeepGroupLengths configures the writer to keep group length tags (GGGG,0000).
// By default, group length tags are removed as they are deprecated in DICOM.
func WithKeepGroupLengths() WriteOption {
	return func(c *writeConfig) {
		c.keepGroupLengths = true
	}
}

// WithLargeObjectSize sets the threshold for what constitutes a "large" object during writing.
// This can be used for optimization purposes (e.g., streaming large pixel data).
// Default is 1MB (1024*1024 bytes). Set to 0 to use the default.
func WithLargeObjectSize(size uint32) WriteOption {
	return func(c *writeConfig) {
		c.largeObjectSize = size
	}
}

// Option is a functional option for Writer (internal use).
type Option func(*Writer)

// New creates a new Writer with the given transfer syntax and options.
// If ts is nil, defaults to Explicit VR Little Endian.
func New(ts *transfer.TransferSyntax, opts ...Option) *Writer {
	// Default to Explicit VR Little Endian if no transfer syntax specified
	if ts == nil {
		ts = transfer.ExplicitVRLittleEndian
	}

	w := &Writer{
		transferSyntax:  ts,
		isExplicitVR:    ts.IsExplicitVR(),
		includePreamble: true, // Default to including preamble
	}

	// Set byte order based on transfer syntax
	if ts.Endian() == 1 { // Big endian
		w.byteOrder = binary.BigEndian
	} else {
		w.byteOrder = binary.LittleEndian
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

// Write writes a DICOM dataset to the writer.
//
// Basic usage (with defaults):
//   writer.Write(w, ds)  // Uses Explicit VR Little Endian, auto-generates File Meta Info
//
// With options:
//   writer.Write(w, ds, writer.WithTransferSyntax(ts))
//   writer.Write(w, ds, writer.WithFileMetaInfo(fmi))
//   writer.Write(w, ds, writer.WithoutPreamble())
//
// The structure written is:
//   - Preamble (128 bytes) + DICM prefix (4 bytes) [default, can be disabled with WithoutPreamble()]
//   - File Meta Information (Group 0002, always Explicit VR Little Endian)
//   - Dataset (encoding depends on Transfer Syntax)
func Write(w io.Writer, ds *dataset.Dataset, opts ...WriteOption) error {
	// Apply options to configuration
	config := &writeConfig{
		transferSyntax:              transfer.ExplicitVRLittleEndian, // Default
		fileMetaInfo:                nil,                              // Will be auto-generated
		includePreamble:             true,                             // Default to including preamble
		explicitLengthSequences:     false,                            // Default: use undefined length
		explicitLengthSequenceItems: false,                            // Default: use undefined length
		keepGroupLengths:            false,                            // Default: remove group lengths
		largeObjectSize:             1024 * 1024,                      // Default: 1MB
	}

	for _, opt := range opts {
		opt(config)
	}

	// If largeObjectSize is explicitly set to 0, use default
	if config.largeObjectSize == 0 {
		config.largeObjectSize = 1024 * 1024
	}

	// Create internal writer
	writer := &Writer{
		writer:                      w,
		transferSyntax:              config.transferSyntax,
		isExplicitVR:                config.transferSyntax.IsExplicitVR(),
		includePreamble:             config.includePreamble,
		explicitLengthSequences:     config.explicitLengthSequences,
		explicitLengthSequenceItems: config.explicitLengthSequenceItems,
		keepGroupLengths:            config.keepGroupLengths,
		largeObjectSize:             config.largeObjectSize,
	}

	// Set byte order based on transfer syntax
	if config.transferSyntax.Endian() == 1 { // Big endian
		writer.byteOrder = binary.BigEndian
	} else {
		writer.byteOrder = binary.LittleEndian
	}

	// Auto-generate File Meta Information if not provided
	fileMetaInfo := config.fileMetaInfo
	if fileMetaInfo == nil {
		fileMetaInfo = writer.generateFileMetaInformation()
	} else {
		// Ensure TransferSyntaxUID is present in the provided fileMetaInfo
		if _, exists := fileMetaInfo.Get(tag.TransferSyntaxUID); !exists {
			fileMetaInfo.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
				[]string{config.transferSyntax.UID().String()}))
		}
	}

	// Write preamble if requested
	if writer.includePreamble {
		if err := writer.writePreamble(); err != nil {
			return fmt.Errorf("failed to write preamble: %w", err)
		}
	}

	// Write File Meta Information (always Explicit VR Little Endian)
	if err := writer.writeFileMetaInformation(fileMetaInfo); err != nil {
		return fmt.Errorf("failed to write file meta information: %w", err)
	}

	// Write main dataset
	if err := writer.writeDataset(ds); err != nil {
		return fmt.Errorf("failed to write dataset: %w", err)
	}

	return nil
}

// WriteFile writes a DICOM dataset to a file.
//
// Basic usage:
//   writer.WriteFile("output.dcm", ds)
//
// With options:
//   writer.WriteFile("output.dcm", ds, writer.WithTransferSyntax(ts))
func WriteFile(path string, ds *dataset.Dataset, opts ...WriteOption) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()

	return Write(file, ds, opts...)
}

// generateFileMetaInformation generates a minimal File Meta Information dataset.
func (w *Writer) generateFileMetaInformation() *dataset.Dataset {
	fileMetaInfo := dataset.New()

	// Add required Transfer Syntax UID (0002,0010)
	fileMetaInfo.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
		[]string{w.transferSyntax.UID().String()}))

	// Note: FileMetaInformationGroupLength (0002,0000) will be calculated
	// automatically in writeFileMetaInformation()

	return fileMetaInfo
}

// writePreamble writes the 128-byte preamble and DICM prefix.
func (w *Writer) writePreamble() error {
	// Write 128-byte preamble (all zeros)
	preamble := make([]byte, 128)
	if _, err := w.writer.Write(preamble); err != nil {
		return fmt.Errorf("failed to write preamble: %w", err)
	}

	// Write DICM prefix
	if _, err := w.writer.Write([]byte("DICM")); err != nil {
		return fmt.Errorf("failed to write DICM prefix: %w", err)
	}

	return nil
}

// writeFileMetaInformation writes Group 0002 elements.
// This is always written as Explicit VR Little Endian.
func (w *Writer) writeFileMetaInformation(ds *dataset.Dataset) error {
	if ds == nil {
		return nil
	}

	// Save current settings
	savedByteOrder := w.byteOrder
	savedIsExplicitVR := w.isExplicitVR

	// File Meta Information is always Explicit VR Little Endian
	w.byteOrder = binary.LittleEndian
	w.isExplicitVR = true

	// Calculate the length of all Group 0002 elements (excluding 0002,0000 itself)
	// We need to write them to a temporary buffer first to calculate the length
	tempBuf := &bytes.Buffer{}
	tempWriter := &Writer{
		writer:       tempBuf,
		byteOrder:    binary.LittleEndian,
		isExplicitVR: true,
	}

	elements := ds.Elements()
	for _, elem := range elements {
		// Only write Group 0002 elements, but skip FileMetaInformationGroupLength if present
		if elem.Tag().Group() == 0x0002 && elem.Tag().Element() != 0x0000 {
			if err := tempWriter.writeElement(elem); err != nil {
				// Restore settings before returning
				w.byteOrder = savedByteOrder
				w.isExplicitVR = savedIsExplicitVR
				return err
			}
		}
	}

	// Now write the FileMetaInformationGroupLength tag first
	groupLengthTag := tag.FileMetaInformationGroupLength
	if err := w.writeTag(groupLengthTag); err != nil {
		w.byteOrder = savedByteOrder
		w.isExplicitVR = savedIsExplicitVR
		return err
	}

	// Write VR (UL)
	if err := w.writeVR(vr.UL); err != nil {
		w.byteOrder = savedByteOrder
		w.isExplicitVR = savedIsExplicitVR
		return err
	}

	// Write length (UL is 16-bit length)
	if err := binary.Write(w.writer, w.byteOrder, uint16(4)); err != nil {
		w.byteOrder = savedByteOrder
		w.isExplicitVR = savedIsExplicitVR
		return err
	}

	// Write the group length value
	// Note: According to DICOM standard, FileMetaInformationGroupLength contains
	// the number of bytes following this element up to and including the last
	// File Meta Information Group element (i.e., excluding the Group Length element itself)
	groupLength := uint32(tempBuf.Len())
	if err := binary.Write(w.writer, w.byteOrder, groupLength); err != nil {
		w.byteOrder = savedByteOrder
		w.isExplicitVR = savedIsExplicitVR
		return err
	}

	// Now write the actual elements
	if _, err := w.writer.Write(tempBuf.Bytes()); err != nil {
		w.byteOrder = savedByteOrder
		w.isExplicitVR = savedIsExplicitVR
		return err
	}

	// Restore original settings
	w.byteOrder = savedByteOrder
	w.isExplicitVR = savedIsExplicitVR

	return nil
}

// writeDataset writes a dataset (collection of elements).
func (w *Writer) writeDataset(ds *dataset.Dataset) error {
	if ds == nil {
		return nil
	}

	elements := ds.Elements()
	for _, elem := range elements {
		// Skip Group 0002 elements (they're written in File Meta Information)
		if elem.Tag().Group() == 0x0002 {
			continue
		}

		// Skip group length tags unless explicitly requested to keep them
		// Group length tags have element number 0x0000
		if !w.keepGroupLengths && elem.Tag().Element() == 0x0000 {
			continue
		}

		if err := w.writeElement(elem); err != nil {
			return err
		}
	}

	return nil
}

// writeElement writes a single DICOM element.
func (w *Writer) writeElement(elem element.Element) error {
	// Check if this is a sequence (handle specially)
	if seq, ok := elem.(*dataset.Sequence); ok {
		return w.writeSequence(seq)
	}

	// Write tag
	if err := w.writeTag(elem.Tag()); err != nil {
		return fmt.Errorf("failed to write tag %s: %w", elem.Tag(), err)
	}

	// Write VR (if Explicit VR)
	elemVR := elem.ValueRepresentation()
	if w.isExplicitVR {
		if err := w.writeVR(elemVR); err != nil {
			return fmt.Errorf("failed to write VR for tag %s: %w", elem.Tag(), err)
		}
	}

	// Get value bytes
	valueBytes := elem.Buffer().Data()

	// Write length
	if err := w.writeLength(elemVR, uint32(len(valueBytes))); err != nil {
		return fmt.Errorf("failed to write length for tag %s: %w", elem.Tag(), err)
	}

	// Write value
	if len(valueBytes) > 0 {
		if _, err := w.writer.Write(valueBytes); err != nil {
			return fmt.Errorf("failed to write value for tag %s: %w", elem.Tag(), err)
		}
	}

	return nil
}

// writeTag writes a DICOM tag (4 bytes).
func (w *Writer) writeTag(t *tag.Tag) error {
	if err := binary.Write(w.writer, w.byteOrder, t.Group()); err != nil {
		return err
	}
	if err := binary.Write(w.writer, w.byteOrder, t.Element()); err != nil {
		return err
	}
	return nil
}

// writeVR writes the Value Representation (2 bytes).
func (w *Writer) writeVR(v *vr.VR) error {
	vrCode := v.Code()
	if len(vrCode) != 2 {
		return fmt.Errorf("invalid VR code length: %d", len(vrCode))
	}
	_, err := w.writer.Write([]byte(vrCode))
	return err
}

// writeLength writes the value length field.
func (w *Writer) writeLength(v *vr.VR, length uint32) error {
	if w.isExplicitVR {
		// Check if VR has 16-bit or 32-bit length
		if v.Is16bitLength() {
			// 16-bit length
			if length > 0xFFFF {
				return fmt.Errorf("length %d exceeds maximum for 16-bit VR %s", length, v.Code())
			}
			return binary.Write(w.writer, w.byteOrder, uint16(length))
		}

		// 32-bit length: write 2 reserved bytes first
		reserved := uint16(0)
		if err := binary.Write(w.writer, w.byteOrder, reserved); err != nil {
			return err
		}
		return binary.Write(w.writer, w.byteOrder, length)
	}

	// Implicit VR: always 32-bit length
	return binary.Write(w.writer, w.byteOrder, length)
}

// writeSequence writes a sequence element (VR=SQ).
func (w *Writer) writeSequence(seq *dataset.Sequence) error {
	// Write the sequence tag first
	if err := w.writeTag(seq.Tag()); err != nil {
		return fmt.Errorf("failed to write sequence tag %s: %w", seq.Tag(), err)
	}

	// Write VR (if explicit)
	if w.isExplicitVR {
		if err := w.writeVR(vr.SQ); err != nil {
			return fmt.Errorf("failed to write SQ VR: %w", err)
		}

		// Write reserved bytes for 32-bit length
		reserved := uint16(0)
		if err := binary.Write(w.writer, w.byteOrder, reserved); err != nil {
			return err
		}
	}

	// Choose between explicit and undefined length
	if w.explicitLengthSequences {
		// Write sequence with explicit length
		// Need to write items to a buffer first to calculate length
		itemsBuf := &bytes.Buffer{}
		itemsWriter := &Writer{
			writer:                      itemsBuf,
			byteOrder:                   w.byteOrder,
			isExplicitVR:                w.isExplicitVR,
			explicitLengthSequenceItems: w.explicitLengthSequenceItems,
			keepGroupLengths:            w.keepGroupLengths,
		}

		for i := 0; i < seq.Count(); i++ {
			item := seq.GetItem(i)
			if err := itemsWriter.writeItem(item); err != nil {
				return fmt.Errorf("failed to write item %d: %w", i, err)
			}
		}

		// Write the explicit length
		seqLength := uint32(itemsBuf.Len())
		if err := binary.Write(w.writer, w.byteOrder, seqLength); err != nil {
			return err
		}

		// Write the items data
		if _, err := w.writer.Write(itemsBuf.Bytes()); err != nil {
			return err
		}
	} else {
		// Write undefined length
		undefinedLength := uint32(0xFFFFFFFF)
		if err := binary.Write(w.writer, w.byteOrder, undefinedLength); err != nil {
			return err
		}

		// Write items
		for i := 0; i < seq.Count(); i++ {
			item := seq.GetItem(i)
			if err := w.writeItem(item); err != nil {
				return fmt.Errorf("failed to write item %d: %w", i, err)
			}
		}

		// Write Sequence Delimitation Item (FFFE,E0DD)
		delimTag := tag.New(0xFFFE, 0xE0DD)
		if err := w.writeTag(delimTag); err != nil {
			return err
		}
		if err := binary.Write(w.writer, w.byteOrder, uint32(0)); err != nil {
			return err
		}
	}

	return nil
}

// writeItem writes a single item within a sequence.
func (w *Writer) writeItem(item *dataset.Dataset) error {
	// Write Item tag (FFFE,E000)
	itemTag := tag.New(0xFFFE, 0xE000)
	if err := w.writeTag(itemTag); err != nil {
		return err
	}

	// Choose between explicit and undefined length
	if w.explicitLengthSequenceItems {
		// Write item with explicit length
		// Need to write elements to a buffer first to calculate length
		elementsBuf := &bytes.Buffer{}
		elementsWriter := &Writer{
			writer:                      elementsBuf,
			byteOrder:                   w.byteOrder,
			isExplicitVR:                w.isExplicitVR,
			explicitLengthSequences:     w.explicitLengthSequences,
			explicitLengthSequenceItems: w.explicitLengthSequenceItems,
			keepGroupLengths:            w.keepGroupLengths,
		}

		// Write all elements in the item
		elements := item.Elements()
		for _, elem := range elements {
			if err := elementsWriter.writeElement(elem); err != nil {
				return err
			}
		}

		// Write the explicit length
		itemLength := uint32(elementsBuf.Len())
		if err := binary.Write(w.writer, w.byteOrder, itemLength); err != nil {
			return err
		}

		// Write the elements data
		if _, err := w.writer.Write(elementsBuf.Bytes()); err != nil {
			return err
		}
	} else {
		// Write undefined length for item
		undefinedLength := uint32(0xFFFFFFFF)
		if err := binary.Write(w.writer, w.byteOrder, undefinedLength); err != nil {
			return err
		}

		// Write all elements in the item
		elements := item.Elements()
		for _, elem := range elements {
			if err := w.writeElement(elem); err != nil {
				return err
			}
		}

		// Write Item Delimitation Item (FFFE,E00D)
		delimTag := tag.New(0xFFFE, 0xE00D)
		if err := w.writeTag(delimTag); err != nil {
			return err
		}
		if err := binary.Write(w.writer, w.byteOrder, uint32(0)); err != nil {
			return err
		}
	}

	return nil
}
