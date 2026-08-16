// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

const originalPatientID = "original"

func TestRemoveAndRemoveMaskedRecordStableTagOrder(t *testing.T) {
	first := tag.New(0x0011, 0x0010)
	second := tag.New(0x0011, 0x0020)
	third := tag.New(0x0013, 0x0010)
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(second, vr.LO, []string{"second"}))
	requireRuleAdd(t, source, element.NewString(first, vr.LO, []string{"first"}))
	requireRuleAdd(t, source, element.NewString(third, vr.LO, []string{"third"}))

	mask := tag.NewMaskedTagWithMask(tag.New(0x0011, 0x0000), 0xffff0000)
	transformer, err := NewTransformer(mustTransformRule(RemoveMasked(mask)))
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if result.Contains(first) || result.Contains(second) || !result.Contains(third) {
		t.Fatal("RemoveMasked removed the wrong tags")
	}
	gotTags := []*tag.Tag{changes[0].Tag, changes[1].Tag}
	if !reflect.DeepEqual(gotTags, []*tag.Tag{first, second}) {
		t.Fatalf("change tags = %v, want sorted %v", gotTags, []*tag.Tag{first, second})
	}
	for _, change := range changes {
		if change.Kind != ChangeRemove || change.Before == nil || change.After != nil {
			t.Fatalf("remove change = %#v", change)
		}
	}

	exact, err := NewTransformer(mustTransformRule(Remove(third)))
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err = exact.Apply(result)
	if err != nil || result.Contains(third) || len(changes) != 1 {
		t.Fatalf("exact remove result=%v changes=%v err=%v", result.Contains(third), changes, err)
	}
}

func TestSetElementClonesConstructorInput(t *testing.T) {
	input := element.NewString(tag.PatientID, vr.LO, []string{originalPatientID})
	rule := mustTransformRule(SetElement(input))
	input.Buffer().Data()[0] = 'X'
	transformer, err := NewTransformer(rule)
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(dataset.New())
	if err != nil {
		t.Fatal(err)
	}
	if value, _ := result.GetString(tag.PatientID); value != originalPatientID {
		t.Fatalf("PatientID = %q, constructor input was not cloned", value)
	}
	if len(changes) != 1 || changes[0].Before != nil || changes[0].Kind != ChangeAssign {
		t.Fatalf("changes = %#v", changes)
	}
}

func TestChangeTagsAndPathsDoNotAliasTransformedDataset(t *testing.T) {
	privateTag := tag.NewWithPrivateCreator(0x0011, 0x1010, tag.NewPrivateCreator(testRuleOriginal))
	transformer, err := NewTransformer(mustTransformRule(SetStrings(privateTag, vr.LO, "value")))
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(dataset.New())
	if err != nil {
		t.Fatal(err)
	}
	changes[0].Tag.SetPrivateCreator(tag.NewPrivateCreator("CHANGE-TAG"))
	changes[0].Path[0].Tag.SetPrivateCreator(tag.NewPrivateCreator("CHANGE-PATH"))
	resultTag := result.GetOrNil(privateTag).Tag()
	if creator := resultTag.PrivateCreator(); creator == nil || creator.Creator() != testRuleOriginal {
		t.Fatalf("change metadata mutation changed result Tag: %v", resultTag)
	}
	if changes[0].After.Tag() == resultTag {
		t.Fatal("Change.After Tag aliases transformed Dataset Tag")
	}
}

func TestSetElementPreservesPopulatedSequenceAndClonesItsItems(t *testing.T) {
	item := dataset.New()
	requireRuleAdd(t, item, element.NewString(tag.PatientID, vr.LO, []string{originalPatientID}))
	sequence := dataset.NewSequenceWithItems(tag.SourceImageSequence, []*dataset.Dataset{item})
	rule := mustTransformRule(SetElement(sequence))
	if err := item.AddOrUpdate(element.NewString(tag.PatientID, vr.LO, []string{"changed"})); err != nil {
		t.Fatal(err)
	}
	transformer, err := NewTransformer(rule)
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(dataset.New())
	if err != nil {
		t.Fatal(err)
	}
	got, sequenceErr := result.GetSequence(tag.SourceImageSequence)
	if sequenceErr != nil {
		t.Fatalf("assigned element = %T, want Sequence: %v", result.GetOrNil(tag.SourceImageSequence), sequenceErr)
	}
	if got.Count() != 1 {
		t.Fatalf("assigned Sequence count = %d, want 1", got.Count())
	}
	if value, _ := got.GetItem(0).GetString(tag.PatientID); value != originalPatientID {
		t.Fatalf("assigned Sequence item PatientID = %q, want constructor snapshot", value)
	}
	changeAfter, isSequence := changes[0].After.(*dataset.Sequence)
	if !isSequence || changeAfter.Count() != 1 {
		t.Fatalf("change After = %T, want populated Sequence", changes[0].After)
	}
	got.GetItem(0).Remove(tag.PatientID)
	if value, _ := changeAfter.GetItem(0).GetString(tag.PatientID); value != originalPatientID {
		t.Fatalf("change snapshot aliased result Sequence: %q", value)
	}
}

