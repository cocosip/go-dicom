// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"sync"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

var (
	// ErrNilCodec identifies an attempted nil codec registration.
	ErrNilCodec = errors.New("codec must not be nil")
	// ErrNilTransferSyntax identifies a codec without a transfer syntax or UID.
	ErrNilTransferSyntax = errors.New("codec transfer syntax must not be nil")
	// ErrCodecAlreadyRegistered identifies a duplicate Register call.
	ErrCodecAlreadyRegistered = errors.New("codec already registered")
)

// Registry manages codecs for different transfer syntaxes.
type Registry struct {
	mu     sync.RWMutex
	codecs map[string]Codec // key: transfer syntax UID
}

func newEmptyRegistry() *Registry {
	return &Registry{
		codecs: make(map[string]Codec),
	}
}

// NewRegistry creates an isolated registry containing the native codecs.
func NewRegistry() *Registry {
	registry := newEmptyRegistry()
	registerBuiltinCodecs(registry)
	return registry
}

// Register adds a codec without replacing an implementation for the same UID.
func (r *Registry) Register(c Codec) error {
	uid, err := registryCodecUID(c)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.codecs[uid]; exists {
		return fmt.Errorf("%w: %s", ErrCodecAlreadyRegistered, uid)
	}
	r.codecs[uid] = c
	return nil
}

// Replace inserts a codec and returns the previous implementation, if any.
func (r *Registry) Replace(c Codec) (Codec, error) {
	uid, err := registryCodecUID(c)
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	previous := r.codecs[uid]
	r.codecs[uid] = c
	return previous, nil
}

// Unregister removes and returns the codec for syntax.
func (r *Registry) Unregister(syntax *transfer.Syntax) (Codec, bool) {
	uid, ok := registrySyntaxUID(syntax)
	if !ok {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	previous, found := r.codecs[uid]
	delete(r.codecs, uid)
	return previous, found
}

// Lookup returns the codec registered for syntax.
func (r *Registry) Lookup(syntax *transfer.Syntax) (Codec, bool) {
	uid, ok := registrySyntaxUID(syntax)
	if !ok {
		return nil, false
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, found := r.codecs[uid]
	return c, found
}

// List returns a UID-sorted snapshot of registered codecs.
func (r *Registry) List() []Codec {
	r.mu.RLock()
	entries := make([]struct {
		uid   string
		codec Codec
	}, 0, len(r.codecs))
	for uid, c := range r.codecs {
		entries = append(entries, struct {
			uid   string
			codec Codec
		}{uid: uid, codec: c})
	}
	r.mu.RUnlock()

	sort.Slice(entries, func(i, j int) bool { return entries[i].uid < entries[j].uid })
	result := make([]Codec, len(entries))
	for index, entry := range entries {
		result[index] = entry.codec
	}
	return result
}

func registryCodecUID(c Codec) (string, error) {
	if codecIsNil(c) {
		return "", ErrNilCodec
	}
	syntax := c.TransferSyntax()
	uid, ok := registrySyntaxUID(syntax)
	if !ok {
		return "", ErrNilTransferSyntax
	}
	return uid, nil
}

func registrySyntaxUID(syntax *transfer.Syntax) (string, bool) {
	if syntax == nil || syntax.UID() == nil || syntax.UID().UID() == "" {
		return "", false
	}
	return syntax.UID().UID(), true
}

func codecIsNil(c Codec) bool {
	if c == nil {
		return true
	}
	value := reflect.ValueOf(c)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// Global codec registry
var (
	globalRegistry     *Registry
	globalRegistryOnce sync.Once
)

// GlobalRegistry returns the mutable, process-wide codec plugin registry.
func GlobalRegistry() *Registry {
	globalRegistryOnce.Do(func() {
		globalRegistry = NewRegistry()
	})
	return globalRegistry
}

// registerBuiltinCodecs registers the built-in codecs.
func registerBuiltinCodecs(registry *Registry) {
	// Register Native codec for uncompressed transfer syntaxes

	// Explicit VR Little Endian
	explicitLECodec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
	if err := registry.Register(explicitLECodec); err != nil {
		panic(err)
	}

	// Implicit VR Little Endian
	implicitLECodec := NewNativeCodec(transfer.ImplicitVRLittleEndian, false)
	if err := registry.Register(implicitLECodec); err != nil {
		panic(err)
	}

	// Explicit VR Big Endian
	explicitBECodec := NewNativeCodec(transfer.ExplicitVRBigEndian, true)
	if err := registry.Register(explicitBECodec); err != nil {
		panic(err)
	}

	// Compressed codecs are supplied by go-dicom-codecs.
}
