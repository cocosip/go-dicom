// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"compress/flate"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// SetOffsetTableForFrames computes and sets the Basic Offset Table (BOT) for a fragment sequence.
// frameStartIndexes must list the starting fragment index for each frame (first index must be 0).
// Offsets are computed against the encoded fragment item stream, per DICOM
// requirements: each preceding fragment contributes its 8-byte item header plus
// its even-length padded value.
func SetOffsetTableForFrames(fs *element.FragmentSequence, frameStartIndexes []int) error {
	fragCount := fs.FragmentCount()
	if fragCount == 0 {
		return fmt.Errorf("cannot build offset table: no fragments")
	}
	if len(frameStartIndexes) == 0 {
		return fmt.Errorf("frameStartIndexes cannot be empty")
	}
	if frameStartIndexes[0] != 0 {
		return fmt.Errorf("frameStartIndexes must start with 0")
	}
	// Verify ascending order and bounds.
	for i := 1; i < len(frameStartIndexes); i++ {
		if frameStartIndexes[i] <= frameStartIndexes[i-1] {
			return fmt.Errorf("frameStartIndexes must be strictly increasing")
		}
	}
	if frameStartIndexes[len(frameStartIndexes)-1] >= fragCount {
		return fmt.Errorf("frameStartIndexes contains out-of-range index")
	}

	// Compute encoded fragment item lengths to align offsets as written.
	encodedLens := make([]uint32, fragCount)
	for i := 0; i < fragCount; i++ {
		frag, err := fs.GetFragment(i)
		if err != nil {
			return err
		}
		size := frag.Size()
		padded := size
		if padded%2 != 0 {
			padded++
		}
		if padded > math.MaxUint32-8 {
			return fmt.Errorf("fragment %d is too large to represent in BOT", i)
		}
		encodedLens[i] = 8 + padded
	}

	// Precompute prefix sums for quick offset lookup.
	prefix := make([]uint32, fragCount+1)
	for i := 0; i < fragCount; i++ {
		if prefix[i] > math.MaxUint32-encodedLens[i] {
			return fmt.Errorf("offset overflow at fragment %d", i)
		}
		prefix[i+1] = prefix[i] + encodedLens[i]
	}

	offsets := make([]uint32, len(frameStartIndexes))
	for i, idx := range frameStartIndexes {
		offsets[i] = prefix[idx]
	}

	fs.SetOffsetTable(offsets)
	return nil
}

// Global configuration for DICOM implementation identification.
// These values are used by default for all DICOM files written by this package.
// They can be customized at application startup using SetDefaultImplementationClassUID()
// and SetDefaultImplementationVersionName().
var (
	globalConfig = struct {
		mu                        sync.RWMutex
		implementationClassUID    string
		implementationVersionName string
	}{
		implementationClassUID:    "1.2.826.0.1.3680043.10.1142", // Default UID
		implementationVersionName: "GO-DICOM_1.0",                // Default version
	}
)

// SetDefaultImplementationClassUID sets the global default Implementation Class UID.
// This UID uniquely identifies your DICOM implementation and should be registered
// with your organization's UID root.
//
// This setting affects all DICOM files written by this package unless overridden
// by WithImplementationClassUID option.
//
// Typical usage at application startup:
//
//	writer.SetDefaultImplementationClassUID("1.2.840.12345.1.2.3")
func SetDefaultImplementationClassUID(uid string) {
	globalConfig.mu.Lock()
	defer globalConfig.mu.Unlock()
	globalConfig.implementationClassUID = uid
}

// SetDefaultImplementationVersionName sets the global default Implementation Version Name.
// This identifies the version of your DICOM implementation (e.g., "MyApp_2.1.0").
//
// This setting affects all DICOM files written by this package unless overridden
// by WithImplementationVersionName option.
//
// Typical usage at application startup:
//
//	writer.SetDefaultImplementationVersionName("MyDicomApp_2.1.0")
func SetDefaultImplementationVersionName(name string) {
	globalConfig.mu.Lock()
	defer globalConfig.mu.Unlock()
	globalConfig.implementationVersionName = name
}

