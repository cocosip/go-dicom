// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package vm implements DICOM Value Multiplicity (VM) types.
//
// VM defines the number of values that can be encoded in a DICOM element.
// Examples:
//   - VM "1": Exactly 1 value
//   - VM "1-n": 1 or more values
//   - VM "2-2n": 2, 4, 6, 8, ... values (multiples of 2)
package vm

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

// VM represents a DICOM Value Multiplicity specification.
type VM struct {
	// minimum is the minimum number of values allowed
	minimum int

	// maximum is the maximum number of values allowed (math.MaxInt for unlimited)
	maximum int

	// multiplicity is the step size (e.g., 2 for "2-2n" means values must be multiples of 2)
	multiplicity int
}

// Minimum returns the minimum number of values allowed.
func (v *VM) Minimum() int {
	return v.minimum
}

// Maximum returns the maximum number of values allowed.
// Returns math.MaxInt for unlimited.
func (v *VM) Maximum() int {
	return v.maximum
}

// Multiplicity returns the value multiplicity step size.
func (v *VM) Multiplicity() int {
	return v.multiplicity
}

// IsUnlimited returns true if there is no maximum limit on the number of values.
func (v *VM) IsUnlimited() bool {
	return v.maximum == math.MaxInt
}

// IsValid checks if a given count satisfies this VM.
func (v *VM) IsValid(count int) bool {
	if count < v.minimum || count > v.maximum {
		return false
	}

	if v.multiplicity > 1 {
		// Check if count is a valid multiple
		return (count-v.minimum)%v.multiplicity == 0
	}

	return true
}

// String returns the string representation of the VM.
func (v *VM) String() string {
	if v.minimum == v.maximum {
		return strconv.Itoa(v.minimum)
	}

	if v.maximum == math.MaxInt {
		if v.multiplicity > 1 {
			return fmt.Sprintf("%d-%dn", v.minimum, v.multiplicity)
		}
		return fmt.Sprintf("%d-n", v.minimum)
	}

	return fmt.Sprintf("%d-%d", v.minimum, v.maximum)
}

// vmCache caches parsed VM instances to avoid repeated parsing
var (
	vmCache   = make(map[string]*VM)
	vmCacheMu sync.RWMutex
)

// Parse parses a VM string specification.
//
// Supported formats:
//   - "1": Exactly 1 value
//   - "1-2": 1 to 2 values
//   - "1-n": 1 or more values (unlimited)
//   - "2-2n": 2, 4, 6, 8, ... values (multiples of 2)
//   - "3-3n": 3, 6, 9, 12, ... values (multiples of 3)
func Parse(s string) (*VM, error) {
	// Check cache first
	vmCacheMu.RLock()
	if cached, ok := vmCache[s]; ok {
		vmCacheMu.RUnlock()
		return cached, nil
	}
	vmCacheMu.RUnlock()

	// Parse the VM string
	vm, err := parseVM(s)
	if err != nil {
		return nil, err
	}

	// Cache the result
	vmCacheMu.Lock()
	vmCache[s] = vm
	vmCacheMu.Unlock()

	return vm, nil
}

// MustParse parses a VM string and panics if there is an error.
// Use this only for static/constant VM values.
func MustParse(s string) *VM {
	vm, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("invalid VM string '%s': %v", s, err))
	}
	return vm
}

// parseVM performs the actual parsing logic.
func parseVM(s string) (*VM, error) {
	// Split by "-", " ", or "or"
	// Examples: "1-2", "1 or 2", "1-n", "2-2n"
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == ' '
	})

	// Remove "or" if present
	filtered := make([]string, 0, len(parts))
	for _, part := range parts {
		if strings.ToLower(part) != "or" {
			filtered = append(filtered, part)
		}
	}
	parts = filtered

	if len(parts) == 0 {
		return nil, fmt.Errorf("empty VM string")
	}

	if len(parts) == 1 {
		// Single value: "1", "2", etc.
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid VM value '%s': %w", parts[0], err)
		}

		return &VM{
			minimum:      val,
			maximum:      val,
			multiplicity: val,
		}, nil
	}

	if len(parts) == 2 {
		// Range: "1-2", "1-n", "2-2n", etc.
		min, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid minimum value '%s': %w", parts[0], err)
		}

		vm := &VM{
			minimum:      min,
			multiplicity: 1,
		}

		maxStr := parts[1]

		// Check for "n" suffix (unlimited)
		if strings.HasSuffix(strings.ToLower(maxStr), "n") {
			vm.maximum = math.MaxInt

			// Remove the "n" suffix
			maxStr = strings.TrimSuffix(strings.ToLower(maxStr), "n")

			// If there's a number before "n", it's the multiplicity
			if maxStr != "" {
				mult, err := strconv.Atoi(maxStr)
				if err != nil {
					return nil, fmt.Errorf("invalid multiplicity value '%s': %w", maxStr, err)
				}
				vm.multiplicity = mult
			}
		} else {
			// Fixed maximum
			max, err := strconv.Atoi(maxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid maximum value '%s': %w", maxStr, err)
			}
			vm.maximum = max
		}

		return vm, nil
	}

	return nil, fmt.Errorf("invalid VM format: '%s'", s)
}

// Standard VM constants
var (
	// VM1 represents exactly 1 value
	VM1 = MustParse("1")

	// VM1_2 represents 1 to 2 values
	VM1_2 = MustParse("1-2")

	// VM1_3 represents 1 to 3 values
	VM1_3 = MustParse("1-3")

	// VM1_8 represents 1 to 8 values
	VM1_8 = MustParse("1-8")

	// VM1_32 represents 1 to 32 values
	VM1_32 = MustParse("1-32")

	// VM1_99 represents 1 to 99 values
	VM1_99 = MustParse("1-99")

	// VM1_n represents 1 or more values (unlimited)
	VM1_n = MustParse("1-n")

	// VM2 represents exactly 2 values
	VM2 = MustParse("2")

	// VM2_n represents 2 or more values (unlimited)
	VM2_n = MustParse("2-n")

	// VM2_2n represents 2, 4, 6, 8, ... values (multiples of 2)
	VM2_2n = MustParse("2-2n")

	// VM3 represents exactly 3 values
	VM3 = MustParse("3")

	// VM3_n represents 3 or more values (unlimited)
	VM3_n = MustParse("3-n")

	// VM3_3n represents 3, 6, 9, 12, ... values (multiples of 3)
	VM3_3n = MustParse("3-3n")

	// VM4 represents exactly 4 values
	VM4 = MustParse("4")

	// VM6 represents exactly 6 values
	VM6 = MustParse("6")

	// VM16 represents exactly 16 values
	VM16 = MustParse("16")
)
