// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// ReadOption controls how large DICOM elements are read.
type ReadOption int

const (
	// ReadDefault reads all tags normally.
	ReadDefault ReadOption = iota

	// ReadLargeOnDemand reads small tags immediately but keeps the stream open
	// to read large tags on demand. The stream must stay open.
	ReadLargeOnDemand

	// SkipLargeTags skips reading large tags entirely. The stream can be closed.
	SkipLargeTags

	// ReadAll reads all tags completely so the stream can be closed.
	ReadAll
)

// FileFormat represents the structure of a DICOM file.
type FileFormat int

const (
	// FormatUnknown indicates the parser could not determine the file format.
	FormatUnknown FileFormat = iota

	// FormatDICOM3 is a valid DICOM file with preamble and file meta info.
	FormatDICOM3

	// FormatDICOM3NoPreamble is a DICOM file without preamble but with file meta info.
	FormatDICOM3NoPreamble

	// FormatDICOM3NoFileMetaInfo is a DICOM file without preamble or file meta info.
	FormatDICOM3NoFileMetaInfo

	// FormatACRNEMA1 is an ACR-NEMA 1.0 file.
	FormatACRNEMA1

	// FormatACRNEMA2 is an ACR-NEMA 2.0 file.
	FormatACRNEMA2
)

// String returns the string representation of FileFormat.
func (f FileFormat) String() string {
	switch f {
	case FormatDICOM3:
		return "DICOM3"
	case FormatDICOM3NoPreamble:
		return "DICOM3NoPreamble"
	case FormatDICOM3NoFileMetaInfo:
		return "DICOM3NoFileMetaInfo"
	case FormatACRNEMA1:
		return "ACRNEMA1"
	case FormatACRNEMA2:
		return "ACRNEMA2"
	default:
		return "Unknown"
	}
}

// ParseResult represents the result of parsing a DICOM file.
//
// It contains both the File Meta Information (Group 0002) and the main Dataset.
// This structure mimics fo-dicom's DicomFile concept.
type ParseResult struct {
	// FileMetaInformation contains Group 0002 elements with convenience accessors.
	// These elements describe the file format and transfer syntax.
	// Always encoded as Explicit VR Little Endian.
	// Use FileMetaInformationDataset() if you need raw dataset access.
	FileMetaInformation *dataset.FileMetaInformation

	// Dataset contains the main DICOM data elements.
	// Encoding depends on the Transfer Syntax specified in FileMetaInformation.
	Dataset *dataset.Dataset

	// TransferSyntax specifies how the dataset is encoded.
	TransferSyntax *transfer.Syntax

	// Format indicates the detected DICOM file format.
	Format FileFormat

	// IsPartial indicates whether parsing ended prematurely
	// (e.g., due to stop criterion or error recovery).
	IsPartial bool
}

// FileMetaInformationDataset returns the underlying dataset of FileMetaInformation.
// This is a convenience method for accessing the raw dataset.
func (pr *ParseResult) FileMetaInformationDataset() *dataset.Dataset {
	if pr.FileMetaInformation == nil {
		return nil
	}
	return pr.FileMetaInformation.Dataset()
}

// parseContext holds the state during DICOM file parsing.
// This is internal and not exposed to users.
type parseContext struct {
	reader         io.Reader
	byteOrder      binary.ByteOrder
	isExplicitVR   bool
	transferSyntax *transfer.Syntax
	dictionary     *dict.Dictionary

	// firstDatasetTag is the first tag read that doesn't belong to Group 0002
	// It's stored here to be used when reading the main dataset
	firstDatasetTag *tag.Tag

	// Configuration options
	maxElementSize  uint32     // Maximum element size to read (0 = unlimited)
	stopAtTag       *tag.Tag   // Stop parsing when this tag is reached
	readOption      ReadOption // How to handle large elements
	largeObjectSize uint32     // Size threshold for "large" objects (default 64KB)

	// File format detection
	detectedFormat FileFormat
	isPartial      bool

	// For lazy loading support
	seekableReader io.ReadSeeker // Set if reader is seekable (for lazy loading)
	file           *os.File      // Set if reader is a file (for FileByteBuffer)
}

