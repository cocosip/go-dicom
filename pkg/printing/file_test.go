// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
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

func TestFilmBoxFolderRoundTrip(t *testing.T) {
	folder := filepath.Join(t.TempDir(), "film-box")
	session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.310", false)
	filmBox := NewFilmBox("2.25.311", testStandardTwoByOne)
	filmBox.AnnotationDisplayFormatID = "ANNOTATION"
	session.AddFilmBox(filmBox)

	for index, instanceUID := range []string{"2.25.312", "2.25.313"} {
		imageBox := NewImageBox(instanceUID, false)
		imageBox.ImageBoxPosition = uint16(index + 1)
		imageItem := dataset.New()
		for _, imageElement := range []element.Element{
			element.NewUnsignedShort(tag.Rows, []uint16{1}),
			element.NewUnsignedShort(tag.Columns, []uint16{2}),
			element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
			element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}),
			element.NewOtherByte(tag.PixelData, []byte{byte(index + 1), byte(index + 2)}),
		} {
			if err := imageItem.Add(imageElement); err != nil {
				t.Fatalf("image item Add() error = %v", err)
			}
		}
		if err := imageBox.SetImageSequence(imageItem); err != nil {
			t.Fatalf("SetImageSequence() error = %v", err)
		}
		filmBox.AddImageBox(imageBox)
	}

	if err := filmBox.Save(folder); err != nil {
		t.Fatalf("FilmBox.Save() error = %v", err)
	}
	for _, relativePath := range []string{
		"FilmBox.dcm",
		filepath.Join("Images", "I000001.dcm"),
		filepath.Join("Images", "I000002.dcm"),
	} {
		if info, err := os.Stat(filepath.Join(folder, relativePath)); err != nil || !info.Mode().IsRegular() {
			t.Fatalf("saved file %q = %#v, %v", relativePath, info, err)
		}
	}

	parsedImage, err := parser.ParseFile(filepath.Join(folder, "Images", "I000001.dcm"))
	if err != nil {
		t.Fatalf("parse saved Image Box: %v", err)
	}
	if parsedImage.TransferSyntax != transfer.ExplicitVRLittleEndian {
		t.Errorf("Image Box transfer syntax = %v", parsedImage.TransferSyntax)
	}
	sequence, err := parsedImage.Dataset.GetSequence(tag.BasicGrayscaleImageSequence)
	if err != nil || sequence.Count() != 1 {
		t.Fatalf("saved image sequence = %#v, %v", sequence, err)
	}
	if rows, err := sequence.GetItem(0).GetUInt16(tag.Rows, 0); err != nil || rows != 1 {
		t.Errorf("saved image Rows = %d, %v", rows, err)
	}
	if pixels := sequence.GetItem(0).GetOrNil(tag.PixelData); pixels == nil || len(pixels.Buffer().Data()) != 2 {
		t.Errorf("saved image Pixel Data = %#v", pixels)
	}

	loadedSession := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.310", false)
	loaded, err := LoadFilmBox(loadedSession, folder)
	if err != nil {
		t.Fatalf("LoadFilmBox() error = %v", err)
	}
	if loaded.FilmSession() != loadedSession {
		t.Error("loaded Film Box parent was not repaired")
	}
	if loaded.SOPInstanceUID != filmBox.SOPInstanceUID || loaded.AnnotationDisplayFormatID != filmBox.AnnotationDisplayFormatID {
		t.Errorf("loaded Film Box = %#v", loaded)
	}
	if len(loaded.BasicImageBoxes) != 2 {
		t.Fatalf("loaded Image Box count = %d", len(loaded.BasicImageBoxes))
	}
	for index, imageBox := range loaded.BasicImageBoxes {
		if imageBox.SOPInstanceUID != filmBox.BasicImageBoxes[index].SOPInstanceUID {
			t.Errorf("loaded Image Box %d UID = %q", index, imageBox.SOPInstanceUID)
		}
		if imageBox.FilmBox() != loaded {
			t.Errorf("loaded Image Box %d parent was not repaired", index)
		}
		item, err := imageBox.ImageSequence()
		if err != nil || item == nil {
			t.Fatalf("loaded Image Box %d sequence = %#v, %v", index, item, err)
		}
		if columns, err := item.GetUInt16(tag.Columns, 0); err != nil || columns != 2 {
			t.Errorf("loaded Image Box %d Columns = %d, %v", index, columns, err)
		}
	}
}

