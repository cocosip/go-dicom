// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestFormatValueSummarizesElementValues(t *testing.T) {
	fragment := element.NewOtherByteFragment(tag.PixelData)
	fragment.AddFragment(buffer.NewMemory([]byte{1, 2, 3}))

	tests := []struct {
		name string
		elem element.Element
		want string
	}{
		{"text", element.NewString(tag.SpecificCharacterSet, vr.CS, []string{"ISO_IR 100"}), "[ISO_IR 100]"},
		{"multi value", element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY"}), `[ORIGINAL\PRIMARY]`},
		{"numeric", element.NewUnsignedShort(tag.Rows, []uint16{512}), "512"},
		{"numeric multi", element.NewUnsignedShort(tag.Rows, []uint16{512, 256}), `512\256`},
		{"date", element.NewDate(tag.ContentDate, []string{"20260101"}), "[20260101]"},
		{"person name", element.NewPersonName(tag.PatientName, []string{"Doe^Jane"}), "[Doe^Jane]"},
		{"attribute tag", element.NewAttributeTag(tag.New(0x0072, 0x0050), []*tag.Tag{tag.PatientName}), "(0010,0010)"},
		{"empty", element.NewOtherByte(tag.PixelData, nil), "(empty)"},
		{"binary", element.NewOtherByte(tag.PixelData, []byte{1, 2, 3, 4}), "(4 bytes)"},
		{"fragment", fragment, "(1 fragment, 0 offsets)"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := element.FormatValue(test.elem, 64); got != test.want {
				t.Errorf("FormatValue() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestFormatValueTruncatesLongText(t *testing.T) {
	long := element.NewString(tag.PatientName, vr.PN, []string{strings.Repeat("A", 80)})
	if got := element.FormatValue(long, 10); got != "[AAAAAAAAAA...]" {
		t.Errorf("FormatValue() = %q, want truncated value", got)
	}
}

func TestSummaryReturnsTagVRAndValue(t *testing.T) {
	text := element.NewString(tag.SpecificCharacterSet, vr.CS, []string{"ISO_IR 100"})
	if got := element.Summary(text); got != "(0008,0005) CS [ISO_IR 100]" {
		t.Errorf("Summary() = %q", got)
	}

	numeric := element.NewUnsignedShort(tag.Rows, []uint16{512})
	if got := element.Summary(numeric); got != "(0028,0010) US 512" {
		t.Errorf("Summary() = %q, want numeric value", got)
	}
}
