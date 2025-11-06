// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package tag_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestNewMaskedTag(t *testing.T) {
	testTag := tag.New(0x0028, 0x0010)
	mt := tag.NewMaskedTag(testTag)

	if mt.Group() != 0x0028 {
		t.Errorf("Group() = 0x%04x, want 0x0028", mt.Group())
	}
	if mt.Element() != 0x0010 {
		t.Errorf("Element() = 0x%04x, want 0x0010", mt.Element())
	}
	if mt.Mask() != tag.FullMask {
		t.Errorf("Mask() = 0x%08x, want 0x%08x", mt.Mask(), tag.FullMask)
	}
}

func TestMaskedTagFormat(t *testing.T) {
	tests := []struct {
		name   string
		mt     *tag.MaskedTag
		format string
		want   string
	}{
		{
			"no wildcards G",
			tag.NewMaskedTag(tag.New(0x0028, 0x0010)),
			"G",
			"(0028,0010)",
		},
		{
			"wildcards in element",
			tag.MustParseMaskedTag("(0028,xx10)"),
			"G",
			"(0028,xx10)",
		},
		{
			"wildcards in group",
			tag.MustParseMaskedTag("(xx28,0010)"),
			"G",
			"(xx28,0010)",
		},
		{
			"wildcards in both",
			tag.MustParseMaskedTag("(xx28,xx10)"),
			"G",
			"(xx28,xx10)",
		},
		{
			"group format",
			tag.MustParseMaskedTag("(xx28,0010)"),
			"g",
			"xx28",
		},
		{
			"element format",
			tag.MustParseMaskedTag("(0028,xx10)"),
			"e",
			"xx10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mt.Format(tt.format); got != tt.want {
				t.Errorf("Format(%q) = %v, want %v", tt.format, got, tt.want)
			}
		})
	}
}

func TestMaskedTagString(t *testing.T) {
	mt := tag.MustParseMaskedTag("(0028,xx10)")
	want := "(0028,xx10)"
	if got := mt.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestMaskedTagIsMatch(t *testing.T) {
	tests := []struct {
		name       string
		maskedTag  string
		testTag    *tag.Tag
		wantMatch  bool
	}{
		{
			"exact match no wildcards",
			"(0028,0010)",
			tag.New(0x0028, 0x0010),
			true,
		},
		{
			"no match no wildcards",
			"(0028,0010)",
			tag.New(0x0028, 0x0011),
			false,
		},
		{
			"match with element wildcard",
			"(0028,xx10)",
			tag.New(0x0028, 0x0010),
			true,
		},
		{
			"match with element wildcard 2",
			"(0028,xx10)",
			tag.New(0x0028, 0x1010),
			true,
		},
		{
			"match with element wildcard 3",
			"(0028,xx10)",
			tag.New(0x0028, 0xFF10),
			true,
		},
		{
			"no match with element wildcard",
			"(0028,xx10)",
			tag.New(0x0028, 0x0011),
			false,
		},
		{
			"match with group wildcard",
			"(xx28,0010)",
			tag.New(0x0028, 0x0010),
			true,
		},
		{
			"match with group wildcard 2",
			"(xx28,0010)",
			tag.New(0x1028, 0x0010),
			true,
		},
		{
			"no match with group wildcard",
			"(xx28,0010)",
			tag.New(0x0029, 0x0010),
			false,
		},
		{
			"match with both wildcards",
			"(xx28,xx10)",
			tag.New(0x0028, 0x0010),
			true,
		},
		{
			"match with both wildcards 2",
			"(xx28,xx10)",
			tag.New(0x1028, 0xFF10),
			true,
		},
		{
			"no match with both wildcards",
			"(xx28,xx10)",
			tag.New(0x0029, 0x0010),
			false,
		},
		{
			"multiple wildcard nibbles",
			"(00xx,00xx)",
			tag.New(0x0028, 0x0010),
			true,
		},
		{
			"multiple wildcard nibbles 2",
			"(00xx,00xx)",
			tag.New(0x00FF, 0x00AA),
			true,
		},
		{
			"no match multiple wildcards",
			"(00xx,00xx)",
			tag.New(0x0128, 0x0010),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := tag.MustParseMaskedTag(tt.maskedTag)
			if got := mt.IsMatch(tt.testTag); got != tt.wantMatch {
				t.Errorf("IsMatch() = %v, want %v (mask=0x%08x, card=0x%08x, testCard=0x%08x)",
					got, tt.wantMatch, mt.Mask(), mt.Card(), tt.testTag.ToUint32())
			}
		})
	}
}

func TestParseMaskedTag(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantG   uint16
		wantE   uint16
		wantMask uint32
		wantErr bool
	}{
		{
			"no wildcards",
			"(0028,0010)",
			0x0028,
			0x0010,
			0xFFFFFFFF,
			false,
		},
		{
			"element wildcard upper byte",
			"(0028,xx10)",
			0x0028,
			0x0010,
			0xFFFF00FF,
			false,
		},
		{
			"element wildcard lower byte",
			"(0028,00xx)",
			0x0028,
			0x0000,
			0xFFFFFF00,
			false,
		},
		{
			"group wildcard",
			"(xx28,0010)",
			0x0028,
			0x0010,
			0x00FFFFFF,
			false,
		},
		{
			"multiple wildcards",
			"(xx28,xx10)",
			0x0028,
			0x0010,
			0x00FF00FF,
			false,
		},
		{
			"all wildcards",
			"(xxxx,xxxx)",
			0x0000,
			0x0000,
			0x00000000,
			false,
		},
		{
			"single nibble wildcard",
			"(0x28,0010)",
			0x0028,
			0x0010,
			0xF0FFFFFF,
			false,
		},
		{
			"without parentheses",
			"0028,xx10",
			0x0028,
			0x0010,
			0xFFFF00FF,
			false,
		},
		{
			"compact format",
			"0028xx10",
			0x0028,
			0x0010,
			0xFFFF00FF,
			false,
		},
		{
			"too short",
			"0028",
			0,
			0,
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tag.ParseMaskedTag(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMaskedTag(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Group() != tt.wantG {
				t.Errorf("Group() = 0x%04x, want 0x%04x", got.Group(), tt.wantG)
			}
			if got.Element() != tt.wantE {
				t.Errorf("Element() = 0x%04x, want 0x%04x", got.Element(), tt.wantE)
			}
			if got.Mask() != tt.wantMask {
				t.Errorf("Mask() = 0x%08x, want 0x%08x", got.Mask(), tt.wantMask)
			}
		})
	}
}

func TestMustParseMaskedTag(t *testing.T) {
	// Should not panic for valid input
	result := tag.MustParseMaskedTag("(0028,xx10)")
	if result.Group() != 0x0028 {
		t.Errorf("Group() = 0x%04x, want 0x0028", result.Group())
	}

	// Should panic for invalid input
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParseMaskedTag with invalid input should panic")
		}
	}()
	tag.MustParseMaskedTag("invalid")
}
