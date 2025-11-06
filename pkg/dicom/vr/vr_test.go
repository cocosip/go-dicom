// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vr_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    *vr.VR
		wantErr bool
	}{
		{"valid AE", "AE", vr.AE, false},
		{"valid PN", "PN", vr.PN, false},
		{"valid US", "US", vr.US, false},
		{"valid SQ", "SQ", vr.SQ, false},
		{"invalid code", "XX", nil, true},
		{"empty code", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vr.Parse(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    *vr.VR
		wantErr bool
	}{
		{"valid AE", []byte("AE"), vr.AE, false},
		{"valid PN", []byte("PN"), vr.PN, false},
		{"valid US", []byte("US"), vr.US, false},
		{"valid OB", []byte("OB"), vr.OB, false},
		{"invalid code", []byte("XX"), nil, true},
		{"too short", []byte("A"), nil, true},
		{"empty", []byte(""), nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vr.ParseBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryParse(t *testing.T) {
	result, ok := vr.TryParse("PN")
	if !ok {
		t.Error("TryParse(\"PN\") should return true")
	}
	if result != vr.PN {
		t.Errorf("TryParse(\"PN\") = %v, want %v", result, vr.PN)
	}

	_, ok = vr.TryParse("XX")
	if ok {
		t.Error("TryParse(\"XX\") should return false")
	}
}

func TestVRProperties(t *testing.T) {
	// Test PN (Person Name) properties
	if vr.PN.Code() != "PN" {
		t.Errorf("PN.Code() = %v, want PN", vr.PN.Code())
	}
	if vr.PN.Name() != "Person Name" {
		t.Errorf("PN.Name() = %v, want Person Name", vr.PN.Name())
	}
	if !vr.PN.IsString() {
		t.Error("PN.IsString() should be true")
	}
	if !vr.PN.IsStringEncoded() {
		t.Error("PN.IsStringEncoded() should be true")
	}
	if !vr.PN.Is16bitLength() {
		t.Error("PN.Is16bitLength() should be true")
	}
	if vr.PN.MaximumLength() != 64 {
		t.Errorf("PN.MaximumLength() = %v, want 64", vr.PN.MaximumLength())
	}

	// Test US (Unsigned Short) properties
	if vr.US.IsString() {
		t.Error("US.IsString() should be false")
	}
	if vr.US.UnitSize() != 2 {
		t.Errorf("US.UnitSize() = %v, want 2", vr.US.UnitSize())
	}
	if vr.US.ByteSwap() != 2 {
		t.Errorf("US.ByteSwap() = %v, want 2", vr.US.ByteSwap())
	}

	// Test OB (Other Byte) properties
	if vr.OB.Is16bitLength() {
		t.Error("OB.Is16bitLength() should be false")
	}
	if vr.OB.MaximumLength() != 0 {
		t.Errorf("OB.MaximumLength() = %v, want 0 (unlimited)", vr.OB.MaximumLength())
	}
}

func TestVRString(t *testing.T) {
	if vr.PN.String() != "PN" {
		t.Errorf("PN.String() = %v, want PN", vr.PN.String())
	}
	if vr.AE.String() != "AE" {
		t.Errorf("AE.String() = %v, want AE", vr.AE.String())
	}
}
