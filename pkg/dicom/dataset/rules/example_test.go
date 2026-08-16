// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func ExampleAll() {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{testRulePatientID}))
	exists, _ := Exists(tag.PatientID)
	pattern, _ := Wildcard(tag.PatientID, "case-*")
	rule, _ := All(exists, pattern)

	matched, _ := rule.Match(ds)
	fmt.Println(matched)

	// Output: true
}

func ExampleTransformer_Apply() {
	source := dataset.New()
	_ = source.Add(element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	prefix, _ := PrefixValue(tag.PatientID, "ID:")
	transformer, _ := NewTransformer(prefix)

	result, changes, _ := transformer.Apply(source)
	sourceValue, _ := source.GetString(tag.PatientID)
	resultValue, _ := result.GetString(tag.PatientID)
	fmt.Println(sourceValue, resultValue, len(changes))

	// Output: 123 ID:123 1
}

func ExampleTransformer_ApplyInPlace() {
	target := dataset.New()
	_ = target.Add(element.NewString(tag.PatientID, vr.LO, []string{"abc"}))
	original := target
	upper, _ := ToUpper(tag.PatientID)
	transformer, _ := NewTransformer(upper)

	changes, _ := transformer.ApplyInPlace(target)
	value, _ := target.GetString(tag.PatientID)
	fmt.Println(target == original, value, len(changes))

	// Output: true ABC 1
}
