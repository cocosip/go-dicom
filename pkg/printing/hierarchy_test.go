// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestFilmSessionHierarchyFindAndDelete(t *testing.T) {
	session := NewFilmSession("", "2.25.200", false)
	boxDataset, err := NewFilmBox("2.25.201", testStandardOneByOne).ToDataset()
	if err != nil {
		t.Fatalf("FilmBox.ToDataset() error = %v", err)
	}
	box, err := session.CreateFilmBox("", boxDataset)
	if err != nil {
		t.Fatalf("CreateFilmBox() error = %v", err)
	}
	image := NewImageBox("2.25.202", false)
	image.ImageBoxPosition = 1
	image.SetImageData([]byte{1, 2})
	box.AddImageBox(image)

	lutDataset, err := NewPresentationLUT("2.25.203").ToDataset()
	if err != nil {
		t.Fatalf("PresentationLUT.ToDataset() error = %v", err)
	}
	lut, err := session.CreatePresentationLUT("", lutDataset)
	if err != nil {
		t.Fatalf("CreatePresentationLUT() error = %v", err)
	}

	if got := session.FindFilmBox("2.25.201"); got != box {
		t.Fatalf("FindFilmBox() = %p, want %p", got, box)
	}
	if got := session.FindImageBox("2.25.202"); got != image {
		t.Fatalf("FindImageBox() = %p, want %p", got, image)
	}
	if got := session.FindPresentationLUT("2.25.203"); got != lut {
		t.Fatalf("FindPresentationLUT() = %p, want %p", got, lut)
	}
	if image.FilmBox() != box || box.FilmSession() != session {
		t.Fatal("hierarchy parent references were not established")
	}

	if !session.DeletePresentationLUT("2.25.203") || session.FindPresentationLUT("2.25.203") != nil {
		t.Fatal("DeletePresentationLUT() did not remove the matching LUT")
	}
	if !session.DeleteFilmBox("2.25.201") || session.FindFilmBox("2.25.201") != nil {
		t.Fatal("DeleteFilmBox() did not remove the matching FilmBox")
	}
	if box.FilmSession() != nil {
		t.Fatal("DeleteFilmBox() retained the removed parent reference")
	}
}

func TestFilmSessionHierarchyRejectsDuplicateUIDs(t *testing.T) {
	session := NewFilmSession("", "2.25.210", false)
	boxDataset, err := NewFilmBox("2.25.211", testStandardOneByOne).ToDataset()
	if err != nil {
		t.Fatalf("FilmBox.ToDataset() error = %v", err)
	}
	if _, err := session.CreateFilmBox("", boxDataset); err != nil {
		t.Fatalf("first CreateFilmBox() error = %v", err)
	}
	if _, err := session.CreateFilmBox("", boxDataset); err == nil {
		t.Fatal("duplicate FilmBox UID was accepted")
	}

	lutDataset, err := NewPresentationLUT("2.25.212").ToDataset()
	if err != nil {
		t.Fatalf("PresentationLUT.ToDataset() error = %v", err)
	}
	if _, err := session.CreatePresentationLUT("", lutDataset); err != nil {
		t.Fatalf("first CreatePresentationLUT() error = %v", err)
	}
	if _, err := session.CreatePresentationLUT("", lutDataset); err == nil {
		t.Fatal("duplicate PresentationLUT UID was accepted")
	}
}

func TestFilmSessionCloneIsIndependentAndReparentsChildren(t *testing.T) {
	source := NewFilmSession("", "2.25.220", true)
	source.FilmSessionLabel = "source"
	box := NewFilmBox("2.25.221", testStandardOneByOne)
	box.ConfigurationInformation = "CS321"
	box.AnnotationDisplayFormatID = "ANNOTATION_2"
	box.ReferencedPresentationLUT = SOPReference{SOPClassUID: presentationLUTSOPClassUID, SOPInstanceUID: "2.25.223"}
	image := NewImageBox("2.25.222", true)
	image.ImageBoxPosition = 1
	maxDensity := uint16(0)
	image.MaxDensity = &maxDensity
	imageItem := dataset.New()
	if err := imageItem.Add(element.NewUnsignedShort(tag.Rows, []uint16{1})); err != nil {
		t.Fatalf("add image Rows: %v", err)
	}
	if err := imageItem.Add(element.NewOtherByte(tag.PixelData, []byte{5, 6, 7})); err != nil {
		t.Fatalf("add image Pixel Data: %v", err)
	}
	if err := image.SetImageSequence(imageItem); err != nil {
		t.Fatalf("SetImageSequence() error = %v", err)
	}
	box.AddImageBox(image)
	source.AddFilmBox(box)
	lut := NewPresentationLUT("2.25.223")
	if err := lut.SetLUT(2, 0, 12, []uint16{11, 22}); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}
	source.AddPresentationLUT(lut)

	clone, err := source.Clone()
	if err != nil {
		t.Fatalf("Clone() error = %v", err)
	}
	cloneBox := clone.FindFilmBox("2.25.221")
	cloneImage := clone.FindImageBox("2.25.222")
	cloneLUT := clone.FindPresentationLUT("2.25.223")
	if clone == source || cloneBox == box || cloneImage == image || cloneLUT == lut {
		t.Fatal("Clone() reused source object pointers")
	}
	if cloneBox == nil || cloneImage == nil || cloneLUT == nil {
		t.Fatal("Clone() omitted hierarchy members")
	}
	if cloneBox.FilmSession() != clone || cloneImage.FilmBox() != cloneBox {
		t.Fatal("Clone() retained source parent references")
	}

	image.PreformattedColorImageSequence[0] = 99
	lut.LUTData[0] = 99
	box.ConfigurationInformation = "changed"
	cloneImageSequence, err := cloneImage.ImageSequence()
	if err != nil {
		t.Fatalf("clone ImageSequence() error = %v", err)
	}
	rows, rowsErr := cloneImageSequence.GetUInt16(tag.Rows, 0)
	if cloneImage.GetImageData()[0] != 5 || cloneImage.MaxDensity == nil || *cloneImage.MaxDensity != 0 ||
		rowsErr != nil || rows != 1 || cloneLUT.LUTData[0] != 11 ||
		cloneBox.ConfigurationInformation != "CS321" || cloneBox.AnnotationDisplayFormatID != "ANNOTATION_2" ||
		cloneBox.ReferencedPresentationLUT.SOPInstanceUID != "2.25.223" {
		t.Fatal("Clone() aliases source state")
	}
}

func TestParseImageDisplayFormatRejectsNonPositiveCounts(t *testing.T) {
	for _, format := range []string{`STANDARD\0,2`, `STANDARD\-1,2`, `ROW\1,0`, `COL\-1,2`} {
		t.Run(format, func(t *testing.T) {
			if _, err := ParseImageDisplayFormat(format); err == nil {
				t.Fatalf("ParseImageDisplayFormat(%q) accepted a non-positive count", format)
			}
		})
	}
}