// GetDefaultImplementationClassUID returns the current global default Implementation Class UID.
func GetDefaultImplementationClassUID() string {
	globalConfig.mu.RLock()
	defer globalConfig.mu.RUnlock()
	return globalConfig.implementationClassUID
}

// GetDefaultImplementationVersionName returns the current global default Implementation Version Name.
func GetDefaultImplementationVersionName() string {
	globalConfig.mu.RLock()
	defer globalConfig.mu.RUnlock()
	return globalConfig.implementationVersionName
}

// Writer writes DICOM files.
type Writer struct {
	writer         io.Writer
	byteOrder      binary.ByteOrder
	isExplicitVR   bool
	transferSyntax *transfer.Syntax

	// Write options
	includePreamble             bool   // Whether to include 128-byte preamble + DICM
	explicitLengthSequences     bool   // Use explicit length for sequences
	explicitLengthSequenceItems bool   // Use explicit length for sequence items
	keepGroupLengths            bool   // Keep group length tags (0xGGGG,0x0000)
	largeObjectSize             uint32 // Threshold for large objects
	implementationClassUID      string // Implementation Class UID for File Meta Information
	implementationVersionName   string // Implementation Version Name for File Meta Information
	sequenceItemObserver        SequenceItemObserver
	positionWriter              *positionWriter
}

// WriteOption is a functional option for Write function.
type WriteOption func(*writeConfig)

// writeConfig holds the configuration for a write operation.
type writeConfig struct {
	transferSyntax              *transfer.Syntax
	transferSyntaxSet           bool
	fileMetaInfo                *dataset.Dataset
	includePreamble             bool
	explicitLengthSequences     bool   // Use explicit length for sequences (default: false, use undefined)
	explicitLengthSequenceItems bool   // Use explicit length for sequence items (default: false, use undefined)
	keepGroupLengths            bool   // Keep group length tags (0xGGGG,0x0000) (default: false)
	largeObjectSize             uint32 // Threshold for large objects (default: 1MB)
	implementationClassUID      string // Implementation Class UID (default: auto-generated)
	implementationVersionName   string // Implementation Version Name (default: GO-DICOM_1.0)
	sequenceItemObserver        SequenceItemObserver
}

