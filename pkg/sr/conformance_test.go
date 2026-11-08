// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package sr

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/anonymizer"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestStructuredReportCreateAndSaveRoundTrip(t *testing.T) {
	finding, err := NewContentItemText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		testFindingText,
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), finding)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	if classUID, _ := report.Dataset().GetString(tag.SOPClassUID); classUID != uid.ComprehensiveSRStorage.UID() {
		t.Fatalf("SOPClassUID = %q, want Comprehensive SR", classUID)
	}
	if instanceUID, _ := report.Dataset().GetString(tag.SOPInstanceUID); !uid.IsValid(instanceUID) {
		t.Fatalf("SOPInstanceUID = %q, want a generated UID", instanceUID)
	}
	if !report.Dataset().Contains(tag.PerformedProcedureCodeSequence) {
		t.Fatal("Performed Procedure Code Sequence is missing")
	}
	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	for _, item := range []element.Element{
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{uid.GenerateDerivedFromUUID().UID()}),
		element.NewString(tag.SeriesInstanceUID, vr.UI, []string{uid.GenerateDerivedFromUUID().UID()}),
		element.NewString(tag.Modality, vr.CS, []string{"SR"}),
	} {
		if err := report.Dataset().AddOrUpdate(item); err != nil {
			t.Fatalf("add %s: %v", item.Tag(), err)
		}
	}

	path := filepath.Join(t.TempDir(), "report.dcm")
	if err := report.Save(path); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	opened, err := Open(path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	if err := opened.Validate(); err != nil {
		t.Fatalf("Open().Validate() error = %v", err)
	}
	children, err := opened.Children()
	if err != nil {
		t.Fatalf("Children() error = %v", err)
	}
	if len(children) != 1 {
		t.Fatalf("children = %d, want 1", len(children))
	}
	if text, err := children[0].GetText(); err != nil || text != testFindingText {
		t.Fatalf("finding text = %q, %v", text, err)
	}
}