// Option is a functional option for configuring the parser.
type Option func(*parseContext)

// WithMaxElementSize sets the maximum element size to read.
// Elements larger than this will cause an error.
func WithMaxElementSize(size uint32) Option {
	return func(ctx *parseContext) {
		ctx.maxElementSize = size
	}
}

// WithStopAtTag sets a tag to stop parsing at.
// Parsing will stop when this tag is encountered.
func WithStopAtTag(t *tag.Tag) Option {
	return func(ctx *parseContext) {
		ctx.stopAtTag = t
	}
}

// WithDictionary sets the DICOM dictionary for implicit VR lookup.
// Required when parsing files with Implicit VR transfer syntax.
func WithDictionary(d *dict.Dictionary) Option {
	return func(ctx *parseContext) {
		ctx.dictionary = d
	}
}

// WithReadOption sets how large elements should be handled during parsing.
//
// Options:
//   - ReadDefault: Read all elements normally
//   - ReadLargeOnDemand: Read large elements on demand (stream must stay open)
//   - SkipLargeTags: Skip large elements entirely
//   - ReadAll: Read all elements including large ones
func WithReadOption(opt ReadOption) Option {
	return func(ctx *parseContext) {
		ctx.readOption = opt
	}
}

// WithLargeObjectSize sets the threshold for what constitutes a "large" object.
// Elements larger than this size are subject to the ReadOption behavior.
// Default is 64KB (65536 bytes). Set to 0 to use the default.
func WithLargeObjectSize(size uint32) Option {
	return func(ctx *parseContext) {
		ctx.largeObjectSize = size
	}
}

// newParseContext creates a new parse context with the given options.
func newParseContext(opts ...Option) *parseContext {
	ctx := &parseContext{
		byteOrder:       binary.LittleEndian,
		isExplicitVR:    true,
		readOption:      ReadDefault,
		largeObjectSize: 65536, // Default 64KB
		detectedFormat:  FormatUnknown,
	}
	for _, opt := range opts {
		opt(ctx)
	}
	// If largeObjectSize is explicitly set to 0, use default
	if ctx.largeObjectSize == 0 {
		ctx.largeObjectSize = 65536
	}
	return ctx
}

// Parse parses a DICOM file from the reader.
// This is the main entry point for reading DICOM files.
//
// Usage:
//
//	result, err := parser.Parse(reader,
//	    parser.WithReadOption(parser.SkipLargeTags),
//	    parser.WithLargeObjectSize(128*1024),
//	)
//
// Returns a ParseResult containing:
//   - FileMetaInformation: Group 0002 elements
//   - Dataset: Main DICOM data
//   - TransferSyntax: How the dataset is encoded
//   - Format: Detected file format
//   - IsPartial: Whether parsing stopped early
func Parse(r io.Reader, opts ...Option) (*ParseResult, error) {
	ctx := newParseContext(opts...)
	return ctx.parse(r)
}

// parse is the internal parsing implementation.
func (p *parseContext) parse(r io.Reader) (*ParseResult, error) {
	p.reader = r
	p.detectedFormat = FormatDICOM3 // Assume DICOM3 initially

	// Check if reader supports seeking (for lazy loading)
	if rs, ok := r.(io.ReadSeeker); ok {
		p.seekableReader = rs
	}

	// Check if reader is an *os.File (for FileByteBuffer)
	if f, ok := r.(*os.File); ok {
		p.file = f
	}

	// Read and validate preamble + DICM prefix
	if err := p.readPreamble(); err != nil {
		return nil, fmt.Errorf("failed to read preamble: %w", err)
	}

	// Read File Meta Information (Group 0002)
	// This is always Explicit VR Little Endian
	p.byteOrder = binary.LittleEndian
	p.isExplicitVR = true

	metaDS, err := p.readFileMetaInformation()
	if err != nil {
		return nil, fmt.Errorf("failed to read file meta information: %w", err)
	}

	// Get Transfer Syntax from meta information
	if err := p.setTransferSyntax(metaDS); err != nil {
		return nil, fmt.Errorf("failed to set transfer syntax: %w", err)
	}

	// Read main dataset with detected transfer syntax
	mainDS, err := p.readDataset()
	if err != nil {
		return nil, fmt.Errorf("failed to read dataset: %w", err)
	}

	// Wrap metaDS in FileMetaInformation for convenience
	fmi := dataset.NewFileMetaInformationFromDataset(metaDS)

	return &ParseResult{
		FileMetaInformation: fmi,
		Dataset:             mainDS,
		TransferSyntax:      p.transferSyntax,
		Format:              p.detectedFormat,
		IsPartial:           p.isPartial,
	}, nil
}

