// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parseable

import (
	"errors"
	"testing"
)

const testValue = "test"

func TestMustParse(t *testing.T) {
	t.Parallel()
	parser := StringParser

	result := MustParse(parser, testValue)
	if result != testValue {
		t.Errorf("MustParse() = %q, want %q", result, testValue)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParse did not panic on error")
		}
	}()

	errorParser := func(_ string) (string, error) {
		return "", errors.New("test error")
	}
	MustParse(errorParser, testValue)
}

func TestTryParse(t *testing.T) {
	t.Parallel()
	parser := StringParser

	result, ok := TryParse(parser, testValue)
	if !ok {
		t.Error("TryParse() ok = false, want true")
	}
	if result != testValue {
		t.Errorf("TryParse() = %q, want %q", result, testValue)
	}

	errorParser := func(_ string) (string, error) {
		return "", errors.New("test error")
	}
	_, ok = TryParse(errorParser, testValue)
	if ok {
		t.Error("TryParse() ok = true on error, want false")
	}
}

func TestParseSlice(t *testing.T) {
	t.Parallel()
	parser := StringParser

	result, err := ParseSlice(parser, []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("ParseSlice() error = %v", err)
	}
	if len(result) != 3 {
		t.Errorf("ParseSlice() len = %d, want 3", len(result))
	}
}

// testType is a test type that implements Parseable
type testType struct {
	value string
}

func (t *testType) Parse(value string) error {
	t.value = value
	return nil
}

func TestParserFor(t *testing.T) {
	t.Parallel()
	newFunc := func() *testType {
		return &testType{}
	}

	parser := ParserFor[*testType](newFunc)

	result, err := parser(testValue)
	if err != nil {
		t.Errorf("ParserFor() error = %v", err)
	}
	if result.value != testValue {
		t.Errorf("ParserFor() value = %q, want %q", result.value, testValue)
	}
}