func TestStructuredReportValidateRequiresSOPCommonAndType2Sequence(t *testing.T) {
	newReport := func(t *testing.T) *StructuredReport {
		t.Helper()
		report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"))
		if err != nil {
			t.Fatalf("NewStructuredReport() error = %v", err)
		}
		return report
	}

	tests := []struct {
		name   string
		remove *tag.Tag
		want   string
	}{
		{name: "SOP Class UID", remove: tag.SOPClassUID, want: "SOP Common Module requires"},
		{name: "SOP Instance UID", remove: tag.SOPInstanceUID, want: "SOP Common Module requires"},
		{name: "Performed Procedure Code Sequence", remove: tag.PerformedProcedureCodeSequence, want: "Performed Procedure Code Sequence"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := newReport(t)
			report.Dataset().Remove(tt.remove)
			err := report.Validate()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func TestAnonymizedStructuredReportStillValidates(t *testing.T) {
	finding, err := NewContentItemText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		"Patient Jane Doe was examined",
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	person, err := NewContentItemPersonName(
		NewCodeItem("113870", "DCM", "Person Name"),
		RelationshipContains,
		testObserverName,
	)
	if err != nil {
		t.Fatalf("NewContentItemPersonName() error = %v", err)
	}
	uidReference, err := NewContentItemUIDReference(
		NewCodeItem("110180", "DCM", "Device Observer UID"),
		RelationshipContains,
		"1.2.826.0.1.3680043.10.999.1",
	)
	if err != nil {
		t.Fatalf("NewContentItemUIDReference() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), finding, person, uidReference)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}

	// Make the report VERIFIED so that the Verifying Observer attributes are
	// required by validation.
	observer := dataset.New()
	for _, item := range []element.Element{
		element.NewPersonName(tag.VerifyingObserverName, []string{testObserverName}),
		element.NewString(tag.VerifyingOrganization, vr.LO, []string{"General Hospital"}),
		element.NewDateTimeFromTime(tag.VerificationDateTime, []time.Time{time.Now()}),
	} {
		if err := observer.Add(item); err != nil {
			t.Fatalf("add %s: %v", item.Tag(), err)
		}
	}
	for _, item := range []element.Element{
		element.NewString(tag.CompletionFlag, vr.CS, []string{"COMPLETE"}),
		element.NewString(tag.VerificationFlag, vr.CS, []string{"VERIFIED"}),
		dataset.NewSequenceWithItems(tag.VerifyingObserverSequence, []*dataset.Dataset{observer}),
	} {
		if err := report.Dataset().AddOrUpdate(item); err != nil {
			t.Fatalf("add %s: %v", item.Tag(), err)
		}
	}
	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() before anonymization error = %v", err)
	}

	if err := anonymizer.NewAnonymizer(anonymizer.NewSecurityProfile(anonymizer.BasicProfile)).AnonymizeInPlace(report.Dataset()); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}

	if flag, ok := report.Dataset().GetString(tag.VerificationFlag); !ok || flag != "VERIFIED" {
		t.Errorf("Verification Flag = %q, %v, want VERIFIED and present", flag, ok)
	}
	if !report.Dataset().Contains(tag.ContentSequence) {
		t.Fatal("Content Sequence was removed by the Basic profile")
	}
	children, err := report.Children()
	if err != nil {
		t.Fatalf("Children() error = %v", err)
	}
	if len(children) != 3 {
		t.Fatalf("children = %d, want 3", len(children))
	}
	if text, err := children[0].GetText(); err != nil || text == "Patient Jane Doe was examined" || text == "" {
		t.Fatalf("text value after anonymization = %q, %v, want a non-empty dummy", text, err)
	}
	if name, err := children[1].GetPersonName(); err != nil || name == testObserverName || name == "" {
		t.Fatalf("person name after anonymization = %q, %v, want a non-empty dummy", name, err)
	}
	if reference, err := children[2].GetUIDReference(); err != nil || reference == "1.2.826.0.1.3680043.10.999.1" {
		t.Fatalf("UID reference after anonymization = %q, %v, want a replaced UID", reference, err)
	}

	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() after anonymization error = %v", err)
	}
}

func TestCleanStructuredContentOptionCleansRatherThanDeletes(t *testing.T) {
	finding, err := NewContentItemText(
		NewCodeItem("999001", "99TEST", "Clinical Note"),
		RelationshipContains,
		testFindingText,
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	person, err := NewContentItemPersonName(
		NewCodeItem("113870", "DCM", "Person Name"),
		RelationshipContains,
		testObserverName,
	)
	if err != nil {
		t.Fatalf("NewContentItemPersonName() error = %v", err)
	}
	subjectSex, err := NewContentItemCode(
		NewCodeItem("121032", "DCM", "Subject Sex"),
		RelationshipContains,
		NewCodeItem("F", "DCM", "Female"),
	)
	if err != nil {
		t.Fatalf("NewContentItemCode() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), finding, person, subjectSex)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}

	profile := anonymizer.NewSecurityProfile(anonymizer.BasicProfile | anonymizer.CleanStructdCont)
	if err := anonymizer.NewAnonymizer(profile).AnonymizeInPlace(report.Dataset()); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	if !report.Dataset().Contains(tag.ContentSequence) {
		t.Fatal("Content Sequence was deleted instead of cleaned")
	}
	children, err := report.Children()
	if err != nil {
		t.Fatalf("Children() error = %v", err)
	}
	// The Subject Sex item is listed in Table E.3.4-1 with action X; the
	// unlisted finding text and the listed Person Name item remain (cleaned).
	if len(children) != 2 {
		t.Fatalf("children = %d, want 2", len(children))
	}
	if text, _ := children[0].GetText(); text != testFindingText {
		t.Fatalf("finding text = %q, want the non-identifying text retained", text)
	}
	if name, _ := children[1].GetPersonName(); name == testObserverName || name == "" {
		t.Fatalf("person name = %q, want a non-empty dummy", name)
	}
	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestNumericContentItemWithQualifier(t *testing.T) {
	code := NewCodeItem("121033", "DCM", "Subject Age")
	item, err := NewContentItemNumeric(code, RelationshipContains, nil)
	if err != nil {
		t.Fatalf("NewContentItemNumeric() error = %v", err)
	}
	if err := item.Validate(); err == nil {
		t.Fatal("Validate() accepted a NUM item without a measured value or qualifier")
	}

	if err := item.Dataset().AddOrUpdate(dataset.NewSequence(tag.MeasuredValueSequence)); err != nil {
		t.Fatalf("add empty measured value sequence: %v", err)
	}
	if err := item.Validate(); err == nil || !strings.Contains(err.Error(), "qualifier") {
		t.Fatalf("Validate() error = %v, want a qualifier requirement", err)
	}

	if err := item.SetNumericQualifier(NewCodeItem("121042", "DCM", "Measurement not available")); err != nil {
		t.Fatalf("SetNumericQualifier() error = %v", err)
	}
	if err := item.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}

	tooMany := dataset.NewSequenceWithItems(tag.MeasuredValueSequence, []*dataset.Dataset{
		NewMeasuredValue(1, NewCodeItem("mm", "UCUM", "millimeter")).Dataset(),
		NewMeasuredValue(2, NewCodeItem("mm", "UCUM", "millimeter")).Dataset(),
	})
	if err := item.Dataset().AddOrUpdate(tooMany); err != nil {
		t.Fatalf("add two measured values: %v", err)
	}
	if err := item.Validate(); err == nil || !strings.Contains(err.Error(), "zero or one") {
		t.Fatalf("Validate() error = %v, want a zero-or-one error", err)
	}
}

func TestMeasuredValueRejectsMultipleNumericValues(t *testing.T) {
	value := NewMeasuredValue(1.5, NewCodeItem("mm", "UCUM", "millimeter"))
	if err := value.Dataset().AddOrUpdate(element.NewString(tag.NumericValue, vr.DS, []string{"1", "2"})); err != nil {
		t.Fatalf("add numeric value: %v", err)
	}
	err := value.Validate()
	if err == nil || !strings.Contains(err.Error(), "want 1") {
		t.Fatalf("Validate() error = %v, want a single-value error", err)
	}
}

func TestTableContentItemValidates(t *testing.T) {
	table := NewTableValue(2, 3)
	item, err := NewContentItemTable(NewCodeItem("121071", "DCM", "Finding"), RelationshipContains, table)
	if err != nil {
		t.Fatalf("NewContentItemTable() error = %v", err)
	}
	if err := item.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
	got, err := item.GetTable()
	if err != nil {
		t.Fatalf("GetTable() error = %v", err)
	}
	if rows, _ := got.Rows(); rows != 2 {
		t.Fatalf("Rows() = %d, want 2", rows)
	}
	if columns, _ := got.Columns(); columns != 3 {
		t.Fatalf("Columns() = %d, want 3", columns)
	}

	missingCounts := NewContentItemFromDataset(item.Dataset().DeepClone())
	_ = missingCounts.Dataset().AddOrUpdate(dataset.NewSequenceWithItems(tag.TabulatedValuesSequence, []*dataset.Dataset{dataset.New()}))
	if err := missingCounts.Validate(); err == nil {
		t.Fatal("Validate() accepted a table without row and column counts")
	}
}

func TestNestedUnheadedContainerIsSupported(t *testing.T) {
	leaf, err := NewContentItemText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		"normal",
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	unheaded, err := NewContentItemContainer(nil, RelationshipContains, ContinuitySeparate, leaf)
	if err != nil {
		t.Fatalf("NewContentItemContainer(nil) error = %v", err)
	}
	if unheaded.Dataset().Contains(tag.ConceptNameCodeSequence) {
		t.Fatal("unheaded container unexpectedly has a Concept Name Code Sequence")
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), unheaded)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	if err := report.Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}

func TestUnicodeStructuredReportDeclaresUTF8(t *testing.T) {
	item, err := NewContentItemText(
		NewCodeItem("121071", "DCM", "Finding"),
		RelationshipContains,
		"患者さんは異常なし",
	)
	if err != nil {
		t.Fatalf("NewContentItemText() error = %v", err)
	}
	report, err := NewStructuredReport(NewCodeItem("113704", "DCM", "SR Document"), item)
	if err != nil {
		t.Fatalf("NewStructuredReport() error = %v", err)
	}
	charset := report.Dataset().SpecificCharacterSet()
	if len(charset) != 1 || charset[0] != "ISO_IR 192" {
		t.Fatalf("SpecificCharacterSet = %v, want [ISO_IR 192]", charset)
	}

	path := filepath.Join(t.TempDir(), "unicode.dcm")
	if err := report.Save(path); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	opened, err := Open(path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	if got := opened.Dataset().SpecificCharacterSet(); len(got) != 1 || got[0] != "ISO_IR 192" {
		t.Fatalf("round-tripped SpecificCharacterSet = %v, want [ISO_IR 192]", got)
	}
	children, err := opened.Children()
	if err != nil || len(children) != 1 {
		t.Fatalf("Children() = %d, %v", len(children), err)
	}
	if text, err := children[0].GetText(); err != nil || text != "患者さんは異常なし" {
		t.Fatalf("text = %q, %v", text, err)
	}
}
