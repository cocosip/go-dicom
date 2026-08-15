// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestCodeItemUsesLongAndURNValueTags(t *testing.T) {
	longValue := "a-code-value-longer-than-16-characters"
	longCode := NewCodeItem(longValue, "99TEST", "Long code")
	if got := longCode.Value(); got != longValue {
		t.Fatalf("long CodeItem value = %q, want %q", got, longValue)
	}
	if !longCode.Dataset().Contains(tag.LongCodeValue) {
		t.Fatal("long CodeItem did not use LongCodeValue")
	}
	if longCode.Dataset().Contains(tag.CodeValue) {
		t.Fatal("long CodeItem also contained CodeValue")
	}

	urnValue := "urn:oid:1.2.840.10008.2.16.4"
	urnCode := NewCodeItem(urnValue, "", "URN code")
	if got := urnCode.Value(); got != urnValue {
		t.Fatalf("URN CodeItem value = %q, want %q", got, urnValue)
	}
	if !urnCode.Dataset().Contains(tag.URNCodeValue) {
		t.Fatal("URN CodeItem did not use URNCodeValue")
	}
}

func TestCodeItemEqualityIncludesCodingSchemeVersion(t *testing.T) {
	first := NewCodeItemWithVersion("12345", "99TEST", "Meaning A", "1")
	same := NewCodeItemWithVersion("12345", "99TEST", "Meaning B", "1")
	differentVersion := NewCodeItemWithVersion("12345", "99TEST", "Meaning A", "2")

	if !first.Equals(same) {
		t.Fatal("CodeItems with the same value, scheme, and version should be equal")
	}
	if first.Equals(differentVersion) {
		t.Fatal("CodeItems with different versions should not be equal")
	}
}

func TestScalarContentItemRoundTrips(t *testing.T) {
	code := NewCodeItem("121071", "DCM", "Finding")
	dateTime := time.Date(2026, time.August, 15, 13, 14, 15, 123456000, time.FixedZone("CST", 8*60*60))

	personName, err := NewContentItemPersonName(code, RelationshipContains, "Doe^Jane")
	if err != nil {
		t.Fatalf("NewContentItemPersonName() error = %v", err)
	}
	if got, err := personName.GetPersonName(); err != nil || got != "Doe^Jane" {
		t.Fatalf("GetPersonName() = %q, %v, want %q, nil", got, err, "Doe^Jane")
	}

	dateItem, err := NewContentItemDate(code, RelationshipContains, dateTime)
	if err != nil {
		t.Fatalf("NewContentItemDate() error = %v", err)
	}
	if got, err := dateItem.GetDate(); err != nil || got.Format("20060102") != "20260815" {
		t.Fatalf("GetDate() = %v, %v, want 20260815", got, err)
	}

	timeItem, err := NewContentItemTime(code, RelationshipContains, dateTime)
	if err != nil {
		t.Fatalf("NewContentItemTime() error = %v", err)
	}
	if got, err := timeItem.GetTime(); err != nil || got.Format("150405.000000") != "131415.123456" {
		t.Fatalf("GetTime() = %v, %v, want 131415.123456", got, err)
	}

	dateTimeItem, err := NewContentItemDateTime(code, RelationshipContains, dateTime)
	if err != nil {
		t.Fatalf("NewContentItemDateTime() error = %v", err)
	}
	if got, err := dateTimeItem.GetDateTime(); err != nil || !got.Equal(dateTime) {
		t.Fatalf("GetDateTime() = %v, %v, want %v", got, err, dateTime)
	}

	uidValue := "1.2.840.10008.5.1.4.1.1.88.11"
	uidItem, err := NewContentItemUIDReference(code, RelationshipContains, uidValue)
	if err != nil {
		t.Fatalf("NewContentItemUIDReference() error = %v", err)
	}
	if got, err := uidItem.GetUIDReference(); err != nil || got != uidValue {
		t.Fatalf("GetUIDReference() = %q, %v, want %q, nil", got, err, uidValue)
	}
}

func TestReferencedSOPContentItemRoundTrips(t *testing.T) {
	code := NewCodeItem("121112", "DCM", "Source of Measurement")
	reference := NewReferencedSOP("1.2.840.10008.1.2.3.4", "1.2.840.10008.5.1.4.1.1.2")

	tests := []struct {
		name      string
		valueType ValueType
		newItem   func(*CodeItem, Relationship, *ReferencedSOP) (*ContentItem, error)
	}{
		{name: "composite", valueType: ValueTypeComposite, newItem: NewContentItemComposite},
		{name: "image", valueType: ValueTypeImage, newItem: NewContentItemImage},
		{name: "waveform", valueType: ValueTypeWaveform, newItem: NewContentItemWaveform},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := tt.newItem(code, RelationshipInferredFrom, reference)
			if err != nil {
				t.Fatalf("constructor error = %v", err)
			}
			if got, err := item.ValueType(); err != nil || got != tt.valueType {
				t.Fatalf("ValueType() = %q, %v, want %q, nil", got, err, tt.valueType)
			}
			got, err := item.GetReferencedSOP()
			if err != nil {
				t.Fatalf("GetReferencedSOP() error = %v", err)
			}
			if got.InstanceUID() != reference.InstanceUID() || got.ClassUID() != reference.ClassUID() {
				t.Fatalf("GetReferencedSOP() = (%q, %q), want (%q, %q)",
					got.InstanceUID(), got.ClassUID(), reference.InstanceUID(), reference.ClassUID())
			}
		})
	}
}
