// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

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
		sopClassUID = "1.2.840.10008.5.1.1.1" // Basic Film Session SOP Class
	}

	if sopInstanceUID == "" {
		// Generate a UID if not provided
		sopInstanceUID = "1.2.840.10008.5.1.1.1.1" // Placeholder
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
