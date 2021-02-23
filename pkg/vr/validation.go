package vr

import (
	"errors"
	"fmt"
	"go-dicom/pkg/util"
	"regexp"
	"strconv"
	"strings"
)

var PerformValidation = true

func ValidateAE(content string) error {
	if len(content) > 16 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 16 characters", content))
	}
	m, err := regexp.MatchString("\\s*$", content)
	if err != nil {
		return err
	}
	if m {
		return errors.New(fmt.Sprintf("%s value may not consist only of spaces", content))
	}
	//Default Character Repertoire excluding character code 5CH (the BACKSLASH "\" in ISO-IR 6), and control characters LF, FF, CR and ESC.
	//content.ToCharArray().Any(char.IsControl)
	//TODO
	if strings.Contains(content, "\\") {
		return errors.New(fmt.Sprintf("%s value contains invalid control character", content))
	}
	return nil
}

func ValidateAS(content string) error {
	if len(content) == 0 {
		return nil
	}
	m, err := regexp.MatchString("^\\d\\d\\d[DWMY]$", content)
	if err != nil {
		return err
	}
	if !m {
		return errors.New(fmt.Sprintf("%s value does not have pattern 000[DWMY]", content))
	}
	return nil
}

func ValidateCS(content string) error {
	if len(content) > 16 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 16 characters", content))
	}
	m, err := regexp.MatchString("^[A-Z0-9_ ]*$", content)
	if err != nil {
		return err
	}
	if !m {
		return errors.New(fmt.Sprintf("%s value contains invalid character. Only uppercase character, digits, space and underscore alre allowed", content))
	}
	return nil
}

func ValidateDA(content string) error {
	/*
	   A string of characters of the format YYYYMMDD; where YYYY shall contain year, MM shall contain the
	   month, and DD shall contain the day, interpreted as a date of the Gregorian calendar system.
	*/
	dateComponents := strings.Split(content, "-")
	if len(dateComponents) > 2 {
		return errors.New(fmt.Sprintf("%s value contains too many range separators '-'", content))
	}
	for _, component := range dateComponents {
		// Trailling spaces are allowed
		trimmedComponent := strings.TrimRight(component, " ")
		if len(trimmedComponent) == 0 {
			continue
		}
		m, err := regexp.MatchString("^\\d{8}$", trimmedComponent);
		if err != nil {
			return err
		}
		if !m {
			return errors.New(fmt.Sprintf("%s one of the date values does not match the pattern yyyyMMdd.", content))
		}
		// The date is in the numeric format, validate the month and day components
		//var month = trimmedComponent.Substring(4, 2);
		//var day = trimmedComponent.Substring(6, 2);

		month, err := strconv.Atoi(trimmedComponent[4:6])
		if err != nil {
			return err
		}
		if month > 12 {
			return errors.New(fmt.Sprintf("%s month component exceeds the value 12", content))
		}

		day, err := strconv.Atoi(trimmedComponent[6:8])
		if err != nil {
			return err
		}
		if day > 31 {
			return errors.New(fmt.Sprintf("%s day component exceeds the value 31", content))
		}
	}
	return nil
}

func ValidateDS(content string) error {
	content = strings.Trim(content, " ")
	m, err := regexp.MatchString("^[+-]?((0|[1-9][0-9]*)([.][0-9]*)?|[.][0-9]+)([eE][-+]?[0-9]+)?$", content)
	if err != nil {
		return err
	}
	if !m {
		return errors.New(fmt.Sprintf("%s value is no decimal string", content))
	}
	return nil
}

