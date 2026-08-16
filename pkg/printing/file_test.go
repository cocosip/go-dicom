// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

func TestFilmSessionFileRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "session.dcm")
	want := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.301", false)
	want.FilmDestination = FilmDestinationMagazine
	want.FilmSessionLabel = "night shift"
	want.MemoryAllocation = 8192
	want.MediumType = MediumTypePaper
	want.PrintPriority = PrintPriorityLow
	want.NumberOfCopies = 2

	if err := want.Save(path); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	parsed, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("parser.ParseFile() error = %v", err)
	}
	if parsed.Format != parser.FormatDICOM3 {
		t.Errorf("parsed format = %s, want DICOM3", parsed.Format)
	}
	if parsed.TransferSyntax != transfer.ExplicitVRLittleEndian {
		t.Errorf("transfer syntax = %v, want Explicit VR Little Endian", parsed.TransferSyntax)
	}
	if got, ok := parsed.FileMetaInformation.MediaStorageSOPClassUID(); !ok || got != want.SOPClassUID {
		t.Errorf("MediaStorageSOPClassUID = %q, %v; want %q, true", got, ok, want.SOPClassUID)
	}
	if got, ok := parsed.FileMetaInformation.MediaStorageSOPInstanceUID(); !ok || got != want.SOPInstanceUID {
		t.Errorf("MediaStorageSOPInstanceUID = %q, %v; want %q, true", got, ok, want.SOPInstanceUID)
	}

	got, err := LoadFilmSession(path)
	if err != nil {
		t.Fatalf("LoadFilmSession() error = %v", err)
	}
	if got.SOPClassUID != want.SOPClassUID || got.SOPInstanceUID != want.SOPInstanceUID ||
		got.FilmDestination != want.FilmDestination || got.FilmSessionLabel != want.FilmSessionLabel ||
		got.MemoryAllocation != want.MemoryAllocation || got.MediumType != want.MediumType ||
		got.PrintPriority != want.PrintPriority || got.NumberOfCopies != want.NumberOfCopies {
		t.Fatalf("loaded session mismatch: got %#v, want %#v", got, want)
	}
}

func TestLoadFilmSessionErrors(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.dcm")
	if _, err := LoadFilmSession(missing); err == nil {
		t.Fatal("LoadFilmSession() accepted a missing file")
	}

	malformed := filepath.Join(t.TempDir(), "malformed.dcm")
	if err := os.WriteFile(malformed, []byte("not dicom"), 0o600); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}
	if _, err := LoadFilmSession(malformed); err == nil {
		t.Fatal("LoadFilmSession() accepted a malformed file")
	}
}

func TestNilFilmSessionSaveFails(t *testing.T) {
	var session *FilmSession
	if err := session.Save(filepath.Join(t.TempDir(), "nil.dcm")); err == nil {
		t.Fatal("nil FilmSession.Save() returned nil")
	}
}
