// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	anonymizedString        = "ANONYMOUS"
	customPatientName       = "ANONYMOUS^PATIENT"
	customPatientID         = "ANON-ID-001"
	explicitlyKeptPatientID = "patient-123"
)

func TestNewAnonymizer(t *testing.T) {
	// Test with nil profile (should use default BasicProfile)
	anon := NewAnonymizer(nil)
	if anon == nil {
		t.Fatal("NewAnonymizer returned nil")
		return
	}
	if anon.Profile == nil {
		t.Fatal("Profile is nil")
	}
	if anon.ReplacedUIDs == nil {
		t.Fatal("ReplacedUIDs map is nil")
	}

	// Test with custom profile
	customProfile := NewSecurityProfile(BasicProfile | RetainUIDs)
	anon2 := NewAnonymizer(customProfile)
	if anon2.Profile != customProfile {
		t.Error("Custom profile not set correctly")
	}
}

func TestNewSecurityProfile(t *testing.T) {
	profile := NewSecurityProfile(BasicProfile)
	if profile == nil {
		t.Fatal("NewSecurityProfile returned nil")
		return
	}
	if len(profile.rules) == 0 {
		t.Error("Profile has no rules loaded")
	}
}

func TestSecurityProfileAddRule(t *testing.T) {
	profile := &SecurityProfile{rules: make([]profileRule, 0)}

	// Test valid pattern
	err := profile.AddRule("0010,0010", ActionX)
	if err != nil {
		t.Errorf("AddRule failed: %v", err)
	}

	// Test invalid pattern
	err = profile.AddRule("[invalid", ActionX)
	if err == nil {
		t.Error("AddRule should fail for invalid regex pattern")
	}
}

func TestSecurityProfileFindAction(t *testing.T) {
	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionX) // PatientName - Remove
	_ = profile.AddRule("0010,0020", ActionZ) // PatientID - Zero
	_ = profile.AddRule("0008,0018", ActionU) // SOPInstanceUID - UID replacement

	tests := []struct {
		name         string
		tag          *tag.Tag
		expectAction SecurityProfileAction
		expectFound  bool
	}{
		{"PatientName", tag.PatientName, ActionX, true},
		{"PatientID", tag.PatientID, ActionZ, true},
		{"SOPInstanceUID", tag.SOPInstanceUID, ActionU, true},
		{"Unmatched", tag.Rows, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			action, found := profile.FindAction(tt.tag)
			if found != tt.expectFound {
				t.Errorf("FindAction found = %v, want %v", found, tt.expectFound)
			}
			if found && action != tt.expectAction {
				t.Errorf("FindAction action = %v, want %v", action, tt.expectAction)
			}
		})
	}
}

func TestAnonymizerRemoveAction(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	_ = ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{testStudyDescription}))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionX) // Remove PatientName

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// PatientName should be removed
	if _, exists := ds.GetString(tag.PatientName); exists {
		t.Error("PatientName should have been removed")
	}

	// PatientID should still exist
	if _, exists := ds.GetString(tag.PatientID); !exists {
		t.Error("PatientID should still exist")
	}
}

func TestAnonymizerReplaceString(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.InstitutionName, vr.LO, []string{"Test Hospital"}))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionD) // Dummy value for PatientName
	_ = profile.AddRule("0008,0080", ActionC) // Clean InstitutionName

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// Both should be replaced with ANONYMOUS
	patientName, _ := ds.GetString(tag.PatientName)
	if patientName != anonymizedString {
		t.Errorf("PatientName = %q, want ANONYMOUS", patientName)
	}

	institutionName, _ := ds.GetString(tag.InstitutionName)
	if institutionName != anonymizedString {
		t.Errorf("InstitutionName = %q, want ANONYMOUS", institutionName)
	}
}

