// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transfer_test

import (
	"sort"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/io/endian"
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
	if ts != transfer.ImplicitVRLittleEndian {
		t.Error("Parse() did not preserve canonical standard object identity")
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

func TestStandardSyntaxParseIsImmutable(t *testing.T) {
	standard := transfer.ImplicitVRLittleEndian
	originalUID := standard.UID().UID()
	originalExplicit := standard.IsExplicitVR()

	if err := standard.Parse(transfer.ExplicitVRLittleEndian.UID().UID()); err == nil {
		t.Fatal("Parse() on a standard syntax should fail")
	}
	if standard.UID().UID() != originalUID || standard.IsExplicitVR() != originalExplicit {
		t.Fatal("Parse() changed the standard syntax")
	}

	custom := transfer.NewBuilder(uid.New("1.2.3.4.5.7", "Custom", uid.TypeTransferSyntax, false)).Build()
	if err := custom.Parse(transfer.ExplicitVRLittleEndian.UID().UID()); err != nil {
		t.Fatalf("Parse() on a non-standard syntax = %v", err)
	}
	if custom.UID() != transfer.ExplicitVRLittleEndian.UID() {
		t.Fatal("Parse() did not apply the canonical syntax to the custom receiver")
	}
}

func TestLookup(t *testing.T) {
	registry := transfer.NewRegistry()
	// Test lookup of a known transfer syntax
	u := uid.ImplicitVRLittleEndian
	ts, err := registry.Lookup(u)
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if !ts.Equals(transfer.ImplicitVRLittleEndian) {
		t.Error("Lookup() did not return ImplicitVRLittleEndian")
	}

	// Test lookup of nil UID
	_, err = registry.Lookup(nil)
	if err == nil {
		t.Error("Lookup(nil) should return error")
	}

	// Test lookup of non-transfer-syntax UID
	nonTS := uid.New("1.2.3.4.5", "Not a Transfer Syntax", uid.TypeSOPClass, false)
	_, err = registry.Lookup(nonTS)
	if err == nil {
		t.Error("Lookup() should return error for non-transfer-syntax UID")
	}
}

func TestNewRegistryIncludesStandardTransferSyntaxes(t *testing.T) {
	registry := transfer.NewRegistry()

	ts, err := registry.Lookup(uid.ImplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if !ts.Equals(transfer.ImplicitVRLittleEndian) {
		t.Fatal("isolated registry should return ImplicitVRLittleEndian for the standard UID")
	}
	if ts.IsExplicitVR() {
		t.Fatal("ImplicitVRLittleEndian from isolated registry should use implicit VR")
	}
	if ts.IsEncapsulated() {
		t.Fatal("ImplicitVRLittleEndian from isolated registry should not be encapsulated")
	}

	if len(registry.List()) == 0 {
		t.Fatal("isolated registry should include standard transfer syntaxes")
	}
}

func TestRegistryStandardIdentityOverlayMaskAndIsolation(t *testing.T) {
	first := transfer.NewRegistry()
	second := transfer.NewRegistry()
	standard := transfer.ImplicitVRLittleEndian

	if got := first.Query(standard.UID()); got != standard {
		t.Fatal("Query() did not return the canonical standard syntax")
	}
	if err := first.Register(standard); err == nil {
		t.Fatal("Register() accepted an effective standard duplicate")
	}

	replacement := transfer.NewBuilder(uid.ImplicitVRLittleEndian).
		SetExplicitVR(true).
		SetEncapsulated(true).
		Build()
	previous, err := first.Replace(replacement)
	if err != nil || previous != standard {
		t.Fatalf("Replace() = %v, %v", previous, err)
	}
	if first.Query(standard.UID()) != replacement {
		t.Fatal("replacement was not visible in the target registry")
	}
	parsed, err := transfer.Parse(standard.UID().UID())
	if err != nil {
		t.Fatalf("Parse(standard) error = %v", err)
	}
	if second.Query(standard.UID()) != standard || parsed != standard {
		t.Fatal("registry replacement changed another registry or the standard catalog")
	}

	removed, found := first.Unregister(standard.UID())
	if !found || removed != replacement {
		t.Fatalf("Unregister() = %v, %v", removed, found)
	}
	if first.Query(standard.UID()) != nil {
		t.Fatal("Unregister() did not mask the standard syntax")
	}
	unknown, err := first.Lookup(standard.UID())
	if err != nil {
		t.Fatalf("Lookup(masked standard) error = %v", err)
	}
	if unknown == standard || !unknown.IsExplicitVR() || !unknown.IsEncapsulated() {
		t.Fatalf("Lookup(masked standard) = %#v", unknown)
	}
}

func TestNewRegistryIncludesAllBuiltInTransferSyntaxes(t *testing.T) {
	registry := transfer.NewRegistry()

	tests := []struct {
		name string
		want *transfer.Syntax
	}{
		{name: "JPEG Lossless SV1", want: transfer.JPEGLosslessSV1},
		{name: "MPEG2", want: transfer.MPEG2},
		{name: "Fragmentable MPEG2", want: transfer.FragmentableMPEG2},
		{name: "HEVC Main Profile Level 5.1", want: transfer.HEVCH265MainProfileLevel51},
		{name: "HEVC Main 10 Profile Level 5.1", want: transfer.HEVCH265Main10ProfileLevel51},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := registry.Lookup(tt.want.UID())
			if err != nil {
				t.Fatalf("Lookup() error = %v", err)
			}
			if !got.Equals(tt.want) {
				t.Fatalf("Lookup() = %v, want %v", got, tt.want)
			}
			if got.IsEncapsulated() != tt.want.IsEncapsulated() {
				t.Fatalf("IsEncapsulated() = %v, want %v", got.IsEncapsulated(), tt.want.IsEncapsulated())
			}
			if got.IsLossy() != tt.want.IsLossy() {
				t.Fatalf("IsLossy() = %v, want %v", got.IsLossy(), tt.want.IsLossy())
			}
			if got.LossyCompressionMethod() != tt.want.LossyCompressionMethod() {
				t.Fatalf("LossyCompressionMethod() = %q, want %q",
					got.LossyCompressionMethod(), tt.want.LossyCompressionMethod())
			}
		})
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
	registry := transfer.NewRegistry()
	// Create a custom transfer syntax
	customUID := uid.New("1.2.3.4.5.6", "Custom Transfer Syntax", uid.TypeTransferSyntax, false)
	customTS := transfer.NewBuilder(customUID).
		SetExplicitVR(true).
		SetEndian(endian.Little).
		Build()

	if registry.Query(customUID) != nil {
		t.Fatal("Builder.Build() registered a transfer syntax")
	}
	if err := registry.Register(customTS); err != nil {
		t.Fatalf("Register() error = %v", err)
	}
	if err := registry.Register(customTS); err == nil {
		t.Fatal("Register() accepted a duplicate")
	}

	// Query should now find it
	queried := registry.Query(customUID)
	if queried == nil {
		t.Fatal("Query() returned nil for registered transfer syntax")
	}
	if !queried.Equals(customTS) {
		t.Error("Query() returned different transfer syntax")
	}

	// Unregister it
	removed, found := registry.Unregister(customUID)
	if !found || removed != customTS {
		t.Errorf("Unregister() = %v, %v", removed, found)
	}

	// Query should not find it anymore
	queried = registry.Query(customUID)
	if queried != nil {
		t.Error("Query() should return nil after unregister")
	}
}

func TestRegisteredSyntaxRejectsParseAndPreservesIdentity(t *testing.T) {
	const registeredUID = "1.2.3.4.5.62"
	registry := transfer.NewRegistry()
	syntax := transfer.NewBuilder(uid.New(registeredUID, "Registered", uid.TypeTransferSyntax, false)).
		SetExplicitVR(true).
		Build()
	if err := registry.Register(syntax); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if err := syntax.Parse(transfer.ImplicitVRLittleEndian.UID().UID()); err == nil {
		t.Fatal("Parse() on a registered syntax succeeded")
	}
	lookupUID := uid.New(registeredUID, "Lookup", uid.TypeTransferSyntax, false)
	if got := registry.Query(lookupUID); got != syntax {
		t.Fatalf("Query() = %v, want registered pointer", got)
	}
}

func TestRegisteredSyntaxDetachesUIDAliases(t *testing.T) {
	const (
		registeredUID = "1.2.3.4.5.63"
		changedUID    = "1.2.3.4.5.64"
	)
	registry := transfer.NewRegistry()
	sourceUID := uid.New(registeredUID, "Registered", uid.TypeTransferSyntax, false)
	syntax := transfer.NewBuilder(sourceUID).SetExplicitVR(true).Build()
	exposedUID := syntax.UID()
	if err := registry.Register(syntax); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if err := sourceUID.Parse(changedUID); err != nil {
		t.Fatalf("Parse() on caller-owned source UID error = %v", err)
	}
	if err := exposedUID.Parse(changedUID); err != nil {
		t.Fatalf("Parse() on pre-registration UID alias error = %v", err)
	}
	returnedUID := syntax.UID()
	if err := returnedUID.Parse(changedUID); err != nil {
		t.Fatalf("Parse() on returned UID snapshot error = %v", err)
	}

	lookupUID := uid.New(registeredUID, "Lookup", uid.TypeTransferSyntax, false)
	if got := registry.Query(lookupUID); got != syntax {
		t.Fatalf("Query(original) = %v, want registered pointer", got)
	}
	if syntax.UID().UID() != registeredUID {
		t.Fatalf("registered syntax UID changed to %q", syntax.UID().UID())
	}
	if got := registry.Query(uid.New(changedUID, "Changed", uid.TypeTransferSyntax, false)); got != nil {
		t.Fatalf("Query(changed) = %v, want nil", got)
	}
}

func TestRegisteredSyntaxBuilderUsesCopyOnWrite(t *testing.T) {
	const registeredUID = "1.2.3.4.5.65"
	registry := transfer.NewRegistry()
	builder := transfer.NewBuilder(uid.New(registeredUID, "Registered", uid.TypeTransferSyntax, false)).
		SetExplicitVR(false)
	registered := builder.Build()
	if err := registry.Register(registered); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	modified := builder.SetExplicitVR(true).Build()
	if modified == registered {
		t.Fatal("Builder modified the registered syntax in place")
	}
	if registered.IsExplicitVR() {
		t.Fatal("registered syntax changed through reused Builder")
	}
	if !modified.IsExplicitVR() {
		t.Fatal("Builder did not apply the requested change to its new syntax")
	}
	if got := registry.Query(uid.New(registeredUID, "Lookup", uid.TypeTransferSyntax, false)); got != registered {
		t.Fatalf("Query() = %v, want original registered pointer", got)
	}
}

func TestList(t *testing.T) {
	registry := transfer.NewRegistry()
	entries := registry.List()
	if len(entries) == 0 {
		t.Error("List() should return non-empty list")
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
		t.Error("List() should contain ImplicitVRLittleEndian")
	}
	values := make([]string, len(entries))
	for i, entry := range entries {
		values[i] = entry.UID().UID()
	}
	if !sort.StringsAreSorted(values) {
		t.Fatal("List() is not sorted by UID")
	}
	entries[0] = nil
	if registry.List()[0] == nil {
		t.Fatal("List() returned shared slice storage")
	}
}

func TestStandardTransferSyntaxListHasUniqueUIDs(t *testing.T) {
	entries := transfer.NewRegistry().List()
	seen := make(map[string]struct{}, len(entries))
	for _, entry := range entries {
		key := entry.UID().UID()
		if _, duplicate := seen[key]; duplicate {
			t.Fatalf("List() contains duplicate UID %q", key)
		}
		seen[key] = struct{}{}
	}
	if len(seen) != len(entries) {
		t.Fatalf("List() returned %d entries but %d unique UIDs", len(entries), len(seen))
	}
}

func TestRegistryConcurrentAccess(t *testing.T) {
	registry := transfer.NewRegistry()
	customUID := uid.New("1.2.3.4.9", "Concurrent Transfer Syntax", uid.TypeTransferSyntax, false)
	custom := transfer.NewBuilder(customUID).SetExplicitVR(true).SetEndian(endian.Little).Build()
	if err := registry.Register(custom); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	var wg sync.WaitGroup
	for range 16 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			registry.Query(customUID)
		}()
		go func() {
			defer wg.Done()
			registry.List()
		}()
		go func() {
			defer wg.Done()
			replacement := transfer.NewBuilder(customUID).SetExplicitVR(true).SetEndian(endian.Little).Build()
			_, _ = registry.Replace(replacement)
		}()
	}
	wg.Wait()

	if got := registry.Query(customUID); got == nil || got.UID().UID() != customUID.UID() {
		t.Fatalf("Query() after concurrent access = %v", got)
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
	if transfer.IsLosslessCompressed(transfer.JPEGBaseline8Bit) {
		t.Error("JPEGBaseline8Bit should not be lossless compressed")
	}
}

func TestIsLossyCompressed(t *testing.T) {
	if transfer.IsLossyCompressed(transfer.ImplicitVRLittleEndian) {
		t.Error("ImplicitVRLittleEndian should not be lossy compressed")
	}
	if !transfer.IsLossyCompressed(transfer.JPEGBaseline8Bit) {
		t.Error("JPEGBaseline8Bit should be lossy compressed")
	}
}
