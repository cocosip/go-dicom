// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"fmt"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// CodecRegistry manages codecs for different transfer syntaxes.
type CodecRegistry struct {
	mu     sync.RWMutex
	codecs map[string]Codec // key: transfer syntax UID
}

// NewCodecRegistry creates a new codec registry.
func NewCodecRegistry() *CodecRegistry {
	return &CodecRegistry{
		codecs: make(map[string]Codec),
	}
}

// RegisterCodec registers a codec for a transfer syntax.
func (r *CodecRegistry) RegisterCodec(ts *transfer.TransferSyntax, codec Codec) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.codecs[ts.UID().UID()] = codec
}

// GetCodec retrieves a codec for a transfer syntax.
// Returns the codec and true if found, nil and false otherwise.
func (r *CodecRegistry) GetCodec(ts *transfer.TransferSyntax) (Codec, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	codec, exists := r.codecs[ts.UID().UID()]
	return codec, exists
}

// HasCodec checks if a codec is registered for a transfer syntax.
func (r *CodecRegistry) HasCodec(ts *transfer.TransferSyntax) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.codecs[ts.UID().UID()]
	return exists
}

// UnregisterCodec removes a codec from the registry.
func (r *CodecRegistry) UnregisterCodec(ts *transfer.TransferSyntax) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.codecs, ts.UID().UID())
}

// ListCodecs returns a list of all registered transfer syntax UIDs.
func (r *CodecRegistry) ListCodecs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	uids := make([]string, 0, len(r.codecs))
	for uid := range r.codecs {
		uids = append(uids, uid)
	}
	return uids
}

// Global codec registry
var (
	globalRegistry     *CodecRegistry
	globalRegistryOnce sync.Once
)

// GetGlobalRegistry returns the global codec registry singleton.
func GetGlobalRegistry() *CodecRegistry {
	globalRegistryOnce.Do(func() {
		globalRegistry = NewCodecRegistry()
		// Register built-in codecs
		registerBuiltinCodecs(globalRegistry)
	})
	return globalRegistry
}

// registerBuiltinCodecs registers the built-in codecs.
func registerBuiltinCodecs(registry *CodecRegistry) {
	// Register Native codec for uncompressed transfer syntaxes

	// Explicit VR Little Endian
	explicitLECodec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
	registry.RegisterCodec(transfer.ExplicitVRLittleEndian, explicitLECodec)

	// Implicit VR Little Endian
	implicitLECodec := NewNativeCodec(transfer.ImplicitVRLittleEndian, false)
	registry.RegisterCodec(transfer.ImplicitVRLittleEndian, implicitLECodec)

	// Explicit VR Big Endian
	explicitBECodec := NewNativeCodec(transfer.ExplicitVRBigEndian, true)
	registry.RegisterCodec(transfer.ExplicitVRBigEndian, explicitBECodec)

	// Register RLE codec
	rleCodec := NewRLECodec()
	registry.RegisterCodec(transfer.RLELossless, rleCodec)

	// Future: Register JPEG, JPEG 2000, JPEG-LS codecs
}

// TranscoderManager provides high-level transcoding operations.
type TranscoderManager struct {
	registry *CodecRegistry
}

// NewTranscoderManager creates a new transcoder manager.
func NewTranscoderManager(registry *CodecRegistry) *TranscoderManager {
	if registry == nil {
		registry = GetGlobalRegistry()
	}

	return &TranscoderManager{
		registry: registry,
	}
}

// CreateTranscoder creates a transcoder for converting between two transfer syntaxes.
func (tm *TranscoderManager) CreateTranscoder(
	inputTS, outputTS *transfer.TransferSyntax,
	opts ...TranscoderOption,
) (*Transcoder, error) {
	// Check if required codecs are available
	if inputTS.IsEncapsulated() && !tm.registry.HasCodec(inputTS) {
		return nil, fmt.Errorf("no codec available for input transfer syntax: %s", inputTS.UID().UID())
	}

	if outputTS.IsEncapsulated() && !tm.registry.HasCodec(outputTS) {
		return nil, fmt.Errorf("no codec available for output transfer syntax: %s", outputTS.UID().UID())
	}

	// Add registry to options
	opts = append(opts, WithCodecRegistry(tm.registry))

	// Create transcoder
	return NewTranscoder(inputTS, outputTS, opts...), nil
}

// CanTranscode checks if transcoding between two transfer syntaxes is supported.
func (tm *TranscoderManager) CanTranscode(inputTS, outputTS *transfer.TransferSyntax) bool {
	// Uncompressed to uncompressed is always supported
	if !inputTS.IsEncapsulated() && !outputTS.IsEncapsulated() {
		return true
	}

	// Check if required codecs are available
	if inputTS.IsEncapsulated() && !tm.registry.HasCodec(inputTS) {
		return false
	}

	if outputTS.IsEncapsulated() && !tm.registry.HasCodec(outputTS) {
		return false
	}

	return true
}

// GetDefaultManager returns a transcoder manager with the global registry.
func GetDefaultManager() *TranscoderManager {
	return NewTranscoderManager(GetGlobalRegistry())
}
