package tag

import (
	"errors"
	"fmt"
	"go-dicom/pkg/util"
	"strconv"
)

//Dicom tag
type DicomTag struct {
	Group   uint16
	Element uint16
}

//IsPrivate indicates if the input group is part of a private tag.
func (tag DicomTag) IsPrivate () bool {
	return util.IsUInt16Odd(tag.Group)
}

// Compare returns -1, 0, or 1 if t<other, t==other, t>other, respectively.
// Tags are ordered first by Group, then by Element.
func (tag DicomTag) Compare (other DicomTag) int {
	if tag.Group < other.Group {
		return -1
	}
	if tag.Group > other.Group {
		return 1
	}
	if tag.Element < other.Element {
		return -1
	}
	if tag.Element > other.Element {
		return 1
	}
	return 0
}

// Equals returns true if this tag equals the provided tag.
func (tag DicomTag) Equals(other DicomTag) bool {
	return tag.Compare(other) == 0
}

// String returns a string of form "(0008,1234)", where 0x0008 is t.Group,
// 0x1234 is t.Element.
func (tag DicomTag) String() string {
	return fmt.Sprintf("(%04x,%04x)", tag.Group, tag.Element)
}

// IsPrivate indicates if the input group is part of a private tag.
func IsPrivate(tag DicomTag) bool {
	return tag.Group%2 == 1
}

// DebugString returns a human-readable diagnostic string for the tag, in format
// "(group, elem)[name]".
func DebugString(tag DicomTag) string {
	if tag.IsPrivate() {
		return fmt.Sprintf("(%04x,%04x)[private]", tag.Group, tag.Element)
	}
	return fmt.Sprintf("(%04x,%04x)", tag.Group, tag.Element)
}

func Parse(s string) (DicomTag,error) {
	if len(s) < 8 {
		return DicomTag{}, errors.New(fmt.Sprintf("%s Expected a string of 8 or more characters", s))
	}
	pos := 0
	if s[pos] == '(' {
		pos++
	}
	group, err := strconv.ParseUint(s[pos:(pos+4)], 10, 16)
	if err != nil {
		return DicomTag{}, err
	}
	pos += 4

	if s[pos] == ',' {
		pos++
	}
	element, err := strconv.ParseUint(s[pos:(pos+4)], 10, 16)
	pos += 4
	//DicomPrivateCreator creator = null;
	if len(s) > pos && s[pos] == ':' {
		pos++
		var c string
		if s[len(s)-1] == ')' {
			c = s[pos:(len(s) - 1)]
		} else {
			c = s[pos:len(s)]
		}
		fmt.Sprintf("c:%s\n", c)
		//creator = DicomDictionary.Default.GetPrivateCreator(c);
	}

	//TODO: get value from related DicomDictionaryEntry
	return DicomTag{Group: uint16(group), Element: uint16(element)}, nil
}


