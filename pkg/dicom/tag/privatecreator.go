// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package tag

import (
    "hash/fnv"
    "strings"
)

// PrivateCreator represents a DICOM private creator identification.
//
// Private creators are used to identify the organization or vendor that
// defines private DICOM tags. The creator string is stored in specific
// tags (Group,0010-00FF) and referenced by private data elements.
type PrivateCreator struct {
	creator string
}

// NewPrivateCreator creates a new private creator with the specified creator string.
func NewPrivateCreator(creator string) *PrivateCreator {
	return &PrivateCreator{
		creator: creator,
	}
}

// Creator returns the creator identification string.
func (pc *PrivateCreator) Creator() string {
	return pc.creator
}

// String returns the string representation of the private creator.
func (pc *PrivateCreator) String() string {
	return pc.creator
}

// Compare compares this private creator to another.
// Returns:
//   - negative if this < other
//   - zero if this == other
//   - positive if this > other
func (pc *PrivateCreator) Compare(other *PrivateCreator) int {
	return strings.Compare(pc.creator, other.creator)
}

// Equals checks if this private creator equals another.
func (pc *PrivateCreator) Equals(other *PrivateCreator) bool {
	if other == nil {
		return false
	}
	return pc.creator == other.creator
}

// Hash returns a hash code for the private creator.
func (pc *PrivateCreator) Hash() uint32 {
    h := fnv.New32a()
    // Write returns an error we should check to satisfy linters and robustness.
    // In practice, fnv.Hash's Write never returns an error, but we still handle it.
    if _, err := h.Write([]byte(pc.creator)); err != nil {
        // Fallback: compute FNV-1a manually if an unexpected error occurs.
        return manualFNV32a(pc.creator)
    }
    return h.Sum32()
}

// manualFNV32a computes the FNV-1a 32-bit hash for a string.
// Params:
//   - s: input string to hash.
// Returns:
//   - uint32: the computed FNV-1a 32-bit hash value.
// Description:
//   Implements FNV-1a (32-bit) in pure Go for the unlikely case when fnv.New32a().Write fails.
func manualFNV32a(s string) uint32 {
    const offset32 uint32 = 2166136261
    const prime32 uint32 = 16777619
    hash := offset32
    for i := 0; i < len(s); i++ {
        hash ^= uint32(s[i])
        hash *= prime32
    }
    return hash
}