// ParseFile parses a DICOM file from a file path.
func ParseFile(path string, opts ...Option) (*ParseResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() { _ = file.Close() }()

	return Parse(file, opts...)
}

// readPreamble reads and validates the 128-byte preamble and DICM prefix.
func (p *parseContext) readPreamble() error {
	// Read 128-byte preamble (usually all zeros, but can be anything)
	preamble := make([]byte, 128)
	if _, err := io.ReadFull(p.reader, preamble); err != nil {
		return fmt.Errorf("failed to read preamble: %w", err)
	}

	// Read DICM prefix
	prefix := make([]byte, 4)
	if _, err := io.ReadFull(p.reader, prefix); err != nil {
		return fmt.Errorf("failed to read DICM prefix: %w", err)
	}

	if string(prefix) != "DICM" {
		return fmt.Errorf("invalid DICM prefix: got %q, want \"DICM\"", prefix)
	}

	return nil
}

// readFileMetaInformation reads Group 0002 elements (File Meta Information).
func (p *parseContext) readFileMetaInformation() (*dataset.Dataset, error) {
	ds := dataset.New()

	// Read elements until we leave Group 0002
	for {
		// Peek at the tag first to check if we're still in Group 0002
		// This is important because if we call readElement() on a non-0002 tag,
		// it might try to read a sequence which uses different encoding rules
		t, err := p.readTag()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Stop when we leave Group 0002
		if t.Group() != 0x0002 {
			// Save this tag - it belongs to the main dataset
			// We need to "unread" it for the next phase
			p.firstDatasetTag = t
			break
		}

		// Now read the VR, length, and value
		// We inline the readElement logic here to avoid re-reading the tag
		vrValue, err := p.readVR(t)
		if err != nil {
			return nil, fmt.Errorf("failed to read VR for tag %s: %w", t, err)
		}

		length, err := p.readLength(vrValue)
		if err != nil {
			return nil, fmt.Errorf("failed to read length for tag %s: %w", t, err)
		}

		// Check max element size
		if p.maxElementSize > 0 && length > p.maxElementSize {
			return nil, fmt.Errorf("element size %d exceeds maximum %d for tag %s", length, p.maxElementSize, t)
		}

		// Read value data
		data := make([]byte, length)
		if _, err := io.ReadFull(p.reader, data); err != nil {
			return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
		}

		// Create element based on VR
		buf := buffer.NewMemory(data)
		elem, err := p.createElement(t, vrValue, buf)
		if err != nil {
			return nil, err
		}

		if err := ds.Add(elem); err != nil {
			return nil, fmt.Errorf("failed to add element %s to dataset: %w", t, err)
		}
	}

	return ds, nil
}

// setTransferSyntax sets the transfer syntax from File Meta Information.
func (p *parseContext) setTransferSyntax(metaDS *dataset.Dataset) error {
	tsUID, exists := metaDS.GetString(tag.TransferSyntaxUID)
	if !exists {
		// Default to Explicit VR Little Endian
		p.transferSyntax = transfer.ExplicitVRLittleEndian
		p.byteOrder = binary.LittleEndian
		p.isExplicitVR = true
		return nil
	}

	// Look up transfer syntax
	ts, err := transfer.Parse(tsUID)
	if err != nil {
		return fmt.Errorf("failed to parse transfer syntax UID %s: %w", tsUID, err)
	}

	p.transferSyntax = ts
	p.isExplicitVR = ts.IsExplicitVR()

	if ts.Endian() == endian.Big {
		p.byteOrder = binary.BigEndian
	} else {
		p.byteOrder = binary.LittleEndian
	}

	return nil
}

