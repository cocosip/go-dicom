// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"errors"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestMissingAndEmptyHaveDistinctMatchSemantics(t *testing.T) {
	missing := dataset.New()
	empty := dataset.New()
	requireRuleAdd(t, empty, element.NewString(tag.PatientName, vr.PN, nil))

	assertRuleMatch(t, mustMatchRule(Exists(tag.PatientName)), missing, false)
	assertRuleMatch(t, mustMatchRule(Exists(tag.PatientName)), empty, true)
	assertRuleMatch(t, mustMatchRule(Empty(tag.PatientName)), missing, true)
	assertRuleMatch(t, mustMatchRule(Empty(tag.PatientName)), empty, true)
	assertRuleMatch(t, mustMatchRule(Equal(tag.PatientName, "")), missing, false)
	assertRuleMatch(t, mustMatchRule(Equal(tag.PatientName, "")), empty, true)
}

func TestContentRulesUseCompleteMultiValueText(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewString(tag.ImageType, vr.CS, []string{"ORIGINAL", "PRIMARY"}))

	assertRuleMatch(t, mustMatchRule(Equal(tag.ImageType, "ORIGINAL\\PRIMARY")), ds, true)
	assertRuleMatch(t, mustMatchRule(Equal(tag.ImageType, "ORIGINAL")), ds, false)
	assertRuleMatch(t, mustMatchRule(StartsWith(tag.ImageType, "ORIGINAL\\")), ds, true)
	assertRuleMatch(t, mustMatchRule(EndsWith(tag.ImageType, "\\PRIMARY")), ds, true)
	assertRuleMatch(t, mustMatchRule(Contains(tag.ImageType, "NAL\\PRI")), ds, true)
	assertRuleMatch(t, mustMatchRule(OneOf(tag.ImageType, "LOCALIZER", "ORIGINAL\\PRIMARY")), ds, true)
}

func TestContentRulesUseCanonicalNumericAndAttributeTagText(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewUnsignedShort(tag.Rows, []uint16{512}))
	requireRuleAdd(t, ds, element.NewAttributeTag(tag.DimensionIndexPointer, []*tag.Tag{tag.Rows}))

	assertRuleMatch(t, mustMatchRule(Equal(tag.Rows, "512")), ds, true)
	assertRuleMatch(t, mustMatchRule(Equal(tag.DimensionIndexPointer, "(0028,0010)")), ds, true)
}

func TestContentRulesRejectUnsupportedValueTypes(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, dataset.NewSequence(tag.ReferencedImageSequence))

	_, err := mustMatchRule(Equal(tag.ReferencedImageSequence, "")).Match(ds)
	var unsupported *element.UnsupportedValueError
	if !errors.As(err, &unsupported) {
		t.Fatalf("error = %v, want UnsupportedValueError", err)
	}
}

func TestAllAnyNotAndBoolComposition(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"}))

	all := mustMatchRule(All(
		mustMatchRule(Exists(tag.PatientName)),
		mustMatchRule(Not(mustMatchRule(Empty(tag.PatientName)))),
		mustMatchRule(Any(Bool(false), mustMatchRule(Equal(tag.PatientName, "Doe^Jane")))),
	))

	assertRuleMatch(t, all, ds, true)
	assertRuleMatch(t, mustMatchRule(All()), ds, true)
	assertRuleMatch(t, mustMatchRule(Any()), ds, false)
}

func TestAllAndAnyShortCircuitErrors(t *testing.T) {
	sentinel := errors.New("must not evaluate")
	ds := dataset.New()

	assertRuleMatch(t, mustMatchRule(All(Bool(false), errorMatchRule{err: sentinel})), ds, false)
	assertRuleMatch(t, mustMatchRule(Any(Bool(true), errorMatchRule{err: sentinel})), ds, true)

	_, err := mustMatchRule(All(Bool(true), errorMatchRule{err: sentinel})).Match(ds)
	if !errors.Is(err, sentinel) {
		t.Fatalf("All error = %v, want sentinel", err)
	}
}

