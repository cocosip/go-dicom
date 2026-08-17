// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

// Save writes the Film Session attributes as a DICOM Part 10 file.
func (fs *FilmSession) Save(path string) error {
	if fs == nil {
		return fmt.Errorf("printing: cannot save a nil FilmSession")
	}
	ds, err := fs.ToDataset()
	if err != nil {
		return fmt.Errorf("printing: encode FilmSession: %w", err)
	}
	if err := writer.WriteFile(path, ds, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		return fmt.Errorf("printing: save FilmSession: %w", err)
	}
	return nil
}

// LoadFilmSession reads a Film Session from a DICOM Part 10 file.
func LoadFilmSession(path string) (*FilmSession, error) {
	result, err := parser.ParseFile(path)
	if err != nil {
		return nil, fmt.Errorf("printing: load FilmSession: %w", err)
	}
	if result == nil || result.Dataset == nil {
		return nil, fmt.Errorf("printing: load FilmSession: file has no Dataset")
	}

	sopClassUID, _ := result.FileMetaInformation.MediaStorageSOPClassUID()
	sopInstanceUID, _ := result.FileMetaInformation.MediaStorageSOPInstanceUID()
	filmSession, err := NewFilmSessionFromDataset(sopClassUID, sopInstanceUID, result.Dataset, false)
	if err != nil {
		return nil, fmt.Errorf("printing: decode FilmSession: %w", err)
	}
	return filmSession, nil
}

// Save writes a Film Box and its Image Boxes to a fo-dicom-compatible folder.
func (fb *FilmBox) Save(folder string) error {
	if fb == nil {
		return fmt.Errorf("printing: cannot save a nil FilmBox")
	}
	if err := os.MkdirAll(folder, 0o755); err != nil {
		return fmt.Errorf("printing: create FilmBox folder %q: %w", folder, err)
	}

	filmBoxDataset, err := fb.ToDataset()
	if err != nil {
		return fmt.Errorf("printing: encode FilmBox: %w", err)
	}
	filmBoxPath := filepath.Join(folder, "FilmBox.dcm")
	if err := writeDatasetAtomically(filmBoxPath, filmBoxDataset); err != nil {
		return fmt.Errorf("printing: save FilmBox %q: %w", filmBoxPath, err)
	}

	imagesFolder := filepath.Join(folder, "Images")
	if err := os.MkdirAll(imagesFolder, 0o755); err != nil {
		return fmt.Errorf("printing: create Image Box folder %q: %w", imagesFolder, err)
	}
	for index, imageBox := range fb.BasicImageBoxes {
		if imageBox == nil {
			return fmt.Errorf("printing: save Image Box %d: nil ImageBox", index+1)
		}
		imageDataset, err := imageBox.ToDataset()
		if err != nil {
			return fmt.Errorf("printing: encode Image Box %d: %w", index+1, err)
		}
		imagePath := filepath.Join(imagesFolder, fmt.Sprintf("I%06d.dcm", index+1))
		if err := writeDatasetAtomically(imagePath, imageDataset); err != nil {
			return fmt.Errorf("printing: save Image Box %d %q: %w", index+1, imagePath, err)
		}
	}
	if err := removeStaleImageBoxFiles(imagesFolder, len(fb.BasicImageBoxes)); err != nil {
		return fmt.Errorf("printing: remove stale Image Box files from %q: %w", imagesFolder, err)
	}
	return nil
}

func removeStaleImageBoxFiles(folder string, keep int) error {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || len(name) != len("I000001.dcm") || name[0] != 'I' || !strings.EqualFold(name[7:], ".dcm") {
			continue
		}
		index, err := strconv.Atoi(name[1:7])
		if err != nil || index <= keep {
			continue
		}
		if err := os.Remove(filepath.Join(folder, name)); err != nil {
			return err
		}
	}
	return nil
}

// LoadFilmBox reads a Film Box folder and attaches the loaded hierarchy to session.
func LoadFilmBox(session *FilmSession, folder string) (*FilmBox, error) {
	if session == nil {
		return nil, fmt.Errorf("printing: FilmSession is required to load a FilmBox")
	}
	filmBoxPath := filepath.Join(folder, "FilmBox.dcm")
	parsedFilmBox, err := parser.ParseFile(filmBoxPath)
	if err != nil {
		return nil, fmt.Errorf("printing: load FilmBox %q: %w", filmBoxPath, err)
	}
	if parsedFilmBox == nil || parsedFilmBox.Dataset == nil {
		return nil, fmt.Errorf("printing: load FilmBox %q: file has no Dataset", filmBoxPath)
	}
	filmBoxInstanceUID := ""
	if parsedFilmBox.FileMetaInformation != nil {
		filmBoxInstanceUID, _ = parsedFilmBox.FileMetaInformation.MediaStorageSOPInstanceUID()
	}
	filmBox, err := NewFilmBoxFromDataset(filmBoxInstanceUID, parsedFilmBox.Dataset)
	if err != nil {
		return nil, fmt.Errorf("printing: decode FilmBox %q: %w", filmBoxPath, err)
	}
	filmBox.filmSession = session

	imagesFolder := filepath.Join(folder, "Images")
	entries, err := os.ReadDir(imagesFolder)
	if err != nil {
		return nil, fmt.Errorf("printing: read Image Box folder %q: %w", imagesFolder, err)
	}
	imagePaths := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".dcm") {
			continue
		}
		imagePaths = append(imagePaths, filepath.Join(imagesFolder, entry.Name()))
	}
	sort.Strings(imagePaths)
	for _, imagePath := range imagePaths {
		parsedImage, err := parser.ParseFile(imagePath)
		if err != nil {
			return nil, fmt.Errorf("printing: load Image Box %q: %w", imagePath, err)
		}
		if parsedImage == nil || parsedImage.Dataset == nil {
			return nil, fmt.Errorf("printing: load Image Box %q: file has no Dataset", imagePath)
		}
		imageClassUID, imageInstanceUID := "", ""
		if parsedImage.FileMetaInformation != nil {
			imageClassUID, _ = parsedImage.FileMetaInformation.MediaStorageSOPClassUID()
			imageInstanceUID, _ = parsedImage.FileMetaInformation.MediaStorageSOPInstanceUID()
		}
		imageBox, err := NewImageBoxFromDataset(imageInstanceUID, parsedImage.Dataset, imageClassUID == SOPClassColorImageBox)
		if err != nil {
			return nil, fmt.Errorf("printing: decode Image Box %q: %w", imagePath, err)
		}
		filmBox.AddImageBox(imageBox)
	}
	return filmBox, nil
}

func writeDatasetAtomically(path string, ds *dataset.Dataset) error {
	temporary, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer func() { _ = os.Remove(temporaryPath) }()

	writeErr := writer.Write(temporary, ds, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian))
	closeErr := temporary.Close()
	if writeErr != nil || closeErr != nil {
		return errors.Join(writeErr, closeErr)
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return err
	}
	return nil
}
