// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// MatchRule evaluates one Dataset without mutating it.
type MatchRule interface {
	Match(*dataset.Dataset) (bool, error)
}

type matchFunc func(*dataset.Dataset) (bool, error)

func (match matchFunc) Match(ds *dataset.Dataset) (bool, error) {
	return match(ds)
}

// Bool returns a rule with a constant result.
func Bool(value bool) MatchRule {
	return matchFunc(func(*dataset.Dataset) (bool, error) { return value, nil })
}

// All matches when every child rule matches. Evaluation short-circuits in
// declaration order. An empty All matches.
func All(rules ...MatchRule) (MatchRule, error) {
	copied, err := copyMatchRules(rules)
	if err != nil {
		return nil, err
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		for _, rule := range copied {
			matched, err := rule.Match(ds)
			if err != nil || !matched {
				return matched, err
			}
		}
		return true, nil
	}), nil
}

// Any matches when at least one child rule matches. Evaluation short-circuits
// in declaration order. An empty Any does not match.
func Any(rules ...MatchRule) (MatchRule, error) {
	copied, err := copyMatchRules(rules)
	if err != nil {
		return nil, err
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		for _, rule := range copied {
			matched, err := rule.Match(ds)
			if err != nil || matched {
				return matched, err
			}
		}
		return false, nil
	}), nil
}

// Not negates a child rule.
func Not(rule MatchRule) (MatchRule, error) {
	if isNilRule(rule) {
		return nil, fmt.Errorf("match rule is nil")
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		matched, err := rule.Match(ds)
		return !matched, err
	}), nil
}

// Exists matches a stored element, including empty values and containers.
func Exists(t *tag.Tag) (MatchRule, error) {
	if t == nil {
		return nil, fmt.Errorf("match tag is nil")
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		if ds == nil {
			return false, fmt.Errorf("match Dataset is nil")
		}
		return ds.Contains(t), nil
	}), nil
}

// Empty matches a missing element, zero-count value, or empty complete text.
func Empty(t *tag.Tag) (MatchRule, error) {
	if t == nil {
		return nil, fmt.Errorf("match tag is nil")
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		if ds == nil {
			return false, fmt.Errorf("match Dataset is nil")
		}
		elem, exists := ds.Get(t)
		if !exists || elem.Count() == 0 {
			return true, nil
		}
		values, err := element.CanonicalStrings(elem)
		if err != nil {
			return false, err
		}
		return strings.Join(values, "\\") == "", nil
	}), nil
}

// Equal matches the complete canonical value case-sensitively.
func Equal(t *tag.Tag, value string) (MatchRule, error) {
	return contentRule(t, func(actual string) bool { return actual == value })
}

// StartsWith matches a case-sensitive complete-value prefix.
func StartsWith(t *tag.Tag, value string) (MatchRule, error) {
	return contentRule(t, func(actual string) bool { return strings.HasPrefix(actual, value) })
}

// EndsWith matches a case-sensitive complete-value suffix.
func EndsWith(t *tag.Tag, value string) (MatchRule, error) {
	return contentRule(t, func(actual string) bool { return strings.HasSuffix(actual, value) })
}

// Contains matches a case-sensitive complete-value substring.
func Contains(t *tag.Tag, value string) (MatchRule, error) {
	return contentRule(t, func(actual string) bool { return strings.Contains(actual, value) })
}

// OneOf matches one of the supplied complete canonical values.
func OneOf(t *tag.Tag, values ...string) (MatchRule, error) {
	copied := append([]string(nil), values...)
	return contentRule(t, func(actual string) bool {
		for _, value := range copied {
			if actual == value {
				return true
			}
		}
		return false
	})
}

// Wildcard matches an entire complete value using case-insensitive * and ?
// wildcard syntax.
func Wildcard(t *tag.Tag, pattern string) (MatchRule, error) {
	quoted := regexp.QuoteMeta(pattern)
	quoted = strings.ReplaceAll(quoted, `\*`, `.*`)
	quoted = strings.ReplaceAll(quoted, `\?`, `.`)
	compiled, err := regexp.Compile(`(?i)^(?:` + quoted + `)$`)
	if err != nil {
		return nil, fmt.Errorf("compile wildcard pattern %q: %w", pattern, err)
	}
	return contentRule(t, compiled.MatchString)
}

// Regex matches a complete canonical value using a compiled Go RE2 regular
// expression. Anchoring and case flags are controlled by the caller's pattern.
func Regex(t *tag.Tag, pattern string) (MatchRule, error) {
	if t == nil {
		return nil, fmt.Errorf("match tag is nil")
	}
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("compile regular expression %q: %w", pattern, err)
	}
	return contentRule(t, compiled.MatchString)
}

func contentRule(t *tag.Tag, predicate func(string) bool) (MatchRule, error) {
	if t == nil {
		return nil, fmt.Errorf("match tag is nil")
	}
	return matchFunc(func(ds *dataset.Dataset) (bool, error) {
		if ds == nil {
			return false, fmt.Errorf("match Dataset is nil")
		}
		elem, exists := ds.Get(t)
		if !exists {
			return false, nil
		}
		values, err := element.CanonicalStrings(elem)
		if err != nil {
			return false, err
		}
		return predicate(strings.Join(values, "\\")), nil
	}), nil
}

func copyMatchRules(rules []MatchRule) ([]MatchRule, error) {
	copied := append([]MatchRule(nil), rules...)
	for index, rule := range copied {
		if isNilRule(rule) {
			return nil, fmt.Errorf("match rule %d is nil", index)
		}
	}
	return copied, nil
}

func isNilRule(rule any) bool {
	if rule == nil {
		return true
	}
	value := reflect.ValueOf(rule)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}
