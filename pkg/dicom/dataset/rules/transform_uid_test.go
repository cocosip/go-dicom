// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestGenerateUIDsUsesStableInjectedMappingAcrossTags(t *testing.T) {
	generator := &deterministicUIDGenerator{
		values: map[string]*uid.UID{testRuleSourceUID: uid.New(testRuleMappedUID, "mapped", uid.TypeUnknown, false)},
	}
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.StudyInstanceUID, vr.UI, []string{testRuleSourceUID}))
	requireRuleAdd(t, source, element.NewString(tag.SeriesInstanceUID, vr.UI, []string{testRuleSourceUID}))
	transformer, err := NewTransformer(
		mustTransformRule(GenerateUIDs(tag.StudyInstanceUID, generator)),
		mustTransformRule(GenerateUIDs(tag.SeriesInstanceUID, generator)),
	)
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	study, _ := result.GetString(tag.StudyInstanceUID)
	series, _ := result.GetString(tag.SeriesInstanceUID)
	if study != testRuleMappedUID || series != testRuleMappedUID || generator.calls[testRuleSourceUID] != 2 {
		t.Fatalf("study=%q series=%q calls=%v", study, series, generator.calls)
	}
	if len(changes) != 2 || changes[0].Kind != ChangeUID || changes[1].Kind != ChangeUID {
		t.Fatalf("changes = %#v", changes)
	}
}

func TestGenerateUIDsPreservesEmptyValues(t *testing.T) {
	generator := &deterministicUIDGenerator{
		values: map[string]*uid.UID{testRuleSourceUID: uid.New(testRuleMappedUID, "mapped", uid.TypeUnknown, false)},
	}
	source := dataset.New()
	source.SetAutoValidate(false)
	requireRuleAdd(t, source, element.NewString(tag.SOPInstanceUID, vr.UI, []string{testRuleSourceUID, ""}))
	transformer, err := NewTransformer(mustTransformRule(GenerateUIDs(tag.SOPInstanceUID, generator)))
	if err != nil {
		t.Fatal(err)
	}
	result, _, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	elem, _ := result.Get(tag.SOPInstanceUID)
	values, err := element.CanonicalStrings(elem)
	if err != nil || !reflect.DeepEqual(values, []string{testRuleMappedUID, ""}) {
		t.Fatalf("values = %#v, err = %v", values, err)
	}
	if len(generator.calls) != 1 || generator.calls[""] != 0 {
		t.Fatalf("generator calls = %v", generator.calls)
	}
}

func TestGenerateUIDsRejectsInvalidConfigurationAndValues(t *testing.T) {
	if _, err := GenerateUIDs(tag.StudyInstanceUID, nil); err == nil {
		t.Fatal("GenerateUIDs accepted nil generator")
	}
	generator := &deterministicUIDGenerator{values: map[string]*uid.UID{}}
	if _, err := GenerateUIDs(nil, generator); err == nil {
		t.Fatal("GenerateUIDs accepted nil tag")
	}

	nonUI := dataset.New()
	requireRuleAdd(t, nonUI, element.NewString(tag.PatientID, vr.LO, []string{testRuleSourceUID}))
	transformer, err := NewTransformer(mustTransformRule(GenerateUIDs(tag.PatientID, generator)))
	if err != nil {
		t.Fatal(err)
	}
	if result, _, err := transformer.Apply(nonUI); err == nil || result != nil {
		t.Fatalf("non-UI result=%v err=%v", result, err)
	}

	invalid := dataset.New()
	invalid.SetAutoValidate(false)
	requireRuleAdd(t, invalid, element.NewString(tag.StudyInstanceUID, vr.UI, []string{"invalid"}))
	transformer, err = NewTransformer(mustTransformRule(GenerateUIDs(tag.StudyInstanceUID, generator)))
	if err != nil {
		t.Fatal(err)
	}
	if result, _, err := transformer.Apply(invalid); err == nil || result != nil {
		t.Fatalf("invalid UID result=%v err=%v", result, err)
	}
}

type deterministicUIDGenerator struct {
	values map[string]*uid.UID
	calls  map[string]int
}

func (generator *deterministicUIDGenerator) Generate(source *uid.UID) *uid.UID {
	if generator.calls == nil {
		generator.calls = make(map[string]int)
	}
	generator.calls[source.UID()]++
	return generator.values[source.UID()]
}
