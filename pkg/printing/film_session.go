// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// FilmDestination represents where the exposed film is stored/processed
type FilmDestination string

const (
	// FilmDestinationMagazine - film stored in magazine
	FilmDestinationMagazine FilmDestination = "MAGAZINE"
	// FilmDestinationProcessor - film developed in processor
	FilmDestinationProcessor FilmDestination = "PROCESSOR"
	// FilmDestinationBin - film deposited in sorter bin (BIN_1, BIN_2, etc.)
)

// MediumType represents the type of medium for printing
type MediumType string

const (
	// MediumTypePaper - paper medium
	MediumTypePaper MediumType = "PAPER"
	// MediumTypeClearFilm - clear film
	MediumTypeClearFilm MediumType = "CLEAR FILM"
	// MediumTypeBlueFilm - blue film
	MediumTypeBlueFilm MediumType = "BLUE FILM"
	// MediumTypeMammoClearFilm - mammography clear film
	MediumTypeMammoClearFilm MediumType = "MAMMO CLEAR FILM"
	// MediumTypeMammoBlueFilm - mammography blue film
	MediumTypeMammoBlueFilm MediumType = "MAMMO BLUE FILM"
)

// PrintPriority represents the priority of the print job
type PrintPriority string

const (
	// PrintPriorityHigh - high priority
	PrintPriorityHigh PrintPriority = "HIGH"
	// PrintPriorityMed - medium priority
	PrintPriorityMed PrintPriority = "MED"
	// PrintPriorityLow - low priority
	PrintPriorityLow PrintPriority = "LOW"
)

// FilmSession represents a Basic Film Session in DICOM Print Management
//
// A Film Session defines the overall printing characteristics and contains
// one or more Film Boxes. It manages properties like film destination,
// medium type, print priority, and number of copies.
//
// Reference: DICOM Part 3, Section C.13.1
type FilmSession struct {
	// SOPClassUID is the Basic Film Session SOP Class UID
	SOPClassUID string

	// SOPInstanceUID is the Basic Film Session SOP Instance UID
	SOPInstanceUID string

	// FilmDestination specifies where the exposed film is stored
	// Defined Terms: MAGAZINE, PROCESSOR, BIN_i
	FilmDestination FilmDestination

	// FilmSessionLabel is a human readable label that identifies the film session
	FilmSessionLabel string

	// MemoryAllocation specifies the amount of memory allocated for the session
	// Value in bytes
	MemoryAllocation int

	// MediumType specifies the type of medium for printing
	// Defined Terms: PAPER, CLEAR FILM, BLUE FILM, etc.
	MediumType MediumType

	// PrintPriority specifies the priority of the print job
	// Enumerated values: HIGH, MED, LOW
	PrintPriority PrintPriority

	// NumberOfCopies is the number of copies to be printed for each film
	NumberOfCopies int

	// IsColor indicates whether this is a color or grayscale session
	IsColor bool

	// BasicFilmBoxes is the list of Film Boxes in this session
	BasicFilmBoxes []*FilmBox

	// PresentationLUTs is the list of Presentation LUTs for this session
	PresentationLUTs []*PresentationLUT
}

// NewFilmSession creates a new Film Session
func NewFilmSession(sopClassUID, sopInstanceUID string, isColor bool) *FilmSession {
	if sopClassUID == "" {
		sopClassUID = basicFilmSessionSOPClassUID
	}

	if sopInstanceUID == "" {
		sopInstanceUID = newSOPInstanceUID()
	}

	return &FilmSession{
		SOPClassUID:      sopClassUID,
		SOPInstanceUID:   sopInstanceUID,
		FilmDestination:  FilmDestinationProcessor,
		FilmSessionLabel: "",
		MemoryAllocation: 0,
		MediumType:       MediumTypeClearFilm,
		PrintPriority:    PrintPriorityMed,
		NumberOfCopies:   1,
		IsColor:          isColor,
		BasicFilmBoxes:   make([]*FilmBox, 0),
		PresentationLUTs: make([]*PresentationLUT, 0),
	}
}

// AddFilmBox adds a Film Box to the session
func (fs *FilmSession) AddFilmBox(filmBox *FilmBox) {
	if filmBox != nil {
		filmBox.filmSession = fs
		fs.BasicFilmBoxes = append(fs.BasicFilmBoxes, filmBox)
	}
}

