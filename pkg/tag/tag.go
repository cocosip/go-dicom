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

func (tag DicomTag) IsPrivate () bool {
	return util.IsUInt16Odd(tag.Group)
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


