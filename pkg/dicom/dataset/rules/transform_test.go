// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"errors"
	"reflect"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestApplyReturnsIndependentCloneOnNoRules(t *testing.T) {
	source := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	requireRuleAdd(t, source, element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"}))
	transformer, err := NewTransformer()
	if err != nil {
		t.Fatal(err)
	}

	got, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if got == source {
		t.Fatal("Apply returned the source Dataset")
	}
	if len(changes) != 0 {
		t.Fatalf("changes = %v, want empty", changes)
	}
	if got.InternalTransferSyntax() != transfer.ExplicitVRBigEndian {
		t.Fatal("Apply did not preserve transfer syntax")
	}
	if err := got.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{"Changed^Clone"})); err != nil {
		t.Fatal(err)
	}
	if value, _ := source.GetString(tag.PatientName); value != "Doe^Jane" {
		t.Fatalf("source changed through result clone: %q", value)
	}
}

func TestConditionalNonMatchReturnsUnchangedClone(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	transformer, err := NewConditionalTransformer(Bool(false), testAssignRule{
		elem: element.NewString(tag.PatientID, vr.LO, []string{"456"}),
	})
	if err != nil {
		t.Fatal(err)
	}

	got, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if got == source || len(changes) != 0 {
		t.Fatalf("result = %p source = %p changes = %v", got, source, changes)
	}
	if value, _ := got.GetString(tag.PatientID); value != "123" {
		t.Fatalf("PatientID = %q, want unchanged", value)
	}
}

func TestApplyFailureNeverMutatesSourceAndReturnsPriorChanges(t *testing.T) {
	sentinel := errors.New("rule failed")
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientName, vr.PN, []string{"Old^Name"}))
	transformer, err := NewTransformer(
		testAssignRule{elem: element.NewString(tag.PatientName, vr.PN, []string{"New^Name"})},
		testFailRule{err: sentinel},
	)
	if err != nil {
		t.Fatal(err)
	}

	got, changes, err := transformer.Apply(source)
	if got != nil {
		t.Fatal("Apply returned a Dataset after rule failure")
	}
	if !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want sentinel", err)
	}
	if len(changes) != 1 || changes[0].Kind != ChangeAssign {
		t.Fatalf("changes = %#v", changes)
	}
	if value, _ := source.GetString(tag.PatientName); value != "Old^Name" {
		t.Fatalf("source changed after failure: %q", value)
	}
}

func TestTransformErrorCarriesNestedRulePath(t *testing.T) {
	sentinel := errors.New("nested failure")
	inner, err := NewTransformer(testFailRule{
		err: &TransformError{
			Stage: StageRule,
			Path:  dataset.Path{{Tag: tag.PatientID}},
			Cause: sentinel,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	outer, err := NewTransformer(testNoopRule{}, inner)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = outer.Apply(dataset.New())
	var transformErr *TransformError
	if !errors.As(err, &transformErr) || !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want nested TransformError", err)
	}
	if transformErr.Stage != StageRule || !reflect.DeepEqual(transformErr.RulePath, []int{1, 0}) {
		t.Fatalf("TransformError = %#v", transformErr)
	}
	if got := dataset.FormatPath(transformErr.Path); got != "(0010,0020)" {
		t.Fatalf("error path = %s", got)
	}
}

func TestConditionalErrorUsesConditionStage(t *testing.T) {
	sentinel := errors.New("condition failed")
	transformer, err := NewConditionalTransformer(errorMatchRule{err: sentinel})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = transformer.Apply(dataset.New())

	var transformErr *TransformError
	if !errors.As(err, &transformErr) || !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want TransformError", err)
	}
	if transformErr.Stage != StageCondition || len(transformErr.RulePath) != 0 {
		t.Fatalf("TransformError = %#v", transformErr)
	}
}

func TestChangeSetDoesNotAliasResultDataset(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientName, vr.PN, []string{"Old^Name"}))
	transformer, err := NewTransformer(testAssignRule{
		elem: element.NewString(tag.PatientName, vr.PN, []string{"New^Name"}),
	})
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := result.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{"Later^Name"})); err != nil {
		t.Fatal(err)
	}

	before, err := element.CanonicalStrings(changes[0].Before)
	if err != nil {
		t.Fatal(err)
	}
	after, err := element.CanonicalStrings(changes[0].After)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(before, []string{"Old^Name"}) || !reflect.DeepEqual(after, []string{"New^Name"}) {
		t.Fatalf("change snapshots before=%v after=%v", before, after)
	}
}