// AddPresentationLUT adds a Presentation LUT to the session
func (fs *FilmSession) AddPresentationLUT(lut *PresentationLUT) {
	if lut != nil {
		fs.PresentationLUTs = append(fs.PresentationLUTs, lut)
	}
}

// CreateFilmBox creates a Film Box from Dataset values and attaches it to the session.
func (fs *FilmSession) CreateFilmBox(sopInstanceUID string, ds *dataset.Dataset) (*FilmBox, error) {
	if fs == nil {
		return nil, fmt.Errorf("printing: nil FilmSession")
	}
	if sopInstanceUID == "" && ds != nil {
		sopInstanceUID, _ = ds.GetString(tag.SOPInstanceUID)
	}
	if sopInstanceUID != "" && fs.FindFilmBox(sopInstanceUID) != nil {
		return nil, fmt.Errorf("printing: FilmBox SOP Instance UID %q already exists", sopInstanceUID)
	}
	filmBox, err := NewFilmBoxFromDataset(sopInstanceUID, ds)
	if err != nil {
		return nil, err
	}
	if fs.FindFilmBox(filmBox.SOPInstanceUID) != nil {
		return nil, fmt.Errorf("printing: FilmBox SOP Instance UID %q already exists", filmBox.SOPInstanceUID)
	}
	fs.AddFilmBox(filmBox)
	return filmBox, nil
}

// FindFilmBox returns the Film Box with the specified SOP Instance UID.
func (fs *FilmSession) FindFilmBox(sopInstanceUID string) *FilmBox {
	if fs == nil || sopInstanceUID == "" {
		return nil
	}
	for _, filmBox := range fs.BasicFilmBoxes {
		if filmBox != nil && filmBox.SOPInstanceUID == sopInstanceUID {
			return filmBox
		}
	}
	return nil
}

// DeleteFilmBox removes the Film Box with the specified SOP Instance UID.
func (fs *FilmSession) DeleteFilmBox(sopInstanceUID string) bool {
	if fs == nil || sopInstanceUID == "" {
		return false
	}
	for index, filmBox := range fs.BasicFilmBoxes {
		if filmBox != nil && filmBox.SOPInstanceUID == sopInstanceUID {
			filmBox.filmSession = nil
			fs.BasicFilmBoxes = append(fs.BasicFilmBoxes[:index], fs.BasicFilmBoxes[index+1:]...)
			return true
		}
	}
	return false
}

// FindImageBox searches every Film Box for an Image Box SOP Instance UID.
func (fs *FilmSession) FindImageBox(sopInstanceUID string) *ImageBox {
	if fs == nil || sopInstanceUID == "" {
		return nil
	}
	for _, filmBox := range fs.BasicFilmBoxes {
		if filmBox == nil {
			continue
		}
		for _, imageBox := range filmBox.BasicImageBoxes {
			if imageBox != nil && imageBox.SOPInstanceUID == sopInstanceUID {
				return imageBox
			}
		}
	}
	return nil
}

// CreatePresentationLUT creates a Presentation LUT from Dataset values and attaches it to the session.
func (fs *FilmSession) CreatePresentationLUT(sopInstanceUID string, ds *dataset.Dataset) (*PresentationLUT, error) {
	if fs == nil {
		return nil, fmt.Errorf("printing: nil FilmSession")
	}
	if sopInstanceUID == "" && ds != nil {
		sopInstanceUID, _ = ds.GetString(tag.SOPInstanceUID)
	}
	if sopInstanceUID != "" && fs.FindPresentationLUT(sopInstanceUID) != nil {
		return nil, fmt.Errorf("printing: PresentationLUT SOP Instance UID %q already exists", sopInstanceUID)
	}
	lut, err := NewPresentationLUTFromDataset(sopInstanceUID, ds)
	if err != nil {
		return nil, err
	}
	if fs.FindPresentationLUT(lut.SOPInstanceUID) != nil {
		return nil, fmt.Errorf("printing: PresentationLUT SOP Instance UID %q already exists", lut.SOPInstanceUID)
	}
	fs.AddPresentationLUT(lut)
	return lut, nil
}

// FindPresentationLUT returns the Presentation LUT with the specified SOP Instance UID.
func (fs *FilmSession) FindPresentationLUT(sopInstanceUID string) *PresentationLUT {
	if fs == nil || sopInstanceUID == "" {
		return nil
	}
	for _, lut := range fs.PresentationLUTs {
		if lut != nil && lut.SOPInstanceUID == sopInstanceUID {
			return lut
		}
	}
	return nil
}