func TestRemoveSequenceRecordsIndependentPopulatedBeforeSnapshot(t *testing.T) {
	item := dataset.New()
	requireRuleAdd(t, item, element.NewString(tag.PatientID, vr.LO, []string{originalPatientID}))
	source := dataset.New()
	requireRuleAdd(t, source, dataset.NewSequenceWithItems(tag.SourceImageSequence, []*dataset.Dataset{item}))
	transformer, err := NewTransformer(mustTransformRule(Remove(tag.SourceImageSequence)))
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if result.Contains(tag.SourceImageSequence) || len(changes) != 1 {
		t.Fatalf("result contains Sequence=%t changes=%v", result.Contains(tag.SourceImageSequence), changes)
	}
	before, isSequence := changes[0].Before.(*dataset.Sequence)
	if !isSequence || before.Count() != 1 {
		t.Fatalf("change Before = %T, want populated Sequence", changes[0].Before)
	}
	item.Remove(tag.PatientID)
	if value, _ := before.GetItem(0).GetString(tag.PatientID); value != originalPatientID {
		t.Fatalf("change snapshot aliased source Sequence: %q", value)
	}
}

func TestSetStringsResolvesExistingExplicitAndDictionaryVR(t *testing.T) {
	tests := []struct {
		name       string
		target     *tag.Tag
		existing   element.Element
		explicitVR *vr.VR
		values     []string
		wantVR     *vr.VR
	}{
		{name: "existing wins", target: tag.PatientID, existing: element.NewString(tag.PatientID, vr.LO, []string{"old"}), explicitVR: vr.SH, values: []string{"new"}, wantVR: vr.LO},
		{name: "explicit private", target: tag.New(0x0011, 0x1010), explicitVR: vr.LO, values: []string{"private"}, wantVR: vr.LO},
		{name: "dictionary", target: tag.StudyDescription, values: []string{"description"}, wantVR: vr.LO},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			source := dataset.New()
			if tc.existing != nil {
				requireRuleAdd(t, source, tc.existing)
			}
			transformer, err := NewTransformer(mustTransformRule(SetStrings(tc.target, tc.explicitVR, tc.values...)))
			if err != nil {
				t.Fatal(err)
			}
			result, changes, err := transformer.Apply(source)
			if err != nil {
				t.Fatal(err)
			}
			elem, _ := result.Get(tc.target)
			if elem.ValueRepresentation() != tc.wantVR || len(changes) != 1 {
				t.Fatalf("VR = %v, changes = %v", elem.ValueRepresentation(), changes)
			}
			got, err := element.CanonicalStrings(elem)
			if err != nil || !reflect.DeepEqual(got, tc.values) {
				t.Fatalf("values = %v, err = %v", got, err)
			}
		})
	}
}

func TestSetStringsUsesDatasetSpecificCharacterSetForMissingText(t *testing.T) {
	source := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	requireRuleAdd(t, source, element.NewString(tag.SpecificCharacterSet, vr.CS, []string{"ISO_IR 192"}))
	transformer, err := NewTransformer(mustTransformRule(SetStrings(tag.PatientName, nil, "张^三")))
	if err != nil {
		t.Fatal(err)
	}

	result, _, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	var encoded bytes.Buffer
	if err := writer.Write(&encoded, result); err != nil {
		t.Fatal(err)
	}
	parsed, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	if value, _ := parsed.Dataset.GetString(tag.PatientName); value != "张^三" {
		t.Fatalf("round-trip PatientName = %q, want %q", value, "张^三")
	}
}

func TestSetStringsUsesDatasetTransferSyntaxForMissingNumeric(t *testing.T) {
	source := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	transformer, err := NewTransformer(mustTransformRule(SetStrings(tag.Rows, vr.US, "4660")))
	if err != nil {
		t.Fatal(err)
	}

	result, _, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	var encoded bytes.Buffer
	if err := writer.Write(&encoded, result); err != nil {
		t.Fatal(err)
	}
	parsed, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	value, err := parsed.Dataset.GetUInt16(tag.Rows, 0)
	if err != nil || value != 0x1234 {
		t.Fatalf("round-trip Rows = %#x, error = %v, want 0x1234", value, err)
	}
}

