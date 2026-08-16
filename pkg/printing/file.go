// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"

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