func ValidateDT(content string) error {
	/*
	   "0"-"9", "+", "-", "." and the SPACE character of Default Character Repertoire

	    26 bytes maximum. In the context of a Query with range matching (see PS3.4), the length is 54 bytes maximum.

	   A concatenated date-time character string in the format:
	   YYYYMMDDHHMMSS.FFFFFF&ZZXX

	   The components of this string, from left to right, are YYYY = Year, MM = Month, DD = Day,
	   HH = Hour (range "00" - "23"), MM = Minute (range "00" - "59"), SS = Second (range "00" - "60").

	   FFFFFF = Fractional Second contains a fractional part of a second as small as 1 millionth of
	   a second (range "000000" - "999999").

	   &ZZXX is an optional suffix for offset from Coordinated Universal Time (UTC), where & = "+"
	   or "-", and ZZ = Hours and XX = Minutes of offset.

	   The year, month, and day shall be interpreted as a date of the Gregorian calendar system.

	   A 24-hour clock is used. Midnight shall be represented by only "0000" since "2400" would
	   violate the hour range.

	   The Fractional Second component, if present, shall contain 1 to 6 digits. If Fractional
	   Second is unspecified the preceding "." shall not be included. The offset suffix, if present,
	   shall contain 4 digits. The string may be padded with trailing SPACE characters. Leading
	   and embedded spaces are not allowed.

	   A component that is omitted from the string is termed a null component. Trailing null
	   components of Date Time indicate that the value is not precise to the precision of those
	   components. The YYYY component shall not be null. Non-trailing null components are prohibited.
	   The optional suffix is not considered as a component.

	   A Date Time value without the optional suffix is interpreted to be in the local time zone
	   of the application creating the Data Element, unless explicitly specified by the Timezone
	   Offset From UTC (0008,0201).

	   UTC offsets are calculated as "local time minus UTC". The offset for a Date Time value in
	   UTC shall be +0000.

	   Used regex checking string: "YYYY[MM[DD[HH[MM[SS[.FFFFFF]]]]]][&ZZXX]"
	   If date is not empty, YYYY should not be null.
	*/

	if strings.Contains(content, "-0000") {
		return errors.New(fmt.Sprintf("%s negative UTC hours component with value -0000 is not allowed", content))
	}

	if strings.Trim(content, " ") == "-" {
		return errors.New(fmt.Sprintf("%s both dateTime components in range cannot be empty", content))
	}

	dateTimeComponents := strings.Split(content, "-")
	// DateTime may contain more than two '-' characters because of the negative UTC suffixes
	if len(dateTimeComponents) > 4 {
		return errors.New(fmt.Sprintf("%s value contains too many range separators '-'", content))
	}

	if len(dateTimeComponents) == 4 {
		// Join 4 range separated components (X,Y,X,Y) into 2 range components with negative UTC (X-Y,X-Y)
		firstComponent := dateTimeComponents[0] + "-" + dateTimeComponents[1]
		secondComponent := dateTimeComponents[2] + "-" + dateTimeComponents[3]
		dateTimeComponents = []string{firstComponent, secondComponent}
	} else if len(dateTimeComponents) == 3 {
		// Join 3 range separated components (X, Y, Z) into 2 range components with negative UTC (X-Y,Z) or (X,Y-Z)
		var firstComponent string
		var secondComponent string

		m1, err := regexp.MatchString("^\\d{4}$", dateTimeComponents[1])
		if err != nil {
			return err
		}

		v1, err := strconv.Atoi(dateTimeComponents[1])
		if err == nil {
			if m1 && v1 <= 1200 {
				// Second component is UTC -> (X-Y,Z)
				firstComponent = dateTimeComponents[0] + "-" + dateTimeComponents[1]
				secondComponent = dateTimeComponents[2]
			}
		}

		m2, err := regexp.MatchString("^\\d{4}$", dateTimeComponents[2])
		if err == nil {
			v2, err := strconv.Atoi(dateTimeComponents[2])
			if err == nil {
				if m2 && v2 <= 1200 {
					// Third component is UTC -> (X,Y-Z)
					firstComponent = dateTimeComponents[0];
					secondComponent = dateTimeComponents[1] + "-" + dateTimeComponents[2];
				}
			}
		}

		if err == nil {
			if len(firstComponent) > 0 || len(secondComponent) > 0 {
				dateTimeComponents = []string{firstComponent, secondComponent}
			} else {
				return errors.New(fmt.Sprintf("%s value is in invalid range format", content))
			}
		} else {
			return err
		}
	} else if len(dateTimeComponents) == 2 {
		m, err := regexp.MatchString("^\\d{4}$", dateTimeComponents[1])
		if err == nil && m {
			newComponent := dateTimeComponents[0] + "-" + dateTimeComponents[1]
			dateTimeComponents = []string{newComponent}
		}
		return err
	}

	for _, component := range dateTimeComponents {
		// Trailling spaces are allowed
		trimmedComponent := strings.TrimRight(component, " ")
		if len(trimmedComponent) == 0 {
			continue
		}
		// Split by optional suffix for UTC +/-ZZXX
		splittedDateTime := util.SplitRunes(trimmedComponent, []rune("+-"))
		if len(splittedDateTime) > 2 {
			return errors.New(fmt.Sprintf("%s value contains too many UTC separators '+' or '-'", content))
		} else if len(splittedDateTime) == 2 {
			utcSuffixString := splittedDateTime[1]
			m, err := regexp.MatchString("^\\d{4}$", utcSuffixString)
			if err != nil {
				return err
			}
			if !m {
				return errors.New(fmt.Sprintf("%s value does not match the UTC pattern &ZZXX", content))
			}
			isPositiveOffset := strings.Contains(trimmedComponent, "+")
			hours := utcSuffixString[0:2]
			minutes := utcSuffixString[2:4]
			hoursValue, err := strconv.Atoi(hours)
			if err != nil {
				return err
			}

			if isPositiveOffset && hoursValue > 14 {
				return errors.New(fmt.Sprintf("%s positive UTC hours component exceeds 14 (allowed range is -1200 to +1400)", content))
			} else if !isPositiveOffset && hoursValue > 12 {
				return errors.New(fmt.Sprintf("%s negative UTC hours component exceeds 12 (allowed range is -1200 to +1400)", content))
			}

			minutesValue, err := strconv.Atoi(minutes)
			if err != nil {
				return err
			}

			if minutesValue > 59 {
				return errors.New(fmt.Sprintf("%s UTC minutes component exceeds 59", content))
			}
		}

		dateTimeString := splittedDateTime[0]
		m, err := regexp.MatchString("^\\d{4}$|^\\d{6}$|^\\d{8}$|^\\d{10}$|^\\d{12}$|^\\d{14}$|^\\d{14}\\.\\d{1,6}$", dateTimeString)
		if err != nil {
			return err
		}
		if !m {
			return errors.New(fmt.Sprintf("%s value does not mach pattern YYYY[MM[DD[HH[MM[SS[.F{1-6}]]]]]]", content))
		}

		if len(dateTimeString) >= 14 {
			seconds := dateTimeString[12:15]
			secondsValue, err := strconv.Atoi(seconds)
			if err != nil {
				return err
			}
			if secondsValue > 59 {
				return errors.New(fmt.Sprintf("%s minutes component exceeds 59", content))
			}
		}
		if len(dateTimeString) >= 12 {
			minutes := dateTimeString[10:12]
			minutesValue, err := strconv.Atoi(minutes)
			if err != nil {
				return err
			}
			if minutesValue > 59 {
				return errors.New(fmt.Sprintf("%s minutes component exceeds 59", content))
			}
		}

		if len(dateTimeString) >= 10 {
			hours := dateTimeString[8:10]
			hoursValue, err := strconv.Atoi(hours)
			if err != nil {
				return err
			}
			if hoursValue > 23 {
				return errors.New(fmt.Sprintf("%s hours component exceeds 23", content))
			}
		}

		if len(dateTimeString) >= 8 {
			day := dateTimeString[6:8]
			dayValue, err := strconv.Atoi(day)
			if err != nil {
				return err
			}
			if dayValue > 31 {
				return errors.New(fmt.Sprintf("%s day component exceeds 31", content))
			} else if dayValue == 0 {
				return errors.New(fmt.Sprintf("%s day component cannot be 0"))
			}
		}

		if len(dateTimeString) >= 6 {
			month := dateTimeString[4:6]
			monthValue, err := strconv.Atoi(month)
			if err != nil {
				return err
			}
			if monthValue > 12 {
				return errors.New(fmt.Sprintf("%s month component exceeds 12", content))
			} else if monthValue == 0 {
				return errors.New(fmt.Sprintf("%s month component cannot be 0", content))
			}
		}

		if len(dateTimeString) > 0 && len(dateTimeString) < 4 {
			return errors.New(fmt.Sprintf("%s year component is too short and not in the correct YYYY format", content))
		}
	}

	return nil
}

