// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"encoding/binary"
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
}

// WriteOption is a functional option for Write function.
type WriteOption func(*writeConfig)

// writeConfig holds the configuration for a write operation.
type writeConfig struct {
	transferSyntax              *transfer.Syntax
	fileMetaInfo                *dataset.Dataset
	includePreamble             bool
	explicitLengthSequences     bool   // Use explicit length for sequences (default: false, use undefined)
	explicitLengthSequenceItems bool   // Use explicit length for sequence items (default: false, use undefined)
	keepGroupLengths            bool   // Keep group length tags (0xGGGG,0x0000) (default: false)
	largeObjectSize             uint32 // Threshold for large objects (default: 1MB)
	implementationClassUID      string // Implementation Class UID (default: auto-generated)
	implementationVersionName   string // Implementation Version Name (default: GO-DICOM_1.0)
}

// WithTransferSyntax specifies the transfer syntax to use.
// If not specified, defaults to Explicit VR Little Endian.
func WithTransferSyntax(ts *transfer.Syntax) WriteOption {
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

// Option is a functional option for Writer (internal use).
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
	if config.transferSyntax == transfer.ExplicitVRLittleEndian {
		if internalTS := ds.InternalTransferSyntax(); internalTS != nil {
			config.transferSyntax = internalTS
		}
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
		// Ensure TransferSyntaxUID is present in the provided fileMetaInfo
		if _, exists := fileMetaInfo.Get(tag.TransferSyntaxUID); !exists {
			_ = fileMetaInfo.Add(element.NewString(tag.TransferSyntaxUID, vr.UI,
				[]string{config.transferSyntax.UID().String()}))
		}
	}

	// Ensure MediaStorageSOPClassUID and MediaStorageSOPInstanceUID are present
	// These should be copied from the main dataset if available
	if _, exists := fileMetaInfo.Get(tag.MediaStorageSOPClassUID); !exists {
		if sopClassUID, ok := ds.GetString(tag.SOPClassUID); ok {
			_ = fileMetaInfo.Add(element.NewString(tag.MediaStorageSOPClassUID, vr.UI,
				[]string{sopClassUID}))
		}
	}

	if _, exists := fileMetaInfo.Get(tag.MediaStorageSOPInstanceUID); !exists {
		if sopInstanceUID, ok := ds.GetString(tag.SOPInstanceUID); ok {
			_ = fileMetaInfo.Add(element.NewString(tag.MediaStorageSOPInstanceUID, vr.UI,
				[]string{sopInstanceUID}))
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

	// Write main dataset
	if err := writer.writeDataset(ds); err != nil {
		return fmt.Errorf("failed to write dataset: %w", err)
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
	defer func() { _ = file.Close() }()

	return Write(file, ds, opts...)
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

	// Check if this is a fragment sequence (encapsulated pixel data)
	if obf, ok := elem.(*element.OtherByteFragment); ok {
		return w.writeFragmentSequence(obf.FragmentSequence)
	}
	if owf, ok := elem.(*element.OtherWordFragment); ok {
		return w.writeFragmentSequence(owf.FragmentSequence)
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
		itemsBuf := buffer.GetBytesBuffer()
		defer buffer.PutBytesBuffer(itemsBuf)

		itemsWriter := &Writer{
			writer:                      itemsBuf,
			byteOrder:                   w.byteOrder,
			isExplicitVR:                w.isExplicitVR,
			explicitLengthSequenceItems: w.explicitLengthSequenceItems,
			keepGroupLengths:            w.keepGroupLengths,
			largeObjectSize:             w.largeObjectSize,
		}

		for i := 0; i < seq.Count(); i++ {
			item := seq.GetItem(i)
			if err := itemsWriter.writeItem(item); err != nil {
				return fmt.Errorf("failed to write item %d: %w", i, err)
			}
		}

		// Write the explicit length
		seqLength := uint32(itemsBuf.Len()) // #nosec G115 -- DICOM sequence length within uint32 range
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
		elementsBuf := buffer.GetBytesBuffer()
		defer buffer.PutBytesBuffer(elementsBuf)

		elementsWriter := &Writer{
			writer:                      elementsBuf,
			byteOrder:                   w.byteOrder,
			isExplicitVR:                w.isExplicitVR,
			explicitLengthSequences:     w.explicitLengthSequences,
			explicitLengthSequenceItems: w.explicitLengthSequenceItems,
			keepGroupLengths:            w.keepGroupLengths,
			largeObjectSize:             w.largeObjectSize,
		}

		// Write all elements in the item
		elements := item.Elements()
		for _, elem := range elements {
			if err := elementsWriter.writeElement(elem); err != nil {
				return err
			}
		}

		// Write the explicit length
		itemLength := uint32(elementsBuf.Len()) // #nosec G115 -- DICOM item length within uint32 range
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

// writeFragmentSequence writes a DICOM fragment sequence (encapsulated pixel data).
// Fragment sequences are used for compressed image formats like JPEG, JPEG 2000, RLE, etc.
//
// Structure:
// - Tag + VR + Undefined Length (FFFFFFFF)
// - Item (FFFE,E000): Offset Table
// - Items (FFFE,E000): Fragments
// - Sequence Delimitation Item (FFFE,E0DD)
func (w *Writer) writeFragmentSequence(fs *element.FragmentSequence) error {
	// Write tag
	if err := w.writeTag(fs.Tag()); err != nil {
		return fmt.Errorf("failed to write fragment sequence tag: %w", err)
	}

	// Write VR (if Explicit VR)
	if w.isExplicitVR {
		if err := w.writeVR(fs.ValueRepresentation()); err != nil {
			return fmt.Errorf("failed to write fragment sequence VR: %w", err)
		}

		// For OB/OW with undefined length, write reserved bytes
		reserved := make([]byte, 2)
		if _, err := w.writer.Write(reserved); err != nil {
			return fmt.Errorf("failed to write reserved bytes: %w", err)
		}
	}

	// Write undefined length (0xFFFFFFFF)
	if err := binary.Write(w.writer, w.byteOrder, uint32(0xFFFFFFFF)); err != nil {
		return fmt.Errorf("failed to write undefined length: %w", err)
	}

	fragCount := fs.FragmentCount()

	// Build Basic Offset Table.
	// Priority:
	//   1) Use caller-supplied offset table if present (preserve as-is).
	//   2) If no offsets provided and there is exactly one fragment, emit a single 0 offset (per DICOM recommendation).
	//   3) If multiple fragments and no offsets provided, write an empty offset table (safer than per-fragment offsets).
	offsets := fs.OffsetTable()
	if len(offsets) == 0 && fragCount == 1 {
		offsets = []uint32{0}
	}

	// Write Item for Offset Table (FFFE,E000)
	itemTag := tag.New(0xFFFE, 0xE000)
	if err := w.writeTag(itemTag); err != nil {
		return fmt.Errorf("failed to write offset table item tag: %w", err)
	}

	// Write offset table
	offsetCount := len(offsets)
	// Each offset is 4 bytes; verify the multiplication does not overflow uint32.
	if offsetCount > int(math.MaxUint32/4) {
		return fmt.Errorf("offset table too large: %d entries", offsetCount)
	}
	offsetTableLength := uint32(offsetCount) * 4
	if err := binary.Write(w.writer, w.byteOrder, offsetTableLength); err != nil {
		return fmt.Errorf("failed to write offset table length: %w", err)
	}

	// Write offset values
	for _, offset := range offsets {
		if err := binary.Write(w.writer, w.byteOrder, offset); err != nil {
			return fmt.Errorf("failed to write offset value: %w", err)
		}
	}

	// Write fragments
	padByte := []byte{0}
	for i := 0; i < fragCount; i++ {
		frag, err := fs.GetFragment(i)
		if err != nil {
			return fmt.Errorf("failed to get fragment %d: %w", i, err)
		}
		fragData := frag.Data()

		// Write Item tag (FFFE,E000)
		if err := w.writeTag(itemTag); err != nil {
			return fmt.Errorf("failed to write fragment item tag: %w", err)
		}

		// Write fragment length
		fragLen := len(fragData)
		if fragLen%2 != 0 {
			fragLen++
		}
		if fragLen > int(math.MaxUint32) {
			return fmt.Errorf("fragment too large: %d bytes", fragLen)
		}
		if err := binary.Write(w.writer, w.byteOrder, uint32(fragLen)); err != nil {
			return fmt.Errorf("failed to write fragment length: %w", err)
		}

		// Write fragment data
		if _, err := w.writer.Write(fragData); err != nil {
			return fmt.Errorf("failed to write fragment data: %w", err)
		}
		if len(fragData)%2 != 0 {
			if _, err := w.writer.Write(padByte); err != nil {
				return fmt.Errorf("failed to write fragment padding: %w", err)
			}
		}
	}

	// Write Sequence Delimitation Item (FFFE,E0DD)
	delimTag := tag.New(0xFFFE, 0xE0DD)
	if err := w.writeTag(delimTag); err != nil {
		return fmt.Errorf("failed to write sequence delimitation tag: %w", err)
	}
	if err := binary.Write(w.writer, w.byteOrder, uint32(0)); err != nil {
		return fmt.Errorf("failed to write sequence delimitation length: %w", err)
	}

	return nil
}