func TestSetStringsRejectsUnknownAndAmbiguousMissingVR(t *testing.T) {
	private := tag.New(0x0011, 0x1010)
	for _, target := range []*tag.Tag{private, tag.SmallestImagePixelValue} {
		transformer, err := NewTransformer(mustTransformRule(SetStrings(target, nil, "1")))
		if err != nil {
			t.Fatal(err)
		}
		if result, _, err := transformer.Apply(dataset.New()); err == nil || result != nil {
			t.Fatalf("SetStrings(%s) result=%v err=%v, want error", target, result, err)
		}
	}
}

func TestNoOpSetAndMissingRemoveProduceNoChanges(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	transformer, err := NewTransformer(
		mustTransformRule(SetStrings(tag.PatientID, nil, "123")),
		mustTransformRule(Remove(tag.PatientName)),
	)
	if err != nil {
		t.Fatal(err)
	}
	_, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if len(changes) != 0 {
		t.Fatalf("changes = %#v, want none", changes)
	}
}

func TestMapValueMatchesCompleteCanonicalText(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.ImageType, vr.CS, []string{testRuleOriginal, "PRIMARY"}))
	transformer, err := NewTransformer(mustTransformRule(MapValue(tag.ImageType, "ORIGINAL\\PRIMARY", "DERIVED\\SECONDARY")))
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	values, _ := result.GetStrings(tag.ImageType)
	if !reflect.DeepEqual(values, []string{"DERIVED", "SECONDARY"}) || len(changes) != 1 || changes[0].Kind != ChangeMap {
		t.Fatalf("values=%v changes=%#v", values, changes)
	}
}

func TestCopyValueConvertsIntoDestinationVR(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewIntegerString(tag.NumberOfFrames, []string{"1024"}))
	requireRuleAdd(t, source, element.NewUnsignedShort(tag.Rows, []uint16{1}))
	transformer, err := NewTransformer(mustTransformRule(CopyValue(tag.NumberOfFrames, tag.Rows, vr.SS)))
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := result.GetUInt16s(tag.Rows)
	if err != nil || !reflect.DeepEqual(rows, []uint16{1024}) {
		t.Fatalf("Rows = %v, err = %v", rows, err)
	}
	rowsElem, _ := result.Get(tag.Rows)
	if rowsElem.ValueRepresentation() != vr.US || len(changes) != 1 || changes[0].Kind != ChangeCopy {
		t.Fatalf("Rows VR = %v, changes = %#v", rowsElem.ValueRepresentation(), changes)
	}
}

func TestCopyValueSupportsExplicitPrivateDestinationAndMissingSourceNoOp(t *testing.T) {
	private := tag.New(0x0011, 0x1010)
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	transformer, err := NewTransformer(
		mustTransformRule(CopyValue(tag.PatientID, private, vr.LO)),
		mustTransformRule(CopyValue(tag.PatientName, tag.StudyDescription, nil)),
	)
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	value, _ := result.GetString(private)
	if value != "123" || len(changes) != 1 {
		t.Fatalf("private value = %q, changes = %#v", value, changes)
	}
}

func TestCopyValueConversionFailureIsTransactional(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"not-a-number"}))
	requireRuleAdd(t, source, element.NewUnsignedShort(tag.Rows, []uint16{512}))
	transformer, err := NewTransformer(mustTransformRule(CopyValue(tag.PatientID, tag.Rows, nil)))
	if err != nil {
		t.Fatal(err)
	}
	result, _, err := transformer.Apply(source)
	if err == nil || result != nil {
		t.Fatalf("result=%v err=%v, want failure", result, err)
	}
	rows, getErr := source.GetUInt16s(tag.Rows)
	if getErr != nil || !reflect.DeepEqual(rows, []uint16{512}) {
		t.Fatalf("source Rows = %v, err = %v", rows, getErr)
	}
}

func TestCoreTransformConstructorsRejectNilInputs(t *testing.T) {
	if _, err := Remove(nil); err == nil {
		t.Fatal("Remove(nil) succeeded")
	}
	if _, err := RemoveMasked(nil); err == nil {
		t.Fatal("RemoveMasked(nil) succeeded")
	}
	if _, err := SetElement(nil); err == nil {
		t.Fatal("SetElement(nil) succeeded")
	}
	if _, err := CopyValue(nil, tag.PatientID, nil); err == nil {
		t.Fatal("CopyValue accepted nil source")
	}
	if _, err := MapValue(nil, "a", "b"); err == nil {
		t.Fatal("MapValue accepted nil tag")
	}
}

func mustTransformRule(rule TransformRule, err error) TransformRule {
	if err != nil {
		panic(err)
	}
	return rule
}