func ValidateIS(content string) error {
	/* "0" - "9", "+", "-" of Default Character Repertoire

	   12 bytes maximum

	   A string of characters representing an Integer in base-10 (decimal), shall contain only the
	   characters 0 - 9, with an optional leading "+" or "-". It may be padded with leading and/or
	   trailing spaces. Embedded spaces are not allowed.

	   The integer, n, represented shall be in the range:

	   -2^31 <= n <= (2^31-1).
	*/
	if len(content) == 0 {
		return nil
	}
	content = strings.Trim(content, " ")
	m, err := regexp.MatchString("^[+-]?\\d+$", content)
	if err != nil {
		return err
	}
	if !m {
		return errors.New(fmt.Sprintf("%s value is not an integer string", content))
	}
	_, err = strconv.Atoi(content)
	if err != nil {
		return errors.New(fmt.Sprintf("%s value too large to fit 32 bit integer", content))
	}
	return nil
}

func ValidateLO(content string) error {
	/*
	 * Default Character Repertoire and/or as defined by (0008,0005).
	 *
	 * 64 chars maximum (see Note in Section 6.2)
	 *
	 * A character string that may be padded with leading and/or trailing spaces. The character
	 * code 5CH (the BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the
	 * delimiter between values in multiple valued data elements. The string shall not have
	 * Control Characters except for ESC.
	 */
	if len(content) == 0 {
		return nil
	}
	if len(content) > 64 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 64 characters", content))
	}
	//if (content.Contains("\\") || content.ToCharArray().Any(IsControlExceptESC))
	//private static bool IsControlExceptESC(char c) => char.IsControl(c) && (c != '\u001b');
	//TODO
	if strings.Contains(content, "\\") {
		return errors.New(fmt.Sprintf("%s value contains invalid character", content))
	}
	return nil
}

