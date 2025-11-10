// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
// can be converted to a single integer value in the OID arc 2.25.
//
// The UUID bytes are reordered to big-endian format and then interpreted
// as a 128-bit unsigned integer.
func convertUUIDToInteger(u uuid.UUID) string {
	// UUID bytes from uuid.UUID are in RFC 4122 format:
	// time_low (4 bytes) | time_mid (2 bytes) | time_hi_and_version (2 bytes) |
	// clock_seq_hi_and_reserved (1 byte) | clock_seq_low (1 byte) | node (6 bytes)
	//
	// To convert to a single integer, we need to reorder these bytes to
	// create a big-endian representation of the 128-bit value.
	//
	// The byte order transformation follows the .NET implementation:
	// Original indices: [0..15]
	// Reordered: [15,14,13,12,11,10,9,8,6,7,4,5,0,1,2,3]
	octets := u[:]

	// Create big-endian byte order for BigInt
	// BigInt expects bytes in big-endian order (most significant byte first)
	bigEndianOrder := []byte{
		octets[15], octets[14], octets[13], octets[12],
		octets[11], octets[10], octets[9], octets[8],
		octets[6], octets[7], octets[4], octets[5],
		octets[0], octets[1], octets[2], octets[3],
	}

	// Convert to big integer
	var intValue big.Int
	intValue.SetBytes(bigEndianOrder)

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
//	rootUID := "1.2.826.0.1.3680043.2.1343.1"
//	instanceUID := GenerateFromRoot(rootUID, 12345)
//	// Returns UID with value "1.2.826.0.1.3680043.2.1343.1.12345"
func GenerateFromRoot(root string, suffix int64) *UID {
	return Append(New(root, "Root UID", TypeUnknown, false), suffix)
}