// readDataset reads a dataset (collection of elements).
func (p *parseContext) readDataset() (*dataset.Dataset, error) {
	ds := dataset.New()

	// Check if we have a saved tag from readFileMetaInformation
	var firstTag *tag.Tag
	if p.firstDatasetTag != nil {
		firstTag = p.firstDatasetTag
		p.firstDatasetTag = nil // Clear it
	}

	for {
		// Check stop condition
		// TODO: Implement stopAtTag functionality for early termination
		_ = p.stopAtTag

		var elem element.Element
		var err error

		// If we have a saved first tag, use it for the first element
		if firstTag != nil {
			elem, err = p.readElementWithTag(firstTag)
			firstTag = nil // Only use it once
		} else {
			elem, err = p.readElement()
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Check if we should stop
		if p.stopAtTag != nil && elem.Tag().ToUint32() >= p.stopAtTag.ToUint32() {
			break
		}

		if err := ds.Add(elem); err != nil {
			return nil, fmt.Errorf("failed to add element to dataset: %w", err)
		}
	}

	return ds, nil
}

// readElementData handles reading element data based on size and read options.
// This is a common helper to avoid code duplication between readElement and readElementWithTag.
func (p *parseContext) readElementData(t *tag.Tag, _ *vr.VR, length uint32) (buffer.ByteBuffer, error) {
	// Handle large objects based on ReadOption
	isLarge := length > p.largeObjectSize

	if isLarge {
		switch p.readOption {
		case SkipLargeTags:
			// Skip the data entirely by discarding it
			if _, err := io.CopyN(io.Discard, p.reader, int64(length)); err != nil {
				return nil, fmt.Errorf("failed to skip large tag %s: %w", t, err)
			}
			// Return an empty buffer
			return buffer.NewMemory([]byte{}), nil

		case ReadLargeOnDemand:
			// Implement lazy loading if possible
			buf, err := p.createLazyBuffer(length)
			if err != nil {
				// If lazy loading not possible, fall back to reading all
				data := make([]byte, length)
				if _, err := io.ReadFull(p.reader, data); err != nil {
					return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
				}
				buf = buffer.NewMemory(data)
			}
			return buf, nil

		case ReadAll, ReadDefault:
			// Read the data normally
			data := make([]byte, length)
			if _, err := io.ReadFull(p.reader, data); err != nil {
				return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
			}
			return buffer.NewMemory(data), nil
		}
	}

	// Read value data normally for non-large elements
	data := make([]byte, length)
	if _, err := io.ReadFull(p.reader, data); err != nil {
		return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
	}

	return buffer.NewMemory(data), nil
}

// readElement reads a single DICOM element.
func (p *parseContext) readElement() (element.Element, error) {
	// Read tag (4 bytes: group + element)
	t, err := p.readTag()
	if err != nil {
		return nil, err
	}

	// Read VR (Value Representation)
	vrValue, err := p.readVR(t)
	if err != nil {
		return nil, fmt.Errorf("failed to read VR for tag %s: %w", t, err)
	}

	// Read value length
	length, err := p.readLength(vrValue)
	if err != nil {
		return nil, fmt.Errorf("failed to read length for tag %s: %w", t, err)
	}

	// Check max element size
	if p.maxElementSize > 0 && length > p.maxElementSize && length != 0xFFFFFFFF {
		return nil, fmt.Errorf("element size %d exceeds maximum %d for tag %s", length, p.maxElementSize, t)
	}

	// Handle special case: Sequence
	if vrValue.Code() == vr.CodeSQ {
		return p.readSequence(t, length)
	}

	// Handle special case: Fragment Sequence (encapsulated pixel data)
	// Fragment sequences have OB or OW VR with undefined length (0xFFFFFFFF)
	if (vrValue.Code() == vr.CodeOB || vrValue.Code() == vr.CodeOW) && length == 0xFFFFFFFF {
		return p.readFragmentSequence(t, vrValue)
	}

	// Read element data
	buf, err := p.readElementData(t, vrValue, length)
	if err != nil {
		return nil, err
	}

	// Create element based on VR
	return p.createElement(t, vrValue, buf)
}

// readElementWithTag reads a single DICOM element when the tag is already known.
// This is used when we've already read the tag (e.g., in readFileMetaInformation).
func (p *parseContext) readElementWithTag(t *tag.Tag) (element.Element, error) {
	// Read VR (Value Representation)
	vrValue, err := p.readVR(t)
	if err != nil {
		return nil, fmt.Errorf("failed to read VR for tag %s: %w", t, err)
	}

	// Read value length
	length, err := p.readLength(vrValue)
	if err != nil {
		return nil, fmt.Errorf("failed to read length for tag %s: %w", t, err)
	}

	// Check max element size
	if p.maxElementSize > 0 && length > p.maxElementSize && length != 0xFFFFFFFF {
		return nil, fmt.Errorf("element size %d exceeds maximum %d for tag %s", length, p.maxElementSize, t)
	}

	// Handle special case: Sequence
	if vrValue.Code() == vr.CodeSQ {
		return p.readSequence(t, length)
	}

	// Handle special case: Fragment Sequence (encapsulated pixel data)
	// Fragment sequences have OB or OW VR with undefined length (0xFFFFFFFF)
	if (vrValue.Code() == vr.CodeOB || vrValue.Code() == vr.CodeOW) && length == 0xFFFFFFFF {
		return p.readFragmentSequence(t, vrValue)
	}

	// Read element data
	buf, err := p.readElementData(t, vrValue, length)
	if err != nil {
		return nil, err
	}

	// Create element based on VR
	return p.createElement(t, vrValue, buf)
}

// readTag reads a DICOM tag (4 bytes).
func (p *parseContext) readTag() (*tag.Tag, error) {
	var group, elem uint16

	if err := binary.Read(p.reader, p.byteOrder, &group); err != nil {
		return nil, err
	}
	if err := binary.Read(p.reader, p.byteOrder, &elem); err != nil {
		return nil, err
	}

	return tag.New(group, elem), nil
}

// readVR reads the Value Representation.
func (p *parseContext) readVR(t *tag.Tag) (*vr.VR, error) {
	if p.isExplicitVR {
		// Read 2-byte VR code
		vrBytes := make([]byte, 2)
		if _, err := io.ReadFull(p.reader, vrBytes); err != nil {
			return nil, err
		}
		return vr.ParseBytes(vrBytes)
	}

	// Implicit VR: look up in dictionary
	if p.dictionary != nil {
		entry := p.dictionary.Lookup(t)
		if entry != nil {
			vrs := entry.ValueRepresentations()
			if len(vrs) > 0 {
				// Return the first VR (most common case has only one VR)
				return vrs[0], nil
			}
		}
	}

	// If no dictionary or tag not found, return UN (Unknown)
	return vr.UN, nil
}

// readLength reads the value length field.
func (p *parseContext) readLength(v *vr.VR) (uint32, error) {
	if p.isExplicitVR {
		// Check if VR has 16-bit or 32-bit length
		if v.Is16bitLength() {
			// 16-bit length
			var length uint16
			if err := binary.Read(p.reader, p.byteOrder, &length); err != nil {
				return 0, err
			}
			return uint32(length), nil
		}

		// 32-bit length: skip 2 reserved bytes first
		reserved := make([]byte, 2)
		if _, err := io.ReadFull(p.reader, reserved); err != nil {
			return 0, err
		}

		var length uint32
		if err := binary.Read(p.reader, p.byteOrder, &length); err != nil {
			return 0, err
		}
		return length, nil
	}

	// Implicit VR: always 32-bit length
	var length uint32
	if err := binary.Read(p.reader, p.byteOrder, &length); err != nil {
		return 0, err
	}
	return length, nil
}

// readSequence reads a sequence element (VR=SQ).
func (p *parseContext) readSequence(t *tag.Tag, length uint32) (*dataset.Sequence, error) {
	seq := dataset.NewSequence(t)

	// Handle undefined length (0xFFFFFFFF)
	if length == 0xFFFFFFFF {
		// Read items until we hit Sequence Delimitation Item
		for {
			itemTag, err := p.readTag()
			if err != nil {
				return nil, err
			}

			// Check for Sequence Delimitation Item (FFFE,E0DD)
			if itemTag.Group() == 0xFFFE && itemTag.Element() == 0xE0DD {
				// Read and discard length
				var delimitLength uint32
				_ = binary.Read(p.reader, p.byteOrder, &delimitLength)
				break
			}

			// Should be Item tag (FFFE,E000)
			if itemTag.Group() != 0xFFFE || itemTag.Element() != 0xE000 {
				return nil, fmt.Errorf("expected Item tag, got %s", itemTag)
			}

			// Read item length
			var itemLength uint32
			if err := binary.Read(p.reader, p.byteOrder, &itemLength); err != nil {
				return nil, err
			}

			// Read item dataset
			item, err := p.readItemDataset(itemLength)
			if err != nil {
				return nil, err
			}

			seq.AddItem(item)
		}
	} else {
		// Defined length sequence
		// Read exactly 'length' bytes and parse items within
		sequenceData := make([]byte, length)
		if _, err := io.ReadFull(p.reader, sequenceData); err != nil {
			return nil, fmt.Errorf("failed to read sequence data: %w", err)
		}

		// Create a reader for the sequence data
		seqReader := bytes.NewReader(sequenceData)
		originalReader := p.reader
		p.reader = seqReader

		// Read items until we've consumed all sequence data
		for seqReader.Len() > 0 {
			itemTag, err := p.readTag()
			if err == io.EOF {
				break
			}
			if err != nil {
				p.reader = originalReader
				return nil, err
			}

			// Should be Item tag (FFFE,E000)
			if itemTag.Group() != 0xFFFE || itemTag.Element() != 0xE000 {
				p.reader = originalReader
				return nil, fmt.Errorf("expected Item tag in sequence, got %s", itemTag)
			}

			// Read item length
			var itemLength uint32
			if err := binary.Read(p.reader, p.byteOrder, &itemLength); err != nil {
				p.reader = originalReader
				return nil, err
			}

			// Read item dataset
			item, err := p.readItemDataset(itemLength)
			if err != nil {
				p.reader = originalReader
				return nil, err
			}

			seq.AddItem(item)
		}

		// Restore original reader
		p.reader = originalReader
	}

	return seq, nil
}

// readItemDataset reads a single item dataset within a sequence.
func (p *parseContext) readItemDataset(length uint32) (*dataset.Dataset, error) {
	item := dataset.New()

	if length == 0xFFFFFFFF {
		// Undefined length item
		for {
			// Peek at the tag to check for Item Delimitation Item
			itemTag, err := p.readTag()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			// Check for Item Delimitation Item (FFFE,E00D)
			if itemTag.Group() == 0xFFFE && itemTag.Element() == 0xE00D {
				// Read and discard the length (should be 0)
				var delimitLength uint32
				_ = binary.Read(p.reader, p.byteOrder, &delimitLength)
				break
			}

			// Read the rest of the element
			elem, err := p.readElementWithTag(itemTag)
			if err != nil {
				return nil, err
			}

			if err := item.Add(elem); err != nil {
				return nil, fmt.Errorf("failed to add element to item: %w", err)
			}
		}
	} else {
		// Defined length item
		// Read exactly 'length' bytes and parse elements within
		itemData := make([]byte, length)
		if _, err := io.ReadFull(p.reader, itemData); err != nil {
			return nil, fmt.Errorf("failed to read item data: %w", err)
		}

		// Create a reader for the item data
		itemReader := bytes.NewReader(itemData)
		originalReader := p.reader
		p.reader = itemReader

		// Read elements until we've consumed all item data
		for itemReader.Len() > 0 {
			elem, err := p.readElement()
			if err == io.EOF {
				break
			}
			if err != nil {
				p.reader = originalReader
				return nil, err
			}

			if err := item.Add(elem); err != nil {
				p.reader = originalReader
				return nil, fmt.Errorf("failed to add element to item: %w", err)
			}
		}

		// Restore original reader
		p.reader = originalReader
	}

	return item, nil
}

// createElement creates an element from tag, VR, and buffer.
func (p *parseContext) createElement(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer) (element.Element, error) {
	// Create appropriate element type based on VR code
	vrCode := v.Code()
	switch vrCode {
	case vr.CodeAE, vr.CodeAS, vr.CodeCS, vr.CodeDA, vr.CodeDS, vr.CodeDT,
		vr.CodeIS, vr.CodeLO, vr.CodeLT, vr.CodePN, vr.CodeSH, vr.CodeST,
		vr.CodeTM, vr.CodeUC, vr.CodeUI, vr.CodeUR, vr.CodeUT:
		return element.NewStringFromBuffer(t, v, buf, nil), nil

	case vr.CodeUS:
		return element.NewUnsignedShortFromBuffer(t, buf), nil
	case vr.CodeUL:
		return element.NewUnsignedLongFromBuffer(t, buf), nil
	case vr.CodeSS:
		return element.NewSignedShortFromBuffer(t, buf), nil
	case vr.CodeSL:
		return element.NewSignedLongFromBuffer(t, buf), nil
	case vr.CodeFL:
		return element.NewFloatFromBuffer(t, buf), nil
	case vr.CodeFD:
		return element.NewDoubleFromBuffer(t, buf), nil

	case vr.CodeOB:
		return element.NewOtherByteFromBuffer(t, buf), nil
	case vr.CodeOW:
		return element.NewOtherWordFromBuffer(t, buf), nil
	case vr.CodeUN:
		return element.NewUnknownFromBuffer(t, buf), nil

	default:
		// Default to Unknown
		return element.NewUnknownFromBuffer(t, buf), nil
	}
}

// createLazyBuffer creates a lazy-loading buffer for large elements.
// Returns an error if lazy loading is not possible (reader is not seekable).
//
// Strategy:
//   - If reader is *os.File: Use FileByteBuffer for efficient file-based access
//   - If reader is io.ReadSeeker: Use LazyByteBuffer with a loader function
//   - Otherwise: Return error (caller should fall back to reading all data)
func (p *parseContext) createLazyBuffer(length uint32) (buffer.ByteBuffer, error) {
	// Strategy 1: If we have a file, use FileByteBuffer
	if p.file != nil {
		// Get current position in file
		currentPos, err := p.file.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to get current file position: %w", err)
		}

		// Create a FileByteBuffer for this range
		fb, err := buffer.NewFile(p.file.Name(), uint32(currentPos), length)
		if err != nil {
			return nil, fmt.Errorf("failed to create file buffer: %w", err)
		}

		// Skip past this data in the stream
		if _, err := p.file.Seek(int64(length), io.SeekCurrent); err != nil {
			return nil, fmt.Errorf("failed to skip data: %w", err)
		}

		return fb, nil
	}

	// Strategy 2: If we have a seekable reader, use LazyByteBuffer
	if p.seekableReader != nil {
		// Get current position
		currentPos, err := p.seekableReader.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to get current position: %w", err)
		}

		// Create a loader function that will read the data when needed
		loader := func() []byte {
			// Save current position
			savedPos, _ := p.seekableReader.Seek(0, io.SeekCurrent)

			// Seek to the data position
			if _, err := p.seekableReader.Seek(currentPos, io.SeekStart); err != nil {
				return nil
			}

			// Read the data
			data := make([]byte, length)
			if _, err := io.ReadFull(p.seekableReader, data); err != nil {
				return nil
			}

			// Restore position
			_, _ = p.seekableReader.Seek(savedPos, io.SeekStart)

			return data
		}

		lb, err := buffer.NewLazy(loader)
		if err != nil {
			return nil, fmt.Errorf("failed to create lazy buffer: %w", err)
		}

		// Skip past this data in the stream
		if _, err := p.seekableReader.Seek(int64(length), io.SeekCurrent); err != nil {
			return nil, fmt.Errorf("failed to skip data: %w", err)
		}

		return lb, nil
	}

	// No seekable reader available
	return nil, fmt.Errorf("lazy loading not supported for non-seekable readers")
}

