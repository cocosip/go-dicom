// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package transfer implements DICOM transfer syntax functionality.
package transfer

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/parseable"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/io/endian"
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
	mu      sync.RWMutex
	overlay map[string]*Syntax
	masked  map[string]struct{}
}

var (
	// ErrNilSyntax indicates that a registry operation received a nil Syntax.
	ErrNilSyntax = errors.New("transfer syntax cannot be nil")
	// ErrInvalidSyntax indicates that a registry operation received an invalid transfer syntax UID.
	ErrInvalidSyntax = errors.New("invalid transfer syntax")
	// ErrSyntaxAlreadyRegistered indicates that Register would replace an effective entry.
	ErrSyntaxAlreadyRegistered = errors.New("transfer syntax already registered")
)

var standardTransferSyntaxes = []*Syntax{
	ImplicitVRLittleEndian,
	ExplicitVRLittleEndian,
	ExplicitVRBigEndian,
	ImplicitVRBigEndian,
	GEPrivateImplicitVRBigEndian,
	Papyrus3ImplicitVRLittleEndianRetired,
	DeflatedExplicitVRLittleEndian,
	JPEGBaseline8Bit,
	JPEGExtended12Bit,
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

var standardTransferSyntaxIndex = buildSyntaxIndex(standardTransferSyntaxes)

// NewRegistry creates an isolated registry layered over the immutable standard catalog.
func NewRegistry() *Registry {
	return &Registry{
		overlay: make(map[string]*Syntax),
		masked:  make(map[string]struct{}),
	}
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
	return lookupSyntax(u, func(value *uid.UID) *Syntax { return standardTransferSyntaxIndex[value.UID()] })
}

// Lookup looks up a transfer syntax by UID in this registry.
//
// If the UID is registered, returns the registered transfer syntax.
// Otherwise, creates a new transfer syntax with default properties:
// - Explicit VR
// - Encapsulated
// - Little Endian
func (r *Registry) Lookup(u *uid.UID) (*Syntax, error) {
	return lookupSyntax(u, r.Query)
}

// Parse normalizes a UID string and resolves it through this Registry.
func (r *Registry) Parse(value string) (*Syntax, error) {
	return r.Lookup(uid.Parse(value, "", uid.TypeTransferSyntax))
}

func lookupSyntax(u *uid.UID, query func(*uid.UID) *Syntax) (*Syntax, error) {
	if u == nil {
		return nil, fmt.Errorf("UID cannot be nil")
	}

	if u.Type() != uid.TypeTransferSyntax {
		return nil, fmt.Errorf("UID %s is not a transfer syntax type", u.UID())
	}

	ts := query(u)
	if ts != nil {
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
	if ts != nil && ts.uid != nil && standardTransferSyntaxIndex[ts.uid.UID()] == ts {
		return fmt.Errorf("cannot modify standard transfer syntax %s", ts.UID().UID())
	}
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*ts = *parsed
	return nil
}

// Register adds a transfer syntax without replacing an effective entry.
func (r *Registry) Register(ts *Syntax) error {
	key, err := syntaxRegistryKey(ts)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.queryLocked(key) != nil {
		return fmt.Errorf("%w: %s", ErrSyntaxAlreadyRegistered, key)
	}
	r.overlay[key] = ts
	delete(r.masked, key)
	return nil
}

// Replace adds or replaces a transfer syntax and returns the previous effective entry.
func (r *Registry) Replace(ts *Syntax) (*Syntax, error) {
	key, err := syntaxRegistryKey(ts)
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	previous := r.queryLocked(key)
	r.overlay[key] = ts
	delete(r.masked, key)
	return previous, nil
}

// Unregister removes the effective syntax and masks a standard syntax in this Registry.
func (r *Registry) Unregister(u *uid.UID) (*Syntax, bool) {
	if u == nil {
		return nil, false
	}
	key := normalizeTransferUID(u.UID())

	r.mu.Lock()
	defer r.mu.Unlock()

	previous := r.queryLocked(key)
	if previous == nil {
		return nil, false
	}
	delete(r.overlay, key)
	r.masked[key] = struct{}{}
	return previous, true
}

// Query returns the transfer syntax registered under the given UID, or nil if not found.
func (r *Registry) Query(u *uid.UID) *Syntax {
	if u == nil {
		return nil
	}
	key := normalizeTransferUID(u.UID())

	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.queryLocked(key)
}

// List returns an independent UID-sorted snapshot of all effective syntaxes.
func (r *Registry) List() []*Syntax {
	r.mu.RLock()
	defer r.mu.RUnlock()
	values := make(map[string]*Syntax, len(standardTransferSyntaxes)+len(r.overlay))
	for _, ts := range standardTransferSyntaxes {
		if _, masked := r.masked[ts.UID().UID()]; !masked {
			values[ts.UID().UID()] = ts
		}
	}
	for key, ts := range r.overlay {
		values[key] = ts
	}
	entries := make([]*Syntax, 0, len(values))
	for _, ts := range values {
		entries = append(entries, ts)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].UID().UID() < entries[j].UID().UID()
	})
	return entries
}

func (r *Registry) queryLocked(value string) *Syntax {
	if ts := r.overlay[value]; ts != nil {
		return ts
	}
	if _, masked := r.masked[value]; masked {
		return nil
	}
	return standardTransferSyntaxIndex[value]
}

func syntaxRegistryKey(ts *Syntax) (string, error) {
	if ts == nil || ts.UID() == nil {
		return "", ErrNilSyntax
	}
	key := normalizeTransferUID(ts.UID().UID())
	if ts.UID().Type() != uid.TypeTransferSyntax || !uid.IsValid(key) {
		return "", fmt.Errorf("%w: %s", ErrInvalidSyntax, ts.UID().UID())
	}
	return key, nil
}

func normalizeTransferUID(value string) string {
	return strings.TrimRight(value, " \x00")
}

func buildSyntaxIndex(values []*Syntax) map[string]*Syntax {
	index := make(map[string]*Syntax, len(values))
	for _, value := range values {
		index[value.UID().UID()] = value
	}
	return index
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

// Build returns the constructed TransferSyntax without registering it.
func (b *Builder) Build() *Syntax {
	return b.ts
}
