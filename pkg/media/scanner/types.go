// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
)

// ResultKind classifies one delivered filesystem scan result.
type ResultKind uint8

const (
	// ResultDICOM identifies a structurally parseable DICOM metadata prefix.
	ResultDICOM ResultKind = iota
	// ResultInvalid identifies a file that cannot be parsed as DICOM metadata.
	ResultInvalid
	// ResultReadError identifies a filesystem traversal, open, or read failure.
	ResultReadError
)

// Result describes one path delivered by a Scan.
type Result struct {
	Root         string
	Path         string
	RelativePath string
	Kind         ResultKind
	File         *parser.ParseResult
	Err          error
}

// Summary contains deterministic counts for delivered and skipped entries.
type Summary struct {
	Results           int
	DICOMFiles        int
	InvalidFiles      int
	ReadErrors        int
	SkippedSymlinks   int
	SkippedNonRegular int
}

// Handler processes one result. Returning an error stops the scan.
type Handler func(Result) error

// ScanError reports the result that stopped a stop-on-error scan.
type ScanError struct {
	Result Result
}

func (e *ScanError) Error() string {
	if e == nil {
		return "scan stopped"
	}
	return fmt.Sprintf("scan stopped at %s: %v", e.Result.Path, e.Result.Err)
}

// Unwrap returns the path-specific error that stopped the scan.
func (e *ScanError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Result.Err
}