func ValidateLT(content string) error {
	if len(content) == 0 {
		return nil
	}
	/*
	   A character string that may contain one or more paragraphs. It may contain the Graphic
	   Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with
	   trailing spaces, which may be ignored, but leading spaces are considered to be significant.
	   Data Elements with this VR shall not be multi - valued and therefore character code
	   5CH(the BACKSLASH "\" in ISO-IR 6) may be used.

	   Default Character Repertoire and / or as defined by(0008, 0005).

	   10240 chars maximum(see Note in Section 6.2) */
	if len(content) > 10240 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 10240 characters", content))
	}
	return nil
}

func ValidatePN(content string) error {
	/*
	   A character string encoded using a 5 component convention. The character code 5CH (the
	   BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the delimiter between
	   values in multiple valued data elements. The string may be padded with trailing spaces. For
	   human use, the five components in their order of occurrence are: family name complex, given
	   name complex, middle name, name prefix, name suffix.

	   Note
	   HL7 prohibits leading spaces within a component; DICOM allows leading and trailing spaces
	   and considers them insignificant.

	    Any of the five components may be an empty string. The component delimiter shall be the
	    caret "^" character (5EH). Delimiters are required for interior null components. Trailing
	    null components and their delimiters may be omitted. Multiple entries are permitted in
	    each component and are encoded as natural text strings, in the format preferred by the
	    named person.

	    For veterinary use, the first two of the five components in their order of occurrence
	    are: responsible party family name or responsible organization name, patient name. The
	    remaining components are not used and shall not be present.

	    This group of five components is referred to as a Person Name component group.

	    For the purpose of writing names in ideographic characters and in phonetic characters,
	    up to 3 groups of components (see Annexes H, I and J) may be used. The delimiter for
	    component groups shall be the equals character "=" (3DH). The three component groups
	    of components in their order of occurrence are: an alphabetic representation, an
	    ideographic representation, and a phonetic representation.

	    Any component group may be absent, including the first component group. In this case,
	    the person name may start with one or more "=" delimiters. Delimiters are required for
	    interior null component groups. Trailing null component groups and their delimiters may
	    be omitted.

	    Precise semantics are defined for each component group. See Section 6.2.1.2.

	    For examples and notes, see Section 6.2.1.1.

	    Default Character Repertoire and/or as defined by (0008,0005) excluding Control Characters
	    LF, FF, and CR but allowing Control Character ESC.

	    64 chars maximum per component group */
	if len(content) == 0 {
		return nil
	}

	groups := strings.Split(content, "=")
	if len(groups) > 3 {
		return errors.New(fmt.Sprintf("%s value contains too many groups", content))
	}
	for _, group := range groups {
		if len(group) > 64 {
			return errors.New(fmt.Sprintf("%s value exceeds maximum length of 64 characters", content))
		}
		// if (group.ToCharArray().Any(IsControlExceptESC))
		//TODO
	}

	var groupcomponents = make([]int, len(groups))
	for i, group := range groups {
		groupcomponents[i] = len(strings.Split(group, "^"))
	}

	for _, l := range groupcomponents {
		if l > 5 {
			return errors.New(fmt.Sprintf("%s value contains too many components", content))
		}
	}
	return nil
}

func ValidateSH(content string) error {
	/*
	   A character string that may be padded with leading and/or trailing spaces. The character
	   code 05CH (the BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the
	   delimiter between values for multiple data elements. The string shall not have Control
	   Characters except ESC.

	   Default Character Repertoire and/or as defined by (0008,0005).

	   16 chars maximum (see Note in Section 6.2)*/

	//if (content.Contains("\\") || content.ToCharArray().Any(IsControlExceptESC))
	//TODO
	if strings.Contains(content, "\\") {
		return errors.New(fmt.Sprintf("%s value contains invalid character", content))
	}
	if len(content) > 10 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 16 characters", content))
	}
	return nil
}


