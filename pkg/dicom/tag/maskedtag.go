// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package tag

import (
	"fmt"
	"strconv"
	"strings"
)

// MaskedTag represents a DICOM tag with a mask for wildcard matching.
//
// A masked tag allows matching tags with wildcards, e.g., "(0028,xx10)"
// matches any tag in group 0028 with element ending in 10.
// The mask determines which bits must match exactly.
type MaskedTag struct {
	tag  *Tag
	card uint32 // Tag value as uint32
	mask uint32 // Mask for matching
}

const (
	// FullMask indicates all bits must match (no wildcards)
	FullMask = 0xFFFFFFFF
)

// NewMaskedTag creates a new masked tag from a regular tag.
// The mask is set to FullMask (no wildcards).
func NewMaskedTag(tag *Tag) *MaskedTag {
	return &MaskedTag{
		tag:  tag,
		card: tag.ToUint32(),
		mask: FullMask,
	}
}

// NewMaskedTagWithMask creates a new masked tag with a specific mask.
func NewMaskedTagWithMask(tag *Tag, mask uint32) *MaskedTag {
	return &MaskedTag{
		tag:  tag,
		card: tag.ToUint32(),
		mask: mask,
	}
}

// Tag returns the underlying DICOM tag.
func (mt *MaskedTag) Tag() *Tag {
	return mt.tag
}

// Group returns the tag's group number.
func (mt *MaskedTag) Group() uint16 {
	return mt.tag.Group()
}

// Element returns the tag's element number.
func (mt *MaskedTag) Element() uint16 {
	return mt.tag.Element()
}

// Card returns the tag as a 32-bit unsigned integer.
func (mt *MaskedTag) Card() uint32 {
	return mt.card
}

// Mask returns the matching mask.
func (mt *MaskedTag) Mask() uint32 {
	return mt.mask
}

// String returns the default string representation of the masked tag.
func (mt *MaskedTag) String() string {
	return mt.Format("G")
}

// Format returns a formatted string representation of the masked tag.
//
// Supported formats:
//   - "G" (default): (0028,xx10) - shows 'x' for wildcard nibbles
//   - "g": group only with wildcards (e.g., "0028" or "xx28")
//   - "e": element only with wildcards (e.g., "0010" or "xx10")
func (mt *MaskedTag) Format(format string) string {
	switch format {
	case "g":
		return mt.formatGroup()
	case "e":
		return mt.formatElement()
	case "G":
		fallthrough
	default:
		return fmt.Sprintf("(%s,%s)", mt.formatGroup(), mt.formatElement())
	}
}

// formatGroup formats the group with wildcards.
func (mt *MaskedTag) formatGroup() string {
	s := fmt.Sprintf("%04x", mt.tag.Group())
	x := ""
	if (mt.mask & 0xF0000000) != 0 {
		x += string(s[0])
	} else {
		x += "x"
	}
	if (mt.mask & 0x0F000000) != 0 {
		x += string(s[1])
	} else {
		x += "x"
	}
	if (mt.mask & 0x00F00000) != 0 {
		x += string(s[2])
	} else {
		x += "x"
	}
	if (mt.mask & 0x000F0000) != 0 {
		x += string(s[3])
	} else {
		x += "x"
	}
	return x
}

// formatElement formats the element with wildcards.
func (mt *MaskedTag) formatElement() string {
	s := fmt.Sprintf("%04x", mt.tag.Element())
	x := ""
	if (mt.mask & 0x0000F000) != 0 {
		x += string(s[0])
	} else {
		x += "x"
	}
	if (mt.mask & 0x00000F00) != 0 {
		x += string(s[1])
	} else {
		x += "x"
	}
	if (mt.mask & 0x000000F0) != 0 {
		x += string(s[2])
	} else {
		x += "x"
	}
	if (mt.mask & 0x0000000F) != 0 {
		x += string(s[3])
	} else {
		x += "x"
	}
	return x
}

// IsMatch checks if the given tag matches this masked tag.
func (mt *MaskedTag) IsMatch(tag *Tag) bool {
	return mt.card == (tag.ToUint32() & mt.mask)
}

// ParseMaskedTag parses a masked tag from a string.
//
// Supported formats:
//   - "(0028,xx10)" - with parentheses and wildcards
//   - "0028,xx10" - without parentheses
//   - "0028xx10" - compact format
//
// Wildcards are represented by 'x' in the hex digits.
func ParseMaskedTag(s string) (*MaskedTag, error) {
	if len(s) < 8 {
		return nil, fmt.Errorf("masked tag string too short: '%s' (expected at least 8 characters)", s)
	}

	// Remove spaces
	s = strings.TrimSpace(s)

	pos := 0

	// Skip opening parenthesis if present
	if s[pos] == '(' {
		pos++
	}

	// Find comma or assume 4 characters for group
	idx := strings.IndexByte(s[pos:], ',')
	if idx == -1 {
		idx = 4
	} else {
		idx = pos + idx
	}

	group := s[pos:idx]
	pos = idx

	// Skip comma if present
	if pos < len(s) && s[pos] == ',' {
		pos++
	}

	// Extract element
	element := s[pos:]
	if strings.HasSuffix(element, ")") {
		element = element[:len(element)-1]
	}

	return parseMaskedTagParts(group, element)
}

// parseMaskedTagParts parses a masked tag from group and element strings.
func parseMaskedTagParts(group, element string) (*MaskedTag, error) {
	// Parse group and element, replacing 'x' with '0'
	groupClean := strings.ToLower(strings.ReplaceAll(group, "x", "0"))
	elementClean := strings.ToLower(strings.ReplaceAll(element, "x", "0"))

	g, err := strconv.ParseUint(groupClean, 16, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid group number '%s': %w", group, err)
	}

	e, err := strconv.ParseUint(elementClean, 16, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid element number '%s': %w", element, err)
	}

	// Build mask: 'x' becomes 0 in mask, hex digit becomes f
	maskStr := strings.ToLower(group + element)
	maskStr = strings.Map(func(r rune) rune {
		if r == 'x' {
			return '0'
		}
		// Any hex digit (0-9, a-f) becomes 'f'
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') {
			return 'f'
		}
		return r
	}, maskStr)

	mask, err := strconv.ParseUint(maskStr, 16, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mask: %w", err)
	}

	tag := New(uint16(g), uint16(e))
	return &MaskedTag{
		tag:  tag,
		card: tag.ToUint32(),
		mask: uint32(mask),
	}, nil
}

// MustParseMaskedTag parses a masked tag from a string and panics if there is an error.
func MustParseMaskedTag(s string) *MaskedTag {
	mt, err := ParseMaskedTag(s)
	if err != nil {
		panic(fmt.Sprintf("invalid masked tag string '%s': %v", s, err))
	}
	return mt
}
