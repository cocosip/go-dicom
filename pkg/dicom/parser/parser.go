// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"compress/flate"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/text/encoding"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
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
	// ReadDefault reads all tags normally. Large elements may use lazy loading
	// when the underlying stream is seekable, deferring data reads until accessed.
	ReadDefault ReadOption = iota

	// ReadLargeOnDemand reads small tags immediately but keeps the stream open
	// to read large tags on demand. The stream must stay open.
	ReadLargeOnDemand

	// SkipLargeTags skips reading large tags entirely. The stream can be closed.
	SkipLargeTags

	// ReadAll forces eager loading of all elements including large ones.
	// LazyByteBuffer is never created; all data is loaded into memory during parsing.
	ReadAll

	// maxSequenceItems is the maximum number of items (fragments or sequence items)
	// allowed in an undefined-length sequence or fragment sequence. This prevents
	// infinite loops on malformed data missing the delimitation item.
	maxSequenceItems = 100000
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
	textEncoding   encoding.Encoding
	textEncodings  []encoding.Encoding

	// firstDatasetTagRaw holds the first tag read that doesn't belong to Group 0002.
	// Raw bytes are preserved because dataset byte order is only known after reading Transfer Syntax.
	firstDatasetTagRaw [4]byte
	hasFirstDatasetTag bool

	// Context for cancellation. Checked before blocking reads.
	ctx context.Context

	// Configuration options
	maxElementSize        uint32           // Maximum element size to read (default 500MB, 0 = unlimited)
	stopAtTag             *tag.Tag         // Stop parsing when this tag is reached
	readOption            ReadOption       // How to handle large elements
	largeObjectSize       uint32           // Size threshold for "large" objects (default 64KB)
	assumedTransferSyntax *transfer.Syntax // Transfer syntax to use for raw datasets without file meta
	sequenceItemObserver  SequenceItemObserver
	position              *readerPosition

	// File format detection
	detectedFormat FileFormat
	isPartial      bool

	// For lazy loading support
	seekableReader io.ReadSeeker // Set if reader is seekable (for lazy loading)
	file           *os.File      // Set if reader is a file (for FileByteBuffer)
	lazyReadMu     sync.Mutex    // Serializes seek/read operations for lazy loaders on shared readers
	datasetCloser  io.Closer     // Closes any wrapper reader created for the dataset stream
	boundedReaders []*io.LimitedReader
}

type privateCreatorScope map[uint32]*tag.PrivateCreator

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
//   - ReadDefault: Read all elements, using lazy loading for large ones when possible
//   - ReadLargeOnDemand: Read large elements on demand (stream must stay open)
//   - SkipLargeTags: Skip large elements entirely
//   - ReadAll: Force eager loading of all elements (no lazy buffers)
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

// WithAssumedTransferSyntax sets the transfer syntax for raw datasets that do not contain file meta information.
func WithAssumedTransferSyntax(ts *transfer.Syntax) Option {
	return func(ctx *parseContext) {
		ctx.assumedTransferSyntax = ts
	}
}

// WithContext sets the context for cancellation during parsing.
// The context is checked before each blocking read operation. If the context
// is cancelled, the parse is aborted with the context error.
// Defaults to context.Background() if not set.
func WithContext(parent context.Context) Option {
	return func(ctx *parseContext) {
		if parent == nil {
			parent = context.Background()
		}
		ctx.ctx = parent
	}
}

