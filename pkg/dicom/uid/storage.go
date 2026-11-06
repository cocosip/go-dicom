// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package uid

import "strings"

// StorageSOPClassUIDs returns the UIDs of the active (non-retired) standard
// Storage SOP Classes in category, in catalog order.
//
// It is the Go equivalent of fo-dicom's
// DicomPresentationContext.GetScpRolePresentationContextsFromStorageUids and is
// typically used to propose the Storage presentation contexts a C-GET client
// can receive back from the SCP.
func StorageSOPClassUIDs(category StorageCategory) []string {
	if category == StorageCategoryNone {
		return nil
	}

	entries := StandardEntries()
	uids := make([]string, 0, len(entries))
	seen := make(map[string]struct{}, len(entries))
	for _, entry := range entries {
		if entry == nil || entry.IsRetired() || entry.StorageCategory() != category {
			continue
		}
		// StorageCategoryPrivate marks every private SOP Class as private, so
		// keep the storage-name filter to exclude private non-storage objects.
		if !strings.Contains(entry.Name(), "Storage") {
			continue
		}
		value := entry.UID()
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		uids = append(uids, value)
	}
	return uids
}
