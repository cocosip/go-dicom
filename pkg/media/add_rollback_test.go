// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

// TestAddFileFailureDoesNotMutateDirectory verifies that a failed AddFile does
// not leave newly discovered PATIENT/STUDY/SERIES records attached to the tree.
func TestAddFileFailureDoesNotMutateDirectory(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}

	first := testParseResult(t, testFileValues{
		PatientID: testPatientID, PatientName: testPatientName, StudyUID: testStudyUID,
		SeriesUID: testSeriesUID, SOPInstanceUID: testSOPInstanceUID,
	})
	entry := addTestFile(t, dir, first, "DIR1/IMAGE001")
	want := countDirectoryRecords(t, dir)

	id, err := ParseFileID("DIR1/IMAGE001")
	if err != nil {
		t.Fatalf("ParseFileID() error = %v", err)
	}
	second := testParseResult(t, testFileValues{
		PatientID: testPatientID, PatientName: testPatientName, StudyUID: "1.2.3.9",
		SeriesUID: "1.2.3.9.1", SOPInstanceUID: testSOPInstanceUID2,
	})
	if _, err := dir.AddFile(second, id); err == nil || !strings.Contains(err.Error(), "already referenced") {
		t.Fatalf("AddFile(duplicate File ID) error = %v, want duplicate error", err)
	}

	if got := countDirectoryRecords(t, dir); got != want {
		t.Fatalf("directory records after failed AddFile = %d, want %d", got, want)
	}
	if got := len(dir.RootRecords()); got != 1 {
		t.Fatalf("root records after failed AddFile = %d, want 1", got)
	}
	if got := len(entry.Patient.Children()); got != 1 {
		t.Fatalf("patient children after failed AddFile = %d, want 1", got)
	}
}

// TestAddFileFailureForNewPatientRemovesCreatedHierarchy verifies that a failure
// after new PATIENT/STUDY/SERIES records were created removes them again and
// leaves the diagnostics count unchanged.
func TestAddFileFailureForNewPatientRemovesCreatedHierarchy(t *testing.T) {
	dir, err := NewDirectory()
	if err != nil {
		t.Fatalf("NewDirectory() error = %v", err)
	}

	first := testParseResult(t, testFileValues{
		PatientID: testPatientID, PatientName: testPatientName, StudyUID: testStudyUID,
		SeriesUID: testSeriesUID, SOPInstanceUID: testSOPInstanceUID,
	})
	addTestFile(t, dir, first, "DIR1/IMAGE001")
	wantRecords := countDirectoryRecords(t, dir)
	wantDiagnostics := len(dir.Diagnostics())

	// The Verification SOP Class has no directory record type, so AddFile fails
	// after the new patient, study, and series records were created.
	unsupported := testParseResult(t, testFileValues{
		PatientID: "NEWPATIENT", PatientName: "NEW^PATIENT", StudyUID: "1.2.3.10",
		SeriesUID: "1.2.3.10.1", SOPInstanceUID: "1.2.3.10.1.1", SOPClass: uid.Verification,
	})
	fileID, err := ParseFileID("DIR1/IMAGE002")
	if err != nil {
		t.Fatalf("ParseFileID() error = %v", err)
	}
	if _, err := dir.AddFile(unsupported, fileID); err == nil {
		t.Fatal("AddFile(unsupported SOP class) succeeded, want error")
	}

	if got := len(dir.RootRecords()); got != 1 {
		t.Fatalf("root records after failed AddFile = %d, want 1", got)
	}
	if got := countDirectoryRecords(t, dir); got != wantRecords {
		t.Fatalf("directory records after failed AddFile = %d, want %d", got, wantRecords)
	}
	if got := len(dir.Diagnostics()); got != wantDiagnostics {
		t.Fatalf("diagnostics after failed AddFile = %d, want %d", got, wantDiagnostics)
	}
}