// newParseContext creates a new parse context with the given options.
func newParseContext(opts ...Option) *parseContext {
	ctx := &parseContext{
		ctx:             context.Background(),
		byteOrder:       binary.LittleEndian,
		isExplicitVR:    true,
		textEncoding:    charset.Default,
		textEncodings:   []encoding.Encoding{charset.Default},
		readOption:      ReadDefault,
		largeObjectSize: 65536,     // Default 64KB
		maxElementSize:  524288000, // Default 500MB (prevent runaway allocation)
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

// ctxReader wraps an io.Reader to check context cancellation before every read.
type ctxReader struct {
	ctx context.Context
	r   io.Reader
}

func (c *ctxReader) Read(p []byte) (int, error) {
	if err := c.ctx.Err(); err != nil {
		return 0, err
	}
	return c.r.Read(p)
}

type ctxReadSeeker struct {
	ctx context.Context
	rs  io.ReadSeeker
}

type readerPosition struct {
	offset uint64
}

type positionReader struct {
	r        io.Reader
	position *readerPosition
}

func (r *positionReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.position.offset += uint64(n) // #nosec G115 -- Read never returns a negative count
	return n, err
}

type positionReadSeeker struct {
	rs       io.ReadSeeker
	position *readerPosition
}

func (r *positionReadSeeker) Read(p []byte) (int, error) {
	n, err := r.rs.Read(p)
	r.position.offset += uint64(n) // #nosec G115 -- Read never returns a negative count
	return n, err
}

func (r *positionReadSeeker) Seek(offset int64, whence int) (int64, error) {
	position, err := r.rs.Seek(offset, whence)
	if err != nil {
		return 0, err
	}
	if position < 0 {
		return 0, fmt.Errorf("reader returned negative position %d", position)
	}
	r.position.offset = uint64(position)
	return position, nil
}

func (c *ctxReadSeeker) Read(p []byte) (int, error) {
	if err := c.ctx.Err(); err != nil {
		return 0, err
	}
	return c.rs.Read(p)
}

func (c *ctxReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return c.rs.Seek(offset, whence)
}

// Parse parses a DICOM file from the reader.
// This is the main entry point for reading DICOM files.
//
// Use WithContext to provide a context for cancellation:
//
//	result, err := parser.Parse(reader,
//	    parser.WithContext(ctx),
//	    parser.WithReadOption(parser.SkipLargeTags),
//	)
//
// Returns a ParseResult containing:
//   - FileMetaInformation: Group 0002 elements
//   - Dataset: Main DICOM data
//   - TransferSyntax: How the dataset is encoded
//   - Format: Detected file format
//   - IsPartial: Whether parsing stopped early
func Parse(r io.Reader, opts ...Option) (*ParseResult, error) {
	pctx := newParseContext(opts...)
	return pctx.parse(r)
}

// parse is the internal parsing implementation.
func (p *parseContext) parse(r io.Reader) (*ParseResult, error) {
	// Wrap reader with context cancellation support so every Read checks ctx.Err().
	p.reader = p.contextAwareReader(r)
	p.detectedFormat = FormatUnknown

	// Check if reader supports seeking (for lazy loading)
	if rs, ok := r.(io.ReadSeeker); ok {
		p.seekableReader = p.contextAwareReadSeeker(rs)
	}

	// Check if reader is an *os.File (for FileByteBuffer)
	if f, ok := r.(*os.File); ok {
		p.file = f
	}

	// Detect preamble and position the stream at the first element.
	if err := p.detectFormatAndPrepareReader(); err != nil {
		return nil, fmt.Errorf("failed to detect file format: %w", err)
	}
	if p.sequenceItemObserver != nil {
		p.startPositionTracking()
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
	if p.datasetCloser != nil {
		defer func() { _ = p.datasetCloser.Close() }()
	}

	// Read main dataset with detected transfer syntax
	mainDS, err := p.readDataset()
	if err != nil {
		return nil, fmt.Errorf("failed to read dataset: %w", err)
	}
	mainDS.SetInternalTransferSyntax(p.transferSyntax)

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

func (p *parseContext) startPositionTracking() {
	position := &readerPosition{}
	if p.detectedFormat == FormatDICOM3 {
		position.offset = 132
	}

	if p.seekableReader != nil {
		tracked := &positionReadSeeker{rs: p.seekableReader, position: position}
		p.seekableReader = tracked
		p.reader = tracked
	} else {
		p.reader = &positionReader{r: p.reader, position: position}
	}
	p.position = position
}

func (p *parseContext) currentPosition() uint64 {
	if p.position == nil {
		return 0
	}
	return p.position.offset
}

func (p *parseContext) contextAwareReader(r io.Reader) io.Reader {
	if rs, ok := r.(io.ReadSeeker); ok {
		return p.contextAwareReadSeeker(rs)
	}
	return &ctxReader{ctx: p.ctx, r: r}
}

func (p *parseContext) contextAwareReadSeeker(rs io.ReadSeeker) io.ReadSeeker {
	return &ctxReadSeeker{ctx: p.ctx, rs: rs}
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

// detectFormatAndPrepareReader detects whether the stream starts with a DICOM preamble
// and makes sure p.reader is positioned at the first data element.
func (p *parseContext) detectFormatAndPrepareReader() error {
	header := make([]byte, 132)
	n, err := io.ReadFull(p.reader, header)
	if err == nil {
		if string(header[128:132]) == "DICM" {
			p.detectedFormat = FormatDICOM3
			return nil
		}

		// Full header was read but no DICM prefix: treat as no-preamble stream.
		if err := p.restoreReaderToStart(header[:n]); err != nil {
			return err
		}
		p.detectedFormat = FormatDICOM3NoPreamble
		return nil
	}

	if err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}

	if n == 0 {
		return io.ErrUnexpectedEOF
	}

	// Short streams can still be valid no-preamble datasets.
	if err := p.restoreReaderToStart(header[:n]); err != nil {
		return err
	}
	p.detectedFormat = FormatDICOM3NoPreamble
	return nil
}

// restoreReaderToStart restores reader position after probing the header.
// For seekable readers we seek back, otherwise we prepend the consumed bytes.
func (p *parseContext) restoreReaderToStart(consumed []byte) error {
	if p.seekableReader != nil {
		if _, err := p.seekableReader.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek reader back to start: %w", err)
		}
		p.reader = p.seekableReader
		return nil
	}

	p.reader = io.MultiReader(bytes.NewReader(consumed), p.reader)
	return nil
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
	ds.SetAutoValidate(false)
	defer ds.SetAutoValidate(true)

	// Read elements until we leave Group 0002
	for {
		// Peek at the tag first to check if we're still in Group 0002.
		// Keep raw bytes so the first dataset tag can be re-decoded with dataset byte order.
		t, rawTag, err := p.readTagWithRaw()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Stop when we leave Group 0002
		if t.Group() != 0x0002 {
			// Save raw bytes and parse this tag later with the dataset byte order.
			p.firstDatasetTagRaw = rawTag
			p.hasFirstDatasetTag = true
			if ds.Count() == 0 && (p.detectedFormat == FormatDICOM3 || p.detectedFormat == FormatDICOM3NoPreamble || p.detectedFormat == FormatUnknown) {
				p.detectedFormat = FormatDICOM3NoFileMetaInfo
			}
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
		p.updateTextEncoding(elem)
	}

	return ds, nil
}

// setTransferSyntax sets the transfer syntax from File Meta Information.
func (p *parseContext) setTransferSyntax(metaDS *dataset.Dataset) error {
	tsUID, exists := metaDS.GetString(tag.TransferSyntaxUID)
	if !exists {
		if p.detectedFormat == FormatDICOM3NoFileMetaInfo {
			if p.assumedTransferSyntax == nil {
				return fmt.Errorf("raw dataset without file meta requires an assumed transfer syntax")
			}
			return p.applyTransferSyntax(p.assumedTransferSyntax)
		}

		// Default to Explicit VR Little Endian
		return p.applyTransferSyntax(transfer.ExplicitVRLittleEndian)
	}

	// Look up transfer syntax
	ts, err := transfer.Parse(tsUID)
	if err != nil {
		return fmt.Errorf("failed to parse transfer syntax UID %s: %w", tsUID, err)
	}

	return p.applyTransferSyntax(ts)
}

func (p *parseContext) applyTransferSyntax(ts *transfer.Syntax) error {
	if ts == nil {
		return fmt.Errorf("transfer syntax cannot be nil")
	}

	p.transferSyntax = ts
	p.isExplicitVR = ts.IsExplicitVR()

	if ts.Endian() == endian.Big {
		p.byteOrder = binary.BigEndian
	} else {
		p.byteOrder = binary.LittleEndian
	}

	if ts.IsDeflate() {
		if p.sequenceItemObserver != nil {
			return fmt.Errorf("sequence item positions are unavailable for deflated transfer syntax")
		}
		if p.hasFirstDatasetTag {
			p.reader = io.MultiReader(bytes.NewReader(p.firstDatasetTagRaw[:]), p.reader)
			p.hasFirstDatasetTag = false
		}
		fr := flate.NewReader(p.reader)
		p.reader = p.contextAwareReader(fr)
		p.datasetCloser = fr
		p.seekableReader = nil
		p.file = nil
	}

	return nil
}

// readDataset reads a dataset (collection of elements).
func (p *parseContext) readDataset() (*dataset.Dataset, error) {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	defer ds.SetAutoValidate(true)
	privateCreators := make(privateCreatorScope)

	var firstTag *tag.Tag
	if p.hasFirstDatasetTag {
		firstTag = decodeTag(p.firstDatasetTagRaw, p.byteOrder)
		p.hasFirstDatasetTag = false
	}

	for {
		// Read the next tag first so stopAtTag can stop before reading value bytes.
		var t *tag.Tag
		var err error
		if firstTag != nil {
			t = firstTag
			firstTag = nil
		} else {
			t, err = p.readTag()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if t.Group() == 0 && t.Element() == 0 {
			if err := p.consumeZeroTrailingPadding(); err != nil {
				return nil, err
			}
			break
		}

		// Stop before reading this element's value payload.
		if p.stopAtTag != nil && t.ToUint32() >= p.stopAtTag.ToUint32() {
			p.isPartial = true
			break
		}

		p.attachPrivateCreator(t, privateCreators)
		elem, err := p.readElementWithTag(t)
		if err != nil {
			return nil, err
		}

		if err := ds.Add(elem); err != nil {
			return nil, fmt.Errorf("failed to add element to dataset: %w", err)
		}
		p.updatePrivateCreator(elem, privateCreators)
		p.updateTextEncoding(elem)
	}

	return ds, nil
}

// consumeZeroTrailingPadding accepts non-conformant zero padding after the dataset.
// A DICOM data element cannot use tag (0000,0000), so the tag can only represent padding here.
func (p *parseContext) consumeZeroTrailingPadding() error {
	buf := make([]byte, 32*1024)
	for {
		n, err := p.reader.Read(buf)
		for _, b := range buf[:n] {
			if b != 0 {
				return fmt.Errorf("non-zero data follows trailing padding tag (0000,0000)")
			}
		}

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read trailing padding: %w", err)
		}
	}
}

// readElementData handles reading element data based on size and read options.
// This is a common helper to avoid code duplication between readElement and readElementWithTag.
func (p *parseContext) readElementData(t *tag.Tag, _ *vr.VR, length uint32) (buffer.ByteBuffer, error) {
	// Guard against undefined length (0xFFFFFFFF) which should have been handled
	// by the caller for SQ/UN/OB/OW sequence and fragment paths. Reaching here
	// with undefined length indicates a malformed DICOM stream.
	if length == 0xFFFFFFFF {
		return nil, fmt.Errorf("element %s has undefined length, expected sequence or fragment handling", t)
	}

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

		case ReadAll:
			// Force eager loading even for large elements; never create lazy buffers.
			data := make([]byte, length)
			if _, err := io.ReadFull(p.reader, data); err != nil {
				return nil, fmt.Errorf("failed to read value data for tag %s: %w", t, err)
			}
			return buffer.NewMemory(data), nil

		case ReadDefault:
			buf, err := p.createLazyBuffer(length)
			if err == nil {
				return buf, nil
			}

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
	return p.readElementWithPrivateCreators(nil)
}

func (p *parseContext) readElementWithPrivateCreators(privateCreators privateCreatorScope) (element.Element, error) {
	// Read tag (4 bytes: group + element)
	t, err := p.readTag()
	if err != nil {
		return nil, err
	}
	p.attachPrivateCreator(t, privateCreators)

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

	// Handle special case: VR=UN with undefined length
	// In DICOM, private tags or unknown tags may be encoded as UN
	// When combined with undefined length, they are typically sequences
	// Treat UN with undefined length as a sequence
	if vrValue.Code() == vr.CodeUN && length == 0xFFFFFFFF {
		return p.readSequence(t, length)
	}

	// Handle special case: Fragment Sequence (encapsulated pixel data)
	// Fragment sequences have OB or OW VR with undefined length (0xFFFFFFFF)
	if (vrValue.Code() == vr.CodeOB || vrValue.Code() == vr.CodeOW) && length == 0xFFFFFFFF {
		if p.readOption == SkipLargeTags {
			if err := p.skipFragmentSequence(); err != nil {
				return nil, err
			}
			return createEmptyFragmentSequence(t, vrValue)
		}
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

	// Handle special case: VR=UN with undefined length
	// In DICOM, private tags or unknown tags may be encoded as UN
	// When combined with undefined length, they are typically sequences
	// Treat UN with undefined length as a sequence
	if vrValue.Code() == vr.CodeUN && length == 0xFFFFFFFF {
		return p.readSequence(t, length)
	}

	// Handle special case: Fragment Sequence (encapsulated pixel data)
	// Fragment sequences have OB or OW VR with undefined length (0xFFFFFFFF)
	if (vrValue.Code() == vr.CodeOB || vrValue.Code() == vr.CodeOW) && length == 0xFFFFFFFF {
		if p.readOption == SkipLargeTags {
			if err := p.skipFragmentSequence(); err != nil {
				return nil, err
			}
			return createEmptyFragmentSequence(t, vrValue)
		}
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
	t, _, err := p.readTagWithRaw()
	return t, err
}

// readTagWithRaw reads a DICOM tag and returns both decoded tag and raw bytes.
func (p *parseContext) readTagWithRaw() (*tag.Tag, [4]byte, error) {
	var raw [4]byte
	if _, err := io.ReadFull(p.reader, raw[:]); err != nil {
		return nil, raw, err
	}
	return decodeTag(raw, p.byteOrder), raw, nil
}

func decodeTag(raw [4]byte, order binary.ByteOrder) *tag.Tag {
	if order == nil {
		order = binary.LittleEndian
	}
	return tag.New(order.Uint16(raw[0:2]), order.Uint16(raw[2:4]))
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
	if isPrivateCreatorTag(t) {
		return vr.LO, nil
	}

	// Implicit VR: look up in dictionary
	// Use provided dictionary, or default dictionary if not provided
	dictionary := p.dictionary
	if dictionary == nil {
		dictionary = dict.Default()
	}

	entry := dictionary.Lookup(t)
	if entry != nil {
		vrs := entry.ValueRepresentations()
		if len(vrs) > 0 {
			// Return the first VR (most common case has only one VR)
			return vrs[0], nil
		}
	}

	// If tag not found in dictionary, return UN (Unknown)
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
		itemCount := 0
		for {
			if itemCount >= maxSequenceItems {
				return nil, fmt.Errorf("exceeded maximum sequence items (%d) for tag %s", maxSequenceItems, t)
			}
			itemCount++

			itemOffset := p.currentPosition()
			itemTag, err := p.readTag()
			if err != nil {
				return nil, err
			}

			// Check for Sequence Delimitation Item (FFFE,E0DD)
			if itemTag.Group() == 0xFFFE && itemTag.Element() == 0xE0DD {
				// Read and discard length
				var delimitLength uint32
				if err := binary.Read(p.reader, p.byteOrder, &delimitLength); err != nil {
					return nil, err
				}
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
			if err := p.observeSequenceItem(t, item, itemOffset); err != nil {
				return nil, err
			}
		}
	} else {
		// Defined-length sequence: parse directly from a bounded stream to avoid
		// allocating and copying the whole sequence payload.
		err := p.withBoundedReader(length, func(lr *io.LimitedReader) error {
			for lr.N > 0 {
				itemOffset := p.currentPosition()
				itemTag, err := p.readTag()
				if err != nil {
					return err
				}

				// Should be Item tag (FFFE,E000)
				if itemTag.Group() != 0xFFFE || itemTag.Element() != 0xE000 {
					return fmt.Errorf("expected Item tag in sequence, got %s", itemTag)
				}

				// Read item length
				var itemLength uint32
				if err := binary.Read(p.reader, p.byteOrder, &itemLength); err != nil {
					return err
				}

				// Read item dataset
				item, err := p.readItemDataset(itemLength)
				if err != nil {
					return err
				}

				seq.AddItem(item)
				if err := p.observeSequenceItem(t, item, itemOffset); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return seq, nil
}

func (p *parseContext) observeSequenceItem(sequenceTag *tag.Tag, item *dataset.Dataset, offset uint64) error {
	if p.sequenceItemObserver == nil {
		return nil
	}
	if err := p.sequenceItemObserver(SequenceItemPosition{
		SequenceTag: sequenceTag,
		Item:        item,
		Offset:      offset,
	}); err != nil {
		return fmt.Errorf("sequence item observer failed for %s at offset %d: %w", sequenceTag, offset, err)
	}
	return nil
}

// readItemDataset reads a single item dataset within a sequence.
func (p *parseContext) readItemDataset(length uint32) (*dataset.Dataset, error) {
	item := dataset.New()
	item.SetAutoValidate(false)
	defer item.SetAutoValidate(true)
	privateCreators := make(privateCreatorScope)
	savedEncoding := p.textEncoding
	savedEncodings := append([]encoding.Encoding(nil), p.textEncodings...)
	defer func() {
		p.textEncoding = savedEncoding
		p.textEncodings = savedEncodings
	}()

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
				if err := binary.Read(p.reader, p.byteOrder, &delimitLength); err != nil {
					return nil, err
				}
				break
			}

			// Read the rest of the element
			p.attachPrivateCreator(itemTag, privateCreators)
			elem, err := p.readElementWithTag(itemTag)
			if err != nil {
				return nil, err
			}

			if err := item.Add(elem); err != nil {
				return nil, fmt.Errorf("failed to add element to item: %w", err)
			}
			p.updatePrivateCreator(elem, privateCreators)
			p.updateTextEncoding(elem)
		}
	} else {
		// Defined-length item: parse directly from a bounded stream to avoid
		// allocating and copying the whole item payload.
		err := p.withBoundedReader(length, func(lr *io.LimitedReader) error {
			for lr.N > 0 {
				elem, err := p.readElementWithPrivateCreators(privateCreators)
				if err != nil {
					return err
				}

				if err := item.Add(elem); err != nil {
					return fmt.Errorf("failed to add element to item: %w", err)
				}
				p.updatePrivateCreator(elem, privateCreators)
				p.updateTextEncoding(elem)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return item, nil
}

func isPrivateCreatorTag(t *tag.Tag) bool {
	return t != nil && t.IsPrivate() && t.Element() >= 0x0010 && t.Element() <= 0x00FF
}

func privateCreatorScopeKey(group, block uint16) uint32 {
	return (uint32(group) << 16) | uint32(block)
}

func (p *parseContext) attachPrivateCreator(t *tag.Tag, creators privateCreatorScope) {
	if t == nil || !t.IsPrivate() || t.Element() < 0x1000 {
		return
	}
	block := t.Element() >> 8
	if creator := creators[privateCreatorScopeKey(t.Group(), block)]; creator != nil {
		t.SetPrivateCreator(creator)
	}
}

func (p *parseContext) updatePrivateCreator(elem element.Element, creators privateCreatorScope) {
	if elem == nil || !isPrivateCreatorTag(elem.Tag()) {
		return
	}
	stringElement, ok := elem.(*element.String)
	if !ok {
		return
	}

	key := privateCreatorScopeKey(elem.Tag().Group(), elem.Tag().Element())
	creatorName := stringElement.GetString()
	if creatorName == "" {
		delete(creators, key)
		return
	}
	dictionary := p.dictionary
	if dictionary == nil {
		dictionary = dict.Default()
	}
	creators[key] = dictionary.GetPrivateCreator(creatorName)
}

func (p *parseContext) updateTextEncoding(elem element.Element) {
	if elem == nil || elem.Tag().ToUint32() != tag.SpecificCharacterSet.ToUint32() {
		return
	}

	strElem, ok := elem.(*element.String)
	if !ok {
		return
	}

	values := strElem.GetValues()
	if len(values) == 0 {
		p.textEncoding = charset.Default
		p.textEncodings = []encoding.Encoding{charset.Default}
		return
	}
	p.textEncodings = charset.GetEncodings(values)
	p.textEncoding = p.textEncodings[0]
}

// withBoundedReader runs fn with p.reader limited to the next 'length' bytes.
// It restores parser state afterward and drains any unread bytes from the bound.
//
// Lazy-loading state (seekable reader/file) remains available inside the bounded
// scope. When lazy buffers seek past bytes instead of reading them, bounded
// readers are adjusted explicitly to keep outer sequence/item lengths aligned.
func (p *parseContext) withBoundedReader(length uint32, fn func(*io.LimitedReader) error) error {
	lr := &io.LimitedReader{R: p.reader, N: int64(length)}

	originalReader := p.reader

	p.reader = lr
	p.boundedReaders = append(p.boundedReaders, lr)

	defer func() {
		p.reader = originalReader
		p.boundedReaders = p.boundedReaders[:len(p.boundedReaders)-1]
	}()

	if err := fn(lr); err != nil {
		return err
	}

	// If the bounded parser didn't consume all bytes, discard the remainder
	// so the outer parser stays aligned at the next element.
	if lr.N > 0 {
		if _, err := io.CopyN(io.Discard, lr, lr.N); err != nil {
			return err
		}
	}

	return nil
}

func (p *parseContext) ensureBoundedAvailable(length uint32) error {
	if len(p.boundedReaders) == 0 {
		return nil
	}

	required := int64(length)
	for _, lr := range p.boundedReaders {
		if lr.N < required {
			return fmt.Errorf("lazy skip %d exceeds bounded reader remaining %d", length, lr.N)
		}
	}

	return nil
}

func (p *parseContext) consumeBounded(length uint32) {
	if len(p.boundedReaders) == 0 {
		return
	}

	consumed := int64(length)
	for _, lr := range p.boundedReaders {
		lr.N -= consumed
	}
}

// createElement creates an element from tag, VR, and buffer.
func (p *parseContext) createElement(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer) (element.Element, error) {
	setOrder := func(elem element.Element) element.Element {
		element.SetByteOrder(elem, p.byteOrder)
		return elem
	}

	// Create appropriate element type based on VR code
	vrCode := v.Code()
	switch vrCode {
	case vr.CodeAE, vr.CodeAS, vr.CodeCS, vr.CodeDA, vr.CodeDS, vr.CodeDT,
		vr.CodeIS, vr.CodeLO, vr.CodeLT, vr.CodePN, vr.CodeSH, vr.CodeST,
		vr.CodeTM, vr.CodeUC, vr.CodeUI, vr.CodeUR, vr.CodeUT:
		return setOrder(element.NewStringFromBufferWithEncodings(t, v, buf, p.textEncodings)), nil

	case vr.CodeUS:
		return setOrder(element.NewUnsignedShortFromBuffer(t, buf)), nil
	case vr.CodeUL:
		return setOrder(element.NewUnsignedLongFromBuffer(t, buf)), nil
	case vr.CodeSS:
		return setOrder(element.NewSignedShortFromBuffer(t, buf)), nil
	case vr.CodeSL:
		return setOrder(element.NewSignedLongFromBuffer(t, buf)), nil
	case vr.CodeFL:
		return setOrder(element.NewFloatFromBuffer(t, buf)), nil
	case vr.CodeFD:
		return setOrder(element.NewDoubleFromBuffer(t, buf)), nil

	case vr.CodeOB:
		return setOrder(element.NewOtherByteFromBuffer(t, buf)), nil
	case vr.CodeOW:
		return setOrder(element.NewOtherWordFromBuffer(t, buf)), nil
	case vr.CodeOD:
		return setOrder(element.NewOtherDoubleFromBuffer(t, buf)), nil
	case vr.CodeOF:
		return setOrder(element.NewOtherFloatFromBuffer(t, buf)), nil
	case vr.CodeOL:
		return setOrder(element.NewOtherLongFromBuffer(t, buf)), nil
	case vr.CodeOV:
		return setOrder(element.NewOtherVeryLongFromBuffer(t, buf)), nil
	case vr.CodeAT:
		return setOrder(element.NewAttributeTagFromBuffer(t, buf)), nil
	case vr.CodeUN:
		return setOrder(element.NewUnknownFromBuffer(t, buf)), nil

	default:
		// Default to Unknown
		return setOrder(element.NewUnknownFromBuffer(t, buf)), nil
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
		if err := p.ensureBoundedAvailable(length); err != nil {
			return nil, err
		}

		// Get current position in file
		currentPos, err := p.file.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to get current file position: %w", err)
		}

		// Create a FileByteBuffer for this range
		fb, err := buffer.NewFileAtWithContext(p.ctx, p.file.Name(), currentPos, length)
		if err != nil {
			return nil, fmt.Errorf("failed to create file buffer: %w", err)
		}

		// Skip past this data in the stream
		if _, err := p.file.Seek(int64(length), io.SeekCurrent); err != nil {
			return nil, fmt.Errorf("failed to skip data: %w", err)
		}
		p.consumeBounded(length)

		return fb, nil
	}

	// Strategy 2: If we have a seekable reader, use LazyByteBuffer
	if p.seekableReader != nil {
		if err := p.ensureBoundedAvailable(length); err != nil {
			return nil, err
		}

		// Get current position
		currentPos, err := p.seekableReader.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to get current position: %w", err)
		}

		// Create a loader function that will read the data when needed
		loader := func() ([]byte, error) {
			p.lazyReadMu.Lock()
			defer p.lazyReadMu.Unlock()
			if p.seekableReader == nil {
				return nil, fmt.Errorf("seekable reader unavailable: previous lazy-load restore seek failed")
			}

			// Save current position
			savedPos, err := p.seekableReader.Seek(0, io.SeekCurrent)
			if err != nil {
				return nil, fmt.Errorf("failed to get current position: %w", err)
			}

			// Seek to the data position
			if _, err := p.seekableReader.Seek(currentPos, io.SeekStart); err != nil {
				return nil, fmt.Errorf("failed to seek to lazy data position: %w", err)
			}
			defer func() {
				if _, err := p.seekableReader.Seek(savedPos, io.SeekStart); err != nil {
					// Mark the reader position as corrupt so subsequent reads fail explicitly.
					// The data for this element was read successfully, but the shared reader
					// position is now unknown and all future lazy/sequential reads are unsafe.
					p.seekableReader = nil
				}
			}()

			// Read the data
			data := make([]byte, length)
			if _, err := io.ReadFull(p.seekableReader, data); err != nil {
				return nil, fmt.Errorf("failed to read lazy data: %w", err)
			}

			return data, nil
		}

		lb, err := buffer.NewLazySizedWithError(length, loader)
		if err != nil {
			return nil, fmt.Errorf("failed to create lazy buffer: %w", err)
		}

		// Skip past this data in the stream
		if _, err := p.seekableReader.Seek(int64(length), io.SeekCurrent); err != nil {
			return nil, fmt.Errorf("failed to skip data: %w", err)
		}
		p.consumeBounded(length)

		return lb, nil
	}

	// No seekable reader available
	return nil, fmt.Errorf("lazy loading not supported for non-seekable readers")
}

func createEmptyFragmentSequence(t *tag.Tag, vrValue *vr.VR) (element.Element, error) {
	switch vrValue.Code() {
	case vr.CodeOB:
		return element.NewOtherByteFragment(t), nil
	case vr.CodeOW:
		return element.NewOtherWordFragment(t), nil
	default:
		return nil, fmt.Errorf("invalid VR for fragment sequence: %v", vrValue)
	}
}

// skipFragmentSequence skips an encapsulated fragment sequence payload.
func (p *parseContext) skipFragmentSequence() error {
	fragmentCount := 0
	for {
		if fragmentCount >= maxSequenceItems {
			return fmt.Errorf("exceeded maximum fragment count (%d)", maxSequenceItems)
		}
		fragmentCount++

		itemTag, err := p.readTag()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("failed to read fragment item tag: %w", err)
		}

		// Sequence Delimitation Item (FFFE,E0DD)
		if itemTag.Group() == 0xFFFE && itemTag.Element() == 0xE0DD {
			var delimLength uint32
			if err := binary.Read(p.reader, p.byteOrder, &delimLength); err != nil {
				return fmt.Errorf("failed to read sequence delimitation length: %w", err)
			}
			return nil
		}

		// Item tag (FFFE,E000)
		if itemTag.Group() != 0xFFFE || itemTag.Element() != 0xE000 {
			return fmt.Errorf("expected Item tag (FFFE,E000) in fragment sequence, got %s", itemTag)
		}

		var itemLength uint32
		if err := binary.Read(p.reader, p.byteOrder, &itemLength); err != nil {
			return fmt.Errorf("failed to read fragment item length: %w", err)
		}
		if itemLength == 0xFFFFFFFF {
			return fmt.Errorf("undefined length fragment items are not supported")
		}
		if p.maxElementSize > 0 && itemLength > p.maxElementSize {
			return fmt.Errorf("fragment item length %d exceeds maximum %d", itemLength, p.maxElementSize)
		}

		if _, err := io.CopyN(io.Discard, p.reader, int64(itemLength)); err != nil {
			return fmt.Errorf("failed to skip fragment item data: %w", err)
		}
	}
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
	fragmentCount := 0

	for {
		if fragmentCount >= maxSequenceItems {
			return nil, fmt.Errorf("exceeded maximum fragment count (%d) for tag %s", maxSequenceItems, t)
		}
		fragmentCount++

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
		if itemLength == 0xFFFFFFFF {
			return nil, fmt.Errorf("undefined length fragment items are not supported")
		}
		if p.maxElementSize > 0 && itemLength > p.maxElementSize {
			return nil, fmt.Errorf("fragment item length %d exceeds maximum %d", itemLength, p.maxElementSize)
		}

		if isFirstItem {
			// The first item is the Basic Offset Table and must be materialized
			// so offsets can be parsed before reading fragments.
			// Zero-length offset table is valid (empty).
			itemData := make([]byte, itemLength)
			if _, err := io.ReadFull(p.reader, itemData); err != nil {
				return nil, fmt.Errorf("failed to read fragment item data: %w", err)
			}

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
			// Per DICOM PS3.5 §7.5, fragment items (non-offset-table) must have
			// a non-zero, even-length value. Zero-length fragments are invalid.
			if itemLength == 0 {
				return nil, fmt.Errorf("fragment %d has zero length (offset table already read)", len(fragSeq.Fragments())+1)
			}
			if itemLength%2 != 0 {
				return nil, fmt.Errorf("fragment %d length %d is not even", len(fragSeq.Fragments())+1, itemLength)
			}
			fragBuf, err := p.readFragmentItemBuffer(itemLength)
			if err != nil {
				return nil, err
			}
			fragSeq.AddFragment(fragBuf)
		}
	}

	return fs, nil
}

func (p *parseContext) readFragmentItemBuffer(itemLength uint32) (buffer.ByteBuffer, error) {
	if p.readOption == ReadLargeOnDemand && itemLength > p.largeObjectSize {
		fragBuf, err := p.createLazyBuffer(itemLength)
		if err == nil {
			return fragBuf, nil
		}
		// Error from createLazyBuffer is intentionally swallowed because lazy loading fails
		// for non-seekable readers (e.g. deflated streams from flate.Reader). Falling through
		// to eager loading here ensures the parse continues and the data is still available,
		// just loaded eagerly instead of on-demand.
	}

	itemData := make([]byte, itemLength)
	if _, err := io.ReadFull(p.reader, itemData); err != nil {
		return nil, fmt.Errorf("failed to read fragment item data: %w", err)
	}
	return buffer.NewMemory(itemData), nil
}
