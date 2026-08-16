// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func BenchmarkMatchRuleSet(b *testing.B) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{testRulePatientID}))
	exists, _ := Exists(tag.PatientID)
	wildcard, _ := Wildcard(tag.PatientID, "case-*")
	contains, _ := Contains(tag.PatientID, "-12")
	rule, _ := All(exists, wildcard, contains)
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_, _ = rule.Match(ds)
	}
}

func BenchmarkTransformRuleSet(b *testing.B) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"case-123"}))
	upper, _ := ToUpper(tag.PatientID)
	prefix, _ := PrefixValue(tag.PatientID, "ID:")
	transformer, _ := NewTransformer(upper, prefix)
	b.ResetTimer()
	for iteration := 0; iteration < b.N; iteration++ {
		_, _, _ = transformer.Apply(ds)
	}
}