func TestAnonymizerReplaceUID(t *testing.T) {
	ds := dataset.New()
	sopInstanceUID := "1.2.3.4.5.6.7"
	studyInstanceUID := "1.2.3.4.5.6.8"
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID}))
	_ = ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{studyInstanceUID}))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0008,0018", ActionU) // Replace SOPInstanceUID
	_ = profile.AddRule("0020,000D", ActionU) // Replace StudyInstanceUID

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// UIDs should be replaced
	newSOPUID, _ := ds.GetString(tag.SOPInstanceUID)
	if newSOPUID == sopInstanceUID {
		t.Error("SOPInstanceUID should have been replaced")
	}
	if newSOPUID == "" {
		t.Error("SOPInstanceUID should not be empty")
	}

	newStudyUID, _ := ds.GetString(tag.StudyInstanceUID)
	if newStudyUID == studyInstanceUID {
		t.Error("StudyInstanceUID should have been replaced")
	}

	// Check UID consistency - same UID should get same replacement
	ds2 := dataset.New()
	_ = ds2.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	err = anon.AnonymizeInPlace(ds2)
	if err != nil {
		t.Fatalf("Second AnonymizeInPlace failed: %v", err)
	}

	newSOPUID2, _ := ds2.GetString(tag.SOPInstanceUID)
	if newSOPUID2 != newSOPUID {
		t.Error("Same UID should get same replacement across multiple calls")
	}
}

func TestAnonymizerBlankItemZero(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0020", ActionZ) // Zero-length PatientID
	_ = profile.AddRule("0028,0010", ActionZ) // Zero-length Rows

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// PatientID should be empty
	patientID, exists := ds.GetString(tag.PatientID)
	if !exists {
		t.Error("PatientID element should exist")
	}
	if patientID != "" {
		t.Errorf("PatientID should be empty, got %q", patientID)
	}
}

func TestAnonymizerSequence(t *testing.T) {
	// Create dataset with sequence
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))

	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	item := dataset.New()
	_ = item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{testStudyUID}))
	_ = item.Add(element.NewString(tag.StudyDescription, vr.LO, []string{testStudyDescription}))
	seq.AddItem(item)
	_ = ds.Add(seq)

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionX) // Remove PatientName
	_ = profile.AddRule("0020,000D", ActionU) // Replace StudyInstanceUID in sequence
	_ = profile.AddRule("0008,1030", ActionC) // Clean StudyDescription

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// PatientName should be removed
	if _, exists := ds.GetString(tag.PatientName); exists {
		t.Error("PatientName should have been removed")
	}

	// Check sequence item was anonymized
	seqElem, exists := ds.Get(tag.ReferencedStudySequence)
	if !exists {
		t.Fatal("Sequence should still exist")
	}

	seqAnon, ok := seqElem.(*dataset.Sequence)
	if !ok {
		t.Fatal("Element is not a sequence")
	}

	if seqAnon.Count() != 1 {
		t.Fatal("Sequence should have 1 item")
	}

	itemAnon := seqAnon.GetItem(0)

	// StudyInstanceUID should be replaced
	studyUID, _ := itemAnon.GetString(tag.StudyInstanceUID)
	if studyUID == testStudyUID {
		t.Error("StudyInstanceUID in sequence should have been replaced")
	}

	// StudyDescription should be cleaned
	studyDesc, _ := itemAnon.GetString(tag.StudyDescription)
	if studyDesc == testStudyDescription {
		t.Error("StudyDescription should have been anonymized")
	}
}

func TestAnonymizerCleanSequencePreservesItems(t *testing.T) {
	ds := dataset.New()
	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	item := dataset.New()
	_ = item.Add(element.NewString(tag.StudyDescription, vr.LO, []string{testStudyDescription}))
	seq.AddItem(item)
	_ = ds.Add(seq)

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0008,1110", ActionC)
	_ = profile.AddRule("0008,1030", ActionD)

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}

	got, ok := ds.GetOrNil(tag.ReferencedStudySequence).(*dataset.Sequence)
	if !ok {
		t.Fatalf("cleaned sequence = %T, want *dataset.Sequence", got)
	}
	if got.Count() != 1 {
		t.Fatalf("cleaned sequence has %d items, want one", got.Count())
	}
	if value, _ := got.GetItem(0).GetString(tag.StudyDescription); value != anonymizedString {
		t.Errorf("StudyDescription = %q, want ANONYMOUS", value)
	}
	if got.GetItem(0).Contains(tag.PatientIdentityRemoved) {
		t.Error("nested sequence item contains top-level de-identification declaration")
	}
}

