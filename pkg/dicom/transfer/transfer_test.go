// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transfer_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

func TestImplicitVRLittleEndian(t *testing.T) {
	ts := transfer.ImplicitVRLittleEndian

	if ts.IsExplicitVR() {
		t.Error("ImplicitVRLittleEndian should have ImplicitVR")
	}
	if ts.Endian() != endian.Little {
		t.Error("ImplicitVRLittleEndian should be Little Endian")
	}
	if ts.IsEncapsulated() {
		t.Error("ImplicitVRLittleEndian should not be encapsulated")
	}
	if ts.IsLossy() {
		t.Error("ImplicitVRLittleEndian should not be lossy")
	}
}

func TestExplicitVRLittleEndian(t *testing.T) {
	ts := transfer.ExplicitVRLittleEndian

	if !ts.IsExplicitVR() {
		t.Error("ExplicitVRLittleEndian should have ExplicitVR")
	}
	if ts.Endian() != endian.Little {
		t.Error("ExplicitVRLittleEndian should be Little Endian")
	}
	if ts.IsEncapsulated() {
		t.Error("ExplicitVRLittleEndian should not be encapsulated")
	}
}

func TestExplicitVRBigEndian(t *testing.T) {
	ts := transfer.ExplicitVRBigEndian

	if !ts.IsExplicitVR() {
		t.Error("ExplicitVRBigEndian should have ExplicitVR")
	}
	if ts.Endian() != endian.Big {
		t.Error("ExplicitVRBigEndian should be Big Endian")
	}
	if !ts.IsRetired() {
		t.Error("ExplicitVRBigEndian should be retired")
	}
}

func TestJPEGBaseline8Bit(t *testing.T) {
	ts := transfer.JPEGBaseline8Bit

	if !ts.IsExplicitVR() {
		t.Error("JPEGBaseline8Bit should have ExplicitVR")
	}
	if !ts.IsEncapsulated() {
		t.Error("JPEGBaseline8Bit should be encapsulated")
	}
	if !ts.IsLossy() {
		t.Error("JPEGBaseline8Bit should be lossy")
	}
	if ts.LossyCompressionMethod() != "ISO_10918_1" {
		t.Errorf("JPEGBaseline8Bit compression method = %q, want ISO_10918_1",
			ts.LossyCompressionMethod())
	}
}

func TestJPEG2000Lossless(t *testing.T) {
	ts := transfer.JPEG2000Lossless

	if !ts.IsExplicitVR() {
		t.Error("JPEG2000Lossless should have ExplicitVR")
	}
	if !ts.IsEncapsulated() {
		t.Error("JPEG2000Lossless should be encapsulated")
	}
	if ts.IsLossy() {
		t.Error("JPEG2000Lossless should not be lossy")
	}
}

func TestRLELossless(t *testing.T) {
	ts := transfer.RLELossless

	if !ts.IsExplicitVR() {
		t.Error("RLELossless should have ExplicitVR")
	}
	if !ts.IsEncapsulated() {
		t.Error("RLELossless should be encapsulated")
	}
	if ts.IsLossy() {
		t.Error("RLELossless should not be lossy")
	}
}

func TestDeflatedExplicitVRLittleEndian(t *testing.T) {
	ts := transfer.DeflatedExplicitVRLittleEndian

	if !ts.IsExplicitVR() {
		t.Error("DeflatedExplicitVRLittleEndian should have ExplicitVR")
	}
	if !ts.IsDeflate() {
		t.Error("DeflatedExplicitVRLittleEndian should use deflate")
	}
}

func TestGEPrivateImplicitVRBigEndian(t *testing.T) {
	ts := transfer.GEPrivateImplicitVRBigEndian

	if ts.IsExplicitVR() {
		t.Error("GEPrivateImplicitVRBigEndian should have ImplicitVR")
	}
	if ts.Endian() != endian.Little {
		t.Error("GEPrivateImplicitVRBigEndian should be Little Endian (for metadata)")
	}
	if !ts.SwapPixelData() {
		t.Error("GEPrivateImplicitVRBigEndian should swap pixel data")
	}
}

func TestParse(t *testing.T) {
	// Test parsing a known transfer syntax
	ts, err := transfer.Parse("1.2.840.10008.1.2")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if !ts.Equals(transfer.ImplicitVRLittleEndian) {
		t.Error("Parse() did not return ImplicitVRLittleEndian")
	}

	// Test parsing explicit VR Little Endian
	ts, err = transfer.Parse("1.2.840.10008.1.2.1")
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if !ts.Equals(transfer.ExplicitVRLittleEndian) {
		t.Error("Parse() did not return ExplicitVRLittleEndian")
	}
}

