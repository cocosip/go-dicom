// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"
	"reflect"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// UIDGenerator provides stable source-to-destination UID mapping. The existing
// *uid.Generator implements this interface.
type UIDGenerator interface {
	Generate(*uid.UID) *uid.UID
}

type generateUIDsRule struct {
	tag       *tag.Tag
	generator UIDGenerator
}

// GenerateUIDs replaces each non-empty UI value through generator.
func GenerateUIDs(t *tag.Tag, generator UIDGenerator) (TransformRule, error) {
	if t == nil {
		return nil, fmt.Errorf("UID tag is nil")
	}
	if isNilRule(generator) {
		return nil, fmt.Errorf("UID generator is nil")
	}
	return generateUIDsRule{tag: t, generator: generator}, nil
}

func (rule generateUIDsRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before, exists := ds.Get(rule.tag)
	if !exists {
		return nil
	}
	if before.ValueRepresentation() == nil || before.ValueRepresentation().Code() != vr.CodeUI {
		return transformValueError(path, rule.tag, fmt.Errorf("UID transform requires UI VR"))
	}
	values, err := element.CanonicalStrings(before)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	mapped := append([]string(nil), values...)
	for index, value := range values {
		if value == "" {
			continue
		}
		if !uid.IsValid(value) {
			return transformValueError(path, rule.tag, fmt.Errorf("UID value[%d] %q is invalid", index, value))
		}
		destination := rule.generator.Generate(uid.Parse(value, "Unknown", uid.TypeUnknown))
		if destination == nil || !uid.IsValid(destination.UID()) {
			return transformValueError(path, rule.tag, fmt.Errorf("UID generator returned an invalid value for value[%d]", index))
		}
		mapped[index] = destination.UID()
	}
	if reflect.DeepEqual(values, mapped) {
		return nil
	}
	after, err := element.ReplaceCanonicalStrings(before, rule.tag, vr.UI, mapped)
	if err != nil {
		return transformValueError(path, rule.tag, err)
	}
	if err := ds.AddOrUpdate(after); err != nil {
		return transformValueError(path, rule.tag, err)
	}
	appendChange(changes, Change{Kind: ChangeUID, Path: changedElementPath(path, rule.tag), Tag: rule.tag, Before: before, After: after})
	return nil
}
