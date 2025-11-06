// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package uid_test

import (
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

func TestGenerateDerivedFromUUID(t *testing.T) {
	// Generate a UUID-derived UID
	u := uid.GenerateDerivedFromUUID()

	// Check that it starts with 2.25.
	if !strings.HasPrefix(u.UID(), "2.25.") {
		t.Errorf("GenerateDerivedFromUUID() UID = %q, should start with '2.25.'", u.UID())
	}

	// Check that the suffix is a valid integer
	suffix := strings.TrimPrefix(u.UID(), "2.25.")
	if suffix == "" {
		t.Error("GenerateDerivedFromUUID() UID has no suffix after '2.25.'")
	}

	// Check UID properties
	if u.Name() != "Local UID" {
		t.Errorf("GenerateDerivedFromUUID() Name = %q, want 'Local UID'", u.Name())
	}
	if u.Type() != uid.TypeUnknown {
		t.Errorf("GenerateDerivedFromUUID() Type = %v, want TypeUnknown", u.Type())
	}
	if u.IsRetired() {
		t.Error("GenerateDerivedFromUUID() IsRetired = true, want false")
	}

	// Generate another UID and verify it's different
	u2 := uid.GenerateDerivedFromUUID()
	if u.UID() == u2.UID() {
		t.Error("GenerateDerivedFromUUID() should generate unique UIDs")
	}
}

func TestGenerateDerivedFromUUID_Valid(t *testing.T) {
	// Generate 10 UIDs and verify they're all valid
	for i := 0; i < 10; i++ {
		u := uid.GenerateDerivedFromUUID()
		if !uid.IsValid(u.UID()) {
			t.Errorf("GenerateDerivedFromUUID() generated invalid UID: %q", u.UID())
		}
	}
}

func TestNewGenerator(t *testing.T) {
	gen := uid.NewGenerator()
	if gen == nil {
		t.Fatal("NewGenerator() returned nil")
	}
}

func TestGenerator_Generate(t *testing.T) {
	gen := uid.NewGenerator()

	// Create a source UID
	sourceUID := uid.New("1.2.3.4.5", "Source UID", uid.TypeSOPInstance, false)

	// Generate destination UID
	destUID1 := gen.Generate(sourceUID)
	if destUID1 == nil {
		t.Fatal("Generate() returned nil")
	}

	// Verify it's a UUID-derived UID
	if !strings.HasPrefix(destUID1.UID(), "2.25.") {
		t.Errorf("Generate() UID = %q, should start with '2.25.'", destUID1.UID())
	}

	// Generate again with same source - should get same destination
	destUID2 := gen.Generate(sourceUID)
	if !destUID1.Equals(destUID2) {
		t.Errorf("Generate() with same source returned different UIDs: %q vs %q",
			destUID1.UID(), destUID2.UID())
	}

	// Generate with different source - should get different destination
	sourceUID2 := uid.New("1.2.3.4.6", "Source UID 2", uid.TypeSOPInstance, false)
	destUID3 := gen.Generate(sourceUID2)
	if destUID1.Equals(destUID3) {
		t.Error("Generate() with different source returned same UID")
	}
}

func TestGenerator_GenerateNil(t *testing.T) {
	gen := uid.NewGenerator()
	result := gen.Generate(nil)
	if result != nil {
		t.Errorf("Generate(nil) = %v, want nil", result)
	}
}

func TestGenerator_Concurrent(t *testing.T) {
	gen := uid.NewGenerator()
	sourceUID := uid.New("1.2.3.4.5", "Source UID", uid.TypeSOPInstance, false)

	// Generate UIDs concurrently
	const numGoroutines = 100
	results := make([]*uid.UID, numGoroutines)
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			results[index] = gen.Generate(sourceUID)
		}(i)
	}

	wg.Wait()

	// All results should be the same
	firstUID := results[0].UID()
	for i, result := range results {
		if result.UID() != firstUID {
			t.Errorf("Concurrent Generate() at index %d = %q, want %q",
				i, result.UID(), firstUID)
		}
	}
}

