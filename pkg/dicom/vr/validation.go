// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package vr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	ageStringPattern     = regexp.MustCompile(`^\d{3}[DWMY]$`)
	codeStringPattern    = regexp.MustCompile(`^[A-Z0-9_ ]*$`)
	datePattern          = regexp.MustCompile(`^\d{8}$`)
	decimalStringPattern = regexp.MustCompile(`^[+-]?((\d+(\.\d*)?)|(\.\d+))([eE][-+]?\d+)?$`)
	integerStringPattern = regexp.MustCompile(`^[+-]?\d+$`)
	timePattern          = regexp.MustCompile(`^\d{2}(\d{2}(\d{2}(\.\d{1,6})?)?)?$`)
	uidPattern           = regexp.MustCompile(`^\d+(\.\d+)*$`)
)

// ValidationError represents a VR validation error.
type ValidationError struct {
	Content string
	VRCode  string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("VR %s validation error: %s (content: '%s')", e.VRCode, e.Message, e.Content)
}

// newValidationError creates a new validation error.
func newValidationError(vrCode, content, message string) error {
	return &ValidationError{
		Content: content,
		VRCode:  vrCode,
		Message: message,
	}
}

// ValidateAE validates Application Entity (AE) values.
func ValidateAE(content string) error {
	// AE values are limited to 16 bytes and always use the Default Character Repertoire.
	if len(content) > 16 {
		return newValidationError("AE", content, "value exceeds maximum length of 16 bytes")
	}

	// may not consist only of spaces
	if strings.TrimSpace(content) == "" && content != "" {
		return newValidationError("AE", content, "value may not consist only of spaces")
	}

	for i := 0; i < len(content); i++ {
		if content[i] < 0x20 || content[i] > 0x7E || content[i] == '\\' {
			return newValidationError("AE", content, "value contains a character outside the permitted Default Character Repertoire")
		}
	}

	return nil
}

// ValidateAS validates Age String (AS) values.
func ValidateAS(content string) error {
	if content == "" {
		return nil
	}

	// Format: nnnD, nnnW, nnnM, nnnY (e.g., 025Y for 25 years)
	if !ageStringPattern.MatchString(content) {
		return newValidationError("AS", content, "value does not match pattern nnnD|nnnW|nnnM|nnnY")
	}

	return nil
}

// ValidateCS validates Code String (CS) values.
func ValidateCS(content string) error {
	// 16 bytes maximum
	if len(content) > 16 {
		return newValidationError("CS", content, "value exceeds maximum length of 16 characters")
	}

	// Uppercase characters, digits, space, and underscore only
	if !codeStringPattern.MatchString(content) {
		return newValidationError("CS", content, "value contains invalid character (only A-Z, 0-9, space, underscore allowed)")
	}

	return nil
}

// ValidateDA validates Date (DA) values.
func ValidateDA(content string) error {
	// Split by range separator if present
	dateComponents := strings.Split(content, "-")

	if len(dateComponents) > 2 {
		return newValidationError("DA", content, "value contains too many range separators '-'")
	}

	for _, component := range dateComponents {
		trimmed := strings.TrimRight(component, " ")
		if trimmed == "" {
			continue
		}

		// Check format: YYYYMMDD
		if !datePattern.MatchString(trimmed) {
			return newValidationError("DA", content, "date does not match pattern YYYYMMDD")
		}

		// Validate month and day components
		month, _ := strconv.Atoi(trimmed[4:6])
		day, _ := strconv.Atoi(trimmed[6:8])

		if month < 1 || month > 12 {
			return newValidationError("DA", content, "month component out of range")
		}
		if day < 1 || day > 31 {
			return newValidationError("DA", content, "day component out of range")
		}
		if _, err := time.Parse("20060102", trimmed); err != nil {
			return newValidationError("DA", content, "date is not a valid calendar date")
		}
	}

	return nil
}

