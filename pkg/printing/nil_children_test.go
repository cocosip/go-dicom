// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

import (
	"strings"
	"testing"
)

func TestPrintValidationRejectsNilHierarchyChildren(t *testing.T) {
	t.Run("nil Film Box", func(t *testing.T) {
		session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.500", false)
		session.BasicFilmBoxes = append(session.BasicFilmBoxes, nil)

		err := validatePrintSession(session)
		if err == nil || !strings.Contains(err.Error(), "FilmBox at index 0 is nil") {
			t.Fatalf("validatePrintSession() error = %v, want nil FilmBox error", err)
		}
	})

	t.Run("nil Image Box", func(t *testing.T) {
		session := NewFilmSession(basicFilmSessionSOPClassUID, "2.25.501", false)
		filmBox := NewFilmBox("2.25.502", "STANDARD\\1,1")
		filmBox.BasicImageBoxes = append(filmBox.BasicImageBoxes, nil)
		session.AddFilmBox(filmBox)

		err := validatePrintSession(session)
		if err == nil || !strings.Contains(err.Error(), "ImageBox at FilmBox") {
			t.Fatalf("validatePrintSession() error = %v, want nil ImageBox error", err)
		}
	})
}

func TestNilHierarchyChildIsInvalidNotPanic(t *testing.T) {
	var filmBox *FilmBox
	if filmBox.IsValid() {
		t.Error("(*FilmBox)(nil).IsValid() = true, want false")
	}
	var imageBox *ImageBox
	if imageBox.IsValid() {
		t.Error("(*ImageBox)(nil).IsValid() = true, want false")
	}
}
