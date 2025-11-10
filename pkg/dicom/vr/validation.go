// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	// Pre-compiled regular expressions for performance
	datePattern = regexp.MustCompile(`^\d{8}$`)
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

// containsControlChars checks if string contains control characters.
func containsControlChars(s string) bool {
	for _, r := range s {
		if unicode.IsControl(r) {
			return true
		}
	}
	return false
}

// ValidateAE validates Application Entity (AE) values.
func ValidateAE(content string) error {
	// may not be longer than 16 characters
	if len(content) > 16 {
		return newValidationError("AE", content, "value exceeds maximum length of 16 characters")
	}

	// may not consist only of spaces
	if strings.TrimSpace(content) == "" && content != "" {
		return newValidationError("AE", content, "value may not consist only of spaces")
	}

	// Default Character Repertoire excluding backslash and control characters
	if strings.Contains(content, "\\") || containsControlChars(content) {
		return newValidationError("AE", content, "value contains invalid control character or backslash")
	}

	return nil
}

// ValidateAS validates Age String (AS) values.
func ValidateAS(content string) error {
	if content == "" {
		return nil
	}

	// Format: nnnD, nnnW, nnnM, nnnY (e.g., 025Y for 25 years)
	matched, _ := regexp.MatchString(`^\d{3}[DWMY]$`, content)
	if !matched {
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
	matched, _ := regexp.MatchString(`^[A-Z0-9_ ]*$`, content)
	if !matched {
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

		if month > 12 {
			return newValidationError("DA", content, "month component exceeds 12")
		}
		if day > 31 {
			return newValidationError("DA", content, "day component exceeds 31")
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
	matched, _ := regexp.MatchString(`^[+-]?((\d+(\.\d*)?)|(\.\d+))([eE][-+]?\d+)?$`, trimmed)
	if !matched {
		return newValidationError("DS", content, "value is not a valid decimal string")
	}

	return nil
}

// ValidateDT validates Date Time (DT) values.
// Format: YYYYMMDDHHMMSS.FFFFFF&ZZXX
func ValidateDT(content string) error {
	// This is a simplified validation. The full validation in fo-dicom is very complex.
	// For now, we'll implement basic checks.

	if content == "" {
		return nil
	}

	// Remove trailing spaces
	trimmed := strings.TrimRight(content, " ")

	// Check if it contains invalid negative UTC
	if strings.Contains(content, "-0000") {
		return newValidationError("DT", content, "negative UTC hours component with value -0000 is not allowed")
	}

	// Basic format check: at minimum YYYY (4 digits)
	if len(trimmed) < 4 {
		return newValidationError("DT", content, "value is too short (minimum YYYY required)")
	}

	// More detailed validation would go here following fo-dicom's complex logic
	// For now, this provides basic validation
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
	matched, _ := regexp.MatchString(`^[+-]?\d+$`, trimmed)
	if !matched {
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
	if len(content) > 64 {
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
	if len(content) > 10240 {
		return newValidationError("LT", content, "value exceeds maximum length of 10240 characters")
	}

	return nil
}

// ValidatePN validates Person Name (PN) values.
func ValidatePN(content string) error {
	// 64 characters maximum per component group
	if len(content) > 64 {
		return newValidationError("PN", content, "value exceeds maximum length of 64 characters")
	}

	// Person names can have up to 5 components separated by ^
	// and up to 3 component groups separated by =
	// For now, basic validation
	return nil
}

// ValidateSH validates Short String (SH) values.
func ValidateSH(content string) error {
	// 16 characters maximum
	if len(content) > 16 {
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
	// 1024 characters maximum
	if len(content) > 1024 {
		return newValidationError("ST", content, "value exceeds maximum length of 1024 characters")
	}

	return nil
}

// ValidateTM validates Time (TM) values.
// Format: HHMMSS.FFFFFF
func ValidateTM(content string) error {
	if content == "" {
		return nil
	}

	trimmed := strings.TrimRight(content, " ")

	// Basic format check
	matched, _ := regexp.MatchString(`^\d{2}(\d{2}(\d{2}(\.\d{1,6})?)?)?$`, trimmed)
	if !matched {
		return newValidationError("TM", content, "value does not match pattern HH[MM[SS[.F{1-6}]]]")
	}

	// Validate hour, minute, second components
	if len(trimmed) >= 2 {
		hour, _ := strconv.Atoi(trimmed[0:2])
		if hour > 23 {
			return newValidationError("TM", content, "hour component exceeds 23")
		}
	}

	if len(trimmed) >= 4 {
		minute, _ := strconv.Atoi(trimmed[2:4])
		if minute > 59 {
			return newValidationError("TM", content, "minute component exceeds 59")
		}
	}

	if len(trimmed) >= 6 {
		second, _ := strconv.Atoi(trimmed[4:6])
		if second > 60 { // 60 allowed for leap second
			return newValidationError("TM", content, "second component exceeds 60")
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
	matched, _ := regexp.MatchString(`^\d+(\.\d+)*$`, trimmed)
	if !matched {
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
func ValidateUR(content string) error {
	// Unlimited length, trailing spaces allowed
	// This is a URL/URI, basic validation
	return nil
}

// ValidateUT validates Unlimited Text (UT) values.
func ValidateUT(content string) error {
	// Unlimited length
	return nil
}
