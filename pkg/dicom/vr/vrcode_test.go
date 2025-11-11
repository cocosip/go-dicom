// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vr_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// vrTestCase holds test data for VR code tests.
type vrTestCase struct {
	name         string
	codeConstant string
	vrInstance   *vr.VR
}

// allVRTestCases returns all standard VR test cases.
func allVRTestCases() []vrTestCase {
	return []vrTestCase{
		{"AE", vr.CodeAE, vr.AE},
		{"AS", vr.CodeAS, vr.AS},
		{"AT", vr.CodeAT, vr.AT},
		{"CS", vr.CodeCS, vr.CS},
		{"DA", vr.CodeDA, vr.DA},
		{"DS", vr.CodeDS, vr.DS},
		{"DT", vr.CodeDT, vr.DT},
		{"FD", vr.CodeFD, vr.FD},
		{"FL", vr.CodeFL, vr.FL},
		{"IS", vr.CodeIS, vr.IS},
		{"LO", vr.CodeLO, vr.LO},
		{"LT", vr.CodeLT, vr.LT},
		{"OB", vr.CodeOB, vr.OB},
		{"OD", vr.CodeOD, vr.OD},
		{"OF", vr.CodeOF, vr.OF},
		{"OL", vr.CodeOL, vr.OL},
		{"OV", vr.CodeOV, vr.OV},
		{"OW", vr.CodeOW, vr.OW},
		{"PN", vr.CodePN, vr.PN},
		{"SH", vr.CodeSH, vr.SH},
		{"SL", vr.CodeSL, vr.SL},
		{"SQ", vr.CodeSQ, vr.SQ},
		{"SS", vr.CodeSS, vr.SS},
		{"ST", vr.CodeST, vr.ST},
		{"SV", vr.CodeSV, vr.SV},
		{"TM", vr.CodeTM, vr.TM},
		{"UC", vr.CodeUC, vr.UC},
		{"UI", vr.CodeUI, vr.UI},
		{"UL", vr.CodeUL, vr.UL},
		{"UN", vr.CodeUN, vr.UN},
		{"UR", vr.CodeUR, vr.UR},
		{"US", vr.CodeUS, vr.US},
		{"UT", vr.CodeUT, vr.UT},
		{"UV", vr.CodeUV, vr.UV},
	}
}

// TestVRCodeConstants verifies that the VR code constants match the VR instance codes
func TestVRCodeConstants(t *testing.T) {
	for _, tt := range allVRTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.codeConstant != tt.vrInstance.Code() {
				t.Errorf("Code constant %q does not match VR instance code %q",
					tt.codeConstant, tt.vrInstance.Code())
			}
		})
	}
}

// TestVRCodeParsing verifies that the code constants can be used to parse VR instances
func TestVRCodeParsing(t *testing.T) {
	for _, tt := range allVRTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := vr.Parse(tt.codeConstant)
			if err != nil {
				t.Errorf("Parse(%q) returned error: %v", tt.codeConstant, err)
				return
			}
			if parsed != tt.vrInstance {
				t.Errorf("Parse(%q) = %v, want %v", tt.codeConstant, parsed, tt.vrInstance)
			}
		})
	}
}

// TestVRCodeUsageExample demonstrates how to use VR code constants
func TestVRCodeUsageExample(t *testing.T) {
	// Example 1: Use code constant to parse VR
	personNameVR, err := vr.Parse(vr.CodePN)
	if err != nil {
		t.Fatalf("Failed to parse PN: %v", err)
	}
	if personNameVR.Name() != "Person Name" {
		t.Errorf("Expected 'Person Name', got %q", personNameVR.Name())
	}

	// Example 2: Use code constant for comparison
	someVR := vr.UI
	if someVR.Code() == vr.CodeUI {
		// This is a Unique Identifier VR
		if !someVR.IsString() {
			t.Error("UI should be a string type")
		}
	}

	// Example 3: Use code constant in a map
	vrMap := map[string]string{
		vr.CodePN: "Patient Name",
		vr.CodeUI: "Study Instance UID",
		vr.CodeDA: "Study Date",
	}
	if vrMap[vr.CodePN] != "Patient Name" {
		t.Error("Map lookup failed")
	}
}
