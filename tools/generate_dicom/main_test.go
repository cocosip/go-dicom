// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	privateFlag        = "-private"
	rootFlag           = "-root"
	standardFlag       = "-standard"
	version2026bMarker = "version 2026b"
)

func TestRunCommandGeneratesAllXMLDerivedData(t *testing.T) {
	inputDir := t.TempDir()
	standard := filepath.Join(inputDir, "DICOM Dictionary.xml")
	private := filepath.Join(inputDir, "Private Dictionary.xml")
	root := t.TempDir()

	const standardXML = `<dictionary version="2026b">
  <tag group="0010" element="0010" keyword="PatientName" vr="PN" vm="1">Patient Name</tag>
  <uid uid="1.2.840.10008.1.1" keyword="Verification" type="SOP Class">Verification SOP Class</uid>
</dictionary>`
	const privateXML = `<dictionaries>
  <dictionary creator="ACME">
    <tag group="0011" element="xx10" keyword="AcmeValue" vr="LO" vm="1">ACME Value</tag>
  </dictionary>
</dictionaries>`
	writeTestFile(t, standard, standardXML)
	writeTestFile(t, private, privateXML)

	var stdout bytes.Buffer
	err := runCommand(
		[]string{standardFlag, standard, privateFlag, private, rootFlag, root},
		&stdout,
	)
	if err != nil {
		t.Fatalf("runCommand() error = %v", err)
	}

	tests := []struct {
		path string
		want []string
	}{
		{
			path: "pkg/dicom/tag/tags_generated.go",
			want: []string{version2026bMarker, "PatientName = New(0x0010, 0x0010)"},
		},
		{
			path: "pkg/dicom/uid/uids_generated.go",
			want: []string{
				version2026bMarker,
				`Verification = New("1.2.840.10008.1.1", "Verification SOP Class", TypeSOPClass, false)`,
			},
		},
		{
			path: "pkg/dicom/dict/dictionary_data.go",
			want: []string{version2026bMarker, "tag.New(0x0010, 0x0010)", `"PatientName"`},
		},
		{
			path: "pkg/dicom/dict/private_dictionary_data.go",
			want: []string{"loadPrivateEntries", `GetPrivateDictionary("ACME")`, "tag.MustParseMaskedTag"},
		},
	}
	for _, test := range tests {
		generated, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(test.path)))
		if err != nil {
			t.Fatalf("ReadFile(%s) error = %v", test.path, err)
		}
		for _, want := range test.want {
			if !strings.Contains(string(generated), want) {
				t.Errorf("generated %s does not contain %q", test.path, want)
			}
		}
	}
	if got := stdout.String(); !strings.Contains(got, "1 tags, 1 UIDs, 1 private creators, 1 private entries") {
		t.Errorf("stdout = %q, want generated counts", got)
	}
}

func TestGetVMCodeUsesCurrentVMSymbolNames(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{value: "1-2", want: "vm.VM12"},
		{value: "1-8", want: "vm.VM18"},
		{value: "1-n", want: "vm.VM1N"},
		{value: "2-2n", want: "vm.VM22N"},
	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			if got := getVMCode(test.value); got != test.want {
				t.Fatalf("getVMCode(%q) = %q, want %q", test.value, got, test.want)
			}
		})
	}
}

func TestRunCommandUsesFODicomRetiredIdentifierConvention(t *testing.T) {
	inputDir := t.TempDir()
	standard := filepath.Join(inputDir, "DICOM Dictionary.xml")
	private := filepath.Join(inputDir, "Private Dictionary.xml")
	root := t.TempDir()

	const standardXML = `<dictionary version="2026b">
  <tag group="0008" element="0001" keyword="LengthToEnd" vr="UL" vm="1" retired="true">Length to End</tag>
  <uid uid="1.2.840.10008.1.2.2" keyword="ExplicitVRBigEndian" type="Transfer Syntax" retired="true">Explicit VR Big Endian (Retired)</uid>
</dictionary>`
	const privateXML = `<dictionaries></dictionaries>`
	writeTestFile(t, standard, standardXML)
	writeTestFile(t, private, privateXML)

	if err := runCommand(
		[]string{standardFlag, standard, privateFlag, private, rootFlag, root},
		io.Discard,
	); err != nil {
		t.Fatalf("runCommand() error = %v", err)
	}

	tags := readTestFile(t, filepath.Join(root, filepath.FromSlash(tagsOutput)))
	if !strings.Contains(tags, "LengthToEndRETIRED = New(0x0008, 0x0001)") {
		t.Errorf("generated tags do not use fo-dicom retired identifier convention")
	}
	uids := readTestFile(t, filepath.Join(root, filepath.FromSlash(uidsOutput)))
	if !strings.Contains(uids, "ExplicitVRBigEndianRETIRED = New(") ||
		!strings.Contains(uids, "Register(ExplicitVRBigEndianRETIRED)") {
		t.Errorf("generated UIDs do not use fo-dicom retired identifier convention")
	}
	dictionary := readTestFile(t, filepath.Join(root, filepath.FromSlash(dictionaryOutput)))
	if !strings.Contains(dictionary, `"LengthToEnd"`) {
		t.Errorf("generated dictionary does not preserve the XML keyword")
	}
}

func TestBundled2026bRegeneratesCommittedData(t *testing.T) {
	repositoryRoot := filepath.Clean(filepath.Join("..", ".."))
	standard := filepath.Join(repositoryRoot, "tools", "data", "2026b", "DICOM Dictionary.xml")
	private := filepath.Join(repositoryRoot, "tools", "data", "2026b", "Private Dictionary.xml")
	generatedRoot := t.TempDir()

	var stdout bytes.Buffer
	if err := runCommand(
		[]string{standardFlag, standard, privateFlag, private, rootFlag, generatedRoot},
		&stdout,
	); err != nil {
		t.Fatalf("runCommand() error = %v", err)
	}

	const wantCounts = "Generated 5347 tags, 1928 UIDs, 235 private creators, 4678 private entries from DICOM 2026b\n"
	if got := stdout.String(); got != wantCounts {
		t.Fatalf("runCommand() output = %q, want %q", got, wantCounts)
	}

	for _, output := range []string{
		tagsOutput,
		uidsOutput,
		dictionaryOutput,
		privateDictionaryOutput,
	} {
		committed, err := os.ReadFile(filepath.Join(repositoryRoot, filepath.FromSlash(output)))
		if err != nil {
			t.Fatalf("ReadFile(committed %s) error = %v", output, err)
		}
		generated, err := os.ReadFile(filepath.Join(generatedRoot, filepath.FromSlash(output)))
		if err != nil {
			t.Fatalf("ReadFile(generated %s) error = %v", output, err)
		}
		if !bytes.Equal(generated, committed) {
			t.Errorf("generated %s differs from committed data", output)
		}
	}
}

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", path, err)
	}
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%s) error = %v", path, err)
	}
	return string(data)
}