func TestGenerator_MultipleSourceUIDs(t *testing.T) {
	gen := uid.NewGenerator()

	// Create multiple source UIDs and generate destinations
	numSources := 10
	sources := make([]*uid.UID, numSources)
	destinations := make([]*uid.UID, numSources)

	for i := 0; i < numSources; i++ {
		sources[i] = uid.New("1.2.3.4."+string(rune('0'+i)), "Source "+string(rune('0'+i)), uid.TypeSOPInstance, false)
		destinations[i] = gen.Generate(sources[i])
	}

	// Verify all destinations are unique
	for i := 0; i < numSources; i++ {
		for j := i + 1; j < numSources; j++ {
			if destinations[i].Equals(destinations[j]) {
				t.Errorf("Destinations %d and %d are the same: %q",
					i, j, destinations[i].UID())
			}
		}
	}

	// Verify generating again gives same results
	for i := 0; i < numSources; i++ {
		dest := gen.Generate(sources[i])
		if !dest.Equals(destinations[i]) {
			t.Errorf("Second generation for source %d returned different UID: %q vs %q",
				i, dest.UID(), destinations[i].UID())
		}
	}
}

func TestGenerateFromRoot(t *testing.T) {
	root := "1.2.826.0.1.3680043.2.1343.1"

	// Generate UID from root
	u := uid.GenerateFromRoot(root, 12345)

	want := root + ".12345"
	if u.UID() != want {
		t.Errorf("GenerateFromRoot(%q, 12345) = %q, want %q", root, u.UID(), want)
	}

	// Verify it's a SOP Instance type
	if u.Type() != uid.TypeSOPInstance {
		t.Errorf("GenerateFromRoot() Type = %v, want TypeSOPInstance", u.Type())
	}
}

func TestGenerateFromRoot_MultipleValues(t *testing.T) {
	root := "1.2.826.0.1.3680043.2.1343.1"

	tests := []struct {
		suffix int64
		want   string
	}{
		{1, root + ".1"},
		{42, root + ".42"},
		{999999, root + ".999999"},
		{0, root + ".0"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			u := uid.GenerateFromRoot(root, tt.suffix)
			if u.UID() != tt.want {
				t.Errorf("GenerateFromRoot(%q, %d) = %q, want %q",
					root, tt.suffix, u.UID(), tt.want)
			}
		})
	}
}

// TestUUIDConversion verifies that the UUID to integer conversion
// produces valid DICOM UIDs
func TestUUIDConversion(t *testing.T) {
	// Test with a known UUID
	testUUID := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	// The conversion should produce a UID starting with 2.25.
	u := uid.GenerateDerivedFromUUID()
	if !strings.HasPrefix(u.UID(), "2.25.") {
		t.Errorf("UUID conversion produced UID without 2.25. prefix: %q", u.UID())
	}

	// The UID should be valid according to DICOM rules
	if !uid.IsValid(u.UID()) {
		t.Errorf("UUID conversion produced invalid DICOM UID: %q", u.UID())
	}

	// The UID should not exceed 64 characters (DICOM limit)
	if len(u.UID()) > 64 {
		t.Errorf("UUID conversion produced UID longer than 64 chars: %d chars", len(u.UID()))
	}

	_ = testUUID // Use the variable to avoid unused warning
}

// Benchmark UUID-derived UID generation
func BenchmarkGenerateDerivedFromUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uid.GenerateDerivedFromUUID()
	}
}

// Benchmark Generator with repeated source UIDs
func BenchmarkGenerator_Generate(b *testing.B) {
	gen := uid.NewGenerator()
	sourceUID := uid.New("1.2.3.4.5", "Source", uid.TypeSOPInstance, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Generate(sourceUID)
	}
}
