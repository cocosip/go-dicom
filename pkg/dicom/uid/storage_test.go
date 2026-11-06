// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package uid

import "testing"

func TestStorageSOPClassUIDs(t *testing.T) {
	images := StorageSOPClassUIDs(StorageCategoryImage)
	if len(images) == 0 {
		t.Fatal("StorageSOPClassUIDs(StorageCategoryImage) returned no UIDs")
	}

	seen := make(map[string]struct{}, len(images))
	for _, value := range images {
		if _, duplicate := seen[value]; duplicate {
			t.Fatalf("duplicate storage UID %s", value)
		}
		seen[value] = struct{}{}

		parsed := Parse(value, "", TypeSOPClass)
		if parsed.IsRetired() {
			t.Fatalf("retired UID %s returned", value)
		}
		if parsed.StorageCategory() != StorageCategoryImage {
			t.Fatalf("UID %s has category %v, want image", value, parsed.StorageCategory())
		}
	}

	if _, ok := seen[CTImageStorage.UID()]; !ok {
		t.Fatal("CT Image Storage missing from image storage UIDs")
	}
	if _, ok := seen[MRImageStorage.UID()]; !ok {
		t.Fatal("MR Image Storage missing from image storage UIDs")
	}
}

func TestStorageSOPClassUIDsNone(t *testing.T) {
	if got := StorageSOPClassUIDs(StorageCategoryNone); got != nil {
		t.Fatalf("StorageSOPClassUIDs(StorageCategoryNone) = %v, want nil", got)
	}
}