func TestLookup(t *testing.T) {
	// Test lookup of a known transfer syntax
	u := uid.ImplicitVRLittleEndian
	ts, err := transfer.Lookup(u)
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if !ts.Equals(transfer.ImplicitVRLittleEndian) {
		t.Error("Lookup() did not return ImplicitVRLittleEndian")
	}

	// Test lookup of nil UID
	_, err = transfer.Lookup(nil)
	if err == nil {
		t.Error("Lookup(nil) should return error")
	}

	// Test lookup of non-transfer-syntax UID
	nonTS := uid.New("1.2.3.4.5", "Not a Transfer Syntax", uid.TypeSOPClass, false)
	_, err = transfer.Lookup(nonTS)
	if err == nil {
		t.Error("Lookup() should return error for non-transfer-syntax UID")
	}
}

func TestEquals(t *testing.T) {
	ts1 := transfer.ImplicitVRLittleEndian
	ts2 := transfer.ImplicitVRLittleEndian
	ts3 := transfer.ExplicitVRLittleEndian

	if !ts1.Equals(ts2) {
		t.Error("Same transfer syntaxes should be equal")
	}
	if ts1.Equals(ts3) {
		t.Error("Different transfer syntaxes should not be equal")
	}
	if ts1.Equals(nil) {
		t.Error("Transfer syntax should not equal nil")
	}
}

func TestString(t *testing.T) {
	ts := transfer.ImplicitVRLittleEndian
	name := ts.String()
	if name == "" {
		t.Error("String() should return non-empty name")
	}
	if name != ts.UID().Name() {
		t.Errorf("String() = %q, want %q", name, ts.UID().Name())
	}
}

func TestRegisterAndQuery(t *testing.T) {
	// Create a custom transfer syntax
	customUID := uid.New("1.2.3.4.5.6", "Custom Transfer Syntax", uid.TypeTransferSyntax, false)
	customTS := transfer.NewBuilder(customUID).
		SetExplicitVR(true).
		SetEndian(endian.Little).
		Build()

	// Query should now find it
	queried := transfer.Query(customUID)
	if queried == nil {
		t.Fatal("Query() returned nil for registered transfer syntax")
	}
	if !queried.Equals(customTS) {
		t.Error("Query() returned different transfer syntax")
	}

	// Unregister it
	if !transfer.Unregister(customUID) {
		t.Error("Unregister() should return true for existing entry")
	}

	// Query should not find it anymore
	queried = transfer.Query(customUID)
	if queried != nil {
		t.Error("Query() should return nil after unregister")
	}
}

func TestKnownEntries(t *testing.T) {
	entries := transfer.KnownEntries()
	if len(entries) == 0 {
		t.Error("KnownEntries() should return non-empty list")
	}

	// Should contain at least the basic transfer syntaxes
	found := false
	for _, ts := range entries {
		if ts.Equals(transfer.ImplicitVRLittleEndian) {
			found = true
			break
		}
	}
	if !found {
		t.Error("KnownEntries() should contain ImplicitVRLittleEndian")
	}
}

func TestIsUncompressed(t *testing.T) {
	if !transfer.IsUncompressed(transfer.ImplicitVRLittleEndian) {
		t.Error("ImplicitVRLittleEndian should be uncompressed")
	}
	if transfer.IsUncompressed(transfer.JPEGBaseline8Bit) {
		t.Error("JPEGBaseline8Bit should be compressed")
	}
}

func TestIsLosslessCompressed(t *testing.T) {
	if transfer.IsLosslessCompressed(transfer.ImplicitVRLittleEndian) {
		t.Error("ImplicitVRLittleEndian should not be lossless compressed")
	}
	if !transfer.IsLosslessCompressed(transfer.RLELossless) {
		t.Error("RLELossless should be lossless compressed")
	}
	if transfer.IsLosslessCompressed(transfer.JPEGBaseline8Bit) {
		t.Error("JPEGBaseline8Bit should not be lossless compressed")
	}
}

func TestIsLossyCompressed(t *testing.T) {
	if transfer.IsLossyCompressed(transfer.ImplicitVRLittleEndian) {
		t.Error("ImplicitVRLittleEndian should not be lossy compressed")
	}
	if transfer.IsLossyCompressed(transfer.RLELossless) {
		t.Error("RLELossless should not be lossy compressed")
	}
	if !transfer.IsLossyCompressed(transfer.JPEGBaseline8Bit) {
		t.Error("JPEGBaseline8Bit should be lossy compressed")
	}
}
