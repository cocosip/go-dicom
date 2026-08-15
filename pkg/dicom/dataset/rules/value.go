// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func resolveDestinationVR(ds *dataset.Dataset, target *tag.Tag, explicit *vr.VR) (*vr.VR, element.Element, error) {
	if existing, ok := ds.Get(target); ok {
		if existing.ValueRepresentation() == nil {
			return nil, existing, fmt.Errorf("existing element %s has nil VR", target)
		}
		return existing.ValueRepresentation(), existing, nil
	}
	if explicit != nil {
		return explicit, nil, nil
	}
	entry := dict.Default().Lookup(target)
	if entry == nil {
		return nil, nil, fmt.Errorf("no dictionary VR for missing element %s; explicit VR is required", target)
	}
	valueRepresentations := entry.ValueRepresentations()
	if len(valueRepresentations) != 1 || valueRepresentations[0] == nil || valueRepresentations[0] == vr.None {
		return nil, nil, fmt.Errorf("dictionary VR for missing element %s is ambiguous; explicit VR is required", target)
	}
	return valueRepresentations[0], nil, nil
}

func replaceCanonicalStrings(
	ds *dataset.Dataset,
	prototype element.Element,
	target *tag.Tag,
	valueRepresentation *vr.VR,
	values []string,
) (element.Element, error) {
	if prototype != nil {
		return element.ReplaceCanonicalStrings(prototype, target, valueRepresentation, values)
	}
	context := element.CanonicalValueContext{TextEncodings: charset.GetEncodings(nil)}
	if charsets, ok := ds.GetStrings(tag.SpecificCharacterSet); ok {
		context.TextEncodings = charset.GetEncodings(charsets)
	}
	if syntax := ds.InternalTransferSyntax(); syntax != nil {
		context.Endian = syntax.Endian()
	}
	return element.ReplaceCanonicalStringsWithContext(target, valueRepresentation, values, context)
}

func changedElementPath(path dataset.Path, target *tag.Tag) dataset.Path {
	return append(dataset.ClonePath(path), dataset.PathSegment{Tag: target})
}

func transformValueError(path dataset.Path, target *tag.Tag, err error) error {
	if err == nil {
		return nil
	}
	return &TransformError{
		Stage: StageRule,
		Path:  changedElementPath(path, target),
		Cause: err,
	}
}

func elementsEquivalent(left, right element.Element) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	if reflect.DeepEqual(left, right) {
		return true
	}
	if left.Tag() == nil || right.Tag() == nil || !left.Tag().Equals(right.Tag()) {
		return false
	}
	if left.ValueRepresentation() == nil || right.ValueRepresentation() == nil ||
		left.ValueRepresentation().Code() != right.ValueRepresentation().Code() || left.Count() != right.Count() {
		return false
	}
	leftBuffer, rightBuffer := left.Buffer(), right.Buffer()
	if leftBuffer == nil || rightBuffer == nil {
		return leftBuffer == nil && rightBuffer == nil
	}
	return bytes.Equal(leftBuffer.Data(), rightBuffer.Data())
}
