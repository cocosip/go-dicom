// Package testutil provides helpers for tests to satisfy linters and improve safety.
// It includes safe numeric conversions and path handling utilities used across tests.
package testutil

import (
	"errors"
	"math"
	"os"
	"path/filepath"
)

// SafeUint16FromInt converts an int to uint16 with bounds checking.
// Params:
//   - v: integer value to convert.
//
// Returns:
//   - uint16: converted value, clamped to [0, math.MaxUint16].
//
// Description:
//
//	Ensures conversion does not overflow. Intended for test data generation.
func SafeUint16FromInt(v int) uint16 { // #nosec G115 -- conversion guarded by explicit bounds
	if v < 0 {
		return 0
	}
	if v > int(math.MaxUint16) {
		return math.MaxUint16
	}
	return uint16(v)
}

// SafeUint32FromInt converts an int to uint32 with bounds checking.
// Params:
//   - v: integer value to convert.
//
// Returns:
//   - uint32: converted value, clamped to [0, math.MaxUint32].
//
// Description:
//
//	Ensures conversion does not overflow. Intended for test data generation.
func SafeUint32FromInt(v int) uint32 { // #nosec G115 -- conversion guarded by explicit bounds
	if v < 0 {
		return 0
	}
	if v > int(math.MaxUint32) {
		return math.MaxUint32
	}
	return uint32(v)
}

// SafeJoin joins and cleans a path ensuring it remains under base.
// Params:
//   - base: base directory path.
//   - parts: path components to join under base.
//
// Returns:
//   - string: cleaned joined path.
//   - error: non-nil if path escapes base.
//
// Description:
//
//	Prevents path traversal; works cross-platform (Windows/Linux/macOS).
func SafeJoin(base string, parts ...string) (string, error) {
	p := filepath.Join(append([]string{base}, parts...)...)
	p = filepath.Clean(p)
	baseClean := filepath.Clean(base)
	// Allow base itself or child paths
	sep := string(os.PathSeparator)
	if p != baseClean && !hasPrefixPath(p, baseClean+sep) {
		return "", errors.New("path escapes base")
	}
	return p, nil
}

// OpenUnder opens a file under base safely.
// Params:
//   - base: base directory path.
//   - name: file name under base.
//
// Returns:
//   - *os.File: opened file handle.
//   - error: opening error or path violation.
//
// Description:
//
//	Uses SafeJoin to validate path; caller must close the file.
func OpenUnder(base, name string) (*os.File, error) {
    p, err := SafeJoin(base, name)
    if err != nil {
        return nil, err
    }
    // Path is validated and cleaned; open file.
    // #nosec G304 -- p comes from SafeJoin which prevents path traversal
    return os.Open(p)
}

// ReadFileUnder reads a file under base safely.
// Params:
//   - base: base directory path.
//   - name: file name under base.
//
// Returns:
//   - []byte: file contents.
//   - error: read error or path violation.
//
// Description:
//
//	Uses SafeJoin to validate path.
func ReadFileUnder(base, name string) ([]byte, error) {
    p, err := SafeJoin(base, name)
    if err != nil {
        return nil, err
    }
    // Path is validated and cleaned; read file.
    // #nosec G304 -- p comes from SafeJoin which prevents path traversal
    return os.ReadFile(p)
}

// hasPrefixPath checks whether path has the given prefix in a path-segment aware way.
// Params:
//   - path: full path to check.
//   - prefix: expected directory prefix ending with path separator.
//
// Returns:
//   - bool: true if path starts with prefix.
//
// Description:
//
//	Avoids false positives where simple string prefix could match unexpectedly.
func hasPrefixPath(path, prefix string) bool {
	// Normalize to platform-specific separators already handled by Clean.
	return len(path) >= len(prefix) && path[:len(prefix)] == prefix
}
