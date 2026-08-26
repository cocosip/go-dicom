// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type splitFormatPart struct {
	literal string
	index   int
	isIndex bool
}

type splitFormatRule struct {
	tag        *tag.Tag
	separators map[rune]struct{}
	parts      []splitFormatPart
}

// SplitFormat splits a tag's complete canonical value on any separator rune
// and builds one replacement canonical value from indexed format placeholders.
// Placeholders use {0}, {1}, and so on; {{ and }} represent literal braces.
func SplitFormat(t *tag.Tag, separators string, format string) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("split format tag is nil")
	}
	if separators == "" {
		return nil, fmt.Errorf("split format separators are empty")
	}
	parts, err := parseSplitFormat(format)
	if err != nil {
		return nil, err
	}
	separatorSet := make(map[rune]struct{}, utf8.RuneCountInString(separators))
	for _, separator := range separators {
		separatorSet[separator] = struct{}{}
	}
	return splitFormatRule{tag: t, separators: separatorSet, parts: parts}, nil
}

func parseSplitFormat(format string) ([]splitFormatPart, error) {
	parts := make([]splitFormatPart, 0)
	var literal strings.Builder
	flushLiteral := func() {
		if literal.Len() > 0 {
			parts = append(parts, splitFormatPart{literal: literal.String()})
			literal.Reset()
		}
	}

	for index := 0; index < len(format); {
		switch format[index] {
		case '{':
			if index+1 < len(format) && format[index+1] == '{' {
				literal.WriteByte('{')
				index += 2
				continue
			}
			flushLiteral()
			end := strings.IndexByte(format[index+1:], '}')
			if end < 0 {
				return nil, fmt.Errorf("split format has unclosed placeholder at byte %d", index)
			}
			end += index + 1
			placeholder := format[index+1 : end]
			if placeholder == "" {
				return nil, fmt.Errorf("split format has empty placeholder at byte %d", index)
			}
			for _, character := range placeholder {
				if character < '0' || character > '9' {
					return nil, fmt.Errorf("split format placeholder %q is not a non-negative index", placeholder)
				}
			}
			value, err := strconv.Atoi(placeholder)
			if err != nil {
				return nil, fmt.Errorf("parse split format placeholder %q: %w", placeholder, err)
			}
			parts = append(parts, splitFormatPart{index: value, isIndex: true})
			index = end + 1
		case '}':
			if index+1 < len(format) && format[index+1] == '}' {
				literal.WriteByte('}')
				index += 2
				continue
			}
			return nil, fmt.Errorf("split format has unmatched closing brace at byte %d", index)
		default:
			runeValue, size := utf8.DecodeRuneInString(format[index:])
			literal.WriteRune(runeValue)
			index += size
		}
	}
	flushLiteral()
	return parts, nil
}

func (rule splitFormatRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before, exists := ds.Get(rule.tag)
	if !exists {
		return nil
	}
	values, err := element.CanonicalStrings(before)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	fields := splitOnRunes(strings.Join(values, "\\"), rule.separators)
	formatted, err := rule.format(fields)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	edited := []string{formatted}
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

func (rule splitFormatRule) format(fields []string) (string, error) {
	var output strings.Builder
	for _, part := range rule.parts {
		if !part.isIndex {
			output.WriteString(part.literal)
			continue
		}
		if part.index >= len(fields) {
			return "", fmt.Errorf("split format index %d is out of range for %d fields", part.index, len(fields))
		}
		output.WriteString(fields[part.index])
	}
	return output.String(), nil
}

func splitOnRunes(value string, separators map[rune]struct{}) []string {
	fields := make([]string, 0, 1)
	start := 0
	for index, character := range value {
		if _, isSeparator := separators[character]; !isSeparator {
			continue
		}
		fields = append(fields, value[start:index])
		start = index + utf8.RuneLen(character)
	}
	return append(fields, value[start:])
}
