// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestPersonNameComponentGroups(t *testing.T) {
	t.Run("GetComponents uses alphabetic group", func(t *testing.T) {
		elem := element.NewPersonName(tag.PatientName, []string{"Yamada^Tarou=Yamada^Taro"})
		components, err := elem.GetComponents(0)
		if err != nil {
			t.Fatalf("GetComponents() error = %v", err)
		}
		if components.FamilyName != "Yamada" || components.GivenName != "Tarou" {
			t.Fatalf("GetComponents() = %+v, want Family=Yamada Given=Tarou", components)
		}
	})

	t.Run("phonetic-only value falls back to first non-empty group", func(t *testing.T) {
		elem := element.NewPersonName(tag.PatientName, []string{"=Yamada^Tarou"})
		components, err := elem.GetComponents(0)
		if err != nil {
			t.Fatalf("GetComponents() error = %v", err)
		}
		if components.FamilyName != "Yamada" || components.GivenName != "Tarou" {
			t.Fatalf("GetComponents() = %+v, want Family=Yamada Given=Tarou", components)
		}
	})

	t.Run("GetFormattedName ignores ideographic group", func(t *testing.T) {
		elem := element.NewPersonName(tag.PatientName, []string{"Yamada^Tarou=Yamada^Taro"})
		name, err := elem.GetFormattedName(0)
		if err != nil {
			t.Fatalf("GetFormattedName() error = %v", err)
		}
		if name != "Tarou Yamada" {
			t.Fatalf("GetFormattedName() = %q, want %q", name, "Tarou Yamada")
		}
	})

	t.Run("GetAllComponents handles groups per value", func(t *testing.T) {
		elem := element.NewPersonName(tag.PatientName, []string{"A^B=C^D", "E^F"})
		all, err := elem.GetAllComponents()
		if err != nil {
			t.Fatalf("GetAllComponents() error = %v", err)
		}
		if len(all) != 2 {
			t.Fatalf("GetAllComponents() returned %d entries, want 2", len(all))
		}
		if all[0].FamilyName != "A" || all[0].GivenName != "B" {
			t.Fatalf("GetAllComponents()[0] = %+v, want Family=A Given=B", all[0])
		}
		if all[1].FamilyName != "E" || all[1].GivenName != "F" {
			t.Fatalf("GetAllComponents()[1] = %+v, want Family=E Given=F", all[1])
		}
	})
}