func TestAnonymizerDummySequencePreservesItems(t *testing.T) {
	ds := dataset.New()
	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	item := dataset.New()
	_ = item.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	seq.AddItem(item)
	_ = ds.Add(seq)

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0008,1110", ActionD)
	_ = profile.AddRule("0010,0010", ActionD)

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}

	got, ok := ds.GetOrNil(tag.ReferencedStudySequence).(*dataset.Sequence)
	if !ok {
		t.Fatalf("dummy sequence = %T, want *dataset.Sequence", got)
	}
	if got.Count() != 1 {
		t.Fatalf("dummy sequence has %d items, want one", got.Count())
	}
	if value, _ := got.GetItem(0).GetString(tag.PatientName); value != anonymizedString {
		t.Errorf("PatientName = %q, want ANONYMOUS", value)
	}
}

func TestAnonymizerDummyStringsAreValidForVR(t *testing.T) {
	ds := dataset.New()
	values := []struct {
		tag   *tag.Tag
		vr    *vr.VR
		input string
		want  string
	}{
		{tag: tag.StudyDate, vr: vr.DA, input: "20260829", want: "19000101"},
		{tag: tag.AcquisitionDateTime, vr: vr.DT, input: "20260829120000", want: "19000101000000"},
		{tag: tag.StudyTime, vr: vr.TM, input: "120000", want: "000000"},
		{tag: tag.PatientAge, vr: vr.AS, input: "042Y", want: "000Y"},
		{tag: tag.PatientWeight, vr: vr.DS, input: "70.5", want: "0"},
		{tag: tag.InstanceNumber, vr: vr.IS, input: "7", want: "0"},
	}
	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	for _, value := range values {
		_ = ds.Add(element.NewString(value.tag, value.vr, []string{value.input}))
		tagText := value.tag.String()
		_ = profile.AddRule(tagText[1:len(tagText)-1], ActionD)
	}

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	for _, value := range values {
		if got, _ := ds.GetString(value.tag); got != value.want {
			t.Errorf("%s = %q, want %q", value.tag, got, value.want)
		}
	}
}

func TestAnonymizerAddsTopLevelDeidentificationDeclaration(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))

	if err := NewAnonymizer(nil).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	if got, _ := ds.GetString(tag.PatientIdentityRemoved); got != "YES" {
		t.Errorf("PatientIdentityRemoved = %q, want YES", got)
	}
	if got, _ := ds.GetString(tag.DeidentificationMethod); got == "" {
		t.Error("DeidentificationMethod is empty")
	}
}

func TestAnonymizeFileInPlaceRebuildsFileMetaInformation(t *testing.T) {
	const (
		sopClassUID       = "1.2.840.10008.5.1.4.1.1.2"
		originalSOPUID    = "1.2.826.0.1.3680043.10.999.1"
		oldImplementation = "1.2.826.0.1.3680043.10.999.2"
	)
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{sopClassUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{originalSOPUID}))
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))

	source := dataset.NewDefaultFileMetaInformation()
	_ = source.SetMediaStorageSOPClassUID(sopClassUID)
	_ = source.SetMediaStorageSOPInstanceUID(originalSOPUID)
	_ = source.SetTransferSyntax(transfer.ExplicitVRLittleEndian)
	_ = source.SetImplementationClassUID(oldImplementation)
	_ = source.SetImplementationVersionName("OLD_VERSION")
	_ = source.SetSourceApplicationEntityTitle("SOURCE_AE")
	_ = source.SetSendingApplicationEntityTitle("SENDING_AE")
	_ = source.SetReceivingApplicationEntityTitle("RECEIVING_AE")
	_ = source.SetPrivateInformationCreatorUID(oldImplementation)
	_ = source.SetPrivateInformation([]byte("patient-linked private metadata"))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0008,0018", ActionU)
	fresh, err := NewAnonymizer(profile).AnonymizeFileInPlace(ds, source)
	if err != nil {
		t.Fatalf("AnonymizeFileInPlace() error = %v", err)
	}
	if fresh == source {
		t.Fatal("AnonymizeFileInPlace() reused input File Meta Information")
	}
	wantSOPUID, _ := ds.GetString(tag.SOPInstanceUID)
	if wantSOPUID == originalSOPUID {
		t.Fatal("SOPInstanceUID was not anonymized")
	}
	if got, _ := fresh.MediaStorageSOPInstanceUID(); got != wantSOPUID {
		t.Errorf("MediaStorageSOPInstanceUID = %q, want %q", got, wantSOPUID)
	}
	if got, _ := fresh.TransferSyntaxUID(); got != transfer.ExplicitVRLittleEndian.UID().UID() {
		t.Errorf("TransferSyntaxUID = %q, want Explicit VR Little Endian", got)
	}
	for _, identifyingTag := range []*tag.Tag{
		tag.SourceApplicationEntityTitle,
		tag.SendingApplicationEntityTitle,
		tag.ReceivingApplicationEntityTitle,
		tag.PrivateInformationCreatorUID,
		tag.PrivateInformation,
	} {
		if fresh.Dataset().Contains(identifyingTag) {
			t.Errorf("fresh File Meta Information contains %s", identifyingTag)
		}
	}
	if got, _ := fresh.ImplementationClassUID(); got == oldImplementation {
		t.Errorf("ImplementationClassUID retained source value %q", got)
	}
}