func TestExactContentRulesAreCaseSensitive(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewString(tag.PatientID, vr.LO, []string{"Patient-ABC"}))

	for _, rule := range []MatchRule{
		mustMatchRule(Equal(tag.PatientID, "patient-abc")),
		mustMatchRule(StartsWith(tag.PatientID, "patient")),
		mustMatchRule(EndsWith(tag.PatientID, "abc")),
		mustMatchRule(Contains(tag.PatientID, "PATIENT")),
		mustMatchRule(OneOf(tag.PatientID, "patient-abc")),
	} {
		assertRuleMatch(t, rule, ds, false)
	}
}

func TestWildcardIsAnchoredAndCaseInsensitive(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewString(tag.PatientID, vr.LO, []string{"Patient-ABC"}))

	assertRuleMatch(t, mustMatchRule(Wildcard(tag.PatientID, "p?tient-*")), ds, true)
	assertRuleMatch(t, mustMatchRule(Wildcard(tag.PatientID, "patient")), ds, false)
	assertRuleMatch(t, mustMatchRule(Wildcard(tag.PatientID, "*ABC")), ds, true)
	assertRuleMatch(t, mustMatchRule(Wildcard(tag.PatientID, "Patient-???")), ds, true)
}

func TestRegexUsesRE2AndInlineFlags(t *testing.T) {
	ds := dataset.New()
	requireRuleAdd(t, ds, element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"}))

	assertRuleMatch(t, mustMatchRule(Regex(tag.PatientName, `^Doe\^J`)), ds, true)
	assertRuleMatch(t, mustMatchRule(Regex(tag.PatientName, `(?i)^doe`)), ds, true)
	assertRuleMatch(t, mustMatchRule(Regex(tag.PatientName, `^doe`)), ds, false)
	if _, err := Regex(tag.PatientName, `(a)\1`); err == nil {
		t.Fatal("Regex accepted a backreference unsupported by RE2")
	}
}

func TestMatchConstructorsRejectNilInputs(t *testing.T) {
	if _, err := Exists(nil); err == nil {
		t.Fatal("Exists(nil) succeeded")
	}
	if _, err := Not(nil); err == nil {
		t.Fatal("Not(nil) succeeded")
	}
	if _, err := All(Bool(true), nil); err == nil {
		t.Fatal("All accepted a nil rule")
	}
	if _, err := Equal(nil, "x"); err == nil {
		t.Fatal("Equal accepted a nil tag")
	}
}

func TestConcurrentMatchRuleReuseOnIndependentDatasets(t *testing.T) {
	rule := mustMatchRule(All(
		mustMatchRule(Exists(tag.PatientID)),
		mustMatchRule(Wildcard(tag.PatientID, "case-*")),
	))

	const workers = 16
	errorsByWorker := make(chan error, workers)
	var wait sync.WaitGroup
	for index := 0; index < workers; index++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			ds := dataset.New()
			if err := ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"CASE-123"})); err != nil {
				errorsByWorker <- err
				return
			}
			matched, err := rule.Match(ds)
			if err != nil {
				errorsByWorker <- err
				return
			}
			if !matched {
				errorsByWorker <- errors.New("rule did not match")
			}
		}()
	}
	wait.Wait()
	close(errorsByWorker)
	for err := range errorsByWorker {
		t.Error(err)
	}
}

type errorMatchRule struct{ err error }

func (rule errorMatchRule) Match(*dataset.Dataset) (bool, error) { return false, rule.err }

func mustMatchRule(rule MatchRule, err error) MatchRule {
	if err != nil {
		panic(err)
	}
	return rule
}

func assertRuleMatch(t *testing.T, rule MatchRule, ds *dataset.Dataset, want bool) {
	t.Helper()
	got, err := rule.Match(ds)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Match() = %t, want %t", got, want)
	}
}

func requireRuleAdd(t *testing.T, ds *dataset.Dataset, elem element.Element) {
	t.Helper()
	if err := ds.Add(elem); err != nil {
		t.Fatal(err)
	}
}
