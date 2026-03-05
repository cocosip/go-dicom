// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package parseable provides the Parseable interface for types that can be parsed from strings.
// This is used in DICOM element value retrieval to support parsing string values into
// specific types like DicomUID, DicomTransferSyntax, etc.
package parseable

import "fmt"

// Parseable is an interface for types that can be parsed from a string.
// Types implementing this interface can be used in DICOM value parsing contexts.
//
// Example usage:
//
//	type MyType struct{}
//	func (m *MyType) Parse(value string) error {
//	    // parse logic
//	    return nil
//	}
type Parseable interface {
	Parse(value string) error
}

// Parser is a function type that parses a string into type T.
type Parser[T any] func(value string) (T, error)

// MustParse parses the value using the provided parser, panicking on error.
// Use this for cases where parsing failure indicates a programming error.
func MustParse[T any](parser Parser[T], value string) T {
	result, err := parser(value)
	if err != nil {
		panic(err)
	}
	return result
}

// TryParse attempts to parse the value, returning the result and a boolean indicating success.
// Use this when you want to handle parsing failures gracefully without error handling.
func TryParse[T any](parser Parser[T], value string) (T, bool) {
	result, err := parser(value)
	if err != nil {
		var zero T
		return zero, false
	}
	return result, true
}

// ParseSlice parses a slice of strings using the provided parser.
// Returns an error if any element fails to parse.
func ParseSlice[T any](parser Parser[T], values []string) ([]T, error) {
	result := make([]T, len(values))
	for i, v := range values {
		parsed, err := parser(v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value at index %d: %w", i, err)
		}
		result[i] = parsed
	}
	return result, nil
}

// StringParser is a Parser that returns the input string unchanged.
func StringParser(value string) (string, error) {
	return value, nil
}

// ParserFor returns a Parser function that calls the Parse method on a new instance of T.
// This is useful for types that implement Parseable.
func ParserFor[T Parseable](newFunc func() T) Parser[T] {
	return func(value string) (T, error) {
		instance := newFunc()
		if err := instance.Parse(value); err != nil {
			var zero T
			return zero, err
		}
		return instance, nil
	}
}
