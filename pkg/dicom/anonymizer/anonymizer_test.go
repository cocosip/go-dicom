// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package anonymizer

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewAnonymizer(t *testing.T) {
	// Test with nil profile (should use default BasicProfile)
	anon := NewAnonymizer(nil)
	if anon == nil {
		t.Fatal("NewAnonymizer returned nil")
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
	_ = profile.AddRule("0010,0010", ActionX)   // PatientName - Remove
	_ = profile.AddRule("0010,0020", ActionZ)   // PatientID - Zero
	_ = profile.AddRule("0008,0018", ActionU)   // SOPInstanceUID - UID replacement

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
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	ds.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Test Study"}))

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
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.InstitutionName, vr.LO, []string{"Test Hospital"}))

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
	if patientName != "ANONYMOUS" {
		t.Errorf("PatientName = %q, want ANONYMOUS", patientName)
	}

	institutionName, _ := ds.GetString(tag.InstitutionName)
	if institutionName != "ANONYMOUS" {
		t.Errorf("InstitutionName = %q, want ANONYMOUS", institutionName)
	}
}

func TestAnonymizerReplaceUID(t *testing.T) {
	ds := dataset.New()
	sopInstanceUID := "1.2.3.4.5.6.7"
	studyInstanceUID := "1.2.3.4.5.6.8"
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID}))
	ds.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{studyInstanceUID}))

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
	ds2.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID}))

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
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))
	ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))

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
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))

	seq := dataset.NewSequence(tag.ReferencedStudySequence)
	item := dataset.New()
	item.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
	item.Add(element.NewString(tag.StudyDescription, vr.LO, []string{"Test Study"}))
	seq.AddItem(item)
	ds.Add(seq)

	profile := &SecurityProfile{rules: make([]profileRule, 0)}
	_ = profile.AddRule("0010,0010", ActionX)   // Remove PatientName
	_ = profile.AddRule("0020,000D", ActionU)   // Replace StudyInstanceUID in sequence
	_ = profile.AddRule("0008,1030", ActionC)   // Clean StudyDescription

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
	if studyUID == "1.2.3.4.5" {
		t.Error("StudyInstanceUID in sequence should have been replaced")
	}

	// StudyDescription should be cleaned
	studyDesc, _ := itemAnon.GetString(tag.StudyDescription)
	if studyDesc == "Test Study" {
		t.Error("StudyDescription should have been anonymized")
	}
}

func TestAnonymizerCustomPatientInfo(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	profile := NewSecurityProfile(BasicProfile)
	profile.PatientName = "ANONYMOUS^PATIENT"
	profile.PatientID = "ANON-ID-001"

	anon := NewAnonymizer(profile)
	err := anon.AnonymizeInPlace(ds)
	if err != nil {
		t.Fatalf("AnonymizeInPlace failed: %v", err)
	}

	// Check custom patient name
	patientName, _ := ds.GetString(tag.PatientName)
	if patientName != "ANONYMOUS^PATIENT" {
		t.Errorf("PatientName = %q, want ANONYMOUS^PATIENT", patientName)
	}

	// Check custom patient ID
	patientID, _ := ds.GetString(tag.PatientID)
	if patientID != "ANON-ID-001" {
		t.Errorf("PatientID = %q, want ANON-ID-001", patientID)
	}
}

func TestAnonymizerClone(t *testing.T) {
	ds := dataset.New()
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^John"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

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
	}
	if len(profile.rules) == 0 {
		t.Error("Profile with combined options should have rules")
	}
}