// DeletePresentationLUT removes the Presentation LUT with the specified SOP Instance UID.
func (fs *FilmSession) DeletePresentationLUT(sopInstanceUID string) bool {
	if fs == nil || sopInstanceUID == "" {
		return false
	}
	for index, lut := range fs.PresentationLUTs {
		if lut != nil && lut.SOPInstanceUID == sopInstanceUID {
			fs.PresentationLUTs = append(fs.PresentationLUTs[:index], fs.PresentationLUTs[index+1:]...)
			return true
		}
	}
	return false
}

// Clone creates a recursively independent Film Session hierarchy.
func (fs *FilmSession) Clone() (*FilmSession, error) {
	if fs == nil {
		return nil, fmt.Errorf("printing: nil FilmSession")
	}
	sessionDataset, err := fs.ToDataset()
	if err != nil {
		return nil, err
	}
	clone, err := NewFilmSessionFromDataset("", "", sessionDataset, fs.IsColor)
	if err != nil {
		return nil, err
	}

	for index, sourceLUT := range fs.PresentationLUTs {
		if sourceLUT == nil {
			return nil, fmt.Errorf("printing: nil PresentationLUT at index %d", index)
		}
		lutDataset, err := sourceLUT.ToDataset()
		if err != nil {
			return nil, fmt.Errorf("clone PresentationLUT %q: %w", sourceLUT.SOPInstanceUID, err)
		}
		lut, err := NewPresentationLUTFromDataset("", lutDataset)
		if err != nil {
			return nil, fmt.Errorf("clone PresentationLUT %q: %w", sourceLUT.SOPInstanceUID, err)
		}
		clone.AddPresentationLUT(lut)
	}

	for boxIndex, sourceBox := range fs.BasicFilmBoxes {
		if sourceBox == nil {
			return nil, fmt.Errorf("printing: nil FilmBox at index %d", boxIndex)
		}
		boxDataset, err := sourceBox.ToDataset()
		if err != nil {
			return nil, fmt.Errorf("clone FilmBox %q: %w", sourceBox.SOPInstanceUID, err)
		}
		box, err := NewFilmBoxFromDataset("", boxDataset)
		if err != nil {
			return nil, fmt.Errorf("clone FilmBox %q: %w", sourceBox.SOPInstanceUID, err)
		}
		clone.AddFilmBox(box)
		for imageIndex, sourceImage := range sourceBox.BasicImageBoxes {
			if sourceImage == nil {
				return nil, fmt.Errorf("printing: nil ImageBox at FilmBox %q index %d", sourceBox.SOPInstanceUID, imageIndex)
			}
			imageDataset, err := sourceImage.ToDataset()
			if err != nil {
				return nil, fmt.Errorf("clone ImageBox %q: %w", sourceImage.SOPInstanceUID, err)
			}
			image, err := NewImageBoxFromDataset("", imageDataset, sourceImage.IsColor)
			if err != nil {
				return nil, fmt.Errorf("clone ImageBox %q: %w", sourceImage.SOPInstanceUID, err)
			}
			box.AddImageBox(image)
		}
	}
	return clone, nil
}

// GetFilmBox returns the Film Box at the specified index
func (fs *FilmSession) GetFilmBox(index int) *FilmBox {
	if index >= 0 && index < len(fs.BasicFilmBoxes) {
		return fs.BasicFilmBoxes[index]
	}
	return nil
}

// GetPresentationLUT returns the Presentation LUT at the specified index
func (fs *FilmSession) GetPresentationLUT(index int) *PresentationLUT {
	if index >= 0 && index < len(fs.PresentationLUTs) {
		return fs.PresentationLUTs[index]
	}
	return nil
}

// FilmBoxCount returns the number of Film Boxes in the session
func (fs *FilmSession) FilmBoxCount() int {
	return len(fs.BasicFilmBoxes)
}

// PresentationLUTCount returns the number of Presentation LUTs in the session
func (fs *FilmSession) PresentationLUTCount() int {
	return len(fs.PresentationLUTs)
}

// IsValid checks if the Film Session configuration is valid
func (fs *FilmSession) IsValid() bool {
	// Must have at least one Film Box
	if len(fs.BasicFilmBoxes) == 0 {
		return false
	}

	// Number of copies must be at least 1
	if fs.NumberOfCopies < 1 {
		return false
	}

	// Validate all Film Boxes
	for _, fb := range fs.BasicFilmBoxes {
		if !fb.IsValid() {
			return false
		}
	}

	return true
}
