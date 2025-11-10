// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package uid implements DICOM Unique Identifier (UID) functionality.
package uid

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// Type represents the type of a DICOM UID.
type Type int

const (
	// TypeTransferSyntax represents a transfer syntax UID
	TypeTransferSyntax Type = iota
	// TypeSOPClass represents a SOP class UID
	TypeSOPClass
	// TypeMetaSOPClass represents a meta SOP class UID
	TypeMetaSOPClass
	// TypeServiceClass represents a service user UID
	TypeServiceClass
	// TypeSOPInstance represents a SOP instance UID
	TypeSOPInstance
	// TypeApplicationContextName represents an application entity UID
	TypeApplicationContextName
	// TypeApplicationHostingModel represents an application hosting model UID
	TypeApplicationHostingModel
	// TypeCodingScheme represents a coding scheme UID
	TypeCodingScheme
	// TypeFrameOfReference represents a code meaning UID
	TypeFrameOfReference
	// TypeLDAP represents an LDAP UID
	TypeLDAP
	// TypeMappingResource represents a mapping resource UID
	TypeMappingResource
	// TypeContextGroupName represents a context group name UID
	TypeContextGroupName
	// TypeUnknown represents an unknown UID type
	TypeUnknown
)

// StorageCategory represents the category of a storage SOP class.
type StorageCategory int

const (
	// StorageCategoryNone represents no specific storage category
	StorageCategoryNone StorageCategory = iota
	// StorageCategoryImage represents image storage
	StorageCategoryImage
	// StorageCategoryPresentationState represents a presentation state storage category
	StorageCategoryPresentationState
	// StorageCategoryStructuredReport represents a structured report storage category
	StorageCategoryStructuredReport
	// StorageCategoryWaveform represents a waveform storage category
	StorageCategoryWaveform
	// StorageCategoryDocument represents a document storage category
	StorageCategoryDocument
	// StorageCategoryRaw represents a raw data storage category
	StorageCategoryRaw
	// StorageCategoryOther represents other storage categories
	StorageCategoryOther
	// StorageCategoryPrivate represents private storage categories
	StorageCategoryPrivate
	// StorageCategoryVolume represents volume storage categories
	StorageCategoryVolume
)

// UID represents a DICOM Unique Identifier.
//
// A UID is a globally unique identifier that identifies various DICOM entities
// such as transfer syntaxes, SOP classes, and instances.
type UID struct {
	uid     string
	name    string
	uidType Type
	retired bool
}

var (
	// RootUID is the base UID for generating new UIDs
	RootUID = "1.2.826.0.1.3680043.2.1343.1"

	// UID registry
	uidRegistry = make(map[string]*UID)
	uidMutex    sync.RWMutex

	// UID validation regex: digits and dots, starting and ending with digit
	uidPattern = regexp.MustCompile(`^[0-9]+(\.[0-9]+)*$`)
)

// New creates a new UID.
func New(uid, name string, uidType Type, retired bool) *UID {
	return &UID{
		uid:     uid,
		name:    name,
		uidType: uidType,
		retired: retired,
	}
}

// UID returns the UID string.
func (u *UID) UID() string {
	return u.uid
}

// Name returns the descriptive name.
func (u *UID) Name() string {
	return u.name
}

// Type returns the UID type.
func (u *UID) Type() Type {
	return u.uidType
}

// IsRetired returns true if the UID is retired.
func (u *UID) IsRetired() bool {
	return u.retired
}

// String returns the UID string representation.
func (u *UID) String() string {
	return u.uid
}

