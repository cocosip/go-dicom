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

// ParseResult represents the result of parsing a DICOM file.
//
// It contains both the File Meta Information (Group 0002) and the main Dataset.
// This structure mimics fo-dicom's DicomFile concept.
type ParseResult struct {
	// FileMetaInformation contains Group 0002 elements.
	// These elements describe the file format and transfer syntax.
	// Always encoded as Explicit VR Little Endian.
	FileMetaInformation *dataset.Dataset

	// Dataset contains the main DICOM data elements.
	// Encoding depends on the Transfer Syntax specified in FileMetaInformation.
	Dataset *dataset.Dataset

	// TransferSyntax specifies how the dataset is encoded.
	TransferSyntax *transfer.TransferSyntax
}

// Parser reads and parses DICOM files.
type Parser struct {
	reader         io.Reader
	byteOrder      binary.ByteOrder
	isExplicitVR   bool
	transferSyntax *transfer.TransferSyntax
	dictionary     *dict.Dictionary

	// firstDatasetTag is the first tag read that doesn't belong to Group 0002
	// It's stored here to be used when reading the main dataset
	firstDatasetTag *tag.Tag

	// Options
	maxElementSize uint32 // Maximum element size to read (0 = unlimited)
	stopAtTag      *tag.Tag
}

// Option is a functional option for Parser.
type Option func(*Parser)

// WithMaxElementSize sets the maximum element size to read.
func WithMaxElementSize(size uint32) Option {
	return func(p *Parser) {
		p.maxElementSize = size
	}
}

// WithStopAtTag sets a tag to stop parsing at.
func WithStopAtTag(t *tag.Tag) Option {
	return func(p *Parser) {
		p.stopAtTag = t
	}
}

// WithDictionary sets the DICOM dictionary for implicit VR lookup.
func WithDictionary(d *dict.Dictionary) Option {
	return func(p *Parser) {
		p.dictionary = d
	}
}

// New creates a new Parser with the given options.
func New(opts ...Option) *Parser {
	p := &Parser{
		byteOrder:    binary.LittleEndian,
		isExplicitVR: true,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Parse parses a DICOM file from the reader.
// This is the main entry point for reading DICOM files.
//
// Returns a ParseResult containing:
//   - FileMetaInformation: Group 0002 elements
//   - Dataset: Main DICOM data
//   - TransferSyntax: How the dataset is encoded
func Parse(r io.Reader, opts ...Option) (*ParseResult, error) {
	p := New(opts...)
	return p.Parse(r)
}

// Parse parses a DICOM file from the reader.
func (p *Parser) Parse(r io.Reader) (*ParseResult, error) {
	p.reader = r

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

	return &ParseResult{
		FileMetaInformation: metaDS,
		Dataset:             mainDS,
		TransferSyntax:      p.transferSyntax,
	}, nil
}

// ParseFile parses a DICOM file from a file path.
func ParseFile(path string, opts ...Option) (*ParseResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return Parse(file, opts...)
}

// readPreamble reads and validates the 128-byte preamble and DICM prefix.
func (p *Parser) readPreamble() error {
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
func (p *Parser) readFileMetaInformation() (*dataset.Dataset, error) {
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

		ds.Add(elem)
	}

	return ds, nil
}

// setTransferSyntax sets the transfer syntax from File Meta Information.
func (p *Parser) setTransferSyntax(metaDS *dataset.Dataset) error {
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
func (p *Parser) readDataset() (*dataset.Dataset, error) {
	ds := dataset.New()

	// Check if we have a saved tag from readFileMetaInformation
	var firstTag *tag.Tag
	if p.firstDatasetTag != nil {
		firstTag = p.firstDatasetTag
		p.firstDatasetTag = nil // Clear it
	}

	for {
		// Check stop condition
		if p.stopAtTag != nil {
			// We need to peek at the next tag
			// For now, we'll read and check
		}

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

		ds.Add(elem)
	}

	return ds, nil
}

// readElement reads a single DICOM element.
func (p *Parser) readElement() (element.Element, error) {
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
	if p.maxElementSize > 0 && length > p.maxElementSize {
		return nil, fmt.Errorf("element size %d exceeds maximum %d for tag %s", length, p.maxElementSize, t)
	}

	// Handle special case: Sequence
	if vrValue.Code() == vr.CodeSQ {
		return p.readSequence(t, length)
	}

	// Read value data
	data := make([]byte, length)
	if _, err := io.ReadFull(p.reader, data); err != nil {
		return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
	}

	// Create element based on VR
	buf := buffer.NewMemory(data)
	return p.createElement(t, vrValue, buf)
}

// readElementWithTag reads a single DICOM element when the tag is already known.
// This is used when we've already read the tag (e.g., in readFileMetaInformation).
func (p *Parser) readElementWithTag(t *tag.Tag) (element.Element, error) {
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
	if p.maxElementSize > 0 && length > p.maxElementSize {
		return nil, fmt.Errorf("element size %d exceeds maximum %d for tag %s", length, p.maxElementSize, t)
	}

	// Handle special case: Sequence
	if vrValue.Code() == vr.CodeSQ {
		return p.readSequence(t, length)
	}

	// Read value data
	data := make([]byte, length)
	if _, err := io.ReadFull(p.reader, data); err != nil {
		return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
	}

	// Create element based on VR
	buf := buffer.NewMemory(data)
	return p.createElement(t, vrValue, buf)
}

// readTag reads a DICOM tag (4 bytes).
func (p *Parser) readTag() (*tag.Tag, error) {
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
func (p *Parser) readVR(t *tag.Tag) (*vr.VR, error) {
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
func (p *Parser) readLength(v *vr.VR) (uint32, error) {
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
func (p *Parser) readSequence(t *tag.Tag, length uint32) (*dataset.Sequence, error) {
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
				binary.Read(p.reader, p.byteOrder, &delimitLength)
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
		for {
			// Check if we've read all the data
			if seqReader.Len() == 0 {
				break
			}

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
func (p *Parser) readItemDataset(length uint32) (*dataset.Dataset, error) {
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
				binary.Read(p.reader, p.byteOrder, &delimitLength)
				break
			}

			// Read the rest of the element
			elem, err := p.readElementWithTag(itemTag)
			if err != nil {
				return nil, err
			}

			item.Add(elem)
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
		for {
			// Check if we've read all the data
			if itemReader.Len() == 0 {
				break
			}

			elem, err := p.readElement()
			if err == io.EOF {
				break
			}
			if err != nil {
				p.reader = originalReader
				return nil, err
			}

			item.Add(elem)
		}

		// Restore original reader
		p.reader = originalReader
	}

	return item, nil
}

// createElement creates an element from tag, VR, and buffer.
func (p *Parser) createElement(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer) (element.Element, error) {
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