// WithTransferSyntax specifies the transfer syntax to use.
// If not specified, defaults to Explicit VR Little Endian.
func WithTransferSyntax(ts *transfer.Syntax) WriteOption {
	return func(c *writeConfig) {
		c.transferSyntax = ts
		c.transferSyntaxSet = true
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
func WithExplicitLengthSequences(explicitLengthSequences bool) WriteOption {
	return func(c *writeConfig) {
		c.explicitLengthSequences = explicitLengthSequences
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

// WithImplementationClassUID sets the Implementation Class UID (0002,0012) for this specific write operation.
// This overrides the global default set by SetDefaultImplementationClassUID().
//
// Use this option when you need to write a file with a different implementation UID
// than your application's default (rare cases).
func WithImplementationClassUID(uid string) WriteOption {
	return func(c *writeConfig) {
		c.implementationClassUID = uid
	}
}

// WithImplementationVersionName sets the Implementation Version Name (0002,0013) for this specific write operation.
// This overrides the global default set by SetDefaultImplementationVersionName().
//
// Use this option when you need to write a file with a different version name
// than your application's default (rare cases).
func WithImplementationVersionName(name string) WriteOption {
	return func(c *writeConfig) {
		c.implementationVersionName = name
	}
}

// Option is a functional option for Writer (internal use, primarily for tests).
type Option func(*Writer)

// New creates a new Writer with the given transfer syntax and options.
// If ts is nil, defaults to Explicit VR Little Endian.
func New(ts *transfer.Syntax, opts ...Option) *Writer {
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
//
//	writer.Write(w, ds)  // Uses Explicit VR Little Endian, auto-generates File Meta Info
//
// With options:
//
//	writer.Write(w, ds, writer.WithTransferSyntax(ts))
//	writer.Write(w, ds, writer.WithFileMetaInfo(fmi))
//	writer.Write(w, ds, writer.WithoutPreamble())
//
// The structure written is:
//   - Preamble (128 bytes) + DICM prefix (4 bytes) [default, can be disabled with WithoutPreamble()]
//   - File Meta Information (Group 0002, always Explicit VR Little Endian)
//   - Dataset (encoding depends on Transfer Syntax)
func Write(w io.Writer, ds *dataset.Dataset, opts ...WriteOption) error {
	if ds == nil {
		return fmt.Errorf("dataset cannot be nil")
	}

	// Apply options to configuration
	// Use global defaults initially
	config := &writeConfig{
		transferSyntax:              transfer.ExplicitVRLittleEndian,       // Default
		fileMetaInfo:                nil,                                   // Will be auto-generated
		includePreamble:             true,                                  // Default to including preamble
		explicitLengthSequences:     false,                                 // Default: use undefined length
		explicitLengthSequenceItems: false,                                 // Default: use undefined length
		keepGroupLengths:            false,                                 // Default: remove group lengths
		largeObjectSize:             1024 * 1024,                           // Default: 1MB
		implementationClassUID:      GetDefaultImplementationClassUID(),    // Use global default
		implementationVersionName:   GetDefaultImplementationVersionName(), // Use global default
	}

	for _, opt := range opts {
		opt(config)
	}

	// If no transfer syntax was explicitly specified via options,
	// try to use the dataset's InternalTransferSyntax (set by transcoder/parser)
	// This allows automatic transfer syntax detection from the dataset
	if !config.transferSyntaxSet {
		if internalTS := ds.InternalTransferSyntax(); internalTS != nil {
			config.transferSyntax = internalTS
		}
	}
	if config.transferSyntax == nil {
		return fmt.Errorf("transfer syntax cannot be nil")
	}
	if config.sequenceItemObserver != nil && config.transferSyntax.IsDeflate() {
		return fmt.Errorf("sequence item positions are unavailable for deflated transfer syntax")
	}
	if err := validatePixelDataTransferSyntax(ds, config.transferSyntax); err != nil {
		return err
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
		sequenceItemObserver:        config.sequenceItemObserver,
	}
	if config.sequenceItemObserver != nil {
		positioned := &positionWriter{writer: w}
		writer.writer = positioned
		writer.positionWriter = positioned
	}

	// Store implementation info in writer for use in generateFileMetaInformation
	writer.implementationClassUID = config.implementationClassUID
	writer.implementationVersionName = config.implementationVersionName

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
		// Work on a clone so callers can reuse their File Meta Information
		// dataset. Its TransferSyntaxUID must always match the dataset bytes
		// written below, even when the caller supplied stale metadata.
		fileMetaInfo = fileMetaInfo.Clone()
		if err := fileMetaInfo.AddOrUpdate(element.NewString(tag.TransferSyntaxUID, vr.UI,
			[]string{config.transferSyntax.UID().String()})); err != nil {
			return fmt.Errorf("failed to set TransferSyntaxUID in file meta information: %w", err)
		}
	}

	if writer.includePreamble {
		if err := synchronizeFileMetaSOPUIDs(fileMetaInfo, ds); err != nil {
			return err
		}
	}

	// Write preamble and File Meta Information if requested
	// For network DIMSE messages, we skip both preamble and File Meta Information
	if writer.includePreamble {
		if err := writer.writePreamble(); err != nil {
			return fmt.Errorf("failed to write preamble: %w", err)
		}

		// Write File Meta Information (always Explicit VR Little Endian)
		if err := writer.writeFileMetaInformation(fileMetaInfo); err != nil {
			return fmt.Errorf("failed to write file meta information: %w", err)
		}
	}

	// Write main dataset. Deflated Explicit VR Little Endian compresses only
	// the dataset following File Meta Information.
	if config.transferSyntax.IsDeflate() {
		fw, err := flate.NewWriter(writer.writer, flate.DefaultCompression)
		if err != nil {
			return fmt.Errorf("failed to create deflate writer: %w", err)
		}
		deflatedWriter := *writer
		deflatedWriter.writer = fw
		if err := deflatedWriter.writeDataset(ds); err != nil {
			_ = fw.Close()
			return fmt.Errorf("failed to write deflated dataset: %w", err)
		}
		if err := fw.Close(); err != nil {
			return fmt.Errorf("failed to finish deflated dataset: %w", err)
		}
	} else if err := writer.writeDataset(ds); err != nil {
		return fmt.Errorf("failed to write dataset: %w", err)
	}

	return nil
}

func validatePixelDataTransferSyntax(ds *dataset.Dataset, ts *transfer.Syntax) error {
	pixelData, exists := ds.Get(tag.PixelData)
	if !exists {
		return nil
	}

	encapsulated := false
	switch pixelData.(type) {
	case *element.FragmentSequence, *element.OtherByteFragment, *element.OtherWordFragment:
		encapsulated = true
	}
	if encapsulated != ts.IsEncapsulated() {
		return fmt.Errorf("pixel data representation is incompatible with transfer syntax %s", ts.UID().String())
	}
	return nil
}

func synchronizeFileMetaSOPUIDs(fileMetaInfo, ds *dataset.Dataset) error {
	sopClassUID, hasSOPClassUID := ds.GetString(tag.SOPClassUID)
	sopInstanceUID, hasSOPInstanceUID := ds.GetString(tag.SOPInstanceUID)
	if (!hasSOPClassUID || sopClassUID == "") && (!hasSOPInstanceUID || sopInstanceUID == "") {
		mediaStorageClassUID, hasMediaStorageClassUID := fileMetaInfo.GetString(tag.MediaStorageSOPClassUID)
		mediaStorageInstanceUID, hasMediaStorageInstanceUID := fileMetaInfo.GetString(tag.MediaStorageSOPInstanceUID)
		if hasMediaStorageClassUID && mediaStorageClassUID == uid.MediaStorageDirectoryStorage.UID() &&
			hasMediaStorageInstanceUID && mediaStorageInstanceUID != "" {
			return nil
		}
	}
	if !hasSOPClassUID || sopClassUID == "" {
		return fmt.Errorf("dataset is missing SOPClassUID required for Part 10 File Meta Information")
	}
	if !hasSOPInstanceUID || sopInstanceUID == "" {
		return fmt.Errorf("dataset is missing SOPInstanceUID required for Part 10 File Meta Information")
	}

	for _, item := range []struct {
		metaTag    *tag.Tag
		datasetUID string
		name       string
	}{
		{tag.MediaStorageSOPClassUID, sopClassUID, "MediaStorageSOPClassUID"},
		{tag.MediaStorageSOPInstanceUID, sopInstanceUID, "MediaStorageSOPInstanceUID"},
	} {
		if metaUID, exists := fileMetaInfo.GetString(item.metaTag); exists {
			if metaUID != item.datasetUID {
				return fmt.Errorf("%s %q does not match dataset UID %q", item.name, metaUID, item.datasetUID)
			}
			continue
		}
		if _, exists := fileMetaInfo.Get(item.metaTag); exists {
			return fmt.Errorf("file meta information contains invalid %s", item.name)
		}
		if err := fileMetaInfo.Add(element.NewString(item.metaTag, vr.UI, []string{item.datasetUID})); err != nil {
			return fmt.Errorf("failed to set %s in file meta information: %w", item.name, err)
		}
	}
	return nil
}

// WriteFile writes a DICOM dataset to a file.
//
// Basic usage:
//
//	writer.WriteFile("output.dcm", ds)
//
// With options:
//
//	writer.WriteFile("output.dcm", ds, writer.WithTransferSyntax(ts))
func WriteFile(path string, ds *dataset.Dataset, opts ...WriteOption) error {
	// Clean the provided path to mitigate path traversal and ensure a canonical form.
	// Cross-platform: filepath.Clean works on Windows, Linux, and macOS.
	cleanPath := filepath.Clean(path)

	// Optional: basic sanity check to avoid extremely suspicious paths.
	// We only permit creating regular files (no device names, etc.).
	// Users may still pass any absolute or relative path; we just canonicalize.

	file, err := os.Create(cleanPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", cleanPath, err)
	}

	return writeAndClose(file, ds, opts...)
}

func writeAndClose(file io.WriteCloser, ds *dataset.Dataset, opts ...WriteOption) error {
	writeErr := Write(file, ds, opts...)
	closeErr := file.Close()
	if writeErr != nil && closeErr != nil {
		return errors.Join(writeErr, closeErr)
	}
	if writeErr != nil {
		return writeErr
	}
	return closeErr
}

// generateFileMetaInformation generates a minimal File Meta Information dataset.
func (w *Writer) generateFileMetaInformation() *dataset.Dataset {
	fileMetaInfo := dataset.New()

	// Add required FileMetaInformationVersion (0002,0001)
	// Value should be 0x00 0x01 according to DICOM standard
	_ = fileMetaInfo.Add(element.NewOtherByte(tag.FileMetaInformationVersion, []byte{0x00, 0x01}))

	// Add required Transfer Syntax UID (0002,0010)
	_ = fileMetaInfo.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
		[]string{w.transferSyntax.UID().String()}))

	// Add required Implementation Class UID (0002,0012)
	// Use configured value or default
	implClassUID := w.implementationClassUID
	if implClassUID == "" {
		implClassUID = "1.2.826.0.1.3680043.10.1142" // Default UID
	}
	_ = fileMetaInfo.Add(element.NewString(tag.ImplementationClassUID, vr.UI,
		[]string{implClassUID}))

	// Add Implementation Version Name (0002,0013) - optional but recommended
	// Use configured value or default
	implVersionName := w.implementationVersionName
	if implVersionName == "" {
		implVersionName = "GO-DICOM_1.0" // Default version
	}
	_ = fileMetaInfo.Add(element.NewString(tag.ImplementationVersionName, vr.SH,
		[]string{implVersionName}))

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
	tempBuf := buffer.GetBytesBuffer()
	defer buffer.PutBytesBuffer(tempBuf)

	tempWriter := &Writer{
		writer:          tempBuf,
		byteOrder:       binary.LittleEndian,
		isExplicitVR:    true,
		largeObjectSize: w.largeObjectSize,
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
	groupLength := uint32(tempBuf.Len()) // #nosec G115 -- DICOM group length within uint32 range
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
	lengths, err := w.calculateWalkLengths(ds)
	if err != nil {
		return err
	}
	return w.writeDatasetWithWalk(ds, lengths)
}

// writeElement writes a single DICOM element.
func (w *Writer) writeElement(elem element.Element) error {
	if elem == nil {
		return fmt.Errorf("cannot write nil element")
	}
	switch elem.(type) {
	case *dataset.Sequence, *element.FragmentSequence, *element.OtherByteFragment, *element.OtherWordFragment:
		return fmt.Errorf("container element %s must be written through Dataset Walk", elem.Tag())
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

	buf := elem.Buffer()
	if buf == nil {
		buf = buffer.Empty
	}
	valueLength := buf.Size()
	paddedLength := valueLength
	needsPadding := valueLength%2 != 0
	if needsPadding {
		if valueLength == math.MaxUint32 {
			return fmt.Errorf("cannot pad maximum-length value for tag %s", elem.Tag())
		}
		paddedLength++
	}

	// Write length
	if err := w.writeLength(elemVR, paddedLength); err != nil {
		return fmt.Errorf("failed to write length for tag %s: %w", elem.Tag(), err)
	}

	// Write value
	if valueLength > 0 {
		streamLargeValue := !buf.IsMemory()
		if !streamLargeValue && w.largeObjectSize > 0 && valueLength > w.largeObjectSize {
			streamLargeValue = true
		}
		if streamLargeValue {
			written, err := buf.WriteTo(w.writer)
			if err != nil {
				return fmt.Errorf("failed to stream value for tag %s: %w", elem.Tag(), err)
			}
			if written != int64(valueLength) {
				return fmt.Errorf("short write for tag %s: wrote %d bytes, expected %d", elem.Tag(), written, valueLength)
			}
		} else {
			valueBytes := buf.Data()
			if uint32(len(valueBytes)) != valueLength { //nolint:gosec // buffer sizes are uint32 by interface contract
				return fmt.Errorf("buffer size mismatch for tag %s: got %d bytes, expected %d", elem.Tag(), len(valueBytes), valueLength)
			}
			if err := writeAll(w.writer, valueBytes); err != nil {
				return fmt.Errorf("failed to write value for tag %s: %w", elem.Tag(), err)
			}
		}
	}
	if needsPadding {
		if err := writeAll(w.writer, []byte{padByteForVR(elemVR)}); err != nil {
			return fmt.Errorf("failed to write padding for tag %s: %w", elem.Tag(), err)
		}
	}

	return nil
}

func padByteForVR(v *vr.VR) byte {
	if v == vr.UI {
		return 0x00
	}
	if isStringVR(v) {
		return ' '
	}
	return 0x00
}

func isStringVR(v *vr.VR) bool {
	switch v.Code() {
	case vr.CodeAE, vr.CodeAS, vr.CodeCS, vr.CodeDA, vr.CodeDS, vr.CodeDT,
		vr.CodeIS, vr.CodeLO, vr.CodeLT, vr.CodePN, vr.CodeSH, vr.CodeST,
		vr.CodeTM, vr.CodeUC, vr.CodeUI, vr.CodeUR, vr.CodeUT:
		return true
	default:
		return false
	}
}

func writeAll(w io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		if n <= 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
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
	if seq == nil {
		return fmt.Errorf("cannot write nil sequence")
	}
	ds := dataset.New()
	ds.SetAutoValidate(false)
	if err := ds.Add(seq); err != nil {
		return err
	}
	return w.writeDataset(ds)
}

func (w *Writer) currentOffset() uint64 {
	if w.positionWriter == nil {
		return 0
	}
	return w.positionWriter.offset
}

func (w *Writer) observeSequenceItem(sequenceTag *tag.Tag, item *dataset.Dataset) error {
	if w.sequenceItemObserver == nil {
		return nil
	}
	offset := w.currentOffset()
	if err := w.sequenceItemObserver(SequenceItemPosition{
		SequenceTag: sequenceTag,
		Item:        item,
		Offset:      offset,
	}); err != nil {
		return fmt.Errorf("sequence item observer failed for %s at offset %d: %w", sequenceTag, offset, err)
	}
	return nil
}

// writeFragmentSequence writes a DICOM fragment sequence (encapsulated pixel data).
// Fragment sequences are used for compressed image formats like JPEG, JPEG 2000, RLE, etc.
//
// Structure:
// - Tag + VR + Undefined Length (FFFFFFFF)
// - Item (FFFE,E000): Offset Table
// - Items (FFFE,E000): Fragments
// - Sequence Delimitation Item (FFFE,E0DD)
func (w *Writer) writeFragmentSequence(fs *element.FragmentSequence) error {
	if fs == nil {
		return fmt.Errorf("cannot write nil fragment sequence")
	}
	ds := dataset.New()
	ds.SetAutoValidate(false)
	if err := ds.Add(fs); err != nil {
		return err
	}
	return w.writeDataset(ds)
}
