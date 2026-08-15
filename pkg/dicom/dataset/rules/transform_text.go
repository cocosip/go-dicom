// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type editValuesRule struct {
	tag  *tag.Tag
	edit func(string) string
}

// ReplaceRegex applies one RE2 replacement independently to each value.
func ReplaceRegex(t *tag.Tag, pattern, replacement string) (TransformRule, error) {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("compile replacement regular expression %q: %w", pattern, err)
	}
	return newEditValuesRule(t, func(value string) string {
		return compiled.ReplaceAllString(value, replacement)
	})
}

// ToUpper converts each value to Unicode upper case.
func ToUpper(t *tag.Tag) (TransformRule, error) {
	return newEditValuesRule(t, strings.ToUpper)
}

// ToLower converts each value to Unicode lower case.
func ToLower(t *tag.Tag) (TransformRule, error) {
	return newEditValuesRule(t, strings.ToLower)
}

// PrefixValue prefixes every value.
func PrefixValue(t *tag.Tag, prefix string) (TransformRule, error) {
	return newEditValuesRule(t, func(value string) string { return prefix + value })
}

// AppendValue appends to every value.
func AppendValue(t *tag.Tag, suffix string) (TransformRule, error) {
	return newEditValuesRule(t, func(value string) string { return value + suffix })
}

// TrimPosition selects which ends a trim rule edits.
type TrimPosition uint8

const (
	// TrimStart edits only the start of each value.
	TrimStart TrimPosition = iota
	// TrimEnd edits only the end of each value.
	TrimEnd
	// TrimBoth edits both ends of each value.
	TrimBoth
)

// TrimSpace removes Unicode whitespace from the selected ends.
func TrimSpace(t *tag.Tag, position TrimPosition) (TransformRule, error) {
	if err := validateTrimPosition(position); err != nil {
		return nil, err
	}
	return newEditValuesRule(t, func(value string) string {
		switch position {
		case TrimStart:
			return strings.TrimLeftFunc(value, unicode.IsSpace)
		case TrimEnd:
			return strings.TrimRightFunc(value, unicode.IsSpace)
		default:
			return strings.TrimFunc(value, unicode.IsSpace)
		}
	})
}

// TrimCharacters removes characters in cutset from the selected ends.
func TrimCharacters(t *tag.Tag, position TrimPosition, cutset string) (TransformRule, error) {
	if err := validateTrimPosition(position); err != nil {
		return nil, err
	}
	if cutset == "" {
		return nil, fmt.Errorf("trim character set is empty")
	}
	return newEditValuesRule(t, func(value string) string {
		switch position {
		case TrimStart:
			return strings.TrimLeft(value, cutset)
		case TrimEnd:
			return strings.TrimRight(value, cutset)
		default:
			return strings.Trim(value, cutset)
		}
	})
}

// TrimRepeatedString removes repeated literal occurrences from the selected
// ends.
func TrimRepeatedString(t *tag.Tag, position TrimPosition, repeated string) (TransformRule, error) {
	if err := validateTrimPosition(position); err != nil {
		return nil, err
	}
	if repeated == "" {
		return nil, fmt.Errorf("repeated trim string is empty")
	}
	return newEditValuesRule(t, func(value string) string {
		if position == TrimStart || position == TrimBoth {
			for strings.HasPrefix(value, repeated) {
				value = strings.TrimPrefix(value, repeated)
			}
		}
		if position == TrimEnd || position == TrimBoth {
			for strings.HasSuffix(value, repeated) {
				value = strings.TrimSuffix(value, repeated)
			}
		}
		return value
	})
}

// PadDirection selects the side padded by PadValue.
type PadDirection uint8

const (
	// PadLeft adds padding before each value.
	PadLeft PadDirection = iota
	// PadRight adds padding after each value.
	PadRight
)

// PadValue pads each value to totalLength Unicode code points.
func PadValue(t *tag.Tag, direction PadDirection, totalLength int, padding rune) (TransformRule, error) {
	if direction != PadLeft && direction != PadRight {
		return nil, fmt.Errorf("invalid pad direction %d", direction)
	}
	if totalLength < 0 {
		return nil, fmt.Errorf("pad length must not be negative")
	}
	return newEditValuesRule(t, func(value string) string {
		missing := totalLength - len([]rune(value))
		if missing <= 0 {
			return value
		}
		paddingText := strings.Repeat(string(padding), missing)
		if direction == PadLeft {
			return paddingText + value
		}
		return value + paddingText
	})
}

// TruncateValues limits every value to length Unicode code points.
func TruncateValues(t *tag.Tag, length int) (TransformRule, error) {
	if length < 0 {
		return nil, fmt.Errorf("truncate length must not be negative")
	}
	return newEditValuesRule(t, func(value string) string {
		runes := []rune(value)
		if len(runes) <= length {
			return value
		}
		return string(runes[:length])
	})
}

func newEditValuesRule(t *tag.Tag, edit func(string) string) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("edit tag is nil")
	}
	return editValuesRule{tag: t, edit: edit}, nil
}

func validateTrimPosition(position TrimPosition) error {
	if position != TrimStart && position != TrimEnd && position != TrimBoth {
		return fmt.Errorf("invalid trim position %d", position)
	}
	return nil
}

func (rule editValuesRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before, exists := ds.Get(rule.tag)
	if !exists {
		return nil
	}
	values, err := element.CanonicalStrings(before)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	edited := make([]string, len(values))
	for index, value := range values {
		edited[index] = rule.edit(value)
	}
	if reflect.DeepEqual(values, edited) {
		return nil
	}
	after, err := element.ReplaceCanonicalStrings(before, rule.tag, before.ValueRepresentation(), edited)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, rule.tag, err)
	}
	return appendChange(changes, Change{Kind: ChangeEdit, Path: changedElementPath(path, rule.tag), Tag: rule.tag, Before: before, After: after})
}
