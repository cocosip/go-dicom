// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"golang.org/x/text/encoding/unicode"
)

func TestTextTransformsPreserveMultiplicityAndOrder(t *testing.T) {
	target := tag.New(0x0011, 0x1010)
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(target, vr.LO, []string{" ab ", "cd"}))
	transformer, err := NewTransformer(
		mustTransformRule(TrimSpace(target, TrimBoth)),
		mustTransformRule(ToUpper(target)),
		mustTransformRule(PrefixValue(target, "[")),
		mustTransformRule(AppendValue(target, "]")),
	)
	if err != nil {
		t.Fatal(err)
	}

	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	values, _ := result.GetStrings(target)
	if !reflect.DeepEqual(values, []string{"[AB]", "[CD]"}) {
		t.Fatalf("values = %v", values)
	}
	if len(changes) != 4 {
		t.Fatalf("changes = %d, want 4", len(changes))
	}
	for _, change := range changes {
		if change.Kind != ChangeEdit {
			t.Fatalf("change kind = %s, want edit", change.Kind)
		}
	}
}

func TestReplaceRegexAppliesToEachValueWithoutConsumingSeparator(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.ImageType, vr.CS, []string{"AB", "CD"}))
	transformer, err := NewTransformer(mustTransformRule(ReplaceRegex(tag.ImageType, `AB\\CD`, "joined")))
	if err != nil {
		t.Fatal(err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	values, _ := result.GetStrings(tag.ImageType)
	if !reflect.DeepEqual(values, []string{"AB", "CD"}) || len(changes) != 0 {
		t.Fatalf("values = %v, changes = %v", values, changes)
	}

	transformer, err = NewTransformer(mustTransformRule(ReplaceRegex(tag.ImageType, `^[AC]`, "X")))
	if err != nil {
		t.Fatal(err)
	}
	result, _, err = transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	values, _ = result.GetStrings(tag.ImageType)
	if !reflect.DeepEqual(values, []string{"XB", "XD"}) {
		t.Fatalf("values = %v", values)
	}
}

func TestTrimRulesHandleEveryPosition(t *testing.T) {
	tests := []struct {
		name  string
		rule  TransformRule
		input string
		want  string
	}{
		{name: "space start", rule: mustTransformRule(TrimSpace(tag.PatientID, TrimStart)), input: "  id", want: "id"},
		{name: "characters end", rule: mustTransformRule(TrimCharacters(tag.PatientID, TrimEnd, "-_")), input: "id_-", want: "id"},
		{name: "characters both", rule: mustTransformRule(TrimCharacters(tag.PatientID, TrimBoth, "-_")), input: "-_id_-", want: "id"},
		{name: "repeated start", rule: mustTransformRule(TrimRepeatedString(tag.PatientID, TrimStart, "ab")), input: "ababidab", want: "idab"},
		{name: "repeated end", rule: mustTransformRule(TrimRepeatedString(tag.PatientID, TrimEnd, "ab")), input: "abidabab", want: "abid"},
		{name: "repeated both", rule: mustTransformRule(TrimRepeatedString(tag.PatientID, TrimBoth, "ab")), input: "ababidabab", want: "id"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			source := dataset.New()
			requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{tc.input}))
			transformer, err := NewTransformer(tc.rule)
			if err != nil {
				t.Fatal(err)
			}
			result, _, err := transformer.Apply(source)
			if err != nil {
				t.Fatal(err)
			}
			if got, _ := result.GetString(tag.PatientID); got != tc.want {
				t.Fatalf("value = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPadAndTruncateUseUnicodeCodePoints(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewStringWithEncoding(tag.PatientID, vr.LO, []string{"汉"}, unicode.UTF8))
	transformer, err := NewTransformer(
		mustTransformRule(PadValue(tag.PatientID, PadLeft, 3, '_')),
		mustTransformRule(AppendValue(tag.PatientID, "字A")),
		mustTransformRule(TruncateValues(tag.PatientID, 4)),
	)
	if err != nil {
		t.Fatal(err)
	}
	result, _, err := transformer.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := result.GetString(tag.PatientID); got != "__汉字" {
		t.Fatalf("value = %q, want %q", got, "__汉字")
	}

	right, err := NewTransformer(mustTransformRule(PadValue(tag.PatientID, PadRight, 3, '_')))
	if err != nil {
		t.Fatal(err)
	}
	result, _, err = right.Apply(source)
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := result.GetString(tag.PatientID); got != "汉__" {
		t.Fatalf("right padded value = %q", got)
	}
}

func TestTextTransformsMissingTagAreNoOp(t *testing.T) {
	transformer, err := NewTransformer(mustTransformRule(ToLower(tag.PatientID)))
	if err != nil {
		t.Fatal(err)
	}
	_, changes, err := transformer.Apply(dataset.New())
	if err != nil || len(changes) != 0 {
		t.Fatalf("changes = %v, err = %v", changes, err)
	}
}

func TestTextTransformConstructorsRejectInvalidConfiguration(t *testing.T) {
	if _, err := ReplaceRegex(tag.PatientID, `[`, "x"); err == nil {
		t.Fatal("ReplaceRegex accepted invalid regular expression")
	}
	if _, err := ToUpper(nil); err == nil {
		t.Fatal("ToUpper accepted nil tag")
	}
	if _, err := TrimSpace(tag.PatientID, TrimPosition(99)); err == nil {
		t.Fatal("TrimSpace accepted invalid position")
	}
	if _, err := TrimCharacters(tag.PatientID, TrimBoth, ""); err == nil {
		t.Fatal("TrimCharacters accepted empty cutset")
	}
	if _, err := TrimRepeatedString(tag.PatientID, TrimBoth, ""); err == nil {
		t.Fatal("TrimRepeatedString accepted empty value")
	}
	if _, err := PadValue(tag.PatientID, PadDirection(99), 2, '_'); err == nil {
		t.Fatal("PadValue accepted invalid direction")
	}
	if _, err := PadValue(tag.PatientID, PadLeft, -1, '_'); err == nil {
		t.Fatal("PadValue accepted negative length")
	}
	if _, err := TruncateValues(tag.PatientID, -1); err == nil {
		t.Fatal("TruncateValues accepted negative length")
	}
}

func TestSplitFormatReordersFullCanonicalValueAndPreservesEmptyFields(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientName, vr.PN, []string{testRuleSplitFormatPatientName}))
	rule, err := SplitFormat(tag.PatientName, "^", "{1} {0} ({3})")
	if err != nil {
		t.Fatalf("SplitFormat() error = %v", err)
	}
	transformer, err := NewTransformer(rule)
	if err != nil {
		t.Fatalf("NewTransformer() error = %v", err)
	}
	result, changes, err := transformer.Apply(source)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	values, err := element.CanonicalStrings(result.GetOrNil(tag.PatientName))
	if err != nil || !reflect.DeepEqual(values, []string{"Jane Doe (Dr)"}) {
		t.Fatalf("PatientName canonical values = %v, error = %v", values, err)
	}
	if len(changes) != 1 || changes[0].Kind != ChangeEdit {
		t.Fatalf("changes = %#v, want one ChangeEdit", changes)
	}
	if got, _ := source.GetString(tag.PatientName); got != "Doe^Jane^^Dr" {
		t.Fatalf("source PatientName mutated to %q", got)
	}
}

func TestSplitFormatSupportsUnicodeSeparatorsEscapedBracesAndMultipleValues(t *testing.T) {
	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.OtherPatientIDsRETIRED, vr.LO, []string{"A", "B|C"}))
	rule, err := SplitFormat(tag.OtherPatientIDsRETIRED, "\\|", "{{{2}}}:{0}:{1}")
	if err != nil {
		t.Fatalf("SplitFormat() error = %v", err)
	}
	transformer, err := NewTransformer(rule)
	if err != nil {
		t.Fatalf("NewTransformer() error = %v", err)
	}
	result, _, err := transformer.Apply(source)
	if err != nil {
		t.Fatalf("Apply() error = %v", err)
	}
	if got, _ := result.GetString(tag.OtherPatientIDsRETIRED); got != "{C}:A:B" {
		t.Fatalf("OtherPatientIDsRETIRED = %q, want %q", got, "{C}:A:B")
	}
}

