// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import "fmt"

// Error represents a structured report error
type Error struct {
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("sr: %s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("sr: %s", e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// NewError creates a new SR error
func NewError(message string) error {
	return &Error{Message: message}
}

// NewErrorf creates a new SR error with formatting
func NewErrorf(format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...)}
}

// WrapError wraps an error with additional context
func WrapError(message string, err error) error {
	return &Error{Message: message, Err: err}
}

// Common errors
var (
	// ErrInvalidValueType indicates an invalid value type for the operation
	ErrInvalidValueType = NewError("invalid value type for operation")

	// ErrInvalidRelationship indicates an invalid relationship type
	ErrInvalidRelationship = NewError("invalid relationship type")

	// ErrEmptySequence indicates an empty sequence where data was expected
	ErrEmptySequence = NewError("empty sequence, no items found")

	// ErrMissingCode indicates a required code item is missing
	ErrMissingCode = NewError("required code item is missing")
)