func TestApplyInPlaceCommitsOnlySuccessfulCompleteResult(t *testing.T) {
	target := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	target.SetAutoValidate(false)
	requireRuleAdd(t, target, element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	targetPointer := target
	transformer, err := NewTransformer(testAssignRule{
		elem: element.NewString(tag.PatientID, vr.LO, []string{"456"}),
	})
	if err != nil {
		t.Fatal(err)
	}

	changes, err := transformer.ApplyInPlace(target)
	if err != nil {
		t.Fatal(err)
	}
	if target != targetPointer || target.AutoValidate() {
		t.Fatal("ApplyInPlace changed pointer or validation mode")
	}
	if target.InternalTransferSyntax() != transfer.ExplicitVRLittleEndian {
		t.Fatal("ApplyInPlace lost transfer syntax")
	}
	if value, _ := target.GetString(tag.PatientID); value != "456" || len(changes) != 1 {
		t.Fatalf("PatientID = %q, changes = %v", value, changes)
	}
}

func TestApplyInPlaceFailureLeavesTargetUnchanged(t *testing.T) {
	sentinel := errors.New("failure")
	target := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	requireRuleAdd(t, target, element.NewString(tag.PatientID, vr.LO, []string{"123"}))
	transformer, err := NewTransformer(
		testAssignRule{elem: element.NewString(tag.PatientID, vr.LO, []string{"456"})},
		testFailRule{err: sentinel},
	)
	if err != nil {
		t.Fatal(err)
	}

	changes, err := transformer.ApplyInPlace(target)
	if !errors.Is(err, sentinel) || len(changes) != 1 {
		t.Fatalf("changes = %v, error = %v", changes, err)
	}
	if value, _ := target.GetString(tag.PatientID); value != "123" {
		t.Fatalf("target changed after failure: %q", value)
	}
	if target.InternalTransferSyntax() != transfer.ExplicitVRBigEndian {
		t.Fatal("target transfer syntax changed after failure")
	}
}

func TestTransformerConstructorsRejectNilRulesAndCondition(t *testing.T) {
	if _, err := NewTransformer(nil); err == nil {
		t.Fatal("NewTransformer accepted nil rule")
	}
	if _, err := NewConditionalTransformer(nil); err == nil {
		t.Fatal("NewConditionalTransformer accepted nil condition")
	}
}

func TestNestedTransformerPropagatesOrderedChanges(t *testing.T) {
	inner, err := NewTransformer(mustTransformRule(AppendValue(tag.PatientID, "]")))
	if err != nil {
		t.Fatal(err)
	}
	outer, err := NewTransformer(
		mustTransformRule(PrefixValue(tag.PatientID, "[")),
		inner,
	)
	if err != nil {
		t.Fatal(err)
	}
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"123"}))

	result, changes, err := outer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if value, _ := result.GetString(tag.PatientID); value != "[123]" {
		t.Fatalf("PatientID = %q", value)
	}
	if len(changes) != 2 || changes[0].Kind != ChangeEdit || changes[1].Kind != ChangeEdit {
		t.Fatalf("changes = %#v", changes)
	}
	firstAfter, _ := element.CanonicalStrings(changes[0].After)
	secondBefore, _ := element.CanonicalStrings(changes[1].Before)
	if !reflect.DeepEqual(firstAfter, []string{"[123"}) || !reflect.DeepEqual(secondBefore, firstAfter) {
		t.Fatalf("nested change sequence first=%v second-before=%v", firstAfter, secondBefore)
	}
}

func TestInvalidReadableDatasetTransformsWithoutValidatingUnrelatedElements(t *testing.T) {
	source := dataset.New()
	source.SetAutoValidate(false)
	requireRuleAdd(t, source, element.NewString(tag.Modality, vr.CS, []string{"invalid-lowercase"}))
	transformer, err := NewTransformer(mustTransformRule(SetStrings(tag.PatientName, vr.PN, "Doe^Jane")))
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if result.AutoValidate() || len(changes) != 1 {
		t.Fatalf("AutoValidate=%t changes=%v", result.AutoValidate(), changes)
	}
	if modality, _ := result.GetString(tag.Modality); modality != "invalid-lowercase" {
		t.Fatalf("unrelated invalid Modality changed: %q", modality)
	}
	patientNameElement, _ := result.Get(tag.PatientName)
	patientNames, valueErr := element.CanonicalStrings(patientNameElement)
	if valueErr != nil || !reflect.DeepEqual(patientNames, []string{"Doe^Jane"}) {
		t.Fatalf("PatientName = %v, err = %v", patientNames, valueErr)
	}
}

func TestConcurrentTransformerReuseOnIndependentDatasets(t *testing.T) {
	condition := mustMatchRule(Wildcard(tag.PatientID, "case-*"))
	transformer, err := NewConditionalTransformer(
		condition,
		mustTransformRule(ToUpper(tag.PatientID)),
		mustTransformRule(PrefixValue(tag.PatientID, "ID:")),
	)
	if err != nil {
		t.Fatal(err)
	}

	const workers = 16
	errorsByWorker := make(chan error, workers)
	var wait sync.WaitGroup
	for index := 0; index < workers; index++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			source := dataset.New()
			if err := source.Add(element.NewString(tag.PatientID, vr.LO, []string{"CASE-123"})); err != nil {
				errorsByWorker <- err
				return
			}
			result, changes, err := transformer.Apply(source)
			if err != nil {
				errorsByWorker <- err
				return
			}
			if value, _ := result.GetString(tag.PatientID); value != "ID:CASE-123" || len(changes) != 1 {
				errorsByWorker <- errors.New("unexpected concurrent transform result")
			}
		}()
	}
	wait.Wait()
	close(errorsByWorker)
	for err := range errorsByWorker {
		t.Error(err)
	}
}

type testAssignRule struct{ elem element.Element }

func (rule testAssignRule) apply(ds *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	before := ds.GetOrNil(rule.elem.Tag())
	assigned := element.DeepClone(rule.elem)
	if err := ds.AddOrUpdate(assigned); err != nil {
		return err
	}
	appendChange(changes, Change{
		Kind:   ChangeAssign,
		Path:   append(dataset.ClonePath(path), dataset.PathSegment{Tag: assigned.Tag()}),
		Tag:    assigned.Tag(),
		Before: before,
		After:  assigned,
	})
	return nil
}

type testFailRule struct{ err error }

func (rule testFailRule) apply(*dataset.Dataset, dataset.Path, *ChangeSet) error { return rule.err }

type testNoopRule struct{}

func (testNoopRule) apply(*dataset.Dataset, dataset.Path, *ChangeSet) error { return nil }