// StorageCategory returns the storage category for SOP Class UIDs.
func (u *UID) StorageCategory() StorageCategory {
	// Private UID check
	if !strings.HasPrefix(u.uid, "1.2.840.10008") && u.uidType == TypeSOPClass {
		return StorageCategoryPrivate
	}

	// Only SOP Classes can have storage categories
	if u.uidType != TypeSOPClass {
		return StorageCategoryNone
	}

	// Check if it's a storage SOP class
	if !strings.Contains(u.name, "Storage") || strings.HasPrefix(u.name, "Storage Commitment") {
		return StorageCategoryNone
	}

	// Categorize by name
	if strings.Contains(u.name, "Image Storage") {
		return StorageCategoryImage
	}
	if strings.Contains(u.name, "Volume Storage") {
		return StorageCategoryVolume
	}
	if strings.Contains(u.name, "Presentation State") {
		return StorageCategoryPresentationState
	}
	if strings.Contains(u.name, "Structured Report") || strings.Contains(u.name, "SR Storage") {
		return StorageCategoryStructuredReport
	}
	if strings.Contains(u.name, "Waveform") {
		return StorageCategoryWaveform
	}
	if strings.Contains(u.name, "Document") || strings.Contains(u.name, "Encapsulated PDF") {
		return StorageCategoryDocument
	}
	if strings.Contains(u.name, "Raw Data") {
		return StorageCategoryRaw
	}

	return StorageCategoryOther
}

// IsImageStorage returns true if this is an image storage SOP class.
func (u *UID) IsImageStorage() bool {
	return u.StorageCategory() == StorageCategoryImage
}

// IsVolumeStorage returns true if this is a volume storage SOP class.
func (u *UID) IsVolumeStorage() bool {
	return u.StorageCategory() == StorageCategoryVolume
}

// Equals checks if two UIDs are equal.
func (u *UID) Equals(other *UID) bool {
	if other == nil {
		return false
	}
	return u.uid == other.uid
}

// Register registers a UID in the global registry.
func Register(u *UID) {
	uidMutex.Lock()
	defer uidMutex.Unlock()
	uidRegistry[u.uid] = u
}

// Parse parses a UID string and returns the corresponding UID.
//
// If the UID is registered, returns the registered instance.
// Otherwise, creates a new UID with the given name and type.
func Parse(s string, name string, uidType Type) *UID {
	// Trim trailing spaces and null characters
	s = strings.TrimRight(s, " \x00")

	// Look up in registry
	uidMutex.RLock()
	registered, found := uidRegistry[s]
	uidMutex.RUnlock()

	if found {
		return registered
	}

	// Create new UID
	return New(s, name, uidType, false)
}

// MustParse parses a UID string and panics if invalid.
func MustParse(s string) *UID {
	if !IsValid(s) {
		panic(fmt.Sprintf("invalid UID: %s", s))
	}
	return Parse(s, "Unknown", TypeUnknown)
}

// IsValid validates a UID string according to DICOM rules.
//
// A valid UID:
//   - Contains only digits (0-9) and dots (.)
//   - Each component is a number (no leading zeros except for "0")
//   - Maximum length is 64 characters
//   - Does not start or end with a dot
func IsValid(s string) bool {
	// Trim trailing spaces and nulls
	s = strings.TrimRight(s, " \x00")

	// Check length
	if len(s) == 0 || len(s) > 64 {
		return false
	}

	// Check pattern
	if !uidPattern.MatchString(s) {
		return false
	}

	// Check each component
	parts := strings.Split(s, ".")
	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
		// No leading zeros except for "0" itself
		if len(part) > 1 && part[0] == '0' {
			return false
		}
	}

	return true
}

// Append appends a sequence number to a base UID.
func Append(baseUID *UID, seq int64) *UID {
	newUID := fmt.Sprintf("%s.%d", baseUID.uid, seq)
	return New(newUID, "SOP Instance UID", TypeSOPInstance, false)
}

// Enumerate returns all registered UIDs.
func Enumerate() []*UID {
	uidMutex.RLock()
	defer uidMutex.RUnlock()

	result := make([]*UID, 0, len(uidRegistry))
	for _, u := range uidRegistry {
		result = append(result, u)
	}
	return result
}