// ValidateDS validates Decimal String (DS) values.
func ValidateDS(content string) error {
	// 16 bytes maximum
	if len(content) > 16 {
		return newValidationError("DS", content, "value exceeds maximum length of 16 characters")
	}

	trimmed := strings.TrimSpace(content)
	if trimmed == "" {
		return nil
	}

	// Match decimal number with optional exponent
	if !decimalStringPattern.MatchString(trimmed) {
		return newValidationError("DS", content, "value is not a valid decimal string")
	}

	return nil
}

// ValidateDT validates Date Time (DT) values.
// Format (PS3.5 Table 6.2-1):
//
//	YYYY[MM[DD[HH[MM[SS[.FFFFFF]]]]]][&ZZXX]
//
// The date/time part must use an even number of digits, fractional seconds are
// only allowed after full seconds, and the optional UTC offset must be prefixed
// with '&' (e.g. "&+0800"). The offset "-0000" is not permitted.
//
// Range matching values (PS3.4 C.2.2.2.5.3) of the form "<datetime1>-<datetime2>",
// "-<datetime1>" and "<datetime1>-" are also accepted.
func ValidateDT(content string) error {
	if content == "" {
		return nil
	}

	trimmed := strings.TrimRight(content, " ")
	singleValueErr := validateDTValue(trimmed)
	if singleValueErr == nil {
		return nil
	}
	if validateDTRange(trimmed) {
		return nil
	}
	return singleValueErr
}

// validateDTRange reports whether value is a PS3.4 C.2.2.2.5.3 DT range. A
// hyphen can also introduce a negative UTC offset, so every hyphen is tried as
// the range separator and the value is accepted if both sides are valid.
func validateDTRange(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] != '-' {
			continue
		}
		// A hyphen followed by exactly four trailing digits after a full date is
		// a UTC offset (e.g. "...-0500"), not a range separator; the single-value
		// parse above decides whether that offset is valid.
		// ponytail: this also rejects the rare 4-digit-year end range such as
		// "20240101-2024"; expand only if such query keys need support.
		if i >= 8 && i+5 == len(value) && allDigits(value[i+1:]) {
			continue
		}
		left := strings.TrimRight(value[:i], " ")
		right := strings.TrimRight(value[i+1:], " ")
		if left != "" && validateDTValue(left) != nil {
			continue
		}
		if right != "" && validateDTValue(right) != nil {
			continue
		}
		return true
	}
	return false
}

func validateDTValue(trimmed string) error {
	main := trimmed
	offset := ""
	hasOffset := false
	// The optional suffix is a signed UTC offset: '+' or '-' followed by 4
	// digits (PS3.5 Table 6.2-1, where the '&' in "&ZZXX" denotes the sign).
	if i := strings.IndexAny(trimmed, "+-"); i >= 0 {
		main = trimmed[:i]
		offset = trimmed[i:]
		hasOffset = true
	}

	// Date/time magnitude: at least YYYY, at most YYYYMMDDHHMMSS.
	dateTime := main
	fraction := ""
	hasFraction := false
	if dot := strings.IndexByte(main, '.'); dot >= 0 {
		dateTime = main[:dot]
		fraction = main[dot+1:]
		hasFraction = true
	}

	if len(dateTime) < 4 || len(dateTime) > 14 || len(dateTime)%2 != 0 {
		return newValidationError("DT", trimmed, "date/time component must be an even number of digits between 4 and 14")
	}
	if !allDigits(dateTime) {
		return newValidationError("DT", trimmed, "date/time component must contain only digits")
	}
	if hasFraction {
		if len(dateTime) != 14 {
			return newValidationError("DT", trimmed, "fractional seconds require a full YYYYMMDDHHMMSS date/time")
		}
		if len(fraction) < 1 || len(fraction) > 6 || !allDigits(fraction) {
			return newValidationError("DT", trimmed, "fractional seconds must be 1 to 6 digits")
		}
	}

	if err := validateDateTimeFields(dateTime); err != nil {
		return newValidationError("DT", trimmed, err.Error())
	}

	if hasOffset {
		if len(offset) != 5 || (offset[0] != '+' && offset[0] != '-') {
			return newValidationError("DT", trimmed, "UTC offset must be a sign followed by 4 digits (ZZXX)")
		}
		if !allDigits(offset[1:]) {
			return newValidationError("DT", trimmed, "UTC offset digits must be numeric")
		}
		if offset == "-0000" {
			return newValidationError("DT", trimmed, "negative UTC hours component with value -0000 is not allowed")
		}
		hours := int(offset[1]-'0')*10 + int(offset[2]-'0')
		minutes := int(offset[3]-'0')*10 + int(offset[4]-'0')
		// PS3.5 Table 6.2-1: UTC offset range is -1200 to +1400.
		if hours > 14 || minutes > 59 || (hours == 14 && minutes != 0) || (offset[0] == '-' && hours > 12) {
			return newValidationError("DT", trimmed, "UTC offset out of range (-1200 to +1400)")
		}
	}

	return nil
}

func allDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func validateDateTimeFields(dateTime string) error {
	atoi2 := func(s string) int { return int(s[0]-'0')*10 + int(s[1]-'0') }
	if len(dateTime) >= 6 {
		if month := atoi2(dateTime[4:6]); month < 1 || month > 12 {
			return fmt.Errorf("month out of range")
		}
	}
	if len(dateTime) >= 8 {
		if day := atoi2(dateTime[6:8]); day < 1 || day > 31 {
			return fmt.Errorf("day out of range")
		}
		if _, err := time.Parse("20060102", dateTime[:8]); err != nil {
			return fmt.Errorf("date is not a valid calendar date")
		}
	}
	if len(dateTime) >= 10 {
		if hour := atoi2(dateTime[8:10]); hour > 23 {
			return fmt.Errorf("hour out of range")
		}
	}
	if len(dateTime) >= 12 {
		if minute := atoi2(dateTime[10:12]); minute > 59 {
			return fmt.Errorf("minute out of range")
		}
	}
	if len(dateTime) >= 14 {
		if second := atoi2(dateTime[12:14]); second > 60 {
			return fmt.Errorf("second out of range")
		}
	}
	return nil
}

// ValidateIS validates Integer String (IS) values.
func ValidateIS(content string) error {
	if content == "" {
		return nil
	}

	// 12 bytes maximum
	if len(content) > 12 {
		return newValidationError("IS", content, "value exceeds maximum length of 12 characters")
	}

	trimmed := strings.TrimSpace(content)
	if trimmed == "" {
		return nil
	}

	// Check format: optional +/- followed by digits
	if !integerStringPattern.MatchString(trimmed) {
		return newValidationError("IS", content, "value is not an integer string")
	}

	// Check if it fits in int32
	_, err := strconv.ParseInt(trimmed, 10, 32)
	if err != nil {
		return newValidationError("IS", content, "value too large to fit in 32-bit integer")
	}

	return nil
}

// ValidateLO validates Long String (LO) values.
func ValidateLO(content string) error {
	// 64 characters maximum per value
	if utf8.RuneCountInString(content) > 64 {
		return newValidationError("LO", content, "value exceeds maximum length of 64 characters")
	}

	// LO uses specific character set, control characters not allowed except ESC
	for _, r := range content {
		if unicode.IsControl(r) && r != '\x1B' { // ESC is allowed
			return newValidationError("LO", content, "value contains invalid control character")
		}
	}

	return nil
}

// ValidateLT validates Long Text (LT) values.
func ValidateLT(content string) error {
	// 10240 characters maximum
	if utf8.RuneCountInString(content) > 10240 {
		return newValidationError("LT", content, "value exceeds maximum length of 10240 characters")
	}

	return nil
}

// ValidatePN validates Person Name (PN) values.
func ValidatePN(content string) error {
	// PS3.5 limits each component group to 64 characters. Component groups are
	// separated by '=': alphabetic=ideographic=phonetic.
	for _, group := range strings.Split(content, "=") {
		if utf8.RuneCountInString(group) > 64 {
			return newValidationError("PN", content, "component group exceeds maximum length of 64 characters")
		}
	}

	// Person names can have up to 5 components separated by ^
	// and up to 3 component groups separated by =
	return nil
}