func TestAnonymizerCustomPatientInfo(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	profile := NewSecurityProfile(BasicProfile)
	profile.PatientName = customPatientName
	profile.PatientID = customPatientID

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// Check custom patient name
	patientName, _ := ds.GetString(tag.PatientName)
	if patientName != customPatientName {
		t.Errorf("PatientName = %q, want ANONYMOUS^PATIENT", patientName)
	}

	// Check custom patient ID
	patientID, _ := ds.GetString(tag.PatientID)
	if patientID != customPatientID {
		t.Errorf("PatientID = %q, want ANON-ID-001", patientID)
	}
}

func TestSecurityProfileOverrideActionKeepsPatientID(t *testing.T) {
	const sopInstanceUID = "1.2.840.10008.1.2.3.4"

	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{explicitlyKeptPatientID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID}))

	profile := NewSecurityProfile(BasicProfile | RetainUIDs)
	profile.OverrideAction(tag.PatientID, ActionK)

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}

	if patientID, _ := ds.GetString(tag.PatientID); patientID != explicitlyKeptPatientID {
		t.Errorf("PatientID = %q, want original value", patientID)
	}
	if patientName, _ := ds.GetString(tag.PatientName); patientName != "" {
		t.Errorf("PatientName = %q, want empty", patientName)
	}
	if got, _ := ds.GetString(tag.SOPInstanceUID); got != sopInstanceUID {
		t.Errorf("SOPInstanceUID = %q, want original value", got)
	}
	if got, _ := ds.GetString(tag.PatientIdentityRemoved); got != "NO" {
		t.Errorf("PatientIdentityRemoved = %q, want NO", got)
	}
	if got, _ := ds.GetString(tag.DeidentificationMethod); !strings.Contains(got, tag.PatientID.String()) {
		t.Errorf("DeidentificationMethod = %q, want retained PatientID override", got)
	}
}

func TestExplicitKeepWithConfiguredReplacementReportsIdentityRemoved(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{explicitlyKeptPatientID}))
	profile := NewSecurityProfile(BasicProfile)
	profile.OverrideAction(tag.PatientID, ActionK)
	profile.PatientID = customPatientID

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	if got, _ := ds.GetString(tag.PatientID); got != profile.PatientID {
		t.Errorf("PatientID = %q, want configured replacement", got)
	}
	if got, _ := ds.GetString(tag.PatientIdentityRemoved); got != "YES" {
		t.Errorf("PatientIdentityRemoved = %q, want YES", got)
	}
}

