// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// ValidationKind identifies the stage that rejected a Dataset value.
type ValidationKind string

const (
	// ValidationStructural identifies an invalid Dataset or element structure.
	ValidationStructural ValidationKind = "structural"
	// ValidationValue identifies an invalid value for the element's actual VR.
	ValidationValue ValidationKind = "value"
	// ValidationVM identifies a dictionary value-multiplicity mismatch.
	ValidationVM ValidationKind = "vm"
)

// ValidationPathSegment identifies an element or a Sequence item in a nested
// Dataset path. ItemIndex is set only for Sequence item segments.
type ValidationPathSegment struct {
	Tag       *tag.Tag
	ItemIndex *int
}

// ValidationError reports a Dataset validation failure and retains its cause.
type ValidationError struct {
	Kind  ValidationKind
	Path  []ValidationPathSegment
	Cause error
}

func (e *ValidationError) Error() string {
	if e == nil {
		return "<nil>"
	}
	path := formatValidationPath(e.Path)
	if e.Cause == nil {
		return fmt.Sprintf("DICOM %s validation failed at %s", e.Kind, path)
	}
	return fmt.Sprintf("DICOM %s validation failed at %s: %v", e.Kind, path, e.Cause)
}

// Unwrap returns the original element, VR, VM, or structural error.
func (e *ValidationError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

// Validate recursively validates every element in tag order. This explicit
// operation always validates, regardless of the Dataset automatic-validation
// setting.
func (ds *Dataset) Validate() error {
	return validateDataset(ds, nil)
}

func validateDataset(ds *Dataset, path []ValidationPathSegment) error {
	if ds == nil {
		return validationError(ValidationStructural, path, fmt.Errorf("dataset is nil"))
	}
	for _, elem := range ds.Elements() {
		if err := validateElement(elem, path); err != nil {
			return err
		}
	}
	return nil
}

func validateElement(elem element.Element, path []ValidationPathSegment) error {
	if elem == nil {
		return validationError(ValidationStructural, path, fmt.Errorf("element is nil"))
	}
	t := elem.Tag()
	if t == nil {
		return validationError(ValidationStructural, path, fmt.Errorf("element tag is nil"))
	}
	if sequence, ok := elem.(*Sequence); ok {
		return validateSequence(sequence, path)
	}

	elementPath := appendValidationPath(path, ValidationPathSegment{Tag: t})
	valueRepresentation := elem.ValueRepresentation()
	if valueRepresentation == nil {
		return validationError(ValidationStructural, elementPath, fmt.Errorf("element VR is nil"))
	}
	if err := element.ValidateValue(elem); err != nil {
		return validationError(ValidationValue, elementPath, err)
	}
	if err := validateElementVM(elem, valueRepresentation); err != nil {
		return validationError(ValidationVM, elementPath, err)
	}
	return nil
}

func validateSequence(sequence *Sequence, path []ValidationPathSegment) error {
	if sequence == nil {
		return validationError(ValidationStructural, path, fmt.Errorf("sequence is nil"))
	}
	if sequence.tag == nil {
		return validationError(ValidationStructural, path, fmt.Errorf("sequence tag is nil"))
	}
	for index, item := range sequence.items {
		if item == nil {
			continue
		}
		itemIndex := index
		itemPath := appendValidationPath(path, ValidationPathSegment{
			Tag:       sequence.tag,
			ItemIndex: &itemIndex,
		})
		if err := validateDataset(item, itemPath); err != nil {
			return err
		}
	}
	return nil
}

func validateElementVM(elem element.Element, valueRepresentation *vr.VR) error {
	if elem.Count() == 0 || elem.Tag().IsPrivate() || isVMExempt(valueRepresentation) {
		return nil
	}
	entry := dict.Default().Lookup(elem.Tag())
	if entry == nil || entry.ValueMultiplicity() == nil {
		return nil
	}
	valueMultiplicity := entry.ValueMultiplicity()
	if valueMultiplicity.IsValid(elem.Count()) {
		return nil
	}
	return fmt.Errorf("number of values %d does not match VM %s", elem.Count(), valueMultiplicity)
}

func isVMExempt(valueRepresentation *vr.VR) bool {
	switch valueRepresentation.Code() {
	case vr.CodeSQ, vr.CodeOB, vr.CodeOW, vr.CodeOL, vr.CodeOD, vr.CodeOF, vr.CodeOV, vr.CodeUC, vr.CodeUN:
		return true
	default:
		return false
	}
}

func validationError(kind ValidationKind, path []ValidationPathSegment, cause error) error {
	return &ValidationError{
		Kind:  kind,
		Path:  append([]ValidationPathSegment(nil), path...),
		Cause: cause,
	}
}

func appendValidationPath(path []ValidationPathSegment, segment ValidationPathSegment) []ValidationPathSegment {
	result := make([]ValidationPathSegment, len(path), len(path)+1)
	copy(result, path)
	return append(result, segment)
}

func formatValidationPath(path []ValidationPathSegment) string {
	if len(path) == 0 {
		return "<dataset>"
	}
	parts := make([]string, 0, len(path))
	for _, segment := range path {
		part := "<unknown-tag>"
		if segment.Tag != nil {
			part = segment.Tag.String()
		}
		if segment.ItemIndex != nil {
			part += fmt.Sprintf("[%d]", *segment.ItemIndex)
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, "/")
}