func ValidateST(content string) error {
	/*
	   A character string that may contain one or more paragraphs. It may contain the Graphic
	   Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with
	   trailing spaces, which may be ignored, but leading spaces are considered to be significant.
	   Data Elements with this VR shall not be multi-valued and therefore character code
	   5CH (the BACKSLASH "\" in ISO-IR 6) may be used.

	   Default Character Repertoire and/or as defined by (0008,0005).

	   1024 chars maximum (see Note in Section 6.2) */
	if len(content) > 1024 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 1024 characters", content))
	}
	return nil
}

func ValidateTM(content string) error {
	/*
	   A string of characters of the format HHMMSS.FFFFFF; where HH contains hours (range "00" - "23"),
	   MM contains minutes (range "00" - "59"), SS contains seconds (range "00" - "60"), and FFFFFF
	   contains a fractional part of a second as small as 1 millionth of a second (range "000000" -
	   "999999"). A 24-hour clock is used. Midnight shall be represented by only "0000" since "2400"
	   would violate the hour range. The string may be padded with trailing spaces. Leading and
	   embedded spaces are not allowed.

	   One or more of the components MM, SS, or FFFFFF may be unspecified as long as every component
	   to the right of an unspecified component is also unspecified, which indicates that the value
	   is not precise to the precision of those unspecified components.

	   The FFFFFF component, if present, shall contain 1 to 6 digits. If FFFFFF is unspecified the
	   preceding "." shall not be included.

	   Examples:

	   "070907.0705 " represents a time of 7 hours, 9 minutes and 7.0705 seconds.

	   "1010" represents a time of 10 hours, and 10 minutes.

	   "021 " is an invalid value.

	   The SS component may have a value of 60 only for a leap second.

	   "0"-"9", "." and the SPACE character of Default Character Repertoire

	   In the context of a Query with range matching (see PS3.4), the character "-" is allowed.

	   16 bytes maximum
	   In the context of a Query with range matching (see PS3.4), the length is 28 bytes maximum.
	*/
	queryComponents := strings.Split(content, "-")
	if len(queryComponents) > 2 {
		return errors.New(fmt.Sprintf("%s value contains too many range separators '-'", content))
	}

	for _, component := range queryComponents {
		cont := strings.TrimRight(component, " ")
		if len(cont) == 0 {
			continue
		}
		m, err := regexp.MatchString("^\\d{2}$|^\\d{4}$|^\\d{6}$|^\\d{6}\\.\\d{1,6}$", cont)
		if err != nil {
			return err
		}
		if !m {
			return errors.New(fmt.Sprintf("%s value does not mach pattern HH or HHMM or HHMMSS or HHMMSS.F{1-6}", content))
		}
		if len(cont) >= 2 {
			hh := cont[0:2]
			hhValue, err := strconv.Atoi(hh)
			if err != nil {
				return err
			}
			if hhValue > 23 {
				return errors.New(fmt.Sprintf("%s hour component exceeds 23", content))
			}
		}
		if len(cont) >= 4 {
			mm := cont[2:4]
			mmValue, err := strconv.Atoi(mm)
			if err != nil {
				return err
			}
			if mmValue > 59 {
				return errors.New(fmt.Sprintf("%s minutes component exceeds 59", content))
			}
		}
		if len(cont) >= 6 {
			ss := cont[4:6]
			ssValue, err := strconv.Atoi(ss)
			if err != nil {
				return err
			}
			if ssValue > 60 {
				return errors.New(fmt.Sprintf("%s seconds component exceeds 60", content))
			}
		}
	}
	return nil
}

func ValidateUI(content string) error {
	/*
	   The UID is a series of numeric components separated by the period "." character.
	   If a Value Field containing one or more UIDs is an odd number of bytes in length,
	   the Value Field shall be padded with a single trailing NULL (00H)
	   character to ensure that the Value Field is an even number of bytes in length.
	   See Section 9 and Annex B for a complete specification and examples.

	   "0"-"9", "." of Default Character Repertoire

	   64 bytes maximum */
	// trailling spaces are allowed
	content = strings.TrimRight(content, " ")
	if len(content) == 0 {
		return nil
	}
	if len(content) > 64 {
		return errors.New(fmt.Sprintf("%s value exceeds maximum length of 64 characters", content))
	}
	m, err := regexp.MatchString("^[0-9\\\\.]*$", content)
	if err != nil {
		return err
	}
	if !m {
		return errors.New(fmt.Sprintf("%s value contains invalid characters other than '0'-'9' and '.'", content))
	}
	if strings.HasPrefix(content, "0") {
		return errors.New(fmt.Sprintf("%s components must not have leading zeros", content))
	}
	m2, err := regexp.MatchString("[.]0\\d", content)
	if err != nil {
		return err
	}
	if m2 {
		return errors.New(fmt.Sprintf("%s components must not have leading zeros", content))
	}
	return nil
}