func TestSplitFormatRejectsInvalidGrammarAndRollsBackOutOfRangeIndex(t *testing.T) {
	for _, format := range []string{"{", "}", "{-1}", "{x}", "{0"} {
		if _, err := SplitFormat(tag.PatientID, "^", format); err == nil {
			t.Fatalf("SplitFormat accepted invalid format %q", format)
		}
	}
	if _, err := SplitFormat(nil, "^", "{0}"); err == nil {
		t.Fatal("SplitFormat accepted nil tag")
	}
	if _, err := SplitFormat(tag.PatientID, "", "{0}"); err == nil {
		t.Fatal("SplitFormat accepted empty separators")
	}

	source := dataset.New()
	requireRuleAdd(t, source, element.NewString(tag.PatientID, vr.LO, []string{"only"}))
	rule := mustTransformRule(SplitFormat(tag.PatientID, "^", "{1}"))
	transformer, err := NewTransformer(rule)
	if err != nil {
		t.Fatalf("NewTransformer() error = %v", err)
	}
	_, changes, err := transformer.Apply(source)
	if err == nil {
		t.Fatal("Apply() error = nil, want out-of-range format index")
	}
	if len(changes) != 0 {
		t.Fatalf("changes = %#v, want no partial change", changes)
	}
	if got, _ := source.GetString(tag.PatientID); got != "only" {
		t.Fatalf("source PatientID mutated to %q", got)
	}
}

func TestSplitFormatRecordsNestedSequencePathAndPreservesTextEncoding(t *testing.T) {
	item := dataset.New()
	requireRuleAdd(t, item, element.NewStringWithEncoding(tag.PatientName, vr.PN, []string{"张^三"}, unicode.UTF8))
	rule := mustTransformRule(SplitFormat(tag.PatientName, "^", "{1}{0}"))
	index := 0
	path := dataset.Path{{Tag: tag.SourceImageSequence}, {ItemIndex: &index}}
	changes := ChangeSet{}

	if err := rule.apply(item, path, &changes); err != nil {
		t.Fatalf("SplitFormat.apply() error = %v", err)
	}
	values, err := element.CanonicalStrings(item.GetOrNil(tag.PatientName))
	if err != nil || !reflect.DeepEqual(values, []string{"三张"}) {
		t.Fatalf("PatientName canonical values = %v, error = %v", values, err)
	}
	if len(changes) != 1 || len(changes[0].Path) != 3 ||
		!changes[0].Path[0].Tag.Equals(tag.SourceImageSequence) || changes[0].Path[1].ItemIndex == nil ||
		*changes[0].Path[1].ItemIndex != 0 || !changes[0].Path[2].Tag.Equals(tag.PatientName) {
		t.Fatalf("changes = %#v, want one nested PatientName edit", changes)
	}
	encoded := item.GetOrNil(tag.PatientName)
	if got := encoded.Buffer().Data(); !reflect.DeepEqual(got, []byte("三张")) {
		t.Fatalf("encoded PN = %x, want UTF-8 bytes", got)
	}
}
