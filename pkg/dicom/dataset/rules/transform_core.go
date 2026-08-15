// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

type removeRule struct {
	tag  *tag.Tag
	mask *tag.MaskedTag
}

// Remove removes one exact Tag.
func Remove(t *tag.Tag) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("remove tag is nil")
	}
	return removeRule{tag: t}, nil
}

// RemoveMasked removes every matching Tag in Dataset tag order.
func RemoveMasked(mask *tag.MaskedTag) (TransformRule, error) {
	if mask == nil || mask.Tag() == nil {
		return nil, fmt.Errorf("remove mask is nil")
	}
	return removeRule{mask: tag.NewMaskedTagWithMask(mask.Tag(), mask.Mask())}, nil
}

func (rule removeRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	if rule.tag != nil {
		before, exists := ds.Get(rule.tag)
		if !exists {
			return nil
		}
		ds.Remove(rule.tag)
		appendChange(changes, Change{Kind: ChangeRemove, Path: changedElementPath(path, rule.tag), Tag: rule.tag, Before: before})
		return nil
	}

	matching := make([]*tag.Tag, 0)
	for _, candidate := range ds.Tags() {
		if rule.mask.IsMatch(candidate) {
			matching = append(matching, candidate)
		}
	}
	for _, candidate := range matching {
		before, _ := ds.Get(candidate)
		ds.Remove(candidate)
		appendChange(changes, Change{Kind: ChangeRemove, Path: changedElementPath(path, candidate), Tag: candidate, Before: before})
	}
	return nil
}

type setElementRule struct{ elem element.Element }

// SetElement assigns an independent clone of elem.
func SetElement(elem element.Element) (TransformRule, error) {
	if elem == nil {
		return nil, fmt.Errorf("set element is nil")
	}
	if elem.Tag() == nil || elem.ValueRepresentation() == nil {
		return nil, fmt.Errorf("set element tag and VR are required")
	}
	if err := element.ValidateValue(elem); err != nil {
		return nil, fmt.Errorf("validate set element: %w", err)
	}
	return setElementRule{elem: element.DeepClone(elem)}, nil
}

func (rule setElementRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before := ds.GetOrNil(rule.elem.Tag())
	after := element.DeepClone(rule.elem)
	if elementsEquivalent(before, after) {
		return nil
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, after.Tag(), err)
	}
	appendChange(changes, Change{Kind: ChangeAssign, Path: changedElementPath(path, after.Tag()), Tag: after.Tag(), Before: before, After: after})
	return nil
}

type setStringsRule struct {
	tag      *tag.Tag
	explicit *vr.VR
	values   []string
}

// SetStrings assigns canonical values using the existing, explicit, or
// unambiguous dictionary VR in that order.
func SetStrings(t *tag.Tag, explicitVR *vr.VR, values ...string) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("set strings tag is nil")
	}
	return setStringsRule{tag: t, explicit: explicitVR, values: append([]string(nil), values...)}, nil
}

func (rule setStringsRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	valueRepresentation, prototype, err := resolveDestinationVR(ds, rule.tag, rule.explicit)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	after, err := element.ReplaceCanonicalStrings(prototype, rule.tag, valueRepresentation, rule.values)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	if elementsEquivalent(prototype, after) {
		return nil
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, rule.tag, err)
	}
	appendChange(changes, Change{Kind: ChangeAssign, Path: changedElementPath(path, rule.tag), Tag: rule.tag, Before: prototype, After: after})
	return nil
}

type mapValueRule struct {
	tag      *tag.Tag
	from, to string
}

// MapValue replaces one complete canonical value.
func MapValue(t *tag.Tag, from, to string) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("map tag is nil")
	}
	return mapValueRule{tag: t, from: from, to: to}, nil
}

func (rule mapValueRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before, exists := ds.Get(rule.tag)
	if !exists {
		return nil
	}
	values, err := element.CanonicalStrings(before)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	if strings.Join(values, "\\") != rule.from || rule.from == rule.to {
		return nil
	}
	after, err := element.ReplaceCanonicalStrings(before, rule.tag, before.ValueRepresentation(), strings.Split(rule.to, "\\"))
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, rule.tag, err)
	}
	appendChange(changes, Change{Kind: ChangeMap, Path: changedElementPath(path, rule.tag), Tag: rule.tag, Before: before, After: after})
	return nil
}

type copyValueRule struct {
	source, destination *tag.Tag
	explicitDestination *vr.VR
}

// CopyValue converts source canonical values into the destination VR.
func CopyValue(source, destination *tag.Tag, explicitDestinationVR *vr.VR) (TransformRule, error) {
	if source == nil || destination == nil {
		return nil, fmt.Errorf("copy source and destination tags are required")
	}
	return copyValueRule{source: source, destination: destination, explicitDestination: explicitDestinationVR}, nil
}

func (rule copyValueRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	source, exists := ds.Get(rule.source)
	if !exists {
		return nil
	}
	values, err := element.CanonicalStrings(source)
	if err != nil {
		return transformValueError(path, rule.source, err)
	}
	valueRepresentation, before, err := resolveDestinationVR(ds, rule.destination, rule.explicitDestination)
	if err != nil {
		return transformValueError(path, rule.destination, err)
	}
	after, err := element.ReplaceCanonicalStrings(before, rule.destination, valueRepresentation, values)
	if err != nil {
		return transformValueError(path, rule.destination, err)
	}
	if elementsEquivalent(before, after) {
		return nil
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, rule.destination, err)
	}
	appendChange(changes, Change{Kind: ChangeCopy, Path: changedElementPath(path, rule.destination), Tag: rule.destination, Before: before, After: after})
	return nil
}
