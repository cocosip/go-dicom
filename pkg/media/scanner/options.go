// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// SymlinkPolicy controls how filesystem links are handled.
type SymlinkPolicy uint8

const (
	// SkipSymlinks ignores every symbolic link.
	SkipSymlinks SymlinkPolicy = iota
	// FollowFileSymlinksWithinRoot follows regular-file links whose target remains under the root.
	FollowFileSymlinksWithinRoot
)

type config struct {
	recursive             bool
	workers               int
	stopOnError           bool
	symlinkPolicy         SymlinkPolicy
	assumedTransferSyntax *transfer.Syntax
}

// Option configures a Scanner.
type Option func(*config)

// WithRecursive enables or disables directory recursion.
func WithRecursive(recursive bool) Option {
	return func(config *config) {
		config.recursive = recursive
	}
}

// WithWorkers sets the maximum number of concurrent file parsers.
func WithWorkers(workers int) Option {
	return func(config *config) {
		config.workers = workers
	}
}

// WithStopOnError stops after the first invalid or unreadable result in discovery order.
func WithStopOnError(stop bool) Option {
	return func(config *config) {
		config.stopOnError = stop
	}
}

// WithSymlinkPolicy selects how symbolic links are handled.
func WithSymlinkPolicy(policy SymlinkPolicy) Option {
	return func(config *config) {
		config.symlinkPolicy = policy
	}
}

// WithAssumedTransferSyntax configures parsing for datasets without File Meta Information.
func WithAssumedTransferSyntax(syntax *transfer.Syntax) Option {
	return func(config *config) {
		config.assumedTransferSyntax = syntax
	}
}

func newConfig(options ...Option) (config, error) {
	value := config{
		recursive:     true,
		workers:       1,
		symlinkPolicy: SkipSymlinks,
	}
	for _, option := range options {
		if option != nil {
			option(&value)
		}
	}
	if value.workers <= 0 {
		return config{}, fmt.Errorf("scanner workers must be positive")
	}
	if value.symlinkPolicy != SkipSymlinks && value.symlinkPolicy != FollowFileSymlinksWithinRoot {
		return config{}, fmt.Errorf("unsupported scanner symlink policy %d", value.symlinkPolicy)
	}
	return value, nil
}
