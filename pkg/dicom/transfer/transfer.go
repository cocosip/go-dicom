// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package transfer implements DICOM transfer syntax functionality.
package transfer

import (
	"fmt"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// TransferSyntax represents a DICOM transfer syntax.
//
// A transfer syntax defines how DICOM data is encoded, including:
// - Value Representation (VR) encoding (Explicit or Implicit)
// - Byte order (Little Endian or Big Endian)
// - Compression (e.g., JPEG, JPEG 2000, RLE)
type TransferSyntax struct {
	uid                   *uid.UID
	isRetired             bool
	isExplicitVR          bool
	isEncapsulated        bool
	isLossy               bool
	lossyCompressionMethod string
	isDeflate             bool
	endian                endian.Endian
	swapPixelData         bool
}

var (
	// Global registry of transfer syntaxes
	registry = make(map[string]*TransferSyntax)
	mu       sync.RWMutex
)

// New creates a new TransferSyntax with the given UID.
func New(u *uid.UID) *TransferSyntax {
	return &TransferSyntax{
		uid:           u,
		isRetired:     u.IsRetired(),
		isExplicitVR:  false,
		isEncapsulated: false,
		isLossy:       false,
		isDeflate:     false,
		endian:        endian.Little,
		swapPixelData: false,
	}
}

// UID returns the unique identifier of the transfer syntax.
func (ts *TransferSyntax) UID() *uid.UID {
	return ts.uid
}

// IsRetired returns true if the transfer syntax is retired.
func (ts *TransferSyntax) IsRetired() bool {
	return ts.isRetired
}

// IsExplicitVR returns true if the transfer syntax uses explicit VR.
func (ts *TransferSyntax) IsExplicitVR() bool {
	return ts.isExplicitVR
}

// IsEncapsulated returns true if the transfer syntax encapsulates pixel data.
func (ts *TransferSyntax) IsEncapsulated() bool {
	return ts.isEncapsulated
}

// IsLossy returns true if the transfer syntax uses lossy compression.
func (ts *TransferSyntax) IsLossy() bool {
	return ts.isLossy
}

// LossyCompressionMethod returns the lossy compression method identifier.
func (ts *TransferSyntax) LossyCompressionMethod() string {
	return ts.lossyCompressionMethod
}

// IsDeflate returns true if the transfer syntax uses deflate compression.
func (ts *TransferSyntax) IsDeflate() bool {
	return ts.isDeflate
}

// Endian returns the byte order of the transfer syntax.
func (ts *TransferSyntax) Endian() endian.Endian {
	return ts.endian
}

// SwapPixelData returns true if pixel data requires byte swapping.
func (ts *TransferSyntax) SwapPixelData() bool {
	return ts.swapPixelData
}

// String returns the name of the transfer syntax.
func (ts *TransferSyntax) String() string {
	return ts.uid.Name()
}

// Equals checks if two transfer syntaxes are equal.
func (ts *TransferSyntax) Equals(other *TransferSyntax) bool {
	if other == nil {
		return false
	}
	return ts.uid.Equals(other.uid)
}

// Parse parses a transfer syntax UID string and returns the corresponding TransferSyntax.
func Parse(uidString string) (*TransferSyntax, error) {
	u := uid.Parse(uidString, "", uid.TypeTransferSyntax)
	return Lookup(u)
}

// Lookup looks up a transfer syntax by UID.
//
// If the UID is registered, returns the registered transfer syntax.
// Otherwise, creates a new transfer syntax with default properties:
// - Explicit VR
// - Encapsulated
// - Little Endian
func Lookup(u *uid.UID) (*TransferSyntax, error) {
	if u == nil {
		return nil, fmt.Errorf("UID cannot be nil")
	}

	if u.Type() != uid.TypeTransferSyntax {
		return nil, fmt.Errorf("UID %s is not a transfer syntax type", u.UID())
	}

	mu.RLock()
	ts, found := registry[u.UID()]
	mu.RUnlock()

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

// Register registers a transfer syntax in the global registry.
func Register(ts *TransferSyntax) {
	if ts == nil || ts.uid == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()
	registry[ts.uid.UID()] = ts
}

// Unregister removes a transfer syntax from the global registry.
func Unregister(u *uid.UID) bool {
	if u == nil {
		return false
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := registry[u.UID()]
	if exists {
		delete(registry, u.UID())
	}
	return exists
}

// Query queries a transfer syntax by UID. Returns nil if not found.
func Query(u *uid.UID) *TransferSyntax {
	if u == nil {
		return nil
	}

	mu.RLock()
	defer mu.RUnlock()
	return registry[u.UID()]
}

// KnownEntries returns all registered transfer syntaxes.
func KnownEntries() []*TransferSyntax {
	mu.RLock()
	defer mu.RUnlock()

	entries := make([]*TransferSyntax, 0, len(registry))
	for _, ts := range registry {
		entries = append(entries, ts)
	}
	return entries
}

// Builder is a helper for constructing TransferSyntax instances with custom properties.
type Builder struct {
	ts *TransferSyntax
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

// Build returns the constructed TransferSyntax and registers it.
func (b *Builder) Build() *TransferSyntax {
	Register(b.ts)
	return b.ts
}