// readFragmentSequence reads a DICOM fragment sequence (encapsulated pixel data).
// Fragment sequences are used for compressed image formats like JPEG, JPEG 2000, RLE, etc.
//
// Structure:
// - First item (FFFE,E000): Offset Table (can be empty)
// - Subsequent items: Compressed frame fragments
// - End marker (FFFE,E0DD): Sequence Delimitation Item
func (p *parseContext) readFragmentSequence(t *tag.Tag, vrValue *vr.VR) (element.Element, error) {
	var fs element.Element

	// Create appropriate fragment sequence type based on VR
	if vrValue.Code() == vr.CodeOB {
		fs = element.NewOtherByteFragment(t)
	} else if vrValue.Code() == vr.CodeOW {
		fs = element.NewOtherWordFragment(t)
	} else {
		return nil, fmt.Errorf("invalid VR for fragment sequence: %v", vrValue)
	}

	// Get the underlying FragmentSequence
	var fragSeq *element.FragmentSequence
	switch v := fs.(type) {
	case *element.OtherByteFragment:
		fragSeq = v.FragmentSequence
	case *element.OtherWordFragment:
		fragSeq = v.FragmentSequence
	default:
		return nil, fmt.Errorf("unexpected fragment sequence type")
	}

	// Read fragments until we hit Sequence Delimitation Item (FFFE,E0DD)
	isFirstItem := true

	for {
		// Read item tag (should be FFFE,E000 for item or FFFE,E0DD for delimitation)
		itemTag, err := p.readTag()
		if err != nil {
			if err == io.EOF {
				// EOF without proper delimitation - some files are like this
				return fs, nil
			}
			return nil, fmt.Errorf("failed to read fragment item tag: %w", err)
		}

		// Check if this is Sequence Delimitation Item (FFFE,E0DD)
		if itemTag.Group() == 0xFFFE && itemTag.Element() == 0xE0DD {
			// Read and discard the length (should be 0)
			var delimLength uint32
			if err := binary.Read(p.reader, p.byteOrder, &delimLength); err != nil {
				return nil, fmt.Errorf("failed to read sequence delimitation length: %w", err)
			}
			// End of fragment sequence
			break
		}

		// Should be Item tag (FFFE,E000)
		if itemTag.Group() != 0xFFFE || itemTag.Element() != 0xE000 {
			return nil, fmt.Errorf("expected Item tag (FFFE,E000) in fragment sequence, got %s", itemTag)
		}

		// Read item length
		var itemLength uint32
		if err := binary.Read(p.reader, p.byteOrder, &itemLength); err != nil {
			return nil, fmt.Errorf("failed to read fragment item length: %w", err)
		}

		// Read item data
		itemData := make([]byte, itemLength)
		if _, err := io.ReadFull(p.reader, itemData); err != nil {
			return nil, fmt.Errorf("failed to read fragment item data: %w", err)
		}

		if isFirstItem {
			// First item is the offset table
			// Parse it as uint32 array if not empty
			if itemLength > 0 {
				if itemLength%4 != 0 {
					return nil, fmt.Errorf("offset table length %d is not a multiple of 4", itemLength)
				}

				numOffsets := itemLength / 4
				offsets := make([]uint32, numOffsets)

				for i := uint32(0); i < numOffsets; i++ {
					offset := p.byteOrder.Uint32(itemData[i*4 : (i+1)*4])
					offsets[i] = offset
				}

				fragSeq.SetOffsetTable(offsets)
			}
			isFirstItem = false
		} else {
			// Subsequent items are fragments
			fragBuf := buffer.NewMemory(itemData)
			fragSeq.AddFragment(fragBuf)
		}
	}

	return fs, nil
}
