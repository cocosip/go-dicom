// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package uid

import (
	"math/big"
	"sync"

	"github.com/google/uuid"
)

// Generator is used for generating DICOM UIDs.
//
// It maintains a mapping from source UIDs to destination UIDs,
// ensuring that the same source UID always maps to the same destination UID.
type Generator struct {
	mu           sync.RWMutex
	sourceUIDMap map[string]*UID
}

// NewGenerator creates a new UID generator.
func NewGenerator() *Generator {
	return &Generator{
		sourceUIDMap: make(map[string]*UID),
	}
}

// Generate returns a destination UID for the given source UID.
//
// If the source UID has been seen before, returns the previously generated
// destination UID. Otherwise, generates a new UUID-derived UID and stores
// the mapping for future calls.
func (g *Generator) Generate(sourceUID *UID) *UID {
	if sourceUID == nil {
		return nil
	}

	sourceUIDStr := sourceUID.UID()

	// Check if we already have a mapping
	g.mu.RLock()
	destUID, exists := g.sourceUIDMap[sourceUIDStr]
	g.mu.RUnlock()

	if exists {
		return destUID
	}

	// Generate new UID and store mapping
	g.mu.Lock()
	defer g.mu.Unlock()

	// Double-check in case another goroutine created it
	if destUID, exists = g.sourceUIDMap[sourceUIDStr]; exists {
		return destUID
	}

	destUID = GenerateDerivedFromUUID()
	g.sourceUIDMap[sourceUIDStr] = destUID
	return destUID
}

// GenerateDerivedFromUUID generates a UUID-derived UID according to
// DICOM PS3.5 Section B.2.
//
// The generated UID has the format "2.25.{integer}" where {integer} is
// the decimal representation of the UUID as a 128-bit integer.
//
// This method is compliant with ISO/IEC 9834-8, paragraph 6.3.
func GenerateDerivedFromUUID() *UID {
	u := uuid.New()
	uidString := convertUUIDToInteger(u)
	return New(uidString, "Local UID", TypeUnknown, false)
}

// convertUUIDToInteger converts a UUID to its integer OID format.
//
// According to ISO/IEC 9834-8 (referenced by DICOM PS3.5 B.2), a UUID
// can be converted to a single integer value in the OID arc 2.25 by
// interpreting the 128-bit UUID value as an unsigned integer.
//
// A google/uuid.UUID already stores the octets in RFC 4122 (network/big-endian)
// order, so the octets are used as-is; interpreting them as a big-endian
// integer yields the UUID's numeric value.
func convertUUIDToInteger(u uuid.UUID) string {
	var intValue big.Int
	intValue.SetBytes(u[:])

	// Return as 2.25.{integer} format
	return "2.25." + intValue.String()
}

// GenerateFromRoot generates a new UID by appending a suffix to the root UID.
//
// This is useful for generating instance UIDs or other derived UIDs
// based on an organization's root UID.
//
// Example:
//
//	instanceUID := GenerateFromRoot(RootUID, 12345)
//	// Returns a UID under the go-dicom placeholder root.
//
// RootUID is a placeholder pending registration; production code should pass
// its own registered organisation root instead.
func GenerateFromRoot(root string, suffix int64) *UID {
	return Append(New(root, "Root UID", TypeUnknown, false), suffix)
}
