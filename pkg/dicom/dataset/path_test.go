// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestFormatPathIncludesItemsAndFragments(t *testing.T) {
	itemIndex := 2
	fragmentIndex := 3
	path := Path{
		{Tag: tag.ReferencedImageSequence, ItemIndex: &itemIndex},
		{Tag: tag.PixelData, FragmentIndex: &fragmentIndex},
	}

	if got, want := FormatPath(path), "(0008,1140)[2]/(7fe0,0010)#3"; got != want {
		t.Fatalf("FormatPath() = %q, want %q", got, want)
	}
	if got := FormatPath(nil); got != "<dataset>" {
		t.Fatalf("FormatPath(nil) = %q, want %q", got, "<dataset>")
	}
}

func TestClonePathDoesNotAliasIndexes(t *testing.T) {
	itemIndex := 1
	fragmentIndex := 4
	original := Path{
		{Tag: tag.ReferencedImageSequence, ItemIndex: &itemIndex},
		{Tag: tag.PixelData, FragmentIndex: &fragmentIndex},
	}

	clone := ClonePath(original)
	*clone[0].ItemIndex = 9
	*clone[1].FragmentIndex = 8
	clone[0].Tag = tag.PatientName

	if *original[0].ItemIndex != 1 || *original[1].FragmentIndex != 4 {
		t.Fatalf("ClonePath shared index storage with original: %#v", original)
	}
	if original[0].Tag != tag.ReferencedImageSequence {
		t.Fatal("ClonePath shared mutable segment storage with original")
	}
}

func TestValidationPathAliasRemainsAssignable(t *testing.T) {
	var path []ValidationPathSegment = Path{{Tag: tag.PatientName}}
	err := &ValidationError{Kind: ValidationValue, Path: path}

	if got := err.Error(); !strings.Contains(got, "(0010,0010)") {
		t.Fatalf("ValidationError.Error() = %q, want Patient Name path", got)
	}
}