// ValidateSH validates Short String (SH) values.
func ValidateSH(content string) error {
	// 16 characters maximum
	if utf8.RuneCountInString(content) > 16 {
		return newValidationError("SH", content, "value exceeds maximum length of 16 characters")
	}

	// SH uses specific character set, control characters not allowed except ESC
	for _, r := range content {
		if unicode.IsControl(r) && r != '\x1B' { // ESC is allowed
			return newValidationError("SH", content, "value contains invalid control character")
		}
	}

	return nil
}

// ValidateST validates Short Text (ST) values.
func ValidateST(content string) error {
	// 1024 characters maximum (count runes, not UTF-8 bytes)
	if utf8.RuneCountInString(content) > 1024 {
		return newValidationError("ST", content, "value exceeds maximum length of 1024 characters")
	}

	return nil
}

// ValidateTM validates Time (TM) values.
// Format: HHMMSS.FFFFFF
//
// Range matching values (PS3.4 C.2.2.2.5.2) of the form "<time1>-<time2>",
// "-<time1>" and "<time1>-" are also accepted.
func ValidateTM(content string) error {
	if content == "" {
		return nil
	}

	trimmed := strings.TrimRight(content, " ")
	if err := validateTMValue(trimmed); err == nil {
		return nil
	}

	parts := strings.Split(trimmed, "-")
	if len(parts) > 2 {
		return newValidationError("TM", content, "value contains too many range separators '-'")
	}
	for _, part := range parts {
		part = strings.TrimRight(part, " ")
		if part == "" {
			continue
		}
		if err := validateTMValue(part); err != nil {
			return err
		}
	}
	return nil
}

func validateTMValue(trimmed string) error {
	// Basic format check
	if !timePattern.MatchString(trimmed) {
		return newValidationError("TM", trimmed, "value does not match pattern HH[MM[SS[.F{1-6}]]]")
	}

	// Validate hour, minute, second components
	if len(trimmed) >= 2 {
		hour, _ := strconv.Atoi(trimmed[0:2])
		if hour > 23 {
			return newValidationError("TM", trimmed, "hour component exceeds 23")
		}
	}

	if len(trimmed) >= 4 {
		minute, _ := strconv.Atoi(trimmed[2:4])
		if minute > 59 {
			return newValidationError("TM", trimmed, "minute component exceeds 59")
		}
	}

	if len(trimmed) >= 6 {
		second, _ := strconv.Atoi(trimmed[4:6])
		if second > 60 { // 60 allowed for leap second
			return newValidationError("TM", trimmed, "second component exceeds 60")
		}
	}

	return nil
}

// ValidateUC validates Unlimited Characters (UC) values.
func ValidateUC(content string) error {
	// Unlimited length, uses specific character set
	// Control characters not allowed except ESC
	for _, r := range content {
		if unicode.IsControl(r) && r != '\x1B' { // ESC is allowed
			return newValidationError("UC", content, "value contains invalid control character")
		}
	}

	return nil
}

// ValidateUI validates Unique Identifier (UI) values.
func ValidateUI(content string) error {
	// 64 characters maximum
	if len(content) > 64 {
		return newValidationError("UI", content, "value exceeds maximum length of 64 characters")
	}

	trimmed := strings.TrimRight(content, " ")
	if trimmed == "" {
		return nil
	}

	// UID format: digits and dots only, no leading/trailing dots
	if !uidPattern.MatchString(trimmed) {
		return newValidationError("UI", content, "value contains invalid characters (only digits and dots allowed)")
	}

	// Components should not start with 0 (except "0" itself)
	components := strings.Split(trimmed, ".")
	for _, comp := range components {
		if len(comp) > 1 && comp[0] == '0' {
			return newValidationError("UI", content, "UID component has leading zero")
		}
	}

	return nil
}

// ValidateUR validates Universal Resource (UR) values.
func ValidateUR(_ string) error {
	// Unlimited length, trailing spaces allowed
	// This is a URL/URI, basic validation
	return nil
}

// ValidateUT validates Unlimited Text (UT) values.
func ValidateUT(_ string) error {
	// Unlimited length
	return nil
}