func TestAnonymizeDeepClonesSequenceItems(t *testing.T) {
	item := dataset.New()
	_ = item.Add(element.NewString(tag.PatientID, vr.LO, []string{explicitlyKeptPatientID}))
	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	seq.AddItem(item)
	ds := dataset.New()
	_ = ds.Add(seq)

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0020", ActionD)
	got, err := NewAnonymizer(profile).Anonymize(ds)
	if err != nil {
		t.Fatalf("Anonymize() error = %v", err)
	}

	gotSeq := got.GetOrNil(tag.ReferencedStudySequence).(*dataset.Sequence)
	if value, _ := gotSeq.GetItem(0).GetString(tag.PatientID); value == explicitlyKeptPatientID {
		t.Fatal("anonymized sequence item retained the original PatientID")
	}
	if value, _ := item.GetString(tag.PatientID); value != explicitlyKeptPatientID {
		t.Errorf("original sequence PatientID = %q, want patient-123", value)
	}
}

func TestCustomPatientReplacementsSurviveRemoveAction(t *testing.T) {
	profile := &SecurityProfile{
		rules:       make([]profileRule, 0),
		PatientName: customPatientName,
		PatientID:   customPatientID,
	}
	_ = profile.AddRule("0010,0010", ActionX)
	_ = profile.AddRule("0010,0020", ActionX)
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{explicitlyKeptPatientID}))

	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	if got, _ := ds.GetString(tag.PatientName); got != profile.PatientName {
		t.Errorf("PatientName = %q, want %q", got, profile.PatientName)
	}
	if got, _ := ds.GetString(tag.PatientID); got != profile.PatientID {
		t.Errorf("PatientID = %q, want %q", got, profile.PatientID)
	}

	absent := dataset.New()
	if err := NewAnonymizer(profile).AnonymizeInPlace(absent); err != nil {
		t.Fatalf("AnonymizeInPlace(absent) error = %v", err)
	}
	if absent.Contains(tag.PatientName) || absent.Contains(tag.PatientID) {
		t.Fatal("custom replacements added patient attributes that were absent")
	}
}

func TestAnonymizeFileInPlaceAppliesSourceTransferSyntaxToDataset(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.826.0.1.3680043.10.999.1"}))
	source := dataset.NewDefaultFileMetaInformation()
	_ = source.SetTransferSyntax(transfer.ExplicitVRBigEndian)

	fmi, err := NewAnonymizer(&SecurityProfile{}).AnonymizeFileInPlace(ds, source)
	if err != nil {
		t.Fatalf("AnonymizeFileInPlace() error = %v", err)
	}
	if ds.InternalTransferSyntax() != transfer.ExplicitVRBigEndian {
		t.Fatalf("Dataset transfer syntax = %v, want Explicit VR Big Endian", ds.InternalTransferSyntax())
	}
	if got, _ := fmi.TransferSyntax(); got != transfer.ExplicitVRBigEndian {
		t.Fatalf("File Meta transfer syntax = %v, want Explicit VR Big Endian", got)
	}
}

func TestAnonymizerClone(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{testPatientName}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionX) // Remove PatientName

	anon := NewAnonymizer(profile)
	anonDS, err := anon.Anonymize(ds)
	if err != nil {
		t.Fatalf("Anonymize failed: %v", err)
	}

	// Original should be unchanged
	if _, exists := ds.GetString(tag.PatientName); !exists {
		t.Error("Original dataset should not be modified")
	}

	// Clone should be anonymized
	if _, exists := anonDS.GetString(tag.PatientName); exists {
		t.Error("Cloned dataset should have PatientName removed")
	}
}

func TestSecurityProfileActions(t *testing.T) {
	tests := []struct {
		action   SecurityProfileAction
		expected string
	}{
		{ActionD, "D"},
		{ActionZ, "Z"},
		{ActionX, "X"},
		{ActionK, "K"},
		{ActionC, "C"},
		{ActionU, "U"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			str := tt.action.String()
			if str != tt.expected {
				t.Errorf("String() = %q, want %q", str, tt.expected)
			}
		})
	}
}

func TestSecurityProfileOptions(t *testing.T) {
	// Test combined options
	options := BasicProfile | RetainUIDs | RetainDeviceIdent

	profile := NewSecurityProfile(options)
	if profile == nil {
		t.Fatal("NewSecurityProfile returned nil")
		return
	}
	if len(profile.rules) == 0 {
		t.Error("Profile with combined options should have rules")
	}
}
