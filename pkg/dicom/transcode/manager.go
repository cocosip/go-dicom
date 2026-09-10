// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"errors"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
)

var (
	// ErrNilRegistry identifies a missing codec Registry dependency.
	ErrNilRegistry = errors.New("codec registry must not be nil")
	// ErrTransferSyntaxRequired identifies a missing input or output syntax.
	ErrTransferSyntaxRequired = errors.New("transfer syntax must not be nil")
	// ErrCodecUnavailable identifies a required compressed codec that is not registered.
	ErrCodecUnavailable = errors.New("codec unavailable")
)

// Manager creates Dataset transcoders from one explicit codec registry.
type Manager struct {
	registry *codec.Registry
}

// NewManager creates a Manager bound to registry.
func NewManager(registry *codec.Registry) (*Manager, error) {
	if registry == nil {
		return nil, ErrNilRegistry
	}
	return &Manager{registry: registry}, nil
}

// CanTranscode reports whether all required compressed codecs are registered.
func (m *Manager) CanTranscode(input, output *transfer.Syntax) bool {
	if m == nil || m.registry == nil || input == nil || output == nil {
		return false
	}
	if input.IsEncapsulated() {
		if _, found := m.registry.Lookup(input); !found {
			return false
		}
	}
	if output.IsEncapsulated() {
		if _, found := m.registry.Lookup(output); !found {
			return false
		}
	}
	return true
}

// NewTranscoder creates a Transcoder using codecs from the Manager registry.
func (m *Manager) NewTranscoder(input, output *transfer.Syntax, options ...Option) (*Transcoder, error) {
	if m == nil || m.registry == nil {
		return nil, ErrNilRegistry
	}
	if input == nil || output == nil {
		return nil, ErrTransferSyntaxRequired
	}
	transcoder := &Transcoder{
		manager:       m,
		inputSyntax:   input,
		outputSyntax:  output,
		strictDICOMVR: true,
	}
	for _, option := range options {
		if option != nil {
			option(transcoder)
		}
	}
	if input.IsEncapsulated() {
		registered, found := m.registry.Lookup(input)
		if !found {
			return nil, fmt.Errorf("%w for input transfer syntax %s", ErrCodecUnavailable, input.UID().UID())
		}
		transcoder.inputCodec = registered
	}
	if output.IsEncapsulated() {
		registered, found := m.registry.Lookup(output)
		if !found {
			return nil, fmt.Errorf("%w for output transfer syntax %s", ErrCodecUnavailable, output.UID().UID())
		}
		transcoder.outputCodec = registered
	}
	return transcoder, nil
}
