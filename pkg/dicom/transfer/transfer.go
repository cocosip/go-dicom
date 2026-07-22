// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package transfer implements DICOM transfer syntax functionality.
package transfer

import (
	"fmt"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/parseable"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

var _ parseable.Parseable = (*Syntax)(nil)

// Syntax represents a DICOM transfer syntax.
//
// A transfer syntax defines how DICOM data is encoded, including:
// - Value Representation (VR) encoding (Explicit or Implicit)
// - Byte order (Little Endian or Big Endian)
// - Compression (e.g., JPEG, JPEG 2000)
type Syntax struct {
	uid                    *uid.UID
	isRetired              bool
	isExplicitVR           bool
	isEncapsulated         bool
	isLossy                bool
	lossyCompressionMethod string
	isDeflate              bool
	endian                 endian.Endian
	swapPixelData          bool
}

// Registry holds a collection of transfer syntaxes keyed by UID string.
// Use NewRegistry to create an isolated registry preloaded with standard
// transfer syntaxes.
type Registry struct {
	mu    sync.RWMutex
	items map[string]*Syntax
}

var standardTransferSyntaxes = []*Syntax{
	ImplicitVRLittleEndian,
	ExplicitVRLittleEndian,
	ExplicitVRBigEndian,
	ImplicitVRBigEndian,
	GEPrivateImplicitVRBigEndian,
	Papyrus3ImplicitVRLittleEndianRetired,
	DeflatedExplicitVRLittleEndian,
	JPEGProcess1,
	JPEGBaseline8Bit,
	JPEGExtended12Bit,
	JPEGProcess2_4,
	JPEGProcess3_5Retired,
	JPEGProcess6_8Retired,
	JPEGProcess7_9Retired,
	JPEGProcess10_12Retired,
	JPEGProcess11_13Retired,
	JPEGProcess16_18Retired,
	JPEGProcess17_19Retired,
	JPEGProcess20_22Retired,
	JPEGProcess21_23Retired,
	JPEGProcess24_26Retired,
	JPEGProcess25_27Retired,
	JPEGLossless,
	JPEGProcess15Retired,
	JPEGProcess28Retired,
	JPEGProcess29Retired,
	JPEGLosslessSV1,
	JPEGLSLossless,
	JPEGLSNearLossless,
	JPEG2000Lossless,
	JPEG2000,
	JPEG2000Part2MultiComponentLosslessOnly,
	JPEG2000Part2MultiComponent,
	MPEG2,
	FragmentableMPEG2,
	MPEG2MainProfileHighLevel,
	FragmentableMPEG2MainProfileHighLevel,
	MPEG4AVCH264HighProfileLevel41,
	FragmentableMPEG4AVCH264HighProfileLevel41,
	MPEG4AVCH264BDCompatibleHighProfileLevel41,
	FragmentableMPEG4AVCH264BDCompatibleHighProfileLevel41,
	MPEG4AVCH264HighProfileLevel42For2DVideo,
	FragmentableMPEG4AVCH264HighProfileLevel42For2DVideo,
	MPEG4AVCH264HighProfileLevel42For3DVideo,
	FragmentableMPEG4AVCH264HighProfileLevel42For3DVideo,
	MPEG4AVCH264StereoHighProfileLevel42,
	FragmentableMPEG4AVCH264StereoHighProfileLevel42,
	HEVCH265MainProfileLevel51,
	HEVCH265Main10ProfileLevel51,
	HTJ2KLossless,
	HTJ2KLosslessRPCL,
	HTJ2K,
	JPIPReferenced,
	JPIPReferencedDeflate,
	JPIPHTJ2KReferenced,
	JPIPHTJ2KReferencedDeflate,
	RLELossless,
	RFC2557MIMEEncapsulation,
	XMLEncoding,
}

// DefaultRegistry is the package-level registry used by Register, Lookup, etc.
var DefaultRegistry = newEmptyRegistry()

// NewRegistry creates a new isolated transfer syntax registry containing the
// standard transfer syntaxes known to the default registry.
func NewRegistry() *Registry {
	r := newEmptyRegistry()
	for _, ts := range standardTransferSyntaxes {
		r.Register(cloneSyntax(ts))
	}
	return r
}

func newEmptyRegistry() *Registry {
	return &Registry{
		items: make(map[string]*Syntax),
	}
}

func cloneSyntax(ts *Syntax) *Syntax {
	if ts == nil {
		return nil
	}

	cloned := *ts
	return &cloned
}

// New creates a new TransferSyntax with the given UID.
func New(u *uid.UID) *Syntax {
	return &Syntax{
		uid:            u,
		isRetired:      u.IsRetired(),
		isExplicitVR:   false,
		isEncapsulated: false,
		isLossy:        false,
		isDeflate:      false,
		endian:         endian.Little,
		swapPixelData:  false,
	}
}

// UID returns the unique identifier of the transfer syntax.
func (ts *Syntax) UID() *uid.UID {
	return ts.uid
}

// IsRetired returns true if the transfer syntax is retired.
func (ts *Syntax) IsRetired() bool {
	return ts.isRetired
}

// IsExplicitVR returns true if the transfer syntax uses explicit VR.
func (ts *Syntax) IsExplicitVR() bool {
	return ts.isExplicitVR
}

// IsEncapsulated returns true if the transfer syntax encapsulates pixel data.
func (ts *Syntax) IsEncapsulated() bool {
	return ts.isEncapsulated
}

// IsLossy returns true if the transfer syntax uses lossy compression.
func (ts *Syntax) IsLossy() bool {
	return ts.isLossy
}

// LossyCompressionMethod returns the lossy compression method identifier.
func (ts *Syntax) LossyCompressionMethod() string {
	return ts.lossyCompressionMethod
}

// IsDeflate returns true if the transfer syntax uses deflate compression.
func (ts *Syntax) IsDeflate() bool {
	return ts.isDeflate
}

// Endian returns the byte order of the transfer syntax.
func (ts *Syntax) Endian() endian.Endian {
	return ts.endian
}

// SwapPixelData returns true if pixel data requires byte swapping.
func (ts *Syntax) SwapPixelData() bool {
	return ts.swapPixelData
}

// String returns the name of the transfer syntax.
func (ts *Syntax) String() string {
	return ts.uid.Name()
}

// Equals checks if two transfer syntaxes are equal.
func (ts *Syntax) Equals(other *Syntax) bool {
	if other == nil {
		return false
	}
	return ts.uid.Equals(other.uid)
}

// Parse parses a transfer syntax UID string and returns the corresponding TransferSyntax.
func Parse(uidString string) (*Syntax, error) {
	u := uid.Parse(uidString, "", uid.TypeTransferSyntax)
	return Lookup(u)
}

// Lookup looks up a transfer syntax by UID in the default registry.
//
// If the UID is registered, returns the registered transfer syntax.
// Otherwise, creates a new transfer syntax with default properties:
// - Explicit VR
// - Encapsulated
// - Little Endian
func Lookup(u *uid.UID) (*Syntax, error) {
	return DefaultRegistry.Lookup(u)
}

// Lookup looks up a transfer syntax by UID in this registry.
//
// If the UID is registered, returns the registered transfer syntax.
// Otherwise, creates a new transfer syntax with default properties:
// - Explicit VR
// - Encapsulated
// - Little Endian
func (r *Registry) Lookup(u *uid.UID) (*Syntax, error) {
	if u == nil {
		return nil, fmt.Errorf("UID cannot be nil")
	}

	if u.Type() != uid.TypeTransferSyntax {
		return nil, fmt.Errorf("UID %s is not a transfer syntax type", u.UID())
	}

	r.mu.RLock()
	ts, found := r.items[u.UID()]
	r.mu.RUnlock()

	if found {
		return ts, nil
	}

	// Create default transfer syntax for unknown UIDs
	ts = New(u)
	ts.isExplicitVR = true
	ts.isEncapsulated = true
	ts.endian = endian.Little

	return ts, nil
}

// Parse implements parseable.Parseable interface.
func (ts *Syntax) Parse(s string) error {
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*ts = *parsed
	return nil
}

// Register registers a transfer syntax in the default registry.
// This is the package-level convenience function.
func Register(ts *Syntax) {
	DefaultRegistry.Register(ts)
}

// Register adds a transfer syntax to this registry.
func (r *Registry) Register(ts *Syntax) {
	if ts == nil || ts.uid == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[ts.uid.UID()] = ts
}

// Unregister removes a transfer syntax from the default registry.
// This is the package-level convenience function.
func Unregister(u *uid.UID) bool {
	return DefaultRegistry.Unregister(u)
}

// Unregister removes a transfer syntax from this registry.
func (r *Registry) Unregister(u *uid.UID) bool {
	if u == nil {
		return false
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.items[u.UID()]
	if exists {
		delete(r.items, u.UID())
	}
	return exists
}

// Query queries a transfer syntax by UID in the default registry. Returns nil if not found.
// This is the package-level convenience function.
func Query(u *uid.UID) *Syntax {
	return DefaultRegistry.Query(u)
}

// Query returns the transfer syntax registered under the given UID, or nil if not found.
func (r *Registry) Query(u *uid.UID) *Syntax {
	if u == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.items[u.UID()]
}

// KnownEntries returns all transfer syntaxes registered in the default registry.
// This is the package-level convenience function.
func KnownEntries() []*Syntax {
	return DefaultRegistry.KnownEntries()
}

// KnownEntries returns all transfer syntaxes in this registry.
func (r *Registry) KnownEntries() []*Syntax {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entries := make([]*Syntax, 0, len(r.items))
	for _, ts := range r.items {
		entries = append(entries, ts)
	}
	return entries
}

// Builder is a helper for constructing TransferSyntax instances with custom properties.
type Builder struct {
	ts *Syntax
}

// NewBuilder creates a new TransferSyntax builder.
func NewBuilder(u *uid.UID) *Builder {
	return &Builder{
		ts: New(u),
	}
}

// SetRetired sets the retired flag.
func (b *Builder) SetRetired(retired bool) *Builder {
	b.ts.isRetired = retired
	return b
}

// SetExplicitVR sets the explicit VR flag.
func (b *Builder) SetExplicitVR(explicitVR bool) *Builder {
	b.ts.isExplicitVR = explicitVR
	return b
}

// SetEncapsulated sets the encapsulated flag.
func (b *Builder) SetEncapsulated(encapsulated bool) *Builder {
	b.ts.isEncapsulated = encapsulated
	return b
}

// SetLossy sets the lossy flag and compression method.
func (b *Builder) SetLossy(lossy bool, method string) *Builder {
	b.ts.isLossy = lossy
	b.ts.lossyCompressionMethod = method
	return b
}

// SetDeflate sets the deflate flag.
func (b *Builder) SetDeflate(deflate bool) *Builder {
	b.ts.isDeflate = deflate
	return b
}

// SetEndian sets the byte order.
func (b *Builder) SetEndian(e endian.Endian) *Builder {
	b.ts.endian = e
	return b
}

// SetSwapPixelData sets the pixel data swapping flag.
func (b *Builder) SetSwapPixelData(swap bool) *Builder {
	b.ts.swapPixelData = swap
	return b
}

// Build returns the constructed TransferSyntax and registers it in the default registry.
func (b *Builder) Build() *Syntax {
	DefaultRegistry.Register(b.ts)
	return b.ts
}