func TestFilmBoxSaveRemovesStaleManagedImageFiles(t *testing.T) {
	folder := filepath.Join(t.TempDir(), "film-box")
	session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.314", false)
	filmBox := NewFilmBox("2.25.315", testStandardTwoByOne)
	session.AddFilmBox(filmBox)
	for index, instanceUID := range []string{"2.25.316", "2.25.317"} {
		imageBox := NewImageBox(instanceUID, false)
		imageBox.ImageBoxPosition = uint16(index + 1)
		filmBox.AddImageBox(imageBox)
	}
	if err := filmBox.Save(folder); err != nil {
		t.Fatalf("first FilmBox.Save() error = %v", err)
	}
	for _, name := range []string{"notes.txt", "X000002.dcm", "I000002.dcm.bak"} {
		if err := os.WriteFile(filepath.Join(folder, "Images", name), []byte("preserve"), 0o600); err != nil {
			t.Fatalf("os.WriteFile(%q) error = %v", name, err)
		}
	}

	filmBox.BasicImageBoxes = filmBox.BasicImageBoxes[:1]
	filmBox.ImageDisplayFormat = testStandardOneByOne
	if err := filmBox.Save(folder); err != nil {
		t.Fatalf("second FilmBox.Save() error = %v", err)
	}
	stalePath := filepath.Join(folder, "Images", "I000002.dcm")
	if _, err := os.Stat(stalePath); !os.IsNotExist(err) {
		t.Fatalf("stale Image Box file still exists: %v", err)
	}
	for _, name := range []string{"notes.txt", "X000002.dcm", "I000002.dcm.bak"} {
		if _, err := os.Stat(filepath.Join(folder, "Images", name)); err != nil {
			t.Errorf("unmanaged file %q was not preserved: %v", name, err)
		}
	}
	if err := os.Remove(filepath.Join(folder, "Images", "X000002.dcm")); err != nil {
		t.Fatalf("os.Remove(X000002.dcm) error = %v", err)
	}

	loaded, err := LoadFilmBox(NewFilmSession(basicFilmSessionSOPClassUID, "2.25.314", false), folder)
	if err != nil {
		t.Fatalf("LoadFilmBox() error = %v", err)
	}
	if len(loaded.BasicImageBoxes) != 1 {
		t.Fatalf("loaded Image Box count = %d, want 1", len(loaded.BasicImageBoxes))
	}
}

func TestFilmBoxSaveDoesNotRemoveStaleFilesWhenCurrentWriteFails(t *testing.T) {
	folder := filepath.Join(t.TempDir(), "film-box")
	imagesFolder := filepath.Join(folder, "Images")
	if err := os.MkdirAll(filepath.Join(imagesFolder, "I000001.dcm"), 0o755); err != nil {
		t.Fatalf("os.MkdirAll() error = %v", err)
	}
	stalePath := filepath.Join(imagesFolder, "I000002.dcm")
	if err := os.WriteFile(stalePath, []byte("stale"), 0o600); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}

	filmBox := NewFilmBox("2.25.318", testStandardOneByOne)
	filmBox.AddImageBox(NewImageBox("2.25.319", false))
	if err := filmBox.Save(folder); err == nil {
		t.Fatal("FilmBox.Save() succeeded with a directory at the current Image Box path")
	}
	if _, err := os.Stat(stalePath); err != nil {
		t.Fatalf("stale Image Box file was removed after a current write failure: %v", err)
	}
}

func TestFilmBoxFolderPersistenceErrors(t *testing.T) {
	t.Run("nil Film Box", func(t *testing.T) {
		var filmBox *FilmBox
		if err := filmBox.Save(filepath.Join(t.TempDir(), "box")); err == nil {
			t.Fatal("nil FilmBox.Save() returned nil")
		}
	})

	t.Run("nil Film Session", func(t *testing.T) {
		if _, err := LoadFilmBox(nil, t.TempDir()); err == nil {
			t.Fatal("LoadFilmBox() accepted a nil Film Session")
		}
	})

	t.Run("missing FilmBox file", func(t *testing.T) {
		session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.320", false)
		if _, err := LoadFilmBox(session, t.TempDir()); err == nil {
			t.Fatal("LoadFilmBox() accepted a missing FilmBox.dcm")
		}
	})

	t.Run("malformed Image Box file", func(t *testing.T) {
		folder := filepath.Join(t.TempDir(), "box")
		session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.321", false)
		filmBox := NewFilmBox("2.25.322", testStandardOneByOne)
		session.AddFilmBox(filmBox)
		filmBox.AddImageBox(NewImageBox("2.25.323", false))
		if err := filmBox.Save(folder); err != nil {
			t.Fatalf("FilmBox.Save() error = %v", err)
		}
		imagePath := filepath.Join(folder, "Images", "I000001.dcm")
		if err := os.WriteFile(imagePath, []byte("not dicom"), 0o600); err != nil {
			t.Fatalf("os.WriteFile() error = %v", err)
		}
		if _, err := LoadFilmBox(session, folder); err == nil {
			t.Fatal("LoadFilmBox() accepted a malformed Image Box")
		}
	})

	t.Run("folder conflicts with file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "box")
		if err := os.WriteFile(path, []byte("file"), 0o600); err != nil {
			t.Fatalf("os.WriteFile() error = %v", err)
		}
		if err := NewFilmBox("2.25.324", testStandardOneByOne).Save(path); err == nil {
			t.Fatal("FilmBox.Save() accepted a file as its target folder")
		}
	})
}
